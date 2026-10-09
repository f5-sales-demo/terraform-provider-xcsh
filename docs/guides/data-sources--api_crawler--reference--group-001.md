---
page_title: "xcsh_api_crawler reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_crawler reference."
---

# xcsh_api_crawler reference

<a id="canonical-0220220322223001-1320313010233203-3300000030200300-0201002302301013-0210212312130310-0300333002303302-1311323021220123-2132303221000023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-3032313323203012-2021301131010010-1211333201122111-1021320230300330-0211003112312030-1200030031330121-3012320020032333-2333221312222211)
- Property reference

<a id="canonical-3002020301212103-3222213312213113-2323231212013103-1021100333231030-0203113301023302-3031031112211022-1322201102032000-1212212310300333"></a>

### Direct properties for `xcsh_api_crawler`

<a id="canonical-0000210022302121-2112003320301222-3132311210310201-0322113212133003-1310112132001033-2120010032011231-2012313121020101-1333002331302010"></a>

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

<a id="canonical-0122000311132013-2031230221310010-2010322103332001-2311303313211030-0201301021301010-2220130201332002-0011111130133233-1213003011233333"></a>

<a id="canonical-1302213121220030-0331022233301302-0322321120033202-3223213223232211-1033100130223210-3013333133223033-1011001021033320-3323212322211030"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the APICrawler.

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

- [domains](data-sources--api_crawler--reference--group-001.md#canonical-3112021300111122-1311200230001132-3332331001022213-2123023231013133-2203120032023202-1313210310120113-3302231313323222-1103212322221010): complete subsection reference.

<a id="canonical-0303030230021221-1311201231331311-3110301003022120-3111331012002111-1120320233333213-0002210100203022-2222211013302302-2221120302032032"></a>

<a id="canonical-0010030011330113-0332122230022301-0133313311100003-3100300202213222-1031303123220033-3332323313020022-3133302033132203-1013222211210130"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2321323123230012-3330210330310002-1320012320113133-0130033001300202-3210210322021333-3223101022310220-0220112201223033-1301012121032100"></a>

<a id="canonical-2021310203320312-2320203312222120-1333023130032122-0123000133112330-3020021223103002-1120000020102110-1321212122131121-2130300110220000"></a>

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

<a id="canonical-3130113210013022-3120032223022131-2000321020321031-0233320120120220-1121200200300222-0202020100200130-0321130132101100-0121320310001233"></a>

<a id="canonical-0231033033123202-0003130113111210-3203020312111203-2222011212211332-2320321313121032-3001023123202010-2011223112130000-1313330022222013"></a>

#### `name` property

Type: `"string"`. Required.

Name of the APICrawler.

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

<a id="canonical-3131013012323232-0203210311000311-3120232013300101-3230213110003023-1122321311320330-2121000231011021-2211230122231000-2310203310322311"></a>

<a id="canonical-1330112201000311-2203220211010301-2302220120323013-3013200300032033-0312333222232213-0302102130020110-2323113322321032-3313230033131330"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the APICrawler exists.

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

<a id="canonical-3011333201112032-3220101023320130-3103101030321130-3003332030331022-1020330122003231-1232133130221002-0330303212223322-0203202012310003"></a>

### All schema paths for `xcsh_api_crawler`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--api_crawler--reference--group-001.md#canonical-0000210022302121-2112003320301222-3132311210310201-0322113212133003-1310112132001033-2120010032011231-2012313121020101-1333002331302010) |
| `description` | [description](data-sources--api_crawler--reference--group-001.md#canonical-0122000311132013-2031230221310010-2010322103332001-2311303313211030-0201301021301010-2220130201332002-0011111130133233-1213003011233333) |
| `domains` | [domains](data-sources--api_crawler--reference--group-001.md#canonical-3131213223332121-0331003330213332-3302303023003323-1012023002021322-3303120200011322-1312321003203322-3223030012221000-0232001211032220) |
| `domains.domain` | [domains.domain](data-sources--api_crawler--reference--group-001.md#canonical-1331220233332203-0221032000012311-1122222212232122-3212322112231133-2111203221312120-3031010130102100-0130003123211013-3103112103102113) |
| `domains.simple_login` | [domains.simple_login](data-sources--api_crawler--reference--group-001.md#canonical-0203113121101133-1221322211102103-0103230323231021-2000013202302002-3010003130023122-2021112201130332-0133203132213121-0030222132030102) |
| `domains.simple_login.password` | [domains.simple_login.password](data-sources--api_crawler--reference--group-001.md#canonical-0110031022131320-2111321212003112-3322303231011330-3100133210011031-3212110032313000-2122222213121213-1322031320013100-2212203310002132) |
| `domains.simple_login.password.blindfold_secret_info` | [domains.simple_login.password.blindfold_secret_info](data-sources--api_crawler--reference--group-001.md#canonical-2323303222110203-2320213030002201-1300130103132213-1110002313213231-2113102301332202-1012200230222303-3311213313201211-0330131002022000) |
| `domains.simple_login.password.blindfold_secret_info.decryption_provider` | [domains.simple_login.password.blindfold_secret_info.decryption_provider](data-sources--api_crawler--reference--group-001.md#canonical-0020313131322033-1321030001002303-0200302200213322-2312213001113322-3222200230332231-3220110221133121-0133130202311221-0210103212201311) |
| `domains.simple_login.password.blindfold_secret_info.location` | [domains.simple_login.password.blindfold_secret_info.location](data-sources--api_crawler--reference--group-001.md#canonical-2321103302212121-1002001001201000-1000013220103133-1201220131033230-1232031121232002-1301201230202121-1020000020123233-3311130123312113) |
| `domains.simple_login.password.blindfold_secret_info.store_provider` | [domains.simple_login.password.blindfold_secret_info.store_provider](data-sources--api_crawler--reference--group-001.md#canonical-3322021133233122-2001232132131321-3000332332110302-0221100301122322-2223302210221120-0111032031303033-2201001032211121-2113333011120022) |
| `domains.simple_login.password.clear_secret_info` | [domains.simple_login.password.clear_secret_info](data-sources--api_crawler--reference--group-001.md#canonical-0000011000200220-2233110320322221-3012303320113123-0001022222122120-3030312031121221-2300320102222300-3312032203013213-3100312032032121) |
| `domains.simple_login.password.clear_secret_info.provider_ref` | [domains.simple_login.password.clear_secret_info.provider_ref](data-sources--api_crawler--reference--group-001.md#canonical-0112131210023300-3100103222232002-1102231332020001-3212010332200131-0112130133303212-2112101132103320-2321130200133331-0302021310111330) |
| `domains.simple_login.password.clear_secret_info.url` | [domains.simple_login.password.clear_secret_info.url](data-sources--api_crawler--reference--group-001.md#canonical-1203310312013210-2231023113331002-1133133012112313-1231003202131220-3131121132023121-1123010220221131-2131221002212122-3320202001212230) |
| `domains.simple_login.user` | [domains.simple_login.user](data-sources--api_crawler--reference--group-001.md#canonical-2200121112301023-3013113201303211-0212011033102221-1021321212002202-1001132003210301-2211323213222123-1102313033202101-2322203310311012) |
| `id` | [ID](data-sources--api_crawler--reference--group-001.md#canonical-0303030230021221-1311201231331311-3110301003022120-3111331012002111-1120320233333213-0002210100203022-2222211013302302-2221120302032032) |
| `labels` | [labels](data-sources--api_crawler--reference--group-001.md#canonical-2321323123230012-3330210330310002-1320012320113133-0130033001300202-3210210322021333-3223101022310220-0220112201223033-1301012121032100) |
| `name` | [name](data-sources--api_crawler--reference--group-001.md#canonical-3130113210013022-3120032223022131-2000321020321031-0233320120120220-1121200200300222-0202020100200130-0321130132101100-0121320310001233) |
| `namespace` | [namespace](data-sources--api_crawler--reference--group-001.md#canonical-3131013012323232-0203210311000311-3120232013300101-3230213110003023-1122321311320330-2121000231011021-2211230122231000-2310203310322311) |

<a id="canonical-3112021300111122-1311200230001132-3332331001022213-2123023231013133-2203120032023202-1313210310120113-3302231313323222-1103212322221010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains` properties

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-3032313323203012-2021301131010010-1211333201122111-1021320230300330-0211003112312030-1200030031330121-3012320020032333-2333221312222211)
- [Property reference](data-sources--api_crawler--reference--group-001.md#canonical-0220220322223001-1320313010233203-3300000030200300-0201002302301013-0210212312130310-0300333002303302-1311323021220123-2132303221000023)
- domains

<a id="canonical-3131213223332121-0331003330213332-3302303023003323-1012023002021322-3303120200011322-1312321003203322-3223030012221000-0232001211032220"></a>

Type: `"list"`. Computed.

API Crawler. API Crawler Configuration.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-3202320200023303-2012103112202213-1313333203213210-3121031312202221-3000130202323132-1112121213113133-2221333223233103-3311103130230021"></a>

### Direct properties for `domains`

<a id="canonical-1331220233332203-0221032000012311-1122222212232122-3212322112231133-2111203221312120-3031010130102100-0130003123211013-3103112103102113"></a>

#### `domains.domain` property

Type: `"string"`. Computed.

Select the domain to execute API Crawling with given credentials.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
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
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](data-sources--api_crawler--reference--group-001.md#canonical-0120231310201203-1120011010012333-3311301110232003-0101122113033103-0010202201301200-1133212333101303-3331202133310203-2001211101310301): complete subsection reference.

<a id="canonical-0120231310201203-1120011010012333-3311301110232003-0101122113033103-0010202201301200-1133212333101303-3331202133310203-2001211101310301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.simple_login` properties

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-3032313323203012-2021301131010010-1211333201122111-1021320230300330-0211003112312030-1200030031330121-3012320020032333-2333221312222211)
- [Property reference](data-sources--api_crawler--reference--group-001.md#canonical-0220220322223001-1320313010233203-3300000030200300-0201002302301013-0210212312130310-0300333002303302-1311323021220123-2132303221000023)
- [domains](data-sources--api_crawler--reference--group-001.md#canonical-3112021300111122-1311200230001132-3332331001022213-2123023231013133-2203120032023202-1313210310120113-3302231313323222-1103212322221010)
- domains.simple_login

<a id="canonical-0203113121101133-1221322211102103-0103230323231021-2000013202302002-3010003130023122-2021112201130332-0133203132213121-0030222132030102"></a>

Type: `"single"`. Computed.

Configuration parameter for simple login.

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

<a id="canonical-3011331132010012-0120310122210002-0323313220332310-3320113230331211-2013131000020232-3102303311313013-0103000001101223-1022102131223231"></a>

### Direct properties for `domains.simple_login`

- [password](data-sources--api_crawler--reference--group-001.md#canonical-1102201101032321-2011133200223223-0230103103000213-3130332132100303-0321113030301111-3221031210122302-1221021320330310-1220220132312302): complete subsection reference.

<a id="canonical-2200121112301023-3013113201303211-0212011033102221-1021321212002202-1001132003210301-2211323213222123-1102313033202101-2322203310311012"></a>

<a id="canonical-3323130300222130-3233020330310311-2230012031222331-1232313222302003-0200321310123312-0202303111222220-0103220222001210-2102022130322303"></a>

#### `domains.simple_login.user` property

Type: `"string"`. Computed.

Enter the username to assign credentials for the selected domain to crawl.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-1102201101032321-2011133200223223-0230103103000213-3130332132100303-0321113030301111-3221031210122302-1221021320330310-1220220132312302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.simple_login.password` properties

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-3032313323203012-2021301131010010-1211333201122111-1021320230300330-0211003112312030-1200030031330121-3012320020032333-2333221312222211)
- [Property reference](data-sources--api_crawler--reference--group-001.md#canonical-0220220322223001-1320313010233203-3300000030200300-0201002302301013-0210212312130310-0300333002303302-1311323021220123-2132303221000023)
- [domains](data-sources--api_crawler--reference--group-001.md#canonical-3112021300111122-1311200230001132-3332331001022213-2123023231013133-2203120032023202-1313210310120113-3302231313323222-1103212322221010)
- [domains.simple_login](data-sources--api_crawler--reference--group-001.md#canonical-0120231310201203-1120011010012333-3311301110232003-0101122113033103-0010202201301200-1133212333101303-3331202133310203-2001211101310301)
- domains.simple_login.password

<a id="canonical-0110031022131320-2111321212003112-3322303231011330-3100133210011031-3212110032313000-2122222213121213-1322031320013100-2212203310002132"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-1220230020310202-0120203330010032-3312103201030323-3032301232303012-1011211023003231-2011122222213213-0011330013201120-0321230201000231"></a>

### Direct properties for `domains.simple_login.password`

- [blindfold_secret_info](data-sources--api_crawler--reference--group-001.md#canonical-1333331002111103-2220222000030111-2023222233111012-1110232131333310-2300000102131230-2111320220303123-2131200203013122-0103221223203110): complete subsection reference.

- [clear_secret_info](data-sources--api_crawler--reference--group-001.md#canonical-0322230302312030-2312120320112312-1302300333332332-1302002330021101-3323331100303033-0200003110113200-0331030113312113-3311312021032233): complete subsection reference.

<a id="canonical-1333331002111103-2220222000030111-2023222233111012-1110232131333310-2300000102131230-2111320220303123-2131200203013122-0103221223203110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.simple_login.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-3032313323203012-2021301131010010-1211333201122111-1021320230300330-0211003112312030-1200030031330121-3012320020032333-2333221312222211)
- [Property reference](data-sources--api_crawler--reference--group-001.md#canonical-0220220322223001-1320313010233203-3300000030200300-0201002302301013-0210212312130310-0300333002303302-1311323021220123-2132303221000023)
- [domains](data-sources--api_crawler--reference--group-001.md#canonical-3112021300111122-1311200230001132-3332331001022213-2123023231013133-2203120032023202-1313210310120113-3302231313323222-1103212322221010)
- [domains.simple_login](data-sources--api_crawler--reference--group-001.md#canonical-0120231310201203-1120011010012333-3311301110232003-0101122113033103-0010202201301200-1133212333101303-3331202133310203-2001211101310301)
- [domains.simple_login.password](data-sources--api_crawler--reference--group-001.md#canonical-1102201101032321-2011133200223223-0230103103000213-3130332132100303-0321113030301111-3221031210122302-1221021320330310-1220220132312302)
- domains.simple_login.password.blindfold_secret_info

<a id="canonical-2323303222110203-2320213030002201-1300130103132213-1110002313213231-2113102301332202-1012200230222303-3311213313201211-0330131002022000"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-1333323030212322-2133301220022201-1232211021330032-0300201230332120-1110021201210113-0330313301013030-3231133200120210-0011123002302322"></a>

### Direct properties for `domains.simple_login.password.blindfold_secret_info`

<a id="canonical-0020313131322033-1321030001002303-0200302200213322-2312213001113322-3222200230332231-3220110221133121-0133130202311221-0210103212201311"></a>

#### `domains.simple_login.password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2321103302212121-1002001001201000-1000013220103133-1201220131033230-1232031121232002-1301201230202121-1020000020123233-3311130123312113"></a>

<a id="canonical-1211111311000010-0302203200000022-1322202203031003-0223023022220131-2233322211212013-0303213232200323-1303121212011013-1303022330000300"></a>

#### `domains.simple_login.password.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3322021133233122-2001232132131321-3000332332110302-0221100301122322-2223302210221120-0111032031303033-2201001032211121-2113333011120022"></a>

<a id="canonical-3211111010201220-3131220230011023-2202212212012023-2300322200011000-2300232213303212-2213112020203212-3123212232103332-1132303022300231"></a>

#### `domains.simple_login.password.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0322230302312030-2312120320112312-1302300333332332-1302002330021101-3323331100303033-0200003110113200-0331030113312113-3311312021032233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.simple_login.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-3032313323203012-2021301131010010-1211333201122111-1021320230300330-0211003112312030-1200030031330121-3012320020032333-2333221312222211)
- [Property reference](data-sources--api_crawler--reference--group-001.md#canonical-0220220322223001-1320313010233203-3300000030200300-0201002302301013-0210212312130310-0300333002303302-1311323021220123-2132303221000023)
- [domains](data-sources--api_crawler--reference--group-001.md#canonical-3112021300111122-1311200230001132-3332331001022213-2123023231013133-2203120032023202-1313210310120113-3302231313323222-1103212322221010)
- [domains.simple_login](data-sources--api_crawler--reference--group-001.md#canonical-0120231310201203-1120011010012333-3311301110232003-0101122113033103-0010202201301200-1133212333101303-3331202133310203-2001211101310301)
- [domains.simple_login.password](data-sources--api_crawler--reference--group-001.md#canonical-1102201101032321-2011133200223223-0230103103000213-3130332132100303-0321113030301111-3221031210122302-1221021320330310-1220220132312302)
- domains.simple_login.password.clear_secret_info

<a id="canonical-0000011000200220-2233110320322221-3012303320113123-0001022222122120-3030312031121221-2300320102222300-3312032203013213-3100312032032121"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-0212012103321233-1303031101132300-1331020002311131-2013330011023021-0002201123102103-0121132331301221-2112000033011213-3123123211113002"></a>

### Direct properties for `domains.simple_login.password.clear_secret_info`

<a id="canonical-0112131210023300-3100103222232002-1102231332020001-3212010332200131-0112130133303212-2112101132103320-2321130200133331-0302021310111330"></a>

#### `domains.simple_login.password.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1203310312013210-2231023113331002-1133133012112313-1231003202131220-3131121132023121-1123010220221131-2131221002212122-3320202001212230"></a>

<a id="canonical-1013222332330312-0112011200112323-1233300103300021-2133311212332210-0222322002320302-0321201112320320-1132133130233120-3332023301300230"></a>

#### `domains.simple_login.password.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```
