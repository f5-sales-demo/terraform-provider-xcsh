---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3021210032222001-3201003210113000-3310200202320013-3233011322102311-1213220210003332-1232213333010022-3011313231331002-1232320210131030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.mandatory_claims` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
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
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
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

<a id="canonical-2032023032321212-2012031323211202-3011133003112013-0313113211020011-1111300032020012-0133300211212001-2022120023223222-2333310322312110"></a>

### Direct properties for `jwt_validation.reserved_claims`

- [audience](resources--http_loadbalancer--reference--group-020.md#canonical-2300303331000122-3121101203310331-1101010002013011-3212310311202230-1120312322133332-0103002110110220-2011100122032132-2222021332021310): complete subsection reference.

- [audience_disable](resources--http_loadbalancer--reference--group-020.md#canonical-2223212220013230-3111321103313220-0222122120332033-0010012000300010-2232103201330012-0212200021021033-2213220202003330-0010230213010222): complete subsection reference.

<a id="canonical-0133300020002310-0112322033132310-3323120231311313-1303020100031303-0012203112320310-2002332313201131-1300100100012120-1210101123203310"></a>

<a id="canonical-1023321010221003-1013100221113321-0112321023211223-0203222221333021-0310321322102203-3223213022231122-0113301230110133-0011021133131313"></a>

#### `jwt_validation.reserved_claims.issuer` property

Type: `"string"`. Optional.

Exact Match. Exclusive with \[issuer\_disable\]

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

- [issuer_disable](resources--http_loadbalancer--reference--group-020.md#canonical-0020321310110001-0231212322321023-2210012110110103-2220201210230313-1001330321023313-0123021031323331-0131330022220103-0012131111000002): complete subsection reference.

- [validate_period_disable](resources--http_loadbalancer--reference--group-020.md#canonical-1212310211203322-2132113320213210-0032103232222000-0003112132223212-3111011312130223-1122212110003000-1023011020311220-0032333301011220): complete subsection reference.

- [validate_period_enable](resources--http_loadbalancer--reference--group-020.md#canonical-1023300203001213-2322122200323300-1330123303122313-0100003200130110-3300321100201032-2123103033121103-0322122213033102-1212003013030233): complete subsection reference.

<a id="canonical-2300303331000122-3121101203310331-1101010002013011-3212310311202230-1120312322133332-0103002110110220-2011100122032132-2222021332021310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.audience` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
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

<a id="canonical-0302302321200010-0120213130100200-2210030010022201-3300133122200003-1132322021203312-1322303332330131-2320221331312033-1110321011213223"></a>

### Direct properties for `jwt_validation.reserved_claims.audience`

<a id="canonical-2022031120030233-1010313220013012-3110013101021231-1031333003211333-0313232223202033-1121230132012131-0212030011112032-0110130111002001"></a>

#### `jwt_validation.reserved_claims.audience.audiences` property

Type: `["list", "string"]`. Optional.

Values. Configuration parameter for audiences

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

<a id="canonical-2223212220013230-3111321103313220-0222122120332033-0010012000300010-2232103201330012-0212200021021033-2213220202003330-0010230213010222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.audience_disable` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222)
- jwt_validation.reserved_claims.audience_disable

<a id="canonical-3212210311213201-0122130111203312-2202302021032112-2313022230320110-3333330131123111-0202323111200310-2300113213301330-2132332120023001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for audience disable.

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
audience_disable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0020321310110001-0231212322321023-2210012110110103-2220201210230313-1001330321023313-0123021031323331-0131330022220103-0012131111000002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.issuer_disable` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222)
- jwt_validation.reserved_claims.issuer_disable

<a id="canonical-1210322203203110-2330001213300031-0201132302122232-2221322202211233-0300022030230131-2023230003312132-0012323003020012-0121103212330223"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for issuer disable.

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
issuer_disable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212310211203322-2132113320213210-0032103232222000-0003112132223212-3111011312130223-1122212110003000-1023011020311220-0032333301011220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.validate_period_disable` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222)
- jwt_validation.reserved_claims.validate_period_disable

<a id="canonical-0113232213113110-1003232330313312-3103100310220103-3011233131010033-2303233101020100-1321120032013313-3013000202111033-0232020322333012"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for validate period disable.

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
validate_period_disable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023300203001213-2322122200323300-1330123303122313-0100003200130110-3300321100201032-2123103033121103-0322122213033102-1212003013030233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.validate_period_enable` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222)
- jwt_validation.reserved_claims.validate_period_enable

<a id="canonical-0021020103322113-0021232223323323-3232131213211023-0131031303313321-1302021230200121-3313300320333222-0200111202330223-2332322000232310"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for validate period enable.

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
validate_period_enable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203010010312213-3220111213032103-0132012333011222-1022230022121330-0102222111013230-1112030311103110-2131210033203332-1122013310120221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.target` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
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

<a id="canonical-0000211333001111-3103220211322313-0330101103330311-1223100312111113-2202330023011202-0031223311300022-2030121022213230-0100121101132332"></a>

### Direct properties for `jwt_validation.target`

- [all_endpoint](resources--http_loadbalancer--reference--group-020.md#canonical-1020320031023321-0320202313120203-1020333130031033-0231120033031022-3001322011013132-1031131331321201-2333121233322230-3033322013201211): complete subsection reference.

- [api_groups](resources--http_loadbalancer--reference--group-020.md#canonical-1222112033113103-0212223103001302-3322022133003223-2323202221310130-2133022333111311-2320312303300010-2023130011233101-3133212112101102): complete subsection reference.

- [base_paths](resources--http_loadbalancer--reference--group-020.md#canonical-1211200313211220-2030022210023120-0330203001332121-0220311123001323-2112313323311233-2013210013013230-3202001132002122-1233232212322110): complete subsection reference.

<a id="canonical-1020320031023321-0320202313120203-1020333130031033-0231120033031022-3001322011013132-1031131331321201-2333121233322230-3033322013201211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.target.all_endpoint` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.target](resources--http_loadbalancer--reference--group-020.md#canonical-3203010010312213-3220111213032103-0132012333011222-1022230022121330-0102222111013230-1112030311103110-2131210033203332-1122013310120221)
- jwt_validation.target.all_endpoint

<a id="canonical-3021120023122321-0213210011132330-0022200203222301-3333110121020311-1010131201232030-3111302321013302-3312003223023131-0320132011202331"></a>

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
all_endpoint = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222112033113103-0212223103001302-3322022133003223-2323202221310130-2133022333111311-2320312303300010-2023130011233101-3133212112101102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.target.api_groups` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
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

<a id="canonical-1033200220313201-2320213113121333-1022022001333130-0002100103132112-0033003312021321-2233201010033120-1203013122221123-3323211331020132"></a>

### Direct properties for `jwt_validation.target.api_groups`

<a id="canonical-0201030213123102-0213312032123202-2100121100332101-2001020212003133-3130113313132332-3003132332202302-1310321010320112-0021212133001101"></a>

#### `jwt_validation.target.api_groups.api_groups` property

Type: `["list", "string"]`. Optional.

API Groups. Group or collection configuration

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

<a id="canonical-1211200313211220-2030022210023120-0330203001332121-0220311123001323-2112313323311233-2013210013013230-3202001132002122-1233232212322110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.target.base_paths` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
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

<a id="canonical-1221122302313122-1222301310133002-2303322213012322-1301032001330000-1300212033123223-3130332123210000-2122210001120101-2002000212231210"></a>

### Direct properties for `jwt_validation.target.base_paths`

<a id="canonical-3223211203303010-0010123313313033-3333221001101111-0001122103123310-2233102023131131-0211333311010133-2230301102023021-1022311310302301"></a>

#### `jwt_validation.target.base_paths.base_paths` property

Type: `["list", "string"]`. Optional.

Prefix Values. File system or URL path

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

<a id="canonical-2100130323133112-2010031310200203-1331011110302222-2122110331203111-1230200312133021-0102032323223320-3132123010020310-1130031012332120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.token_location` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- jwt_validation.token_location

<a id="canonical-0110122110101020-2213022213101220-3001021132321201-1101100011110303-3110023220212110-3302233031320030-0133200002223203-2231132023300230"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for token location.

Additional upstream details:

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

<a id="canonical-3202220000300021-0311100310210101-3310320131301122-0120103200310301-0203202330311122-3020003301021232-3003110031020031-2222221110233110"></a>

### Direct properties for `jwt_validation.token_location`

- [bearer_token](resources--http_loadbalancer--reference--group-020.md#canonical-3321120100031033-3323212200212200-0123023031330003-1203231013131122-2011301000100231-3131211333010001-3010002000331213-3110110300331302): complete subsection reference.

<a id="canonical-3321120100031033-3323212200212200-0123023031330003-1203231013131122-2011301000100231-3131211333010001-3010002000331213-3110110300331302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.token_location.bearer_token` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.token_location](resources--http_loadbalancer--reference--group-020.md#canonical-2100130323133112-2010031310200203-1331011110302222-2122110331203111-1230200312133021-0102032323223320-3132123010020310-1130031012332120)
- jwt_validation.token_location.bearer_token

<a id="canonical-0211221203032321-1212111002310033-1001220110022211-1001330230131100-2020103120323113-1020311300200203-3213000230222222-1030331313231012"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bearer token.

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
bearer_token {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320110132020001-0022333031222332-0332001223201230-0000302202010121-2213203130001312-3221132020131213-3301310301230123-2220221132033300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_action_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- l7_ddos_action_block

<a id="canonical-0033232131220130-3322300300000223-1323101032222302-2111321001222301-0312212303003130-3002120021323011-2320000031121121-3130133311130202"></a>

Type: `["object", {}]`. Optional.

\[OneOf: l7\_ddos\_action\_block, l7\_ddos\_action\_default, l7\_ddos\_action\_js\_challenge;
Default: l7\_ddos\_action\_default\] Enable this option

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

- [l7_ddos_action_block](resources--http_loadbalancer--reference--group-020.md#canonical-0033232131220130-3322300300000223-1323101032222302-2111321001222301-0312212303003130-3002120021323011-2320000031121121-3130133311130202)
- [l7_ddos_action_default](resources--http_loadbalancer--reference--group-020.md#canonical-0310222221220320-0030213003223012-0003112031003100-3000300312023111-1011003103012311-3231220231213221-2220232130310122-0120103032031120)
- [l7_ddos_action_js_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-3201213003032333-3301120031030220-0131331332333232-1133310130122013-1002021113313300-0022133200331212-2232013002123020-1201313313200322)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
l7_ddos_action_block = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122303101123332-1003322121300112-3112102023031013-1001133230202001-1121101002133313-0121023032001313-1200303023002210-0310210303322333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_action_default` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- l7_ddos_action_default

<a id="canonical-0310222221220320-0030213003223012-0003112031003100-3000300312023111-1011003103012311-3231220231213221-2220232130310122-0120103032031120"></a>

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
l7_ddos_action_default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130022200313200-0012221203233323-2031132210123022-0122213131223223-0033223300233010-3030023013112322-3321030210121223-0130102321222031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_action_js_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- l7_ddos_action_js_challenge

<a id="canonical-3201213003032333-3301120031030220-0131331332333232-1133310130122013-1002021113313300-0022133200331212-2232013002123020-1201313313200322"></a>

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

<a id="canonical-2313122011032321-1332310221022022-0330302212112302-0201332330113031-2232133313030223-1030113003020012-1220332202011202-1123032122012303"></a>

### Direct properties for `l7_ddos_action_js_challenge`

<a id="canonical-3131023030010021-2210303302333303-0031031011112301-0003211203002000-3201031203101110-2021130310210002-2110222203101022-0232212001213110"></a>

#### `l7_ddos_action_js_challenge.cookie_expiry` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0221310110312322-0131302002123110-0321131022310312-0312331210232021-3213022310000123-1222102033031100-0232133222022310-0102231212303220"></a>

#### `l7_ddos_action_js_challenge.custom_page` property

Type: `"string"`. Optional.

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

<a id="canonical-3101120320330203-1202322132032123-0123131303211230-0020112113031132-2130203233001220-2021101010210200-3030100331020211-3011013213012123"></a>

#### `l7_ddos_action_js_challenge.js_script_delay` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- l7_ddos_protection

<a id="canonical-3110302123103331-1102003311011121-0222232232200030-1232201323102322-2000300322231132-2303121211321100-2001322232123021-2312323033012310"></a>

Type: `"object"`. single nested block, Optional.

L7 DDoS protection is critical for safeguarding web applications, APIs, and services that are
exposed to the internet from sophisticated, volumetric, application-level threats. Configure
actions, thresholds and policies to apply during L7 DDoS attack. Defaults to \`map\[\]\`. Server
applies default when omitted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("clientside_action_captcha_challenge",
    "clientside_action_js_challenge"),
  validators.ConflictingObjectAttributes("clientside_action_captcha_challenge",
    "clientside_action_none"),
  validators.ConflictingObjectAttributes("clientside_action_js_challenge",
    "clientside_action_none"),
  validators.ConflictingObjectAttributes("ddos_policy_custom",
    "ddos_policy_none"),
  validators.ConflictingObjectAttributes("default_rps_threshold",
    "rps_threshold"),
  validators.ConflictingObjectAttributes("mitigation_block",
    "mitigation_captcha_challenge"),
  validators.ConflictingObjectAttributes("mitigation_block",
    "mitigation_js_challenge"),
  validators.ConflictingObjectAttributes("mitigation_captcha_challenge",
    "mitigation_js_challenge")}
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
  "x-ves-oneof-field-clientside_action_choice": "[\"clientside_action_captcha_challenge\",\"clientside_action_js_challenge\",\"clientside_action_none\"]",
  "x-ves-oneof-field-ddos_policy_choice": "[\"ddos_policy_custom\",\"ddos_policy_none\"]",
  "x-ves-oneof-field-mitigation_action_choice": "[\"mitigation_block\",\"mitigation_captcha_challenge\",\"mitigation_js_challenge\"]",
  "x-ves-oneof-field-rps_threshold_choice": "[\"default_rps_threshold\",\"rps_threshold\"]"
}
```

Terraform syntax:

```terraform
l7_ddos_protection {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212223013233032-0020100231300122-3220213331313303-0001011202311001-2012202023310202-2330333220211310-0332120013222300-0212330203332123"></a>

### Direct properties for `l7_ddos_protection`

- [clientside_action_captcha_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-2132322130132131-0221002210111202-0312210122200110-2021110221003130-1200330031302010-3221030211111131-1333301121201332-1122121202211132): complete subsection reference.

- [clientside_action_js_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-0333120302000032-2111010020220213-3001222100021202-3303111122030223-3301012233231223-2023303300030103-1333103133330100-2013231002002123): complete subsection reference.

- [clientside_action_none](resources--http_loadbalancer--reference--group-020.md#canonical-0232020023213211-0010012133112133-3201232132101111-0122310232212020-2332312100133201-3002012000121223-2333311311203210-2200010031330232): complete subsection reference.

- [ddos_policy_custom](resources--http_loadbalancer--reference--group-020.md#canonical-2032302331132210-0023203122022322-3032330330303001-2020223202033300-2030111322330001-0312200322113211-3000033301113203-3303002121321102): complete subsection reference.

- [ddos_policy_none](resources--http_loadbalancer--reference--group-020.md#canonical-2300010132330322-2103003000322002-3021312112020131-2303000032213033-3013331212212202-0213203100113003-2201112331020132-0300213200302321): complete subsection reference.

- [default_rps_threshold](resources--http_loadbalancer--reference--group-020.md#canonical-0301032010330233-2032013210320130-3210101323222232-0012021320302330-0223331231012320-2131320332211100-2001103200321213-1330302002102102): complete subsection reference.

- [mitigation_block](resources--http_loadbalancer--reference--group-020.md#canonical-3102200321321113-1320220022300332-1021202111011303-2133113221331112-1002233311030020-3000303031300302-1122210131112013-1111010212132100): complete subsection reference.

- [mitigation_captcha_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-2311312103021000-2231011313233313-0000032301302200-1110230312213123-3030113001001113-1103000230122301-3102220030133023-1103012032130232): complete subsection reference.

- [mitigation_js_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-1100330131220011-2223323303122231-3000301210002122-0002323302333200-2222221113020033-0312010003231303-0010020200123030-3130033132221212): complete subsection reference.

<a id="canonical-2002330212323213-0100112213231011-3123131011303313-2310031230020023-0011113331032322-0112221031301200-1003130210230312-0302032110031033"></a>

<a id="canonical-2212021031012313-3313213223320303-1201011311002132-3003211222022201-1203311333321333-0102022300121330-2133322110223211-3001120331223000"></a>

#### `l7_ddos_protection.rps_threshold` property

Type: `"number"`. Optional.

Exclusive with \[default\_rps\_threshold\] Configure custom RPS threshold.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 50000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 50000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "50000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "50000"
  }
}
```

<a id="canonical-2132322130132131-0221002210111202-0312210122200110-2021110221003130-1200330031302010-3221030211111131-1333301121201332-1122121202211132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.clientside_action_captcha_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.clientside_action_captcha_challenge

<a id="canonical-2023210103313231-3210213220322303-1011012013100101-0000123023000001-2233003113301111-2133300030202301-2010133312123110-2002021000231331"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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
clientside_action_captcha_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033003230033313-2002101233302200-2031302221021030-3311013123213200-3320100000323002-2211113001311120-2222111100123222-3111200322230102"></a>

### Direct properties for `l7_ddos_protection.clientside_action_captcha_challenge`

<a id="canonical-0332300012103103-0322121323333323-0232032122003001-2202113120313301-3303322110312003-3221330300031330-0311222013332230-0232120023310032"></a>

#### `l7_ddos_protection.clientside_action_captcha_challenge.cookie_expiry` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3301012132023002-2031313133022300-0011032233023022-2022332310020333-0333002322032032-0132321321101313-3222200000200023-1333023101100012"></a>

<a id="canonical-3221010131033100-2113313121030210-1131111121113310-1020233102311213-3110111213001231-0132100211331121-1221322312202322-3300311303332102"></a>

#### `l7_ddos_protection.clientside_action_captcha_challenge.custom_page` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0333120302000032-2111010020220213-3001222100021202-3303111122030223-3301012233231223-2023303300030103-1333103133330100-2013231002002123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.clientside_action_js_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.clientside_action_js_challenge

<a id="canonical-1002123331011332-0120233123332130-1132301100121330-0322323221331331-2320302021300213-0311123102121223-1020232112112213-1311132103003323"></a>

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
clientside_action_js_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020210230032130-0031012331311311-1330022031221031-2012103332110332-1310001212113330-1123300320100332-0330012020220311-2323110300033321"></a>

### Direct properties for `l7_ddos_protection.clientside_action_js_challenge`

<a id="canonical-1130330300313000-2222122322011021-0300022102312223-3323100302030303-2213012220002120-3120302120212213-1030302101312220-2103313112120001"></a>

#### `l7_ddos_protection.clientside_action_js_challenge.cookie_expiry` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0221220121011300-1002313101300121-0211102222321213-0201232030321000-1001311012001030-1213203321203222-3033230033003012-2311331122331311"></a>

<a id="canonical-3222233230033332-1102221231013311-3120131130333202-0212310133031131-0023222002113110-1023313010223220-2031002103313200-0320231331331221"></a>

#### `l7_ddos_protection.clientside_action_js_challenge.custom_page` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3120133331312301-3010330013330133-2123130112023331-0220031111231102-0323032003300130-3022320121200321-2200100303010023-2010121123231212"></a>

<a id="canonical-3300300123000303-2111111333300232-2103131320110221-0110003002133022-3112212231321320-2122330332212122-0203022331230123-2322221323100300"></a>

#### `l7_ddos_protection.clientside_action_js_challenge.js_script_delay` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0232020023213211-0010012133112133-3201232132101111-0122310232212020-2332312100133201-3002012000121223-2333311311203210-2200010031330232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.clientside_action_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.clientside_action_none

<a id="canonical-2232303310220223-2121002300221020-3223110321223121-2020311130310110-0313211200232231-3232133320113212-3312231203201003-2131321132222333"></a>

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
clientside_action_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032302331132210-0023203122022322-3032330330303001-2020223202033300-2030111322330001-0312200322113211-3000033301113203-3303002121321102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.ddos_policy_custom` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.ddos_policy_custom

<a id="canonical-2023112333321212-2312212002203210-3330322233210023-0111300222230213-3130333002130102-2232322023303133-2313112231313113-0310221300132313"></a>

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
ddos_policy_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-2003103211101021-1033030222130023-1332120222012020-0111033233101102-0203302110103121-1132123122301023-0123331103222201-2203310201002311"></a>

### Direct properties for `l7_ddos_protection.ddos_policy_custom`

<a id="canonical-0330031133112302-3000220200103000-3110231312300030-1020012000213102-0101013233333113-1301322210230233-2032033001132021-0211230121313310"></a>

#### `l7_ddos_protection.ddos_policy_custom.name` property

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

<a id="canonical-2103002231123130-1213230101000033-0002130112221212-2110001232211100-3333332111002003-2202010110331332-3203233112233113-0112312021230030"></a>

<a id="canonical-1200220033330223-0203200220211032-2101221312222331-3000222231112133-3001110330210320-3112133003201030-1102120103320211-3211110233033100"></a>

#### `l7_ddos_protection.ddos_policy_custom.namespace` property

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

<a id="canonical-2313231032032102-3233331232313203-3222203321231312-2022203011101212-1323332001322310-1130010312302311-2231001122220312-0210123001113010"></a>

<a id="canonical-0031001100213120-1302000133313013-0220010111002022-3201033021000221-1231310322131213-3100223333300323-1111332233110131-1203232210200122"></a>

#### `l7_ddos_protection.ddos_policy_custom.tenant` property

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

<a id="canonical-2300010132330322-2103003000322002-3021312112020131-2303000032213033-3013331212212202-0213203100113003-2201112331020132-0300213200302321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.ddos_policy_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.ddos_policy_none

<a id="canonical-1030211220130021-0203311002323303-3232020002112332-0101123321211012-0001200001303032-3033202122103211-2231313012322020-1012222232213021"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ddos policy none.

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
ddos_policy_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301032010330233-2032013210320130-3210101323222232-0012021320302330-0223331231012320-2131320332211100-2001103200321213-1330302002102102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.default_rps_threshold` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.default_rps_threshold

<a id="canonical-1222233123130302-3312313101101003-3032322213223300-3313002231231311-0011303300213303-3120011303321131-3231322021330020-1323122023001311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default rps threshold.

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
default_rps_threshold = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102200321321113-1320220022300332-1021202111011303-2133113221331112-1002233311030020-3000303031300302-1122210131112013-1111010212132100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.mitigation_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.mitigation_block

<a id="canonical-1302232310102120-3321021001223231-2300000031130123-3202303131232133-2010332011100130-0213311302321221-0323230323313212-2103123231230133"></a>

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
mitigation_block = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311312103021000-2231011313233313-0000032301302200-1110230312213123-3030113001001113-1103000230122301-3102220030133023-1103012032130232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.mitigation_captcha_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.mitigation_captcha_challenge

<a id="canonical-2103221032121123-0022001013003020-2022320011012222-0100103012123123-1022202201202322-2120223133101130-1122000312303023-0111010011020312"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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
mitigation_captcha_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103003102021100-1100303230102101-0200221230122111-1031210331013210-0031001331312301-3122303020011002-1113133121332221-1122020311311210"></a>

### Direct properties for `l7_ddos_protection.mitigation_captcha_challenge`

<a id="canonical-3203200311032323-2001300020100202-2311102030030231-1323223202333033-0110113301032002-2221201000022311-2202202223102211-3020312333003213"></a>

#### `l7_ddos_protection.mitigation_captcha_challenge.cookie_expiry` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0032302013030130-1000033113220121-0202312013223001-3020312230131213-0321312312121230-3122131323032300-2320002330112121-1103223230121333"></a>

<a id="canonical-1312322133323101-0332213031303210-0011100321312133-1232130011220221-3032232001203020-0303033300300232-3133230002113323-1103330023202323"></a>

#### `l7_ddos_protection.mitigation_captcha_challenge.custom_page` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1100330131220011-2223323303122231-3000301210002122-0002323302333200-2222221113020033-0312010003231303-0010020200123030-3130033132221212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.mitigation_js_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.mitigation_js_challenge

<a id="canonical-1100012301120112-2202032112113213-3303220332023333-1011032322131031-0333322112122232-0001212113322233-3223212123022221-1010203133103031"></a>

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
mitigation_js_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230320300313223-1002120200320213-2103020330323311-3230223213321211-3301031030032302-1212203300132032-3003122002121220-1000232310131202"></a>

### Direct properties for `l7_ddos_protection.mitigation_js_challenge`

<a id="canonical-3120202110231203-0312120011330330-3323303123222302-2022101112002113-3122302020210000-0201220332103330-3031031220302121-3100303012230130"></a>

#### `l7_ddos_protection.mitigation_js_challenge.cookie_expiry` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3130011210203130-1231200123133311-3112101002120320-1211011323312020-2111123010332100-3120232102323200-0013112132110313-0312330030003320"></a>

<a id="canonical-0203330321122021-2001022332020233-3030333102213233-1100013121232331-1210133200230112-2010102230233122-2002332321333232-0100001300301211"></a>

#### `l7_ddos_protection.mitigation_js_challenge.custom_page` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1233300201023201-3220122320031202-3010123233301300-3022210321203212-1233222013012202-3020233333001230-0203333131310113-0312120132020003"></a>

<a id="canonical-1033122232223112-0200122011313111-3103200201111113-3223112301231021-3103003310131212-3211031111120111-2121010011000331-3321200203310302"></a>

#### `l7_ddos_protection.mitigation_js_challenge.js_script_delay` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3100013201201320-1030023222021000-1020231012221032-2100130020230332-3130310121332223-2110221222102103-3201213033132002-0230230101332320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `least_active` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- least_active

<a id="canonical-0012031220101203-1331010113102002-1112333232131023-3022102022010211-3120131001301000-1311230213201110-2003321233300023-2003020103022211"></a>

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
least_active = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- malware_protection_settings

<a id="canonical-2131132110220032-0231001231223002-1030113021221223-1220220323132330-0320302100310030-1022231121013021-1003021302003300-2223310300300221"></a>

Type: `"object"`. single nested block, Optional.

Malware Protection protects Web Apps and APIs, from malicious file uploads by scanning files in
real-time.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("malware_protection_rules")}
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
malware_protection_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303222130003012-0020023231330333-2022230212123302-2010130300202011-2322003310302230-3030220031111220-0110123200330112-2120122300033203"></a>

### Direct properties for `malware_protection_settings`

- [malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002): complete subsection reference.

<a id="canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- malware_protection_settings.malware_protection_rules

<a id="canonical-2003313323213130-3200012223333213-2123023201133202-1021313232303312-1323230003030111-3211121322032311-2213132022312010-1003333123212030"></a>

Type: `"object"`. list nested block, Optional.

Configure the match criteria to trigger Malware Protection Scan.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
malware_protection_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320122323321302-0320020131300120-3212331330001010-2121012030303131-3303201220123033-3001201331223300-0203132131202120-3103331333313101"></a>

### Direct properties for `malware_protection_settings.malware_protection_rules`

- [action](resources--http_loadbalancer--reference--group-020.md#canonical-1103131210031313-3311231021030110-0032103021313030-3111311021013102-2030230131230231-0312110100211132-0221333000302223-0213201313023331): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-020.md#canonical-1111301311211300-2021032002103123-2030010313113220-0322312002103113-1323320333032031-3311300130232232-2120101201300121-2120202200222202): complete subsection reference.

<a id="canonical-1123211222000221-3330110210331021-3100302111013330-1320012313300220-2001132002000300-3021000111301002-1223013203011020-2312001330202201"></a>

<a id="canonical-3232100202010220-0123323122033302-2330333321201030-2112212201301313-1120222321000302-2020012020011011-0003311032133001-2213003331203322"></a>

#### `malware_protection_settings.malware_protection_rules.http_methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] HTTP Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-020.md#canonical-0223322011233333-2023010130231330-0010321002000323-0122011111220022-3003332303022132-2110332221322031-3030112211110311-0332132113211233): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-020.md#canonical-3322003033222212-3013222122332232-2330201330212203-1131100100112220-3111212202333120-2013103011123302-3212033033303033-1113102122300202): complete subsection reference.

<a id="canonical-1103131210031313-3311231021030110-0032103021313030-3111311021013102-2030230131230231-0312110100211132-0221333000302223-0213201313023331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules.action` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- malware_protection_settings.malware_protection_rules.action

<a id="canonical-2210212321221320-0001032121031230-1220020331000013-2320312003013120-0221222100102310-3331001130213223-2322210130011310-1302131003222021"></a>

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

<a id="canonical-2120312102110020-2112303203133210-3202032321312202-2131113002300231-3233231330111223-2123123130101222-3101330331031031-2233211330003103"></a>

### Direct properties for `malware_protection_settings.malware_protection_rules.action`

- [block](resources--http_loadbalancer--reference--group-020.md#canonical-0112031212302003-3133023110301111-0110110200302313-0313313333322201-1122032312202132-1222020022232311-1220002101220110-0021030120230233): complete subsection reference.

- [report](resources--http_loadbalancer--reference--group-020.md#canonical-3112330000013332-1132020133131113-1130320100133102-1320011030300130-1201330121232231-1331233000130002-2301130201122113-0023013303320102): complete subsection reference.

<a id="canonical-0112031212302003-3133023110301111-0110110200302313-0313313333322201-1122032312202132-1222020022232311-1220002101220110-0021030120230233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules.action.block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- [malware_protection_settings.malware_protection_rules.action](resources--http_loadbalancer--reference--group-020.md#canonical-1103131210031313-3311231021030110-0032103021313030-3111311021013102-2030230131230231-0312110100211132-0221333000302223-0213201313023331)
- malware_protection_settings.malware_protection_rules.action.block

<a id="canonical-2201020301210300-1131232200010033-0221011120232302-0002230200030322-2001031000023301-0133311301213023-1320222301120313-0333201123331111"></a>

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

<a id="canonical-3112330000013332-1132020133131113-1130320100133102-1320011030300130-1201330121232231-1331233000130002-2301130201122113-0023013303320102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules.action.report` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- [malware_protection_settings.malware_protection_rules.action](resources--http_loadbalancer--reference--group-020.md#canonical-1103131210031313-3311231021030110-0032103021313030-3111311021013102-2030230131230231-0312110100211132-0221333000302223-0213201313023331)
- malware_protection_settings.malware_protection_rules.action.report

<a id="canonical-2212332212203203-1212320102213321-1311013232130133-0110022201032212-1013203321232313-2320220023301231-2232032022002133-3111133300103200"></a>

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

<a id="canonical-1111301311211300-2021032002103123-2030010313113220-0322312002103113-1323320333032031-3311300130232232-2120101201300121-2120202200222202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- malware_protection_settings.malware_protection_rules.domain

<a id="canonical-1003330200023001-2210003333123122-0122013213303230-0112311201111033-1303210113110223-1211222211210032-0000332302311112-0301313310321100"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Additional upstream details:

Domain to be matched.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_domain",
    "domain")}
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
  "x-ves-oneof-field-domain_matcher": "[\"any_domain\",\"domain\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320230111031033-0021130023313321-2231132321121202-0220212001113202-3103011102032113-2010231300310330-2303101311003131-2130221310120022"></a>

### Direct properties for `malware_protection_settings.malware_protection_rules.domain`

- [any_domain](resources--http_loadbalancer--reference--group-020.md#canonical-2131102033323300-2113201202300021-3330301302112221-1222123122210123-2210311031002332-0212011032023222-1332220103131100-3120311313223133): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-020.md#canonical-3101221001133011-1222331203111231-3210213103012220-0303200021313013-1100223121022033-2131002210230220-2011121030102021-3301233222323300): complete subsection reference.

<a id="canonical-2131102033323300-2113201202300021-3330301302112221-1222123122210123-2210311031002332-0212011032023222-1332220103131100-3120311313223133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules.domain.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- [malware_protection_settings.malware_protection_rules.domain](resources--http_loadbalancer--reference--group-020.md#canonical-1111301311211300-2021032002103123-2030010313113220-0322312002103113-1323320333032031-3311300130232232-2120101201300121-2120202200222202)
- malware_protection_settings.malware_protection_rules.domain.any_domain

<a id="canonical-0322031120230010-0132231233010131-1130223030320130-0012013111312031-2102300232213321-1223332233210022-3022333032000113-0213202321122233"></a>

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

<a id="canonical-3101221001133011-1222331203111231-3210213103012220-0303200021313013-1100223121022033-2131002210230220-2011121030102021-3301233222323300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules.domain.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- [malware_protection_settings.malware_protection_rules.domain](resources--http_loadbalancer--reference--group-020.md#canonical-1111301311211300-2021032002103123-2030010313113220-0322312002103113-1323320333032031-3311300130232232-2120101201300121-2120202200222202)
- malware_protection_settings.malware_protection_rules.domain.domain

<a id="canonical-3013210230031230-0133322212032112-1131033220101310-1331302201011331-3201033311123113-1021310321302300-0223121103321312-0010011123020213"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-3000002313312123-1202320023020000-0103001103320320-3002230111331002-1121211312131330-2020223313033110-1110200211200331-2020113201222121"></a>

### Direct properties for `malware_protection_settings.malware_protection_rules.domain.domain`

<a id="canonical-0211210032303032-3211321012301132-0032322101132302-2320133313332202-3032212000021000-0301113201220332-2103311121003120-1221321023000100"></a>

#### `malware_protection_settings.malware_protection_rules.domain.domain.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0022000221230100-0102203232003301-1303022013320203-2331022002321003-1133110232202033-1323300311131122-2203210301320111-1203321200033120"></a>

<a id="canonical-2333121013000221-1020001210111023-1202011122211212-2111131223210310-2121003121121030-3100032122232100-2302122231023302-3132302103000200"></a>

#### `malware_protection_settings.malware_protection_rules.domain.domain.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0010010030103301-2023113003223320-2301322311213132-3202000113112200-0033310003213221-3330200010123222-3123323002130032-2323001330233021"></a>

<a id="canonical-0302012021101102-0131332000013101-0023112012131030-1330112230110030-0311300301301102-0301200212130012-3101322013230223-0200003033310233"></a>

#### `malware_protection_settings.malware_protection_rules.domain.domain.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0223322011233333-2023010130231330-0010321002000323-0122011111220022-3003332303022132-2110332221322031-3030112211110311-0332132113211233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- malware_protection_settings.malware_protection_rules.metadata

<a id="canonical-3011002222322020-1030203011213323-0012132331130010-3300233120200330-1133300022131013-1012000101213112-1123303230102322-2311000301203013"></a>

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

<a id="canonical-1123001330313320-2132202000120231-1103130003310213-1130303311032203-3110012102330023-1131012321330202-2120223130110103-3030330231212203"></a>

### Direct properties for `malware_protection_settings.malware_protection_rules.metadata`

<a id="canonical-3120131232311310-3031022212001111-0301030021022322-2231111321233331-0302131330211212-2233103023202232-2023012122102030-2213301200310230"></a>

#### `malware_protection_settings.malware_protection_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1112012021101222-0213121100311203-2320323103333001-0303333110200111-0112232232101200-2110120122031331-1032012221033332-0123302000021301"></a>

<a id="canonical-2023300012012323-3012103223223223-0312200033131221-1232033103112223-2302321200333010-3203113310321001-1123202003232303-2312120113002200"></a>

#### `malware_protection_settings.malware_protection_rules.metadata.name` property

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

<a id="canonical-3322003033222212-3013222122332232-2330201330212203-1131100100112220-3111212202333120-2013103011123302-3212033033303033-1113102122300202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-020.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- malware_protection_settings.malware_protection_rules.path

<a id="canonical-3202311131310211-3011313221120113-2012111202020300-3220012011020032-3301301020212321-0333231121311020-1102232223332222-1131221100230002"></a>

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

<a id="canonical-0123222031031323-0001232120022310-0003220002232311-3002030300112320-1213130330133220-3323201201313001-1323100212113010-0231130323131132"></a>

### Direct properties for `malware_protection_settings.malware_protection_rules.path`

<a id="canonical-0133012320321310-0000331202130002-3332210113201130-3003232221023320-1330322333321321-0021031203113320-1333212002102112-0223210112310000"></a>

#### `malware_protection_settings.malware_protection_rules.path.path` property

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

<a id="canonical-0022220321020031-3331133032103013-0103200003322011-2323001103313301-1120020211230213-0101100321132020-1112101110113332-1303030303203311"></a>

<a id="canonical-0310020012313132-0102133333020013-1222230120000013-2300310031101301-3203001013202301-2113010321110113-2020202231230312-3130230112030120"></a>

#### `malware_protection_settings.malware_protection_rules.path.prefix` property

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

<a id="canonical-2223200020130020-3022200320323301-3012131323102100-1303210220121021-2320233332303330-3120222133032102-1010222103330223-2233221202322132"></a>

<a id="canonical-0202333333033322-2233112221211111-2111011221112123-0230000031002311-2203231232101110-0202103100010222-3202320230303323-2023010231300020"></a>

#### `malware_protection_settings.malware_protection_rules.path.regex` property

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

<a id="canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- more_option

<a id="canonical-0301121120121203-3222003213101121-0113110333321300-0301211131322210-1010300021322222-1320321112322013-1333031323000321-2130301001121303"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to define a route.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("max_requests_per_connection",
    "no_request_limit_per_connection")}
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
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-strict_sni_host_header_check_choice": "[]"
}
```

Terraform syntax:

```terraform
more_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-0202220202023023-1123321120321133-2321233102021112-0212022121012210-3002103202322130-3331231010302201-2032220001010132-2210203003200202"></a>

### Direct properties for `more_option`

- [buffer_policy](resources--http_loadbalancer--reference--group-021.md#canonical-3001230121131210-2320111332013012-2111012221211123-0201000121203103-2231133211130003-1333223331020330-3002111133222000-2333201100021120): complete subsection reference.

- [compression_params](resources--http_loadbalancer--reference--group-021.md#canonical-2210002000001212-0010203202310332-1010121210231301-3211000110202332-0022230020101320-1312133232033210-1220213012122321-0223023202033001): complete subsection reference.

<a id="canonical-1232310212203031-1011003112231311-0301012021213023-0202203302011012-3110031020213333-0221322012311302-0331132310033000-0133130012030103"></a>

<a id="canonical-2323102323023010-1322000032200312-1030333302003230-0231333311000000-2300000130001220-1012302131213000-0201201312110112-3201321101311223"></a>

#### `more_option.custom_errors` property

Type: `["map", "string"]`. Optional.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response
code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP
responses. Specific response code takes preference when both response code and response code class
matches for a request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"ranges\":[[3,3],[4,4],[5,5],[300,599]],\"type\":\"uint32-string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.uint32.ranges\":\"3,4,5,300-599\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"65536\",\"ves.io.schema.rules.map.values.string.uri_ref\":\"true\"},\"values\":{\"format\":\"uri-reference\",\"maxLength\":65536,\"type\":\"string\"}}")}
```

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
    "keys": {
      "ranges": [
        [
          3,
          3
        ],
        [
          4,
          4
        ],
        [
          5,
          5
        ],
        [
          300,
          599
        ]
      ],
      "type": "uint32-string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "65536",
      "ves.io.schema.rules.map.values.string.uri_ref": "true"
    },
    "values": {
      "format": "uri-reference",
      "maxLength": 65536,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  }
}
```

<a id="canonical-0003010020000331-3022131020303020-0030122112122032-1231222310330031-2331302102321230-3303231000100222-0322021313023123-0203330321022032"></a>

<a id="canonical-2010213111320221-0230000102013223-0102112331232220-3320122300210003-0101111022023033-0223201222231020-2201211300132132-1003111030320213"></a>

#### `more_option.disable_default_error_pages` property

Type: `"bool"`. Optional.

Disable the use of default F5XC error pages.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [disable_path_normalize](resources--http_loadbalancer--reference--group-021.md#canonical-1033132001213130-2001323231130303-3302133101202300-2012320223111211-2113022313122123-3231031321132200-0022333302032210-1001120012321201): complete subsection reference.

- [enable_path_normalize](resources--http_loadbalancer--reference--group-021.md#canonical-0200102323330033-2123023013233122-3220321300123212-3020110001131003-2313312012331231-1030020232012220-3302210111313313-3333332202102203): complete subsection reference.

<a id="canonical-3133133130103310-2221113213321333-2100330301230320-3121330330021302-1012302031310233-1212213230000320-2313023232031322-1200013331201100"></a>

<a id="canonical-2021200011132202-0133302000310103-1201132311200223-0123323223131033-3212113030102122-2031012232321112-1123210232211311-3331131321222003"></a>

#### `more_option.idle_timeout` property

Type: `"number"`. Optional.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with an HTTP 504 (Gateway Timeout) error code if no upstream response
header has been received, otherwise the stream is reset.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(3600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3600000,
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
    "ves.io.schema.rules.uint32.lte": "3600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "3600000"
  }
}
```

<a id="canonical-2320201121311013-3222113310331220-2022031231232100-2123200131322331-0022100121101013-2013333110300032-2010331000023223-0300102132301330"></a>

<a id="canonical-2121223031310030-1222133101220332-3311323101100233-0320301232213222-0120133002322103-0211002223233310-0310222223002011-3103313203111033"></a>

#### `more_option.max_request_header_size` property

Type: `"number"`. Optional.

The maximum request header size for downstream connections, in KiB. An HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(96),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 96,
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
    "ves.io.schema.rules.uint32.lte": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "96"
  }
}
```

<a id="canonical-3210333230300103-1110210333313003-2032322200012210-1233030221201010-1233313323033321-3011031120001131-2301022303033130-2303121213302013"></a>

<a id="canonical-3332002333330303-1122111131131222-1301320323301303-1201031233222010-1332102101220211-0302101322311001-2120103130000001-2020230111001003"></a>

#### `more_option.max_requests_per_connection` property

Type: `"number"`. Optional.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [no_request_limit_per_connection](resources--http_loadbalancer--reference--group-021.md#canonical-1023102010030230-2332223133022232-3123113033013210-3231031103313320-3102130133132220-1233310023222222-3310320102211232-1232211101323103): complete subsection reference.

- [request_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-1002021132200112-1120200221201022-0220211301011220-0101022100322133-2202230030112132-3203321123003202-3201221131021020-2201202320300321): complete subsection reference.

<a id="canonical-0001203023021003-2201113102202312-1201013113011320-0023211001113321-3323330130202112-1231201112202203-0021331321032321-1311201312313220"></a>

<a id="canonical-2331331312212022-3101101033232301-3001200303323031-0312203332021332-0223201302320132-1123030031201003-2112212023310013-3111201102310320"></a>

#### `more_option.request_cookies_to_remove` property

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

- [request_headers_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-0212010211023031-3133301021032123-0202022321022330-0302322100113223-2120010011130022-1322200102120032-0121222313312022-2202030212133020): complete subsection reference.

<a id="canonical-2000001000331333-3102211222113332-2222103130112201-2313221022102210-2322133333023210-2223320322213100-3300301032032203-0303020131202333"></a>

<a id="canonical-1032312010120111-2301033231133103-1012121302112132-0301123012120223-0332310232101102-1121030101033333-3231313210322323-2033201031030311"></a>

#### `more_option.request_headers_to_remove` property

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

- [response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322): complete subsection reference.

<a id="canonical-1110233201010302-3032033303013001-2003113233111332-2123112122113013-1110033223033131-1123332232233112-1231013300301132-2321231333012323"></a>
