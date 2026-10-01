---
page_title: "xcsh_fast_acl_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl_rule reference."
---

# xcsh_fast_acl_rule reference

<a id="canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130321201223112-0232101021101102-3110122212010231-1203032011322022-3200012331101221-0123232013103011-2122022010033220-1122212002121133"></a>

## Property reference — Property reference / 233113233113 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- Property reference

<a id="canonical-3021121212122122-1000302123132023-2000312231112313-2230233111123231-3202223010210311-3102000312233233-0023031320100211-1023131211221202"></a>

## Direct properties — Property reference / 233113233113 / 3

- [action](resources--fast_acl_rule--reference--group-001.md#canonical-2323132233031011-0321032123110200-3322222321020213-0222230201011121-2112122222031112-1110200012230031-0103000130300210-2301310300233232): complete subsection reference.

<a id="canonical-0113221031200012-1202301231132131-1130220033031212-0231113330023213-3212000023313100-2032311313312021-1033123331202132-2032100102210332"></a>

<a id="canonical-3200030202302331-3323210223233113-0100230013033332-2112320100131132-2232300110013001-1022111113112330-1022221023330021-3300100201223222"></a>

## annotations property — Property reference / 233113233113 / 4

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

<a id="canonical-3020020303211333-0221323131001223-0000121010112103-1023303321203322-1200231200100312-0211001131020320-0111320120231303-2211223121201013"></a>

<a id="canonical-2330303201102311-3123002223132010-1332030202020132-1021320201333011-2113001333100103-1033103121231221-2202001220211103-0330220012013102"></a>

## description property — Property reference / 233113233113 / 5

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

<a id="canonical-1230123102233312-2331323101311003-2221220301030130-3320223300332330-0312311332031122-3331231022212312-2102232101303222-2112023002003333"></a>

<a id="canonical-2233231132323100-3230022220321212-0332031100011102-2112230301012310-3021003211130032-0030301333113302-2122332032111321-1201133230330101"></a>

## disable property — Property reference / 233113233113 / 6

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

<a id="canonical-2223312003031111-0132220013112010-3022323120013300-3011201323003212-0232213310032222-0100023200121200-0320321010133311-3230132333331031"></a>

<a id="canonical-3211222333131111-2131333221131122-2322101211211212-3110312013232313-1210300132012112-2131203102102222-1220332310100012-3212223130031210"></a>

## ID property — Property reference / 233113233113 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_prefix_set](resources--fast_acl_rule--reference--group-001.md#canonical-0313300332123332-0030203111330313-3121323032102000-0301210002212311-2030021003133210-1230233230030123-3221202110201133-2030302010102232): complete subsection reference.

<a id="canonical-3323203030002121-3132332311111111-0012332213033321-2132001203010323-0013333221010133-0122202001110222-2003030311133002-3231031203021203"></a>

<a id="canonical-3221233302130213-0100311100131013-2113313200123032-3331113221232002-3031211202212001-1212232311031131-2000132232303300-2003132203331111"></a>

## labels property — Property reference / 233113233113 / 8

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

<a id="canonical-1313311202301313-0121212113210012-0101031333332010-2111113011002231-1212003131003202-1223013332033223-1122231122111111-3132330003213302"></a>

<a id="canonical-0112233121230223-0113332010300103-2210322101303303-2201021303023002-1302200110002331-0300102200212013-3030012221011303-0201022110132301"></a>

## name property — Property reference / 233113233113 / 9

Type: `"string"`. Required.

Name of the Fast ACL Rule. Must be unique within the namespace.

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

<a id="canonical-1213220010311213-2230311201331303-0032030020120001-2121103301022320-1003202312233331-2010332221212223-2200031310321313-0203220232101131"></a>

<a id="canonical-0312213100213132-0122023012203302-3233222020010333-2031333202230333-3233020001011202-0313130111201031-0123131023201010-0003222123323102"></a>

## namespace property — Property reference / 233113233113 / 10

Type: `"string"`. Required.

Namespace where the Fast ACL Rule is created.

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

- [port](resources--fast_acl_rule--reference--group-001.md#canonical-1032003000303102-1033202210203210-3222020302212013-0321300133103321-3312031111132110-3202311220103302-2233113211303003-3331013000330130): complete subsection reference.

- [prefix](resources--fast_acl_rule--reference--group-001.md#canonical-0313132020210223-3030132113131010-0123233323202310-2312303132222323-1203330311231211-3212223222232203-1303200300003303-2200112333000313): complete subsection reference.

- [timeouts](resources--fast_acl_rule--reference--group-001.md#canonical-3100132302101032-0231233113011020-2113300231211030-0113331001112220-2233213301003123-2112012222201133-2123110313013103-2022223302030201): complete subsection reference.

<a id="canonical-0212213003203220-3020311112021013-1211100132131001-3132103031010111-1111222132302332-0130322000003320-1223122020333310-1011303212301000"></a>

## All schema paths — Property reference / 233113233113 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](resources--fast_acl_rule--reference--group-001.md#canonical-1332003121200110-0200311202012002-1113032233023111-3313131123332212-1221100202132211-1022220221321123-0331210313121120-0230131233202221) |
| `action.policer_action` | [action.policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-2320112213211132-1133303313301300-0013201101031320-3330322322232131-1203202301322303-0230201112213122-3222320120110021-0213030331232020) |
| `action.policer_action.ref` | [action.policer_action.ref](resources--fast_acl_rule--reference--group-001.md#canonical-0311231102221330-2200312233230133-0231102310303000-0332001323231122-1321330121323130-2110122023113302-2131101000003131-3000122210121013) |
| `action.policer_action.ref.kind` | [action.policer_action.ref.kind](resources--fast_acl_rule--reference--group-001.md#canonical-2122203023031231-1031100333333102-3130112001012033-0013201010231200-2301002033000033-2201012120033133-0131233201332130-0012010130213102) |
| `action.policer_action.ref.name` | [action.policer_action.ref.name](resources--fast_acl_rule--reference--group-001.md#canonical-3333132321201311-2020321331221030-1113300121301302-1121200111012013-2212301001321331-2231331033110020-2310013210023232-1112130101110320) |
| `action.policer_action.ref.namespace` | [action.policer_action.ref.namespace](resources--fast_acl_rule--reference--group-001.md#canonical-2310231220201010-0212110033202132-1031220333330033-1013103223011311-1212333202321232-2121310130113111-3322302030033120-1032033331021000) |
| `action.policer_action.ref.tenant` | [action.policer_action.ref.tenant](resources--fast_acl_rule--reference--group-001.md#canonical-3002132333001311-0131012213123200-1110211132301101-2021302210020102-3022330013021212-0212021301100012-3103121230000102-3033212011102212) |
| `action.policer_action.ref.uid` | [action.policer_action.ref.uid](resources--fast_acl_rule--reference--group-001.md#canonical-2300233332021233-3303001000031101-2120301033201132-3212332330210313-2033000321300001-2331213333233321-1130023012013331-3320213302331030) |
| `action.protocol_policer_action` | [action.protocol_policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-3200311321001222-2110102012310333-1020232113103130-1112232022013000-3323123031321112-0111330120111010-1110220200001232-0131201030303330) |
| `action.protocol_policer_action.ref` | [action.protocol_policer_action.ref](resources--fast_acl_rule--reference--group-001.md#canonical-2310232220113133-1202003002003230-2222132233121120-1020311020131130-2312011111110213-1211330113123123-1011210103301230-0011121033011321) |
| `action.protocol_policer_action.ref.kind` | [action.protocol_policer_action.ref.kind](resources--fast_acl_rule--reference--group-001.md#canonical-1232112130232332-1210201103221102-1111023203332103-2033310021000301-2100112022301332-0102233020320113-1001022212321011-0031100103020120) |
| `action.protocol_policer_action.ref.name` | [action.protocol_policer_action.ref.name](resources--fast_acl_rule--reference--group-001.md#canonical-1100222002320333-1121302311200303-3202003330012230-0332220103223022-0132331230331311-2210020110310321-2210123210221013-2120033112110203) |
| `action.protocol_policer_action.ref.namespace` | [action.protocol_policer_action.ref.namespace](resources--fast_acl_rule--reference--group-001.md#canonical-0123232223212031-0031222221003301-1323223322021103-3012001002201320-3313111210102002-1321020203321021-3320022230033111-0022000302112333) |
| `action.protocol_policer_action.ref.tenant` | [action.protocol_policer_action.ref.tenant](resources--fast_acl_rule--reference--group-001.md#canonical-2212132131020321-2002320011331223-0332211111120122-3202012330120202-2210212302320311-0330313331033333-0013101212321100-0302323210010131) |
| `action.protocol_policer_action.ref.uid` | [action.protocol_policer_action.ref.uid](resources--fast_acl_rule--reference--group-001.md#canonical-0012131111323212-3001031123132010-2022232212210021-2003021131313020-0202002313321121-1023202020123213-3131320303030323-3002113013212232) |
| `action.simple_action` | [action.simple_action](resources--fast_acl_rule--reference--group-001.md#canonical-1301031000221133-2031130030311320-3311022211211101-0211203002003203-0103111110222013-3013010020131330-2001211311131112-1102131111112211) |
| `annotations` | [annotations](resources--fast_acl_rule--reference--group-001.md#canonical-0113221031200012-1202301231132131-1130220033031212-0231113330023213-3212000023313100-2032311313312021-1033123331202132-2032100102210332) |
| `description` | [description](resources--fast_acl_rule--reference--group-001.md#canonical-3020020303211333-0221323131001223-0000121010112103-1023303321203322-1200231200100312-0211001131020320-0111320120231303-2211223121201013) |
| `disable` | [disable](resources--fast_acl_rule--reference--group-001.md#canonical-1230123102233312-2331323101311003-2221220301030130-3320223300332330-0312311332031122-3331231022212312-2102232101303222-2112023002003333) |
| `id` | [ID](resources--fast_acl_rule--reference--group-001.md#canonical-2223312003031111-0132220013112010-3022323120013300-3011201323003212-0232213310032222-0100023200121200-0320321010133311-3230132333331031) |
| `ip_prefix_set` | [ip_prefix_set](resources--fast_acl_rule--reference--group-001.md#canonical-1320112323231321-2121002110321111-0233021012110202-2220030202011211-1112223021101131-1030323031232111-3021223101223331-3022110302210120) |
| `ip_prefix_set.ref` | [ip_prefix_set.ref](resources--fast_acl_rule--reference--group-001.md#canonical-3112101300321133-1200000002220023-1122333112030221-0200132013112300-1231013101332030-0220230320111323-1000102022300100-0023312212000321) |
| `ip_prefix_set.ref.kind` | [ip_prefix_set.ref.kind](resources--fast_acl_rule--reference--group-001.md#canonical-3313231112233233-1121130120330103-2131300001030003-1130031303001322-0233103212211023-0321202123002212-0232112021002333-0002212100300102) |
| `ip_prefix_set.ref.name` | [ip_prefix_set.ref.name](resources--fast_acl_rule--reference--group-001.md#canonical-3223203213020200-1021330232220111-2102232122133120-3233301230113320-2311131012212300-3013133321003001-1120111122001231-1121120301101302) |
| `ip_prefix_set.ref.namespace` | [ip_prefix_set.ref.namespace](resources--fast_acl_rule--reference--group-001.md#canonical-3030201130322322-2033332122103002-1103120011123323-2101300003130130-1330332121120222-3220012303033333-2003021133303231-0102100023201001) |
| `ip_prefix_set.ref.tenant` | [ip_prefix_set.ref.tenant](resources--fast_acl_rule--reference--group-001.md#canonical-0213301331220121-1020333033201102-2312022302201112-1012033010010232-1221011203101231-2123312223122130-0203010210111032-3032011131113211) |
| `ip_prefix_set.ref.uid` | [ip_prefix_set.ref.uid](resources--fast_acl_rule--reference--group-001.md#canonical-2233210120332330-1032131213113312-2200113230023300-3130332200033202-1232213031233120-1000231133031213-1111323022311202-1021230201123311) |
| `labels` | [labels](resources--fast_acl_rule--reference--group-001.md#canonical-3323203030002121-3132332311111111-0012332213033321-2132001203010323-0013333221010133-0122202001110222-2003030311133002-3231031203021203) |
| `name` | [name](resources--fast_acl_rule--reference--group-001.md#canonical-1313311202301313-0121212113210012-0101031333332010-2111113011002231-1212003131003202-1223013332033223-1122231122111111-3132330003213302) |
| `namespace` | [namespace](resources--fast_acl_rule--reference--group-001.md#canonical-1213220010311213-2230311201331303-0032030020120001-2121103301022320-1003202312233331-2010332221212223-2200031310321313-0203220232101131) |
| `port` | [port](resources--fast_acl_rule--reference--group-001.md#canonical-1313032123001021-1103223311313212-2331130300201220-2203011000332020-0102231322112133-3021310203101000-1112123013023201-2022023322101020) |
| `port.all` | [port.all](resources--fast_acl_rule--reference--group-001.md#canonical-2020132012123031-1220032122300120-2123000232101310-1012211003222100-0203030133033111-2032033131012203-3220232012200311-2322322332112012) |
| `port.dns` | [port.dns](resources--fast_acl_rule--reference--group-001.md#canonical-1103202110132303-1202020222111000-0103132330112102-2332001000103202-0332131002323022-3102000310333020-1012312203332103-1022120311330320) |
| `port.user_defined` | [port.user_defined](resources--fast_acl_rule--reference--group-001.md#canonical-2320201230030033-2212123101313022-0012200112232230-3030321330122321-2233311332213211-3031301031012003-1330112003003032-0313212322333121) |
| `prefix` | [prefix](resources--fast_acl_rule--reference--group-001.md#canonical-2102023111201032-1020201013120222-0200213223122302-1321112313032131-1232022301020120-2000122331213102-1321033020231223-1013000020122322) |
| `prefix.prefix` | [prefix.prefix](resources--fast_acl_rule--reference--group-001.md#canonical-3310212211111330-2003112000211132-2001022120231333-0113221321312022-2103200030203233-0131130201310123-0133300232333002-3211330100221121) |
| `timeouts` | [timeouts](resources--fast_acl_rule--reference--group-001.md#canonical-2023333201122001-1203023103001130-1321231333032111-2231333013222111-0123323011023002-3023022100111213-1101233113223212-2120201312002320) |
| `timeouts.create` | [timeouts.create](resources--fast_acl_rule--reference--group-001.md#canonical-2322120320123002-1220330330003030-0223330103001200-1220123232021303-3222112001132321-2000003302113032-0102311022100022-1300210103301302) |
| `timeouts.delete` | [timeouts.delete](resources--fast_acl_rule--reference--group-001.md#canonical-2203012020032322-0321010312110233-1311211013110233-3222331002002220-2211111000110011-0030023222333100-3032213020110220-1121110132333221) |
| `timeouts.read` | [timeouts.read](resources--fast_acl_rule--reference--group-001.md#canonical-0103022100003233-1211202002331230-3100222222320123-0123302103100301-1301313012301312-2333000222212131-1203201132133330-2223302210010313) |
| `timeouts.update` | [timeouts.update](resources--fast_acl_rule--reference--group-001.md#canonical-2333110030210331-3022301322212322-1321301123021232-1202231213331231-3333212231011331-1301233200222103-2023022133030000-1332010220330213) |

<a id="canonical-1111200300233112-0023023303322320-0220033113202310-1200002311321003-3103131233223203-2313123220113031-3130032303012121-2101102032102212"></a>

## Next pages — Property reference / 233113233113 / 12

- [action](resources--fast_acl_rule--reference--group-001.md#canonical-2323132233031011-0321032123110200-3322222321020213-0222230201011121-2112122222031112-1110200012230031-0103000130300210-2301310300233232)
- [ip_prefix_set](resources--fast_acl_rule--reference--group-001.md#canonical-0313300332123332-0030203111330313-3121323032102000-0301210002212311-2030021003133210-1230233230030123-3221202110201133-2030302010102232)
- [port](resources--fast_acl_rule--reference--group-001.md#canonical-1032003000303102-1033202210203210-3222020302212013-0321300133103321-3312031111132110-3202311220103302-2233113211303003-3331013000330130)
- [prefix](resources--fast_acl_rule--reference--group-001.md#canonical-0313132020210223-3030132113131010-0123233323202310-2312303132222323-1203330311231211-3212223222232203-1303200300003303-2200112333000313)
- [timeouts](resources--fast_acl_rule--reference--group-001.md#canonical-3100132302101032-0231233113011020-2113300231211030-0113331001112220-2233213301003123-2112012222201133-2123110313013103-2022223302030201)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)

<a id="canonical-2323132233031011-0321032123110200-3322222321020213-0222230201011121-2112122222031112-1110200012230031-0103000130300210-2301310300233232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123302020201221-0122210221200330-0011220131323311-0130033102021111-1213032202033201-1000031033023033-1110123313110320-2233320013233220"></a>

## action — action / 013001003021 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- action

<a id="canonical-1332003121200110-0200311202012002-1113032233023111-3313131123332212-1221100202132211-1022220221321123-0331210313121120-0230131233202221"></a>

Type: `"object"`. single nested block, Optional.

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Upstream description:

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("policer_action",
    "protocol_policer_action"),
  validators.ConflictingObjectAttributes("policer_action",
    "simple_action"),
  validators.ConflictingObjectAttributes("protocol_policer_action",
    "simple_action")}
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
  "x-ves-oneof-field-action": "[\"policer_action\",\"protocol_policer_action\",\"simple_action\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332200233010121-1332111021012330-2123321210100212-3323323010231021-2302202200201013-0312012023232012-0031332113221232-2333000003013330"></a>

## Direct properties — action / 013001003021 / 3

- [policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-1220210013232110-1103010033322231-2200331031121032-1010021122312012-0112233031130010-2331101033211202-1022320330321321-0221101310331003): complete subsection reference.

- [protocol_policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-0032023332120030-1021231301112230-2223021031203310-2321231220233200-0221103002122230-3331100012103221-1230100032311323-2310012320223022): complete subsection reference.

<a id="canonical-1301031000221133-2031130030311320-3311022211211101-0211203002003203-0103111110222013-3013010020131330-2001211311131112-1102131111112211"></a>

<a id="canonical-3220322323223202-2122103032311112-1212302013300333-1101000010101300-2022011020323130-2132113330211302-1103010111220003-2332112010003113"></a>

## simple_action property — action / 013001003021 / 4

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

FastAclRuleSimpleAction specifies simple action like PASS or DENY

Drop the traffic Forward the traffic.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3031021230330301-2201203330000311-0301122223332120-0220302112111330-0131233013233002-1032333323002331-3323020332111023-0011122102120230"></a>

## Next pages — action / 013001003021 / 5

- [action.policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-1220210013232110-1103010033322231-2200331031121032-1010021122312012-0112233031130010-2331101033211202-1022320330321321-0221101310331003)
- [action.protocol_policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-0032023332120030-1021231301112230-2223021031203310-2321231220233200-0221103002122230-3331100012103221-1230100032311323-2310012320223022)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)

<a id="canonical-1220210013232110-1103010033322231-2200331031121032-1010021122312012-0112233031130010-2331101033211202-1022320330321321-0221101310331003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001220120031113-1012332211132021-2130230123303211-3321131211102123-2331032020021020-3300332212121132-3331033230312033-2302302322001332"></a>

## action.policer_action — policer_action / 222023211030 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [action](resources--fast_acl_rule--reference--group-001.md#canonical-2323132233031011-0321032123110200-3322222321020213-0222230201011121-2112122222031112-1110200012230031-0103000130300210-2301310300233232)
- action.policer_action

<a id="canonical-2320112213211132-1133303313301300-0013201101031320-3330322322232131-1203202301322303-0230201112213122-3222320120110021-0213030331232020"></a>

Type: `"object"`. single nested block, Optional.

Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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
policer_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230013010123011-2023011101333313-0030200323022100-0203002103331001-0023012200022122-3323202322003222-1303011211023003-3200023033002212"></a>

## Direct properties — policer_action / 222023211030 / 3

- [ref](resources--fast_acl_rule--reference--group-001.md#canonical-2302001012022000-2303112203202222-3013211332132120-2322210113311201-1022312311001010-3202311330001201-0121100002223231-0030112310132320): complete subsection reference.

<a id="canonical-3023113133022233-0111123013131001-3230312031030021-2200202312101013-2321123113031202-2122220030213202-1300120033113103-3222221101023301"></a>

## Next pages — policer_action / 222023211030 / 4

- [action.policer_action.ref](resources--fast_acl_rule--reference--group-001.md#canonical-2302001012022000-2303112203202222-3013211332132120-2322210113311201-1022312311001010-3202311330001201-0121100002223231-0030112310132320)
- [action](resources--fast_acl_rule--reference--group-001.md#canonical-2323132233031011-0321032123110200-3322222321020213-0222230201011121-2112122222031112-1110200012230031-0103000130300210-2301310300233232)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)

<a id="canonical-2302001012022000-2303112203202222-3013211332132120-2322210113311201-1022312311001010-3202311330001201-0121100002223231-0030112310132320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222003000133330-1212322000012201-0000302011021010-3303312333220003-0332012203111222-1031201003301330-3021113003130131-0132111012101222"></a>

## action.policer_action.ref — ref / 221223112222 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [action](resources--fast_acl_rule--reference--group-001.md#canonical-2323132233031011-0321032123110200-3322222321020213-0222230201011121-2112122222031112-1110200012230031-0103000130300210-2301310300233232)
- [action.policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-1220210013232110-1103010033322231-2200331031121032-1010021122312012-0112233031130010-2331101033211202-1022320330321321-0221101310331003)
- action.policer_action.ref

<a id="canonical-0311231102221330-2200312233230133-0231102310303000-0332001323231122-1321330121323130-2110122023113302-2131101000003131-3000122210121013"></a>

Type: `"object"`. list nested block, Optional.

Reference. A policer direct reference.

Upstream description:

A policer direct reference.

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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-1101132222122101-0030333021320110-1121100321311233-3023201210133123-1131231310333001-1101302200301313-0210031020312131-0332333313303202"></a>

## Direct properties — ref / 221223112222 / 3

<a id="canonical-2122203023031231-1031100333333102-3130112001012033-0013201010231200-2301002033000033-2201012120033133-0131233201332130-0012010130213102"></a>

<a id="canonical-2220032000222032-2102232211103022-1013111301031032-0203233231303012-0222122111121033-1201001222002310-2010130002202233-1003121032323002"></a>

## kind property — ref / 221223112222 / 4

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

<a id="canonical-3333132321201311-2020321331221030-1113300121301302-1121200111012013-2212301001321331-2231331033110020-2310013210023232-1112130101110320"></a>

<a id="canonical-1000023003120133-2022022230133211-0311012001132221-1010311301012010-1030212331130020-0013333301020030-3112223223302302-3201001332222220"></a>

## name property — ref / 221223112222 / 5

Type: `"string"`. Optional.

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

<a id="canonical-2310231220201010-0212110033202132-1031220333330033-1013103223011311-1212333202321232-2121310130113111-3322302030033120-1032033331021000"></a>

<a id="canonical-3100133113212301-0023111110320112-2030202020013302-3313230102001133-1301222110313021-2323030222312130-0302133223033233-2110301020122200"></a>

## namespace property — ref / 221223112222 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3002132333001311-0131012213123200-1110211132301101-2021302210020102-3022330013021212-0212021301100012-3103121230000102-3033212011102212"></a>

<a id="canonical-3003100222130101-1002333202022321-1131223112132313-2011022203033120-2320101030211203-0031313122211313-1223332001121021-0330112100011310"></a>

## tenant property — ref / 221223112222 / 7

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

<a id="canonical-2300233332021233-3303001000031101-2120301033201132-3212332330210313-2033000321300001-2331213333233321-1130023012013331-3320213302331030"></a>

<a id="canonical-2002231020003222-0322321333213220-3320231332013312-0332233333103233-1203133132131121-2211110332223013-2232021020222210-1323302331222031"></a>

## uid property — ref / 221223112222 / 8

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

<a id="canonical-0331300123102030-3030201000220002-0121120333300031-3132012331213310-2123111303321122-1233313122131210-2322212122201003-1132310011132130"></a>

## Next pages — ref / 221223112222 / 9

- [action.policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-1220210013232110-1103010033322231-2200331031121032-1010021122312012-0112233031130010-2331101033211202-1022320330321321-0221101310331003)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)

<a id="canonical-0032023332120030-1021231301112230-2223021031203310-2321231220233200-0221103002122230-3331100012103221-1230100032311323-2310012320223022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112121312231122-1223301001113101-2213201300000123-0123121033313200-0331000132012202-0323133303331332-0222012320121000-3220320130223210"></a>

## action.protocol_policer_action — protocol_policer_action / 122100022003 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [action](resources--fast_acl_rule--reference--group-001.md#canonical-2323132233031011-0321032123110200-3322222321020213-0222230201011121-2112122222031112-1110200012230031-0103000130300210-2301310300233232)
- action.protocol_policer_action

<a id="canonical-3200311321001222-2110102012310333-1020232113103130-1112232022013000-3323123031321112-0111330120111010-1110220200001232-0131201030303330"></a>

Type: `"object"`. single nested block, Optional.

Protocol Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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
protocol_policer_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313200210003320-0321011111010000-1120130312323203-1212101010032211-3021203032102000-0202300100302303-0310123132231023-1221230320101221"></a>

## Direct properties — protocol_policer_action / 122100022003 / 3

- [ref](resources--fast_acl_rule--reference--group-001.md#canonical-3210323231330213-1130110002032031-0222001321332110-0130201210000031-2122213332202100-1002103320213103-0030002211000120-2002230112113012): complete subsection reference.

<a id="canonical-2110232233003322-0103001033033230-1201012122300303-3012123121310333-0303221211033331-0113323012013213-2021201212002122-0112001033002230"></a>

## Next pages — protocol_policer_action / 122100022003 / 4

- [action.protocol_policer_action.ref](resources--fast_acl_rule--reference--group-001.md#canonical-3210323231330213-1130110002032031-0222001321332110-0130201210000031-2122213332202100-1002103320213103-0030002211000120-2002230112113012)
- [action](resources--fast_acl_rule--reference--group-001.md#canonical-2323132233031011-0321032123110200-3322222321020213-0222230201011121-2112122222031112-1110200012230031-0103000130300210-2301310300233232)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)

<a id="canonical-3210323231330213-1130110002032031-0222001321332110-0130201210000031-2122213332202100-1002103320213103-0030002211000120-2002230112113012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220201200330202-1200300321201002-1231002013203122-3201122123133212-3231312222303220-0130320010312301-0101301120112021-0220021200013012"></a>

## action.protocol_policer_action.ref — ref / 202233313331 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [action](resources--fast_acl_rule--reference--group-001.md#canonical-2323132233031011-0321032123110200-3322222321020213-0222230201011121-2112122222031112-1110200012230031-0103000130300210-2301310300233232)
- [action.protocol_policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-0032023332120030-1021231301112230-2223021031203310-2321231220233200-0221103002122230-3331100012103221-1230100032311323-2310012320223022)
- action.protocol_policer_action.ref

<a id="canonical-2310232220113133-1202003002003230-2222132233121120-1020311020131130-2312011111110213-1211330113123123-1011210103301230-0011121033011321"></a>

Type: `"object"`. list nested block, Optional.

Protocol policer Reference. Reference to protocol policer object.

Upstream description:

Reference to protocol policer object.

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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3021012013233021-1000223311030331-3313231302010320-3122023111131202-3030222010123321-1321212313100203-0312222230120221-2033303211102313"></a>

## Direct properties — ref / 202233313331 / 3

<a id="canonical-1232112130232332-1210201103221102-1111023203332103-2033310021000301-2100112022301332-0102233020320113-1001022212321011-0031100103020120"></a>

<a id="canonical-3310221002103122-2312301322311110-1131111222332102-0100221201320302-2333111212331223-2131202221121130-3322312130232022-2310302300123222"></a>

## kind property — ref / 202233313331 / 4

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

<a id="canonical-1100222002320333-1121302311200303-3202003330012230-0332220103223022-0132331230331311-2210020110310321-2210123210221013-2120033112110203"></a>

<a id="canonical-0230011320020222-1021023301211010-2100021313002003-1212023300223213-1313320011021033-2112112030023002-3111011331312213-1110102212231223"></a>

## name property — ref / 202233313331 / 5

Type: `"string"`. Optional.

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

<a id="canonical-0123232223212031-0031222221003301-1323223322021103-3012001002201320-3313111210102002-1321020203321021-3320022230033111-0022000302112333"></a>

<a id="canonical-3121333213010131-0123301003223201-2122112301230133-0302001310100213-3231201201000202-3333320021222212-2222321021302020-0111212333231331"></a>

## namespace property — ref / 202233313331 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2212132131020321-2002320011331223-0332211111120122-3202012330120202-2210212302320311-0330313331033333-0013101212321100-0302323210010131"></a>

<a id="canonical-3212102330312021-0123310331202322-2110201210012332-2132100301003010-1323333110222123-2202302003101220-2121123110323111-2233103321012100"></a>

## tenant property — ref / 202233313331 / 7

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

<a id="canonical-0012131111323212-3001031123132010-2022232212210021-2003021131313020-0202002313321121-1023202020123213-3131320303030323-3002113013212232"></a>

<a id="canonical-0301223120021111-3100121232233303-2203311111021021-3132003023023022-1231001232120331-1200123303033033-0032033101303330-2323021020022110"></a>

## uid property — ref / 202233313331 / 8

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

<a id="canonical-0323110033103230-2301203101110121-2301222213000100-2130312013032221-1311230200333011-3313013112210222-3111302113323202-2312031101212212"></a>

## Next pages — ref / 202233313331 / 9

- [action.protocol_policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-0032023332120030-1021231301112230-2223021031203310-2321231220233200-0221103002122230-3331100012103221-1230100032311323-2310012320223022)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)

<a id="canonical-0313300332123332-0030203111330313-3121323032102000-0301210002212311-2030021003133210-1230233230030123-3221202110201133-2030302010102232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301113231120212-1130221203131112-2301001310210123-0012133332001310-0003131320213312-3131001133201100-2112301023212230-0313120100122133"></a>

## ip_prefix_set — ip_prefix_set / 230322010101 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- ip_prefix_set

<a id="canonical-1320112323231321-2121002110321111-0233021012110202-2220030202011211-1112223021101131-1030323031232111-3021223101223331-3022110302210120"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ip\_prefix\_set, prefix\] List of references to ip\_prefix\_set objects.

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

- [ip_prefix_set](resources--fast_acl_rule--reference--group-001.md#canonical-1320112323231321-2121002110321111-0233021012110202-2220030202011211-1112223021101131-1030323031232111-3021223101223331-3022110302210120)
- [prefix](resources--fast_acl_rule--reference--group-001.md#canonical-2102023111201032-1020201013120222-0200213223122302-1321112313032131-1232022301020120-2000122331213102-1321033020231223-1013000020122322)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011323100100303-2013230011011101-0211311201033012-0330230222101100-1100011233133022-3131312102103030-3123331300120223-2031120330200232"></a>

## Direct properties — ip_prefix_set / 230322010101 / 3

- [ref](resources--fast_acl_rule--reference--group-001.md#canonical-0021312121120211-1213221230010330-2130010031132123-3033012030021231-0022113010211312-3022200211131201-0113111322321131-2313112221210332): complete subsection reference.

<a id="canonical-1200010301111312-3123121123021012-1321221223023201-3010222200330332-2030102222130312-0013031031232203-3201003123310100-1121200011222000"></a>

## Next pages — ip_prefix_set / 230322010101 / 4

- [ip_prefix_set.ref](resources--fast_acl_rule--reference--group-001.md#canonical-0021312121120211-1213221230010330-2130010031132123-3033012030021231-0022113010211312-3022200211131201-0113111322321131-2313112221210332)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)

<a id="canonical-0021312121120211-1213221230010330-2130010031132123-3033012030021231-0022113010211312-3022200211131201-0113111322321131-2313112221210332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122223220113222-1332012310310233-3120322123313002-3002010101112013-0133101333003323-1021022021113202-3123003322011303-0302002103331313"></a>

## ip_prefix_set.ref — ref / 300312222022 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [ip_prefix_set](resources--fast_acl_rule--reference--group-001.md#canonical-0313300332123332-0030203111330313-3121323032102000-0301210002212311-2030021003133210-1230233230030123-3221202110201133-2030302010102232)
- ip_prefix_set.ref

<a id="canonical-3112101300321133-1200000002220023-1122333112030221-0200132013112300-1231013101332030-0220230320111323-1000102022300100-0023312212000321"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021000033231212-0221312321330022-2300023030022312-1102313221130313-1311210111300021-0320112210003220-0213002021111100-1123023110003213"></a>

## Direct properties — ref / 300312222022 / 3

<a id="canonical-3313231112233233-1121130120330103-2131300001030003-1130031303001322-0233103212211023-0321202123002212-0232112021002333-0002212100300102"></a>

<a id="canonical-0312202110200212-0031211332033030-0010021033211030-1320021331213331-2322222233020323-2030221121333011-2322110313220223-0321203213300030"></a>

## kind property — ref / 300312222022 / 4

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

<a id="canonical-3223203213020200-1021330232220111-2102232122133120-3233301230113320-2311131012212300-3013133321003001-1120111122001231-1121120301101302"></a>

<a id="canonical-2000331320131103-2211220302001011-2322231212031133-0030120221130320-1330203120203022-1112011011100012-1130010100112230-2102313303003011"></a>

## name property — ref / 300312222022 / 5

Type: `"string"`. Optional.

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

<a id="canonical-3030201130322322-2033332122103002-1103120011123323-2101300003130130-1330332121120222-3220012303033333-2003021133303231-0102100023201001"></a>

<a id="canonical-2133110011001332-1032213230322220-1202120023011331-0120100320112123-2233123003103330-0222231013222120-2103100232230222-3202021211311323"></a>

## namespace property — ref / 300312222022 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0213301331220121-1020333033201102-2312022302201112-1012033010010232-1221011203101231-2123312223122130-0203010210111032-3032011131113211"></a>

<a id="canonical-3100303310223103-3030131231112021-1321100201010220-0013130102102210-0221031002220022-2130322131023322-2313310333001011-2130321103111321"></a>

## tenant property — ref / 300312222022 / 7

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

<a id="canonical-2233210120332330-1032131213113312-2200113230023300-3130332200033202-1232213031233120-1000231133031213-1111323022311202-1021230201123311"></a>

<a id="canonical-3220021000001203-2021312132332013-1033100132000033-0231333221313011-2032031213120333-3203223022031333-0133132031002312-2101100321233213"></a>

## uid property — ref / 300312222022 / 8

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

<a id="canonical-0122113012221302-1113230203101210-2021022023222213-2102023113320111-0201032332100102-2200001230300003-0311010132332301-2000013101133023"></a>

## Next pages — ref / 300312222022 / 9

- [ip_prefix_set](resources--fast_acl_rule--reference--group-001.md#canonical-0313300332123332-0030203111330313-3121323032102000-0301210002212311-2030021003133210-1230233230030123-3221202110201133-2030302010102232)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)

<a id="canonical-1032003000303102-1033202210203210-3222020302212013-0321300133103321-3312031111132110-3202311220103302-2233113211303003-3331013000330130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213132202232222-2032200311022210-1023010202022000-0020033010100333-3133022313121133-3133132110212130-0010301022130322-3203102221110032"></a>

## port — port / 002022100032 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- port

<a id="canonical-1313032123001021-1103223311313212-2331130300201220-2203011000332020-0102231322112133-3021310203101000-1112123013023201-2022023322101020"></a>

Type: `"object"`. list nested block, Optional.

Source Ports. L4 port numbers to match.

Upstream description:

L4 port numbers to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all",
    "dns"),
  validators.ConflictingListObjectAttributes("all",
    "user_defined"),
  validators.ConflictingListObjectAttributes("dns",
    "user_defined")}
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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112000011233011-3130313203110220-2121133023330301-1110120032230232-1011022330122212-1033101221101122-3202103320012001-2211302023222210"></a>

## Direct properties — port / 002022100032 / 3

- [all](resources--fast_acl_rule--reference--group-001.md#canonical-3322223020322000-0200121011322320-0111201032210033-0031023301033131-1331231031112112-2022122302112031-3001301012202031-3121220113211011): complete subsection reference.

- [DNS](resources--fast_acl_rule--reference--group-001.md#canonical-2321310033213222-1210123211133230-0310201203322011-0113122020022133-0311122321230130-3131223320320133-1311233320003010-3300220333302233): complete subsection reference.

<a id="canonical-2320201230030033-2212123101313022-0012200112232230-3030321330122321-2233311332213211-3031301031012003-1330112003003032-0313212322333121"></a>

<a id="canonical-0033320131213101-2333130300112123-1322320131113122-2010020322113311-1223223132213230-0203003310110310-1012201131021302-2223333120033003"></a>

## user_defined property — port / 002022100032 / 4

Type: `"number"`. Optional.

Exclusive with \[all DNS\] Matches the user defined port.

Upstream description:

Exclusive with \[all DNS\] Matches the user defined port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3032121320103200-3122203131011111-0103113103011030-1000011013023332-2332122322213003-2033113100321312-0231003313122032-2200032301333230"></a>

## Next pages — port / 002022100032 / 5

- [port.all](resources--fast_acl_rule--reference--group-001.md#canonical-3322223020322000-0200121011322320-0111201032210033-0031023301033131-1331231031112112-2022122302112031-3001301012202031-3121220113211011)
- [port.dns](resources--fast_acl_rule--reference--group-001.md#canonical-2321310033213222-1210123211133230-0310201203322011-0113122020022133-0311122321230130-3131223320320133-1311233320003010-3300220333302233)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)

<a id="canonical-3322223020322000-0200121011322320-0111201032210033-0031023301033131-1331231031112112-2022122302112031-3001301012202031-3121220113211011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213003103122130-0312213222303311-3102223012020210-1203002131322321-0123211322202221-0303311321333310-0012311122032003-1211130002102033"></a>

## port.all — all / 200312211021 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [port](resources--fast_acl_rule--reference--group-001.md#canonical-1032003000303102-1033202210203210-3222020302212013-0321300133103321-3312031111132110-3202311220103302-2233113211303003-3331013000330130)
- port.all

<a id="canonical-2020132012123031-1220032122300120-2123000232101310-1012211003222100-0203030133033111-2032033131012203-3220232012200311-2322322332112012"></a>

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
all = {}
```

<a id="canonical-2203321321200123-1012101233311302-1103210102003300-1031211021032032-0130021000111331-2201011210111133-2021222323123123-0111222101131312"></a>

## Direct properties — all / 200312211021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331333230103033-1300213203323002-1320103332031310-0201201100112302-2211001100011301-1313013021012110-3312310202013002-3103203302100300"></a>

## Next pages — all / 200312211021 / 4

- [port](resources--fast_acl_rule--reference--group-001.md#canonical-1032003000303102-1033202210203210-3222020302212013-0321300133103321-3312031111132110-3202311220103302-2233113211303003-3331013000330130)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)

<a id="canonical-2321310033213222-1210123211133230-0310201203322011-0113122020022133-0311122321230130-3131223320320133-1311233320003010-3300220333302233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010121301213011-3021033223211100-3330012330110033-0212202020302233-1202003222212130-3010202120020300-2102232030210332-0220133223212301"></a>

## port.DNS — DNS / 222112131301 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [port](resources--fast_acl_rule--reference--group-001.md#canonical-1032003000303102-1033202210203210-3222020302212013-0321300133103321-3312031111132110-3202311220103302-2233113211303003-3331013000330130)
- port.DNS

<a id="canonical-1103202110132303-1202020222111000-0103132330112102-2332001000103202-0332131002323022-3102000310333020-1012312203332103-1022120311330320"></a>

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
dns = {}
```

<a id="canonical-2222301301301000-2232200223311013-1233233203130312-3310021231123300-1212201300123203-3100320221221002-1013030321310301-0033300020033321"></a>

## Direct properties — DNS / 222112131301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200202313321020-2210211101013122-0032303001230121-1033030130220222-2113133322221001-3013003132231021-0332011111030001-0200223113011101"></a>

## Next pages — DNS / 222112131301 / 4

- [port](resources--fast_acl_rule--reference--group-001.md#canonical-1032003000303102-1033202210203210-3222020302212013-0321300133103321-3312031111132110-3202311220103302-2233113211303003-3331013000330130)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)

<a id="canonical-0313132020210223-3030132113131010-0123233323202310-2312303132222323-1203330311231211-3212223222232203-1303200300003303-2200112333000313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333233030132120-1020323201130221-3101330232132120-0123120101220102-0322002221221231-2312030220132220-2223333012002101-2132233320300321"></a>

## prefix — prefix / 030320203331 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- prefix

<a id="canonical-2102023111201032-1020201013120222-0200213223122302-1321112313032131-1232022301020120-2000122331213102-1321033020231223-1013000020122322"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
prefix {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201223123232213-3330320131211032-3211130113213112-2313231032032333-3120223022322333-2320213330102223-2313003203332302-1210130300310000"></a>

## Direct properties — prefix / 030320203331 / 3

<a id="canonical-3310212211111330-2003112000211132-2001022120231333-0113221321312022-2103200030203233-0131130201310123-0133300232333002-3211330100221121"></a>

<a id="canonical-2233331220031020-3101133332302323-2103303110233110-0200300010030133-3131120030030303-3201023012321010-0100323030323113-3001210121033203"></a>

## prefix property — prefix / 030320203331 / 4

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-1302233120312102-0031033333103021-1002322331032030-3200213211122000-3030221222001002-0203320320303212-0313321021001033-3311123010101122"></a>

## Next pages — prefix / 030320203331 / 5

- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)

<a id="canonical-3100132302101032-0231233113011020-2113300231211030-0113331001112220-2233213301003123-2112012222201133-2123110313013103-2022223302030201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000021233323120-3333201232201122-2110013313312110-1302131101023331-2113210003221110-2121201303021223-0122231313320301-3121011222100201"></a>

## timeouts — timeouts / 110212300330 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- timeouts

<a id="canonical-2023333201122001-1203023103001130-1321231333032111-2231333013222111-0123323011023002-3023022100111213-1101233113223212-2120201312002320"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203302310121022-0010112300313100-0020002020012011-1213030332010213-2320213130203100-2320123030012332-1021130022303321-3333003011000323"></a>

## Direct properties — timeouts / 110212300330 / 3

<a id="canonical-2322120320123002-1220330330003030-0223330103001200-1220123232021303-3222112001132321-2000003302113032-0102311022100022-1300210103301302"></a>

<a id="canonical-0102210303320011-2311132120312201-2132220110230010-1312203221032111-2313301100130010-2020331311203032-2013222001132100-3033230233320113"></a>

## create property — timeouts / 110212300330 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2203012020032322-0321010312110233-1311211013110233-3222331002002220-2211111000110011-0030023222333100-3032213020110220-1121110132333221"></a>

<a id="canonical-3231013332103311-1121220221300032-1320222321011322-3210210233113322-2003102232112203-3311333322233100-2011300023023023-3020220330220333"></a>

## delete property — timeouts / 110212300330 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0103022100003233-1211202002331230-3100222222320123-0123302103100301-1301313012301312-2333000222212131-1203201132133330-2223302210010313"></a>

<a id="canonical-1030310120211301-0302010122103321-3022202302112323-3113213313030322-3302202223113132-2023010332230113-3300323230023333-0213101311101313"></a>

## read property — timeouts / 110212300330 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2333110030210331-3022301322212322-1321301123021232-1202231213331231-3333212231011331-1301233200222103-2023022133030000-1332010220330213"></a>

<a id="canonical-0001133133113303-3303300201030133-1110010201200233-1013122011132202-2112013112122111-2233222230311002-0313300203331032-2200202232321100"></a>

## update property — timeouts / 110212300330 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0210000203010311-0232333222300230-0023132313121323-3200102200120111-3232002121303212-2003212001110113-1101131111020300-0233131332112312"></a>

## Next pages — timeouts / 110212300330 / 8

- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
