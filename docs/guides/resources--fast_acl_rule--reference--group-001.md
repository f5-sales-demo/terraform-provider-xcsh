---
page_title: "xcsh_fast_acl_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl_rule reference."
---

# xcsh_fast_acl_rule reference

<a id="canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- Property reference

<a id="canonical-3130321201223112-0232101021101102-3110122212010231-1203032011322022-3200012331101221-0123232013103011-2122022010033220-1122212002121133"></a>

### Direct properties for `xcsh_fast_acl_rule`

- [action](resources--fast_acl_rule--reference--group-001.md#canonical-2323132233031011-0321032123110200-3322222321020213-0222230201011121-2112122222031112-1110200012230031-0103000130300210-2301310300233232): complete subsection reference.

<a id="canonical-0113221031200012-1202301231132131-1130220033031212-0231113330023213-3212000023313100-2032311313312021-1033123331202132-2032100102210332"></a>

<a id="canonical-3021121212122122-1000302123132023-2000312231112313-2230233111123231-3202223010210311-3102000312233233-0023031320100211-1023131211221202"></a>

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

<a id="canonical-3020020303211333-0221323131001223-0000121010112103-1023303321203322-1200231200100312-0211001131020320-0111320120231303-2211223121201013"></a>

<a id="canonical-3200030202302331-3323210223233113-0100230013033332-2112320100131132-2232300110013001-1022111113112330-1022221023330021-3300100201223222"></a>

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

<a id="canonical-1230123102233312-2331323101311003-2221220301030130-3320223300332330-0312311332031122-3331231022212312-2102232101303222-2112023002003333"></a>

<a id="canonical-2330303201102311-3123002223132010-1332030202020132-1021320201333011-2113001333100103-1033103121231221-2202001220211103-0330220012013102"></a>

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

<a id="canonical-2223312003031111-0132220013112010-3022323120013300-3011201323003212-0232213310032222-0100023200121200-0320321010133311-3230132333331031"></a>

<a id="canonical-2233231132323100-3230022220321212-0332031100011102-2112230301012310-3021003211130032-0030301333113302-2122332032111321-1201133230330101"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_prefix_set](resources--fast_acl_rule--reference--group-001.md#canonical-0313300332123332-0030203111330313-3121323032102000-0301210002212311-2030021003133210-1230233230030123-3221202110201133-2030302010102232): complete subsection reference.

<a id="canonical-3323203030002121-3132332311111111-0012332213033321-2132001203010323-0013333221010133-0122202001110222-2003030311133002-3231031203021203"></a>

<a id="canonical-3211222333131111-2131333221131122-2322101211211212-3110312013232313-1210300132012112-2131203102102222-1220332310100012-3212223130031210"></a>

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

<a id="canonical-1313311202301313-0121212113210012-0101031333332010-2111113011002231-1212003131003202-1223013332033223-1122231122111111-3132330003213302"></a>

<a id="canonical-3221233302130213-0100311100131013-2113313200123032-3331113221232002-3031211202212001-1212232311031131-2000132232303300-2003132203331111"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Fast ACL Rule. Must be unique within the namespace.

Additional upstream details:

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

<a id="canonical-1213220010311213-2230311201331303-0032030020120001-2121103301022320-1003202312233331-2010332221212223-2200031310321313-0203220232101131"></a>

<a id="canonical-0112233121230223-0113332010300103-2210322101303303-2201021303023002-1302200110002331-0300102200212013-3030012221011303-0201022110132301"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Fast ACL Rule is created.

Additional upstream details:

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

- [port](resources--fast_acl_rule--reference--group-001.md#canonical-1032003000303102-1033202210203210-3222020302212013-0321300133103321-3312031111132110-3202311220103302-2233113211303003-3331013000330130): complete subsection reference.

- [prefix](resources--fast_acl_rule--reference--group-001.md#canonical-0313132020210223-3030132113131010-0123233323202310-2312303132222323-1203330311231211-3212223222232203-1303200300003303-2200112333000313): complete subsection reference.

- [timeouts](resources--fast_acl_rule--reference--group-001.md#canonical-3100132302101032-0231233113011020-2113300231211030-0113331001112220-2233213301003123-2112012222201133-2123110313013103-2022223302030201): complete subsection reference.

<a id="canonical-0312213100213132-0122023012203302-3233222020010333-2031333202230333-3233020001011202-0313130111201031-0123131023201010-0003222123323102"></a>

### All schema paths for `xcsh_fast_acl_rule`

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

<a id="canonical-2323132233031011-0321032123110200-3322222321020213-0222230201011121-2112122222031112-1110200012230031-0103000130300210-2301310300233232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `action` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- action

<a id="canonical-1332003121200110-0200311202012002-1113032233023111-3313131123332212-1221100202132211-1022220221321123-0331210313121120-0230131233202221"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-2123302020201221-0122210221200330-0011220131323311-0130033102021111-1213032202033201-1000031033023033-1110123313110320-2233320013233220"></a>

### Direct properties for `action`

- [policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-1220210013232110-1103010033322231-2200331031121032-1010021122312012-0112233031130010-2331101033211202-1022320330321321-0221101310331003): complete subsection reference.

- [protocol_policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-0032023332120030-1021231301112230-2223021031203310-2321231220233200-0221103002122230-3331100012103221-1230100032311323-2310012320223022): complete subsection reference.

<a id="canonical-1301031000221133-2031130030311320-3311022211211101-0211203002003203-0103111110222013-3013010020131330-2001211311131112-1102131111112211"></a>

<a id="canonical-0332200233010121-1332111021012330-2123321210100212-3323323010231021-2302202200201013-0312012023232012-0031332113221232-2333000003013330"></a>

#### `action.simple_action` property

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

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

<a id="canonical-1220210013232110-1103010033322231-2200331031121032-1010021122312012-0112233031130010-2331101033211202-1022320330321321-0221101310331003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `action.policer_action` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [action](resources--fast_acl_rule--reference--group-001.md#canonical-2323132233031011-0321032123110200-3322222321020213-0222230201011121-2112122222031112-1110200012230031-0103000130300210-2301310300233232)
- action.policer_action

<a id="canonical-2320112213211132-1133303313301300-0013201101031320-3330322322232131-1203202301322303-0230201112213122-3222320120110021-0213030331232020"></a>

Type: `"object"`. single nested block, Optional.

Policer Reference. Reference to policer object.

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

<a id="canonical-1001220120031113-1012332211132021-2130230123303211-3321131211102123-2331032020021020-3300332212121132-3331033230312033-2302302322001332"></a>

### Direct properties for `action.policer_action`

- [ref](resources--fast_acl_rule--reference--group-001.md#canonical-2302001012022000-2303112203202222-3013211332132120-2322210113311201-1022312311001010-3202311330001201-0121100002223231-0030112310132320): complete subsection reference.

<a id="canonical-2302001012022000-2303112203202222-3013211332132120-2322210113311201-1022312311001010-3202311330001201-0121100002223231-0030112310132320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `action.policer_action.ref` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [action](resources--fast_acl_rule--reference--group-001.md#canonical-2323132233031011-0321032123110200-3322222321020213-0222230201011121-2112122222031112-1110200012230031-0103000130300210-2301310300233232)
- [action.policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-1220210013232110-1103010033322231-2200331031121032-1010021122312012-0112233031130010-2331101033211202-1022320330321321-0221101310331003)
- action.policer_action.ref

<a id="canonical-0311231102221330-2200312233230133-0231102310303000-0332001323231122-1321330121323130-2110122023113302-2131101000003131-3000122210121013"></a>

Type: `"object"`. list nested block, Optional.

Reference. A policer direct reference.

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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-2222003000133330-1212322000012201-0000302011021010-3303312333220003-0332012203111222-1031201003301330-3021113003130131-0132111012101222"></a>

### Direct properties for `action.policer_action.ref`

<a id="canonical-2122203023031231-1031100333333102-3130112001012033-0013201010231200-2301002033000033-2201012120033133-0131233201332130-0012010130213102"></a>

#### `action.policer_action.ref.kind` property

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

<a id="canonical-3333132321201311-2020321331221030-1113300121301302-1121200111012013-2212301001321331-2231331033110020-2310013210023232-1112130101110320"></a>

<a id="canonical-1101132222122101-0030333021320110-1121100321311233-3023201210133123-1131231310333001-1101302200301313-0210031020312131-0332333313303202"></a>

#### `action.policer_action.ref.name` property

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

<a id="canonical-2310231220201010-0212110033202132-1031220333330033-1013103223011311-1212333202321232-2121310130113111-3322302030033120-1032033331021000"></a>

<a id="canonical-2220032000222032-2102232211103022-1013111301031032-0203233231303012-0222122111121033-1201001222002310-2010130002202233-1003121032323002"></a>

#### `action.policer_action.ref.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-3002132333001311-0131012213123200-1110211132301101-2021302210020102-3022330013021212-0212021301100012-3103121230000102-3033212011102212"></a>

<a id="canonical-1000023003120133-2022022230133211-0311012001132221-1010311301012010-1030212331130020-0013333301020030-3112223223302302-3201001332222220"></a>

#### `action.policer_action.ref.tenant` property

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

<a id="canonical-2300233332021233-3303001000031101-2120301033201132-3212332330210313-2033000321300001-2331213333233321-1130023012013331-3320213302331030"></a>

<a id="canonical-3100133113212301-0023111110320112-2030202020013302-3313230102001133-1301222110313021-2323030222312130-0302133223033233-2110301020122200"></a>

#### `action.policer_action.ref.uid` property

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

<a id="canonical-0032023332120030-1021231301112230-2223021031203310-2321231220233200-0221103002122230-3331100012103221-1230100032311323-2310012320223022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `action.protocol_policer_action` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [action](resources--fast_acl_rule--reference--group-001.md#canonical-2323132233031011-0321032123110200-3322222321020213-0222230201011121-2112122222031112-1110200012230031-0103000130300210-2301310300233232)
- action.protocol_policer_action

<a id="canonical-3200311321001222-2110102012310333-1020232113103130-1112232022013000-3323123031321112-0111330120111010-1110220200001232-0131201030303330"></a>

Type: `"object"`. single nested block, Optional.

Protocol Policer Reference. Reference to policer object.

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

<a id="canonical-0112121312231122-1223301001113101-2213201300000123-0123121033313200-0331000132012202-0323133303331332-0222012320121000-3220320130223210"></a>

### Direct properties for `action.protocol_policer_action`

- [ref](resources--fast_acl_rule--reference--group-001.md#canonical-3210323231330213-1130110002032031-0222001321332110-0130201210000031-2122213332202100-1002103320213103-0030002211000120-2002230112113012): complete subsection reference.

<a id="canonical-3210323231330213-1130110002032031-0222001321332110-0130201210000031-2122213332202100-1002103320213103-0030002211000120-2002230112113012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `action.protocol_policer_action.ref` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [action](resources--fast_acl_rule--reference--group-001.md#canonical-2323132233031011-0321032123110200-3322222321020213-0222230201011121-2112122222031112-1110200012230031-0103000130300210-2301310300233232)
- [action.protocol_policer_action](resources--fast_acl_rule--reference--group-001.md#canonical-0032023332120030-1021231301112230-2223021031203310-2321231220233200-0221103002122230-3331100012103221-1230100032311323-2310012320223022)
- action.protocol_policer_action.ref

<a id="canonical-2310232220113133-1202003002003230-2222132233121120-1020311020131130-2312011111110213-1211330113123123-1011210103301230-0011121033011321"></a>

Type: `"object"`. list nested block, Optional.

Protocol policer Reference. Reference to protocol policer object.

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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220201200330202-1200300321201002-1231002013203122-3201122123133212-3231312222303220-0130320010312301-0101301120112021-0220021200013012"></a>

### Direct properties for `action.protocol_policer_action.ref`

<a id="canonical-1232112130232332-1210201103221102-1111023203332103-2033310021000301-2100112022301332-0102233020320113-1001022212321011-0031100103020120"></a>

#### `action.protocol_policer_action.ref.kind` property

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

<a id="canonical-1100222002320333-1121302311200303-3202003330012230-0332220103223022-0132331230331311-2210020110310321-2210123210221013-2120033112110203"></a>

<a id="canonical-3021012013233021-1000223311030331-3313231302010320-3122023111131202-3030222010123321-1321212313100203-0312222230120221-2033303211102313"></a>

#### `action.protocol_policer_action.ref.name` property

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

<a id="canonical-0123232223212031-0031222221003301-1323223322021103-3012001002201320-3313111210102002-1321020203321021-3320022230033111-0022000302112333"></a>

<a id="canonical-3310221002103122-2312301322311110-1131111222332102-0100221201320302-2333111212331223-2131202221121130-3322312130232022-2310302300123222"></a>

#### `action.protocol_policer_action.ref.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-2212132131020321-2002320011331223-0332211111120122-3202012330120202-2210212302320311-0330313331033333-0013101212321100-0302323210010131"></a>

<a id="canonical-0230011320020222-1021023301211010-2100021313002003-1212023300223213-1313320011021033-2112112030023002-3111011331312213-1110102212231223"></a>

#### `action.protocol_policer_action.ref.tenant` property

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

<a id="canonical-0012131111323212-3001031123132010-2022232212210021-2003021131313020-0202002313321121-1023202020123213-3131320303030323-3002113013212232"></a>

<a id="canonical-3121333213010131-0123301003223201-2122112301230133-0302001310100213-3231201201000202-3333320021222212-2222321021302020-0111212333231331"></a>

#### `action.protocol_policer_action.ref.uid` property

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

<a id="canonical-0313300332123332-0030203111330313-3121323032102000-0301210002212311-2030021003133210-1230233230030123-3221202110201133-2030302010102232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ip_prefix_set` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- ip_prefix_set

<a id="canonical-1320112323231321-2121002110321111-0233021012110202-2220030202011211-1112223021101131-1030323031232111-3021223101223331-3022110302210120"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ip\_prefix\_set, prefix\] List of references to ip\_prefix\_set objects.

Additional upstream details:

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

<a id="canonical-0301113231120212-1130221203131112-2301001310210123-0012133332001310-0003131320213312-3131001133201100-2112301023212230-0313120100122133"></a>

### Direct properties for `ip_prefix_set`

- [ref](resources--fast_acl_rule--reference--group-001.md#canonical-0021312121120211-1213221230010330-2130010031132123-3033012030021231-0022113010211312-3022200211131201-0113111322321131-2313112221210332): complete subsection reference.

<a id="canonical-0021312121120211-1213221230010330-2130010031132123-3033012030021231-0022113010211312-3022200211131201-0113111322321131-2313112221210332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ip_prefix_set.ref` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [ip_prefix_set](resources--fast_acl_rule--reference--group-001.md#canonical-0313300332123332-0030203111330313-3121323032102000-0301210002212311-2030021003133210-1230233230030123-3221202110201133-2030302010102232)
- ip_prefix_set.ref

<a id="canonical-3112101300321133-1200000002220023-1122333112030221-0200132013112300-1231013101332030-0220230320111323-1000102022300100-0023312212000321"></a>

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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122223220113222-1332012310310233-3120322123313002-3002010101112013-0133101333003323-1021022021113202-3123003322011303-0302002103331313"></a>

### Direct properties for `ip_prefix_set.ref`

<a id="canonical-3313231112233233-1121130120330103-2131300001030003-1130031303001322-0233103212211023-0321202123002212-0232112021002333-0002212100300102"></a>

#### `ip_prefix_set.ref.kind` property

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

<a id="canonical-3223203213020200-1021330232220111-2102232122133120-3233301230113320-2311131012212300-3013133321003001-1120111122001231-1121120301101302"></a>

<a id="canonical-2021000033231212-0221312321330022-2300023030022312-1102313221130313-1311210111300021-0320112210003220-0213002021111100-1123023110003213"></a>

#### `ip_prefix_set.ref.name` property

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

<a id="canonical-3030201130322322-2033332122103002-1103120011123323-2101300003130130-1330332121120222-3220012303033333-2003021133303231-0102100023201001"></a>

<a id="canonical-0312202110200212-0031211332033030-0010021033211030-1320021331213331-2322222233020323-2030221121333011-2322110313220223-0321203213300030"></a>

#### `ip_prefix_set.ref.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-0213301331220121-1020333033201102-2312022302201112-1012033010010232-1221011203101231-2123312223122130-0203010210111032-3032011131113211"></a>

<a id="canonical-2000331320131103-2211220302001011-2322231212031133-0030120221130320-1330203120203022-1112011011100012-1130010100112230-2102313303003011"></a>

#### `ip_prefix_set.ref.tenant` property

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

<a id="canonical-2233210120332330-1032131213113312-2200113230023300-3130332200033202-1232213031233120-1000231133031213-1111323022311202-1021230201123311"></a>

<a id="canonical-2133110011001332-1032213230322220-1202120023011331-0120100320112123-2233123003103330-0222231013222120-2103100232230222-3202021211311323"></a>

#### `ip_prefix_set.ref.uid` property

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

<a id="canonical-1032003000303102-1033202210203210-3222020302212013-0321300133103321-3312031111132110-3202311220103302-2233113211303003-3331013000330130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `port` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- port

<a id="canonical-1313032123001021-1103223311313212-2331130300201220-2203011000332020-0102231322112133-3021310203101000-1112123013023201-2022023322101020"></a>

Type: `"object"`. list nested block, Optional.

Source Ports. L4 port numbers to match.

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

<a id="canonical-1213132202232222-2032200311022210-1023010202022000-0020033010100333-3133022313121133-3133132110212130-0010301022130322-3203102221110032"></a>

### Direct properties for `port`

- [all](resources--fast_acl_rule--reference--group-001.md#canonical-3322223020322000-0200121011322320-0111201032210033-0031023301033131-1331231031112112-2022122302112031-3001301012202031-3121220113211011): complete subsection reference.

- [DNS](resources--fast_acl_rule--reference--group-001.md#canonical-2321310033213222-1210123211133230-0310201203322011-0113122020022133-0311122321230130-3131223320320133-1311233320003010-3300220333302233): complete subsection reference.

<a id="canonical-2320201230030033-2212123101313022-0012200112232230-3030321330122321-2233311332213211-3031301031012003-1330112003003032-0313212322333121"></a>

<a id="canonical-3112000011233011-3130313203110220-2121133023330301-1110120032230232-1011022330122212-1033101221101122-3202103320012001-2211302023222210"></a>

#### `port.user_defined` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3322223020322000-0200121011322320-0111201032210033-0031023301033131-1331231031112112-2022122302112031-3001301012202031-3121220113211011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `port.all` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [port](resources--fast_acl_rule--reference--group-001.md#canonical-1032003000303102-1033202210203210-3222020302212013-0321300133103321-3312031111132110-3202311220103302-2233113211303003-3331013000330130)
- port.all

<a id="canonical-2020132012123031-1220032122300120-2123000232101310-1012211003222100-0203030133033111-2032033131012203-3220232012200311-2322322332112012"></a>

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
all = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321310033213222-1210123211133230-0310201203322011-0113122020022133-0311122321230130-3131223320320133-1311233320003010-3300220333302233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `port.dns` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [port](resources--fast_acl_rule--reference--group-001.md#canonical-1032003000303102-1033202210203210-3222020302212013-0321300133103321-3312031111132110-3202311220103302-2233113211303003-3331013000330130)
- port.DNS

<a id="canonical-1103202110132303-1202020222111000-0103132330112102-2332001000103202-0332131002323022-3102000310333020-1012312203332103-1022120311330320"></a>

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
dns = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313132020210223-3030132113131010-0123233323202310-2312303132222323-1203330311231211-3212223222232203-1303200300003303-2200112333000313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `prefix` properties

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Property reference](resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- prefix

<a id="canonical-2102023111201032-1020201013120222-0200213223122302-1321112313032131-1232022301020120-2000122331213102-1321033020231223-1013000020122322"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-0333233030132120-1020323201130221-3101330232132120-0123120101220102-0322002221221231-2312030220132220-2223333012002101-2132233320300321"></a>

### Direct properties for `prefix`

<a id="canonical-3310212211111330-2003112000211132-2001022120231333-0113221321312022-2103200030203233-0131130201310123-0133300232333002-3211330100221121"></a>

#### `prefix.prefix` property

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

<a id="canonical-3100132302101032-0231233113011020-2113300231211030-0113331001112220-2233213301003123-2112012222201133-2123110313013103-2022223302030201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

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

<a id="canonical-0000021233323120-3333201232201122-2110013313312110-1302131101023331-2113210003221110-2121201303021223-0122231313320301-3121011222100201"></a>

### Direct properties for `timeouts`

<a id="canonical-2322120320123002-1220330330003030-0223330103001200-1220123232021303-3222112001132321-2000003302113032-0102311022100022-1300210103301302"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2203012020032322-0321010312110233-1311211013110233-3222331002002220-2211111000110011-0030023222333100-3032213020110220-1121110132333221"></a>

<a id="canonical-3203302310121022-0010112300313100-0020002020012011-1213030332010213-2320213130203100-2320123030012332-1021130022303321-3333003011000323"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0103022100003233-1211202002331230-3100222222320123-0123302103100301-1301313012301312-2333000222212131-1203201132133330-2223302210010313"></a>

<a id="canonical-0102210303320011-2311132120312201-2132220110230010-1312203221032111-2313301100130010-2020331311203032-2013222001132100-3033230233320113"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2333110030210331-3022301322212322-1321301123021232-1202231213331231-3333212231011331-1301233200222103-2023022133030000-1332010220330213"></a>

<a id="canonical-3231013332103311-1121220221300032-1320222321011322-3210210233113322-2003102232112203-3311333322233100-2011300023023023-3020220330220333"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
