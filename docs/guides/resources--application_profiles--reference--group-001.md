---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

<a id="canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300123300313211-3331231322100132-0103012232333131-0211013132202020-0132131121130212-2213113310011132-3211132230033320-2232122113022232"></a>

## Property reference — Property reference / 300231030013 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- Property reference

<a id="canonical-0220220031223201-0030030310003230-0330301011032203-0322313212231111-3202002231311003-1210210231133332-2122001230230120-3100310322001100"></a>

## Direct properties — Property reference / 300231030013 / 3

- [advanced_tcp_profile](resources--application_profiles--reference--group-001.md#canonical-0102012130211220-0033013023011312-0020312201233333-0110230303020102-0021300320132223-2232021231133230-0130333111122200-0300200330323322): complete subsection reference.

<a id="canonical-3003320301233333-2023332201113331-3130303002102301-3130303332021221-2023001303313210-0210311100333130-2220011013230121-0233012332032203"></a>

<a id="canonical-3320201100230013-3223203111232133-2213311320013032-2013032200211031-1012122000130333-1012333021303332-2211202321030002-3332311110220021"></a>

## annotations property — Property reference / 300231030013 / 4

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

- [ddos_profile](resources--application_profiles--reference--group-001.md#canonical-2202332031223123-2301302010321012-1011102311002222-1100030332313031-3311332110021011-2231133020100211-2122201323020110-1321212330120300): complete subsection reference.

<a id="canonical-1022323010112121-3202310321313202-0131332233302200-3023201113230020-0213120101331323-2302332003230011-0223110013112132-1103300310002121"></a>

<a id="canonical-1300332102131223-1112202302113220-2032133031231323-3121333123210330-3100210030201100-3011111210322020-3020102200032130-3310121102022001"></a>

## description property — Property reference / 300231030013 / 5

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

<a id="canonical-1002112012100332-2313122210012121-3002212310321333-2032312303233022-3032021220112221-3322211012122011-0112123101323233-3301303212022223"></a>

<a id="canonical-2101223110323123-0322001121323230-3212103021222010-2212312021112213-3023210331323211-3200230130020230-1221030312213200-3331120311131223"></a>

## disable property — Property reference / 300231030013 / 6

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

<a id="canonical-1100111001102022-3132002211330013-0012023030112101-1103031023223102-0202331231123311-2200121200000100-1222310030100030-2130212012100332"></a>

<a id="canonical-0331331123101112-2100233120332021-2021003130023231-0103332312002300-0230020213023000-0131001013010032-3201233211222112-3023203101032213"></a>

## ID property — Property reference / 300231030013 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [irules](resources--application_profiles--reference--group-001.md#canonical-0333211213222003-2111200010301230-0111200020033200-1110123010330321-0030333121010031-2213122011223020-1010213020122303-1230123111121033): complete subsection reference.

<a id="canonical-3220002310030311-0231010003303233-1122000121222331-2203011331122332-0111321130100102-0023122231311110-0230112230210222-1002202311122131"></a>

<a id="canonical-2001332311300213-1032202211313212-0221001333313013-0033020012101221-3103130212111111-3131010310311210-1312210320321132-0311131001320231"></a>

## labels property — Property reference / 300231030013 / 8

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

<a id="canonical-1321132013002021-3102303312301013-0200130312121313-1110032112003021-1100120222203003-1331222202213103-3100231110220021-3202121113003301"></a>

<a id="canonical-3033122102200333-0000130002001212-3210202321301031-2302032103112213-2031131302030121-3301112200231311-2001121312012033-3230022303110012"></a>

## name property — Property reference / 300231030013 / 9

Type: `"string"`. Required.

Name of the Application Profiles. Must be unique within the namespace.

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

<a id="canonical-1002333012222021-3022113002232333-0222303311100022-3333121003003131-3011121332110113-2232112312331001-3031221301122321-0231333313011321"></a>

<a id="canonical-1331303300020012-2222231220221313-1330302303230321-3321011032232120-1230321113032132-1023233310131102-1010002320330232-0020222002122311"></a>

## namespace property — Property reference / 300231030013 / 10

Type: `"string"`. Required.

Namespace where the Application Profiles is created.

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

- [timeouts](resources--application_profiles--reference--group-001.md#canonical-2100312103022331-2033132121021132-3011311223212110-1112030031330113-3322301201010000-3123302212130112-2110221332213112-2221112310232000): complete subsection reference.

- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312): complete subsection reference.

<a id="canonical-1212332132003020-1320210113313312-2012232011332330-1012102033022211-3201302230302111-3310113110002320-0321212232020130-0321133301331331"></a>

## All schema paths — Property reference / 300231030013 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_tcp_profile` | [advanced_tcp_profile](resources--application_profiles--reference--group-001.md#canonical-0131203132032110-0212101023311010-2032110211331123-0332202113033003-3002223023320122-1001030130001112-1112130311233202-1012311211333220) |
| `advanced_tcp_profile.disable_tcp_advanced_profile` | [advanced_tcp_profile.disable_tcp_advanced_profile](resources--application_profiles--reference--group-001.md#canonical-2213222212222231-2332222120213322-2130201032201311-1130332103012133-1103103330320222-3200122030123032-2032230233200130-2130230212302232) |
| `advanced_tcp_profile.enable_tcp_advanced_profile` | [advanced_tcp_profile.enable_tcp_advanced_profile](resources--application_profiles--reference--group-001.md#canonical-3111111131310203-2102020011312223-1201321213320001-0122120131322231-2310131133200331-0012000132222202-2233101310322133-0312022030301332) |
| `annotations` | [annotations](resources--application_profiles--reference--group-001.md#canonical-3003320301233333-2023332201113331-3130303002102301-3130303332021221-2023001303313210-0210311100333130-2220011013230121-0233012332032203) |
| `ddos_profile` | [ddos_profile](resources--application_profiles--reference--group-001.md#canonical-1223320121231003-3321303031010020-2312302332001103-1011301210323311-1331330011110330-1032011211322112-2032321111130203-0202211101010030) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](resources--application_profiles--reference--group-001.md#canonical-1321321102100022-1102333210032332-1120201113020323-3123202112031320-0010111123030133-3003001332111312-0103220223131301-3331212132311330) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](resources--application_profiles--reference--group-001.md#canonical-1220102123222302-2012021211123101-3110323021201220-1113031033222330-0021303312223003-0331312133330202-2010121203222003-2232210111101111) |
| `description` | [description](resources--application_profiles--reference--group-001.md#canonical-1022323010112121-3202310321313202-0131332233302200-3023201113230020-0213120101331323-2302332003230011-0223110013112132-1103300310002121) |
| `disable` | [disable](resources--application_profiles--reference--group-001.md#canonical-1002112012100332-2313122210012121-3002212310321333-2032312303233022-3032021220112221-3322211012122011-0112123101323233-3301303212022223) |
| `id` | [id](resources--application_profiles--reference--group-001.md#canonical-1100111001102022-3132002211330013-0012023030112101-1103031023223102-0202331231123311-2200121200000100-1222310030100030-2130212012100332) |
| `irules` | [irules](resources--application_profiles--reference--group-001.md#canonical-3201312303230210-3330121210002003-3201220220001113-3211311212101311-1223121230302030-2333311023002030-2203031222103002-0303020210202302) |
| `irules.kind` | [irules.kind](resources--application_profiles--reference--group-001.md#canonical-0213302010001001-3313301203133010-2303233322203311-2301300223202302-3120130111112033-1333221000023102-0123012202322320-1011200301031023) |
| `irules.name` | [irules.name](resources--application_profiles--reference--group-001.md#canonical-0232230210112222-1112102331103011-1232232303300223-0321102033301332-1211122302310022-2002113123331003-3012300133131231-0230101120233102) |
| `irules.namespace` | [irules.namespace](resources--application_profiles--reference--group-001.md#canonical-3132302021003023-3331333131300100-0211122111220101-2113322111121312-1010322133230311-1321031311323100-0111311113113133-2113131101221123) |
| `irules.tenant` | [irules.tenant](resources--application_profiles--reference--group-001.md#canonical-3201302212023101-1012123130132013-1231131003001012-3001130313230022-0121203321112021-2031022120303011-1000120011010102-3022232023113301) |
| `irules.uid` | [irules.uid](resources--application_profiles--reference--group-001.md#canonical-0122003020322120-1213002321111221-0201203003311301-3310231002022130-2132133231200301-0013002103213220-0110311122330132-2202212303202301) |
| `labels` | [labels](resources--application_profiles--reference--group-001.md#canonical-3220002310030311-0231010003303233-1122000121222331-2203011331122332-0111321130100102-0023122231311110-0230112230210222-1002202311122131) |
| `name` | [name](resources--application_profiles--reference--group-001.md#canonical-1321132013002021-3102303312301013-0200130312121313-1110032112003021-1100120222203003-1331222202213103-3100231110220021-3202121113003301) |
| `namespace` | [namespace](resources--application_profiles--reference--group-001.md#canonical-1002333012222021-3022113002232333-0222303311100022-3333121003003131-3011121332110113-2232112312331001-3031221301122321-0231333313011321) |
| `timeouts` | [timeouts](resources--application_profiles--reference--group-001.md#canonical-3023002233200132-1221101032202122-0203100022131133-0333230103323011-2133121300332321-2201131022300002-3101100102330133-2321212230102310) |
| `timeouts.create` | [timeouts.create](resources--application_profiles--reference--group-001.md#canonical-3100210310022100-3300213130331313-2202212203203103-1303310122020313-3220202020323103-3110303031303213-0231001302031132-3120100330301021) |
| `timeouts.delete` | [timeouts.delete](resources--application_profiles--reference--group-001.md#canonical-0233230313012211-2111222111213103-2110311121322213-0210332310020332-2122213210003130-1303223110111320-3313032330112322-0321322130231111) |
| `timeouts.read` | [timeouts.read](resources--application_profiles--reference--group-001.md#canonical-1321021001133122-2313303100033103-2322110102232232-2010121001320310-0221130230133211-2122332312122102-2203002322231120-0121312002002110) |
| `timeouts.update` | [timeouts.update](resources--application_profiles--reference--group-001.md#canonical-1130210232203300-2200030213200132-1310331333020003-0022011021010303-3321203300301301-3301010010123223-0302302013223310-0101233323032223) |
| `virtual_server` | [virtual_server](resources--application_profiles--reference--group-001.md#canonical-0223123303001200-0331130002020103-2301300000320300-2323002001312002-0013113220333013-1301230220223200-1311211110203213-3222203102122101) |
| `virtual_server.access_profile` | [virtual_server.access_profile](resources--application_profiles--reference--group-001.md#canonical-2223203132121233-0130230302003020-0022132130321212-1323133010230220-3133112020312012-2213310032031021-2221020320021022-1322020220333033) |
| `virtual_server.access_profile.kind` | [virtual_server.access_profile.kind](resources--application_profiles--reference--group-001.md#canonical-2312120313230122-0103133201013300-3020301012103301-2120001020102000-1130303212230010-2101320023030220-2233002110122232-3010231223330130) |
| `virtual_server.access_profile.name` | [virtual_server.access_profile.name](resources--application_profiles--reference--group-001.md#canonical-3333000021223122-1231320332233023-3321322021230303-1130012313132100-3033230200332103-3231222200333121-1021233212002330-3230103331230230) |
| `virtual_server.access_profile.namespace` | [virtual_server.access_profile.namespace](resources--application_profiles--reference--group-001.md#canonical-0102333102213322-3333132033032130-0322003333120312-2102231201112322-1033133130100203-0132022001303300-2310100322220313-1012302130202001) |
| `virtual_server.access_profile.tenant` | [virtual_server.access_profile.tenant](resources--application_profiles--reference--group-001.md#canonical-1311231113222022-0201313212032203-0001230023321233-0033222132211312-3033321230330210-2030002120222303-1023303112000121-1230201121221112) |
| `virtual_server.access_profile.uid` | [virtual_server.access_profile.uid](resources--application_profiles--reference--group-001.md#canonical-0021311103300301-2023023131211133-3123101003001111-3213030232303001-3001001321222230-2221011230103332-3013300113202012-2012301323130030) |
| `virtual_server.address_translation` | [virtual_server.address_translation](resources--application_profiles--reference--group-001.md#canonical-0300032312110001-3232021231013033-1202233220030012-2010213320123213-1112002012130110-0220131220101001-3220320331312001-0011330303333322) |
| `virtual_server.address_translation.address_translation_disable` | [virtual_server.address_translation.address_translation_disable](resources--application_profiles--reference--group-001.md#canonical-0012120321210313-1132301022231031-3132112031313330-1101133222033112-2212223011012323-3023223302313030-1030201213021112-2011302032200230) |
| `virtual_server.address_translation.address_translation_enable` | [virtual_server.address_translation.address_translation_enable](resources--application_profiles--reference--group-001.md#canonical-0331233123202101-2133022100220113-1033121202222303-0322123011213123-1011231330332022-3320310111310102-3033102012312201-1023303213331222) |
| `virtual_server.auto_last_hop` | [virtual_server.auto_last_hop](resources--application_profiles--reference--group-002.md#canonical-2102003333002100-3120130233030201-2230023011033122-1122220031100031-0002012133223230-1220000021212230-2313302313111212-3212210233202022) |
| `virtual_server.auto_last_hop.auto_last_hop_default` | [virtual_server.auto_last_hop.auto_last_hop_default](resources--application_profiles--reference--group-002.md#canonical-0200312103221313-1103111200001032-1221002220011310-0112113223110233-3233322010331203-3230013203122122-1130012233102331-2313231022313131) |
| `virtual_server.auto_last_hop.auto_last_hop_disable` | [virtual_server.auto_last_hop.auto_last_hop_disable](resources--application_profiles--reference--group-002.md#canonical-3333323311300221-1032210323101131-3110013321303220-0201100332302123-0222310211320100-1300320120332311-1223000303202000-2331032230221011) |
| `virtual_server.auto_last_hop.auto_last_hop_enable` | [virtual_server.auto_last_hop.auto_last_hop_enable](resources--application_profiles--reference--group-002.md#canonical-2123101121211332-1123033021312131-0220230032033320-1312120221300022-0313300221331012-3230112211311023-0332021203320110-3123130113132311) |
| `virtual_server.clone_pool_client` | [virtual_server.clone_pool_client](resources--application_profiles--reference--group-002.md#canonical-1033030230211202-1021333323012333-3323221312331011-3320122331112203-2220120321232231-0010221120323210-1310313301100022-2200300322112102) |
| `virtual_server.clone_pool_client.kind` | [virtual_server.clone_pool_client.kind](resources--application_profiles--reference--group-002.md#canonical-3120313123311212-3000231030202030-1210031020310120-3333210210332233-3202133133011221-1030021010303212-2120312233210302-2023202020020100) |
| `virtual_server.clone_pool_client.name` | [virtual_server.clone_pool_client.name](resources--application_profiles--reference--group-002.md#canonical-0103112213301230-3333211011322130-0231332103311020-1231211330030103-2010001223211303-1011311231010010-1231100100213312-0233122323023011) |
| `virtual_server.clone_pool_client.namespace` | [virtual_server.clone_pool_client.namespace](resources--application_profiles--reference--group-002.md#canonical-0301313033003103-2131023200312212-1301112022013032-2211213122102220-1122022010322210-2031320113103333-3012330020320002-1103031303111023) |
| `virtual_server.clone_pool_client.tenant` | [virtual_server.clone_pool_client.tenant](resources--application_profiles--reference--group-002.md#canonical-1131200111132033-1211033311321020-0022320230302121-2202023021000220-0103131210321232-2001310012032033-0311010322302310-3110220220220131) |
| `virtual_server.clone_pool_client.uid` | [virtual_server.clone_pool_client.uid](resources--application_profiles--reference--group-002.md#canonical-3123233321330320-0331033131032210-3232200313222301-2030302122301231-2203003022113133-2100331021232122-3022132133013223-0201330110103311) |
| `virtual_server.clone_pool_server` | [virtual_server.clone_pool_server](resources--application_profiles--reference--group-002.md#canonical-1013221131231303-0110323201210021-1301332311032223-1333023030001320-2100220021203031-0202002323323223-3032101233123232-2120000203203121) |
| `virtual_server.clone_pool_server.kind` | [virtual_server.clone_pool_server.kind](resources--application_profiles--reference--group-002.md#canonical-1223121310000131-0022301003011121-2221122103302113-0030111003021112-1230302210221101-0012300132033202-0002323331303012-2030330031211032) |
| `virtual_server.clone_pool_server.name` | [virtual_server.clone_pool_server.name](resources--application_profiles--reference--group-002.md#canonical-2302010132332301-2211231221023230-1110110111231030-1302333311012223-0130311111330321-2322033223012313-0310332212302203-0320330101122133) |
| `virtual_server.clone_pool_server.namespace` | [virtual_server.clone_pool_server.namespace](resources--application_profiles--reference--group-002.md#canonical-3030102120302331-0130002010320330-3332321233111002-1313330233130311-2122320233302211-3211030301202321-3021002013022323-3220202312102131) |
| `virtual_server.clone_pool_server.tenant` | [virtual_server.clone_pool_server.tenant](resources--application_profiles--reference--group-002.md#canonical-3113330030320212-2210233311203202-2321120111202211-0231111010022231-0210132132323001-0323002123010031-3211213020010312-0130123230233000) |
| `virtual_server.clone_pool_server.uid` | [virtual_server.clone_pool_server.uid](resources--application_profiles--reference--group-002.md#canonical-2020312020003200-3122200122222001-3131220100232322-2021232032133202-1003332112322122-2313112313203103-1332023213112111-1003312303323102) |
| `virtual_server.connection_limit` | [virtual_server.connection_limit](resources--application_profiles--reference--group-001.md#canonical-3302010203222333-0123332120302230-1133001202003313-2012021112101211-2220012001012020-2111220133210220-2123201312330130-1200112232202230) |
| `virtual_server.connection_rate_limit` | [virtual_server.connection_rate_limit](resources--application_profiles--reference--group-001.md#canonical-3212101123123123-2202303100203133-3103230100003322-0303123102100200-0230003130122322-2111302331112002-3110320121113212-2203202230321303) |
| `virtual_server.connection_rate_limit_mode` | [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-3301023333300032-2210013303130033-3332030022313210-1311132001212223-3022013121111032-0221303131332122-2021123212320013-0221332020333311) |
| `virtual_server.connection_rate_limit_mode.per_destination_address` | [virtual_server.connection_rate_limit_mode.per_destination_address](resources--application_profiles--reference--group-002.md#canonical-3313232302320002-1322131102311312-3202132010023031-2321130021120013-2003133200130013-2122320003212321-3020011300103011-3302222023132003) |
| `virtual_server.connection_rate_limit_mode.per_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_destination_address.destination_mask](resources--application_profiles--reference--group-002.md#canonical-1011231323321110-0220010012111331-0001312032130211-2230012002022120-2203231322201101-2003032111331110-0222310201000131-1002031220203333) |
| `virtual_server.connection_rate_limit_mode.per_source_address` | [virtual_server.connection_rate_limit_mode.per_source_address](resources--application_profiles--reference--group-002.md#canonical-1032313202203331-2213320301221002-2211210101321220-3133123230322220-1301210201332131-0321031301121220-2332221203212311-1323130120131323) |
| `virtual_server.connection_rate_limit_mode.per_source_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_source_address.source_mask](resources--application_profiles--reference--group-002.md#canonical-3030132331223231-2311121121311021-0023330023232220-2130321003300202-0321323130233121-2333301200322333-2132323112102003-1021023322120103) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address` | [virtual_server.connection_rate_limit_mode.per_source_destination_address](resources--application_profiles--reference--group-002.md#canonical-1121321010302331-3210012021133201-1311000332111023-0102212333311320-1302133003221333-3332301131221312-1322321301203313-2200323113202003) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_source_destination_address.destination_mask](resources--application_profiles--reference--group-002.md#canonical-3130120312101122-1301101010333203-1321211202203331-3233022013120312-1312002132212303-3103202112111212-2310220211133201-1222332310111100) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_source_destination_address.source_mask](resources--application_profiles--reference--group-002.md#canonical-3131112320301313-1101232211113011-3001121131001302-0101312032323111-1000000131131032-2202312000332200-1131023233110021-3031121111323311) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server` | [virtual_server.connection_rate_limit_mode.per_virtual_server](resources--application_profiles--reference--group-002.md#canonical-3130113030210232-3312223213230030-2322321202130103-1302122101233331-0120310132320003-1010003012102203-1003023002131030-1022333022323012) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address](resources--application_profiles--reference--group-002.md#canonical-2012022113300113-2301301113011120-2020123111303122-1010131133200020-1212100222003022-2110131011230123-2303210200233110-1300223013230020) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address.destination_mask](resources--application_profiles--reference--group-002.md#canonical-1000011102212103-3113131323031300-1033321133312130-3011111121312303-0032332330103213-2130301222003102-0232130310131322-0012233033333233) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address](resources--application_profiles--reference--group-002.md#canonical-2101303130210111-1220102123110101-0231331023010311-1203131013010310-3023132221000003-3032200132201300-3222303021122200-0133330120331001) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address.source_mask](resources--application_profiles--reference--group-002.md#canonical-0333231212030220-1021031331111010-0321211202222011-3221220033112032-2233122210303320-0132212011202312-3223331303230003-0201112110123333) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address](resources--application_profiles--reference--group-002.md#canonical-0331310010022310-2302213032223221-1223232003201100-3021010332102311-2211130113110022-0223001011223003-3212310223321212-0023232130313231) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.destination_mask](resources--application_profiles--reference--group-002.md#canonical-1131102111332120-3131131320131203-3203120021321221-0233121222103203-2213311330001211-2100023100332312-3133011120012220-3321112230102313) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.source_mask](resources--application_profiles--reference--group-002.md#canonical-1231033213102212-0020130033122133-0330222102122121-2211323131230012-0200201101111020-1020212310313223-3322201022120002-1000312130312113) |
| `virtual_server.default_persistence_profile` | [virtual_server.default_persistence_profile](resources--application_profiles--reference--group-002.md#canonical-3003330212123022-1202023302201020-2222131130302220-3103112023103232-0120013021311210-1301120010220333-0031231203030233-2001000121103120) |
| `virtual_server.default_persistence_profile.kind` | [virtual_server.default_persistence_profile.kind](resources--application_profiles--reference--group-002.md#canonical-1032030113301233-3012310320011110-2321311321123130-1010033223211110-3002330031100210-1033101231011032-0021331301222213-1312000100200102) |
| `virtual_server.default_persistence_profile.name` | [virtual_server.default_persistence_profile.name](resources--application_profiles--reference--group-002.md#canonical-1013302000100300-0300223331100030-2301030003001313-3211212333200301-0220220033023021-2132120331210312-3201021221330201-0231103211111301) |
| `virtual_server.default_persistence_profile.namespace` | [virtual_server.default_persistence_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-0123110133310023-2330220232011220-2010113312321210-1213112030223011-2031102300213301-3331013321030323-0212110123021301-2022130110313013) |
| `virtual_server.default_persistence_profile.tenant` | [virtual_server.default_persistence_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-2132010111023003-1310103300310111-0023222132330330-3203010001133001-1112031101200301-3312011210302002-1201102312131100-2131122132111300) |
| `virtual_server.default_persistence_profile.uid` | [virtual_server.default_persistence_profile.uid](resources--application_profiles--reference--group-002.md#canonical-2302120021022132-1013123122013002-2213233113310232-3033120213200102-1032003123003033-2103113033103000-3223303121211332-1320121123222331) |
| `virtual_server.default_pool` | [virtual_server.default_pool](resources--application_profiles--reference--group-002.md#canonical-3323300103233100-2200233303311301-1122232232202202-1300201201233111-3320200010030210-3001213020011320-2330222130011223-0013330220133002) |
| `virtual_server.default_pool.kind` | [virtual_server.default_pool.kind](resources--application_profiles--reference--group-002.md#canonical-3233011302111323-3010331222123322-2113221113201212-0300101020323000-3113313031321130-0131030001233023-0213100200220203-2230003310200201) |
| `virtual_server.default_pool.name` | [virtual_server.default_pool.name](resources--application_profiles--reference--group-002.md#canonical-1002132220121110-1130032020003033-2002033121132232-2120023033003231-0022110202310230-3332022011103131-0213321001113130-0200221001031032) |
| `virtual_server.default_pool.namespace` | [virtual_server.default_pool.namespace](resources--application_profiles--reference--group-002.md#canonical-0110023303023330-3310132013030213-2011022132022302-3230220213233313-1301133323333101-2033203332112300-3333200010023021-0213013001211022) |
| `virtual_server.default_pool.tenant` | [virtual_server.default_pool.tenant](resources--application_profiles--reference--group-002.md#canonical-3302123123312220-1320322313001213-0101323301020221-0202312333130231-1110112320010232-1200110201300211-1012213220310203-0033010123332220) |
| `virtual_server.default_pool.uid` | [virtual_server.default_pool.uid](resources--application_profiles--reference--group-002.md#canonical-0012023333200002-2030020322213323-3012222013333221-1221232000013020-3222212301313023-1313110223331002-1302313223301110-3332300221000103) |
| `virtual_server.fallback_persistence_profile` | [virtual_server.fallback_persistence_profile](resources--application_profiles--reference--group-002.md#canonical-3021300023002001-2000330111223011-0331013022010230-3330313113132111-3323032232021003-3122301013122023-2212102223011130-3300130111221331) |
| `virtual_server.fallback_persistence_profile.kind` | [virtual_server.fallback_persistence_profile.kind](resources--application_profiles--reference--group-002.md#canonical-1330112230020101-0021310312221331-0103121112122020-0210321112111223-3103001323110320-2032210302321233-0011320310231023-1030033230022203) |
| `virtual_server.fallback_persistence_profile.name` | [virtual_server.fallback_persistence_profile.name](resources--application_profiles--reference--group-002.md#canonical-2030221001130301-3122031011310203-0303101300031310-0000132113321210-1223121302130113-2012233200300030-2301032112033000-0112101012230311) |
| `virtual_server.fallback_persistence_profile.namespace` | [virtual_server.fallback_persistence_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-1123323201121001-2013133301202012-0113102110322300-3331131322120030-3231222132013022-2323000332133031-1301002101110230-3121330230102010) |
| `virtual_server.fallback_persistence_profile.tenant` | [virtual_server.fallback_persistence_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-1333323303123233-2212133033033032-2120232102123002-1220002122200002-1123310211023213-1031113002012221-0120113120223010-1111003031133230) |
| `virtual_server.fallback_persistence_profile.uid` | [virtual_server.fallback_persistence_profile.uid](resources--application_profiles--reference--group-002.md#canonical-3301210113012320-2212320330000330-1011212000223311-1302222001220112-2032220300132021-0121012331200331-0131023223302302-3332022203003302) |
| `virtual_server.fix_profile` | [virtual_server.fix_profile](resources--application_profiles--reference--group-002.md#canonical-2220330223223032-2003200010320322-3132202123100121-2211120320000322-1013120130311031-3101111201103333-3311130121031320-2122213302333231) |
| `virtual_server.fix_profile.kind` | [virtual_server.fix_profile.kind](resources--application_profiles--reference--group-002.md#canonical-2332022332123013-0013231020120012-1231010201023101-2132002021110223-0001312112002302-1330333131031221-1210320030303330-2122211033200233) |
| `virtual_server.fix_profile.name` | [virtual_server.fix_profile.name](resources--application_profiles--reference--group-002.md#canonical-2233022010323031-2220201322202013-1002120001010322-0010201330310101-0200120100333021-1333030301222310-3033101302323022-0321333123322100) |
| `virtual_server.fix_profile.namespace` | [virtual_server.fix_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-0203032130000311-2321212301033301-2331313011220000-1003032130330230-3330310231333222-3121233002203331-2231213300212122-3213330112031323) |
| `virtual_server.fix_profile.tenant` | [virtual_server.fix_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-2121121222100033-1120220212111320-2310301202130120-1222201030031022-0130123120131231-2101223003221201-0301320022220202-3323110232102322) |
| `virtual_server.fix_profile.uid` | [virtual_server.fix_profile.uid](resources--application_profiles--reference--group-002.md#canonical-1202223323102200-0012100002112132-0012200300132221-1313220330133232-2130023120233002-2230212102100031-3023333312023321-0022200132010232) |
| `virtual_server.http` | [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2332010000332020-0110331212332300-3330311200330200-3222000010231313-1221311303330220-0313132223022231-3301130303131001-0033103233213202) |
| `virtual_server.http.client_ssl_profile` | [virtual_server.http.client_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-2003221132223331-1212333232201023-3211311333310031-2221202111301221-1323210013230332-3312220301221300-0303321020102203-2122120110233333) |
| `virtual_server.http.client_ssl_profile.kind` | [virtual_server.http.client_ssl_profile.kind](resources--application_profiles--reference--group-002.md#canonical-3300332013002312-3210111301031111-3012202101030120-3211323013131232-1312222001320310-2211103133103111-0100030301231002-0323210022012010) |
| `virtual_server.http.client_ssl_profile.name` | [virtual_server.http.client_ssl_profile.name](resources--application_profiles--reference--group-002.md#canonical-2031111110320233-0033303223312000-0013213013110100-2001102310113302-2100311030223303-0200333220310101-3222220300022122-2131230211322113) |
| `virtual_server.http.client_ssl_profile.namespace` | [virtual_server.http.client_ssl_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-3333032332002022-0023210022110332-3202002121330201-3123123001320133-3000221013122223-1222221202331031-3211103312010311-1101121200102033) |
| `virtual_server.http.client_ssl_profile.tenant` | [virtual_server.http.client_ssl_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-0201312331013222-3021233111130210-0310113000300213-2310122230233121-1211210333210031-0320203032213200-1312031203132210-2200022313012022) |
| `virtual_server.http.client_ssl_profile.uid` | [virtual_server.http.client_ssl_profile.uid](resources--application_profiles--reference--group-002.md#canonical-2032321300031201-1102321000211001-2000213130132310-3231330313312130-3012223333302123-3102320013222011-3002233220302122-1212100222200310) |
| `virtual_server.http.http2_client_profile` | [virtual_server.http.http2_client_profile](resources--application_profiles--reference--group-002.md#canonical-1022312112332200-0222223130011023-1211321323203210-1222020210311023-1120011101100232-2023112312302133-0012200131203330-1232112223113013) |
| `virtual_server.http.http2_client_profile.kind` | [virtual_server.http.http2_client_profile.kind](resources--application_profiles--reference--group-002.md#canonical-3101002331323112-2223133331130222-3111323123123223-3111021303031302-1213213010321303-3300003323313303-3101322130311202-0103130121323332) |
| `virtual_server.http.http2_client_profile.name` | [virtual_server.http.http2_client_profile.name](resources--application_profiles--reference--group-002.md#canonical-3322131032121130-0313332202013030-1032331120313213-1011200231110213-1103222011113212-1013311131311203-1032300021000333-2033300322023010) |
| `virtual_server.http.http2_client_profile.namespace` | [virtual_server.http.http2_client_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-0333330322231131-1123322131032031-2311313301022110-1110211102003001-3211120023113332-2233233300331313-1302302210003320-1301121102112330) |
| `virtual_server.http.http2_client_profile.tenant` | [virtual_server.http.http2_client_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-2103320101311310-1013130302023321-1102003220102122-0132001321212111-3022030210201221-3211331321013022-1011331003111231-0032320230233002) |
| `virtual_server.http.http2_client_profile.uid` | [virtual_server.http.http2_client_profile.uid](resources--application_profiles--reference--group-002.md#canonical-1121111232030031-2313113321020320-3211322311103002-0321310331030301-3201231100210230-3002203221331333-2001030022122330-2321022033330322) |
| `virtual_server.http.http2_server_profile` | [virtual_server.http.http2_server_profile](resources--application_profiles--reference--group-002.md#canonical-1021113223331302-0220202131331210-2123132120213301-3310212101201220-2111233331121033-0022122331121111-3201000211233123-3120021232320232) |
| `virtual_server.http.http2_server_profile.kind` | [virtual_server.http.http2_server_profile.kind](resources--application_profiles--reference--group-002.md#canonical-2221023020223311-0113331122232003-3123201221211132-0000330112100001-1313311120012110-2000001110113330-2112022021323332-3310022020330032) |
| `virtual_server.http.http2_server_profile.name` | [virtual_server.http.http2_server_profile.name](resources--application_profiles--reference--group-002.md#canonical-3032011303010020-2103020203002013-0100103303220001-2203233132100002-2312302030322320-0132213311010201-2211233302013030-2302330322212133) |
| `virtual_server.http.http2_server_profile.namespace` | [virtual_server.http.http2_server_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-3001103022033230-1333001103131111-1022222132003100-0113211001231031-2121321120001103-3220003103310033-2101212200111231-1010221213303121) |
| `virtual_server.http.http2_server_profile.tenant` | [virtual_server.http.http2_server_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-1031220312220123-1011233100132310-1230331123101311-1233120220201033-3223312101201130-0031002211102313-1312010113110301-0111010221020213) |
| `virtual_server.http.http2_server_profile.uid` | [virtual_server.http.http2_server_profile.uid](resources--application_profiles--reference--group-002.md#canonical-0330130332102210-1110033022220012-0031132320113020-3303220320210011-2132210110312230-2003232211203320-3103023211322002-1010220333130301) |
| `virtual_server.http.http_client_profile` | [virtual_server.http.http_client_profile](resources--application_profiles--reference--group-002.md#canonical-1113100222021131-2011310210010022-0120023101322100-0330202311321000-3132200232113230-2323110010233002-3321311201232300-1222113233120233) |
| `virtual_server.http.http_client_profile.kind` | [virtual_server.http.http_client_profile.kind](resources--application_profiles--reference--group-002.md#canonical-0101021333322322-3303132201311230-2100031032312123-1333110312111133-2220121011121020-1301131002333022-1313130302222233-2223123131313002) |
| `virtual_server.http.http_client_profile.name` | [virtual_server.http.http_client_profile.name](resources--application_profiles--reference--group-002.md#canonical-3220303301223023-2123211100031220-0200233111011000-1111223322003323-0100301203001321-3133110102230110-2101212210012010-0230313301323011) |
| `virtual_server.http.http_client_profile.namespace` | [virtual_server.http.http_client_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-1030111303102331-1302321220333001-1201010133021102-3213303223210011-3231003312330102-1332331322310330-2201211132010113-1312111020233022) |
| `virtual_server.http.http_client_profile.tenant` | [virtual_server.http.http_client_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-3300100233013102-1012010320002031-3302131002012222-1210310113300010-0220032232010232-1103131003232011-0321201022302000-2112030031233100) |
| `virtual_server.http.http_client_profile.uid` | [virtual_server.http.http_client_profile.uid](resources--application_profiles--reference--group-002.md#canonical-0311132203322311-0233130001201300-1233022103211333-2133013320020233-0111320303133032-2212132112122010-0233220013131303-2002022300220012) |
| `virtual_server.http.http_server_profile` | [virtual_server.http.http_server_profile](resources--application_profiles--reference--group-002.md#canonical-1213110300000223-2122321330223221-3123002011000213-0021132300332002-0032022023223023-3133330032100221-1300020310321322-0223322210021032) |
| `virtual_server.http.http_server_profile.kind` | [virtual_server.http.http_server_profile.kind](resources--application_profiles--reference--group-002.md#canonical-0123112123120200-0102021220033201-1200130211101330-1032020001301303-0032013331130301-2130212010113221-1112202111232321-1123232123133233) |
| `virtual_server.http.http_server_profile.name` | [virtual_server.http.http_server_profile.name](resources--application_profiles--reference--group-002.md#canonical-1300111231210032-1020010010131011-3012310202020113-1312211021223123-3133001111312332-0201300000232011-0133320013331312-2302133101002003) |
| `virtual_server.http.http_server_profile.namespace` | [virtual_server.http.http_server_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-3121232322132020-3332023201332123-0233203032212023-0220012213023322-2313011333231313-2313001112232003-3300222222031020-3100101322120233) |
| `virtual_server.http.http_server_profile.tenant` | [virtual_server.http.http_server_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-2323312323003313-3033312331322211-3203313301020221-2132111232212003-1333322322230012-3332113030213222-0231223211023212-0121210300123022) |
| `virtual_server.http.http_server_profile.uid` | [virtual_server.http.http_server_profile.uid](resources--application_profiles--reference--group-002.md#canonical-0122133312001331-1022310203010302-0233012103332213-1311201022311230-0112112322122231-0231010330131111-0203100211213130-2232202123322202) |
| `virtual_server.http.ocsp_profile` | [virtual_server.http.ocsp_profile](resources--application_profiles--reference--group-002.md#canonical-3100301002030221-1333133211330332-3023213213021130-1030100001100123-2231320331330011-2103212330310100-0332130113313302-2323130220000223) |
| `virtual_server.http.ocsp_profile.kind` | [virtual_server.http.ocsp_profile.kind](resources--application_profiles--reference--group-002.md#canonical-0213213000023030-2031213303111300-2301011013312220-3233030330333331-2303320233112223-2302213020122111-2022131003003223-0212121233202012) |
| `virtual_server.http.ocsp_profile.name` | [virtual_server.http.ocsp_profile.name](resources--application_profiles--reference--group-002.md#canonical-0130120301011100-3200223212121012-2202001302003132-0303010233320232-0011223121021233-0301013012310000-2003313030001033-1213100102333200) |
| `virtual_server.http.ocsp_profile.namespace` | [virtual_server.http.ocsp_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-0300202233010331-2211211132210022-3103231202213033-2333120002112232-2102132133211300-3331102322031120-0030033330013213-3222201230021123) |
| `virtual_server.http.ocsp_profile.tenant` | [virtual_server.http.ocsp_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-0311301120303300-3322213301301012-2130303131032301-0032202323001220-1212123331330100-0101022012112321-3020003212023300-0300201210012211) |
| `virtual_server.http.ocsp_profile.uid` | [virtual_server.http.ocsp_profile.uid](resources--application_profiles--reference--group-002.md#canonical-3001331120213330-1222231232223200-1110021010300110-0130323311033013-0013120233222121-3123011202013232-3111003112311002-2100103111312221) |
| `virtual_server.http.server_ssl_profile` | [virtual_server.http.server_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-3010133303300011-3222112300303320-1201223020122330-2133200201010012-2123103122120202-1121303321332202-3320312331123202-2000333302030120) |
| `virtual_server.http.server_ssl_profile.kind` | [virtual_server.http.server_ssl_profile.kind](resources--application_profiles--reference--group-002.md#canonical-1012310300203232-2210311031133003-1322030232220323-2001000301031231-3120022313323033-2022301133213233-3330000101120030-0303300012102210) |
| `virtual_server.http.server_ssl_profile.name` | [virtual_server.http.server_ssl_profile.name](resources--application_profiles--reference--group-002.md#canonical-0113130033330003-2001102213130102-1213331212222101-2132022032322212-2310003201003202-3321223130121131-1321110310331233-3103312332332022) |
| `virtual_server.http.server_ssl_profile.namespace` | [virtual_server.http.server_ssl_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-3103311222322332-2313101313122333-3303311220230231-0012211131102130-2213100321122230-3332020232033121-1031011203311033-0122011212102310) |
| `virtual_server.http.server_ssl_profile.tenant` | [virtual_server.http.server_ssl_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-0311331321301221-2032133113211303-2112213231301110-0201103110220200-3003031010213011-1133332202222133-2110211303220333-3312310130001133) |
| `virtual_server.http.server_ssl_profile.uid` | [virtual_server.http.server_ssl_profile.uid](resources--application_profiles--reference--group-002.md#canonical-3000002222332213-0130233303102212-3121023212100231-2021021123110012-1303203313201212-1130232013301210-2011121003131103-0312332303203130) |
| `virtual_server.http.stream_profile` | [virtual_server.http.stream_profile](resources--application_profiles--reference--group-002.md#canonical-3102013030110002-1100221012323030-3312013033112103-3012023300122022-0132211310102010-3301002032211022-2110002101003312-1033331011012212) |
| `virtual_server.http.stream_profile.kind` | [virtual_server.http.stream_profile.kind](resources--application_profiles--reference--group-002.md#canonical-0033323323112113-3121331332122111-1311120203312000-3203113121021202-1322103133020021-3203230201010012-0001223001120103-2232223001321020) |
| `virtual_server.http.stream_profile.name` | [virtual_server.http.stream_profile.name](resources--application_profiles--reference--group-002.md#canonical-2333233112332102-0020003323010312-1121221332221311-0012213223102203-2311322033020333-2130013013203211-3313231221300103-0132200032003133) |
| `virtual_server.http.stream_profile.namespace` | [virtual_server.http.stream_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-2020130232330213-2101101323212120-0010000313203013-2002103202003130-3233023012013332-2122102130203323-2321210222122320-3011210113103112) |
| `virtual_server.http.stream_profile.tenant` | [virtual_server.http.stream_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-3133301103203232-3203121101200023-3133323320130221-1121000120110131-0131212201012111-0220131222022302-2233121332221033-2300300131010313) |
| `virtual_server.http.stream_profile.uid` | [virtual_server.http.stream_profile.uid](resources--application_profiles--reference--group-002.md#canonical-0102131022332311-1301310130321013-0101322320030322-0130333101203301-2331030033010331-2310202032210231-2301021110032231-0131310220211212) |
| `virtual_server.http.tcp_client_profile` | [virtual_server.http.tcp_client_profile](resources--application_profiles--reference--group-002.md#canonical-2222103002213222-1321203331302022-1212333301101011-1232112221113131-2021123322210003-3330013232203221-1323211110021210-3203312102321133) |
| `virtual_server.http.tcp_client_profile.kind` | [virtual_server.http.tcp_client_profile.kind](resources--application_profiles--reference--group-002.md#canonical-3100210000103122-1130131103112020-3303032232032300-1233220333223213-0023323313231002-1203010022010012-1001123231120003-3331313322111012) |
| `virtual_server.http.tcp_client_profile.name` | [virtual_server.http.tcp_client_profile.name](resources--application_profiles--reference--group-002.md#canonical-2101110333101233-1300123303133222-1202110230223302-1321102021300203-3002310000031003-2322223030230020-0202121331130120-1023311102201100) |
| `virtual_server.http.tcp_client_profile.namespace` | [virtual_server.http.tcp_client_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-0013311201323312-2211312312312022-1022302111030002-1331110120012212-0023300031202200-2032313321123311-1030033232201332-0020120103230000) |
| `virtual_server.http.tcp_client_profile.tenant` | [virtual_server.http.tcp_client_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-0321200223302101-0100131032311111-0222202033010030-3033233333010002-1332322212302022-2311011010122100-3020331310201322-2132231001330320) |
| `virtual_server.http.tcp_client_profile.uid` | [virtual_server.http.tcp_client_profile.uid](resources--application_profiles--reference--group-002.md#canonical-2003010303310013-0031210011230133-0321133231132022-3110212212101030-2200021033221232-3213310211310021-0301001302231330-3112322321131231) |
| `virtual_server.http.tcp_server_profile` | [virtual_server.http.tcp_server_profile](resources--application_profiles--reference--group-002.md#canonical-0032322120000231-2113212203313122-3322322010120231-3301111113111021-2031230313322100-2220221313011021-3303121310113132-0022300100310032) |
| `virtual_server.http.tcp_server_profile.kind` | [virtual_server.http.tcp_server_profile.kind](resources--application_profiles--reference--group-002.md#canonical-3330231201203003-1031222202310302-0120220020202323-2322323012032011-3311300120301120-0222301302122033-2301321023001223-0011000202001000) |
| `virtual_server.http.tcp_server_profile.name` | [virtual_server.http.tcp_server_profile.name](resources--application_profiles--reference--group-002.md#canonical-3101211303223032-1231020320220223-3022020310322030-3133221330020101-1201001203121132-2101202221332231-0103020131322130-3012002001120232) |
| `virtual_server.http.tcp_server_profile.namespace` | [virtual_server.http.tcp_server_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-3021032011003322-1332330220020233-1023122033131112-1312120123200023-0213320013313013-3132003313111030-2213012133020212-0220220010130321) |
| `virtual_server.http.tcp_server_profile.tenant` | [virtual_server.http.tcp_server_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-0223011121300023-2212232213313220-2202230031232023-1332223320023032-2211330212023002-1322301003313011-1013332330122023-1313100031320223) |
| `virtual_server.http.tcp_server_profile.uid` | [virtual_server.http.tcp_server_profile.uid](resources--application_profiles--reference--group-002.md#canonical-3012110322231301-2021220300300201-0311100000133220-3032032033330021-2132321233031013-3232202101330013-2012220323221322-2023123201012322) |
| `virtual_server.http.websocket_client_profile` | [virtual_server.http.websocket_client_profile](resources--application_profiles--reference--group-002.md#canonical-2223110220120301-1213223332122020-1320112133333333-0201211010202031-1022132031331300-0332032202323012-1303202122321332-2231221120220032) |
| `virtual_server.http.websocket_client_profile.kind` | [virtual_server.http.websocket_client_profile.kind](resources--application_profiles--reference--group-002.md#canonical-3030032010131311-1001103223310323-0011013221123013-0013102111032322-2023230033303313-3033131033103131-0013000113201101-1313021313110222) |
| `virtual_server.http.websocket_client_profile.name` | [virtual_server.http.websocket_client_profile.name](resources--application_profiles--reference--group-002.md#canonical-1122302322322001-2102320321021121-3110330220332222-0011032330322310-0311312112011030-0003212023003011-3103113013001330-3120132110033101) |
| `virtual_server.http.websocket_client_profile.namespace` | [virtual_server.http.websocket_client_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-3101102123202032-2113102303113130-3020312322030012-3101312332022031-1100313233303333-1131310021301322-2233032011011230-0001033110103012) |
| `virtual_server.http.websocket_client_profile.tenant` | [virtual_server.http.websocket_client_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-3200203130131213-2211031120033003-1103221001102300-0202201030013133-0023201102112131-1110003113111032-3301000331321032-3010303322220222) |
| `virtual_server.http.websocket_client_profile.uid` | [virtual_server.http.websocket_client_profile.uid](resources--application_profiles--reference--group-002.md#canonical-1331111301122210-1023033323332121-3112312330002232-1223210212221120-2332030012232213-1102232323103320-2032322201002202-3221031013320323) |
| `virtual_server.http.websocket_server_profile` | [virtual_server.http.websocket_server_profile](resources--application_profiles--reference--group-002.md#canonical-1202232302031303-0013323212010003-2020001202002320-1103221331022212-0113213022111122-2102330111012133-0022131122303201-0103012213310211) |
| `virtual_server.http.websocket_server_profile.kind` | [virtual_server.http.websocket_server_profile.kind](resources--application_profiles--reference--group-002.md#canonical-2131021222322302-1310112022132011-1112213013312312-0223320031030002-2101020321001130-1121113321202122-1212200122031320-2311103013020331) |
| `virtual_server.http.websocket_server_profile.name` | [virtual_server.http.websocket_server_profile.name](resources--application_profiles--reference--group-002.md#canonical-1231113113020320-1002213012123130-3102331221323121-2101130331232211-2001000333102302-0032102100002210-1210213211303030-3310211303022230) |
| `virtual_server.http.websocket_server_profile.namespace` | [virtual_server.http.websocket_server_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-3333101202000331-0331320333223120-0310321333023120-1330033023011123-2300203111011200-2121130100230100-2133221300332113-2231023220210223) |
| `virtual_server.http.websocket_server_profile.tenant` | [virtual_server.http.websocket_server_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-1111213312131202-0002000100232122-2301033120001112-1233023323203101-0230010112233302-0321021031112231-1123033221223311-0312211321133331) |
| `virtual_server.http.websocket_server_profile.uid` | [virtual_server.http.websocket_server_profile.uid](resources--application_profiles--reference--group-002.md#canonical-0031201230102221-3321111012120001-3223303333010333-2303231200023111-3333102132002302-0000301232013021-0231233121031122-3021213231131011) |
| `virtual_server.http3` | [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-2002231130013121-1320001010323121-2331302320313322-2221020113232332-1113213211110201-1022230001113101-1132301132322222-3120032223001131) |
| `virtual_server.http3.client_ssl_profile` | [virtual_server.http3.client_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-1120110303201101-1301001103033203-2001103223211123-1023320032013330-1112100022313220-3011010013203212-3110122001132223-3111011223120333) |
| `virtual_server.http3.client_ssl_profile.kind` | [virtual_server.http3.client_ssl_profile.kind](resources--application_profiles--reference--group-002.md#canonical-2200200203103102-3133013030202230-1021203002003221-2110333022012111-3110302013321330-1331302101000003-1311221130300200-1333203223031111) |
| `virtual_server.http3.client_ssl_profile.name` | [virtual_server.http3.client_ssl_profile.name](resources--application_profiles--reference--group-002.md#canonical-3221132330131232-2332303212210212-0110223210030001-1211100200023300-2132000210320310-1133332222003221-0013313032201203-0121233203313220) |
| `virtual_server.http3.client_ssl_profile.namespace` | [virtual_server.http3.client_ssl_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-1031323230323132-3303012013032303-3003133221121022-2022000000202030-3203111232200222-1311120211232033-0021311303102132-3021101232132321) |
| `virtual_server.http3.client_ssl_profile.tenant` | [virtual_server.http3.client_ssl_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-1133102131011033-0311201313100103-3322323331312302-0223000103022232-2320103302021210-0011233002222300-3202321012210232-3221020223320230) |
| `virtual_server.http3.client_ssl_profile.uid` | [virtual_server.http3.client_ssl_profile.uid](resources--application_profiles--reference--group-002.md#canonical-2133030102220021-3331133303233312-0010033013023023-1221011200002020-0232203313301020-0021310222210131-3203120311121221-3010003211132132) |
| `virtual_server.http3.http3_profile` | [virtual_server.http3.http3_profile](resources--application_profiles--reference--group-002.md#canonical-0322200003130310-2100000313011310-3310110111203133-0322333010302213-2113101012322222-3021021122212110-1320200333300022-1010100310322312) |
| `virtual_server.http3.http3_profile.kind` | [virtual_server.http3.http3_profile.kind](resources--application_profiles--reference--group-002.md#canonical-0223023331021000-2223031010030023-2311021301121001-1130222000011003-2303133002113130-0102302331300022-2202122200232233-1221133222302313) |
| `virtual_server.http3.http3_profile.name` | [virtual_server.http3.http3_profile.name](resources--application_profiles--reference--group-002.md#canonical-1222200231311100-1101302000230030-2231211101101303-2310111330033023-2222021031303132-2212333002031030-0300302330332212-1333010032233210) |
| `virtual_server.http3.http3_profile.namespace` | [virtual_server.http3.http3_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-0322112021132133-3021032010020012-2023301002300111-2202223300111312-0010102130331332-3312122322132123-3221222230131233-2231132201033101) |
| `virtual_server.http3.http3_profile.tenant` | [virtual_server.http3.http3_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-1022330111013223-3102311311132110-1012010030003301-0120212110223111-3012101001002300-3101120312120021-0111103220303012-3131302201321022) |
| `virtual_server.http3.http3_profile.uid` | [virtual_server.http3.http3_profile.uid](resources--application_profiles--reference--group-002.md#canonical-1322323202021031-2012010100131333-2233031213321333-0123121021232122-3331333123011201-2133232131310333-0110012022023220-2030022300001301) |
| `virtual_server.http3.http_client_profile` | [virtual_server.http3.http_client_profile](resources--application_profiles--reference--group-002.md#canonical-2111231120012330-1233222013020212-2013003201033331-3302211331120231-1313132022313113-0110012313212012-0131120302222113-2233021322200321) |
| `virtual_server.http3.http_client_profile.kind` | [virtual_server.http3.http_client_profile.kind](resources--application_profiles--reference--group-002.md#canonical-1232323312312320-0203200313010222-0202101012120212-3320021331120322-3212201230230303-0132321030302112-2332132213012301-1300221230320211) |
| `virtual_server.http3.http_client_profile.name` | [virtual_server.http3.http_client_profile.name](resources--application_profiles--reference--group-002.md#canonical-3133323121002120-0220100010031021-1132302332013120-0001301021111323-3222323212312302-0331213103333032-0133000103102232-1300003332222123) |
| `virtual_server.http3.http_client_profile.namespace` | [virtual_server.http3.http_client_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-1123133003010012-1021313333201313-1003311301312113-3133120023301203-2233121221331021-2203201101011233-3120010332132013-3211111122131330) |
| `virtual_server.http3.http_client_profile.tenant` | [virtual_server.http3.http_client_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-0321320130021102-3330121302313022-1001312023200033-3023322222003110-1221000110301300-3123222120321033-3000012022011133-0210012333233313) |
| `virtual_server.http3.http_client_profile.uid` | [virtual_server.http3.http_client_profile.uid](resources--application_profiles--reference--group-002.md#canonical-1102013031132323-0002311212321333-0011103210021022-2100323313321233-1202311320131202-1012003310103111-2113202221132212-3030221000201110) |
| `virtual_server.http3.http_server_profile` | [virtual_server.http3.http_server_profile](resources--application_profiles--reference--group-002.md#canonical-2133130302111323-0311332113010102-2203021301212221-3311003113123201-2002013122221310-0203221012012023-0203001021303203-1012221323112022) |
| `virtual_server.http3.http_server_profile.kind` | [virtual_server.http3.http_server_profile.kind](resources--application_profiles--reference--group-002.md#canonical-2302321123330311-2011003131132333-0323321010202100-0021123221310302-3032012110320221-0331213222200032-3232003103320030-1120123022311111) |
| `virtual_server.http3.http_server_profile.name` | [virtual_server.http3.http_server_profile.name](resources--application_profiles--reference--group-002.md#canonical-2213310223202332-0121320111231201-3003031213003203-2333233003010103-0313000100201311-0013013223110221-2223330033223010-1002031312311323) |
| `virtual_server.http3.http_server_profile.namespace` | [virtual_server.http3.http_server_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-1022320011002333-0023330332201230-3113033012211223-1020201320031100-3012002112022130-3031220222302121-0333011313030101-3013021320020321) |
| `virtual_server.http3.http_server_profile.tenant` | [virtual_server.http3.http_server_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-2031231210110123-0203123100120331-3000311221212021-1120222023232111-1200323023213332-3030321303333012-0210333021313222-1031111112113330) |
| `virtual_server.http3.http_server_profile.uid` | [virtual_server.http3.http_server_profile.uid](resources--application_profiles--reference--group-002.md#canonical-3330310132210002-0302201100211032-3012312202022322-0211103022100022-0310313332103213-3302222311122223-3312010012100002-0030313103221213) |
| `virtual_server.http3.quic_profile` | [virtual_server.http3.quic_profile](resources--application_profiles--reference--group-002.md#canonical-1302232321131103-0021111103200313-2331011300003103-0001032113331103-2123300232330232-0202033332003011-0301212020002233-3203103023000203) |
| `virtual_server.http3.quic_profile.kind` | [virtual_server.http3.quic_profile.kind](resources--application_profiles--reference--group-002.md#canonical-2103301310301303-2213000303330332-1210031331121212-0211120322231010-2002323220303330-3322331223133220-2021000022032210-0200001200132003) |
| `virtual_server.http3.quic_profile.name` | [virtual_server.http3.quic_profile.name](resources--application_profiles--reference--group-002.md#canonical-1011131312312103-2332110030312221-1332211112000033-2122120023202330-0030201011213303-1223101230202303-1102102121010022-1013202121303000) |
| `virtual_server.http3.quic_profile.namespace` | [virtual_server.http3.quic_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-3231202020111232-2023212123121022-0013223123210310-1000103102302133-1001211103302122-1233230200223323-3022002231312230-2320032132321003) |
| `virtual_server.http3.quic_profile.tenant` | [virtual_server.http3.quic_profile.tenant](resources--application_profiles--reference--group-002.md#canonical-0331231031100010-3132103210320110-0022122310201313-3321332303100102-3011332133013120-0320112123112312-0310101003223203-1010123301102220) |
| `virtual_server.http3.quic_profile.uid` | [virtual_server.http3.quic_profile.uid](resources--application_profiles--reference--group-002.md#canonical-1010230110221332-1202213231322301-2211100303001133-1233022232001102-3013100232223000-1131211102001003-3020100002231201-1313212101102330) |
| `virtual_server.http3.server_ssl_profile` | [virtual_server.http3.server_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-3011020032120132-2302133210113203-2231330200220020-2011300122312220-2130310201130222-1203123211131010-0021313120101311-3230333210021311) |
| `virtual_server.http3.server_ssl_profile.kind` | [virtual_server.http3.server_ssl_profile.kind](resources--application_profiles--reference--group-002.md#canonical-3123101132120311-1121210230331331-1233123311113011-3322112010310111-0002012322121331-3020203322032132-3113233321330202-1312213011201132) |
| `virtual_server.http3.server_ssl_profile.name` | [virtual_server.http3.server_ssl_profile.name](resources--application_profiles--reference--group-002.md#canonical-3303131301110201-1021303100323033-0101202321000101-3210132321223212-3200121332322323-1033002303110310-2002211221031022-2132320312200213) |
| `virtual_server.http3.server_ssl_profile.namespace` | [virtual_server.http3.server_ssl_profile.namespace](resources--application_profiles--reference--group-002.md#canonical-3330222313300323-1300303111103222-1212310233133110-2310133303300312-0010323332131012-3320202010313020-1311231111223120-0230231231131333) |
| `virtual_server.http3.server_ssl_profile.tenant` | [virtual_server.http3.server_ssl_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-2330313112311101-2323331322211132-3032313321111113-2133000113202002-2300220332223110-2230130012233031-2010002010020222-0003033320100202) |
| `virtual_server.http3.server_ssl_profile.uid` | [virtual_server.http3.server_ssl_profile.uid](resources--application_profiles--reference--group-003.md#canonical-0203003212200112-3310300301310211-2323330302211331-3322110132233122-1003021030333111-1012130322132301-0003003311211121-0001103031022320) |
| `virtual_server.http3.tcp_server_profile` | [virtual_server.http3.tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-2112022202030023-0100133022230321-0003322233201012-1233302133322131-1221012120113011-1300301322102201-3220232232012300-2313211223230131) |
| `virtual_server.http3.tcp_server_profile.kind` | [virtual_server.http3.tcp_server_profile.kind](resources--application_profiles--reference--group-003.md#canonical-3233212222020113-2013030122200212-1312021020332233-2213131313303312-2022033212012200-1332231110032222-2330112220031320-1203213110201331) |
| `virtual_server.http3.tcp_server_profile.name` | [virtual_server.http3.tcp_server_profile.name](resources--application_profiles--reference--group-003.md#canonical-1313331100013012-1030031223311021-3012301331111302-3102010301010222-2331303203022011-0123303021133010-1311032111222033-1321301300000111) |
| `virtual_server.http3.tcp_server_profile.namespace` | [virtual_server.http3.tcp_server_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-1222310012213002-3123101310310103-1303100221030210-3010102203123330-3223033230301311-3312210132221300-0223301222313211-0320132001222301) |
| `virtual_server.http3.tcp_server_profile.tenant` | [virtual_server.http3.tcp_server_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-3302321003121013-1303320221101203-3330120030100003-3322220212132311-2131210331022232-2123000202110111-1112211223002332-3120001103303301) |
| `virtual_server.http3.tcp_server_profile.uid` | [virtual_server.http3.tcp_server_profile.uid](resources--application_profiles--reference--group-003.md#canonical-2011231123300211-1110331203033123-3133331233211012-3302200221020313-0202100131330100-3311102111013213-2021210000213321-1302330013220212) |
| `virtual_server.http3.udp_client_profile` | [virtual_server.http3.udp_client_profile](resources--application_profiles--reference--group-003.md#canonical-1301232231102003-3331013000233123-3210231321202022-1323203122210233-1102103310230133-3132113123220232-2201022010231232-0120001122103011) |
| `virtual_server.http3.udp_client_profile.kind` | [virtual_server.http3.udp_client_profile.kind](resources--application_profiles--reference--group-003.md#canonical-3013223112100102-3013001111311323-1231002102311223-0011023031023330-3213100330021322-0112130103212321-3230013311130223-1203212231210101) |
| `virtual_server.http3.udp_client_profile.name` | [virtual_server.http3.udp_client_profile.name](resources--application_profiles--reference--group-003.md#canonical-1121320020312332-1110231313033021-3303210021322300-2102203323220310-1021020102232111-3231001101021221-1133231201032011-1322313032310000) |
| `virtual_server.http3.udp_client_profile.namespace` | [virtual_server.http3.udp_client_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-2322102203111312-1033312323301130-2003312130120211-0130002323321213-2010113120023333-3330323300321233-0112113203200133-1030200000111310) |
| `virtual_server.http3.udp_client_profile.tenant` | [virtual_server.http3.udp_client_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-1103011102221312-1320032331123122-3213233021233110-0223223312220023-2120020230011321-2031131321131120-0222301222320002-3210233123322330) |
| `virtual_server.http3.udp_client_profile.uid` | [virtual_server.http3.udp_client_profile.uid](resources--application_profiles--reference--group-003.md#canonical-1011023113133001-0310312311131200-1032221021310230-0213203012203220-0232131313113202-0030201032312003-3311211120310032-3202202133302212) |
| `virtual_server.http3.udp_server_profile` | [virtual_server.http3.udp_server_profile](resources--application_profiles--reference--group-003.md#canonical-2101303221030311-3333130023032321-2133221203131113-0122010021103001-3012322030021323-1021120122000203-0313111000212222-0032022121330320) |
| `virtual_server.http3.udp_server_profile.kind` | [virtual_server.http3.udp_server_profile.kind](resources--application_profiles--reference--group-003.md#canonical-0000131221231020-1310212223312213-0331013132010110-1311233302130323-3013200313320131-3223120132012132-1120120113023012-1221233222201021) |
| `virtual_server.http3.udp_server_profile.name` | [virtual_server.http3.udp_server_profile.name](resources--application_profiles--reference--group-003.md#canonical-2300230332231111-1013120233103012-0203213202332000-1232000321230111-2212310102023203-2001023000301211-3120203330302323-0330122212330033) |
| `virtual_server.http3.udp_server_profile.namespace` | [virtual_server.http3.udp_server_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-0003000132010202-3210313322011303-2310323320002310-0322300310101032-2030311333313110-2021322120301300-3203311300000122-3021203332112323) |
| `virtual_server.http3.udp_server_profile.tenant` | [virtual_server.http3.udp_server_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-2222313230003331-2033032001033011-0320333200333303-1130122013322021-0020211312201313-2000301122120130-3122220220333010-2012302012302320) |
| `virtual_server.http3.udp_server_profile.uid` | [virtual_server.http3.udp_server_profile.uid](resources--application_profiles--reference--group-003.md#canonical-0203011010032210-2131222022303211-3221221203322223-3230300330201001-1010302121021133-0110123110333002-2232000111331000-0012232023221200) |
| `virtual_server.https` | [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-3001023313121321-0102311133101033-0023030120103121-2330130013111333-2123113113011130-0202212203232022-1123300210111113-1023120132333222) |
| `virtual_server.https.client_ssl_profile` | [virtual_server.https.client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-0310121121023300-0202320010121122-2230211332011001-3202032212312211-1102322000011001-0120031112230010-0003210133022323-3330320012220231) |
| `virtual_server.https.client_ssl_profile.kind` | [virtual_server.https.client_ssl_profile.kind](resources--application_profiles--reference--group-003.md#canonical-3032231110110323-1033321021030322-0123301303001130-3223330130020011-3221120132020132-1223221002031223-2210010320200032-3313013112301010) |
| `virtual_server.https.client_ssl_profile.name` | [virtual_server.https.client_ssl_profile.name](resources--application_profiles--reference--group-003.md#canonical-0003003301000222-0301100210030302-3132310302032101-0220320330301001-3021220010212333-3131332111311002-0220322002310223-3032220023200211) |
| `virtual_server.https.client_ssl_profile.namespace` | [virtual_server.https.client_ssl_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-1311323220110203-2133123021312233-2023032031232113-3203001331011010-0132130330102233-1101300101220331-1031203301132233-3201010223100210) |
| `virtual_server.https.client_ssl_profile.tenant` | [virtual_server.https.client_ssl_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-2022202320302123-1133322223320312-1230332000322123-3022112012223322-2212321330310320-1100322000232323-2221111330332110-0021110023012030) |
| `virtual_server.https.client_ssl_profile.uid` | [virtual_server.https.client_ssl_profile.uid](resources--application_profiles--reference--group-003.md#canonical-2223310102001111-0022200133320100-0231300013011313-0311010133030002-3211032220131323-1313213312231102-3331320220023212-2210302232123231) |
| `virtual_server.https.http2_client_profile` | [virtual_server.https.http2_client_profile](resources--application_profiles--reference--group-003.md#canonical-1003123300312211-1200121022121330-1013021323113111-0210100223202000-2320003131022321-2223221010130012-2100302320320122-2133311112211120) |
| `virtual_server.https.http2_client_profile.kind` | [virtual_server.https.http2_client_profile.kind](resources--application_profiles--reference--group-003.md#canonical-0020210321321021-2030112332221211-2120011022223230-0311333010303030-2330123101122300-1220123112032232-1123022003131133-1331113200330021) |
| `virtual_server.https.http2_client_profile.name` | [virtual_server.https.http2_client_profile.name](resources--application_profiles--reference--group-003.md#canonical-3122002103220133-0131023110022301-0321122302302113-0003311233310333-3332032033033013-1032010113313330-0322012321230310-0300222310200331) |
| `virtual_server.https.http2_client_profile.namespace` | [virtual_server.https.http2_client_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-1211331331323102-0210322300233311-3033002012102113-1121231111200120-2102220120111131-1221011020303012-2102131130211330-0123321132001200) |
| `virtual_server.https.http2_client_profile.tenant` | [virtual_server.https.http2_client_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-1322300310122122-0003132023032102-1033301113123323-2021133331010012-1131013212021122-2333023300110121-1203302301210210-1302330200320033) |
| `virtual_server.https.http2_client_profile.uid` | [virtual_server.https.http2_client_profile.uid](resources--application_profiles--reference--group-003.md#canonical-1103130313031020-3223210221132310-2031233210210213-2200232120203323-3001232212322123-3123211010323320-1302003030312312-1001020323122201) |
| `virtual_server.https.http2_server_profile` | [virtual_server.https.http2_server_profile](resources--application_profiles--reference--group-003.md#canonical-1301311030022121-3200103330220220-3210233112201320-0010303121231223-0313222113100020-3112130201100122-3322123010223020-1002303231002203) |
| `virtual_server.https.http2_server_profile.kind` | [virtual_server.https.http2_server_profile.kind](resources--application_profiles--reference--group-003.md#canonical-1113211231233032-0222313101010330-2331111002231120-0232200310102012-3230200321113121-1320033113303311-1211123020303130-0123033131000123) |
| `virtual_server.https.http2_server_profile.name` | [virtual_server.https.http2_server_profile.name](resources--application_profiles--reference--group-003.md#canonical-3023233033020131-1321311130111021-2331012103212023-0013301133200031-2211210012330022-3322300200021012-0021231010122232-3021231002013233) |
| `virtual_server.https.http2_server_profile.namespace` | [virtual_server.https.http2_server_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-2301320332123011-1301313120311322-0000211013023311-0223020320313313-1002031022231321-2133313022023200-2103102102320121-1110311123032110) |
| `virtual_server.https.http2_server_profile.tenant` | [virtual_server.https.http2_server_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-0320221013202203-2010221331110003-3203132200302111-3020303301233130-3202231323130001-2023012112122322-0020320110203112-1020232301110003) |
| `virtual_server.https.http2_server_profile.uid` | [virtual_server.https.http2_server_profile.uid](resources--application_profiles--reference--group-003.md#canonical-0103301103002230-0031310120131100-0331022101323301-3211200320000320-0110100321023120-2333130113200321-1310221113101303-3300232120132233) |
| `virtual_server.https.http_client_profile` | [virtual_server.https.http_client_profile](resources--application_profiles--reference--group-003.md#canonical-0032323313122011-1113013231233012-1211201132213211-1120003003232021-1113331031023332-2003213113121200-0303213030113103-0221312101011203) |
| `virtual_server.https.http_client_profile.kind` | [virtual_server.https.http_client_profile.kind](resources--application_profiles--reference--group-003.md#canonical-2023113210200000-0200112222213111-0231203130300021-0131103012301303-1233202032321121-2222322012120130-1300000323021303-2103213221211221) |
| `virtual_server.https.http_client_profile.name` | [virtual_server.https.http_client_profile.name](resources--application_profiles--reference--group-003.md#canonical-0303323220213020-2000222113103230-0223320012112210-3230021003123010-3131020110011331-3001220030012321-1131211301030010-1331231133232233) |
| `virtual_server.https.http_client_profile.namespace` | [virtual_server.https.http_client_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-0022233331201131-2312323210310333-1233301230322012-3231321322022312-3120100003010302-3002200323200103-0220321012113221-0301132013330101) |
| `virtual_server.https.http_client_profile.tenant` | [virtual_server.https.http_client_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-2021313123201301-3203133211212100-1320023310123102-0211302230010300-0323012020300220-2130133103331020-0313321003222233-1103022003202320) |
| `virtual_server.https.http_client_profile.uid` | [virtual_server.https.http_client_profile.uid](resources--application_profiles--reference--group-003.md#canonical-2233033110310331-2303202013321113-3112300221103022-0223232133221001-0130323302312003-2013100003333210-1103321033320232-0322230010012321) |
| `virtual_server.https.http_server_profile` | [virtual_server.https.http_server_profile](resources--application_profiles--reference--group-003.md#canonical-0022113000123221-1310311301020202-3033300230312220-3112103313120203-3220123203100231-2233231201330130-2013212103023230-2303102003312012) |
| `virtual_server.https.http_server_profile.kind` | [virtual_server.https.http_server_profile.kind](resources--application_profiles--reference--group-003.md#canonical-3203212101112002-2302130332100001-2312322011031023-3123133031311232-0031333131031233-1102130110320233-0002102301013232-2202131022300202) |
| `virtual_server.https.http_server_profile.name` | [virtual_server.https.http_server_profile.name](resources--application_profiles--reference--group-003.md#canonical-0203320031122111-0012323323321330-0232021203130212-1333013233130302-3221201132300102-2112331122111320-1320322122120031-1303201303101220) |
| `virtual_server.https.http_server_profile.namespace` | [virtual_server.https.http_server_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-0301311010130332-2013223110331003-2302223131330010-1000133101312133-3033012102320203-1313133132032131-2300021120322213-3333203323131311) |
| `virtual_server.https.http_server_profile.tenant` | [virtual_server.https.http_server_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-1033322130333333-3310212312100302-1222121103220102-3011321303001233-2010212113130103-0210303033332102-2001213320123001-0120101101032032) |
| `virtual_server.https.http_server_profile.uid` | [virtual_server.https.http_server_profile.uid](resources--application_profiles--reference--group-003.md#canonical-0332102231332221-1010303130330101-1123001332332033-2302311132213211-2233102332022310-3332221230333310-1033323303233022-0030111223231310) |
| `virtual_server.https.ocsp_profile` | [virtual_server.https.ocsp_profile](resources--application_profiles--reference--group-003.md#canonical-1132012023312201-2221332332121311-0112130333130300-0020032030200233-3313013312303212-1330022213030112-2313113303201211-0103121230002203) |
| `virtual_server.https.ocsp_profile.kind` | [virtual_server.https.ocsp_profile.kind](resources--application_profiles--reference--group-003.md#canonical-0331303310021132-3120032200223020-3333010133320103-0333233232201121-2322310013000031-2110223020321130-2201332003010223-2321120231031112) |
| `virtual_server.https.ocsp_profile.name` | [virtual_server.https.ocsp_profile.name](resources--application_profiles--reference--group-003.md#canonical-3230223322021133-0322233321110023-2103131011132232-0230121003000232-3033332313333333-1023332232331022-2311212200220013-3003223211220211) |
| `virtual_server.https.ocsp_profile.namespace` | [virtual_server.https.ocsp_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-0200000111002232-0131311002102013-1013002000322012-3301003301302030-2210110001102331-0232022112113211-1320020222210102-0323100232321312) |
| `virtual_server.https.ocsp_profile.tenant` | [virtual_server.https.ocsp_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-2022323331132330-0120331103211023-0233322101011010-0311120122030322-2223033322002110-2002132003022002-2033223321312220-2011311222000203) |
| `virtual_server.https.ocsp_profile.uid` | [virtual_server.https.ocsp_profile.uid](resources--application_profiles--reference--group-003.md#canonical-0201011233111201-3003223310323100-0301020223011200-0223333112123322-1120321301331013-0222011321001200-2310330030023103-0321301122013302) |
| `virtual_server.https.server_ssl_profile` | [virtual_server.https.server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-1302131030322013-3023030122211333-1303321201211310-3023221200331023-3301001001112232-3322321222032130-3131022020013032-2011300332200312) |
| `virtual_server.https.server_ssl_profile.kind` | [virtual_server.https.server_ssl_profile.kind](resources--application_profiles--reference--group-003.md#canonical-0121330131230122-2132203220311032-2203212221212202-1231331211233321-1232022031133022-2010321122311320-3301002322212330-2113103130122030) |
| `virtual_server.https.server_ssl_profile.name` | [virtual_server.https.server_ssl_profile.name](resources--application_profiles--reference--group-003.md#canonical-3011030022032230-1232001313203022-1020123021201331-2022011220011003-2003220100103320-1222302222100331-3302101302312002-0020300312000130) |
| `virtual_server.https.server_ssl_profile.namespace` | [virtual_server.https.server_ssl_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-0011221132001012-1112311330330132-0322031302020303-2002211120002132-0012323003303223-2323003000001203-1211210300223012-0030333133311120) |
| `virtual_server.https.server_ssl_profile.tenant` | [virtual_server.https.server_ssl_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-0310330122030031-0101222311320002-1030110132331323-2303313130222120-3111101011303331-0121313301233201-1320011100133302-1013032111302132) |
| `virtual_server.https.server_ssl_profile.uid` | [virtual_server.https.server_ssl_profile.uid](resources--application_profiles--reference--group-003.md#canonical-0302101201032331-3220323320312022-0010120103003331-0312000331223203-3101231030132310-3113322310313322-3310222102210333-1220222303022232) |
| `virtual_server.https.stream_profile` | [virtual_server.https.stream_profile](resources--application_profiles--reference--group-003.md#canonical-1323133221230101-1211220021102331-0311111233130122-3331120312301222-3323200113222330-1002010112200121-2320201121033111-1031323222012022) |
| `virtual_server.https.stream_profile.kind` | [virtual_server.https.stream_profile.kind](resources--application_profiles--reference--group-003.md#canonical-1100212030322331-1233021212131001-3203103033222110-0312311010323000-3003033133101213-0032013230210332-2011103220222221-3300000123311330) |
| `virtual_server.https.stream_profile.name` | [virtual_server.https.stream_profile.name](resources--application_profiles--reference--group-003.md#canonical-3211331103210323-2221332333220231-3201101220133002-3003330330201213-3332332132020110-0003231230031121-2011321101120303-3032200311112121) |
| `virtual_server.https.stream_profile.namespace` | [virtual_server.https.stream_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-3101131301333110-2002313032232303-3321100022111132-1222010332102301-0102222223130221-0330330012032010-0210332100223200-1030130331021130) |
| `virtual_server.https.stream_profile.tenant` | [virtual_server.https.stream_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-2303132231321321-2323232030330122-0033333030001002-3231312310010231-3120002201121233-0313033030003223-3303322230230330-0111332030130232) |
| `virtual_server.https.stream_profile.uid` | [virtual_server.https.stream_profile.uid](resources--application_profiles--reference--group-003.md#canonical-2300211022301120-1311323011231211-2332020310032031-2023203333203001-2101302313213330-1222121123103302-2310010011232003-2301123322331020) |
| `virtual_server.https.tcp_client_profile` | [virtual_server.https.tcp_client_profile](resources--application_profiles--reference--group-003.md#canonical-3303020301320113-0013112032021200-1322222023000302-2120131003333003-0321131323223231-3111231031023213-3310322101031110-3311133000000013) |
| `virtual_server.https.tcp_client_profile.kind` | [virtual_server.https.tcp_client_profile.kind](resources--application_profiles--reference--group-003.md#canonical-3032300212330101-2020003032001222-2321233223030001-0031220233120212-1033211301113110-0130301220303322-0002023030331023-3232012223031132) |
| `virtual_server.https.tcp_client_profile.name` | [virtual_server.https.tcp_client_profile.name](resources--application_profiles--reference--group-003.md#canonical-0201133010110031-2110033022201102-3032130132113332-0321133230123002-1211121022122113-2230033223023021-2032311302311230-3010331321330103) |
| `virtual_server.https.tcp_client_profile.namespace` | [virtual_server.https.tcp_client_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-2003221130112030-2123320021131023-1001330331110302-2311013101330310-2033020220213132-1102200102312230-0101121120323213-2300121302010132) |
| `virtual_server.https.tcp_client_profile.tenant` | [virtual_server.https.tcp_client_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-3022200102020212-2102312321101213-3332122213120300-3030201033213310-1303231031001323-3123211112200022-0103303321333321-3331211000121300) |
| `virtual_server.https.tcp_client_profile.uid` | [virtual_server.https.tcp_client_profile.uid](resources--application_profiles--reference--group-003.md#canonical-1132330013210000-2001200002021010-2203022121320113-0321020220212100-0230320312221200-3020202033303133-3220122133111023-0031301233003033) |
| `virtual_server.https.tcp_server_profile` | [virtual_server.https.tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-2210332031001101-3032113311222310-0111231003010232-0333222211331022-3130220330300311-0323123333023213-3122230031130333-0322301330020202) |
| `virtual_server.https.tcp_server_profile.kind` | [virtual_server.https.tcp_server_profile.kind](resources--application_profiles--reference--group-003.md#canonical-2123101320202333-1230313111120003-3321213300233003-0333021020221313-0122301022221021-2011302311110330-2310110030312331-1202330211330231) |
| `virtual_server.https.tcp_server_profile.name` | [virtual_server.https.tcp_server_profile.name](resources--application_profiles--reference--group-003.md#canonical-1211213332131223-3321032031122311-1302230121123021-1110132313232011-3021320321330033-0230000301011221-3023032210133220-0033132122300102) |
| `virtual_server.https.tcp_server_profile.namespace` | [virtual_server.https.tcp_server_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-0223010133022003-0003310213232130-2003103333303231-0330000211011332-1331110033200231-2113210322321010-2011001012221320-2200330310211322) |
| `virtual_server.https.tcp_server_profile.tenant` | [virtual_server.https.tcp_server_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-3332301021033311-2001013131001110-1310310331323133-0333313333111103-1020230100232310-3213023121121103-1233001321330130-2122023113332033) |
| `virtual_server.https.tcp_server_profile.uid` | [virtual_server.https.tcp_server_profile.uid](resources--application_profiles--reference--group-003.md#canonical-2303310310321332-2011003112300000-2223123033322230-0021233113230202-1111111213121203-1000020113011011-1300203323021233-0101033021023022) |
| `virtual_server.https.websocket_client_profile` | [virtual_server.https.websocket_client_profile](resources--application_profiles--reference--group-003.md#canonical-0321132123300232-1313031133033202-3120303020320213-2103110113311202-3101300112300221-2302231110003323-3111332212002330-0213011031011312) |
| `virtual_server.https.websocket_client_profile.kind` | [virtual_server.https.websocket_client_profile.kind](resources--application_profiles--reference--group-003.md#canonical-1230101103122213-2220322121321111-3110001101012202-0213003332023021-0001321201121302-0231013122310220-3332231112010002-2203130122311302) |
| `virtual_server.https.websocket_client_profile.name` | [virtual_server.https.websocket_client_profile.name](resources--application_profiles--reference--group-003.md#canonical-3033102210031010-2220102133122020-0131311020202300-1032310000002131-2000132121113122-1331110023120303-2132220233022210-0312331110020210) |
| `virtual_server.https.websocket_client_profile.namespace` | [virtual_server.https.websocket_client_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-1223302120121302-2000212220312303-1013001023021311-1330223230130130-1300112122300003-1020220232230130-3110333110113001-0210310321100201) |
| `virtual_server.https.websocket_client_profile.tenant` | [virtual_server.https.websocket_client_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-0031233030311313-3233012110320121-2133020221203023-0321220300100330-3000010113121113-1231133223320003-3131320303322210-3133300213302100) |
| `virtual_server.https.websocket_client_profile.uid` | [virtual_server.https.websocket_client_profile.uid](resources--application_profiles--reference--group-003.md#canonical-0202100213232022-2112002130110032-1233320311121300-3300312012203332-2322110313201220-0232102103220211-2120313022103033-0302001030300111) |
| `virtual_server.https.websocket_server_profile` | [virtual_server.https.websocket_server_profile](resources--application_profiles--reference--group-003.md#canonical-1211103303022013-3011120031202312-3312010310110313-0331213122130111-0103113120321002-0100023210232022-3213113010130310-3303210301033233) |
| `virtual_server.https.websocket_server_profile.kind` | [virtual_server.https.websocket_server_profile.kind](resources--application_profiles--reference--group-003.md#canonical-2002132132300230-0201310132212033-1322112023211312-3023130003100103-0032230023013302-3202130322133201-0032303112130113-3220211000320232) |
| `virtual_server.https.websocket_server_profile.name` | [virtual_server.https.websocket_server_profile.name](resources--application_profiles--reference--group-003.md#canonical-3103101222123011-0220212113302331-0203032132132012-3213123211011220-2011233033232100-2131022112212000-1123201031011012-2303232003210310) |
| `virtual_server.https.websocket_server_profile.namespace` | [virtual_server.https.websocket_server_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-0213002233302131-0000211332221202-3301223000223001-3332223012330000-3312002103321233-3033122021300301-0201032310203311-0332030202202133) |
| `virtual_server.https.websocket_server_profile.tenant` | [virtual_server.https.websocket_server_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-1023131000113303-0330032322331212-3210211211332231-2003202330121003-3301000110113230-2220122100102220-2330201220000131-3011121012130012) |
| `virtual_server.https.websocket_server_profile.uid` | [virtual_server.https.websocket_server_profile.uid](resources--application_profiles--reference--group-003.md#canonical-3131223303321333-0012221311011021-3031000011303211-0111313020323322-1030020110220032-3133220121310033-2231220333020022-2312010310032002) |
| `virtual_server.immediate_action_on_service_down` | [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-3230213032221222-0013333133210033-1222132013313112-2322100010110313-1330313131112110-0230121231012010-0232033001322133-0310202132120231) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop](resources--application_profiles--reference--group-003.md#canonical-1221013321201031-0111102300222301-3311010131213030-0001200212301332-0200022011103130-1230030322301001-1210102033020203-0310010232233012) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none](resources--application_profiles--reference--group-003.md#canonical-2003131000131233-3100223321020233-1201100121133302-3200203320011111-0112132020230001-2020331313201112-0001223210133112-2031122322303203) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset](resources--application_profiles--reference--group-003.md#canonical-3302211020211201-1110322130010101-0223133020003013-1123322322203021-3312323132031330-3200020132302312-0000211003333112-3033230302112203) |
| `virtual_server.last_hop_pool` | [virtual_server.last_hop_pool](resources--application_profiles--reference--group-003.md#canonical-3130100312210212-0020201222232323-0102223330023013-2103121210010212-2120033023232203-2312233331101003-1200120122133211-3232032330003110) |
| `virtual_server.last_hop_pool.kind` | [virtual_server.last_hop_pool.kind](resources--application_profiles--reference--group-003.md#canonical-0133220332112203-2120203331023303-0103103201001100-0020310302031313-1322121321222110-2223230321020200-2323310102331130-2110322010101133) |
| `virtual_server.last_hop_pool.name` | [virtual_server.last_hop_pool.name](resources--application_profiles--reference--group-003.md#canonical-0131123000130202-2001333301303310-3310133113312212-2211331301120031-3021020201333212-1023003310010321-3102323310010113-1303302321031230) |
| `virtual_server.last_hop_pool.namespace` | [virtual_server.last_hop_pool.namespace](resources--application_profiles--reference--group-003.md#canonical-3110220010333333-0132110232200320-2333301323012210-1012002310320022-1210302020303131-0312011232331102-3203201301113230-2213010313021030) |
| `virtual_server.last_hop_pool.tenant` | [virtual_server.last_hop_pool.tenant](resources--application_profiles--reference--group-003.md#canonical-1020121201321111-0321222303220311-3031220120003133-1023233230022012-2021133322232211-3002130130133210-0313013101231031-1310120223223031) |
| `virtual_server.last_hop_pool.uid` | [virtual_server.last_hop_pool.uid](resources--application_profiles--reference--group-003.md#canonical-2003030013011001-2001200312132330-3012223122303112-2103222123001133-3212331021031331-1222010300102330-0130102200233130-0100223322332213) |
| `virtual_server.nat64` | [virtual_server.nat64](resources--application_profiles--reference--group-003.md#canonical-1003101233121130-3210201232023112-3310130202030221-2000220312333131-1103301120112202-2232110110103211-3021002033300032-3230222222320030) |
| `virtual_server.nat64.nat64_disable` | [virtual_server.nat64.nat64_disable](resources--application_profiles--reference--group-003.md#canonical-0032200222113123-3223310301322110-3102323310301321-0323320301120132-3213301222232333-1303323203000301-1230002321301113-0200223100132103) |
| `virtual_server.nat64.nat64_enable` | [virtual_server.nat64.nat64_enable](resources--application_profiles--reference--group-003.md#canonical-3213312001302303-0123300102331001-2011233202113212-3023132011020311-0222123030123111-1023322311011003-1003301320323332-1231200122323330) |
| `virtual_server.port_translation` | [virtual_server.port_translation](resources--application_profiles--reference--group-003.md#canonical-3300211131303022-0213330231302003-0323211023223112-0221132220220312-1121333121213210-2133331232122332-0010300301232312-0122121000212121) |
| `virtual_server.port_translation.port_translation_disable` | [virtual_server.port_translation.port_translation_disable](resources--application_profiles--reference--group-003.md#canonical-0331201330112220-3322000201302223-3302323211013310-1011022233300321-0102203301310010-0301110102203200-0222131221201313-0130021112300230) |
| `virtual_server.port_translation.port_translation_enable` | [virtual_server.port_translation.port_translation_enable](resources--application_profiles--reference--group-003.md#canonical-2110210322123011-1100230132320211-0101312110330201-0010313202120310-2302300301132231-3303200010330110-1020120100133223-0330011233231021) |
| `virtual_server.request_logging_profile` | [virtual_server.request_logging_profile](resources--application_profiles--reference--group-003.md#canonical-3021100321112301-1233301012311200-0312330023111200-2001301020303311-2233302130132031-2113023101013111-0133021233213323-3232312330213023) |
| `virtual_server.request_logging_profile.kind` | [virtual_server.request_logging_profile.kind](resources--application_profiles--reference--group-003.md#canonical-2233232030023313-3133202323110131-3131322232122022-1020123321200333-2321211132333330-0330210122202000-2232233121020312-1121133312313321) |
| `virtual_server.request_logging_profile.name` | [virtual_server.request_logging_profile.name](resources--application_profiles--reference--group-003.md#canonical-2330113122030033-3333310220311332-0223021101213303-1323103321221121-3201211312121333-0033032101131310-2332022230133023-3322223132003220) |
| `virtual_server.request_logging_profile.namespace` | [virtual_server.request_logging_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-0233120103100302-1311201110110101-1120331232001210-0000121221300121-2120013031003101-0213331233212322-0010322230122333-3123123332012210) |
| `virtual_server.request_logging_profile.tenant` | [virtual_server.request_logging_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-0203102211223131-3103102013132000-3130123210303311-2100310013322023-2003321300202302-0101001321232302-0232203321100000-2310013110330001) |
| `virtual_server.request_logging_profile.uid` | [virtual_server.request_logging_profile.uid](resources--application_profiles--reference--group-003.md#canonical-1031020323010220-1310111021303131-1211130301120110-0120303302330312-0301202230332110-0122211111020020-2323102113133330-0220213033212231) |
| `virtual_server.source_port` | [virtual_server.source_port](resources--application_profiles--reference--group-003.md#canonical-2212010332223131-2111323213003221-3122213303023310-2020333131322323-0103201100101002-1020010232021313-0111203202021022-2223132001320213) |
| `virtual_server.source_port.source_port_change` | [virtual_server.source_port.source_port_change](resources--application_profiles--reference--group-003.md#canonical-1311330323021032-2210013333333302-1331301003130203-0003333220201000-3131331012012000-2013032302131030-2303202113122003-3323113201112320) |
| `virtual_server.source_port.source_port_preserve` | [virtual_server.source_port.source_port_preserve](resources--application_profiles--reference--group-003.md#canonical-1022313012110112-0320200311020023-1202203230011030-1020100221232300-1322001023000120-3123313223012212-0303223030221123-1102121132013322) |
| `virtual_server.source_port.source_port_preserve_strict` | [virtual_server.source_port.source_port_preserve_strict](resources--application_profiles--reference--group-003.md#canonical-2212111003100111-1020312302333103-1300101311131310-0311112203132102-2032131012231213-1123102122310031-2310031230010301-0130002031012300) |
| `virtual_server.statistics_profile` | [virtual_server.statistics_profile](resources--application_profiles--reference--group-003.md#canonical-2033031113031132-0023012103300133-1113311132332021-2001201030102232-3111110013031123-1322303301011332-3021010132203122-3121230300333000) |
| `virtual_server.statistics_profile.kind` | [virtual_server.statistics_profile.kind](resources--application_profiles--reference--group-003.md#canonical-1003022020103133-2131111133003331-0330033120200303-2201030200030322-2303321002002211-0202203103201322-2310331120023300-2022202210203031) |
| `virtual_server.statistics_profile.name` | [virtual_server.statistics_profile.name](resources--application_profiles--reference--group-003.md#canonical-1300331213131100-1210231223321210-0020103121333312-3101130222010232-3132202211311101-3010011311133122-0031101310031213-1302031013313323) |
| `virtual_server.statistics_profile.namespace` | [virtual_server.statistics_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-3223323302013221-2212221320311113-0100312022002101-0330320332101201-0100233201322021-1030312112310232-0011022110313021-3300023211200223) |
| `virtual_server.statistics_profile.tenant` | [virtual_server.statistics_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-2221133111123231-3110101301023112-3231003013000012-2122230101110303-3212013230220030-2232232033012220-1223233112020031-0002030333311121) |
| `virtual_server.statistics_profile.uid` | [virtual_server.statistics_profile.uid](resources--application_profiles--reference--group-003.md#canonical-0132320122003211-0231231122332331-1003010313310213-2123301122301231-0322332221310022-0203123211123220-1222031302133221-1200331122023110) |
| `virtual_server.tcp` | [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-1310210103201212-1320301132201233-1210200013231223-1101230001132123-3302033232011132-2332321312221012-2323132221331212-0130323232233301) |
| `virtual_server.tcp.client_ssl_profile` | [virtual_server.tcp.client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-1012331213102311-2311220320111322-0123323233221201-0303203130002202-2000012130110011-0131133020313202-2111122321103312-1332310121103313) |
| `virtual_server.tcp.client_ssl_profile.kind` | [virtual_server.tcp.client_ssl_profile.kind](resources--application_profiles--reference--group-003.md#canonical-1233012303331002-1301300223130020-1210102110333010-1201103032033230-1030230201231331-2330230321233222-2123313200310302-3010200130331201) |
| `virtual_server.tcp.client_ssl_profile.name` | [virtual_server.tcp.client_ssl_profile.name](resources--application_profiles--reference--group-003.md#canonical-3312021332221321-2210030022101301-1223102233113103-1310220121202012-0202103103302322-0320203233310212-2322212223322033-1233030323030321) |
| `virtual_server.tcp.client_ssl_profile.namespace` | [virtual_server.tcp.client_ssl_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-2231021120022302-3031120322012220-3123202230023202-1201233022301012-3112121200011123-3331223131023202-1230310323230232-3220010310211111) |
| `virtual_server.tcp.client_ssl_profile.tenant` | [virtual_server.tcp.client_ssl_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-3120212012111032-1013103233222010-1203000331212110-1133102031232311-3121223232313333-0233221303232133-0323231331110220-1201223021000213) |
| `virtual_server.tcp.client_ssl_profile.uid` | [virtual_server.tcp.client_ssl_profile.uid](resources--application_profiles--reference--group-003.md#canonical-1121233010211331-3031222020231121-2312000122013110-1220221302100230-3200321013220312-3331200330132012-3001310002311313-3322122323132033) |
| `virtual_server.tcp.ocsp_profile` | [virtual_server.tcp.ocsp_profile](resources--application_profiles--reference--group-003.md#canonical-2201330102321311-3113102120311103-0132000300300123-2232331231112021-0030311232033330-1030333220331133-2233212221203000-2220021013322130) |
| `virtual_server.tcp.ocsp_profile.kind` | [virtual_server.tcp.ocsp_profile.kind](resources--application_profiles--reference--group-003.md#canonical-1201102121120222-1022222303212032-0322203300233201-2321122032331211-1312031123202300-3123131021001101-0031312312010221-3032313300130002) |
| `virtual_server.tcp.ocsp_profile.name` | [virtual_server.tcp.ocsp_profile.name](resources--application_profiles--reference--group-003.md#canonical-3220000202123301-2133312020312122-2213001131121321-2300122321203311-1333130233201213-1203023101110133-0311130020302200-2133220203102110) |
| `virtual_server.tcp.ocsp_profile.namespace` | [virtual_server.tcp.ocsp_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-1030330131020231-3012302212102302-0131222233220221-2313221220330210-1113333230120223-3223121032303210-3022220012221013-3021233212311311) |
| `virtual_server.tcp.ocsp_profile.tenant` | [virtual_server.tcp.ocsp_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-0032230003100312-1121023211302231-3120031120223200-2212022112333313-2120130201310201-1132102223112330-0303301332002311-0001202023003323) |
| `virtual_server.tcp.ocsp_profile.uid` | [virtual_server.tcp.ocsp_profile.uid](resources--application_profiles--reference--group-003.md#canonical-2221023120123210-3013033120202100-1222203231232011-0210232310230203-3321200030133200-2100022322310200-3100301220121010-1120022013123311) |
| `virtual_server.tcp.server_ssl_profile` | [virtual_server.tcp.server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-0313332111330002-2300020230310231-2231001022303321-0233332113220203-3231013103110021-2131313133233102-1133110013032032-3130221031223302) |
| `virtual_server.tcp.server_ssl_profile.kind` | [virtual_server.tcp.server_ssl_profile.kind](resources--application_profiles--reference--group-003.md#canonical-3000301213030020-0232223320210223-1100203023111111-2200123330230123-0220113113122020-0033322302122200-1020203113131031-2033323320221123) |
| `virtual_server.tcp.server_ssl_profile.name` | [virtual_server.tcp.server_ssl_profile.name](resources--application_profiles--reference--group-003.md#canonical-0331011311300103-2303020313011023-0323311131001332-2232012123332212-1213101230013110-2203232012211103-0312112021022210-2131202323331112) |
| `virtual_server.tcp.server_ssl_profile.namespace` | [virtual_server.tcp.server_ssl_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-2211313211222321-1200130133311001-1321223203123312-2322013012003030-2101320031032233-3021131130011110-3320231203110033-2110132130210033) |
| `virtual_server.tcp.server_ssl_profile.tenant` | [virtual_server.tcp.server_ssl_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-0320303333101212-1300001123203212-1120030002323031-3032033320113002-2232031333132003-1033022112033322-1300013131312023-1210002101332203) |
| `virtual_server.tcp.server_ssl_profile.uid` | [virtual_server.tcp.server_ssl_profile.uid](resources--application_profiles--reference--group-003.md#canonical-2301220023033330-2023131320213003-2131230213320311-1300222113211131-1112331122121130-2211030220030333-2112302322201101-3201301223333221) |
| `virtual_server.tcp.tcp_client_profile` | [virtual_server.tcp.tcp_client_profile](resources--application_profiles--reference--group-003.md#canonical-2323011223033022-3013123232310303-2200023233330130-2110012032322320-2211000322003323-1132330113030022-1210110101322320-2011132102222300) |
| `virtual_server.tcp.tcp_client_profile.kind` | [virtual_server.tcp.tcp_client_profile.kind](resources--application_profiles--reference--group-003.md#canonical-1202333133013223-0321300232132300-3322332332223322-3310100210232000-2313000123313022-0133031012031030-0111033001113023-1333023133112133) |
| `virtual_server.tcp.tcp_client_profile.name` | [virtual_server.tcp.tcp_client_profile.name](resources--application_profiles--reference--group-003.md#canonical-2313321202202003-2211103320031111-1320213310113001-1231201312123112-0133033013113213-3001313330131133-2200001003013203-2233312331323120) |
| `virtual_server.tcp.tcp_client_profile.namespace` | [virtual_server.tcp.tcp_client_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-2111120120313200-3222101022030010-1301020013112011-2231120323103120-2010320021303330-3022113313332101-0230303203133131-2231323202010312) |
| `virtual_server.tcp.tcp_client_profile.tenant` | [virtual_server.tcp.tcp_client_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-1031111320132203-0200012003211020-3021103333133103-2210212320101013-1000010323023303-2000010111101101-0230111331202123-3231120012033013) |
| `virtual_server.tcp.tcp_client_profile.uid` | [virtual_server.tcp.tcp_client_profile.uid](resources--application_profiles--reference--group-003.md#canonical-2031013330200223-0133032220302133-3223210223232202-1110331220121021-0213221001321033-0112020211120203-3232331310200220-1120233212320331) |
| `virtual_server.tcp.tcp_server_profile` | [virtual_server.tcp.tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-3122102111112221-0210010201323231-2021200021100112-1020010323310032-1011031101220320-1111000012210123-0010330130312223-1231232220122133) |
| `virtual_server.tcp.tcp_server_profile.kind` | [virtual_server.tcp.tcp_server_profile.kind](resources--application_profiles--reference--group-003.md#canonical-1303000301321302-3323023113031001-2131111011202020-1113011322321101-3101312301111011-2100030020031300-2221222311010002-0030123213213302) |
| `virtual_server.tcp.tcp_server_profile.name` | [virtual_server.tcp.tcp_server_profile.name](resources--application_profiles--reference--group-003.md#canonical-1010002221322231-2200211120100222-3223323222032100-0030000033123203-1201012112031311-1201213011010333-0222033300320320-2032120100321330) |
| `virtual_server.tcp.tcp_server_profile.namespace` | [virtual_server.tcp.tcp_server_profile.namespace](resources--application_profiles--reference--group-003.md#canonical-2120132210000323-2102211130312203-1223202021100201-2131133002303321-3203123113221012-2213313132001133-1102132201232002-1021032020110230) |
| `virtual_server.tcp.tcp_server_profile.tenant` | [virtual_server.tcp.tcp_server_profile.tenant](resources--application_profiles--reference--group-003.md#canonical-2030130300331200-1133103021231303-0312010033233120-1332321301230300-2002110230112212-2010022302000312-0122331020010202-3031032010011133) |
| `virtual_server.tcp.tcp_server_profile.uid` | [virtual_server.tcp.tcp_server_profile.uid](resources--application_profiles--reference--group-003.md#canonical-3222222311333301-2121222012111010-1002102211212032-2321032311302023-2312132301021221-2013023103122003-2023330233313221-3202312331301211) |
| `virtual_server.udp` | [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-0202310031202010-2013230330033220-3203120030300320-3233311000332012-2333000000230213-3011230113132321-2311033300201330-2320323313023303) |
| `virtual_server.udp.client_ssl_profile` | [virtual_server.udp.client_ssl_profile](resources--application_profiles--reference--group-004.md#canonical-3033031223130232-3113223130200330-2021133322032131-2121022022233233-0110300131213332-1221213010021110-2033230033101033-1301003103210032) |
| `virtual_server.udp.client_ssl_profile.kind` | [virtual_server.udp.client_ssl_profile.kind](resources--application_profiles--reference--group-004.md#canonical-0112221330330033-0123100222330113-0000131223033231-3122001033223013-2313131213121103-3032233332110211-0112002133302311-0012332122121212) |
| `virtual_server.udp.client_ssl_profile.name` | [virtual_server.udp.client_ssl_profile.name](resources--application_profiles--reference--group-004.md#canonical-3310003111312011-0110110302200230-1310313220211111-1320110221232332-1232133312013202-1032321212101131-3233002001333230-0003032200000230) |
| `virtual_server.udp.client_ssl_profile.namespace` | [virtual_server.udp.client_ssl_profile.namespace](resources--application_profiles--reference--group-004.md#canonical-2203021033323210-3131000030120021-3332320122220013-1032233323200031-1123322132130213-3232103210212020-0312111222223001-3103011111101010) |
| `virtual_server.udp.client_ssl_profile.tenant` | [virtual_server.udp.client_ssl_profile.tenant](resources--application_profiles--reference--group-004.md#canonical-3303230300320023-2312211010220130-3120021233220312-2311103213322001-0320231313302113-3200110203101102-1230220101132132-3111200323330002) |
| `virtual_server.udp.client_ssl_profile.uid` | [virtual_server.udp.client_ssl_profile.uid](resources--application_profiles--reference--group-004.md#canonical-1021230312010001-3323222330013223-0312201330133212-0310300310330323-1130122012220120-3013113132311122-2320000302220020-1022320232202003) |
| `virtual_server.udp.server_ssl_profile` | [virtual_server.udp.server_ssl_profile](resources--application_profiles--reference--group-004.md#canonical-3003003201122202-0223010232323030-0120011011231320-2210202332111112-1123030111021102-0231120121330001-1203103020331123-0003003303322302) |
| `virtual_server.udp.server_ssl_profile.kind` | [virtual_server.udp.server_ssl_profile.kind](resources--application_profiles--reference--group-004.md#canonical-3330011101321333-2102122213332322-3331123203310012-1320120133210330-2131113210033133-0021313322130030-2021220000200022-2211021220032103) |
| `virtual_server.udp.server_ssl_profile.name` | [virtual_server.udp.server_ssl_profile.name](resources--application_profiles--reference--group-004.md#canonical-1313322120200101-0001100011203130-2222023031332012-1313211011110230-2100322022323232-2300311032020022-3131301121032133-3302213120130011) |
| `virtual_server.udp.server_ssl_profile.namespace` | [virtual_server.udp.server_ssl_profile.namespace](resources--application_profiles--reference--group-004.md#canonical-1221003120222303-2113123323330112-1120232022033123-2211003222320221-1200120022301111-2023021012211121-1322023011013201-0022313301010113) |
| `virtual_server.udp.server_ssl_profile.tenant` | [virtual_server.udp.server_ssl_profile.tenant](resources--application_profiles--reference--group-004.md#canonical-0222121030223213-1110131020220332-3323322211320002-3131033333022111-0132230222133020-2103303023322220-3022021002231122-3302011210201131) |
| `virtual_server.udp.server_ssl_profile.uid` | [virtual_server.udp.server_ssl_profile.uid](resources--application_profiles--reference--group-004.md#canonical-3013133102112210-1320121112233011-1133130310232133-1032320211021023-3300021022330023-3221103221010123-1102013113013303-3003200012300220) |
| `virtual_server.udp.udp_client_profile` | [virtual_server.udp.udp_client_profile](resources--application_profiles--reference--group-004.md#canonical-3102232222323031-2323123101221120-2230023032121103-1011120012111311-1103321001112222-0301133311233110-0213100312222021-0122013311111330) |
| `virtual_server.udp.udp_client_profile.kind` | [virtual_server.udp.udp_client_profile.kind](resources--application_profiles--reference--group-004.md#canonical-2030030302310122-3211100003112200-1223203112013030-1113312033320103-2123330331330303-0211100120030331-0231211210212131-1113230110001003) |
| `virtual_server.udp.udp_client_profile.name` | [virtual_server.udp.udp_client_profile.name](resources--application_profiles--reference--group-004.md#canonical-3102132213113330-1100000322332332-0213231020133021-0013313220102203-3333123202303023-0003131110020201-1103123231120101-3012111333313211) |
| `virtual_server.udp.udp_client_profile.namespace` | [virtual_server.udp.udp_client_profile.namespace](resources--application_profiles--reference--group-004.md#canonical-2321111021331113-1232122033013100-3333200333220210-0001322102210200-1331333111123130-2033012002030031-1320102202202203-0323020011312012) |
| `virtual_server.udp.udp_client_profile.tenant` | [virtual_server.udp.udp_client_profile.tenant](resources--application_profiles--reference--group-004.md#canonical-0032100333201103-0220001121020133-2221320033230312-2210200323112023-2001120130030213-0302222113003011-0023310301120230-3321030330121022) |
| `virtual_server.udp.udp_client_profile.uid` | [virtual_server.udp.udp_client_profile.uid](resources--application_profiles--reference--group-004.md#canonical-1300012133023202-0031103230103113-2020122322002101-3100100020132020-2233112012330303-0201332200033202-2130111122123301-0301023110220103) |
| `virtual_server.udp.udp_server_profile` | [virtual_server.udp.udp_server_profile](resources--application_profiles--reference--group-004.md#canonical-0230123203221322-0330012221231312-0321110103233220-1311131131303031-0212322030321232-1001002021323123-1333323130323301-0303213012323011) |
| `virtual_server.udp.udp_server_profile.kind` | [virtual_server.udp.udp_server_profile.kind](resources--application_profiles--reference--group-004.md#canonical-3333223030203310-3022013212003103-1220121302230210-2023313013012312-0112010202020300-0302002003213321-2311120013310122-3023312002011220) |
| `virtual_server.udp.udp_server_profile.name` | [virtual_server.udp.udp_server_profile.name](resources--application_profiles--reference--group-004.md#canonical-2120003111113001-3231302300033013-2112210312301322-3030223033023233-2100013120213022-3111230331012331-2022020002012113-0033331232310103) |
| `virtual_server.udp.udp_server_profile.namespace` | [virtual_server.udp.udp_server_profile.namespace](resources--application_profiles--reference--group-004.md#canonical-0001023012113110-0133300100233311-2203122230012223-0000220021112210-3002112020313031-2322210333210311-2301113100333313-1203301001031003) |
| `virtual_server.udp.udp_server_profile.tenant` | [virtual_server.udp.udp_server_profile.tenant](resources--application_profiles--reference--group-004.md#canonical-0231311131102033-1200223220033132-2301230000211130-1321232200332202-0030230310020323-2012031321313002-3113310002201220-2022013320310022) |
| `virtual_server.udp.udp_server_profile.uid` | [virtual_server.udp.udp_server_profile.uid](resources--application_profiles--reference--group-004.md#canonical-0100303333012021-2200032023233320-2121312121001322-0031313023212230-1202011231110003-3301033311302232-3320333220313223-2012322201011332) |
| `virtual_server.virtual_server_state` | [virtual_server.virtual_server_state](resources--application_profiles--reference--group-004.md#canonical-0122303122010231-1321031133131130-1202302030103331-3111033233131023-3101020302031303-3022112323200233-2003020011121311-2213000201203322) |
| `virtual_server.virtual_server_state.state_disabled` | [virtual_server.virtual_server_state.state_disabled](resources--application_profiles--reference--group-004.md#canonical-0120301130000121-0203210122211133-1021132101231303-0101332022233021-2133030330213310-3301330232222331-0202031022230300-1121322301311323) |
| `virtual_server.virtual_server_state.state_enabled` | [virtual_server.virtual_server_state.state_enabled](resources--application_profiles--reference--group-004.md#canonical-2030200221303101-1201013133000300-1132311030000301-3132012300330200-2133023103032200-2032232030013012-1332111122103212-3201113121302101) |
| `virtual_server.vs_score` | [virtual_server.vs_score](resources--application_profiles--reference--group-001.md#canonical-2132131111031230-1300312200033023-0231302111133020-2130131320133212-0213221211030321-2013120203302123-1330110231323203-1202031133321302) |

<a id="canonical-2233123010013332-0221321003320332-1320203331332121-2013113003311312-1232021012010302-1313313213330122-3111320020330022-1330332312230332"></a>

## Next pages — Property reference / 300231030013 / 12

- [advanced_tcp_profile](resources--application_profiles--reference--group-001.md#canonical-0102012130211220-0033013023011312-0020312201233333-0110230303020102-0021300320132223-2232021231133230-0130333111122200-0300200330323322)
- [ddos_profile](resources--application_profiles--reference--group-001.md#canonical-2202332031223123-2301302010321012-1011102311002222-1100030332313031-3311332110021011-2231133020100211-2122201323020110-1321212330120300)
- [irules](resources--application_profiles--reference--group-001.md#canonical-0333211213222003-2111200010301230-0111200020033200-1110123010330321-0030333121010031-2213122011223020-1010213020122303-1230123111121033)
- [timeouts](resources--application_profiles--reference--group-001.md#canonical-2100312103022331-2033132121021132-3011311223212110-1112030031330113-3322301201010000-3123302212130112-2110221332213112-2221112310232000)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0102012130211220-0033013023011312-0020312201233333-0110230303020102-0021300320132223-2232021231133230-0130333111122200-0300200330323322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131000222223202-3323230121330331-3130130233321300-2031010131302032-2010301200113230-0032222032123002-0200031320212313-1131013333020102"></a>

## advanced_tcp_profile — advanced_tcp_profile / 331100030312 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- advanced_tcp_profile

<a id="canonical-0131203132032110-0212101023311010-2032110211331123-0332202113033003-3002223023320122-1001030130001112-1112130311233202-1012311211333220"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for advanced tcp profile.

Upstream description:

BIG-IP Advanced TCP Profile.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_tcp_advanced_profile",
    "enable_tcp_advanced_profile")}
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
  "x-ves-oneof-field-tcp_advanced_profile_choice": "[\"disable_tcp_advanced_profile\",\"enable_tcp_advanced_profile\"]"
}
```

Terraform syntax:

```terraform
advanced_tcp_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0023300203000011-0212003022300332-2010212001223100-0023110210231221-2321133032102323-0111300233231331-0023211112210313-3120030000330123"></a>

## Direct properties — advanced_tcp_profile / 331100030312 / 3

- [disable_tcp_advanced_profile](resources--application_profiles--reference--group-001.md#canonical-0201221302013100-2013130322130320-1200303113112222-1312203330331102-2031323221233100-0000021120300111-3303331233222322-3330311103011330): complete subsection reference.

- [enable_tcp_advanced_profile](resources--application_profiles--reference--group-001.md#canonical-0100120012303011-1032233223333112-2201311322200322-0132213030222110-0322131323122212-0010221001100313-1133303313112133-1222010102330302): complete subsection reference.

<a id="canonical-2313003132310100-1300003030233002-0112203132322233-2231320011201232-1120100013103103-1100020210133000-1200330001020303-2333230302102302"></a>

## Next pages — advanced_tcp_profile / 331100030312 / 4

- [advanced_tcp_profile.disable_tcp_advanced_profile](resources--application_profiles--reference--group-001.md#canonical-0201221302013100-2013130322130320-1200303113112222-1312203330331102-2031323221233100-0000021120300111-3303331233222322-3330311103011330)
- [advanced_tcp_profile.enable_tcp_advanced_profile](resources--application_profiles--reference--group-001.md#canonical-0100120012303011-1032233223333112-2201311322200322-0132213030222110-0322131323122212-0010221001100313-1133303313112133-1222010102330302)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0201221302013100-2013130322130320-1200303113112222-1312203330331102-2031323221233100-0000021120300111-3303331233222322-3330311103011330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203331302233331-3213021202020201-3011231303303313-3013000320013212-1000022030133102-2300312101102210-0203133223013132-0022012221123322"></a>

## advanced_tcp_profile.disable_tcp_advanced_profile — disable_tcp_advanced_profile / 320132312111 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [advanced_tcp_profile](resources--application_profiles--reference--group-001.md#canonical-0102012130211220-0033013023011312-0020312201233333-0110230303020102-0021300320132223-2232021231133230-0130333111122200-0300200330323322)
- advanced_tcp_profile.disable_tcp_advanced_profile

<a id="canonical-2213222212222231-2332222120213322-2130201032201311-1130332103012133-1103103330320222-3200122030123032-2032230233200130-2130230212302232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable tcp advanced profile.

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
disable_tcp_advanced_profile = {}
```

<a id="canonical-3200101201221230-2222322110203213-0321131132311300-3233230230213301-0132311321203321-2203210002122001-3123330220200223-0210022203132200"></a>

## Direct properties — disable_tcp_advanced_profile / 320132312111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002312020131210-3201012111010120-2111010332230100-0333220102001023-1331203133303102-3002031313133302-3211213100202320-2201020110131031"></a>

## Next pages — disable_tcp_advanced_profile / 320132312111 / 4

- [advanced_tcp_profile](resources--application_profiles--reference--group-001.md#canonical-0102012130211220-0033013023011312-0020312201233333-0110230303020102-0021300320132223-2232021231133230-0130333111122200-0300200330323322)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0100120012303011-1032233223333112-2201311322200322-0132213030222110-0322131323122212-0010221001100313-1133303313112133-1222010102330302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213032211032301-2221201030022301-2333222331132210-2231310021030122-1000002301221202-0213210233302300-2331323310322301-1232113011023221"></a>

## advanced_tcp_profile.enable_tcp_advanced_profile — enable_tcp_advanced_profile / 101021212133 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [advanced_tcp_profile](resources--application_profiles--reference--group-001.md#canonical-0102012130211220-0033013023011312-0020312201233333-0110230303020102-0021300320132223-2232021231133230-0130333111122200-0300200330323322)
- advanced_tcp_profile.enable_tcp_advanced_profile

<a id="canonical-3111111131310203-2102020011312223-1201321213320001-0122120131322231-2310131133200331-0012000132222202-2233101310322133-0312022030301332"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable tcp advanced profile.

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
enable_tcp_advanced_profile = {}
```

<a id="canonical-1323012222011033-2311301300312022-2130100220210130-3312121031130231-3231130032031321-2021212131222220-1113033000222001-0111012231101331"></a>

## Direct properties — enable_tcp_advanced_profile / 101021212133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312112213011132-3123222231123200-1321100202320030-0003130222231200-0102322312130322-0112122113020111-0011203230110102-3132013133201303"></a>

## Next pages — enable_tcp_advanced_profile / 101021212133 / 4

- [advanced_tcp_profile](resources--application_profiles--reference--group-001.md#canonical-0102012130211220-0033013023011312-0020312201233333-0110230303020102-0021300320132223-2232021231133230-0130333111122200-0300200330323322)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2202332031223123-2301302010321012-1011102311002222-1100030332313031-3311332110021011-2231133020100211-2122201323020110-1321212330120300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233001111011310-1000112003110022-3220200101131133-2331312213230332-2122003333030021-3120131003100331-2101313013003120-1000200112333110"></a>

## ddos_profile — ddos_profile / 330101212131 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- ddos_profile

<a id="canonical-1223320121231003-3321303031010020-2312302332001103-1011301210323311-1331330011110330-1032011211322112-2032321111130203-0202211101010030"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ddos profile.

Upstream description:

BIG-IP DDoS Protection Rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_ddos_mitigation",
    "enable_ddos_mitigation")}
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
  "x-ves-oneof-field-ddos_mitigation_choice": "[\"disable_ddos_mitigation\",\"enable_ddos_mitigation\"]"
}
```

Terraform syntax:

```terraform
ddos_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111230023200002-3130021133310122-0232003101131011-1312230130132110-2033122303313031-3122001132131101-0322212030033131-1223111022033321"></a>

## Direct properties — ddos_profile / 330101212131 / 3

- [disable_ddos_mitigation](resources--application_profiles--reference--group-001.md#canonical-2221001322213022-1103211321113101-1030110302100001-2123120100221211-3203131103303331-1003200210030103-2300322223030120-0212020131300221): complete subsection reference.

- [enable_ddos_mitigation](resources--application_profiles--reference--group-001.md#canonical-3012211310012110-3310211322133123-1010031131031103-3023210111312002-3021122233030130-2013113200222011-2123123110111001-1010111110221022): complete subsection reference.

<a id="canonical-1223322320022210-1233203200013212-3120212111310011-2321100123210001-0110111131001332-0120122100000322-1300113212002300-0221210120233230"></a>

## Next pages — ddos_profile / 330101212131 / 4

- [ddos_profile.disable_ddos_mitigation](resources--application_profiles--reference--group-001.md#canonical-2221001322213022-1103211321113101-1030110302100001-2123120100221211-3203131103303331-1003200210030103-2300322223030120-0212020131300221)
- [ddos_profile.enable_ddos_mitigation](resources--application_profiles--reference--group-001.md#canonical-3012211310012110-3310211322133123-1010031131031103-3023210111312002-3021122233030130-2013113200222011-2123123110111001-1010111110221022)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2221001322213022-1103211321113101-1030110302100001-2123120100221211-3203131103303331-1003200210030103-2300322223030120-0212020131300221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000132222220022-2001130122100311-0112233103202321-2211310011233032-3210320303202130-3001031033103203-2023123101013230-0123310122321302"></a>

## ddos_profile.disable_ddos_mitigation — disable_ddos_mitigation / 212033311223 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [ddos_profile](resources--application_profiles--reference--group-001.md#canonical-2202332031223123-2301302010321012-1011102311002222-1100030332313031-3311332110021011-2231133020100211-2122201323020110-1321212330120300)
- ddos_profile.disable_ddos_mitigation

<a id="canonical-1321321102100022-1102333210032332-1120201113020323-3123202112031320-0010111123030133-3003001332111312-0103220223131301-3331212132311330"></a>

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
disable_ddos_mitigation = {}
```

<a id="canonical-1012000012021032-1111202231221002-0220212133320113-1023010220312233-2300232302213321-0212100112332233-3012223202320223-1232002103313030"></a>

## Direct properties — disable_ddos_mitigation / 212033311223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301230022323313-1210300310023332-0120122231303330-2022033312330112-3331221333120302-1131000323011311-2233111103021101-1223321332222232"></a>

## Next pages — disable_ddos_mitigation / 212033311223 / 4

- [ddos_profile](resources--application_profiles--reference--group-001.md#canonical-2202332031223123-2301302010321012-1011102311002222-1100030332313031-3311332110021011-2231133020100211-2122201323020110-1321212330120300)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3012211310012110-3310211322133123-1010031131031103-3023210111312002-3021122233030130-2013113200222011-2123123110111001-1010111110221022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333111303013312-2210203301103100-3321031121011300-2331011012211203-0122300323012212-2222113230333003-3213320333000000-2211203222232100"></a>

## ddos_profile.enable_ddos_mitigation — enable_ddos_mitigation / 123321303102 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [ddos_profile](resources--application_profiles--reference--group-001.md#canonical-2202332031223123-2301302010321012-1011102311002222-1100030332313031-3311332110021011-2231133020100211-2122201323020110-1321212330120300)
- ddos_profile.enable_ddos_mitigation

<a id="canonical-1220102123222302-2012021211123101-3110323021201220-1113031033222330-0021303312223003-0331312133330202-2010121203222003-2232210111101111"></a>

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
enable_ddos_mitigation = {}
```

<a id="canonical-2323123113020220-1333310121222033-2132110121012010-1120332230011013-2333030220122103-3122111020033300-2220033211111121-1113200300230310"></a>

## Direct properties — enable_ddos_mitigation / 123321303102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022311113003010-0123201332001213-1333130201331332-0110003001031202-3133301223000123-0033332333332333-2012331022301131-3011233310301222"></a>

## Next pages — enable_ddos_mitigation / 123321303102 / 4

- [ddos_profile](resources--application_profiles--reference--group-001.md#canonical-2202332031223123-2301302010321012-1011102311002222-1100030332313031-3311332110021011-2231133020100211-2122201323020110-1321212330120300)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0333211213222003-2111200010301230-0111200020033200-1110123010330321-0030333121010031-2213122011223020-1010213020122303-1230123111121033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022322322220323-3112222203220110-2010102013231030-3301233130133210-3123120022100220-3031301311020230-0021021003010021-1112133331031311"></a>

## irules — irules / 020111013332 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- irules

<a id="canonical-3201312303230210-3330121210002003-3201220220001113-3211311212101311-1223121230302030-2333311023002030-2203031222103002-0303020210202302"></a>

Type: `"object"`. list nested block, Optional.

OPTIONS for attaching iRules to BIG-IP Proxy.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
irules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333330331011032-2230323020213312-1333300230231100-2112123313122321-2213000230321111-3003211321222313-3111003022231313-0313321013002310"></a>

## Direct properties — irules / 020111013332 / 3

<a id="canonical-0213302010001001-3313301203133010-2303233322203311-2301300223202302-3120130111112033-1333221000023102-0123012202322320-1011200301031023"></a>

<a id="canonical-2002320123221000-0100203331003121-0101120202312321-1223323023210200-2133311003112031-3302033202300113-2333012012321032-3123110311031202"></a>

## kind property — irules / 020111013332 / 4

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

<a id="canonical-0232230210112222-1112102331103011-1232232303300223-0321102033301332-1211122302310022-2002113123331003-3012300133131231-0230101120233102"></a>

<a id="canonical-2023010131101330-1233323221110012-1301010023223331-1312101312313230-2133201320000331-1203112231030211-0012303201003233-1120021302111321"></a>

## name property — irules / 020111013332 / 5

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

<a id="canonical-3132302021003023-3331333131300100-0211122111220101-2113322111121312-1010322133230311-1321031311323100-0111311113113133-2113131101221123"></a>

<a id="canonical-0322123003202221-3212301333301301-3323302123122113-2100121212013000-2203210331333220-1102330300110020-2312201332011313-2020223310033000"></a>

## namespace property — irules / 020111013332 / 6

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

<a id="canonical-3201302212023101-1012123130132013-1231131003001012-3001130313230022-0121203321112021-2031022120303011-1000120011010102-3022232023113301"></a>

<a id="canonical-0323011011002220-3233211220031323-2032012313112310-3102131112033311-2203131333210233-1230220113013003-1220101131301123-1333303201320330"></a>

## tenant property — irules / 020111013332 / 7

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

<a id="canonical-0122003020322120-1213002321111221-0201203003311301-3310231002022130-2132133231200301-0013002103213220-0110311122330132-2202212303202301"></a>

<a id="canonical-2200130333011311-3131230220111121-3030000332020301-0300122121132211-1022110330201010-0002222121103310-1232232000203233-3112311310123211"></a>

## uid property — irules / 020111013332 / 8

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

<a id="canonical-0113133221003003-1301201231312131-1101030012000222-0010030332022131-1313121223131313-2121232122212031-0233310102130131-2033113002011312"></a>

## Next pages — irules / 020111013332 / 9

- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2100312103022331-2033132121021132-3011311223212110-1112030031330113-3322301201010000-3123302212130112-2110221332213112-2221112310232000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101110012332302-3020120303312011-0113131111130212-0002210003230102-0021223013303322-1212312320020233-2220020133202032-2302233230023302"></a>

## timeouts — timeouts / 312302003210 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- timeouts

<a id="canonical-3023002233200132-1221101032202122-0203100022131133-0333230103323011-2133121300332321-2201131022300002-3101100102330133-2321212230102310"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103112121020010-1323123233221200-3333212200031032-2231323333220210-1100133213213002-0021220003303001-2000113313330102-2010332212101313"></a>

## Direct properties — timeouts / 312302003210 / 3

<a id="canonical-3100210310022100-3300213130331313-2202212203203103-1303310122020313-3220202020323103-3110303031303213-0231001302031132-3120100330301021"></a>

<a id="canonical-2133313300133200-3002003031200022-0310120133000120-3020222003231111-0222310220000223-1003021232123211-2220131202003303-3222112110133202"></a>

## create property — timeouts / 312302003210 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0233230313012211-2111222111213103-2110311121322213-0210332310020332-2122213210003130-1303223110111320-3313032330112322-0321322130231111"></a>

<a id="canonical-3203223122111233-3031113003212122-3033003302020021-2031331111001233-0032101032202311-1130230332001232-1333100231222222-2121010200311211"></a>

## delete property — timeouts / 312302003210 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1321021001133122-2313303100033103-2322110102232232-2010121001320310-0221130230133211-2122332312122102-2203002322231120-0121312002002110"></a>

<a id="canonical-3102003203131121-0001313123113230-2200000123130211-2012211133211332-3122121322210022-3133321211132211-2220001320123112-3032112201111020"></a>

## read property — timeouts / 312302003210 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1130210232203300-2200030213200132-1310331333020003-0022011021010303-3321203300301301-3301010010123223-0302302013223310-0101233323032223"></a>

<a id="canonical-0021132313113230-3313000303301011-3102123302102122-1100021223222310-2320132010213010-3130000213102310-2120002002121023-3112301130332130"></a>

## update property — timeouts / 312302003210 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2101312320210333-0122012123032303-1003113032131322-3100001313120010-0003311322030013-3311220222221101-0011000230223332-3013123220103221"></a>

## Next pages — timeouts / 312302003210 / 8

- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303130112131130-0220332210101011-3231330321220010-0133023020033220-1300020330231020-0321112312001111-2201302111120101-2003033323110123"></a>

## virtual_server — virtual_server / 322332232032 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- virtual_server

<a id="canonical-0223123303001200-0331130002020103-2301300000320300-2323002001312002-0013113220333013-1301230220223200-1311211110203213-3222203102122101"></a>

Type: `"object"`. single nested block, Optional.

Specifies configuration related to virtual server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http",
    "http3"),
  validators.ConflictingObjectAttributes("http",
    "https"),
  validators.ConflictingObjectAttributes("http",
    "tcp"),
  validators.ConflictingObjectAttributes("http",
    "udp"),
  validators.ConflictingObjectAttributes("http3",
    "https"),
  validators.ConflictingObjectAttributes("http3",
    "tcp"),
  validators.ConflictingObjectAttributes("http3",
    "udp"),
  validators.ConflictingObjectAttributes("https",
    "tcp"),
  validators.ConflictingObjectAttributes("https",
    "udp"),
  validators.ConflictingObjectAttributes("tcp",
    "udp")}
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
  "x-ves-oneof-field-virtual_server_type": "[\"http\",\"http3\",\"https\",\"tcp\",\"udp\"]"
}
```

Terraform syntax:

```terraform
virtual_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-3001220232230113-3232103101231203-1322222210221001-2013002103130312-0221033313013103-3331132301302210-0231021113013022-3202113210311233"></a>

## Direct properties — virtual_server / 322332232032 / 3

- [access_profile](resources--application_profiles--reference--group-001.md#canonical-1103220102313320-1103021311012130-1222022231332111-2130321202133312-3132210032030313-2232033220322030-3132200001032210-3301032131013002): complete subsection reference.

- [address_translation](resources--application_profiles--reference--group-001.md#canonical-3021112011013001-3231012202002003-2211233300032133-1313022100111002-1312313102121203-1032312311312031-2201100330021010-0102103013113132): complete subsection reference.

- [auto_last_hop](resources--application_profiles--reference--group-002.md#canonical-0121310003002130-3003332201223302-1200332002132003-2333323132231110-0132023021202330-1302303220222033-1033311330201233-3023023232102123): complete subsection reference.

- [clone_pool_client](resources--application_profiles--reference--group-002.md#canonical-1021302330213202-2303301231320102-2211323122010333-2021300332322031-3212031203232320-0311110303022111-1313022210003033-3300113110112122): complete subsection reference.

- [clone_pool_server](resources--application_profiles--reference--group-002.md#canonical-0110030313223221-0212311332031102-3220000320023132-0113100203203210-2300323131111132-1320212032201203-2100233233233312-1111220311120132): complete subsection reference.

<a id="canonical-3302010203222333-0123332120302230-1133001202003313-2012021112101211-2220012001012020-2111220133210220-2123201312330130-1200112232202230"></a>

<a id="canonical-3112111323120332-1303130112120113-2032022102201300-0231010331120020-0300103231333000-3132010213101222-2223203320021220-3333023130003102"></a>

## connection_limit property — virtual_server / 322332232032 / 4

Type: `"number"`. Optional.

Specifies the maximum number of concurrent connections allowed for the virtual server. Setting this
to 0 turns off connection limits. The.

Upstream description:

Specifies the maximum number of concurrent connections allowed for the virtual server. Setting this
to 0 turns off connection limits. The default is 0.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-3212101123123123-2202303100203133-3103230100003322-0303123102100200-0230003130122322-2111302331112002-3110320121113212-2203202230321303"></a>

<a id="canonical-2010030023113013-3022113000110323-0000001013132002-3203321323121013-3131023233003232-0203222321212020-3011121223312132-3222212130230222"></a>

## connection_rate_limit property — virtual_server / 322332232032 / 5

Type: `"number"`. Optional.

Specifies the maximum number of connections-per-second allowed for a virtual server. When the number
of connections-per-second reaches the limit for a given virtual server, the system drops (UDP) or
resets (TCP) additional connection requests. This helps detect Denial of Service attacks, where..

Upstream description:

Specifies the maximum number of connections-per-second allowed for a virtual server. When the number
of connections-per-second reaches the limit for a given virtual server, the system drops (UDP) or
resets (TCP) additional connection requests. This helps detect Denial of Service attacks, where
connection requests flood a virtual server. Setting this to 0 turns off connection limits. The
default is 0.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202): complete subsection reference.

- [default_persistence_profile](resources--application_profiles--reference--group-002.md#canonical-1130200120303132-1310331020320222-2223332303331002-1001311303211331-1130223022103331-3332031011100321-0310011023011303-0330123013203021): complete subsection reference.

- [default_pool](resources--application_profiles--reference--group-002.md#canonical-0220212320310233-0020030230010211-1101200002230023-3233033023031122-2320003100323101-1300230000023000-1322202320211120-0101202210330311): complete subsection reference.

- [fallback_persistence_profile](resources--application_profiles--reference--group-002.md#canonical-3103321022122021-1320313333200212-3310120231330011-3010133213120203-2222103302302221-3011000212110223-0013202223013102-1230032013233120): complete subsection reference.

- [fix_profile](resources--application_profiles--reference--group-002.md#canonical-0300032301211222-3312312313011013-2022223230130323-3110133300220133-1121320101021121-2303320221132211-1112303221011231-2020012203010100): complete subsection reference.

- [http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322): complete subsection reference.

- [http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132): complete subsection reference.

- [https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031): complete subsection reference.

- [immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-2120313011001333-0332132133112211-2032113202213200-1133333320210120-0120003121121103-1110233000210030-2313321130012321-0301012230001201): complete subsection reference.

- [last_hop_pool](resources--application_profiles--reference--group-003.md#canonical-1133301311230011-1310300321322210-2110201322213000-0102311121311103-1133120133100303-3022030212232231-3000032201020203-0023300101231100): complete subsection reference.

- [nat64](resources--application_profiles--reference--group-003.md#canonical-2002031331123221-1331221331321023-3111232122011101-0320122210120313-0212322110001221-0230222031021312-0203110233131210-3013333012331002): complete subsection reference.

- [port_translation](resources--application_profiles--reference--group-003.md#canonical-3121121231001031-3330300000133002-1031230222111031-0110330122312033-2031202031313323-3102103200300201-2230013301021233-3121002202231030): complete subsection reference.

- [request_logging_profile](resources--application_profiles--reference--group-003.md#canonical-3210202300313301-0310203232211033-1321103003021001-0212002312311332-1112331303033331-3323000030313321-0300000231230113-3220231130103020): complete subsection reference.

- [source_port](resources--application_profiles--reference--group-003.md#canonical-1232230330103113-1011323112020110-2112322333333013-3302112210103221-2002022032123220-0023200230233130-2102030102320103-0013123112003122): complete subsection reference.

- [statistics_profile](resources--application_profiles--reference--group-003.md#canonical-2230332010120210-3032201031320120-1320321332001001-1230210320213223-3313130100303010-3030221032130120-0211333311033312-3013103322102113): complete subsection reference.

- [tcp](resources--application_profiles--reference--group-003.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231): complete subsection reference.

- [udp](resources--application_profiles--reference--group-003.md#canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313): complete subsection reference.

- [virtual_server_state](resources--application_profiles--reference--group-004.md#canonical-0323021232302203-3330212203221312-0021222111110121-2230020130302203-1311230313200231-3310002022201103-2331300211333032-3021322223300011): complete subsection reference.

<a id="canonical-2132131111031230-1300312200033023-0231302111133020-2130131320133212-0213221211030321-2013120203302123-1330110231323203-1202031133321302"></a>

<a id="canonical-2310200210120102-3203303313112312-3310021300030311-2233211031121310-0221302212023030-2110130121123111-0232010313202101-1210302133120111"></a>

## vs_score property — virtual_server / 322332232032 / 6

Type: `"number"`. Optional.

Specifies the virtual server score in percent. Global Traffic Manager (GTM) can rely on this value
to load balance traffic in a proportional manner. The , meaning that no additional metric is applied
for the virtual server.

Upstream description:

Specifies the virtual server score in percent. Global Traffic Manager (GTM) can rely on this value
to load balance traffic in a proportional manner. The default is 0, meaning that no additional
metric is applied for the virtual server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-3023231010333333-3111013121313011-1111323322221321-3303120110310223-3321030312212132-2033330111202113-0001010021202333-0031223202322322"></a>

## Next pages — virtual_server / 322332232032 / 7

- [virtual_server.access_profile](resources--application_profiles--reference--group-001.md#canonical-1103220102313320-1103021311012130-1222022231332111-2130321202133312-3132210032030313-2232033220322030-3132200001032210-3301032131013002)
- [virtual_server.address_translation](resources--application_profiles--reference--group-001.md#canonical-3021112011013001-3231012202002003-2211233300032133-1313022100111002-1312313102121203-1032312311312031-2201100330021010-0102103013113132)
- [virtual_server.auto_last_hop](resources--application_profiles--reference--group-002.md#canonical-0121310003002130-3003332201223302-1200332002132003-2333323132231110-0132023021202330-1302303220222033-1033311330201233-3023023232102123)
- [virtual_server.clone_pool_client](resources--application_profiles--reference--group-002.md#canonical-1021302330213202-2303301231320102-2211323122010333-2021300332322031-3212031203232320-0311110303022111-1313022210003033-3300113110112122)
- [virtual_server.clone_pool_server](resources--application_profiles--reference--group-002.md#canonical-0110030313223221-0212311332031102-3220000320023132-0113100203203210-2300323131111132-1320212032201203-2100233233233312-1111220311120132)
- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- [virtual_server.default_persistence_profile](resources--application_profiles--reference--group-002.md#canonical-1130200120303132-1310331020320222-2223332303331002-1001311303211331-1130223022103331-3332031011100321-0310011023011303-0330123013203021)
- [virtual_server.default_pool](resources--application_profiles--reference--group-002.md#canonical-0220212320310233-0020030230010211-1101200002230023-3233033023031122-2320003100323101-1300230000023000-1322202320211120-0101202210330311)
- [virtual_server.fallback_persistence_profile](resources--application_profiles--reference--group-002.md#canonical-3103321022122021-1320313333200212-3310120231330011-3010133213120203-2222103302302221-3011000212110223-0013202223013102-1230032013233120)
- [virtual_server.fix_profile](resources--application_profiles--reference--group-002.md#canonical-0300032301211222-3312312313011013-2022223230130323-3110133300220133-1121320101021121-2303320221132211-1112303221011231-2020012203010100)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-003.md#canonical-2120313011001333-0332132133112211-2032113202213200-1133333320210120-0120003121121103-1110233000210030-2313321130012321-0301012230001201)
- [virtual_server.last_hop_pool](resources--application_profiles--reference--group-003.md#canonical-1133301311230011-1310300321322210-2110201322213000-0102311121311103-1133120133100303-3022030212232231-3000032201020203-0023300101231100)
- [virtual_server.nat64](resources--application_profiles--reference--group-003.md#canonical-2002031331123221-1331221331321023-3111232122011101-0320122210120313-0212322110001221-0230222031021312-0203110233131210-3013333012331002)
- [virtual_server.port_translation](resources--application_profiles--reference--group-003.md#canonical-3121121231001031-3330300000133002-1031230222111031-0110330122312033-2031202031313323-3102103200300201-2230013301021233-3121002202231030)
- [virtual_server.request_logging_profile](resources--application_profiles--reference--group-003.md#canonical-3210202300313301-0310203232211033-1321103003021001-0212002312311332-1112331303033331-3323000030313321-0300000231230113-3220231130103020)
- [virtual_server.source_port](resources--application_profiles--reference--group-003.md#canonical-1232230330103113-1011323112020110-2112322333333013-3302112210103221-2002022032123220-0023200230233130-2102030102320103-0013123112003122)
- [virtual_server.statistics_profile](resources--application_profiles--reference--group-003.md#canonical-2230332010120210-3032201031320120-1320321332001001-1230210320213223-3313130100303010-3030221032130120-0211333311033312-3013103322102113)
- [virtual_server.tcp](resources--application_profiles--reference--group-003.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231)
- [virtual_server.udp](resources--application_profiles--reference--group-003.md#canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313)
- [virtual_server.virtual_server_state](resources--application_profiles--reference--group-004.md#canonical-0323021232302203-3330212203221312-0021222111110121-2230020130302203-1311230313200231-3310002022201103-2331300211333032-3021322223300011)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1103220102313320-1103021311012130-1222022231332111-2130321202133312-3132210032030313-2232033220322030-3132200001032210-3301032131013002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333112230302030-0131021033210203-0122331133220202-3030003333323112-3212333231301001-2331210232220312-3110132121331020-0300030130110223"></a>

## virtual_server.access_profile — access_profile / 001333120102 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.access_profile

<a id="canonical-2223203132121233-0130230302003020-0022132130321212-1323133010230220-3133112020312012-2213310032031021-2221020320021022-1322020220333033"></a>

Type: `"object"`. list nested block, Optional.

Specifies an access policy that determines the authentication rules and access controls applied to
user sessions for this virtual server.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
access_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0312203301302033-1231000031101110-2023020200200210-0110330202021023-2011231133003023-1222222111223013-0122033221331122-0121032231013030"></a>

## Direct properties — access_profile / 001333120102 / 3

<a id="canonical-2312120313230122-0103133201013300-3020301012103301-2120001020102000-1130303212230010-2101320023030220-2233002110122232-3010231223330130"></a>

<a id="canonical-1100322010220111-3332313320110212-2213232310202000-0322102000023312-1220101223023230-0111013002010301-1032011103123000-1130103213232010"></a>

## kind property — access_profile / 001333120102 / 4

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

<a id="canonical-3333000021223122-1231320332233023-3321322021230303-1130012313132100-3033230200332103-3231222200333121-1021233212002330-3230103331230230"></a>

<a id="canonical-3022131202123023-0120021112012000-1313133332323022-0102030203302320-0002212232233300-3222012033110130-0122032211231311-1001213323001001"></a>

## name property — access_profile / 001333120102 / 5

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

<a id="canonical-0102333102213322-3333132033032130-0322003333120312-2102231201112322-1033133130100203-0132022001303300-2310100322220313-1012302130202001"></a>

<a id="canonical-3310001212311202-2221133123011030-1223321201012301-1201011121133100-0113021223133300-0012022220200303-2200131323211323-0132013302223231"></a>

## namespace property — access_profile / 001333120102 / 6

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

<a id="canonical-1311231113222022-0201313212032203-0001230023321233-0033222132211312-3033321230330210-2030002120222303-1023303112000121-1230201121221112"></a>

<a id="canonical-1321012210012022-1130220313011123-0122121031233202-2313131021101133-1303230210313012-0212331010030332-3031002033121313-0013013330213011"></a>

## tenant property — access_profile / 001333120102 / 7

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

<a id="canonical-0021311103300301-2023023131211133-3123101003001111-3213030232303001-3001001321222230-2221011230103332-3013300113202012-2012301323130030"></a>

<a id="canonical-2103213112301330-1313200330113020-1230213101201303-1332111300302120-3230103100212321-2230300023032210-3211003021021113-0130110033301232"></a>

## uid property — access_profile / 001333120102 / 8

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

<a id="canonical-0023111031032131-0321121232100231-2213220001113301-1002011102102223-2321302231021312-1121010231320012-1120110302133033-3110111101311110"></a>

## Next pages — access_profile / 001333120102 / 9

- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3021112011013001-3231012202002003-2211233300032133-1313022100111002-1312313102121203-1032312311312031-2201100330021010-0102103013113132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021011303333302-3300121000311200-3231020013131002-0311002320022210-0212031100133322-2133131132101022-2220231110301201-0201013322322030"></a>

## virtual_server.address_translation — address_translation / 222202331311 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.address_translation

<a id="canonical-0300032312110001-3232021231013033-1202233220030012-2010213320123213-1112002012130110-0220131220101001-3220320331312001-0011330303333322"></a>

Type: `"object"`. single nested block, Optional.

Specifies, when checked (enabled), that the system translates the address of the virtual server.
When cleared (disabled), specifies that the system uses the address without translation. This option
is useful when the system is load balancing devices that have the same IP address.

Upstream description:

Specifies, when checked (enabled), that the system translates the address of the virtual server.
When cleared (disabled), specifies that the system uses the address without translation. This option
is useful when the system is load balancing devices that have the same IP address. The default is
enabled.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("address_translation_disable",
    "address_translation_enable")}
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
  "x-ves-oneof-field-address_translation_choice": "[\"address_translation_disable\",\"address_translation_enable\"]"
}
```

Terraform syntax:

```terraform
address_translation {
  # Configure direct properties listed below.
}
```

<a id="canonical-3033210020033221-1201302121123010-3212113010001122-3203323112312123-2011321010003221-1120323011231323-0203111130133220-3201113323223201"></a>

## Direct properties — address_translation / 222202331311 / 3

- [address_translation_disable](resources--application_profiles--reference--group-001.md#canonical-2000122210122111-0121133123022303-3113103102233210-2320301213121222-0213023033210230-3201210133032233-1310202323213311-2321331201032302): complete subsection reference.

- [address_translation_enable](resources--application_profiles--reference--group-001.md#canonical-3012201011311202-0012320220102023-1003023003223210-1132033230031311-0311321131123233-1210111000312012-0322110301130230-1203113031212032): complete subsection reference.

<a id="canonical-1301310111102001-2322020313333322-0032302333110013-2220311320203333-0101322101012210-0201230213311233-2011223233111203-1300132223210220"></a>

## Next pages — address_translation / 222202331311 / 4

- [virtual_server.address_translation.address_translation_disable](resources--application_profiles--reference--group-001.md#canonical-2000122210122111-0121133123022303-3113103102233210-2320301213121222-0213023033210230-3201210133032233-1310202323213311-2321331201032302)
- [virtual_server.address_translation.address_translation_enable](resources--application_profiles--reference--group-001.md#canonical-3012201011311202-0012320220102023-1003023003223210-1132033230031311-0311321131123233-1210111000312012-0322110301130230-1203113031212032)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2000122210122111-0121133123022303-3113103102233210-2320301213121222-0213023033210230-3201210133032233-1310202323213311-2321331201032302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002103012321200-2311103333201101-0002302000102013-2302321202303021-3123331231212032-3200130310023020-3023202210203301-3233322131210232"></a>

## virtual_server.address_translation.address_translation_disable — address_translation_disable / 331111330003 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.address_translation](resources--application_profiles--reference--group-001.md#canonical-3021112011013001-3231012202002003-2211233300032133-1313022100111002-1312313102121203-1032312311312031-2201100330021010-0102103013113132)
- virtual_server.address_translation.address_translation_disable

<a id="canonical-0012120321210313-1132301022231031-3132112031313330-1101133222033112-2212223011012323-3023223302313030-1030201213021112-2011302032200230"></a>

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
address_translation_disable = {}
```

<a id="canonical-3212121332301102-0311012100002232-1101233103200103-2201311021321103-3222132001132031-2332121322132012-0320300210112232-3131302111320210"></a>

## Direct properties — address_translation_disable / 331111330003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011310300310021-3312212000121300-1302103103002231-2023333322200133-2030103300311232-1122112110323003-0230023122121312-3200000210333122"></a>

## Next pages — address_translation_disable / 331111330003 / 4

- [virtual_server.address_translation](resources--application_profiles--reference--group-001.md#canonical-3021112011013001-3231012202002003-2211233300032133-1313022100111002-1312313102121203-1032312311312031-2201100330021010-0102103013113132)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3012201011311202-0012320220102023-1003023003223210-1132033230031311-0311321131123233-1210111000312012-0322110301130230-1203113031212032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322023020022301-2132113303330020-1313301112213113-2001332203322333-2130133301231312-0200300132302131-2311032211102103-3101130011330232"></a>

## virtual_server.address_translation.address_translation_enable — address_translation_enable / 321203012100 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.address_translation](resources--application_profiles--reference--group-001.md#canonical-3021112011013001-3231012202002003-2211233300032133-1313022100111002-1312313102121203-1032312311312031-2201100330021010-0102103013113132)
- virtual_server.address_translation.address_translation_enable

<a id="canonical-0331233123202101-2133022100220113-1033121202222303-0322123011213123-1011231330332022-3320310111310102-3033102012312201-1023303213331222"></a>

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
address_translation_enable = {}
```

<a id="canonical-2200221323112111-1121120032212010-1212100101222202-2331322211133230-2130323031123313-0203133121011123-1130303231121110-1020131302300332"></a>

## Direct properties — address_translation_enable / 321203012100 / 3

This is an empty object or choice marker. It has no direct properties.
