---
page_title: "xcsh_network_policy_view reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_view reference."
---

# xcsh_network_policy_view reference

<a id="canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123103312213333-0200332301133101-2030212133023131-3310203202012223-1131313032032222-1023030211122130-2001020331102030-3301201011312031"></a>

## Property reference — Property reference / 211210010121 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- Property reference

<a id="canonical-2130220000101302-3221323211121302-0302122011221232-1333003103201312-1210202221332311-0200100100021302-1010000312000000-2012332201320100"></a>

## Direct properties — Property reference / 211210010121 / 3

<a id="canonical-0112023300132001-0131032212213212-2203123001210332-3232102010313113-1302030022330032-1130233001201111-2320300310233032-3002113032112312"></a>

<a id="canonical-1211220201120202-2032232011103000-3130022103221213-1220310311101003-2122200011233011-1133033231003133-0133222303203222-1231232003331201"></a>

## annotations property — Property reference / 211210010121 / 4

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

<a id="canonical-1012020213021330-3103030022102122-0200222201120331-1010321333033030-0002003223130132-1013012023311201-3111000020000131-0011233023020223"></a>

<a id="canonical-3132200232003200-0121103303033113-2010112031112132-2101223002203211-1210132100022121-0011022111011131-0013131233200303-2220132201201022"></a>

## description property — Property reference / 211210010121 / 5

Type: `"string"`. Computed.

Description of the NetworkPolicyView.

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

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022): complete subsection reference.

- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120): complete subsection reference.

<a id="canonical-1212231112313013-0030020201312101-2232200333002003-3021222231101031-1222031002201233-0201230231121301-0102120200003032-1000230233110322"></a>

<a id="canonical-0020003012002213-3001102210201333-3313121200231133-0301203130011311-2011131112110210-3332203110301022-0322311102002020-2210132031321330"></a>

## ID property — Property reference / 211210010121 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230): complete subsection reference.

<a id="canonical-3301230033121220-3133023010111112-1313113310031131-0111023010223223-2020130013323223-3003021233232111-1123001031300030-0113111011230103"></a>

<a id="canonical-1313320320203123-1210203211000113-1102200223330202-1233113313211110-2310310130122132-3002111200033121-1033303003001112-0123000110311200"></a>

## labels property — Property reference / 211210010121 / 7

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

<a id="canonical-2031220122032332-1013330130211303-1312003231021332-0100013003020011-1021331110221202-2112321133331022-3210122133121232-0222111311010022"></a>

<a id="canonical-3311312013002130-2002101113110121-0221003011002001-2222002322333002-1020030231021233-3020330323310303-1012023203121310-1123302033102303"></a>

## name property — Property reference / 211210010121 / 8

Type: `"string"`. Required.

Name of the NetworkPolicyView.

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

<a id="canonical-1003123010310220-0211121011003203-3021002303232312-1020201211121320-3223202113220033-3011332123002130-2130003322212001-3332322213020010"></a>

<a id="canonical-3121200321233110-1221223323213133-3030021002002233-0221322122233020-2001320210200200-0102020130031122-0100120133321222-2311131221011022"></a>

## namespace property — Property reference / 211210010121 / 9

Type: `"string"`. Optional, Computed.

Namespace where the NetworkPolicyView exists.

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

<a id="canonical-1100031210321012-0210302133101301-0332323030331322-3321330010200131-1032022000302023-1032011100033131-3323202321030213-0130331203332103"></a>

## All schema paths — Property reference / 211210010121 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--network_policy_view--reference--group-001.md#canonical-0112023300132001-0131032212213212-2203123001210332-3232102010313113-1302030022330032-1130233001201111-2320300310233032-3002113032112312) |
| `description` | [description](data-sources--network_policy_view--reference--group-001.md#canonical-1012020213021330-3103030022102122-0200222201120331-1010321333033030-0002003223130132-1013012023311201-3111000020000131-0011233023020223) |
| `egress_rules` | [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0210012222310121-0203301010320030-1112112210033300-3320100030220311-2030020333321220-1212031200230132-1201211202302112-1220000120221030) |
| `egress_rules.action` | [egress_rules.action](data-sources--network_policy_view--reference--group-001.md#canonical-0103320330222110-2300210011021101-0301021221000332-0120112110222003-1123331023130300-3112302011121010-3301112311222331-1212131321013112) |
| `egress_rules.adv_action` | [egress_rules.adv_action](data-sources--network_policy_view--reference--group-001.md#canonical-2213210032120033-2331233212230021-0103012103113302-0010331132221033-1110013031320230-1101010200201231-3220332031113030-0132330100130230) |
| `egress_rules.adv_action.action` | [egress_rules.adv_action.action](data-sources--network_policy_view--reference--group-001.md#canonical-1010222112000120-2130230103022001-0302003333320212-2223021012210000-0213223102203001-3011011211003010-0101101323310210-2211102100321303) |
| `egress_rules.all_tcp_traffic` | [egress_rules.all_tcp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-1313110331133132-1333211021110032-2221111030023323-2200030022313122-3213221010120103-3113122323233223-0130112321200023-2002303120221320) |
| `egress_rules.all_traffic` | [egress_rules.all_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-0112130123320212-3202203011303200-2102022121022221-1222121221022320-0323332013320202-0220221221210100-0123032111013321-3011130101101311) |
| `egress_rules.all_udp_traffic` | [egress_rules.all_udp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-0322213201221013-2213313120310323-0203011232310232-0102021113003312-0201020013323233-2112213033033211-0200203011102003-0102033302112033) |
| `egress_rules.any` | [egress_rules.any](data-sources--network_policy_view--reference--group-001.md#canonical-2202030120110030-0203300212211122-3212322200102200-2233223123112331-1202121302100333-2310010303102102-2223301120103300-2210130031000101) |
| `egress_rules.applications` | [egress_rules.applications](data-sources--network_policy_view--reference--group-001.md#canonical-1223201121101002-0103322132222021-2101222100231011-3010000230310310-0200031330330301-2320203100211102-2122030122211310-3323010203131103) |
| `egress_rules.applications.applications` | [egress_rules.applications.applications](data-sources--network_policy_view--reference--group-001.md#canonical-2100233213332031-1120313112230313-3203220310121023-2023031230202331-0000023220001010-3130220212021020-2221120131100102-0313033231103232) |
| `egress_rules.inside_endpoints` | [egress_rules.inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-0310302311323130-2311120110022000-0133102233323202-0010023330211100-0103022121322221-2100302110333313-1113032130110303-0220103231020011) |
| `egress_rules.ip_prefix_set` | [egress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-1331001321030200-0030121330120112-2321321022021031-1031011022031300-3302303122333213-1102022332103233-3211321303121103-0023223111311222) |
| `egress_rules.ip_prefix_set.ref` | [egress_rules.ip_prefix_set.ref](data-sources--network_policy_view--reference--group-001.md#canonical-3133120030310023-2332021013303131-2031310003001213-1112331211020210-3031121211323301-0120022110010322-1023131133002223-0101111213121021) |
| `egress_rules.ip_prefix_set.ref.kind` | [egress_rules.ip_prefix_set.ref.kind](data-sources--network_policy_view--reference--group-001.md#canonical-2110113203222003-3222010330201233-1012021311001200-3121312130200332-2003103202333333-3231230222132013-3301213233302003-1233102213321102) |
| `egress_rules.ip_prefix_set.ref.name` | [egress_rules.ip_prefix_set.ref.name](data-sources--network_policy_view--reference--group-001.md#canonical-3211001301121203-3023312033300200-2113100202301211-3321211003033103-1012113212330210-0000132022322130-2010312031121003-2322203012133130) |
| `egress_rules.ip_prefix_set.ref.namespace` | [egress_rules.ip_prefix_set.ref.namespace](data-sources--network_policy_view--reference--group-001.md#canonical-0030230233113221-0303322120030000-1110223300121020-0232120123222110-0203220202002000-3201230333320332-1313120020030100-3222023233321112) |
| `egress_rules.ip_prefix_set.ref.tenant` | [egress_rules.ip_prefix_set.ref.tenant](data-sources--network_policy_view--reference--group-001.md#canonical-1123132303103132-0211210102133103-3003021020100320-3000102331201110-0103201033230203-1112102030123103-3020323322321222-0000223302333330) |
| `egress_rules.ip_prefix_set.ref.uid` | [egress_rules.ip_prefix_set.ref.uid](data-sources--network_policy_view--reference--group-001.md#canonical-2023132230032131-1033301213220302-3013003100030113-3222232120133010-1233123121213333-1032001121023300-1011131113033101-2322210131030200) |
| `egress_rules.label_matcher` | [egress_rules.label_matcher](data-sources--network_policy_view--reference--group-001.md#canonical-1123103330112233-3123012000323111-3122320002302323-1103003012103303-0303313020312010-3003133022300302-1001020113020020-1101101000320100) |
| `egress_rules.label_matcher.keys` | [egress_rules.label_matcher.keys](data-sources--network_policy_view--reference--group-001.md#canonical-3233330330233301-0311132323031302-0011231130331232-3131002011112222-2211231000132031-3112211103320200-0113122031022102-2031333210332111) |
| `egress_rules.label_selector` | [egress_rules.label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-2112113302032332-2332013313133332-3203032320332002-1133111212101203-2320120030103300-1201113003012203-0113211013302132-0200321210323230) |
| `egress_rules.label_selector.expressions` | [egress_rules.label_selector.expressions](data-sources--network_policy_view--reference--group-001.md#canonical-0332001021330121-0303201320110011-2312313312220131-3331023310222222-2032320210301122-3100100230002110-0033001213333230-1323110002211031) |
| `egress_rules.metadata` | [egress_rules.metadata](data-sources--network_policy_view--reference--group-001.md#canonical-0033200123132322-1012110012010101-2210211211320001-1202310100232121-2323103201323100-0232302212121322-0323102012232113-2032323311022132) |
| `egress_rules.metadata.description_spec` | [egress_rules.metadata.description_spec](data-sources--network_policy_view--reference--group-001.md#canonical-0202003022031032-0131330020303110-3313203312102333-0103123233021223-1321030022003222-2020231131312202-2320211020103320-2111033001331111) |
| `egress_rules.metadata.name` | [egress_rules.metadata.name](data-sources--network_policy_view--reference--group-001.md#canonical-2322121111031231-3322111332012112-2331210320303022-1210132311033320-3330102313201120-3030100310123013-0131320221023112-1131021113222320) |
| `egress_rules.outside_endpoints` | [egress_rules.outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-2122232203000202-2032130131301110-3002003111321011-0221032031323332-1022323130001301-2011110022122133-0032032201312230-0320103321222203) |
| `egress_rules.prefix_list` | [egress_rules.prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-2101133031032302-3332123110333122-3122203033331122-2013011212023302-2223020333201112-2211110032312313-1133020002220210-2103332213233020) |
| `egress_rules.prefix_list.prefixes` | [egress_rules.prefix_list.prefixes](data-sources--network_policy_view--reference--group-001.md#canonical-2230112310221321-0330301101013132-1331100010310221-3000211232012213-1222200313011200-1322121211023211-2112133322032112-2321032031312312) |
| `egress_rules.protocol_port_range` | [egress_rules.protocol_port_range](data-sources--network_policy_view--reference--group-001.md#canonical-0112212333132020-1100131000320202-1310311211232022-0012230030320113-0330310201201312-0321221322323231-0302232213311333-3002001232300233) |
| `egress_rules.protocol_port_range.port_ranges` | [egress_rules.protocol_port_range.port_ranges](data-sources--network_policy_view--reference--group-001.md#canonical-0231111230202323-0323322230131020-2222212210331301-2130313200210223-0233011332132201-0300223221321230-2023000103312123-2131223201223233) |
| `egress_rules.protocol_port_range.protocol` | [egress_rules.protocol_port_range.protocol](data-sources--network_policy_view--reference--group-001.md#canonical-3011121303012233-0203111132331100-0202202220110030-1331110323031032-3010223230331232-0311020100012120-1301323233131101-3122311011001222) |
| `endpoint` | [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2133300101133331-2011203300030330-0300111120220132-1113223003100030-0321032032131221-3232201000311001-3032002132231123-0331331310031213) |
| `endpoint.any` | [endpoint.any](data-sources--network_policy_view--reference--group-001.md#canonical-0003303203102212-1313100020211301-3013200231103132-0120312210030013-0030333300032033-1000232311130223-2112323313103111-0121203103133032) |
| `endpoint.inside_endpoints` | [endpoint.inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-2001110021201323-1121211011322122-0103132012111333-1210231332121103-2231033120133001-1321233023332011-3131032013311232-1303020011313322) |
| `endpoint.label_selector` | [endpoint.label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-3200303103301110-3302231110332302-3331000110012202-2320313233100320-3232303200020302-1232301322223112-1111313013202122-0003313221102310) |
| `endpoint.label_selector.expressions` | [endpoint.label_selector.expressions](data-sources--network_policy_view--reference--group-001.md#canonical-1120330110122121-1330131120010001-0332200120203030-1230013301212210-2232010011323031-2302131230011133-3212231112211311-2211312100011202) |
| `endpoint.outside_endpoints` | [endpoint.outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-2001121112001120-2310221020010032-3113233331332302-1333323002232111-1112220102002322-3020310312123201-2230122121322312-1133110011230322) |
| `endpoint.prefix_list` | [endpoint.prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-0002300303300122-3030331030031221-2130231032032221-3132112320123312-3112012110302033-1102333023302023-0320122223023020-3233322310223031) |
| `endpoint.prefix_list.prefixes` | [endpoint.prefix_list.prefixes](data-sources--network_policy_view--reference--group-001.md#canonical-0232111023211330-2030012022301212-2010303110311220-1322221023102310-0231112330201101-0102033102301021-2020321220111221-0333121121021012) |
| `id` | [id](data-sources--network_policy_view--reference--group-001.md#canonical-1212231112313013-0030020201312101-2232200333002003-3021222231101031-1222031002201233-0201230231121301-0102120200003032-1000230233110322) |
| `ingress_rules` | [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-3113032312302112-3033313213100000-3101221020021133-3302003211300123-1300200230011322-1321121031030310-2112123212201201-1020321103322230) |
| `ingress_rules.action` | [ingress_rules.action](data-sources--network_policy_view--reference--group-001.md#canonical-0223013212010310-2213321322221112-0112132110221332-1220331133130001-3110021221112220-1212220211222023-0000022012130123-3330212022202122) |
| `ingress_rules.adv_action` | [ingress_rules.adv_action](data-sources--network_policy_view--reference--group-001.md#canonical-2310022020022201-1311110213123132-2121013313302211-0110032110011313-0030323021221203-3101022201233002-2310131022020220-3312301001121200) |
| `ingress_rules.adv_action.action` | [ingress_rules.adv_action.action](data-sources--network_policy_view--reference--group-001.md#canonical-0111100203321331-0201221011331102-3032102220303113-2200222311132331-3302000303112103-0330333330120313-3203310233020310-3233130301000220) |
| `ingress_rules.all_tcp_traffic` | [ingress_rules.all_tcp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-1221003022212223-3230312303323013-3032100022030321-2103310330112322-2013003013331223-3121201322223000-1001030300212032-2313202021120202) |
| `ingress_rules.all_traffic` | [ingress_rules.all_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-1132222022101100-3232002100020221-2023022123233002-1011132020032011-3101322320322232-2101001333120120-2121030320000032-3022331221230302) |
| `ingress_rules.all_udp_traffic` | [ingress_rules.all_udp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-0030102320000212-3212200013330112-2200021202030003-0131131112230003-0022310100232302-1211202123133221-1033113132230021-0033311010333200) |
| `ingress_rules.any` | [ingress_rules.any](data-sources--network_policy_view--reference--group-001.md#canonical-1122032021000011-2120033132021133-2302203132021113-3013311200230130-1220013130000302-3030012331023230-2301120333213222-2332031311102112) |
| `ingress_rules.applications` | [ingress_rules.applications](data-sources--network_policy_view--reference--group-001.md#canonical-3113030113133301-3222100112201023-3212323131031113-2110321111333303-0131001322220023-3122133331213033-1003213320023310-3130002231222302) |
| `ingress_rules.applications.applications` | [ingress_rules.applications.applications](data-sources--network_policy_view--reference--group-001.md#canonical-2323331102232311-2210033313221213-2310021121121232-3122010203030103-0201120303013021-1303221202311231-2002111113323322-0022221110311231) |
| `ingress_rules.inside_endpoints` | [ingress_rules.inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-1311212122331323-1332203103121012-1113202333212033-1111000233131110-3101322301023332-1322033322121000-1203203331332000-2032102002332300) |
| `ingress_rules.ip_prefix_set` | [ingress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-3331320012003021-2310020312333030-2301332201301303-2102212101212213-3220110121333112-2123122103022331-1302211230332332-0011311022200323) |
| `ingress_rules.ip_prefix_set.ref` | [ingress_rules.ip_prefix_set.ref](data-sources--network_policy_view--reference--group-001.md#canonical-0111302112300013-0003132032211121-2120220130113112-0112332312112133-2130211100000210-1032233322022233-1332313312311313-3003113011112002) |
| `ingress_rules.ip_prefix_set.ref.kind` | [ingress_rules.ip_prefix_set.ref.kind](data-sources--network_policy_view--reference--group-001.md#canonical-2233112311313022-1000120331333301-0212221023302002-0013000333020312-3122312302331311-2130101110300102-1030221003312131-2100020030220011) |
| `ingress_rules.ip_prefix_set.ref.name` | [ingress_rules.ip_prefix_set.ref.name](data-sources--network_policy_view--reference--group-001.md#canonical-0322332331101021-3331020223220222-3233232020011200-2033300032323310-3020032133021312-3220033022302133-1332300321301302-0110332113202332) |
| `ingress_rules.ip_prefix_set.ref.namespace` | [ingress_rules.ip_prefix_set.ref.namespace](data-sources--network_policy_view--reference--group-001.md#canonical-3030100320200323-2010303122023013-0103212003223201-0131303110101301-2113301133210210-0102123221022020-0002210000021133-1300302110021112) |
| `ingress_rules.ip_prefix_set.ref.tenant` | [ingress_rules.ip_prefix_set.ref.tenant](data-sources--network_policy_view--reference--group-001.md#canonical-3000333012002132-1023223311130101-1301223201100322-3100230013231211-2001330133310202-0000010202003003-2323121121300303-3010030212010032) |
| `ingress_rules.ip_prefix_set.ref.uid` | [ingress_rules.ip_prefix_set.ref.uid](data-sources--network_policy_view--reference--group-001.md#canonical-0121233302302011-2023301121022231-0331000220121100-1033232212031000-0103322030223122-0010313222000032-2100121111210011-0232223202022022) |
| `ingress_rules.label_matcher` | [ingress_rules.label_matcher](data-sources--network_policy_view--reference--group-001.md#canonical-3102000212100033-3030132311311203-1012100130331301-0233333020202013-2200212201103220-2130332020123110-1321110233321223-3132201133303313) |
| `ingress_rules.label_matcher.keys` | [ingress_rules.label_matcher.keys](data-sources--network_policy_view--reference--group-001.md#canonical-1011213233010311-1233033233212130-3110321220032333-3222313101231320-2331123130320312-2320310310013323-3130232200230023-1022231031030030) |
| `ingress_rules.label_selector` | [ingress_rules.label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-2110021000201020-2031310000201302-1203210030010100-1012021212310010-0300010001310133-1300200120331001-0200033121331103-2130131303330302) |
| `ingress_rules.label_selector.expressions` | [ingress_rules.label_selector.expressions](data-sources--network_policy_view--reference--group-001.md#canonical-0302032121101130-3023012223231033-3301110111010131-1202032002322220-1222131011120323-3010213202311023-3310111021230101-2022112020020123) |
| `ingress_rules.metadata` | [ingress_rules.metadata](data-sources--network_policy_view--reference--group-001.md#canonical-2133131322333313-1103002021110002-1032203233123021-2100312221301322-0022200213311002-3300322002120231-3021313001230213-3312032112233331) |
| `ingress_rules.metadata.description_spec` | [ingress_rules.metadata.description_spec](data-sources--network_policy_view--reference--group-001.md#canonical-1332201332333032-0120110032313320-1323321020333332-0131031130212312-0112232023121212-3110031031020020-1012310320031321-1133122312313023) |
| `ingress_rules.metadata.name` | [ingress_rules.metadata.name](data-sources--network_policy_view--reference--group-001.md#canonical-2031111132010112-0320112130302302-2201130211232021-1302100031331203-3330203300000202-3031033122022013-2312230013312333-1303302332133133) |
| `ingress_rules.outside_endpoints` | [ingress_rules.outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-3011033200020130-1002102113212113-2010010331000320-2100203200232232-1211201100121112-3100333030313221-3023010101323222-1012232032310030) |
| `ingress_rules.prefix_list` | [ingress_rules.prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-1323303230320033-1311130331030120-0212202333131210-2000131131202022-2220100331211201-1213033232013132-0201120331122310-2213010323022033) |
| `ingress_rules.prefix_list.prefixes` | [ingress_rules.prefix_list.prefixes](data-sources--network_policy_view--reference--group-001.md#canonical-2011313131121021-1322010123002022-1231203011303133-2301322231030220-2221010331311223-2301230210230012-1330220311300113-0131302011020111) |
| `ingress_rules.protocol_port_range` | [ingress_rules.protocol_port_range](data-sources--network_policy_view--reference--group-001.md#canonical-2232122101222311-2322222102310230-2210101230021302-3320211001322123-0200130200303203-1120310310000210-2020311100232230-2212232203230023) |
| `ingress_rules.protocol_port_range.port_ranges` | [ingress_rules.protocol_port_range.port_ranges](data-sources--network_policy_view--reference--group-001.md#canonical-0121211002312013-1023113102332022-2101232233002101-0022001332302020-0120000112133132-2332002201210303-1301320020032222-2331300222100122) |
| `ingress_rules.protocol_port_range.protocol` | [ingress_rules.protocol_port_range.protocol](data-sources--network_policy_view--reference--group-001.md#canonical-2230221211233233-1330112133203220-3233020313130320-1123122132303132-3321111300102021-0110101210023221-2011102001330233-2122320213010001) |
| `labels` | [labels](data-sources--network_policy_view--reference--group-001.md#canonical-3301230033121220-3133023010111112-1313113310031131-0111023010223223-2020130013323223-3003021233232111-1123001031300030-0113111011230103) |
| `name` | [name](data-sources--network_policy_view--reference--group-001.md#canonical-2031220122032332-1013330130211303-1312003231021332-0100013003020011-1021331110221202-2112321133331022-3210122133121232-0222111311010022) |
| `namespace` | [namespace](data-sources--network_policy_view--reference--group-001.md#canonical-1003123010310220-0211121011003203-3021002303232312-1020201211121320-3223202113220033-3011332123002130-2130003322212001-3332322213020010) |

<a id="canonical-3003300313101223-3120312211021223-2113323333310102-1222011310223031-3031123121233023-3031331122112200-1120030201101233-1031331223330101"></a>

## Next pages — Property reference / 211210010121 / 11

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303131230000110-0200103033323310-3133201133321223-1033321021123232-0213102013023121-1002032302003002-1233311311311112-1132320331233101"></a>

## egress_rules — egress_rules / 321131321110 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- egress_rules

<a id="canonical-0210012222310121-0203301010320030-1112112210033300-3320100030220311-2030020333321220-1212031200230132-1201211202302112-1220000120221030"></a>

Type: `"list"`. Computed.

Ordered list of rules applied to connections from policy endpoints.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-2333113033003321-2233301003020031-1131202311231011-3301013231123113-3130101223333321-3130010202232002-0013030002323013-3331233100200301"></a>

## Direct properties — egress_rules / 321131321110 / 3

<a id="canonical-0103320330222110-2300210011021101-0301021221000332-0120112110222003-1123331023130300-3112302011121010-3301112311222331-1212131321013112"></a>

<a id="canonical-1312102311111101-1033212321200220-3031022032031230-2000020322220222-2321122021303112-0100111031233333-1200113312301032-2033011301131031"></a>

## action property — egress_rules / 321131321110 / 4

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [adv_action](data-sources--network_policy_view--reference--group-001.md#canonical-3033020000112100-2221203131311100-0310103311321113-0300133013210303-0210001102310321-1332131003313012-2200310312001013-0220033302332332): complete subsection reference.

- [all_tcp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-3131220310002121-0333202203303100-0220020120133330-3201130223321132-3223132233123212-0321231210313322-2010320122231322-0311232310301013): complete subsection reference.

- [all_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-3032300120113323-3310112203301102-3110003000102022-1123313020113303-0321223321230103-2003300330302011-1201213330103311-1201001132112301): complete subsection reference.

- [all_udp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-0200311213333321-1330232333122100-3230101223121111-3222003211313000-1321321112022031-0010003311313210-1312021230100313-0302022003312333): complete subsection reference.

- [any](data-sources--network_policy_view--reference--group-001.md#canonical-1332112231101121-2133231101212101-0022022313312300-2220003013000111-2131001023111033-0031223320010103-1200310110021033-0021030221330323): complete subsection reference.

- [applications](data-sources--network_policy_view--reference--group-001.md#canonical-0230320310311312-2233000100332133-2332231001331212-2131233221220022-0301113223132121-3201121201103111-1303300131013332-1110122212323112): complete subsection reference.

- [inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-3102312011022130-0311203322112022-0111001223000212-0100220133112133-0201030303303222-3313120232112133-3113102101130133-2212113132333302): complete subsection reference.

- [ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-2300101113211222-0321120310000113-0300010211332333-0101331103203200-1301230331121333-0331330012300000-3010030201232013-0321230110202323): complete subsection reference.

- [label_matcher](data-sources--network_policy_view--reference--group-001.md#canonical-0323131022301030-3200102201333212-1103330201131221-3130311110123303-2331232101203231-1333110232211320-0200213010123130-3203001212101130): complete subsection reference.

- [label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-2302031122222020-1121201013111113-3203002211310103-0223203322231131-2313202131023232-2000330211100220-1003130013232022-0221130003232320): complete subsection reference.

- [metadata](data-sources--network_policy_view--reference--group-001.md#canonical-1322322222202103-1231102103200131-3210202132100221-1323301120222012-3031213212221322-1132221230021120-2030230103322321-1333310001130312): complete subsection reference.

- [outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-1033003313030113-1203303121020303-0130233220121120-0201031121310003-2302130002023111-2220300030320300-1232322121310132-2320303313331113): complete subsection reference.

- [prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-1300023113210112-1312310002030100-0013030010321211-3000022210032002-2020300031132211-2300112121321330-1331020111323332-0122221331233322): complete subsection reference.

- [protocol_port_range](data-sources--network_policy_view--reference--group-001.md#canonical-1222022232132112-2220111000131211-1002311200331321-3232230300023030-2203001102301031-3120302213221033-0101011233011031-3012110023331302): complete subsection reference.

<a id="canonical-2312100123200323-1132213010313200-1001210220311321-3313031321222303-0300230133301301-1111332113323020-2202011321230003-2010303302323301"></a>

## Next pages — egress_rules / 321131321110 / 5

- [egress_rules.adv_action](data-sources--network_policy_view--reference--group-001.md#canonical-3033020000112100-2221203131311100-0310103311321113-0300133013210303-0210001102310321-1332131003313012-2200310312001013-0220033302332332)
- [egress_rules.all_tcp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-3131220310002121-0333202203303100-0220020120133330-3201130223321132-3223132233123212-0321231210313322-2010320122231322-0311232310301013)
- [egress_rules.all_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-3032300120113323-3310112203301102-3110003000102022-1123313020113303-0321223321230103-2003300330302011-1201213330103311-1201001132112301)
- [egress_rules.all_udp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-0200311213333321-1330232333122100-3230101223121111-3222003211313000-1321321112022031-0010003311313210-1312021230100313-0302022003312333)
- [egress_rules.any](data-sources--network_policy_view--reference--group-001.md#canonical-1332112231101121-2133231101212101-0022022313312300-2220003013000111-2131001023111033-0031223320010103-1200310110021033-0021030221330323)
- [egress_rules.applications](data-sources--network_policy_view--reference--group-001.md#canonical-0230320310311312-2233000100332133-2332231001331212-2131233221220022-0301113223132121-3201121201103111-1303300131013332-1110122212323112)
- [egress_rules.inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-3102312011022130-0311203322112022-0111001223000212-0100220133112133-0201030303303222-3313120232112133-3113102101130133-2212113132333302)
- [egress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-2300101113211222-0321120310000113-0300010211332333-0101331103203200-1301230331121333-0331330012300000-3010030201232013-0321230110202323)
- [egress_rules.label_matcher](data-sources--network_policy_view--reference--group-001.md#canonical-0323131022301030-3200102201333212-1103330201131221-3130311110123303-2331232101203231-1333110232211320-0200213010123130-3203001212101130)
- [egress_rules.label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-2302031122222020-1121201013111113-3203002211310103-0223203322231131-2313202131023232-2000330211100220-1003130013232022-0221130003232320)
- [egress_rules.metadata](data-sources--network_policy_view--reference--group-001.md#canonical-1322322222202103-1231102103200131-3210202132100221-1323301120222012-3031213212221322-1132221230021120-2030230103322321-1333310001130312)
- [egress_rules.outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-1033003313030113-1203303121020303-0130233220121120-0201031121310003-2302130002023111-2220300030320300-1232322121310132-2320303313331113)
- [egress_rules.prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-1300023113210112-1312310002030100-0013030010321211-3000022210032002-2020300031132211-2300112121321330-1331020111323332-0122221331233322)
- [egress_rules.protocol_port_range](data-sources--network_policy_view--reference--group-001.md#canonical-1222022232132112-2220111000131211-1002311200331321-3232230300023030-2203001102301031-3120302213221033-0101011233011031-3012110023331302)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-3033020000112100-2221203131311100-0310103311321113-0300133013210303-0210001102310321-1332131003313012-2200310312001013-0220033302332332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330011332211000-1322110020021121-3021103311220100-0011033113123022-1022111322112022-1302030121313003-3300312302113102-0032210033112212"></a>

## egress_rules.adv_action — adv_action / 110033121302 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.adv_action

<a id="canonical-2213210032120033-2331233212230021-0103012103113302-0010331132221033-1110013031320230-1101010200201231-3220332031113030-0132330100130230"></a>

Type: `"single"`. Computed.

Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and
PBRRuleAction.

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

<a id="canonical-1233032001112132-0003322013032330-3001112133033310-2333221012300033-1311023022320321-1222331101313130-1021323131332133-3113321222332012"></a>

## Direct properties — adv_action / 110033121302 / 3

<a id="canonical-1010222112000120-2130230103022001-0302003333320212-2223021012210000-0213223102203001-3011011211003010-0101101323310210-2211102100321303"></a>

<a id="canonical-1233001012231302-3000012021213220-1220101133023200-3102210003022031-3003010130120011-0112130013233310-0113312322011212-1021001333101012"></a>

## action property — adv_action / 110033121302 / 4

Type: `"string"`. Computed.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Upstream description:

Choice to choose logging or no logging This works together with option selected via
NetworkPolicyRuleAction or any other action specified x-example: (No Selection in
NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction
&#8203;+ AdvancedAction as LOG) = Log and Allow/Deny, (ALLOW/DENY in NetworkPolicyRuleAction + NOLOG in
AdvancedAction) = Allow/Deny with no log

Don't sample the traffic hitting the rule Sample the traffic hitting the rule.

Receipt-pinned upstream constraints:

```json
{
  "default": "NOLOG",
  "enum": [
    "NOLOG",
    "LOG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3310230022213022-3301012030101120-0300213311221313-0120212330312310-3231202312133223-1312332001321220-1203000220331303-1211113100333120"></a>

## Next pages — adv_action / 110033121302 / 5

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-3131220310002121-0333202203303100-0220020120133330-3201130223321132-3223132233123212-0321231210313322-2010320122231322-0311232310301013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330210023223302-1100200321221032-1221133321030023-1000332332313330-1332033103110330-1221101133113122-1101030303102310-0220013000023013"></a>

## egress_rules.all_tcp_traffic — all_tcp_traffic / 302212320331 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.all_tcp_traffic

<a id="canonical-1313110331133132-1333211021110032-2221111030023323-2200030022313122-3213221010120103-3113122323233223-0130112321200023-2002303120221320"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all tcp traffic.

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

<a id="canonical-2331331211113222-2320213112023032-2202232211022003-1120021231322133-3102220211230103-2310201002301020-1330123113100111-0201100102101001"></a>

## Direct properties — all_tcp_traffic / 302212320331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3110030122103112-3220003330310031-3323100203111023-0032332012101020-3030111122331101-3002021331210203-1333113111100133-1002303021102310"></a>

## Next pages — all_tcp_traffic / 302212320331 / 4

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-3032300120113323-3310112203301102-3110003000102022-1123313020113303-0321223321230103-2003300330302011-1201213330103311-1201001132112301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232331213123120-2311313131332302-2333300210123032-1231000133110223-0201033232210323-2013332232011120-0001213221103003-1133323003303211"></a>

## egress_rules.all_traffic — all_traffic / 212121233122 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.all_traffic

<a id="canonical-0112130123320212-3202203011303200-2102022121022221-1222121221022320-0323332013320202-0220221221210100-0123032111013321-3011130101101311"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all traffic.

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

<a id="canonical-0000112123211222-1200230131322212-3030210332201302-0110231310110022-1320333210320200-1011202331230131-3021322211020001-3003121031113333"></a>

## Direct properties — all_traffic / 212121233122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220321120030131-0310211113103232-0212131330212023-3332203200101031-1220103210212323-1120322302323221-3111122312312223-0230322033233120"></a>

## Next pages — all_traffic / 212121233122 / 4

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-0200311213333321-1330232333122100-3230101223121111-3222003211313000-1321321112022031-0010003311313210-1312021230100313-0302022003312333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212313133101233-2001210102003303-1101032023232201-0233113212232201-2023233030321210-0010021121020201-0012231112312021-3203220323033332"></a>

## egress_rules.all_udp_traffic — all_udp_traffic / 213122110123 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.all_udp_traffic

<a id="canonical-0322213201221013-2213313120310323-0203011232310232-0102021113003312-0201020013323233-2112213033033211-0200203011102003-0102033302112033"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all udp traffic.

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

<a id="canonical-3110231121000310-0011223233020300-1311013021230132-3020022232233033-2001232312201332-2202232221211001-1230031110020210-3023033221311032"></a>

## Direct properties — all_udp_traffic / 213122110123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113030212322213-3000131201222220-1133301311310121-0111002013331233-1321133023303211-0131003233011123-0111031210031300-0203122203311101"></a>

## Next pages — all_udp_traffic / 213122110123 / 4

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-1332112231101121-2133231101212101-0022022313312300-2220003013000111-2131001023111033-0031223320010103-1200310110021033-0021030221330323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312302232203213-2130310000121132-3213222000011112-2000221113101332-2321100201000022-1112130333132132-2020213322202222-1212031230321210"></a>

## egress_rules.any — any / 100030313031 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.any

<a id="canonical-2202030120110030-0203300212211122-3212322200102200-2233223123112331-1202121302100333-2310010303102102-2223301120103300-2210130031000101"></a>

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

<a id="canonical-2011302013122301-1001123113221222-3010011132332123-0033211302232022-0022013212110130-0333121232203313-3023031131110133-3000010210100220"></a>

## Direct properties — any / 100030313031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200003210230323-1202120303110210-3311111330010032-1100222210300301-1223001223120120-2312233221123222-1002120233101120-3020022313203113"></a>

## Next pages — any / 100030313031 / 4

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-0230320310311312-2233000100332133-2332231001331212-2131233221220022-0301113223132121-3201121201103111-1303300131013332-1110122212323112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220103011231323-1212322201011020-0231120303210201-0331221032000222-0330221221030012-3231211111231230-1122131110103332-1112111011132210"></a>

## egress_rules.applications — applications / 323320002321 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.applications

<a id="canonical-1223201121101002-0103322132222021-2101222100231011-3010000230310310-0200031330330301-2320203100211102-2122030122211310-3323010203131103"></a>

Type: `"single"`. Computed.

Configuration parameter for applications.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="canonical-0120331011031211-2121003121021030-3321011132110321-1232303120213000-1322333223033000-0221311231333201-3023122020133223-1321221210031131"></a>

## Direct properties — applications / 323320002321 / 3

<a id="canonical-2100233213332031-1120313112230313-3203220310121023-2023031230202331-0000023220001010-3130220212021020-2221120131100102-0313033231103232"></a>

<a id="canonical-1010023223013002-0202233220310222-3223123323111320-3303212001311321-0231101332033100-2130030312233113-0230332230103112-1033123200302103"></a>

## applications property — applications / 323320002321 / 4

Type: `["list", "string"]`. Computed.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="canonical-2002200221033231-3210312013033330-1323221121203312-1302312121000112-0122123103030222-3012000223030000-2011311133023111-2033021031302332"></a>

## Next pages — applications / 323320002321 / 5

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-3102312011022130-0311203322112022-0111001223000212-0100220133112133-0201030303303222-3313120232112133-3113102101130133-2212113132333302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201222320331230-0010303101303230-3231330131303113-3212002122120121-2230213201001230-1110313302210233-1300102100003003-0032221203200013"></a>

## egress_rules.inside_endpoints — inside_endpoints / 002230222211 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.inside_endpoints

<a id="canonical-0310302311323130-2311120110022000-0133102233323202-0010023330211100-0103022121322221-2100302110333313-1113032130110303-0220103231020011"></a>

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

<a id="canonical-1021221300112231-1333220313311101-3030002310113213-0123013010213202-0123100001233130-3202130103003332-0301303032121102-2312322103003122"></a>

## Direct properties — inside_endpoints / 002230222211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000012103311012-3030021033112313-1010202202213332-0210210202310022-3212023331110000-0202310121200031-1302302121130023-0203120303131301"></a>

## Next pages — inside_endpoints / 002230222211 / 4

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-2300101113211222-0321120310000113-0300010211332333-0101331103203200-1301230331121333-0331330012300000-3010030201232013-0321230110202323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001211203110232-3332211122212120-2123230213132303-1013030003001320-1203131112022322-3033122113113201-0220122210211122-3102332121120310"></a>

## egress_rules.ip_prefix_set — ip_prefix_set / 201211030100 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.ip_prefix_set

<a id="canonical-1331001321030200-0030121330120112-2321321022021031-1031011022031300-3302303122333213-1102022332103233-3211321303121103-0023223111311222"></a>

Type: `"single"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

<a id="canonical-3123223012210303-3123122011111300-3233210123131020-2021233201011301-3123330113013211-0020301033100203-2112030101112233-2033301032111200"></a>

## Direct properties — ip_prefix_set / 201211030100 / 3

- [ref](data-sources--network_policy_view--reference--group-001.md#canonical-0122032033213213-2330231233312003-0220213031123200-3323121112320310-1102312230201223-0301000311321031-3203111021233023-1223120313201231): complete subsection reference.

<a id="canonical-1203223131312321-1130020103131003-3223223111121121-0231011210010012-0302102123202131-1322013231323123-0000222131321022-0311312203120021"></a>

## Next pages — ip_prefix_set / 201211030100 / 4

- [egress_rules.ip_prefix_set.ref](data-sources--network_policy_view--reference--group-001.md#canonical-0122032033213213-2330231233312003-0220213031123200-3323121112320310-1102312230201223-0301000311321031-3203111021233023-1223120313201231)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-0122032033213213-2330231233312003-0220213031123200-3323121112320310-1102312230201223-0301000311321031-3203111021233023-1223120313201231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121032220033301-2020123211312313-3300200200020031-2011322001213331-0021003023120023-1320033002120100-2302213111130032-0013212112001323"></a>

## egress_rules.ip_prefix_set.ref — ref / 101301211130 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [egress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-2300101113211222-0321120310000113-0300010211332333-0101331103203200-1301230331121333-0331330012300000-3010030201232013-0321230110202323)
- egress_rules.ip_prefix_set.ref

<a id="canonical-3133120030310023-2332021013303131-2031310003001213-1112331211020210-3031121211323301-0120022110010322-1023131133002223-0101111213121021"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1032200202012003-0103311022320330-2212201312200123-3201201321333332-0122223032011100-1233233102011032-0130213032031211-3203133202331002"></a>

## Direct properties — ref / 101301211130 / 3

<a id="canonical-2110113203222003-3222010330201233-1012021311001200-3121312130200332-2003103202333333-3231230222132013-3301213233302003-1233102213321102"></a>

<a id="canonical-3031110131132211-3123022030231333-3301201332023301-3112233020300300-1320223031120230-2231103130233031-0331311202001123-2203133120223032"></a>

## kind property — ref / 101301211130 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-3211001301121203-3023312033300200-2113100202301211-3321211003033103-1012113212330210-0000132022322130-2010312031121003-2322203012133130"></a>

<a id="canonical-2011131123111333-1301130332023223-0032110130011333-3121132230010023-3203330023131302-3223320300221112-2103001020030321-0122111010320320"></a>

## name property — ref / 101301211130 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-0030230233113221-0303322120030000-1110223300121020-0232120123222110-0203220202002000-3201230333320332-1313120020030100-3222023233321112"></a>

<a id="canonical-0320203211022013-0321212300123103-0010322213133213-3300130101203213-3212310122030210-0232013232011312-0212121023233210-2000111102203130"></a>

## namespace property — ref / 101301211130 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-1123132303103132-0211210102133103-3003021020100320-3000102331201110-0103201033230203-1112102030123103-3020323322321222-0000223302333330"></a>

<a id="canonical-1023303102021321-1122122320002003-3210312203122003-3131122030332101-3320123113010210-2332120031302311-2011300201330113-2021003111330231"></a>

## tenant property — ref / 101301211130 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2023132230032131-1033301213220302-3013003100030113-3222232120133010-1233123121213333-1032001121023300-1011131113033101-2322210131030200"></a>

<a id="canonical-1032232012123130-3300032111030121-3231211023023200-3032013123302312-1010021113300032-3011232133302303-0231003311031111-0213012023133312"></a>

## uid property — ref / 101301211130 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-2103330313100202-1021131321103131-2003121321023300-1320000031111113-1333121020020013-2222023312131011-2030310133132312-3313322002230230"></a>

## Next pages — ref / 101301211130 / 9

- [egress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-2300101113211222-0321120310000113-0300010211332333-0101331103203200-1301230331121333-0331330012300000-3010030201232013-0321230110202323)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-0323131022301030-3200102201333212-1103330201131221-3130311110123303-2331232101203231-1333110232211320-0200213010123130-3203001212101130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013320203230132-2303320011310333-2201302323320012-2102200132103303-3120332312122020-1100201223312333-3020330230302331-0302203220312032"></a>

## egress_rules.label_matcher — label_matcher / 200101011233 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.label_matcher

<a id="canonical-1123103330112233-3123012000323111-3122320002302323-1103003012103303-0303313020312010-3003133022300302-1001020113020020-1101101000320100"></a>

Type: `"single"`. Computed.

Label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Upstream description:

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

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

<a id="canonical-2122003310010110-0112200311033320-0221333130011331-2002302122311203-0131232212332302-0301303131230111-1023130131120123-1220202232031332"></a>

## Direct properties — label_matcher / 200101011233 / 3

<a id="canonical-3233330330233301-0311132323031302-0011231130331232-3131002011112222-2211231000132031-3112211103320200-0113122031022102-2031333210332111"></a>

<a id="canonical-3310122313323222-0302331130330221-0123121322311332-2300210323101210-1002211001202121-3121032202301103-3013200221323233-0231010100123011"></a>

## keys property — label_matcher / 200101011233 / 4

Type: `["list", "string"]`. Computed.

The list of label key names that have to match.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2232331120332332-1122113102300102-0113010121202302-3212133332212111-1121333021031231-1110200223100311-1000103032101213-3331321312021032"></a>

## Next pages — label_matcher / 200101011233 / 5

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-2302031122222020-1121201013111113-3203002211310103-0223203322231131-2313202131023232-2000330211100220-1003130013232022-0221130003232320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303111102013231-2211113331220120-0212000300022030-1011303111210103-1032032021001230-3121203112323133-3223033000302212-0300001033330223"></a>

## egress_rules.label_selector — label_selector / 121322210321 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.label_selector

<a id="canonical-2112113302032332-2332013313133332-3203032320332002-1133111212101203-2320120030103300-1201113003012203-0113211013302132-0200321210323230"></a>

Type: `"single"`. Computed.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-2313010210313011-2220113001133102-2010323212211133-1020322313030223-3131300201021001-1330100002100322-2133002020013112-2231302310213303"></a>

## Direct properties — label_selector / 121322210321 / 3

<a id="canonical-0332001021330121-0303201320110011-2312313312220131-3331023310222222-2032320210301122-3100100230002110-0033001213333230-1323110002211031"></a>

<a id="canonical-1200001132313112-3032010102212011-1031331202033321-2332112101011131-2023201002323112-0131203203210133-1023312133230121-1111333021222001"></a>

## expressions property — label_selector / 121322210321 / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2221211332122232-0123301130311322-2023011222310100-3333221012021213-0231201122120201-0202320332231112-2001100132122123-1233111123301222"></a>

## Next pages — label_selector / 121322210321 / 5

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-1322322222202103-1231102103200131-3210202132100221-1323301120222012-3031213212221322-1132221230021120-2030230103322321-1333310001130312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131103022213301-2003121003223002-3033020332212102-2103312100101033-0123211031330113-1223313202000133-2232023220032021-3220231123111220"></a>

## egress_rules.metadata — metadata / 313031102312 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.metadata

<a id="canonical-0033200123132322-1012110012010101-2210211211320001-1202310100232121-2323103201323100-0232302212121322-0323102012232113-2032323311022132"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-3311313221213103-2011102310311002-2123003020032003-2203011230132322-0232200220002310-2100032031232010-1210011331200313-1222331223112332"></a>

## Direct properties — metadata / 313031102312 / 3

<a id="canonical-0202003022031032-0131330020303110-3313203312102333-0103123233021223-1321030022003222-2020231131312202-2320211020103320-2111033001331111"></a>

<a id="canonical-3303112132123323-2101323021202122-1313333102203220-1332103131031110-1300212002011030-3300323101101220-0323011312300032-0030310101331101"></a>

## description_spec property — metadata / 313031102312 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2322121111031231-3322111332012112-2331210320303022-1210132311033320-3330102313201120-3030100310123013-0131320221023112-1131021113222320"></a>

<a id="canonical-3322210222000010-2021203212012313-2211130220333202-1223330302022222-1112130231032110-0123220011023111-0310200202332321-1312222032320100"></a>

## name property — metadata / 313031102312 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-3120030230222211-1011030012211102-1023311323001330-1333012231100232-1023012302002022-3031000121123001-2112212210220022-3113000200331323"></a>

## Next pages — metadata / 313031102312 / 6

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-1033003313030113-1203303121020303-0130233220121120-0201031121310003-2302130002023111-2220300030320300-1232322121310132-2320303313331113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303132110130311-0220012002212113-1210321211310222-1001311113210322-1211022102133120-2210112322221212-3122312212203112-1022102200201021"></a>

## egress_rules.outside_endpoints — outside_endpoints / 010133010031 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.outside_endpoints

<a id="canonical-2122232203000202-2032130131301110-3002003111321011-0221032031323332-1022323130001301-2011110022122133-0032032201312230-0320103321222203"></a>

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

<a id="canonical-3103012133211132-0312233113302021-0022120130113000-0103111123111201-2022303010133111-3000233330123112-2032101313302212-2210031101120320"></a>

## Direct properties — outside_endpoints / 010133010031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333330222223020-3211303232023220-2012022002300330-1133020322002310-2020013021222031-3320031302213032-2122121103221320-2211201021122232"></a>

## Next pages — outside_endpoints / 010133010031 / 4

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-1300023113210112-1312310002030100-0013030010321211-3000022210032002-2020300031132211-2300112121321330-1331020111323332-0122221331233322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231123031112131-3221222220130222-0102022332123000-2111102002120312-3023012020221133-2012300223103212-1132133123210202-3101020313223110"></a>

## egress_rules.prefix_list — prefix_list / 031330230232 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.prefix_list

<a id="canonical-2101133031032302-3332123110333122-3122203033331122-2013011212023302-2223020333201112-2211110032312313-1133020002220210-2103332213233020"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-1011223023100132-2101312133212033-3233023212122232-1120333230111203-2133203113302000-2201301330321002-1031113100111310-2110120200122013"></a>

## Direct properties — prefix_list / 031330230232 / 3

<a id="canonical-2230112310221321-0330301101013132-1331100010310221-3000211232012213-1222200313011200-1322121211023211-2112133322032112-2321032031312312"></a>

<a id="canonical-2020132223323030-0221133211323213-0022210001101020-1032010130332130-2220311112110130-0023033102323031-2310212000313303-0210211003031223"></a>

## prefixes property — prefix_list / 031330230232 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0322012313033320-2302131331320220-3000303100330213-1200222132200010-0020311330123011-2030123322012322-3200102021003101-2022021020301022"></a>

## Next pages — prefix_list / 031330230232 / 5

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-1222022232132112-2220111000131211-1002311200331321-3232230300023030-2203001102301031-3120302213221033-0101011233011031-3012110023331302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302110023220112-0232230310103013-1033000301130012-3220330233320000-2022312201200332-3223230332000313-3031010311310012-1020101013320020"></a>

## egress_rules.protocol_port_range — protocol_port_range / 221232200213 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.protocol_port_range

<a id="canonical-0112212333132020-1100131000320202-1310311211232022-0012230030320113-0330310201201312-0321221322323231-0302232213311333-3002001232300233"></a>

Type: `"single"`. Computed.

Protocol and Port. Protocol and Port ranges.

Upstream description:

Protocol and Port ranges.

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

<a id="canonical-3323221123023002-0211111121110312-1231102222321202-2312230311203100-0031301323321031-2113102023333310-2312210201001021-3303033011333201"></a>

## Direct properties — protocol_port_range / 221232200213 / 3

<a id="canonical-0231111230202323-0323322230131020-2222212210331301-2130313200210223-0233011332132201-0300223221321230-2023000103312123-2131223201223233"></a>

<a id="canonical-3320201303022123-3310323021300020-0030003102232032-1301211232031123-0131011221210120-3011133111023123-1122220010303022-0231321211331013"></a>

## port_ranges property — protocol_port_range / 221232200213 / 4

Type: `["list", "string"]`. Computed.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-3011121303012233-0203111132331100-0202202220110030-1331110323031032-3010223230331232-0311020100012120-1301323233131101-3122311011001222"></a>

<a id="canonical-2110320012121312-1013000312023033-1123233333311332-0032103103201010-2120201102020003-0233132332232131-1331310213301023-3033313001020203"></a>

## protocol property — protocol_port_range / 221232200213 / 5

Type: `"string"`. Computed.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-0011210110022121-2110202132000022-0212301030211210-1330320111030130-1002233330131122-2002230100112220-3223333121013300-2310310320120103"></a>

## Next pages — protocol_port_range / 221232200213 / 6

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113233021202000-0123031213001103-0103012233301200-0032113211132223-1130030001222211-1100100211320002-0201031303110230-0102233303211032"></a>

## endpoint — endpoint / 102220211131 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- endpoint

<a id="canonical-2133300101133331-2011203300030330-0300111120220132-1113223003100030-0321032032131221-3232201000311001-3032002132231123-0331331310031213"></a>

Type: `"single"`. Computed.

Shape of the endpoint choices for a view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-endpoint_choice": "[\"any\",\"inside_endpoints\",\"label_selector\",\"outside_endpoints\",\"prefix_list\"]"
}
```

<a id="canonical-1331311121321123-0210013322213021-0101130313001310-2320222020013022-1233330100022003-3100000002333032-3130102332132232-2102003223233111"></a>

## Direct properties — endpoint / 102220211131 / 3

- [any](data-sources--network_policy_view--reference--group-001.md#canonical-0131333302333112-2310121311021213-2211101133233031-2220103131121112-0310103131133202-3330221100323231-0113133000021230-0132230333033333): complete subsection reference.

- [inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-2202321203030200-3022202322001032-3123200130203231-3211312133323331-2231003332000102-0323103023332220-0032132213310313-0302132320021122): complete subsection reference.

- [label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-1030323021230021-0313132002133212-2033222012202013-2021133100011213-1210303022033023-1221030213232212-1303110011330202-0223303322201010): complete subsection reference.

- [outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-3220110110203232-2020332111132301-2002312333231111-2330322032112311-1122122222212222-0211201023101222-0101312120133001-0130221313012322): complete subsection reference.

- [prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-1302102000330332-3102200010120011-1333131331310032-1003202300332230-1123221331302211-2211011020023102-1020300023010231-3211120010200310): complete subsection reference.

<a id="canonical-3001321220123111-1101222033301223-3122312222122110-0032011131333201-1311032210113312-1103332122210220-3103211222313310-3233301120023301"></a>

## Next pages — endpoint / 102220211131 / 4

- [endpoint.any](data-sources--network_policy_view--reference--group-001.md#canonical-0131333302333112-2310121311021213-2211101133233031-2220103131121112-0310103131133202-3330221100323231-0113133000021230-0132230333033333)
- [endpoint.inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-2202321203030200-3022202322001032-3123200130203231-3211312133323331-2231003332000102-0323103023332220-0032132213310313-0302132320021122)
- [endpoint.label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-1030323021230021-0313132002133212-2033222012202013-2021133100011213-1210303022033023-1221030213232212-1303110011330202-0223303322201010)
- [endpoint.outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-3220110110203232-2020332111132301-2002312333231111-2330322032112311-1122122222212222-0211201023101222-0101312120133001-0130221313012322)
- [endpoint.prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-1302102000330332-3102200010120011-1333131331310032-1003202300332230-1123221331302211-2211011020023102-1020300023010231-3211120010200310)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-0131333302333112-2310121311021213-2211101133233031-2220103131121112-0310103131133202-3330221100323231-0113133000021230-0132230333033333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333001212133230-3322112133223132-3311100023110322-1330113221223203-3030310221213202-1210200031213110-0003332203231122-1012002122120213"></a>

## endpoint.any — any / 213323130333 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120)
- endpoint.any

<a id="canonical-0003303203102212-1313100020211301-3013200231103132-0120312210030013-0030333300032033-1000232311130223-2112323313103111-0121203103133032"></a>

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

<a id="canonical-0021310212123330-3332111302310332-3310020103320023-0332323300002212-0132222220020112-0020133110011230-2302103023223111-3021031120020323"></a>

## Direct properties — any / 213323130333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221133332302110-2032212231101021-0212333133200311-3332320100223321-0111132112320302-3331312102333111-2202012123321120-3222131311311220"></a>

## Next pages — any / 213323130333 / 4

- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-2202321203030200-3022202322001032-3123200130203231-3211312133323331-2231003332000102-0323103023332220-0032132213310313-0302132320021122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012001130101203-3321133320011221-2001303022222232-1231030100300111-2132033231121102-0322322132122320-2213010123200200-2331210223031302"></a>

## endpoint.inside_endpoints — inside_endpoints / 321322321200 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120)
- endpoint.inside_endpoints

<a id="canonical-2001110021201323-1121211011322122-0103132012111333-1210231332121103-2231033120133001-1321233023332011-3131032013311232-1303020011313322"></a>

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

<a id="canonical-2122000121333221-2301021200300133-2111232231120012-0130321300333213-3232221100231313-2303101020131202-1020133101210012-2001203323201102"></a>

## Direct properties — inside_endpoints / 321322321200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330131120213020-3123012032332020-0032131233101032-1102011320212223-2022110103310311-2312303331013222-0202032302310022-0332111311313120"></a>

## Next pages — inside_endpoints / 321322321200 / 4

- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-1030323021230021-0313132002133212-2033222012202013-2021133100011213-1210303022033023-1221030213232212-1303110011330202-0223303322201010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020210110321033-0132212021110330-0031132323102020-3111302100300323-2211302300003010-2211130101223302-0301331123000010-2330102231010313"></a>

## endpoint.label_selector — label_selector / 003133133133 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120)
- endpoint.label_selector

<a id="canonical-3200303103301110-3302231110332302-3331000110012202-2320313233100320-3232303200020302-1232301322223112-1111313013202122-0003313221102310"></a>

Type: `"single"`. Computed.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-3012311021102011-3103331123130322-1132031222102232-3322231230011010-3013333323031220-1013233321122331-3202230233001333-2000113203130311"></a>

## Direct properties — label_selector / 003133133133 / 3

<a id="canonical-1120330110122121-1330131120010001-0332200120203030-1230013301212210-2232010011323031-2302131230011133-3212231112211311-2211312100011202"></a>

<a id="canonical-3013203333321031-2300132023213112-2323001202213232-0312320313210102-1323012031203011-0012120202033032-2232010333301232-1222300223100121"></a>

## expressions property — label_selector / 003133133133 / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1021102110201102-3223113322012113-0110211222321132-3210301301013233-3300132332323133-3122100031031100-3033013101102232-2001030020310203"></a>

## Next pages — label_selector / 003133133133 / 5

- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-3220110110203232-2020332111132301-2002312333231111-2330322032112311-1122122222212222-0211201023101222-0101312120133001-0130221313012322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122220323023122-2013113222120200-3101200213002203-2303210302133103-3033333130223202-1111200201202110-0001023022100212-0203331220113302"></a>

## endpoint.outside_endpoints — outside_endpoints / 033320000311 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120)
- endpoint.outside_endpoints

<a id="canonical-2001121112001120-2310221020010032-3113233331332302-1333323002232111-1112220102002322-3020310312123201-2230122121322312-1133110011230322"></a>

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

<a id="canonical-3333121122220310-3323311123233330-0312131302110321-1002100231011313-2013331012032032-2203030103232203-2002013011101001-3103001130300323"></a>

## Direct properties — outside_endpoints / 033320000311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002313313200210-1333322201202033-3133301320202223-2031310013300002-0233223210313203-1130321120312101-1032331212013003-0112120303223301"></a>

## Next pages — outside_endpoints / 033320000311 / 4

- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-1302102000330332-3102200010120011-1333131331310032-1003202300332230-1123221331302211-2211011020023102-1020300023010231-3211120010200310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121231022312131-1111323012003001-0123023332002101-0320113211320322-1233033103002012-0210200333112020-3333112300302311-2310000130303221"></a>

## endpoint.prefix_list — prefix_list / 212122020133 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120)
- endpoint.prefix_list

<a id="canonical-0002300303300122-3030331030031221-2130231032032221-3132112320123312-3112012110302033-1102333023302023-0320122223023020-3233322310223031"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-2201203123333103-0113213022010213-0002212233132102-3102310201103030-2022030311112302-2123032000012033-2222002022302102-1000311010320302"></a>

## Direct properties — prefix_list / 212122020133 / 3

<a id="canonical-0232111023211330-2030012022301212-2010303110311220-1322221023102310-0231112330201101-0102033102301021-2020321220111221-0333121121021012"></a>

<a id="canonical-0231301113310132-0001213022310132-2131112033131030-0110113311211233-3303313211222003-0020230113222020-0121200232200320-0110323322132300"></a>

## prefixes property — prefix_list / 212122020133 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3200320033003203-0001201123001100-3311201112013001-2331200231211211-0123000032222220-1000331020130302-1310231202223320-3000211012211032"></a>

## Next pages — prefix_list / 212122020133 / 5

- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210030312102331-0000312130100002-0332022113010313-1030212010021300-2010200201131233-2333303123232101-0013212321202001-2011113312303122"></a>

## ingress_rules — ingress_rules / 123301032220 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- ingress_rules

<a id="canonical-3113032312302112-3033313213100000-3101221020021133-3302003211300123-1300200230011322-1321121031030310-2112123212201201-1020321103322230"></a>

Type: `"list"`. Computed.

Ordered list of rules applied to connections to policy endpoints.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-1220221322002122-1231000310211113-2221012030111022-2112021101002012-3030222022312032-3021333110101233-3210203112030330-0102001111031011"></a>

## Direct properties — ingress_rules / 123301032220 / 3

<a id="canonical-0223013212010310-2213321322221112-0112132110221332-1220331133130001-3110021221112220-1212220211222023-0000022012130123-3330212022202122"></a>

<a id="canonical-1332232020102130-0320103320010331-0300003232213131-2310200130303202-0133123130023301-1132303013003011-2202132231213121-0003130303123320"></a>

## action property — ingress_rules / 123301032220 / 4

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [adv_action](data-sources--network_policy_view--reference--group-001.md#canonical-3030110321201232-1333103111021021-1013122302111032-2100203310301332-1032123102300313-1120303310012033-1002130200122222-1132100313233010): complete subsection reference.

- [all_tcp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-0333020103300021-1330000322120200-3111103330222033-0033220132100300-1013303111230120-0111031211321123-1112231000030320-3130011010032011): complete subsection reference.

- [all_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-1021232322221302-2311331101213203-3203323120012311-0013311101301120-0203030302130332-1221023203000220-3031023101131120-3303131302032320): complete subsection reference.

- [all_udp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-3130000110212210-1331231110311102-0310121222230033-0020010303000203-2121111001001220-1333212331101210-3112110332010022-0100321310332010): complete subsection reference.

- [any](data-sources--network_policy_view--reference--group-001.md#canonical-3333023301201321-0311011101001132-2200200320012032-3333120203310203-2201203330123033-0331333120123030-0020021202012011-3030203001133020): complete subsection reference.

- [applications](data-sources--network_policy_view--reference--group-001.md#canonical-3332100121000323-0023323322202301-2001212311212200-3120203333313201-3313131212113023-3221133103211103-2132013313322301-3212330332120302): complete subsection reference.

- [inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-3013210110031313-3121232321321212-0321000003332322-3113102131130032-3322312013210103-0132010210333031-3212130330311012-0000023311020130): complete subsection reference.

- [ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-3112010312320322-3011123232321023-2203213132012001-0013212231233201-3020330201212031-3123231121312022-2231213231022020-2011120212032020): complete subsection reference.

- [label_matcher](data-sources--network_policy_view--reference--group-001.md#canonical-0002213203230022-1013101101002020-1331132233231101-3212230020000002-3000002303100102-3332033230023323-1100022202203223-0313313330203303): complete subsection reference.

- [label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-0003203132122230-2030201311201020-1232213332133022-2222321300022321-0003311132113333-1101221030020232-3221103332020201-1222320232020310): complete subsection reference.

- [metadata](data-sources--network_policy_view--reference--group-001.md#canonical-3102200022002323-0102113103000221-0130320232122200-2101133310130230-0133110112333021-2002230311020332-3000020101032101-0111330330323311): complete subsection reference.

- [outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-0212231310120212-2131122310002101-0331331212312032-1031300013231001-3032312123121210-2120033110322300-0323330022133333-3012000201301322): complete subsection reference.

- [prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-0032110322313321-1323000222322302-1132111322200130-1102212012210020-0003221121223000-0100203000323223-3113132222323013-1331002322311022): complete subsection reference.

- [protocol_port_range](data-sources--network_policy_view--reference--group-001.md#canonical-2103010300101312-2113023010012333-2012111120001202-2032101331122223-1333121202302020-2201110223000222-3010303021001333-2021320103130002): complete subsection reference.

<a id="canonical-1213113022101130-3322020301321030-2300301120031131-0102211102102111-0301122103133233-0220302001102211-0103102322310030-3321033011122231"></a>

## Next pages — ingress_rules / 123301032220 / 5

- [ingress_rules.adv_action](data-sources--network_policy_view--reference--group-001.md#canonical-3030110321201232-1333103111021021-1013122302111032-2100203310301332-1032123102300313-1120303310012033-1002130200122222-1132100313233010)
- [ingress_rules.all_tcp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-0333020103300021-1330000322120200-3111103330222033-0033220132100300-1013303111230120-0111031211321123-1112231000030320-3130011010032011)
- [ingress_rules.all_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-1021232322221302-2311331101213203-3203323120012311-0013311101301120-0203030302130332-1221023203000220-3031023101131120-3303131302032320)
- [ingress_rules.all_udp_traffic](data-sources--network_policy_view--reference--group-001.md#canonical-3130000110212210-1331231110311102-0310121222230033-0020010303000203-2121111001001220-1333212331101210-3112110332010022-0100321310332010)
- [ingress_rules.any](data-sources--network_policy_view--reference--group-001.md#canonical-3333023301201321-0311011101001132-2200200320012032-3333120203310203-2201203330123033-0331333120123030-0020021202012011-3030203001133020)
- [ingress_rules.applications](data-sources--network_policy_view--reference--group-001.md#canonical-3332100121000323-0023323322202301-2001212311212200-3120203333313201-3313131212113023-3221133103211103-2132013313322301-3212330332120302)
- [ingress_rules.inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-3013210110031313-3121232321321212-0321000003332322-3113102131130032-3322312013210103-0132010210333031-3212130330311012-0000023311020130)
- [ingress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-3112010312320322-3011123232321023-2203213132012001-0013212231233201-3020330201212031-3123231121312022-2231213231022020-2011120212032020)
- [ingress_rules.label_matcher](data-sources--network_policy_view--reference--group-001.md#canonical-0002213203230022-1013101101002020-1331132233231101-3212230020000002-3000002303100102-3332033230023323-1100022202203223-0313313330203303)
- [ingress_rules.label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-0003203132122230-2030201311201020-1232213332133022-2222321300022321-0003311132113333-1101221030020232-3221103332020201-1222320232020310)
- [ingress_rules.metadata](data-sources--network_policy_view--reference--group-001.md#canonical-3102200022002323-0102113103000221-0130320232122200-2101133310130230-0133110112333021-2002230311020332-3000020101032101-0111330330323311)
- [ingress_rules.outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-0212231310120212-2131122310002101-0331331212312032-1031300013231001-3032312123121210-2120033110322300-0323330022133333-3012000201301322)
- [ingress_rules.prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-0032110322313321-1323000222322302-1132111322200130-1102212012210020-0003221121223000-0100203000323223-3113132222323013-1331002322311022)
- [ingress_rules.protocol_port_range](data-sources--network_policy_view--reference--group-001.md#canonical-2103010300101312-2113023010012333-2012111120001202-2032101331122223-1333121202302020-2201110223000222-3010303021001333-2021320103130002)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-3030110321201232-1333103111021021-1013122302111032-2100203310301332-1032123102300313-1120303310012033-1002130200122222-1132100313233010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203233202312031-0302311002201001-0230212001010012-3021322013323211-1303302202213101-1033033122323021-2112323012323313-2110123220203231"></a>

## ingress_rules.adv_action — adv_action / 311101130121 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.adv_action

<a id="canonical-2310022020022201-1311110213123132-2121013313302211-0110032110011313-0030323021221203-3101022201233002-2310131022020220-3312301001121200"></a>

Type: `"single"`. Computed.

Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and
PBRRuleAction.

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

<a id="canonical-3320303313302303-3013130032002021-3330301310302132-0313103211030201-1013013330210231-2013113330321210-3120301021312003-1031011310101312"></a>

## Direct properties — adv_action / 311101130121 / 3

<a id="canonical-0111100203321331-0201221011331102-3032102220303113-2200222311132331-3302000303112103-0330333330120313-3203310233020310-3233130301000220"></a>

<a id="canonical-1221110332023311-0331123310111301-2001123222020202-2301101110321012-1213323003103020-1211032132131033-3313100121030333-2132132033022320"></a>

## action property — adv_action / 311101130121 / 4

Type: `"string"`. Computed.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Upstream description:

Choice to choose logging or no logging This works together with option selected via
NetworkPolicyRuleAction or any other action specified x-example: (No Selection in
NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction
&#8203;+ AdvancedAction as LOG) = Log and Allow/Deny, (ALLOW/DENY in NetworkPolicyRuleAction + NOLOG in
AdvancedAction) = Allow/Deny with no log

Don't sample the traffic hitting the rule Sample the traffic hitting the rule.

Receipt-pinned upstream constraints:

```json
{
  "default": "NOLOG",
  "enum": [
    "NOLOG",
    "LOG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2332321211213113-0001113200230100-1113213030311101-0031021030330222-2211313230230033-2333221001303120-0333310121013011-2103110311120321"></a>

## Next pages — adv_action / 311101130121 / 5

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-0333020103300021-1330000322120200-3111103330222033-0033220132100300-1013303111230120-0111031211321123-1112231000030320-3130011010032011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200220012022032-1132010221033003-3210233011123020-3330023100131222-3101130123203122-0130330013330113-2310213101010113-1332010332032123"></a>

## ingress_rules.all_tcp_traffic — all_tcp_traffic / 031032013133 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.all_tcp_traffic

<a id="canonical-1221003022212223-3230312303323013-3032100022030321-2103310330112322-2013003013331223-3121201322223000-1001030300212032-2313202021120202"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all tcp traffic.

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

<a id="canonical-2121222222332203-0133333102311303-3112303202013323-0013032312202032-2220302130331231-2113002002310320-3200311022310121-3302020211222332"></a>

## Direct properties — all_tcp_traffic / 031032013133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301033210311222-3221200321112221-3313223113033113-1010010201222202-2120323002232132-1131302230112112-3130100001330323-3222010013210012"></a>

## Next pages — all_tcp_traffic / 031032013133 / 4

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-1021232322221302-2311331101213203-3203323120012311-0013311101301120-0203030302130332-1221023203000220-3031023101131120-3303131302032320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303323122213322-3002102201021213-2130302102211010-0020012000113332-1000310010223032-0122100130213201-2211210003210031-3330221233132112"></a>

## ingress_rules.all_traffic — all_traffic / 320013131221 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.all_traffic

<a id="canonical-1132222022101100-3232002100020221-2023022123233002-1011132020032011-3101322320322232-2101001333120120-2121030320000032-3022331221230302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all traffic.

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

<a id="canonical-2210013223322222-1132030223333300-3001321031222003-1310230223001311-2113011302320100-2323222203331203-0113303210312132-3312302012133212"></a>

## Direct properties — all_traffic / 320013131221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2312200010023012-2022211320023200-1321202232110320-0223322330130031-3223011102011130-0200111031133000-2103310231312000-3030102322101311"></a>

## Next pages — all_traffic / 320013131221 / 4

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-3130000110212210-1331231110311102-0310121222230033-0020010303000203-2121111001001220-1333212331101210-3112110332010022-0100321310332010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220123211311012-2311231301030220-2203001012120013-0010110323111020-1211111302002110-0011313030031032-3202113332033033-0101322020213022"></a>

## ingress_rules.all_udp_traffic — all_udp_traffic / 202323302303 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.all_udp_traffic

<a id="canonical-0030102320000212-3212200013330112-2200021202030003-0131131112230003-0022310100232302-1211202123133221-1033113132230021-0033311010333200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all udp traffic.

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

<a id="canonical-3213123211032130-1110331313233132-0001311301302320-1310103212121031-3232011113120030-2010232112013122-1320021110210312-2201321232023120"></a>

## Direct properties — all_udp_traffic / 202323302303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003301132000012-1121020222122030-3030212112120030-1030001331133200-1013123221213021-0112021000332233-3311021001012230-1110311333132322"></a>

## Next pages — all_udp_traffic / 202323302303 / 4

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-3333023301201321-0311011101001132-2200200320012032-3333120203310203-2201203330123033-0331333120123030-0020021202012011-3030203001133020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201102002331220-1312320212021330-1321200322231131-0322311112122333-1120020023121122-3221121133011131-0001202020030323-1320312330210001"></a>

## ingress_rules.any — any / 323103220330 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.any

<a id="canonical-1122032021000011-2120033132021133-2302203132021113-3013311200230130-1220013130000302-3030012331023230-2301120333213222-2332031311102112"></a>

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

<a id="canonical-0000012230303021-0010131113112003-0303000300012200-1312020112003111-2010231100112112-1301332232211232-2023312332002313-3220220221012233"></a>

## Direct properties — any / 323103220330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002102010301102-2022031031220123-1302030012130130-0330110113310332-3123233222222002-2021013332233310-1333330333022300-2210031313011202"></a>

## Next pages — any / 323103220330 / 4

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-3332100121000323-0023323322202301-2001212311212200-3120203333313201-3313131212113023-3221133103211103-2132013313322301-3212330332120302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302231102302010-3102331211232212-3022130232130302-0120332102130320-0212202132121201-0231133130032013-0003112013001032-2232021003033022"></a>

## ingress_rules.applications — applications / 200023201201 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.applications

<a id="canonical-3113030113133301-3222100112201023-3212323131031113-2110321111333303-0131001322220023-3122133331213033-1003213320023310-3130002231222302"></a>

Type: `"single"`. Computed.

Configuration parameter for applications.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="canonical-3312131122313210-3130120023001232-1323020112202212-3123100312321203-2010311322102312-1311123122332101-2120201213021000-1233221203023330"></a>

## Direct properties — applications / 200023201201 / 3

<a id="canonical-2323331102232311-2210033313221213-2310021121121232-3122010203030103-0201120303013021-1303221202311231-2002111113323322-0022221110311231"></a>

<a id="canonical-1033311231320311-2313233221020303-1203323222210301-3100003312013121-2301003322101320-2100313032300112-3031123331221322-1033233332100221"></a>

## applications property — applications / 200023201201 / 4

Type: `["list", "string"]`. Computed.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="canonical-1202321200321332-0000200032322323-0200211331312333-1030130313111332-1311323320212332-3123123233223232-0010103312313103-3321111122003210"></a>

## Next pages — applications / 200023201201 / 5

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-3013210110031313-3121232321321212-0321000003332322-3113102131130032-3322312013210103-0132010210333031-3212130330311012-0000023311020130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330223303101321-1320323113323131-3130121223330111-3301012000023320-0232001330330311-1020230112233331-2030213021320300-3032322210313022"></a>

## ingress_rules.inside_endpoints — inside_endpoints / 120311232331 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.inside_endpoints

<a id="canonical-1311212122331323-1332203103121012-1113202333212033-1111000233131110-3101322301023332-1322033322121000-1203203331332000-2032102002332300"></a>

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

<a id="canonical-1300232303210203-3333023011013301-0102200322122131-2122103001110301-2133201020023122-1332320233023200-2001331120102012-2113131110300200"></a>

## Direct properties — inside_endpoints / 120311232331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302131122301313-0103223020030000-3021330012203201-1011122330300020-2301123212321213-1302011232210112-0030123320022330-3000032112202320"></a>

## Next pages — inside_endpoints / 120311232331 / 4

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-3112010312320322-3011123232321023-2203213132012001-0013212231233201-3020330201212031-3123231121312022-2231213231022020-2011120212032020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033213202331300-1020113200100330-2013231332303102-1321112200022321-3021323212310201-1231310122003120-1132022200020232-3113131212210221"></a>

## ingress_rules.ip_prefix_set — ip_prefix_set / 223220221332 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.ip_prefix_set

<a id="canonical-3331320012003021-2310020312333030-2301332201301303-2102212101212213-3220110121333112-2123122103022331-1302211230332332-0011311022200323"></a>

Type: `"single"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

<a id="canonical-3021121332320122-0031303202220320-1121301321030230-2030211012120101-2112221011230022-1232313030230321-2030323030120303-1230001132101111"></a>

## Direct properties — ip_prefix_set / 223220221332 / 3

- [ref](data-sources--network_policy_view--reference--group-001.md#canonical-2203303213220001-3230310310200310-0021000201202332-1122302222101003-0301021112020231-1010221310202101-3302202021331131-2313223002130223): complete subsection reference.

<a id="canonical-0133321021333212-0000213110211230-3330300120331102-3232323313033000-2010233230021103-3110210200300001-2111010112330330-2213221213100323"></a>

## Next pages — ip_prefix_set / 223220221332 / 4

- [ingress_rules.ip_prefix_set.ref](data-sources--network_policy_view--reference--group-001.md#canonical-2203303213220001-3230310310200310-0021000201202332-1122302222101003-0301021112020231-1010221310202101-3302202021331131-2313223002130223)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-2203303213220001-3230310310200310-0021000201202332-1122302222101003-0301021112020231-1010221310202101-3302202021331131-2313223002130223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312323120213001-2110233200001232-0320331311222130-3300130113020110-0100301123122023-2020122212223333-0201203230210302-2103203232310321"></a>

## ingress_rules.ip_prefix_set.ref — ref / 103123021001 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [ingress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-3112010312320322-3011123232321023-2203213132012001-0013212231233201-3020330201212031-3123231121312022-2231213231022020-2011120212032020)
- ingress_rules.ip_prefix_set.ref

<a id="canonical-0111302112300013-0003132032211121-2120220130113112-0112332312112133-2130211100000210-1032233322022233-1332313312311313-3003113011112002"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1210013302121233-3111011230320110-0301122103010330-1233003332210112-0122132032003313-0232030103001322-1211312331220021-1301322222030010"></a>

## Direct properties — ref / 103123021001 / 3

<a id="canonical-2233112311313022-1000120331333301-0212221023302002-0013000333020312-3122312302331311-2130101110300102-1030221003312131-2100020030220011"></a>

<a id="canonical-0312130211333220-1111210021011132-1023303202220010-0032322220312220-1331301330320320-3012133133222300-2300312200032113-1220302120212113"></a>

## kind property — ref / 103123021001 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-0322332331101021-3331020223220222-3233232020011200-2033300032323310-3020032133021312-3220033022302133-1332300321301302-0110332113202332"></a>

<a id="canonical-3003301211330003-1132313000011201-2123202312210131-2323310230330211-3003222101303201-2221122021213103-1201033332023002-2122311032202221"></a>

## name property — ref / 103123021001 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-3030100320200323-2010303122023013-0103212003223201-0131303110101301-2113301133210210-0102123221022020-0002210000021133-1300302110021112"></a>

<a id="canonical-3200012022311031-3212221120122322-0010101031013121-2112233123212121-0301220210130313-1012110022232211-0312110202001010-1300130020202013"></a>

## namespace property — ref / 103123021001 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3000333012002132-1023223311130101-1301223201100322-3100230013231211-2001330133310202-0000010202003003-2323121121300303-3010030212010032"></a>

<a id="canonical-3232213030113120-3011310303320303-1331111121132000-3211011221102330-0022133220231112-2213112123120130-1323233130031123-0200332313011022"></a>

## tenant property — ref / 103123021001 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0121233302302011-2023301121022231-0331000220121100-1033232212031000-0103322030223122-0010313222000032-2100121111210011-0232223202022022"></a>

<a id="canonical-3313111121221130-2132132311331331-3220033203030112-0032321313130021-2332303200220210-2320020210332202-2122002131003322-1310000332320232"></a>

## uid property — ref / 103123021001 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-3221113332020231-3322201122322231-0221133100201033-0202022002021020-3130220100320210-0031221000223033-1120221023331100-2211313200121001"></a>

## Next pages — ref / 103123021001 / 9

- [ingress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-3112010312320322-3011123232321023-2203213132012001-0013212231233201-3020330201212031-3123231121312022-2231213231022020-2011120212032020)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-0002213203230022-1013101101002020-1331132233231101-3212230020000002-3000002303100102-3332033230023323-1100022202203223-0313313330203303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123203000123012-2330002102210331-3101203002332331-2021313012201201-1012323321010013-2121012123133302-0320302211132000-3112030332330022"></a>

## ingress_rules.label_matcher — label_matcher / 013001122213 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.label_matcher

<a id="canonical-3102000212100033-3030132311311203-1012100130331301-0233333020202013-2200212201103220-2130332020123110-1321110233321223-3132201133303313"></a>

Type: `"single"`. Computed.

Label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Upstream description:

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

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

<a id="canonical-0203100202020011-2030002333113113-1230213333010122-3203100120113203-0100302301220012-0331021030331233-0023102233303012-2113102123200012"></a>

## Direct properties — label_matcher / 013001122213 / 3

<a id="canonical-1011213233010311-1233033233212130-3110321220032333-3222313101231320-2331123130320312-2320310310013323-3130232200230023-1022231031030030"></a>

<a id="canonical-3032012112320232-3102001102211011-3231220011232233-1122311220231132-0002002101331311-0002230322013031-1221013230111301-0130223203313220"></a>

## keys property — label_matcher / 013001122213 / 4

Type: `["list", "string"]`. Computed.

The list of label key names that have to match.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0232101103330210-3330103213313200-3210030023320311-0012202010221312-0332030123312300-2310133221201211-2030022220331322-3310202232321120"></a>

## Next pages — label_matcher / 013001122213 / 5

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-0003203132122230-2030201311201020-1232213332133022-2222321300022321-0003311132113333-1101221030020232-3221103332020201-1222320232020310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202130022113113-2013023012121213-3213321231013000-2113233223220330-3000323231032023-2330130032221022-3101111103213113-2211121111101310"></a>

## ingress_rules.label_selector — label_selector / 202312311302 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.label_selector

<a id="canonical-2110021000201020-2031310000201302-1203210030010100-1012021212310010-0300010001310133-1300200120331001-0200033121331103-2130131303330302"></a>

Type: `"single"`. Computed.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-2111010032230133-1013311130211121-1000203332023310-3000310122020223-2232013220003023-2122003113011101-2130132010123223-2030122110330002"></a>

## Direct properties — label_selector / 202312311302 / 3

<a id="canonical-0302032121101130-3023012223231033-3301110111010131-1202032002322220-1222131011120323-3010213202311023-3310111021230101-2022112020020123"></a>

<a id="canonical-1230101022300000-3121233031013231-0030233113313331-1130111313021222-0220020120202211-3003031202320010-2100003030232332-2200101010332130"></a>

## expressions property — label_selector / 202312311302 / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-3313000210032212-3133323223323210-2120023232123011-1100230001033311-1123030313131133-2221300131002310-1221221033210232-1112231301031011"></a>

## Next pages — label_selector / 202312311302 / 5

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-3102200022002323-0102113103000221-0130320232122200-2101133310130230-0133110112333021-2002230311020332-3000020101032101-0111330330323311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232320231131313-3200330031200113-3122033111220013-2000211321220132-3103200213320030-2331133112211203-2201102031203001-2220130001001002"></a>

## ingress_rules.metadata — metadata / 112022112222 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.metadata

<a id="canonical-2133131322333313-1103002021110002-1032203233123021-2100312221301322-0022200213311002-3300322002120231-3021313001230213-3312032112233331"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-0131110012112020-3231011131210231-3122221032100113-2113011203320000-3323302033213300-0121120120202123-2312200203101301-2003123313021320"></a>

## Direct properties — metadata / 112022112222 / 3

<a id="canonical-1332201332333032-0120110032313320-1323321020333332-0131031130212312-0112232023121212-3110031031020020-1012310320031321-1133122312313023"></a>

<a id="canonical-3310301213032200-2202303121213320-1311213101033101-2000330103213332-1313221211230333-2012332302003233-3022312220113300-2020212132112101"></a>

## description_spec property — metadata / 112022112222 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2031111132010112-0320112130302302-2201130211232021-1302100031331203-3330203300000202-3031033122022013-2312230013312333-1303302332133133"></a>

<a id="canonical-3330001122011031-2123302331311222-0000301033110121-1213301332013220-0003202233202121-1110130211012330-3102002300303203-1120312202222003"></a>

## name property — metadata / 112022112222 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0010331303022333-2130320212300000-1023111022310201-3021320003021132-0300223131322202-2321001111330330-0111301011301320-1011021223032033"></a>

## Next pages — metadata / 112022112222 / 6

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-0212231310120212-2131122310002101-0331331212312032-1031300013231001-3032312123121210-2120033110322300-0323330022133333-3012000201301322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132333210233111-2012323011002320-3311113310030232-0103023223201023-1112102111200233-3131311001101230-0202110132223130-0003122213300312"></a>

## ingress_rules.outside_endpoints — outside_endpoints / 020032320022 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.outside_endpoints

<a id="canonical-3011033200020130-1002102113212113-2010010331000320-2100203200232232-1211201100121112-3100333030313221-3023010101323222-1012232032310030"></a>

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

<a id="canonical-1123001001222331-0030213020112113-1033020210213120-2002321312312030-0002203113300020-0131322200201121-3031301121022133-3320112031100202"></a>

## Direct properties — outside_endpoints / 020032320022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312022030233131-3213221302112301-2032200232323010-2023110222201031-3022330212202113-3000232201100201-1303033213310121-0310002333320132"></a>

## Next pages — outside_endpoints / 020032320022 / 4

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-0032110322313321-1323000222322302-1132111322200130-1102212012210020-0003221121223000-0100203000323223-3113132222323013-1331002322311022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313113020333023-2320002231332003-3322031302232332-2331010003111200-1212001021012200-2322322203122221-1033320131220221-3200031230312111"></a>

## ingress_rules.prefix_list — prefix_list / 320032002303 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.prefix_list

<a id="canonical-1323303230320033-1311130331030120-0212202333131210-2000131131202022-2220100331211201-1213033232013132-0201120331122310-2213010323022033"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-0222212003031002-0301011101010211-3333220122200320-2122303110031323-0101333122121202-1110200102121111-1003212032113313-0312202001210012"></a>

## Direct properties — prefix_list / 320032002303 / 3

<a id="canonical-2011313131121021-1322010123002022-1231203011303133-2301322231030220-2221010331311223-2301230210230012-1330220311300113-0131302011020111"></a>

<a id="canonical-1322122123312320-2301110333023230-3133332010130310-0200130320332312-3211033033231232-1213313121023223-1123212202011111-3022303321201132"></a>

## prefixes property — prefix_list / 320032002303 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0220001102133220-1201010003113101-3322130302102122-2123130123122310-3202202301103213-1132131112312213-1332022033333210-3002122112031233"></a>

## Next pages — prefix_list / 320032002303 / 5

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)

<a id="canonical-2103010300101312-2113023010012333-2012111120001202-2032101331122223-1333121202302020-2201110223000222-3010303021001333-2021320103130002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323233322330302-2031301102313222-1013000220303022-0110231131303302-2320213222010210-3300103313332002-0200323210031030-1200323313210211"></a>

## ingress_rules.protocol_port_range — protocol_port_range / 120112220333 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.protocol_port_range

<a id="canonical-2232122101222311-2322222102310230-2210101230021302-3320211001322123-0200130200303203-1120310310000210-2020311100232230-2212232203230023"></a>

Type: `"single"`. Computed.

Protocol and Port. Protocol and Port ranges.

Upstream description:

Protocol and Port ranges.

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

<a id="canonical-3313112000113311-0010300221131032-2210110011132201-3221332220300232-2010113212230031-0131023323003333-1032031301013033-0332233113303222"></a>

## Direct properties — protocol_port_range / 120112220333 / 3

<a id="canonical-0121211002312013-1023113102332022-2101232233002101-0022001332302020-0120000112133132-2332002201210303-1301320020032222-2331300222100122"></a>

<a id="canonical-0321011310113303-3233333300130111-0232201232013302-0131201233021103-2133022120121001-3111111311102022-0232100332320121-3031000030131023"></a>

## port_ranges property — protocol_port_range / 120112220333 / 4

Type: `["list", "string"]`. Computed.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-2230221211233233-1330112133203220-3233020313130320-1123122132303132-3321111300102021-0110101210023221-2011102001330233-2122320213010001"></a>

<a id="canonical-0302102131302311-1312022100133023-1322131013232103-2220013102023332-3030332211101203-2232111031213102-2120131321003313-2123030223212102"></a>

## protocol property — protocol_port_range / 120112220333 / 5

Type: `"string"`. Computed.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-3302330221312023-0003013103133231-0212311103312302-3023320320022113-0130230131233030-0012111113232031-0020012303010102-1100021302223020"></a>

## Next pages — protocol_port_range / 120112220333 / 6

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
