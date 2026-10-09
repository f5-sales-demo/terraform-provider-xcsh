---
page_title: "xcsh_authorization_server reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authorization_server reference."
---

# xcsh_authorization_server reference

<a id="canonical-1033321203032211-3120322231302331-2201330321022031-2230022113210210-1333130133303221-2233013123102231-1302133330012130-3023121200332031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_authorization_server](../data-sources/authorization_server.md#canonical-3021220012103201-0130032102230113-0221231210301022-2323112200300110-3212010013210200-3300200232211023-2113102021313101-1310130133021322)
- Property reference

<a id="canonical-2121230030233013-1323032130010322-2111211233010011-2122032120120303-1110223322201323-3220030231112011-2130002303022202-2312023231111221"></a>

### Direct properties for `xcsh_authorization_server`

<a id="canonical-3310022020202112-3213220033232233-1233333221232323-0323233311223100-0312000031202210-2032121233330021-1230031233210023-2111322000131212"></a>

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

<a id="canonical-2013213100001310-0312123011112130-2000222311000032-2211120011121012-2131221213000001-3100002302102213-1202020220300110-3210201022013222"></a>

<a id="canonical-0121012223333100-1123102211032213-0312130132023132-1121200000032121-3223021323122000-1023113230300030-3202221312031133-0332120210323232"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the AuthorizationServer.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2110132302200331-2012302201203210-2231221232121310-1313210221311002-1232030210233210-0000200000231203-1130203223103233-3112113102003000"></a>

<a id="canonical-3201021320332302-3131302122333233-2220211332203313-0012111230330022-0302022022312120-0231132010302103-0013120301313100-0201301020210201"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3010100021002013-2321112223222110-3221233301133112-3322021023121101-1300132032300120-3223201031112122-3230322321100302-3011232303122000"></a>

<a id="canonical-1120032233332331-3033003033122001-1212013212033231-1200120120302300-0231230203123110-0100022002323121-1200023301312132-0213313323133310"></a>

#### `jwks_uri` property

Type: `"string"`. Computed.

X-textBlockContent: Automatic fetching of JWKS will happen once daily. You can also do it manually
from the list of Authorization Servers at any time.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.uri": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.uri": "true"
  }
}
```

<a id="canonical-3323233303121301-2023021013300200-1232333311321113-3200212200301311-1333232332332321-2102122303302001-2103001031033311-1000222302212210"></a>

<a id="canonical-3202101323013332-3023302323031320-0201333022201100-1112100311313322-2033311333212132-3211210101302013-2212100122330100-2312123330011001"></a>

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

<a id="canonical-1331200023123111-2023201220302003-3001021203303211-1302332010031201-3113331111233312-0312023233111121-0303300211000300-0022120101011132"></a>

<a id="canonical-3312312312130221-3333300031113012-0101221131021312-1223323320101230-0223333201010113-1120232112332133-2223033003200233-1012102330101201"></a>

#### `name` property

Type: `"string"`. Required.

Name of the AuthorizationServer.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1230210102320300-1301011011321131-0203131301210110-0302010202302330-3310101133013301-2112223103113200-1012202003211202-2033130031021032"></a>

<a id="canonical-3201012131322321-3310012232003132-1000121323300000-3003022302101211-2320231031123100-3220002101223322-0320332110330111-0222300213223123"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the AuthorizationServer exists.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2100212123331303-1202130220110030-1200123030223122-1003221013001120-3010210321300322-2323301012311002-3011031023212020-3112302222013301"></a>

### All schema paths for `xcsh_authorization_server`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--authorization_server--reference--group-001.md#canonical-3310022020202112-3213220033232233-1233333221232323-0323233311223100-0312000031202210-2032121233330021-1230031233210023-2111322000131212) |
| `description` | [description](data-sources--authorization_server--reference--group-001.md#canonical-2013213100001310-0312123011112130-2000222311000032-2211120011121012-2131221213000001-3100002302102213-1202020220300110-3210201022013222) |
| `id` | [ID](data-sources--authorization_server--reference--group-001.md#canonical-2110132302200331-2012302201203210-2231221232121310-1313210221311002-1232030210233210-0000200000231203-1130203223103233-3112113102003000) |
| `jwks_uri` | [jwks_uri](data-sources--authorization_server--reference--group-001.md#canonical-3010100021002013-2321112223222110-3221233301133112-3322021023121101-1300132032300120-3223201031112122-3230322321100302-3011232303122000) |
| `labels` | [labels](data-sources--authorization_server--reference--group-001.md#canonical-3323233303121301-2023021013300200-1232333311321113-3200212200301311-1333232332332321-2102122303302001-2103001031033311-1000222302212210) |
| `name` | [name](data-sources--authorization_server--reference--group-001.md#canonical-1331200023123111-2023201220302003-3001021203303211-1302332010031201-3113331111233312-0312023233111121-0303300211000300-0022120101011132) |
| `namespace` | [namespace](data-sources--authorization_server--reference--group-001.md#canonical-1230210102320300-1301011011321131-0203131301210110-0302010202302330-3310101133013301-2112223103113200-1012202003211202-2033130031021032) |
