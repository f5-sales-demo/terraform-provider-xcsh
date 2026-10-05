---
page_title: "xcsh_waf_exclusion_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_exclusion_policy reference."
---

# xcsh_waf_exclusion_policy reference

<a id="canonical-3201012220022020-0302300332132202-1210300000101133-0132320021112000-0303020230312320-2033302130202231-1302212000313220-0123110313003311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031011100003223-3212322201122020-0210103112131202-3021323013312223-3332132123233132-2121202101230013-2113331121101203-3303103201102332"></a>

## Property reference — Property reference / 103030012210 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)
- Property reference

<a id="canonical-2302313122331011-2323230300300222-1211203131102311-1323322213331213-0332030130320033-0103001133001132-2323221210120201-1322300232013130"></a>

## Direct properties — Property reference / 103030012210 / 3

<a id="canonical-2111111101303332-1320201132211301-0111030122132121-2332100333313333-0213012022312132-3122303222333112-0311301000233101-1023233110332301"></a>

<a id="canonical-2210302001321131-3110023323233310-0300202022223120-1310033031332011-1221003303322133-2202133022010200-3323310320333121-2032333322003311"></a>

## annotations property — Property reference / 103030012210 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

<a id="canonical-2023231233220131-1003020311130200-3223130211330331-2030232210002113-3113012303321021-0211030001220203-0001213211311230-3231233110002232"></a>

<a id="canonical-0233111111103010-0232212200111313-3011132001012200-2332300310133301-1222331021230100-1220132213301301-2302131333302310-3122301120233333"></a>

## description property — Property reference / 103030012210 / 5

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

<a id="canonical-2102100000030202-3122323000230323-3311121131132021-1230333301312323-0131100023331022-0023023211322030-3222012103010302-1211010112313203"></a>

<a id="canonical-2010102111231222-2220202130113103-0031101121131010-2012300320321011-1101033113032301-1330231123111121-1131101213203131-1200300113320003"></a>

## disable property — Property reference / 103030012210 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

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

<a id="canonical-0122312111333011-0300123113230310-3000021123313100-2020300011313213-0223330232321220-1122012332301222-0102312100013302-3331012030101012"></a>

<a id="canonical-3313213320011021-0110112231112003-3001010323111121-1132033103011213-0301011310212031-3023010213121003-3011003012031331-0320032122000322"></a>

## ID property — Property reference / 103030012210 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3223120010300110-3202331200232133-2310133203301203-3200020112112030-0220112203231322-2112021212022332-3012210300320130-3323123220132132"></a>

<a id="canonical-2323021003233002-3321133031103021-0122310023312331-0313003022313022-2203033330213230-1212120013001010-0100223032310103-2323133212233122"></a>

## labels property — Property reference / 103030012210 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-2020013231010210-1122031222030213-0232120132010030-2231013333123221-3121001112113200-1102023220201223-0110300221233000-0002310200100211"></a>

<a id="canonical-2101120311111031-3100011321223323-1110103312310112-3200200002132113-2121103111022303-2323110200121023-2222133000023031-1021203330030030"></a>

## name property — Property reference / 103030012210 / 9

Type: `"string"`. Required.

Name of the WAF Exclusion Policy. Must be unique within the namespace.

Upstream description:

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

<a id="canonical-2122012101301110-0212222212222201-0202023321230232-1022212102002031-0202232311133101-3330100202201033-0101113023022300-1231031311233020"></a>

<a id="canonical-3102320300112320-2221101023220333-2331332230010221-0223312223330101-2301212103303303-1022113113233101-1212221121030103-0000320022132011"></a>

## namespace property — Property reference / 103030012210 / 10

Type: `"string"`. Required.

Namespace where the WAF Exclusion Policy is created.

Upstream description:

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

- [timeouts](resources--waf_exclusion_policy--reference--group-001.md#canonical-3210220302330331-3113032111202011-3101313322200302-0303022000112330-3131010020211100-2031013130321333-1022303122302202-3121110220200132): complete subsection reference.

- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012): complete subsection reference.

<a id="canonical-2200300132311301-1223323032122010-0201032111011112-0312032233312033-3211021123222001-0313333313011122-0311113332122132-3132301101030110"></a>

## All schema paths — Property reference / 103030012210 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--waf_exclusion_policy--reference--group-001.md#canonical-2111111101303332-1320201132211301-0111030122132121-2332100333313333-0213012022312132-3122303222333112-0311301000233101-1023233110332301) |
| `description` | [description](resources--waf_exclusion_policy--reference--group-001.md#canonical-2023231233220131-1003020311130200-3223130211330331-2030232210002113-3113012303321021-0211030001220203-0001213211311230-3231233110002232) |
| `disable` | [disable](resources--waf_exclusion_policy--reference--group-001.md#canonical-2102100000030202-3122323000230323-3311121131132021-1230333301312323-0131100023331022-0023023211322030-3222012103010302-1211010112313203) |
| `id` | [ID](resources--waf_exclusion_policy--reference--group-001.md#canonical-0122312111333011-0300123113230310-3000021123313100-2020300011313213-0223330232321220-1122012332301222-0102312100013302-3331012030101012) |
| `labels` | [labels](resources--waf_exclusion_policy--reference--group-001.md#canonical-3223120010300110-3202331200232133-2310133203301203-3200020112112030-0220112203231322-2112021212022332-3012210300320130-3323123220132132) |
| `name` | [name](resources--waf_exclusion_policy--reference--group-001.md#canonical-2020013231010210-1122031222030213-0232120132010030-2231013333123221-3121001112113200-1102023220201223-0110300221233000-0002310200100211) |
| `namespace` | [namespace](resources--waf_exclusion_policy--reference--group-001.md#canonical-2122012101301110-0212222212222201-0202023321230232-1022212102002031-0202232311133101-3330100202201033-0101113023022300-1231031311233020) |
| `timeouts` | [timeouts](resources--waf_exclusion_policy--reference--group-001.md#canonical-1011103003202233-2123030332323333-0301330001020220-0201021213231221-0223101033032312-0212220310100002-1233320300232220-3221012331223010) |
| `timeouts.create` | [timeouts.create](resources--waf_exclusion_policy--reference--group-001.md#canonical-0303131012323133-2010310323231210-3221221313011211-0133003011210323-2133211113331203-0333320133232021-2300222232311221-3013211123120122) |
| `timeouts.delete` | [timeouts.delete](resources--waf_exclusion_policy--reference--group-001.md#canonical-1323110033322001-0003320322011120-0330003111131232-2330122112333310-1032211001120200-2312221123130232-2310011103102111-1202032301133032) |
| `timeouts.read` | [timeouts.read](resources--waf_exclusion_policy--reference--group-001.md#canonical-1200020131011102-1112122301220100-0300331033333331-2201030313002203-3320232210223030-3213000210011010-1231202130333130-2213122102322100) |
| `timeouts.update` | [timeouts.update](resources--waf_exclusion_policy--reference--group-001.md#canonical-1311122333231213-2020001211312033-2122103101321321-3210130031011231-1120221222122132-1131112103021221-1331003333102121-2323233113201032) |
| `waf_exclusion_rules` | [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-3332113103211223-2300330311302021-3330213301303301-3133223020122320-3233030003113001-3123023113122100-0230331112013113-2222030033022321) |
| `waf_exclusion_rules.any_domain` | [waf_exclusion_rules.any_domain](resources--waf_exclusion_policy--reference--group-001.md#canonical-0330022102223303-2301220233021200-2203221132130013-3303002033333202-3123322323001030-2013123331021132-0311002300101301-2302202023100213) |
| `waf_exclusion_rules.any_path` | [waf_exclusion_rules.any_path](resources--waf_exclusion_policy--reference--group-001.md#canonical-3010023131102133-0123110320231210-1133230001121010-1331203100120303-2311323311321002-1233122120003332-3032012123213323-2221222320010000) |
| `waf_exclusion_rules.app_firewall_detection_control` | [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-1010320301133303-2303213222110112-2032222100122302-2123201220321330-0330113111111210-2312101211223112-2030023022222200-3022030001302303) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-1203120133113303-1012333021123120-1001133013102131-1010021010321320-2113022021130333-3232312031011332-1013012333202012-1001223332100113) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context](resources--waf_exclusion_policy--reference--group-001.md#canonical-0120120230020030-2021230002021111-2330131330223013-0202010021020330-1113102323221212-0320033333332020-3022101320333022-2032310303321012) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context_name](resources--waf_exclusion_policy--reference--group-001.md#canonical-3302231012012301-3323212021321330-0131312211130313-0030213133202303-0030021310202101-2122123222221203-0311031203022231-3132013032323020) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](resources--waf_exclusion_policy--reference--group-001.md#canonical-3320302202313031-2013123030023112-0201002011023000-1223331233201101-0312021312000302-1132212023013301-0321011302201012-1331103120310122) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-3233013231103321-2303230333231210-1031310133111320-2033003132011121-3231023222332331-0121000101002111-2001202323122011-1302213101030022) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](resources--waf_exclusion_policy--reference--group-001.md#canonical-0201001000130110-1122100213320313-0200202101020233-0302100230123321-0310102203130123-1221033303120133-0210122201320032-3113113010331021) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-3133112302031223-1031010302332123-3323113012022203-0310111312111113-0021013213000111-0200010302313200-2330230312020231-0330102102022121) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context](resources--waf_exclusion_policy--reference--group-001.md#canonical-2320020320313212-3123201323012221-2232322303023303-3121202110202322-1022231111012111-0031211122203203-1231331313222231-2203133331031332) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context_name](resources--waf_exclusion_policy--reference--group-001.md#canonical-3232233033122201-0212012001112010-2331000333000210-3200323032121222-1211000303231211-2012220122211021-1230023201300203-2103003023021113) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.signature_id](resources--waf_exclusion_policy--reference--group-001.md#canonical-3201230222213130-2320111200231032-1111113113330102-1223113303210131-3101201220132312-0112333020031203-2112120033020111-1213202011301310) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-0021333230120303-3123010113311332-3011331011231301-0103330131022002-1323313231123021-1202303010213012-3312301322130121-3211301000333020) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context](resources--waf_exclusion_policy--reference--group-001.md#canonical-1200213100310013-3233010322020033-3133013033203120-0322033013122132-2313300321133310-2321022311113201-2112213222131210-2103002031132110) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context_name](resources--waf_exclusion_policy--reference--group-001.md#canonical-0133312033223022-3221102200122133-1000233323323322-0203020003022221-0020100303103322-1030013112331330-1220330322112113-2020110331101303) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](resources--waf_exclusion_policy--reference--group-001.md#canonical-1123302001223300-1030001203030132-0230131130211131-0213313310133332-3222223223302223-2121002333022230-1122223121302312-3121332132110301) |
| `waf_exclusion_rules.exact_value` | [waf_exclusion_rules.exact_value](resources--waf_exclusion_policy--reference--group-001.md#canonical-0101033213203211-2122220012112323-3330201303210133-2211212023133333-0302100212033133-2132212133310030-3131233133110213-2010120130210200) |
| `waf_exclusion_rules.expiration_timestamp` | [waf_exclusion_rules.expiration_timestamp](resources--waf_exclusion_policy--reference--group-001.md#canonical-3210031313003010-2203301202012022-0120321022300133-1001103332211322-3001210330113021-3032102330103321-2033323313011232-0133223302022022) |
| `waf_exclusion_rules.metadata` | [waf_exclusion_rules.metadata](resources--waf_exclusion_policy--reference--group-001.md#canonical-3222311333123220-2323130011012232-1022013201000031-1231122013322201-3211101332010132-0212110232332001-3211111210330331-3133000201230112) |
| `waf_exclusion_rules.metadata.description_spec` | [waf_exclusion_rules.metadata.description_spec](resources--waf_exclusion_policy--reference--group-001.md#canonical-2310232312021203-2211320102111002-0100210202013023-1323333103311111-2122033000122331-0101301032013200-2013311030232130-0133121320112031) |
| `waf_exclusion_rules.metadata.name` | [waf_exclusion_rules.metadata.name](resources--waf_exclusion_policy--reference--group-001.md#canonical-2200010132323313-2112223221302021-3110203111211101-0231302233011120-2113112033223100-0113023020211230-2002210220312331-2312123123313101) |
| `waf_exclusion_rules.methods` | [waf_exclusion_rules.methods](resources--waf_exclusion_policy--reference--group-001.md#canonical-3303331010302223-3331300231002302-0122102113213031-2211032102313301-1023203323313032-0331032331031232-2220223021211312-2210223230032313) |
| `waf_exclusion_rules.path_prefix` | [waf_exclusion_rules.path_prefix](resources--waf_exclusion_policy--reference--group-001.md#canonical-0221122130300002-3131313203021220-0211033021101321-0312133220221230-3021133330111133-0132213130231011-1322020020022121-2010200201230233) |
| `waf_exclusion_rules.path_regex` | [waf_exclusion_rules.path_regex](resources--waf_exclusion_policy--reference--group-001.md#canonical-3222332021011312-0103012001222033-2230002200332300-3000010302102001-2023320112100320-3133332322101131-3320231000122233-3012003100233122) |
| `waf_exclusion_rules.suffix_value` | [waf_exclusion_rules.suffix_value](resources--waf_exclusion_policy--reference--group-001.md#canonical-0012002103012310-1330232202201301-1133131230033020-0033232112012212-0331113021011021-0210032123322000-0320200221202232-1200102133001332) |
| `waf_exclusion_rules.waf_skip_processing` | [waf_exclusion_rules.waf_skip_processing](resources--waf_exclusion_policy--reference--group-001.md#canonical-1111223003123022-3231322113112213-0321013203322223-2320322022002211-2301310312330021-2333032303320033-3032010110203221-1223132121320221) |

<a id="canonical-3212330332122201-2233122103020011-1102203313312010-1303330312132131-0011003103332312-0001022313100133-2330120032032122-0321013223210202"></a>

## Next pages — Property reference / 103030012210 / 12

- [timeouts](resources--waf_exclusion_policy--reference--group-001.md#canonical-3210220302330331-3113032111202011-3101313322200302-0303022000112330-3131010020211100-2031013130321333-1022303122302202-3121110220200132)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)

<a id="canonical-3210220302330331-3113032111202011-3101313322200302-0303022000112330-3131010020211100-2031013130321333-1022303122302202-3121110220200132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003212123020133-1311313220322330-0033332312121120-1002110301031331-2203002213321001-2232333212011100-0102111220331132-0201030321210332"></a>

## timeouts — timeouts / 310230010033 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-3201012220022020-0302300332132202-1210300000101133-0132320021112000-0303020230312320-2033302130202231-1302212000313220-0123110313003311)
- timeouts

<a id="canonical-1011103003202233-2123030332323333-0301330001020220-0201021213231221-0223101033032312-0212220310100002-1233320300232220-3221012331223010"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220000103232202-0232203331132223-2010110010020222-2021032201210113-3302222302023113-0122310203310030-2132201121223100-3030020332212330"></a>

## Direct properties — timeouts / 310230010033 / 3

<a id="canonical-0303131012323133-2010310323231210-3221221313011211-0133003011210323-2133211113331203-0333320133232021-2300222232311221-3013211123120122"></a>

<a id="canonical-3203312023001333-0312320120210021-3332201310130323-2100100302003330-3132103320321201-1300130011213112-3022213103100332-1131330320210123"></a>

## create property — timeouts / 310230010033 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1323110033322001-0003320322011120-0330003111131232-2330122112333310-1032211001120200-2312221123130232-2310011103102111-1202032301133032"></a>

<a id="canonical-1011100130220213-3123211133303011-3133113303132033-2212212303200310-2233303130320101-2012020113121022-1023101003232330-3103011200322222"></a>

## delete property — timeouts / 310230010033 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1200020131011102-1112122301220100-0300331033333331-2201030313002203-3320232210223030-3213000210011010-1231202130333130-2213122102322100"></a>

<a id="canonical-2130233032121202-3230231300100302-2302320120131000-1113331133222011-3303333113133300-0121213001122210-3313201223212002-0030202132022132"></a>

## read property — timeouts / 310230010033 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1311122333231213-2020001211312033-2122103101321321-3210130031011231-1120221222122132-1131112103021221-1331003333102121-2323233113201032"></a>

<a id="canonical-3302111233100321-0012232013230222-1022321203312011-2132002111102311-0122123022233223-3311102100110023-3320300012310021-1120132320003301"></a>

## update property — timeouts / 310230010033 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2312310313001210-2230103032021230-3030130223133020-3132110323111101-2023313003033020-2331333003122220-0031013322203203-3112131131132102"></a>

## Next pages — timeouts / 310230010033 / 8

- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-3201012220022020-0302300332132202-1210300000101133-0132320021112000-0303020230312320-2033302130202231-1302212000313220-0123110313003311)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)

<a id="canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011302200112121-1323223120210121-3321120222210023-3102010333131203-1200113302322022-0320103131322313-3103002311100033-1302210233032020"></a>

## waf_exclusion_rules — waf_exclusion_rules / 213231313030 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-3201012220022020-0302300332132202-1210300000101133-0132320021112000-0303020230312320-2033302130202231-1302212000313220-0123110313003311)
- waf_exclusion_rules

<a id="canonical-3332113103211223-2300330311302021-3330213301303301-3133223020122320-3233030003113001-3123023113122100-0230331112013113-2222030033022321"></a>

Type: `"object"`. list nested block, Optional.

WAF Exclusion Rules. An ordered list of rules.

Upstream description:

An ordered list of rules.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "exact_value"),
  validators.ConflictingListObjectAttributes("any_domain",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_prefix"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_regex"),
  validators.ConflictingListObjectAttributes("app_firewall_detection_control",
    "waf_skip_processing"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("path_prefix",
    "path_regex")}
```

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

Terraform syntax:

```terraform
waf_exclusion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010312200230100-1231313122230130-3322000220103022-2123032311023011-2221213113322102-1130323312210330-1213123300010210-0012122121012321"></a>

## Direct properties — waf_exclusion_rules / 213231313030 / 3

- [any_domain](resources--waf_exclusion_policy--reference--group-001.md#canonical-3330010210021233-0130201323101213-1021112103023031-1310312230022330-1310321103013102-2302123020203211-2321110302321031-3320223133122302): complete subsection reference.

- [any_path](resources--waf_exclusion_policy--reference--group-001.md#canonical-0201023020120113-1323000123311302-1333100013102233-3012111200320221-0302222032330133-3020013302232311-2310331323321222-0300132323112330): complete subsection reference.

- [app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-3033331200333021-0211200202303333-3133032223121003-1130303000300131-2233330202331212-1210230101113303-0211123033312221-0120132322022313): complete subsection reference.

<a id="canonical-0101033213203211-2122220012112323-3330201303210133-2211212023133333-0302100212033133-2132212133310030-3131233133110213-2010120130210200"></a>

<a id="canonical-0232232321233313-1121233032212122-1330223231200011-3320023201132100-3300000121203222-0100221312021312-3013231000330031-2130133032113213"></a>

## exact_value property — waf_exclusion_rules / 213231313030 / 4

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-3210031313003010-2203301202012022-0120321022300133-1001103332211322-3001210330113021-3032102330103321-2033323313011232-0133223302022022"></a>

<a id="canonical-3102232303021101-0302302323303222-1010113031301021-3321103201000200-2202300211310201-3321000013121112-1111233300023111-2000120200210123"></a>

## expiration_timestamp property — waf_exclusion_rules / 213231313030 / 5

Type: `"string"`. Optional.

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

- [metadata](resources--waf_exclusion_policy--reference--group-001.md#canonical-2311323130333222-3113330103001312-3020301011221112-3103322023021012-1303023301123211-2131321231201031-2030103213222113-0302001231101103): complete subsection reference.

<a id="canonical-3303331010302223-3331300231002302-0122102113213031-2211032102313301-1023203323313032-0331032331031232-2220223021211312-2210223230032313"></a>

<a id="canonical-1022132031033131-2323133231212013-0012032122031212-3220102023310103-0010203203123310-2211110002232212-0102111103122111-3202003230020020"></a>

## methods property — waf_exclusion_rules / 213231313030 / 6

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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

<a id="canonical-0221122130300002-3131313203021220-0211033021101321-0312133220221230-3021133330111133-0132213130231011-1322020020022121-2010200201230233"></a>

<a id="canonical-0330230033111212-2321130031120301-0213031120210302-0200120132313302-0323322220323121-2001311001213023-0111310003231122-2102011300221131"></a>

## path_prefix property — waf_exclusion_rules / 213231313030 / 7

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

Upstream description:

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-3222332021011312-0103012001222033-2230002200332300-3000010302102001-2023320112100320-3133332322101131-3320231000122233-3012003100233122"></a>

<a id="canonical-3233320310201031-3021311332322200-2223102131031021-0102102303131312-3332330232110313-2232133322113232-1032323313121203-1113120123102033"></a>

## path_regex property — waf_exclusion_rules / 213231313030 / 8

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\] Define the regular expression for the path. For example, the regular expression
^/.\*$ will match on all paths.

Upstream description:

Exclusive with \[any\_path path\_prefix\] Define the regular expression for the path. For example, the regular expression
^/.\*$ will match on all paths.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-0012002103012310-1330232202201301-1133131230033020-0033232112012212-0331113021011021-0210032123322000-0320200221202232-1200102133001332"></a>

<a id="canonical-3032330010323111-1222211123320213-3220103122323130-0000003122222233-3011332300331110-1323133331121020-3021101220220321-1023331121323022"></a>

## suffix_value property — waf_exclusion_rules / 213231313030 / 9

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

- [waf_skip_processing](resources--waf_exclusion_policy--reference--group-001.md#canonical-1211112303030110-1100130231023332-3201000020101201-3013232112230302-1210330310233010-0233333332101112-2213132333012113-2213220303222123): complete subsection reference.

<a id="canonical-3222000211320320-2133111200331100-2023123122303031-2232331020211310-0221001130200220-2321332011232101-0330102130113020-0112220101313132"></a>

## Next pages — waf_exclusion_rules / 213231313030 / 10

- [waf_exclusion_rules.any_domain](resources--waf_exclusion_policy--reference--group-001.md#canonical-3330010210021233-0130201323101213-1021112103023031-1310312230022330-1310321103013102-2302123020203211-2321110302321031-3320223133122302)
- [waf_exclusion_rules.any_path](resources--waf_exclusion_policy--reference--group-001.md#canonical-0201023020120113-1323000123311302-1333100013102233-3012111200320221-0302222032330133-3020013302232311-2310331323321222-0300132323112330)
- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-3033331200333021-0211200202303333-3133032223121003-1130303000300131-2233330202331212-1210230101113303-0211123033312221-0120132322022313)
- [waf_exclusion_rules.metadata](resources--waf_exclusion_policy--reference--group-001.md#canonical-2311323130333222-3113330103001312-3020301011221112-3103322023021012-1303023301123211-2131321231201031-2030103213222113-0302001231101103)
- [waf_exclusion_rules.waf_skip_processing](resources--waf_exclusion_policy--reference--group-001.md#canonical-1211112303030110-1100130231023332-3201000020101201-3013232112230302-1210330310233010-0233333332101112-2213132333012113-2213220303222123)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-3201012220022020-0302300332132202-1210300000101133-0132320021112000-0303020230312320-2033302130202231-1302212000313220-0123110313003311)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)

<a id="canonical-3330010210021233-0130201323101213-1021112103023031-1310312230022330-1310321103013102-2302123020203211-2321110302321031-3320223133122302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313032032021103-2102330223101202-0301322013211330-0331132110122000-2022313221021220-1322121213200100-0022100203031301-1031331133122022"></a>

## waf_exclusion_rules.any_domain — any_domain / 122102311310 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-3201012220022020-0302300332132202-1210300000101133-0132320021112000-0303020230312320-2033302130202231-1302212000313220-0123110313003311)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012)
- waf_exclusion_rules.any_domain

<a id="canonical-0330022102223303-2301220233021200-2203221132130013-3303002033333202-3123322323001030-2013123331021132-0311002300101301-2302202023100213"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
any_domain = {}
```

<a id="canonical-0212202322311131-1011332320132223-0003000230330223-2320230123001230-1221202122102220-0221123320121133-1330010102013333-2312100231131201"></a>

## Direct properties — any_domain / 122102311310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213032330021133-3211113113323120-0223112021112131-1000003033010103-0323131333111212-3212012021002112-3101211311000000-1003302301303131"></a>

## Next pages — any_domain / 122102311310 / 4

- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)

<a id="canonical-0201023020120113-1323000123311302-1333100013102233-3012111200320221-0302222032330133-3020013302232311-2310331323321222-0300132323112330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022311230102112-2113313330002220-2333120001210011-2312022321303322-1221021303021330-1330031001210221-1012312033003332-2010101233011120"></a>

## waf_exclusion_rules.any_path — any_path / 110231330321 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-3201012220022020-0302300332132202-1210300000101133-0132320021112000-0303020230312320-2033302130202231-1302212000313220-0123110313003311)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012)
- waf_exclusion_rules.any_path

<a id="canonical-3010023131102133-0123110320231210-1133230001121010-1331203100120303-2311323311321002-1233122120003332-3032012123213323-2221222320010000"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
any_path = {}
```

<a id="canonical-2003221131123330-3123121200133310-0021311032302312-2203120131000212-1133331310232311-3331202020000302-0330010231010302-2302101311101200"></a>

## Direct properties — any_path / 110231330321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033213023213013-1112230220030011-0011323030101030-3200011010000111-3231332120021222-3220220112300330-2300211301322021-2220200021101101"></a>

## Next pages — any_path / 110231330321 / 4

- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)

<a id="canonical-3033331200333021-0211200202303333-3133032223121003-1130303000300131-2233330202331212-1210230101113303-0211123033312221-0120132322022313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221022300221203-1201312223103320-1110120223310113-3032222320322031-1323331331331201-3221001123123212-3221333222331012-0011332221000332"></a>

## waf_exclusion_rules.app_firewall_detection_control — app_firewall_detection_control / 100310333220 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-3201012220022020-0302300332132202-1210300000101133-0132320021112000-0303020230312320-2033302130202231-1302212000313220-0123110313003311)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012)
- waf_exclusion_rules.app_firewall_detection_control

<a id="canonical-1010320301133303-2303213222110112-2032222100122302-2123201220321330-0330113111111210-2312101211223112-2030023022222200-3022030001302303"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
app_firewall_detection_control {
  # Configure direct properties listed below.
}
```

<a id="canonical-3012312203211020-3303200211203023-0220313300122012-0211233103313312-3320310120222320-1023033310220202-2203010131123300-2202122232303201"></a>

## Direct properties — app_firewall_detection_control / 100310333220 / 3

- [exclude_attack_type_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-0032113131223101-2032333302300312-1121021132010023-3320132100212210-0132013223320201-2130301030020123-3032310333221123-1330112012131203): complete subsection reference.

- [exclude_bot_name_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-2023312020310013-0223003320122301-3302023002030222-3221000023213303-1213233220000111-2003211211003301-0232203033320200-1310220320231132): complete subsection reference.

- [exclude_signature_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-1230032303210112-2223112101330230-0001013102313322-0212121030010122-0332122123222130-2303132023312331-3300322011100022-2100120012211320): complete subsection reference.

- [exclude_violation_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-1133011201100000-0311130202022202-0130130030303013-3102132212031332-0132022203132311-2111222113330230-2100120032000032-2013211131031130): complete subsection reference.

<a id="canonical-2223123333312322-0310111203203212-0212031012222223-3130302301032210-0033331020212020-2110312212132302-0032031010100312-2222012200013310"></a>

## Next pages — app_firewall_detection_control / 100310333220 / 4

- [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-0032113131223101-2032333302300312-1121021132010023-3320132100212210-0132013223320201-2130301030020123-3032310333221123-1330112012131203)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-2023312020310013-0223003320122301-3302023002030222-3221000023213303-1213233220000111-2003211211003301-0232203033320200-1310220320231132)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-1230032303210112-2223112101330230-0001013102313322-0212121030010122-0332122123222130-2303132023312331-3300322011100022-2100120012211320)
- [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts](resources--waf_exclusion_policy--reference--group-001.md#canonical-1133011201100000-0311130202022202-0130130030303013-3102132212031332-0132022203132311-2111222113330230-2100120032000032-2013211131031130)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)

<a id="canonical-0032113131223101-2032333302300312-1121021132010023-3320132100212210-0132013223320201-2130301030020123-3032310333221123-1330112012131203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000300310323320-3003112012223210-2110210212320022-2112012100233322-3231032011200303-3001033131120200-2003121032023303-3010320200313133"></a>

## waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts — exclude_attack_type_contexts / 232300321130 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-3201012220022020-0302300332132202-1210300000101133-0132320021112000-0303020230312320-2033302130202231-1302212000313220-0123110313003311)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012)
- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-3033331200333021-0211200202303333-3133032223121003-1130303000300131-2233330202331212-1210230101113303-0211123033312221-0120132322022313)
- waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-1203120133113303-1012333021123120-1001133013102131-1010021010321320-2113022021130333-3232312031011332-1013012333202012-1001223332100113"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
exclude_attack_type_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1230220230333322-2021011121102201-2303123320021332-1213000123320122-1331212101130102-0231201201231002-2131133021021001-0130202202303020"></a>

## Direct properties — exclude_attack_type_contexts / 232300321130 / 3

<a id="canonical-0120120230020030-2021230002021111-2330131330223013-0202010021020330-1113102323221212-0320033333332020-3022101320333022-2032310303321012"></a>

<a id="canonical-1111300330021230-2122320201311112-1212022122330210-1220210023313123-3322213330100220-0313100013120233-2321111320211022-1233320203013122"></a>

## context property — exclude_attack_type_contexts / 232300321130 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CONTEXT_ANY","CONTEXT_BODY","CONTEXT_COOKIE","CONTEXT_HEADER","CONTEXT_PARAMETER","CONTEXT_REQUEST","CONTEXT_RESPONSE","CONTEXT_URI","CONTEXT_URL"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

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

<a id="canonical-3302231012012301-3323212021321330-0131312211130313-0030213133202303-0030021310202101-2122123222221203-0311031203022231-3132013032323020"></a>

<a id="canonical-1132112000222333-2210032232110010-0301220123302000-1020033300031310-3233202203212230-3230233212302221-1233301322233033-3303211211121230"></a>

## context_name property — exclude_attack_type_contexts / 232300321130 / 5

Type: `"string"`. Optional.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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

<a id="canonical-3320302202313031-2013123030023112-0201002011023000-1223331233201101-0312021312000302-1132212023013301-0321011302201012-1331103120310122"></a>

<a id="canonical-0321110223212013-0120202311311023-1113002331201120-1300033220202312-3102130211131102-0130002001200130-0000030303131322-2101310032122112"></a>

## exclude_attack_type property — exclude_attack_type_contexts / 232300321130 / 6

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY","ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS","ATTACK_TYPE_BUFFER_OVERFLOW","ATTACK_TYPE_COMMAND_EXECUTION","ATTACK_TYPE_CROSS_SITE_SCRIPTING","ATTACK_TYPE_DENIAL_OF_SERVICE","ATTACK_TYPE_DETECTION_EVASION","ATTACK_TYPE_DIRECTORY_INDEXING","ATTACK_TYPE_FORCEFUL_BROWSING","ATTACK_TYPE_GRAPHQL_PARSER_ATTACK","ATTACK_TYPE_HTTP_PARSER_ATTACK","ATTACK_TYPE_HTTP_RESPONSE_SPLITTING","ATTACK_TYPE_INFORMATION_LEAKAGE","ATTACK_TYPE_LDAP_INJECTION","ATTACK_TYPE_MALICIOUS_FILE_UPLOAD","ATTACK_TYPE_NONE","ATTACK_TYPE_NON_BROWSER_CLIENT","ATTACK_TYPE_OTHER_APPLICATION_ATTACKS","ATTACK_TYPE_PATH_TRAVERSAL","ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION","ATTACK_TYPE_REMOTE_FILE_INCLUDE","ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION","ATTACK_TYPE_SESSION_HIJACKING","ATTACK_TYPE_SQL_INJECTION","ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE","ATTACK_TYPE_VULNERABILITY_SCAN","ATTACK_TYPE_XPATH_INJECTION"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ATTACK_TYPE_NONE",
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
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"),
}
```

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

<a id="canonical-3330211210031002-2133312013312113-3320113201301130-2122301011302202-0002300300223121-2130211310111111-3021213000102120-3121131112100121"></a>

## Next pages — exclude_attack_type_contexts / 232300321130 / 7

- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-3033331200333021-0211200202303333-3133032223121003-1130303000300131-2233330202331212-1210230101113303-0211123033312221-0120132322022313)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)

<a id="canonical-2023312020310013-0223003320122301-3302023002030222-3221000023213303-1213233220000111-2003211211003301-0232203033320200-1310220320231132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011210010200123-3103303021333121-1123220313101322-2102330122003320-0003303112110321-0021230030020221-1021122112133203-0022121221331320"></a>

## waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts — exclude_bot_name_contexts / 013012310301 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-3201012220022020-0302300332132202-1210300000101133-0132320021112000-0303020230312320-2033302130202231-1302212000313220-0123110313003311)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012)
- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-3033331200333021-0211200202303333-3133032223121003-1130303000300131-2233330202331212-1210230101113303-0211123033312221-0120132322022313)
- waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-3233013231103321-2303230333231210-1031310133111320-2033003132011121-3231023222332331-0121000101002111-2001202323122011-1302213101030022"></a>

Type: `"object"`. list nested block, Optional.

Bot Names to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("bot_name")}
```

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

Terraform syntax:

```terraform
exclude_bot_name_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233332213101222-1311003120000030-3101122200330323-1102332113112033-3002111003211232-2010302120230101-3131312121021003-3130300101131102"></a>

## Direct properties — exclude_bot_name_contexts / 013012310301 / 3

<a id="canonical-0201001000130110-1122100213320313-0200202101020233-0302100230123321-0310102203130123-1221033303120133-0210122201320032-3113113010331021"></a>

<a id="canonical-2120011132032013-2301113212231222-0133230232120202-1131113130002130-2021031303213230-1233031313212301-0112323210322222-0221001201221203"></a>

## bot_name property — exclude_bot_name_contexts / 013012310301 / 4

Type: `"string"`. Optional.

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

<a id="canonical-3210323221110111-2032230332013130-0100323310132132-3231223033222212-1130222020202322-1032301132020220-1231223201103220-2033320002133331"></a>

## Next pages — exclude_bot_name_contexts / 013012310301 / 5

- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-3033331200333021-0211200202303333-3133032223121003-1130303000300131-2233330202331212-1210230101113303-0211123033312221-0120132322022313)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)

<a id="canonical-1230032303210112-2223112101330230-0001013102313322-0212121030010122-0332122123222130-2303132023312331-3300322011100022-2100120012211320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133233112100100-1312010303020130-1011000200323231-0312311323303330-1313202003032021-3233202301022201-2210200113210030-1021212100312332"></a>

## waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts — exclude_signature_contexts / 310331002331 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-3201012220022020-0302300332132202-1210300000101133-0132320021112000-0303020230312320-2033302130202231-1302212000313220-0123110313003311)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012)
- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-3033331200333021-0211200202303333-3133032223121003-1130303000300131-2233330202331212-1210230101113303-0211123033312221-0120132322022313)
- waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-3133112302031223-1031010302332123-3323113012022203-0310111312111113-0021013213000111-0200010302313200-2330230312020231-0330102102022121"></a>

Type: `"object"`. list nested block, Optional.

Signature IDs to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("signature_id")}
```

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

Terraform syntax:

```terraform
exclude_signature_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320310232033221-1130312111100112-3233030310333031-2213102120012121-2013111320310332-2200021230332023-1030230233232000-2200111312030003"></a>

## Direct properties — exclude_signature_contexts / 310331002331 / 3

<a id="canonical-2320020320313212-3123201323012221-2232322303023303-3121202110202322-1022231111012111-0031211122203203-1231331313222231-2203133331031332"></a>

<a id="canonical-2312022233130331-0213320021103303-2030231110110332-1221202031032013-1310201312310312-3212101032201023-2220012000113313-0213002130232332"></a>

## context property — exclude_signature_contexts / 310331002331 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CONTEXT_ANY","CONTEXT_BODY","CONTEXT_COOKIE","CONTEXT_HEADER","CONTEXT_PARAMETER","CONTEXT_REQUEST","CONTEXT_RESPONSE","CONTEXT_URI","CONTEXT_URL"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

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

<a id="canonical-3232233033122201-0212012001112010-2331000333000210-3200323032121222-1211000303231211-2012220122211021-1230023201300203-2103003023021113"></a>

<a id="canonical-2033013021103213-1133020320312122-1031202322003001-2331113130131120-2233113212001311-2211323111333332-2323130113231123-1003030020023200"></a>

## context_name property — exclude_signature_contexts / 310331002331 / 5

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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

<a id="canonical-3201230222213130-2320111200231032-1111113113330102-1223113303210131-3101201220132312-0112333020031203-2112120033020111-1213202011301310"></a>

<a id="canonical-2030022203310020-2011313200033232-1010231232233000-1201220321201003-1312332202310203-0013012223331232-0222212310230330-3102023111213113"></a>

## signature_id property — exclude_signature_contexts / 310331002331 / 6

Type: `"number"`. Optional.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 299999999),
}
```

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

<a id="canonical-0121310110111010-1203131231111113-1120120313303232-3122331123111033-1233203130203332-2330311212200332-0130003312020323-1320300031133133"></a>

## Next pages — exclude_signature_contexts / 310331002331 / 7

- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-3033331200333021-0211200202303333-3133032223121003-1130303000300131-2233330202331212-1210230101113303-0211123033312221-0120132322022313)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)

<a id="canonical-1133011201100000-0311130202022202-0130130030303013-3102132212031332-0132022203132311-2111222113330230-2100120032000032-2013211131031130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331002201111131-0321103332002331-1002212230330103-2100331330233303-2231231113123313-1231123321231310-1002032303203203-0313031011303130"></a>

## waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts — exclude_violation_contexts / 230022111122 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-3201012220022020-0302300332132202-1210300000101133-0132320021112000-0303020230312320-2033302130202231-1302212000313220-0123110313003311)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012)
- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-3033331200333021-0211200202303333-3133032223121003-1130303000300131-2233330202331212-1210230101113303-0211123033312221-0120132322022313)
- waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-0021333230120303-3123010113311332-3011331011231301-0103330131022002-1323313231123021-1202303010213012-3312301322130121-3211301000333020"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
exclude_violation_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322333133330023-3212300101133111-3022032133120022-1010033332131030-2212330301210213-0320200011003322-3220031233100201-0303013000333131"></a>

## Direct properties — exclude_violation_contexts / 230022111122 / 3

<a id="canonical-1200213100310013-3233010322020033-3133013033203120-0322033013122132-2313300321133310-2321022311113201-2112213222131210-2103002031132110"></a>

<a id="canonical-2120223301033232-0130201213303112-2033301201133022-2301301222323032-0030122023003103-2010221231303323-3202200230020020-0212221300133123"></a>

## context property — exclude_violation_contexts / 230022111122 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CONTEXT_ANY","CONTEXT_BODY","CONTEXT_COOKIE","CONTEXT_HEADER","CONTEXT_PARAMETER","CONTEXT_REQUEST","CONTEXT_RESPONSE","CONTEXT_URI","CONTEXT_URL"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

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

<a id="canonical-0133312033223022-3221102200122133-1000233323323322-0203020003022221-0020100303103322-1030013112331330-1220330322112113-2020110331101303"></a>

<a id="canonical-0230131300230032-0212312111031233-0021302221202200-0011112212213020-3122303222203111-0202200113122322-1102000113112301-2020123212003101"></a>

## context_name property — exclude_violation_contexts / 230022111122 / 5

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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

<a id="canonical-1123302001223300-1030001203030132-0230131130211131-0213313310133332-3222223223302223-2121002333022230-1122223121302312-3121332132110301"></a>

<a id="canonical-1313311021312123-3323220323312311-0321022130220113-1331332311321011-1213222110133203-3010002101302230-2021221111000230-3101121202230130"></a>

## exclude_violation property — exclude_violation_contexts / 230022111122 / 6

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VIOL_ASM_COOKIE_MODIFIED","VIOL_COOKIE_MALFORMED","VIOL_COOKIE_MODIFIED","VIOL_DATA_GUARD","VIOL_ENCODING","VIOL_EVASION_APACHE_WHITESPACE","VIOL_EVASION_BAD_UNESCAPE","VIOL_EVASION_BARE_BYTE_DECODING","VIOL_EVASION_DIRECTORY_TRAVERSALS","VIOL_EVASION_IIS_BACKSLASHES","VIOL_EVASION_IIS_UNICODE_CODEPOINTS","VIOL_EVASION_MULTIPLE_DECODING","VIOL_EVASION_PERCENT_U_DECODING","VIOL_FILETYPE","VIOL_FILE_UPLOAD","VIOL_FILE_UPLOAD_IN_BODY","VIOL_GRAPHQL_FORMAT","VIOL_GRAPHQL_INTROSPECTION_QUERY","VIOL_GRAPHQL_MALFORMED","VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE","VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION","VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST","VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS","VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST","VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS","VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT","VIOL_HTTP_RESPONSE_STATUS","VIOL_JSON_MALFORMED","VIOL_MALFORMED_REQUEST","VIOL_MANDATORY_HEADER","VIOL_METHOD","VIOL_NONE","VIOL_REQUEST_MAX_LENGTH","VIOL_XML_MALFORMED"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("VIOL_NONE",
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
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
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
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"),
}
```

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

<a id="canonical-3322232333023313-3322103010202032-2022213012330322-3123020321313213-2211232030003120-2212023320201203-0121233203110013-0233002102131212"></a>

## Next pages — exclude_violation_contexts / 230022111122 / 7

- [waf_exclusion_rules.app_firewall_detection_control](resources--waf_exclusion_policy--reference--group-001.md#canonical-3033331200333021-0211200202303333-3133032223121003-1130303000300131-2233330202331212-1210230101113303-0211123033312221-0120132322022313)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)

<a id="canonical-2311323130333222-3113330103001312-3020301011221112-3103322023021012-1303023301123211-2131321231201031-2030103213222113-0302001231101103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200301200113031-0013120211100032-3301200310221323-3233113211320321-0123133220233123-1111031301112212-2131121002203331-3113110123122023"></a>

## waf_exclusion_rules.metadata — metadata / 111333101312 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-3201012220022020-0302300332132202-1210300000101133-0132320021112000-0303020230312320-2033302130202231-1302212000313220-0123110313003311)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012)
- waf_exclusion_rules.metadata

<a id="canonical-3222311333123220-2323130011012232-1022013201000031-1231122013322201-3211101332010132-0212110232332001-3211111210330331-3133000201230112"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

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

<a id="canonical-2112231310232100-1213011220000303-2021113320120111-1312232003210021-3110011213231302-1310010302133300-1222103201123321-2230220333321003"></a>

## Direct properties — metadata / 111333101312 / 3

<a id="canonical-2310232312021203-2211320102111002-0100210202013023-1323333103311111-2122033000122331-0101301032013200-2013311030232130-0133121320112031"></a>

<a id="canonical-2201022201332200-0202132202100021-1130323201311220-1212021231303203-3200312311321222-1003100212111101-3321303133320233-0330222310101032"></a>

## description_spec property — metadata / 111333101312 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2200010132323313-2112223221302021-3110203111211101-0231302233011120-2113112033223100-0113023020211230-2002210220312331-2312123123313101"></a>

<a id="canonical-3001211321220120-0100100232021111-3202303120021112-1002302210320012-0223003012332013-1312310030101212-1133231213122233-3221031012311300"></a>

## name property — metadata / 111333101312 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

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

<a id="canonical-3211112320030211-2222220210001032-3323122123300213-0301001113203011-1302131310302102-3210202020310312-1121201201113111-0213020022010310"></a>

## Next pages — metadata / 111333101312 / 6

- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)

<a id="canonical-1211112303030110-1100130231023332-3201000020101201-3013232112230302-1210330310233010-0233333332101112-2213132333012113-2213220303222123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000232102200233-2030222032333301-0332331321100113-2320100310323112-0021232000323221-1203010323203311-0000112110030123-2220111102032233"></a>

## waf_exclusion_rules.waf_skip_processing — waf_skip_processing / 002202021203 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)
- [Property reference](resources--waf_exclusion_policy--reference--group-001.md#canonical-3201012220022020-0302300332132202-1210300000101133-0132320021112000-0303020230312320-2033302130202231-1302212000313220-0123110313003311)
- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012)
- waf_exclusion_rules.waf_skip_processing

<a id="canonical-1111223003123022-3231322113112213-0321013203322223-2320322022002211-2301310312330021-2333032303320033-3032010110203221-1223132121320221"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
waf_skip_processing = {}
```

<a id="canonical-3220311202232232-0130221231323200-2102030323320133-3211131010300231-3000230211223322-3323232021121131-0201130322333300-2003220212233221"></a>

## Direct properties — waf_skip_processing / 002202021203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201222102101100-0332201013023223-0030220130213032-3201131202033003-1331030121021323-2310113031332322-1012011201011233-0011122000130020"></a>

## Next pages — waf_skip_processing / 002202021203 / 4

- [waf_exclusion_rules](resources--waf_exclusion_policy--reference--group-001.md#canonical-0322320101202230-2231013321001031-0201113210112101-3002102211332203-0013201213213233-2211232131201233-1132103130123231-2100101032201012)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)
