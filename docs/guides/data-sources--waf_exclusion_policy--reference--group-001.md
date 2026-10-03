---
page_title: "xcsh_waf_exclusion_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_exclusion_policy reference."
---

# xcsh_waf_exclusion_policy reference

<a id="canonical-3020311023020023-3232003230320301-1322203311203211-2020223132023101-2231320232012112-0103201103323111-3033312300021222-0203023101301112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132222131032033-3132302010202233-2023120211330033-3213300211131111-2311331002323133-2201323321200001-3323031031100102-1101321230131203"></a>

## Property reference — Property reference / 323121221321 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)
- Property reference

<a id="canonical-0200332112221320-3322323011023101-2232113233203121-3033332322213302-1001100103230212-0001320203200021-0121120120130333-0233210001032122"></a>

## Direct properties — Property reference / 323121221321 / 3

<a id="canonical-0130001111321331-1303033332110321-1203032310122003-0020230001300132-3331020033211012-1133021012202323-0000101000033323-3132333203133303"></a>

<a id="canonical-1330210003221000-1320302011122200-1301233330212212-1320031310011333-3300233333200332-1331230132102102-0302213010311200-2121201313021311"></a>

## annotations property — Property reference / 323121221321 / 4

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

<a id="canonical-3003113012001232-3102222021331202-1211310010300021-1131010232231333-2310030200112122-1121230113100212-1122331230002013-2231130012233100"></a>

<a id="canonical-3212301030223230-0032000201313330-1023333330303212-1201003333333123-1210032323320103-0230231332323321-1102211020013302-2000201302313103"></a>

## description property — Property reference / 323121221321 / 5

Type: `"string"`. Computed.

Description of the WAFExclusionPolicy.

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

<a id="canonical-0032102311322332-1231122033122212-2210101212300313-2320232223321221-1130013120013202-2231012221201003-0212002310210322-1030302001230212"></a>

<a id="canonical-1002033000300003-2323231313000322-2100302212113133-3002013203331033-1220103130331012-2031303312310002-2201322132323232-1101132002113113"></a>

## ID property — Property reference / 323121221321 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0023333222133320-1312120122311022-3230232131112233-0220321111001031-3111201001232011-3221223112103213-3023203231103313-3032123212010220"></a>

<a id="canonical-1300022203220323-0212033021231332-2210001232200312-2103000122232303-3313312020020313-0112300122213133-0133030111021030-3313231111221301"></a>

## labels property — Property reference / 323121221321 / 7

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

<a id="canonical-3001222021101330-1111120013333312-3022310033312111-0333300112200010-3122212130203233-0103330112033211-0003222233223222-3301133203131202"></a>

<a id="canonical-0301333030101201-2033030102312113-0311012233031000-2122332111302013-2231221230110102-3001330201223021-0331312032233031-1001222110232122"></a>

## name property — Property reference / 323121221321 / 8

Type: `"string"`. Required.

Name of the WAFExclusionPolicy.

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

<a id="canonical-3021003121220302-2001231102232113-1333020212323212-0211030301303303-3232220323033203-2210101301311103-1020010013211113-2130000033132021"></a>

<a id="canonical-2302320020133001-1011122123001133-1112222133203223-1301321133321000-2301230300331330-1210322132332313-3211303230121023-0010100031030121"></a>

## namespace property — Property reference / 323121221321 / 9

Type: `"string"`. Required.

Namespace where the WAFExclusionPolicy exists.

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

- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230): complete subsection reference.

<a id="canonical-2100112303311011-1202203320100300-3030200000010322-2000211333031011-2200332323200313-2031212033033032-1110011310302023-2211230120321030"></a>

## All schema paths — Property reference / 323121221321 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0130001111321331-1303033332110321-1203032310122003-0020230001300132-3331020033211012-1133021012202323-0000101000033323-3132333203133303) |
| `description` | [description](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3003113012001232-3102222021331202-1211310010300021-1131010232231333-2310030200112122-1121230113100212-1122331230002013-2231130012233100) |
| `id` | [ID](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0032102311322332-1231122033122212-2210101212300313-2320232223321221-1130013120013202-2231012221201003-0212002310210322-1030302001230212) |
| `labels` | [labels](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0023333222133320-1312120122311022-3230232131112233-0220321111001031-3111201001232011-3221223112103213-3023203231103313-3032123212010220) |
| `name` | [name](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3001222021101330-1111120013333312-3022310033312111-0333300112200010-3122212130203233-0103330112033211-0003222233223222-3301133203131202) |
| `namespace` | [namespace](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3021003121220302-2001231102232113-1333020212323212-0211030301303303-3232220323033203-2210101301311103-1020010013211113-2130000033132021) |
| `waf_exclusion_rules` | [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-2300023020213231-2113103331320300-1221101312332333-2323103323300231-1020211020201033-2310222122020000-3111033022121312-1312030200120213) |
| `waf_exclusion_rules.any_domain` | [waf_exclusion_rules.any_domain](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3212032233000201-2323130121200201-0301000130023023-3300221311122120-3313300110121213-3103003033210132-0322033320232323-1332030031012112) |
| `waf_exclusion_rules.any_path` | [waf_exclusion_rules.any_path](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-2223303100131111-2133312220002113-2331101310200313-2310213032023021-0323000213000000-1013202232003010-0202300233310001-2321103130033231) |
| `waf_exclusion_rules.app_firewall_detection_control` | [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-1030321332302120-2122131120211112-1131021323330121-3002012200001032-2323201103133032-2320001201133310-2110111200201230-3003030303020203) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3323221330223030-2210210233232203-0201033313310312-0001112320111023-2012233121232111-2200310101231031-0230103331101000-3210110310122220) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-1010022322100312-3213031023110100-2313120020112103-1012213003221221-2220003230031203-0122020001001313-3131303303303221-2332100110300200) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context_name](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0001012122322310-1110302121212110-0032020223101221-2000100000212132-1300311213003331-2301123303112312-1220022020032120-1222212231132003) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-2321223012120123-0330102002132321-0001333320201221-2003003221113202-1332332313010132-3331233011221032-1132112100001132-1310330202020103) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-2112221212230230-1021113312131301-2312022311023200-2231230333331100-0203010023330011-0311122010002211-3212320200011322-1102331310311113) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-1001211022030100-2133232220322102-0231110011232331-1310303302203231-2202023313210313-3311031011213123-3021203020100003-1001011012211130) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3321023203130003-0320103130103010-3210021133203212-0301320022201331-2323112232132112-2032233210321103-1001131221010320-0322001313133113) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-1122321010320020-2221212300303112-3121232122211011-3120000221221232-3332201212302300-1013303031221230-0101312300202332-3330232020233230) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context_name](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-2331300210003013-1111110211233031-1022130222121322-1211333233031101-2132203233021301-0222201221221123-0303133030023302-3311213033031211) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.signature_id](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3313111120332230-1101300212133032-0220001130112113-2312201103303303-2221232231002313-1311301002003000-0102310322112133-3210002013202011) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0111320303113202-1230321003331321-3330003301123330-1321332101210220-0010110230001310-0301311233100332-0300321311332032-2200110031011102) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-2210001322133010-3232220311311133-0003001221301122-1003103132122301-0031030303120320-2311223300001032-3011013022023330-2003022200211123) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context_name](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3223310233303101-0001213302233032-3103210223223110-1131120000303310-2311213230201021-0320010332110002-2130023033302133-0112230002302020) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3313121111111002-1112202132333232-0332311023302232-0002102013212221-1103230332303310-2111020032030302-2331130022011123-1303230002011201) |
| `waf_exclusion_rules.exact_value` | [waf_exclusion_rules.exact_value](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3032030300201113-2130131111222112-2012231311023323-1310002233113211-0233012211311011-2123101222013120-3302213232030123-0320100021110030) |
| `waf_exclusion_rules.expiration_timestamp` | [waf_exclusion_rules.expiration_timestamp](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-1300130021123303-1313003203123221-0212112231010232-1211101213210332-3230001231033001-0103212113113001-0010100232232313-2201003122100002) |
| `waf_exclusion_rules.metadata` | [waf_exclusion_rules.metadata](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-1003322300220231-1021232233031232-2123132033201111-1223012220112011-1011320011321103-1221100212032033-0111322323202002-0023022112013100) |
| `waf_exclusion_rules.metadata.description_spec` | [waf_exclusion_rules.metadata.description_spec](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0000321331222023-3333311010313310-1131012310220033-3330022213132220-3002000120123330-1332200122320310-0112321003210111-0011101313101321) |
| `waf_exclusion_rules.metadata.name` | [waf_exclusion_rules.metadata.name](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3033333130030323-1111312111233020-3102202003233233-2222133111222323-2303200002023300-0120011301230111-2030203321132103-2300310233003123) |
| `waf_exclusion_rules.methods` | [waf_exclusion_rules.methods](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-2201000000301130-1313001013222013-1303033213222031-0312002313020213-3301120113202123-3133121012203120-2231212311022110-1220200120012030) |
| `waf_exclusion_rules.path_prefix` | [waf_exclusion_rules.path_prefix](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3112023113011103-3001213021121211-0110231103313120-1012002130300220-2303033322133331-2112233212103313-0323221213033023-0001220113103221) |
| `waf_exclusion_rules.path_regex` | [waf_exclusion_rules.path_regex](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0232123302122302-3021000212200201-2102110230232321-2133131021302010-0201213133010100-3203203320312232-2112010232300132-0010230012212123) |
| `waf_exclusion_rules.suffix_value` | [waf_exclusion_rules.suffix_value](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-1203321101230310-3022130300203023-0231333221002101-3010232130101213-0203111220200103-1211111031311333-1022022123200133-3130113111213033) |
| `waf_exclusion_rules.waf_skip_processing` | [waf_exclusion_rules.waf_skip_processing](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0213322110323021-0203220220111102-0030033020223331-0023302131120121-2233123101112201-0223212011121323-0212212120022233-1333200113302000) |

<a id="canonical-0313232222133300-3313120310223213-3220033112133131-3100032302332330-2020020233232031-0032123113023300-3102312031201100-3332303122323302"></a>

## Next pages — Property reference / 323121221321 / 11

- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)

<a id="canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233020332133033-3123132011103301-3003321330033220-0321232110321103-2330000312310231-2202211231023312-0200333201202212-3231020020232102"></a>

## waf_exclusion_rules — waf_exclusion_rules / 202033003231 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3020311023020023-3232003230320301-1322203311203211-2020223132023101-2231320232012112-0103201103323111-3033312300021222-0203023101301112)
- waf_exclusion_rules

<a id="canonical-2300023020213231-2113103331320300-1221101312332333-2323103323300231-1020211020201033-2310222122020000-3111033022121312-1312030200120213"></a>

Type: `"list"`. Computed.

WAF Exclusion Rules. An ordered list of rules.

Upstream description:

An ordered list of rules.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-2022122233232200-2213201101110033-0332102112230003-3132021031110120-3313300013203201-1100010032333321-1231333231133000-2111223233232322"></a>

## Direct properties — waf_exclusion_rules / 202033003231 / 3

- [any_domain](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-1031232131210220-2211302202113123-3233111102133200-1330130220133003-3220011231102021-1212032133222001-0021320311212320-2013012221100200): complete subsection reference.

- [any_path](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3121333301001221-0231212110202122-3212200130311301-2222021030032211-3101222201102231-3313232232231001-1320110233330010-2300310203210203): complete subsection reference.

- [app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3322312002010201-0212312321023111-3101111001131303-0123323202330030-2003331030333323-0332032131131313-1233312300222123-0002322310131112): complete subsection reference.

<a id="canonical-3032030300201113-2130131111222112-2012231311023323-1310002233113211-0233012211311011-2123101222013120-3302213232030123-0320100021110030"></a>

<a id="canonical-1301300112100210-1021220032000122-3103113002302123-2313220202003032-1310200131222002-1212012001031011-2312130222133102-3321020212322302"></a>

## exact_value property — waf_exclusion_rules / 202033003231 / 4

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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

<a id="canonical-1300130021123303-1313003203123221-0212112231010232-1211101213210332-3230001231033001-0103212113113001-0010100232232313-2201003122100002"></a>

<a id="canonical-2013103021200212-1302113010112222-1013002300022003-2132123112023331-3233102131032203-0321132320103222-1331321113213333-1120213233001001"></a>

## expiration_timestamp property — waf_exclusion_rules / 202033003231 / 5

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

- [metadata](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-1112230130323111-0132203102323222-2133211322112310-2113302232320311-2210033002012321-3103321123222011-3333323322210233-2013321221020222): complete subsection reference.

<a id="canonical-2201000000301130-1313001013222013-1303033213222031-0312002313020213-3301120113202123-3133121012203120-2231212311022110-1220200120012030"></a>

<a id="canonical-0321330230210111-2332200130100323-3011320121022230-3302112123223121-2211103331020102-1201321110002203-2003233033111331-1201011333120011"></a>

## methods property — waf_exclusion_rules / 202033003231 / 6

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3112023113011103-3001213021121211-0110231103313120-1012002130300220-2303033322133331-2112233212103313-0323221213033023-0001220113103221"></a>

<a id="canonical-0123002313033022-0103231001001300-1201223331113131-0310011221220313-2111203321010202-3110102011203230-1311311203301301-0311102131300131"></a>

## path_prefix property — waf_exclusion_rules / 202033003231 / 7

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

Upstream description:

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0232123302122302-3021000212200201-2102110230232321-2133131021302010-0201213133010100-3203203320312232-2112010232300132-0010230012212123"></a>

<a id="canonical-1131121212330011-2232030000100231-0203230120232222-0133131122332131-3120022003330223-3122233331132221-0221322310131320-0032311302133123"></a>

## path_regex property — waf_exclusion_rules / 202033003231 / 8

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\] Define the regular expression for the path. For example, the regular expression
^/.\*$ will match on all paths.

Upstream description:

Exclusive with \[any\_path path\_prefix\] Define the regular expression for the path. For example, the regular expression
^/.\*$ will match on all paths.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1203321101230310-3022130300203023-0231333221002101-3010232130101213-0203111220200103-1211111031311333-1022022123200133-3130113111213033"></a>

<a id="canonical-0130303230233002-2021310132310300-2132302013131130-0110313330300212-3300121203203103-0121232121113102-3300021002112222-1022332330132200"></a>

## suffix_value property — waf_exclusion_rules / 202033003231 / 9

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
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

- [waf_skip_processing](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0030000200213223-3131213201031122-3113331101003021-1213130322032301-0320232032122333-0203323212323100-2200102023333332-0101231222323112): complete subsection reference.

<a id="canonical-3322021323312033-3121100030201132-3220301011322221-1020323131001023-3311000022013002-1100013202121223-3231333132311331-3132223203131113"></a>

## Next pages — waf_exclusion_rules / 202033003231 / 10

- [waf_exclusion_rules.any_domain](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-1031232131210220-2211302202113123-3233111102133200-1330130220133003-3220011231102021-1212032133222001-0021320311212320-2013012221100200)
- [waf_exclusion_rules.any_path](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3121333301001221-0231212110202122-3212200130311301-2222021030032211-3101222201102231-3313232232231001-1320110233330010-2300310203210203)
- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3322312002010201-0212312321023111-3101111001131303-0123323202330030-2003331030333323-0332032131131313-1233312300222123-0002322310131112)
- [waf_exclusion_rules.metadata](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-1112230130323111-0132203102323222-2133211322112310-2113302232320311-2210033002012321-3103321123222011-3333323322210233-2013321221020222)
- [waf_exclusion_rules.waf_skip_processing](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0030000200213223-3131213201031122-3113331101003021-1213130322032301-0320232032122333-0203323212323100-2200102023333332-0101231222323112)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3020311023020023-3232003230320301-1322203311203211-2020223132023101-2231320232012112-0103201103323111-3033312300021222-0203023101301112)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)

<a id="canonical-1031232131210220-2211302202113123-3233111102133200-1330130220133003-3220011231102021-1212032133222001-0021320311212320-2013012221100200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221300023011031-1233013033021113-1323133100202200-0303033023223220-2021332303331133-0123312111303220-0312312202122200-3133300020023223"></a>

## waf_exclusion_rules.any_domain — any_domain / 220031201103 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3020311023020023-3232003230320301-1322203311203211-2020223132023101-2231320232012112-0103201103323111-3033312300021222-0203023101301112)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230)
- waf_exclusion_rules.any_domain

<a id="canonical-3212032233000201-2323130121200201-0301000130023023-3300221311122120-3313300110121213-3103003033210132-0322033320232323-1332030031012112"></a>

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

<a id="canonical-1111320200032211-3110313201333032-2033021020113300-1110020022200012-0330331332332322-3130010110333023-2131230112233122-3331003110103001"></a>

## Direct properties — any_domain / 220031201103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210303012322122-1110310030232223-1231332003231232-1230332332022133-2131321220200121-3311232121311001-2301223300131020-3301300101313011"></a>

## Next pages — any_domain / 220031201103 / 4

- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)

<a id="canonical-3121333301001221-0231212110202122-3212200130311301-2222021030032211-3101222201102231-3313232232231001-1320110233330010-2300310203210203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231303011122110-3030000121232213-1001010312003032-2111202110331003-0100110103220001-1303033000101001-2020233013321313-1122322223102032"></a>

## waf_exclusion_rules.any_path — any_path / 103333301011 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3020311023020023-3232003230320301-1322203311203211-2020223132023101-2231320232012112-0103201103323111-3033312300021222-0203023101301112)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230)
- waf_exclusion_rules.any_path

<a id="canonical-2223303100131111-2133312220002113-2331101310200313-2310213032023021-0323000213000000-1013202232003010-0202300233310001-2321103130033231"></a>

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

<a id="canonical-0130222212020032-0112101021301102-3222121231022220-2100033323003231-1012011030111102-0010021111100002-2032121311333023-2233233113221002"></a>

## Direct properties — any_path / 103333301011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013111212330230-2213121300000200-1011203313320100-1120301200123333-2323210033132230-2202320131023102-3212130221102303-1221111120231100"></a>

## Next pages — any_path / 103333301011 / 4

- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)

<a id="canonical-3322312002010201-0212312321023111-3101111001131303-0123323202330030-2003331030333323-0332032131131313-1233312300222123-0002322310131112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121320030333032-3331123320322030-2103211103201121-1000033101232131-2111202122301200-0231103233101122-2010111212302112-2013131222313112"></a>

## waf_exclusion_rules.app_firewall_detection_control — app_firewall_detection_control / 210031131223 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3020311023020023-3232003230320301-1322203311203211-2020223132023101-2231320232012112-0103201103323111-3033312300021222-0203023101301112)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230)
- waf_exclusion_rules.app_firewall_detection_control

<a id="canonical-1030321332302120-2122131120211112-1131021323330121-3002012200001032-2323201103133032-2320001201133310-2110111200201230-3003030303020203"></a>

Type: `"single"`. Computed.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

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

<a id="canonical-2333131330332321-2013023131312321-1122133311310010-3232032130031003-0330231213210102-3210031121113000-2232113320231300-3221032030230112"></a>

## Direct properties — app_firewall_detection_control / 210031131223 / 3

- [exclude_attack_type_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3030121321113100-2103102300213232-0322300123023223-0102103330302031-1120022233111313-3222012101113010-0323132313121002-0023331020223113): complete subsection reference.

- [exclude_bot_name_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-2122322231300101-0321033301023303-0100300320301033-3100303100200110-3021123320211231-3101300330302011-1230321210312023-3120220320222220): complete subsection reference.

- [exclude_signature_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3123010201210221-2232013322033120-2302013100133031-0132121120300123-2033023321220211-3230031213013302-1220221121313113-2010022230022302): complete subsection reference.

- [exclude_violation_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3111102103213331-3311003302110203-1132220322110302-0300311333133121-1112302213213121-1323212201111031-3131202320232003-0232320133130130): complete subsection reference.

<a id="canonical-1030213132210033-0333232110102010-1120212332033313-3020203330201310-2222100311321023-3111211112120222-2113202012302012-0121303323201011"></a>

## Next pages — app_firewall_detection_control / 210031131223 / 4

- [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3030121321113100-2103102300213232-0322300123023223-0102103330302031-1120022233111313-3222012101113010-0323132313121002-0023331020223113)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-2122322231300101-0321033301023303-0100300320301033-3100303100200110-3021123320211231-3101300330302011-1230321210312023-3120220320222220)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3123010201210221-2232013322033120-2302013100133031-0132121120300123-2033023321220211-3230031213013302-1220221121313113-2010022230022302)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3111102103213331-3311003302110203-1132220322110302-0300311333133121-1112302213213121-1323212201111031-3131202320232003-0232320133130130)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)

<a id="canonical-3030121321113100-2103102300213232-0322300123023223-0102103330302031-1120022233111313-3222012101113010-0323132313121002-0023331020223113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322211310222103-2133332001011323-1132230211131310-1323321001303111-3303013003220301-1111223002133210-1133210321132301-0220331222000030"></a>

## waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts — exclude_attack_type_contexts / 033020321112 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3020311023020023-3232003230320301-1322203311203211-2020223132023101-2231320232012112-0103201103323111-3033312300021222-0203023101301112)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230)
- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3322312002010201-0212312321023111-3101111001131303-0123323202330030-2003331030333323-0332032131131313-1233312300222123-0002322310131112)
- waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-3323221330223030-2210210233232203-0201033313310312-0001112320111023-2012233121232111-2200310101231031-0230103331101000-3210110310122220"></a>

Type: `"list"`. Computed.

Exclude an entire attack type only in the named context. For migrated per-parameter exceptions,
prefer this over signature-ID exclusions because one payload can trigger several signatures;
unrelated parameters and attack types remain protected.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1132102311022202-1331010133333110-3310311301030123-2003233233031103-3220211031300022-2022031003203232-1322030133310033-3301021300301312"></a>

## Direct properties — exclude_attack_type_contexts / 033020321112 / 3

<a id="canonical-1010022322100312-3213031023110100-2313120020112103-1012213003221221-2220003230031203-0122020001001313-3131303303303221-2332100110300200"></a>

<a id="canonical-3002221020300012-1212213102232302-1130213233203200-3312323012132302-1303033231230210-3201122313222222-3213012312300321-0012032030031301"></a>

## context property — exclude_attack_type_contexts / 033020321112 / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

Exclusion scope. Use CONTEXT\_PARAMETER with context\_name for one parameter, CONTEXT\_COOKIE for
one cookie, or CONTEXT\_ANY only for an intentionally global scope.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0001012122322310-1110302121212110-0032020223101221-2000100000212132-1300311213003331-2301123303112312-1220022020032120-1222212231132003"></a>

<a id="canonical-2220032231011013-0210102312120031-1112102310013301-1330213122033203-1203222032320132-1331130313320111-3010103203333231-2023030302021023"></a>

## context_name property — exclude_attack_type_contexts / 033020321112 / 5

Type: `"string"`. Computed.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2321223012120123-0330102002132321-0001333320201221-2003003221113202-1332332313010132-3331233011221032-1132112100001132-1310330202020103"></a>

<a id="canonical-2331233123202112-2200121112222202-0130112113001032-0302233332111323-3130301101302303-2123110332212030-3032203003120233-1001203023133013"></a>

## exclude_attack_type property — exclude_attack_type_contexts / 033020321112 / 6

Type: `"string"`. Computed.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of all Attack Types ATTACK\_TYPE\_NONE ATTACK\_TYPE\_NON\_BROWSER\_CLIENT
ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE
ATTACK\_TYPE\_DETECTION\_EVASION ATTACK\_TYPE\_VULNERABILITY\_SCAN
ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS..
Possible values are \`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Upstream description:

Attack-type enum excluded in this context, for example ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING. Other
attack types remain enforced.

Receipt-pinned upstream constraints:

```json
{
  "default": "ATTACK_TYPE_NONE",
  "enum": [
    "ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1330020102303213-1333313221331202-1203300111101313-1020223012031013-3110023303230210-3323031231023003-1312101221233323-1303033020022200"></a>

## Next pages — exclude_attack_type_contexts / 033020321112 / 7

- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3322312002010201-0212312321023111-3101111001131303-0123323202330030-2003331030333323-0332032131131313-1233312300222123-0002322310131112)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)

<a id="canonical-2122322231300101-0321033301023303-0100300320301033-3100303100200110-3021123320211231-3101300330302011-1230321210312023-3120220320222220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021131100032313-0311120233312332-2121222131131213-0320332122222021-3322201010220012-3012133312311122-1202322023303232-3320213320123012"></a>

## waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts — exclude_bot_name_contexts / 123001111032 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3020311023020023-3232003230320301-1322203311203211-2020223132023101-2231320232012112-0103201103323111-3033312300021222-0203023101301112)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230)
- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3322312002010201-0212312321023111-3101111001131303-0123323202330030-2003331030333323-0332032131131313-1233312300222123-0002322310131112)
- waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-2112221212230230-1021113312131301-2312022311023200-2231230333331100-0203010023330011-0311122010002211-3212320200011322-1102331310311113"></a>

Type: `"list"`. Computed.

Bot Names to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3313102323323033-2320033031311010-1032210130002001-0112332203203000-1203020303030011-2300211110331310-3000211301102331-3212112331101320"></a>

## Direct properties — exclude_bot_name_contexts / 123001111032 / 3

<a id="canonical-1001211022030100-2133232220322102-0231110011232331-1310303302203231-2202023313210313-3311031011213123-3021203020100003-1001011012211130"></a>

<a id="canonical-0221323233022303-2323200321333101-2120113020001112-2010000112200230-1233111123100031-1112123313231332-1101032021231303-1113031323322023"></a>

## bot_name property — exclude_bot_name_contexts / 123001111032 / 4

Type: `"string"`. Computed.

Bot Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1023332223230113-1302202133221223-1011102112322111-3123320031120321-0333113001333320-1021202131032320-2321231213222033-2002213100233320"></a>

## Next pages — exclude_bot_name_contexts / 123001111032 / 5

- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3322312002010201-0212312321023111-3101111001131303-0123323202330030-2003331030333323-0332032131131313-1233312300222123-0002322310131112)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)

<a id="canonical-3123010201210221-2232013322033120-2302013100133031-0132121120300123-2033023321220211-3230031213013302-1220221121313113-2010022230022302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320131231012232-2223221113221112-3211211022130221-3023123321132121-3132200032110301-3030133031021331-3333031023220000-0110133130233023"></a>

## waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts — exclude_signature_contexts / 002020321203 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3020311023020023-3232003230320301-1322203311203211-2020223132023101-2231320232012112-0103201103323111-3033312300021222-0203023101301112)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230)
- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3322312002010201-0212312321023111-3101111001131303-0123323202330030-2003331030333323-0332032131131313-1233312300222123-0002322310131112)
- waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-3321023203130003-0320103130103010-3210021133203212-0301320022201331-2323112232132112-2032233210321103-1001131221010320-0322001313133113"></a>

Type: `"list"`. Computed.

Signature IDs to be excluded for the defined match criteria.

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

<a id="canonical-2021100031201000-1300323231302000-0322231300303111-1222023200312310-0011010132120200-0120102202031321-2003330300311223-0323120331033303"></a>

## Direct properties — exclude_signature_contexts / 002020321203 / 3

<a id="canonical-1122321010320020-2221212300303112-3121232122211011-3120000221221232-3332201212302300-1013303031221230-0101312300202332-3330232020233230"></a>

<a id="canonical-3323000022003133-2221121213010101-1212011031123132-2121201133121332-1133100201121201-2020302232211012-1210020230312220-1311112000100332"></a>

## context property — exclude_signature_contexts / 002020321203 / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2331300210003013-1111110211233031-1022130222121322-1211333233031101-2132203233021301-0222201221221123-0303133030023302-3311213033031211"></a>

<a id="canonical-1333030030130003-2003101202011121-0200011320312210-1023133221030302-3032003230102112-0310322211120102-0322031023323002-1010303303303223"></a>

## context_name property — exclude_signature_contexts / 002020321203 / 5

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-3313111120332230-1101300212133032-0220001130112113-2312201103303303-2221232231002313-1311301002003000-0102310322112133-3210002013202011"></a>

<a id="canonical-0321111031312333-0021210220223310-3000031203330323-3023113110200113-3300213311130203-2322023121213010-0223322101012103-2100133022110322"></a>

## signature_id property — exclude_signature_contexts / 002020321203 / 6

Type: `"number"`. Computed.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 299999999,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  }
}
```

<a id="canonical-0300221310023303-0103112032100200-1020331202003302-2031222222131012-0010000200002213-3130022222100312-0000312222011232-2313310102133103"></a>

## Next pages — exclude_signature_contexts / 002020321203 / 7

- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3322312002010201-0212312321023111-3101111001131303-0123323202330030-2003331030333323-0332032131131313-1233312300222123-0002322310131112)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)

<a id="canonical-3111102103213331-3311003302110203-1132220322110302-0300311333133121-1112302213213121-1323212201111031-3131202320232003-0232320133130130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100101003020123-0003131222303322-1200033332113210-1122022111223321-0001222003132301-2012030130021002-3212302232033311-1303100022231312"></a>

## waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts — exclude_violation_contexts / 200131331130 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3020311023020023-3232003230320301-1322203311203211-2020223132023101-2231320232012112-0103201103323111-3033312300021222-0203023101301112)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230)
- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3322312002010201-0212312321023111-3101111001131303-0123323202330030-2003331030333323-0332032131131313-1233312300222123-0002322310131112)
- waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-0111320303113202-1230321003331321-3330003301123330-1321332101210220-0010110230001310-0301311233100332-0300321311332032-2200110031011102"></a>

Type: `"list"`. Computed.

Violations to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3022022313213112-3330232303133131-1131333123132102-0301111101101321-3300213202122020-2031220110332323-3333132030331213-0221123131211331"></a>

## Direct properties — exclude_violation_contexts / 200131331130 / 3

<a id="canonical-2210001322133010-3232220311311133-0003001221301122-1003103132122301-0031030303120320-2311223300001032-3011013022023330-2003022200211123"></a>

<a id="canonical-3032331232230111-3212003120231103-2131121032133300-0003113301013103-1100220000132312-3112233312112000-0232100120201223-1133130030203201"></a>

## context property — exclude_violation_contexts / 200131331130 / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3223310233303101-0001213302233032-3103210223223110-1131120000303310-2311213230201021-0320010332110002-2130023033302133-0112230002302020"></a>

<a id="canonical-3121300232220222-2122311021010033-2210100023013233-2102112023113030-1013123231123231-3231030031100202-3321021031002011-3310231032332232"></a>

## context_name property — exclude_violation_contexts / 200131331130 / 5

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-3313121111111002-1112202132333232-0332311023302232-0002102013212221-1103230332303310-2111020032030302-2331130022011123-1303230002011201"></a>

<a id="canonical-0321011020322123-0101213213321332-3320133021031122-1202133220100023-3102133101210320-2003313011100222-2200233203201031-1333121212121213"></a>

## exclude_violation property — exclude_violation_contexts / 200131331130 / 6

Type: `"string"`. Computed.

\[Enum:
VIOL\_NONE|VIOL\_FILETYPE|VIOL\_METHOD|VIOL\_MANDATORY\_HEADER|VIOL\_HTTP\_RESPONSE\_STATUS|VIOL\_REQUEST\_MAX\_LENGTH|VIOL\_FILE\_UPLOAD|VIOL\_FILE\_UPLOAD\_IN\_BODY|VIOL\_XML\_MALFORMED|VIOL\_JSON\_MALFORMED|VIOL\_ASM\_COOKIE\_MODIFIED|VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS|VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE|VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT|VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST|VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION|VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS|VIOL\_EVASION\_DIRECTORY\_TRAVERSALS|VIOL\_MALFORMED\_REQUEST|VIOL\_EVASION\_MULTIPLE\_DECODING|VIOL\_DATA\_GUARD|VIOL\_EVASION\_APACHE\_WHITESPACE|VIOL\_COOKIE\_MODIFIED|VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS|VIOL\_EVASION\_IIS\_BACKSLASHES|VIOL\_EVASION\_PERCENT\_U\_DECODING|VIOL\_EVASION\_BARE\_BYTE\_DECODING|VIOL\_EVASION\_BAD\_UNESCAPE|VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST|VIOL\_ENCODING|VIOL\_COOKIE\_MALFORMED|VIOL\_GRAPHQL\_FORMAT|VIOL\_GRAPHQL\_MALFORMED|VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\]
List of all supported Violation Types VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER
VIOL\_HTTP\_RESPONSE\_STATUS VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD
VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED
VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS.. Possible values are \`VIOL\_NONE\`,
\`VIOL\_FILETYPE\`, \`VIOL\_METHOD\`, \`VIOL\_MANDATORY\_HEADER\`, \`VIOL\_HTTP\_RESPONSE\_STATUS\`,
\`VIOL\_REQUEST\_MAX\_LENGTH\`, \`VIOL\_FILE\_UPLOAD\`, \`VIOL\_FILE\_UPLOAD\_IN\_BODY\`,
\`VIOL\_XML\_MALFORMED\`, \`VIOL\_JSON\_MALFORMED\`, \`VIOL\_ASM\_COOKIE\_MODIFIED\`,
\`VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE\`,
\`VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT\`, \`VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION\`,
\`VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS\`,
\`VIOL\_EVASION\_DIRECTORY\_TRAVERSALS\`, \`VIOL\_MALFORMED\_REQUEST\`,
\`VIOL\_EVASION\_MULTIPLE\_DECODING\`, \`VIOL\_DATA\_GUARD\`, \`VIOL\_EVASION\_APACHE\_WHITESPACE\`,
\`VIOL\_COOKIE\_MODIFIED\`, \`VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS\`,
\`VIOL\_EVASION\_IIS\_BACKSLASHES\`, \`VIOL\_EVASION\_PERCENT\_U\_DECODING\`,
\`VIOL\_EVASION\_BARE\_BYTE\_DECODING\`, \`VIOL\_EVASION\_BAD\_UNESCAPE\`,
\`VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST\`, \`VIOL\_ENCODING\`,
\`VIOL\_COOKIE\_MALFORMED\`, \`VIOL\_GRAPHQL\_FORMAT\`, \`VIOL\_GRAPHQL\_MALFORMED\`,
\`VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\`. Defaults to \`VIOL\_NONE\`.

Upstream description:

List of all supported Violation Types

VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER VIOL\_HTTP\_RESPONSE\_STATUS
VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED
VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS
VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT
VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION
VIOL\_HTTP\_PROTOCOL\_CRLF\_CHARACTERS\_BEFORE\_REQUEST\_START
VIOL\_HTTP\_PROTOCOL\_NO\_HOST\_HEADER\_IN\_HTTP\_1\_1\_REQUEST
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_PARAMETERS\_PARSING
VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS
VIOL\_HTTP\_PROTOCOL\_CONTENT\_LENGTH\_SHOULD\_BE\_A\_POSITIVE\_NUMBER
VIOL\_EVASION\_DIRECTORY\_TRAVERSALS VIOL\_MALFORMED\_REQUEST VIOL\_EVASION\_MULTIPLE\_DECODING
VIOL\_DATA\_GUARD VIOL\_EVASION\_APACHE\_WHITESPACE VIOL\_COOKIE\_MODIFIED
VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS VIOL\_EVASION\_IIS\_BACKSLASHES
VIOL\_EVASION\_PERCENT\_U\_DECODING VIOL\_EVASION\_BARE\_BYTE\_DECODING VIOL\_EVASION\_BAD\_UNESCAPE
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_FORMDATA\_REQUEST\_PARSING
VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST
VIOL\_HTTP\_PROTOCOL\_HIGH\_ASCII\_CHARACTERS\_IN\_HEADERS VIOL\_ENCODING VIOL\_COOKIE\_MALFORMED
VIOL\_GRAPHQL\_FORMAT VIOL\_GRAPHQL\_MALFORMED VIOL\_GRAPHQL\_INTROSPECTION\_QUERY.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIOL_NONE",
  "enum": [
    "VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START",
    "VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3102020023020100-2221103023103123-1102330302333201-0210330311310212-2312203223013220-3113230122030233-3210311012032302-2013200321102333"></a>

## Next pages — exclude_violation_contexts / 200131331130 / 7

- [waf_exclusion_rules.app_firewall_detection_control](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3322312002010201-0212312321023111-3101111001131303-0123323202330030-2003331030333323-0332032131131313-1233312300222123-0002322310131112)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)

<a id="canonical-1112230130323111-0132203102323222-2133211322112310-2113302232320311-2210033002012321-3103321123222011-3333323322210233-2013321221020222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310120121220331-1200113200121221-3121013222123033-2330033332020301-0110113221112230-1332302221322222-2020313201011233-0223002123212211"></a>

## waf_exclusion_rules.metadata — metadata / 031012021113 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3020311023020023-3232003230320301-1322203311203211-2020223132023101-2231320232012112-0103201103323111-3033312300021222-0203023101301112)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230)
- waf_exclusion_rules.metadata

<a id="canonical-1003322300220231-1021232233031232-2123132033201111-1223012220112011-1011320011321103-1221100212032033-0111322323202002-0023022112013100"></a>

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

<a id="canonical-1302120130221332-1212020321011202-2211013011022033-2223333123221033-3111211221000131-3321330113331122-0221112330331332-2032032010110003"></a>

## Direct properties — metadata / 031012021113 / 3

<a id="canonical-0000321331222023-3333311010313310-1131012310220033-3330022213132220-3002000120123330-1332200122320310-0112321003210111-0011101313101321"></a>

<a id="canonical-1022210002112213-2023103003012201-1013022231203102-3003220123323221-2303322302133030-2301110121022303-0301323202102333-2102130300323202"></a>

## description_spec property — metadata / 031012021113 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3033333130030323-1111312111233020-3102202003233233-2222133111222323-2303200002023300-0120011301230111-2030203321132103-2300310233003123"></a>

<a id="canonical-0221212330011033-3112110302220112-0203321012230333-2220211011222122-3021012032321221-0031122211330223-2321122132323133-0202323230000310"></a>

## name property — metadata / 031012021113 / 5

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

<a id="canonical-3211321223003001-2301332230313020-1031131123322001-3000130122303312-3232230132303031-1202311020300301-1122313032233220-3301331231203002"></a>

## Next pages — metadata / 031012021113 / 6

- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)

<a id="canonical-0030000200213223-3131213201031122-3113331101003021-1213130322032301-0320232032122333-0203323212323100-2200102023333332-0101231222323112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020120302302032-1321333120310030-2303323001032132-1301103210231310-1003101120001121-3023330211220031-0200210132132301-0301130023020120"></a>

## waf_exclusion_rules.waf_skip_processing — waf_skip_processing / 102212322101 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)
- [Property reference](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-3020311023020023-3232003230320301-1322203311203211-2020223132023101-2231320232012112-0103201103323111-3033312300021222-0203023101301112)
- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230)
- waf_exclusion_rules.waf_skip_processing

<a id="canonical-0213322110323021-0203220220111102-0030033020223331-0023302131120121-2233123101112201-0223212011121323-0212212120022233-1333200113302000"></a>

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

<a id="canonical-0310220033231312-2123003131000000-2332221031331201-0300011231332312-3031113201331000-0033113230301033-1123000302110022-2010303101210333"></a>

## Direct properties — waf_skip_processing / 102212322101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323131312222022-0033312313221332-3133123331313332-0010112222331103-3010313001213201-0013121230320200-0220213013022023-1220200323303000"></a>

## Next pages — waf_skip_processing / 102212322101 / 4

- [waf_exclusion_rules](data-sources--waf_exclusion_policy--reference--group-001.md#canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)
