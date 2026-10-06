---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0300011310133020-1213220302100321-0202203230002231-3202211301302123-1112322110332100-1330031303301101-2230123202021332-1021323102321011"></a>

## `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.path.prefix` property

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

<a id="canonical-2012130221111232-2220212320323022-0201111301031233-1310102101320012-2003130022132223-1232120202000110-3320112202001331-3122200310333133"></a>

<a id="canonical-0312103320012101-3021031301311310-3332221112232333-0032313301332131-2233123120111301-1030113200210201-1130301201323113-3002001022221320"></a>

## `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.path.regex` property

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

<a id="canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules

<a id="canonical-2111232211320012-0121011021031310-3003001111013212-0233323122133211-1313233003100322-3210022221232021-1332322333012331-2302213031303002"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Bot Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130010030230120-0201033123300012-0202202302112212-0120313001021133-0230220310223131-1002011133221130-0121101322002223-1310322131332220"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules`

- [any_domain](resources--http_loadbalancer--reference--group-012.md#canonical-3302323211110130-0110203013301300-3010220320201231-0331123130023032-3201333011023011-2000220011201021-1111210110312033-2201010102300002): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-012.md#canonical-3210233113300130-0110011020100103-3022032333031122-2030033023200330-1132020102012221-0300221030131132-2302312020000213-2302130212021310): complete subsection reference.

<a id="canonical-2332011103221221-0221011211120120-3221021112332030-0003110301323122-3320331330332030-2322100031212013-3230012130103103-2022001333112112"></a>

<a id="canonical-1302230221130312-2321002110302230-0333230121003233-0302211022103103-3233332132231220-3201032312113230-1203030122232300-0210022110330212"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.javascript_location` property

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["AFTER_HEAD","AFTER_TITLE_END","BEFORE_SCRIPT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-012.md#canonical-0203223320013121-3331011111301320-0022103222030001-2121021302203113-2233320132221203-1211113232203231-0023203103000333-0130021320133222): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-012.md#canonical-0222233300033223-0331203312001231-1021001121003333-1013320231220321-2323211102211312-2321101311030333-0213323310233321-0022010023322101): complete subsection reference.

<a id="canonical-3302323211110130-0110203013301300-3010220320201231-0331123130023032-3201333011023011-2000220011201021-1111210110312033-2201010102300002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-012.md#canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.any_domain

<a id="canonical-1321133220010122-2112123203100123-0321321322231113-1330111201310000-0133211313201123-3322210230320330-2332202322101002-2000301113122113"></a>

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

<a id="canonical-3210233113300130-0110011020100103-3022032333031122-2030033023200330-1132020102012221-0300221030131132-2302312020000213-2302130212021310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-012.md#canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.domain

<a id="canonical-2030030301022001-1220112212022232-3332330322101001-3332010030021010-2333302131303013-1110230021110320-2000121010012312-2221201333133002"></a>

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

<a id="canonical-3220202120001011-3203313002211322-2232011033100313-3233130331011221-0321020332011011-2113010012300102-3000330033032310-2320203330002321"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.domain`

<a id="canonical-2321322123203311-2331310131213011-0232323122121201-0010221332123111-1201031212233023-2333323122221102-1313122312330303-0203120001122320"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.domain.exact_value` property

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

<a id="canonical-2321023313333310-3333100000333201-2222010113123102-1323302120100021-2333010212303331-1002120022221133-3100103322020301-2312200203333020"></a>

<a id="canonical-2320230021131320-3312203021000011-3311003303000320-3001230112020320-1221021232001222-1010000203222301-2132333033011200-1300123001210312"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.domain.regex_value` property

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

<a id="canonical-1020013022103212-3223332102230101-0323000312110113-1001003220211030-2102001300031012-0132303001231311-3033120001030130-1302332210322201"></a>

<a id="canonical-3231221023111130-2111300322110302-3323301202330012-3203312332312021-1023310220100032-1203303123311232-1112232122131232-2330002201233231"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.domain.suffix_value` property

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

<a id="canonical-0203223320013121-3331011111301320-0022103222030001-2121021302203113-2233320132221203-1211113232203231-0023203103000333-0130021320133222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-012.md#canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.metadata

<a id="canonical-1033311100102013-3313221313033310-2231312112200301-1111212020112313-2000332100330201-3022110303121330-2000211012313221-2011221013233103"></a>

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

<a id="canonical-0003303013031003-1010032322330323-0112101011201223-2111322103200112-1330101212012003-0212020200221332-0100122113021011-2112233201103133"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.metadata`

<a id="canonical-3232030033201202-0310002331011203-0002233213032032-3013120220222003-3212021023303312-2102103023332133-0131323201223123-0032323123002101"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3010233320133002-0312023131021012-3112222312320120-3313232121030200-2021313302201322-3031123133101112-1333100000300302-3132122001001220"></a>

<a id="canonical-2222323203032111-1301011231131231-0323110313032222-0200133111120112-3312031321032331-3033010103303310-2223022030113013-2221030021320111"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.metadata.name` property

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

<a id="canonical-0222233300033223-0331203312001231-1021001121003333-1013320231220321-2323211102211312-2321101311030333-0213323310233321-0022010023322101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-012.md#canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.path

<a id="canonical-1221030210321211-0331302030330331-3300102113321001-1321310312022211-0301320132203031-3123210332302132-3002231123331332-1020323212103330"></a>

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

<a id="canonical-0111122233202121-3131332103320331-1322023320313312-1033333223013221-0102010101002110-1121302021122132-3220023203301312-2010011223113100"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.path`

<a id="canonical-3202113103023002-2332202323030323-3110112332231303-1200030121001201-2011310131030321-0302321012121213-3301130132311320-1022230200030200"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.path.path` property

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

<a id="canonical-3303202330102231-2032103230311303-2212130231231100-1212210332310100-1133022300000020-2312311013101301-0132121133230113-0202033232112220"></a>

<a id="canonical-3111302101130200-3333332033301020-1311331232022201-3223111001002231-3031103101130021-0213012102233221-2121300001102120-0301203302022023"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.path.prefix` property

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

<a id="canonical-1312031212002001-3301330231230111-1021222303120321-0002303311102331-2002210121202222-0312133333022230-2010010001330221-2102330022302021"></a>

<a id="canonical-2130232202131133-1111112221202110-3303022131112000-3000310111223322-0000123201313223-2330300102303001-3212130220202123-0102210123022300"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules.path.regex` property

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

<a id="canonical-3113113321230312-1011132221021022-2222200122232321-1121313322131023-3010032012223210-0112301112312333-2003213221232012-3022332221010201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.mobile` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- bot_defense_advanced_protection.both_web_and_mobile.mobile

<a id="canonical-0103132030310130-0213012332313021-1120313223301212-1223332013211033-0310310010021311-0033102202320013-0321131020320302-2021022333103322"></a>

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
mobile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012133030103323-0222123320333322-1002020111112000-1201103010033312-2121203000201012-0320230331121302-3011112323230033-0311011111101100"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.mobile`

<a id="canonical-2033002331311212-3102321010021111-1011133002302310-3021213102311132-1112312210112220-0030330112000003-3113122230121021-2113112020301232"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.mobile.name` property

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

<a id="canonical-1232113220032222-1331003211210033-3330231222233212-0033112123202331-0110031133131302-3310323131123312-3300101233003200-0210031032131130"></a>

<a id="canonical-2011023302322323-3113112012031131-3133310321302232-0013100010113002-3113111230031210-3212222123221210-3131321300020120-2221233302133102"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.mobile.namespace` property

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

<a id="canonical-3031000220333222-2032302201313313-2321302001030020-1123302113222231-1322321200032200-1112321000213103-1021000032000011-3203103223313302"></a>

<a id="canonical-2131103100331003-0131310033122102-0123121032221000-2033031310210123-1321123310133300-1133203110221011-2311122001121330-0022203010222220"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.mobile.tenant` property

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

<a id="canonical-0021022102130333-0233002330221112-1311012030002331-0321313122100100-2122112223222023-2200212300000310-2001022021132003-1100033212221121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config

<a id="canonical-0010120112031102-3233302001013301-2313122201220120-1322233310310120-2230331302012223-2011202111101313-3030333121333111-1232233121200312"></a>

Type: `"object"`. single nested block, Optional.

Mobile Request Identifier Headers. Mobile Request Identifier Headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
mobile_sdk_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2332001120133120-3120103102320013-1031101211110202-0001131010222120-3102320111233302-0210222230201210-3301301222130231-0211220210233330"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config`

- [mobile_identifier](resources--http_loadbalancer--reference--group-012.md#canonical-3202310312233203-1200311110012031-2133032021222203-0212130332032121-0003130233333023-0331013200220133-0313300333313322-0100333032003022): complete subsection reference.

<a id="canonical-3202310312233203-1200311110012031-2133032021222203-0212130332032121-0003130233333023-0331013200220133-0313300333313322-0100333032003022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](resources--http_loadbalancer--reference--group-012.md#canonical-0021022102130333-0233002330221112-1311012030002331-0321313122100100-2122112223222023-2200212300000310-2001022021132003-1100033212221121)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier

<a id="canonical-2101113021320121-3110022033123233-0032111223023232-2210320113012310-2020332331210230-0123000003203012-3312233113101323-3133321203103102"></a>

Type: `"object"`. single nested block, Optional.

Mobile Traffic Identifier. Mobile traffic identifier type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
mobile_identifier {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300221030120300-0230330321123011-1303330020100201-2132201320102022-2012321310231112-3222331321331300-2001032110312210-3331020211210312"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier`

- [headers](resources--http_loadbalancer--reference--group-012.md#canonical-3033100203033131-3022001131301000-1001333110101301-1202221220323013-2303123030120131-2313032102231230-3322121122103221-3030132333323120): complete subsection reference.

<a id="canonical-3033100203033131-3022001131301000-1001333110101301-1202221220323013-2303123030120131-2313032102231230-3322121122103221-3030132333323120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](resources--http_loadbalancer--reference--group-012.md#canonical-0021022102130333-0233002330221112-1311012030002331-0321313122100100-2122112223222023-2200212300000310-2001022021132003-1100033212221121)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-012.md#canonical-3202310312233203-1200311110012031-2133032021222203-0212130332032121-0003130233333023-0331013200220133-0313300333313322-0100333032003022)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-1022121230111131-2323023221232310-0030100322210120-1232031123322231-0112230002101231-2222312231000301-2222132021011230-1003121131001221"></a>

Type: `"object"`. list nested block, Optional.

Headers that can be used to identify mobile traffic.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
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

<a id="canonical-2001313203033303-1032223122222311-0113320103023111-1331313132220111-0230221111333111-3020122333003223-0000120300202233-1331110112220022"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers`

- [check_not_present](resources--http_loadbalancer--reference--group-012.md#canonical-3131222020202133-1120203030323000-1033223001030230-3013311033123100-0030131120203303-3131313120131012-1003222231012233-1220112131132230): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-012.md#canonical-3032222213131332-3322000020230202-0323332232103122-1111121320101031-3020033200313011-1011031210211310-3310211100330130-1203112102123111): complete subsection reference.

- [item](resources--http_loadbalancer--reference--group-012.md#canonical-2031320001212222-2121021321103321-1203322320003022-2232221122330301-2301113302230013-1000202333210022-1002213013222023-2133003201221102): complete subsection reference.

<a id="canonical-3032322300122221-0311021301110211-3133013020103120-0202202133310312-3102120322012201-2111330233230300-3311313012120030-0321311033022202"></a>

<a id="canonical-3333222231330212-2301311221203311-0320310330322122-0012131111120211-1032113120203311-0313131103122002-3230113203031223-1102311022132200"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.name` property

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3131222020202133-1120203030323000-1033223001030230-3013311033123100-0030131120203303-3131313120131012-1003222231012233-1220112131132230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](resources--http_loadbalancer--reference--group-012.md#canonical-0021022102130333-0233002330221112-1311012030002331-0321313122100100-2122112223222023-2200212300000310-2001022021132003-1100033212221121)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-012.md#canonical-3202310312233203-1200311110012031-2133032021222203-0212130332032121-0003130233333023-0331013200220133-0313300333313322-0100333032003022)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-012.md#canonical-3033100203033131-3022001131301000-1001333110101301-1202221220323013-2303123030120131-2313032102231230-3322121122103221-3030132333323120)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_not_present

<a id="canonical-2301100120031121-3211210313003122-3332122123102113-3231202120100013-1031212230013122-2021012330021030-3233323123111320-0331133111221313"></a>

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

<a id="canonical-3032222213131332-3322000020230202-0323332232103122-1111121320101031-3020033200313011-1011031210211310-3310211100330130-1203112102123111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](resources--http_loadbalancer--reference--group-012.md#canonical-0021022102130333-0233002330221112-1311012030002331-0321313122100100-2122112223222023-2200212300000310-2001022021132003-1100033212221121)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-012.md#canonical-3202310312233203-1200311110012031-2133032021222203-0212130332032121-0003130233333023-0331013200220133-0313300333313322-0100333032003022)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-012.md#canonical-3033100203033131-3022001131301000-1001333110101301-1202221220323013-2303123030120131-2313032102231230-3322121122103221-3030132333323120)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_present

<a id="canonical-1121230300112211-2002311113201330-1200100203303103-3020112301120213-3001233303222221-2011303203221323-0322311220120002-2333113221012230"></a>

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

<a id="canonical-2031320001212222-2121021321103321-1203322320003022-2232221122330301-2301113302230013-1000202333210022-1002213013222023-2133003201221102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](resources--http_loadbalancer--reference--group-012.md#canonical-0021022102130333-0233002330221112-1311012030002331-0321313122100100-2122112223222023-2200212300000310-2001022021132003-1100033212221121)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-012.md#canonical-3202310312233203-1200311110012031-2133032021222203-0212130332032121-0003130233333023-0331013200220133-0313300333313322-0100333032003022)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-012.md#canonical-3033100203033131-3022001131301000-1001333110101301-1202221220323013-2303123030120131-2313032102231230-3322121122103221-3030132333323120)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.item

<a id="canonical-1103003211132012-0330120213213331-0322000322320020-3123233203222001-3312102112030030-0223323101112221-1023310313301310-0131012220300233"></a>

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

<a id="canonical-1330302201023213-0320112232111130-3233330220222103-1203222223301011-2130302301000011-1130303021110320-3233010211203003-2230330102211033"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.item`

<a id="canonical-3200033033202021-3013212300121233-1111202232013010-1221230200022230-3201230220313213-1230022300123031-2102022223302111-2231103322022112"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.item.exact_values` property

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

<a id="canonical-3230123320230333-0313031012232200-0313322121332003-0120200031123110-2313210100033301-1111331123021013-3022130332013001-3202330301221033"></a>

<a id="canonical-2211131100210101-3031032032032003-0223330200030203-2102312230121013-0133112323202023-2223022230310133-2200200033032130-3232322323010221"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.item.regex_values` property

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

<a id="canonical-0333031112030320-2313101133320031-1000010123013310-0220030010123211-2321212031230123-0101000203232020-1313102201122210-0023300213122313"></a>

<a id="canonical-2220030003133312-0010231230233033-2030203032223021-1213222122033211-1033311010000310-1100032223201021-1011010113203320-0323331011321011"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.item.transformers` property

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3200000311231133-3323200203131332-0220323231120332-1111202111132333-0131021232301123-1131022130213101-1010110003233102-3213003220120002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.web` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- bot_defense_advanced_protection.both_web_and_mobile.web

<a id="canonical-1323222133101321-1201302222222323-0320011321302213-0223020201232322-3331323001321030-0033222020100003-0122002113001313-3103000023312331"></a>

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
web {
  # Configure direct properties listed below.
}
```

<a id="canonical-1122032020103303-1321101122012132-0202001000032322-0000222323212333-0113212032100130-3222131323101113-1321320231111222-2112303122110013"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.web`

<a id="canonical-0003330231232031-0021200100322001-2000203023310301-0000011103102122-1230231013221322-3311000023031311-0233323011132333-3210112102033331"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.web.name` property

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

<a id="canonical-2301321103020023-3122002030230222-2111223012111220-2312100032131130-2331220131302323-1022312312221312-2132010010203033-0122321011231013"></a>

<a id="canonical-2310121030200102-1002312000302210-2110021013112133-2013021310111031-0230303113020201-3123323322100201-1311331212211010-3131233003010030"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.web.namespace` property

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

<a id="canonical-1212323323313233-3130313020110121-0220021120121030-0130230310202123-1200003121200110-0100220123120021-0132113103331113-0200010102221223"></a>

<a id="canonical-0002001033121100-1332002102313320-3322223311301021-1221110023222001-0333221020023100-1313130301221030-1310001320222231-3311122313233110"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.web.tenant` property

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

<a id="canonical-1313221110001232-1000111131110001-2120333202122021-0101332112201311-3303230010022023-1222103103001210-1000210111303323-2003031202312033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.mobile_only` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- bot_defense_advanced_protection.mobile_only

<a id="canonical-1212113101121233-0310101112200203-1332201203302303-0132333133221113-3303211010210110-1002213112031200-1131010100131313-1331310321312020"></a>

Type: `"object"`. single nested block, Optional.

Mobile. Mobile only configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
mobile_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-2031020132310000-1021320301033000-1020013133001111-2321112233221030-3111332023032000-2300030203220231-2112133121111213-2310300102222013"></a>

### Direct properties for `bot_defense_advanced_protection.mobile_only`

- [mobile](resources--http_loadbalancer--reference--group-012.md#canonical-2302013123321210-2310223331012130-2033001202111121-2211013221120030-0200121213132111-0120112120311120-0233312220213230-0032133101312331): complete subsection reference.

<a id="canonical-2302013123321210-2310223331012130-2033001202111121-2211013221120030-0200121213132111-0120112120311120-0233312220213230-0032133101312331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.mobile_only.mobile` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.mobile_only](resources--http_loadbalancer--reference--group-012.md#canonical-1313221110001232-1000111131110001-2120333202122021-0101332112201311-3303230010022023-1222103103001210-1000210111303323-2003031202312033)
- bot_defense_advanced_protection.mobile_only.mobile

<a id="canonical-0232130003231201-3220210130102201-2122111202221120-1031030312020333-1021022130320200-1201000133223211-1310113312330121-0312213333201122"></a>

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
mobile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200002110210112-2223133333230001-0321112311120101-2211233031001021-1023110031102010-2202310132023022-0321311332213331-0211121332111233"></a>

### Direct properties for `bot_defense_advanced_protection.mobile_only.mobile`

<a id="canonical-3000211301121102-3000223202232130-2223333303021211-2130332100303313-3322121231130023-0123201133231133-3330230133010232-0133312123020332"></a>

#### `bot_defense_advanced_protection.mobile_only.mobile.name` property

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

<a id="canonical-0321303331220312-3111212010130023-1233022303110231-1233123213322333-3031202030320120-1002103220302122-0010232003013311-1210130223312323"></a>

<a id="canonical-1220110330332323-0022100313222133-1303023221001132-3322010202210333-0101010013330033-0002200003322313-3323231013320232-0320323020031132"></a>

#### `bot_defense_advanced_protection.mobile_only.mobile.namespace` property

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

<a id="canonical-2302000202020022-1203001333221302-3113102203210333-1333312313113021-3203001020203121-0033003332103031-2103132311101332-2220133223221332"></a>

<a id="canonical-3131113222302332-0032113331203101-2012213333232132-3230112002331100-3011232230221311-2312311130022000-1002210113123333-1320221331023222"></a>

#### `bot_defense_advanced_protection.mobile_only.mobile.tenant` property

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

<a id="canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- bot_defense_advanced_protection.web_only

<a id="canonical-2320021130120100-0133032320321101-0330301303303031-0110300000121012-1313332000311202-2221223013201220-0203123001320033-3001230100033211"></a>

Type: `"object"`. single nested block, Optional.

Web. Web only configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages_except",
    "js_insertion_rules")}
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]"
}
```

Terraform syntax:

```terraform
web_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-3133133323323102-2211100122320210-2310213330222110-3233011100201300-3021002033223233-1011131000323300-0230103122230112-2003312212221313"></a>

### Direct properties for `bot_defense_advanced_protection.web_only`

- [disable_js_insert](resources--http_loadbalancer--reference--group-012.md#canonical-2130220313321211-2102111000221003-0230003023002230-1022110112002202-0210133133123211-2020131103323312-0213032332130032-3121033302113020): complete subsection reference.

- [js_insert_all_pages](resources--http_loadbalancer--reference--group-012.md#canonical-0122213003032102-0300021333330122-1010112231112101-3320330201022123-2110002313133303-0022030322303010-2010202000233202-3012132210200330): complete subsection reference.

- [js_insert_all_pages_except](resources--http_loadbalancer--reference--group-012.md#canonical-2113112203330321-3303103230122333-0233222132021112-1213133102232030-1013130211111132-2221210301022210-2332303022031323-0222200102100113): complete subsection reference.

- [js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301): complete subsection reference.

- [web](resources--http_loadbalancer--reference--group-013.md#canonical-0032332312200101-3012322220030112-1320110202102312-0220201232102020-2202130130221013-0301003122332030-0331121213303131-0133333302230210): complete subsection reference.

<a id="canonical-2130220313321211-2102111000221003-0230003023002230-1022110112002202-0210133133123211-2020131103323312-0213032332130032-3121033302113020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.disable_js_insert` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- bot_defense_advanced_protection.web_only.disable_js_insert

<a id="canonical-1201213102332322-2301233202011023-1131220320302012-3021113231203003-3001303033003301-3300233312120033-1313210111122020-3003131032311201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

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
disable_js_insert = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122213003032102-0300021333330122-1010112231112101-3320330201022123-2110002313133303-0022030322303010-2010202000233202-3012132210200330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insert_all_pages` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- bot_defense_advanced_protection.web_only.js_insert_all_pages

<a id="canonical-0121323202031021-2133322013002231-1212033102002201-3120033301123111-3220232313030223-1220213210310220-0110331010113212-2232131001212033"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
js_insert_all_pages {
  # Configure direct properties listed below.
}
```

<a id="canonical-1330013111020313-3212130311213111-0233302100221012-2030212212033123-1122130211000312-2302122322013321-0223332210211122-2103121332220011"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insert_all_pages`

<a id="canonical-2020111132120030-3210133222210330-3302002130132101-3021310330223300-2233010300123001-1130033322111131-3313013031230031-1122010001112123"></a>

#### `bot_defense_advanced_protection.web_only.js_insert_all_pages.javascript_location` property

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["AFTER_HEAD","AFTER_TITLE_END","BEFORE_SCRIPT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2113112203330321-3303103230122333-0233222132021112-1213133102232030-1013130211111132-2221210301022210-2332303022031323-0222200102100113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insert_all_pages_except` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except

<a id="canonical-0321221331303132-0033121320311213-3101202032321212-2132010312203112-1022330322222131-1233012111122113-0220000103022123-3223212003331101"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages with the exceptions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
js_insert_all_pages_except {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132332102301022-2202222321110212-3332121201013230-3332332322200022-0131000002002030-1012321212222311-2021323130132210-3233131012130111"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insert_all_pages_except`

- [exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230): complete subsection reference.

<a id="canonical-2032220013323220-1023102333021201-2321023110122020-2232320221301132-2021303312132103-2231120100322112-1233110223310033-2210213332100232"></a>

<a id="canonical-1333123023033102-2233032221210002-1232030201032211-1313113032103022-1203001132302233-0311103202210100-3010230003212230-0220232200012021"></a>

#### `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.javascript_location` property

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["AFTER_HEAD","AFTER_TITLE_END","BEFORE_SCRIPT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-012.md#canonical-2113112203330321-3303103230122333-0233222132021112-1213133102232030-1013130211111132-2221210301022210-2332303022031323-0222200102100113)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list

<a id="canonical-2332000031113210-1032230132322013-0033100112312022-1003203323123102-1301000112222210-2313211233022221-2212012012230000-2122223313330100"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320023100301330-3023130300132312-3011313000030223-3023020231201130-2301030331230111-0011322111300003-2322330121333011-1222303322031231"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list`

- [any_domain](resources--http_loadbalancer--reference--group-012.md#canonical-1111031202211010-3323001121320012-0130021100321300-2230112001313000-2023113210032031-1221312330230001-1022331132320301-0200112130020321): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-012.md#canonical-2231200103232023-1011223122003131-2200002033011013-0023233101132321-0311331112213033-2132103130103322-3301131033013033-3201203300233202): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-012.md#canonical-1110200221023030-3221021100012023-0013102202120010-1223210220031131-3303123230130212-2123133302300120-1331321123220130-0102201222112033): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-012.md#canonical-2111201011110321-1110000313021120-2210300001203111-0301100200223101-1222232001301202-3300100100301021-0322011020201020-2120000131322203): complete subsection reference.

<a id="canonical-1111031202211010-3323001121320012-0130021100321300-2230112001313000-2023113210032031-1221312330230001-1022331132320301-0200112130020321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-012.md#canonical-2113112203330321-3303103230122333-0233222132021112-1213133102232030-1013130211111132-2221210301022210-2332303022031323-0222200102100113)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-1002313001231332-2323013230112022-1023303122303002-2122112131033101-0100211200012212-0200310021230303-3122003023312103-0110230023312112"></a>

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

<a id="canonical-2231200103232023-1011223122003131-2200002033011013-0023233101132321-0311331112213033-2132103130103322-3301131033013033-3201203300233202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-012.md#canonical-2113112203330321-3303103230122333-0233222132021112-1213133102232030-1013130211111132-2221210301022210-2332303022031323-0222200102100113)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-0212323120121122-3022103010030120-3103200010222322-3201230020200203-2013010131202132-2331221031031133-3102010103011233-1000011323332203"></a>

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

<a id="canonical-1130120131320012-0021212232101303-2011302120101123-2032120110012103-3011300023132120-3303121213320301-0321210301212020-2022200113022203"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.domain`

<a id="canonical-0131100322221233-2011301310203020-0133223312030001-3132110133103130-2133121131212230-0031101213132113-2110321133331230-1322213112310030"></a>

#### `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.domain.exact_value` property

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

<a id="canonical-3112020122022223-2233133132110312-3200322212333133-2211112302100233-3333220201312131-3111210103333130-2212022102203000-3201301302123202"></a>

<a id="canonical-2020002003121312-2031130220302312-0313233320133001-0133302111211213-1010222300131031-1303012011101023-1110003300003031-3300231131111123"></a>

#### `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.domain.regex_value` property

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

<a id="canonical-3320030033111002-1333302131101121-2112113212321210-2003132210332121-2333332100320203-1002231130100313-3333112010002023-2320122313330022"></a>

<a id="canonical-0232303112300221-3100330311000202-2010011103033102-3132331030333302-3232133111322012-2002330213203230-1112301333211101-1032011201233002"></a>

#### `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.domain.suffix_value` property

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

<a id="canonical-1110200221023030-3221021100012023-0013102202120010-1223210220031131-3303123230130212-2123133302300120-1331321123220130-0102201222112033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-012.md#canonical-2113112203330321-3303103230122333-0233222132021112-1213133102232030-1013130211111132-2221210301022210-2332303022031323-0222200102100113)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-0033200220033023-3111222310212132-2210331121000220-1010113010103132-3230132320322123-0131012231203112-0002201130202322-2311310312230032"></a>

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

<a id="canonical-0211032210101113-3010220101332011-2322230312303002-2121111230320021-0023102020101111-3121033010022130-1023003002311300-3203111332013331"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.metadata`

<a id="canonical-3232322131320000-1032203021012030-2012332100030120-1031221113310231-1111030101130130-0001212222200203-2303313122233233-0110131202312032"></a>

#### `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2302212123230300-0123312332032313-3323101233100201-3110333111330213-0201233312123321-0322111112233011-1123223330232011-0222113301303131"></a>

<a id="canonical-3032332233011320-1332303013013103-0302121100300233-1101031033230002-1212200202220333-0301101113003100-2131033330203102-0331220110233230"></a>

#### `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.metadata.name` property

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

<a id="canonical-2111201011110321-1110000313021120-2210300001203111-0301100200223101-1222232001301202-3300100100301021-0322011020201020-2120000131322203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-012.md#canonical-2113112203330321-3303103230122333-0233222132021112-1213133102232030-1013130211111132-2221210301022210-2332303022031323-0222200102100113)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path

<a id="canonical-0303132232102310-1220032203332202-2313100201312221-0331001022211011-2002322013222013-3032123022121302-3331230122202002-0120312321002300"></a>

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

<a id="canonical-2213002011020122-3313213310200010-0300103202103302-3322011210022303-2303033200022231-0322100131332211-0323320131033011-1222322233100222"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path`

<a id="canonical-0322300023320031-3100020322113101-1213210330133220-0010023132200003-0312200100301311-1200110332133300-3011322031003322-1302121330232030"></a>

#### `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path.path` property

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

<a id="canonical-2322302312131112-0310103132310331-3103312002102213-3031031301121210-1330333032132023-3032211000030033-0112030330101122-3000030001002301"></a>

<a id="canonical-1013031020033310-2121331130021121-0230121310332031-1001122122100003-3220302021021323-2333022320110233-0221122203033332-3122111120213011"></a>

#### `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path.prefix` property

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

<a id="canonical-1310212313001131-3122311320332202-2301321310312010-0002333132231211-0213111021232011-0010120310131023-1223132110203003-1130123112003321"></a>

<a id="canonical-0213211200132230-2022030113330201-1102100130000300-3210301211221230-3031302000330212-1300133312230112-0103231312300233-1101122220111130"></a>

#### `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path.regex` property

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

<a id="canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- bot_defense_advanced_protection.web_only.js_insertion_rules

<a id="canonical-2311031300232101-3112032200333120-3211031000121011-1013103222113100-1010311333003000-0212033211010000-1203010210113312-0322312300023132"></a>

Type: `"object"`. single nested block, Optional.

This defines custom JavaScript insertion rules for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130321033320020-0220023032313311-3103121131302201-0323110221231222-0303220230023310-3330222012212113-1110230101011212-2230323303220101"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules`

- [exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320): complete subsection reference.

- [rules](resources--http_loadbalancer--reference--group-013.md#canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303): complete subsection reference.

<a id="canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list

<a id="canonical-3222332222303021-1002220010121332-1201303221300222-1131311220102122-1233010211310233-0123111033120030-0013311323112001-1221113212101013"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2031201022003110-0112021210310102-2021333002312200-3111002113010133-1010221030230211-2331303223211123-0003023120221313-3023112312223333"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list`

- [any_domain](resources--http_loadbalancer--reference--group-012.md#canonical-2011103010011123-3322332303233223-2320021013221023-1013330012013130-2223103000200000-1202032102332203-0122322222221113-0010112130122100): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-012.md#canonical-1131122213030023-0310320003132323-2032213312222222-1230011110032223-2111221023221312-3000012001330103-1133212233103221-2120031232010313): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-012.md#canonical-1011020020300202-2332203021211013-0023222120212201-0101123013013310-1311310301020303-3021200102103310-1103312210021013-3102330112113101): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-013.md#canonical-1112201310023113-1302120212211300-1210321011312300-2032300112230112-1003233022203030-2331212302112213-2313130121113332-2211112020112331): complete subsection reference.

<a id="canonical-2011103010011123-3322332303233223-2320021013221023-1013330012013130-2223103000200000-1202032102332203-0122322222221113-0010112130122100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.any_domain

<a id="canonical-3322213322100020-1021322200323302-2032112201000133-1212300130032130-1023023112210112-0301000211030120-1032223333021313-3112202220313320"></a>

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

<a id="canonical-1131122213030023-0310320003132323-2032213312222222-1230011110032223-2111221023221312-3000012001330103-1133212233103221-2120031232010313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain

<a id="canonical-2031320111230100-1101300032133131-0033320130220000-2200030101013121-0010130133122032-2321113211333101-0323023111311133-1333202311202112"></a>

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

<a id="canonical-1133330223010032-1330303222130022-0030133130120230-1020202132121001-1313123303323230-0331132213332323-1020311302223101-3131112223331332"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain`

<a id="canonical-3303003131321232-0222132233200222-3201013203021332-2202101330230213-2011203330312211-0003133233010033-0030030313320330-2221302312200102"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain.exact_value` property

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

<a id="canonical-0313212120131221-1131320332231033-3112121031021232-1102200113033212-3032100222023121-3223123233031030-0033320011123233-3333122031322033"></a>

<a id="canonical-1231032203012212-0021211122203032-2032102332002201-2122103313313213-2003011333210021-3311100231011130-3031023123012333-2303112302300212"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain.regex_value` property

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

<a id="canonical-2221100332102103-1332302121220322-3111132323111230-1313023132113003-1333132120030010-3230311203320023-2033021111330210-3130322301001232"></a>

<a id="canonical-0223330212032233-3300032233033321-0202331313110012-1100310211012113-0322001222230232-3201220313113123-0231331220133233-0000332332230101"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain.suffix_value` property

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

<a id="canonical-1011020020300202-2332203021211013-0023222120212201-0101123013013310-1311310301020303-3021200102103310-1103312210021013-3102330112113101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata

<a id="canonical-3001023303310003-0101303210201203-1121210003203311-0310233122100213-3330131201103112-3202220311311110-3010312002212303-1010302311330020"></a>

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

<a id="canonical-0020222201221131-2102132332211120-0012313021202202-2331233002230222-1131313133212201-3010210112021020-1233011112321213-0201322031131013"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata`

<a id="canonical-1233003133220230-1232110011332210-1102000310132332-3203013112300213-3032100331311232-2222320323210112-1001123212322211-1310212320100131"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3032003012021333-1302233030133312-0130013312231312-0112013133313011-2102132130022100-3122032333203322-0132021320033300-2223221233003020"></a>
