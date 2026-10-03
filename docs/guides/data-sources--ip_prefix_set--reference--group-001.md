---
page_title: "xcsh_ip_prefix_set reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ip_prefix_set reference."
---

# xcsh_ip_prefix_set reference

<a id="canonical-0210223101032303-3303221220210133-3110331001212210-2021220232323132-2011212301131311-0200220122100200-1103212333012112-1201121010111102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122301013320013-3101100332130210-3000123003312002-2021001010103002-3230221123002112-2111113022112202-0303302322031232-2302310333313132"></a>

## Property reference — Property reference / 212120202033 / 2

Breadcrumbs:

- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md#canonical-2223223220333133-0031221032103123-3133100020012123-1131313331133102-3100111213100301-1210002313103313-0210003100132003-2132031023323130)
- Property reference

<a id="canonical-1012103121130030-2221203231121210-2233132233120200-0333100200310000-1112213323213120-1133230313200333-3032320201332303-0211323100320332"></a>

## Direct properties — Property reference / 212120202033 / 3

<a id="canonical-0022013222232322-3102121032300210-3023032003203032-2023223302002201-3023010330332220-1223302222331203-1312123320011201-1120222000023212"></a>

<a id="canonical-2110322113303102-0003100011213033-2301030001202232-0020221203131222-1001301203212022-1213003313103032-3210000332131200-2313101023311202"></a>

## annotations property — Property reference / 212120202033 / 4

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

<a id="canonical-3032230223222331-1122030213111020-1132303123200031-2310201201003213-2322231213002122-2112300321133130-0311202121022311-1122310202323212"></a>

<a id="canonical-1301302112332030-0211213222011011-2121030032111203-1322313332322331-3322301113210222-1033103202023323-0103020310020103-3301133132232310"></a>

## description property — Property reference / 212120202033 / 5

Type: `"string"`. Computed.

Description of the IPPrefixSet.

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

<a id="canonical-3320001311303121-0330032130103333-2322112332313312-2332101303123333-3302330103013333-0131312111210221-1023000032210030-0320111212001333"></a>

<a id="canonical-3310002202332233-1123123232110230-0332211223202223-1121322113131310-0123211132212200-1000331232120222-1010001230203021-0031211220203323"></a>

## ID property — Property reference / 212120202033 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipv4_prefixes](data-sources--ip_prefix_set--reference--group-001.md#canonical-2230202101210010-2331000201030332-0321131110201103-3123221101201312-1323101201322322-0023022032221311-2000312233223302-0121112122331233): complete subsection reference.

<a id="canonical-3232012012333112-0212212302003011-3120333230121222-3321203303013103-1301203201113022-1311031231302030-2001102022213021-3311320301302221"></a>

<a id="canonical-0132211202012321-3321212113210223-1131132313100011-1221310231101102-0002000122003030-1203010332132032-2303233210100131-2112202312230212"></a>

## labels property — Property reference / 212120202033 / 7

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

<a id="canonical-0221000102033201-0232122120223202-2320130331333010-1103122111023333-3023030122100010-0331123120232030-3032133102020111-3110201211103110"></a>

<a id="canonical-0322123003302302-3233030031300200-0132201301022131-2322223002300133-3132012001303023-2230012133130111-2220022032322212-1221333002220131"></a>

## name property — Property reference / 212120202033 / 8

Type: `"string"`. Required.

Name of the IPPrefixSet.

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

<a id="canonical-0100233101131222-2103030032211013-2122222113101120-3030332213001213-0233310300103323-2222001000302111-2223112101010200-3210312121032022"></a>

<a id="canonical-0123202000221123-0110333200223122-0002330210230132-0013200113110012-0021020131121222-2112022100323322-3312201213021013-2112210203201112"></a>

## namespace property — Property reference / 212120202033 / 9

Type: `"string"`. Required.

Namespace where the IPPrefixSet exists.

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

<a id="canonical-0112300103222310-2131100121010210-2320201010201130-2333202203333213-3033200233221130-0322103223201333-2113213232311012-1212013122130112"></a>

## All schema paths — Property reference / 212120202033 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--ip_prefix_set--reference--group-001.md#canonical-0022013222232322-3102121032300210-3023032003203032-2023223302002201-3023010330332220-1223302222331203-1312123320011201-1120222000023212) |
| `description` | [description](data-sources--ip_prefix_set--reference--group-001.md#canonical-3032230223222331-1122030213111020-1132303123200031-2310201201003213-2322231213002122-2112300321133130-0311202121022311-1122310202323212) |
| `id` | [ID](data-sources--ip_prefix_set--reference--group-001.md#canonical-3320001311303121-0330032130103333-2322112332313312-2332101303123333-3302330103013333-0131312111210221-1023000032210030-0320111212001333) |
| `ipv4_prefixes` | [ipv4_prefixes](data-sources--ip_prefix_set--reference--group-001.md#canonical-3321020212332020-3222230332103300-1223323233223030-2002132200110312-0033323232003330-0033102321320223-1000302201223030-1002301100131021) |
| `ipv4_prefixes.description_spec` | [ipv4_prefixes.description_spec](data-sources--ip_prefix_set--reference--group-001.md#canonical-2203021210211331-1010101013201121-0030132230003221-1231113100020012-1003113221112130-3002000221022231-2330110121000132-3003123200031001) |
| `ipv4_prefixes.ipv4_prefix` | [ipv4_prefixes.ipv4_prefix](data-sources--ip_prefix_set--reference--group-001.md#canonical-2203303310000200-1113232300023322-1133301211230331-0120220113230311-0201033223311101-1312231131221330-0212020133133012-1302120013200212) |
| `labels` | [labels](data-sources--ip_prefix_set--reference--group-001.md#canonical-3232012012333112-0212212302003011-3120333230121222-3321203303013103-1301203201113022-1311031231302030-2001102022213021-3311320301302221) |
| `name` | [name](data-sources--ip_prefix_set--reference--group-001.md#canonical-0221000102033201-0232122120223202-2320130331333010-1103122111023333-3023030122100010-0331123120232030-3032133102020111-3110201211103110) |
| `namespace` | [namespace](data-sources--ip_prefix_set--reference--group-001.md#canonical-0100233101131222-2103030032211013-2122222113101120-3030332213001213-0233310300103323-2222001000302111-2223112101010200-3210312121032022) |

<a id="canonical-1010122113001301-3213202023213202-3103320321330233-1310102133220330-3002100031303030-1210221020112101-2103003132310103-3013012210023021"></a>

## Next pages — Property reference / 212120202033 / 11

- [ipv4_prefixes](data-sources--ip_prefix_set--reference--group-001.md#canonical-2230202101210010-2331000201030332-0321131110201103-3123221101201312-1323101201322322-0023022032221311-2000312233223302-0121112122331233)
- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md#canonical-2223223220333133-0031221032103123-3133100020012123-1131313331133102-3100111213100301-1210002313103313-0210003100132003-2132031023323130)

<a id="canonical-2230202101210010-2331000201030332-0321131110201103-3123221101201312-1323101201322322-0023022032221311-2000312233223302-0121112122331233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303002233112310-0230033123122132-2200302110332032-0333110120220303-0133300003011113-0213311012102002-3331101122103200-1001010203020212"></a>

## ipv4_prefixes — ipv4_prefixes / 211202131030 / 2

Breadcrumbs:

- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md#canonical-2223223220333133-0031221032103123-3133100020012123-1131313331133102-3100111213100301-1210002313103313-0210003100132003-2132031023323130)
- [Property reference](data-sources--ip_prefix_set--reference--group-001.md#canonical-0210223101032303-3303221220210133-3110331001212210-2021220232323132-2011212301131311-0200220122100200-1103212333012112-1201121010111102)
- ipv4_prefixes

<a id="canonical-3321020212332020-3222230332103300-1223323233223030-2002132200110312-0033323232003330-0033102321320223-1000302201223030-1002301100131021"></a>

Type: `"list"`. Computed.

IPv4 Prefixes. List of IPv4 prefixes with description.

Upstream description:

List of IPv4 prefixes with description.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2232102232201031-0211031102210310-1121101312222222-3202202013213301-2321002332302000-0330001203122200-0312010110110021-0112131312010332"></a>

## Direct properties — ipv4_prefixes / 211202131030 / 3

<a id="canonical-2203021210211331-1010101013201121-0030132230003221-1231113100020012-1003113221112130-3002000221022231-2330110121000132-3003123200031001"></a>

<a id="canonical-0213101033211003-2320302031001300-3200011102003310-2231111031322333-2133012103021331-0001312031100321-0101312113123321-2032110312023002"></a>

## description_spec property — ipv4_prefixes / 211202131030 / 4

Type: `"string"`. Computed.

Description. Human-readable description text

<a id="canonical-2203303310000200-1113232300023322-1133301211230331-0120220113230311-0201033223311101-1312231131221330-0212020133133012-1302120013200212"></a>

<a id="canonical-2000321321010012-0232132002012112-1313121112001223-3120232323303121-3331102001123312-1220201210311132-0111100200111011-2103330200103333"></a>

## ipv4_prefix property — ipv4_prefixes / 211202131030 / 5

Type: `"string"`. Computed.

IPv4 Prefix. IP address configuration

Upstream description:

IP address configuration

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-3213012322133013-2220330001101012-1320322212120323-2212320201100301-0231210313010123-1103222202122233-2220021111203213-0111003120033200"></a>

## Next pages — ipv4_prefixes / 211202131030 / 6

- [Property reference](data-sources--ip_prefix_set--reference--group-001.md#canonical-0210223101032303-3303221220210133-3110331001212210-2021220232323132-2011212301131311-0200220122100200-1103212333012112-1201121010111102)
- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md#canonical-2223223220333133-0031221032103123-3133100020012123-1131313331133102-3100111213100301-1210002313103313-0210003100132003-2132031023323130)
