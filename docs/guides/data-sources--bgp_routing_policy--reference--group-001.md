---
page_title: "xcsh_bgp_routing_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_routing_policy reference."
---

# xcsh_bgp_routing_policy reference

<a id="canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020323101200333-1212103213313102-2320212001302231-0213130031231321-2323200101312101-0020333021030331-0301011013333000-0000321030133101"></a>

## Property reference — Property reference / 311211130321 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- Property reference

<a id="canonical-3013220002312213-2201021113202003-3312021303300022-0030210233201310-1120131020000101-2330211011313121-2000011332223332-3111202123010030"></a>

## Direct properties — Property reference / 311211130321 / 3

<a id="canonical-0211033203101211-1211102111210332-0031022130020230-0212102020302312-2223011133210311-2213303311011001-0300001200212000-3010303100310123"></a>

<a id="canonical-1000213013023122-0121001031303031-2300120230202031-0332133131103301-3122231311210113-1302222030010302-0232012201013201-1301310232302000"></a>

## annotations property — Property reference / 311211130321 / 4

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

<a id="canonical-0100232332112110-3111003100120032-2221302023013111-0133332110031012-1233021302023101-1330211122321122-3320020332131202-1331203322001331"></a>

<a id="canonical-3333223210011030-2123002031223120-1200121121332303-0232033130212332-1113302312303202-0303120212333201-1121231023301313-1133033232213123"></a>

## description property — Property reference / 311211130321 / 5

Type: `"string"`. Computed.

Description of the BGPRoutingPolicy.

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

<a id="canonical-0312222131203120-3313322313330022-2212322101001213-1033132231220112-3222320321000013-1110213012221323-3023131212200020-1222312011021232"></a>

<a id="canonical-3310003203111222-0211221003000112-2133301301011210-3221331112110201-3010330022312131-3121113310232231-3013200133110011-0301130023101113"></a>

## ID property — Property reference / 311211130321 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0303113323122031-2210311222103031-3212330121023130-2021133030311300-3003130232013103-1111223203010203-2002002101032312-3232201111013200"></a>

<a id="canonical-0111002011332233-2103132101010002-2033330133023122-0102333113130100-3233302012233231-0221123310232321-1220220301133101-0110023232102130"></a>

## labels property — Property reference / 311211130321 / 7

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

<a id="canonical-0322102320101223-3111221031121123-1212201301331330-1200203302211313-3021301223103233-3203312010020333-3110312313122302-3133121123312321"></a>

<a id="canonical-1320311323011230-0123001310213112-2202203101333233-2312230012122020-0021123331002320-0000320221300233-1330321000131133-1312210103211120"></a>

## name property — Property reference / 311211130321 / 8

Type: `"string"`. Required.

Name of the BGPRoutingPolicy.

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

<a id="canonical-1221122101012233-1011313022221322-3330301003011321-0132022201011102-0312213001312212-1123110230011313-3211302230033231-3022210212323310"></a>

<a id="canonical-2313133113031122-1203232102101302-1033321001031231-2012120332102203-0003111102001133-1211231030120332-3232003110332023-1021012232103302"></a>

## namespace property — Property reference / 311211130321 / 9

Type: `"string"`. Required.

Namespace where the BGPRoutingPolicy exists.

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

- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101): complete subsection reference.

<a id="canonical-2330123332132333-2110323001121131-1333332003232330-1201303332230333-3021321201201032-0300300122221110-2303022233100312-3000102132200003"></a>

## All schema paths — Property reference / 311211130321 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0211033203101211-1211102111210332-0031022130020230-0212102020302312-2223011133210311-2213303311011001-0300001200212000-3010303100310123) |
| `description` | [description](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0100232332112110-3111003100120032-2221302023013111-0133332110031012-1233021302023101-1330211122321122-3320020332131202-1331203322001331) |
| `id` | [id](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0312222131203120-3313322313330022-2212322101001213-1033132231220112-3222320321000013-1110213012221323-3023131212200020-1222312011021232) |
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

<a id="canonical-2203133123301020-3101323321131130-3302231112213102-0021031110030331-0020110001032231-0010230002200301-1302032101033103-0232102322020200"></a>

## Next pages — Property reference / 311211130321 / 11

- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)

<a id="canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021103020300113-0310013300022320-2132123322030100-2313030320111201-1121132310013302-2032222121200203-3022202210123331-2310011313222113"></a>

## rules — rules / 113033300301 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- rules

<a id="canonical-1033310312020001-2303221011331201-1003323313112123-3020020212013201-1133022020212010-2313333203120133-2213002112123312-0303113331233021"></a>

Type: `"list"`. Computed.

BGP Routing policy is composed of one or more rules. Note that the order of rules is critical as
rules are applied top to bottom.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1123012110120101-2200110332011000-2101200213011313-2100132111320013-3011131003320121-1222312320130322-3131222222220020-1021320323321132"></a>

## Direct properties — rules / 113033300301 / 3

- [action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3213010232033132-2233122120322102-2201313030203000-0333323221221012-1233023012003330-0312312003320132-1320330121302120-2110113132311321): complete subsection reference.

- [match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1322101220011103-0313013332201213-2323000301233023-3230120113201302-3001323212210311-2030201201311323-1231001323010013-0331000221102012): complete subsection reference.

<a id="canonical-2202202123222303-3321310132101323-2120210112123021-3101111022022132-2231303303331012-1232033022321033-3212330121321301-1221131310130111"></a>

## Next pages — rules / 113033300301 / 4

- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3213010232033132-2233122120322102-2201313030203000-0333323221221012-1233023012003330-0312312003320132-1320330121302120-2110113132311321)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1322101220011103-0313013332201213-2323000301233023-3230120113201302-3001323212210311-2030201201311323-1231001323010013-0331000221102012)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)

<a id="canonical-3213010232033132-2233122120322102-2201313030203000-0333323221221012-1233023012003330-0312312003320132-1320330121302120-2110113132311321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221220131322131-2230113312123010-2200013313222333-1132103223312113-0031013201320333-0031100032022233-3232231132322000-1132033221130201"></a>

## rules.action — action / 212211201230 / 2

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

<a id="canonical-2201233001030321-3311022101002021-2300123320221201-3200323331322002-1202103002313310-1230333102101313-2200301030302002-2310213202001222"></a>

## Direct properties — action / 212211201230 / 3

- [allow](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1111230013213002-2323230010031231-1222020311012212-2101103330020123-0013303203002013-0233231003323332-1010303330121303-1011202023110123): complete subsection reference.

<a id="canonical-0323310110221332-2213320303030111-1313030332011223-1202320022301331-0301231030322213-0131012330203200-3023300111012001-2230131000112322"></a>

<a id="canonical-3030123303123022-2221200231310202-0130111230212330-3131030322101110-0021113020210100-0220001230102132-0220203332030213-0013200200331122"></a>

## as_path property — action / 212211201230 / 4

Type: `"string"`. Computed.

Exclusive with \[allow community deny local\_preference metric\] AS-Path Prepending is generally
used to influence incoming traffic.

Upstream description:

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

- [community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0033110321000302-2132303301212021-1332331010232320-3101332112131311-0322321331333201-1321010001213133-3222300233100233-1220322110221031): complete subsection reference.

- [deny](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1312023331110120-2333223332331130-3100300032033130-2013112001020131-1033133132030202-0300033113331302-3132001322132020-1112201001322101): complete subsection reference.

<a id="canonical-1132202222313232-0220202112213232-0020110113221212-2222203301313200-2313001203003323-1312223200011221-0111300223010301-0000203123301110"></a>

<a id="canonical-2230311022133133-2222302113221303-3100321230021112-0323201311111211-3011322322233110-0022312131332023-0312103033022332-1202020021111312"></a>

## local_preference property — action / 212211201230 / 5

Type: `"number"`. Computed.

Exclusive with \[allow as\_path community deny metric\] BGP Local Preference is generally used to
influence outgoing traffic.

Upstream description:

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

<a id="canonical-3302301332200211-1322033330023110-3212131202011322-2102031023222130-1030110200003033-2323230103213013-2111301313222123-1331112330103221"></a>

## metric property — action / 212211201230 / 6

Type: `"number"`. Computed.

Exclusive with \[allow as\_path community deny local\_preference\] The Multi-Exit Discriminator
metric to indicate the preferred path to AS.

Upstream description:

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

<a id="canonical-0022320010123212-2321323223023331-1201220111211103-0131133011031031-0020300333010223-2231212130301320-2210231001311222-3022221000002023"></a>

## Next pages — action / 212211201230 / 7

- [rules.action.allow](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1111230013213002-2323230010031231-1222020311012212-2101103330020123-0013303203002013-0233231003323332-1010303330121303-1011202023110123)
- [rules.action.community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0033110321000302-2132303301212021-1332331010232320-3101332112131311-0322321331333201-1321010001213133-3222300233100233-1220322110221031)
- [rules.action.deny](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1312023331110120-2333223332331130-3100300032033130-2013112001020131-1033133132030202-0300033113331302-3132001322132020-1112201001322101)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)

<a id="canonical-1111230013213002-2323230010031231-1222020311012212-2101103330020123-0013303203002013-0233231003323332-1010303330121303-1011202023110123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103011111303311-3310322202212011-2200102323312310-1101220121201323-0220121332311331-2000132031030312-0322231032001112-2023233002223210"></a>

## rules.action.allow — allow / 320221312321 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3213010232033132-2233122120322102-2201313030203000-0333323221221012-1233023012003330-0312312003320132-1320330121302120-2110113132311321)
- rules.action.allow

<a id="canonical-0111222201122101-0031021333310120-0232332122030013-1232232201301001-1020322100230000-3202131210100103-1221301202013031-2321130102032001"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-2120213212313202-3233223332221232-3333301100203312-1210023013332021-1033033110113233-0031313100331320-3233122010212221-1330202013123321"></a>

## Direct properties — allow / 320221312321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323200013121113-3331020111331233-2300031312332020-2021210000110302-0122323121333200-1021022311233333-3230233220001300-0111231121232202"></a>

## Next pages — allow / 320221312321 / 4

- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3213010232033132-2233122120322102-2201313030203000-0333323221221012-1233023012003330-0312312003320132-1320330121302120-2110113132311321)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)

<a id="canonical-0033110321000302-2132303301212021-1332331010232320-3101332112131311-0322321331333201-1321010001213133-3222300233100233-1220322110221031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023003011332311-2300111200033012-2101220022233321-0020303201320031-0220302132120230-1002003100122002-3301131222302102-0123212223120123"></a>

## rules.action.community — community / 301030110100 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3213010232033132-2233122120322102-2201313030203000-0333323221221012-1233023012003330-0312312003320132-1320330121302120-2110113132311321)
- rules.action.community

<a id="canonical-3200221222013233-2111122033321223-3100213310321030-1321330310222230-3321001213100202-1300210302021212-2220330321122113-2122101211020132"></a>

Type: `"single"`. Computed.

BGP Community list. List of BGP communities.

Upstream description:

List of BGP communities.

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

<a id="canonical-0322323313230331-0311111312000032-1203232011110112-1313030230122303-2020033230332131-3333111202210132-1132212000121011-2133213121313220"></a>

## Direct properties — community / 301030110100 / 3

<a id="canonical-1123202202001331-0121023020103021-3223200003110323-3121112300103233-2010301330131203-0030131033020232-3301203322011230-1312001130033302"></a>

<a id="canonical-3230121223000030-1131122331131031-1101103221203302-2103133210330011-3332223331212320-1130011032022231-1223223331330131-1303222130021001"></a>

## community property — community / 301030110100 / 4

Type: `["list", "string"]`. Computed.

Unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits being
value.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3001212322011123-1303033222110220-1133321100331210-1121023112210133-2323231011131120-3101230202102003-1211111313223002-3031231011303001"></a>

## Next pages — community / 301030110100 / 5

- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3213010232033132-2233122120322102-2201313030203000-0333323221221012-1233023012003330-0312312003320132-1320330121302120-2110113132311321)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)

<a id="canonical-1312023331110120-2333223332331130-3100300032033130-2013112001020131-1033133132030202-0300033113331302-3132001322132020-1112201001322101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321310130120021-3032310223032102-3110312313310032-2323113230131312-3200003021003333-1113210131113132-3331331303203320-1331112333131230"></a>

## rules.action.deny — deny / 011003010330 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3213010232033132-2233122120322102-2201313030203000-0333323221221012-1233023012003330-0312312003320132-1320330121302120-2110113132311321)
- rules.action.deny

<a id="canonical-1332001302113231-0320030003302203-2320032111200231-2210112312332123-3333100322230112-1100202332102203-2301123200230031-0100231211011131"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-0220000332200332-2210020112133230-2200120131110213-3013333013220130-0132222131030131-3002230200302030-1212323010122001-0300211212323003"></a>

## Direct properties — deny / 011003010330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311212323331302-0010320013002230-0010303231032112-2333032023033210-3311220303101230-3032031131210020-3013102221232103-0202332320221231"></a>

## Next pages — deny / 011003010330 / 4

- [rules.action](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3213010232033132-2233122120322102-2201313030203000-0333323221221012-1233023012003330-0312312003320132-1320330121302120-2110113132311321)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)

<a id="canonical-1322101220011103-0313013332201213-2323000301233023-3230120113201302-3001323212210311-2030201201311323-1231001323010013-0331000221102012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012231023223312-0332321221111303-1221130303221130-3312132300012331-0030311021002133-3211031102022231-3312210131302020-2333322301212233"></a>

## rules.match — match / 322200211302 / 2

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

<a id="canonical-1113111210212030-3020130023021113-1200232103310312-1203113103312332-1322012121300031-1320200322033020-2310221110220012-3013013020002120"></a>

## Direct properties — match / 322200211302 / 3

<a id="canonical-0010311113023333-0321310213112020-3320022101222122-2031013010032022-3110132112312301-2112320322233023-0311033103330013-0033002030003011"></a>

<a id="canonical-2210210202133202-3000313310222030-1012032301102120-3223220302312330-3202011010103100-0131220123100000-2011102113010231-0132103332223310"></a>

## as_path property — match / 322200211302 / 4

Type: `"string"`. Computed.

Exclusive with \[community ip\_prefixes\] AS path can also be a regex, which will be matched against
route information.

Upstream description:

Exclusive with \[community ip\_prefixes\] AS path can also be a regex, which will be matched against
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

- [community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0001202131112321-1213030200213233-2221231022133021-3303221103331100-3203223332203301-0123123200123230-1103133212100301-1213312033210322): complete subsection reference.

- [ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3232120320131320-1221111231000331-2033031200002022-2231123202131301-3023212210322112-3131202030202103-2111032111021232-0012100132103323): complete subsection reference.

<a id="canonical-2323231003303132-1020213013123112-2331212223132233-2003113121223203-2101120123323131-1312320320102333-0001031012332033-3322232333121222"></a>

## Next pages — match / 322200211302 / 5

- [rules.match.community](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0001202131112321-1213030200213233-2221231022133021-3303221103331100-3203223332203301-0123123200123230-1103133212100301-1213312033210322)
- [rules.match.ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3232120320131320-1221111231000331-2033031200002022-2231123202131301-3023212210322112-3131202030202103-2111032111021232-0012100132103323)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)

<a id="canonical-0001202131112321-1213030200213233-2221231022133021-3303221103331100-3203223332203301-0123123200123230-1103133212100301-1213312033210322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331222213332033-0022201330010303-2013321332012331-1221113101210030-0111020301213321-0301300120213032-2002333120233332-3130213230313330"></a>

## rules.match.community — community / 030101012213 / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Property reference](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [rules](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2012101133313130-2303020032221131-3321331301330221-3211032200210031-3203200121201330-1011012010003121-2210122112102310-0013000232211101)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1322101220011103-0313013332201213-2323000301233023-3230120113201302-3001323212210311-2030201201311323-1231001323010013-0331000221102012)
- rules.match.community

<a id="canonical-1003213020122333-2003232223222131-2221213313303210-3023211131201302-3333313232001032-1131221201321110-2203023010123311-0203230233230322"></a>

Type: `"single"`. Computed.

BGP Community list. List of BGP communities.

Upstream description:

List of BGP communities.

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

<a id="canonical-0010303320100131-0330032213011132-1312313102211310-3210001121322010-0020022021301201-3111230233122121-1131132113221011-2130332233010310"></a>

## Direct properties — community / 030101012213 / 3

<a id="canonical-3210031300210313-1003001011132232-2130022310003000-0101033013122110-0220013212030221-3101102020033012-3200100001202330-2101132132320220"></a>

<a id="canonical-1032130303333321-3312321223300200-3221130010311332-1333320131311313-0200031201022203-1222323113222210-1110132332201121-3221333022311322"></a>

## community property — community / 030101012213 / 4

Type: `["list", "string"]`. Computed.

Unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits being
value.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1230123122030233-1233130033002222-0130202032200012-3231322132120231-1213130301030133-3230023301222213-2001231300202222-3110230330213230"></a>

## Next pages — community / 030101012213 / 5

- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1322101220011103-0313013332201213-2323000301233023-3230120113201302-3001323212210311-2030201201311323-1231001323010013-0331000221102012)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)

<a id="canonical-3232120320131320-1221111231000331-2033031200002022-2231123202131301-3023212210322112-3131202030202103-2111032111021232-0012100132103323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231321301013013-3313220110230232-0121313121112302-1011233013012113-0302003322033212-1020111233333311-1313101202303130-3212023123222001"></a>

## rules.match.ip_prefixes — ip_prefixes / 201233300002 / 2

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

<a id="canonical-0113323212323323-1222232300020312-1013001302202112-0332131333202313-1323111320123000-3302311133131031-1202331313012110-2032202113130310"></a>

## Direct properties — ip_prefixes / 201233300002 / 3

- [prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2302223130322231-3123200002133201-2210033123113033-3323300033012323-2223223121202002-2303002113320113-2211031100001131-0103201001002000): complete subsection reference.

<a id="canonical-2111023132333230-0102232132322021-2103332211011103-3222323302223133-0010330013010302-1030233122112121-0220221232233030-0013332333100132"></a>

## Next pages — ip_prefixes / 201233300002 / 4

- [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2302223130322231-3123200002133201-2210033123113033-3323300033012323-2223223121202002-2303002113320113-2211031100001131-0103201001002000)
- [rules.match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-1322101220011103-0313013332201213-2323000301233023-3230120113201302-3001323212210311-2030201201311323-1231001323010013-0331000221102012)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)

<a id="canonical-2302223130322231-3123200002133201-2210033123113033-3323300033012323-2223223121202002-2303002113320113-2211031100001131-0103201001002000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032122201302211-0300332333212132-1322301111320312-1101130113202013-3310003333020001-0123002100001330-3312300122302320-1310131013212110"></a>

## rules.match.ip_prefixes.prefixes — prefixes / 231320132001 / 2

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

Upstream description:

List of IP prefix.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1332012120010032-0002213031232113-0121123030030203-3202013032031001-1031311212203110-1100002210330233-3333113132131322-3222113013332133"></a>

## Direct properties — prefixes / 231320132001 / 3

- [equal_or_longer_than](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3123103233200300-2320313010131123-1213211222320332-3223201113300320-1302022033320222-1110032313110002-0211302323330213-1213112000233102): complete subsection reference.

- [exact_match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2131231201023113-0002123003002130-0233312121323313-3100212010313210-1302202333011203-0323002002211210-1200332212313333-3103130200020130): complete subsection reference.

<a id="canonical-3230211210103020-2311030012010132-1020121110123032-1122221033323102-3331021112031110-3200223312202322-2200101202301010-1123031003302303"></a>

<a id="canonical-1210311233231000-1212311020023020-1023002012322202-3221200030122022-3032121020001320-1321022112330112-3013022330031100-2210130323201301"></a>

## ip_prefixes property — prefixes / 231320132001 / 4

Type: `"string"`. Computed.

IP Prefix. IP prefix to match on BGP route.

Upstream description:

IP prefix to match on BGP route.

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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [longer_than](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0320113330111233-2020000223303110-0232033020102033-1103213213312211-2032030302323322-2100232020020003-1033002201110100-3022022113231333): complete subsection reference.

<a id="canonical-3011200011223211-0121332200023130-0111323022213113-0120001213330002-2330130301223210-0002231033212001-3030311210123221-2011102213322120"></a>

## Next pages — prefixes / 231320132001 / 5

- [rules.match.ip_prefixes.prefixes.equal_or_longer_than](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3123103233200300-2320313010131123-1213211222320332-3223201113300320-1302022033320222-1110032313110002-0211302323330213-1213112000233102)
- [rules.match.ip_prefixes.prefixes.exact_match](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2131231201023113-0002123003002130-0233312121323313-3100212010313210-1302202333011203-0323002002211210-1200332212313333-3103130200020130)
- [rules.match.ip_prefixes.prefixes.longer_than](data-sources--bgp_routing_policy--reference--group-001.md#canonical-0320113330111233-2020000223303110-0232033020102033-1103213213312211-2032030302323322-2100232020020003-1033002201110100-3022022113231333)
- [rules.match.ip_prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-3232120320131320-1221111231000331-2033031200002022-2231123202131301-3023212210322112-3131202030202103-2111032111021232-0012100132103323)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)

<a id="canonical-3123103233200300-2320313010131123-1213211222320332-3223201113300320-1302022033320222-1110032313110002-0211302323330213-1213112000233102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010131323110323-3310030222303301-0012311001132220-1233101131310110-1013210101101313-0023030210231200-2202103112221310-2022011020222221"></a>

## rules.match.ip_prefixes.prefixes.equal_or_longer_than — equal_or_longer_than / 213100010111 / 2

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

Upstream description:

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

<a id="canonical-3321320022110131-1331131101002220-2110321121312021-2032112233102203-3332213100133031-0011321133013121-0322232031212203-2232310201120211"></a>

## Direct properties — equal_or_longer_than / 213100010111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132132120222213-3112312000330111-0000103112321000-1331310213102233-1000222133233322-2210200022130121-2330212302113332-2101003223001230"></a>

## Next pages — equal_or_longer_than / 213100010111 / 4

- [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2302223130322231-3123200002133201-2210033123113033-3323300033012323-2223223121202002-2303002113320113-2211031100001131-0103201001002000)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)

<a id="canonical-2131231201023113-0002123003002130-0233312121323313-3100212010313210-1302202333011203-0323002002211210-1200332212313333-3103130200020130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112213213031202-2202020231331330-3321230232233001-0111111120320011-3330103310213310-2112233100123121-1032030133301330-0030210223221210"></a>

## rules.match.ip_prefixes.prefixes.exact_match — exact_match / 313103112223 / 2

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

Upstream description:

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

<a id="canonical-3322302022030110-3313212222000022-0333330122132111-3332312200331021-1310030100212302-2012332021023112-0033013010120103-1220113322321320"></a>

## Direct properties — exact_match / 313103112223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313112303200030-0132130011023300-1313302210201032-3321100203001103-1030120203111021-0022103331200120-3121130322223322-3332113023220303"></a>

## Next pages — exact_match / 313103112223 / 4

- [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2302223130322231-3123200002133201-2210033123113033-3323300033012323-2223223121202002-2303002113320113-2211031100001131-0103201001002000)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)

<a id="canonical-0320113330111233-2020000223303110-0232033020102033-1103213213312211-2032030302323322-2100232020020003-1033002201110100-3022022113231333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132300032121021-2122223300032003-0102213320123212-0220302000033233-3130233112130320-1020101311012003-0212323333102130-2021303220332121"></a>

## rules.match.ip_prefixes.prefixes.longer_than — longer_than / 122300101301 / 2

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

Upstream description:

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

<a id="canonical-2012321300312221-1103013030031302-1001022320201021-3033110300133232-3210011001233022-3021030133003223-1012101130011123-2020223131202032"></a>

## Direct properties — longer_than / 122300101301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2132031013212221-2011003003130210-3322203202022113-3112310123113103-3200030322131023-2330303222223031-0011020221031132-2320312012310133"></a>

## Next pages — longer_than / 122300101301 / 4

- [rules.match.ip_prefixes.prefixes](data-sources--bgp_routing_policy--reference--group-001.md#canonical-2302223130322231-3123200002133201-2210033123113033-3323300033012323-2223223121202002-2303002113320113-2211031100001131-0103201001002000)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
