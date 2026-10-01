---
page_title: "xcsh_allowed_domain reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_allowed_domain reference."
---

# xcsh_allowed_domain reference

<a id="canonical-2001103322233010-3101211230020031-2213110101011302-0313200303102033-3001012300210310-2320011103111122-1213233212110011-2230323220112210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033232010333110-3310233200202211-2212013021012230-3221131030001220-2321113113132323-2331301332112131-1000312110220130-3202132322011312"></a>

## Property reference — Property reference / 003013323130 / 2

Breadcrumbs:

- [xcsh_allowed_domain](../data-sources/allowed_domain.md#canonical-3031012320322023-2113120332301001-0010123230112302-0303132031320332-1032231003302322-1130311111131220-1203223310030220-1032110013113220)
- Property reference

<a id="canonical-3203331311023110-1102333111313011-3232223213213330-1221103133323302-1303313120210030-0112231103323011-3202131300011303-0311313103201310"></a>

## Direct properties — Property reference / 003013323130 / 3

<a id="canonical-0122132001311102-0221022000012223-3220301200313320-2333012311003103-3323002330101303-1320120030112030-3001233221013110-0130210310133321"></a>

<a id="canonical-2202203013232212-0223001301223120-1112331210300311-3233111200130021-1223113013133303-2332301311320100-2103023010312213-3101302120230103"></a>

## allowed_domain property — Property reference / 003013323130 / 4

Type: `"string"`. Computed.

Enter root domain or domain to be entered to allow list below. Domains can be entered only one at a
time. In case of conflicting entries, the domain entry takes precedence over the root domain entry.

Upstream description:

Enter root domain or domain to be entered to allow list below. Domains can be entered only one at a
time. In case of conflicting entries, the domain entry takes precedence over the root domain entry.
Example: if you are adding Client-Side Defense JS on checkout.example.com, you should enter
example.com here.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-2130133001202003-3203321120230203-1023131323203300-2123133131321100-0003033323100131-0221322021000103-3202211213011000-3002111003301313"></a>

<a id="canonical-0321013132121221-1300231103301221-1022211203122230-3222021133122123-1122131323030200-0110213313322020-1311220002200123-0230203133231321"></a>

## annotations property — Property reference / 003013323130 / 5

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

<a id="canonical-0121112233223320-3102231222232303-2102320313111011-3300113121210210-0011011033303003-2012103112032302-3203221110120302-0111212311333203"></a>

<a id="canonical-3031321301201331-3122200003110032-2030323310203102-0213230221012310-1230323122012233-3021003102133230-3213303001120130-3033111021130313"></a>

## description property — Property reference / 003013323130 / 6

Type: `"string"`. Computed.

Description of the AllowedDomain.

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

<a id="canonical-2123022310123202-2012013000203220-3223131233220323-3111333220210102-0103302232213032-0321032332030333-1032202210223023-1232123320221011"></a>

<a id="canonical-2002020100311310-0100211213033001-3210210132011320-1130000200120200-2032323200112202-0231332100121101-2113102330003312-1312013032232011"></a>

## ID property — Property reference / 003013323130 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1201132101232132-2120213332301020-0303321230201312-1221013000112031-3303200100200232-3323321332333333-0101002311121320-3111101021023122"></a>

<a id="canonical-3031322020113230-2111210223011131-0011313130002122-1220213322102121-3321102210330223-0100312002011220-2333222220010001-1303200131120311"></a>

## labels property — Property reference / 003013323130 / 8

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

<a id="canonical-3331212003022002-1311031222230221-1122320203222331-1331233223231002-1201221132301313-2103233130200002-2023202201210333-2120013100311113"></a>

<a id="canonical-1012032233303110-0012030303300121-2311121310220110-0201210220202133-1203012103333100-1210102301313003-2012103322101233-3330231020220102"></a>

## name property — Property reference / 003013323130 / 9

Type: `"string"`. Required.

Name of the AllowedDomain.

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

<a id="canonical-2230202122203220-2002332102130101-0320133102310102-2322332031031111-3332322002303300-0103130021212121-1122223022311000-1032001222311233"></a>

<a id="canonical-1203021203000111-2232100022022330-3103120002212332-3110320022211032-3333112233300012-0330302111002112-1021220220113031-1133232133301321"></a>

## namespace property — Property reference / 003013323130 / 10

Type: `"string"`. Required.

Namespace where the AllowedDomain exists.

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

<a id="canonical-2001203332331102-0011212232003103-2113100032222133-2030330132111030-0321010022231221-1133200211031311-2332101303223010-2100113122113312"></a>

## All schema paths — Property reference / 003013323130 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allowed_domain` | [allowed_domain](data-sources--allowed_domain--reference--group-001.md#canonical-0122132001311102-0221022000012223-3220301200313320-2333012311003103-3323002330101303-1320120030112030-3001233221013110-0130210310133321) |
| `annotations` | [annotations](data-sources--allowed_domain--reference--group-001.md#canonical-2130133001202003-3203321120230203-1023131323203300-2123133131321100-0003033323100131-0221322021000103-3202211213011000-3002111003301313) |
| `description` | [description](data-sources--allowed_domain--reference--group-001.md#canonical-0121112233223320-3102231222232303-2102320313111011-3300113121210210-0011011033303003-2012103112032302-3203221110120302-0111212311333203) |
| `id` | [id](data-sources--allowed_domain--reference--group-001.md#canonical-2123022310123202-2012013000203220-3223131233220323-3111333220210102-0103302232213032-0321032332030333-1032202210223023-1232123320221011) |
| `labels` | [labels](data-sources--allowed_domain--reference--group-001.md#canonical-1201132101232132-2120213332301020-0303321230201312-1221013000112031-3303200100200232-3323321332333333-0101002311121320-3111101021023122) |
| `name` | [name](data-sources--allowed_domain--reference--group-001.md#canonical-3331212003022002-1311031222230221-1122320203222331-1331233223231002-1201221132301313-2103233130200002-2023202201210333-2120013100311113) |
| `namespace` | [namespace](data-sources--allowed_domain--reference--group-001.md#canonical-2230202122203220-2002332102130101-0320133102310102-2322332031031111-3332322002303300-0103130021212121-1122223022311000-1032001222311233) |

<a id="canonical-1211311321030313-2132103320313213-2120131221321200-0021210102003321-1000320333133313-3112102330130231-1323103332113330-0013230010231313"></a>

## Next pages — Property reference / 003013323130 / 12

- [xcsh_allowed_domain](../data-sources/allowed_domain.md#canonical-3031012320322023-2113120332301001-0010123230112302-0303132031320332-1032231003302322-1130311111131220-1203223310030220-1032110013113220)
