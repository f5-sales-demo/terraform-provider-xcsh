---
page_title: "xcsh_bgp_asn_set reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_asn_set reference."
---

# xcsh_bgp_asn_set reference

<a id="canonical-3313133110321332-0200122330231112-0330010031231222-0131121232003313-0202132231312021-2212233212012031-1311232213212101-1310201022103223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_bgp_asn_set](../data-sources/bgp_asn_set.md#canonical-0101132211310203-2233232230121033-0201030033302303-2112033202323033-1200310311200222-2223131321000021-3211120032230011-1102013013320333)
- Property reference

<a id="canonical-2112210102220313-1111101023032132-0200113122212023-2310231111102012-1102133210033133-0230330301201013-0201313002200232-2210000132202011"></a>

### Direct properties for `xcsh_bgp_asn_set`

<a id="canonical-3022132132321021-1230033310122333-2202112332221123-3103233311123121-3300330022102011-3100322312320032-2313101031232003-3232230132130302"></a>

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

<a id="canonical-1100323213213131-3102021031323022-0002021320122022-2323223131202323-1320331331212133-1130123331202310-2320212322130120-0200202201233302"></a>

<a id="canonical-1230230010323021-0103311303333222-2032110123200230-2223002200323303-3233330332233131-2303321330311122-0022321213312211-3202221010123302"></a>

#### `as_numbers` property

Type: `["list", "number"]`. Computed.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create whitelists or
blacklists for use in network policy or service policy.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2020201011102021-3312120200133112-2120012332131000-1231103132132013-0122300012030101-1113303202313032-0212030302022101-2311002012133200"></a>

<a id="canonical-1001023110032100-2320221003002300-0000012200111213-0131203320011110-3010303111232220-3330231201320301-1023002213312301-3333221110001112"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the BGPAsnSet.

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

<a id="canonical-0223023010213002-3012322013222123-2022222010220121-2313200112203032-2000200103021102-1020312030113030-0320332233221112-0301002022123113"></a>

<a id="canonical-1331030231203032-0313220321012302-2312201113330313-3311212022020013-2103220212320133-1223020213332322-3313032003110131-2302312220233112"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1023313233222120-3132000330203023-1130012331300133-0002132332013010-1132311222133120-2101233123003033-2300112312222220-0202333111121133"></a>

<a id="canonical-1313002212321112-3212233111232233-2111000202331101-2321223112232221-0301210212332231-2112032312133112-3001100233030010-2132213001232130"></a>

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

<a id="canonical-1113002302100313-3323101312031033-0131022011313121-1000212301313221-2133003300130100-1331111020202321-1020103133033322-0203032213300132"></a>

<a id="canonical-2133232313223230-3130023232023010-2130021230221110-3312310320013330-3212133033133220-3230103332322330-0200103131330213-0100322010233331"></a>

#### `name` property

Type: `"string"`. Required.

Name of the BGPAsnSet.

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

<a id="canonical-1333313202202133-3010110132222333-0300321010321301-0300322012112110-0202133233222200-3301203103133103-3011113100002002-2300211132011112"></a>

<a id="canonical-3122021112132210-2032301302102031-3302012122031200-0233223212010232-1310022131033223-2112012201212013-1330212310012230-2032201231212303"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the BGPAsnSet exists.

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

<a id="canonical-2101003010003110-3000210331022023-1232022112322101-3211220302232301-3312300121213231-3320023330131312-1202123333002100-3033123011022112"></a>

### All schema paths for `xcsh_bgp_asn_set`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--bgp_asn_set--reference--group-001.md#canonical-3022132132321021-1230033310122333-2202112332221123-3103233311123121-3300330022102011-3100322312320032-2313101031232003-3232230132130302) |
| `as_numbers` | [as_numbers](data-sources--bgp_asn_set--reference--group-001.md#canonical-1100323213213131-3102021031323022-0002021320122022-2323223131202323-1320331331212133-1130123331202310-2320212322130120-0200202201233302) |
| `description` | [description](data-sources--bgp_asn_set--reference--group-001.md#canonical-2020201011102021-3312120200133112-2120012332131000-1231103132132013-0122300012030101-1113303202313032-0212030302022101-2311002012133200) |
| `id` | [ID](data-sources--bgp_asn_set--reference--group-001.md#canonical-0223023010213002-3012322013222123-2022222010220121-2313200112203032-2000200103021102-1020312030113030-0320332233221112-0301002022123113) |
| `labels` | [labels](data-sources--bgp_asn_set--reference--group-001.md#canonical-1023313233222120-3132000330203023-1130012331300133-0002132332013010-1132311222133120-2101233123003033-2300112312222220-0202333111121133) |
| `name` | [name](data-sources--bgp_asn_set--reference--group-001.md#canonical-1113002302100313-3323101312031033-0131022011313121-1000212301313221-2133003300130100-1331111020202321-1020103133033322-0203032213300132) |
| `namespace` | [namespace](data-sources--bgp_asn_set--reference--group-001.md#canonical-1333313202202133-3010110132222333-0300321010321301-0300322012112110-0202133233222200-3301203103133103-3011113100002002-2300211132011112) |
