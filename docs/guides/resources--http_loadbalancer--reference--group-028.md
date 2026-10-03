---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2120013222300121-1021302113002212-1112113123112023-2130200303000000-3000220132101001-2333212302222012-1320331300211331-0222011012331023"></a>

## name property — metadata / 000130032202 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3002102300032310-2031320312322031-0132013331122222-2211230112320223-0302023332203133-0120102130123321-3220102230000300-0023103223212233"></a>

## Next pages — metadata / 000130032202 / 6

- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-027.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1202001010122130-1232112232211102-0102102132023332-3020023330320231-3120312230233312-0111230122132322-0202011203322330-2223113000302022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223232030300201-1321220231230010-3230001103122323-2332122320321333-0300113032132000-3123033100103023-1310301230323203-1220230213033320"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing — waf_skip_processing / 132000231022 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-027.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-027.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-027.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing

<a id="canonical-2103331220022032-1303010102302121-3012001201200032-2322133213002220-2200112100332331-2021001210131233-1332322132013133-2303021032100233"></a>

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
waf_skip_processing = {}
```

<a id="canonical-3321230023120011-0000321330303210-0002300122011002-2210103303322320-0220123231220123-0230213111203302-1003113212031232-3131322210300310"></a>

## Direct properties — waf_skip_processing / 132000231022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110011311323223-2311030110113201-3331302220331211-0202010203010203-2211330220021303-0103103320013310-0220223333202111-1001003033303121"></a>

## Next pages — waf_skip_processing / 132000231022 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-027.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2300231203030031-1332023310203311-2011202020030201-3003033103130323-3301122311203303-1021210131330310-0012111123120220-1030113003313033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023000310333012-2010130020333100-3002330112122013-2200030113130103-3320101013021102-3300020011230321-2013203113300100-1333111212013210"></a>

## waf_exclusion.waf_exclusion_policy — waf_exclusion_policy / 200101110102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-027.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- waf_exclusion.waf_exclusion_policy

<a id="canonical-3103213001110000-2300333222120102-0301100113022213-2010111332313102-3302200112113110-0102131201033031-0201110021302021-3213232131103113"></a>

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
waf_exclusion_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222113023321331-3121200212231031-3102111323221312-1133203120303323-1131320132202033-2021312233222231-1001300301022313-0202101033112103"></a>

## Direct properties — waf_exclusion_policy / 200101110102 / 3

<a id="canonical-1133001222213233-1131031202331200-2300212112020303-2031003200020131-3121332133120320-2212310300032100-0320112030023301-3301310031231201"></a>

<a id="canonical-3320100100021012-0001323013312133-3330233303332122-0300302333320131-2011230300231310-2202331333332133-1332201123333123-2120333210103302"></a>

## name property — waf_exclusion_policy / 200101110102 / 4

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

<a id="canonical-0213300212012230-0331013120103222-0323220202202312-1001303031103013-2103313330223233-2103112223032203-1200301223212102-3123320300100023"></a>

<a id="canonical-2211003231113221-0000000122012311-2322302213133211-3111302201011123-0031211103113023-2023211321213213-0013131230230311-0201221130210223"></a>

## namespace property — waf_exclusion_policy / 200101110102 / 5

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

<a id="canonical-1013311010201121-2020010103011113-1330323222321131-0123320302232200-2111310111130311-1303033210000310-0110120303003203-3130000201031310"></a>

<a id="canonical-2111302230303323-3202011310203101-1300001300303102-2110023210203320-3111231032222223-1211331022003003-2233230032023300-3333100112230011"></a>

## tenant property — waf_exclusion_policy / 200101110102 / 6

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

<a id="canonical-1013313101131021-0103331201310331-1320321201012330-3130213132211201-0100323331221233-1003230213322213-3300021130220312-2102222312103111"></a>

## Next pages — waf_exclusion_policy / 200101110102 / 7

- [waf_exclusion](resources--http_loadbalancer--reference--group-027.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
