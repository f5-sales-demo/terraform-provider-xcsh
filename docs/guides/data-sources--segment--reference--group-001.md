---
page_title: "xcsh_segment reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_segment reference."
---

# xcsh_segment reference

<a id="canonical-1110023031331322-0210111201230333-1322101312131120-3013310220132303-3223212001102311-1012321231332121-2223100100323021-0213000200231233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_segment](../data-sources/segment.md#canonical-1322202313211133-3232200032023212-3030131113213211-1232130001132103-2232301233230330-2022002332021123-3031321121233232-3101311003121100)
- Property reference

<a id="canonical-2111001210230002-1323101010100020-3012112111101313-1301022221030013-2222023223221302-3100210130101213-3002022013001121-0320300131202231"></a>

### Direct properties for `xcsh_segment`

<a id="canonical-3031121321320122-3322221200110103-2231003320032022-0322000033012020-2201010302320332-1210031331000030-1110231132110000-0010103330113103"></a>

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

<a id="canonical-0230201313012103-3330311331010301-0033000310113212-2321103020031132-1221213110031003-2010320321233021-1030312221203111-0311011232222223"></a>

<a id="canonical-0232200303021320-0230112301200111-0303232002300303-1213111210100310-3220232332213010-2202312102111110-0132211330312330-3000301310333001"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the Segment.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [disable_spec](data-sources--segment--reference--group-001.md#canonical-2101300001211001-2112111023033023-1022321222302201-1003133311123320-2122122230330002-2223331111220223-2010323313120033-2031031033100031): complete subsection reference.

- [enable](data-sources--segment--reference--group-001.md#canonical-1312112111201330-1021022320303231-2103100332202213-3032302001320023-2131210010133312-3320330301231032-2211223033011320-1311210010220212): complete subsection reference.

<a id="canonical-0300111000103132-3003311210001312-0121221221213223-0302232301120221-2211230020123132-2330002231111122-3313022302031202-2311231020323201"></a>

<a id="canonical-0333122221110131-2313123033013210-3231023003223023-3102102000132312-1213230333012011-2331031101112121-1120233203003110-0313002210020112"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2020312213031033-0203223002302121-2102000133331312-1121310230223131-2233000130110213-1201311133133213-1121023022320332-0230320111203020"></a>

<a id="canonical-1232112033330300-3131323100303102-2002201130123230-2120123020103010-0002332210213332-1330031122220100-1011132232210231-3210331012331302"></a>

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

<a id="canonical-3000233121012333-3232033121023300-1303121332123033-1220203012011100-3212220002230312-0011111011233133-2233302203102211-0033033112220133"></a>

<a id="canonical-0221221233233210-1112201120222202-1020123211121133-0302121123321330-3121210131311331-1332011122100131-0332011113012103-3020320212222130"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Segment.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1232201013130023-0332313000222212-3113112010112312-1023012201102022-0002312110011321-2032202123133021-3223021300333031-0102100003230030"></a>

<a id="canonical-2331102210023123-2112101123122231-1010033221101201-3222331231133201-0112110213002101-0221120020030302-0331232210103131-1202131300022020"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace where the Segment exists.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3320221232122010-0113012112223301-3320232112010222-3132200221320101-0012221330233030-0132312101223220-2102133030201000-1112013213313002"></a>

### All schema paths for `xcsh_segment`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--segment--reference--group-001.md#canonical-3031121321320122-3322221200110103-2231003320032022-0322000033012020-2201010302320332-1210031331000030-1110231132110000-0010103330113103) |
| `description` | [description](data-sources--segment--reference--group-001.md#canonical-0230201313012103-3330311331010301-0033000310113212-2321103020031132-1221213110031003-2010320321233021-1030312221203111-0311011232222223) |
| `disable_spec` | [disable_spec](data-sources--segment--reference--group-001.md#canonical-3321201122210102-1203312133320232-1223222331133013-1330301333232003-3013020132021303-0220013011330031-2110010103111020-3331322202203133) |
| `enable` | [enable](data-sources--segment--reference--group-001.md#canonical-3201131110323221-0131211122000222-1132213103223321-3003312110121222-3333300011012330-1001321203220232-0333302320100232-2203132231331113) |
| `id` | [ID](data-sources--segment--reference--group-001.md#canonical-0300111000103132-3003311210001312-0121221221213223-0302232301120221-2211230020123132-2330002231111122-3313022302031202-2311231020323201) |
| `labels` | [labels](data-sources--segment--reference--group-001.md#canonical-2020312213031033-0203223002302121-2102000133331312-1121310230223131-2233000130110213-1201311133133213-1121023022320332-0230320111203020) |
| `name` | [name](data-sources--segment--reference--group-001.md#canonical-3000233121012333-3232033121023300-1303121332123033-1220203012011100-3212220002230312-0011111011233133-2233302203102211-0033033112220133) |
| `namespace` | [namespace](data-sources--segment--reference--group-001.md#canonical-1232201013130023-0332313000222212-3113112010112312-1023012201102022-0002312110011321-2032202123133021-3223021300333031-0102100003230030) |

<a id="canonical-2101300001211001-2112111023033023-1022321222302201-1003133311123320-2122122230330002-2223331111220223-2010323313120033-2031031033100031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_spec` properties

Breadcrumbs:

- [xcsh_segment](../data-sources/segment.md#canonical-1322202313211133-3232200032023212-3030131113213211-1232130001132103-2232301233230330-2022002332021123-3031321121233232-3101311003121100)
- [Property reference](data-sources--segment--reference--group-001.md#canonical-1110023031331322-0210111201230333-1322101312131120-3013310220132303-3223212001102311-1012321231332121-2223100100323021-0213000200231233)
- disable_spec

<a id="canonical-3321201122210102-1203312133320232-1223222331133013-1330301333232003-3013020132021303-0220013011330031-2110010103111020-3331322202203133"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable, enable\] Enable this option

OneOf alternatives in this subsection:

- `disable`
- [enable](data-sources--segment--reference--group-001.md#canonical-3201131110323221-0131211122000222-1132213103223321-3003312110121222-3333300011012330-1001321203220232-0333302320100232-2203132231331113)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312112111201330-1021022320303231-2103100332202213-3032302001320023-2131210010133312-3320330301231032-2211223033011320-1311210010220212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable` properties

Breadcrumbs:

- [xcsh_segment](../data-sources/segment.md#canonical-1322202313211133-3232200032023212-3030131113213211-1232130001132103-2232301233230330-2022002332021123-3031321121233232-3101311003121100)
- [Property reference](data-sources--segment--reference--group-001.md#canonical-1110023031331322-0210111201230333-1322101312131120-3013310220132303-3223212001102311-1012321231332121-2223100100323021-0213000200231233)
- enable

<a id="canonical-3201131110323221-0131211122000222-1132213103223321-3003312110121222-3333300011012330-1001321203220232-0333302320100232-2203132231331113"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.
