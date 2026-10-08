---
page_title: "xcsh_bgp_routing_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_routing_policy reference."
---

# xcsh_bgp_routing_policy reference

<a id="canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- Property reference

<a id="canonical-3020323101200333-1212103213313102-2320212001302231-0213130031231321-2323200101312101-0020333021030331-0301011013333000-0000321030133101"></a>

### Direct properties for `xcsh_bgp_routing_policy`

<a id="canonical-0211033203101211-1211102111210332-0031022130020230-0212102020302312-2223011133210311-2213303311011001-0300001200212000-3010303100310123"></a>

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

<a id="canonical-0100232332112110-3111003100120032-2221302023013111-0133332110031012-1233021302023101-1330211122321122-3320020332131202-1331203322001331"></a>

<a id="canonical-3013220002312213-2201021113202003-3312021303300022-0030210233201310-1120131020000101-2330211011313121-2000011332223332-3111202123010030"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the BGPRoutingPolicy.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0312222131203120-3313322313330022-2212322101001213-1033132231220112-3222320321000013-1110213012221323-3023131212200020-1222312011021232"></a>

<a id="canonical-1000213013023122-0121001031303031-2300120230202031-0332133131103301-3122231311210113-1302222030010302-0232012201013201-1301310232302000"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0303113323122031-2210311222103031-3212330121023130-2021133030311300-3003130232013103-1111223203010203-2002002101032312-3232201111013200"></a>

<a id="canonical-3333223210011030-2123002031223120-1200121121332303-0232033130212332-1113302312303202-0303120212333201-1121231023301313-1133033232213123"></a>

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

<a id="canonical-0322102320101223-3111221031121123-1212201301331330-1200203302211313-3021301223103233-3203312010020333-3110312313122302-3133121123312321"></a>

<a id="canonical-3310003203111222-0211221003000112-2133301301011210-3221331112110201-3010330022312131-3121113310232231-3013200133110011-0301130023101113"></a>

#### `name` property

Type: `"string"`. Required.

Name of the BGPRoutingPolicy.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1221122101012233-1011313022221322-3330301003011321-0132022201011102-0312213001312212-1123110230011313-3211302230033231-3022210212323310"></a>

<a id="canonical-0111002011332233-2103132101010002-2033330133023122-0102333113130100-3233302012233231-0221123310232321-1220220301133101-0110023232102130"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the BGPRoutingPolicy exists.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101): complete subsection reference.

<a id="canonical-1320311323011230-0123001310213112-2202203101333233-2312230012122020-0021123331002320-0000320221300233-1330321000131133-1312210103211120"></a>

### All schema paths for `xcsh_bgp_routing_policy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0211033203101211-1211102111210332-0031022130020230-0212102020302312-2223011133210311-2213303311011001-0300001200212000-3010303100310123) |
| `description` | [description](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0100232332112110-3111003100120032-2221302023013111-0133332110031012-1233021302023101-1330211122321122-3320020332131202-1331203322001331) |
| `id` | [ID](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0312222131203120-3313322313330022-2212322101001213-1033132231220112-3222320321000013-1110213012221323-3023131212200020-1222312011021232) |
| `labels` | [labels](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0303113323122031-2210311222103031-3212330121023130-2021133030311300-3003130232013103-1111223203010203-2002002101032312-3232201111013200) |
| `name` | [name](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0322102320101223-3111221031121123-1212201301331330-1200203302211313-3021301223103233-3203312010020333-3110312313122302-3133121123312321) |
| `namespace` | [namespace](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1221122101012233-1011313022221322-3330301003011321-0132022201011102-0312213001312212-1123110230011313-3211302230033231-3022210212323310) |
| `rules` | [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1033310312020001-2303221011331201-1003323313112123-3020020212013201-1133022020212010-2313333203120133-2213002112123312-0303113331233021) |
| `rules.action` | [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3103200213003003-0312323323033133-3013013030213331-2011322123130223-2130112100222203-1333231001002323-1301032221312123-2102313100332101) |
| `rules.action.allow` | [rules.action.allow](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0111222201122101-0031021333310120-0232332122030013-1232232201301001-1020322100230000-3202131210100103-1221301202013031-2321130102032001) |
| `rules.action.as_path` | [rules.action.as_path](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0323310110221332-2213320303030111-1313030332011223-1202320022301331-0301231030322213-0131012330203200-3023300111012001-2230131000112322) |
| `rules.action.community` | [rules.action.community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3200221222013233-2111122033321223-3100213310321030-1321330310222230-3321001213100202-1300210302021212-2220330321122113-2122101211020132) |
| `rules.action.community.community` | [rules.action.community.community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1123202202001331-0121023020103021-3223200003110323-3121112300103233-2010301330131203-0030131033020232-3301203322011230-1312001130033302) |
| `rules.action.deny` | [rules.action.deny](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1332001302113231-0320030003302203-2320032111200231-2210112312332123-3333100322230112-1100202332102203-2301123200230031-0100231211011131) |
| `rules.action.local_preference` | [rules.action.local_preference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1132202222313232-0220202112213232-0020110113221212-2222203301313200-2313001203003323-1312223200011221-0111300223010301-0000203123301110) |
| `rules.action.metric` | [rules.action.metric](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0300010002320230-0013203320233013-0311213213300023-0223132102311011-1233220213021301-1210213210112111-3011220231111212-1331220032322003) |
| `rules.match` | [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0123133001211110-1123211331222032-1113211222002201-3332312021001323-1321132000213320-0111331010323230-3233002233132120-2232221212301133) |
| `rules.match.as_path` | [rules.match.as_path](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0010311113023333-0321310213112020-3320022101222122-2031013010032022-3110132112312301-2112320322233023-0311033103330013-0033002030003011) |
| `rules.match.community` | [rules.match.community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1003213020122333-2003232223222131-2221213313303210-3023211131201302-3333313232001032-1131221201321110-2203023010123311-0203230233230322) |
| `rules.match.community.community` | [rules.match.community.community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3210031300210313-1003001011132232-2130022310003000-0101033013122110-0220013212030221-3101102020033012-3200100001202330-2101132132320220) |
| `rules.match.ip_prefixes` | [rules.match.ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3101210331202332-1330332011122302-1103123211223123-2030131312001131-1233223311230230-2230111213001123-3131120201020101-0203021200301013) |
| `rules.match.ip_prefixes.prefixes` | [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0001231202211003-2332021203320130-3333121121033223-1130002331311132-2303223332221120-0113302032001223-1122020001001031-2203211032120202) |
| `rules.match.ip_prefixes.prefixes.equal_or_longer_than` | [rules.match.ip_prefixes.prefixes.equal_or_longer_than](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0310202003133331-1132133232212223-0221300301031322-3001302330222203-2133112323112032-3300131022122000-1223202233303131-3232020002313032) |
| `rules.match.ip_prefixes.prefixes.exact_match` | [rules.match.ip_prefixes.prefixes.exact_match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2033203123202113-3122111202322230-2223331132311003-3131331031010103-3200223021120101-3233322231303021-1200001210113113-1312113201300302) |
| `rules.match.ip_prefixes.prefixes.ip_prefixes` | [rules.match.ip_prefixes.prefixes.ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230211210103020-2311030012010132-1020121110123032-1122221033323102-3331021112031110-3200223312202322-2200101202301010-1123031003302303) |
| `rules.match.ip_prefixes.prefixes.longer_than` | [rules.match.ip_prefixes.prefixes.longer_than](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3231202123112302-0322110120223210-2011131031233012-2222100001213322-0323332301203230-0321300110032033-2120110201202332-2131013130010133) |

<a id="canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules` properties

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- rules

<a id="canonical-1033310312020001-2303221011331201-1003323313112123-3020020212013201-1133022020212010-2313333203120133-2213002112123312-0303113331233021"></a>

Type: `"list"`. Computed.

A BGP Routing policy is composed of one or more rules. Note that the order of rules is critical as
rules are applied top to bottom.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minItems": 1,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3021103020300113-0310013300022320-2132123322030100-2313030320111201-1121132310013302-2032222121200203-3022202210123331-2310011313222113"></a>

### Direct properties for `rules`

- [action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3213010232033132-2233122120322102-2201313030203000-0333323221221012-1233023012003330-0312312003320132-1320330121302120-2110113132311321): complete subsection reference.

- [match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1322101220011103-0313013332201213-2323000301233023-3230120113201302-3001323212210311-2030201201311323-1231001323010013-0331000221102012): complete subsection reference.

<a id="canonical-3213010232033132-2233122120322102-2201313030203000-0333323221221012-1233023012003330-0312312003320132-1320330121302120-2110113132311321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.action` properties

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- rules.action

<a id="canonical-3103200213003003-0312323323033133-3013013030213331-2011322123130223-2130112100222203-1333231001002323-1301032221312123-2102313100332101"></a>

Type: `"single"`. Computed.

Action to be enforced if the BGP route matches the rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"allow\",\"as_path\",\"community\",\"deny\",\"local_preference\",\"metric\"]"
}
```

<a id="canonical-2221220131322131-2230113312123010-2200013313222333-1132103223312113-0031013201320333-0031100032022233-3232231132322000-1132033221130201"></a>

### Direct properties for `rules.action`

- [allow](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1111230013213002-2323230010031231-1222020311012212-2101103330020123-0013303203002013-0233231003323332-1010303330121303-1011202023110123): complete subsection reference.

<a id="canonical-0323310110221332-2213320303030111-1313030332011223-1202320022301331-0301231030322213-0131012330203200-3023300111012001-2230131000112322"></a>

<a id="canonical-2201233001030321-3311022101002021-2300123320221201-3200323331322002-1202103002313310-1230333102101313-2200301030302002-2310213202001222"></a>

#### `rules.action.as_path` property

Type: `"string"`. Computed.

Exclusive with \[allow community deny local\_preference metric\] AS-Path Prepending is generally
used to influence incoming traffic.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0033110321000302-2132303301212021-1332331010232320-3101332112131311-0322321331333201-1321010001213133-3222300233100233-1220322110221031): complete subsection reference.

- [deny](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1312023331110120-2333223332331130-3100300032033130-2013112001020131-1033133132030202-0300033113331302-3132001322132020-1112201001322101): complete subsection reference.

<a id="canonical-1132202222313232-0220202112213232-0020110113221212-2222203301313200-2313001203003323-1312223200011221-0111300223010301-0000203123301110"></a>

<a id="canonical-3030123303123022-2221200231310202-0130111230212330-3131030322101110-0021113020210100-0220001230102132-0220203332030213-0013200200331122"></a>

#### `rules.action.local_preference` property

Type: `"number"`. Computed.

Exclusive with \[allow as\_path community deny metric\] BGP Local Preference is generally used to
influence outgoing traffic.

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

<a id="canonical-0300010002320230-0013203320233013-0311213213300023-0223132102311011-1233220213021301-1210213210112111-3011220231111212-1331220032322003"></a>

<a id="canonical-2230311022133133-2222302113221303-3100321230021112-0323201311111211-3011322322233110-0022312131332023-0312103033022332-1202020021111312"></a>

#### `rules.action.metric` property

Type: `"number"`. Computed.

Exclusive with \[allow as\_path community deny local\_preference\] The Multi-Exit Discriminator
metric to indicate the preferred path to AS.

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

<a id="canonical-1111230013213002-2323230010031231-1222020311012212-2101103330020123-0013303203002013-0233231003323332-1010303330121303-1011202023110123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.action.allow` properties

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3213010232033132-2233122120322102-2201313030203000-0333323221221012-1233023012003330-0312312003320132-1320330121302120-2110113132311321)
- rules.action.allow

<a id="canonical-0111222201122101-0031021333310120-0232332122030013-1232232201301001-1020322100230000-3202131210100103-1221301202013031-2321130102032001"></a>

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

<a id="canonical-0033110321000302-2132303301212021-1332331010232320-3101332112131311-0322321331333201-1321010001213133-3222300233100233-1220322110221031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.action.community` properties

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3213010232033132-2233122120322102-2201313030203000-0333323221221012-1233023012003330-0312312003320132-1320330121302120-2110113132311321)
- rules.action.community

<a id="canonical-3200221222013233-2111122033321223-3100213310321030-1321330310222230-3321001213100202-1300210302021212-2220330321122113-2122101211020132"></a>

Type: `"single"`. Computed.

BGP Community list. List of BGP communities.

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

<a id="canonical-2023003011332311-2300111200033012-2101220022233321-0020303201320031-0220302132120230-1002003100122002-3301131222302102-0123212223120123"></a>

### Direct properties for `rules.action.community`

<a id="canonical-1123202202001331-0121023020103021-3223200003110323-3121112300103233-2010301330131203-0030131033020232-3301203322011230-1312001130033302"></a>

#### `rules.action.community.community` property

Type: `["list", "string"]`. Computed.

An unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits
being value.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1312023331110120-2333223332331130-3100300032033130-2013112001020131-1033133132030202-0300033113331302-3132001322132020-1112201001322101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.action.deny` properties

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3213010232033132-2233122120322102-2201313030203000-0333323221221012-1233023012003330-0312312003320132-1320330121302120-2110113132311321)
- rules.action.deny

<a id="canonical-1332001302113231-0320030003302203-2320032111200231-2210112312332123-3333100322230112-1100202332102203-2301123200230031-0100231211011131"></a>

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

<a id="canonical-1322101220011103-0313013332201213-2323000301233023-3230120113201302-3001323212210311-2030201201311323-1231001323010013-0331000221102012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.match` properties

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- rules.match

<a id="canonical-0123133001211110-1123211331222032-1113211222002201-3332312021001323-1321132000213320-0111331010323230-3233002233132120-2232221212301133"></a>

Type: `"single"`. Computed.

Predicates which have to match information in route for action to be applied.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type_of_match": "[\"as_path\",\"community\",\"ip_prefixes\"]"
}
```

<a id="canonical-2012231023223312-0332321221111303-1221130303221130-3312132300012331-0030311021002133-3211031102022231-3312210131302020-2333322301212233"></a>

### Direct properties for `rules.match`

<a id="canonical-0010311113023333-0321310213112020-3320022101222122-2031013010032022-3110132112312301-2112320322233023-0311033103330013-0033002030003011"></a>

#### `rules.match.as_path` property

Type: `"string"`. Computed.

Exclusive with \[community ip\_prefixes\] AS path can also be a regular expression, which will be matched against
route information.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0001202131112321-1213030200213233-2221231022133021-3303221103331100-3203223332203301-0123123200123230-1103133212100301-1213312033210322): complete subsection reference.

- [ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3232120320131320-1221111231000331-2033031200002022-2231123202131301-3023212210322112-3131202030202103-2111032111021232-0012100132103323): complete subsection reference.

<a id="canonical-0001202131112321-1213030200213233-2221231022133021-3303221103331100-3203223332203301-0123123200123230-1103133212100301-1213312033210322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.match.community` properties

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1322101220011103-0313013332201213-2323000301233023-3230120113201302-3001323212210311-2030201201311323-1231001323010013-0331000221102012)
- rules.match.community

<a id="canonical-1003213020122333-2003232223222131-2221213313303210-3023211131201302-3333313232001032-1131221201321110-2203023010123311-0203230233230322"></a>

Type: `"single"`. Computed.

BGP Community list. List of BGP communities.

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

<a id="canonical-2331222213332033-0022201330010303-2013321332012331-1221113101210030-0111020301213321-0301300120213032-2002333120233332-3130213230313330"></a>

### Direct properties for `rules.match.community`

<a id="canonical-3210031300210313-1003001011132232-2130022310003000-0101033013122110-0220013212030221-3101102020033012-3200100001202330-2101132132320220"></a>

#### `rules.match.community.community` property

Type: `["list", "string"]`. Computed.

An unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits
being value.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3232120320131320-1221111231000331-2033031200002022-2231123202131301-3023212210322112-3131202030202103-2111032111021232-0012100132103323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.match.ip_prefixes` properties

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1322101220011103-0313013332201213-2323000301233023-3230120113201302-3001323212210311-2030201201311323-1231001323010013-0331000221102012)
- rules.match.ip_prefixes

<a id="canonical-3101210331202332-1330332011122302-1103123211223123-2030131312001131-1233223311230230-2230111213001123-3131120201020101-0203021200301013"></a>

Type: `"single"`. Computed.

List of IP prefix and prefix length range match condition.

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

<a id="canonical-1231321301013013-3313220110230232-0121313121112302-1011233013012113-0302003322033212-1020111233333311-1313101202303130-3212023123222001"></a>

### Direct properties for `rules.match.ip_prefixes`

- [prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2302223130322231-3123200002133201-2210033123113033-3323300033012323-2223223121202002-2303002113320113-2211031100001131-0103201001002000): complete subsection reference.

<a id="canonical-2302223130322231-3123200002133201-2210033123113033-3323300033012323-2223223121202002-2303002113320113-2211031100001131-0103201001002000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.match.ip_prefixes.prefixes` properties

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1322101220011103-0313013332201213-2323000301233023-3230120113201302-3001323212210311-2030201201311323-1231001323010013-0331000221102012)
- [rules.match.ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3232120320131320-1221111231000331-2033031200002022-2231123202131301-3023212210322112-3131202030202103-2111032111021232-0012100132103323)
- rules.match.ip_prefixes.prefixes

<a id="canonical-0001231202211003-2332021203320130-3333121121033223-1130002331311132-2303223332221120-0113302032001223-1122020001001031-2203211032120202"></a>

Type: `"list"`. Computed.

Prefix list. List of IP prefix.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3032122201302211-0300332333212132-1322301111320312-1101130113202013-3310003333020001-0123002100001330-3312300122302320-1310131013212110"></a>

### Direct properties for `rules.match.ip_prefixes.prefixes`

- [equal_or_longer_than](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3123103233200300-2320313010131123-1213211222320332-3223201113300320-1302022033320222-1110032313110002-0211302323330213-1213112000233102): complete subsection reference.

- [exact_match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2131231201023113-0002123003002130-0233312121323313-3100212010313210-1302202333011203-0323002002211210-1200332212313333-3103130200020130): complete subsection reference.

<a id="canonical-3230211210103020-2311030012010132-1020121110123032-1122221033323102-3331021112031110-3200223312202322-2200101202301010-1123031003302303"></a>

<a id="canonical-1332012120010032-0002213031232113-0121123030030203-3202013032031001-1031311212203110-1100002210330233-3333113132131322-3222113013332133"></a>

#### `rules.match.ip_prefixes.prefixes.ip_prefixes` property

Type: `"string"`. Computed.

IP Prefix. IP prefix to match on BGP route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [longer_than](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0320113330111233-2020000223303110-0232033020102033-1103213213312211-2032030302323322-2100232020020003-1033002201110100-3022022113231333): complete subsection reference.

<a id="canonical-3123103233200300-2320313010131123-1213211222320332-3223201113300320-1302022033320222-1110032313110002-0211302323330213-1213112000233102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.match.ip_prefixes.prefixes.equal_or_longer_than` properties

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1322101220011103-0313013332201213-2323000301233023-3230120113201302-3001323212210311-2030201201311323-1231001323010013-0331000221102012)
- [rules.match.ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3232120320131320-1221111231000331-2033031200002022-2231123202131301-3023212210322112-3131202030202103-2111032111021232-0012100132103323)
- [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2302223130322231-3123200002133201-2210033123113033-3323300033012323-2223223121202002-2303002113320113-2211031100001131-0103201001002000)
- rules.match.ip_prefixes.prefixes.equal_or_longer_than

<a id="canonical-0310202003133331-1132133232212223-0221300301031322-3001302330222203-2133112323112032-3300131022122000-1223202233303131-3232020002313032"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for equal or longer than.

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

<a id="canonical-2131231201023113-0002123003002130-0233312121323313-3100212010313210-1302202333011203-0323002002211210-1200332212313333-3103130200020130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.match.ip_prefixes.prefixes.exact_match` properties

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1322101220011103-0313013332201213-2323000301233023-3230120113201302-3001323212210311-2030201201311323-1231001323010013-0331000221102012)
- [rules.match.ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3232120320131320-1221111231000331-2033031200002022-2231123202131301-3023212210322112-3131202030202103-2111032111021232-0012100132103323)
- [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2302223130322231-3123200002133201-2210033123113033-3323300033012323-2223223121202002-2303002113320113-2211031100001131-0103201001002000)
- rules.match.ip_prefixes.prefixes.exact_match

<a id="canonical-2033203123202113-3122111202322230-2223331132311003-3131331031010103-3200223021120101-3233322231303021-1200001210113113-1312113201300302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for exact match.

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

<a id="canonical-0320113330111233-2020000223303110-0232033020102033-1103213213312211-2032030302323322-2100232020020003-1033002201110100-3022022113231333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.match.ip_prefixes.prefixes.longer_than` properties

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1322101220011103-0313013332201213-2323000301233023-3230120113201302-3001323212210311-2030201201311323-1231001323010013-0331000221102012)
- [rules.match.ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3232120320131320-1221111231000331-2033031200002022-2231123202131301-3023212210322112-3131202030202103-2111032111021232-0012100132103323)
- [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2302223130322231-3123200002133201-2210033123113033-3323300033012323-2223223121202002-2303002113320113-2211031100001131-0103201001002000)
- rules.match.ip_prefixes.prefixes.longer_than

<a id="canonical-3231202123112302-0322110120223210-2011131031233012-2222100001213322-0323332301203230-0321300110032033-2120110201202332-2131013130010133"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for longer than.

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
