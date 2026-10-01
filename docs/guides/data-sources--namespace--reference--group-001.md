---
page_title: "xcsh_namespace reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_namespace reference."
---

# xcsh_namespace reference

<a id="canonical-1233101021033332-1222000332100203-2301103231011223-1201002301311201-3021202223213233-0113330233212033-2200103302332302-0032213221333113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223220101222331-3102133001220333-0333130113223102-1123303030103020-1010110200132112-3122213333231301-3132321100200020-0030233101312012"></a>

## Property reference — Property reference / 213233022211 / 2

Breadcrumbs:

- [xcsh_namespace](../data-sources/namespace.md#canonical-2013030003332230-3323112220123010-1133000333310232-3133330332200222-0123031332332100-3300310323320232-0203323000033132-1310310302002112)
- Property reference

<a id="canonical-2100011220001221-3003102223103033-3132100122203031-3013320312002002-3010102232222033-3111133100022011-3012011022001030-0201202131132223"></a>

## Direct properties — Property reference / 213233022211 / 3

<a id="canonical-2221101320222002-0331313011200322-2000223231312321-3211313223032103-2311131302211132-1221202333301303-1100302202003311-1020311200303032"></a>

<a id="canonical-0100023120121203-2023223332233012-0233320112221232-0010013330011330-0311023030233030-3112023021223231-3330223222211220-0320321300123100"></a>

## annotations property — Property reference / 213233022211 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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

<a id="canonical-1023222200302311-0121323002123330-1020022130202131-0110023222313022-2021003120221210-2033233320112213-2202313123333011-3220112311130023"></a>

<a id="canonical-1033320033332222-1111222003023303-1133031002021220-2123021112313113-3123223132232222-1323033303001330-0001100210221313-1303130330000232"></a>

## description property — Property reference / 213233022211 / 5

Type: `"string"`. Computed.

Description of the Namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0012201031203223-0230300132203121-1200111220320112-0020201333011200-1211230310200022-1322122002210010-1112130130331000-2122001323300333"></a>

<a id="canonical-1123110312212320-2131023030112033-2002113232220110-0333020201321302-0210203103003212-1201301331030122-0220001320003331-1211303023013123"></a>

## ID property — Property reference / 213233022211 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3232203101312223-0220212000200302-3121232013230002-2300322123113022-2121211301330133-3320312303020222-2233221323300312-0313313312221030"></a>

<a id="canonical-1122321031031220-0101001120110033-3133331203121101-2233101321312312-1211021020233201-2300222132213001-2112020301023122-2312030202000130"></a>

## labels property — Property reference / 213233022211 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

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

<a id="canonical-2003300101030120-1131033020003000-2010303200032200-1121203312120130-1131002310312331-0212100012321102-1112211330223131-2302022233200320"></a>

<a id="canonical-2320033103111011-1130121011200222-1032102012102031-0000031031131131-2222110200103322-1323112133222203-0223132021010133-0323230132132331"></a>

## name property — Property reference / 213233022211 / 8

Type: `"string"`. Required.

Name of the Namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2232322231130113-1000130121202230-3021200330323323-2030313013113001-3020002013222012-2312330323032213-0313200320122013-1003120313012321"></a>

<a id="canonical-2222321223323121-3321033232301111-0132323020210123-2122310012002011-2133032133212230-2130202111033230-1321130212201023-2212300222001123"></a>

## namespace property — Property reference / 213233022211 / 9

Type: `"string"`. Optional.

Namespaces are tenant-level objects. Omit this argument.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0231032302323310-1000331232020022-2221231020303331-3300201231010310-3212023133310200-1012003131122311-1111110230030130-3312133232033310"></a>

## All schema paths — Property reference / 213233022211 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--namespace--reference--group-001.md#canonical-2221101320222002-0331313011200322-2000223231312321-3211313223032103-2311131302211132-1221202333301303-1100302202003311-1020311200303032) |
| `description` | [description](data-sources--namespace--reference--group-001.md#canonical-1023222200302311-0121323002123330-1020022130202131-0110023222313022-2021003120221210-2033233320112213-2202313123333011-3220112311130023) |
| `id` | [id](data-sources--namespace--reference--group-001.md#canonical-0012201031203223-0230300132203121-1200111220320112-0020201333011200-1211230310200022-1322122002210010-1112130130331000-2122001323300333) |
| `labels` | [labels](data-sources--namespace--reference--group-001.md#canonical-3232203101312223-0220212000200302-3121232013230002-2300322123113022-2121211301330133-3320312303020222-2233221323300312-0313313312221030) |
| `name` | [name](data-sources--namespace--reference--group-001.md#canonical-2003300101030120-1131033020003000-2010303200032200-1121203312120130-1131002310312331-0212100012321102-1112211330223131-2302022233200320) |
| `namespace` | [namespace](data-sources--namespace--reference--group-001.md#canonical-2232322231130113-1000130121202230-3021200330323323-2030313013113001-3020002013222012-2312330323032213-0313200320122013-1003120313012321) |

<a id="canonical-0030021301331210-3012030233213122-0113013200033202-1211320000032221-1000212113320220-3221201312020123-0000322320232031-3020122230311312"></a>

## Next pages — Property reference / 213233022211 / 11

- [xcsh_namespace](../data-sources/namespace.md#canonical-2013030003332230-3323112220123010-1133000333310232-3133330332200222-0123031332332100-3300310323320232-0203323000033132-1310310302002112)
