---
page_title: "xcsh_filter_set reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_filter_set reference."
---

# xcsh_filter_set reference

<a id="canonical-1023200121033312-1222030021033022-2031210010033023-1200303020100031-1303323030030032-3011012223301222-1332302023110212-1031302002022311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-3033322202313333-3021122111030012-3023212212330123-0300102111110132-3312003112301221-1222013003030021-1210202002210210-0310110313101001)
- Property reference

<a id="canonical-3211110010333020-0003103030120100-3120133122121231-1323033010133330-1312032220012113-3133321333002023-3223310123230313-3213210121013021"></a>

### Direct properties for `xcsh_filter_set`

<a id="canonical-3003101322200333-0201203221030101-1302122323022313-2023000121220000-2103022300101223-2222332111210321-3001101310320032-2011313331222023"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-1320223131312210-2011323030111010-2003311323301312-1320213311313230-3021203131302223-2322303320210130-3100031220010302-3003320120103022"></a>

<a id="canonical-0302132322200321-2021233221131301-1013321312003301-2111200333013010-0332111321113212-1032301012101020-3132311321011033-2120222122122213"></a>

#### `context_key` property

Type: `"string"`. Required.

Indexable context key that identifies a page or page type for which the FilterSet is applicable.

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

<a id="canonical-1233013332130113-1100302031232303-2301100112121220-3332332320322102-2123002301032012-2203223213020132-3122220313012132-1103012113302301"></a>

<a id="canonical-3310012133132311-1313333103023133-1012220311200021-1231013312203330-2300103302112021-1300321123001320-1010032212100112-1202221221122013"></a>

#### `description` property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-3030212233123323-2200322233031000-0320301101131303-2101322101311223-2203212022320101-0232200001212202-0112311210222122-3123122103213210"></a>

<a id="canonical-1111033213312103-0130320222022112-1301211122300133-3333033312131203-3211021223310321-3133023013120213-0320023032100022-0303210333121313"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [filter_fields](resources--filter_set--reference--group-001.md#canonical-1202233312203020-2310221132102333-1202000021102222-0221311311322020-1220222202313033-2021111203221223-0120231120131101-2201031233202332): complete subsection reference.

<a id="canonical-3213330033033011-2031022210330301-0203201213033101-2123110033010120-1122312231101112-2122200001302201-3223013032010120-0103120320232231"></a>

<a id="canonical-0103332130230233-2123232203020322-0202210000001300-2320123023202220-2301003313300333-3223003302301332-3121011230312123-0023313100011113"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3221011312301231-1320133231210023-2033023302020312-0202013311203112-0230023122333003-2312321321021210-3202332122003022-2022231100120122"></a>

<a id="canonical-3220320121321212-2313200131103302-1023200103130031-1333211211120020-2033030002110130-3000302321001232-3322023313000022-0132003122013320"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0300202223203321-2120211023233011-2111130120113021-3112002203210231-3323303310102133-0333313100331202-3111120222011033-3311110321121222"></a>

<a id="canonical-0231311211321222-2032320112320233-1311022023010321-0031210002200201-3020202203300302-1203130033231303-3001321022232132-1012023132201033"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Filter Set. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NameValidator(),
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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

<a id="canonical-0313023320301001-2022122211323223-1222322233011101-2313102331330130-0321021322230120-2321130100230223-2202001333023001-2000313112102311"></a>

<a id="canonical-1001023313323302-3331233121132012-2202331032303012-1311110021311110-2101321033031202-2021320220220321-3233230012120033-0332033330031222"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Filter Set is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NamespaceValidator(),
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

- [timeouts](resources--filter_set--reference--group-001.md#canonical-0011101021101013-2313133130213201-3103023032131122-3002021001111101-3210233223100111-0320222302200020-3022031033233103-0101020301213031): complete subsection reference.

<a id="canonical-1000331302010111-2222133231201121-3231332130133021-0102223110320022-3320331033032102-1320303230000101-0101322233132103-3323300131132203"></a>

### All schema paths for `xcsh_filter_set`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--filter_set--reference--group-001.md#canonical-3003101322200333-0201203221030101-1302122323022313-2023000121220000-2103022300101223-2222332111210321-3001101310320032-2011313331222023) |
| `context_key` | [context_key](resources--filter_set--reference--group-001.md#canonical-1320223131312210-2011323030111010-2003311323301312-1320213311313230-3021203131302223-2322303320210130-3100031220010302-3003320120103022) |
| `description` | [description](resources--filter_set--reference--group-001.md#canonical-1233013332130113-1100302031232303-2301100112121220-3332332320322102-2123002301032012-2203223213020132-3122220313012132-1103012113302301) |
| `disable` | [disable](resources--filter_set--reference--group-001.md#canonical-3030212233123323-2200322233031000-0320301101131303-2101322101311223-2203212022320101-0232200001212202-0112311210222122-3123122103213210) |
| `filter_fields` | [filter_fields](resources--filter_set--reference--group-001.md#canonical-3331310130020331-0201232123323322-1223300031033211-0221223030200033-0022013002303232-3213312123332310-3112011231033110-1232202320300301) |
| `filter_fields.date_field` | [filter_fields.date_field](resources--filter_set--reference--group-001.md#canonical-3130122311312120-3113330001033031-3200320233011312-0200123332231223-1003013002220302-1233233121233000-1111130101230330-3020313122000132) |
| `filter_fields.date_field.absolute` | [filter_fields.date_field.absolute](resources--filter_set--reference--group-001.md#canonical-0333102120202031-0312232313113132-1213301100130021-3223010311010101-3222122211211330-1323002131222320-3033331230003012-0030131131030123) |
| `filter_fields.date_field.absolute.end_date` | [filter_fields.date_field.absolute.end_date](resources--filter_set--reference--group-001.md#canonical-1302022301213301-1111010230220013-0201231111323301-2011312233220330-3231332330232013-3131000113032332-2221200010312321-3100202310303223) |
| `filter_fields.date_field.absolute.start_date` | [filter_fields.date_field.absolute.start_date](resources--filter_set--reference--group-001.md#canonical-0011110033310100-3102132313033022-2313222233330211-2231303120320122-3021030201112331-0332310323330010-1122331121112001-0002223200010230) |
| `filter_fields.date_field.relative` | [filter_fields.date_field.relative](resources--filter_set--reference--group-001.md#canonical-3121201301102033-2003130122203200-1010330332210321-2313330123110121-1203220230021101-3011121331111021-0101112010303001-3002312023202110) |
| `filter_fields.field_id` | [filter_fields.field_id](resources--filter_set--reference--group-001.md#canonical-2312031120302310-0303130121031012-0013012331233012-1330301010331020-1311221021203222-3300333322202131-3311333012303332-0331203200103113) |
| `filter_fields.filter_expression_field` | [filter_fields.filter_expression_field](resources--filter_set--reference--group-001.md#canonical-2303211002220303-2021333310320120-0031131002202303-3201131310100313-1230003201023013-2030211212310032-1013210301312003-1102002101203223) |
| `filter_fields.filter_expression_field.expression` | [filter_fields.filter_expression_field.expression](resources--filter_set--reference--group-001.md#canonical-3102010021122101-3121110231300101-0031220012220031-1001201020102031-0112301323011323-3203303100022330-2021121033020310-0130010212313230) |
| `filter_fields.string_field` | [filter_fields.string_field](resources--filter_set--reference--group-001.md#canonical-0201331012101301-0003022032310230-2131010100322233-2222032311003033-3022233333032301-3001121121000320-3323303233030210-1001200122101233) |
| `filter_fields.string_field.field_values` | [filter_fields.string_field.field_values](resources--filter_set--reference--group-001.md#canonical-0102221120230231-0011101323232020-3011303110021201-1120310230101313-0122200320120311-0033001200133012-3033303302302213-0212210212200132) |
| `id` | [ID](resources--filter_set--reference--group-001.md#canonical-3213330033033011-2031022210330301-0203201213033101-2123110033010120-1122312231101112-2122200001302201-3223013032010120-0103120320232231) |
| `labels` | [labels](resources--filter_set--reference--group-001.md#canonical-3221011312301231-1320133231210023-2033023302020312-0202013311203112-0230023122333003-2312321321021210-3202332122003022-2022231100120122) |
| `name` | [name](resources--filter_set--reference--group-001.md#canonical-0300202223203321-2120211023233011-2111130120113021-3112002203210231-3323303310102133-0333313100331202-3111120222011033-3311110321121222) |
| `namespace` | [namespace](resources--filter_set--reference--group-001.md#canonical-0313023320301001-2022122211323223-1222322233011101-2313102331330130-0321021322230120-2321130100230223-2202001333023001-2000313112102311) |
| `timeouts` | [timeouts](resources--filter_set--reference--group-001.md#canonical-3201133020322002-2301300123220201-0113201211212212-2133332122030103-1211032331213200-3010110113312313-2032231300110112-1110010012033303) |
| `timeouts.create` | [timeouts.create](resources--filter_set--reference--group-001.md#canonical-2300213230310131-1200221012130221-0111300313321330-1223201212222103-2301031213233223-1120111331113310-1220113303330231-2211000111312002) |
| `timeouts.delete` | [timeouts.delete](resources--filter_set--reference--group-001.md#canonical-3012232110212023-1030203231203002-2312213300202131-3130310032132211-1122331010013231-0311031301111230-2101030223203121-3233133100100232) |
| `timeouts.read` | [timeouts.read](resources--filter_set--reference--group-001.md#canonical-3323001003030022-1100210212010300-0102231031102220-2322313122020112-3112210200122010-1011101202030003-1030122223233011-3202223232012301) |
| `timeouts.update` | [timeouts.update](resources--filter_set--reference--group-001.md#canonical-0110032012220222-3023102331123112-1320122033303312-2231321101300010-0212103232012211-1023313032113012-0131123100121002-3320020003231112) |

<a id="canonical-1202233312203020-2310221132102333-1202000021102222-0221311311322020-1220222202313033-2021111203221223-0120231120131101-2201031233202332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `filter_fields` properties

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-3033322202313333-3021122111030012-3023212212330123-0300102111110132-3312003112301221-1222013003030021-1210202002210210-0310110313101001)
- [Property reference](resources--filter_set--reference--group-001.md#canonical-1023200121033312-1222030021033022-2031210010033023-1200303020100031-1303323030030032-3011012223301222-1332302023110212-1031302002022311)
- filter_fields

<a id="canonical-3331310130020331-0201232123323322-1223300031033211-0221223030200033-0022013002303232-3213312123332310-3112011231033110-1232202320300301"></a>

Type: `"object"`. list nested block, Optional.

List of fields and their values selected by the user.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("field_id"),
  validators.ConflictingListObjectAttributes("date_field",
    "filter_expression_field"),
  validators.ConflictingListObjectAttributes("date_field",
    "string_field"),
  validators.ConflictingListObjectAttributes("filter_expression_field",
    "string_field")}
```

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
filter_fields {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130021222330331-2222301013220330-2200232231103100-1033320011223310-0212233201331223-2033021232002031-1313113233312121-3312001213222213"></a>

### Direct properties for `filter_fields`

- [date_field](resources--filter_set--reference--group-001.md#canonical-0311201023033312-3031321322112321-3013102220300333-3123311121222133-1332333132233322-0013022102112313-1223331032212111-3221231001222032): complete subsection reference.

<a id="canonical-2312031120302310-0303130121031012-0013012331233012-1330301010331020-1311221021203222-3300333322202131-3311333012303332-0331203200103113"></a>

<a id="canonical-2122230310132133-3132033300123203-3322213233001122-2330111203031032-0123232110320032-0103102212310110-1313203301323022-3323233322212321"></a>

#### `filter_fields.field_id` property

Type: `"string"`. Optional.

An identifier for the field that maps to some UI filter component.

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

- [filter_expression_field](resources--filter_set--reference--group-001.md#canonical-1013212232102003-0023331222000202-3103321201011201-1312112112231332-1102300321031130-3222323231013203-3222002300201330-2013102110132122): complete subsection reference.

- [string_field](resources--filter_set--reference--group-001.md#canonical-0112121032331333-1130312233210221-0212330223212300-2113011231302323-3112202120102200-1033220201313333-2331112121021010-0323232030100213): complete subsection reference.

<a id="canonical-0311201023033312-3031321322112321-3013102220300333-3123311121222133-1332333132233322-0013022102112313-1223331032212111-3221231001222032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `filter_fields.date_field` properties

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-3033322202313333-3021122111030012-3023212212330123-0300102111110132-3312003112301221-1222013003030021-1210202002210210-0310110313101001)
- [Property reference](resources--filter_set--reference--group-001.md#canonical-1023200121033312-1222030021033022-2031210010033023-1200303020100031-1303323030030032-3011012223301222-1332302023110212-1031302002022311)
- [filter_fields](resources--filter_set--reference--group-001.md#canonical-1202233312203020-2310221132102333-1202000021102222-0221311311322020-1220222202313033-2021111203221223-0120231120131101-2201031233202332)
- filter_fields.date_field

<a id="canonical-3130122311312120-3113330001033031-3200320233011312-0200123332231223-1003013002220302-1233233121233000-1111130101230330-3020313122000132"></a>

Type: `"object"`. single nested block, Optional.

Either an absolute time range or a relative time interval.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("absolute",
    "relative")}
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
  "x-ves-oneof-field-range_type": "[\"absolute\",\"relative\"]"
}
```

Terraform syntax:

```terraform
date_field {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032313211301033-1100020032320202-0013213303132130-0003201012302223-1001012201011301-3330000211210110-1322032230200022-3333312133013120"></a>

### Direct properties for `filter_fields.date_field`

- [absolute](resources--filter_set--reference--group-001.md#canonical-0021010123302132-3213302002223301-0231220023101303-0202300022203023-0202032002220103-2203201313310012-3033202331321311-3100221122333303): complete subsection reference.

<a id="canonical-3121201301102033-2003130122203200-1010330332210321-2313330123110121-1203220230021101-3011121331111021-0101112010303001-3002312023202110"></a>

<a id="canonical-1120323312031102-2033330230220023-0123123132023222-1202323301230232-1211112113313220-0210202113232323-3020110011030102-1312223330321312"></a>

#### `filter_fields.date_field.relative` property

Type: `"string"`. Optional.

Exclusive with \[absolute\] relative time duration.

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

<a id="canonical-0021010123302132-3213302002223301-0231220023101303-0202300022203023-0202032002220103-2203201313310012-3033202331321311-3100221122333303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `filter_fields.date_field.absolute` properties

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-3033322202313333-3021122111030012-3023212212330123-0300102111110132-3312003112301221-1222013003030021-1210202002210210-0310110313101001)
- [Property reference](resources--filter_set--reference--group-001.md#canonical-1023200121033312-1222030021033022-2031210010033023-1200303020100031-1303323030030032-3011012223301222-1332302023110212-1031302002022311)
- [filter_fields](resources--filter_set--reference--group-001.md#canonical-1202233312203020-2310221132102333-1202000021102222-0221311311322020-1220222202313033-2021111203221223-0120231120131101-2201031233202332)
- [filter_fields.date_field](resources--filter_set--reference--group-001.md#canonical-0311201023033312-3031321322112321-3013102220300333-3123311121222133-1332333132233322-0013022102112313-1223331032212111-3221231001222032)
- filter_fields.date_field.absolute

<a id="canonical-0333102120202031-0312232313113132-1213301100130021-3223010311010101-3222122211211330-1323002131222320-3033331230003012-0030131131030123"></a>

Type: `"object"`. single nested block, Optional.

Date range is for selecting a date range.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("end_date",
    "start_date")}
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
absolute {
  # Configure direct properties listed below.
}
```

<a id="canonical-1023123123000132-0120331021331300-0020010010311003-0130101032233132-1003101211233133-1030030033302232-2312021223010100-0311210110133212"></a>

### Direct properties for `filter_fields.date_field.absolute`

<a id="canonical-1302022301213301-1111010230220013-0201231111323301-2011312233220330-3231332330232013-3131000113032332-2221200010312321-3100202310303223"></a>

#### `filter_fields.date_field.absolute.end_date` property

Type: `"string"`. Optional.

End Date. Contains end date.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0011110033310100-3102132313033022-2313222233330211-2231303120320122-3021030201112331-0332310323330010-1122331121112001-0002223200010230"></a>

<a id="canonical-0102111331232000-3323323022301110-0223331203130321-1100020233023130-3222230113311112-1122030031302320-1010202202321300-1231020330311200"></a>

#### `filter_fields.date_field.absolute.start_date` property

Type: `"string"`. Optional.

Start Date. Contains start date.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1013212232102003-0023331222000202-3103321201011201-1312112112231332-1102300321031130-3222323231013203-3222002300201330-2013102110132122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `filter_fields.filter_expression_field` properties

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-3033322202313333-3021122111030012-3023212212330123-0300102111110132-3312003112301221-1222013003030021-1210202002210210-0310110313101001)
- [Property reference](resources--filter_set--reference--group-001.md#canonical-1023200121033312-1222030021033022-2031210010033023-1200303020100031-1303323030030032-3011012223301222-1332302023110212-1031302002022311)
- [filter_fields](resources--filter_set--reference--group-001.md#canonical-1202233312203020-2310221132102333-1202000021102222-0221311311322020-1220222202313033-2021111203221223-0120231120131101-2201031233202332)
- filter_fields.filter_expression_field

<a id="canonical-2303211002220303-2021333310320120-0031131002202303-3201131310100313-1230003201023013-2030211212310032-1013210301312003-1102002101203223"></a>

Type: `"object"`. single nested block, Optional.

Filter Expression Field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expression")}
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
filter_expression_field {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233311212023310-1123203212013031-0132010333333313-2111230203300332-1211130003213003-3003123033203212-0331002332221302-2113112020120211"></a>

### Direct properties for `filter_fields.filter_expression_field`

<a id="canonical-3102010021122101-3121110231300101-0031220012220031-1001201020102031-0112301323011323-3203303100022330-2021121033020310-0130010212313230"></a>

#### `filter_fields.filter_expression_field.expression` property

Type: `"string"`. Optional.

Expression is a Kubernetes style label expression for selections, but differs in that it allows
special characters in the keys and values.

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

<a id="canonical-0112121032331333-1130312233210221-0212330223212300-2113011231302323-3112202120102200-1033220201313333-2331112121021010-0323232030100213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `filter_fields.string_field` properties

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-3033322202313333-3021122111030012-3023212212330123-0300102111110132-3312003112301221-1222013003030021-1210202002210210-0310110313101001)
- [Property reference](resources--filter_set--reference--group-001.md#canonical-1023200121033312-1222030021033022-2031210010033023-1200303020100031-1303323030030032-3011012223301222-1332302023110212-1031302002022311)
- [filter_fields](resources--filter_set--reference--group-001.md#canonical-1202233312203020-2310221132102333-1202000021102222-0221311311322020-1220222202313033-2021111203221223-0120231120131101-2201031233202332)
- filter_fields.string_field

<a id="canonical-0201331012101301-0003022032310230-2131010100322233-2222032311003033-3022233333032301-3001121121000320-3323303233030210-1001200122101233"></a>

Type: `"object"`. single nested block, Optional.

Filter String Field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("field_values")}
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
string_field {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220011103020001-3123312230213332-0203101333020332-1020312201100221-0001321212323103-0021000110121021-3313232211021330-1301103323210333"></a>

### Direct properties for `filter_fields.string_field`

<a id="canonical-0102221120230231-0011101323232020-3011303110021201-1120310230101313-0122200320120311-0033001200133012-3033303302302213-0212210212200132"></a>

#### `filter_fields.string_field.field_values` property

Type: `["list", "string"]`. Optional.

String Value(s). Field specification or configuration

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

<a id="canonical-0011101021101013-2313133130213201-3103023032131122-3002021001111101-3210233223100111-0320222302200020-3022031033233103-0101020301213031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-3033322202313333-3021122111030012-3023212212330123-0300102111110132-3312003112301221-1222013003030021-1210202002210210-0310110313101001)
- [Property reference](resources--filter_set--reference--group-001.md#canonical-1023200121033312-1222030021033022-2031210010033023-1200303020100031-1303323030030032-3011012223301222-1332302023110212-1031302002022311)
- timeouts

<a id="canonical-3201133020322002-2301300123220201-0113201211212212-2133332122030103-1211032331213200-3010110113312313-2032231300110112-1110010012033303"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1301223232033010-1232103131203122-3213203110300222-1132011131320012-0131301022122022-1322220101302320-1130330202220220-1223311112133012"></a>

### Direct properties for `timeouts`

<a id="canonical-2300213230310131-1200221012130221-0111300313321330-1223201212222103-2301031213233223-1120111331113310-1220113303330231-2211000111312002"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3012232110212023-1030203231203002-2312213300202131-3130310032132211-1122331010013231-0311031301111230-2101030223203121-3233133100100232"></a>

<a id="canonical-0100303020322112-2212313211212313-3032132302221130-0122220302103300-2112312031010223-1222102113000130-2011323121213222-1032310031322333"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3323001003030022-1100210212010300-0102231031102220-2322313122020112-3112210200122010-1011101202030003-1030122223233011-3202223232012301"></a>

<a id="canonical-0122101200030221-3311310233303211-3013011230103003-1001101213103031-2210020133102233-3200331120200210-0310300320023013-0310301003121313"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0110032012220222-3023102331123112-1320122033303312-2231321101300010-0212103232012211-1023313032113012-0131123100121002-3320020003231112"></a>

<a id="canonical-1033120221100222-1120103333221023-1031102100323221-3130122321011113-2133032133202321-3303003312112032-3232212031313130-3200031003102201"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
