---
page_title: "xcsh_malicious_user_mitigation reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_malicious_user_mitigation reference."
---

# xcsh_malicious_user_mitigation reference

<a id="canonical-1110001002333302-2300113210021312-0300220333020331-3110232013111231-2133223211023113-1021322132102302-1212032231131231-3211212320213101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210222022123301-1131102011223303-1021233010100211-0221010113011221-0233030202322300-0330001023032300-0232211000102100-2300301232323203"></a>

## Property reference — Property reference / 120111001303 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- Property reference

<a id="canonical-2003121213221210-3102222001221332-3323212111213320-3113120220311110-1330111100110233-2030223101210103-2301220012122032-3303023123323202"></a>

## Direct properties — Property reference / 120111001303 / 3

<a id="canonical-3023022010123102-3112031023202033-0332111102203233-3202232202323200-2021321013222101-2002312333211122-0112200220133300-2011112020132000"></a>

<a id="canonical-0032332022133233-1300330322023312-2233010030101220-0311311001323332-3131220331313131-2320331122230220-0123330311023032-0101220033012212"></a>

## annotations property — Property reference / 120111001303 / 4

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

<a id="canonical-1011102333031332-0121311121100312-1013212201122020-2302111100231301-2231020320133221-3302013100120112-2133231231100020-3200130302110303"></a>

<a id="canonical-2101010331030223-3003333200123030-3003302033000303-2222032031033212-3031002203310330-0203210023220212-3232031330020231-2011011021121033"></a>

## description property — Property reference / 120111001303 / 5

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

<a id="canonical-3233320011112020-1323212220122112-1203122033102030-3303200031313332-1101121030201103-1230132010220201-2022013203033312-3131202221311100"></a>

<a id="canonical-0222133231112213-3003231030231132-2013230003212323-1301130203120021-1221023122230130-2220103220232021-1010121333333021-3122330322210301"></a>

## disable property — Property reference / 120111001303 / 6

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

<a id="canonical-2202231020123100-2300101321223113-3112311033211203-0131310223022121-2222320313232210-1312312023102133-0312110232223033-0002103233200321"></a>

<a id="canonical-0213301102121332-3101130331211221-1220230302132313-0033232011222002-2322102331203320-1033311312213123-3233331220210210-0321331113323330"></a>

## ID property — Property reference / 120111001303 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2203303331212211-1001102322120221-1010020023000223-2303002222110111-2021110103112200-2111231232312132-3023002301012012-3030300203230032"></a>

<a id="canonical-2112333121330210-0112301101110333-2202010122330103-0302302321301022-2231232203003210-1233211302112210-0122112020302230-0313230202120230"></a>

## labels property — Property reference / 120111001303 / 8

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

- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-0120200313002001-2220011100300123-2211210022323122-3031231321312001-2101212221230203-1133011303122233-1121112112303313-3011200201333230): complete subsection reference.

<a id="canonical-3011031332131312-3000320012323132-2032111031300103-2123330301232110-2002102031011311-1102310130233323-1210130110002002-0233311132300130"></a>

<a id="canonical-3031331111302103-0011212130201233-1122102220310023-3200012120121323-0310331333023002-0021312020111003-1031123212102201-0210230333300010"></a>

## name property — Property reference / 120111001303 / 9

Type: `"string"`. Required.

Name of the Malicious User Mitigation. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1010111022201310-3332012221030131-0313023000002123-0322332023023012-1113201131032012-0100312030012310-0202111133233213-0301311130032010"></a>

<a id="canonical-1301202230121311-0032131200200001-2133200100322030-0233302013033030-2032221130011202-3320133110321001-0222031130101212-2003303032110033"></a>

## namespace property — Property reference / 120111001303 / 10

Type: `"string"`. Required.

Namespace where the Malicious User Mitigation is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
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

- [timeouts](resources--malicious_user_mitigation--reference--group-001.md#canonical-2111112231103223-3302110213200023-0210302320312310-2312101001023031-2031202320032003-3310223203030203-0311101322111032-2021032200130112): complete subsection reference.

<a id="canonical-0302011131302303-0113013012332203-1302203323321213-0020131010230323-0030133000021333-3101002313331032-1130313121202210-2330122002020111"></a>

## All schema paths — Property reference / 120111001303 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--malicious_user_mitigation--reference--group-001.md#canonical-3023022010123102-3112031023202033-0332111102203233-3202232202323200-2021321013222101-2002312333211122-0112200220133300-2011112020132000) |
| `description` | [description](resources--malicious_user_mitigation--reference--group-001.md#canonical-1011102333031332-0121311121100312-1013212201122020-2302111100231301-2231020320133221-3302013100120112-2133231231100020-3200130302110303) |
| `disable` | [disable](resources--malicious_user_mitigation--reference--group-001.md#canonical-3233320011112020-1323212220122112-1203122033102030-3303200031313332-1101121030201103-1230132010220201-2022013203033312-3131202221311100) |
| `id` | [id](resources--malicious_user_mitigation--reference--group-001.md#canonical-2202231020123100-2300101321223113-3112311033211203-0131310223022121-2222320313232210-1312312023102133-0312110232223033-0002103233200321) |
| `labels` | [labels](resources--malicious_user_mitigation--reference--group-001.md#canonical-2203303331212211-1001102322120221-1010020023000223-2303002222110111-2021110103112200-2111231232312132-3023002301012012-3030300203230032) |
| `mitigation_type` | [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-0120112022322100-0322213112003022-1020200202221132-3000121213033330-2331122301202221-2121202030222232-3212230223122311-3110230321201201) |
| `mitigation_type.rules` | [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-1302310011231120-1313311020310223-3011023010213013-2011302202013131-0203230002232210-0200020011202101-2131321120220233-0301232200133322) |
| `mitigation_type.rules.mitigation_action` | [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-3132301103222330-1212112331230312-0030102320103201-1300023013133232-2230313210130203-2003130212020212-0011332102032102-2310232133021023) |
| `mitigation_type.rules.mitigation_action.block_temporarily` | [mitigation_type.rules.mitigation_action.block_temporarily](resources--malicious_user_mitigation--reference--group-001.md#canonical-0113000222002302-3301330302232222-1023231310313002-1131232030231120-0003233101221213-0102022031030011-3203110320233300-2212300023023330) |
| `mitigation_type.rules.mitigation_action.captcha_challenge` | [mitigation_type.rules.mitigation_action.captcha_challenge](resources--malicious_user_mitigation--reference--group-001.md#canonical-0033012101123221-0310230331333101-0123301000001100-0330210131102122-3122333023311200-3101020212001310-2313231011200011-0232032323003022) |
| `mitigation_type.rules.mitigation_action.javascript_challenge` | [mitigation_type.rules.mitigation_action.javascript_challenge](resources--malicious_user_mitigation--reference--group-001.md#canonical-1101100322022312-2113012333103100-2333200010133202-1121022330001220-3322213123221122-2003231303223212-2322323023320000-0220120200122322) |
| `mitigation_type.rules.threat_level` | [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-0323131130213120-0000202320023202-3102110310322203-0320222021212211-3313323021120220-3321311201013112-0333200320333313-2023311230211230) |
| `mitigation_type.rules.threat_level.high` | [mitigation_type.rules.threat_level.high](resources--malicious_user_mitigation--reference--group-001.md#canonical-2101330131011321-1212311033201203-1202210001121230-1100221202033120-0300130013231221-3212222311202232-0110331112232310-3122122310231031) |
| `mitigation_type.rules.threat_level.low` | [mitigation_type.rules.threat_level.low](resources--malicious_user_mitigation--reference--group-001.md#canonical-3210202122221201-3221213313031303-1230300131021323-3300033130130311-0032033002230001-2311020100313232-3232330132130000-1000101210010302) |
| `mitigation_type.rules.threat_level.medium` | [mitigation_type.rules.threat_level.medium](resources--malicious_user_mitigation--reference--group-001.md#canonical-3212101012232030-3113011201322200-0201311312333202-1222031013122321-1003231300103110-0222201032023033-2203001302103320-3220100232202320) |
| `name` | [name](resources--malicious_user_mitigation--reference--group-001.md#canonical-3011031332131312-3000320012323132-2032111031300103-2123330301232110-2002102031011311-1102310130233323-1210130110002002-0233311132300130) |
| `namespace` | [namespace](resources--malicious_user_mitigation--reference--group-001.md#canonical-1010111022201310-3332012221030131-0313023000002123-0322332023023012-1113201131032012-0100312030012310-0202111133233213-0301311130032010) |
| `timeouts` | [timeouts](resources--malicious_user_mitigation--reference--group-001.md#canonical-3213312030200131-1112033000203322-1303013333033313-0131303322011300-3101331111121102-1331331310310203-0222221023311323-0211003030303112) |
| `timeouts.create` | [timeouts.create](resources--malicious_user_mitigation--reference--group-001.md#canonical-2300311120113112-2123121031312323-2010022221223201-1210113012010321-3321122201111310-1132332020311302-3310122231200011-2232012220111211) |
| `timeouts.delete` | [timeouts.delete](resources--malicious_user_mitigation--reference--group-001.md#canonical-3213220323103223-0320202210331210-2132101130131003-2311211011020023-1302003232121100-2111300003132033-1112212002131000-3210220201010221) |
| `timeouts.read` | [timeouts.read](resources--malicious_user_mitigation--reference--group-001.md#canonical-1332110120031020-0021300310302231-3033021123203112-2101312312223103-2021120020223210-0222033223212320-2111113300233203-2003011102133222) |
| `timeouts.update` | [timeouts.update](resources--malicious_user_mitigation--reference--group-001.md#canonical-0203022210120231-2212311232000330-0322222131023311-3102123201000103-2001331312323021-3322320110332030-3211203022022323-3333001300312120) |

<a id="canonical-3333210013021033-3212300001000300-0300202011000223-0033003212202200-1023223100102331-1112221001203022-2111231113131203-1331330113032321"></a>

## Next pages — Property reference / 120111001303 / 12

- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-0120200313002001-2220011100300123-2211210022323122-3031231321312001-2101212221230203-1133011303122233-1121112112303313-3011200201333230)
- [timeouts](resources--malicious_user_mitigation--reference--group-001.md#canonical-2111112231103223-3302110213200023-0210302320312310-2312101001023031-2031202320032003-3310223203030203-0311101322111032-2021032200130112)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)

<a id="canonical-0120200313002001-2220011100300123-2211210022323122-3031231321312001-2101212221230203-1133011303122233-1121112112303313-3011200201333230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113311022303123-3103310111333012-3220003311021232-1230330100212021-1303213311321212-0011312331321001-0301132032000111-0113321322300322"></a>

## mitigation_type — mitigation_type / 202032200002 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-1110001002333302-2300113210021312-0300220333020331-3110232013111231-2133223211023113-1021322132102302-1212032231131231-3211212320213101)
- mitigation_type

<a id="canonical-0120112022322100-0322213112003022-1020200202221132-3000121213033330-2331122301202221-2121202030222232-3212230223122311-3110230321201201"></a>

Type: `"object"`. single nested block, Optional.

Settings that specify the actions to be taken when malicious users are determined to be at different
threat levels. User's activity is monitored and continuously analyzed for malicious behavior. From
this analysis, a threat-level is assigned to each user. Server applies default when omitted.

Upstream description:

Settings that specify the actions to be taken when malicious users are determined to be at different
threat levels. User's activity is monitored and continuously analyzed for malicious behavior. From
this analysis, a threat-level is assigned to each user. The settings defined in malicious user
mitigation specify what mitigation actions to take for user determined to be at different threat
levels.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
mitigation_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323330203321201-0101001002203013-0212002121123012-0130013121033031-0011001033330331-2330000220303131-3010223102222111-2220020002033133"></a>

## Direct properties — mitigation_type / 202032200002 / 3

- [rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-2203232003033103-0203333122231231-1322022310231230-0112032113103131-2230132122210032-0000332021330012-2222332212022102-1333332133331303): complete subsection reference.

<a id="canonical-1112211013023210-1301021120323201-1213120232211232-3123021300302322-3232321131112020-3221010010031303-2223033032222123-1213301330130210"></a>

## Next pages — mitigation_type / 202032200002 / 4

- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-2203232003033103-0203333122231231-1322022310231230-0112032113103131-2230132122210032-0000332021330012-2222332212022102-1333332133331303)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-1110001002333302-2300113210021312-0300220333020331-3110232013111231-2133223211023113-1021322132102302-1212032231131231-3211212320213101)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)

<a id="canonical-2203232003033103-0203333122231231-1322022310231230-0112032113103131-2230132122210032-0000332021330012-2222332212022102-1333332133331303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001223222132122-2310030013030313-1210332013110311-0001133111033111-2002013032230323-2221012322230000-0110032033210321-1113021011210013"></a>

## mitigation_type.rules — rules / 033110012220 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-1110001002333302-2300113210021312-0300220333020331-3110232013111231-2133223211023113-1021322132102302-1212032231131231-3211212320213101)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-0120200313002001-2220011100300123-2211210022323122-3031231321312001-2101212221230203-1133011303122233-1121112112303313-3011200201333230)
- mitigation_type.rules

<a id="canonical-1302310011231120-1313311020310223-3011023010213013-2011302202013131-0203230002232210-0200020011202101-2131321120220233-0301232200133322"></a>

Type: `"object"`. list nested block, Optional.

Define the threat levels and the corresponding mitigation actions to be taken.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true",
    "ves.io.schema.rules.repeated.unique_threat_level": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true",
    "ves.io.schema.rules.repeated.unique_threat_level": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023133132210000-1101133102331001-3012132312233301-0033121002032100-1200211011130113-0313110221123011-3210332030112011-3101131123011202"></a>

## Direct properties — rules / 033110012220 / 3

- [mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-1210132312032130-2021211012101311-2121023301132323-0220200312030001-1302213233333000-0202022130200033-3300201122103120-0222231120011323): complete subsection reference.

- [threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-3331212021222023-0103201222303231-1032210101232212-3220023023001111-1221021033133130-3032030322013121-0302133310013322-3030121313211021): complete subsection reference.

<a id="canonical-1012001032330101-2312311211021112-1323221303302113-1003202023023123-2003032002121021-2022223212211323-2033110131111132-2310331212012221"></a>

## Next pages — rules / 033110012220 / 4

- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-1210132312032130-2021211012101311-2121023301132323-0220200312030001-1302213233333000-0202022130200033-3300201122103120-0222231120011323)
- [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-3331212021222023-0103201222303231-1032210101232212-3220023023001111-1221021033133130-3032030322013121-0302133310013322-3030121313211021)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-0120200313002001-2220011100300123-2211210022323122-3031231321312001-2101212221230203-1133011303122233-1121112112303313-3011200201333230)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)

<a id="canonical-1210132312032130-2021211012101311-2121023301132323-0220200312030001-1302213233333000-0202022130200033-3300201122103120-0222231120011323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321203033320011-3131120120330023-2300020120213121-1200332032201210-3333300032322211-2210231013130302-3222020232302001-1303112033223111"></a>

## mitigation_type.rules.mitigation_action — mitigation_action / 232133133031 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-1110001002333302-2300113210021312-0300220333020331-3110232013111231-2133223211023113-1021322132102302-1212032231131231-3211212320213101)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-0120200313002001-2220011100300123-2211210022323122-3031231321312001-2101212221230203-1133011303122233-1121112112303313-3011200201333230)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-2203232003033103-0203333122231231-1322022310231230-0112032113103131-2230132122210032-0000332021330012-2222332212022102-1333332133331303)
- mitigation_type.rules.mitigation_action

<a id="canonical-3132301103222330-1212112331230312-0030102320103201-1300023013133232-2230313210130203-2003130212020212-0011332102032102-2310232133021023"></a>

Type: `"object"`. single nested block, Optional.

Supported actions that can be taken to mitigate malicious activity from a user.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block_temporarily",
    "captcha_challenge"),
  validators.ConflictingObjectAttributes("block_temporarily",
    "javascript_challenge"),
  validators.ConflictingObjectAttributes("captcha_challenge",
    "javascript_challenge")}
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
  "x-ves-oneof-field-mitigation_action": "[\"block_temporarily\",\"captcha_challenge\",\"javascript_challenge\"]"
}
```

Terraform syntax:

```terraform
mitigation_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131013001233220-3111030221131312-2313012213313332-0103233032330223-2101202020333322-2013332130020301-2223232321302113-3021201131002033"></a>

## Direct properties — mitigation_action / 232133133031 / 3

- [block_temporarily](resources--malicious_user_mitigation--reference--group-001.md#canonical-1112200131021313-1313332202001231-0211021013100110-0303320231320220-1203330010122310-3030030010133010-3310110210033231-0102120232312111): complete subsection reference.

- [captcha_challenge](resources--malicious_user_mitigation--reference--group-001.md#canonical-3133010200312323-2110201210333230-1201032313103030-3320230112030131-1010022120220322-0330110220201022-1221310233221320-3330121210131121): complete subsection reference.

- [javascript_challenge](resources--malicious_user_mitigation--reference--group-001.md#canonical-3132003212133220-0231332030021231-3220101320311012-3100110221203231-1232233130110023-2111333121001131-3020130133133301-0023321330103123): complete subsection reference.

<a id="canonical-0222110013002221-2313301132130222-2330120100003111-2132031033303002-2111201031131003-2123122133302112-3121103131102021-3121303111002223"></a>

## Next pages — mitigation_action / 232133133031 / 4

- [mitigation_type.rules.mitigation_action.block_temporarily](resources--malicious_user_mitigation--reference--group-001.md#canonical-1112200131021313-1313332202001231-0211021013100110-0303320231320220-1203330010122310-3030030010133010-3310110210033231-0102120232312111)
- [mitigation_type.rules.mitigation_action.captcha_challenge](resources--malicious_user_mitigation--reference--group-001.md#canonical-3133010200312323-2110201210333230-1201032313103030-3320230112030131-1010022120220322-0330110220201022-1221310233221320-3330121210131121)
- [mitigation_type.rules.mitigation_action.javascript_challenge](resources--malicious_user_mitigation--reference--group-001.md#canonical-3132003212133220-0231332030021231-3220101320311012-3100110221203231-1232233130110023-2111333121001131-3020130133133301-0023321330103123)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-2203232003033103-0203333122231231-1322022310231230-0112032113103131-2230132122210032-0000332021330012-2222332212022102-1333332133331303)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)

<a id="canonical-1112200131021313-1313332202001231-0211021013100110-0303320231320220-1203330010122310-3030030010133010-3310110210033231-0102120232312111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323001302200022-2321310013323011-1000031211233220-0122120302233331-2220033100022001-3002320313102110-2102231013211020-2100202100022003"></a>

## mitigation_type.rules.mitigation_action.block_temporarily — block_temporarily / 202122112323 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-1110001002333302-2300113210021312-0300220333020331-3110232013111231-2133223211023113-1021322132102302-1212032231131231-3211212320213101)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-0120200313002001-2220011100300123-2211210022323122-3031231321312001-2101212221230203-1133011303122233-1121112112303313-3011200201333230)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-2203232003033103-0203333122231231-1322022310231230-0112032113103131-2230132122210032-0000332021330012-2222332212022102-1333332133331303)
- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-1210132312032130-2021211012101311-2121023301132323-0220200312030001-1302213233333000-0202022130200033-3300201122103120-0222231120011323)
- mitigation_type.rules.mitigation_action.block_temporarily

<a id="canonical-0113000222002302-3301330302232222-1023231310313002-1131232030231120-0003233101221213-0102022031030011-3203110320233300-2212300023023330"></a>

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
block_temporarily = {}
```

<a id="canonical-3200133323003220-2223001010330323-3211001100213303-0213202000102300-2123101032313222-0020013002013122-1102333123311022-0231103333022121"></a>

## Direct properties — block_temporarily / 202122112323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110000020112101-0123122221210111-1330133030123100-0022113030110213-0223313333101330-3221312220333121-2321232231021203-0210012122202022"></a>

## Next pages — block_temporarily / 202122112323 / 4

- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-1210132312032130-2021211012101311-2121023301132323-0220200312030001-1302213233333000-0202022130200033-3300201122103120-0222231120011323)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)

<a id="canonical-3133010200312323-2110201210333230-1201032313103030-3320230112030131-1010022120220322-0330110220201022-1221310233221320-3330121210131121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032131010113302-2312233211101312-3313312131001300-3023121130221121-2332333130030233-0202131222201322-1303022003213031-2022301223002212"></a>

## mitigation_type.rules.mitigation_action.captcha_challenge — captcha_challenge / 222230233022 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-1110001002333302-2300113210021312-0300220333020331-3110232013111231-2133223211023113-1021322132102302-1212032231131231-3211212320213101)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-0120200313002001-2220011100300123-2211210022323122-3031231321312001-2101212221230203-1133011303122233-1121112112303313-3011200201333230)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-2203232003033103-0203333122231231-1322022310231230-0112032113103131-2230132122210032-0000332021330012-2222332212022102-1333332133331303)
- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-1210132312032130-2021211012101311-2121023301132323-0220200312030001-1302213233333000-0202022130200033-3300201122103120-0222231120011323)
- mitigation_type.rules.mitigation_action.captcha_challenge

<a id="canonical-0033012101123221-0310230331333101-0123301000001100-0330210131102122-3122333023311200-3101020212001310-2313231011200011-0232032323003022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for captcha challenge.

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
captcha_challenge = {}
```

<a id="canonical-3210120200013002-3222001233332131-3131021333321011-2003220111211100-2112331300221020-3321132112113000-3311311113333123-1203003330331230"></a>

## Direct properties — captcha_challenge / 222230233022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320302020033131-0200121121321202-1021103323022122-1011220302020230-3132311032231212-2301121231001333-2301022131130030-0032223320122233"></a>

## Next pages — captcha_challenge / 222230233022 / 4

- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-1210132312032130-2021211012101311-2121023301132323-0220200312030001-1302213233333000-0202022130200033-3300201122103120-0222231120011323)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)

<a id="canonical-3132003212133220-0231332030021231-3220101320311012-3100110221203231-1232233130110023-2111333121001131-3020130133133301-0023321330103123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110112320123231-1122003313020223-2301122322203320-3011022211110222-3220201000210201-0303110101210202-1200310202132121-0213313203333312"></a>

## mitigation_type.rules.mitigation_action.javascript_challenge — javascript_challenge / 021101001130 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-1110001002333302-2300113210021312-0300220333020331-3110232013111231-2133223211023113-1021322132102302-1212032231131231-3211212320213101)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-0120200313002001-2220011100300123-2211210022323122-3031231321312001-2101212221230203-1133011303122233-1121112112303313-3011200201333230)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-2203232003033103-0203333122231231-1322022310231230-0112032113103131-2230132122210032-0000332021330012-2222332212022102-1333332133331303)
- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-1210132312032130-2021211012101311-2121023301132323-0220200312030001-1302213233333000-0202022130200033-3300201122103120-0222231120011323)
- mitigation_type.rules.mitigation_action.javascript_challenge

<a id="canonical-1101100322022312-2113012333103100-2333200010133202-1121022330001220-3322213123221122-2003231303223212-2322323023320000-0220120200122322"></a>

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
javascript_challenge = {}
```

<a id="canonical-0322233133210122-1311023031300203-2321203023203120-3112212320220313-2010223321223012-1110301223220320-2231131322102020-3003332001300230"></a>

## Direct properties — javascript_challenge / 021101001130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2312201321113032-1300133231122201-1310321213311101-2220013300103120-0003300213000300-3311201201332202-3123133001220321-3201312100223120"></a>

## Next pages — javascript_challenge / 021101001130 / 4

- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--reference--group-001.md#canonical-1210132312032130-2021211012101311-2121023301132323-0220200312030001-1302213233333000-0202022130200033-3300201122103120-0222231120011323)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)

<a id="canonical-3331212021222023-0103201222303231-1032210101232212-3220023023001111-1221021033133130-3032030322013121-0302133310013322-3030121313211021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133322231213033-3211130100012001-0022023303012202-3102212201130223-0121220313120031-1202012112233010-3211023103000201-0301303320202320"></a>

## mitigation_type.rules.threat_level — threat_level / 221120222022 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-1110001002333302-2300113210021312-0300220333020331-3110232013111231-2133223211023113-1021322132102302-1212032231131231-3211212320213101)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-0120200313002001-2220011100300123-2211210022323122-3031231321312001-2101212221230203-1133011303122233-1121112112303313-3011200201333230)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-2203232003033103-0203333122231231-1322022310231230-0112032113103131-2230132122210032-0000332021330012-2222332212022102-1333332133331303)
- mitigation_type.rules.threat_level

<a id="canonical-0323131130213120-0000202320023202-3102110310322203-0320222021212211-3313323021120220-3321311201013112-0333200320333313-2023311230211230"></a>

Type: `"object"`. single nested block, Optional.

Threat level estimated for each user based on the user's activity and reputation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("high",
    "low"),
  validators.ConflictingObjectAttributes("high",
    "medium"),
  validators.ConflictingObjectAttributes("low",
    "medium")}
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
  "x-ves-oneof-field-threat_level": "[\"high\",\"low\",\"medium\"]"
}
```

Terraform syntax:

```terraform
threat_level {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211213101322132-0001221312310200-2200300322103313-0112032211020201-3222001132012211-3233102320230221-3123331333131132-3000112111323330"></a>

## Direct properties — threat_level / 221120222022 / 3

- [high](resources--malicious_user_mitigation--reference--group-001.md#canonical-0100333213231221-1212123231222232-3300213123021331-3231220222100010-1000030133330202-2322230311001001-3233110213020132-3201012211231103): complete subsection reference.

- [low](resources--malicious_user_mitigation--reference--group-001.md#canonical-3012112202333023-2122323122221310-0222011233102321-1010020120022130-0031131121123210-0001301012312323-3133210120302100-0132120031211012): complete subsection reference.

- [medium](resources--malicious_user_mitigation--reference--group-001.md#canonical-0031123301312110-0031012210220021-1310100331233200-2013301013211302-1301103313203033-0023200130110330-2313010200221133-1000113330320202): complete subsection reference.

<a id="canonical-3022300002020300-0020031131001200-2003300200322030-3132223311000320-1011022031210213-1103302232310120-1003011110220221-1122232300200030"></a>

## Next pages — threat_level / 221120222022 / 4

- [mitigation_type.rules.threat_level.high](resources--malicious_user_mitigation--reference--group-001.md#canonical-0100333213231221-1212123231222232-3300213123021331-3231220222100010-1000030133330202-2322230311001001-3233110213020132-3201012211231103)
- [mitigation_type.rules.threat_level.low](resources--malicious_user_mitigation--reference--group-001.md#canonical-3012112202333023-2122323122221310-0222011233102321-1010020120022130-0031131121123210-0001301012312323-3133210120302100-0132120031211012)
- [mitigation_type.rules.threat_level.medium](resources--malicious_user_mitigation--reference--group-001.md#canonical-0031123301312110-0031012210220021-1310100331233200-2013301013211302-1301103313203033-0023200130110330-2313010200221133-1000113330320202)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-2203232003033103-0203333122231231-1322022310231230-0112032113103131-2230132122210032-0000332021330012-2222332212022102-1333332133331303)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)

<a id="canonical-0100333213231221-1212123231222232-3300213123021331-3231220222100010-1000030133330202-2322230311001001-3233110213020132-3201012211231103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132230113213313-1100120032132322-0231111132132221-3120221030203113-2132332332000021-0032312130211130-3303201122300000-2110210021001023"></a>

## mitigation_type.rules.threat_level.high — high / 030301331220 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-1110001002333302-2300113210021312-0300220333020331-3110232013111231-2133223211023113-1021322132102302-1212032231131231-3211212320213101)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-0120200313002001-2220011100300123-2211210022323122-3031231321312001-2101212221230203-1133011303122233-1121112112303313-3011200201333230)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-2203232003033103-0203333122231231-1322022310231230-0112032113103131-2230132122210032-0000332021330012-2222332212022102-1333332133331303)
- [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-3331212021222023-0103201222303231-1032210101232212-3220023023001111-1221021033133130-3032030322013121-0302133310013322-3030121313211021)
- mitigation_type.rules.threat_level.high

<a id="canonical-2101330131011321-1212311033201203-1202210001121230-1100221202033120-0300130013231221-3212222311202232-0110331112232310-3122122310231031"></a>

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
high = {}
```

<a id="canonical-1132112212011113-2131202330101000-0000332303133221-2113010201002323-3221210102132233-3132201223200132-0101100120313233-2221131001110010"></a>

## Direct properties — high / 030301331220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131133310122222-2303003200022000-0120003323230202-0200323003002011-1102100110202100-2112331302122120-2213322101302322-3200121102103102"></a>

## Next pages — high / 030301331220 / 4

- [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-3331212021222023-0103201222303231-1032210101232212-3220023023001111-1221021033133130-3032030322013121-0302133310013322-3030121313211021)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)

<a id="canonical-3012112202333023-2122323122221310-0222011233102321-1010020120022130-0031131121123210-0001301012312323-3133210120302100-0132120031211012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022010223221131-3222310221313230-0120202022100220-1122032031110312-3103020010010022-1003031111322002-3232222232103120-1002220330221130"></a>

## mitigation_type.rules.threat_level.low — low / 102012132231 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-1110001002333302-2300113210021312-0300220333020331-3110232013111231-2133223211023113-1021322132102302-1212032231131231-3211212320213101)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-0120200313002001-2220011100300123-2211210022323122-3031231321312001-2101212221230203-1133011303122233-1121112112303313-3011200201333230)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-2203232003033103-0203333122231231-1322022310231230-0112032113103131-2230132122210032-0000332021330012-2222332212022102-1333332133331303)
- [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-3331212021222023-0103201222303231-1032210101232212-3220023023001111-1221021033133130-3032030322013121-0302133310013322-3030121313211021)
- mitigation_type.rules.threat_level.low

<a id="canonical-3210202122221201-3221213313031303-1230300131021323-3300033130130311-0032033002230001-2311020100313232-3232330132130000-1000101210010302"></a>

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
low = {}
```

<a id="canonical-1000330303003323-2311033033101200-3220231020201133-1001010121000032-1200031222032032-3123101031203321-3021102103122020-3213202201122113"></a>

## Direct properties — low / 102012132231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202323103122302-2331300323101012-0233203002331311-3033310200031101-1130123033132010-3020330001300312-3320133200023311-2123010131331303"></a>

## Next pages — low / 102012132231 / 4

- [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-3331212021222023-0103201222303231-1032210101232212-3220023023001111-1221021033133130-3032030322013121-0302133310013322-3030121313211021)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)

<a id="canonical-0031123301312110-0031012210220021-1310100331233200-2013301013211302-1301103313203033-0023200130110330-2313010200221133-1000113330320202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100203102331231-3303020221230103-2202132231120222-0022302331231223-3300220211221131-1313130130233213-0203121011010123-0333333120230302"></a>

## mitigation_type.rules.threat_level.medium — medium / 111220002003 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-1110001002333302-2300113210021312-0300220333020331-3110232013111231-2133223211023113-1021322132102302-1212032231131231-3211212320213101)
- [mitigation_type](resources--malicious_user_mitigation--reference--group-001.md#canonical-0120200313002001-2220011100300123-2211210022323122-3031231321312001-2101212221230203-1133011303122233-1121112112303313-3011200201333230)
- [mitigation_type.rules](resources--malicious_user_mitigation--reference--group-001.md#canonical-2203232003033103-0203333122231231-1322022310231230-0112032113103131-2230132122210032-0000332021330012-2222332212022102-1333332133331303)
- [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-3331212021222023-0103201222303231-1032210101232212-3220023023001111-1221021033133130-3032030322013121-0302133310013322-3030121313211021)
- mitigation_type.rules.threat_level.medium

<a id="canonical-3212101012232030-3113011201322200-0201311312333202-1222031013122321-1003231300103110-0222201032023033-2203001302103320-3220100232202320"></a>

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
medium = {}
```

<a id="canonical-0202231132130231-3011302112031233-2130301022012120-3122332210101200-2103120300103110-1123222201333023-3232200103122110-1330011133022233"></a>

## Direct properties — medium / 111220002003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311322100003322-3322332021301122-0313112010020123-0321213220302133-2013313022330033-1312312210230231-3321323220000302-2021313331010300"></a>

## Next pages — medium / 111220002003 / 4

- [mitigation_type.rules.threat_level](resources--malicious_user_mitigation--reference--group-001.md#canonical-3331212021222023-0103201222303231-1032210101232212-3220023023001111-1221021033133130-3032030322013121-0302133310013322-3030121313211021)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)

<a id="canonical-2111112231103223-3302110213200023-0210302320312310-2312101001023031-2031202320032003-3310223203030203-0311101322111032-2021032200130112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032210100212333-0330021120303210-0101022123212113-1133320330123222-3203102202030112-2000112100222031-2013202211001123-1000120330111121"></a>

## timeouts — timeouts / 301312022022 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-1110001002333302-2300113210021312-0300220333020331-3110232013111231-2133223211023113-1021322132102302-1212032231131231-3211212320213101)
- timeouts

<a id="canonical-3213312030200131-1112033000203322-1303013333033313-0131303322011300-3101331111121102-1331331310310203-0222221023311323-0211003030303112"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311331100310323-0222112100032323-2333012223102001-3203000200010222-1333011013121233-2023023113033030-2000011201002132-1322031130132213"></a>

## Direct properties — timeouts / 301312022022 / 3

<a id="canonical-2300311120113112-2123121031312323-2010022221223201-1210113012010321-3321122201111310-1132332020311302-3310122231200011-2232012220111211"></a>

<a id="canonical-1302103231121222-1230101301110032-3132203103031323-0323010230011232-3021321003202022-3220320300101212-1313031331011222-1130121121021023"></a>

## create property — timeouts / 301312022022 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3213220323103223-0320202210331210-2132101130131003-2311211011020023-1302003232121100-2111300003132033-1112212002131000-3210220201010221"></a>

<a id="canonical-1021320213301312-3020303133333300-3212033101300312-1013330330232033-0233323133311223-1132322133300011-1332200221322022-3113230011102323"></a>

## delete property — timeouts / 301312022022 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1332110120031020-0021300310302231-3033021123203112-2101312312223103-2021120020223210-0222033223212320-2111113300233203-2003011102133222"></a>

<a id="canonical-2320003222013123-1212111033012111-1113001230110223-0101110223130031-2233103101131311-1002311120221331-0210003002010320-3303113120312133"></a>

## read property — timeouts / 301312022022 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0203022210120231-2212311232000330-0322222131023311-3102123201000103-2001331312323021-3322320110332030-3211203022022323-3333001300312120"></a>

<a id="canonical-1200103100030012-2220023021132131-0032100301023333-3010323033311232-3222132031313121-0220301103003101-2230222233130313-1332330112212300"></a>

## update property — timeouts / 301312022022 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3302321101023302-2200121020121111-3112330233121201-3321313120032021-2020320001312032-0002023030313330-0323011230222222-0221011333101230"></a>

## Next pages — timeouts / 301312022022 / 8

- [Property reference](resources--malicious_user_mitigation--reference--group-001.md#canonical-1110001002333302-2300113210021312-0300220333020331-3110232013111231-2133223211023113-1021322132102302-1212032231131231-3211212320213101)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
