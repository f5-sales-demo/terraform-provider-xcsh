---
page_title: "xcsh_network_policy_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_rule reference."
---

# xcsh_network_policy_rule reference

<a id="canonical-0111202221110032-0130300212200201-2331112010002030-3130213113210211-3232312202131201-0020002303223310-0331212101231323-1120301311233300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120001220021113-1333003233331300-0013311021213210-0300122221022300-1132111230302202-0310333230320202-0010233313303230-3010001110021322"></a>

## Property reference — Property reference / 231030100212 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103)
- Property reference

<a id="canonical-3001132121032220-1010223100001023-3332030030312223-0112003112331313-0132320302022312-1322013202130203-0023102302333113-3011130021011121"></a>

## Direct properties — Property reference / 231030100212 / 3

<a id="canonical-2111130320210011-3132110012002231-1111201132103002-3022320230010030-1102110010103303-1302300212202131-2203331311033203-1323203323332333"></a>

<a id="canonical-3202212023102030-1133230133012131-0213013212220013-1121220003210221-3203131123333003-2030131231223232-0102122211010023-2100222331303303"></a>

## action property — Property reference / 231030100212 / 4

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

- [advanced_action](data-sources--network_policy_rule--reference--group-001.md#canonical-1030200213123211-2113320222320331-0012331231020202-3203310011003323-2202131023321332-3032000031012111-1321000332222320-0330311313333210): complete subsection reference.

<a id="canonical-3232031130333120-0001033033031023-1212210022210020-0231331031200010-1132333120223203-2131310011331313-2321013013103000-2222032032133330"></a>

<a id="canonical-1202300330322330-0002010211102302-0133220211100130-0112002133321333-0211303012022022-1200031120230203-3110223021200110-2010122021300021"></a>

## annotations property — Property reference / 231030100212 / 5

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

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

<a id="canonical-3203003300011112-0322332323011011-0302102020102133-3113333113323120-1331020132333111-1010310330111120-3300030123201112-1312120323210131"></a>

<a id="canonical-3212330232233311-1312333002303223-1220120301312310-0311122213220021-0133312213110202-0312213222203211-0020101013202030-3100202111300333"></a>

## description property — Property reference / 231030100212 / 6

Type: `"string"`. Computed.

Description of the NetworkPolicyRule.

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

<a id="canonical-2112032321102010-2122133231102121-0010101201331011-2112221302012112-0123212120003100-0011123001002213-0312012023120020-2022013303121313"></a>

<a id="canonical-0213213231010013-0030332033333332-0313003133013000-3232230130000010-0030022203210212-0311122013333113-2202000313203200-0310202013300130"></a>

## ID property — Property reference / 231030100212 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_prefix_set](data-sources--network_policy_rule--reference--group-001.md#canonical-2323013311103120-1113111211332202-0302230120021000-1232111100213302-2012320223231123-3031312301013322-2311112123113030-0120133330022311): complete subsection reference.

- [label_matcher](data-sources--network_policy_rule--reference--group-001.md#canonical-2011312111210001-1011010132212110-0023201030131020-3113102121332313-3012300000321313-0023333132133022-3222220031110122-0201112203212032): complete subsection reference.

<a id="canonical-3022110001320322-1030122331112102-3003223023330020-0111032132031122-1233011213202211-0032213121020331-1332001022131023-1232102230101331"></a>

<a id="canonical-2120311031020231-2230223312212311-1032033023113303-3203012100033121-1332223200330320-2021213021021103-0121201303231111-0100203000333103"></a>

## labels property — Property reference / 231030100212 / 8

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

<a id="canonical-2231330123300333-0211321230012200-2200013302102233-3112223230211313-3202203030302031-0202311323310012-1230203012023122-2201002323121020"></a>

<a id="canonical-2132213113330021-2020021202323001-1212110011231032-2132212033011032-2300000321002113-2023103101032331-3032120332132102-2330320110203101"></a>

## name property — Property reference / 231030100212 / 9

Type: `"string"`. Required.

Name of the NetworkPolicyRule.

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

<a id="canonical-2110302233310012-2010232113112301-1102101132020122-0203132002013213-0010330033323022-0112100112213210-2022011010001331-3222330101021311"></a>

<a id="canonical-3322320131203113-1020013330131211-1223121211332301-2121033202013103-0311300203201031-2213131230020332-0002100123302332-2030231020233121"></a>

## namespace property — Property reference / 231030100212 / 10

Type: `"string"`. Required.

Namespace where the NetworkPolicyRule exists.

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

<a id="canonical-1031320030011111-0303203201031222-1010232030233032-3213021300101201-1203030000322333-2210022210323021-1120232233210202-0330021312221122"></a>

<a id="canonical-3313320113203102-2023333222002013-2031110320230110-2213210220323231-1220333121201120-2300132131002110-2223131100203020-2203330100223103"></a>

## ports property — Property reference / 231030100212 / 11

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [prefix](data-sources--network_policy_rule--reference--group-001.md#canonical-1012230302210131-1012031320322031-1112203010211230-0330130301211000-1100203023100022-3303020011130313-1201120321112232-0021323001131330): complete subsection reference.

- [prefix_selector](data-sources--network_policy_rule--reference--group-001.md#canonical-3103133210210131-3111230202200002-1023002330233210-1022310220213223-0321323033330030-0203020201030121-1031130113032211-0303302132110012): complete subsection reference.

<a id="canonical-0321120122213220-2332311232223000-2222112310303311-3231211023321101-0111321230030021-3320222223103133-2213212210001100-2323011030322211"></a>

<a id="canonical-3110301323323332-1100312021130123-1002021321222212-2002330320220113-1131001112221130-3202310321321021-1001000230121303-2030111331031102"></a>

## protocol property — Property reference / 231030100212 / 12

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1013311300110120-0103111232020230-0201102030203032-1323300311331020-0321300231112112-0332030222101012-2103330103220230-1311110333113103"></a>

## All schema paths — Property reference / 231030100212 / 13

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](data-sources--network_policy_rule--reference--group-001.md#canonical-2111130320210011-3132110012002231-1111201132103002-3022320230010030-1102110010103303-1302300212202131-2203331311033203-1323203323332333) |
| `advanced_action` | [advanced_action](data-sources--network_policy_rule--reference--group-001.md#canonical-3200330121131222-1320101011213001-0003333322130100-2213011133010130-2133001132332222-0012102301123112-0322332122012123-1030332200122333) |
| `advanced_action.action` | [advanced_action.action](data-sources--network_policy_rule--reference--group-001.md#canonical-0110312010223323-3213003002301212-2300132320022213-2033120333330133-3223112203002332-0231203231221013-2131323311003310-3013200113133033) |
| `annotations` | [annotations](data-sources--network_policy_rule--reference--group-001.md#canonical-3232031130333120-0001033033031023-1212210022210020-0231331031200010-1132333120223203-2131310011331313-2321013013103000-2222032032133330) |
| `description` | [description](data-sources--network_policy_rule--reference--group-001.md#canonical-3203003300011112-0322332323011011-0302102020102133-3113333113323120-1331020132333111-1010310330111120-3300030123201112-1312120323210131) |
| `id` | [ID](data-sources--network_policy_rule--reference--group-001.md#canonical-2112032321102010-2122133231102121-0010101201331011-2112221302012112-0123212120003100-0011123001002213-0312012023120020-2022013303121313) |
| `ip_prefix_set` | [ip_prefix_set](data-sources--network_policy_rule--reference--group-001.md#canonical-3200232010221010-3000222320330001-0220022131302323-0123313131312222-1103221123301130-2230301011030101-3010220201031121-3000111011313331) |
| `ip_prefix_set.ref` | [ip_prefix_set.ref](data-sources--network_policy_rule--reference--group-001.md#canonical-0012003201211213-0112103333010121-0021033023300212-1300013121210001-0223201033121010-2220302102212311-0021102331132020-3230213300232322) |
| `ip_prefix_set.ref.kind` | [ip_prefix_set.ref.kind](data-sources--network_policy_rule--reference--group-001.md#canonical-0032200312001122-0212002000031132-0211221333302000-2021101302031301-0031322201000312-2120213013123031-2332230332203022-1101211320110033) |
| `ip_prefix_set.ref.name` | [ip_prefix_set.ref.name](data-sources--network_policy_rule--reference--group-001.md#canonical-3223032322122201-1223212200020111-3320132221323032-2302132113203213-2033000013031200-2020111032201031-1102030221002220-0322103311021323) |
| `ip_prefix_set.ref.namespace` | [ip_prefix_set.ref.namespace](data-sources--network_policy_rule--reference--group-001.md#canonical-0120001030300233-2112032302222320-1330213001300121-1313212000033001-0333310302103011-2201012033110311-0332311031103010-1101020300220311) |
| `ip_prefix_set.ref.tenant` | [ip_prefix_set.ref.tenant](data-sources--network_policy_rule--reference--group-001.md#canonical-3133300200122113-3303223301312013-2012301023132230-3021123323211010-3203322131300131-1202301020032310-1021011112233332-1132322023213112) |
| `ip_prefix_set.ref.uid` | [ip_prefix_set.ref.uid](data-sources--network_policy_rule--reference--group-001.md#canonical-3012103132021230-3131023230102100-1202311301230121-3323122221132332-0210121212312003-0131232223310231-2311211133010033-0210013001231033) |
| `label_matcher` | [label_matcher](data-sources--network_policy_rule--reference--group-001.md#canonical-3100002112001311-3211130101010223-2013311101133202-0000203012111130-0223211332232222-2222030303000201-1202013102121221-2333011011303023) |
| `label_matcher.keys` | [label_matcher.keys](data-sources--network_policy_rule--reference--group-001.md#canonical-1323122211220121-0310220331222311-1230000100130131-0311321202100023-3001001230112330-0203101220011120-1011320130101033-3331031132100201) |
| `labels` | [labels](data-sources--network_policy_rule--reference--group-001.md#canonical-3022110001320322-1030122331112102-3003223023330020-0111032132031122-1233011213202211-0032213121020331-1332001022131023-1232102230101331) |
| `name` | [name](data-sources--network_policy_rule--reference--group-001.md#canonical-2231330123300333-0211321230012200-2200013302102233-3112223230211313-3202203030302031-0202311323310012-1230203012023122-2201002323121020) |
| `namespace` | [namespace](data-sources--network_policy_rule--reference--group-001.md#canonical-2110302233310012-2010232113112301-1102101132020122-0203132002013213-0010330033323022-0112100112213210-2022011010001331-3222330101021311) |
| `ports` | [ports](data-sources--network_policy_rule--reference--group-001.md#canonical-1031320030011111-0303203201031222-1010232030233032-3213021300101201-1203030000322333-2210022210323021-1120232233210202-0330021312221122) |
| `prefix` | [prefix](data-sources--network_policy_rule--reference--group-001.md#canonical-0110312120031203-1131222133222001-3230323302302122-0133223132122223-3112220313221201-2023201030301202-2220003012322231-1033221312310011) |
| `prefix.prefix` | [prefix.prefix](data-sources--network_policy_rule--reference--group-001.md#canonical-1020100132020200-0013333313321220-2331103012330310-3301013333213100-2201012112122023-1332300032111300-0212002131230203-1310313011031001) |
| `prefix_selector` | [prefix_selector](data-sources--network_policy_rule--reference--group-001.md#canonical-1010212332313010-1313332221121110-3013100132031013-0202310000321100-3010013021211131-0330330331322123-0021023013011221-1033323013302100) |
| `prefix_selector.expressions` | [prefix_selector.expressions](data-sources--network_policy_rule--reference--group-001.md#canonical-1133112011031311-0323203300032311-1132303021011111-1223020330032231-2321201301310030-0023212113203133-2102002101130003-1332033313130112) |
| `protocol` | [protocol](data-sources--network_policy_rule--reference--group-001.md#canonical-0321120122213220-2332311232223000-2222112310303311-3231211023321101-0111321230030021-3320222223103133-2213212210001100-2323011030322211) |

<a id="canonical-3032001001121120-0311033233013232-2233133010200301-3320102121011300-2202310102300331-0113230020230302-1223222323320230-3223101123303011"></a>

## Next pages — Property reference / 231030100212 / 14

- [advanced_action](data-sources--network_policy_rule--reference--group-001.md#canonical-1030200213123211-2113320222320331-0012331231020202-3203310011003323-2202131023321332-3032000031012111-1321000332222320-0330311313333210)
- [ip_prefix_set](data-sources--network_policy_rule--reference--group-001.md#canonical-2323013311103120-1113111211332202-0302230120021000-1232111100213302-2012320223231123-3031312301013322-2311112123113030-0120133330022311)
- [label_matcher](data-sources--network_policy_rule--reference--group-001.md#canonical-2011312111210001-1011010132212110-0023201030131020-3113102121332313-3012300000321313-0023333132133022-3222220031110122-0201112203212032)
- [prefix](data-sources--network_policy_rule--reference--group-001.md#canonical-1012230302210131-1012031320322031-1112203010211230-0330130301211000-1100203023100022-3303020011130313-1201120321112232-0021323001131330)
- [prefix_selector](data-sources--network_policy_rule--reference--group-001.md#canonical-3103133210210131-3111230202200002-1023002330233210-1022310220213223-0321323033330030-0203020201030121-1031130113032211-0303302132110012)
- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103)

<a id="canonical-1030200213123211-2113320222320331-0012331231020202-3203310011003323-2202131023321332-3032000031012111-1321000332222320-0330311313333210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033023310010213-1031002202222220-3233021112130210-0020031030023300-1223110010322000-2210300310213230-2122101100102222-1020033130032000"></a>

## advanced_action — advanced_action / 323000002300 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103)
- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-0111202221110032-0130300212200201-2331112010002030-3130213113210211-3232312202131201-0020002303223310-0331212101231323-1120301311233300)
- advanced_action

<a id="canonical-3200330121131222-1320101011213001-0003333322130100-2213011133010130-2133001132332222-0012102301123112-0322332122012123-1030332200122333"></a>

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

<a id="canonical-1023021002300023-3203012133333020-0311030233201130-2322013121303021-2003010111200223-2013033023322131-0211013011021203-1232020302330323"></a>

## Direct properties — advanced_action / 323000002300 / 3

<a id="canonical-0110312010223323-3213003002301212-2300132320022213-2033120333330133-3223112203002332-0231203231221013-2131323311003310-3013200113133033"></a>

<a id="canonical-0002232330111213-1111013130330011-3011221230321020-3300220310201101-3323333120123302-1131332030101112-2100332320220310-3030100210012010"></a>

## action property — advanced_action / 323000002300 / 4

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

<a id="canonical-3133123302013200-2202230213021030-0330331231323022-0233130300112233-1013231200302130-2013033122310030-3110001131221230-2131002301221320"></a>

## Next pages — advanced_action / 323000002300 / 5

- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-0111202221110032-0130300212200201-2331112010002030-3130213113210211-3232312202131201-0020002303223310-0331212101231323-1120301311233300)
- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103)

<a id="canonical-2323013311103120-1113111211332202-0302230120021000-1232111100213302-2012320223231123-3031312301013322-2311112123113030-0120133330022311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132300033331211-0311210333023010-3101103213221113-3213000023200120-3331220232233222-1322210021133001-2322113111313100-3212101113033021"></a>

## ip_prefix_set — ip_prefix_set / 122030130030 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103)
- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-0111202221110032-0130300212200201-2331112010002030-3130213113210211-3232312202131201-0020002303223310-0331212101231323-1120301311233300)
- ip_prefix_set

<a id="canonical-3200232010221010-3000222320330001-0220022131302323-0123313131312222-1103221123301130-2230301011030101-3010220201031121-3000111011313331"></a>

Type: `"single"`. Computed.

\[OneOf: ip\_prefix\_set, prefix, prefix\_selector\] List of references to ip\_prefix\_set objects.

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

OneOf alternatives in this subsection:

- [ip_prefix_set](data-sources--network_policy_rule--reference--group-001.md#canonical-3200232010221010-3000222320330001-0220022131302323-0123313131312222-1103221123301130-2230301011030101-3010220201031121-3000111011313331)
- [prefix](data-sources--network_policy_rule--reference--group-001.md#canonical-0110312120031203-1131222133222001-3230323302302122-0133223132122223-3112220313221201-2023201030301202-2220003012322231-1033221312310011)
- [prefix_selector](data-sources--network_policy_rule--reference--group-001.md#canonical-1010212332313010-1313332221121110-3013100132031013-0202310000321100-3010013021211131-0330330331322123-0021023013011221-1033323013302100)

Select alternatives according to the provider validators above.

<a id="canonical-2102203032223113-3313102113320313-2321212100322323-3121102221031110-1120000121321012-2131123002013033-2120221023001011-3203323003111310"></a>

## Direct properties — ip_prefix_set / 122030130030 / 3

- [ref](data-sources--network_policy_rule--reference--group-001.md#canonical-1110101121332321-1213123122223031-1101113212102002-0130302313101332-3003012033003200-0020320121233203-3232333321333221-3202132010331310): complete subsection reference.

<a id="canonical-2222332331223102-0022200113203122-1132202300110202-0030002000302111-3321132301203032-2001223331122311-1102302131021220-1320213133211110"></a>

## Next pages — ip_prefix_set / 122030130030 / 4

- [ip_prefix_set.ref](data-sources--network_policy_rule--reference--group-001.md#canonical-1110101121332321-1213123122223031-1101113212102002-0130302313101332-3003012033003200-0020320121233203-3232333321333221-3202132010331310)
- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-0111202221110032-0130300212200201-2331112010002030-3130213113210211-3232312202131201-0020002303223310-0331212101231323-1120301311233300)
- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103)

<a id="canonical-1110101121332321-1213123122223031-1101113212102002-0130302313101332-3003012033003200-0020320121233203-3232333321333221-3202132010331310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201121300222311-3000221101302101-0222000030010012-3123332211033331-3001030131301233-0013131033002302-3130322000030210-1003231221311213"></a>

## ip_prefix_set.ref — ref / 001000032130 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103)
- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-0111202221110032-0130300212200201-2331112010002030-3130213113210211-3232312202131201-0020002303223310-0331212101231323-1120301311233300)
- [ip_prefix_set](data-sources--network_policy_rule--reference--group-001.md#canonical-2323013311103120-1113111211332202-0302230120021000-1232111100213302-2012320223231123-3031312301013322-2311112123113030-0120133330022311)
- ip_prefix_set.ref

<a id="canonical-0012003201211213-0112103333010121-0021033023300212-1300013121210001-0223201033121010-2220302102212311-0021102331132020-3230213300232322"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2203203120003003-1130100030232122-1300200002031203-1130110302130200-1302320221132313-1100301101032311-3321102113313330-3000121223202203"></a>

## Direct properties — ref / 001000032130 / 3

<a id="canonical-0032200312001122-0212002000031132-0211221333302000-2021101302031301-0031322201000312-2120213013123031-2332230332203022-1101211320110033"></a>

<a id="canonical-3111320122213021-1211201312020003-0233103220333032-0022331031032312-0002120023310203-2111302302322103-0231220320320030-0133120021130102"></a>

## kind property — ref / 001000032130 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3223032322122201-1223212200020111-3320132221323032-2302132113203213-2033000013031200-2020111032201031-1102030221002220-0322103311021323"></a>

<a id="canonical-1131113131222310-1100123131022013-2203033321103012-0103203130032222-2213202330022033-1111232112330301-2000003111100123-2231212321303030"></a>

## name property — ref / 001000032130 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0120001030300233-2112032302222320-1330213001300121-1313212000033001-0333310302103011-2201012033110311-0332311031103010-1101020300220311"></a>

<a id="canonical-3023032333332230-3311122322322223-3102213103320003-3203013323021131-1312220011213300-2302103331013312-3203030333022133-2220103102210322"></a>

## namespace property — ref / 001000032130 / 6

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

<a id="canonical-3133300200122113-3303223301312013-2012301023132230-3021123323211010-3203322131300131-1202301020032310-1021011112233332-1132322023213112"></a>

<a id="canonical-2013330113221313-1223302323000211-2212013132213300-1113221332321211-3000103100230211-2331001230321330-3031000022221023-1200132113321010"></a>

## tenant property — ref / 001000032130 / 7

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3012103132021230-3131023230102100-1202311301230121-3323122221132332-0210121212312003-0131232223310231-2311211133010033-0210013001231033"></a>

<a id="canonical-1022233023012133-1010111231332102-1101022022003302-3131122230302313-0233232310222001-1130222311320303-1331131330210211-1030131111011211"></a>

## uid property — ref / 001000032130 / 8

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2330033000010121-3123133022221132-1121313200012232-2133320301002023-1130332102313010-0111310213322333-1112132303301133-1230232012321211"></a>

## Next pages — ref / 001000032130 / 9

- [ip_prefix_set](data-sources--network_policy_rule--reference--group-001.md#canonical-2323013311103120-1113111211332202-0302230120021000-1232111100213302-2012320223231123-3031312301013322-2311112123113030-0120133330022311)
- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103)

<a id="canonical-2011312111210001-1011010132212110-0023201030131020-3113102121332313-3012300000321313-0023333132133022-3222220031110122-0201112203212032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131120021310103-1023230211211101-3133022113031030-3121222213002000-1232321112222023-3313001000001311-3131131332200230-3133003203312301"></a>

## label_matcher — label_matcher / 323002202203 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103)
- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-0111202221110032-0130300212200201-2331112010002030-3130213113210211-3232312202131201-0020002303223310-0331212101231323-1120301311233300)
- label_matcher

<a id="canonical-3100002112001311-3211130101010223-2013311101133202-0000203012111130-0223211332232222-2222030303000201-1202013102121221-2333011011303023"></a>

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

<a id="canonical-3101113033300121-2132211130123202-0313323030222332-0103211100222133-1012101300113032-2211211111121001-0002100323130332-0002301100032231"></a>

## Direct properties — label_matcher / 323002202203 / 3

<a id="canonical-1323122211220121-0310220331222311-1230000100130131-0311321202100023-3001001230112330-0203101220011120-1011320130101033-3331031132100201"></a>

<a id="canonical-3111201030110210-3311233220323113-2230220321000303-3303212202021310-1103103300221211-3113032312320221-1112133323230213-0133202020022213"></a>

## keys property — label_matcher / 323002202203 / 4

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

<a id="canonical-2001333031323021-2112310203110331-1230231100330111-0010303221233000-1120200012222031-0001310223112213-0001133132330121-3113212320112011"></a>

## Next pages — label_matcher / 323002202203 / 5

- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-0111202221110032-0130300212200201-2331112010002030-3130213113210211-3232312202131201-0020002303223310-0331212101231323-1120301311233300)
- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103)

<a id="canonical-1012230302210131-1012031320322031-1112203010211230-0330130301211000-1100203023100022-3303020011130313-1201120321112232-0021323001131330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303012130111100-2130331003220322-3320011313133203-2333301133300123-1013300032321111-1221230322302011-2201103330002133-1112122202033112"></a>

## prefix — prefix / 333130303230 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103)
- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-0111202221110032-0130300212200201-2331112010002030-3130213113210211-3232312202131201-0020002303223310-0331212101231323-1120301311233300)
- prefix

<a id="canonical-0110312120031203-1131222133222001-3230323302302122-0133223132122223-3112220313221201-2023201030301202-2220003012322231-1033221312310011"></a>

Type: `"single"`. Computed.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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

<a id="canonical-3210211111320311-0303213021110223-0130120102332030-2011311131103230-0310333112322123-0212301030323333-2100111203113211-1110030101232033"></a>

## Direct properties — prefix / 333130303230 / 3

<a id="canonical-1020100132020200-0013333313321220-2331103012330310-3301013333213100-2201012112122023-1332300032111300-0212002131230203-1310313011031001"></a>

<a id="canonical-2010322322323321-2310201002022221-2030201033011330-2120003121333113-3123003233120013-3131120110311213-3022131322230300-2121102213330221"></a>

## prefix property — prefix / 333130303230 / 4

Type: `["list", "string"]`. Computed.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-2231122012213113-3010301033101130-0130212103111212-0120221013110113-0223110000210120-0112203010213102-0000010212003000-3100311133210300"></a>

## Next pages — prefix / 333130303230 / 5

- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-0111202221110032-0130300212200201-2331112010002030-3130213113210211-3232312202131201-0020002303223310-0331212101231323-1120301311233300)
- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103)

<a id="canonical-3103133210210131-3111230202200002-1023002330233210-1022310220213223-0321323033330030-0203020201030121-1031130113032211-0303302132110012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023322223232322-1000312331100132-1230123200331323-1132023233321033-0123230333133221-3110233011021031-3323333122031110-0320301202202201"></a>

## prefix_selector — prefix_selector / 322102203023 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103)
- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-0111202221110032-0130300212200201-2331112010002030-3130213113210211-3232312202131201-0020002303223310-0331212101231323-1120301311233300)
- prefix_selector

<a id="canonical-1010212332313010-1313332221121110-3013100132031013-0202310000321100-3010013021211131-0330330331322123-0021023013011221-1033323013302100"></a>

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

<a id="canonical-2113123132201313-1222000001112011-3102003300011102-1320201002313332-3312122311012111-3010103010003311-0131122101230012-2322211333330330"></a>

## Direct properties — prefix_selector / 322102203023 / 3

<a id="canonical-1133112011031311-0323203300032311-1132303021011111-1223020330032231-2321201301310030-0023212113203133-2102002101130003-1332033313130112"></a>

<a id="canonical-1133023210123132-3311032100011321-3310010111123313-1230012001131123-1321112220130232-3332012001120103-2302111210101210-3030111022130003"></a>

## expressions property — prefix_selector / 322102203023 / 4

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

<a id="canonical-1201313232232212-1231203003122222-3231112120323122-2300312132232310-2213103212321023-0013101030112203-0033200030230010-3102331322101023"></a>

## Next pages — prefix_selector / 322102203023 / 5

- [Property reference](data-sources--network_policy_rule--reference--group-001.md#canonical-0111202221110032-0130300212200201-2331112010002030-3130213113210211-3232312202131201-0020002303223310-0331212101231323-1120301311233300)
- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103)
