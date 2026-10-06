---
page_title: "xcsh_irule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_irule reference."
---

# xcsh_irule reference

<a id="canonical-2012220332011020-0110013123030123-1221233222021110-1212001032200110-0201033020011033-3011302200213220-2332100202322100-1123030132312010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_irule](../data-sources/irule.md#canonical-0231101210120303-0011000002331110-3333032201223210-2323130033302300-2222120120302330-2301200112302321-0023231230022212-3231111021211202)
- Property reference

<a id="canonical-2312111213211121-3212213303113200-1320030221220102-2222311232331012-3130031123122212-1003333132103330-3110122333002201-2332300200130031"></a>

### Direct properties for `xcsh_irule`

<a id="canonical-2300332002213320-0100121322333311-2122132011121202-2113012032011300-2120023222120220-0321233212232223-2010223310003311-2113023222212220"></a>

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

<a id="canonical-0101131102121220-1131022320313132-1320130101303011-3203231111301333-2221210220210130-2300300300001310-2032301120120013-1223232330321003"></a>

<a id="canonical-0112323211102202-3202122001201013-1130033001120302-3231231133220200-2112211330011021-1322310222120323-3232120220310202-3211011112320131"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the Irule.

Additional upstream details:

Specify Description for iRule.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2333122111021310-0323313211223223-2330213033011023-3121221202300313-0111022103003020-3220102122030332-3013301211101110-3020021201020303"></a>

<a id="canonical-3012202103031130-0332002031000222-3121231221001101-2222122101012033-2220310001312012-1002210222120233-2220132132223332-0212333213330322"></a>

#### `description_spec` property

Type: `"string"`. Computed.

Description for iRule. Specify Description for iRule.

<a id="canonical-0031301000321332-3311320213233212-1101003102121113-1301102002001203-1112330121332230-3020330203023022-1201031103120201-2322133022120100"></a>

<a id="canonical-3322202311233230-3001210332012102-0320011030012032-2100300310021102-3220301133322100-1300010013310222-0331022013201230-1202311102110310"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3303313032233203-1112213110222100-3300002131311222-0033223030331100-2333300023101033-3322021130221322-0113012200133301-1010311022200123"></a>

<a id="canonical-1111131031010120-3112123102030300-2103011312301000-0303031030021001-3301110213103121-1000013010011302-0332113022210132-3010001330123030"></a>

#### `irule` property

Type: `"string"`. Computed.

www&#46;internal.example.f5.com')\} DNS::drop\} irule content.

Additional upstream details:

www&#46;internal.example.f5.com")\} DNS::drop\} irule content.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 24576,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 24576,
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
    "ves.io.schema.rules.string.max_len": "24576"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "24576"
  }
}
```

<a id="canonical-0021302033112333-2230203300301300-1221021101003213-0021331111022233-0032302013210112-3202330032330001-1232003033211032-3011302330102330"></a>

<a id="canonical-0322332331200222-1100230322130022-0331030232112002-1232111031010030-1023022023222003-2131212110102230-2021222230033012-2012310330300320"></a>

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

<a id="canonical-0113221123031131-0211032202331013-3302012023220221-0213130100223101-2212110132331111-0313130100333101-3112232221023330-3031300330011030"></a>

<a id="canonical-1332131220020202-1203211020102121-1300213332022132-0011032130223021-3323202320003303-2312102121332210-3111003202200032-1333031210321200"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Irule.

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

<a id="canonical-3310113200300222-3303203123023033-2311333012122022-2331312131230103-2103002323321130-3300313222311323-3010201210001133-0233202110213123"></a>

<a id="canonical-2132210222100113-1311220102100231-0031001231122221-2223231003001210-1032102001211013-2003100013312321-1302121221031311-1302332120110013"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Irule exists.

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

<a id="canonical-3013303201022310-2001113310213123-2301210331231213-2111233002031011-3031303111113301-0221001002123311-0011331012132130-1111102133300113"></a>

### All schema paths for `xcsh_irule`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--irule--reference--group-001.md#canonical-2300332002213320-0100121322333311-2122132011121202-2113012032011300-2120023222120220-0321233212232223-2010223310003311-2113023222212220) |
| `description` | [description](data-sources--irule--reference--group-001.md#canonical-0101131102121220-1131022320313132-1320130101303011-3203231111301333-2221210220210130-2300300300001310-2032301120120013-1223232330321003) |
| `description_spec` | [description_spec](data-sources--irule--reference--group-001.md#canonical-2333122111021310-0323313211223223-2330213033011023-3121221202300313-0111022103003020-3220102122030332-3013301211101110-3020021201020303) |
| `id` | [ID](data-sources--irule--reference--group-001.md#canonical-0031301000321332-3311320213233212-1101003102121113-1301102002001203-1112330121332230-3020330203023022-1201031103120201-2322133022120100) |
| `irule` | [irule](data-sources--irule--reference--group-001.md#canonical-3303313032233203-1112213110222100-3300002131311222-0033223030331100-2333300023101033-3322021130221322-0113012200133301-1010311022200123) |
| `labels` | [labels](data-sources--irule--reference--group-001.md#canonical-0021302033112333-2230203300301300-1221021101003213-0021331111022233-0032302013210112-3202330032330001-1232003033211032-3011302330102330) |
| `name` | [name](data-sources--irule--reference--group-001.md#canonical-0113221123031131-0211032202331013-3302012023220221-0213130100223101-2212110132331111-0313130100333101-3112232221023330-3031300330011030) |
| `namespace` | [namespace](data-sources--irule--reference--group-001.md#canonical-3310113200300222-3303203123023033-2311333012122022-2331312131230103-2103002323321130-3300313222311323-3010201210001133-0233202110213123) |
