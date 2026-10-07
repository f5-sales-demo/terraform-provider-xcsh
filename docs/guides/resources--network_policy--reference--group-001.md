---
page_title: "xcsh_network_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_policy reference."
---

# xcsh_network_policy reference

<a id="canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- Property reference

<a id="canonical-0113221102011330-3220200022103230-1230200212011102-1311220033000203-1223032020031032-0110110300233333-3022001102330303-0213000011111022"></a>

### Direct properties for `xcsh_network_policy`

<a id="canonical-1020332133010030-2010001331112322-0010132233311002-2330031021012231-2212120200000031-2212023212010221-0011033023110203-0320011011223230"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

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

<a id="canonical-1321100013012323-3020101201232222-2011122323020112-1322123010113011-2212112023212012-3312020130113102-0030233320330130-3013300333021223"></a>

<a id="canonical-0131323300100223-1101200033110332-3231100022010122-1322131233323101-3000232200322202-2030333221313210-1101100111323000-0002133003002321"></a>

#### `description` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2031031323203130-0022110003203332-0110003212330102-0022233230203011-0003122100301231-2110231300312330-1020023210112131-0131133311211100"></a>

<a id="canonical-0123102323222001-3100112200120222-3030331233023332-2113202231123132-1300131330202231-0121331023131033-0212333233333131-0302302133212110"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

A value of true will administratively disable the object.

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

- [endpoint](resources--network_policy--reference--group-001.md#canonical-2230123131001221-0003020123323130-2210111031030032-0120322100010011-3122331012333110-1213102313210220-0223333202132233-0212103232303013): complete subsection reference.

<a id="canonical-1132031220132013-0330130110303332-0330121322233201-0121202001003211-0330103101230213-2200301132232222-2131023331120233-0232130113123001"></a>

<a id="canonical-3302032303230322-3212111233330310-1120000023000303-3313032311130102-3200100321201021-0021320333130012-2231012312010202-2312330021311131"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1113132321131230-3110233221301011-2302003230101233-3212132000003303-2302013301213033-1220312122303002-1301232111001122-3220333232022023"></a>

<a id="canonical-2021031323331312-2113323220022320-2320202321202120-3002123032202331-1211302333213202-2232331211010111-2130203203211003-3012310312211301"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-1021312302320030-3202201322302330-1231310110013013-3022320312130130-0301321001302101-1013210002231003-2023102203210330-2000010321000132"></a>

<a id="canonical-0022133022021102-3130323300032002-2322230030300100-1303300100102113-3313232213031113-0211203012203101-1301020002101311-0133210000200023"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Network Policy. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NameValidator(),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2220321011122310-3100212203331003-2103310223130021-1220130130312001-1320022211032000-1222000301121202-3221121210202323-1202222011332230"></a>

<a id="canonical-0001302020013210-3113120321100132-1323323232220001-1022223002121012-1113103103010201-3013220231222103-1320132020312132-1330331333131021"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Network Policy is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301): complete subsection reference.

- [timeouts](resources--network_policy--reference--group-001.md#canonical-0123011311123100-3002220220320132-2020303012211002-0200230132211012-1010123002001012-1310123220320332-0113122121212101-2123300222322320): complete subsection reference.

<a id="canonical-3200101030233033-1123203202103130-1221021101221020-0213023222200203-3121321123110203-2111321102030031-1233331311321012-3010013032033100"></a>

### All schema paths for `xcsh_network_policy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--network_policy--reference--group-001.md#canonical-1020332133010030-2010001331112322-0010132233311002-2330031021012231-2212120200000031-2212023212010221-0011033023110203-0320011011223230) |
| `description` | [description](resources--network_policy--reference--group-001.md#canonical-1321100013012323-3020101201232222-2011122323020112-1322123010113011-2212112023212012-3312020130113102-0030233320330130-3013300333021223) |
| `disable` | [disable](resources--network_policy--reference--group-001.md#canonical-2031031323203130-0022110003203332-0110003212330102-0022233230203011-0003122100301231-2110231300312330-1020023210112131-0131133311211100) |
| `endpoint` | [endpoint](resources--network_policy--reference--group-001.md#canonical-3002333012010020-3232201213010033-3112313233202322-0000103133301013-2002131011010311-0013122300213133-0030122122333002-2333320122330111) |
| `endpoint.any` | [endpoint.any](resources--network_policy--reference--group-001.md#canonical-0022121012313020-1131121122021120-1122233311123100-1013221001211300-1121303020031301-1003211113003202-0111213013130200-1000111121033322) |
| `endpoint.inside_endpoints` | [endpoint.inside_endpoints](resources--network_policy--reference--group-001.md#canonical-1313313123032012-0200310310031111-2012302001123121-0200122222323310-3232233223321102-1223013311002231-1012033120012311-0301211033220231) |
| `endpoint.label_selector` | [endpoint.label_selector](resources--network_policy--reference--group-001.md#canonical-0232302221330031-2020101322302333-2100211020033021-3110200110222132-2302203211322032-0012210322231001-0203103221221133-1101012233320210) |
| `endpoint.label_selector.expressions` | [endpoint.label_selector.expressions](resources--network_policy--reference--group-001.md#canonical-2021221122230001-2031011311302323-0332323233100010-3032222232123020-2212020130011001-2121312103311113-2231210203030123-2310131223133103) |
| `endpoint.outside_endpoints` | [endpoint.outside_endpoints](resources--network_policy--reference--group-001.md#canonical-3322232021230102-1231211112130303-3213210322033120-3123131200013012-3132313202312101-3331021112222321-1200032310001101-0301103103100202) |
| `endpoint.prefix_list` | [endpoint.prefix_list](resources--network_policy--reference--group-001.md#canonical-2011000313012231-1220130233301332-3323002103002122-0230322320123021-1222120130101110-2220322331210010-0332213300303032-1213210200210000) |
| `endpoint.prefix_list.prefixes` | [endpoint.prefix_list.prefixes](resources--network_policy--reference--group-001.md#canonical-1100123202010001-2112132023302222-2320201310110013-0212020311323220-3000221330213201-0311320112230122-0102111201010211-1313002101022100) |
| `id` | [ID](resources--network_policy--reference--group-001.md#canonical-1132031220132013-0330130110303332-0330121322233201-0121202001003211-0330103101230213-2200301132232222-2131023331120233-0232130113123001) |
| `labels` | [labels](resources--network_policy--reference--group-001.md#canonical-1113132321131230-3110233221301011-2302003230101233-3212132000003303-2302013301213033-1220312122303002-1301232111001122-3220333232022023) |
| `name` | [name](resources--network_policy--reference--group-001.md#canonical-1021312302320030-3202201322302330-1231310110013013-3022320312130130-0301321001302101-1013210002231003-2023102203210330-2000010321000132) |
| `namespace` | [namespace](resources--network_policy--reference--group-001.md#canonical-2220321011122310-3100212203331003-2103310223130021-1220130130312001-1320022211032000-1222000301121202-3221121210202323-1202222011332230) |
| `rules` | [rules](resources--network_policy--reference--group-001.md#canonical-3111020031233311-2310232212311231-3021102331013110-1111010022222213-3131210233111100-2023303112131333-2130100222211030-0201022030003012) |
| `rules.egress_rules` | [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-0222111333300322-3221111020331212-0323032023323201-3010133212233202-0301030331230000-1223202010013223-2221012222120100-0301200002130113) |
| `rules.egress_rules.action` | [rules.egress_rules.action](resources--network_policy--reference--group-001.md#canonical-1232032331330033-3103233012101211-2213121033210323-0333222013200201-1110020331311013-1300321130112120-3021023132303102-3220302323120132) |
| `rules.egress_rules.adv_action` | [rules.egress_rules.adv_action](resources--network_policy--reference--group-001.md#canonical-3203032021230120-1331201132000113-0301113003211021-0100222020021303-3101112013211232-0311330301032320-2332232312232002-0033133312001103) |
| `rules.egress_rules.adv_action.action` | [rules.egress_rules.adv_action.action](resources--network_policy--reference--group-001.md#canonical-1023333011113010-3103212020110223-0201012103120230-3231310000000303-2200222300230201-2020313111011101-0313202302211032-0220323111020002) |
| `rules.egress_rules.all_tcp_traffic` | [rules.egress_rules.all_tcp_traffic](resources--network_policy--reference--group-001.md#canonical-3300230133121031-2203233003322233-0031302202333213-2202333212120320-1201230300123001-0200300203100330-2222013103001300-1320011122333001) |
| `rules.egress_rules.all_traffic` | [rules.egress_rules.all_traffic](resources--network_policy--reference--group-001.md#canonical-1223311201110300-1331200300010323-1112331032212130-3203131300320223-2132330302021112-0131123132203312-0302010001002303-2212100233000200) |
| `rules.egress_rules.all_udp_traffic` | [rules.egress_rules.all_udp_traffic](resources--network_policy--reference--group-001.md#canonical-1102330320111233-0012111322231002-1003333223200201-2320032303001301-2112233323302021-0123210201030323-3221321311002012-2131330321030112) |
| `rules.egress_rules.any` | [rules.egress_rules.any](resources--network_policy--reference--group-001.md#canonical-3233233310330001-3101300031003010-3200032201111230-1010111020210332-3213200210011120-3320022013120200-0211300033201232-3220003112031000) |
| `rules.egress_rules.applications` | [rules.egress_rules.applications](resources--network_policy--reference--group-001.md#canonical-2231112103212301-3101333133200201-2123230021333031-2210103111020223-2133113030013133-1221221100202120-2133022310312112-1113032010100330) |
| `rules.egress_rules.applications.applications` | [rules.egress_rules.applications.applications](resources--network_policy--reference--group-001.md#canonical-2131221231313121-1010031200003031-1233002232202112-0100121023303102-3232220002313130-1312320201003110-2331010022120310-3333310200121321) |
| `rules.egress_rules.inside_endpoints` | [rules.egress_rules.inside_endpoints](resources--network_policy--reference--group-001.md#canonical-0131131301033003-3013003011012111-3011010123320303-1033300322333103-3120231233210130-3130103201202030-3322123131223330-0201230220220130) |
| `rules.egress_rules.ip_prefix_set` | [rules.egress_rules.ip_prefix_set](resources--network_policy--reference--group-001.md#canonical-3313132012301331-3333220231123132-0103030203131202-2202203230330332-3132210033311320-1332121101213300-1121113200320222-0313333112232233) |
| `rules.egress_rules.ip_prefix_set.ref` | [rules.egress_rules.ip_prefix_set.ref](resources--network_policy--reference--group-001.md#canonical-1232033133200101-2030311000211333-0123203202302333-0201232313321022-3300103020221210-3031222211002232-2112121030120003-1201102323010000) |
| `rules.egress_rules.ip_prefix_set.ref.kind` | [rules.egress_rules.ip_prefix_set.ref.kind](resources--network_policy--reference--group-001.md#canonical-2333011002131000-1320031203022231-1131231303032223-2300132301032121-0312330211003100-2111012312222132-0012103111132101-1121331322200032) |
| `rules.egress_rules.ip_prefix_set.ref.name` | [rules.egress_rules.ip_prefix_set.ref.name](resources--network_policy--reference--group-001.md#canonical-3101330323030221-0113311300313020-2123103333003313-2223231331211013-3220010010333111-1020213102203002-0312013211300122-2301013121010130) |
| `rules.egress_rules.ip_prefix_set.ref.namespace` | [rules.egress_rules.ip_prefix_set.ref.namespace](resources--network_policy--reference--group-001.md#canonical-1022201331000103-1221112331001102-2030113230301121-2200013221202021-3123103231032200-2310333201221030-0003210102222202-3232310110303033) |
| `rules.egress_rules.ip_prefix_set.ref.tenant` | [rules.egress_rules.ip_prefix_set.ref.tenant](resources--network_policy--reference--group-001.md#canonical-0220011022201303-1202303232303300-0321223112100023-3330001132330323-0122030133222122-2012130033100232-3310323310111120-3120222021023212) |
| `rules.egress_rules.ip_prefix_set.ref.uid` | [rules.egress_rules.ip_prefix_set.ref.uid](resources--network_policy--reference--group-001.md#canonical-2121212310200012-0200113330230322-1201213330102002-1313000010010013-2223001002013232-3111303220312311-2022123333000230-0003002013201211) |
| `rules.egress_rules.label_matcher` | [rules.egress_rules.label_matcher](resources--network_policy--reference--group-001.md#canonical-0120023103213100-3322220320321212-2132130000122030-1112110002310231-1101002221023032-1213333030012313-1103200333232003-3111030321103023) |
| `rules.egress_rules.label_matcher.keys` | [rules.egress_rules.label_matcher.keys](resources--network_policy--reference--group-001.md#canonical-1312331330100221-0223123032010220-0320110332203223-0333000001213022-2212020012203122-1203233213332033-1121210230333111-1331002323202203) |
| `rules.egress_rules.label_selector` | [rules.egress_rules.label_selector](resources--network_policy--reference--group-001.md#canonical-0332213110021323-2100000311311003-0001233301222221-1000033211110230-0323300220310123-0102012222113122-3001321332222131-0322200222331012) |
| `rules.egress_rules.label_selector.expressions` | [rules.egress_rules.label_selector.expressions](resources--network_policy--reference--group-001.md#canonical-0113111320332202-2221333230331230-1023123320222302-0012032001323002-2001100011003301-0103331121123320-1301120012113033-3233301330322320) |
| `rules.egress_rules.metadata` | [rules.egress_rules.metadata](resources--network_policy--reference--group-001.md#canonical-1232031010230011-2011013233200021-0111223212330122-2003132000130220-1111023300213222-0000201230112030-0123132210132333-3233101302103323) |
| `rules.egress_rules.metadata.description_spec` | [rules.egress_rules.metadata.description_spec](resources--network_policy--reference--group-001.md#canonical-3301231122103222-2323010020130101-1320222230312200-2012301100033013-2110000122023032-1223310102211333-2230330232022102-2123323100112210) |
| `rules.egress_rules.metadata.name` | [rules.egress_rules.metadata.name](resources--network_policy--reference--group-001.md#canonical-3112232122220113-0022031121100313-2101123121101120-2020122012332120-0032101001330202-2103303200111301-1033333200220232-1023221002010231) |
| `rules.egress_rules.outside_endpoints` | [rules.egress_rules.outside_endpoints](resources--network_policy--reference--group-001.md#canonical-3010221002300013-2030312313011331-0102001313020113-3301200321000102-2131333230203232-3112202021111221-0303023331333303-0300131030203322) |
| `rules.egress_rules.prefix_list` | [rules.egress_rules.prefix_list](resources--network_policy--reference--group-001.md#canonical-2223223230313301-2300130323000023-3303321032323132-0112332020202221-2020120002321110-0023213213101000-1321201111321133-2121300301310311) |
| `rules.egress_rules.prefix_list.prefixes` | [rules.egress_rules.prefix_list.prefixes](resources--network_policy--reference--group-001.md#canonical-0220030332211201-0032233312113033-3130113112211113-1321003230302200-1123000330013103-0032323330133222-3131123000313031-3131313332133300) |
| `rules.egress_rules.protocol_port_range` | [rules.egress_rules.protocol_port_range](resources--network_policy--reference--group-001.md#canonical-2111133011031032-0032003203302101-0321320013030110-2301312102220022-2000312030222110-2111131033132302-0010001202232330-3021212102233233) |
| `rules.egress_rules.protocol_port_range.port_ranges` | [rules.egress_rules.protocol_port_range.port_ranges](resources--network_policy--reference--group-001.md#canonical-3120003032233123-3221310221331202-0223121330121232-0230232013223312-0301103110300211-0102202333003020-1330220130132112-1312203102030333) |
| `rules.egress_rules.protocol_port_range.protocol` | [rules.egress_rules.protocol_port_range.protocol](resources--network_policy--reference--group-001.md#canonical-1111103130230011-0131013333201100-3221232120121233-3221021332322230-1121000200133333-1223123333230323-2222201302201013-0100013300202330) |
| `rules.ingress_rules` | [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-3110332322122213-0300133202112202-2210212303331212-2132302022002321-3111320030211012-2310031203102101-3231213122023013-0032003303331023) |
| `rules.ingress_rules.action` | [rules.ingress_rules.action](resources--network_policy--reference--group-001.md#canonical-0021320201303023-3111302322323230-1331100012023213-0131101232020301-2310331002103331-1202033320232230-0032131203210301-3110012203332032) |
| `rules.ingress_rules.adv_action` | [rules.ingress_rules.adv_action](resources--network_policy--reference--group-001.md#canonical-1220013122332021-0213331030221201-3212022022232001-1122320031231130-1013303303320330-1012011102202202-0310110133031123-0131203231230200) |
| `rules.ingress_rules.adv_action.action` | [rules.ingress_rules.adv_action.action](resources--network_policy--reference--group-001.md#canonical-3310211212221032-3310321330001132-1010001332111123-1332112201312221-0113100130332231-3002212333120203-1212332031133001-1022302111010010) |
| `rules.ingress_rules.all_tcp_traffic` | [rules.ingress_rules.all_tcp_traffic](resources--network_policy--reference--group-001.md#canonical-2302103301120130-2303332132211100-2220122300311203-2222200013301200-2013231113310123-3210130123233001-3231303122210302-1121332022010211) |
| `rules.ingress_rules.all_traffic` | [rules.ingress_rules.all_traffic](resources--network_policy--reference--group-001.md#canonical-1312111031123211-0330311330122213-1012110301332121-2310011112330032-0010100022212231-1120332213231200-1011213202222203-0303021013213123) |
| `rules.ingress_rules.all_udp_traffic` | [rules.ingress_rules.all_udp_traffic](resources--network_policy--reference--group-001.md#canonical-3031100133331303-1303003003103303-1012212302122032-2013130001321333-3210030130320302-0322322103012203-2333201303101010-3020001022222123) |
| `rules.ingress_rules.any` | [rules.ingress_rules.any](resources--network_policy--reference--group-001.md#canonical-2122201312333031-3100010130010322-0103211233103322-2021002011110013-2113012331010103-2101023010111320-2103103220033100-3112321323323000) |
| `rules.ingress_rules.applications` | [rules.ingress_rules.applications](resources--network_policy--reference--group-001.md#canonical-3313211102033112-2032133221121312-3031121301013310-0010123132021101-2313230022300210-1112301320300020-3110220212300302-0002232033032010) |
| `rules.ingress_rules.applications.applications` | [rules.ingress_rules.applications.applications](resources--network_policy--reference--group-001.md#canonical-3300032130130222-0312332133032200-3332223101102003-1002121023021211-3220002333130122-2100230031333010-1031013100230133-0033030033313000) |
| `rules.ingress_rules.inside_endpoints` | [rules.ingress_rules.inside_endpoints](resources--network_policy--reference--group-001.md#canonical-0203233132201312-2123333312210201-0101023100322010-0012230321302033-2312210231212123-2000333330110303-3212022200323132-2100332333332023) |
| `rules.ingress_rules.ip_prefix_set` | [rules.ingress_rules.ip_prefix_set](resources--network_policy--reference--group-001.md#canonical-2222112312031322-3313312203031222-2323222023133201-2101223323110331-2202131022201201-3021211231223112-0110032302123311-0031023311313230) |
| `rules.ingress_rules.ip_prefix_set.ref` | [rules.ingress_rules.ip_prefix_set.ref](resources--network_policy--reference--group-001.md#canonical-3222030110232132-0322220102322012-3130032222032020-2133101130202230-3000011030221020-0130321030313000-1321030323111130-0102211203201321) |
| `rules.ingress_rules.ip_prefix_set.ref.kind` | [rules.ingress_rules.ip_prefix_set.ref.kind](resources--network_policy--reference--group-001.md#canonical-0123233230323131-0031112011003011-3200301032332022-2233000233301301-3012202113123301-3202113231120020-2010132010101110-2223120120321010) |
| `rules.ingress_rules.ip_prefix_set.ref.name` | [rules.ingress_rules.ip_prefix_set.ref.name](resources--network_policy--reference--group-001.md#canonical-3311133102010333-3220323301202210-2330031301330221-2020232031020001-3303011120333320-2131113011310133-0023202012232233-2130203110001133) |
| `rules.ingress_rules.ip_prefix_set.ref.namespace` | [rules.ingress_rules.ip_prefix_set.ref.namespace](resources--network_policy--reference--group-001.md#canonical-0121130131131333-0132200032322211-3120221301100232-0002022222122112-1122320203222103-3223103310310313-1331212213131103-2213032022113301) |
| `rules.ingress_rules.ip_prefix_set.ref.tenant` | [rules.ingress_rules.ip_prefix_set.ref.tenant](resources--network_policy--reference--group-001.md#canonical-1211112323132123-1302223303112220-0310323031202031-2012202323201120-1123001101203203-1133330210302321-2230032200312003-1322203020131213) |
| `rules.ingress_rules.ip_prefix_set.ref.uid` | [rules.ingress_rules.ip_prefix_set.ref.uid](resources--network_policy--reference--group-001.md#canonical-0021031122211003-3133023121233032-3203321101032321-0320212003022102-3322133012211312-1113013130103021-1321023230121131-1121300123030111) |
| `rules.ingress_rules.label_matcher` | [rules.ingress_rules.label_matcher](resources--network_policy--reference--group-001.md#canonical-1333130130321112-2010123200001331-0120200031030320-2212333001103023-2303301111012331-0212132201000333-1130312133021010-2202302230301012) |
| `rules.ingress_rules.label_matcher.keys` | [rules.ingress_rules.label_matcher.keys](resources--network_policy--reference--group-001.md#canonical-2113020322012100-1020011031101100-0231013202020002-1313032312022020-0212310221231101-2320320113021013-1323212032111201-0200113022301133) |
| `rules.ingress_rules.label_selector` | [rules.ingress_rules.label_selector](resources--network_policy--reference--group-001.md#canonical-3222022203011122-0123101120332020-2233331101010011-3202211323201302-3123101203300110-3313223232003200-1230131211321020-0313330330013133) |
| `rules.ingress_rules.label_selector.expressions` | [rules.ingress_rules.label_selector.expressions](resources--network_policy--reference--group-001.md#canonical-1300122112113230-1022332310001330-2333133300022132-1302302120303311-2320221331032020-0310011313322001-0210132200321312-1001130212111111) |
| `rules.ingress_rules.metadata` | [rules.ingress_rules.metadata](resources--network_policy--reference--group-001.md#canonical-0303332303103102-3001032322121223-2320302023230010-0300100133232131-1020322111121022-2300121101332312-0212000010220003-3103313013031322) |
| `rules.ingress_rules.metadata.description_spec` | [rules.ingress_rules.metadata.description_spec](resources--network_policy--reference--group-001.md#canonical-3303101230212330-2322221010332000-3032202030213301-0022022221233212-3231202121121010-1103332331312022-1002121033031300-2233121231212112) |
| `rules.ingress_rules.metadata.name` | [rules.ingress_rules.metadata.name](resources--network_policy--reference--group-001.md#canonical-1331112110102031-2022203033201100-3123223132220030-1302002011301113-0001233020012120-1131133310121222-2202321021323203-1312013320130320) |
| `rules.ingress_rules.outside_endpoints` | [rules.ingress_rules.outside_endpoints](resources--network_policy--reference--group-001.md#canonical-2212230201212211-1211033233210022-3210202103103122-2333103103311120-1321203023020311-3221203230310002-3223201312221321-0123110123103322) |
| `rules.ingress_rules.prefix_list` | [rules.ingress_rules.prefix_list](resources--network_policy--reference--group-001.md#canonical-2012022101012232-2303001031023333-2000022133211112-3033002122213212-2111113320133232-2030332213012331-3301101103200111-3320202210302330) |
| `rules.ingress_rules.prefix_list.prefixes` | [rules.ingress_rules.prefix_list.prefixes](resources--network_policy--reference--group-001.md#canonical-0112130100210313-0232213113202220-3021200221201123-2032000120313111-3133021303000211-3002111030310303-3212321230112322-0111333230302133) |
| `rules.ingress_rules.protocol_port_range` | [rules.ingress_rules.protocol_port_range](resources--network_policy--reference--group-001.md#canonical-0020122023310003-2100330120110231-0103302131121033-3130220021112013-0010333022300322-2001111132310031-2212221330122133-2113331212021331) |
| `rules.ingress_rules.protocol_port_range.port_ranges` | [rules.ingress_rules.protocol_port_range.port_ranges](resources--network_policy--reference--group-001.md#canonical-1321121233203032-3023221101331111-2223000302010230-2000303030202203-0011113132010331-3020220100320212-2311333012101113-0023333303203223) |
| `rules.ingress_rules.protocol_port_range.protocol` | [rules.ingress_rules.protocol_port_range.protocol](resources--network_policy--reference--group-001.md#canonical-0111020233113013-2111132312010233-3121123120122232-3132231021220200-2132001121033102-3301230110003323-1130312121021011-0013100200021030) |
| `timeouts` | [timeouts](resources--network_policy--reference--group-001.md#canonical-1131121301101001-0223322023322110-3132001220000313-2303003301030112-0110133222031321-2011203020111003-3013213031211203-0120023110201211) |
| `timeouts.create` | [timeouts.create](resources--network_policy--reference--group-001.md#canonical-3301300023023131-2223230301131313-2211020332000002-0331331203021202-0003213212211322-2211332031010113-1003112022213203-3303113311003030) |
| `timeouts.delete` | [timeouts.delete](resources--network_policy--reference--group-001.md#canonical-0331222113122301-3203323211212032-1333221011331121-0022212032120023-3121212033201212-1021300200112012-0311033100022232-1023033030013210) |
| `timeouts.read` | [timeouts.read](resources--network_policy--reference--group-001.md#canonical-1321110213331313-0002012201130302-0331103130122112-2032021232102111-3333110213331323-0301031213330202-0233321033121031-2330010020030001) |
| `timeouts.update` | [timeouts.update](resources--network_policy--reference--group-001.md#canonical-1220210332121332-2220232123232320-0232120331200021-3110113012203333-2021002321221032-2103233012223301-3023020333232310-3022202313113203) |

<a id="canonical-2230123131001221-0003020123323130-2210111031030032-0120322100010011-3122331012333110-1213102313210220-0223333202132233-0212103232303013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- endpoint

<a id="canonical-3002333012010020-3232201213010033-3112313233202322-0000103133301013-2002131011010311-0013122300213133-0030122122333002-2333320122330111"></a>

Type: `"object"`. single nested block, Optional.

Shape of the endpoint choices for a view.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("any",
    "inside_endpoints"),
  validators.ConflictingObjectAttributes("any",
    "label_selector"),
  validators.ConflictingObjectAttributes("any",
    "outside_endpoints"),
  validators.ConflictingObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingObjectAttributes("inside_endpoints",
    "label_selector"),
  validators.ConflictingObjectAttributes("inside_endpoints",
    "outside_endpoints"),
  validators.ConflictingObjectAttributes("inside_endpoints",
    "prefix_list"),
  validators.ConflictingObjectAttributes("label_selector",
    "outside_endpoints"),
  validators.ConflictingObjectAttributes("label_selector",
    "prefix_list"),
  validators.ConflictingObjectAttributes("outside_endpoints",
    "prefix_list")}
```

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

Terraform syntax:

```terraform
endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002212011123130-2102200131102101-2323310200101031-0032020200013323-1310133002220113-1111301330103300-3210111311202322-3003002223122010"></a>

### Direct properties for `endpoint`

- [any](resources--network_policy--reference--group-001.md#canonical-0121302023111133-0330100131013110-3210002030111101-1133332330132321-2312031223131003-3323232221023130-3132222310232331-3213102030300113): complete subsection reference.

- [inside_endpoints](resources--network_policy--reference--group-001.md#canonical-0232222033301310-1011011021331010-0122120033333302-3300213020031332-2302231000210000-1231102022100310-3232223322113101-3313131200300031): complete subsection reference.

- [label_selector](resources--network_policy--reference--group-001.md#canonical-3023232323102223-2311322302302201-2120320103200030-3330210023113100-1231322131310332-3023001023233323-2101032201200320-2021320023132122): complete subsection reference.

- [outside_endpoints](resources--network_policy--reference--group-001.md#canonical-2321133302230113-0210102311233333-2231232132213123-3310301132230122-0211100232123003-3000032000021030-3302010321231202-2002033112122001): complete subsection reference.

- [prefix_list](resources--network_policy--reference--group-001.md#canonical-0101101133233231-3100131210320010-1203111023320221-3321020312231322-0220311230111103-1131333211323122-1203321221111130-1032112133233200): complete subsection reference.

<a id="canonical-0121302023111133-0330100131013110-3210002030111101-1133332330132321-2312031223131003-3323232221023130-3132222310232331-3213102030300113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint.any` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [endpoint](resources--network_policy--reference--group-001.md#canonical-2230123131001221-0003020123323130-2210111031030032-0120322100010011-3122331012333110-1213102313210220-0223333202132233-0212103232303013)
- endpoint.any

<a id="canonical-0022121012313020-1131121122021120-1122233311123100-1013221001211300-1121303020031301-1003211113003202-0111213013130200-1000111121033322"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
any = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232222033301310-1011011021331010-0122120033333302-3300213020031332-2302231000210000-1231102022100310-3232223322113101-3313131200300031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint.inside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [endpoint](resources--network_policy--reference--group-001.md#canonical-2230123131001221-0003020123323130-2210111031030032-0120322100010011-3122331012333110-1213102313210220-0223333202132233-0212103232303013)
- endpoint.inside_endpoints

<a id="canonical-1313313123032012-0200310310031111-2012302001123121-0200122222323310-3232233223321102-1223013311002231-1012033120012311-0301211033220231"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
inside_endpoints = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023232323102223-2311322302302201-2120320103200030-3330210023113100-1231322131310332-3023001023233323-2101032201200320-2021320023132122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint.label_selector` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [endpoint](resources--network_policy--reference--group-001.md#canonical-2230123131001221-0003020123323130-2210111031030032-0120322100010011-3122331012333110-1213102313210220-0223333202132233-0212103232303013)
- endpoint.label_selector

<a id="canonical-0232302221330031-2020101322302333-2100211020033021-3110200110222132-2302203211322032-0012210322231001-0203103221221133-1101012233320210"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
```

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

Terraform syntax:

```terraform
label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332112020331223-0312011112321032-3200110021220120-0031011332223103-3333032022010021-3113101221320000-1313001303122131-0123123002032321"></a>

### Direct properties for `endpoint.label_selector`

<a id="canonical-2021221122230001-2031011311302323-0332323233100010-3032222232123020-2212020130011001-2121312103311113-2231210203030123-2310131223133103"></a>

#### `endpoint.label_selector.expressions` property

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2321133302230113-0210102311233333-2231232132213123-3310301132230122-0211100232123003-3000032000021030-3302010321231202-2002033112122001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint.outside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [endpoint](resources--network_policy--reference--group-001.md#canonical-2230123131001221-0003020123323130-2210111031030032-0120322100010011-3122331012333110-1213102313210220-0223333202132233-0212103232303013)
- endpoint.outside_endpoints

<a id="canonical-3322232021230102-1231211112130303-3213210322033120-3123131200013012-3132313202312101-3331021112222321-1200032310001101-0301103103100202"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
outside_endpoints = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101101133233231-3100131210320010-1203111023320221-3321020312231322-0220311230111103-1131333211323122-1203321221111130-1032112133233200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint.prefix_list` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [endpoint](resources--network_policy--reference--group-001.md#canonical-2230123131001221-0003020123323130-2210111031030032-0120322100010011-3122331012333110-1213102313210220-0223333202132233-0212103232303013)
- endpoint.prefix_list

<a id="canonical-2011000313012231-1220130233301332-3323002103002122-0230322320123021-1222120130101110-2220322331210010-0332213300303032-1213210200210000"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110322220133002-0110210231333323-2103220111221033-2111011220100310-0121100222231311-1210210130013001-0233220220302001-0031210013232213"></a>

### Direct properties for `endpoint.prefix_list`

<a id="canonical-1100123202010001-2112132023302222-2320201310110013-0212020311323220-3000221330213201-0311320112230122-0102111201010211-1313002101022100"></a>

#### `endpoint.prefix_list.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- rules

<a id="canonical-3111020031233311-2310232212311231-3021102331013110-1111010022222213-3131210233111100-2023303112131333-2130100222211030-0201022030003012"></a>

Type: `"object"`. single nested block, Optional.

Rule Choice. Shape of Rule Choice.

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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121012232232032-3103101323132322-0021123233120110-1033023301222331-3002022303021303-0112320211001023-2122030212201013-2021031010200300"></a>

### Direct properties for `rules`

- [egress_rules](resources--network_policy--reference--group-001.md#canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203): complete subsection reference.

- [ingress_rules](resources--network_policy--reference--group-001.md#canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133): complete subsection reference.

<a id="canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.egress_rules` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- rules.egress_rules

<a id="canonical-0222111333300322-3221111020331212-0323032023323201-3010133212233202-0301030331230000-1223202010013223-2221012222120100-0301200002130113"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of rules applied to connections from policy endpoints.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("any",
    "inside_endpoints"),
  validators.ConflictingListObjectAttributes("any",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("any",
    "label_selector"),
  validators.ConflictingListObjectAttributes("any",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("applications",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "label_selector"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "label_selector"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("label_selector",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("label_selector",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("outside_endpoints",
    "prefix_list")}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
egress_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3120233121111222-2121013100211222-0321231132333200-2320103120030003-3102231131213321-0111301003132102-1032201012022132-1210231221202002"></a>

### Direct properties for `rules.egress_rules`

<a id="canonical-1232032331330033-3103233012101211-2213121033210323-0333222013200201-1110020331311013-1300321130112120-3021023132303102-3220302323120132"></a>

#### `rules.egress_rules.action` property

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ALLOW","DENY"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW"),
}
```

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

- [adv_action](resources--network_policy--reference--group-001.md#canonical-3113033020313203-3103001031013322-1333022213323021-3310121100033131-2000233301323330-1302300101320012-2211100010101311-0011101033003313): complete subsection reference.

- [all_tcp_traffic](resources--network_policy--reference--group-001.md#canonical-0320210210333303-2300210121233012-0023202231311231-0313023121323213-3133303212220203-3031310113312211-3133023332310222-3113011102102323): complete subsection reference.

- [all_traffic](resources--network_policy--reference--group-001.md#canonical-3031111320200023-0311321120321113-0201111001202213-3301103300102221-0200120012110000-3323030121230301-1132130201133332-1103200313020230): complete subsection reference.

- [all_udp_traffic](resources--network_policy--reference--group-001.md#canonical-3122203212023100-0323303031302322-1110302012130131-0303330321001013-0333212110323233-0233000301011311-3103202300223330-1310131031213330): complete subsection reference.

- [any](resources--network_policy--reference--group-001.md#canonical-3311111111300210-2212311131133300-3303113211233233-1203132023301032-0102100032222233-3312011031331013-3031322201230021-0212000322011312): complete subsection reference.

- [applications](resources--network_policy--reference--group-001.md#canonical-2322320101222210-1332012011330001-3102221003230012-3021330221030223-3103003001320303-0103002232130102-1311102033220300-0012002023130210): complete subsection reference.

- [inside_endpoints](resources--network_policy--reference--group-001.md#canonical-3132211001312020-3013331233311131-2231331331022110-2030001210132220-0201113220233313-3332300032032033-2211300222303111-1211201221121213): complete subsection reference.

- [ip_prefix_set](resources--network_policy--reference--group-001.md#canonical-0132103010030311-1033233100012302-0121033213310011-2312021232331203-1311112012232203-3311122002212323-1333112021330102-3222210311012011): complete subsection reference.

- [label_matcher](resources--network_policy--reference--group-001.md#canonical-0313303001231300-1213322213333202-1101303032130212-0111233101123313-1111033113303113-2011331133323133-0102020011023111-2103320112031303): complete subsection reference.

- [label_selector](resources--network_policy--reference--group-001.md#canonical-0313200233022122-1011111121221330-2312232210123301-3133103211220033-2131200223302212-0320221220020302-2210313132123030-1030313120102212): complete subsection reference.

- [metadata](resources--network_policy--reference--group-001.md#canonical-0313312122321203-1303132013202013-0110131111110032-3232320002323023-2120320132120313-3312123033220001-2121201303203222-2323303003031231): complete subsection reference.

- [outside_endpoints](resources--network_policy--reference--group-001.md#canonical-2110323021333030-1213332000232321-2013312202033032-1011202303331301-2100120233103222-0020122011212320-3001100002023303-2032322223102223): complete subsection reference.

- [prefix_list](resources--network_policy--reference--group-001.md#canonical-2110233120231311-2223011002301312-1320100323032222-3131311313212032-3232103221213103-0000131103222233-1332300102302301-0101323320333100): complete subsection reference.

- [protocol_port_range](resources--network_policy--reference--group-001.md#canonical-2012213131312232-1113312032121130-0210323122310013-2201130220112023-3013002302203233-3301202123303110-2223102202313111-2323300301101122): complete subsection reference.

<a id="canonical-3113033020313203-3103001031013322-1333022213323021-3310121100033131-2000233301323330-1302300101320012-2211100010101311-0011101033003313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.egress_rules.adv_action` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203)
- rules.egress_rules.adv_action

<a id="canonical-3203032021230120-1331201132000113-0301113003211021-0100222020021303-3101112013211232-0311330301032320-2332232312232002-0033133312001103"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
adv_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102333202021323-0320100222310213-0120031022322011-0311203000210110-0213012023002231-2003102033303011-3221102213302333-0103220231130313"></a>

### Direct properties for `rules.egress_rules.adv_action`

<a id="canonical-1023333011113010-3103212020110223-0201012103120230-3231310000000303-2200222300230201-2020313111011101-0313202302211032-0220323111020002"></a>

#### `rules.egress_rules.adv_action.action` property

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["LOG","NOLOG"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("NOLOG",
    "LOG"),
}
```

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

<a id="canonical-0320210210333303-2300210121233012-0023202231311231-0313023121323213-3133303212220203-3031310113312211-3133023332310222-3113011102102323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.egress_rules.all_tcp_traffic` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203)
- rules.egress_rules.all_tcp_traffic

<a id="canonical-3300230133121031-2203233003322233-0031302202333213-2202333212120320-1201230300123001-0200300203100330-2222013103001300-1320011122333001"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
all_tcp_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031111320200023-0311321120321113-0201111001202213-3301103300102221-0200120012110000-3323030121230301-1132130201133332-1103200313020230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.egress_rules.all_traffic` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203)
- rules.egress_rules.all_traffic

<a id="canonical-1223311201110300-1331200300010323-1112331032212130-3203131300320223-2132330302021112-0131123132203312-0302010001002303-2212100233000200"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
all_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122203212023100-0323303031302322-1110302012130131-0303330321001013-0333212110323233-0233000301011311-3103202300223330-1310131031213330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.egress_rules.all_udp_traffic` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203)
- rules.egress_rules.all_udp_traffic

<a id="canonical-1102330320111233-0012111322231002-1003333223200201-2320032303001301-2112233323302021-0123210201030323-3221321311002012-2131330321030112"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
all_udp_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311111111300210-2212311131133300-3303113211233233-1203132023301032-0102100032222233-3312011031331013-3031322201230021-0212000322011312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.egress_rules.any` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203)
- rules.egress_rules.any

<a id="canonical-3233233310330001-3101300031003010-3200032201111230-1010111020210332-3213200210011120-3320022013120200-0211300033201232-3220003112031000"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
any = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322320101222210-1332012011330001-3102221003230012-3021330221030223-3103003001320303-0103002232130102-1311102033220300-0012002023130210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.egress_rules.applications` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203)
- rules.egress_rules.applications

<a id="canonical-2231112103212301-3101333133200201-2123230021333031-2210103111020223-2133113030013133-1221221100202120-2133022310312112-1113032010100330"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
applications {
  # Configure direct properties listed below.
}
```

<a id="canonical-2022200133131231-2312023300301121-2212223330023110-0200300310331031-2122323013020013-1023233110221233-3233013213213132-1132203213300223"></a>

### Direct properties for `rules.egress_rules.applications`

<a id="canonical-2131221231313121-1010031200003031-1233002232202112-0100121023303102-3232220002313130-1312320201003110-2331010022120310-3333310200121321"></a>

#### `rules.egress_rules.applications.applications` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-3132211001312020-3013331233311131-2231331331022110-2030001210132220-0201113220233313-3332300032032033-2211300222303111-1211201221121213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.egress_rules.inside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203)
- rules.egress_rules.inside_endpoints

<a id="canonical-0131131301033003-3013003011012111-3011010123320303-1033300322333103-3120231233210130-3130103201202030-3322123131223330-0201230220220130"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
inside_endpoints = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132103010030311-1033233100012302-0121033213310011-2312021232331203-1311112012232203-3311122002212323-1333112021330102-3222210311012011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.egress_rules.ip_prefix_set` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203)
- rules.egress_rules.ip_prefix_set

<a id="canonical-3313132012301331-3333220231123132-0103030203131202-2202203230330332-3132210033311320-1332121101213300-1121113200320222-0313333112232233"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233300313312320-3001322321201003-1212213213130003-3101130232221122-3031220020003132-0001010013121230-0012001223122021-0223230012311220"></a>

### Direct properties for `rules.egress_rules.ip_prefix_set`

- [ref](resources--network_policy--reference--group-001.md#canonical-0012211012202032-2211230110131312-0122310000131223-0201201322312111-2220113023111332-2020310112130333-0011132132332110-1200112123122330): complete subsection reference.

<a id="canonical-0012211012202032-2211230110131312-0122310000131223-0201201322312111-2220113023111332-2020310112130333-0011132132332110-1200112123122330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.egress_rules.ip_prefix_set.ref` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203)
- [rules.egress_rules.ip_prefix_set](resources--network_policy--reference--group-001.md#canonical-0132103010030311-1033233100012302-0121033213310011-2312021232331203-1311112012232203-3311122002212323-1333112021330102-3222210311012011)
- rules.egress_rules.ip_prefix_set.ref

<a id="canonical-1232033133200101-2030311000211333-0123203202302333-0201232313321022-3300103020221210-3031222211002232-2112121030120003-1201102323010000"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302331101330000-2300011002323220-0111121232211021-2323303010310311-2010030333333010-0312330112000011-2123032112120332-3222332011132230"></a>

### Direct properties for `rules.egress_rules.ip_prefix_set.ref`

<a id="canonical-2333011002131000-1320031203022231-1131231303032223-2300132301032121-0312330211003100-2111012312222132-0012103111132101-1121331322200032"></a>

#### `rules.egress_rules.ip_prefix_set.ref.kind` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3101330323030221-0113311300313020-2123103333003313-2223231331211013-3220010010333111-1020213102203002-0312013211300122-2301013121010130"></a>

<a id="canonical-2003112301130130-2101120032032000-3303203301233230-2322020331131300-1131130200032231-0323303013003032-0320321113232302-3323113011032212"></a>

#### `rules.egress_rules.ip_prefix_set.ref.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1022201331000103-1221112331001102-2030113230301121-2200013221202021-3123103231032200-2310333201221030-0003210102222202-3232310110303033"></a>

<a id="canonical-0030122330303213-0333221332101210-0311332121112330-0002010321322222-3332132102133222-2333320201032312-0223032010302211-0322203203222030"></a>

#### `rules.egress_rules.ip_prefix_set.ref.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0220011022201303-1202303232303300-0321223112100023-3330001132330323-0122030133222122-2012130033100232-3310323310111120-3120222021023212"></a>

<a id="canonical-2211212210332121-3330210222203000-1222332031113030-3011200113321202-0113311322020122-2122332210122020-2212223321013230-1300301320212301"></a>

#### `rules.egress_rules.ip_prefix_set.ref.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2121212310200012-0200113330230322-1201213330102002-1313000010010013-2223001002013232-3111303220312311-2022123333000230-0003002013201211"></a>

<a id="canonical-3011111322003100-0233301210111220-3020220013111003-0002330002203111-2111021131113002-2011231130103112-2130001311200122-2131003301131131"></a>

#### `rules.egress_rules.ip_prefix_set.ref.uid` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0313303001231300-1213322213333202-1101303032130212-0111233101123313-1111033113303113-2011331133323133-0102020011023111-2103320112031303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.egress_rules.label_matcher` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203)
- rules.egress_rules.label_matcher

<a id="canonical-0120023103213100-3322220320321212-2132130000122030-1112110002310231-1101002221023032-1213333030012313-1103200333232003-3111030321103023"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
label_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121030002310112-2321203213222000-2320003100320330-3122233313231103-0013002032321230-1000002111121220-2302331131022103-0311011203001321"></a>

### Direct properties for `rules.egress_rules.label_matcher`

<a id="canonical-1312331330100221-0223123032010220-0320110332203223-0333000001213022-2212020012203122-1203233213332033-1121210230333111-1331002323202203"></a>

#### `rules.egress_rules.label_matcher.keys` property

Type: `["list", "string"]`. Optional.

The list of label key names that have to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0313200233022122-1011111121221330-2312232210123301-3133103211220033-2131200223302212-0320221220020302-2210313132123030-1030313120102212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.egress_rules.label_selector` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203)
- rules.egress_rules.label_selector

<a id="canonical-0332213110021323-2100000311311003-0001233301222221-1000033211110230-0323300220310123-0102012222113122-3001321332222131-0322200222331012"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
```

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

Terraform syntax:

```terraform
label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130122203213202-3013323103030122-1012322031201011-0203232320323122-1223001130010032-1110303332132312-3311012033113022-2113100033132101"></a>

### Direct properties for `rules.egress_rules.label_selector`

<a id="canonical-0113111320332202-2221333230331230-1023123320222302-0012032001323002-2001100011003301-0103331121123320-1301120012113033-3233301330322320"></a>

#### `rules.egress_rules.label_selector.expressions` property

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0313312122321203-1303132013202013-0110131111110032-3232320002323023-2120320132120313-3312123033220001-2121201303203222-2323303003031231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.egress_rules.metadata` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203)
- rules.egress_rules.metadata

<a id="canonical-1232031010230011-2011013233200021-0111223212330122-2003132000130220-1111023300213222-0000201230112030-0123132210132333-3233101302103323"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2320211331200110-0211101112303320-2013323132233113-0321310203313210-1322211111231132-3303022002021320-0320131020101122-2310120320023100"></a>

### Direct properties for `rules.egress_rules.metadata`

<a id="canonical-3301231122103222-2323010020130101-1320222230312200-2012301100033013-2110000122023032-1223310102211333-2230330232022102-2123323100112210"></a>

#### `rules.egress_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3112232122220113-0022031121100313-2101123121101120-2020122012332120-0032101001330202-2103303200111301-1033333200220232-1023221002010231"></a>

<a id="canonical-1331330132000232-1303200302001121-3001120132201111-1223310031222012-1121220310013300-1113231331133320-3001220103233213-3222323020212321"></a>

#### `rules.egress_rules.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2110323021333030-1213332000232321-2013312202033032-1011202303331301-2100120233103222-0020122011212320-3001100002023303-2032322223102223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.egress_rules.outside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203)
- rules.egress_rules.outside_endpoints

<a id="canonical-3010221002300013-2030312313011331-0102001313020113-3301200321000102-2131333230203232-3112202021111221-0303023331333303-0300131030203322"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
outside_endpoints = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110233120231311-2223011002301312-1320100323032222-3131311313212032-3232103221213103-0000131103222233-1332300102302301-0101323320333100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.egress_rules.prefix_list` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203)
- rules.egress_rules.prefix_list

<a id="canonical-2223223230313301-2300130323000023-3303321032323132-0112332020202221-2020120002321110-0023213213101000-1321201111321133-2121300301310311"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2320123033102220-1313331300031212-1003211133132121-0303230331300120-1120223233001112-2133222123130033-3010010020013213-3002203321000021"></a>

### Direct properties for `rules.egress_rules.prefix_list`

<a id="canonical-0220030332211201-0032233312113033-3130113112211113-1321003230302200-1123000330013103-0032323330133222-3131123000313031-3131313332133300"></a>

#### `rules.egress_rules.prefix_list.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2012213131312232-1113312032121130-0210323122310013-2201130220112023-3013002302203233-3301202123303110-2223102202313111-2323300301101122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.egress_rules.protocol_port_range` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-1320231220302030-1011120023201130-2333312202130003-3033223101012031-3220313233031211-1002322230320332-2221030010200112-1330113202032203)
- rules.egress_rules.protocol_port_range

<a id="canonical-2111133011031032-0032003203302101-0321320013030110-2301312102220022-2000312030222110-2111131033132302-0010001202232330-3021212102233233"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
protocol_port_range {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231313011210001-0123133030311311-2133120322221100-0230132203320310-2030211021202221-1321212031310222-0331333020320102-1033331301122033"></a>

### Direct properties for `rules.egress_rules.protocol_port_range`

<a id="canonical-3120003032233123-3221310221331202-0223121330121232-0230232013223312-0301103110300211-0102202333003020-1330220130132112-1312203102030333"></a>

#### `rules.egress_rules.protocol_port_range.port_ranges` property

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1111103130230011-0131013333201100-3221232120121233-3221021332322230-1121000200133333-1223123333230323-2222201302201013-0100013300202330"></a>

<a id="canonical-3212232101100011-3212031313031000-1210030300310111-2001220302233331-3111001302301332-2132230033303022-0110011121311100-0003131121232102"></a>

#### `rules.egress_rules.protocol_port_range.protocol` property

Type: `"string"`. Optional.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ALL","ICMP","TCP","UDP"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ALL",
    "TCP",
    "UDP",
    "ICMP"),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ingress_rules` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- rules.ingress_rules

<a id="canonical-3110332322122213-0300133202112202-2210212303331212-2132302022002321-3111320030211012-2310031203102101-3231213122023013-0032003303331023"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of rules applied to connections to policy endpoints.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("any",
    "inside_endpoints"),
  validators.ConflictingListObjectAttributes("any",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("any",
    "label_selector"),
  validators.ConflictingListObjectAttributes("any",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("applications",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "label_selector"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "label_selector"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("label_selector",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("label_selector",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("outside_endpoints",
    "prefix_list")}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
ingress_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232213133300133-3021210312202132-2303001322033030-2112130022321002-1132302123233301-3231022200012233-1030011220313021-3122123233023313"></a>

### Direct properties for `rules.ingress_rules`

<a id="canonical-0021320201303023-3111302322323230-1331100012023213-0131101232020301-2310331002103331-1202033320232230-0032131203210301-3110012203332032"></a>

#### `rules.ingress_rules.action` property

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ALLOW","DENY"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW"),
}
```

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

- [adv_action](resources--network_policy--reference--group-001.md#canonical-3011133033230230-0023221201013332-3321030102233121-2003202301310033-3313303233233231-2102221331031201-1302322123130321-1332310001313122): complete subsection reference.

- [all_tcp_traffic](resources--network_policy--reference--group-001.md#canonical-0122003121103201-1130111020320230-2232101203100331-2032213103220320-2230121202313023-3321320123323200-2211211322310122-2103202322112001): complete subsection reference.

- [all_traffic](resources--network_policy--reference--group-001.md#canonical-1301333310103112-2211013211100133-2121100110212111-1231211102113201-0232132211030333-0003132331021033-0220313223311310-2210202002233302): complete subsection reference.

- [all_udp_traffic](resources--network_policy--reference--group-001.md#canonical-1313013113302021-0001010112110332-2333020300322230-1100201212111200-1121223133311210-1110101122200113-2333031230202002-2302321112223211): complete subsection reference.

- [any](resources--network_policy--reference--group-001.md#canonical-1021312211022322-0312033210230202-0130132330333001-3123311010310331-3211133113123211-2310302232222101-1003002311233122-0130011211301033): complete subsection reference.

- [applications](resources--network_policy--reference--group-001.md#canonical-0103013002311201-1111301220102001-2001003022223121-2210020222011312-0101330220102312-2101332000113102-0130210031111213-3212202101221212): complete subsection reference.

- [inside_endpoints](resources--network_policy--reference--group-001.md#canonical-0132102032130200-2112031332310303-3232333332130312-3321203310212133-0211330101321011-1021231210222312-3103302101022022-0220102333302212): complete subsection reference.

- [ip_prefix_set](resources--network_policy--reference--group-001.md#canonical-0023333031132130-1100013001112130-0030201103122211-0022200330023310-1220102022312303-2310032133113000-1333332000200231-2113220132203232): complete subsection reference.

- [label_matcher](resources--network_policy--reference--group-001.md#canonical-1122003323300302-1113001123221231-3322000223332123-0013300200210322-0203330023023021-1133301023122322-1210221320033001-3222301112032320): complete subsection reference.

- [label_selector](resources--network_policy--reference--group-001.md#canonical-1103310133123113-0220113002310210-2100111302200121-3323033312132230-0333122021103202-0020101302003221-2211232120202211-1331030121233110): complete subsection reference.

- [metadata](resources--network_policy--reference--group-001.md#canonical-1201031220031002-1232321102232232-2122230310123211-1211120113023320-3300121230233031-2323130012130301-2031221232310110-2330131301231331): complete subsection reference.

- [outside_endpoints](resources--network_policy--reference--group-001.md#canonical-0103111233320202-0110211203220112-0333203201223122-1111112300213201-3032310200203223-3103001323322010-1121211110022320-0313010023323301): complete subsection reference.

- [prefix_list](resources--network_policy--reference--group-001.md#canonical-1303211102123202-2102300100212202-1303123332001102-3001232222033323-0033110212232002-1203223112110323-1131123030102322-2231200300332003): complete subsection reference.

- [protocol_port_range](resources--network_policy--reference--group-001.md#canonical-3330233332010023-1112220302213302-3012311122301122-2212301222011310-3201221331312022-0133002310031233-2020100003302322-1030123323233320): complete subsection reference.

<a id="canonical-3011133033230230-0023221201013332-3321030102233121-2003202301310033-3313303233233231-2102221331031201-1302322123130321-1332310001313122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ingress_rules.adv_action` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133)
- rules.ingress_rules.adv_action

<a id="canonical-1220013122332021-0213331030221201-3212022022232001-1122320031231130-1013303303320330-1012011102202202-0310110133031123-0131203231230200"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
adv_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132132121131211-2020103111302312-1220123233131223-0030311011312312-0022113220021320-2113000120112120-1223022023201300-1233300031103213"></a>

### Direct properties for `rules.ingress_rules.adv_action`

<a id="canonical-3310211212221032-3310321330001132-1010001332111123-1332112201312221-0113100130332231-3002212333120203-1212332031133001-1022302111010010"></a>

#### `rules.ingress_rules.adv_action.action` property

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["LOG","NOLOG"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("NOLOG",
    "LOG"),
}
```

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

<a id="canonical-0122003121103201-1130111020320230-2232101203100331-2032213103220320-2230121202313023-3321320123323200-2211211322310122-2103202322112001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ingress_rules.all_tcp_traffic` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133)
- rules.ingress_rules.all_tcp_traffic

<a id="canonical-2302103301120130-2303332132211100-2220122300311203-2222200013301200-2013231113310123-3210130123233001-3231303122210302-1121332022010211"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
all_tcp_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301333310103112-2211013211100133-2121100110212111-1231211102113201-0232132211030333-0003132331021033-0220313223311310-2210202002233302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ingress_rules.all_traffic` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133)
- rules.ingress_rules.all_traffic

<a id="canonical-1312111031123211-0330311330122213-1012110301332121-2310011112330032-0010100022212231-1120332213231200-1011213202222203-0303021013213123"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
all_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313013113302021-0001010112110332-2333020300322230-1100201212111200-1121223133311210-1110101122200113-2333031230202002-2302321112223211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ingress_rules.all_udp_traffic` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133)
- rules.ingress_rules.all_udp_traffic

<a id="canonical-3031100133331303-1303003003103303-1012212302122032-2013130001321333-3210030130320302-0322322103012203-2333201303101010-3020001022222123"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
all_udp_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021312211022322-0312033210230202-0130132330333001-3123311010310331-3211133113123211-2310302232222101-1003002311233122-0130011211301033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ingress_rules.any` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133)
- rules.ingress_rules.any

<a id="canonical-2122201312333031-3100010130010322-0103211233103322-2021002011110013-2113012331010103-2101023010111320-2103103220033100-3112321323323000"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
any = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103013002311201-1111301220102001-2001003022223121-2210020222011312-0101330220102312-2101332000113102-0130210031111213-3212202101221212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ingress_rules.applications` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133)
- rules.ingress_rules.applications

<a id="canonical-3313211102033112-2032133221121312-3031121301013310-0010123132021101-2313230022300210-1112301320300020-3110220212300302-0002232033032010"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
applications {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210321120210131-1032202110033131-1013102113113111-1102130233221010-2003122302222133-0302210301020213-0201233111211000-1000121332021322"></a>

### Direct properties for `rules.ingress_rules.applications`

<a id="canonical-3300032130130222-0312332133032200-3332223101102003-1002121023021211-3220002333130122-2100230031333010-1031013100230133-0033030033313000"></a>

#### `rules.ingress_rules.applications.applications` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-0132102032130200-2112031332310303-3232333332130312-3321203310212133-0211330101321011-1021231210222312-3103302101022022-0220102333302212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ingress_rules.inside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133)
- rules.ingress_rules.inside_endpoints

<a id="canonical-0203233132201312-2123333312210201-0101023100322010-0012230321302033-2312210231212123-2000333330110303-3212022200323132-2100332333332023"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
inside_endpoints = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023333031132130-1100013001112130-0030201103122211-0022200330023310-1220102022312303-2310032133113000-1333332000200231-2113220132203232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ingress_rules.ip_prefix_set` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133)
- rules.ingress_rules.ip_prefix_set

<a id="canonical-2222112312031322-3313312203031222-2323222023133201-2101223323110331-2202131022201201-3021211231223112-0110032302123311-0031023311313230"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-0120100202201133-1110121301020300-1103310313101103-1111323131233023-2212322322210201-1121112320303201-3033220020322133-0012303022233220"></a>

### Direct properties for `rules.ingress_rules.ip_prefix_set`

- [ref](resources--network_policy--reference--group-001.md#canonical-3330111322222011-0102210120232113-1131220322222213-3201131103133022-2320300130231223-0031310121011020-0221010231301010-2310213133112222): complete subsection reference.

<a id="canonical-3330111322222011-0102210120232113-1131220322222213-3201131103133022-2320300130231223-0031310121011020-0221010231301010-2310213133112222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ingress_rules.ip_prefix_set.ref` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133)
- [rules.ingress_rules.ip_prefix_set](resources--network_policy--reference--group-001.md#canonical-0023333031132130-1100013001112130-0030201103122211-0022200330023310-1220102022312303-2310032133113000-1333332000200231-2113220132203232)
- rules.ingress_rules.ip_prefix_set.ref

<a id="canonical-3222030110232132-0322220102322012-3130032222032020-2133101130202230-3000011030221020-0130321030313000-1321030323111130-0102211203201321"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-2133311002121200-2131001321330200-0013310222022010-2133133310310312-0000120332102223-0311103122032310-3232230221022110-1232001103033103"></a>

### Direct properties for `rules.ingress_rules.ip_prefix_set.ref`

<a id="canonical-0123233230323131-0031112011003011-3200301032332022-2233000233301301-3012202113123301-3202113231120020-2010132010101110-2223120120321010"></a>

#### `rules.ingress_rules.ip_prefix_set.ref.kind` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3311133102010333-3220323301202210-2330031301330221-2020232031020001-3303011120333320-2131113011310133-0023202012232233-2130203110001133"></a>

<a id="canonical-1322033322213232-1101313310200003-2320330211313330-3222310220001012-3023111222200031-1031332100032231-1333302020023231-1302212002220321"></a>

#### `rules.ingress_rules.ip_prefix_set.ref.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0121130131131333-0132200032322211-3120221301100232-0002022222122112-1122320203222103-3223103310310313-1331212213131103-2213032022113301"></a>

<a id="canonical-2113333302330033-1032121030111320-0101131313331021-1311001122323302-1320121011233111-1223031313013132-3221010031113031-2201003200203301"></a>

#### `rules.ingress_rules.ip_prefix_set.ref.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1211112323132123-1302223303112220-0310323031202031-2012202323201120-1123001101203203-1133330210302321-2230032200312003-1322203020131213"></a>

<a id="canonical-1203002111323113-3300000030331313-3100103210222020-0113201332123131-2330020330020223-0133032031130033-0113133203133131-0111112130032132"></a>

#### `rules.ingress_rules.ip_prefix_set.ref.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0021031122211003-3133023121233032-3203321101032321-0320212003022102-3322133012211312-1113013130103021-1321023230121131-1121300123030111"></a>

<a id="canonical-0100230001132121-1022301032211033-1123200033132313-3312333322330021-1010212133000301-2021113132220112-3220122110002301-2211331033332223"></a>

#### `rules.ingress_rules.ip_prefix_set.ref.uid` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1122003323300302-1113001123221231-3322000223332123-0013300200210322-0203330023023021-1133301023122322-1210221320033001-3222301112032320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ingress_rules.label_matcher` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133)
- rules.ingress_rules.label_matcher

<a id="canonical-1333130130321112-2010123200001331-0120200031030320-2212333001103023-2303301111012331-0212132201000333-1130312133021010-2202302230301012"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
label_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131030323002233-1212232200133203-3010020211031133-1023213322303223-3320111000331021-1101111030223011-1322031022010100-3123211031030013"></a>

### Direct properties for `rules.ingress_rules.label_matcher`

<a id="canonical-2113020322012100-1020011031101100-0231013202020002-1313032312022020-0212310221231101-2320320113021013-1323212032111201-0200113022301133"></a>

#### `rules.ingress_rules.label_matcher.keys` property

Type: `["list", "string"]`. Optional.

The list of label key names that have to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1103310133123113-0220113002310210-2100111302200121-3323033312132230-0333122021103202-0020101302003221-2211232120202211-1331030121233110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ingress_rules.label_selector` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133)
- rules.ingress_rules.label_selector

<a id="canonical-3222022203011122-0123101120332020-2233331101010011-3202211323201302-3123101203300110-3313223232003200-1230131211321020-0313330330013133"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
```

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

Terraform syntax:

```terraform
label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212101132032100-0133022303031000-3013231323100312-3202311033321313-1312113303033313-1330132131010330-1133221323032002-2311020010321333"></a>

### Direct properties for `rules.ingress_rules.label_selector`

<a id="canonical-1300122112113230-1022332310001330-2333133300022132-1302302120303311-2320221331032020-0310011313322001-0210132200321312-1001130212111111"></a>

#### `rules.ingress_rules.label_selector.expressions` property

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1201031220031002-1232321102232232-2122230310123211-1211120113023320-3300121230233031-2323130012130301-2031221232310110-2330131301231331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ingress_rules.metadata` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133)
- rules.ingress_rules.metadata

<a id="canonical-0303332303103102-3001032322121223-2320302023230010-0300100133232131-1020322111121022-2300121101332312-0212000010220003-3103313013031322"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122220103031333-1323311000221300-2131331010010121-3122122312030331-1303302110032222-3303122221130023-0201232112010100-1011301322332301"></a>

### Direct properties for `rules.ingress_rules.metadata`

<a id="canonical-3303101230212330-2322221010332000-3032202030213301-0022022221233212-3231202121121010-1103332331312022-1002121033031300-2233121231212112"></a>

#### `rules.ingress_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1331112110102031-2022203033201100-3123223132220030-1302002011301113-0001233020012120-1131133310121222-2202321021323203-1312013320130320"></a>

<a id="canonical-0112021300113332-2122221101023323-2303003110230222-2210033223300233-1211132102222133-1113333000112033-0202231310323200-2322220321130112"></a>

#### `rules.ingress_rules.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0103111233320202-0110211203220112-0333203201223122-1111112300213201-3032310200203223-3103001323322010-1121211110022320-0313010023323301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ingress_rules.outside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133)
- rules.ingress_rules.outside_endpoints

<a id="canonical-2212230201212211-1211033233210022-3210202103103122-2333103103311120-1321203023020311-3221203230310002-3223201312221321-0123110123103322"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
outside_endpoints = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303211102123202-2102300100212202-1303123332001102-3001232222033323-0033110212232002-1203223112110323-1131123030102322-2231200300332003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ingress_rules.prefix_list` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133)
- rules.ingress_rules.prefix_list

<a id="canonical-2012022101012232-2303001031023333-2000022133211112-3033002122213212-2111113320133232-2030332213012331-3301101103200111-3320202210302330"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111013022322222-1201110032231211-3201203120202232-2030111111220323-3323103011021303-1112331213003300-0301200122131101-0012322221030200"></a>

### Direct properties for `rules.ingress_rules.prefix_list`

<a id="canonical-0112130100210313-0232213113202220-3021200221201123-2032000120313111-3133021303000211-3002111030310303-3212321230112322-0111333230302133"></a>

#### `rules.ingress_rules.prefix_list.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3330233332010023-1112220302213302-3012311122301122-2212301222011310-3201221331312022-0133002310031233-2020100003302322-1030123323233320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.ingress_rules.protocol_port_range` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [rules](resources--network_policy--reference--group-001.md#canonical-2110113303212330-0213122231023011-3312111220103320-1123002101332130-1221320030322012-1021032001120221-2222221120333002-2121113033332301)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0030021231313103-2113212123030202-3223011201100310-1223222321103032-1301000012013112-3212122233002012-3203311002030201-2103213132113133)
- rules.ingress_rules.protocol_port_range

<a id="canonical-0020122023310003-2100330120110231-0103302131121033-3130220021112013-0010333022300322-2001111132310031-2212221330122133-2113331212021331"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
protocol_port_range {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122220000111233-1002333222302230-0130121330233132-3230001012120021-2033300010322333-0201011123113122-1100300233111022-3110022301210230"></a>

### Direct properties for `rules.ingress_rules.protocol_port_range`

<a id="canonical-1321121233203032-3023221101331111-2223000302010230-2000303030202203-0011113132010331-3020220100320212-2311333012101113-0023333303203223"></a>

#### `rules.ingress_rules.protocol_port_range.port_ranges` property

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0111020233113013-2111132312010233-3121123120122232-3132231021220200-2132001121033102-3301230110003323-1130312121021011-0013100200021030"></a>

<a id="canonical-3323000200013331-0203121132310022-1220203332122122-0002002022132221-0012300332112022-0332303301302233-3310232111100020-3301301003000213"></a>

#### `rules.ingress_rules.protocol_port_range.protocol` property

Type: `"string"`. Optional.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ALL","ICMP","TCP","UDP"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ALL",
    "TCP",
    "UDP",
    "ICMP"),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0123011311123100-3002220220320132-2020303012211002-0200230132211012-1010123002001012-1310123220320332-0113122121212101-2123300222322320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- timeouts

<a id="canonical-1131121301101001-0223322023322110-3132001220000313-2303003301030112-0110133222031321-2011203020111003-3013213031211203-0120023110201211"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3220003000303121-2032132221011222-1231220002300323-1233130233331131-1233130113321222-3120201222100211-0321333210311133-3111311033111031"></a>

### Direct properties for `timeouts`

<a id="canonical-3301300023023131-2223230301131313-2211020332000002-0331331203021202-0003213212211322-2211332031010113-1003112022213203-3303113311003030"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0331222113122301-3203323211212032-1333221011331121-0022212032120023-3121212033201212-1021300200112012-0311033100022232-1023033030013210"></a>

<a id="canonical-3210033222303011-0211121310131312-1031131113133202-0101220323120011-0203222132103000-3213000130300031-2222221133110121-1000302233221202"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1321110213331313-0002012201130302-0331103130122112-2032021232102111-3333110213331323-0301031213330202-0233321033121031-2330010020030001"></a>

<a id="canonical-1021023100131020-3032320301312013-3301001222122120-1112123010303111-3232333331103010-1113320110212000-3003230223001201-3122323133233010"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1220210332121332-2220232123232320-0232120331200021-3110113012203333-2021002321221032-2103233012223301-3023020333232310-3022202313113203"></a>

<a id="canonical-3101220312203030-3112232323023100-3032233123302310-0001013121201033-1013213313311303-1303112302320002-1200000211031213-0233030002012220"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
