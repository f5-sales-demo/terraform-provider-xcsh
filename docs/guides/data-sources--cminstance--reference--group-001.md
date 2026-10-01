---
page_title: "xcsh_cminstance reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cminstance reference."
---

# xcsh_cminstance reference

<a id="canonical-3021233103221211-3000330212303311-3123011101210133-2233111111213112-2200323123322113-3330123132032201-0000230310101123-2230201321203202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110210113321012-3033213013330231-2030020201111310-3220232323203230-0111022332333222-1211200203213031-3101221001203223-1233231313303022"></a>

## Property reference — Property reference / 210101111222 / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020)
- Property reference

<a id="canonical-2112100233020103-0011112333111132-3100212233123120-0131131032113130-2100101132210101-2212002102133330-0031003202230102-2002130132310103"></a>

## Direct properties — Property reference / 210101111222 / 3

<a id="canonical-3220131023300001-1223111030101203-2133302311303111-0313321212203023-3102213313030333-1020031113333112-2322300323310031-2012230103113113"></a>

<a id="canonical-0012321302321330-2003112111223022-1213321033012231-0232323023123322-0123230321222332-2120231100301022-2020223202122223-1331013110211002"></a>

## annotations property — Property reference / 210101111222 / 4

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

- [api_token](data-sources--cminstance--reference--group-001.md#canonical-2122011332310321-2012020322322212-2103013120112311-1301000111201000-3210011003323231-1202212322022311-0001312112333103-1011331210231031): complete subsection reference.

<a id="canonical-3213321301001000-2100030022203133-0033111002302233-0101303103002002-3233311223011211-3033312023223320-2230010022121233-1310321011213210"></a>

<a id="canonical-0030112013322023-3322131102211111-0230300100021212-3130002102330023-0320301000030331-0103023113132010-0302333222031211-2321303033201122"></a>

## description property — Property reference / 210101111222 / 5

Type: `"string"`. Computed.

Description of the Cminstance.

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

<a id="canonical-3303022120310213-0020100303332010-2112100221333103-1320003203333020-2011301221202231-1031321031123001-0302022201113310-2111100021303032"></a>

<a id="canonical-1211113321121232-3001221032032013-3223030301121223-1221023132320130-1303302302101322-3133102020330121-3003112003302223-2010212032320311"></a>

## ID property — Property reference / 210101111222 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip](data-sources--cminstance--reference--group-001.md#canonical-0330120303210200-1132021301220102-0312131322232112-1133122030220320-2203003331333221-0221002022303031-2220220300100212-0232002201101111): complete subsection reference.

<a id="canonical-0120112313211110-3102133011313001-1211112030020113-3013200211100202-0230202003313312-1221321200332011-2223111133200203-1313022011021011"></a>

<a id="canonical-3000000001100021-3311213133201033-2120012221312210-1011322112132033-3123220231302031-1022203320020331-3013113200321132-1330112033320131"></a>

## labels property — Property reference / 210101111222 / 7

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

<a id="canonical-0232123221020031-0130321203210310-2012303131331120-3323103000111022-3212101010320321-3210123123302230-1023320213100332-0033223223232202"></a>

<a id="canonical-2013111210320031-3301213213021123-0133213121101112-0130131303230103-3321023331021320-1132311002132101-0111231233130233-0301220111100333"></a>

## name property — Property reference / 210101111222 / 8

Type: `"string"`. Required.

Name of the Cminstance.

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

<a id="canonical-2103212201201323-3223033202121111-1200202110031001-1121120211231011-2012123103110123-3201311023301100-3302311231310102-2023223311202130"></a>

<a id="canonical-3303130213232113-0211230202130010-1020331323210311-1020003000200313-0301211131011133-3131323223323031-3023012220323133-2321201021203123"></a>

## namespace property — Property reference / 210101111222 / 9

Type: `"string"`. Required.

Namespace where the Cminstance exists.

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

- [password](data-sources--cminstance--reference--group-001.md#canonical-1322111331201233-3313303002011133-2023312201301311-0330011002333312-0223110211233200-0010213202031321-1023122120220031-0213320030103121): complete subsection reference.

<a id="canonical-1120001102212231-1231113313022032-0331230030203221-2022303020321010-2200331112022313-2001132211303223-1301002232312123-0312233320323101"></a>

<a id="canonical-2000322032130311-2100111310213310-1311123312010221-2212233131030303-1120000311320310-0202222212223021-2033221122012221-1313102212003330"></a>

## port property — Property reference / 210101111222 / 10

Type: `"number"`. Computed.

Port of the Central Manager instance to connect to.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3231310311131313-0322301201301323-3211002311311220-0333303310321332-2222233301021102-2213002022103313-3001310310310112-2332322101230123"></a>

<a id="canonical-0023330100321322-1200100303100223-1320121302003313-0323023103213031-1112131022320333-1333033311222021-2011033321002303-0322322322013032"></a>

## username property — Property reference / 210101111222 / 11

Type: `"string"`. Computed.

Username for the Central Manager instance.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "4"
  }
}
```

<a id="canonical-0221120103303333-3012100001102100-2113100213023323-1311200330320020-1112331022100013-1200032303013200-2203001320221111-3231103110303030"></a>

## All schema paths — Property reference / 210101111222 / 12

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cminstance--reference--group-001.md#canonical-3220131023300001-1223111030101203-2133302311303111-0313321212203023-3102213313030333-1020031113333112-2322300323310031-2012230103113113) |
| `api_token` | [api_token](data-sources--cminstance--reference--group-001.md#canonical-1233111010230322-0010103131120012-3002200212332133-0201301023022031-2121211203021222-3111120303021021-2232001310100030-1222023330130010) |
| `api_token.blindfold_secret_info` | [api_token.blindfold_secret_info](data-sources--cminstance--reference--group-001.md#canonical-3022031303123012-2331110020123133-3232033102123003-3202222211013122-1203320230033113-2231203200010031-2022233111123100-0032132300220122) |
| `api_token.blindfold_secret_info.decryption_provider` | [api_token.blindfold_secret_info.decryption_provider](data-sources--cminstance--reference--group-001.md#canonical-1022313203121031-0233322103121032-2020230101100120-3220331202213012-0300303323010310-3102332300001312-2013013320030020-1201121003012011) |
| `api_token.blindfold_secret_info.location` | [api_token.blindfold_secret_info.location](data-sources--cminstance--reference--group-001.md#canonical-3020320123303100-0323033010311231-3030212313320012-0223122102120021-2003211001112103-3011303232233011-3320301133122203-0003301221202100) |
| `api_token.blindfold_secret_info.store_provider` | [api_token.blindfold_secret_info.store_provider](data-sources--cminstance--reference--group-001.md#canonical-0223223003003330-2011302330332232-1030121030020122-1201120221321033-2003303003221323-2012330130131231-0100021201301320-3133010311022302) |
| `api_token.clear_secret_info` | [api_token.clear_secret_info](data-sources--cminstance--reference--group-001.md#canonical-1322010130102032-1013113313330133-1022012203102132-3002133113111320-2112302223112233-3232300103222232-3100320332132011-1120032010032310) |
| `api_token.clear_secret_info.provider_ref` | [api_token.clear_secret_info.provider_ref](data-sources--cminstance--reference--group-001.md#canonical-1022313133102003-3003111023323010-1131203111313110-0122211000122003-0121312103312020-3031312033112032-3331122313222011-1031110120000212) |
| `api_token.clear_secret_info.url` | [api_token.clear_secret_info.url](data-sources--cminstance--reference--group-001.md#canonical-1202030021033112-3310223002221220-2300202330033232-1131011020203213-0301310203222112-3312011030013231-1331031021202110-1332000031231100) |
| `description` | [description](data-sources--cminstance--reference--group-001.md#canonical-3213321301001000-2100030022203133-0033111002302233-0101303103002002-3233311223011211-3033312023223320-2230010022121233-1310321011213210) |
| `id` | [ID](data-sources--cminstance--reference--group-001.md#canonical-3303022120310213-0020100303332010-2112100221333103-1320003203333020-2011301221202231-1031321031123001-0302022201113310-2111100021303032) |
| `ip` | [ip](data-sources--cminstance--reference--group-001.md#canonical-3122233212032132-3001133012021100-0022103113320003-3323300120031202-3230020002331021-1020103233003222-3333110120330322-1212320013203220) |
| `ip.addr` | [ip.addr](data-sources--cminstance--reference--group-001.md#canonical-2031201220130322-2232131201212313-1320210100123103-2001121220210333-0010121132011310-0113300333310330-1322301211021002-0233200033020231) |
| `labels` | [labels](data-sources--cminstance--reference--group-001.md#canonical-0120112313211110-3102133011313001-1211112030020113-3013200211100202-0230202003313312-1221321200332011-2223111133200203-1313022011021011) |
| `name` | [name](data-sources--cminstance--reference--group-001.md#canonical-0232123221020031-0130321203210310-2012303131331120-3323103000111022-3212101010320321-3210123123302230-1023320213100332-0033223223232202) |
| `namespace` | [namespace](data-sources--cminstance--reference--group-001.md#canonical-2103212201201323-3223033202121111-1200202110031001-1121120211231011-2012123103110123-3201311023301100-3302311231310102-2023223311202130) |
| `password` | [password](data-sources--cminstance--reference--group-001.md#canonical-2320223110001200-3223033130302012-1000323232300113-1021210322111121-3113210101211002-3331133200011002-0123221333032103-0001212132003333) |
| `password.blindfold_secret_info` | [password.blindfold_secret_info](data-sources--cminstance--reference--group-001.md#canonical-2110231003111110-0032011300211322-3001331010333200-1223021113330302-2313230230032303-0301313332110013-0222323222323202-1222220012001110) |
| `password.blindfold_secret_info.decryption_provider` | [password.blindfold_secret_info.decryption_provider](data-sources--cminstance--reference--group-001.md#canonical-0000202220021232-1222222320233113-0120021310311310-0223311030331023-0000001233333223-3312232003132321-1120323202213301-0232031031010033) |
| `password.blindfold_secret_info.location` | [password.blindfold_secret_info.location](data-sources--cminstance--reference--group-001.md#canonical-2033301332031130-0311210310310313-0022122333100201-0231231103111031-3311220020220222-0301102111100211-2313213302111300-0301232022203131) |
| `password.blindfold_secret_info.store_provider` | [password.blindfold_secret_info.store_provider](data-sources--cminstance--reference--group-001.md#canonical-1121202110212302-2023322210120311-1030311002231033-2323201011000300-2111133110112103-2203220020122212-0013001223200312-0120302110200133) |
| `password.clear_secret_info` | [password.clear_secret_info](data-sources--cminstance--reference--group-001.md#canonical-3103330232230132-2202313331013021-1213113230132131-0213321021300233-1100130110002202-3010222022030301-0110321032003211-0232231321200123) |
| `password.clear_secret_info.provider_ref` | [password.clear_secret_info.provider_ref](data-sources--cminstance--reference--group-001.md#canonical-3302023222100233-2011113130110333-1303310000103113-0111003001022101-3133321320222322-3030231201033133-0131300022303023-0113132223303311) |
| `password.clear_secret_info.url` | [password.clear_secret_info.url](data-sources--cminstance--reference--group-001.md#canonical-3010022113032233-3312111231001001-3332031213302022-2130200102321330-0020333120002023-3013223233320033-0002123133012001-2113022330302112) |
| `port` | [port](data-sources--cminstance--reference--group-001.md#canonical-1120001102212231-1231113313022032-0331230030203221-2022303020321010-2200331112022313-2001132211303223-1301002232312123-0312233320323101) |
| `username` | [username](data-sources--cminstance--reference--group-001.md#canonical-3231310311131313-0322301201301323-3211002311311220-0333303310321332-2222233301021102-2213002022103313-3001310310310112-2332322101230123) |

<a id="canonical-1301022102022133-1120021000112222-0113301303322203-2001133333133113-3123222330110233-1032232113323301-3222013330033022-0212200110023010"></a>

## Next pages — Property reference / 210101111222 / 13

- [api_token](data-sources--cminstance--reference--group-001.md#canonical-2122011332310321-2012020322322212-2103013120112311-1301000111201000-3210011003323231-1202212322022311-0001312112333103-1011331210231031)
- [ip](data-sources--cminstance--reference--group-001.md#canonical-0330120303210200-1132021301220102-0312131322232112-1133122030220320-2203003331333221-0221002022303031-2220220300100212-0232002201101111)
- [password](data-sources--cminstance--reference--group-001.md#canonical-1322111331201233-3313303002011133-2023312201301311-0330011002333312-0223110211233200-0010213202031321-1023122120220031-0213320030103121)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020)

<a id="canonical-2122011332310321-2012020322322212-2103013120112311-1301000111201000-3210011003323231-1202212322022311-0001312112333103-1011331210231031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122322002231222-2003301120300123-2312212212223333-3023111132101212-3102321200132002-3111202220301102-3303233332312010-0212001031002331"></a>

## api_token — api_token / 320220010121 / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-3021233103221211-3000330212303311-3123011101210133-2233111111213112-2200323123322113-3330123132032201-0000230310101123-2230201321203202)
- api_token

<a id="canonical-1233111010230322-0010103131120012-3002200212332133-0201301023022031-2121211203021222-3111120303021021-2232001310100030-1222023330130010"></a>

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

<a id="canonical-0210311210102232-2200301021012123-0111113120312021-3210220120002222-3021212003103022-1113020120230131-2300200111003101-1103303033100002"></a>

## Direct properties — api_token / 320220010121 / 3

- [blindfold_secret_info](data-sources--cminstance--reference--group-001.md#canonical-0030312203301333-2111213232003200-0320113212211311-2321000123301011-3033203333000131-0303332211330110-2203311001133332-2202303110302221): complete subsection reference.

- [clear_secret_info](data-sources--cminstance--reference--group-001.md#canonical-3033022130110220-2333320220110233-0313333202133102-1133302010330010-3012310010010232-3212311020100000-1021223320021223-2233303312100320): complete subsection reference.

<a id="canonical-3011203020231203-0203310023132100-0113233223030123-0302020010122032-2102123100213203-2211201330010200-3311000022311211-0102020321113220"></a>

## Next pages — api_token / 320220010121 / 4

- [api_token.blindfold_secret_info](data-sources--cminstance--reference--group-001.md#canonical-0030312203301333-2111213232003200-0320113212211311-2321000123301011-3033203333000131-0303332211330110-2203311001133332-2202303110302221)
- [api_token.clear_secret_info](data-sources--cminstance--reference--group-001.md#canonical-3033022130110220-2333320220110233-0313333202133102-1133302010330010-3012310010010232-3212311020100000-1021223320021223-2233303312100320)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-3021233103221211-3000330212303311-3123011101210133-2233111111213112-2200323123322113-3330123132032201-0000230310101123-2230201321203202)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020)

<a id="canonical-0030312203301333-2111213232003200-0320113212211311-2321000123301011-3033203333000131-0303332211330110-2203311001133332-2202303110302221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320021103312310-1121200110112322-0133101022110230-2331131003101332-0221013113312122-3122211101320003-2323020110211322-1333323000112003"></a>

## api_token.blindfold_secret_info — blindfold_secret_info / 230300311023 / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-3021233103221211-3000330212303311-3123011101210133-2233111111213112-2200323123322113-3330123132032201-0000230310101123-2230201321203202)
- [api_token](data-sources--cminstance--reference--group-001.md#canonical-2122011332310321-2012020322322212-2103013120112311-1301000111201000-3210011003323231-1202212322022311-0001312112333103-1011331210231031)
- api_token.blindfold_secret_info

<a id="canonical-3022031303123012-2331110020123133-3232033102123003-3202222211013122-1203320230033113-2231203200010031-2022233111123100-0032132300220122"></a>

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

<a id="canonical-2122200233330122-2221112200110300-0032223011011122-0212233002111100-2230123333321021-2001030020300320-1313331002111300-2102320013312211"></a>

## Direct properties — blindfold_secret_info / 230300311023 / 3

<a id="canonical-1022313203121031-0233322103121032-2020230101100120-3220331202213012-0300303323010310-3102332300001312-2013013320030020-1201121003012011"></a>

<a id="canonical-0230011031130101-0220232320032111-3203112232012123-2201213103220201-1001001122231112-2233232131122311-3111213112321232-3130111020202321"></a>

## decryption_provider property — blindfold_secret_info / 230300311023 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3020320123303100-0323033010311231-3030212313320012-0223122102120021-2003211001112103-3011303232233011-3320301133122203-0003301221202100"></a>

<a id="canonical-0130232310333211-1320330322111121-2022333022103130-0101330313122032-3200021002003332-1332110001130033-0312303013332031-2312331321033333"></a>

## location property — blindfold_secret_info / 230300311023 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0223223003003330-2011302330332232-1030121030020122-1201120221321033-2003303003221323-2012330130131231-0100021201301320-3133010311022302"></a>

<a id="canonical-1201012332120020-0300320013002120-0012210321202100-2101112303332321-2133332201013111-2232000200323033-2003123123133220-2212330303231202"></a>

## store_provider property — blindfold_secret_info / 230300311023 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0033223112130210-3013232121113011-2120121331032210-1331300002202030-1311201120131321-1332123010022001-1023110322033133-0233013131102320"></a>

## Next pages — blindfold_secret_info / 230300311023 / 7

- [api_token](data-sources--cminstance--reference--group-001.md#canonical-2122011332310321-2012020322322212-2103013120112311-1301000111201000-3210011003323231-1202212322022311-0001312112333103-1011331210231031)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020)

<a id="canonical-3033022130110220-2333320220110233-0313333202133102-1133302010330010-3012310010010232-3212311020100000-1021223320021223-2233303312100320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121330323011121-3110310013111222-1131001023321211-2122011312012301-3223332221202322-2100012320130020-3200002031020233-3120200222022001"></a>

## api_token.clear_secret_info — clear_secret_info / 222030120231 / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-3021233103221211-3000330212303311-3123011101210133-2233111111213112-2200323123322113-3330123132032201-0000230310101123-2230201321203202)
- [api_token](data-sources--cminstance--reference--group-001.md#canonical-2122011332310321-2012020322322212-2103013120112311-1301000111201000-3210011003323231-1202212322022311-0001312112333103-1011331210231031)
- api_token.clear_secret_info

<a id="canonical-1322010130102032-1013113313330133-1022012203102132-3002133113111320-2112302223112233-3232300103222232-3100320332132011-1120032010032310"></a>

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

<a id="canonical-1300231211211112-3113030031321023-2331020123333030-1332112312210100-2313002202111212-2032012131102032-2233323201212012-1112021130102002"></a>

## Direct properties — clear_secret_info / 222030120231 / 3

<a id="canonical-1022313133102003-3003111023323010-1131203111313110-0122211000122003-0121312103312020-3031312033112032-3331122313222011-1031110120000212"></a>

<a id="canonical-0120031000003131-3022020220321220-2203021132223303-0223331323002113-1012102012132113-0121033332111121-0011300302033020-3323203231132230"></a>

## provider_ref property — clear_secret_info / 222030120231 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1202030021033112-3310223002221220-2300202330033232-1131011020203213-0301310203222112-3312011030013231-1331031021202110-1332000031231100"></a>

<a id="canonical-0332202031031223-3100010023203210-2100312010113223-0113120213120211-1233301023212321-2103333002323213-2210112003101330-1101111300020002"></a>

## URL property — clear_secret_info / 222030120231 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0102323110113021-0020110311100320-2300202123303213-0010030130212232-3222012100311202-1323330100012322-1313101232303020-1033312211212221"></a>

## Next pages — clear_secret_info / 222030120231 / 6

- [api_token](data-sources--cminstance--reference--group-001.md#canonical-2122011332310321-2012020322322212-2103013120112311-1301000111201000-3210011003323231-1202212322022311-0001312112333103-1011331210231031)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020)

<a id="canonical-0330120303210200-1132021301220102-0312131322232112-1133122030220320-2203003331333221-0221002022303031-2220220300100212-0232002201101111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031211133133221-3203321302031021-2302130132123212-2320030122301210-1132003130102210-1310102223011302-2202330032320101-3020302221000232"></a>

## ip — ip / 110133002122 / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-3021233103221211-3000330212303311-3123011101210133-2233111111213112-2200323123322113-3330123132032201-0000230310101123-2230201321203202)
- ip

<a id="canonical-3122233212032132-3001133012021100-0022103113320003-3323300120031202-3230020002331021-1020103233003222-3333110120330322-1212320013203220"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

<a id="canonical-2120032233021210-3323023211102212-3231032223123012-2002011333222123-2132232101020013-2110223102111022-0301230223113303-1333301230100202"></a>

## Direct properties — ip / 110133002122 / 3

<a id="canonical-2031201220130322-2232131201212313-1320210100123103-2001121220210333-0010121132011310-0113300333310330-1322301211021002-0233200033020231"></a>

<a id="canonical-3201320102022001-3230012031121300-3120212231303123-3333201220232331-0310113101221301-2300222232312323-2220021312002131-2223122331222123"></a>

## addr property — ip / 110133002122 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1213111002133333-0311311112023102-1032000203122333-3233331323302003-2103300321101112-2312001030320302-3332001122202301-3132113003030111"></a>

## Next pages — ip / 110133002122 / 5

- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-3021233103221211-3000330212303311-3123011101210133-2233111111213112-2200323123322113-3330123132032201-0000230310101123-2230201321203202)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020)

<a id="canonical-1322111331201233-3313303002011133-2023312201301311-0330011002333312-0223110211233200-0010213202031321-1023122120220031-0213320030103121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121320210013321-2312031212233130-3211003031222001-2320321031121211-2220310333113302-2301230310321232-0012330111000002-1232332000013000"></a>

## password — password / 331130002120 / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-3021233103221211-3000330212303311-3123011101210133-2233111111213112-2200323123322113-3330123132032201-0000230310101123-2230201321203202)
- password

<a id="canonical-2320223110001200-3223033130302012-1000323232300113-1021210322111121-3113210101211002-3331133200011002-0123221333032103-0001212132003333"></a>

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

<a id="canonical-0300000113213330-1003010021010103-0112122300202131-2200033100330323-3220203231123002-1023101320232202-2000321313131111-0330111211033112"></a>

## Direct properties — password / 331130002120 / 3

- [blindfold_secret_info](data-sources--cminstance--reference--group-001.md#canonical-1010111322030023-2303132122020322-1221022133322033-0132312012101202-2211120010330213-3013213010000330-3210032000220122-0323211313113103): complete subsection reference.

- [clear_secret_info](data-sources--cminstance--reference--group-001.md#canonical-0032113001311202-1303000030003130-3102321013000012-0110310030002223-3131332311120332-2131320301213020-3213113133223320-1001121330001300): complete subsection reference.

<a id="canonical-3202230321012033-1231321010012331-2132202210230220-3333330321331103-2033231131332131-1010332130101302-3123032211031032-2320121332311332"></a>

## Next pages — password / 331130002120 / 4

- [password.blindfold_secret_info](data-sources--cminstance--reference--group-001.md#canonical-1010111322030023-2303132122020322-1221022133322033-0132312012101202-2211120010330213-3013213010000330-3210032000220122-0323211313113103)
- [password.clear_secret_info](data-sources--cminstance--reference--group-001.md#canonical-0032113001311202-1303000030003130-3102321013000012-0110310030002223-3131332311120332-2131320301213020-3213113133223320-1001121330001300)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-3021233103221211-3000330212303311-3123011101210133-2233111111213112-2200323123322113-3330123132032201-0000230310101123-2230201321203202)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020)

<a id="canonical-1010111322030023-2303132122020322-1221022133322033-0132312012101202-2211120010330213-3013213010000330-3210032000220122-0323211313113103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020231010333001-1101030020023031-2213300330211310-2020210310022000-1133320031230112-0302210212231020-0003110322013220-3110030202303122"></a>

## password.blindfold_secret_info — blindfold_secret_info / 023232333331 / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-3021233103221211-3000330212303311-3123011101210133-2233111111213112-2200323123322113-3330123132032201-0000230310101123-2230201321203202)
- [password](data-sources--cminstance--reference--group-001.md#canonical-1322111331201233-3313303002011133-2023312201301311-0330011002333312-0223110211233200-0010213202031321-1023122120220031-0213320030103121)
- password.blindfold_secret_info

<a id="canonical-2110231003111110-0032011300211322-3001331010333200-1223021113330302-2313230230032303-0301313332110013-0222323222323202-1222220012001110"></a>

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

<a id="canonical-3313103311032131-1212021021203311-2112211010210220-3110233311222310-2220211303323333-3203232120111213-3210001231103311-3121120320113211"></a>

## Direct properties — blindfold_secret_info / 023232333331 / 3

<a id="canonical-0000202220021232-1222222320233113-0120021310311310-0223311030331023-0000001233333223-3312232003132321-1120323202213301-0232031031010033"></a>

<a id="canonical-2021113300100302-1300011322213221-2323220312100111-0132203023032311-2331311110122233-2023302022332000-1010003103010113-0122310120301101"></a>

## decryption_provider property — blindfold_secret_info / 023232333331 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2033301332031130-0311210310310313-0022122333100201-0231231103111031-3311220020220222-0301102111100211-2313213302111300-0301232022203131"></a>

<a id="canonical-3103211321212330-0102231112000132-2310123100111020-0311323131322201-2212220212300322-0333301323102212-0301033122303310-0313203230131203"></a>

## location property — blindfold_secret_info / 023232333331 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1121202110212302-2023322210120311-1030311002231033-2323201011000300-2111133110112103-2203220020122212-0013001223200312-0120302110200133"></a>

<a id="canonical-0023133323201003-0210212131122111-3123130222211003-3203030230200031-2210223313120001-3122210113221232-3221002132113101-1030122001321221"></a>

## store_provider property — blindfold_secret_info / 023232333331 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2313023002320312-3321222213121031-3022130332312213-2012110222013203-3111133120223311-1113201200200222-3033033102001102-0330301223022210"></a>

## Next pages — blindfold_secret_info / 023232333331 / 7

- [password](data-sources--cminstance--reference--group-001.md#canonical-1322111331201233-3313303002011133-2023312201301311-0330011002333312-0223110211233200-0010213202031321-1023122120220031-0213320030103121)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020)

<a id="canonical-0032113001311202-1303000030003130-3102321013000012-0110310030002223-3131332311120332-2131320301213020-3213113133223320-1001121330001300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122220010100230-1222021013003011-3032332212101122-2232203311031322-0303132301001020-1011103120010200-0120131032023332-1012113312220331"></a>

## password.clear_secret_info — clear_secret_info / 120021320021 / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-3021233103221211-3000330212303311-3123011101210133-2233111111213112-2200323123322113-3330123132032201-0000230310101123-2230201321203202)
- [password](data-sources--cminstance--reference--group-001.md#canonical-1322111331201233-3313303002011133-2023312201301311-0330011002333312-0223110211233200-0010213202031321-1023122120220031-0213320030103121)
- password.clear_secret_info

<a id="canonical-3103330232230132-2202313331013021-1213113230132131-0213321021300233-1100130110002202-3010222022030301-0110321032003211-0232231321200123"></a>

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

<a id="canonical-3000212230300332-0011003211101112-0133223010230102-0110101123132320-1002120033212332-2300103300322030-2302102313132021-0131111211112130"></a>

## Direct properties — clear_secret_info / 120021320021 / 3

<a id="canonical-3302023222100233-2011113130110333-1303310000103113-0111003001022101-3133321320222322-3030231201033133-0131300022303023-0113132223303311"></a>

<a id="canonical-0110021203212332-1103110203001031-1222333022313110-2230100323201033-2301132322223213-0031311003332111-3011112023303202-3022222002223300"></a>

## provider_ref property — clear_secret_info / 120021320021 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3010022113032233-3312111231001001-3332031213302022-2130200102321330-0020333120002023-3013223233320033-0002123133012001-2113022330302112"></a>

<a id="canonical-3022022231113212-1220300302233301-1012313231300031-1100110123321101-3123120323131333-1332032121311032-1013030122001230-1110211033033211"></a>

## URL property — clear_secret_info / 120021320021 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2220330100131121-3112032012332211-1211010102110032-2232200201000112-3212332110223200-0232000313322230-0011210003021123-3103303201132332"></a>

## Next pages — clear_secret_info / 120021320021 / 6

- [password](data-sources--cminstance--reference--group-001.md#canonical-1322111331201233-3313303002011133-2023312201301311-0330011002333312-0223110211233200-0010213202031321-1023122120220031-0213320030103121)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020)
