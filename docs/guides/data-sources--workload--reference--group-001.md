---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- Property reference

<a id="canonical-3230202210333311-1132133213222003-2311023322023202-1210201232021313-2122031223003033-2331030113232221-0121233130211130-3031202320200113"></a>

### Direct properties for `xcsh_workload`

<a id="canonical-3331101321312310-2132331323323130-3322013230311302-1223002323233320-3121102032103302-3113332230303011-0112030330200111-2103100021010322"></a>

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

<a id="canonical-3011323232031133-1123230012010032-1111202101230331-0221213233113231-0121212221133011-0221333231310021-2031000100122213-1330122130101232"></a>

<a id="canonical-3033031123001232-2201322200310110-1211030012323212-2231022223010301-0213030033332230-0222323222210320-0013103002221021-1330113232303012"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the Workload.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1100313110202220-0101323003322323-1101213223200212-0212003302101200-1133011101203200-0213133312013220-0303120031211223-0323223211110232"></a>

<a id="canonical-3113333133131333-3001112112021311-2101333202330111-3133011213210232-2011031331131332-2012313103030310-3001002103122312-2213211210211110"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131): complete subsection reference.

<a id="canonical-3010112001223203-0230322012312002-1022122022230333-1111012101201010-0111110310021300-2100202012311003-2020333123310112-0303110013312221"></a>

<a id="canonical-3033013303123331-3113133212121002-3203030323022231-3221330332230311-1313223202203120-2331201100013102-0313003232223210-1313132303033220"></a>

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

<a id="canonical-1112230021323200-0323233231021132-3002213212023302-0300333330130322-3131032201032210-3011233022032120-0110323330033123-0233131011100133"></a>

<a id="canonical-2132131120322132-3021221003203032-3222030203130121-0022112321123210-3101133001322321-2321022312300312-3301232220302112-1222333201001031"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Workload.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3030311032201033-0033122332232100-2221121122231020-3311013122311022-3132003012021030-2103001112030202-2031203123310121-2311030110220020"></a>

<a id="canonical-3211112202111110-0012120233313201-3232302311021333-3230120023301130-3330112332231331-2220303203023222-2003103103323303-3110323203121303"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Workload exists.

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
  }
}
```

- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211): complete subsection reference.

- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010): complete subsection reference.

- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113): complete subsection reference.
