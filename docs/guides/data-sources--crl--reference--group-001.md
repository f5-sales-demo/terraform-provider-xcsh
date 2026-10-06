---
page_title: "xcsh_crl reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_crl reference."
---

# xcsh_crl reference

<a id="canonical-3002300001121302-2031123030132100-2121300311223202-1012300002331221-2202200233021133-2322302101223023-1223032000220320-2322010130131013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_crl](../data-sources/crl.md#canonical-1303332111131203-2002232212001132-0030110222210023-0210330330101332-1012330233220233-3212220211023233-3303111232113130-2302300032211003)
- Property reference

<a id="canonical-0213101200300031-1321212121211221-2133322310333320-0320101330032112-0003130133012133-3202112233213112-2333113003120322-1123031100002122"></a>

### Direct properties for `xcsh_crl`

<a id="canonical-0033333021101221-3300002112303133-1201131131313211-2031020221133021-0201132110120120-3021101002132031-0222130023132200-1110131033221331"></a>

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

<a id="canonical-3103312101010323-2100322302101333-3121032012110030-0331001021220020-1011121113302211-3223220213331312-2230302310022033-0123121300131110"></a>

<a id="canonical-3212031002233011-1203121111322102-3210013211012110-0232220221332112-0032110031301200-2301101130331132-1232011321223030-1101220123013323"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the CRL.

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

- [http_access](data-sources--crl--reference--group-001.md#canonical-1022222013321321-2111220010000022-2232313002220321-3121302011213012-0210002333232201-0302202003231231-1033320101311130-3211121130113211): complete subsection reference.

<a id="canonical-1132320213320231-2222310222300021-1231103333322122-2300022301312302-3122113012033311-2300131011300223-1313223122013013-2030102101102320"></a>

<a id="canonical-0311322323031031-2212303011011323-1132122131323102-0213312330213210-3030002311230222-2203302322110132-1113310133330103-2230332331011001"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3113123311132111-2221003213132203-3033202322212103-3001231023113033-3311211322013001-1333021231101201-3130113200030213-2230011233212232"></a>

<a id="canonical-2321102200202202-0022310320123323-0312232221200113-3133333113333002-1003201013312003-3332202311331013-3003021100330311-0230323003112023"></a>

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

<a id="canonical-0230030131230032-1220322130113023-3011102202132123-0130212210213120-3123131111313320-1111221200012220-2332121303132200-0113331100123030"></a>

<a id="canonical-3231000033121210-2000132001301001-0302221313130032-1221222132311201-2213221002001022-3100120232222120-3000222113120202-3220302330213311"></a>

#### `name` property

Type: `"string"`. Required.

Name of the CRL.

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

<a id="canonical-2101201223311232-3213332210102210-2103021103201222-2223132012331122-3312323000303003-0221001231320112-1123121332101131-1110210213230102"></a>

<a id="canonical-0222211022330301-3223302333330132-0320211011112313-2312122011131010-0323102020230033-0211113320321323-0302032322132302-3131131201101110"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the CRL exists.

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

<a id="canonical-1230301301232033-3201113213103201-2331213211310113-0010132112021221-2223312122010122-1100221210303030-3333232311212001-1112011203300023"></a>

<a id="canonical-1031122101230011-3012003133333022-2122122121323011-3220301020230330-0301003022302011-1133232301111012-1323013213203103-2233110032031030"></a>

#### `refresh_interval` property

Type: `"number"`. Computed.

CRL Refresh interval. CRL refresh interval, in hours.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 168,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 6
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "6",
    "ves.io.schema.rules.uint32.lte": "168"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "6",
    "ves.io.schema.rules.uint32.lte": "168"
  }
}
```

<a id="canonical-3232031130120231-0232112121310231-0220131322012221-2333101110023202-0321210103102320-2010123003200103-0233210200020321-2121301110010210"></a>

<a id="canonical-2332322033132211-1231111001022301-0302131011133003-0000102201013222-3300202131211313-2133110132210130-1222103222301310-3221301230231031"></a>

#### `server_address` property

Type: `"string"`. Computed.

CRL Server address. CRL server address or hostname.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
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
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-1131231113002202-1213101322101310-3111031133031121-2301320301120233-1102123001000302-3212323020233210-2021313110223312-3021002333312030"></a>

<a id="canonical-1032033031202032-0210220311022301-0010211122033032-0132103331302200-0302322130332300-0322022200131012-3302200021301121-1011032322022212"></a>

#### `server_port` property

Type: `"number"`. Computed.

CRL Server Port. Set CRL Server port number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3201313113303232-1212013101330210-1030113203210312-2101221020231210-3131200132103101-1100232331122131-2000130210220123-0222322030332200"></a>

<a id="canonical-3120110031222210-3221333021022021-3302021032332301-3312211133113333-0030001310121013-0133332023231003-0102111210220132-2003130132030310"></a>

#### `timeout` property

Type: `"number"`. Computed.

CRL download timeout. CRL download wait time, in seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "180"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "180"
  }
}
```

<a id="canonical-1022012130110032-2031113110020213-0211300021031313-3231000101230103-0033120211233130-0113031222120311-0032200102010011-3120121313002011"></a>

### All schema paths for `xcsh_crl`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--crl--reference--group-001.md#canonical-0033333021101221-3300002112303133-1201131131313211-2031020221133021-0201132110120120-3021101002132031-0222130023132200-1110131033221331) |
| `description` | [description](data-sources--crl--reference--group-001.md#canonical-3103312101010323-2100322302101333-3121032012110030-0331001021220020-1011121113302211-3223220213331312-2230302310022033-0123121300131110) |
| `http_access` | [http_access](data-sources--crl--reference--group-001.md#canonical-1223312300103023-1010123131221102-1132110322023001-0303331102213212-3111211020302233-1303123231113230-2322311331120303-1231022022023120) |
| `http_access.path` | [http_access.path](data-sources--crl--reference--group-001.md#canonical-1001120131221323-3320010000010202-0032213011030121-2211300110113230-3313030313031123-0222322033212023-1203223000330132-2330301202331223) |
| `id` | [ID](data-sources--crl--reference--group-001.md#canonical-1132320213320231-2222310222300021-1231103333322122-2300022301312302-3122113012033311-2300131011300223-1313223122013013-2030102101102320) |
| `labels` | [labels](data-sources--crl--reference--group-001.md#canonical-3113123311132111-2221003213132203-3033202322212103-3001231023113033-3311211322013001-1333021231101201-3130113200030213-2230011233212232) |
| `name` | [name](data-sources--crl--reference--group-001.md#canonical-0230030131230032-1220322130113023-3011102202132123-0130212210213120-3123131111313320-1111221200012220-2332121303132200-0113331100123030) |
| `namespace` | [namespace](data-sources--crl--reference--group-001.md#canonical-2101201223311232-3213332210102210-2103021103201222-2223132012331122-3312323000303003-0221001231320112-1123121332101131-1110210213230102) |
| `refresh_interval` | [refresh_interval](data-sources--crl--reference--group-001.md#canonical-1230301301232033-3201113213103201-2331213211310113-0010132112021221-2223312122010122-1100221210303030-3333232311212001-1112011203300023) |
| `server_address` | [server_address](data-sources--crl--reference--group-001.md#canonical-3232031130120231-0232112121310231-0220131322012221-2333101110023202-0321210103102320-2010123003200103-0233210200020321-2121301110010210) |
| `server_port` | [server_port](data-sources--crl--reference--group-001.md#canonical-1131231113002202-1213101322101310-3111031133031121-2301320301120233-1102123001000302-3212323020233210-2021313110223312-3021002333312030) |
| `timeout` | [timeout](data-sources--crl--reference--group-001.md#canonical-3201313113303232-1212013101330210-1030113203210312-2101221020231210-3131200132103101-1100232331122131-2000130210220123-0222322030332200) |

<a id="canonical-1022222013321321-2111220010000022-2232313002220321-3121302011213012-0210002333232201-0302202003231231-1033320101311130-3211121130113211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_access` properties

Breadcrumbs:

- [xcsh_crl](../data-sources/crl.md#canonical-1303332111131203-2002232212001132-0030110222210023-0210330330101332-1012330233220233-3212220211023233-3303111232113130-2302300032211003)
- [Property reference](data-sources--crl--reference--group-001.md#canonical-3002300001121302-2031123030132100-2121300311223202-1012300002331221-2202200233021133-2322302101223023-1223032000220320-2322010130131013)
- http_access

<a id="canonical-1223312300103023-1010123131221102-1132110322023001-0303331102213212-3111211020302233-1303123231113230-2322311331120303-1231022022023120"></a>

Type: `"single"`. Computed.

Configuration parameter for http access.

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

<a id="canonical-0212021033321222-0003130320122330-1023132210201013-0323002102022203-2333111003120130-0023122321211233-1102022221320101-2210013031121210"></a>

### Direct properties for `http_access`

<a id="canonical-1001120131221323-3320010000010202-0032213011030121-2211300110113230-3313030313031123-0222322033212023-1203223000330132-2330301202331223"></a>

#### `http_access.path` property

Type: `"string"`. Computed.

CRL File path. CRL file location.

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
