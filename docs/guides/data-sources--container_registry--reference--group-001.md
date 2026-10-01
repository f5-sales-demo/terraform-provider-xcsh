---
page_title: "xcsh_container_registry reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_container_registry reference."
---

# xcsh_container_registry reference

<a id="canonical-3211202121023222-3310130210222030-3300301333003322-1210013130013311-0121323231231103-3102100321133312-2232203033100102-2132331202300121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121213020301211-3023311331220123-2130031022002020-3220033021223220-1132210133003300-0232220131113021-2133201112311110-3033331000020201"></a>

## Property reference — Property reference / 302321131213 / 2

Breadcrumbs:

- [xcsh_container_registry](../data-sources/container_registry.md#canonical-0113331231101032-2331101103012211-0031031222323020-2332013123121012-2010031020013003-3001130201213030-1222020321220312-1223011211021133)
- Property reference

<a id="canonical-3210120030332223-2332233331113201-2020132121131133-1011331110221032-2001122013111122-3332311002331103-0001032322221312-0211102010233320"></a>

## Direct properties — Property reference / 302321131213 / 3

<a id="canonical-1000332112300023-2332002210222212-0011232110302201-0200032020210033-0130102013102022-0131030032333001-0132330321213313-2312222211021101"></a>

<a id="canonical-1000232211320123-2001010200313110-3131013121303231-2001102201022210-0313310010322133-2132110323111302-0300301122113120-2223220130102103"></a>

## annotations property — Property reference / 302321131213 / 4

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

<a id="canonical-1221313310002123-1101232322013321-1100020310203031-3303001300220003-0010322212003101-3311131003033012-1112113332123033-3203103200221020"></a>

<a id="canonical-2313023033031132-2202000202310311-0111222133130322-3201011313012203-0111322102312010-0210323113313222-0323332311111321-2220132102222010"></a>

## description property — Property reference / 302321131213 / 5

Type: `"string"`. Computed.

Description of the ContainerRegistry.

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

<a id="canonical-1132122213333112-2122212222332013-1231133010310102-0311003020213232-1133221110112331-1320032323200000-0201122110223131-3322011133113203"></a>

<a id="canonical-2321011110031312-0113313001202120-0330322023103000-2001000111031330-0222022111210331-2221110120120101-3232231300132202-2032002201002001"></a>

## email property — Property reference / 302321131213 / 6

Type: `"string"`. Computed.

Email. Email used for the registry.

Upstream description:

Email used for the registry.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "email",
    "formatDescription": "RFC 5322 email address, max 254 characters",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
    "validation": {
      "rfc": "RFC 5322"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.email": "true"
  }
}
```

<a id="canonical-0301310003131203-0122221032122110-2212313101033223-0133332000102313-3112300023133232-0332111133102231-2103103001213102-0311033320023333"></a>

<a id="canonical-2311301213130200-0230202311020222-3323333113110223-1312031200130133-2323332110202300-1233201123322323-0201100131302303-0033213320020112"></a>

## ID property — Property reference / 302321131213 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3003202003012223-0313103311030002-1111211122231210-2113220302113322-3012011331001221-1313300033320312-1132303100112333-1323233201011102"></a>

<a id="canonical-3033100023133130-2330213120332231-2330021010222012-1323003200001100-1301022002302332-3031113323203030-1201133100120001-3330121310332310"></a>

## labels property — Property reference / 302321131213 / 8

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

<a id="canonical-2302130331013332-3023311101131132-3101310122223102-1133113011320102-2100210310011321-3312122011222030-0002331231202110-2002221331331203"></a>

<a id="canonical-2001333021211020-3230131133321012-0010302213322213-1313212013123321-3301222133021203-0322313323332103-2330302130102212-3222002310011230"></a>

## name property — Property reference / 302321131213 / 9

Type: `"string"`. Required.

Name of the ContainerRegistry.

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

<a id="canonical-2230130120000310-2300100331001122-0011032222231311-0113120202222133-1232230333010122-3013033120200101-2302030222010031-0213210210011021"></a>

<a id="canonical-3231002003301032-0223301301110210-2331020220333303-0000020022323113-1101202331330203-3321322023012012-3211122220303222-1003220002012203"></a>

## namespace property — Property reference / 302321131213 / 10

Type: `"string"`. Required.

Namespace where the ContainerRegistry exists.

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

- [password](data-sources--container_registry--reference--group-001.md#canonical-1212000100200100-1302221200112101-1300300033113211-1202131302222323-1032221332111232-3310112110033231-3210221320333321-2320303221321233): complete subsection reference.

<a id="canonical-1221001203303311-0320122321110002-0231113133012000-0003111202211222-0322322312103003-0003313212321233-1123210012022233-1110201022213302"></a>

<a id="canonical-3133111020322002-1301103201332212-3221200201120122-0223322032023201-1112312021113120-1301301000202312-2002101303120323-1032231100112020"></a>

## registry property — Property reference / 302321131213 / 11

Type: `"string"`. Computed.

Fully qualified name of the registry login server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-2202003311230201-2213003022112320-1303132331102111-2301211020020103-3011132311101330-3220033003210011-2113012313203201-2021210013101232"></a>

<a id="canonical-0313121211221221-3321000022100100-0221133122101332-3132202212201323-2322313302013320-3202211303021231-3232303030130000-2201202122120010"></a>

## user_name property — Property reference / 302321131213 / 12

Type: `"string"`. Computed.

User Name. Username used to access the registry.

Upstream description:

Username used to access the registry.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2313013103113101-0201203221321312-0000320033232303-2121300312023131-3233010313333231-1211102012103021-3130322032022313-3203233102113230"></a>

## All schema paths — Property reference / 302321131213 / 13

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--container_registry--reference--group-001.md#canonical-1000332112300023-2332002210222212-0011232110302201-0200032020210033-0130102013102022-0131030032333001-0132330321213313-2312222211021101) |
| `description` | [description](data-sources--container_registry--reference--group-001.md#canonical-1221313310002123-1101232322013321-1100020310203031-3303001300220003-0010322212003101-3311131003033012-1112113332123033-3203103200221020) |
| `email` | [email](data-sources--container_registry--reference--group-001.md#canonical-1132122213333112-2122212222332013-1231133010310102-0311003020213232-1133221110112331-1320032323200000-0201122110223131-3322011133113203) |
| `id` | [id](data-sources--container_registry--reference--group-001.md#canonical-0301310003131203-0122221032122110-2212313101033223-0133332000102313-3112300023133232-0332111133102231-2103103001213102-0311033320023333) |
| `labels` | [labels](data-sources--container_registry--reference--group-001.md#canonical-3003202003012223-0313103311030002-1111211122231210-2113220302113322-3012011331001221-1313300033320312-1132303100112333-1323233201011102) |
| `name` | [name](data-sources--container_registry--reference--group-001.md#canonical-2302130331013332-3023311101131132-3101310122223102-1133113011320102-2100210310011321-3312122011222030-0002331231202110-2002221331331203) |
| `namespace` | [namespace](data-sources--container_registry--reference--group-001.md#canonical-2230130120000310-2300100331001122-0011032222231311-0113120202222133-1232230333010122-3013033120200101-2302030222010031-0213210210011021) |
| `password` | [password](data-sources--container_registry--reference--group-001.md#canonical-0231132103013311-3321311323013320-1312231131133222-3322023111120323-1212221213201022-1210002202102133-3320203232123303-2321122020010112) |
| `password.blindfold_secret_info` | [password.blindfold_secret_info](data-sources--container_registry--reference--group-001.md#canonical-3003220203223303-1131222330201230-2321333310102022-2201301330012310-2022221302130130-2110232011013311-3221331310103133-2333222211323120) |
| `password.blindfold_secret_info.decryption_provider` | [password.blindfold_secret_info.decryption_provider](data-sources--container_registry--reference--group-001.md#canonical-3233131032012002-1102002332003102-1122302321212130-1010010231000003-3103102022313211-1211313110112001-3313101331101223-0312202103131212) |
| `password.blindfold_secret_info.location` | [password.blindfold_secret_info.location](data-sources--container_registry--reference--group-001.md#canonical-1320310301132010-1301221230101200-1300230030311032-3201211321300030-1122233232130310-3023231220311323-0011031331303100-0330202133201230) |
| `password.blindfold_secret_info.store_provider` | [password.blindfold_secret_info.store_provider](data-sources--container_registry--reference--group-001.md#canonical-3201133211311123-3011102112021333-2031030201213300-3310230010021022-2333233211111113-3012001233130232-1031033120111210-3310102322033212) |
| `password.clear_secret_info` | [password.clear_secret_info](data-sources--container_registry--reference--group-001.md#canonical-1302301311311130-2312010302331323-2123202233231232-0001131003202012-2322123111103310-3233323223022221-0231320222121013-1332020333313201) |
| `password.clear_secret_info.provider_ref` | [password.clear_secret_info.provider_ref](data-sources--container_registry--reference--group-001.md#canonical-2133332000033103-1332002322320200-1022103000132221-3303103033233000-1301310230032133-1330003002311120-0112031012303033-3120131033112313) |
| `password.clear_secret_info.url` | [password.clear_secret_info.url](data-sources--container_registry--reference--group-001.md#canonical-3310030332333232-1002301212012311-3131002102320111-3123311230123112-1130111313232013-3222110020003101-2233111100023122-0322120100100330) |
| `registry` | [registry](data-sources--container_registry--reference--group-001.md#canonical-1221001203303311-0320122321110002-0231113133012000-0003111202211222-0322322312103003-0003313212321233-1123210012022233-1110201022213302) |
| `user_name` | [user_name](data-sources--container_registry--reference--group-001.md#canonical-2202003311230201-2213003022112320-1303132331102111-2301211020020103-3011132311101330-3220033003210011-2113012313203201-2021210013101232) |

<a id="canonical-0310100213123012-0020001121102122-1132010231222022-2213030011120123-2110231200100201-0031003312232301-1230133312120032-3230203030203133"></a>

## Next pages — Property reference / 302321131213 / 14

- [password](data-sources--container_registry--reference--group-001.md#canonical-1212000100200100-1302221200112101-1300300033113211-1202131302222323-1032221332111232-3310112110033231-3210221320333321-2320303221321233)
- [xcsh_container_registry](../data-sources/container_registry.md#canonical-0113331231101032-2331101103012211-0031031222323020-2332013123121012-2010031020013003-3001130201213030-1222020321220312-1223011211021133)

<a id="canonical-1212000100200100-1302221200112101-1300300033113211-1202131302222323-1032221332111232-3310112110033231-3210221320333321-2320303221321233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023210003322022-2333313123133322-1102231132222221-1231132323023103-3300210130000013-1212233333032112-1003221011003121-2022313012303211"></a>

## password — password / 302233221131 / 2

Breadcrumbs:

- [xcsh_container_registry](../data-sources/container_registry.md#canonical-0113331231101032-2331101103012211-0031031222323020-2332013123121012-2010031020013003-3001130201213030-1222020321220312-1223011211021133)
- [Property reference](data-sources--container_registry--reference--group-001.md#canonical-3211202121023222-3310130210222030-3300301333003322-1210013130013311-0121323231231103-3102100321133312-2232203033100102-2132331202300121)
- password

<a id="canonical-0231132103013311-3321311323013320-1312231131133222-3322023111120323-1212221213201022-1210002202102133-3320203232123303-2321122020010112"></a>

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

<a id="canonical-3323232302001011-1211222311020210-2333232332200102-2201001132123020-1020102123301112-2032022201210323-0232111231230322-2133001031220001"></a>

## Direct properties — password / 302233221131 / 3

- [blindfold_secret_info](data-sources--container_registry--reference--group-001.md#canonical-2112320202132103-3123111011221102-1201221112301331-1013202231231331-1310330230002200-1203121121322302-1132021333121330-2002131213001232): complete subsection reference.

- [clear_secret_info](data-sources--container_registry--reference--group-001.md#canonical-2231031222232103-1222133202000220-2031021322312303-0213122032121312-1102031022120002-2112311103203110-1022310010100300-0212012233011003): complete subsection reference.

<a id="canonical-0303313321133013-3101030013210112-1332100321313221-3101232311221321-3032332203320012-3302332203230312-2201313311113110-0330101201222023"></a>

## Next pages — password / 302233221131 / 4

- [password.blindfold_secret_info](data-sources--container_registry--reference--group-001.md#canonical-2112320202132103-3123111011221102-1201221112301331-1013202231231331-1310330230002200-1203121121322302-1132021333121330-2002131213001232)
- [password.clear_secret_info](data-sources--container_registry--reference--group-001.md#canonical-2231031222232103-1222133202000220-2031021322312303-0213122032121312-1102031022120002-2112311103203110-1022310010100300-0212012233011003)
- [Property reference](data-sources--container_registry--reference--group-001.md#canonical-3211202121023222-3310130210222030-3300301333003322-1210013130013311-0121323231231103-3102100321133312-2232203033100102-2132331202300121)
- [xcsh_container_registry](../data-sources/container_registry.md#canonical-0113331231101032-2331101103012211-0031031222323020-2332013123121012-2010031020013003-3001130201213030-1222020321220312-1223011211021133)

<a id="canonical-2112320202132103-3123111011221102-1201221112301331-1013202231231331-1310330230002200-1203121121322302-1132021333121330-2002131213001232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101002020223022-0213112233301003-2223323003003332-3121023123332021-2020321132200003-1301212033300311-0003011311122120-1313022323101331"></a>

## password.blindfold_secret_info — blindfold_secret_info / 111031303021 / 2

Breadcrumbs:

- [xcsh_container_registry](../data-sources/container_registry.md#canonical-0113331231101032-2331101103012211-0031031222323020-2332013123121012-2010031020013003-3001130201213030-1222020321220312-1223011211021133)
- [Property reference](data-sources--container_registry--reference--group-001.md#canonical-3211202121023222-3310130210222030-3300301333003322-1210013130013311-0121323231231103-3102100321133312-2232203033100102-2132331202300121)
- [password](data-sources--container_registry--reference--group-001.md#canonical-1212000100200100-1302221200112101-1300300033113211-1202131302222323-1032221332111232-3310112110033231-3210221320333321-2320303221321233)
- password.blindfold_secret_info

<a id="canonical-3003220203223303-1131222330201230-2321333310102022-2201301330012310-2022221302130130-2110232011013311-3221331310103133-2333222211323120"></a>

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

<a id="canonical-1311102232111111-2332021013100312-3333000012013303-0100131312213031-0320013132022100-3212003300033113-3232101120111123-1122013321030322"></a>

## Direct properties — blindfold_secret_info / 111031303021 / 3

<a id="canonical-3233131032012002-1102002332003102-1122302321212130-1010010231000003-3103102022313211-1211313110112001-3313101331101223-0312202103131212"></a>

<a id="canonical-1100013230332301-0000013221131303-0133101323000220-0102113123223031-1103020320011011-3131223022303131-1010230230303001-3233122011111332"></a>

## decryption_provider property — blindfold_secret_info / 111031303021 / 4

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

<a id="canonical-1320310301132010-1301221230101200-1300230030311032-3201211321300030-1122233232130310-3023231220311323-0011031331303100-0330202133201230"></a>

<a id="canonical-3031023003103012-1031132222111223-0100132330133213-0033032003131000-1213232000220132-3321303011023122-2103203202223023-1101103013031212"></a>

## location property — blindfold_secret_info / 111031303021 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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

<a id="canonical-3201133211311123-3011102112021333-2031030201213300-3310230010021022-2333233211111113-3012001233130232-1031033120111210-3310102322033212"></a>

<a id="canonical-1201300312232031-3300120213101103-3131113222020001-3010200030303203-1133230213021030-1111120130133000-0021310102120311-2203230003131030"></a>

## store_provider property — blindfold_secret_info / 111031303021 / 6

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

<a id="canonical-0132002323100213-3003220103012132-0232010120112011-3322032333011223-0031111033332100-2102132302321333-0101223113121131-2321213301001031"></a>

## Next pages — blindfold_secret_info / 111031303021 / 7

- [password](data-sources--container_registry--reference--group-001.md#canonical-1212000100200100-1302221200112101-1300300033113211-1202131302222323-1032221332111232-3310112110033231-3210221320333321-2320303221321233)
- [xcsh_container_registry](../data-sources/container_registry.md#canonical-0113331231101032-2331101103012211-0031031222323020-2332013123121012-2010031020013003-3001130201213030-1222020321220312-1223011211021133)

<a id="canonical-2231031222232103-1222133202000220-2031021322312303-0213122032121312-1102031022120002-2112311103203110-1022310010100300-0212012233011003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332123320101031-1331213303022303-2131013232112100-1102103123210210-0132123330300121-0002333222131231-3011001111100103-3301230120101210"></a>

## password.clear_secret_info — clear_secret_info / 303303220132 / 2

Breadcrumbs:

- [xcsh_container_registry](../data-sources/container_registry.md#canonical-0113331231101032-2331101103012211-0031031222323020-2332013123121012-2010031020013003-3001130201213030-1222020321220312-1223011211021133)
- [Property reference](data-sources--container_registry--reference--group-001.md#canonical-3211202121023222-3310130210222030-3300301333003322-1210013130013311-0121323231231103-3102100321133312-2232203033100102-2132331202300121)
- [password](data-sources--container_registry--reference--group-001.md#canonical-1212000100200100-1302221200112101-1300300033113211-1202131302222323-1032221332111232-3310112110033231-3210221320333321-2320303221321233)
- password.clear_secret_info

<a id="canonical-1302301311311130-2312010302331323-2123202233231232-0001131003202012-2322123111103310-3233323223022221-0231320222121013-1332020333313201"></a>

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

<a id="canonical-0233023232103321-2103330223311130-0212203023023113-3021130022000100-0310000013001130-3100322211302330-1102003020112103-0120300311122312"></a>

## Direct properties — clear_secret_info / 303303220132 / 3

<a id="canonical-2133332000033103-1332002322320200-1022103000132221-3303103033233000-1301310230032133-1330003002311120-0112031012303033-3120131033112313"></a>

<a id="canonical-0203213113131312-3130203211012103-2321213202310101-3212201312212103-2222331010102312-0121213001233112-0211210311102330-3332312003130103"></a>

## provider_ref property — clear_secret_info / 303303220132 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3310030332333232-1002301212012311-3131002102320111-3123311230123112-1130111313232013-3222110020003101-2233111100023122-0322120100100330"></a>

<a id="canonical-0030023102202212-0323012113322131-1031112302121133-3120202020313021-3311013312033320-3322321000300311-1333303312331212-2200200231313312"></a>

## URL property — clear_secret_info / 303303220132 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-2010100030112122-2310103020202003-0302120222230011-0023332131110122-3131200312033330-3200221021112301-0031111113232301-2120323300303323"></a>

## Next pages — clear_secret_info / 303303220132 / 6

- [password](data-sources--container_registry--reference--group-001.md#canonical-1212000100200100-1302221200112101-1300300033113211-1202131302222323-1032221332111232-3310112110033231-3210221320333321-2320303221321233)
- [xcsh_container_registry](../data-sources/container_registry.md#canonical-0113331231101032-2331101103012211-0031031222323020-2332013123121012-2010031020013003-3001130201213030-1222020321220312-1223011211021133)
