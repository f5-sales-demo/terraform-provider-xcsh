---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1121233023131220-1003111220221011-0133310311113000-1021131213123031-3323030010330300-0200200213101021-1233030130032333-0013322100312031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.inline_rate_limiter` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter

<a id="canonical-3032133021102233-1331030100031322-3030202220021113-1223322100201011-3110301231002132-0101010211022101-2100000313002220-0202120031122033"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inline rate limiter.

Additional upstream details:

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("threshold"),
  validators.ConflictingObjectAttributes("ref_user_id",
    "use_http_lb_user_id")}
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
  "x-ves-oneof-field-count_by_choice": "[\"ref_user_id\",\"use_http_lb_user_id\"]"
}
```

Terraform syntax:

```terraform
inline_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-3323223100331201-2030110331220232-3100302003012123-1213110000013232-2300230131222131-3120012132201221-0110101310201221-3101233023211300"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.inline_rate_limiter`

- [ref_user_id](resources--http_loadbalancer--reference--group-007.md#canonical-1322000210130123-0113311310103133-3112203131210121-0320302333110311-0113031113200002-3223121312302200-1113233233203222-3000010213202003): complete subsection reference.

<a id="canonical-1311232223200320-3301220321321211-1112320123123300-0113023000002303-1012120320103000-1232021121221211-0002111333301032-3231202000323111"></a>

<a id="canonical-3332012121233030-1201333121111211-3313002113033302-2330003323232123-0023213330212102-1110123203322201-2100221132211112-3231023230110220"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.threshold` property

Type: `"number"`. Optional.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-2130211202000102-1122131022031000-1302113200201232-0220323001230131-0330332013020312-3032211201003303-1002000131331113-0110000231130312"></a>

<a id="canonical-2230211213011333-2203130221132121-2212222221321110-1213312130210133-1031133220013233-2023002210230022-3031233323330323-3313103100031221"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.unit` property

Type: `"string"`. Optional.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["HOUR","MINUTE","SECOND"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SECOND",
    "MINUTE",
    "HOUR"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [use_http_lb_user_id](resources--http_loadbalancer--reference--group-007.md#canonical-3330031100020111-3030203233132221-1000321133200203-1321313131010003-3122213211133010-3300211303020132-1200031300300131-2210032020311321): complete subsection reference.

<a id="canonical-1322000210130123-0113311310103133-3112203131210121-0320302333110311-0113031113200002-3223121312302200-1113233233203222-3000010213202003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-007.md#canonical-1121233023131220-1003111220221011-0133310311113000-1021131213123031-3323030010330300-0200200213101021-1233030130032333-0013322100312031)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id

<a id="canonical-0232331112003121-0302311101311331-2331313120203301-1212132313310033-1202003333201312-2300103323213001-0033223300222010-0302110032012222"></a>

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
ref_user_id {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301110103323102-1232122011123301-0220223023223323-0301021000111232-0103333313111023-2223302300232011-2121310122311101-0002202100223123"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id`

<a id="canonical-1313010013102033-2130020321321311-1210233100113032-3103323032023110-2333010023121012-3230133132223220-0001201101231200-0110323322302131"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0120213103130302-1201010301322301-1221020302323121-2211000313023332-2021213221122122-2231312103230020-1133021322211001-1311311033100013"></a>

<a id="canonical-1222302331203102-0333302210231033-2000133012013122-3203133033212131-3032222230210320-1322230231211222-2310300322232323-1330221032103320"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1102030003221110-3032012132102022-3202301330132322-0210113301221213-2031110021122230-3221132132321111-2101000132101320-2202301212302032"></a>

<a id="canonical-0010230132013112-1012111231310300-3303112233302303-3112233223000201-2130022233332231-2000011103331212-3211020000310231-1113110133111012"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3330031100020111-3030203233132221-1000321133200203-1321313131010003-3122213211133010-3300211303020132-1200031300300131-2210032020311321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-007.md#canonical-1121233023131220-1003111220221011-0133310311113000-1021131213123031-3323030010330300-0200200213101021-1233030130032333-0013322100312031)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-2203110110331322-3032012202233323-1113021333010211-1232210210220321-1332202303223133-0001323102023022-0221103020311310-3331000102013303"></a>

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
use_http_lb_user_id = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031121110211321-0222223200032133-1223103211312023-1012222201100010-0323202231130000-2323201330321110-1222232311021302-3202202311112033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.ref_rate_limiter` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- api_rate_limit.api_endpoint_rules.ref_rate_limiter

<a id="canonical-3112213233313300-2121330033100010-0303132133030301-3320002023002222-2232220020303001-3300000201222122-2210013213011210-3033321100331201"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Additional upstream details:

Reference to a stored rate-limiter object for this scoped rule. Select exactly one of
ref\_rate\_limiter and inline\_rate\_limiter.

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
ref_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-3021303010000330-3201110102231212-0231311322320232-2020300200311102-1132232230311013-0321223012302123-1302203001200010-3123030200102300"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.ref_rate_limiter`

<a id="canonical-2100123312323022-2130220123200132-3201213220130011-3122103112103300-1032030132220232-3212003001102133-1203232212032221-0310033112223322"></a>

#### `api_rate_limit.api_endpoint_rules.ref_rate_limiter.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1302033100220222-2131010303211321-1211122221202311-1232112033101330-2012322002323201-3010130332322321-3301232332132111-2122333101223031"></a>

<a id="canonical-3131021130030200-1222223001000203-3003132131002010-2022022112111010-0220233022013311-1002122210013303-1210130100333113-2222112111021001"></a>

#### `api_rate_limit.api_endpoint_rules.ref_rate_limiter.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1222332210122300-3120310010333223-2032021230211232-3231123131001020-0111331013132320-1213103020100023-0213111231301032-1111312213330321"></a>

<a id="canonical-1223100302310010-0122030302102321-2232321032220311-2220211100323301-2003131321003230-3011321331020230-0002210230132012-3303001320323312"></a>

#### `api_rate_limit.api_endpoint_rules.ref_rate_limiter.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- api_rate_limit.api_endpoint_rules.request_matcher

<a id="canonical-1320020013211132-2031013002002013-1203201303313130-1213200320011033-2103201322111302-0133123010012232-3213222101200313-3302231313332121"></a>

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

<a id="canonical-0012101011333320-2312001320032131-1123020300220002-1112121103001102-3020013131033233-3012130310233311-1320213033110110-0021113230011231"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher`

- [cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-007.md#canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221): complete subsection reference.

- [jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-007.md#canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321): complete subsection reference.

<a id="canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers

<a id="canonical-1312232002030230-3031023203302003-2233331213111113-3230312233323132-0322031211203333-1313120203133101-1023232211001021-1132122212221201"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2020312320222103-0332033302031213-2313130333022222-3202213001323220-2332022203221102-0103113000321203-1112203311022013-1201001210131123"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers`

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-2131023022232100-1313031230102113-1322331330011233-2003100010212013-1221031002330332-2130233033302100-1301100011232323-1232303001203211): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-007.md#canonical-2122101123213330-3233231220232002-1332013023321002-0223022233200113-3332002102113313-1330000032111110-3200220223030122-2200323201130011): complete subsection reference.

<a id="canonical-2013322200103021-2332012002310130-3302020130113310-1310032220221231-2133100321132131-0322310111003223-2333013332212003-2113103000213000"></a>

<a id="canonical-3110132300130231-1001323012021233-1322002021013221-2222010013012322-0222103321332112-1030322002210110-1200022230012302-2210330102321322"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.invert_matcher` property

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

- [item](resources--http_loadbalancer--reference--group-007.md#canonical-1012000023300311-2100310300202101-0330200133301120-2132232320103013-2222300100123211-2232201001100122-1233230020202012-1012030213200011): complete subsection reference.

<a id="canonical-0132130122112023-0001303101133310-3313223232121300-2023203133000321-1002030020230220-1130121023321113-1331102323332203-2023323231312321"></a>

<a id="canonical-3001101202131231-0110023320103220-2220020210200023-1330031100200021-2111303332203312-0113312133301301-3221312320310023-3131130210013001"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2131023022232100-1313031230102113-1322331330011233-2003100010212013-1221031002330332-2130233033302100-1301100011232323-1232303001203211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-0223031032222220-2011222120210323-0313210222321312-3332020121102012-3023233220310222-3333231323122322-3031110121133130-0032123332230133"></a>

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

<a id="canonical-2122101123213330-3233231220232002-1332013023321002-0223022233200113-3332002102113313-1330000032111110-3200220223030122-2200323201130011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-0221221121110211-0011320032123021-3132211231112112-1031110122320121-2302012101123220-2031032020302320-0032120323300121-1110020102332120"></a>

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

<a id="canonical-1012000023300311-2100310300202101-0330200133301120-2132232320103013-2222300100123211-2232201001100122-1233230020202012-1012030213200011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item

<a id="canonical-3323300233200130-3222030201211321-2102131320212131-2203103333102010-1213303132000312-3132301002333323-3130123102102012-1122222003121301"></a>

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

<a id="canonical-3213103232011222-0022332000102330-0120032000211003-0113331123000202-3330222032132310-0310302300233131-2303222212011000-3300221323001312"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item`

<a id="canonical-1021212103122213-2110101122232131-1201331031002232-3210213002003300-2220212033131120-1133212213010032-1033032323322001-0203211012020313"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item.exact_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0130310310022322-0312321331113202-3112300130321311-1100122300220232-3201231332322113-2001102301211233-2322233020121011-3000122121033013"></a>

<a id="canonical-3302213221000303-3002001133003121-3222323312231013-0210131010333122-1032322000322021-1032130203132033-2201001113113333-0100030123021332"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item.regex_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3231101030222201-0003100232211013-2132030230202233-3111120130101223-0113103210130100-1331210113001201-2122302122332323-3302230331202303"></a>

<a id="canonical-1332222213233103-3202310121301312-2033102231100323-1310221222013030-2000322232220133-3233310311011123-0002302003212101-1021210023210012"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item.transformers` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- api_rate_limit.api_endpoint_rules.request_matcher.headers

<a id="canonical-1112032022303031-1310321332323132-3301000302203120-0220003021312212-3112123202311302-2320130023332300-0123211232001320-1203213232220003"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1211031310010032-2130333323210311-3010331113031231-3211203030320113-1202211232323113-0201302313022322-3232111132100011-0312203333311032"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.headers`

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-2110322001203333-3210020311020010-3132330023321212-0020011103201020-2012321031321121-3333022200211103-0001113303232322-2131012131203131): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-007.md#canonical-0131011301333221-1213032033122212-3133122030003223-3330230002221201-0010033011321103-0222021111101122-1333312230221231-1320212320131220): complete subsection reference.

<a id="canonical-0003311303201101-3001301131210220-0113133231023113-0100230323012101-3202102331123210-3131130302323002-0121012211100310-3203331201110020"></a>

<a id="canonical-2113003113023300-2333121211303002-2121303322202103-2213101220112000-0333102030120022-1311303121102033-1212320023130211-2031221011330013"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.headers.invert_matcher` property

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

- [item](resources--http_loadbalancer--reference--group-007.md#canonical-1110320202331212-3101201330223333-2012013200101223-0212203002330323-0320122032313113-2301211211232203-0202030211022332-0311030010103023): complete subsection reference.

<a id="canonical-0211313013102023-0210200002230113-2132112032011012-3200111020100310-2311202323323100-1003110023012231-1111100200313033-1020303131112123"></a>

<a id="canonical-2102103113310011-0220023030312103-1233331122133332-3222131220031311-3120031311222113-3221221321130110-2311121030132022-0201303022001120"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.headers.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2110322001203333-3210020311020010-3132330023321212-0020011103201020-2012321031321121-3333022200211103-0001113303232322-2131012131203131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present

<a id="canonical-0002133200213133-2001303203103011-3121130030231223-0101333322221100-2210330112200201-1220020122131101-3311021202011110-1103333230223202"></a>

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

<a id="canonical-0131011301333221-1213032033122212-3133122030003223-3330230002221201-0010033011321103-0222021111101122-1333312230221231-1320212320131220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present

<a id="canonical-3103033330003201-2033103203322000-0110123113003200-1110120011121023-3322310311023002-1221122230002003-2031021332022101-0320223210203032"></a>

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

<a id="canonical-1110320202331212-3101201330223333-2012013200101223-0212203002330323-0320122032313113-2301211211232203-0202030211022332-0311030010103023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.item

<a id="canonical-1310102202200201-0012131313200012-2310132112111013-2313203023312123-2012111100002032-2230122032011323-3221101013332230-3023120021133201"></a>

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

<a id="canonical-1332213103223311-1130212010212112-3103111002311032-1133113201211233-2330130223122032-2223300202023303-1330321032223133-3322311303133301"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.headers.item`

<a id="canonical-2133210111223122-0211131002310221-0100231003132032-1121203031210100-1002012323121002-2200101203201021-3202022203132200-1311233011122213"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.headers.item.exact_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1331210103033311-0231111001123001-0010123322232320-2100333310010120-0111213212313100-3022131011233301-3113212323213110-3122111033333110"></a>

<a id="canonical-2023113111020323-3002012212210320-0300213120111233-2312033212311020-2202122211312301-2111101332310100-3332330112230301-2333231001232320"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.headers.item.regex_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1331332311312310-0212023323302022-1103300200320031-3203022133313333-3103111102232130-0010123322311122-1030100302332030-0123231323232303"></a>

<a id="canonical-2030011030201203-3213103310201312-2320311011220113-3232003200201033-3113020033313010-1133322221131021-3310000221320302-3013020011122233"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.headers.item.transformers` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims

<a id="canonical-0033131120213101-2021101213122313-1112301101110023-0020303203303030-2330012203010213-0231302000220110-0012000120012123-0022210112113322"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1102303103120111-1033221103122221-0110023031300311-2111300010311303-1102310331312003-1210002120321220-3002213232122302-3112320320301033"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims`

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-2202013113121212-0312302301312321-2203331102213021-3321211020110121-1301132230300211-1031011303212011-1100112222210020-3132032331330122): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-007.md#canonical-2012002332011020-0333231311122122-2012110033333113-0120300002132332-0203321002111313-2130200102000022-0200023110302321-2323001020122133): complete subsection reference.

<a id="canonical-2122231301331010-1320312333312003-1333001210032332-3133321021031021-1121310322232221-1102212011201031-1210332132002020-2122103210333102"></a>

<a id="canonical-0330200333130200-2232333110333122-3302202120023122-1131320000120122-2231031021103310-2220202222322202-2012023101012221-0211020213301323"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.invert_matcher` property

Type: `"bool"`. Optional.

Invert Matcher. Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--http_loadbalancer--reference--group-007.md#canonical-2120111123131100-2123300020302112-2331201211201300-3223211132210032-3100021311223202-2012020210031032-1230012021112223-2200322103311310): complete subsection reference.

<a id="canonical-3132211312321102-1113230313203121-0100021113222112-1033010200002002-3110030330231230-2012320131232123-1333002110200220-1322013302223211"></a>

<a id="canonical-0023321122032003-3132221333203213-0232302330321203-1233132332300332-2312011033023311-2101130213212313-3001232032203010-1201301001122002"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.name` property

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2202013113121212-0312302301312321-2203331102213021-3321211020110121-1301132230300211-1031011303212011-1100112222210020-3132032331330122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-2213212220111130-2210120211112303-1130310210031133-2121103010330113-0332100101232323-2000333022101030-0313022222222000-0333312321011210"></a>

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

<a id="canonical-2012002332011020-0333231311122122-2012110033333113-0120300002132332-0203321002111313-2130200102000022-0200023110302321-2323001020122133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_present

<a id="canonical-2312021030100012-0322330333213110-0001221031311302-2110023221303221-1121020023310031-3001112312023122-2321222121023113-0331010320331021"></a>

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

<a id="canonical-2120111123131100-2123300020302112-2331201211201300-3223211132210032-3100021311223202-2012020210031032-1230012021112223-2200322103311310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item

<a id="canonical-0111033122123022-1130321103113202-0223200230000033-2013111311033021-3323213123011032-2033001202310023-2111203212310123-3210111313320030"></a>

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

<a id="canonical-1330122212003101-1332121032223100-1111011213000222-1210012220203032-1123023300301200-1213333233002013-3113123311120321-2013221123212110"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item`

<a id="canonical-2312202120212310-1202200320110121-2020332330333023-3011320030031232-1320013020130022-0102132033232012-0130032203311322-3232121101130111"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item.exact_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3203213223302010-0032123323231331-3102323133330002-2112103131130121-2321312120231221-1020303310122111-0200311213203031-1323013020200323"></a>

<a id="canonical-0231213313201313-2230303030122302-3020100022310111-3023220032020331-1023300332331001-1230020333033233-1003231312221220-2130221203202023"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item.regex_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1111203301223120-2132332121032212-1221022221300332-1330302310322310-1020122230012110-2333032321233302-2223233203022200-3010213001123030"></a>

<a id="canonical-2202111113000123-2100132322100112-3230220232300311-2321123031333331-0021111011023023-3213031302202101-0023212032332233-2112233001203331"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item.transformers` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params

<a id="canonical-0320330213122122-1313120301101122-0000330133300220-1332310200101313-1303200103110131-2102031203002021-1330111302122001-3102101320132001"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1112332003110330-0033023331230301-1100310133233211-2321032233131102-2133200232032312-1031012311300222-2303102133303020-0113221031121001"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.query_params`

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-2000312303011131-1310130332222303-0331322333203013-1230121012023223-3311022133012103-2121021032023033-3132320321213311-3011310032330101): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-007.md#canonical-2200220323213322-2000010132222101-2221223131230120-0012101310002203-2111131300000130-1311012332213010-1021120223102320-2220102221332102): complete subsection reference.

<a id="canonical-3021330220001013-1331033222021121-3102121302132200-0112203312110112-2221102232021211-2023301211321311-1131310210303222-3022201123003223"></a>

<a id="canonical-1301122131111121-3033033231233232-2110203110112132-0302022103001200-0222212311100103-1221130231210333-3312012013232310-0011030133303101"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.query_params.invert_matcher` property

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--http_loadbalancer--reference--group-007.md#canonical-0302201210130133-3213120002023132-1131210033000211-1302122002210033-2333012132111202-2011200200332122-3002022131222223-0231101230312232): complete subsection reference.

<a id="canonical-3103003112002332-3010112013033031-2213210133023300-1223031322303030-2032012020130210-3221130312213030-3111010012323013-1023202312321113"></a>

<a id="canonical-1210013120021332-2310221011103010-0310333201233213-1332012122113013-0203131011303231-0213032321111130-2101302121100230-3100223130023313"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.query_params.key` property

Type: `"string"`. Optional.

A case-sensitive HTTP query parameter name.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2000312303011131-1310130332222303-0331322333203013-1230121012023223-3311022133012103-2121021032023033-3132320321213311-3011310032330101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-007.md#canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_not_present

<a id="canonical-0111010011022133-3123323033111131-0230120131203231-1233010222002033-1203123113023311-3111202132303132-2201301302301110-1330133332110123"></a>

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

<a id="canonical-2200220323213322-2000010132222101-2221223131230120-0012101310002203-2111131300000130-1311012332213010-1021120223102320-2220102221332102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-007.md#canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_present

<a id="canonical-3102033011333210-0311103333120023-0030100000101321-0102331030113201-0113233231232310-1212033103322112-3321121222123011-3130101220111121"></a>

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

<a id="canonical-0302201210130133-3213120002023132-1131210033000211-1302122002210033-2333012132111202-2011200200332122-3002022131222223-0231101230312232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.query_params.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-007.md#canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params.item

<a id="canonical-2220000320223113-0302112200020230-3121020032101303-1122203032333332-1211301032133233-0101312312101122-3220021212201311-1123323001210111"></a>

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

<a id="canonical-1001013132003023-3232021101232202-2320013211011133-1133120333101230-0220102022133323-3303133101001012-3330321002011220-2223111233232000"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.query_params.item`

<a id="canonical-1313330222022100-3312122223233133-2323300210201222-0113021130333313-3333221322230110-1233231113002100-2213201021311000-1313032000133200"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.query_params.item.exact_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1331100200123200-3031033231133032-3121032231302023-3100202200021330-0331332112321133-1000122203013232-0120233101033010-0112320112003010"></a>

<a id="canonical-2123313122123013-1310320131031301-1031331031202121-0120112013111331-1001033303220313-1321023233022013-3220031011003213-3131203300111300"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.query_params.item.regex_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0120313202221230-3131012213211000-1303232203300113-2020313300110121-1211211332033013-1203202221023021-2111130112113102-2113102011213210"></a>

<a id="canonical-0233311202011130-1021103322302132-1121033111220210-1010022112101022-0303031020001013-3220220231133000-0003132221230323-3120031002112023"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.query_params.item.transformers` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- api_rate_limit.bypass_rate_limiting_rules

<a id="canonical-0123312320102111-3213033002332023-0232203123102221-1001130102221023-2113023301230111-1233313013310122-3100113001110231-2301001312200003"></a>

Type: `"object"`. single nested block, Optional.

This category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

Receipt-pinned upstream constraints:

```json
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
bypass_rate_limiting_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231010331002021-1220201223030103-2022123033023120-3333333201322020-3300330032201230-2331303322002223-2021212003323113-1123222322100211"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules`

- [bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033): complete subsection reference.

<a id="canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules

<a id="canonical-3210011231220003-3320121223332000-3302121222323300-0320131302113210-3110323330132221-1022123010110213-2231120123020032-0312230313123121"></a>

Type: `"object"`. list nested block, Optional.

This category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("any_url",
    "api_endpoint"),
  validators.ConflictingListObjectAttributes("any_url",
    "api_groups"),
  validators.ConflictingListObjectAttributes("any_url",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "api_groups"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_groups",
    "base_path")}
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
bypass_rate_limiting_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112113103223303-1312021011130332-0033130010200032-2022110023223322-2131200131110002-1312200311000000-0301330313010123-1031233122031320"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules`

- [any_domain](resources--http_loadbalancer--reference--group-007.md#canonical-0132120010232023-2120333332011212-3022310321203131-1202230000000123-2210313112230121-0000313103000112-0013311222103313-0033021121121131): complete subsection reference.

- [any_url](resources--http_loadbalancer--reference--group-007.md#canonical-0103121101322112-2100332012000200-0012003012122131-0033010133230322-3021013003201133-3020103323322013-1201301211323101-0220113202302020): complete subsection reference.

- [api_endpoint](resources--http_loadbalancer--reference--group-007.md#canonical-2110031310101110-1303103222112222-1320122301213130-0011200231101121-2302201210300100-2322110233030333-2120311323010233-3231113023031103): complete subsection reference.

- [api_groups](resources--http_loadbalancer--reference--group-007.md#canonical-3200113020100230-1001333030320020-2031211032320122-1230230223121200-3021001330030102-3221120020020100-3200001032100131-3330002000333201): complete subsection reference.

<a id="canonical-1330133330101211-2113321332210330-3321303312302210-1232030312033201-2232010301013322-0113322231123010-3030011121031102-3033123021132323"></a>

<a id="canonical-3022332121031311-1032012031333310-0110030033001320-1130220333322203-1301020010211332-1110300221310001-1030222310122211-2332130212001211"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.base_path` property

Type: `"string"`. Optional.

Exclusive with \[any\_url api\_endpoint api\_groups\] The base path which this validation applies
to.

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
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223): complete subsection reference.

- [request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222): complete subsection reference.

<a id="canonical-2113330011220321-2003310231110000-1003000031333203-3112102001233311-3101310013212322-1202202111330133-2100210231230203-3121300021220302"></a>

<a id="canonical-3000310023303321-3221030022033313-2033133002331201-2100330100210102-1220331110021213-1103301012311131-0230200001212002-0211330203200312"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.specific_domain` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0132120010232023-2120333332011212-3022310321203131-1202230000000123-2210313112230121-0000313103000112-0013311222103313-0033021121121131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain

<a id="canonical-0133111030321200-1213122130100333-1131321000131330-0302332132012332-3031101323310320-3313201212300100-0312003030132330-0220221113211122"></a>

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

<a id="canonical-0103121101322112-2100332012000200-0012003012122131-0033010133230322-3021013003201133-3020103323322013-1201301211323101-0220113202302020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url

<a id="canonical-3321200132022202-3131231010032230-2313331022202330-3022022123133200-3012220230012121-1021110331022320-3213211131023332-0102013322030322"></a>

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
any_url = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110031310101110-1303103222112222-1320122301213130-0011200231101121-2302201210300100-2322110233030333-2120311323010233-3231113023031103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint

<a id="canonical-3230030233302333-0213010233312000-3001233322002120-2300212233212310-2030212203233213-2031102103222111-3302311120213111-1210233320330212"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
api_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203233203012211-3011302202030310-2103013122122000-2121103200332313-1212013132032001-0223313322131212-2333133313033213-1312230302331100"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint`

<a id="canonical-3032121031003022-1102221111012322-3300231132001003-0030111303103100-0313000002311112-1221312101332331-0133120123110300-3212010330210220"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1321111203032120-1312202221100022-3213210112233010-2011323011111133-0023133312210003-0313313203221030-1202222223312003-1020103030113301"></a>

<a id="canonical-0133132313113331-0202303001313012-0203020221002333-2013232212202220-3212312301023300-2012212333212312-3001033220030133-3230311220211210"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint.path` property

Type: `"string"`. Optional.

Path. Path to be matched.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-3200113020100230-1001333030320020-2031211032320122-1230230223121200-3021001330030102-3221120020020100-3200001032100131-3330002000333201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups

<a id="canonical-3203002232121121-2223200201110121-2001233012033212-2020101010301032-3300220121121300-2213303003122112-3213031311010012-2030333301313220"></a>

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

<a id="canonical-0030032233132000-2030012021313311-1100202120320200-0020032223200013-0201223311130123-3203210321113310-2301202333101233-3310021000302330"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups`

<a id="canonical-0013112121233013-0112330220303020-3203300332333320-2200322123022103-3022020202232100-0010302121022320-2120133221101230-2023021332323203"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups.api_groups` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher

<a id="canonical-0313112301223312-3013231202200230-0302022303222131-3031221132223213-1212220001301313-2000303310231210-0122313233131012-0202023200303113"></a>

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

<a id="canonical-3303212231110200-0010323032012301-2300130213103202-1030002212230123-1232110023132213-1212211210320312-3313011130230223-0001232101133301"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`

- [any_client](resources--http_loadbalancer--reference--group-007.md#canonical-2333010110323301-3130300032202001-3020212211033302-3200322232000223-0223330222020103-0322211201320330-0220120330030111-1303123033210123): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-007.md#canonical-1313032013200331-1303110023313222-1201330133202132-1111003201020323-3220033221302323-2122331002111013-0332321330110212-3122301001302210): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-007.md#canonical-3103323302220132-1003100003223021-0223121213302210-3222013302313130-0200313320322222-0003120333232123-0201323021310312-3320113320330012): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2310212103013113-1100303220012111-1332102213102101-2312202131121300-1230013103032011-3320023310303221-2002113131100121-0013210210133211): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-007.md#canonical-3113312003231312-0131331320123120-2112122033223023-3200031103311113-1301033132002210-3220200030300103-2132303312313212-0212003033102003): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2112210201211101-0212311022213221-1022100111213003-0100111233033032-2110021210200322-3102121001111103-1331231323211323-2223320033022133): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-007.md#canonical-1200113333332102-3331022332002031-1120103010102301-2121331130312211-2212212010121333-2200210300321022-1003200210000330-3113200103222123): complete subsection reference.

- [ip_threat_category_list](resources--http_loadbalancer--reference--group-007.md#canonical-3312120300100220-2330131103032312-0202010330213232-3020322331303310-0212132301011221-1033003031230310-3321101123123011-1122312120111220): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3031203222211302-2303123322301120-3302130021100031-1000101233131220-3101201020012210-1301302133100302-2200220030010101-2113133013012220): complete subsection reference.

<a id="canonical-2333010110323301-3130300032202001-3020212211033302-3200322232000223-0223330222020103-0322211201320330-0220120330030111-1303123033210123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client

<a id="canonical-2122303312133030-0311200301201012-3300202320103110-2210120103232031-0303330202101111-3331031110002031-1201023021100231-1023321331032203"></a>

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

<a id="canonical-1313032013200331-1303110023313222-1201330133202132-1111003201020323-3220033221302323-2122331002111013-0332321330110212-3122301001302210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip

<a id="canonical-0013300312001301-2302301331003330-2003110011133013-0233332312032230-1020131121002112-3330223122203030-2133010310003221-2111230001210210"></a>

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

<a id="canonical-3103323302220132-1003100003223021-0223121213302210-3222013302313130-0200313320322222-0003120333232123-0201323021310312-3320113320330012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list

<a id="canonical-2113330312123003-3202321033220022-0203323130222003-1130311221332331-2122220011212103-3131121310313121-1023030200133032-2002023001232311"></a>

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

<a id="canonical-1202332221133222-2100220103202131-1111310102233231-3122012220302101-0300120012100133-2011313101123320-2000232001013012-0032221232212121"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list`

<a id="canonical-0110002022202321-2110101332320121-1031320102021331-3222131023013212-1013102310301023-0000132030111101-2200211133123010-3333022120102032"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list.as_numbers` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2310212103013113-1100303220012111-1332102213102101-2312202131121300-1230013103032011-3320023310303221-2002113131100121-0013210210133211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher

<a id="canonical-0312220303220013-2030123102021332-1031220202130320-1113122000123121-3002123002301302-0332003313212222-1210022010023023-1010331112011121"></a>

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

<a id="canonical-3303122322002010-3212303200212122-1322111231032311-2202313112201133-1301211033222011-1111201331311211-0330303300132012-2231111213113013"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher`

- [asn_sets](resources--http_loadbalancer--reference--group-007.md#canonical-0300233232013012-1313310113122010-1002110223320210-1201121002333210-0012131002021002-2023121313300101-0133300103102101-3013120320132201): complete subsection reference.

<a id="canonical-0300233232013012-1313310113122010-1002110223320210-1201121002333210-0012131002021002-2023121313300101-0133300103102101-3013120320132201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2310212103013113-1100303220012111-1332102213102101-2312202131121300-1230013103032011-3320023310303221-2002113131100121-0013210210133211)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-2031332202110200-3210321322222121-3333101320000332-0202301113311222-0122101033121021-2003321212022200-3011123003130123-1003330331302200"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0220331311302233-3133322033301020-3312111200332030-1313333210010130-0213303301221311-3120100233310230-0133232230201103-1103011303002302"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets`

<a id="canonical-3023303322122121-0333313023033203-3132311331023311-1233303331331330-2132201202012001-1322002201002321-2002232031311103-0303323021202132"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets.kind` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1013023200132221-0021010311121333-2230310103112030-2303100021333021-1023103021330301-2321132112130312-2213222113333200-1130030220312111"></a>

<a id="canonical-3120201212323010-2122130132022021-0102132112030013-0223111111211310-2303333220202023-0120123332220103-0031120312102320-1111321223212221"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0213120113311302-0110231210120211-3331233033112000-0313133022232032-1010332310232023-3301221003203003-1310233232300310-3101032320101333"></a>

<a id="canonical-2313002122303133-3111110313011310-1132311212103031-1013012213013020-2233013120221210-0320333330103312-3311100330011301-0100031112321031"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0030330110113131-3032130230302102-3111202232212021-1210122021112110-1000233001210112-0323122322230020-3022121322202212-2200210022301100"></a>

<a id="canonical-0032111212210213-0121320230301200-3323132333211303-2001300232310012-3303222113022320-3322301331300033-1211032230301323-2201131010311100"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0133122021221021-3120313110121001-0223200022312211-0222212310200233-2110220112311130-1210110213302301-2330003030133122-3000112033011111"></a>

<a id="canonical-0023102131231331-2003213021000331-0220103222233030-2000233023110003-3322021220120031-2322023221230023-1123312320003032-1222312333020032"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets.uid` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3113312003231312-0131331320123120-2112122033223023-3200031103311113-1301033132002210-3220200030300103-2132303312313212-0212003033102003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector

<a id="canonical-1332002202232331-1012330000010133-3310103332023221-2032020323332301-3123130000121223-3003323021233020-3223133321302001-0100331112233102"></a>

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

<a id="canonical-1112211322222100-2223213203331120-1022032332313123-2003102311130310-1130131203222210-2200133132100223-0001102032013213-2331230111130020"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector`

<a id="canonical-3021100201132312-0122311133130223-3011323113010323-0020000303210231-0120120131122322-3103210301000133-0130021032113213-3233020132312302"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector.expressions` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2112210201211101-0212311022213221-1022100111213003-0100111233033032-2110021210200322-3102121001111103-1331231323211323-2223320033022133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher

<a id="canonical-3333122010313100-3120120131332123-2233321323230230-2030230130130012-0022331122132223-2022133302313030-0331121330202121-3022230201213113"></a>

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

<a id="canonical-0102101011201203-0021011330300233-1133231032010012-1211331330302233-2112022232013223-2323030133032333-1230323022033022-0102021123012201"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher`

<a id="canonical-1002231312030322-3311131333200230-1020110032102021-2330233333130032-3023210123120323-0332011131213231-3231131010331311-2211021220320233"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.invert_matcher` property

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

- [prefix_sets](resources--http_loadbalancer--reference--group-007.md#canonical-3313202030031003-1123103301111002-0002111323113100-1022023213020212-3231321131131320-2222210011100203-2313232332132003-2111232003313111): complete subsection reference.

<a id="canonical-3313202030031003-1123103301111002-0002111323113100-1022023213020212-3231321131131320-2222210011100203-2313232332132003-2111232003313111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2112210201211101-0212311022213221-1022100111213003-0100111233033032-2110021210200322-3102121001111103-1331231323211323-2223320033022133)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-2013323000223021-2120210302002010-1033120102210120-0213011210100302-0120313323210331-2323200233033110-0332303002331332-2032332033101221"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2321211201211121-1300002233110111-0133003100120232-2013322021131021-3100003220133022-0012323100223320-1132221333300203-0113330112023133"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets`

<a id="canonical-1031113320111020-2022132222013330-3313011013300330-1321022310302211-2102212230120122-3311201202330231-3010220222231233-2322100200132321"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets.kind` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2202220001130022-0202022200100031-0330312213030123-1213221231311033-1233001223133302-3330100323121131-1332120211212320-1232223103200212"></a>

<a id="canonical-0323011220101201-2112220022310331-3302030110032002-2002132123233002-0233310301113111-3010133021232011-2332220213011233-3121021331210132"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1122231112133130-1003001300221001-2030232220002331-3022221121332101-2010211020000313-2132230300310022-0233230310113333-0010223020330003"></a>

<a id="canonical-3112010330210302-0313221023132223-0130122033101200-1023202120002003-3001200102023302-1020001013312130-2212130000210300-0321030333030230"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2032322312011002-1201210300012222-2302230020210011-3033313312000131-0010100331233002-1001003331203032-0000020000202013-0033111002122122"></a>

<a id="canonical-3221011223013323-1021220001211231-0102301031002233-0210032201211301-0331110100310321-1312132311301131-3321222211310013-0120212310310121"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1102001221022131-0132213131000330-2213121230102311-3003220222030020-2122112010010301-3101211133133231-0230303101321030-2211232002330200"></a>

<a id="canonical-0103023121313130-1013013023203100-1223303103303020-1102320332313233-1310110232112222-2111202310232122-0212221200220132-1032003323001130"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets.uid` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1200113333332102-3331022332002031-1120103010102301-2121331130312211-2212212010121333-2200210300321022-1003200210000330-3113200103222123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list

<a id="canonical-0202312213312000-3203203322023010-0011230313132331-3112023010323212-2113132222003010-1320121310320300-1030031333320102-1303123120302212"></a>

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

<a id="canonical-1103313100230223-3102111002100312-1310003033221001-1233101010022222-0210223103322010-0000303200201320-2111113321233321-0321231112332023"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list`

<a id="canonical-2202103202101002-3320000021333022-2213333103201221-3121120321212102-2303312203113213-3000322320103323-3231223121021133-3020303201332212"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list.invert_match` property

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

<a id="canonical-2213330313001113-0111112013131331-1111020311100030-0031030013011202-3321333333210302-2220222121121211-2330122131303212-0111201011311203"></a>

<a id="canonical-2330222020103032-0002120022221100-0312033323032200-2320321010120221-3223322121033301-1330031313003222-2111220233112220-1323121013313222"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list.ip_prefixes` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3312120300100220-2330131103032312-0202010330213232-3020322331303310-0212132301011221-1033003031230310-3321101123123011-1122312120111220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list

<a id="canonical-0302332011311321-0221033331302301-0221021130101110-0033201112323320-3023013313331103-0333333322032211-3022120332322131-2232331331212000"></a>

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

<a id="canonical-2000223310121233-1213232012311322-1101323223120323-2233310223102301-2222002313123210-0122133303312030-3320232102201202-1231331103201311"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list`

<a id="canonical-1133220333300300-1012001201111020-2000203021233221-3313322102331220-0221021303111203-2200331323010032-0201132211231201-2230121010020211"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3031203222211302-2303123322301120-3302130021100031-1000101233131220-3101201020012210-1301302133100302-2200220030010101-2113133013012220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-0031120021232233-0233123101022221-2112113211301033-3201302022332302-1121122132102213-0113031020213320-2001333022103103-0303230330013103"></a>

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

<a id="canonical-1330210233110210-1330020321223101-0100100312100112-1230203030022200-1212110001100133-3300132320133032-1120022033011222-2131312311213222"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-0320213132131201-3131223031311303-3112212312100200-3122111213112330-3231331003231131-2102213231132323-3002000223020203-0011230233121120"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher.classes` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1021101222131230-3302130312002103-0332220303232302-3000331102323210-1203301333212300-2012121110002130-1030023011120132-1230323230100013"></a>

<a id="canonical-1301321001010223-3030211210030312-3220321322310322-1102110231123023-0020302202303020-3120113011133310-0011322000030012-0232023330021000"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3022201232220002-1131320312131203-3022011332132331-0003213332110130-2121230200223320-1033211003011222-0100300132220302-3201333201323100"></a>

<a id="canonical-1221320033033021-0331210312201213-0300313033223323-3130330231201321-2030102320133210-0332022321103103-0330201210333023-0320131311121001"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher.excluded_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher

<a id="canonical-1332010230112030-1013220102023001-2131300200110331-2020002313312011-1022021101231311-1232101230032222-0203330321332031-1331230121003211"></a>

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

<a id="canonical-0203232312002020-0320310221331103-1010212312221211-0313130112222123-1110133222003130-2120103231333123-2311010002200133-0202110001023112"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher`

- [cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-007.md#canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311): complete subsection reference.

- [jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-007.md#canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011): complete subsection reference.

<a id="canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers

<a id="canonical-3131223312031330-2131301312303203-2103203112011103-0103130213121212-3020203110122103-2210323330120212-0023211113003303-0201233200102301"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0022221320003013-2311013202102213-1230300000220033-2101213022033001-0211220233030103-2302002220312123-1300221123010313-3213122101100020"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers`

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-0332001211220202-2222302130003203-0200000303223110-2320212201333033-2003221000321000-1111332320123003-2223120102320032-2021301300331321): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-007.md#canonical-1013311121031001-0323310112313032-0200311210032213-2003222203010331-2122132302232200-3003203030221220-2310303333330010-2333120002303232): complete subsection reference.

<a id="canonical-0330123202222313-1212123103203130-3032023303022212-2101001233132022-2201130322210202-1012311100003211-0112211011310330-0320121322230202"></a>

<a id="canonical-2122230213320030-3002112313031232-1231110332211330-0010021213211301-1112313012122330-0210032020011320-2033203102312301-1003112003222012"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.invert_matcher` property

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

- [item](resources--http_loadbalancer--reference--group-007.md#canonical-0310011312213322-1120200011320002-1111312033130310-3102020121101100-3002323002130121-1130010313020013-3130121101032131-0020313110132231): complete subsection reference.

<a id="canonical-3022320002033011-1210210311211011-0220012000200201-1011302021021110-0231202230001111-2133313211011232-1213222033020031-0003331330010222"></a>

<a id="canonical-1133110203110023-1002333300222303-0000103323103321-2303133120102200-0032020333330300-2013111120331003-2123302032030133-0210112110223230"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0332001211220202-2222302130003203-0200000303223110-2320212201333033-2003221000321000-1111332320123003-2223120102320032-2021301300331321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-2031121112020331-3302223131021322-0023321220220321-2100131103122313-3013232300101001-3122332122022333-1303103223202323-1113221312333232"></a>

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

<a id="canonical-1013311121031001-0323310112313032-0200311210032213-2003222203010331-2122132302232200-3003203030221220-2310303333330010-2333120002303232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-3101031320010310-3103310000233320-3120022022311012-0033322331002211-0203230132132112-2013333021123021-3123203000102100-1112000231102220"></a>

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

<a id="canonical-0310011312213322-1120200011320002-1111312033130310-3102020121101100-3002323002130121-1130010313020013-3130121101032131-0020313110132231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item

<a id="canonical-0322112323232031-0311132313001003-3031222022213013-0003021002023003-1002211230002032-0210320122023312-0000133323000221-3213200012022310"></a>

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

<a id="canonical-1233101322133310-2201300001132012-0322100202112002-2211202303303121-2102033002101013-1023203212211320-1013122023221032-0002203211232333"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item`

<a id="canonical-0211131030230313-2002320012220223-0130113020030010-1131003320022312-2123221021030233-2201212122331212-0232013302132022-0311202303102220"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item.exact_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2132310332120021-1202210100203312-0330302021221322-2213013303312323-1313012130213022-1333122012213122-2103023023313003-1123101311123303"></a>

<a id="canonical-3301211220230310-2312302100001333-2201233122310231-0331232331211030-0002032002312013-3211220030322233-2303303032010131-0301133302112202"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item.regex_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3023203023101113-3133111313211022-1201101232023301-2321233120232012-1113010020233013-2012310001020120-3012120021033032-2031001221300332"></a>

<a id="canonical-0232123220112311-2202331321332013-3133200130322200-3210012103310131-0231212200202302-0033001200311101-3030233011333110-3211001010013303"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item.transformers` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers

<a id="canonical-0100022210022131-3130131232331323-2322121320013101-0303232103132330-3102201322120011-2303302320202210-3221223231202122-1230303101203233"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3101132012220211-2101003013321211-1202230210012120-1011131013103322-1220331000200311-2021203330003232-3030120003011132-1322121030103121"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers`

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-1122130132031003-3223101202211302-2201131001300001-1122323213110303-0030221023012002-3321022031222123-1213320211233011-1013123122303033): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-007.md#canonical-0320220010023331-2302202111322120-0120310011031302-0232113023120201-3001011012322033-0301313233320222-3103310230023331-0323010203001222): complete subsection reference.

<a id="canonical-0212000221221212-3311231333313022-3123321331310120-3203013101333330-2020123320123010-3001223333222022-0200010220010121-3211222000003312"></a>

<a id="canonical-2100211323031123-3013100003111123-3000123033331033-1010202102133310-2021202200113231-1022013133203231-0321233220023320-2211300332121201"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.invert_matcher` property

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

- [item](resources--http_loadbalancer--reference--group-007.md#canonical-2312310310233132-3312120113120301-2102131130101312-0113220232131123-0233110022113003-2010102233201332-3031202331220220-3013210133102313): complete subsection reference.

<a id="canonical-0121100023210323-2122023002111211-0103102233031111-0213000222000332-0123232333031111-0022222223023222-3023010201032102-0301102000313003"></a>

<a id="canonical-2130223112330010-2001132020200133-1113120030101011-3213301301300300-0301211032032330-0210213213122303-3032202312130300-3331302303113212"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1122130132031003-3223101202211302-2201131001300001-1122323213110303-0030221023012002-3321022031222123-1213320211233011-1013123122303033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present

<a id="canonical-2110023000333300-2111303303020220-2012003311332112-0100323131021013-1233303112002022-0011312123233222-1312021310221330-1030302232110312"></a>

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

<a id="canonical-0320220010023331-2302202111322120-0120310011031302-0232113023120201-3001011012322033-0301313233320222-3103310230023331-0323010203001222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present

<a id="canonical-0230101211001311-1020011021210311-1201223130201220-0212020100010331-1312212131210120-1231233200303203-3322032220330012-0231101212023210"></a>

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

<a id="canonical-2312310310233132-3312120113120301-2102131130101312-0113220232131123-0233110022113003-2010102233201332-3031202331220220-3013210133102313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item

<a id="canonical-2232311113030211-2030310331202002-0203131102122220-2030002002320103-3021122002102223-1301202233212122-1201132221203200-2023020310223133"></a>

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

<a id="canonical-3301120021012012-1020313000003012-1123203210313221-1322230031210222-1111012123310020-3223010332201220-0212332222112100-3133231233032301"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item`

<a id="canonical-2101222210121330-3321320211222010-0301012111212121-0323102121212032-0100032232223200-3312212032310210-3303013310213121-0030021301321013"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item.exact_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3131002031111121-3023100021131011-3012133331202130-2210211001202010-3311031111123023-1303022333203021-2230222003313201-2031323012323120"></a>

<a id="canonical-2211211310332120-2321110301321203-1311230133021333-3002130210210013-1023321310310103-2113213220210230-0200211103000113-1110013101022030"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item.regex_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0022220002020230-1003010030131333-3303110321122313-2213133011022103-2003221333303101-2020202103302021-0022330130301031-1203130102230330"></a>

<a id="canonical-3102333232120331-1231301131302132-0123331230102323-3223331332323213-2211310203110102-2211220023323210-1301131123031103-1001111311210032"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item.transformers` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims

<a id="canonical-2131222033203211-3130123231030020-0222211332333330-0010122200133101-1331332022230023-2333001203131011-3033303223200320-1110210131130032"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2330032300200001-2002231213322231-0030021030102230-2112000130023222-0013100201203103-1131221031131203-0130212103111212-1033111102230313"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims`

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-2030312121020031-0233023120302331-1311101330233222-2010021032222120-2101010233312112-3122310011100011-0101001332310321-2111122212212302): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-007.md#canonical-3310131023313222-0323311302301330-0330121123220032-1103002110303210-1132022103030201-1222113010022210-2003300020022001-0030231132332212): complete subsection reference.

<a id="canonical-2033320033213331-2132313032212122-0011231123030200-3122010210022023-2221203322113032-0211322231302033-0132022102300020-2133133311232320"></a>

<a id="canonical-0322120111321202-2121302011200123-0123131332021231-1200233133101300-0130011111110222-1102113021300133-0033302031012311-1112303322220013"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.invert_matcher` property

Type: `"bool"`. Optional.

Invert Matcher. Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--http_loadbalancer--reference--group-007.md#canonical-0312023210333011-3220122310300130-1332132131103031-3011312321333103-3130103313320121-3331322220220322-3332331211132303-1113203221103111): complete subsection reference.

<a id="canonical-2020323023320221-1300312210331021-2122123102102202-3132220030330233-1232223303211210-3223000310130101-0030201010223033-2320313003112130"></a>

<a id="canonical-1031212221000130-0310123030200210-2333103331002002-3210010222021202-2123023023022213-1010221121111033-3302012320133131-0322020132323300"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.name` property

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2030312121020031-0233023120302331-1311101330233222-2010021032222120-2101010233312112-3122310011100011-0101001332310321-2111122212212302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-0201321003202122-3122222320310313-1111310222320002-0031310012211013-1013321002032321-1122222133300300-0001113130031013-3011120231100022"></a>

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

<a id="canonical-3310131023313222-0323311302301330-0330121123220032-1103002110303210-1132022103030201-1222113010022210-2003300020022001-0030231132332212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present

<a id="canonical-2312301311303133-2211312233203211-3010313022023001-3121302021100211-2323033122003020-2012101303322103-2320011133203121-0100102020103101"></a>

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

<a id="canonical-0312023210333011-3220122310300130-1332132131103031-3011312321333103-3130103313320121-3331322220220322-3332331211132303-1113203221103111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item

<a id="canonical-3121021123020123-3210011130010302-1122120113033220-1312301031022020-2332111201220030-2121223000331212-1012233213313300-0132121301130312"></a>

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

<a id="canonical-0033313212210320-3203333020320132-2030010122300033-1213110333302203-0023203122322321-0031130001302001-2010212201310330-2032031300003000"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item`

<a id="canonical-2013301233100303-0321202212312013-1202223313200212-1222111221231233-3310110013213021-3331213101013320-0021311120332032-1133303022312011"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item.exact_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2212032231020022-2303223201233120-1031310000201211-2213322002033031-2022113313123000-1003213232222112-3010133311023011-1011203311131303"></a>

<a id="canonical-2230033021230313-0220033013321113-1030230031020103-3013020331031013-1312301130332113-3230313203221012-2020032303210032-0003331113313121"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item.regex_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3010001313020230-0221212322202202-3010122100322121-1023112131312231-2311003203201013-2101231232310303-3000102301030302-3330021103030330"></a>

<a id="canonical-0332010231333211-0222222232323201-0230211110323133-1033203020300233-3033330021321221-0001300001230310-1011101031022131-0121113021123213"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item.transformers` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params

<a id="canonical-0230100302023013-2001113023300110-1232120030000011-2001231222132202-0233333320233231-2023000211212022-0333210033200211-1221213002030203"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310010123120233-3023330101330132-0111331213223032-3311202203113023-2000300003202100-2231031203103010-3330313123203300-2311132311210120"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params`

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-2133331220212010-1131020000300212-0000202201222330-2102033122223221-3002330212312012-0110323033001201-2323121022032323-2121221021033131): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-008.md#canonical-1030312120322111-0033230321102320-1211023133031131-2303322232000221-3210112012331103-2210022021220122-0302033212333022-2010323221130123): complete subsection reference.

<a id="canonical-3203003323010322-1020002033023023-2321322202113113-3132100022200211-1321211123002000-3101111110133110-1111222031022300-1211310133320010"></a>

<a id="canonical-1332231333212033-2001302303103330-2020030131130021-1000322021113331-0213223123013001-0303311133113323-0203033000112011-2330033110300222"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.invert_matcher` property

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--http_loadbalancer--reference--group-008.md#canonical-3202300301312031-0020030211222211-0313121300121311-2031212131321320-1102131113201100-3312312301133303-3130013111111322-1113322333210333): complete subsection reference.

<a id="canonical-2032101200220213-2122331222030110-1123103312122312-1333223010320302-3211102230312002-2212102212312202-3211033301232031-2133220300000023"></a>

<a id="canonical-0121300013213013-0011113230031110-0100301211121223-0020323232002300-0320123003002232-1211213120331231-2002331012023103-2220033021021313"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.key` property

Type: `"string"`. Optional.

A case-sensitive HTTP query parameter name.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2133331220212010-1131020000300212-0000202201222330-2102033122223221-3002330212312012-0110323033001201-2323121022032323-2121221021033131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-007.md#canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present

<a id="canonical-3302111023300033-0212112200122102-2331123213320131-1211020210333302-1022122103313112-3230133020301312-2130201310233312-2002010033201122"></a>

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
