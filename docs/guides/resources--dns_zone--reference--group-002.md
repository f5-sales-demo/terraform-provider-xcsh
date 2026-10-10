---
page_title: "xcsh_dns_zone reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone reference."
---

# xcsh_dns_zone reference

<a id="canonical-0310232010230002-2032333311200131-2210000030122303-1120200331023020-3101110012303221-2302222310112000-2232130122311323-0203010122113311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.eui64_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.eui64_record

<a id="canonical-3223320120003011-3103320130031032-3311330233013221-1313223103010030-1113212203021121-0022012230030112-1113312132032223-1322223101102132"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for eui64 record.

Additional upstream details:

DNS EUI64 Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("value")}
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
eui64_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232013220223120-3033303010111331-2011321232202322-2101320330111333-3100121112320113-1120002100301332-2022033302031032-2301032111023122"></a>

### Direct properties for `primary.default_rr_set_group.eui64_record`

<a id="canonical-1110123123300000-2102133000131311-2223001311203032-1012132033120023-1113322230311331-3332032332102122-1302103333220210-3311012123110000"></a>

#### `primary.default_rr_set_group.eui64_record.name` property

Type: `"string"`. Optional.

EUI64 Record name, please provide only the specific subdomain or record name without the base
domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-3030022232113003-1122010013003310-0022001100231103-3021222011200202-0130033322011032-2331223311112000-3231101330333020-1213331123313213"></a>

<a id="canonical-3202011330031020-1121110221212320-3231201232233320-3100122220213201-2333212121223132-0031133300120022-1200320332213233-0320303111022321"></a>

#### `primary.default_rr_set_group.eui64_record.value` property

Type: `"string"`. Optional.

EUI64 Identifier. A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(23, 23),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 23,
  "minLength": 23,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 23,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 23,
    "pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-2000023011031312-2003213123011132-0300331211333233-1003131020211202-0213013000200212-3132103320112330-2130111300321023-0221122033031030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.lb_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.lb_record

<a id="canonical-0123002111321001-2130320220120210-1321310003303003-2100203031031100-2233221131213003-3032110020220011-1013001123223221-1101033233232310"></a>

Type: `"object"`. single nested block, Optional.

DNS Load Balancer Record. DNS Load Balancer Record.

Receipt-pinned upstream constraints:

```json
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
lb_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0123232021311300-1313321221012121-3132010030013312-0202033231111112-1123212221313201-1100013010301021-0000102233001222-1310200212321201"></a>

### Direct properties for `primary.default_rr_set_group.lb_record`

<a id="canonical-1202012023213130-2103102332233010-0230120330112132-1212223011111002-2022131301321010-1003300222312330-0103010213100322-1201001121033001"></a>

#### `primary.default_rr_set_group.lb_record.name` property

Type: `"string"`. Optional.

Load Balancer record name (except for SRV DNS Load balancer record) should be a simple record name
and not a subdomain of a subdomain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
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
    "maxLength": 255,
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
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

- [value](resources--dns_zone--reference--group-002.md#canonical-1130010332130022-0221033112200030-1231013030221321-1331113132211300-1132223203031221-0333312030123323-2123213012013111-0031011201002313): complete subsection reference.

<a id="canonical-1130010332130022-0221033112200030-1231013030221321-1331113132211300-1132223203031221-0333312030123323-2123213012013111-0031011201002313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.lb_record.value` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.lb_record](resources--dns_zone--reference--group-002.md#canonical-2000023011031312-2003213123011132-0300331211333233-1003131020211202-0213013000200212-3132103320112330-2130111300321023-0221122033031030)
- primary.default_rr_set_group.lb_record.value

<a id="canonical-3323100013132323-2203130020313231-2231013033222222-0303012302232332-1033302010331212-3320121211101001-0313002232032000-2111101103302032"></a>

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
value {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131001300103000-3110132102133232-0033011323023221-1130133303220211-2100033211130102-0033212130303211-1023003313322332-1320230033021313"></a>

### Direct properties for `primary.default_rr_set_group.lb_record.value`

<a id="canonical-0210222223002130-2230001023300130-0101330120013331-2003001200212201-2330301013202032-1020331323121313-2330111010203310-2330210310031302"></a>

#### `primary.default_rr_set_group.lb_record.value.name` property

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

<a id="canonical-2131202322332212-3002101101303303-0113312022321123-0022323100111111-3222131232310203-3312003313113011-2211113303301103-1310021021213013"></a>

<a id="canonical-0332330130211020-3020111212332330-3102101120210221-0030213020003032-1101211000213232-2020021032321102-3311022223220212-0102223232332212"></a>

#### `primary.default_rr_set_group.lb_record.value.namespace` property

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

<a id="canonical-2311232221230032-1222033332001211-3213122203121101-1111312321031020-1230020213322220-0203213030132232-2020220311001121-3131322201223123"></a>

<a id="canonical-1303031112223022-3102030202032131-0202332000013032-3012013032302322-1100112322002123-1230322013233211-3210223012212322-0021323110332300"></a>

#### `primary.default_rr_set_group.lb_record.value.tenant` property

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

<a id="canonical-0001331322032313-1331233300221231-2021032123331122-1313301111122120-1113123132222021-0013311333002210-3021113111120010-3121010323223333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.loc_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.loc_record

<a id="canonical-2132111113030021-3321322230122210-2330022320201032-0022111313131111-3112122321323003-3321113201120112-2331221333120033-0232020030113010"></a>

Type: `"object"`. single nested block, Optional.

DNS LOC Record. DNS LOC Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
loc_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320001301011213-3120311233302313-1201212230131332-2113103122323202-0022331022310310-2120002332112030-0131232123302001-1213313033003031"></a>

### Direct properties for `primary.default_rr_set_group.loc_record`

<a id="canonical-0012123332333200-2002113200301200-3333120212310101-1112202231311122-3313130131110103-3032330100233230-3031333110113100-2122330012023202"></a>

#### `primary.default_rr_set_group.loc_record.name` property

Type: `"string"`. Optional.

LOC Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-2112121111131001-3303232101130313-1333110202023323-2101112231121012-1103313011133322-1033130001232233-0232023330011221-1223033200111230): complete subsection reference.

<a id="canonical-2112121111131001-3303232101130313-1333110202023323-2101112231121012-1103313011133322-1033130001232233-0232023330011221-1223033200111230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.loc_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.loc_record](resources--dns_zone--reference--group-002.md#canonical-0001331322032313-1331233300221231-2021032123331122-1313301111122120-1113123132222021-0013311333002210-3021113111120010-3121010323223333)
- primary.default_rr_set_group.loc_record.values

<a id="canonical-0000303313130103-3211013021131123-2212112102013332-2313032301323331-2113023333303021-2122000331232010-1033313323030300-0231003010101302"></a>

Type: `"object"`. list nested block, Optional.

LOC Value. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("altitude",
    "latitude_degree",
    "longitude_degree")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-1031221320001033-0022320101121320-1120021020202102-3130210022123102-2031200123202123-0330031220233002-3200310113013031-3223331323233321"></a>

### Direct properties for `primary.default_rr_set_group.loc_record.values`

<a id="canonical-2131202112103221-2211310101233311-3013113321130321-3223213123102001-1223001232012103-2111100230303110-2011211320302332-0030000332201231"></a>

#### `primary.default_rr_set_group.loc_record.values.altitude` property

Type: `"number"`. Optional.

Altitude. Altitude in meters.

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
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0223310311003231-2130112230301222-1200002110121021-3023210233223033-2013112222210132-3110100112113332-0022003203112030-0013132220111321"></a>

<a id="canonical-1230300022122321-3102301121101213-0311030321110322-3310121100331213-3232302131220322-3332222312200102-2312022021100212-0321200121131000"></a>

#### `primary.default_rr_set_group.loc_record.values.horizontal_precision` property

Type: `"number"`. Optional.

Horizontal Precision. Horizontal Precision in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-3101333230031331-3203232301103020-0123322112112103-3013313202211231-2002101312323213-2033001121131010-3313131200130212-1330321222233011"></a>

<a id="canonical-1331110213012320-3132311211200202-0203220021123230-1313221302013200-3010330102001221-3201203132033130-0011133211111002-0112303031233223"></a>

#### `primary.default_rr_set_group.loc_record.values.latitude_degree` property

Type: `"number"`. Optional.

Latitude degree, an integer between 0 and 90, including 0 and 90.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 90),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 90,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3010322301212100-3221131321000232-2003311301111310-0231101220002001-1030003102211020-0301203102230320-3221023213233001-1002210230231013"></a>

<a id="canonical-3310102010131301-3210100123210211-2300323020123233-0020333230122213-1102233332212013-1213030333300131-3211303032213311-0310102013323121"></a>

#### `primary.default_rr_set_group.loc_record.values.latitude_hemisphere` property

Type: `"string"`. Optional.

\[Enum: N|S\] Latitude hemisphere can only be N or S - N: North Hemisphere - S: South Hemisphere.
Possible values are \`N\`, \`S\`. Defaults to \`N\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["N","S"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("N",
    "S"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "N",
  "enum": [
    "N",
    "S"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3003200010121011-3122020100020232-3120300303303100-1300330210031023-1101122321033011-3320131033011310-1113300233220011-1130013003311130"></a>

<a id="canonical-3031212112222331-1000003122221211-0233013213333333-2100331132222303-3030003013232132-1010023013103102-3331230330321321-2310331102133220"></a>

#### `primary.default_rr_set_group.loc_record.values.latitude_minute` property

Type: `"number"`. Optional.

Latitude minute, an integer between 0 and 59, including 0 and 59.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 59),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-1132311121131210-0032113300321100-2230002011302100-1312130031130022-0023131012023001-3011231130111301-1310210111302010-1112100100223123"></a>

<a id="canonical-0302013330223032-3001301313020103-0300212220130320-3331231313230232-2102320333220031-0021200011210132-1310011011331013-0200102321100032"></a>

#### `primary.default_rr_set_group.loc_record.values.latitude_second` property

Type: `"number"`. Optional.

Latitude second, an decimal between 0 and 59.999, including 0 and 59.999.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-2001302000003023-0110113301110220-3022310200322013-1322311000011013-2203323331133222-1332122010032033-1022302221123222-2310221233123231"></a>

<a id="canonical-2302213112033022-1103311032133303-3100120221221222-1231123100122223-0313321012031333-3321133202113230-3013111123002132-1031100100320110"></a>

#### `primary.default_rr_set_group.loc_record.values.location_diameter` property

Type: `"number"`. Optional.

Diameter of a sphere enclosing the described entity, in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-0123032002122113-1101033300112321-3102333111012213-3111120331331330-0211113310100321-1201301111002221-2113201211123032-2231310312113203"></a>

<a id="canonical-1231301212101220-0122213020231202-0220202230221302-2311221032111301-2231321132033113-0203300130113021-2131010033312210-3122213213212233"></a>

#### `primary.default_rr_set_group.loc_record.values.longitude_degree` property

Type: `"number"`. Optional.

Longitude degree, an integer between 0 and 180, including 0 and 180.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 180),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1033200313203132-3302301021222000-2012320320112210-3201220231010301-0013323120103203-1232023331020331-2301011211131303-2013000022210032"></a>

<a id="canonical-0020003313201313-3302002211211230-0102302113022213-0201200322001312-1301201032032323-3130231121103223-2112000222101331-1212301211202102"></a>

#### `primary.default_rr_set_group.loc_record.values.longitude_hemisphere` property

Type: `"string"`. Optional.

\[Enum: E|W\] Longitude hemisphere can only be E or W - E: East Hemisphere - W: West Hemisphere.
Possible values are \`E\`, \`W\`. Defaults to \`E\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["E","W"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("E",
    "W"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "E",
  "enum": [
    "E",
    "W"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3101310132203301-1232323301110103-1023221213131013-0330301121231202-1033222223012023-0221300103211313-3200213130320330-0201323023322231"></a>

<a id="canonical-0300120212301023-1231000310101032-1010022102221132-2121132032013303-2032100033212133-3321323033330032-0123112100203312-2303221311321212"></a>

#### `primary.default_rr_set_group.loc_record.values.longitude_minute` property

Type: `"number"`. Optional.

Longitude minute, an integer between 0 and 59, including 0 and 59.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 59),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-2023230311100231-3322110003113122-1033331333231311-0332023131230210-2233133003101311-2033101222223113-3110013203002202-0322120201123130"></a>

<a id="canonical-3022032313313213-2301203201001201-2032020131122011-1222202311323321-3221000033333313-1212031113112101-3212102002003330-3230100200121201"></a>

#### `primary.default_rr_set_group.loc_record.values.longitude_second` property

Type: `"number"`. Optional.

Longitude second, an decimal between 0 and 59.999, including 0 and 59.999.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-2303210211122011-3132033003000301-0102311002020001-0200122221213003-1021221132220102-1322301213122313-2220221303100310-1220013111321300"></a>

<a id="canonical-1103102032011100-2203022122200120-2313030333000000-2303220033200113-1012303132223100-2200021020102012-3101302201203222-3000100233311121"></a>

#### `primary.default_rr_set_group.loc_record.values.vertical_precision` property

Type: `"number"`. Optional.

Vertical Precision. Vertical Precision in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-1002110103301213-0230131110201331-3223012003111321-0300033220002210-3000130132020202-0011011200330101-0120220030322210-3213201300203312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.mx_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.mx_record

<a id="canonical-3031022012121010-1033000003212220-1333213133232001-2301203330223011-1202332101303120-0230200230301011-1311002130031131-0113212112300021"></a>

Type: `"object"`. single nested block, Optional.

DNSMXResourceRecord.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
mx_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230032010113001-1102301231000013-1300013210320113-2012002111202210-3202022210230012-2100321213231313-2103300321333231-2130222331033330"></a>

### Direct properties for `primary.default_rr_set_group.mx_record`

<a id="canonical-1321002202112310-2312002321302110-1130310200333323-3211330321331112-1031301011303103-1103133023220232-0302202132003202-0230330012010013"></a>

#### `primary.default_rr_set_group.mx_record.name` property

Type: `"string"`. Optional.

MX Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-2132212202122320-2030321100131003-0133003301133102-0220221112310121-3122231332022001-3032232202002320-0113111230100302-3313121002202311): complete subsection reference.

<a id="canonical-2132212202122320-2030321100131003-0133003301133102-0220221112310121-3122231332022001-3032232202002320-0113111230100302-3313121002202311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.mx_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.mx_record](resources--dns_zone--reference--group-002.md#canonical-1002110103301213-0230131110201331-3223012003111321-0300033220002210-3000130132020202-0011011200330101-0120220030322210-3213201300203312)
- primary.default_rr_set_group.mx_record.values

<a id="canonical-0003030310133033-1301312201120100-1102301001211011-1101203121013200-3232121203223313-3200201303223333-3212301101202311-2110102332330230"></a>

Type: `"object"`. list nested block, Optional.

MX Record Value. Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-1311000231202133-1320022022133002-2101231302110033-2000030201313130-3132303331202021-1000333333301333-1110012023020231-0013323102030301"></a>

### Direct properties for `primary.default_rr_set_group.mx_record.values`

<a id="canonical-3212223332222301-2111210213031321-1302011002231203-2203220223311120-2233203000210201-2221000222013101-0013120102130102-2020000013302301"></a>

#### `primary.default_rr_set_group.mx_record.values.domain` property

Type: `"string"`. Optional.

Mail exchanger domain name, please provide the full hostname, for example: mail.example.com.

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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-0301021110310111-1322301311102201-1231000121230232-1333010102002301-2202120233313002-3303212202312212-3030323011221232-2032301101201300"></a>

<a id="canonical-3221112310133230-2330000110223312-2010222001102211-1003330123221003-2211222023100020-2203130103323023-2231331300100021-2332111203102323"></a>

#### `primary.default_rr_set_group.mx_record.values.priority` property

Type: `"number"`. Optional.

Priority. Mail exchanger priority code.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0221201201321311-3033210312022203-0010211233110120-0301032022010211-2322320001313111-2002020120203123-1333000323301120-3330303320333302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.naptr_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.naptr_record

<a id="canonical-3302201020210130-3022011223202021-0033303101220201-0002010112333122-0031033012232221-3313032200103232-3102321021201211-1101130203111310"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for naptr record.

Additional upstream details:

DNS NAPTR Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
naptr_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301313031010320-3100310100312111-3311113222131313-3030231202100013-2123302102211113-2033131303332012-3302201111322200-0031123233203133"></a>

### Direct properties for `primary.default_rr_set_group.naptr_record`

<a id="canonical-3303200222203311-3322303122022321-0310111001311232-1311233100222002-2000210202131011-3322313011110023-3133012122231120-2120030332332222"></a>

#### `primary.default_rr_set_group.naptr_record.name` property

Type: `"string"`. Optional.

NAPTR Record name, please provide only the specific subdomain or record name without the base
domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-1111030123330233-3012003313010210-1310201033210001-1202023220222030-2122100010113323-2232130003103221-0031211122130333-0020302323221311): complete subsection reference.

<a id="canonical-1111030123330233-3012003313010210-1310201033210001-1202023220222030-2122100010113323-2232130003103221-0031211122130333-0020302323221311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.naptr_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.naptr_record](resources--dns_zone--reference--group-002.md#canonical-0221201201321311-3033210312022203-0010211233110120-0301032022010211-2322320001313111-2002020120203123-1333000323301120-3330303320333302)
- primary.default_rr_set_group.naptr_record.values

<a id="canonical-0012133211132101-0222112211233003-0202312222331133-1111020032121221-1311303333202031-3112101213011201-2030200321213312-3020100000132311"></a>

Type: `"object"`. list nested block, Optional.

NAPTR Value. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("flags",
    "order",
    "preference")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010233333033332-0310230310301002-1303002232123202-1323000330111201-1232000023012013-3211200222201112-1101031130321020-2133333210012312"></a>

### Direct properties for `primary.default_rr_set_group.naptr_record.values`

<a id="canonical-1033021233210122-3200232001131212-0130330200111101-2002003333123000-3310031000110031-3230321033232230-0130211012022123-3303002312103123"></a>

#### `primary.default_rr_set_group.naptr_record.values.flags` property

Type: `"string"`. Optional.

Flag to control aspects of the rewriting and interpretation of the fields in the record. At this
time only four flags, S/A/U/P, are defined.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  }
}
```

<a id="canonical-0003322302231222-1002002201313010-3201230211230220-1223201122320320-3311330200200203-0000121233330012-0311232032101331-3123313013332232"></a>

<a id="canonical-2202132021311320-1111321010323002-0021103113302322-0131320330002211-2202222122313221-2330132002332013-3113221202003000-1123003012010022"></a>

#### `primary.default_rr_set_group.naptr_record.values.order` property

Type: `"number"`. Optional.

Order in which the NAPTR records must be processed. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0203233120120121-1223303131330102-1223232310132002-3301221113103100-3312233033302311-0011020332332101-0011323233301132-3221221100312303"></a>

<a id="canonical-1203302011203323-3102030021212031-0303201030202331-3102300130222203-0132130220023103-1311300312010132-0132013023303321-2312220030230200"></a>

#### `primary.default_rr_set_group.naptr_record.values.preference` property

Type: `"number"`. Optional.

Preference when records have the same order. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1133101022121030-3222111130232323-3003312223130130-3000111311111320-0102021122022012-3123102310331021-2331310210210022-1100132112013120"></a>

<a id="canonical-3322230013113012-3033120023022231-2002220300133201-1313210201213320-1132311322201012-0203100102031131-2331132011203222-1332102010300320"></a>

#### `primary.default_rr_set_group.naptr_record.values.regexp` property

Type: `"string"`. Optional.

Regular expression to construct the next domain name to lookup.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
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
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-1303002220120012-3033311223202121-2210031032110320-0220031103301301-1110120230031311-0310212111211132-0101022331313302-1221303110022131"></a>

<a id="canonical-3202312033223000-1122023220102013-0301201232001320-1023330000031202-1222220133013103-3033221112233303-2201222013120221-0123120120031020"></a>

#### `primary.default_rr_set_group.naptr_record.values.replacement` property

Type: `"string"`. Optional.

The next NAME to query for NAPTR, SRV, or address records depending on the value of the flags field.

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

<a id="canonical-0100322000333330-1013302333103132-2202202223102312-1101112312031213-3310021222110212-1220033222200200-3132030223101200-0120130112301221"></a>

<a id="canonical-2200302133200300-2212002213023203-3130033132101333-3210210211101233-2221102001220120-3230031022203320-1131212310102003-3003012200333011"></a>

#### `primary.default_rr_set_group.naptr_record.values.service` property

Type: `"string"`. Optional.

Specifies the service(s) available down this rewrite path.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  }
}
```

<a id="canonical-1320012000010221-1013101202011330-0222031102033131-0333201222231300-2310103201032220-3222020113330211-0230001303312231-1302222302130302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.ns_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.ns_record

<a id="canonical-3133323100030001-3323320220032012-1023332123310223-2133233223111122-2110331310221011-3223131111103203-2022210101000332-0012033013311031"></a>

Type: `"object"`. single nested block, Optional.

DNSNSResourceRecord.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
ns_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232133033111132-1031011000322101-2233021312212003-0110021101033033-0322233313312030-0310231011123013-3003213001213121-1011132001203321"></a>

### Direct properties for `primary.default_rr_set_group.ns_record`

<a id="canonical-2033003111331232-2100321300133202-0112031132303032-0321230200103202-3333010122311223-3320231030313320-2300021021310111-2111212112023110"></a>

#### `primary.default_rr_set_group.ns_record.name` property

Type: `"string"`. Optional.

NS Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-3113333313223202-3322202032213010-1202233000332200-3222103322113102-2222133303301233-0320203210302131-3231111130301002-2312123022210102"></a>

<a id="canonical-1111023131013111-3031223021023031-1010120233000133-3320202231012111-1020001303030020-0002332311101032-2302223222312020-0201222223131312"></a>

#### `primary.default_rr_set_group.ns_record.values` property

Type: `["list", "string"]`. Optional.

Name Servers. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3030100213211033-2113011011200322-2233222032221031-1220020031011201-1132132112023130-3211231011202332-2123323330002213-2200232102012230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.ptr_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.ptr_record

<a id="canonical-2000130221010012-1133232130332203-2023123012103123-1213333232011033-3021323331322132-0220320211113333-2311120321002201-0202020210201231"></a>

Type: `"object"`. single nested block, Optional.

DNSPTRResourceRecord.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
ptr_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2332311012223110-2123121003300022-2223001210303000-2331010212303231-0231001223130331-3231021110031101-0121100010232013-3330133020310331"></a>

### Direct properties for `primary.default_rr_set_group.ptr_record`

<a id="canonical-1102100003332030-3003322001321331-2232010102001313-3223110102112111-1320323213032321-1132212230311002-3130231300002110-0132033303303000"></a>

#### `primary.default_rr_set_group.ptr_record.name` property

Type: `"string"`. Optional.

PTR Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-0212323031323213-3310023311233130-3013102103220130-2030220310203210-3221031003330200-3230100102132031-0110120031033122-3022123022013230"></a>

<a id="canonical-0030300311002122-0032302212233012-0233331213313213-1332303032231303-0310030302101022-0212003210321031-1102322103122003-0202310111302031"></a>

#### `primary.default_rr_set_group.ptr_record.values` property

Type: `["list", "string"]`. Optional.

Domain Name. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0113001230120133-2303121012102031-1200311122002232-3032331232033323-1200010123122221-1212320123121123-0013213201323320-3313313211120233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.srv_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.srv_record

<a id="canonical-3232132023001031-0110132101031030-2312320323010000-0302113021120313-2200303232110120-0032332303330320-2200231003220012-3002111130330010"></a>

Type: `"object"`. single nested block, Optional.

DNSSRVResourceRecord.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name",
    "values")}
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
srv_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0033203220331123-2013010000311112-1102323331331032-2121203130200131-2010002203123311-1113001113220021-3010303232223223-3030210302012210"></a>

### Direct properties for `primary.default_rr_set_group.srv_record`

<a id="canonical-2021022302320110-0232311121203113-0131321030222332-3321300123301022-2202212013111310-3002012230220200-3010121102310002-1211021110021100"></a>

#### `primary.default_rr_set_group.srv_record.name` property

Type: `"string"`. Optional.

SRV Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-0220200230321313-1301130010021321-0111211323133021-2233302021322332-0311101110010303-1222201032221100-1000122112231033-2103231133330232): complete subsection reference.

<a id="canonical-0220200230321313-1301130010021321-0111211323133021-2233302021322332-0311101110010303-1222201032221100-1000122112231033-2103231133330232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.srv_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.srv_record](resources--dns_zone--reference--group-002.md#canonical-0113001230120133-2303121012102031-1200311122002232-3032331232033323-1200010123122221-1212320123121123-0013213201323320-3313313211120233)
- primary.default_rr_set_group.srv_record.values

<a id="canonical-1233110110121133-0211332212112022-0220333223110020-0222101213000123-1223321332033030-1212201202333331-0222210032003303-0133220311010220"></a>

Type: `"object"`. list nested block, Optional.

SRV Value. Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-3321121013031301-3333132000111030-2230321310013030-1221322211321003-2313320120203002-3030131030111033-0000231320230321-1301233113332223"></a>

### Direct properties for `primary.default_rr_set_group.srv_record.values`

<a id="canonical-3020210212331312-2113033213301003-3332030210023002-1222331100023311-3213033310033211-3020101212033020-3002301000121233-0023302121210313"></a>

#### `primary.default_rr_set_group.srv_record.values.port` property

Type: `"number"`. Optional.

Port. Port on which the service can be found.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0212122213022002-1102130330230022-2113103310311202-3022213003321111-3031122031223210-2302032331122032-0033223320030102-3021033322001002"></a>

<a id="canonical-0101331133301333-1030111000213221-0302310033011120-0020121213021113-3131032200013011-1020030231100301-0013122212201002-1102021131101311"></a>

#### `primary.default_rr_set_group.srv_record.values.priority` property

Type: `"number"`. Optional.

Priority of the target. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3223322100303321-0031100122012000-0231120301220323-3101323020211220-3321023031332030-2000233003301022-1012301232212321-0223131202102230"></a>

<a id="canonical-1330303021221322-3332031022230010-0132130233221123-2300133022023330-1210132330020013-2232012132031302-2323321202131310-3220233323030013"></a>

#### `primary.default_rr_set_group.srv_record.values.target` property

Type: `"string"`. Optional.

Hostname of the machine providing the service.

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
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  }
}
```

<a id="canonical-0223320020032303-2313320021010002-1100112012113203-2222112110223313-0220113331223131-1330112001302113-1102223203003133-2303132101301102"></a>

<a id="canonical-2212002001232202-1210001330202100-3210322121023021-2313120222113212-0323030103020231-3012113111310020-0323212331312023-3130022202320332"></a>

#### `primary.default_rr_set_group.srv_record.values.weight` property

Type: `"number"`. Optional.

Weight of the target. A higher number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0002133302013323-3312120311230021-3113212100312321-1000222020201120-3100200121220311-1202130100221130-3332302111201101-1102012023123221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.sshfp_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.sshfp_record

<a id="canonical-0033310300233313-3001201311021021-3123210133213011-0313232002021102-0111202110303313-0230133031020323-2110230212303300-3320313013021301"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sshfp record.

Additional upstream details:

DNS SSHFP Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
sshfp_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-1012022302233122-1112222301103030-0133133123002302-0133323030300321-3213200311001023-2100223021010122-1100311332223311-2300330030323130"></a>

### Direct properties for `primary.default_rr_set_group.sshfp_record`

<a id="canonical-1312323000120211-3122313122033001-2301100232310012-0010001313331121-1103323211133023-1111210113022110-0113201100200021-1323210100121211"></a>

#### `primary.default_rr_set_group.sshfp_record.name` property

Type: `"string"`. Optional.

SSHFP Record name, please provide only the specific subdomain or record name without the base
domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-1101120321300303-2320122100231000-1311002212133031-0121103313331023-3103122323300332-0202111210102033-0301330202021300-0012010120131212): complete subsection reference.

<a id="canonical-1101120321300303-2320122100231000-1311002212133031-0121103313331023-3103122323300332-0202111210102033-0301330202021300-0012010120131212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.sshfp_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.sshfp_record](resources--dns_zone--reference--group-002.md#canonical-0002133302013323-3312120311230021-3113212100312321-1000222020201120-3100200121220311-1202130100221130-3332302111201101-1102012023123221)
- primary.default_rr_set_group.sshfp_record.values

<a id="canonical-0323200310303211-3133202223300230-3333231133230112-0322212032230011-1120330010031221-3103010302313102-3310223021330202-0222321323003333"></a>

Type: `"object"`. list nested block, Optional.

SSHFP Value. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("sha1_fingerprint",
    "sha256_fingerprint")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222222210111320-0012013103231030-1112001111121300-2000023322221032-0102012210132233-2200013002032013-3202233331113200-3222232331203213"></a>

### Direct properties for `primary.default_rr_set_group.sshfp_record.values`

<a id="canonical-0100312300001122-1100322030101033-3013022100332131-0233211220102221-3132132323131321-3013000222230301-1021002323200020-1320101113003203"></a>

#### `primary.default_rr_set_group.sshfp_record.values.algorithm` property

Type: `"string"`. Optional.

\[Enum: UNSPECIFIEDALGORITHM|RSA|DSA|ECDSA|Ed25519|Ed448\] SSHFP algorithm value must be compatible
with the specified algorithm. - UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM - RSA: RSA - DSA: DSA -
ECDSA: ECDSA - Ed25519: Ed25519 - Ed448: Ed448. Possible values are \`UNSPECIFIEDALGORITHM\`,
\`RSA\`, \`DSA\`, \`ECDSA\`, \`Ed25519\`, \`Ed448\`. Defaults to \`UNSPECIFIEDALGORITHM\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DSA","ECDSA","Ed25519","Ed448","RSA","UNSPECIFIEDALGORITHM"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("UNSPECIFIEDALGORITHM",
    "RSA",
    "DSA",
    "ECDSA",
    "Ed25519",
    "Ed448"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIEDALGORITHM",
  "enum": [
    "UNSPECIFIEDALGORITHM",
    "RSA",
    "DSA",
    "ECDSA",
    "Ed25519",
    "Ed448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [sha1_fingerprint](resources--dns_zone--reference--group-002.md#canonical-2023231033131113-2110202123131020-1203111201302332-1313212001113301-3333113320331021-2020121001100123-0220312010123310-2000300312120110): complete subsection reference.

- [sha256_fingerprint](resources--dns_zone--reference--group-002.md#canonical-3103022332110213-3100031323212012-0221003020003232-0230303211030301-0033003030113002-0032113003232322-2323001210333103-1021113222213111): complete subsection reference.

<a id="canonical-2023231033131113-2110202123131020-1203111201302332-1313212001113301-3333113320331021-2020121001100123-0220312010123310-2000300312120110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.sshfp_record](resources--dns_zone--reference--group-002.md#canonical-0002133302013323-3312120311230021-3113212100312321-1000222020201120-3100200121220311-1202130100221130-3332302111201101-1102012023123221)
- [primary.default_rr_set_group.sshfp_record.values](resources--dns_zone--reference--group-002.md#canonical-1101120321300303-2320122100231000-1311002212133031-0121103313331023-3103122323300332-0202111210102033-0301330202021300-0012010120131212)
- primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint

<a id="canonical-3300230102310032-0013311210030213-3111312011020222-2122123002032112-2221302013330301-3202101323320302-3302012213021222-2320211131333010"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha1 fingerprint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("fingerprint")}
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
sha1_fingerprint {
  # Configure direct properties listed below.
}
```

<a id="canonical-1321101202102302-1303110031311001-3110010132031131-2222032110202322-3300332221320300-1332313000123122-3223321013033300-1321201321321111"></a>

### Direct properties for `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint`

<a id="canonical-3230330302020020-1213132222131323-0002010120323031-3233011323002013-2201200301032200-0121322201011222-1313221131123210-2001212211321122"></a>

#### `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint.fingerprint` property

Type: `"string"`. Optional.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(40, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 40,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-3103022332110213-3100031323212012-0221003020003232-0230303211030301-0033003030113002-0032113003232322-2323001210333103-1021113222213111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.sshfp_record](resources--dns_zone--reference--group-002.md#canonical-0002133302013323-3312120311230021-3113212100312321-1000222020201120-3100200121220311-1202130100221130-3332302111201101-1102012023123221)
- [primary.default_rr_set_group.sshfp_record.values](resources--dns_zone--reference--group-002.md#canonical-1101120321300303-2320122100231000-1311002212133031-0121103313331023-3103122323300332-0202111210102033-0301330202021300-0012010120131212)
- primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint

<a id="canonical-0123302233213302-3311232213203031-2013230021213311-3000122000201211-0020103220202111-1121230212202312-2201032200130123-3303320103313102"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha256 fingerprint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("fingerprint")}
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
sha256_fingerprint {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303030123311212-3210203101120013-2312203332122322-0333031223103011-0000222121002200-3321032300031003-3000320213133321-0013321002322213"></a>

### Direct properties for `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint`

<a id="canonical-1302201230021103-1232233201132323-1030012333213200-1133320003102230-3323030011231103-1230022311102032-2000000122202223-3302112310002200"></a>

#### `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint.fingerprint` property

Type: `"string"`. Optional.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(64, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 64,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-1110210111012102-2031202010320123-1132320121232332-2002023031230032-2022213032232233-1212213101223203-1103321022220000-0330112323013122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.tlsa_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.tlsa_record

<a id="canonical-1320110233003023-3320023020133213-1311023002130001-3013232233301000-2313223031331311-1202330011123222-3011223300111332-1022132230100103"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tlsa record.

Additional upstream details:

DNS TLSA Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
tlsa_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333113123231031-0323200223031110-3021133003001212-3211110031103220-1222301311031110-1123232202013202-2000211233203223-3020302111022103"></a>

### Direct properties for `primary.default_rr_set_group.tlsa_record`

<a id="canonical-1223313213332221-2321231333010013-2102333210131220-2031033201102231-0113101310133310-1212302101010313-2333033303031031-1302331113230213"></a>

#### `primary.default_rr_set_group.tlsa_record.name` property

Type: `"string"`. Optional.

TLSA Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-1103233213133131-2320013020321000-0101132130130100-1302121013001001-0101011131022333-1213332023301201-3301000123113300-0202231132001031): complete subsection reference.

<a id="canonical-1103233213133131-2320013020321000-0101132130130100-1302121013001001-0101011131022333-1213332023301201-3301000123113300-0202231132001031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.tlsa_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.tlsa_record](resources--dns_zone--reference--group-002.md#canonical-1110210111012102-2031202010320123-1132320121232332-2002023031230032-2022213032232233-1212213101223203-1103321022220000-0330112323013122)
- primary.default_rr_set_group.tlsa_record.values

<a id="canonical-2033223312322302-1020120221212021-1021023130230232-3102322210033211-1133210131303000-2202231002002201-2100002112220313-2113311310333233"></a>

Type: `"object"`. list nested block, Optional.

TLSA Value. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_association_data")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-3033101133030313-2221020232202330-0110201211113023-3022002112002310-0200331200210001-1230133320303110-1112211223330113-1001232112013232"></a>

### Direct properties for `primary.default_rr_set_group.tlsa_record.values`

<a id="canonical-2202311121123332-2023231031233301-3303012300022301-0020033213311212-2210302332102100-2311212113013020-3223323000212021-1313100321202233"></a>

#### `primary.default_rr_set_group.tlsa_record.values.certificate_association_data` property

Type: `"string"`. Optional.

The actual data to be matched given the settings of the other fields.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0113022033122333-3313311131112032-3030022320121300-2333302321230022-1112323100213131-1232222311222310-2222211013312320-0032120012312110"></a>

<a id="canonical-3303023023213033-3012200032210013-2330303232001323-0323112031232122-1022220103122321-0312112323330201-0022212133122013-0133323332020103"></a>

#### `primary.default_rr_set_group.tlsa_record.values.certificate_usage` property

Type: `"string"`. Optional.

\[Enum:
CertificateAuthorityConstraint|ServiceCertificateConstraint|TrustAnchorAssertion|DomainIssuedCertificate\]
&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint - ServiceCertificateConstraint:
Service Certificate Constraint - TrustAnchorAssertion: Trust Anchor Assertion -
DomainIssuedCertificate: Domain Issued Certificate. Possible values are
\`CertificateAuthorityConstraint\`, \`ServiceCertificateConstraint\`, \`TrustAnchorAssertion\`,
\`DomainIssuedCertificate\`. Defaults to \`CertificateAuthorityConstraint\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CertificateAuthorityConstraint","DomainIssuedCertificate","ServiceCertificateConstraint","TrustAnchorAssertion"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CertificateAuthorityConstraint",
    "ServiceCertificateConstraint",
    "TrustAnchorAssertion",
    "DomainIssuedCertificate"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CertificateAuthorityConstraint",
  "enum": [
    "CertificateAuthorityConstraint",
    "ServiceCertificateConstraint",
    "TrustAnchorAssertion",
    "DomainIssuedCertificate"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1023202102001012-3112230200010101-2321201132100033-0011113231330002-1323022102101133-1232311132322203-2003302222023300-3112032022303023"></a>

<a id="canonical-0210000321212133-3013020221121203-0101103012122201-3211101302113323-1121100012321333-0020233220020232-3323013310310212-0333300003321221"></a>

#### `primary.default_rr_set_group.tlsa_record.values.matching_type` property

Type: `"string"`. Optional.

\[Enum: NoHash|SHA256|SHA512\] - NoHash: No Hash - SHA256: SHA-256 - SHA512: SHA-512. Possible
values are \`NoHash\`, \`SHA256\`, \`SHA512\`. Defaults to \`NoHash\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["NoHash","SHA256","SHA512"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("NoHash",
    "SHA256",
    "SHA512"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NoHash",
  "enum": [
    "NoHash",
    "SHA256",
    "SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1203222200330322-1300312001022303-0100122323211131-2122203330333103-0300310121201012-2003000231312033-2322200333001333-1300020113311311"></a>

<a id="canonical-1100323110332011-0213021023302020-1221100313213212-0222100123003323-2012011303100010-2331213103330120-3221101231030111-0221011213132013"></a>

#### `primary.default_rr_set_group.tlsa_record.values.selector` property

Type: `"string"`. Optional.

\[Enum: FullCertificate|UseSubjectPublicKey\] - FullCertificate: Full Certificate -
UseSubjectPublicKey: Use Subject Public Key. Possible values are \`FullCertificate\`,
\`UseSubjectPublicKey\`. Defaults to \`FullCertificate\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["FullCertificate","UseSubjectPublicKey"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("FullCertificate",
    "UseSubjectPublicKey"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "FullCertificate",
  "enum": [
    "FullCertificate",
    "UseSubjectPublicKey"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3331231312030132-3010311201030021-0330032020003003-1111312303203000-3232301002223130-3103200221113020-1202122130031211-1301211330201120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_rr_set_group.txt_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.txt_record

<a id="canonical-1101322310230312-3222303222113203-2322030321021101-3321113220132023-2012323110330200-0103131230030110-3300121021102002-1022211223103031"></a>

Type: `"object"`. single nested block, Optional.

DNSTXTResourceRecord.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
txt_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200300103031113-0122102002132021-1103032232121120-1012023311310323-3201022132300310-3023030021132001-2212223010013201-1222003133003223"></a>

### Direct properties for `primary.default_rr_set_group.txt_record`

<a id="canonical-1021122013021120-1111100310231012-1003331323020003-1022023030313013-3203012320111313-0312211310312032-2000112201121211-3300232121200001"></a>

#### `primary.default_rr_set_group.txt_record.name` property

Type: `"string"`. Optional.

TXT Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-3203321001320331-0232131310021333-3320331222001200-1010030112223331-2002103203121112-0011011311332130-1010001210210030-1012311123120023"></a>

<a id="canonical-0222220101012001-0210111023320311-1313022322121022-1323322330102023-0102333201030203-1122001030112210-2303202032231101-3232120123220220"></a>

#### `primary.default_rr_set_group.txt_record.values` property

Type: `["list", "string"]`. Optional.

Text. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0001013031002232-2323220130321002-0110031123323021-1302303333101221-0212302022021323-3331121010231120-3131130311320003-3312131103003301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.default_soa_parameters` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- primary.default_soa_parameters

<a id="canonical-3200211123121101-0222032221231303-2032000002003211-0010311331231033-3032330230003031-1133110003111120-2222022011330333-3102013131201222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default soa parameters.

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
default_soa_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200011203232322-0232022120232023-3232232023010300-3212102331200100-1003120100321102-2232232021023203-1013232302222223-0200033311033102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.dnssec_mode` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- primary.dnssec_mode

<a id="canonical-2010020000212012-0312302321031122-0031232110200121-2021212131320032-1030321233203212-1122212300013010-0302322103301010-2301132300330100"></a>

Type: `"object"`. single nested block, Optional.

DNSSEC Mode.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-mode": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
dnssec_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-1301321130322033-2220232123213303-1231103221301133-3100211132111130-3211332333021103-2321303331110122-1102111133310010-2223001111100221"></a>

### Direct properties for `primary.dnssec_mode`

- [disable_spec](resources--dns_zone--reference--group-002.md#canonical-2110111011300011-0110222200202000-2331001010300133-2101102011023011-3210111013013002-2001331000032223-1122311331130203-1030210321032303): complete subsection reference.

- [enable](resources--dns_zone--reference--group-002.md#canonical-1310313012332122-1030002133303121-1003102013010323-3332112323301223-3110312030122203-0312021133201032-1100323202001132-2131130301303301): complete subsection reference.

<a id="canonical-2110111011300011-0110222200202000-2331001010300133-2101102011023011-3210111013013002-2001331000032223-1122311331130203-1030210321032303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.dnssec_mode.disable_spec` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.dnssec_mode](resources--dns_zone--reference--group-002.md#canonical-2200011203232322-0232022120232023-3232232023010300-3212102331200100-1003120100321102-2232232021023203-1013232302222223-0200033311033102)
- primary.dnssec_mode.disable_spec

<a id="canonical-0102011201220321-3223102212323133-1131213003012321-2300222021123323-2111313310131321-1033100302132121-0201310110333002-1330301122030221"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310313012332122-1030002133303121-1003102013010323-3332112323301223-3110312030122203-0312021133201032-1100323202001132-2131130301303301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.dnssec_mode.enable` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.dnssec_mode](resources--dns_zone--reference--group-002.md#canonical-2200011203232322-0232022120232023-3232232023010300-3212102331200100-1003120100321102-2232232021023203-1013232302222223-0200033311033102)
- primary.dnssec_mode.enable

<a id="canonical-2203011132121032-2323023131312322-1021021231331220-0131021102332033-3200331321131112-1130030220013101-2311023022113003-0200222031102223"></a>

Type: `["object", {}]`. Optional.

Enable. DNSSEC enable.

Receipt-pinned upstream constraints:

```json
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
enable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- primary.rr_set_group

<a id="canonical-2322001101312122-2303120132022113-0220120321000013-2022230220112311-0030333011013203-2232310232111211-1302303330103133-2221200303231123"></a>

Type: `"object"`. list nested block, Optional.

Create and manage set groups, and resource record sets within them, x-VES-I/O-managed set is managed
by F5.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
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
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rr_set_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-2030023201310033-0321023311211103-2202033122323221-3220020131231210-1101032100131021-3220110112333000-1022212130212101-2103103303331013"></a>

### Direct properties for `primary.rr_set_group`

- [metadata](resources--dns_zone--reference--group-002.md#canonical-0030212320120101-2321031322102202-0232302322102120-0030103101210333-2323021020120003-3300023212022220-2003202133231201-1301212100013303): complete subsection reference.

- [rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313): complete subsection reference.

<a id="canonical-0030212320120101-2321031322102202-0232302322102120-0030103101210333-2323021020120003-3300023212022220-2003202133231201-1301212100013303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.metadata` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- primary.rr_set_group.metadata

<a id="canonical-1002010330023123-0133011030001021-0201203112331312-3203023122201120-3010100112102002-2211331230303303-3100201021201011-0301202201111300"></a>

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

<a id="canonical-2301113111202230-1120310012130123-0131012132310031-1121301122113102-0233011331123300-0310001011331122-1321230000031330-3333000000333130"></a>

### Direct properties for `primary.rr_set_group.metadata`

<a id="canonical-1322010030310310-0013023022313121-1321130201203213-1020303010211132-1333232101323021-3001232033002200-3021331312303202-0200220301133203"></a>

#### `primary.rr_set_group.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2013210032112020-2012212313210202-3310000230131101-3102233100030210-0333311232132132-3021313113233010-0110022333233002-2300213020220012"></a>

<a id="canonical-1112023223321121-0122300231232023-0133221111303203-2133120313221130-2111230233313213-0323320231033320-0002313303101110-1130002212023003"></a>

#### `primary.rr_set_group.metadata.name` property

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

<a id="canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- primary.rr_set_group.rr_set

<a id="canonical-1321132121303313-2303003201012122-1132223021003222-3003122211220200-2032321003203202-2212100213021102-0323333113103201-2011103021121101"></a>

Type: `"object"`. list nested block, Optional.

Resource Record Sets. Collection of DNS resource record sets.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("ttl"),
  validators.ConflictingListObjectAttributes("a_record",
    "aaaa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "afsdb_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "afsdb_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("sshfp_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("sshfp_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("tlsa_record",
    "txt_record")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50000,
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
    "ves.io.schema.rules.repeated.max_items": "50000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50000"
  }
}
```

Terraform syntax:

```terraform
rr_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110033313021321-3330300303320203-0133311002331003-1000311201130030-3132022212333232-0203322112323123-0202023331133032-3233022222212331"></a>

### Direct properties for `primary.rr_set_group.rr_set`

- [a_record](resources--dns_zone--reference--group-002.md#canonical-1120220211221011-3230322322033011-1233032203220321-1102111001122332-1332121133201301-1033322220200120-2313023233322321-3010230212113003): complete subsection reference.

- [aaaa_record](resources--dns_zone--reference--group-002.md#canonical-1333023102122010-3110332232301001-2120203012203220-3011102200230000-3223333103313110-1331230111222313-1203023033000321-0112021323323303): complete subsection reference.

- [afsdb_record](resources--dns_zone--reference--group-002.md#canonical-1031301030212300-0033032201222102-0233230113301302-0110223101223112-3232032331132310-0021013222101111-0321311010333002-3311213030201330): complete subsection reference.

- [alias_record](resources--dns_zone--reference--group-002.md#canonical-0233221001322012-3112103303232212-1312101110101110-3121331300013112-0011002002223132-1231003102222231-1212201121230333-3331020002012320): complete subsection reference.

- [caa_record](resources--dns_zone--reference--group-002.md#canonical-3033201020333330-0032100001223222-2131222112130203-2010223121133122-3220201301130013-1211100133100131-1111022033130010-3311121023031100): complete subsection reference.

- [cds_record](resources--dns_zone--reference--group-002.md#canonical-0212131200323302-0300321013230033-2121301221232203-0233133001113103-3201313000002021-1131310322331320-1013231302130100-1033333322030133): complete subsection reference.

- [cert_record](resources--dns_zone--reference--group-002.md#canonical-2122132002202221-2302210131231003-3030120300120123-2133331322231013-1130023111130312-3031311111321320-2202301111312003-0301302003020101): complete subsection reference.

- [cname_record](resources--dns_zone--reference--group-002.md#canonical-1103130121131032-2002012232031213-3122130221012010-2331212231230031-3110031301310331-0321330030021320-3212001113023221-1011201311323322): complete subsection reference.

<a id="canonical-2133332032212210-0123213230222121-3123113232120122-1010332211233120-3100320331130120-0330130000220032-2113103331312030-1322231031232211"></a>

<a id="canonical-3300323211013030-2011320002002230-1002333121130213-1322113223023013-2033023212112221-2333313002222110-1201033132100102-1000321300002010"></a>

#### `primary.rr_set_group.rr_set.description_spec` property

Type: `"string"`. Optional.

Comment. Human-readable description text

- [ds_record](resources--dns_zone--reference--group-002.md#canonical-0133111032113033-3001031330210022-3133010321131111-3030032331330223-1030103230301101-1222213021103231-1113333233030203-1300220320123223): complete subsection reference.

- [eui48_record](resources--dns_zone--reference--group-002.md#canonical-3030020123311102-0330002300121201-0201032021303132-0331310302301121-2002232330201131-3201010010133002-1121220200310131-0121101001003201): complete subsection reference.

- [eui64_record](resources--dns_zone--reference--group-002.md#canonical-1123232112333312-0301023300332122-0100233233030001-1023002032221030-0013031220223001-2301310111033331-0120133033311203-0112320013203033): complete subsection reference.

- [lb_record](resources--dns_zone--reference--group-002.md#canonical-2011000133312033-0111302203233320-2010201212131212-2202122323010311-1322123301132002-1213120331011210-3311300321113201-3203231313323203): complete subsection reference.

- [loc_record](resources--dns_zone--reference--group-002.md#canonical-3003222202202221-1033200330333101-1301030003310233-1331212131201210-3302032301321211-0112023010221213-1200003133200302-0012102223210202): complete subsection reference.

- [mx_record](resources--dns_zone--reference--group-003.md#canonical-2003001322011302-1301231333320300-1203302312221330-2033303022131132-3121320130110031-2332100310010333-2130011113103032-1022320310102202): complete subsection reference.

- [naptr_record](resources--dns_zone--reference--group-003.md#canonical-0031001220121130-2232011132130122-3223223310210021-2102322233011223-2233030012200001-3232113133123013-2230020011101011-3203003021223223): complete subsection reference.

- [ns_record](resources--dns_zone--reference--group-003.md#canonical-3112102233211200-0313123203112300-1130013102230113-2011023232120220-1302012230210220-3202131201133110-0011021300201313-1102311332102132): complete subsection reference.

- [ptr_record](resources--dns_zone--reference--group-003.md#canonical-3232000310031110-1111032130320130-1330010220023223-2023200213300002-3123322311111130-2121011102311023-1203013003000322-3300000321210110): complete subsection reference.

- [srv_record](resources--dns_zone--reference--group-003.md#canonical-3331101323210103-3033131313122302-0130321211322012-2233211131121323-1203222333302212-2332301313331233-3123010110333331-2023103210211121): complete subsection reference.

- [sshfp_record](resources--dns_zone--reference--group-003.md#canonical-3021303131230121-2130011003320213-1320312132022020-1030220220200211-1220030223000022-0033023310312020-1222101103301213-2133330022031311): complete subsection reference.

- [tlsa_record](resources--dns_zone--reference--group-003.md#canonical-0031230210322331-3321210223301010-0132333332013133-0020032020210333-3021212331022022-3312031103113322-3133100300013111-0030030101213110): complete subsection reference.

<a id="canonical-2212321033023303-0323001231312003-3210032310003003-0201001200102131-0131103232010111-2122312100213023-1303001300022011-0200102110001300"></a>

<a id="canonical-0011311101130102-2123210002031033-2031213001223031-1022112232002011-1221010200002211-2311131331032220-3322323302210001-3312032303101320"></a>

#### `primary.rr_set_group.rr_set.ttl` property

Type: `"number"`. Optional.

Time to live. Time-to-live duration in seconds

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(60, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 60
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

- [txt_record](resources--dns_zone--reference--group-003.md#canonical-1121131111220220-0333230311100110-3220333331113330-3330020320131230-1033022131302013-0333232321103112-3113013123231300-0230232233102113): complete subsection reference.

<a id="canonical-1120220211221011-3230322322033011-1233032203220321-1102111001122332-1332121133201301-1033322220200120-2313023233322321-3010230212113003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.a_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.a_record

<a id="canonical-2031310110000111-1322322130201330-3323301001110212-3133310330133013-0210013033102131-2203230200200030-2213013220132213-1301102103210031"></a>

Type: `"object"`. single nested block, Optional.

DNSAResourceRecord. A Records

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
a_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032100101230123-1113210131021132-3022010123232223-2101002320202122-1123112233202222-1031331022100011-3033301133112311-3101312302021003"></a>

### Direct properties for `primary.rr_set_group.rr_set.a_record`

<a id="canonical-0220023032013002-1221230021322120-1213110113102212-2212322313110302-0021203310120013-0333222231333213-3301220311113111-3020010330101031"></a>

#### `primary.rr_set_group.rr_set.a_record.name` property

Type: `"string"`. Optional.

A Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-0300022002231102-2302030030032203-3030233303233321-0223032331312033-2323030002001021-3310200220101300-3022110323013323-0333223132133122"></a>

<a id="canonical-2231230332033100-0213130330312302-1200032101032022-3030022122131311-0121123210132122-2303102332120333-3130332000200120-0220101020101203"></a>

#### `primary.rr_set_group.rr_set.a_record.values` property

Type: `["list", "string"]`. Optional.

IPv4 Addresses. A valid IPv4 address, for example: 192.0.2.242.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1333023102122010-3110332232301001-2120203012203220-3011102200230000-3223333103313110-1331230111222313-1203023033000321-0112021323323303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.aaaa_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.aaaa_record

<a id="canonical-1323010102202331-0031123232222202-0211213021200023-3012223001120200-1132111003100012-2302000011113313-2120122120121020-0333031000330003"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for aaaa record.

Additional upstream details:

RecordSet for AAAA Records.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
aaaa_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210322033030330-1213002110230331-0023002132001303-1032113010231000-1012001221331111-2032000330300200-0332003131013000-1001123110321330"></a>

### Direct properties for `primary.rr_set_group.rr_set.aaaa_record`

<a id="canonical-3223312000221212-3003122023110122-2003002023300333-0211011333112000-0012300333010312-2222110201231331-0203200221131130-0132121232031223"></a>

#### `primary.rr_set_group.rr_set.aaaa_record.name` property

Type: `"string"`. Optional.

AAAA Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-3020221301220010-1032211021132313-1123220110010130-0122113131112121-2231112232001330-3103313032100111-2232111032011301-1022233132013200"></a>

<a id="canonical-2312203323033111-3101303233223321-3211202313003332-0113200131303130-3000333030203130-1110313131220003-2131003201211303-3211113211033023"></a>

#### `primary.rr_set_group.rr_set.aaaa_record.values` property

Type: `["list", "string"]`. Optional.

IPv6 Addresses. A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1031301030212300-0033032201222102-0233230113301302-0110223101223112-3232032331132310-0021013222101111-0321311010333002-3311213030201330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.afsdb_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.afsdb_record

<a id="canonical-1212310322223223-2312001020221020-2230222312010033-2320000220311213-0003133123122130-2310303301101003-3111300121022223-2213303111133022"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for afsdb record.

Additional upstream details:

DNS AFSDB Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
afsdb_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0123133312121111-2303330101313103-3222133322112100-2122131033131323-2302310120222122-0113130010222211-1313103021302003-3210010000031313"></a>

### Direct properties for `primary.rr_set_group.rr_set.afsdb_record`

<a id="canonical-1213302211210031-2133320200203000-0111102230320010-1221332300120121-0223003031321222-0303132322330003-0030220030203103-1112223023220031"></a>

#### `primary.rr_set_group.rr_set.afsdb_record.name` property

Type: `"string"`. Optional.

AFSDB Record name, please provide only the specific subdomain or record name without the base
domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-3301223013203213-3100232023203123-0323311031321001-0003112110023320-2303310020231103-2322133102323010-1130031222122123-0110302230131311): complete subsection reference.

<a id="canonical-3301223013203213-3100232023203123-0323311031321001-0003112110023320-2303310020231103-2322133102323010-1130031222122123-0110302230131311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.afsdb_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.afsdb_record](resources--dns_zone--reference--group-002.md#canonical-1031301030212300-0033032201222102-0233230113301302-0110223101223112-3232032331132310-0021013222101111-0321311010333002-3311213030201330)
- primary.rr_set_group.rr_set.afsdb_record.values

<a id="canonical-2302200222330321-1300110110323123-1303233000130322-3201222130200231-0003332331333121-3201222231130021-0232113000202130-0232132312031203"></a>

Type: `"object"`. list nested block, Optional.

AFSDB Value. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("hostname")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320223310331303-0310102023310002-0323130202110321-3233300312220333-0110032231100032-3201103311020102-3311213321122223-2210133032213331"></a>

### Direct properties for `primary.rr_set_group.rr_set.afsdb_record.values`

<a id="canonical-2102002201003330-2002101221012223-3001203121312002-3211211121013330-3111312122230022-1210303123021302-3111023333013132-2313221302033331"></a>

#### `primary.rr_set_group.rr_set.afsdb_record.values.hostname` property

Type: `"string"`. Optional.

Server name of the AFS cell database server or the DCE name server.

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
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-2000111133310013-2233120201210110-2220002303030010-2110000110303223-3010023113213023-3123100010310020-1201230020100323-3032301122200321"></a>

<a id="canonical-2200013021210211-0322301020020131-0003132233212010-1121211010210310-2021130212013001-3130311101333232-1222020202231311-0130100023200102"></a>

#### `primary.rr_set_group.rr_set.afsdb_record.values.subtype` property

Type: `"string"`. Optional.

\[Enum: NONE|AFSVolumeLocationServer|DCEAuthenticationServer\] AFS Volume Location Server or DCE
Authentication Server. - NONE: NONE - AFSVolumeLocationServer: AFS Volume Location Server -
DCEAuthenticationServer: DCE Authentication Server. Possible values are \`NONE\`,
\`AFSVolumeLocationServer\`, \`DCEAuthenticationServer\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["AFSVolumeLocationServer","DCEAuthenticationServer","NONE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("NONE",
    "AFSVolumeLocationServer",
    "DCEAuthenticationServer"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NONE",
  "enum": [
    "NONE",
    "AFSVolumeLocationServer",
    "DCEAuthenticationServer"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0233221001322012-3112103303232212-1312101110101110-3121331300013112-0011002002223132-1231003102222231-1212201121230333-3331020002012320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.alias_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.alias_record

<a id="canonical-3023031303112002-1011220313330031-3300012333020203-1103120213122123-0332110100122323-0321332213233010-1003301332021321-3322231133202100"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for alias record.

Receipt-pinned upstream constraints:

```json
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
alias_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303310201112310-3110332010032330-2133231202030131-3233233133302131-2210102233300311-0003313303021022-2333320200032101-3322210003100330"></a>

### Direct properties for `primary.rr_set_group.rr_set.alias_record`

<a id="canonical-3203103303223000-2221133301210312-1313103330113003-1000031030201221-0122121222313102-1221020030310013-3021011011232033-3231211111121212"></a>

#### `primary.rr_set_group.rr_set.alias_record.value` property

Type: `"string"`. Optional.

Domain. A valid domain name, for example: example.com.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-3033201020333330-0032100001223222-2131222112130203-2010223121133122-3220201301130013-1211100133100131-1111022033130010-3311121023031100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.caa_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.caa_record

<a id="canonical-3203022223320231-1032100010021023-3102132232202011-3221031132212002-2332032001000210-3210130210033031-0230230202331030-0132232201123103"></a>

Type: `"object"`. single nested block, Optional.

DNSCAAResourceRecord.

Receipt-pinned upstream constraints:

```json
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
caa_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-1023203333202013-1011032320303301-0201232002233202-1232221021301002-2321212310231211-1132310331003322-2202331203312022-2300221332220120"></a>

### Direct properties for `primary.rr_set_group.rr_set.caa_record`

<a id="canonical-0202232310230231-0100333010221323-2311013202202210-1230220200223232-0130303212131312-1131222311113120-1033123120003232-3010203312311203"></a>

#### `primary.rr_set_group.rr_set.caa_record.name` property

Type: `"string"`. Optional.

CAA Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-3233301110320230-2121202201132103-3002310301013122-2200203211110111-0012201220000310-0102313030033200-1131210233232132-0110201023022101): complete subsection reference.

<a id="canonical-3233301110320230-2121202201132103-3002310301013122-2200203211110111-0012201220000310-0102313030033200-1131210233232132-0110201023022101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.caa_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.caa_record](resources--dns_zone--reference--group-002.md#canonical-3033201020333330-0032100001223222-2131222112130203-2010223121133122-3220201301130013-1211100133100131-1111022033130010-3311121023031100)
- primary.rr_set_group.rr_set.caa_record.values

<a id="canonical-2113020310310122-0010131022220111-2010230133103022-0213103110113121-1002000132023232-2111320322131310-0112002102320313-1102302232201030"></a>

Type: `"object"`. list nested block, Optional.

CAA Record Value. Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201202000202320-0321002213121023-2211333321100303-2313001022133123-2021301033111221-2101100313303001-3022021111222022-0202202121122020"></a>

### Direct properties for `primary.rr_set_group.rr_set.caa_record.values`

<a id="canonical-0301013201232232-3231012001032130-2102210023001223-3113003232330333-1230031010002122-3123133202333113-0212022022112022-2103100200100311"></a>

#### `primary.rr_set_group.rr_set.caa_record.values.flags` property

Type: `"number"`. Optional.

This flag should be an integer between 0 and 255.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-0330230220122010-2001122300032033-0202022222323200-0233331130103322-1131320300310111-1332321010302331-0203103301013320-3231111120230011"></a>

<a id="canonical-1222211110312010-2300032322200311-2302332102010112-0112213110323012-3221211120321021-3111100103331310-3021312233302122-0031003022020312"></a>

#### `primary.rr_set_group.rr_set.caa_record.values.tag` property

Type: `"string"`. Optional.

\[Enum: issue|issuewild|iodef\] Tag. Tag for categorization and filtering. Possible values are
\`issue\`, \`issuewild\`, \`iodef\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["iodef","issue","issuewild"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("issue",
    "issuewild",
    "iodef"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "issue",
    "issuewild",
    "iodef"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  }
}
```

<a id="canonical-0000131020323322-3312101311333222-3310002101231232-2213233023311213-3002212311223300-0310121023110323-1301230212012130-0211303130103132"></a>

<a id="canonical-0331302203333120-2223302031012212-0312000131213211-1001121032000313-3011010200133113-3321121203311021-0000122200103200-1012001301310133"></a>

#### `primary.rr_set_group.rr_set.caa_record.values.value` property

Type: `"string"`. Optional.

Value. Configuration parameter for value

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0212131200323302-0300321013230033-2121301221232203-0233133001113103-3201313000002021-1131310322331320-1013231302130100-1033333322030133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.cds_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.cds_record

<a id="canonical-3012220321222313-2102100303223132-3320001303231312-1113220202122233-3233033202202311-3303302130320322-3130001301210120-3230110112300123"></a>

Type: `"object"`. single nested block, Optional.

DNS CDS Record. DNS CDS Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
cds_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-1230200010210003-2223301210223102-1103002303320013-2120023000212111-0321233110112030-2313230233232022-1200231211230102-3022001212231321"></a>

### Direct properties for `primary.rr_set_group.rr_set.cds_record`

<a id="canonical-1202020213032001-1221001000120313-0200233210221010-1120001323311232-0231212323131302-1220021102102332-1111103201120221-0023103032201132"></a>

#### `primary.rr_set_group.rr_set.cds_record.name` property

Type: `"string"`. Optional.

CDS Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033): complete subsection reference.

<a id="canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.cds_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-0212131200323302-0300321013230033-2121301221232203-0233133001113103-3201313000002021-1131310322331320-1013231302130100-1033333322030133)
- primary.rr_set_group.rr_set.cds_record.values

<a id="canonical-2321310003223320-1321332333030003-2322033021331223-1330322002030000-3221210101203310-1200302130113113-0210313223323033-2023023212230202"></a>

Type: `"object"`. list nested block, Optional.

DS Value. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("key_tag"),
  validators.ConflictingListObjectAttributes("sha1_digest",
    "sha256_digest"),
  validators.ConflictingListObjectAttributes("sha1_digest",
    "sha384_digest"),
  validators.ConflictingListObjectAttributes("sha256_digest",
    "sha384_digest")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001300212220323-1333231020303312-0212021133010032-0022311233222100-2002222030332102-3303013031010311-3223231310113220-2121113133112121"></a>

### Direct properties for `primary.rr_set_group.rr_set.cds_record.values`

<a id="canonical-3221313023302011-1302331132022102-2102313232003220-1011111302030030-2132232010020200-3212131221320011-0101312021001313-2011310131011331"></a>

#### `primary.rr_set_group.rr_set.cds_record.values.ds_key_algorithm` property

Type: `"string"`. Optional.

\[Enum:
UNSPECIFIED|RSASHA1|RSASHA1NSEC3SHA1|RSASHA256|RSASHA512|ECDSAP256SHA256|ECDSAP384SHA384|ED25519|ED448\]
DS key-value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1:
RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 -
ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448:
ED448. Possible values are \`UNSPECIFIED\`, \`RSASHA1\`, \`RSASHA1NSEC3SHA1\`, \`RSASHA256\`,
\`RSASHA512\`, \`ECDSAP256SHA256\`, \`ECDSAP384SHA384\`, \`ED25519\`, \`ED448\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ECDSAP256SHA256","ECDSAP384SHA384","ED25519","ED448","RSASHA1","RSASHA1NSEC3SHA1","RSASHA256","RSASHA512","UNSPECIFIED"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIED",
  "enum": [
    "UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3231012122013300-0201102022233213-3230031002000002-1022233323000132-1132223033013323-1210102310333003-0201113120223330-1102111303211100"></a>

<a id="canonical-0020312121300020-0030303011322313-3233123330123121-3001310330032321-0021212321223122-1022301001111311-3021200232111231-3312123320300330"></a>

#### `primary.rr_set_group.rr_set.cds_record.values.key_tag` property

Type: `"number"`. Optional.

A short numeric value which can help quickly identify the referenced DNSKEY-record.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [sha1_digest](resources--dns_zone--reference--group-002.md#canonical-3100120201012212-2003022220310300-3201231103212000-2113003003200130-0200023032301110-3232330220223333-2033203112000101-0103201010021303): complete subsection reference.

- [sha256_digest](resources--dns_zone--reference--group-002.md#canonical-2302310320330210-0313333002331000-2112300301232232-3322031231010011-2101010213322220-3111320011201223-3300032310020120-3321323123001031): complete subsection reference.

- [sha384_digest](resources--dns_zone--reference--group-002.md#canonical-2032220222001301-1021003012000102-0101002330300323-0312203302002020-0003200112112211-3021032200112002-2003312333133200-3113331200110221): complete subsection reference.

<a id="canonical-3100120201012212-2003022220310300-3201231103212000-2113003003200130-0200023032301110-3232330220223333-2033203112000101-0103201010021303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.cds_record.values.sha1_digest` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-0212131200323302-0300321013230033-2121301221232203-0233133001113103-3201313000002021-1131310322331320-1013231302130100-1033333322030133)
- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-002.md#canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033)
- primary.rr_set_group.rr_set.cds_record.values.sha1_digest

<a id="canonical-3033132103122202-2320322021002103-2212203030021003-3313002010003333-0320010103312010-3332021022111200-0130221032312321-1230113302010031"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha1 digest.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha1_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231033010133020-1322210223030022-2130020213030200-2202133323312322-3232300103132021-3211202010311000-2232322123231103-2300233303002311"></a>

### Direct properties for `primary.rr_set_group.rr_set.cds_record.values.sha1_digest`

<a id="canonical-2132020012111003-1303102023010130-1010202313311110-2100201300203203-2201121330113311-0220010203302022-2223132130313211-3312311030322213"></a>

#### `primary.rr_set_group.rr_set.cds_record.values.sha1_digest.digest` property

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(40, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 40
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-2302310320330210-0313333002331000-2112300301232232-3322031231010011-2101010213322220-3111320011201223-3300032310020120-3321323123001031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.cds_record.values.sha256_digest` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-0212131200323302-0300321013230033-2121301221232203-0233133001113103-3201313000002021-1131310322331320-1013231302130100-1033333322030133)
- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-002.md#canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033)
- primary.rr_set_group.rr_set.cds_record.values.sha256_digest

<a id="canonical-0010233123223233-2012102031300132-3312323022113110-3031311102102121-2101311323203202-2131032213233033-1131311220010212-2221222002001002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha256 digest.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha256_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202331012030300-1203311110323110-2213310203233112-1102301301032023-3102313310110112-1012222012023000-1102213130102322-2222301122132302"></a>

### Direct properties for `primary.rr_set_group.rr_set.cds_record.values.sha256_digest`

<a id="canonical-0221230202212112-0211131000331112-2321230213212000-3122310223312310-0210202320203302-2231231020101232-3230200313113132-3000021201202213"></a>

#### `primary.rr_set_group.rr_set.cds_record.values.sha256_digest.digest` property

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(64, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 64
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-2032220222001301-1021003012000102-0101002330300323-0312203302002020-0003200112112211-3021032200112002-2003312333133200-3113331200110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.cds_record.values.sha384_digest` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-0212131200323302-0300321013230033-2121301221232203-0233133001113103-3201313000002021-1131310322331320-1013231302130100-1033333322030133)
- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-002.md#canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033)
- primary.rr_set_group.rr_set.cds_record.values.sha384_digest

<a id="canonical-3120330220013230-0303113122202021-2201232202113323-2022113002210013-0330202233323122-0233122303011112-1210231220322032-0101310120112201"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha384 digest.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha384_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302221100020130-3330110313211132-0322013001310031-0101331211122021-3330032122300233-1332222121032313-3011003000203121-0110131222202302"></a>

### Direct properties for `primary.rr_set_group.rr_set.cds_record.values.sha384_digest`

<a id="canonical-1130113223110200-0120022101323323-2031221211033022-0120123132011321-2021010302130322-0331110032010222-3300220222212311-0312312030112311"></a>

#### `primary.rr_set_group.rr_set.cds_record.values.sha384_digest.digest` property

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(96, 96),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-2122132002202221-2302210131231003-3030120300120123-2133331322231013-1130023111130312-3031311111321320-2202301111312003-0301302003020101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.cert_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.cert_record

<a id="canonical-3012301320210302-1333331200213200-3010310033130210-1223002010211130-0022101013011131-0230313332121221-0332120101200133-1231213322232330"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for cert record.

Additional upstream details:

DNS CERT Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
cert_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002112322021130-3101223213133010-3330003311021321-2331333110033012-2021211221201031-3322112002032032-1020203211223123-2222323203010023"></a>

### Direct properties for `primary.rr_set_group.rr_set.cert_record`

<a id="canonical-2020032320331031-0011123303232211-0321013130110322-0013222113013223-0111012212223311-1111032321332332-0222332031302331-0230333013301020"></a>

#### `primary.rr_set_group.rr_set.cert_record.name` property

Type: `"string"`. Optional.

CERT Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-1110212302332301-1330131123332013-1122230122101023-3000000013013121-2213021011231101-1100133232033221-1123322331330012-3012232330233330): complete subsection reference.

<a id="canonical-1110212302332301-1330131123332013-1122230122101023-3000000013013121-2213021011231101-1100133232033221-1123322331330012-3012232330233330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.cert_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.cert_record](resources--dns_zone--reference--group-002.md#canonical-2122132002202221-2302210131231003-3030120300120123-2133331322231013-1130023111130312-3031311111321320-2202301111312003-0301302003020101)
- primary.rr_set_group.rr_set.cert_record.values

<a id="canonical-2321003200110130-1233110210113031-1223210301301013-2121223122301131-1323123030313000-3023232220110202-0033103212123212-1003302122000220"></a>

Type: `"object"`. list nested block, Optional.

CERT Value. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("cert_key_tag",
    "certificate")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310220020233220-0333023320012021-3112111110113223-2302010130212201-0232331233122102-2212323122021200-1320001301222030-1312111132113010"></a>

### Direct properties for `primary.rr_set_group.rr_set.cert_record.values`

<a id="canonical-1211012122010310-0201120313120013-1032322211321110-3133121300311113-3210223231022230-1311321311232212-2301212110321012-1112220013310023"></a>

#### `primary.rr_set_group.rr_set.cert_record.values.algorithm` property

Type: `"string"`. Optional.

\[Enum: RESERVEDALGORITHM|RSAMD5|DH|DSASHA1|ECC|RSASHA1ALGORITHM|INDIRECT|PRIVATEDNS|PRIVATEOID\]
CERT algorithm value must be compatible with the specified algorithm. - RESERVEDALGORITHM:
RESERVEDALGORITHM - RSAMD5: RSAMD5 - DH: DH - DSASHA1: DSASHA1 - ECC: ECC - RSASHA1ALGORITHM:
RSA-SHA1 - INDIRECT: INDIRECT - PRIVATEDNS: PRIVATEDNS - PRIVATEOID: PRIVATEOID. Possible values are
\`RESERVEDALGORITHM\`, \`RSAMD5\`, \`DH\`, \`DSASHA1\`, \`ECC\`, \`RSASHA1ALGORITHM\`, \`INDIRECT\`,
\`PRIVATEDNS\`, \`PRIVATEOID\`. Defaults to \`RESERVEDALGORITHM\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DH","DSASHA1","ECC","INDIRECT","PRIVATEDNS","PRIVATEOID","RESERVEDALGORITHM","RSAMD5","RSASHA1ALGORITHM"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("RESERVEDALGORITHM",
    "RSAMD5",
    "DH",
    "DSASHA1",
    "ECC",
    "RSASHA1ALGORITHM",
    "INDIRECT",
    "PRIVATEDNS",
    "PRIVATEOID"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "RESERVEDALGORITHM",
  "enum": [
    "RESERVEDALGORITHM",
    "RSAMD5",
    "DH",
    "DSASHA1",
    "ECC",
    "RSASHA1ALGORITHM",
    "INDIRECT",
    "PRIVATEDNS",
    "PRIVATEOID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1303010300120031-1323220122303030-2132211322123221-0022313212032000-0322201331020131-1120312020210320-2122311112312121-1212323022302322"></a>

<a id="canonical-1210322033311010-1220332333011331-2033011011302221-2321103221022121-0110100211002313-2121122020221312-1203111320220300-0133313233012032"></a>

#### `primary.rr_set_group.rr_set.cert_record.values.cert_key_tag` property

Type: `"number"`. Optional.

Key Tag. Tag for categorization and filtering

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2131303313320230-3121011203223201-3112232213321203-2021210323002222-3133020003302001-3033100333133332-3230300002002000-2210211112033030"></a>

<a id="canonical-3223013032003100-0222301223120100-1021131300002001-1023023333302103-0321032002010021-0020013122303201-1313110323200230-2111033221011122"></a>

#### `primary.rr_set_group.rr_set.cert_record.values.cert_type` property

Type: `"string"`. Optional.

\[Enum: INVALIDCERTTYPE|PKIX|SPKI|PGP|IPKIX|ISPKI|IPGP|ACPKIX|IACPKIX|URI\_|OID\] CERT type value
must be compatible with the specified types. - INVALIDCERTTYPE: INVALIDCERTTYPE - PKIX: PKIX - SPKI:
SPKI - PGP: PGP - IPKIX: IPKIX - ISPKI: ISPKI - IPGP: IPGP - ACPKIX: ACPKIX - IACPKIX: IACPKIX -
URI\_: URI - OID: OID. Possible values are \`INVALIDCERTTYPE\`, \`PKIX\`, \`SPKI\`, \`PGP\`,
\`IPKIX\`, \`ISPKI\`, \`IPGP\`, \`ACPKIX\`, \`IACPKIX\`, \`URI\_\`, \`OID\`. Defaults to
\`INVALIDCERTTYPE\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ACPKIX","IACPKIX","INVALIDCERTTYPE","IPGP","IPKIX","ISPKI","OID","PGP","PKIX","SPKI","URI_"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("INVALIDCERTTYPE",
    "PKIX",
    "SPKI",
    "PGP",
    "IPKIX",
    "ISPKI",
    "IPGP",
    "ACPKIX",
    "IACPKIX",
    "URI_",
    "OID"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INVALIDCERTTYPE",
  "enum": [
    "INVALIDCERTTYPE",
    "PKIX",
    "SPKI",
    "PGP",
    "IPKIX",
    "ISPKI",
    "IPGP",
    "ACPKIX",
    "IACPKIX",
    "URI_",
    "OID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1030030023212121-1310122023033021-1222333022110320-2030112330130132-2021320300013223-3231113123013013-0221132013100002-0330200333032321"></a>

<a id="canonical-3132201310332023-2230022211120212-3100300112102231-2111003031220222-3221333130212101-0312130022100001-3321001131222013-1323013200000332"></a>

#### `primary.rr_set_group.rr_set.cert_record.values.certificate` property

Type: `"string"`. Optional.

Certificate. Certificate in base 64 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 5242880,
      "min": 100
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "pem",
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
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
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1103130121131032-2002012232031213-3122130221012010-2331212231230031-3110031301310331-0321330030021320-3212001113023221-1011201311323322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.cname_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.cname_record

<a id="canonical-0310311121001122-2211200002312221-3310123210132123-1111003013102101-0220013300100011-3323021133223210-1203332212331011-2031212020103123"></a>

Type: `"object"`. single nested block, Optional.

DNSCNAMEResourceRecord.

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
cname_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321230311311111-1110100300310031-2220021000022300-2211202113333322-2032220231002110-1330310100130213-3133031231000131-1213130012013201"></a>

### Direct properties for `primary.rr_set_group.rr_set.cname_record`

<a id="canonical-1220111321122301-3001030230300300-0313021211213223-3313012203222220-3233002233010220-3300121311110312-2112210113022301-3022113321331133"></a>

#### `primary.rr_set_group.rr_set.cname_record.name` property

Type: `"string"`. Optional.

CName Record name, please provide only the specific subdomain or record name without the base
domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-2120103131210132-3011322322030331-3011211302212322-0232133033012312-0020100212002222-2321013301033303-3231312232333112-3002112120212101"></a>

<a id="canonical-3033120303031332-0321103023121021-0002030332330212-2321123112112310-3013132230223121-0101232033203023-0312231301003310-1000321121321332"></a>

#### `primary.rr_set_group.rr_set.cname_record.value` property

Type: `"string"`. Optional.

Domain. Configuration parameter for value

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-0133111032113033-3001031330210022-3133010321131111-3030032331330223-1030103230301101-1222213021103231-1113333233030203-1300220320123223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.ds_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.ds_record

<a id="canonical-1312320202330202-3022222003331233-1020030132132200-1011023111332103-3330200202323012-1123022031311220-2221113301031332-2302121100021102"></a>

Type: `"object"`. single nested block, Optional.

DNS DS Record. DNS DS Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
ds_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011222200100310-1010333112211321-3200202033332130-0300030030131011-2302032013132302-2131132331221122-2030312310210223-0013013203032003"></a>

### Direct properties for `primary.rr_set_group.rr_set.ds_record`

<a id="canonical-0032321310113023-1133320130332132-2121301000120130-1023313113323311-0330201003120100-2023003231303311-3101000322230333-1220332110021031"></a>

#### `primary.rr_set_group.rr_set.ds_record.name` property

Type: `"string"`. Optional.

DS Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-3002100320102213-3300132232333231-3121011221223032-1102300223033120-0011311221003330-3021330103033101-0000320030230013-0310012032010331): complete subsection reference.

<a id="canonical-3002100320102213-3300132232333231-3121011221223032-1102300223033120-0011311221003330-3021330103033101-0000320030230013-0310012032010331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.ds_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-002.md#canonical-0133111032113033-3001031330210022-3133010321131111-3030032331330223-1030103230301101-1222213021103231-1113333233030203-1300220320123223)
- primary.rr_set_group.rr_set.ds_record.values

<a id="canonical-2100203210333023-2131000233103233-3110200332300003-0202103210212130-0311330213233211-3003301122200002-0210011133323020-3132332211132201"></a>

Type: `"object"`. list nested block, Optional.

DS Value. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("key_tag"),
  validators.ConflictingListObjectAttributes("sha1_digest",
    "sha256_digest"),
  validators.ConflictingListObjectAttributes("sha1_digest",
    "sha384_digest"),
  validators.ConflictingListObjectAttributes("sha256_digest",
    "sha384_digest")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220210121330331-2223100011211011-3122012133011330-3313120011211203-3202302101030333-2332300223130132-2000020300122000-2022022331311320"></a>

### Direct properties for `primary.rr_set_group.rr_set.ds_record.values`

<a id="canonical-2323221310033131-0222331312210010-1322132223111200-0020031323310313-1203021201000031-0100113321133012-0303303100011131-1330230201310300"></a>

#### `primary.rr_set_group.rr_set.ds_record.values.ds_key_algorithm` property

Type: `"string"`. Optional.

\[Enum:
UNSPECIFIED|RSASHA1|RSASHA1NSEC3SHA1|RSASHA256|RSASHA512|ECDSAP256SHA256|ECDSAP384SHA384|ED25519|ED448\]
DS key-value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1:
RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 -
ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448:
ED448. Possible values are \`UNSPECIFIED\`, \`RSASHA1\`, \`RSASHA1NSEC3SHA1\`, \`RSASHA256\`,
\`RSASHA512\`, \`ECDSAP256SHA256\`, \`ECDSAP384SHA384\`, \`ED25519\`, \`ED448\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ECDSAP256SHA256","ECDSAP384SHA384","ED25519","ED448","RSASHA1","RSASHA1NSEC3SHA1","RSASHA256","RSASHA512","UNSPECIFIED"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIED",
  "enum": [
    "UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1022013312210330-3111132222231013-0131033022321103-1203322330311311-0233320321200221-2103232330233210-1131331232030031-1303232033110231"></a>

<a id="canonical-0201312311201221-1313130300112002-1223132323212133-3220101020023113-0202333001102031-1011230211001213-1113322013023131-1321303003001030"></a>

#### `primary.rr_set_group.rr_set.ds_record.values.key_tag` property

Type: `"number"`. Optional.

A short numeric value which can help quickly identify the referenced DNSKEY-record.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [sha1_digest](resources--dns_zone--reference--group-002.md#canonical-1221213211021330-1020231010121110-1213003033310222-2002101033112223-2321201230033323-0103212031133023-0030002331221301-0111033300002012): complete subsection reference.

- [sha256_digest](resources--dns_zone--reference--group-002.md#canonical-1102203212111212-0030020311132130-0021222332121310-3011012022233211-0322312122333202-2332120213121001-0010201033302330-1100120312132210): complete subsection reference.

- [sha384_digest](resources--dns_zone--reference--group-002.md#canonical-1311123030010222-0013302212333220-2102221011302212-0333230023003222-3000320213321330-3113123310312010-3331130023103201-2223231323312320): complete subsection reference.

<a id="canonical-1221213211021330-1020231010121110-1213003033310222-2002101033112223-2321201230033323-0103212031133023-0030002331221301-0111033300002012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.ds_record.values.sha1_digest` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-002.md#canonical-0133111032113033-3001031330210022-3133010321131111-3030032331330223-1030103230301101-1222213021103231-1113333233030203-1300220320123223)
- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-3002100320102213-3300132232333231-3121011221223032-1102300223033120-0011311221003330-3021330103033101-0000320030230013-0310012032010331)
- primary.rr_set_group.rr_set.ds_record.values.sha1_digest

<a id="canonical-0011133131321023-3130203301120032-0110112220131102-2111121011201021-2211331322120130-2002031220021330-2112231221322020-0223322123323200"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha1 digest.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha1_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220021012032001-3130123031020232-2132202002121122-1331123033031021-2031200220133013-2321212100323210-3001131300321013-1231320221122001"></a>

### Direct properties for `primary.rr_set_group.rr_set.ds_record.values.sha1_digest`

<a id="canonical-2232131303232033-0333312022022312-1111110010021003-2101222120032030-3330123100230122-2130222210233312-0022103013111032-0202231013211033"></a>

#### `primary.rr_set_group.rr_set.ds_record.values.sha1_digest.digest` property

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(40, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 40
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-1102203212111212-0030020311132130-0021222332121310-3011012022233211-0322312122333202-2332120213121001-0010201033302330-1100120312132210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.ds_record.values.sha256_digest` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-002.md#canonical-0133111032113033-3001031330210022-3133010321131111-3030032331330223-1030103230301101-1222213021103231-1113333233030203-1300220320123223)
- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-3002100320102213-3300132232333231-3121011221223032-1102300223033120-0011311221003330-3021330103033101-0000320030230013-0310012032010331)
- primary.rr_set_group.rr_set.ds_record.values.sha256_digest

<a id="canonical-3223300111103300-2131113003121202-0300202322223220-1212302211030213-1133121200002131-2300100110232331-0220122010230232-2213220212023120"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha256 digest.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha256_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-2030221320012002-1311213200202002-3331210201222023-0132032021220220-0032211112300031-2210132200312230-0130103013230200-3333133031232011"></a>

### Direct properties for `primary.rr_set_group.rr_set.ds_record.values.sha256_digest`

<a id="canonical-1102232200130203-3213023331321211-0313023210120322-2121210303232303-3020110200013132-3312310210123331-3231002011333230-1213030122123000"></a>

#### `primary.rr_set_group.rr_set.ds_record.values.sha256_digest.digest` property

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(64, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 64
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-1311123030010222-0013302212333220-2102221011302212-0333230023003222-3000320213321330-3113123310312010-3331130023103201-2223231323312320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.ds_record.values.sha384_digest` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-002.md#canonical-0133111032113033-3001031330210022-3133010321131111-3030032331330223-1030103230301101-1222213021103231-1113333233030203-1300220320123223)
- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-3002100320102213-3300132232333231-3121011221223032-1102300223033120-0011311221003330-3021330103033101-0000320030230013-0310012032010331)
- primary.rr_set_group.rr_set.ds_record.values.sha384_digest

<a id="canonical-2020321033011002-0102000003310030-2202320000020333-0010133010030020-0130313010330210-3311003231003101-0212100220121010-2220132120202300"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha384 digest.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha384_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133223123322230-3333112312031000-0320010333123200-2001200211120021-0132211110202311-3020131111000303-1122012232323003-2302112012103010"></a>

### Direct properties for `primary.rr_set_group.rr_set.ds_record.values.sha384_digest`

<a id="canonical-1320313210322131-2231002220303302-0332223110112201-2232123012321012-1221031322020202-0210103332300123-2223332212300301-1131002122200030"></a>

#### `primary.rr_set_group.rr_set.ds_record.values.sha384_digest.digest` property

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(96, 96),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-3030020123311102-0330002300121201-0201032021303132-0331310302301121-2002232330201131-3201010010133002-1121220200310131-0121101001003201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.eui48_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.eui48_record

<a id="canonical-0201123302222103-2312013113132012-0122000120310320-0320022121033100-3013022323220030-3331223213221333-2321013100233321-3302330110121232"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for eui48 record.

Additional upstream details:

DNS EUI48 Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("value")}
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
eui48_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301221301212033-2310113210303003-1331033233012320-2131220020211113-0220230322032001-3010121320103031-0223133331320330-3113101230222003"></a>

### Direct properties for `primary.rr_set_group.rr_set.eui48_record`

<a id="canonical-2002122332210301-1330301002010310-3110211312012330-2323213122132332-1333031220133212-1220032013012131-0133100103102002-1331021133023020"></a>

#### `primary.rr_set_group.rr_set.eui48_record.name` property

Type: `"string"`. Optional.

EUI48 Record name, please provide only the specific subdomain or record name without the base
domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-0002023132231011-0303112102101232-1021233301002103-3320333213213013-1212111020313131-1311033201013113-2110200100221133-2121213220021303"></a>

<a id="canonical-2310322123231103-1132330030023021-0221322022303130-0211030103331231-0100232310010333-1133223201210312-3232302130213022-0212103020030221"></a>

#### `primary.rr_set_group.rr_set.eui48_record.value` property

Type: `"string"`. Optional.

EUI48 Identifier. A valid eui48 identifier, for example: 01-23-45-67-89-ab.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(17, 17),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 17,
  "minLength": 17,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 17,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 17,
    "pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-1123232112333312-0301023300332122-0100233233030001-1023002032221030-0013031220223001-2301310111033331-0120133033311203-0112320013203033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.eui64_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.eui64_record

<a id="canonical-2013211131032121-0312023111000010-0122020330121311-3122211222020033-0030010323222303-0111203231222112-2211112301312120-1210203033031233"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for eui64 record.

Additional upstream details:

DNS EUI64 Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("value")}
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
eui64_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-1331211311221011-0012123223033310-1010001313202120-1132222132101020-3032102022120021-3331310010111232-0033010002033120-3223123123030101"></a>

### Direct properties for `primary.rr_set_group.rr_set.eui64_record`

<a id="canonical-3003013232120122-2111212112102312-2002331003000031-0223123230200221-1022202212233002-3101121031230310-3230022312231202-2221012103233032"></a>

#### `primary.rr_set_group.rr_set.eui64_record.name` property

Type: `"string"`. Optional.

EUI64 Record name, please provide only the specific subdomain or record name without the base
domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-1102123323213102-3112220233132131-2232323223003302-0101010323201231-3130012311112111-3100222311132122-3113323010332001-1221321213220212"></a>

<a id="canonical-3322213320132210-2023122012112112-3331202122013330-0103011310330313-1201333230230020-1311232100303222-0003202331111133-3233220201233110"></a>

#### `primary.rr_set_group.rr_set.eui64_record.value` property

Type: `"string"`. Optional.

EUI64 Identifier. A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(23, 23),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 23,
  "minLength": 23,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 23,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 23,
    "pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-2011000133312033-0111302203233320-2010201212131212-2202122323010311-1322123301132002-1213120331011210-3311300321113201-3203231313323203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.lb_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.lb_record

<a id="canonical-0021120310010231-0211130000112132-3000333101232313-3320012233310210-2300011022210313-2233120301332012-2023332030010130-0213202201302233"></a>

Type: `"object"`. single nested block, Optional.

DNS Load Balancer Record. DNS Load Balancer Record.

Receipt-pinned upstream constraints:

```json
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
lb_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200122333301331-3101203222111302-3233331333010203-0320202013023300-0213203110301210-1320332033203110-1220200123212003-3310103303202300"></a>

### Direct properties for `primary.rr_set_group.rr_set.lb_record`

<a id="canonical-0021020300011012-2022131312323132-3001202303113200-2131321121113030-1312131110310332-0120032231133222-1210312011100312-3321113221002200"></a>

#### `primary.rr_set_group.rr_set.lb_record.name` property

Type: `"string"`. Optional.

Load Balancer record name (except for SRV DNS Load balancer record) should be a simple record name
and not a subdomain of a subdomain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
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
    "maxLength": 255,
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
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

- [value](resources--dns_zone--reference--group-002.md#canonical-2223311132322230-3320322312300132-0302230100002202-3110210102103332-0101321332303223-1231303031322303-3020002233313333-0103313232301202): complete subsection reference.

<a id="canonical-2223311132322230-3320322312300132-0302230100002202-3110210102103332-0101321332303223-1231303031322303-3020002233313333-0103313232301202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.lb_record.value` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.lb_record](resources--dns_zone--reference--group-002.md#canonical-2011000133312033-0111302203233320-2010201212131212-2202122323010311-1322123301132002-1213120331011210-3311300321113201-3203231313323203)
- primary.rr_set_group.rr_set.lb_record.value

<a id="canonical-2030002130231132-2000012332213113-1213331012213220-2121101123302313-2103300331310021-3120010212222111-2022230201301023-1002333113300020"></a>

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
value {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102223202010023-0221022110120123-1210012321122013-1202111321322300-1103022102332320-3223110233030302-0003012123132032-3132120000021320"></a>

### Direct properties for `primary.rr_set_group.rr_set.lb_record.value`

<a id="canonical-2130011331331021-0303222032023313-0222032030222131-1133331012200100-0032122011332321-3023221320120133-2130121033230223-0103230120303301"></a>

#### `primary.rr_set_group.rr_set.lb_record.value.name` property

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

<a id="canonical-2112210102230222-2032201010213103-3312310231113002-3130221302303220-2312300113013012-3023233212002331-0230213132012331-1230101230212321"></a>

<a id="canonical-1230322332032130-0122200030123102-0201202013120112-2321111033201120-0102213103013333-2030130211123311-0120122133331021-3021130310010010"></a>

#### `primary.rr_set_group.rr_set.lb_record.value.namespace` property

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

<a id="canonical-1322131212320031-0201330021321131-0313321131201233-1132132212021322-2000322002013013-1130310102311223-1033033120020102-1002010220201202"></a>

<a id="canonical-2010220232033002-1032202230032223-2310233130300200-1110032201323031-1311131211010233-3123102331033102-0302013302132023-3011302013233223"></a>

#### `primary.rr_set_group.rr_set.lb_record.value.tenant` property

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

<a id="canonical-3003222202202221-1033200330333101-1301030003310233-1331212131201210-3302032301321211-0112023010221213-1200003133200302-0012102223210202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.loc_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.loc_record

<a id="canonical-1203123000020330-2113200013101102-1113123013131013-1210020133113223-1223230311323133-1001123200011213-0103212102012311-0322023223102232"></a>

Type: `"object"`. single nested block, Optional.

DNS LOC Record. DNS LOC Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
loc_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3102021332003123-2313113003103232-0103332111211001-3010031113220220-0001002003130212-0010023011000022-2211300322321032-1010110110333100"></a>

### Direct properties for `primary.rr_set_group.rr_set.loc_record`

<a id="canonical-2201011333233103-3121332211033331-2231300010202123-0332212313332201-2113003330013321-2100003331321020-0332101220101331-3102323322321131"></a>

#### `primary.rr_set_group.rr_set.loc_record.name` property

Type: `"string"`. Optional.

LOC Record name, please provide only the specific subdomain or record name without the base domain.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-1323013302112203-3032212023103013-1011322030210112-1133333031121233-3020121122120021-2303321121202030-3222123211002022-1012320101111312): complete subsection reference.

<a id="canonical-1323013302112203-3032212023103013-1011322030210112-1133333031121233-3020121122120021-2303321121202030-3222123211002022-1012320101111312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.loc_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.loc_record](resources--dns_zone--reference--group-002.md#canonical-3003222202202221-1033200330333101-1301030003310233-1331212131201210-3302032301321211-0112023010221213-1200003133200302-0012102223210202)
- primary.rr_set_group.rr_set.loc_record.values

<a id="canonical-3313311200021130-1131002203100230-2031300301002002-3101101210022010-3311213223203013-2322301230200210-3020231212101013-2213131332202110"></a>

Type: `"object"`. list nested block, Optional.

LOC Value. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("altitude",
    "latitude_degree",
    "longitude_degree")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012110001022013-1010001002110313-2111103032323332-0323331223030330-2112033221233312-2211100110203230-2330030110230330-3323322033220230"></a>

### Direct properties for `primary.rr_set_group.rr_set.loc_record.values`

<a id="canonical-0013112313120332-3320132333222132-1331001033001122-0033001311322013-0302211032210011-3131023033300010-3000323211033002-0313030203101130"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.altitude` property

Type: `"number"`. Optional.

Altitude. Altitude in meters.

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
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0302000232331101-3133201211201122-3132233113031112-2323313130202220-3321020231222211-2130033321232001-0012111322120301-0313020120221001"></a>

<a id="canonical-1320131330110320-3211001112212232-2333312131113110-1131011011103030-1212100032313122-0122322320033201-1131331130033113-3013001321102301"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.horizontal_precision` property

Type: `"number"`. Optional.

Horizontal Precision. Horizontal Precision in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-3022310322123313-0311101020222012-1020230200300031-2322110031022310-2102101131200133-0132113323121131-3310021130310013-1123200031113211"></a>

<a id="canonical-3322313010213013-2323111131222303-0033220011221323-3201303210311031-2121122133002100-1312201231133122-0003223310312221-3112122303220013"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.latitude_degree` property

Type: `"number"`. Optional.

Latitude degree, an integer between 0 and 90, including 0 and 90.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 90),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 90,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3310130223121023-2221330010131111-2001312003111220-0232232312132011-1310113100303011-1321212133120332-2030301033020023-3123133000211103"></a>

<a id="canonical-2120301333321012-2131030301031233-0030223310030023-0331303201121200-2321220133112100-0211110200100021-2032110232201101-3330230000200011"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.latitude_hemisphere` property

Type: `"string"`. Optional.

\[Enum: N|S\] Latitude hemisphere can only be N or S - N: North Hemisphere - S: South Hemisphere.
Possible values are \`N\`, \`S\`. Defaults to \`N\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["N","S"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("N",
    "S"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "N",
  "enum": [
    "N",
    "S"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2000212200110013-0000201320333332-0113331020020110-2102202223021330-2132110102021133-1310031021020032-3133313213032313-2330021201020213"></a>

<a id="canonical-3013322230131113-3300110311021302-2223231223000023-0231030233131301-3232011320121122-2311222021202122-3303103202111013-0100010312312032"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.latitude_minute` property

Type: `"number"`. Optional.

Latitude minute, an integer between 0 and 59, including 0 and 59.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 59),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-0321123302220223-1213203131102300-1000112213311010-3302322030020202-0111130331232310-0002103001322223-1112231210023113-1223123312232322"></a>

<a id="canonical-2012222002210100-2310013131021301-3130222032103132-2010223111112111-1101210303320301-3111211302130202-0210100011230213-1323201223310321"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.latitude_second` property

Type: `"number"`. Optional.

Latitude second, an decimal between 0 and 59.999, including 0 and 59.999.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-0000130111101322-2233132231333001-2302120220310101-3312021312221011-0112012322212233-2003103030011131-0030222122223300-2132111303102301"></a>

<a id="canonical-3111010121300333-3023301303030332-0013020301232133-0302231022313132-3020221331330010-2222032120011330-2121001300120023-2232100330313200"></a>

#### `primary.rr_set_group.rr_set.loc_record.values.location_diameter` property

Type: `"number"`. Optional.

Diameter of a sphere enclosing the described entity, in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-1112023302312232-1331003021013200-1021302200332101-2110023122323012-2123311211120002-0332122102330131-3121320123011311-3020232223323122"></a>
