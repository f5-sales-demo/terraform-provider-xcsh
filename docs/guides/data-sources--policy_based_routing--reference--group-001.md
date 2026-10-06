---
page_title: "xcsh_policy_based_routing reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policy_based_routing reference."
---

# xcsh_policy_based_routing reference

<a id="canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- Property reference

<a id="canonical-3030020300331332-3112100233133011-2000122311122311-0001103222322211-1303102103102122-3202332232010032-3300222132003120-0122300012103021"></a>

### Direct properties for `xcsh_policy_based_routing`

<a id="canonical-3112112311032233-2113003320300201-1330030322111232-0010200130002131-3213031300311022-0300332010130131-3131020100003231-0122112101111132"></a>

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

<a id="canonical-2031313220202300-2221312212011230-0002001202110322-1300312113223320-2210032320101133-3232233302303103-1033203112023120-1001020202100231"></a>

<a id="canonical-3031302232231212-1230030023200013-1100212211322120-1031303130002212-3112230020101023-0320311121113301-0303223130230121-2131100232110321"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the PolicyBasedRouting.

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

- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3223233112321212-1021022011013010-3323201332221223-2030103110313232-0332221032132101-0032001220122113-2313122122020010-3321223211310302): complete subsection reference.

- [forwarding_class_list](data-sources--policy_based_routing--reference--group-001.md#canonical-0000313321010012-3122033121021020-3003202200223311-1012313310330211-1020301120121202-0010021123132011-2130133313112003-2201320000010232): complete subsection reference.

<a id="canonical-3112212000310233-2022322202201320-2032113100112212-2313233033301332-3300300111112220-1113000221022302-0303221001303323-0101000131203012"></a>

<a id="canonical-1330331112302232-3303201232013012-2121001220103321-0232233102010230-2102032300111313-0112003031030033-1213212331102131-2323011201010131"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0330313330032031-0113303020332110-1310032122333200-3110022011313021-3211021220300010-2212230323031323-0321022111300221-3212011312230003"></a>

<a id="canonical-1001103321200110-3100003030021212-1233002230022001-1020100120221122-3133112110130310-0102002011331210-3203032300211101-1103221320312013"></a>

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

<a id="canonical-2231031221013132-2100302222220001-0130310333133221-1100003333302202-3330212013023133-2101202331331113-2112332001131212-1200133301012201"></a>

<a id="canonical-3202222332100012-2301020132100333-1223011232311322-2122111101321332-2103203212332211-1332310103102230-1221101120311332-1201103230112100"></a>

#### `name` property

Type: `"string"`. Required.

Name of the PolicyBasedRouting.

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

<a id="canonical-0233103022321233-2112333013313113-1332012100021321-1200133033012001-1112311101333212-0012020222010302-1102221100011311-1301112000213201"></a>

<a id="canonical-1131301030100310-3201303133233023-1030230201012332-2301320211012130-0320113323303310-0101211323101001-0230332303333201-1311123332130112"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the PolicyBasedRouting exists.

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

- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011): complete subsection reference.

<a id="canonical-1323103133231020-1033100000321300-1032322203110031-1131103012212012-2003202003100330-3333311033100230-1110112111111122-0232103103301223"></a>

### All schema paths for `xcsh_policy_based_routing`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--policy_based_routing--reference--group-001.md#canonical-3112112311032233-2113003320300201-1330030322111232-0010200130002131-3213031300311022-0300332010130131-3131020100003231-0122112101111132) |
| `description` | [description](data-sources--policy_based_routing--reference--group-001.md#canonical-2031313220202300-2221312212011230-0002001202110322-1300312113223320-2210032320101133-3232233302303103-1033203112023120-1001020202100231) |
| `forward_proxy_pbr` | [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3221222020111032-3201323030232013-0001022322021330-2220133011233223-0030110210213002-2131123321312012-2003211103300201-0022303213300201) |
| `forward_proxy_pbr.forward_proxy_pbr_rules` | [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-3321000112303200-1211100020211103-0302111232133313-0313312303211231-3102020331101012-1210002002111120-1222030302231322-1201333231231220) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations` | [forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations](data-sources--policy_based_routing--reference--group-001.md#canonical-0030312222220013-3020103222333000-3302300201020202-3030120313313212-3010302320323133-2020102311210222-2221111332332310-2133102002012213) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.all_sources` | [forward_proxy_pbr.forward_proxy_pbr_rules.all_sources](data-sources--policy_based_routing--reference--group-001.md#canonical-3222000123130021-3022000103322213-1333233003010012-0311331022131333-1021232032100122-1003201333232200-2132121313321313-3330023213133033) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list](data-sources--policy_based_routing--reference--group-001.md#canonical-0103111111302122-0321010211301332-3100320311311133-2022213032112302-3120223130303301-2303111102123002-3010300212033000-1333233323221223) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.name](data-sources--policy_based_routing--reference--group-001.md#canonical-3113233211031303-1220230122202021-2312222032212021-2111120333011103-3211021031130020-1221022033331110-0200000330121202-2203113211332020) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.namespace` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.namespace](data-sources--policy_based_routing--reference--group-001.md#canonical-0312332220031233-1310330110120132-1223003103130031-1221201311300222-0303332132122220-2121133131313232-0200012223232113-0322121023113110) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.tenant` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.tenant](data-sources--policy_based_routing--reference--group-001.md#canonical-0323201332022333-1233122210130101-2131033233212002-2320232132303123-3323122331012303-0303323121130011-2300023321123211-1220000221321210) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-0120102310322232-2320222102320011-1001133000010003-0331010013202030-1130333321012220-0030103212111310-3212031311210112-0123133000113300) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-2012112131013310-2103010130013321-2220230102200200-3100320131013123-0333132302212322-2302212020203131-3333123030222320-2211331131222120) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path](data-sources--policy_based_routing--reference--group-001.md#canonical-2033233333200132-1212213320122310-1000300302203212-1111022311213212-0120312023123230-0212312203101221-0013211011313122-2002112210030032) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.exact_value](data-sources--policy_based_routing--reference--group-001.md#canonical-2130230233101101-2312022110332032-2110103031123131-3322213133110301-1300103301210121-0212032122210312-2120112232101001-2023113103301111) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_exact_value](data-sources--policy_based_routing--reference--group-001.md#canonical-3202103213321313-0032210103121202-1113201332211232-3210221110202310-1001302111110203-2120303100203311-0320211120232133-0032131001030001) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_prefix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_prefix_value](data-sources--policy_based_routing--reference--group-001.md#canonical-1332232021302313-1231013132023000-1231213323223331-2003212213322133-1211100020220212-0013332031333232-2203130121011110-0201021312322122) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_regex_value](data-sources--policy_based_routing--reference--group-001.md#canonical-1131110231310032-3113123232102102-2313122131313023-3333003110333120-3203101211311013-1232321300023133-1021200312030030-3030120032223313) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.regex_value](data-sources--policy_based_routing--reference--group-001.md#canonical-0203123230311202-3123333030312321-0032032200100322-1222312000111221-3130303323331133-0010003113303223-2011011311232220-3101303032332130) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.suffix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.suffix_value](data-sources--policy_based_routing--reference--group-001.md#canonical-2031211021123121-3323233022113213-2131003203030010-2201013123000021-2110321102101331-1333222200320113-2100111321000223-0303110100100101) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set](data-sources--policy_based_routing--reference--group-001.md#canonical-3330131102202302-1303000230023323-0221323131322313-1002302000323100-1023233301032331-1111222213103023-1021102010223213-0332310112200101) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.name](data-sources--policy_based_routing--reference--group-001.md#canonical-3101203201012331-0233303302200113-3121003322211030-0201302300110022-1223022312302122-2321212103102100-3000210210300101-2002312200023321) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.namespace` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.namespace](data-sources--policy_based_routing--reference--group-001.md#canonical-2021221233302033-0113030220133122-2000021102310100-0320032102031312-2031303011032221-2200023122310323-3023223211311131-3332033201112303) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.tenant` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.tenant](data-sources--policy_based_routing--reference--group-001.md#canonical-3021202330222010-0011010232020021-3132011120011322-2030111000001002-2231100303223113-3302331012021221-1202012330112233-3323022131113133) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector` | [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector](data-sources--policy_based_routing--reference--group-001.md#canonical-3202301020010233-1201001212232331-1030213011321222-2233011310012210-3121223100130011-2103130001322100-2101232002301233-2331031323203311) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector.expressions` | [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector.expressions](data-sources--policy_based_routing--reference--group-001.md#canonical-1322110303222023-3111321333023333-3112133221011122-3210320210010223-0032332010322003-2122320031200200-0103032322112301-1102311320131123) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata](data-sources--policy_based_routing--reference--group-001.md#canonical-2020023202023303-0320113203112312-0020230112131313-3300220213230010-3012001200011320-3020330113023230-3203102310321320-0130311010130103) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata.description_spec` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata.description_spec](data-sources--policy_based_routing--reference--group-001.md#canonical-3210131001331113-3311121333033131-3232313020233123-2333002002000110-1331202203311020-3230301032130120-0002202120010332-2121102330200112) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata.name](data-sources--policy_based_routing--reference--group-001.md#canonical-0011323201021033-3111030130233121-0220221000001212-2201103212213122-0132103123311322-0211022311003032-1220103300132313-2200323201320230) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list](data-sources--policy_based_routing--reference--group-001.md#canonical-1132301233131202-3223122203320313-2202100311302013-3330332303030301-2330123302203120-2210320213103331-3203300130000330-0310002222222033) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list.prefixes` | [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list.prefixes](data-sources--policy_based_routing--reference--group-001.md#canonical-0322003202212132-3211020113003033-2202330203220012-0302130130333003-1111203103321320-0112332012212001-0101121032213331-3011020321131311) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](data-sources--policy_based_routing--reference--group-001.md#canonical-3330122032010110-0002321321310021-0322330101101322-3120102113300120-0233232022021033-3232103013011200-0310103131132322-2233023102311120) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list](data-sources--policy_based_routing--reference--group-001.md#canonical-3031102333101032-1210113022122103-2222200102313132-2013111233311020-2021322132001113-1331133200323023-0230213231202332-1002320210302120) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.exact_value](data-sources--policy_based_routing--reference--group-001.md#canonical-0230111301330111-0132120010031120-3011222203123312-2222022032211131-0133311220211003-1002311233310110-1123111111003023-0031200033320323) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.regex_value](data-sources--policy_based_routing--reference--group-001.md#canonical-2303322022000333-3221022131102033-1021211132111232-0000320211330022-2311021232120330-3030321311312113-1230111102333110-3023331333332120) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.suffix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.suffix_value](data-sources--policy_based_routing--reference--group-001.md#canonical-0213221011231123-1110231031103230-0330112010230322-0010211031201033-1200311311212213-1103013031013102-3020130332113212-0202302323203312) |
| `forwarding_class_list` | [forwarding_class_list](data-sources--policy_based_routing--reference--group-001.md#canonical-1301021100101323-3200212102002132-1103111003023133-3211313230031323-1022332320303023-3330022330313032-0210203122302310-1323302032113120) |
| `forwarding_class_list.name` | [forwarding_class_list.name](data-sources--policy_based_routing--reference--group-001.md#canonical-3013031330003313-3323323132301321-1313203332211102-1313021111213210-2333201130123330-2120000313210223-1121202203322313-3312030123203102) |
| `forwarding_class_list.namespace` | [forwarding_class_list.namespace](data-sources--policy_based_routing--reference--group-001.md#canonical-3033332323013311-0230012000033021-1002011211211013-2030111130101223-3210221003111121-0111331011100113-0212130133213332-3130003313320012) |
| `forwarding_class_list.tenant` | [forwarding_class_list.tenant](data-sources--policy_based_routing--reference--group-001.md#canonical-1213233100110320-2033333003221011-0322032101200120-3322002333333000-0310332112330231-1000302310033111-0201201031332031-3212112013233220) |
| `id` | [ID](data-sources--policy_based_routing--reference--group-001.md#canonical-3112212000310233-2022322202201320-2032113100112212-2313233033301332-3300300111112220-1113000221022302-0303221001303323-0101000131203012) |
| `labels` | [labels](data-sources--policy_based_routing--reference--group-001.md#canonical-0330313330032031-0113303020332110-1310032122333200-3110022011313021-3211021220300010-2212230323031323-0321022111300221-3212011312230003) |
| `name` | [name](data-sources--policy_based_routing--reference--group-001.md#canonical-2231031221013132-2100302222220001-0130310333133221-1100003333302202-3330212013023133-2101202331331113-2112332001131212-1200133301012201) |
| `namespace` | [namespace](data-sources--policy_based_routing--reference--group-001.md#canonical-0233103022321233-2112333013313113-1332012100021321-1200133033012001-1112311101333212-0012020222010302-1102221100011311-1301112000213201) |
| `network_pbr` | [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3331001030303012-0003001332202031-2110120322211301-3311311331330122-2011011221312002-0111202131230112-3110122333022321-2320312323220003) |
| `network_pbr.any` | [network_pbr.any](data-sources--policy_based_routing--reference--group-001.md#canonical-1003002300100132-0312322232230322-2301320020133101-1322213003110331-2100102012100300-0332022233320133-1133103123320002-2221321010110231) |
| `network_pbr.label_selector` | [network_pbr.label_selector](data-sources--policy_based_routing--reference--group-001.md#canonical-0111122201132010-3303201213011313-2113100003302110-3313233201301123-2100021103321300-2111301000331333-2231200003113123-2223323020101303) |
| `network_pbr.label_selector.expressions` | [network_pbr.label_selector.expressions](data-sources--policy_based_routing--reference--group-001.md#canonical-0301212100320300-1102212331233032-1202300312332203-0100123311200112-1232032221211321-0231000132333111-2120123123222122-3023021030323233) |
| `network_pbr.network_pbr_rules` | [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-2203310300222123-0123002223021120-3120023210223002-3233010220102122-0201110231322303-3000212333031132-2202113222000111-0130023010311310) |
| `network_pbr.network_pbr_rules.all_tcp_traffic` | [network_pbr.network_pbr_rules.all_tcp_traffic](data-sources--policy_based_routing--reference--group-001.md#canonical-2002201111112010-3300303300010122-3033220010210213-0323221113322103-3020030310221300-0310012013022200-0320033112030223-3210022111002132) |
| `network_pbr.network_pbr_rules.all_traffic` | [network_pbr.network_pbr_rules.all_traffic](data-sources--policy_based_routing--reference--group-001.md#canonical-1032230231223000-2232212230112201-0013021130110330-2210313323032110-3201311110300312-0103331310313203-1122333202323003-0023002322212133) |
| `network_pbr.network_pbr_rules.all_udp_traffic` | [network_pbr.network_pbr_rules.all_udp_traffic](data-sources--policy_based_routing--reference--group-001.md#canonical-0311332320012103-2222100300221320-1122210313203301-3003113010210213-0232003112133001-1313020322230221-2121131200322313-3223021222002103) |
| `network_pbr.network_pbr_rules.any` | [network_pbr.network_pbr_rules.any](data-sources--policy_based_routing--reference--group-001.md#canonical-2333202103032000-2330113230011012-3113322123121211-3002032201210230-2002323321120203-1112002101321112-3130012131222300-3133031323303023) |
| `network_pbr.network_pbr_rules.applications` | [network_pbr.network_pbr_rules.applications](data-sources--policy_based_routing--reference--group-001.md#canonical-0113112200120333-2132110202032013-2332122312130133-0310010012201023-3123220131022100-1232112030020320-2002213203033201-3302322021333130) |
| `network_pbr.network_pbr_rules.applications.applications` | [network_pbr.network_pbr_rules.applications.applications](data-sources--policy_based_routing--reference--group-001.md#canonical-3031121012133021-3022123232031231-0123121031321012-1231013131033010-0202130003310201-3030232011013103-2013123130011300-3002010031220131) |
| `network_pbr.network_pbr_rules.dns_name` | [network_pbr.network_pbr_rules.dns_name](data-sources--policy_based_routing--reference--group-001.md#canonical-2201223111010222-2233210200313010-0122312133120112-3321333203122320-3230330023223331-3123113131033101-0230101030112003-0311011121203010) |
| `network_pbr.network_pbr_rules.forwarding_class_list` | [network_pbr.network_pbr_rules.forwarding_class_list](data-sources--policy_based_routing--reference--group-001.md#canonical-0212320032233232-0103312113032133-3230130012330212-2310313100100320-0130303323213103-2200331002123211-3000101230302133-0102333310210123) |
| `network_pbr.network_pbr_rules.forwarding_class_list.name` | [network_pbr.network_pbr_rules.forwarding_class_list.name](data-sources--policy_based_routing--reference--group-001.md#canonical-1002310212001301-0321200133221322-2231231023021221-1120002323333020-2123200110023100-0103323223101101-1210302230121131-0130312031000030) |
| `network_pbr.network_pbr_rules.forwarding_class_list.namespace` | [network_pbr.network_pbr_rules.forwarding_class_list.namespace](data-sources--policy_based_routing--reference--group-001.md#canonical-1201233030323312-3032002201133330-2220302301321013-3020232331303032-0003021020101103-1312030013110112-2020301002122202-0023200132010312) |
| `network_pbr.network_pbr_rules.forwarding_class_list.tenant` | [network_pbr.network_pbr_rules.forwarding_class_list.tenant](data-sources--policy_based_routing--reference--group-001.md#canonical-3301223003100103-2010030231002303-1010220111113300-3223302333301223-1031203230302111-3211003112311332-1012112130220023-0023131002022303) |
| `network_pbr.network_pbr_rules.ip_prefix_set` | [network_pbr.network_pbr_rules.ip_prefix_set](data-sources--policy_based_routing--reference--group-001.md#canonical-1322312020212233-1032130331003333-0221132330320112-0003113333212021-2100113330313012-3300333312333112-3212230231100010-0130132031010003) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref` | [network_pbr.network_pbr_rules.ip_prefix_set.ref](data-sources--policy_based_routing--reference--group-001.md#canonical-2111101312230220-0330213323023330-2331323002211020-1333211310013133-3213223112203210-2330101232010122-2310102013101311-0132332203313111) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.kind` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.kind](data-sources--policy_based_routing--reference--group-001.md#canonical-2101023031100013-0101212322200232-2321133211130113-0210131113001301-0231021200033000-1300130133012101-3031002100010220-0322003021213222) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.name` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.name](data-sources--policy_based_routing--reference--group-001.md#canonical-2132213231302311-0011023231312011-0123013302131022-3313030030000223-2221102312113000-2311111012020211-3220100030311021-1212113211333330) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.namespace` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.namespace](data-sources--policy_based_routing--reference--group-001.md#canonical-2021021213020110-1021300121303332-2131331000020021-2213310032130302-0321311233023102-1030221001123221-2213123210033103-3011003222111310) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.tenant` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.tenant](data-sources--policy_based_routing--reference--group-001.md#canonical-1212003033230213-1323110220231310-2111121102000023-1300123212133230-0331123102331110-3212211313010212-3101033201232002-3221201032231223) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.uid` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.uid](data-sources--policy_based_routing--reference--group-001.md#canonical-1133222212130211-3210003133000220-1033300222110233-0303002130003100-1032123101321021-0100100032212023-1112303100202031-0210112300333303) |
| `network_pbr.network_pbr_rules.metadata` | [network_pbr.network_pbr_rules.metadata](data-sources--policy_based_routing--reference--group-001.md#canonical-0313332130213011-2311301131301200-0201232103110212-0100233031332202-3132323001332112-1323322333300121-0003303031321202-2222033323333123) |
| `network_pbr.network_pbr_rules.metadata.description_spec` | [network_pbr.network_pbr_rules.metadata.description_spec](data-sources--policy_based_routing--reference--group-001.md#canonical-2222022021031032-0203231020330100-3102012000112300-0311103120320112-3103023221332031-0222230112003203-3333021013330223-3030102333303003) |
| `network_pbr.network_pbr_rules.metadata.name` | [network_pbr.network_pbr_rules.metadata.name](data-sources--policy_based_routing--reference--group-001.md#canonical-2000111021021330-2102323123203222-0201102133131112-3231001223032221-2010123010120231-2230201000110000-0201003032220013-3311121233012031) |
| `network_pbr.network_pbr_rules.prefix_list` | [network_pbr.network_pbr_rules.prefix_list](data-sources--policy_based_routing--reference--group-001.md#canonical-1000200332003223-1232001202333132-1332203121100302-3221103211200202-1002221021212110-0302001113103132-1111103223323021-3032030012210130) |
| `network_pbr.network_pbr_rules.prefix_list.prefixes` | [network_pbr.network_pbr_rules.prefix_list.prefixes](data-sources--policy_based_routing--reference--group-001.md#canonical-1113210021230001-2230132103013021-2203132002333311-2311012201013132-2031231310331202-0122330103002230-1222123121312220-2221221032320012) |
| `network_pbr.network_pbr_rules.protocol_port_range` | [network_pbr.network_pbr_rules.protocol_port_range](data-sources--policy_based_routing--reference--group-001.md#canonical-0303020020330122-1213123023321203-0320101013012131-2222102302001221-0132010231333202-3333003301211303-2331011231212101-0023133211313110) |
| `network_pbr.network_pbr_rules.protocol_port_range.port_ranges` | [network_pbr.network_pbr_rules.protocol_port_range.port_ranges](data-sources--policy_based_routing--reference--group-001.md#canonical-2320311100312331-1312132000310302-2301013101023212-0202103123001022-1302302010132300-2322200320021112-2321031303133030-1130202100123213) |
| `network_pbr.network_pbr_rules.protocol_port_range.protocol` | [network_pbr.network_pbr_rules.protocol_port_range.protocol](data-sources--policy_based_routing--reference--group-001.md#canonical-3000123220230320-2022002230113321-1133003120211132-3330012021032212-0130302231211220-3030002113131332-2303220131230332-2202001020132101) |
| `network_pbr.prefix_list` | [network_pbr.prefix_list](data-sources--policy_based_routing--reference--group-001.md#canonical-2031120311002120-3223203222031001-3113232201313233-2013200123212120-0230112203312300-1333330110201302-1232220010313212-0202132021120220) |
| `network_pbr.prefix_list.prefixes` | [network_pbr.prefix_list.prefixes](data-sources--policy_based_routing--reference--group-001.md#canonical-2133022100213211-1000311303021011-0011330310201233-3203001223300200-1222133331110003-1121202112320300-0021120302002113-0100202232103202) |

<a id="canonical-3223233112321212-1021022011013010-3323201332221223-2030103110313232-0332221032132101-0032001220122113-2313122122020010-3321223211310302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- forward_proxy_pbr

<a id="canonical-3221222020111032-3201323030232013-0001022322021330-2220133011233223-0030110210213002-2131123321312012-2003211103300201-0022303213300201"></a>

Type: `"single"`. Computed.

\[OneOf: forward\_proxy\_pbr, network\_pbr\] Configuration parameter for forward proxy pbr.

Additional upstream details:

Network(L3/L4) routing policy rule.

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

- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3221222020111032-3201323030232013-0001022322021330-2220133011233223-0030110210213002-2131123321312012-2003211103300201-0022303213300201)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3331001030303012-0003001332202031-2110120322211301-3311311331330122-2011011221312002-0111202131230112-3110122333022321-2320312323220003)

Select alternatives according to the provider validators above.

<a id="canonical-0300333313222110-2132333302122110-1012203123011113-0203323210302012-0332210230330001-2222310013113222-1130323331031112-1301103202033102"></a>

### Direct properties for `forward_proxy_pbr`

- [forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-2033231232023322-1123220021132001-1113001322220122-3322102000322103-3113211023303333-3022331000233322-1230123031120013-0020212321011122): complete subsection reference.

<a id="canonical-2033231232023322-1123220021132001-1113001322220122-3322102000322103-3113211023303333-3022331000233322-1230123031120013-0020212321011122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3223233112321212-1021022011013010-3323201332221223-2030103110313232-0332221032132101-0032001220122113-2313122122020010-3321223211310302)
- forward_proxy_pbr.forward_proxy_pbr_rules

<a id="canonical-3321000112303200-1211100020211103-0302111232133313-0313312303211231-3102020331101012-1210002002111120-1222030302231322-1201333231231220"></a>

Type: `"list"`. Computed.

L3/L4 routing rules. Network(L3/L4) routing policy rules.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-0030021231110323-3022232321030110-3233200312322303-3300012201100320-1121110212100211-2203320032020203-0020200110103322-3031230302323130"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules`

- [all_destinations](data-sources--policy_based_routing--reference--group-001.md#canonical-1201130123230132-2331331213231103-2121332003113021-0122030122313311-2202232322011113-1001212010111301-3101113122120320-2200100222120213): complete subsection reference.

- [all_sources](data-sources--policy_based_routing--reference--group-001.md#canonical-3200200332230033-1102022201300300-3303032103312000-0310022103202011-2023233333310222-3103201210123332-2230033312101303-1231132301301222): complete subsection reference.

- [forwarding_class_list](data-sources--policy_based_routing--reference--group-001.md#canonical-1302020001301312-0222223033012131-1120121310032012-1300131022330110-1121012100112031-1000221300223122-2213300322211021-1022002122123222): complete subsection reference.

- [http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-1323233330201000-1123120330313023-2232010013232112-0133020112302123-0333100330230222-0111103031022313-2310300210200300-0003022222231013): complete subsection reference.

- [ip_prefix_set](data-sources--policy_based_routing--reference--group-001.md#canonical-0230201101301320-1013303212311213-1010202310302221-1213223232321200-3210332221211020-3231200221203111-0231230123201131-3110010031102221): complete subsection reference.

- [label_selector](data-sources--policy_based_routing--reference--group-001.md#canonical-3303021203132210-2011310232100120-2331201113033331-0020311331120101-1222201221322030-2132230231011122-3110031321023110-2220000311300312): complete subsection reference.

- [metadata](data-sources--policy_based_routing--reference--group-001.md#canonical-1220321013131012-0021022112023111-3032000132200103-3101312300230022-1131211103133102-2111332302230001-1013312222010331-3103311231223201): complete subsection reference.

- [prefix_list](data-sources--policy_based_routing--reference--group-001.md#canonical-0012110202312313-3331231321212310-0123201110332020-0201121021220031-1101032110123113-3303130232010133-3111212110120332-2020000310112111): complete subsection reference.

- [tls_list](data-sources--policy_based_routing--reference--group-001.md#canonical-1100123132223230-2111231033322130-3110112110301010-3111032233331122-2303101101303011-2323021020130212-0231110110132033-2030132133131220): complete subsection reference.

<a id="canonical-1201130123230132-2331331213231103-2121332003113021-0122030122313311-2202232322011113-1001212010111301-3101113122120320-2200100222120213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3223233112321212-1021022011013010-3323201332221223-2030103110313232-0332221032132101-0032001220122113-2313122122020010-3321223211310302)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-2033231232023322-1123220021132001-1113001322220122-3322102000322103-3113211023303333-3022331000233322-1230123031120013-0020212321011122)
- forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations

<a id="canonical-0030312222220013-3020103222333000-3302300201020202-3030120313313212-3010302320323133-2020102311210222-2221111332332310-2133102002012213"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all destinations.

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

<a id="canonical-3200200332230033-1102022201300300-3303032103312000-0310022103202011-2023233333310222-3103201210123332-2230033312101303-1231132301301222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.all_sources` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3223233112321212-1021022011013010-3323201332221223-2030103110313232-0332221032132101-0032001220122113-2313122122020010-3321223211310302)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-2033231232023322-1123220021132001-1113001322220122-3322102000322103-3113211023303333-3022331000233322-1230123031120013-0020212321011122)
- forward_proxy_pbr.forward_proxy_pbr_rules.all_sources

<a id="canonical-3222000123130021-3022000103322213-1333233003010012-0311331022131333-1021232032100122-1003201333232200-2132121313321313-3330023213133033"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all sources.

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

<a id="canonical-1302020001301312-0222223033012131-1120121310032012-1300131022330110-1121012100112031-1000221300223122-2213300322211021-1022002122123222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3223233112321212-1021022011013010-3323201332221223-2030103110313232-0332221032132101-0032001220122113-2313122122020010-3321223211310302)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-2033231232023322-1123220021132001-1113001322220122-3322102000322103-3113211023303333-3022331000233322-1230123031120013-0020212321011122)
- forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list

<a id="canonical-0103111111302122-0321010211301332-3100320311311133-2022213032112302-3120223130303301-2303111102123002-3010300212033000-1333233323221223"></a>

Type: `"list"`. Computed.

Ordered list of forwarding Class to be used if no rule match.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2133211201002013-0310023210130020-2321011133313312-1102232203321121-1033133230200133-3012122223301011-2233122313330203-1031110231020003"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list`

<a id="canonical-3113233211031303-1220230122202021-2312222032212021-2111120333011103-3211021031130020-1221022033331110-0200000330121202-2203113211332020"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0312332220031233-1310330110120132-1223003103130031-1221201311300222-0303332132122220-2121133131313232-0200012223232113-0322121023113110"></a>

<a id="canonical-1030112220120112-0313323200102032-1033012130101022-2131100101213223-3032331211123332-1312031221230002-1230001323330121-0211102303010003"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0323201332022333-1233122210130101-2131033233212002-2320232132303123-3323122331012303-0303323121130011-2300023321123211-1220000221321210"></a>

<a id="canonical-2001021321113033-2030210000133212-2221310233300201-1211011011002131-3230332003130300-1002033302213101-2313210230311123-0110310110311223"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1323233330201000-1123120330313023-2232010013232112-0133020112302123-0333100330230222-0111103031022313-2310300210200300-0003022222231013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.http_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3223233112321212-1021022011013010-3323201332221223-2030103110313232-0332221032132101-0032001220122113-2313122122020010-3321223211310302)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-2033231232023322-1123220021132001-1113001322220122-3322102000322103-3113211023303333-3022331000233322-1230123031120013-0020212321011122)
- forward_proxy_pbr.forward_proxy_pbr_rules.http_list

<a id="canonical-0120102310322232-2320222102320011-1001133000010003-0331010013202030-1130333321012220-0030103212111310-3212031311210112-0123133000113300"></a>

Type: `"single"`. Computed.

URLListType.

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

<a id="canonical-2333210223333220-0120033010000101-2132233220122331-0320133110111300-3133230312121333-3012001103103202-1020123113302332-0002112021103020"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.http_list`

- [http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-1123032221210301-3302230211300103-2111023222211022-2301111133303113-3312110221032023-1313200212123033-0223330102133021-2131233122113020): complete subsection reference.

<a id="canonical-1123032221210301-3302230211300103-2111023222211022-2301111133303113-3312110221032023-1313200212123033-0223330102133021-2131233122113020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3223233112321212-1021022011013010-3323201332221223-2030103110313232-0332221032132101-0032001220122113-2313122122020010-3321223211310302)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-2033231232023322-1123220021132001-1113001322220122-3322102000322103-3113211023303333-3022331000233322-1230123031120013-0020212321011122)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-1323233330201000-1123120330313023-2232010013232112-0133020112302123-0333100330230222-0111103031022313-2310300210200300-0003022222231013)
- forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list

<a id="canonical-2012112131013310-2103010130013321-2220230102200200-3100320131013123-0333132302212322-2302212020203131-3333123030222320-2211331131222120"></a>

Type: `"list"`. Computed.

HTTP URLs. URLs for HTTP connections.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1101331001321133-0130033100232033-3232210230001212-2011213103100111-1320222323033210-1213030310330013-2103332022021333-2132330131310013"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list`

- [any_path](data-sources--policy_based_routing--reference--group-001.md#canonical-3331120323021312-2122230121231211-1332313233031110-1010210100013030-0222203211123311-1333220030230222-0202120201221312-0212032111131330): complete subsection reference.

<a id="canonical-2130230233101101-2312022110332032-2110103031123131-3322213133110301-1300103301210121-0212032122210312-2120112232101001-2023113103301111"></a>

<a id="canonical-2000100300221030-1020212020212221-3310333021331102-3023022032232122-0110113032211301-3030220011023001-0103010301222300-0201113321321112"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3202103213321313-0032210103121202-1113201332211232-3210221110202310-1001302111110203-2120303100203311-0320211120232133-0032131001030001"></a>

<a id="canonical-1333102010230201-2300300313021202-2013311130220303-3121203311303310-0001331303113213-1213132121011302-0232332222112312-1101200213302231"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_exact_value` property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1332232021302313-1231013132023000-1231213323223331-2003212213322133-1211100020220212-0013332031333232-2203130121011110-0201021312322122"></a>

<a id="canonical-1222323111323132-1223232100101033-3033332020332300-0203020001033100-2122010033130300-0230232203323032-2002300033031310-2203200313112130"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_prefix_value` property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Additional upstream details:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1131110231310032-3113123232102102-2313122131313023-3333003110333120-3203101211311013-1232321300023133-1021200312030030-3030120032223313"></a>

<a id="canonical-0300201321213322-1221313130121033-3031122301231123-1022002120310130-0223203123031112-2011101200011223-3100221111201331-3013230313313301"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_regex_value` property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0203123230311202-3123333030312321-0032032200100322-1222312000111221-3130303323331133-0010003113303223-2011011311232220-3101303032332130"></a>

<a id="canonical-2033003002221303-0323320020210230-3232012332131233-2312213122221011-3011223231300021-1211001122320113-2122113212220301-2100020111013000"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.regex_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2031211021123121-3323233022113213-2131003203030010-2201013123000021-2110321102101331-1333222200320113-2100111321000223-0303110100100101"></a>

<a id="canonical-1331002333001110-1322033020020222-1001031102331231-3103121212132103-3230130221322201-2211201102131030-2301222323323111-2311130112323121"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3331120323021312-2122230121231211-1332313233031110-1010210100013030-0222203211123311-1333220030230222-0202120201221312-0212032111131330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3223233112321212-1021022011013010-3323201332221223-2030103110313232-0332221032132101-0032001220122113-2313122122020010-3321223211310302)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-2033231232023322-1123220021132001-1113001322220122-3322102000322103-3113211023303333-3022331000233322-1230123031120013-0020212321011122)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-1323233330201000-1123120330313023-2232010013232112-0133020112302123-0333100330230222-0111103031022313-2310300210200300-0003022222231013)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list](data-sources--policy_based_routing--reference--group-001.md#canonical-1123032221210301-3302230211300103-2111023222211022-2301111133303113-3312110221032023-1313200212123033-0223330102133021-2131233122113020)
- forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path

<a id="canonical-2033233333200132-1212213320122310-1000300302203212-1111022311213212-0120312023123230-0212312203101221-0013211011313122-2002112210030032"></a>

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

<a id="canonical-0230201101301320-1013303212311213-1010202310302221-1213223232321200-3210332221211020-3231200221203111-0231230123201131-3110010031102221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3223233112321212-1021022011013010-3323201332221223-2030103110313232-0332221032132101-0032001220122113-2313122122020010-3321223211310302)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-2033231232023322-1123220021132001-1113001322220122-3322102000322103-3113211023303333-3022331000233322-1230123031120013-0020212321011122)
- forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set

<a id="canonical-3330131102202302-1303000230023323-0221323131322313-1002302000323100-1023233301032331-1111222213103023-1021102010223213-0332310112200101"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2331033230122023-2012220101103220-3022010012222303-3211012011331010-1013233332301120-1301320212130000-3021203201102023-0013200231223111"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set`

<a id="canonical-3101203201012331-0233303302200113-3121003322211030-0201302300110022-1223022312302122-2321212103102100-3000210210300101-2002312200023321"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2021221233302033-0113030220133122-2000021102310100-0320032102031312-2031303011032221-2200023122310323-3023223211311131-3332033201112303"></a>

<a id="canonical-0100321033132212-2031113200032223-3003131110130220-2303201101312001-3231310311113230-0230132030132330-3133111221123022-3330221321033011"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3021202330222010-0011010232020021-3132011120011322-2030111000001002-2231100303223113-3302331012021221-1202012330112233-3323022131113133"></a>

<a id="canonical-2200220011212200-1330001113323302-0133211122330221-3312131303003332-0311332233121000-2221122212130203-3133120113032010-1333102310012212"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3303021203132210-2011310232100120-2331201113033331-0020311331120101-1222201221322030-2132230231011122-3110031321023110-2220000311300312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3223233112321212-1021022011013010-3323201332221223-2030103110313232-0332221032132101-0032001220122113-2313122122020010-3321223211310302)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-2033231232023322-1123220021132001-1113001322220122-3322102000322103-3113211023303333-3022331000233322-1230123031120013-0020212321011122)
- forward_proxy_pbr.forward_proxy_pbr_rules.label_selector

<a id="canonical-3202301020010233-1201001212232331-1030213011321222-2233011310012210-3121223100130011-2103130001322100-2101232002301233-2331031323203311"></a>

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

<a id="canonical-1001303210111122-2212022313113101-1320113201002211-2001303300320033-0102302013101012-1112121233012121-1122031211331312-3032332032133220"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector`

<a id="canonical-1322110303222023-3111321333023333-3112133221011122-3210320210010223-0032332010322003-2122320031200200-0103032322112301-1102311320131123"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector.expressions` property

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

<a id="canonical-1220321013131012-0021022112023111-3032000132200103-3101312300230022-1131211103133102-2111332302230001-1013312222010331-3103311231223201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.metadata` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3223233112321212-1021022011013010-3323201332221223-2030103110313232-0332221032132101-0032001220122113-2313122122020010-3321223211310302)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-2033231232023322-1123220021132001-1113001322220122-3322102000322103-3113211023303333-3022331000233322-1230123031120013-0020212321011122)
- forward_proxy_pbr.forward_proxy_pbr_rules.metadata

<a id="canonical-2020023202023303-0320113203112312-0020230112131313-3300220213230010-3012001200011320-3020330113023230-3203102310321320-0130311010130103"></a>

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

<a id="canonical-3222310210000223-0330322030101120-0032231002021220-1203110233100302-0130123332310332-2221132121013130-3332120032333012-3123233210232222"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.metadata`

<a id="canonical-3210131001331113-3311121333033131-3232313020233123-2333002002000110-1331202203311020-3230301032130120-0002202120010332-2121102330200112"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0011323201021033-3111030130233121-0220221000001212-2201103212213122-0132103123311322-0211022311003032-1220103300132313-2200323201320230"></a>

<a id="canonical-1221011122303231-3230212113322003-2123002032123231-1301023001312313-0002101323000211-3311233210220210-2022130110302211-3322212210203132"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.metadata.name` property

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

<a id="canonical-0012110202312313-3331231321212310-0123201110332020-0201121021220031-1101032110123113-3303130232010133-3111212110120332-2020000310112111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3223233112321212-1021022011013010-3323201332221223-2030103110313232-0332221032132101-0032001220122113-2313122122020010-3321223211310302)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-2033231232023322-1123220021132001-1113001322220122-3322102000322103-3113211023303333-3022331000233322-1230123031120013-0020212321011122)
- forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list

<a id="canonical-1132301233131202-3223122203320313-2202100311302013-3330332303030301-2330123302203120-2210320213103331-3203300130000330-0310002222222033"></a>

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

<a id="canonical-3010010022013301-1013012330100231-2313200203131003-0103313032321023-2301120301220231-0231111013130100-0333011233320031-1013302103112113"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list`

<a id="canonical-0322003202212132-3211020113003033-2202330203220012-0302130130333003-1111203103321320-0112332012212001-0101121032213331-3011020321131311"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list.prefixes` property

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

<a id="canonical-1100123132223230-2111231033322130-3110112110301010-3111032233331122-2303101101303011-2323021020130212-0231110110132033-2030132133131220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3223233112321212-1021022011013010-3323201332221223-2030103110313232-0332221032132101-0032001220122113-2313122122020010-3321223211310302)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-2033231232023322-1123220021132001-1113001322220122-3322102000322103-3113211023303333-3022331000233322-1230123031120013-0020212321011122)
- forward_proxy_pbr.forward_proxy_pbr_rules.tls_list

<a id="canonical-3330122032010110-0002321321310021-0322330101101322-3120102113300120-0233232022021033-3232103013011200-0310103131132322-2233023102311120"></a>

Type: `"single"`. Computed.

DomainListType.

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

<a id="canonical-3010102333122023-3003111323211031-0320200113123012-0312331211220211-1020000321332201-1232002201001103-3330230031113203-2001333220112233"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list`

- [tls_list](data-sources--policy_based_routing--reference--group-001.md#canonical-3312033210132312-1220122223101220-3300122022320102-2333033110011030-2211113332133220-0133003033301223-3210020200120210-0230010201321303): complete subsection reference.

<a id="canonical-3312033210132312-1220122223101220-3300122022320102-2333033110011030-2211113332133220-0133003033301223-3210020200120210-0230010201321303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [forward_proxy_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-3223233112321212-1021022011013010-3323201332221223-2030103110313232-0332221032132101-0032001220122113-2313122122020010-3321223211310302)
- [forward_proxy_pbr.forward_proxy_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-2033231232023322-1123220021132001-1113001322220122-3322102000322103-3113211023303333-3022331000233322-1230123031120013-0020212321011122)
- [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](data-sources--policy_based_routing--reference--group-001.md#canonical-1100123132223230-2111231033322130-3110112110301010-3111032233331122-2303101101303011-2323021020130212-0231110110132033-2030132133131220)
- forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list

<a id="canonical-3031102333101032-1210113022122103-2222200102313132-2013111233311020-2021322132001113-1331133200323023-0230213231202332-1002320210302120"></a>

Type: `"list"`. Computed.

TLS Domains. Domains in SNI for TLS connections.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1312231333233002-2230000020310022-1132313022022221-3011020321133122-3100230300120233-1002133211210320-3110302133123102-0102232023113232"></a>

### Direct properties for `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list`

<a id="canonical-0230111301330111-0132120010031120-3011222203123312-2222022032211131-0133311220211003-1002311233310110-1123111111003023-0031200033320323"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2303322022000333-3221022131102033-1021211132111232-0000320211330022-2311021232120330-3030321311312113-1230111102333110-3023331333332120"></a>

<a id="canonical-3122110230211022-1001320023333011-3021210133232230-3021021030313220-0123003210330322-3200132031111300-1023310012131111-1213323102213311"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.regex_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0213221011231123-1110231031103230-0330112010230322-0010211031201033-1200311311212213-1103013031013102-3020130332113212-0202302323203312"></a>

<a id="canonical-1120133321020211-1312322120311231-1310132220203000-0000001011210123-2233012022331303-1333213312310203-0210313000320302-3112121131102200"></a>

#### `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0000313321010012-3122033121021020-3003202200223311-1012313310330211-1020301120121202-0010021123132011-2130133313112003-2201320000010232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `forwarding_class_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- forwarding_class_list

<a id="canonical-1301021100101323-3200212102002132-1103111003023133-3211313230031323-1022332320303023-3330022330313032-0210203122302310-1323302032113120"></a>

Type: `"list"`. Computed.

Ordered list of forwarding Class to be used if source application match and no rule match.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-1322222211230322-2222033312100232-0212100002020102-1211321033100103-0133332012023111-3131111303123232-1300110310112033-1320010133113132"></a>

### Direct properties for `forwarding_class_list`

<a id="canonical-3013031330003313-3323323132301321-1313203332211102-1313021111213210-2333201130123330-2120000313210223-1121202203322313-3312030123203102"></a>

#### `forwarding_class_list.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3033332323013311-0230012000033021-1002011211211013-2030111130101223-3210221003111121-0111331011100113-0212130133213332-3130003313320012"></a>

<a id="canonical-2011012332120311-1330203133012210-2202220121303101-2003310213003211-3231202212120201-1321230213323223-1101031212310023-1032033321101003"></a>

#### `forwarding_class_list.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1213233100110320-2033333003221011-0322032101200120-3322002333333000-0310332112330231-1000302310033111-0201201031332031-3212112013233220"></a>

<a id="canonical-3133130233002220-0103020102220320-2020331203333213-0222232011303323-0300210003321231-0101202122101130-2331310222132003-0112103013103001"></a>

#### `forwarding_class_list.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- network_pbr

<a id="canonical-3331001030303012-0003001332202031-2110120322211301-3311311331330122-2011011221312002-0111202131230112-3110122333022321-2320312323220003"></a>

Type: `"single"`. Computed.

Configuration parameter for network pbr.

Additional upstream details:

Network(L3/L4) routing policy rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-source_choice": "[\"any\",\"label_selector\",\"prefix_list\"]"
}
```

<a id="canonical-3322021013333002-3133233302110121-2112131323202223-3110100030231112-2200311201021211-2200121122210313-0102101322221332-0022001010033003"></a>

### Direct properties for `network_pbr`

- [any](data-sources--policy_based_routing--reference--group-001.md#canonical-3013300330132021-3323013123323013-1203200002111203-1331001112132111-0213331101233011-0231321101312202-3213012113020230-3330122302332303): complete subsection reference.

- [label_selector](data-sources--policy_based_routing--reference--group-001.md#canonical-0122322112001130-0213130203210221-0302000123333310-3230311012002213-1031132132102010-1233312020012200-1233223233103103-0320303130031021): complete subsection reference.

- [network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-0112333310121320-1011211132322313-2120120101012110-3101011022221203-2221023311032232-2122212210313003-1022010133030330-2320323000210001): complete subsection reference.

- [prefix_list](data-sources--policy_based_routing--reference--group-001.md#canonical-0232121212310000-1320100010331121-1320312232123020-2202103002223032-3002001120033200-1112310321101002-0023031200202202-3313213332021310): complete subsection reference.

<a id="canonical-3013300330132021-3323013123323013-1203200002111203-1331001112132111-0213331101233011-0231321101312202-3213012113020230-3330122302332303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.any` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011)
- network_pbr.any

<a id="canonical-1003002300100132-0312322232230322-2301320020133101-1322213003110331-2100102012100300-0332022233320133-1133103123320002-2221321010110231"></a>

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

<a id="canonical-0122322112001130-0213130203210221-0302000123333310-3230311012002213-1031132132102010-1233312020012200-1233223233103103-0320303130031021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.label_selector` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011)
- network_pbr.label_selector

<a id="canonical-0111122201132010-3303201213011313-2113100003302110-3313233201301123-2100021103321300-2111301000331333-2231200003113123-2223323020101303"></a>

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

<a id="canonical-2031103331121020-3000002332331031-2022111011131231-2311330220213333-2302133313133131-2032321200103332-2011032320003310-3021231020101130"></a>

### Direct properties for `network_pbr.label_selector`

<a id="canonical-0301212100320300-1102212331233032-1202300312332203-0100123311200112-1232032221211321-0231000132333111-2120123123222122-3023021030323233"></a>

#### `network_pbr.label_selector.expressions` property

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

<a id="canonical-0112333310121320-1011211132322313-2120120101012110-3101011022221203-2221023311032232-2122212210313003-1022010133030330-2320323000210001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011)
- network_pbr.network_pbr_rules

<a id="canonical-2203310300222123-0123002223021120-3120023210223002-3233010220102122-0201110231322303-3000212333031132-2202113222000111-0130023010311310"></a>

Type: `"list"`. Computed.

L3/L4 Destination Routing Rules. Network(L3/L4) routing policy rule.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-0330221101232222-2211123213302232-3003203330131002-0002101332022001-0012223311311201-1110112301203220-3032301211311210-0020102211023302"></a>

### Direct properties for `network_pbr.network_pbr_rules`

- [all_tcp_traffic](data-sources--policy_based_routing--reference--group-001.md#canonical-2202123332101011-3102103012100011-1330301011300233-1001202133210131-1222120213311010-0112320200012230-2322221212230023-0203212333011301): complete subsection reference.

- [all_traffic](data-sources--policy_based_routing--reference--group-001.md#canonical-2320000212133031-0033030233112031-3230031123122322-2132022333220021-0103311122212230-2213000032302023-3303020020110012-1023130020232220): complete subsection reference.

- [all_udp_traffic](data-sources--policy_based_routing--reference--group-001.md#canonical-0130003212211122-0020132221110222-2311323333112312-0322313313231312-2220311012200210-3012201010120320-0103010321102111-1101033301031321): complete subsection reference.

- [any](data-sources--policy_based_routing--reference--group-001.md#canonical-3332210023010303-0002101311231111-2002331313010021-2210230203320111-2103203322012212-1310323032223013-0101210200320320-2322132123200133): complete subsection reference.

- [applications](data-sources--policy_based_routing--reference--group-001.md#canonical-2123103013333023-2200103322031003-2131210132220321-0132320102222023-1303212210303332-0213130220213323-0101013320122210-3002203001113210): complete subsection reference.

<a id="canonical-2201223111010222-2233210200313010-0122312133120112-3321333203122320-3230330023223331-3123113131033101-0230101030112003-0311011121203010"></a>

<a id="canonical-2000100001322023-2210012303203100-1001010200001101-2320302223323210-0301120202312020-3323001111311301-0132231020320231-3102130322303212"></a>

#### `network_pbr.network_pbr_rules.dns_name` property

Type: `"string"`. Computed.

Exclusive with \[any ip\_prefix\_set prefix\_list\] Resolve hostname to GET the IP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [forwarding_class_list](data-sources--policy_based_routing--reference--group-001.md#canonical-0101230202100221-1133031033022222-1211002213230201-1132113023113330-0003113212001303-1311101000301212-1002103111323112-1223312130322020): complete subsection reference.

- [ip_prefix_set](data-sources--policy_based_routing--reference--group-001.md#canonical-3132011220012130-0013120222210331-1321100223223310-0331212023012023-0321123303312223-2032302123310313-1131133233133112-1223301230221033): complete subsection reference.

- [metadata](data-sources--policy_based_routing--reference--group-001.md#canonical-0211202313010223-2323110333012010-1222210320213102-2223310100101322-0312322303012031-3033313032032220-2231033322202011-0222130330021322): complete subsection reference.

- [prefix_list](data-sources--policy_based_routing--reference--group-001.md#canonical-1201000231313320-1202003331031013-3333003220333122-2113101321133322-2203113110220102-3311300221221333-3310233023301102-2221101202022022): complete subsection reference.

- [protocol_port_range](data-sources--policy_based_routing--reference--group-001.md#canonical-1301113023123301-0131022111331211-3323121033210233-3223310322301103-3103021312200020-2332122321310013-1023300331013102-1101212211021110): complete subsection reference.

<a id="canonical-2202123332101011-3102103012100011-1330301011300233-1001202133210131-1222120213311010-0112320200012230-2322221212230023-0203212333011301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.all_tcp_traffic` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-0112333310121320-1011211132322313-2120120101012110-3101011022221203-2221023311032232-2122212210313003-1022010133030330-2320323000210001)
- network_pbr.network_pbr_rules.all_tcp_traffic

<a id="canonical-2002201111112010-3300303300010122-3033220010210213-0323221113322103-3020030310221300-0310012013022200-0320033112030223-3210022111002132"></a>

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

<a id="canonical-2320000212133031-0033030233112031-3230031123122322-2132022333220021-0103311122212230-2213000032302023-3303020020110012-1023130020232220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.all_traffic` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-0112333310121320-1011211132322313-2120120101012110-3101011022221203-2221023311032232-2122212210313003-1022010133030330-2320323000210001)
- network_pbr.network_pbr_rules.all_traffic

<a id="canonical-1032230231223000-2232212230112201-0013021130110330-2210313323032110-3201311110300312-0103331310313203-1122333202323003-0023002322212133"></a>

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

<a id="canonical-0130003212211122-0020132221110222-2311323333112312-0322313313231312-2220311012200210-3012201010120320-0103010321102111-1101033301031321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.all_udp_traffic` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-0112333310121320-1011211132322313-2120120101012110-3101011022221203-2221023311032232-2122212210313003-1022010133030330-2320323000210001)
- network_pbr.network_pbr_rules.all_udp_traffic

<a id="canonical-0311332320012103-2222100300221320-1122210313203301-3003113010210213-0232003112133001-1313020322230221-2121131200322313-3223021222002103"></a>

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

<a id="canonical-3332210023010303-0002101311231111-2002331313010021-2210230203320111-2103203322012212-1310323032223013-0101210200320320-2322132123200133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.any` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-0112333310121320-1011211132322313-2120120101012110-3101011022221203-2221023311032232-2122212210313003-1022010133030330-2320323000210001)
- network_pbr.network_pbr_rules.any

<a id="canonical-2333202103032000-2330113230011012-3113322123121211-3002032201210230-2002323321120203-1112002101321112-3130012131222300-3133031323303023"></a>

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

<a id="canonical-2123103013333023-2200103322031003-2131210132220321-0132320102222023-1303212210303332-0213130220213323-0101013320122210-3002203001113210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.applications` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-0112333310121320-1011211132322313-2120120101012110-3101011022221203-2221023311032232-2122212210313003-1022010133030330-2320323000210001)
- network_pbr.network_pbr_rules.applications

<a id="canonical-0113112200120333-2132110202032013-2332122312130133-0310010012201023-3123220131022100-1232112030020320-2002213203033201-3302322021333130"></a>

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

<a id="canonical-2300312302033220-3311001301301023-1023131001303010-2311032310223330-2101300022321100-1200322011201020-2201003023111323-3231202113100030"></a>

### Direct properties for `network_pbr.network_pbr_rules.applications`

<a id="canonical-3031121012133021-3022123232031231-0123121031321012-1231013131033010-0202130003310201-3030232011013103-2013123130011300-3002010031220131"></a>

#### `network_pbr.network_pbr_rules.applications.applications` property

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

<a id="canonical-0101230202100221-1133031033022222-1211002213230201-1132113023113330-0003113212001303-1311101000301212-1002103111323112-1223312130322020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.forwarding_class_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-0112333310121320-1011211132322313-2120120101012110-3101011022221203-2221023311032232-2122212210313003-1022010133030330-2320323000210001)
- network_pbr.network_pbr_rules.forwarding_class_list

<a id="canonical-0212320032233232-0103312113032133-3230130012330212-2310313100100320-0130303323213103-2200331002123211-3000101230302133-0102333310210123"></a>

Type: `"list"`. Computed.

Ordered list of forwarding Class to be used if rule match.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2233313313300020-3310222033122210-0100021220011013-1233332221012331-1201130210122213-2321022131120023-0022310220022313-1203031311301210"></a>

### Direct properties for `network_pbr.network_pbr_rules.forwarding_class_list`

<a id="canonical-1002310212001301-0321200133221322-2231231023021221-1120002323333020-2123200110023100-0103323223101101-1210302230121131-0130312031000030"></a>

#### `network_pbr.network_pbr_rules.forwarding_class_list.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1201233030323312-3032002201133330-2220302301321013-3020232331303032-0003021020101103-1312030013110112-2020301002122202-0023200132010312"></a>

<a id="canonical-0200121321311000-1102001202330012-3012301123033011-1221221002221210-0102301112120200-1021210211230002-2313232312032002-2203220313003013"></a>

#### `network_pbr.network_pbr_rules.forwarding_class_list.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3301223003100103-2010030231002303-1010220111113300-3223302333301223-1031203230302111-3211003112311332-1012112130220023-0023131002022303"></a>

<a id="canonical-2020203300021132-2032310111133021-0111330213123133-2100311313123321-0012103123303210-3133313101200011-2221123100121333-2130130113030200"></a>

#### `network_pbr.network_pbr_rules.forwarding_class_list.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3132011220012130-0013120222210331-1321100223223310-0331212023012023-0321123303312223-2032302123310313-1131133233133112-1223301230221033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.ip_prefix_set` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-0112333310121320-1011211132322313-2120120101012110-3101011022221203-2221023311032232-2122212210313003-1022010133030330-2320323000210001)
- network_pbr.network_pbr_rules.ip_prefix_set

<a id="canonical-1322312020212233-1032130331003333-0221132330320112-0003113333212021-2100113330313012-3300333312333112-3212230231100010-0130132031010003"></a>

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

<a id="canonical-0111001002310110-3321102230022023-2200213322111221-0001313332221302-0020012220212032-0210130023330321-1101201330222122-0222132333322311"></a>

### Direct properties for `network_pbr.network_pbr_rules.ip_prefix_set`

- [ref](data-sources--policy_based_routing--reference--group-001.md#canonical-1211303122021021-3200210331232023-1020010012120330-0312221300110312-3001233311033321-3123122310020300-1332032313333123-1302020222023110): complete subsection reference.

<a id="canonical-1211303122021021-3200210331232023-1020010012120330-0312221300110312-3001233311033321-3123122310020300-1332032313333123-1302020222023110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.ip_prefix_set.ref` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-0112333310121320-1011211132322313-2120120101012110-3101011022221203-2221023311032232-2122212210313003-1022010133030330-2320323000210001)
- [network_pbr.network_pbr_rules.ip_prefix_set](data-sources--policy_based_routing--reference--group-001.md#canonical-3132011220012130-0013120222210331-1321100223223310-0331212023012023-0321123303312223-2032302123310313-1131133233133112-1223301230221033)
- network_pbr.network_pbr_rules.ip_prefix_set.ref

<a id="canonical-2111101312230220-0330213323023330-2331323002211020-1333211310013133-3213223112203210-2330101232010122-2310102013101311-0132332203313111"></a>

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

<a id="canonical-3031130010112111-0301232302331131-0322312322301001-1020203032011200-0110032201201031-0212220112110013-1303230021013222-2232133102230011"></a>

### Direct properties for `network_pbr.network_pbr_rules.ip_prefix_set.ref`

<a id="canonical-2101023031100013-0101212322200232-2321133211130113-0210131113001301-0231021200033000-1300130133012101-3031002100010220-0322003021213222"></a>

#### `network_pbr.network_pbr_rules.ip_prefix_set.ref.kind` property

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

<a id="canonical-2132213231302311-0011023231312011-0123013302131022-3313030030000223-2221102312113000-2311111012020211-3220100030311021-1212113211333330"></a>

<a id="canonical-0213231110330330-1100010020011233-3203320300313221-2032313123002031-2221113000212221-0010222031023300-2303113330300332-2312330021302232"></a>

#### `network_pbr.network_pbr_rules.ip_prefix_set.ref.name` property

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

<a id="canonical-2021021213020110-1021300121303332-2131331000020021-2213310032130302-0321311233023102-1030221001123221-2213123210033103-3011003222111310"></a>

<a id="canonical-3232230230023100-3311310003303211-3112323312112011-3011013320220313-3120022013000110-2130321301103131-3211122000021122-0222011020333322"></a>

#### `network_pbr.network_pbr_rules.ip_prefix_set.ref.namespace` property

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

<a id="canonical-1212003033230213-1323110220231310-2111121102000023-1300123212133230-0331123102331110-3212211313010212-3101033201232002-3221201032231223"></a>

<a id="canonical-0301112010120200-3002330222102233-2113303301011123-3330310313330331-0300200233030010-0221113303203200-1302103201033303-3121021303113110"></a>

#### `network_pbr.network_pbr_rules.ip_prefix_set.ref.tenant` property

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

<a id="canonical-1133222212130211-3210003133000220-1033300222110233-0303002130003100-1032123101321021-0100100032212023-1112303100202031-0210112300333303"></a>

<a id="canonical-3322232123100020-2130203011311322-2322311330133000-2330000102332321-1133223222322132-0303132230130221-2032112011011331-2200210223323002"></a>

#### `network_pbr.network_pbr_rules.ip_prefix_set.ref.uid` property

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

<a id="canonical-0211202313010223-2323110333012010-1222210320213102-2223310100101322-0312322303012031-3033313032032220-2231033322202011-0222130330021322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.metadata` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-0112333310121320-1011211132322313-2120120101012110-3101011022221203-2221023311032232-2122212210313003-1022010133030330-2320323000210001)
- network_pbr.network_pbr_rules.metadata

<a id="canonical-0313332130213011-2311301131301200-0201232103110212-0100233031332202-3132323001332112-1323322333300121-0003303031321202-2222033323333123"></a>

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

<a id="canonical-0110203232131101-0120310010111220-3023230203301020-1102033211330003-2321322201132202-1032122322120103-0010223211231312-3130321210200130"></a>

### Direct properties for `network_pbr.network_pbr_rules.metadata`

<a id="canonical-2222022021031032-0203231020330100-3102012000112300-0311103120320112-3103023221332031-0222230112003203-3333021013330223-3030102333303003"></a>

#### `network_pbr.network_pbr_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2000111021021330-2102323123203222-0201102133131112-3231001223032221-2010123010120231-2230201000110000-0201003032220013-3311121233012031"></a>

<a id="canonical-3210230110221031-1131110010003012-2200000333231310-3033210210203032-0032131010321333-3333203003013030-2130222210211131-1011123332131112"></a>

#### `network_pbr.network_pbr_rules.metadata.name` property

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

<a id="canonical-1201000231313320-1202003331031013-3333003220333122-2113101321133322-2203113110220102-3311300221221333-3310233023301102-2221101202022022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.prefix_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-0112333310121320-1011211132322313-2120120101012110-3101011022221203-2221023311032232-2122212210313003-1022010133030330-2320323000210001)
- network_pbr.network_pbr_rules.prefix_list

<a id="canonical-1000200332003223-1232001202333132-1332203121100302-3221103211200202-1002221021212110-0302001113103132-1111103223323021-3032030012210130"></a>

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

<a id="canonical-1122311112133031-3010012300003020-0013201331331322-3231102301200031-0033313211102112-2320233313322133-1003332023222121-3133130123023221"></a>

### Direct properties for `network_pbr.network_pbr_rules.prefix_list`

<a id="canonical-1113210021230001-2230132103013021-2203132002333311-2311012201013132-2031231310331202-0122330103002230-1222123121312220-2221221032320012"></a>

#### `network_pbr.network_pbr_rules.prefix_list.prefixes` property

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

<a id="canonical-1301113023123301-0131022111331211-3323121033210233-3223310322301103-3103021312200020-2332122321310013-1023300331013102-1101212211021110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.network_pbr_rules.protocol_port_range` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011)
- [network_pbr.network_pbr_rules](data-sources--policy_based_routing--reference--group-001.md#canonical-0112333310121320-1011211132322313-2120120101012110-3101011022221203-2221023311032232-2122212210313003-1022010133030330-2320323000210001)
- network_pbr.network_pbr_rules.protocol_port_range

<a id="canonical-0303020020330122-1213123023321203-0320101013012131-2222102302001221-0132010231333202-3333003301211303-2331011231212101-0023133211313110"></a>

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

<a id="canonical-2023031011311301-1103003102211213-2212130102222030-2301230320202111-1312330211113030-0113201122330013-1031213032230322-0302320031331023"></a>

### Direct properties for `network_pbr.network_pbr_rules.protocol_port_range`

<a id="canonical-2320311100312331-1312132000310302-2301013101023212-0202103123001022-1302302010132300-2322200320021112-2321031303133030-1130202100123213"></a>

#### `network_pbr.network_pbr_rules.protocol_port_range.port_ranges` property

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

<a id="canonical-3000123220230320-2022002230113321-1133003120211132-3330012021032212-0130302231211220-3030002113131332-2303220131230332-2202001020132101"></a>

<a id="canonical-1201220003322302-0032232123331232-2020212011233130-1130000112010332-3200120110212220-0220023231232000-2302312322232303-0110210112120302"></a>

#### `network_pbr.network_pbr_rules.protocol_port_range.protocol` property

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

<a id="canonical-0232121212310000-1320100010331121-1320312232123020-2202103002223032-3002001120033200-1112310321101002-0023031200202202-3313213332021310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_pbr.prefix_list` properties

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Property reference](data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [network_pbr](data-sources--policy_based_routing--reference--group-001.md#canonical-2310123313133330-3300331130112200-3022301223200233-0133100230311323-0311303113010300-1101101122210220-0033203100001012-2213012302100011)
- network_pbr.prefix_list

<a id="canonical-2031120311002120-3223203222031001-3113232201313233-2013200123212120-0230112203312300-1333330110201302-1232220010313212-0202132021120220"></a>

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

<a id="canonical-3201123232320002-1022013233010312-2010221130121010-3312221312322333-2023220231132133-1113110013021301-2130300202220132-1321313122200333"></a>

### Direct properties for `network_pbr.prefix_list`

<a id="canonical-2133022100213211-1000311303021011-0011330310201233-3203001223300200-1222133331110003-1121202112320300-0021120302002113-0100202232103202"></a>

#### `network_pbr.prefix_list.prefixes` property

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
