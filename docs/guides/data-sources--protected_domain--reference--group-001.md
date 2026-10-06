---
page_title: "xcsh_protected_domain reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_domain reference."
---

# xcsh_protected_domain reference

<a id="canonical-0100212211301123-0131201333311010-0332012111101032-1203123003220200-0132003223232101-2021312112301101-0321113310100202-2032033113330123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_protected_domain](../data-sources/protected_domain.md#canonical-0332111003332231-3231333320302220-0013131312310303-3021021201222321-0302322122212110-0022133313313221-1010303303020331-2102022101100033)
- Property reference

<a id="canonical-3111131022031233-2333021133231123-3121122002310122-1030301133101331-0302022130031101-2222003310203033-3023312033312212-2131211232200221"></a>

### Direct properties for `xcsh_protected_domain`

<a id="canonical-2232301303123230-2022112001213103-2031303210300003-0010221102000013-3121010030231203-2213321123011201-1110203201233031-1310201121101031"></a>

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

<a id="canonical-0011102201130230-1002021133000212-3200331030000031-0123233320301333-1021021123000221-1301311331003130-3023201013010231-1210101100310132"></a>

<a id="canonical-1322311302023000-0331330302302113-3200230122332032-1333111001302020-0311133230110231-0311111220220100-3322303101023130-1232211211000010"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the ProtectedDomain.

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

<a id="canonical-1301222310010222-2110311001130212-2110111003020103-2220231300331231-3121022312323103-3313120032103310-1232021333300321-0322321303121122"></a>

<a id="canonical-1333131302030232-0320011030201223-0021330013010013-1112132123231312-2221323023003232-3121220220321013-3002303323103211-3312031200001211"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0213230111010232-2022201220213102-3202121010002002-3010203113100211-0012001022132211-0002320303023202-0031212301200132-2113130322113322"></a>

<a id="canonical-1033012213102210-1200310021201300-0123030031110322-0303210301321110-2212111230222321-3120203223201300-0200311120131003-3333210000232230"></a>

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

<a id="canonical-1312231112330313-2103020111003333-3300310311030222-2112311310010131-0312022013212213-3011111230230231-3222032101211122-2121011220222113"></a>

<a id="canonical-0021312030013131-3112312003312300-2121202211102302-0223330322131130-0000202212011032-3232231303002111-2333233303232232-3101103230231202"></a>

#### `name` property

Type: `"string"`. Required.

Name of the ProtectedDomain.

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

<a id="canonical-1113313330100100-3002230211201111-0120210201303232-1002013031002310-2123313033312310-3302121000003231-3311121021113020-3323103212133000"></a>

<a id="canonical-2101113311320333-0313232330312311-2113330003230231-1011323213022320-3110312312001311-0322303112120002-1022033302323200-3311022200000030"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the ProtectedDomain exists.

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

<a id="canonical-1020332230133120-2130313010012222-1030333332200302-1000331310330201-3230133222332323-0220020212302230-2212130031210123-3113320113021112"></a>

<a id="canonical-0310010112032022-0120112130013131-0313121210103132-3330321021102303-3003022100102303-0322110022310121-2111113220301222-2230200100133021"></a>

#### `protected_domain` property

Type: `"string"`. Computed.

For Client-Side Defense to work on the web pages where you injected the JS, you need to enter the
root domain below. Example: if you are adding Client-Side Defense JS on checkout.example.com, you
should enter example.com here.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.etld_plus_one": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.etld_plus_one": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2211010300232203-1320213300231202-3330133023301111-3000121233211312-3002310213231012-1320222120103200-0022020211022102-3303102013312223"></a>

### All schema paths for `xcsh_protected_domain`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--protected_domain--reference--group-001.md#canonical-2232301303123230-2022112001213103-2031303210300003-0010221102000013-3121010030231203-2213321123011201-1110203201233031-1310201121101031) |
| `description` | [description](data-sources--protected_domain--reference--group-001.md#canonical-0011102201130230-1002021133000212-3200331030000031-0123233320301333-1021021123000221-1301311331003130-3023201013010231-1210101100310132) |
| `id` | [ID](data-sources--protected_domain--reference--group-001.md#canonical-1301222310010222-2110311001130212-2110111003020103-2220231300331231-3121022312323103-3313120032103310-1232021333300321-0322321303121122) |
| `labels` | [labels](data-sources--protected_domain--reference--group-001.md#canonical-0213230111010232-2022201220213102-3202121010002002-3010203113100211-0012001022132211-0002320303023202-0031212301200132-2113130322113322) |
| `name` | [name](data-sources--protected_domain--reference--group-001.md#canonical-1312231112330313-2103020111003333-3300310311030222-2112311310010131-0312022013212213-3011111230230231-3222032101211122-2121011220222113) |
| `namespace` | [namespace](data-sources--protected_domain--reference--group-001.md#canonical-1113313330100100-3002230211201111-0120210201303232-1002013031002310-2123313033312310-3302121000003231-3311121021113020-3323103212133000) |
| `protected_domain` | [protected_domain](data-sources--protected_domain--reference--group-001.md#canonical-1020332230133120-2130313010012222-1030333332200302-1000331310330201-3230133222332323-0220020212302230-2212130031210123-3113320113021112) |
