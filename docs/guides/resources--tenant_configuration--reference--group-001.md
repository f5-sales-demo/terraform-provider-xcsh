---
page_title: "xcsh_tenant_configuration reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tenant_configuration reference."
---

# xcsh_tenant_configuration reference

<a id="canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010321020302211-0023223322301312-3211122223121320-2232002320323000-2032312211233111-0100233110331131-0101021000223221-1100031122202332"></a>

## Property reference — Property reference / 323210111002 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)
- Property reference

<a id="canonical-0211230220200213-2110100233222331-0233112313222332-1032232231020323-3300313223313322-1203012001211030-3113221213322101-2223230200133010"></a>

## Direct properties — Property reference / 323210111002 / 3

<a id="canonical-3021103322123321-2321310311110323-2332302022002223-0303333222102311-2110233310031201-3303330220013120-0312021011203132-1233323301122203"></a>

<a id="canonical-3012003103120202-1032112200013301-0101123103002202-3111202213301130-0331323230221333-3300032231132332-0321221123020112-2221212303121201"></a>

## annotations property — Property reference / 323210111002 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

- [brute_force_detection](resources--tenant_configuration--reference--group-001.md#canonical-2002011201213020-3120102323023032-3331212102222001-0213001223111132-0330203000010333-3320213330231102-1100132103102210-3220032120100020): complete subsection reference.

<a id="canonical-1111233011010123-0222022100003313-1030320210101131-0020120132322120-2300023212332100-0322032001012310-0113132303112222-1231200132310203"></a>

<a id="canonical-3000011000123220-0210003333200021-0033211203033012-1033202321130023-1010301000223123-3211102100111000-1300313132133122-1100311022220303"></a>

## description property — Property reference / 323210111002 / 5

Type: `"string"`. Optional.

Human readable description for the object.

<a id="canonical-2021003021322001-0103202321230012-3001312322210133-3202110130222200-2311123221003222-3130310231231001-2231010101313131-3031032333120222"></a>

<a id="canonical-0200230132013332-0030022032111223-2021310013001223-0030300203003000-1110221332201230-3320130320002321-0331233312011301-2233310332312231"></a>

## disable property — Property reference / 323210111002 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

<a id="canonical-3121120232032013-1320001233012001-1103223233313031-1132102121113312-1003220020100012-3201210112301202-3130023002120133-2032313220023301"></a>

<a id="canonical-3331111001233101-3303230311133130-2223223300200122-1010212201300131-1210030113230112-0322100322112203-1122200131033302-2312213221200032"></a>

## ID property — Property reference / 323210111002 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2011302231010330-1333002312323133-2010022021310011-1202000203212021-3122100121233302-1233220133021321-3131110311223323-1112110132003100"></a>

<a id="canonical-2023132133002003-0221231332200220-0210320312003120-1312302310013112-0101313113223220-0120332112011013-0302213331221330-3000221101210013"></a>

## labels property — Property reference / 323210111002 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

<a id="canonical-1031003323020231-0102322233311120-3301110300322313-0000230003330110-2001232022032221-1221012223302100-3221013230023223-0202111100032033"></a>

<a id="canonical-2230010202323212-3221010110022131-2032032212211222-1121223103222020-0133220030310112-0000331013012032-3131030212212020-2122320323133301"></a>

## name property — Property reference / 323210111002 / 9

Type: `"string"`. Required.

Name of the Tenant Configuration. Must be unique within the namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

<a id="canonical-2233223301123110-2100013031031233-0020323310102232-2010002013022302-1112210322302202-3002031100230313-0003031220331313-3330223211231301"></a>

<a id="canonical-3331303322330310-3100103211101001-1232303321100102-1101212332300100-0013203201111123-1121220002103020-3100233122233021-3213301232031300"></a>

## namespace property — Property reference / 323210111002 / 10

Type: `"string"`. Required.

Namespace where the Tenant Configuration is created.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

- [password_policy](resources--tenant_configuration--reference--group-001.md#canonical-3102022230023020-1021013032122012-3020123221123311-3313322012121231-3031033300031300-3312331012000020-2031003212112110-2203300222010022): complete subsection reference.

- [tenant_details](resources--tenant_configuration--reference--group-001.md#canonical-1011022202202111-0102220102231000-2232303112122022-0132312233310312-2213203313133203-1210320213013321-3222001331131113-3331130300130333): complete subsection reference.

- [timeouts](resources--tenant_configuration--reference--group-001.md#canonical-0211120030131233-0221232021112002-3012103322332003-1021113132232201-2312122022030320-1231200221002212-0122322000130230-2300121212012010): complete subsection reference.

- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-3002032020221331-0023102232020311-2221003000102010-1231312030013330-1231033101320300-2312123103202330-1322112233100103-2133031101202322): complete subsection reference.

<a id="canonical-2201133202131211-1321101333112211-3220210223323211-0023033113033123-2201103003032012-2031322012033121-0310221112123132-3122011320201000"></a>

## All schema paths — Property reference / 323210111002 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--tenant_configuration--reference--group-001.md#canonical-3021103322123321-2321310311110323-2332302022002223-0303333222102311-2110233310031201-3303330220013120-0312021011203132-1233323301122203) |
| `brute_force_detection` | [brute_force_detection](resources--tenant_configuration--reference--group-001.md#canonical-1020332221103212-3021131122203101-1013301032133200-2032311123031210-0322032303003120-2033022230333022-2131020200033000-3312230222211130) |
| `brute_force_detection.max_login_failures` | [brute_force_detection.max_login_failures](resources--tenant_configuration--reference--group-001.md#canonical-0120213103212012-0310313212312131-0301300131231002-2000002021120221-2130020120311200-0221103223132233-0110213001112231-0110132230000200) |
| `description` | [description](resources--tenant_configuration--reference--group-001.md#canonical-1111233011010123-0222022100003313-1030320210101131-0020120132322120-2300023212332100-0322032001012310-0113132303112222-1231200132310203) |
| `disable` | [disable](resources--tenant_configuration--reference--group-001.md#canonical-2021003021322001-0103202321230012-3001312322210133-3202110130222200-2311123221003222-3130310231231001-2231010101313131-3031032333120222) |
| `id` | [id](resources--tenant_configuration--reference--group-001.md#canonical-3121120232032013-1320001233012001-1103223233313031-1132102121113312-1003220020100012-3201210112301202-3130023002120133-2032313220023301) |
| `labels` | [labels](resources--tenant_configuration--reference--group-001.md#canonical-2011302231010330-1333002312323133-2010022021310011-1202000203212021-3122100121233302-1233220133021321-3131110311223323-1112110132003100) |
| `name` | [name](resources--tenant_configuration--reference--group-001.md#canonical-1031003323020231-0102322233311120-3301110300322313-0000230003330110-2001232022032221-1221012223302100-3221013230023223-0202111100032033) |
| `namespace` | [namespace](resources--tenant_configuration--reference--group-001.md#canonical-2233223301123110-2100013031031233-0020323310102232-2010002013022302-1112210322302202-3002031100230313-0003031220331313-3330223211231301) |
| `password_policy` | [password_policy](resources--tenant_configuration--reference--group-001.md#canonical-3022112100000001-1030120011031300-0313030102002100-1313130130120201-2011030231103023-0201322121310103-3112322203113122-0032120011202033) |
| `password_policy.digits` | [password_policy.digits](resources--tenant_configuration--reference--group-001.md#canonical-1330232212122300-0331013323013211-3112101020302220-2203332032301131-1321103003123321-0003303201331022-0211033011120231-3023301303001220) |
| `password_policy.expire_password` | [password_policy.expire_password](resources--tenant_configuration--reference--group-001.md#canonical-2311232122303323-1220332223132321-2313210123112132-2311233212002233-1120332002011312-3110211313220023-2011311330211011-1223001120101030) |
| `password_policy.lowercase_characters` | [password_policy.lowercase_characters](resources--tenant_configuration--reference--group-001.md#canonical-3331321310123223-0122233230313000-1000030300221020-1131110302233000-2301101210210300-0131112232211133-1301013033011323-0012212003233320) |
| `password_policy.minimum_length` | [password_policy.minimum_length](resources--tenant_configuration--reference--group-001.md#canonical-1233002113021031-1122102133103310-3032321000321231-2220323010323323-1020030330220313-3110000022023201-3220211221103002-3020110010221132) |
| `password_policy.not_recently_used` | [password_policy.not_recently_used](resources--tenant_configuration--reference--group-001.md#canonical-2330000232103132-2001212130232210-3133112031123132-1220223222212200-1203221331201030-2002321031200222-0322213220110112-1122130220001010) |
| `password_policy.not_username` | [password_policy.not_username](resources--tenant_configuration--reference--group-001.md#canonical-1221101110133111-1303303033201102-0300210111033300-2212013321231321-0231311203133011-0103323013211022-2122230001032213-0320203003130122) |
| `password_policy.special_characters` | [password_policy.special_characters](resources--tenant_configuration--reference--group-001.md#canonical-0311330330000203-2132120100110022-0310223221010120-0220303111322012-3130233212021033-2113123122011012-0313203200202001-0230112101300203) |
| `password_policy.uppercase_characters` | [password_policy.uppercase_characters](resources--tenant_configuration--reference--group-001.md#canonical-3311003131210010-3033212320333220-2303120210131031-1100102112331130-1300033223320301-0303210000001313-0100302322113101-1012333201221130) |
| `tenant_details` | [tenant_details](resources--tenant_configuration--reference--group-001.md#canonical-2223012132011230-2321231303113020-2201102130002012-2112211222133231-2032323223320211-3001132010330202-0032112201022211-0132230233131231) |
| `tenant_details.display_name` | [tenant_details.display_name](resources--tenant_configuration--reference--group-001.md#canonical-2012121132033311-1101102113131323-2230022231210302-0121301030311023-3010021131101032-3131102231031123-3132101013302232-2101022101022131) |
| `timeouts` | [timeouts](resources--tenant_configuration--reference--group-001.md#canonical-1300230233221312-1213330312333323-1323102311230113-3112131012111032-2303133203210110-1232313333133213-0310030033011212-0020320020112313) |
| `timeouts.create` | [timeouts.create](resources--tenant_configuration--reference--group-001.md#canonical-1133122021001301-1212301202001332-3203333130230222-1031331301231003-3230121011311103-2301000301130222-0030300030211013-2112101103132311) |
| `timeouts.delete` | [timeouts.delete](resources--tenant_configuration--reference--group-001.md#canonical-1113121012113000-1201202121003222-0021101213020132-2031321303001223-0313010033333233-1132123112113123-2333130233203220-2121100001331021) |
| `timeouts.read` | [timeouts.read](resources--tenant_configuration--reference--group-001.md#canonical-0030111303112312-2310222303112320-0000220221311133-3000213111331222-3200023032113213-0313013033132221-0110230222303001-3131330001321310) |
| `timeouts.update` | [timeouts.update](resources--tenant_configuration--reference--group-001.md#canonical-1013210300323121-1022300131032301-2220222311103301-0122101000030230-1232232100002333-3032222311131132-3233030303123202-2313130300112131) |
| `user_session_expiration` | [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-2101223000110110-0300010313112300-2200021311002331-1121011122003303-3123103133301311-3132200231333221-2103111001001011-0302033223101302) |
| `user_session_expiration.absolute_timeout` | [user_session_expiration.absolute_timeout](resources--tenant_configuration--reference--group-001.md#canonical-3101003133012123-2303213123233200-0321113211113001-3031111323010000-2033211311330323-0202232022120101-1130301203202313-2131220132221323) |
| `user_session_expiration.absolute_timeout.hours` | [user_session_expiration.absolute_timeout.hours](resources--tenant_configuration--reference--group-001.md#canonical-2333001321320010-2131301012010203-1203202013122020-3320333210013020-0001211200033102-2200003321231220-0102232323023123-0011111331320111) |
| `user_session_expiration.absolute_timeout.hours.duration` | [user_session_expiration.absolute_timeout.hours.duration](resources--tenant_configuration--reference--group-001.md#canonical-3112030302313023-3033320220100203-2131123113001233-3100101320032231-3011100233113011-0330032232323002-1201332232213113-1222311221122113) |
| `user_session_expiration.absolute_timeout.minutes` | [user_session_expiration.absolute_timeout.minutes](resources--tenant_configuration--reference--group-001.md#canonical-3333210001012231-0000121110310111-3023221201302332-1002330232332021-0002021221212102-0230220330131111-1200300101000111-1233033322333003) |
| `user_session_expiration.absolute_timeout.minutes.duration` | [user_session_expiration.absolute_timeout.minutes.duration](resources--tenant_configuration--reference--group-001.md#canonical-3313123133320102-1321031013000203-0122213201122102-2011010133210000-3311303203321212-0301033223113120-3000213312332330-1033203200113300) |
| `user_session_expiration.idle_timeout` | [user_session_expiration.idle_timeout](resources--tenant_configuration--reference--group-001.md#canonical-1130010230321001-0021001133000010-0200313133311301-3101120120030311-1032201133103200-3003021322200030-2131120003033020-0132202223332212) |
| `user_session_expiration.idle_timeout.hours` | [user_session_expiration.idle_timeout.hours](resources--tenant_configuration--reference--group-001.md#canonical-1221030310003031-1122022113230232-1310010122210220-1123312100230133-1120030133232221-0003220132101333-0122131330113322-2003223001330322) |
| `user_session_expiration.idle_timeout.hours.duration` | [user_session_expiration.idle_timeout.hours.duration](resources--tenant_configuration--reference--group-001.md#canonical-3023122233113030-1220333122330320-0130202311013213-0300103013221333-0213131203210101-3030100331321131-0031213201320333-0201312103333222) |
| `user_session_expiration.idle_timeout.minutes` | [user_session_expiration.idle_timeout.minutes](resources--tenant_configuration--reference--group-001.md#canonical-1013032210221213-0112203001130101-3033212020023113-2021033011001100-2330001100322230-0213221220032131-1022123332223101-1002023220211222) |
| `user_session_expiration.idle_timeout.minutes.duration` | [user_session_expiration.idle_timeout.minutes.duration](resources--tenant_configuration--reference--group-001.md#canonical-2021023230133101-1003312231213220-1012200212131102-3332210012100303-1122103310323212-2201012113103101-1101300131001010-1311233113123232) |

<a id="canonical-0021211302231233-1210210202002320-0112121000202021-1111131323330211-0131223210310233-0213123302123010-2211232133011013-2333311213020303"></a>

## Next pages — Property reference / 323210111002 / 12

- [brute_force_detection](resources--tenant_configuration--reference--group-001.md#canonical-2002011201213020-3120102323023032-3331212102222001-0213001223111132-0330203000010333-3320213330231102-1100132103102210-3220032120100020)
- [password_policy](resources--tenant_configuration--reference--group-001.md#canonical-3102022230023020-1021013032122012-3020123221123311-3313322012121231-3031033300031300-3312331012000020-2031003212112110-2203300222010022)
- [tenant_details](resources--tenant_configuration--reference--group-001.md#canonical-1011022202202111-0102220102231000-2232303112122022-0132312233310312-2213203313133203-1210320213013321-3222001331131113-3331130300130333)
- [timeouts](resources--tenant_configuration--reference--group-001.md#canonical-0211120030131233-0221232021112002-3012103322332003-1021113132232201-2312122022030320-1231200221002212-0122322000130230-2300121212012010)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-3002032020221331-0023102232020311-2221003000102010-1231312030013330-1231033101320300-2312123103202330-1322112233100103-2133031101202322)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)

<a id="canonical-2002011201213020-3120102323023032-3331212102222001-0213001223111132-0330203000010333-3320213330231102-1100132103102210-3220032120100020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333131222103123-3213123212113321-2002320013022301-0103120200121311-1130330122013333-3123020310230233-1310312311121133-1310200110003210"></a>

## brute_force_detection — brute_force_detection / 031221332131 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- brute_force_detection

<a id="canonical-1020332221103212-3021131122203101-1013301032133200-2032311123031210-0322032303003120-2033022230333022-2131020200033000-3312230222211130"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for brute force detection.

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
brute_force_detection {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012210122033300-0322120212310202-3023201312130001-3310103332012323-0032211022101101-3303211301011000-0021302030202021-0022331033212101"></a>

## Direct properties — brute_force_detection / 031221332131 / 3

<a id="canonical-0120213103212012-0310313212312131-0301300131231002-2000002021120221-2130020120311200-0221103223132233-0110213001112231-0110132230000200"></a>

<a id="canonical-2300322331101101-1203232102133121-2200113321132133-3022231103020303-2000333232332130-0131003233223332-2332333221122030-1313102101211210"></a>

## max_login_failures property — brute_force_detection / 031221332131 / 4

Type: `"number"`. Optional.

How many failures before wait is triggered. When login failure count is hit, user will be
temporarily locked for a max duration of 15 minutes.

Upstream description:

How many failures before wait is triggered. When login failure count is hit, user will be
temporarily locked for a max duration of 15 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-2011301310031123-2021323031101100-1231202113201333-3030023012331020-3000322012012110-0233230303013023-2320131222011313-3112012030100033"></a>

## Next pages — brute_force_detection / 031221332131 / 5

- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)

<a id="canonical-3102022230023020-1021013032122012-3020123221123311-3313322012121231-3031033300031300-3312331012000020-2031003212112110-2203300222010022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313101310223231-3203332202330130-2123201230002022-1333312211312233-1213103011013001-2000210200112101-2123220332101321-1202313323121031"></a>

## password_policy — password_policy / 223013301201 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- password_policy

<a id="canonical-3022112100000001-1030120011031300-0313030102002100-1313130130120201-2011030231103023-0201322121310103-3112322203113122-0032120011202033"></a>

Type: `"object"`. single nested block, Optional.

Policy configuration for this feature.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("minimum_length")}
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
password_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2100130213330120-3211211031031120-1233120310322100-3101123300332012-0303101110313300-0222201202222231-3333302011223321-3231213100320121"></a>

## Direct properties — password_policy / 223013301201 / 3

<a id="canonical-1330232212122300-0331013323013211-3112101020302220-2203332032301131-1321103003123321-0003303201331022-0211033011120231-3023301303001220"></a>

<a id="canonical-1000201222322332-1300102001020100-3010213302020000-3102020120101012-0203132233123000-1220231002111210-2230313222333212-0012213222211000"></a>

## digits property — password_policy / 223013301201 / 4

Type: `"number"`. Optional.

The number of digits required to be in the password string.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-2311232122303323-1220332223132321-2313210123112132-2311233212002233-1120332002011312-3110211313220023-2011311330211011-1223001120101030"></a>

<a id="canonical-0131033103231020-2021200132313033-1200222332321102-3022131120332300-3001223300032010-0133120033230122-1300120011220131-3233022312323202"></a>

## expire_password property — password_policy / 223013301201 / 5

Type: `"number"`. Optional.

The number of days for which the password is valid. After the number of days has expired, the user
is required to change their password.

Upstream description:

The number of days for which the password is valid. After the number of days has expired, the user
is required to change their password.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1080),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1080,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "1080"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1080"
  }
}
```

<a id="canonical-3331321310123223-0122233230313000-1000030300221020-1131110302233000-2301101210210300-0131112232211133-1301013033011323-0012212003233320"></a>

<a id="canonical-3101102100211030-0331213330011303-3021020123221312-2001312303211212-2211001222311132-2133021001111233-0130012310311322-2113133301110111"></a>

## lowercase_characters property — password_policy / 223013301201 / 6

Type: `"number"`. Optional.

The number of lower case letters required to be in the password string.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-1233002113021031-1122102133103310-3032321000321231-2220323010323323-1020030330220313-3110000022023201-3220211221103002-3020110010221132"></a>

<a id="canonical-1033302103101200-3203032130113020-2022120100032133-2232303003031123-3031123002100011-2221030022333020-0131000302012200-2102031201300211"></a>

## minimum_length property — password_policy / 223013301201 / 7

Type: `"number"`. Optional.

Minimum Length. Minimum length of password.

Upstream description:

Minimum length of password.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(7),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 7
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "7"
  }
}
```

<a id="canonical-2330000232103132-2001212130232210-3133112031123132-1220223222212200-1203221331201030-2002321031200222-0322213220110112-1122130220001010"></a>

<a id="canonical-1232220013023223-0311003202122333-1022223202220101-2202022323212130-3002130331231231-0212313023330213-1322222033231321-0101330300223131"></a>

## not_recently_used property — password_policy / 223013301201 / 8

Type: `"number"`. Optional.

Policy is used to restrict user from using previously used passwords. Number that's set determines
number of last passwords which user cannot use as new password.

Upstream description:

This policy is used to restrict user from using previously used passwords. Number that's set
determines number of last passwords which user cannot use as new password.

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

<a id="canonical-1221101110133111-1303303033201102-0300210111033300-2212013321231321-0231311203133011-0103323013211022-2122230001032213-0320203003130122"></a>

<a id="canonical-2331001023232323-0032001232200200-1300132223213100-2003201302330210-0332131120212013-0331020013233321-0311221313010102-2000222321330211"></a>

## not_username property — password_policy / 223013301201 / 9

Type: `"bool"`. Optional.

When set, the password is not allowed to be the same as the username.

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

<a id="canonical-0311330330000203-2132120100110022-0310223221010120-0220303111322012-3130233212021033-2113123122011012-0313203200202001-0230112101300203"></a>

<a id="canonical-0012212132212201-1203332333331020-3001031123330130-3322000313200000-2310102031120221-1201020301201132-0033322203012222-2202232123102113"></a>

## special_characters property — password_policy / 223013301201 / 10

Type: `"number"`. Optional.

The number of special characters like '?!\#%$' required to be in the password string.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-3311003131210010-3033212320333220-2303120210131031-1100102112331130-1300033223320301-0303210000001313-0100302322113101-1012333201221130"></a>

<a id="canonical-2112300320010203-1221021211012300-1120011211201012-3012201011113233-2213323033212003-0011123220203123-2013021312311133-3223010230011112"></a>

## uppercase_characters property — password_policy / 223013301201 / 11

Type: `"number"`. Optional.

The number of upper case letters required to be in the password string.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-1030222100011222-3011223330121230-1212113310230201-0312013323231030-2001233113030111-1000022112122202-2130321010120110-0221231230102232"></a>

## Next pages — password_policy / 223013301201 / 12

- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)

<a id="canonical-1011022202202111-0102220102231000-2232303112122022-0132312233310312-2213203313133203-1210320213013321-3222001331131113-3331130300130333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211313001022311-1332301031320123-2333301311301332-0102233111233201-1312222002203112-2222111020332311-1312033131000020-2122031012223101"></a>

## tenant_details — tenant_details / 021020313303 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- tenant_details

<a id="canonical-2223012132011230-2321231303113020-2201102130002012-2112211222133231-2032323223320211-3001132010330202-0032112201022211-0132230233131231"></a>

Type: `"object"`. single nested block, Optional.

BasicConfiguration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("display_name")}
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
tenant_details {
  # Configure direct properties listed below.
}
```

<a id="canonical-3321212231101323-0210030231310231-2000203101013301-0101133200203322-0000100133213021-1200112220122202-2022333133230330-3133100032022210"></a>

## Direct properties — tenant_details / 021020313303 / 3

<a id="canonical-2012121132033311-1101102113131323-2230022231210302-0121301030311023-3010021131101032-3131102231031123-3132101013302232-2101022101022131"></a>

<a id="canonical-3021210121220012-1120130111103302-3011110233031202-0020010330301230-0202010110000112-2103023111303110-0220010333110301-0313201133321112"></a>

## display_name property — tenant_details / 021020313303 / 4

Type: `"string"`. Optional.

Changes the tenant name displayed during login without affecting your company’s domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[^<>&\\\"]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^[^<>&\\\"]+$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^[^<>&\\\"]+$"
  }
}
```

<a id="canonical-0223001121233122-1030211220010213-0012112001332103-0333201321122033-0223132332332031-1212311003322200-1121022203313113-0231031300120320"></a>

## Next pages — tenant_details / 021020313303 / 5

- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)

<a id="canonical-0211120030131233-0221232021112002-3012103322332003-1021113132232201-2312122022030320-1231200221002212-0122322000130230-2300121212012010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102011330233212-0233210130121131-3033112221011021-0102332313100003-3120033330213212-0113221303310303-1110002103320302-1302223101210233"></a>

## timeouts — timeouts / 321020223012 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- timeouts

<a id="canonical-1300230233221312-1213330312333323-1323102311230113-3112131012111032-2303133203210110-1232313333133213-0310030033011212-0020320020112313"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0331032120111223-2211021333333330-1301311322231011-0300213122201113-2100202112200222-1220310131230110-2010300032100212-2002102102010312"></a>

## Direct properties — timeouts / 321020223012 / 3

<a id="canonical-1133122021001301-1212301202001332-3203333130230222-1031331301231003-3230121011311103-2301000301130222-0030300030211013-2112101103132311"></a>

<a id="canonical-1222121231003230-3231233030333012-3031333110232101-1103221033010100-3320203032101112-3203102121003212-3320131101230220-2112313012030303"></a>

## create property — timeouts / 321020223012 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1113121012113000-1201202121003222-0021101213020132-2031321303001223-0313010033333233-1132123112113123-2333130233203220-2121100001331021"></a>

<a id="canonical-3023320311021211-0002332113332131-3302030220120132-3222313331222132-2012301321121322-1231100222301303-3030003332111202-1221312000332012"></a>

## delete property — timeouts / 321020223012 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0030111303112312-2310222303112320-0000220221311133-3000213111331222-3200023032113213-0313013033132221-0110230222303001-3131330001321310"></a>

<a id="canonical-2110333223003113-0320223021130012-2100213210330021-0311102103223133-1123332333031133-0110130210103212-0200222113011230-2030033221322122"></a>

## read property — timeouts / 321020223012 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1013210300323121-1022300131032301-2220222311103301-0122101000030230-1232232100002333-3032222311131132-3233030303123202-2313130300112131"></a>

<a id="canonical-1130133213123100-0111300211010331-2200203302030032-2202221012003213-3110223320303300-3212012233021111-3003002010303131-1331123323200332"></a>

## update property — timeouts / 321020223012 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2121110100311303-2211022031202233-0000210111111130-1112330313330313-1122102110223331-0020202313331300-1113330320130121-3003122103200212"></a>

## Next pages — timeouts / 321020223012 / 8

- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)

<a id="canonical-3002032020221331-0023102232020311-2221003000102010-1231312030013330-1231033101320300-2312123103202330-1322112233100103-2133031101202322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311233002232202-3112213130232002-3030121331030011-0113311003211322-2323211332023131-3031310122032122-1330132212301121-2002110300232110"></a>

## user_session_expiration — user_session_expiration / 310031301223 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- user_session_expiration

<a id="canonical-2101223000110110-0300010313112300-2200021311002331-1121011122003303-3123103133301311-3132200231333221-2103111001001011-0302033223101302"></a>

Type: `"object"`. single nested block, Optional.

Defines all session-related expiration for user sessions within a tenant's environment. Relationship
between session\_expiry and cookie\_expiry: - session\_expiry defines the 'absolute maximum
duration' of a session and enforces RE-authentication after this time. - cookie\_expiry defines
the..

Upstream description:

Defines all session-related expiration for user sessions within a tenant's environment. Relationship
between session\_expiry and cookie\_expiry: &#8203;- session\_expiry defines the 'absolute maximum
duration' of a session and enforces RE-authentication after this time. &#8203;- cookie\_expiry
defines the 'inactivity timeout', which resets on user activity and only logs out users after idle
periods. Together, these ensure the user is logged out when either the session reaches its maximum
age or the user is inactive for too long.

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
user_session_expiration {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230023212333303-1333303101131332-1101110202133123-0100110011231322-1333321310020312-1133311112121300-3230300232222231-2113013013223001"></a>

## Direct properties — user_session_expiration / 310031301223 / 3

- [absolute_timeout](resources--tenant_configuration--reference--group-001.md#canonical-0131112302222330-0332200030123212-0100031133203323-0333330313033111-2211333123210323-0231123233100022-3202031211002022-2310113122310211): complete subsection reference.

- [idle_timeout](resources--tenant_configuration--reference--group-001.md#canonical-2013211033012201-1303322322012201-0030031200023130-3301130101231210-0113013331113100-0310232130230220-1232000031122322-1301202130122312): complete subsection reference.

<a id="canonical-3221201313201221-2023212103213302-2022302222011122-0221200222001130-2300230023113021-0212311220130313-0312122033022332-2110033230313101"></a>

## Next pages — user_session_expiration / 310031301223 / 4

- [user_session_expiration.absolute_timeout](resources--tenant_configuration--reference--group-001.md#canonical-0131112302222330-0332200030123212-0100031133203323-0333330313033111-2211333123210323-0231123233100022-3202031211002022-2310113122310211)
- [user_session_expiration.idle_timeout](resources--tenant_configuration--reference--group-001.md#canonical-2013211033012201-1303322322012201-0030031200023130-3301130101231210-0113013331113100-0310232130230220-1232000031122322-1301202130122312)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)

<a id="canonical-0131112302222330-0332200030123212-0100031133203323-0333330313033111-2211333123210323-0231123233100022-3202031211002022-2310113122310211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312021221313113-0203301320212211-0301100022302132-3132202301231222-1120200310202020-3322102031110002-1332113110011010-1011100112110000"></a>

## user_session_expiration.absolute_timeout — absolute_timeout / 222023230002 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-3002032020221331-0023102232020311-2221003000102010-1231312030013330-1231033101320300-2312123103202330-1322112233100103-2133031101202322)
- user_session_expiration.absolute_timeout

<a id="canonical-3101003133012123-2303213123233200-0321113211113001-3031111323010000-2033211311330323-0202232022120101-1130301203202313-2131220132221323"></a>

Type: `"object"`. single nested block, Optional.

Represents the session expiration duration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("hours",
    "minutes")}
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
  "x-ves-oneof-field-unit_of_time": "[\"hours\",\"minutes\"]"
}
```

Terraform syntax:

```terraform
absolute_timeout {
  # Configure direct properties listed below.
}
```

<a id="canonical-0023231310010213-0333003212003321-2211300230300231-3132101200133033-2321312113031122-2323131301021221-1033013331132333-1230311100113111"></a>

## Direct properties — absolute_timeout / 222023230002 / 3

- [hours](resources--tenant_configuration--reference--group-001.md#canonical-0330301320200331-2313011321301203-3313122002101001-1302223220323120-2213020111132231-1312031033032000-2333023000333322-1021102201003003): complete subsection reference.

- [minutes](resources--tenant_configuration--reference--group-001.md#canonical-2121112223201202-1330003023120300-2323300211223301-3001203311001012-1003200020001023-0202313310203112-0320132201113220-0012211331121033): complete subsection reference.

<a id="canonical-1230100200110010-0112020111310303-3323021202222221-2032313331130201-0313100102332231-1211023031103332-1331232011200101-1002310313113202"></a>

## Next pages — absolute_timeout / 222023230002 / 4

- [user_session_expiration.absolute_timeout.hours](resources--tenant_configuration--reference--group-001.md#canonical-0330301320200331-2313011321301203-3313122002101001-1302223220323120-2213020111132231-1312031033032000-2333023000333322-1021102201003003)
- [user_session_expiration.absolute_timeout.minutes](resources--tenant_configuration--reference--group-001.md#canonical-2121112223201202-1330003023120300-2323300211223301-3001203311001012-1003200020001023-0202313310203112-0320132201113220-0012211331121033)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-3002032020221331-0023102232020311-2221003000102010-1231312030013330-1231033101320300-2312123103202330-1322112233100103-2133031101202322)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)

<a id="canonical-0330301320200331-2313011321301203-3313122002101001-1302223220323120-2213020111132231-1312031033032000-2333023000333322-1021102201003003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100100120003103-3003002030230231-1232013230331132-3220123013320323-3231320030312233-1010232010201000-3320221002321022-0113011310001212"></a>

## user_session_expiration.absolute_timeout.hours — hours / 311022302030 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-3002032020221331-0023102232020311-2221003000102010-1231312030013330-1231033101320300-2312123103202330-1322112233100103-2133031101202322)
- [user_session_expiration.absolute_timeout](resources--tenant_configuration--reference--group-001.md#canonical-0131112302222330-0332200030123212-0100031133203323-0333330313033111-2211333123210323-0231123233100022-3202031211002022-2310113122310211)
- user_session_expiration.absolute_timeout.hours

<a id="canonical-2333001321320010-2131301012010203-1203202013122020-3320333210013020-0001211200033102-2200003321231220-0102232323023123-0011111331320111"></a>

Type: `"object"`. single nested block, Optional.

Represents the session duration in hours.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-3031310033120201-3023321322233103-2012003300113232-2030121102300332-2123312303031033-1332232013020200-0300232133102332-3211220222303030"></a>

## Direct properties — hours / 311022302030 / 3

<a id="canonical-3112030302313023-3033320220100203-2131123113001233-3100101320032231-3011100233113011-0330032232323002-1201332232213113-1222311221122113"></a>

<a id="canonical-0312123003011011-2220002113332120-2230213002302110-0321122222022332-1212032213310213-3120000320121212-0323300202230210-1203121202333012"></a>

## duration property — hours / 311022302030 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 720),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 720,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "720"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "720"
  }
}
```

<a id="canonical-0222333101021100-1300231203222200-2313021000023000-3312320001003022-1010022211003022-2132221023320220-2113321132010130-3222211101030202"></a>

## Next pages — hours / 311022302030 / 5

- [user_session_expiration.absolute_timeout](resources--tenant_configuration--reference--group-001.md#canonical-0131112302222330-0332200030123212-0100031133203323-0333330313033111-2211333123210323-0231123233100022-3202031211002022-2310113122310211)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)

<a id="canonical-2121112223201202-1330003023120300-2323300211223301-3001203311001012-1003200020001023-0202313310203112-0320132201113220-0012211331121033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220022310111023-0300002032001033-1020323323121022-3021132212113001-3232332011303231-2203011310122222-1210313331203120-0132003003221011"></a>

## user_session_expiration.absolute_timeout.minutes — minutes / 011221032203 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-3002032020221331-0023102232020311-2221003000102010-1231312030013330-1231033101320300-2312123103202330-1322112233100103-2133031101202322)
- [user_session_expiration.absolute_timeout](resources--tenant_configuration--reference--group-001.md#canonical-0131112302222330-0332200030123212-0100031133203323-0333330313033111-2211333123210323-0231123233100022-3202031211002022-2310113122310211)
- user_session_expiration.absolute_timeout.minutes

<a id="canonical-3333210001012231-0000121110310111-3023221201302332-1002330232332021-0002021221212102-0230220330131111-1200300101000111-1233033322333003"></a>

Type: `"object"`. single nested block, Optional.

Represents the session duration in minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011101321113311-2003211302200321-0000000200231331-2332330011031303-3231232013301032-0230133130212112-2332210222212233-0122313120013000"></a>

## Direct properties — minutes / 011221032203 / 3

<a id="canonical-3313123133320102-1321031013000203-0122213201122102-2011010133210000-3311303203321212-0301033223113120-3000213312332330-1033203200113300"></a>

<a id="canonical-3210012130201011-2102032201320321-2103112200300331-0203031223321021-1121322330110311-1131211310111322-1012110101133301-1130301333230131"></a>

## duration property — minutes / 011221032203 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(5, 43200),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 43200,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 5
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "5",
    "ves.io.schema.rules.uint32.lte": "43200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "5",
    "ves.io.schema.rules.uint32.lte": "43200"
  }
}
```

<a id="canonical-1000033322123301-2013030121111130-3113233301032030-2310122120132210-0030103330023301-1321232320330300-3200111121103323-3021100300030122"></a>

## Next pages — minutes / 011221032203 / 5

- [user_session_expiration.absolute_timeout](resources--tenant_configuration--reference--group-001.md#canonical-0131112302222330-0332200030123212-0100031133203323-0333330313033111-2211333123210323-0231123233100022-3202031211002022-2310113122310211)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)

<a id="canonical-2013211033012201-1303322322012201-0030031200023130-3301130101231210-0113013331113100-0310232130230220-1232000031122322-1301202130122312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120201321120301-3330132300322023-0221220200330121-2122313130230121-0031101303000132-0320302303103231-1223031200032121-0102121000033131"></a>

## user_session_expiration.idle_timeout — idle_timeout / 100133301232 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-3002032020221331-0023102232020311-2221003000102010-1231312030013330-1231033101320300-2312123103202330-1322112233100103-2133031101202322)
- user_session_expiration.idle_timeout

<a id="canonical-1130010230321001-0021001133000010-0200313133311301-3101120120030311-1032201133103200-3003021322200030-2131120003033020-0132202223332212"></a>

Type: `"object"`. single nested block, Optional.

Represents the cookie expiration duration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("hours",
    "minutes")}
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
  "x-ves-oneof-field-unit_of_time": "[\"hours\",\"minutes\"]"
}
```

Terraform syntax:

```terraform
idle_timeout {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213220112231002-3322213000110131-3123202303011121-0021110020130021-3322010010132200-0330030300231232-0101130310123030-0102112022102312"></a>

## Direct properties — idle_timeout / 100133301232 / 3

- [hours](resources--tenant_configuration--reference--group-001.md#canonical-0012002003312021-2323222033233212-3003131023020020-1220011333001232-3120011001131131-3130331100010322-3110102020220131-0230010021201010): complete subsection reference.

- [minutes](resources--tenant_configuration--reference--group-001.md#canonical-3020320022111123-3123203123322032-0000323233022332-0022001003012102-1213020330012010-3011220232013300-0003202002001000-3130030112121001): complete subsection reference.

<a id="canonical-3121210010031201-3201122202010330-3231321320322320-1313331021312203-2303223101101310-3310130100222020-1002210011013100-3330012230313201"></a>

## Next pages — idle_timeout / 100133301232 / 4

- [user_session_expiration.idle_timeout.hours](resources--tenant_configuration--reference--group-001.md#canonical-0012002003312021-2323222033233212-3003131023020020-1220011333001232-3120011001131131-3130331100010322-3110102020220131-0230010021201010)
- [user_session_expiration.idle_timeout.minutes](resources--tenant_configuration--reference--group-001.md#canonical-3020320022111123-3123203123322032-0000323233022332-0022001003012102-1213020330012010-3011220232013300-0003202002001000-3130030112121001)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-3002032020221331-0023102232020311-2221003000102010-1231312030013330-1231033101320300-2312123103202330-1322112233100103-2133031101202322)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)

<a id="canonical-0012002003312021-2323222033233212-3003131023020020-1220011333001232-3120011001131131-3130331100010322-3110102020220131-0230010021201010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331020132012100-3222112023131003-1213111303003020-3021222212303122-2323120111132103-3131030313031123-3331303200103200-2103130231323101"></a>

## user_session_expiration.idle_timeout.hours — hours / 021020213323 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-3002032020221331-0023102232020311-2221003000102010-1231312030013330-1231033101320300-2312123103202330-1322112233100103-2133031101202322)
- [user_session_expiration.idle_timeout](resources--tenant_configuration--reference--group-001.md#canonical-2013211033012201-1303322322012201-0030031200023130-3301130101231210-0113013331113100-0310232130230220-1232000031122322-1301202130122312)
- user_session_expiration.idle_timeout.hours

<a id="canonical-1221030310003031-1122022113230232-1310010122210220-1123312100230133-1120030133232221-0003220132101333-0122131330113322-2003223001330322"></a>

Type: `"object"`. single nested block, Optional.

Represents the cookie duration in hours.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301032001130232-1331020232201121-2223232003232012-2210020313111121-3222321321221230-1312130030200101-1200231303002312-2002311230010021"></a>

## Direct properties — hours / 021020213323 / 3

<a id="canonical-3023122233113030-1220333122330320-0130202311013213-0300103013221333-0213131203210101-3030100331321131-0031213201320333-0201312103333222"></a>

<a id="canonical-3323031112113232-1112321001113333-0133332203111223-1223210230213132-1212320223022110-3231210211230131-2002332230131120-0013123212231003"></a>

## duration property — hours / 021020213323 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 720),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 720,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "720"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "720"
  }
}
```

<a id="canonical-0203230223011121-0203023311322021-3112100312203323-2011122111031030-3033030211120202-2322103112030111-0132122303220221-2121001103300300"></a>

## Next pages — hours / 021020213323 / 5

- [user_session_expiration.idle_timeout](resources--tenant_configuration--reference--group-001.md#canonical-2013211033012201-1303322322012201-0030031200023130-3301130101231210-0113013331113100-0310232130230220-1232000031122322-1301202130122312)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)

<a id="canonical-3020320022111123-3123203123322032-0000323233022332-0022001003012102-1213020330012010-3011220232013300-0003202002001000-3130030112121001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321003201102332-3123132321120310-2332101201311310-3000322102032332-3003333113130321-0012001213103021-2231211021020321-1233320100221111"></a>

## user_session_expiration.idle_timeout.minutes — minutes / 033232112330 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)
- [Property reference](resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- [user_session_expiration](resources--tenant_configuration--reference--group-001.md#canonical-3002032020221331-0023102232020311-2221003000102010-1231312030013330-1231033101320300-2312123103202330-1322112233100103-2133031101202322)
- [user_session_expiration.idle_timeout](resources--tenant_configuration--reference--group-001.md#canonical-2013211033012201-1303322322012201-0030031200023130-3301130101231210-0113013331113100-0310232130230220-1232000031122322-1301202130122312)
- user_session_expiration.idle_timeout.minutes

<a id="canonical-1013032210221213-0112203001130101-3033212020023113-2021033011001100-2330001100322230-0213221220032131-1022123332223101-1002023220211222"></a>

Type: `"object"`. single nested block, Optional.

Represents the cookie duration in minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1032132021010122-3333221100202031-1010120203033300-0231222300121013-1312201312322131-2021330112310002-1201223312311111-1030231023331313"></a>

## Direct properties — minutes / 033232112330 / 3

<a id="canonical-2021023230133101-1003312231213220-1012200212131102-3332210012100303-1122103310323212-2201012113103101-1101300131001010-1311233113123232"></a>

<a id="canonical-0312113232030302-0010330002030010-3221302303323312-0310022201123113-3032303332220103-1033002322312010-0112300303033130-3211313122211023"></a>

## duration property — minutes / 033232112330 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(5, 43200),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 43200,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 5
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "5",
    "ves.io.schema.rules.uint32.lte": "43200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "5",
    "ves.io.schema.rules.uint32.lte": "43200"
  }
}
```

<a id="canonical-3000020331311031-1230211133233000-2032201101003112-0220021212102031-3001022131013113-1123122021012013-3313211300131102-1012220132101333"></a>

## Next pages — minutes / 033232112330 / 5

- [user_session_expiration.idle_timeout](resources--tenant_configuration--reference--group-001.md#canonical-2013211033012201-1303322322012201-0030031200023130-3301130101231210-0113013331113100-0310232130230220-1232000031122322-1301202130122312)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)
