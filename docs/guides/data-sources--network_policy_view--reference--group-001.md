---
page_title: "xcsh_network_policy_view reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_view reference."
---

# xcsh_network_policy_view reference

<a id="canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- Property reference

<a id="canonical-1123103312213333-0200332301133101-2030212133023131-3310203202012223-1131313032032222-1023030211122130-2001020331102030-3301201011312031"></a>

### Direct properties for `xcsh_network_policy_view`

<a id="canonical-0112023300132001-0131032212213212-2203123001210332-3232102010313113-1302030022330032-1130233001201111-2320300310233032-3002113032112312"></a>

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

<a id="canonical-1012020213021330-3103030022102122-0200222201120331-1010321333033030-0002003223130132-1013012023311201-3111000020000131-0011233023020223"></a>

<a id="canonical-2130220000101302-3221323211121302-0302122011221232-1333003103201312-1210202221332311-0200100100021302-1010000312000000-2012332201320100"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the NetworkPolicyView.

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

- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022): complete subsection reference.

- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120): complete subsection reference.

<a id="canonical-1212231112313013-0030020201312101-2232200333002003-3021222231101031-1222031002201233-0201230231121301-0102120200003032-1000230233110322"></a>

<a id="canonical-1211220201120202-2032232011103000-3130022103221213-1220310311101003-2122200011233011-1133033231003133-0133222303203222-1231232003331201"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230): complete subsection reference.

<a id="canonical-3301230033121220-3133023010111112-1313113310031131-0111023010223223-2020130013323223-3003021233232111-1123001031300030-0113111011230103"></a>

<a id="canonical-3132200232003200-0121103303033113-2010112031112132-2101223002203211-1210132100022121-0011022111011131-0013131233200303-2220132201201022"></a>

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

<a id="canonical-2031220122032332-1013330130211303-1312003231021332-0100013003020011-1021331110221202-2112321133331022-3210122133121232-0222111311010022"></a>

<a id="canonical-0020003012002213-3001102210201333-3313121200231133-0301203130011311-2011131112110210-3332203110301022-0322311102002020-2210132031321330"></a>

#### `name` property

Type: `"string"`. Required.

Name of the NetworkPolicyView.

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

<a id="canonical-1003123010310220-0211121011003203-3021002303232312-1020201211121320-3223202113220033-3011332123002130-2130003322212001-3332322213020010"></a>

<a id="canonical-1313320320203123-1210203211000113-1102200223330202-1233113313211110-2310310130122132-3002111200033121-1033303003001112-0123000110311200"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace where the NetworkPolicyView exists.

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

<a id="canonical-3311312013002130-2002101113110121-0221003011002001-2222002322333002-1020030231021233-3020330323310303-1012023203121310-1123302033102303"></a>

### All schema paths for `xcsh_network_policy_view`

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
| `id` | [ID](data-sources--network_policy_view--reference--group-001.md#canonical-1212231112313013-0030020201312101-2232200333002003-3021222231101031-1222031002201233-0201230231121301-0102120200003032-1000230233110322) |
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

<a id="canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules` properties

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-1303131230000110-0200103033323310-3133201133321223-1033321021123232-0213102013023121-1002032302003002-1233311311311112-1132320331233101"></a>

### Direct properties for `egress_rules`

<a id="canonical-0103320330222110-2300210011021101-0301021221000332-0120112110222003-1123331023130300-3112302011121010-3301112311222331-1212131321013112"></a>

#### `egress_rules.action` property

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

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

<a id="canonical-3033020000112100-2221203131311100-0310103311321113-0300133013210303-0210001102310321-1332131003313012-2200310312001013-0220033302332332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.adv_action` properties

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

<a id="canonical-0330011332211000-1322110020021121-3021103311220100-0011033113123022-1022111322112022-1302030121313003-3300312302113102-0032210033112212"></a>

### Direct properties for `egress_rules.adv_action`

<a id="canonical-1010222112000120-2130230103022001-0302003333320212-2223021012210000-0213223102203001-3011011211003010-0101101323310210-2211102100321303"></a>

#### `egress_rules.adv_action.action` property

Type: `"string"`. Computed.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Additional upstream details:

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

<a id="canonical-3131220310002121-0333202203303100-0220020120133330-3201130223321132-3223132233123212-0321231210313322-2010320122231322-0311232310301013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.all_tcp_traffic` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.all_tcp_traffic

<a id="canonical-1313110331133132-1333211021110032-2221111030023323-2200030022313122-3213221010120103-3113122323233223-0130112321200023-2002303120221320"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all tcp traffic.

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

<a id="canonical-3032300120113323-3310112203301102-3110003000102022-1123313020113303-0321223321230103-2003300330302011-1201213330103311-1201001132112301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.all_traffic` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.all_traffic

<a id="canonical-0112130123320212-3202203011303200-2102022121022221-1222121221022320-0323332013320202-0220221221210100-0123032111013321-3011130101101311"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all traffic.

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

<a id="canonical-0200311213333321-1330232333122100-3230101223121111-3222003211313000-1321321112022031-0010003311313210-1312021230100313-0302022003312333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.all_udp_traffic` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.all_udp_traffic

<a id="canonical-0322213201221013-2213313120310323-0203011232310232-0102021113003312-0201020013323233-2112213033033211-0200203011102003-0102033302112033"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all udp traffic.

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

<a id="canonical-1332112231101121-2133231101212101-0022022313312300-2220003013000111-2131001023111033-0031223320010103-1200310110021033-0021030221330323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.any` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.any

<a id="canonical-2202030120110030-0203300212211122-3212322200102200-2233223123112331-1202121302100333-2310010303102102-2223301120103300-2210130031000101"></a>

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

<a id="canonical-0230320310311312-2233000100332133-2332231001331212-2131233221220022-0301113223132121-3201121201103111-1303300131013332-1110122212323112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.applications` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.applications

<a id="canonical-1223201121101002-0103322132222021-2101222100231011-3010000230310310-0200031330330301-2320203100211102-2122030122211310-3323010203131103"></a>

Type: `"single"`. Computed.

Configuration parameter for applications.

Additional upstream details:

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

<a id="canonical-3220103011231323-1212322201011020-0231120303210201-0331221032000222-0330221221030012-3231211111231230-1122131110103332-1112111011132210"></a>

### Direct properties for `egress_rules.applications`

<a id="canonical-2100233213332031-1120313112230313-3203220310121023-2023031230202331-0000023220001010-3130220212021020-2221120131100102-0313033231103232"></a>

#### `egress_rules.applications.applications` property

Type: `["list", "string"]`. Computed.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

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

<a id="canonical-3102312011022130-0311203322112022-0111001223000212-0100220133112133-0201030303303222-3313120232112133-3113102101130133-2212113132333302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.inside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.inside_endpoints

<a id="canonical-0310302311323130-2311120110022000-0133102233323202-0010023330211100-0103022121322221-2100302110333313-1113032130110303-0220103231020011"></a>

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

<a id="canonical-2300101113211222-0321120310000113-0300010211332333-0101331103203200-1301230331121333-0331330012300000-3010030201232013-0321230110202323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.ip_prefix_set` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.ip_prefix_set

<a id="canonical-1331001321030200-0030121330120112-2321321022021031-1031011022031300-3302303122333213-1102022332103233-3211321303121103-0023223111311222"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3001211203110232-3332211122212120-2123230213132303-1013030003001320-1203131112022322-3033122113113201-0220122210211122-3102332121120310"></a>

### Direct properties for `egress_rules.ip_prefix_set`

- [ref](data-sources--network_policy_view--reference--group-001.md#canonical-0122032033213213-2330231233312003-0220213031123200-3323121112320310-1102312230201223-0301000311321031-3203111021233023-1223120313201231): complete subsection reference.

<a id="canonical-0122032033213213-2330231233312003-0220213031123200-3323121112320310-1102312230201223-0301000311321031-3203111021233023-1223120313201231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.ip_prefix_set.ref` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- [egress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-2300101113211222-0321120310000113-0300010211332333-0101331103203200-1301230331121333-0331330012300000-3010030201232013-0321230110202323)
- egress_rules.ip_prefix_set.ref

<a id="canonical-3133120030310023-2332021013303131-2031310003001213-1112331211020210-3031121211323301-0120022110010322-1023131133002223-0101111213121021"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0121032220033301-2020123211312313-3300200200020031-2011322001213331-0021003023120023-1320033002120100-2302213111130032-0013212112001323"></a>

### Direct properties for `egress_rules.ip_prefix_set.ref`

<a id="canonical-2110113203222003-3222010330201233-1012021311001200-3121312130200332-2003103202333333-3231230222132013-3301213233302003-1233102213321102"></a>

#### `egress_rules.ip_prefix_set.ref.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-3211001301121203-3023312033300200-2113100202301211-3321211003033103-1012113212330210-0000132022322130-2010312031121003-2322203012133130"></a>

<a id="canonical-1032200202012003-0103311022320330-2212201312200123-3201201321333332-0122223032011100-1233233102011032-0130213032031211-3203133202331002"></a>

#### `egress_rules.ip_prefix_set.ref.name` property

Type: `"string"`. Computed.

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

<a id="canonical-0030230233113221-0303322120030000-1110223300121020-0232120123222110-0203220202002000-3201230333320332-1313120020030100-3222023233321112"></a>

<a id="canonical-3031110131132211-3123022030231333-3301201332023301-3112233020300300-1320223031120230-2231103130233031-0331311202001123-2203133120223032"></a>

#### `egress_rules.ip_prefix_set.ref.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-1123132303103132-0211210102133103-3003021020100320-3000102331201110-0103201033230203-1112102030123103-3020323322321222-0000223302333330"></a>

<a id="canonical-2011131123111333-1301130332023223-0032110130011333-3121132230010023-3203330023131302-3223320300221112-2103001020030321-0122111010320320"></a>

#### `egress_rules.ip_prefix_set.ref.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-2023132230032131-1033301213220302-3013003100030113-3222232120133010-1233123121213333-1032001121023300-1011131113033101-2322210131030200"></a>

<a id="canonical-0320203211022013-0321212300123103-0010322213133213-3300130101203213-3212310122030210-0232013232011312-0212121023233210-2000111102203130"></a>

#### `egress_rules.ip_prefix_set.ref.uid` property

Type: `"string"`. Computed.

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

<a id="canonical-0323131022301030-3200102201333212-1103330201131221-3130311110123303-2331232101203231-1333110232211320-0200213010123130-3203001212101130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.label_matcher` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.label_matcher

<a id="canonical-1123103330112233-3123012000323111-3122320002302323-1103003012103303-0303313020312010-3003133022300302-1001020113020020-1101101000320100"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2013320203230132-2303320011310333-2201302323320012-2102200132103303-3120332312122020-1100201223312333-3020330230302331-0302203220312032"></a>

### Direct properties for `egress_rules.label_matcher`

<a id="canonical-3233330330233301-0311132323031302-0011231130331232-3131002011112222-2211231000132031-3112211103320200-0113122031022102-2031333210332111"></a>

#### `egress_rules.label_matcher.keys` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2302031122222020-1121201013111113-3203002211310103-0223203322231131-2313202131023232-2000330211100220-1003130013232022-0221130003232320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.label_selector` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.label_selector

<a id="canonical-2112113302032332-2332013313133332-3203032320332002-1133111212101203-2320120030103300-1201113003012203-0113211013302132-0200321210323230"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2303111102013231-2211113331220120-0212000300022030-1011303111210103-1032032021001230-3121203112323133-3223033000302212-0300001033330223"></a>

### Direct properties for `egress_rules.label_selector`

<a id="canonical-0332001021330121-0303201320110011-2312313312220131-3331023310222222-2032320210301122-3100100230002110-0033001213333230-1323110002211031"></a>

#### `egress_rules.label_selector.expressions` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1322322222202103-1231102103200131-3210202132100221-1323301120222012-3031213212221322-1132221230021120-2030230103322321-1333310001130312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.metadata` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.metadata

<a id="canonical-0033200123132322-1012110012010101-2210211211320001-1202310100232121-2323103201323100-0232302212121322-0323102012232113-2032323311022132"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2131103022213301-2003121003223002-3033020332212102-2103312100101033-0123211031330113-1223313202000133-2232023220032021-3220231123111220"></a>

### Direct properties for `egress_rules.metadata`

<a id="canonical-0202003022031032-0131330020303110-3313203312102333-0103123233021223-1321030022003222-2020231131312202-2320211020103320-2111033001331111"></a>

#### `egress_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2322121111031231-3322111332012112-2331210320303022-1210132311033320-3330102313201120-3030100310123013-0131320221023112-1131021113222320"></a>

<a id="canonical-3311313221213103-2011102310311002-2123003020032003-2203011230132322-0232200220002310-2100032031232010-1210011331200313-1222331223112332"></a>

#### `egress_rules.metadata.name` property

Type: `"string"`. Computed.

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

<a id="canonical-1033003313030113-1203303121020303-0130233220121120-0201031121310003-2302130002023111-2220300030320300-1232322121310132-2320303313331113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.outside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.outside_endpoints

<a id="canonical-2122232203000202-2032130131301110-3002003111321011-0221032031323332-1022323130001301-2011110022122133-0032032201312230-0320103321222203"></a>

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

<a id="canonical-1300023113210112-1312310002030100-0013030010321211-3000022210032002-2020300031132211-2300112121321330-1331020111323332-0122221331233322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.prefix_list` properties

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

<a id="canonical-2231123031112131-3221222220130222-0102022332123000-2111102002120312-3023012020221133-2012300223103212-1132133123210202-3101020313223110"></a>

### Direct properties for `egress_rules.prefix_list`

<a id="canonical-2230112310221321-0330301101013132-1331100010310221-3000211232012213-1222200313011200-1322121211023211-2112133322032112-2321032031312312"></a>

#### `egress_rules.prefix_list.prefixes` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1222022232132112-2220111000131211-1002311200331321-3232230300023030-2203001102301031-3120302213221033-0101011233011031-3012110023331302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.protocol_port_range` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [egress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-1223310232310203-1312220302320002-1002233232003003-0333313301113132-0001212022122011-1012122221122002-1100032130331230-3033222231013022)
- egress_rules.protocol_port_range

<a id="canonical-0112212333132020-1100131000320202-1310311211232022-0012230030320113-0330310201201312-0321221322323231-0302232213311333-3002001232300233"></a>

Type: `"single"`. Computed.

Protocol and Port. Protocol and Port ranges.

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

<a id="canonical-0302110023220112-0232230310103013-1033000301130012-3220330233320000-2022312201200332-3223230332000313-3031010311310012-1020101013320020"></a>

### Direct properties for `egress_rules.protocol_port_range`

<a id="canonical-0231111230202323-0323322230131020-2222212210331301-2130313200210223-0233011332132201-0300223221321230-2023000103312123-2131223201223233"></a>

#### `egress_rules.protocol_port_range.port_ranges` property

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

<a id="canonical-3323221123023002-0211111121110312-1231102222321202-2312230311203100-0031301323321031-2113102023333310-2312210201001021-3303033011333201"></a>

#### `egress_rules.protocol_port_range.protocol` property

Type: `"string"`. Computed.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

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
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint` properties

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

<a id="canonical-2113233021202000-0123031213001103-0103012233301200-0032113211132223-1130030001222211-1100100211320002-0201031303110230-0102233303211032"></a>

### Direct properties for `endpoint`

- [any](data-sources--network_policy_view--reference--group-001.md#canonical-0131333302333112-2310121311021213-2211101133233031-2220103131121112-0310103131133202-3330221100323231-0113133000021230-0132230333033333): complete subsection reference.

- [inside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-2202321203030200-3022202322001032-3123200130203231-3211312133323331-2231003332000102-0323103023332220-0032132213310313-0302132320021122): complete subsection reference.

- [label_selector](data-sources--network_policy_view--reference--group-001.md#canonical-1030323021230021-0313132002133212-2033222012202013-2021133100011213-1210303022033023-1221030213232212-1303110011330202-0223303322201010): complete subsection reference.

- [outside_endpoints](data-sources--network_policy_view--reference--group-001.md#canonical-3220110110203232-2020332111132301-2002312333231111-2330322032112311-1122122222212222-0211201023101222-0101312120133001-0130221313012322): complete subsection reference.

- [prefix_list](data-sources--network_policy_view--reference--group-001.md#canonical-1302102000330332-3102200010120011-1333131331310032-1003202300332230-1123221331302211-2211011020023102-1020300023010231-3211120010200310): complete subsection reference.

<a id="canonical-0131333302333112-2310121311021213-2211101133233031-2220103131121112-0310103131133202-3330221100323231-0113133000021230-0132230333033333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint.any` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120)
- endpoint.any

<a id="canonical-0003303203102212-1313100020211301-3013200231103132-0120312210030013-0030333300032033-1000232311130223-2112323313103111-0121203103133032"></a>

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

<a id="canonical-2202321203030200-3022202322001032-3123200130203231-3211312133323331-2231003332000102-0323103023332220-0032132213310313-0302132320021122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint.inside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120)
- endpoint.inside_endpoints

<a id="canonical-2001110021201323-1121211011322122-0103132012111333-1210231332121103-2231033120133001-1321233023332011-3131032013311232-1303020011313322"></a>

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

<a id="canonical-1030323021230021-0313132002133212-2033222012202013-2021133100011213-1210303022033023-1221030213232212-1303110011330202-0223303322201010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint.label_selector` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120)
- endpoint.label_selector

<a id="canonical-3200303103301110-3302231110332302-3331000110012202-2320313233100320-3232303200020302-1232301322223112-1111313013202122-0003313221102310"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0020210110321033-0132212021110330-0031132323102020-3111302100300323-2211302300003010-2211130101223302-0301331123000010-2330102231010313"></a>

### Direct properties for `endpoint.label_selector`

<a id="canonical-1120330110122121-1330131120010001-0332200120203030-1230013301212210-2232010011323031-2302131230011133-3212231112211311-2211312100011202"></a>

#### `endpoint.label_selector.expressions` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3220110110203232-2020332111132301-2002312333231111-2330322032112311-1122122222212222-0211201023101222-0101312120133001-0130221313012322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint.outside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [endpoint](data-sources--network_policy_view--reference--group-001.md#canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120)
- endpoint.outside_endpoints

<a id="canonical-2001121112001120-2310221020010032-3113233331332302-1333323002232111-1112220102002322-3020310312123201-2230122121322312-1133110011230322"></a>

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

<a id="canonical-1302102000330332-3102200010120011-1333131331310032-1003202300332230-1123221331302211-2211011020023102-1020300023010231-3211120010200310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint.prefix_list` properties

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

<a id="canonical-3121231022312131-1111323012003001-0123023332002101-0320113211320322-1233033103002012-0210200333112020-3333112300302311-2310000130303221"></a>

### Direct properties for `endpoint.prefix_list`

<a id="canonical-0232111023211330-2030012022301212-2010303110311220-1322221023102310-0231112330201101-0102033102301021-2020321220111221-0333121121021012"></a>

#### `endpoint.prefix_list.prefixes` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules` properties

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-3210030312102331-0000312130100002-0332022113010313-1030212010021300-2010200201131233-2333303123232101-0013212321202001-2011113312303122"></a>

### Direct properties for `ingress_rules`

<a id="canonical-0223013212010310-2213321322221112-0112132110221332-1220331133130001-3110021221112220-1212220211222023-0000022012130123-3330212022202122"></a>

#### `ingress_rules.action` property

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

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

<a id="canonical-3030110321201232-1333103111021021-1013122302111032-2100203310301332-1032123102300313-1120303310012033-1002130200122222-1132100313233010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.adv_action` properties

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

<a id="canonical-0203233202312031-0302311002201001-0230212001010012-3021322013323211-1303302202213101-1033033122323021-2112323012323313-2110123220203231"></a>

### Direct properties for `ingress_rules.adv_action`

<a id="canonical-0111100203321331-0201221011331102-3032102220303113-2200222311132331-3302000303112103-0330333330120313-3203310233020310-3233130301000220"></a>

#### `ingress_rules.adv_action.action` property

Type: `"string"`. Computed.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Additional upstream details:

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

<a id="canonical-0333020103300021-1330000322120200-3111103330222033-0033220132100300-1013303111230120-0111031211321123-1112231000030320-3130011010032011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.all_tcp_traffic` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.all_tcp_traffic

<a id="canonical-1221003022212223-3230312303323013-3032100022030321-2103310330112322-2013003013331223-3121201322223000-1001030300212032-2313202021120202"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all tcp traffic.

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

<a id="canonical-1021232322221302-2311331101213203-3203323120012311-0013311101301120-0203030302130332-1221023203000220-3031023101131120-3303131302032320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.all_traffic` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.all_traffic

<a id="canonical-1132222022101100-3232002100020221-2023022123233002-1011132020032011-3101322320322232-2101001333120120-2121030320000032-3022331221230302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all traffic.

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

<a id="canonical-3130000110212210-1331231110311102-0310121222230033-0020010303000203-2121111001001220-1333212331101210-3112110332010022-0100321310332010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.all_udp_traffic` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.all_udp_traffic

<a id="canonical-0030102320000212-3212200013330112-2200021202030003-0131131112230003-0022310100232302-1211202123133221-1033113132230021-0033311010333200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all udp traffic.

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

<a id="canonical-3333023301201321-0311011101001132-2200200320012032-3333120203310203-2201203330123033-0331333120123030-0020021202012011-3030203001133020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.any` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.any

<a id="canonical-1122032021000011-2120033132021133-2302203132021113-3013311200230130-1220013130000302-3030012331023230-2301120333213222-2332031311102112"></a>

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

<a id="canonical-3332100121000323-0023323322202301-2001212311212200-3120203333313201-3313131212113023-3221133103211103-2132013313322301-3212330332120302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.applications` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.applications

<a id="canonical-3113030113133301-3222100112201023-3212323131031113-2110321111333303-0131001322220023-3122133331213033-1003213320023310-3130002231222302"></a>

Type: `"single"`. Computed.

Configuration parameter for applications.

Additional upstream details:

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

<a id="canonical-2302231102302010-3102331211232212-3022130232130302-0120332102130320-0212202132121201-0231133130032013-0003112013001032-2232021003033022"></a>

### Direct properties for `ingress_rules.applications`

<a id="canonical-2323331102232311-2210033313221213-2310021121121232-3122010203030103-0201120303013021-1303221202311231-2002111113323322-0022221110311231"></a>

#### `ingress_rules.applications.applications` property

Type: `["list", "string"]`. Computed.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

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

<a id="canonical-3013210110031313-3121232321321212-0321000003332322-3113102131130032-3322312013210103-0132010210333031-3212130330311012-0000023311020130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.inside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.inside_endpoints

<a id="canonical-1311212122331323-1332203103121012-1113202333212033-1111000233131110-3101322301023332-1322033322121000-1203203331332000-2032102002332300"></a>

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

<a id="canonical-3112010312320322-3011123232321023-2203213132012001-0013212231233201-3020330201212031-3123231121312022-2231213231022020-2011120212032020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.ip_prefix_set` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.ip_prefix_set

<a id="canonical-3331320012003021-2310020312333030-2301332201301303-2102212101212213-3220110121333112-2123122103022331-1302211230332332-0011311022200323"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1033213202331300-1020113200100330-2013231332303102-1321112200022321-3021323212310201-1231310122003120-1132022200020232-3113131212210221"></a>

### Direct properties for `ingress_rules.ip_prefix_set`

- [ref](data-sources--network_policy_view--reference--group-001.md#canonical-2203303213220001-3230310310200310-0021000201202332-1122302222101003-0301021112020231-1010221310202101-3302202021331131-2313223002130223): complete subsection reference.

<a id="canonical-2203303213220001-3230310310200310-0021000201202332-1122302222101003-0301021112020231-1010221310202101-3302202021331131-2313223002130223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.ip_prefix_set.ref` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- [ingress_rules.ip_prefix_set](data-sources--network_policy_view--reference--group-001.md#canonical-3112010312320322-3011123232321023-2203213132012001-0013212231233201-3020330201212031-3123231121312022-2231213231022020-2011120212032020)
- ingress_rules.ip_prefix_set.ref

<a id="canonical-0111302112300013-0003132032211121-2120220130113112-0112332312112133-2130211100000210-1032233322022233-1332313312311313-3003113011112002"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1312323120213001-2110233200001232-0320331311222130-3300130113020110-0100301123122023-2020122212223333-0201203230210302-2103203232310321"></a>

### Direct properties for `ingress_rules.ip_prefix_set.ref`

<a id="canonical-2233112311313022-1000120331333301-0212221023302002-0013000333020312-3122312302331311-2130101110300102-1030221003312131-2100020030220011"></a>

#### `ingress_rules.ip_prefix_set.ref.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-0322332331101021-3331020223220222-3233232020011200-2033300032323310-3020032133021312-3220033022302133-1332300321301302-0110332113202332"></a>

<a id="canonical-1210013302121233-3111011230320110-0301122103010330-1233003332210112-0122132032003313-0232030103001322-1211312331220021-1301322222030010"></a>

#### `ingress_rules.ip_prefix_set.ref.name` property

Type: `"string"`. Computed.

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

<a id="canonical-3030100320200323-2010303122023013-0103212003223201-0131303110101301-2113301133210210-0102123221022020-0002210000021133-1300302110021112"></a>

<a id="canonical-0312130211333220-1111210021011132-1023303202220010-0032322220312220-1331301330320320-3012133133222300-2300312200032113-1220302120212113"></a>

#### `ingress_rules.ip_prefix_set.ref.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-3000333012002132-1023223311130101-1301223201100322-3100230013231211-2001330133310202-0000010202003003-2323121121300303-3010030212010032"></a>

<a id="canonical-3003301211330003-1132313000011201-2123202312210131-2323310230330211-3003222101303201-2221122021213103-1201033332023002-2122311032202221"></a>

#### `ingress_rules.ip_prefix_set.ref.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-0121233302302011-2023301121022231-0331000220121100-1033232212031000-0103322030223122-0010313222000032-2100121111210011-0232223202022022"></a>

<a id="canonical-3200012022311031-3212221120122322-0010101031013121-2112233123212121-0301220210130313-1012110022232211-0312110202001010-1300130020202013"></a>

#### `ingress_rules.ip_prefix_set.ref.uid` property

Type: `"string"`. Computed.

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

<a id="canonical-0002213203230022-1013101101002020-1331132233231101-3212230020000002-3000002303100102-3332033230023323-1100022202203223-0313313330203303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.label_matcher` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.label_matcher

<a id="canonical-3102000212100033-3030132311311203-1012100130331301-0233333020202013-2200212201103220-2130332020123110-1321110233321223-3132201133303313"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3123203000123012-2330002102210331-3101203002332331-2021313012201201-1012323321010013-2121012123133302-0320302211132000-3112030332330022"></a>

### Direct properties for `ingress_rules.label_matcher`

<a id="canonical-1011213233010311-1233033233212130-3110321220032333-3222313101231320-2331123130320312-2320310310013323-3130232200230023-1022231031030030"></a>

#### `ingress_rules.label_matcher.keys` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0003203132122230-2030201311201020-1232213332133022-2222321300022321-0003311132113333-1101221030020232-3221103332020201-1222320232020310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.label_selector` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.label_selector

<a id="canonical-2110021000201020-2031310000201302-1203210030010100-1012021212310010-0300010001310133-1300200120331001-0200033121331103-2130131303330302"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2202130022113113-2013023012121213-3213321231013000-2113233223220330-3000323231032023-2330130032221022-3101111103213113-2211121111101310"></a>

### Direct properties for `ingress_rules.label_selector`

<a id="canonical-0302032121101130-3023012223231033-3301110111010131-1202032002322220-1222131011120323-3010213202311023-3310111021230101-2022112020020123"></a>

#### `ingress_rules.label_selector.expressions` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3102200022002323-0102113103000221-0130320232122200-2101133310130230-0133110112333021-2002230311020332-3000020101032101-0111330330323311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.metadata` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.metadata

<a id="canonical-2133131322333313-1103002021110002-1032203233123021-2100312221301322-0022200213311002-3300322002120231-3021313001230213-3312032112233331"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3232320231131313-3200330031200113-3122033111220013-2000211321220132-3103200213320030-2331133112211203-2201102031203001-2220130001001002"></a>

### Direct properties for `ingress_rules.metadata`

<a id="canonical-1332201332333032-0120110032313320-1323321020333332-0131031130212312-0112232023121212-3110031031020020-1012310320031321-1133122312313023"></a>

#### `ingress_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2031111132010112-0320112130302302-2201130211232021-1302100031331203-3330203300000202-3031033122022013-2312230013312333-1303302332133133"></a>

<a id="canonical-0131110012112020-3231011131210231-3122221032100113-2113011203320000-3323302033213300-0121120120202123-2312200203101301-2003123313021320"></a>

#### `ingress_rules.metadata.name` property

Type: `"string"`. Computed.

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

<a id="canonical-0212231310120212-2131122310002101-0331331212312032-1031300013231001-3032312123121210-2120033110322300-0323330022133333-3012000201301322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.outside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.outside_endpoints

<a id="canonical-3011033200020130-1002102113212113-2010010331000320-2100203200232232-1211201100121112-3100333030313221-3023010101323222-1012232032310030"></a>

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

<a id="canonical-0032110322313321-1323000222322302-1132111322200130-1102212012210020-0003221121223000-0100203000323223-3113132222323013-1331002322311022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.prefix_list` properties

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

<a id="canonical-0313113020333023-2320002231332003-3322031302232332-2331010003111200-1212001021012200-2322322203122221-1033320131220221-3200031230312111"></a>

### Direct properties for `ingress_rules.prefix_list`

<a id="canonical-2011313131121021-1322010123002022-1231203011303133-2301322231030220-2221010331311223-2301230210230012-1330220311300113-0131302011020111"></a>

#### `ingress_rules.prefix_list.prefixes` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2103010300101312-2113023010012333-2012111120001202-2032101331122223-1333121202302020-2201110223000222-3010303021001333-2021320103130002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.protocol_port_range` properties

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Property reference](data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [ingress_rules](data-sources--network_policy_view--reference--group-001.md#canonical-0103121203102322-0310232133121202-2211323331121011-0313300122110111-2201222221132100-1302333303021030-1130312230232002-2322101333003230)
- ingress_rules.protocol_port_range

<a id="canonical-2232122101222311-2322222102310230-2210101230021302-3320211001322123-0200130200303203-1120310310000210-2020311100232230-2212232203230023"></a>

Type: `"single"`. Computed.

Protocol and Port. Protocol and Port ranges.

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

<a id="canonical-2323233322330302-2031301102313222-1013000220303022-0110231131303302-2320213222010210-3300103313332002-0200323210031030-1200323313210211"></a>

### Direct properties for `ingress_rules.protocol_port_range`

<a id="canonical-0121211002312013-1023113102332022-2101232233002101-0022001332302020-0120000112133132-2332002201210303-1301320020032222-2331300222100122"></a>

#### `ingress_rules.protocol_port_range.port_ranges` property

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

<a id="canonical-3313112000113311-0010300221131032-2210110011132201-3221332220300232-2010113212230031-0131023323003333-1032031301013033-0332233113303222"></a>

#### `ingress_rules.protocol_port_range.protocol` property

Type: `"string"`. Computed.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

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
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```
