---
page_title: "xcsh_data_group reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_group reference."
---

# xcsh_data_group reference

<a id="canonical-2122332000002322-0030233311323023-1021312312333322-3212232000101113-1300221310101123-3101301013233100-3021022110110100-0302132122222001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_data_group](../data-sources/data_group.md#canonical-0222002101202123-1221202102003112-2210120001000231-0023132202013201-3203321232210332-0212210130133110-2010020210021103-2203321203313002)
- Property reference

<a id="canonical-1210221313110310-1223103123112123-2233110310321333-3103231220220332-3133233102102131-0212220011200221-1301233112313332-3011001323313130"></a>

### Direct properties for `xcsh_data_group`

- [address_records](data-sources--data_group--reference--group-001.md#canonical-0000231131200300-2003120323302033-2322221110023130-3302222302303321-3302001123203133-2101100110313300-0212130202203331-3021200300120022): complete subsection reference.

<a id="canonical-3002103000030031-3223101311122101-1310322203303032-2032322000310333-2120033133131121-2131323311221131-0011002230312320-3113002113232200"></a>

<a id="canonical-3100212021211303-1133232112222233-1032202020032213-1311301231013122-3123332110212201-0323322113200322-3112111003313132-1123221223131030"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

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

<a id="canonical-1211000022101030-1012203230010312-1122032211010001-1133300033123210-2230123123022230-1331312001322202-1012333220101312-3100331332231103"></a>

<a id="canonical-0011023222022121-0000221121232311-3301130323012200-3032220303112300-3211101122000300-0101323110002301-2230320131110020-0231310020202031"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the DataGroup.

Additional upstream details:

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0011110301313001-2231232010021301-2320313200322130-0012231112301300-1131123333220022-2110303112320231-0203223111001321-1333101213132212"></a>

<a id="canonical-3031130312211012-2201021020330022-1132310011001232-0211303232112333-3312013202312233-1121220303031210-2301000321121331-3302111031102032"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [integer_records](data-sources--data_group--reference--group-001.md#canonical-3130103021011010-1212133122221030-0133012312320022-2113020102301010-3112302302102112-1211031023130212-1232203010203111-2212113222031213): complete subsection reference.

<a id="canonical-0021200222010222-1021220001320200-3213203011100321-0220200231003001-3130320020100110-2331201223333202-2112231013232131-3003020203230020"></a>

<a id="canonical-3111200011323003-1012302012011111-1130222210321322-3201103033331132-2000121022312101-3210013321113131-0232121132112121-1120123331320233"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-2031013133322330-0311320022330123-1213113030011310-0011231122130320-3231023232003322-2121330231112230-0011210320031021-1111200020211301"></a>

<a id="canonical-0120100131311020-2101012232101233-1210130200200112-1321032003210023-3222202330232011-1002233230111033-1310200111020210-3023022310233220"></a>

#### `name` property

Type: `"string"`. Required.

Name of the DataGroup.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1003101133223023-2310021221120011-0122313032133130-2112132211011031-1012323111101132-0111100321010320-0102210202022300-1223010111323300"></a>

<a id="canonical-2221002320013011-1032300223031001-2033322003312320-1132331313231310-2221200003020221-0230201333323121-2200321301022103-0212100332013010"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the DataGroup exists.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [string_records](data-sources--data_group--reference--group-001.md#canonical-3000201011020031-0300333202003013-1032212333232010-1130110032033211-3013220313313233-3311123021230233-3301233232230021-1302012323321013): complete subsection reference.

<a id="canonical-3130120203232303-3330202021102330-2322221202133103-0332012030120112-3122001023202123-0231103322300313-1233022011321230-3330231211223221"></a>

### All schema paths for `xcsh_data_group`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address_records` | [address_records](data-sources--data_group--reference--group-001.md#canonical-2211323301223212-1132113211301212-0022201030211213-2222321200223012-0201223112131121-1322211020030201-1220130301000010-1021123203012133) |
| `address_records.records` | [address_records.records](data-sources--data_group--reference--group-001.md#canonical-3031110322200022-2122123311323021-3203322102111212-3102123323032030-3133320203131031-3110011030331222-3233221300022233-3000333233232212) |
| `annotations` | [annotations](data-sources--data_group--reference--group-001.md#canonical-3002103000030031-3223101311122101-1310322203303032-2032322000310333-2120033133131121-2131323311221131-0011002230312320-3113002113232200) |
| `description` | [description](data-sources--data_group--reference--group-001.md#canonical-1211000022101030-1012203230010312-1122032211010001-1133300033123210-2230123123022230-1331312001322202-1012333220101312-3100331332231103) |
| `id` | [ID](data-sources--data_group--reference--group-001.md#canonical-0011110301313001-2231232010021301-2320313200322130-0012231112301300-1131123333220022-2110303112320231-0203223111001321-1333101213132212) |
| `integer_records` | [integer_records](data-sources--data_group--reference--group-001.md#canonical-0301002013030332-0123021100000221-2013310123021330-3011120303013003-3200332133010103-0223303321031201-0333303323003323-0122331002232030) |
| `integer_records.records` | [integer_records.records](data-sources--data_group--reference--group-001.md#canonical-3022112111003021-0201013331213032-1011032311130323-2033112013211323-0131331313131011-2132312120223131-0020203032201331-1001330221220100) |
| `labels` | [labels](data-sources--data_group--reference--group-001.md#canonical-0021200222010222-1021220001320200-3213203011100321-0220200231003001-3130320020100110-2331201223333202-2112231013232131-3003020203230020) |
| `name` | [name](data-sources--data_group--reference--group-001.md#canonical-2031013133322330-0311320022330123-1213113030011310-0011231122130320-3231023232003322-2121330231112230-0011210320031021-1111200020211301) |
| `namespace` | [namespace](data-sources--data_group--reference--group-001.md#canonical-1003101133223023-2310021221120011-0122313032133130-2112132211011031-1012323111101132-0111100321010320-0102210202022300-1223010111323300) |
| `string_records` | [string_records](data-sources--data_group--reference--group-001.md#canonical-3031033201110120-1101031030221030-2101112133323313-0302301113210030-1202221031102000-0323202100203030-1123100101222002-0020313211121111) |
| `string_records.records` | [string_records.records](data-sources--data_group--reference--group-001.md#canonical-2301102111303010-0200011212003032-2003130230211312-1001122013311132-2122001102223211-2331303210110110-3030223100022120-2102122130331023) |

<a id="canonical-0000231131200300-2003120323302033-2322221110023130-3302222302303321-3302001123203133-2101100110313300-0212130202203331-3021200300120022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `address_records` properties

Breadcrumbs:

- [xcsh_data_group](../data-sources/data_group.md#canonical-0222002101202123-1221202102003112-2210120001000231-0023132202013201-3203321232210332-0212210130133110-2010020210021103-2203321203313002)
- [Property reference](data-sources--data_group--reference--group-001.md#canonical-2122332000002322-0030233311323023-1021312312333322-3212232000101113-1300221310101123-3101301013233100-3021022110110100-0302132122222001)
- address_records

<a id="canonical-2211323301223212-1132113211301212-0022201030211213-2222321200223012-0201223112131121-1322211020030201-1220130301000010-1021123203012133"></a>

Type: `"single"`. Computed.

\[OneOf: address\_records, integer\_records, string\_records\] Address Record. Data group with
address record List.

Receipt-pinned upstream constraints:

```json
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

- [address_records](data-sources--data_group--reference--group-001.md#canonical-2211323301223212-1132113211301212-0022201030211213-2222321200223012-0201223112131121-1322211020030201-1220130301000010-1021123203012133)
- [integer_records](data-sources--data_group--reference--group-001.md#canonical-0301002013030332-0123021100000221-2013310123021330-3011120303013003-3200332133010103-0223303321031201-0333303323003323-0122331002232030)
- [string_records](data-sources--data_group--reference--group-001.md#canonical-3031033201110120-1101031030221030-2101112133323313-0302301113210030-1202221031102000-0323202100203030-1123100101222002-0020313211121111)

Select alternatives according to the provider validators above.

<a id="canonical-3113131202000222-3230033202331030-3103330223100030-1300120332022101-3020101031203222-1110231110020310-2211301120233221-2303333300101023"></a>

### Direct properties for `address_records`

<a id="canonical-3031110322200022-2122123311323021-3203322102111212-3102123323032030-3133320203131031-3110011030331222-3233221300022233-3000333233232212"></a>

#### `address_records.records` property

Type: `["map", "string"]`. Computed.

Address records. Configuration parameter for records

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 4096
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "format": "ip-address",
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.ip": "true",
      "ves.io.schema.rules.map.max_pairs": "4096",
      "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
    },
    "values": {
      "pattern": "^[0-9a-zA-Z._-]*$",
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
    "ves.io.schema.rules.map.keys.string.ip": "true",
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.ip": "true",
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  }
}
```

<a id="canonical-3130103021011010-1212133122221030-0133012312320022-2113020102301010-3112302302102112-1211031023130212-1232203010203111-2212113222031213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `integer_records` properties

Breadcrumbs:

- [xcsh_data_group](../data-sources/data_group.md#canonical-0222002101202123-1221202102003112-2210120001000231-0023132202013201-3203321232210332-0212210130133110-2010020210021103-2203321203313002)
- [Property reference](data-sources--data_group--reference--group-001.md#canonical-2122332000002322-0030233311323023-1021312312333322-3212232000101113-1300221310101123-3101301013233100-3021022110110100-0302132122222001)
- integer_records

<a id="canonical-0301002013030332-0123021100000221-2013310123021330-3011120303013003-3200332133010103-0223303321031201-0333303323003323-0122331002232030"></a>

Type: `"single"`. Computed.

Configuration parameter for integer records.

Additional upstream details:

Data group with integer record List.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2031002330122322-2212222200233310-0102322133311010-1202212001323233-2303332122322232-3110213013300333-0202021322032100-0011331231113232"></a>

### Direct properties for `integer_records`

<a id="canonical-3022112111003021-0201013331213032-1011032311130323-2033112013211323-0131331313131011-2132312120223131-0020203032201331-1001330221220100"></a>

#### `integer_records.records` property

Type: `["map", "string"]`. Computed.

Integer records. Configuration parameter for records

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 4096
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "4096",
      "ves.io.schema.rules.map.values.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$"
    },
    "values": {
      "pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
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
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$"
  }
}
```

<a id="canonical-3000201011020031-0300333202003013-1032212333232010-1130110032033211-3013220313313233-3311123021230233-3301233232230021-1302012323321013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `string_records` properties

Breadcrumbs:

- [xcsh_data_group](../data-sources/data_group.md#canonical-0222002101202123-1221202102003112-2210120001000231-0023132202013201-3203321232210332-0212210130133110-2010020210021103-2203321203313002)
- [Property reference](data-sources--data_group--reference--group-001.md#canonical-2122332000002322-0030233311323023-1021312312333322-3212232000101113-1300221310101123-3101301013233100-3021022110110100-0302132122222001)
- string_records

<a id="canonical-3031033201110120-1101031030221030-2101112133323313-0302301113210030-1202221031102000-0323202100203030-1123100101222002-0020313211121111"></a>

Type: `"single"`. Computed.

Configuration parameter for string records.

Additional upstream details:

Data group with strings record List.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2013031033210111-0000133220112200-2210212130031012-2121310311003113-2030210212010133-3211221013130111-1121003131231011-0201133021221111"></a>

### Direct properties for `string_records`

<a id="canonical-2301102111303010-0200011212003032-2003130230211312-1001122013311132-2122001102223211-2331303210110110-3030223100022120-2102122130331023"></a>

#### `string_records.records` property

Type: `["map", "string"]`. Computed.

String records. Configuration parameter for records

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 4096
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "4096",
      "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
    },
    "values": {
      "pattern": "^[0-9a-zA-Z._-]*$",
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
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "4096",
    "ves.io.schema.rules.map.values.string.pattern": "^[0-9a-zA-Z._-]*$"
  }
}
```
