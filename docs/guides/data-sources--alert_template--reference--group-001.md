---
page_title: "xcsh_alert_template reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_template reference."
---

# xcsh_alert_template reference

<a id="canonical-3212223031032323-0020230301002003-0212033002000323-0113313121231200-0223110303221321-0303212201023311-3001013320003100-1303010003111200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011313233301203-0330213311010333-3331211331321023-2333000302032030-0113001232032311-3210303020332021-2110231231001222-0102130303332332"></a>

## Property reference — Property reference / 213323212232 / 2

Breadcrumbs:

- [xcsh_alert_template](../data-sources/alert_template.md#canonical-1003131020323200-2102322313132332-1330100113320011-3213311001122330-0122322231213130-1212212131231012-1332020002133101-0232021330332020)
- Property reference

<a id="canonical-1013310331001112-0220201313303202-1111030130202312-0130033311010301-0202232132313313-2012203021003301-0130223133202311-2200300321133031"></a>

## Direct properties — Property reference / 213323212232 / 3

<a id="canonical-2101120201132223-1112003102210033-1001230101033101-0321231110031232-0331321221133332-0210221021223211-2300313121221110-3310310030231310"></a>

<a id="canonical-1322333333210312-0332301333000302-3323010222312110-1011102332110101-1202002211110010-3200003302201312-2011131012030323-3001233223000130"></a>

## alert_message property — Property reference / 213323212232 / 4

Type: `"string"`. Computed.

Alert Message. Alert Message.

Upstream description:

Alert Message.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="canonical-3320020211011223-2303130302033310-0231022310120321-0311201110202121-1112000323010301-0202113310133201-3131130323230021-3332010123223030"></a>

<a id="canonical-2020303311030030-2102000001330223-3222023000202221-1113312201023221-2300212320330030-0002002103102230-2232020220020332-2111320030012002"></a>

## alert_message_details property — Property reference / 213323212232 / 5

Type: `"string"`. Computed.

Alert Message Details. Detailed message of the alert.

Upstream description:

Detailed message of the alert.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2302013201022302-3111131113110000-2001230132320001-0221221320331031-2303232330301123-0101220330220332-1021101132322232-0302023321202311"></a>

<a id="canonical-1210200103022232-0133102311203103-2323033132210321-0333112100200000-2232332011001012-0101212301103300-0221222031010023-0120302031331322"></a>

## alert_name property — Property reference / 213323212232 / 6

Type: `"string"`. Computed.

Alert Name. Alert Name.

Upstream description:

Alert Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 16,
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
    "ves.io.schema.rules.string.max_len": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "16"
  }
}
```

<a id="canonical-2002311120121233-0011301023131312-3300102301220011-2211231303112122-1231331001031002-1011103022331321-1133331133033130-1223233302011011"></a>

<a id="canonical-2331121303031002-3103311031221022-0031202003330313-2300321333131010-0222203222011203-1002223112320120-0201102301321131-1111023003202110"></a>

## annotations property — Property reference / 213323212232 / 7

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

<a id="canonical-0021220110131003-1233232320013100-1302022221021331-0003313000120330-0322112231322331-0123120111330321-2213011330210122-3022113321200220"></a>

<a id="canonical-1303130132131120-0123323110100000-0230200303132012-3200312032120303-1101313322000001-1132011033310103-3032231010202021-3120020331213021"></a>

## description property — Property reference / 213323212232 / 8

Type: `"string"`. Computed.

Description of the AlertTemplate.

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

<a id="canonical-3331031022130211-3132032123121012-0201010030311210-1311131301030323-1123100212103031-0030231123200332-1033201003031111-3010123132321121"></a>

<a id="canonical-2211212230032212-0012102003233031-1200023230311100-0122232323203312-2013333323330130-3101013130223323-3000233030213003-2000301232211111"></a>

## ID property — Property reference / 213323212232 / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0312312233030130-2132200130033121-0330321000031013-2000312112012300-0012003331330001-2311202013010212-0102031133121122-2312232022212022"></a>

<a id="canonical-1331213323023100-1223330332132120-3203031130022222-2032001011123303-2321300300012122-1003100112201300-0332102221230032-1300031102010301"></a>

## labels property — Property reference / 213323212232 / 10

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

<a id="canonical-3022210213131001-2033332221300233-1003120023322211-2310030031003223-2030002333201021-0210200323323221-2001220301332033-1113111133033100"></a>

<a id="canonical-1330012031320231-1001231122311301-3132202003221322-1303021031102122-3020102103200303-0000211002200330-2030133313112101-3121010223002232"></a>

## name property — Property reference / 213323212232 / 11

Type: `"string"`. Required.

Name of the AlertTemplate.

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

<a id="canonical-2321032021101100-0213112101130300-0210301111110122-2323331323321112-2230131030210030-3133013000332223-0121222102000222-3301331211232321"></a>

<a id="canonical-0023232331313020-0331100202320112-1232312233221021-0031033303201322-0012233102133000-2203222323101201-2101032213132212-1333132222100220"></a>

## namespace property — Property reference / 213323212232 / 12

Type: `"string"`. Required.

Namespace where the AlertTemplate exists.

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

<a id="canonical-3003223023211131-3231100300001322-3002133100100030-0221233023011022-2132110301113223-1123230122201330-0300030112313100-3121212032202231"></a>

<a id="canonical-2130330131110011-1233122321213030-0033202133321300-2101322202102133-0303331120013011-1010023220022220-1021001333033323-1313002031020103"></a>

## severity property — Property reference / 213323212232 / 13

Type: `"string"`. Computed.

\[Enum: MINOR|MAJOR|CRITICAL\] List of alert severities Minor Major Critical. Possible values are
\`MINOR\`, \`MAJOR\`, \`CRITICAL\`. Defaults to \`MINOR\`.

Upstream description:

List of alert severities

Minor Major Critical.

Receipt-pinned upstream constraints:

```json
{
  "default": "MINOR",
  "enum": [
    "MINOR",
    "MAJOR",
    "CRITICAL"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1211010121030131-2032030032112013-2310322013013312-1020321221231212-3023232122031202-1103031032133012-2102110213222332-1232010033021012"></a>

## All schema paths — Property reference / 213323212232 / 14

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `alert_message` | [alert_message](data-sources--alert_template--reference--group-001.md#canonical-2101120201132223-1112003102210033-1001230101033101-0321231110031232-0331321221133332-0210221021223211-2300313121221110-3310310030231310) |
| `alert_message_details` | [alert_message_details](data-sources--alert_template--reference--group-001.md#canonical-3320020211011223-2303130302033310-0231022310120321-0311201110202121-1112000323010301-0202113310133201-3131130323230021-3332010123223030) |
| `alert_name` | [alert_name](data-sources--alert_template--reference--group-001.md#canonical-2302013201022302-3111131113110000-2001230132320001-0221221320331031-2303232330301123-0101220330220332-1021101132322232-0302023321202311) |
| `annotations` | [annotations](data-sources--alert_template--reference--group-001.md#canonical-2002311120121233-0011301023131312-3300102301220011-2211231303112122-1231331001031002-1011103022331321-1133331133033130-1223233302011011) |
| `description` | [description](data-sources--alert_template--reference--group-001.md#canonical-0021220110131003-1233232320013100-1302022221021331-0003313000120330-0322112231322331-0123120111330321-2213011330210122-3022113321200220) |
| `id` | [ID](data-sources--alert_template--reference--group-001.md#canonical-3331031022130211-3132032123121012-0201010030311210-1311131301030323-1123100212103031-0030231123200332-1033201003031111-3010123132321121) |
| `labels` | [labels](data-sources--alert_template--reference--group-001.md#canonical-0312312233030130-2132200130033121-0330321000031013-2000312112012300-0012003331330001-2311202013010212-0102031133121122-2312232022212022) |
| `name` | [name](data-sources--alert_template--reference--group-001.md#canonical-3022210213131001-2033332221300233-1003120023322211-2310030031003223-2030002333201021-0210200323323221-2001220301332033-1113111133033100) |
| `namespace` | [namespace](data-sources--alert_template--reference--group-001.md#canonical-2321032021101100-0213112101130300-0210301111110122-2323331323321112-2230131030210030-3133013000332223-0121222102000222-3301331211232321) |
| `severity` | [severity](data-sources--alert_template--reference--group-001.md#canonical-3003223023211131-3231100300001322-3002133100100030-0221233023011022-2132110301113223-1123230122201330-0300030112313100-3121212032202231) |

<a id="canonical-3021002333330111-2313032110330023-1010032223121120-3110111003013213-1333333111030131-0322100322020131-1032031023023003-3013211303231010"></a>

## Next pages — Property reference / 213323212232 / 15

- [xcsh_alert_template](../data-sources/alert_template.md#canonical-1003131020323200-2102322313132332-1330100113320011-3213311001122330-0122322231213130-1212212131231012-1332020002133101-0232021330332020)
