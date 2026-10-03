---
page_title: "xcsh_fast_acl reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl reference."
---

# xcsh_fast_acl reference

<a id="canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213103113313201-0003332323113202-2220330021312223-2131231022232332-2220303310203133-0001210321211120-2110010003311222-0001223231203210"></a>

## Property reference — Property reference / 020322101001 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- Property reference

<a id="canonical-1000000103333031-1002200310330000-3023200101321021-2312233333130120-1200320131210113-3331013310303210-1101131011011321-0331011323002111"></a>

## Direct properties — Property reference / 020322101001 / 3

<a id="canonical-0031122312202013-2013213200001003-0011103210331100-2320220100123101-0032020310313020-3131102233022221-3211331012303331-2203330131110013"></a>

<a id="canonical-2130031301131103-1012303300020332-3311211312321130-0203310222013111-2333320002223220-3013330320222021-2000301323031002-2320311212303310"></a>

## annotations property — Property reference / 020322101001 / 4

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

<a id="canonical-2311201130333220-1012302302202002-3222023113002132-1332210022131032-3132110223023001-0130220310110122-2232323303121333-2001231203120012"></a>

<a id="canonical-0112321003223310-3310300102201101-0020333200033001-3203103112202112-1120013303201220-3322320211000133-0022110320133131-1231222200332030"></a>

## description property — Property reference / 020322101001 / 5

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

<a id="canonical-3132232033232033-3211202030011312-1331330322103323-1312132200201322-3130021203023323-1021221121300133-0030330120003022-3331333330023232"></a>

<a id="canonical-0103012300123331-0022012211220020-0213300022032132-0211333122230002-1221001332321201-2103012231022312-1231331003313000-0312013010013101"></a>

## disable property — Property reference / 020322101001 / 6

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

<a id="canonical-3202112020133102-2112013133321013-1003221310330000-0132113031103130-1200330323222331-1330031321230220-1222032300220131-2132300102111233"></a>

<a id="canonical-2113103332121313-3321013221223013-3333031223312132-1110203012033001-0310230132030303-0022331211132122-3133122132003322-2302203321020000"></a>

## ID property — Property reference / 020322101001 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1131021122203012-1202313220321321-3033232323201312-0311211233103133-2013212031000302-1030013320003023-3333312010112013-1222000330110100"></a>

<a id="canonical-0331022133213312-0030001002320013-2212001221320331-0221021122233013-1103331331332222-1222131212013300-0010300230000311-3301120010033022"></a>

## labels property — Property reference / 020322101001 / 8

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

<a id="canonical-0310102203230223-2320001023100032-2331130012300121-3300302023203101-0000100201010220-3031100202103233-0030303122203131-2221101301011303"></a>

<a id="canonical-1202221001321130-2011233003130323-2213032320022320-3010010123202231-1021001102321212-2122020002021112-3010103313130323-3120121311322223"></a>

## name property — Property reference / 020322101001 / 9

Type: `"string"`. Required.

Name of the Fast ACL. Must be unique within the namespace.

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

<a id="canonical-3221210312201210-3133122322231012-3320231230313022-0110301012130103-1022203202103002-0001103003022220-2301120031033211-3202202303232221"></a>

<a id="canonical-3030033202331331-2311110320133110-0113321131010131-2210332011010001-0020233333132301-1130333000031222-1122130232022220-0212002332320311"></a>

## namespace property — Property reference / 020322101001 / 10

Type: `"string"`. Optional, Computed.

Namespace for the Fast ACL. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [protocol_policer](resources--fast_acl--reference--group-001.md#canonical-3213203231322131-2302120011233022-1101333002320310-0010010133131023-1310101021330233-0023121220213301-1032021020200222-1000230202033211): complete subsection reference.

- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221): complete subsection reference.

- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322): complete subsection reference.

- [timeouts](resources--fast_acl--reference--group-001.md#canonical-2222200211213322-0210312320233133-3101222313001002-3200023112003003-2231122211030313-0030223121111320-3321112132323230-1211100021300311): complete subsection reference.

<a id="canonical-0211323130221033-1002323203000021-0322301220310332-3312023033123002-1032310232001233-1223230120200222-3312203311213011-3212122211213020"></a>

## All schema paths — Property reference / 020322101001 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--fast_acl--reference--group-001.md#canonical-0031122312202013-2013213200001003-0011103210331100-2320220100123101-0032020310313020-3131102233022221-3211331012303331-2203330131110013) |
| `description` | [description](resources--fast_acl--reference--group-001.md#canonical-2311201130333220-1012302302202002-3222023113002132-1332210022131032-3132110223023001-0130220310110122-2232323303121333-2001231203120012) |
| `disable` | [disable](resources--fast_acl--reference--group-001.md#canonical-3132232033232033-3211202030011312-1331330322103323-1312132200201322-3130021203023323-1021221121300133-0030330120003022-3331333330023232) |
| `id` | [ID](resources--fast_acl--reference--group-001.md#canonical-3202112020133102-2112013133321013-1003221310330000-0132113031103130-1200330323222331-1330031321230220-1222032300220131-2132300102111233) |
| `labels` | [labels](resources--fast_acl--reference--group-001.md#canonical-1131021122203012-1202313220321321-3033232323201312-0311211233103133-2013212031000302-1030013320003023-3333312010112013-1222000330110100) |
| `name` | [name](resources--fast_acl--reference--group-001.md#canonical-0310102203230223-2320001023100032-2331130012300121-3300302023203101-0000100201010220-3031100202103233-0030303122203131-2221101301011303) |
| `namespace` | [namespace](resources--fast_acl--reference--group-001.md#canonical-3221210312201210-3133122322231012-3320231230313022-0110301012130103-1022203202103002-0001103003022220-2301120031033211-3202202303232221) |
| `protocol_policer` | [protocol_policer](resources--fast_acl--reference--group-001.md#canonical-3102103313301300-2231233023321201-0113333233012333-2300001103232101-0003223113230333-2312121230013321-0003103221000030-3210320201230331) |
| `protocol_policer.name` | [protocol_policer.name](resources--fast_acl--reference--group-001.md#canonical-1231302002211331-3110223031112310-2003231201012233-0122222001221232-3102231113120003-0211111123033233-2002132102112232-0212313100133100) |
| `protocol_policer.namespace` | [protocol_policer.namespace](resources--fast_acl--reference--group-001.md#canonical-1012111003333223-2201311203302303-1000333300232030-0212023021122031-1321000300021033-3023113003030300-1103312010103302-2322010123332003) |
| `protocol_policer.tenant` | [protocol_policer.tenant](resources--fast_acl--reference--group-001.md#canonical-3110322121023032-0123303332020013-2013212112120120-2111303001230123-3133132101112333-1211220222102122-2110331102023232-2323220301313201) |
| `re_acl` | [re_acl](resources--fast_acl--reference--group-001.md#canonical-2031331223221133-0333222313002321-3131001313230032-1033233112222331-1323223232332312-3313200032102321-3200103002320021-1213213032213311) |
| `re_acl.all_public_vips` | [re_acl.all_public_vips](resources--fast_acl--reference--group-001.md#canonical-1123232130300230-3222212333303332-3312301131002112-3312133033012133-3011222332311033-0203310222123233-0112010312311310-3121200101221100) |
| `re_acl.default_tenant_vip` | [re_acl.default_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-0102101121211100-2020132333233111-2303310112003122-0031210132022231-0023210020112100-1232313210012011-1311223213230222-1311112131230201) |
| `re_acl.fast_acl_rules` | [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1110303332131332-3201323100322330-3320332003010313-2332003330311003-3313221232220322-1102031232130311-0332323211032322-0230122010223113) |
| `re_acl.fast_acl_rules.action` | [re_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-3100003210032133-1300133020003332-0022121332120103-3233013230103232-2212322111212030-1123301021202133-1202132322321030-2322032000320200) |
| `re_acl.fast_acl_rules.action.policer_action` | [re_acl.fast_acl_rules.action.policer_action](resources--fast_acl--reference--group-001.md#canonical-2221321200021220-1021231010323222-0003223032203322-1301103222010312-0323000011012021-2222223120033131-2323321000032131-1331020003220232) |
| `re_acl.fast_acl_rules.action.policer_action.ref` | [re_acl.fast_acl_rules.action.policer_action.ref](resources--fast_acl--reference--group-001.md#canonical-3303301232300322-0201201333333222-3332111123121223-1311201100222110-1330311333201000-1203211322003001-0100113022120332-1233110002311323) |
| `re_acl.fast_acl_rules.action.policer_action.ref.kind` | [re_acl.fast_acl_rules.action.policer_action.ref.kind](resources--fast_acl--reference--group-001.md#canonical-2030102221303120-1330302012323120-0013121121202231-0223203203300230-3023330313110000-2202032131133102-2033112111323322-0210223230133333) |
| `re_acl.fast_acl_rules.action.policer_action.ref.name` | [re_acl.fast_acl_rules.action.policer_action.ref.name](resources--fast_acl--reference--group-001.md#canonical-2131201312023011-3020110200120121-1123133003112312-2222013230222021-2031222003313121-1220003333310023-0321202102220132-1212310132102303) |
| `re_acl.fast_acl_rules.action.policer_action.ref.namespace` | [re_acl.fast_acl_rules.action.policer_action.ref.namespace](resources--fast_acl--reference--group-001.md#canonical-0212033232012031-0132220230222233-3010321012021031-2331233313130333-0031000301013121-0311111133133233-2102011031300002-0303013100100200) |
| `re_acl.fast_acl_rules.action.policer_action.ref.tenant` | [re_acl.fast_acl_rules.action.policer_action.ref.tenant](resources--fast_acl--reference--group-001.md#canonical-2001232321123122-1210101010132331-3120011303302333-3012300203020213-0311030232321213-1100132321331013-3021000222010331-3200122110020132) |
| `re_acl.fast_acl_rules.action.policer_action.ref.uid` | [re_acl.fast_acl_rules.action.policer_action.ref.uid](resources--fast_acl--reference--group-001.md#canonical-0330020210222020-3210300232202230-3301201031313123-1000033012111230-2011202300202220-3311220023231031-3212013103310222-3223200323111203) |
| `re_acl.fast_acl_rules.action.protocol_policer_action` | [re_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-2133210311112322-2222201220111231-0222111210202131-0101312010212230-0203131023311230-3110320011320201-1003233223221013-0001123133101010) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref](resources--fast_acl--reference--group-001.md#canonical-1313330221332023-1321131011332000-3122301210211113-1032322032310213-0202032132013312-2010323120133031-2032233013202223-2311313013213033) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.kind` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.kind](resources--fast_acl--reference--group-001.md#canonical-0300310132113201-3212321121021121-1010201132020212-0120013301011321-3000010331321312-1310012110211102-1310301211212023-3101131213102020) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.name` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.name](resources--fast_acl--reference--group-001.md#canonical-3233302112110031-3023130001003022-0332022330113212-1323030010113200-0230030023202211-3300011100323102-2223210020121202-0101301002000322) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace](resources--fast_acl--reference--group-001.md#canonical-0322020220223110-3200311120222213-0113022320112210-0110111331213001-2121031122103211-3013123212312113-3321301232030313-3330320030301211) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant](resources--fast_acl--reference--group-001.md#canonical-3001023011110333-0200122312322010-2101203011312121-3130120322122231-0210002120330300-0301320300313001-3001102312020020-1311100310321200) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.uid` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.uid](resources--fast_acl--reference--group-001.md#canonical-0310003232210223-1022322020312002-2113220122310112-3311012231122302-3003203012333002-2113023030331220-2301203011310200-3133220210122103) |
| `re_acl.fast_acl_rules.action.simple_action` | [re_acl.fast_acl_rules.action.simple_action](resources--fast_acl--reference--group-001.md#canonical-1001120102321220-1011322121321310-2111200113012021-3103221202030201-0032232332123003-3012232310032122-0333133013100223-3230310130021100) |
| `re_acl.fast_acl_rules.ip_prefix_set` | [re_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-2301212323203202-0133322022203122-1020010002002110-0311011201131113-1021222003022201-2233122201303010-3021132200033332-1130121312023311) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref` | [re_acl.fast_acl_rules.ip_prefix_set.ref](resources--fast_acl--reference--group-001.md#canonical-0131113102320202-1313021131330100-0130312331220113-0303132111101031-1102213231030312-2132202213010112-1011203303313201-1132213011010022) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.kind` | [re_acl.fast_acl_rules.ip_prefix_set.ref.kind](resources--fast_acl--reference--group-001.md#canonical-3111300002213303-2301331033022321-3210210223201200-1000221033001320-3200022023132210-0000301310212103-3100132030013313-2123232321133022) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.name` | [re_acl.fast_acl_rules.ip_prefix_set.ref.name](resources--fast_acl--reference--group-001.md#canonical-0010013321332001-2300112233102110-1030130103122223-0302123231301113-0212221032323111-1212300202311331-2113312220112133-0030313012010023) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.namespace` | [re_acl.fast_acl_rules.ip_prefix_set.ref.namespace](resources--fast_acl--reference--group-001.md#canonical-1120101320310120-0031130030313101-1011221222000030-2202233231110022-2010130223230120-2131010331223032-0212231323331223-2203210110213022) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.tenant` | [re_acl.fast_acl_rules.ip_prefix_set.ref.tenant](resources--fast_acl--reference--group-001.md#canonical-1213320112131032-2011101221133011-0320012230333203-2303130011302001-2103200232111101-0323201002001222-1310101302213302-1123023203021203) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.uid` | [re_acl.fast_acl_rules.ip_prefix_set.ref.uid](resources--fast_acl--reference--group-001.md#canonical-0320310111322103-3221320030213012-3333322033312110-2123001210200300-3031213122132030-1233333132130003-2121122300222022-0301020102320031) |
| `re_acl.fast_acl_rules.metadata` | [re_acl.fast_acl_rules.metadata](resources--fast_acl--reference--group-001.md#canonical-2033031123311032-2220212123122120-2100323130003020-2303202202300111-1201130132120023-2131220031311320-2312023220103312-3010312111102020) |
| `re_acl.fast_acl_rules.metadata.description_spec` | [re_acl.fast_acl_rules.metadata.description_spec](resources--fast_acl--reference--group-001.md#canonical-2322310021321012-3312021012312223-1032321001322301-2033231010213000-3003122021020300-0313233000223232-0030101023323033-0233033111333112) |
| `re_acl.fast_acl_rules.metadata.name` | [re_acl.fast_acl_rules.metadata.name](resources--fast_acl--reference--group-001.md#canonical-1002131002303103-0302120012022022-0321132232110031-1202020200211313-1021100223200221-3202100011323232-2303111330033220-0230331121021003) |
| `re_acl.fast_acl_rules.port` | [re_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-3022310110023111-3200133200322213-2232223122230201-0330002020330223-3023132212322310-1032321110131212-2203013300012111-2313231110323002) |
| `re_acl.fast_acl_rules.port.all` | [re_acl.fast_acl_rules.port.all](resources--fast_acl--reference--group-001.md#canonical-2131230023003202-0130322001023200-0021100003102210-1201322101203202-2130231121223212-2310031132300322-3133110320222102-1222220023233013) |
| `re_acl.fast_acl_rules.port.dns` | [re_acl.fast_acl_rules.port.dns](resources--fast_acl--reference--group-001.md#canonical-2032313011223231-1223223312333021-3310033132010001-2233300230121322-2313130010033300-1012021233111330-0113011132313113-2121223221331213) |
| `re_acl.fast_acl_rules.port.user_defined` | [re_acl.fast_acl_rules.port.user_defined](resources--fast_acl--reference--group-001.md#canonical-3232212313000203-3223032223313312-0212202211122222-3322020200020010-1331001132231311-3012022033303033-0021211022331333-0121003302110130) |
| `re_acl.fast_acl_rules.prefix` | [re_acl.fast_acl_rules.prefix](resources--fast_acl--reference--group-001.md#canonical-1012333012203023-1201331111213100-0131110023002222-0100223113221120-3013221023310303-0012131331222023-2001300131312310-3101003023222231) |
| `re_acl.fast_acl_rules.prefix.prefix` | [re_acl.fast_acl_rules.prefix.prefix](resources--fast_acl--reference--group-001.md#canonical-2303131030001221-3111320203111332-0220033222312301-3313132230101120-1100330112331031-2100021113010212-1000212031131132-3212013211310120) |
| `re_acl.selected_tenant_vip` | [re_acl.selected_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-0033230212102001-0013122110231101-2021311311032200-3212000132322202-0120230311112211-3330132001223300-1302020003311012-1211110232122210) |
| `re_acl.selected_tenant_vip.default_tenant_vip` | [re_acl.selected_tenant_vip.default_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-1131333232223032-0011332132331201-0112001000011101-0123301213200303-1023211021120133-2320122101012033-2123111333302203-2221312212022223) |
| `re_acl.selected_tenant_vip.public_ip_refs` | [re_acl.selected_tenant_vip.public_ip_refs](resources--fast_acl--reference--group-001.md#canonical-3220331011030310-1022210133001210-1300333123001300-3011220013331220-2231302121303331-0320322031000011-2212320120013301-3330220102033211) |
| `re_acl.selected_tenant_vip.public_ip_refs.name` | [re_acl.selected_tenant_vip.public_ip_refs.name](resources--fast_acl--reference--group-001.md#canonical-2221023021030002-0333223230110311-0022220302223122-0013220313000322-1301333232201021-0131223210302311-1301001012310203-3013101323001323) |
| `re_acl.selected_tenant_vip.public_ip_refs.namespace` | [re_acl.selected_tenant_vip.public_ip_refs.namespace](resources--fast_acl--reference--group-001.md#canonical-3301301000130003-2100322102330301-1321032301022322-3133302120023020-0330213012101322-1203111201002332-3312123233010331-2211330203312110) |
| `re_acl.selected_tenant_vip.public_ip_refs.tenant` | [re_acl.selected_tenant_vip.public_ip_refs.tenant](resources--fast_acl--reference--group-001.md#canonical-2022332020333202-2000012312232102-2131112322230232-2322312021201001-3030311212111213-0032313110010131-3020123323231321-0012023312320110) |
| `site_acl` | [site_acl](resources--fast_acl--reference--group-001.md#canonical-0210112023132223-2100131101110110-3332213000200131-2011231123303112-3110103303220200-2121123213022301-0002112303001221-3213202223312213) |
| `site_acl.all_services` | [site_acl.all_services](resources--fast_acl--reference--group-001.md#canonical-2001202110321132-0323203012111101-1223021000320122-0002313213030130-2232321231322220-0021023313222232-0202030202332233-2301330223313113) |
| `site_acl.fast_acl_rules` | [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1213000133313233-3120011203322202-3030010321322320-2103231231130020-1222211223113301-2220031231333320-2312120202103030-0101221321123213) |
| `site_acl.fast_acl_rules.action` | [site_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-3302002002122313-2301103031020011-3200320132101311-3311322130112011-2321022303120000-2313213031222000-2232130330302012-3220102332200313) |
| `site_acl.fast_acl_rules.action.policer_action` | [site_acl.fast_acl_rules.action.policer_action](resources--fast_acl--reference--group-001.md#canonical-3111013313030110-0111031233222000-3330113210231300-3332210132330121-3231231233022013-3020013123103203-1313122013232012-2302031213100110) |
| `site_acl.fast_acl_rules.action.policer_action.ref` | [site_acl.fast_acl_rules.action.policer_action.ref](resources--fast_acl--reference--group-001.md#canonical-1201013223301020-1312101233320332-0331023212112111-2211210112231210-0030322123113311-2311301111331033-3310320113220003-3321012030300020) |
| `site_acl.fast_acl_rules.action.policer_action.ref.kind` | [site_acl.fast_acl_rules.action.policer_action.ref.kind](resources--fast_acl--reference--group-001.md#canonical-0222010221330122-2002222312203000-3032023020230131-1231102123032313-1332001110232022-3233311113320313-1302033013100322-2322223110000200) |
| `site_acl.fast_acl_rules.action.policer_action.ref.name` | [site_acl.fast_acl_rules.action.policer_action.ref.name](resources--fast_acl--reference--group-001.md#canonical-2111110023001022-1010010321011231-0330211202211301-0332101201330203-1103232312010331-0001332333311230-2020112013113332-2113000233113322) |
| `site_acl.fast_acl_rules.action.policer_action.ref.namespace` | [site_acl.fast_acl_rules.action.policer_action.ref.namespace](resources--fast_acl--reference--group-001.md#canonical-3210211210303003-3210113010113122-2200030031230301-3032032213032031-1003100333003133-2122132320320101-1003223220120310-2232322032121131) |
| `site_acl.fast_acl_rules.action.policer_action.ref.tenant` | [site_acl.fast_acl_rules.action.policer_action.ref.tenant](resources--fast_acl--reference--group-001.md#canonical-2210132000102221-2001231122301220-2122013220031120-0101220310130223-3133132312001100-3101223333101011-1223123112021211-2330202013013012) |
| `site_acl.fast_acl_rules.action.policer_action.ref.uid` | [site_acl.fast_acl_rules.action.policer_action.ref.uid](resources--fast_acl--reference--group-001.md#canonical-3100232000311330-1032330330321321-2000332012200102-2231031303001101-0010011331103010-0332123132222223-1123303130100113-1112302010002331) |
| `site_acl.fast_acl_rules.action.protocol_policer_action` | [site_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-0313322010100321-1213120033203313-3200332302223001-3133033102131000-1222020201332000-3111100120020303-0213111131022110-1210312110323312) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref](resources--fast_acl--reference--group-001.md#canonical-0330001223012233-2000201022122200-2231130010330203-1312320003111220-3013013031133101-3203220110331033-2113212332132331-1212323122032023) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.kind` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.kind](resources--fast_acl--reference--group-001.md#canonical-0301101323331131-2033112101223110-1220201002212012-1030312030100010-3233311023131112-1003001121130213-3001200121213102-1100023102203002) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.name` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.name](resources--fast_acl--reference--group-001.md#canonical-1000211132011122-3132211131231032-3203112023002312-0201331020301312-0023331002032301-3131321023003131-0112121002210123-2220322032223221) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace](resources--fast_acl--reference--group-001.md#canonical-3132110133030000-2300312231123233-1111303000200021-2301032222111220-0211122022030022-1001022111203210-1000302230100023-1202223222213003) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant](resources--fast_acl--reference--group-001.md#canonical-1123122123100302-1003222021313003-1030202211121333-1231222322121223-0330031212112112-3331133132031312-0020201112002000-1011011212130123) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.uid` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.uid](resources--fast_acl--reference--group-001.md#canonical-0332313201222132-1231123303101310-1032022221002333-3233202312211002-3222112010302102-1221323201021133-0032030322110120-0231030203133031) |
| `site_acl.fast_acl_rules.action.simple_action` | [site_acl.fast_acl_rules.action.simple_action](resources--fast_acl--reference--group-001.md#canonical-0121302222110021-3233031022213201-0002232121311121-0332300313032133-0223330022223032-0202132312111030-2323123301102223-0130322121211322) |
| `site_acl.fast_acl_rules.ip_prefix_set` | [site_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-3322223021113030-3221231231223211-1001203202213311-2313013022323321-3230012332021120-3030223003020130-0232121013332323-3131000111112300) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref` | [site_acl.fast_acl_rules.ip_prefix_set.ref](resources--fast_acl--reference--group-001.md#canonical-0010103121201203-1033221132212222-3313210132012122-0201111203210232-0223333220123132-0123012122321310-3003220301323101-1322121011113012) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.kind` | [site_acl.fast_acl_rules.ip_prefix_set.ref.kind](resources--fast_acl--reference--group-001.md#canonical-1323233213001133-0232101133032310-1001333230102300-2100130212211020-3321131213100101-1313311123221300-0133202131202102-2020323203012011) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.name` | [site_acl.fast_acl_rules.ip_prefix_set.ref.name](resources--fast_acl--reference--group-001.md#canonical-1332300001323031-1022121320322213-1223313302330202-0012101210330022-0112120023021020-3011311333221300-3203023130002331-0203332031033001) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.namespace` | [site_acl.fast_acl_rules.ip_prefix_set.ref.namespace](resources--fast_acl--reference--group-001.md#canonical-2200031130302301-2133330102120230-0221333300003312-2120200311132202-2120322323012032-1013202323130121-2113232033020232-3131001132032123) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.tenant` | [site_acl.fast_acl_rules.ip_prefix_set.ref.tenant](resources--fast_acl--reference--group-001.md#canonical-1230210312001010-3000032110221032-1103031301202202-1322213322102213-1222130231031011-0333333312113103-2321003101113112-3213233311013030) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.uid` | [site_acl.fast_acl_rules.ip_prefix_set.ref.uid](resources--fast_acl--reference--group-001.md#canonical-2000331131312131-1001113220101312-0312033311220331-3120022332222222-0023020230230123-3231311220031011-1021330212131323-0011323131133230) |
| `site_acl.fast_acl_rules.metadata` | [site_acl.fast_acl_rules.metadata](resources--fast_acl--reference--group-001.md#canonical-2313322301233213-0302103210100232-0111200002311223-3222232011020210-1000321030331301-2000110110122030-0003003100132010-3311202113212203) |
| `site_acl.fast_acl_rules.metadata.description_spec` | [site_acl.fast_acl_rules.metadata.description_spec](resources--fast_acl--reference--group-001.md#canonical-2010001000033113-2010322302111313-1100000022222213-0213320001302122-3030323223012022-3332130100021320-3113311121003120-3222200030201002) |
| `site_acl.fast_acl_rules.metadata.name` | [site_acl.fast_acl_rules.metadata.name](resources--fast_acl--reference--group-001.md#canonical-2213011320312132-3001232021111310-0212310022123300-2332321133101320-2322233111312113-1120022120000132-3332102110120300-2031030320020203) |
| `site_acl.fast_acl_rules.port` | [site_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-1211210303302232-1322333111030122-3121032331200033-3020000003333130-1132213123011322-2000022231131003-1211130321331320-2313221022132001) |
| `site_acl.fast_acl_rules.port.all` | [site_acl.fast_acl_rules.port.all](resources--fast_acl--reference--group-001.md#canonical-3320013122130033-2010013312330311-2213023333210300-0130200212010333-2211311203000222-2310212030223113-0312313220020203-3302020123002220) |
| `site_acl.fast_acl_rules.port.dns` | [site_acl.fast_acl_rules.port.dns](resources--fast_acl--reference--group-001.md#canonical-1200123303201100-0332222121333010-1331130230201230-2221123331312320-1101010021301322-3012233200031203-2132111303010132-2012001001100232) |
| `site_acl.fast_acl_rules.port.user_defined` | [site_acl.fast_acl_rules.port.user_defined](resources--fast_acl--reference--group-001.md#canonical-1033013212131230-1010311212231102-2310001230113123-1011101223022013-1222023112221232-2212231211231011-2211200001120030-3210311001103101) |
| `site_acl.fast_acl_rules.prefix` | [site_acl.fast_acl_rules.prefix](resources--fast_acl--reference--group-001.md#canonical-2023023130320121-0232021103113002-3222220012201110-3231300013211102-1302313100100113-2333300030023132-1002131023110020-0123120332020001) |
| `site_acl.fast_acl_rules.prefix.prefix` | [site_acl.fast_acl_rules.prefix.prefix](resources--fast_acl--reference--group-001.md#canonical-0101112312101310-3222221003300130-1302003223220231-1300010033221303-2113100132312110-0021220303130210-3103101123220301-1131132310031112) |
| `site_acl.inside_network` | [site_acl.inside_network](resources--fast_acl--reference--group-001.md#canonical-1001013123000211-3010211112110201-2111001023021302-1321100212000233-0200031120331212-1231223221000031-0121121333320020-2133312020231321) |
| `site_acl.interface_services` | [site_acl.interface_services](resources--fast_acl--reference--group-001.md#canonical-3322231330113110-2203003221011110-2120332013131202-3121302103222030-0203013001130012-3322120023122022-2221221201122120-1120213010313233) |
| `site_acl.outside_network` | [site_acl.outside_network](resources--fast_acl--reference--group-001.md#canonical-1001130101230122-3111133131222303-2232330233111011-1220030303010101-1030203100323310-0120023202202303-2201322000030322-3002233223123112) |
| `site_acl.vip_services` | [site_acl.vip_services](resources--fast_acl--reference--group-001.md#canonical-1000222221131001-0201030101213220-2213322200112130-2230120001030211-1111132130201300-1222302121021031-1002222123033222-2133013300133021) |
| `timeouts` | [timeouts](resources--fast_acl--reference--group-001.md#canonical-2113103320231323-2313313023211221-3121130102023332-2003111211203121-2220212210302132-2100311120322312-2230102101110220-1202231002210231) |
| `timeouts.create` | [timeouts.create](resources--fast_acl--reference--group-001.md#canonical-3123113132301203-2100023011311002-0021332200320330-1313111333002100-0221123100122212-2223302221333201-3111222112333310-1002121103110221) |
| `timeouts.delete` | [timeouts.delete](resources--fast_acl--reference--group-001.md#canonical-1021112223333130-1130210232233233-2223012030321310-3013111103201312-1331022232132323-2331130002202301-1130302333101331-2332110311130222) |
| `timeouts.read` | [timeouts.read](resources--fast_acl--reference--group-001.md#canonical-3323003310010333-1222202331110112-0212023320123300-0230130320130210-2213203231231132-1011103103130320-1321210301321202-0333310120021300) |
| `timeouts.update` | [timeouts.update](resources--fast_acl--reference--group-001.md#canonical-0112100001222103-2303331111231020-3330110301213321-0310132100010131-3220322012131130-3333200103303233-0322013321001211-2111001033311233) |

<a id="canonical-2002112330130200-1301310100313023-1331330100113213-1020232032003102-1120111301131203-3131131331122331-3110001201323211-3131031122303130"></a>

## Next pages — Property reference / 020322101001 / 12

- [protocol_policer](resources--fast_acl--reference--group-001.md#canonical-3213203231322131-2302120011233022-1101333002320310-0010010133131023-1310101021330233-0023121220213301-1032021020200222-1000230202033211)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [timeouts](resources--fast_acl--reference--group-001.md#canonical-2222200211213322-0210312320233133-3101222313001002-3200023112003003-2231122211030313-0030223121111320-3321112132323230-1211100021300311)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-3213203231322131-2302120011233022-1101333002320310-0010010133131023-1310101021330233-0023121220213301-1032021020200222-1000230202033211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121003133103121-0330320020222012-2123300131001210-0231000300302210-2130020003313110-3101303021201203-1223013112031001-0111313313213100"></a>

## protocol_policer — protocol_policer / 312110002230 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- protocol_policer

<a id="canonical-3102103313301300-2231233023321201-0113333233012333-2300001103232101-0003223113230333-2312121230013321-0003103221000030-3210320201230331"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
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
protocol_policer {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232333030130121-2032232001022323-3213131021001322-3103211132301110-2012031010303022-1132313103033113-1001322211220210-0102122030032222"></a>

## Direct properties — protocol_policer / 312110002230 / 3

<a id="canonical-1231302002211331-3110223031112310-2003231201012233-0122222001221232-3102231113120003-0211111123033233-2002132102112232-0212313100133100"></a>

<a id="canonical-2202122232233023-2113021211021013-3313022201320303-1230102102300023-3110300121330002-3032102230012311-1003110010100020-3031010323102111"></a>

## name property — protocol_policer / 312110002230 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-1012111003333223-2201311203302303-1000333300232030-0212023021122031-1321000300021033-3023113003030300-1103312010103302-2322010123332003"></a>

<a id="canonical-3030313220023232-2003320002233010-3311212001131321-2102102213112120-2203311303223200-3122310030000232-1321231100202120-2320113233101131"></a>

## namespace property — protocol_policer / 312110002230 / 5

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
}
```

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

<a id="canonical-3110322121023032-0123303332020013-2013212112120120-2111303001230123-3133132101112333-1211220222102122-2110331102023232-2323220301313201"></a>

<a id="canonical-3302332220030323-1300003020123203-0003001011012002-3002102103303111-3213313321202001-0322333312111022-2203211010200330-3201311231301121"></a>

## tenant property — protocol_policer / 312110002230 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-1223201001101322-3121123332030003-1112223111030220-2020130233031103-3323333232002031-2103003120121221-1222332213300231-3303011203201320"></a>

## Next pages — protocol_policer / 312110002230 / 7

- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122203123100022-0023101331212021-2303200220203122-2310310013331320-3330102311132002-0201213031113332-2111330121001303-2011310001231032"></a>

## re_acl — re_acl / 202332212130 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- re_acl

<a id="canonical-2031331223221133-0333222313002321-3131001313230032-1033233112222331-1323223232332312-3313200032102321-3200103002320021-1213213032213311"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: re\_acl, site\_acl\] Fast ACL for RE. Fast ACL definition for RE.

Upstream description:

Fast ACL definition for RE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_public_vips",
    "default_tenant_vip"),
  validators.ConflictingObjectAttributes("all_public_vips",
    "selected_tenant_vip"),
  validators.ConflictingObjectAttributes("default_tenant_vip",
    "selected_tenant_vip")}
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
  "x-ves-oneof-field-vip_choice": "[\"all_public_vips\",\"default_tenant_vip\",\"selected_tenant_vip\"]"
}
```

OneOf alternatives in this subsection:

- [re_acl](resources--fast_acl--reference--group-001.md#canonical-2031331223221133-0333222313002321-3131001313230032-1033233112222331-1323223232332312-3313200032102321-3200103002320021-1213213032213311)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0210112023132223-2100131101110110-3332213000200131-2011231123303112-3110103303220200-2121123213022301-0002112303001221-3213202223312213)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
re_acl {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012310331203120-1123233101210321-1310332233002310-0331213021222101-2131022030001131-1330000120311333-3021003222001102-2300013130311303"></a>

## Direct properties — re_acl / 202332212130 / 3

- [all_public_vips](resources--fast_acl--reference--group-001.md#canonical-2000312210120233-1310230112223212-1030202203110302-3031330303231331-3321133323031310-1322201012310302-3030022202220232-1313132121221211): complete subsection reference.

- [default_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-1012033020221202-2222110213033032-3311323232211331-0212303300232303-2011322233023322-0231130110100223-1333133302302011-1020022033322122): complete subsection reference.

- [fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010): complete subsection reference.

- [selected_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-2313112310232031-0132231301303321-1222013103002123-1203001002202003-2121231321230200-2210112100123321-0011302310000002-2202122102102002): complete subsection reference.

<a id="canonical-2222003213031312-1332000323233033-0131111023232123-1011323013320111-0003131233233000-0233322310311203-1103202211323320-0310322333233000"></a>

## Next pages — re_acl / 202332212130 / 4

- [re_acl.all_public_vips](resources--fast_acl--reference--group-001.md#canonical-2000312210120233-1310230112223212-1030202203110302-3031330303231331-3321133323031310-1322201012310302-3030022202220232-1313132121221211)
- [re_acl.default_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-1012033020221202-2222110213033032-3311323232211331-0212303300232303-2011322233023322-0231130110100223-1333133302302011-1020022033322122)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- [re_acl.selected_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-2313112310232031-0132231301303321-1222013103002123-1203001002202003-2121231321230200-2210112100123321-0011302310000002-2202122102102002)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-2000312210120233-1310230112223212-1030202203110302-3031330303231331-3321133323031310-1322201012310302-3030022202220232-1313132121221211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212332213000310-2313120031332332-1233010332302333-1011312130203332-0022211010201313-2303202002012310-3323200103210230-2200020332303021"></a>

## re_acl.all_public_vips — all_public_vips / 313031313022 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- re_acl.all_public_vips

<a id="canonical-1123232130300230-3222212333303332-3312301131002112-3312133033012133-3011222332311033-0203310222123233-0112010312311310-3121200101221100"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
all_public_vips = {}
```

<a id="canonical-1333321322010300-0313302010330210-0311131310120021-1002023230033301-1003220021111122-2212123020333011-0323301123010232-0301310101101232"></a>

## Direct properties — all_public_vips / 313031313022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333122222022313-3023130213112313-1210131333010102-1330131223113020-0223103023121132-3331033323120023-0112321301032323-2133222330112303"></a>

## Next pages — all_public_vips / 313031313022 / 4

- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-1012033020221202-2222110213033032-3311323232211331-0212303300232303-2011322233023322-0231130110100223-1333133302302011-1020022033322122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213130002030300-3302303333003321-2022112100110323-3201311122112031-1033331231311333-3003031120121201-3231023303122032-0312313210013100"></a>

## re_acl.default_tenant_vip — default_tenant_vip / 012102232331 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- re_acl.default_tenant_vip

<a id="canonical-0102101121211100-2020132333233111-2303310112003122-0031210132022231-0023210020112100-1232313210012011-1311223213230222-1311112131230201"></a>

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
default_tenant_vip = {}
```

<a id="canonical-1320031300203020-3100303121133001-1000301202013203-2002232231302312-3233320213131123-0023203020220120-1321110131021213-2032033312210021"></a>

## Direct properties — default_tenant_vip / 012102232331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301032310120203-1232322130223323-1132100320102230-2332121031331323-0003222100033220-1133021200311113-3030122303333302-1131103002312030"></a>

## Next pages — default_tenant_vip / 012102232331 / 4

- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023322312110102-3123303130210012-0010231202120133-2021003312303301-2110000202213013-0321230122210131-0000233103311010-2301010020223211"></a>

## re_acl.fast_acl_rules — fast_acl_rules / 120033333112 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- re_acl.fast_acl_rules

<a id="canonical-1110303332131332-3201323100322330-3320332003010313-2332003330311003-3313221232220322-1102031232130311-0332323211032322-0230122010223113"></a>

Type: `"object"`. list nested block, Optional.

Rules. Fast ACL rules to match. Defaults to \`\[\]\`. Server applies default when omitted.

Upstream description:

Fast ACL rules to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix")}
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
fast_acl_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1113221202213033-2111222102211202-1010222120100220-2133201132132222-2001303322331223-2110110033203321-3302023220330221-0030030012230230"></a>

## Direct properties — fast_acl_rules / 120033333112 / 3

- [action](resources--fast_acl--reference--group-001.md#canonical-3122332100322313-2031131132103113-3311023121020300-2310111001133321-0120011322032101-3230323102111030-0111011300102110-0311030112121100): complete subsection reference.

- [ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-1311020211222203-0131210113113321-1223212120122230-0001123311013200-0102302122213323-0133232111010113-1301220312311300-3210332033012101): complete subsection reference.

- [metadata](resources--fast_acl--reference--group-001.md#canonical-1003102013211211-0022312332203312-3322132221023230-3231230011221311-0311032111302220-0000200231022301-3100103010001033-2001300202332000): complete subsection reference.

- [port](resources--fast_acl--reference--group-001.md#canonical-0100122032000213-2012010313310033-1011023301101021-2123211122003310-0322300001013201-2323000322233123-1313222213331131-2232231202013120): complete subsection reference.

- [prefix](resources--fast_acl--reference--group-001.md#canonical-0313320122213021-1012310012121113-3302332132321101-1331303011330201-2020030010012301-0203310111202103-0313123301020202-3132213032111200): complete subsection reference.

<a id="canonical-1120112121132321-3211312222022132-0332220030303020-0123121032301312-2201130122330213-1000033131012333-0332332212000001-1201121332113111"></a>

## Next pages — fast_acl_rules / 120033333112 / 4

- [re_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-3122332100322313-2031131132103113-3311023121020300-2310111001133321-0120011322032101-3230323102111030-0111011300102110-0311030112121100)
- [re_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-1311020211222203-0131210113113321-1223212120122230-0001123311013200-0102302122213323-0133232111010113-1301220312311300-3210332033012101)
- [re_acl.fast_acl_rules.metadata](resources--fast_acl--reference--group-001.md#canonical-1003102013211211-0022312332203312-3322132221023230-3231230011221311-0311032111302220-0000200231022301-3100103010001033-2001300202332000)
- [re_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-0100122032000213-2012010313310033-1011023301101021-2123211122003310-0322300001013201-2323000322233123-1313222213331131-2232231202013120)
- [re_acl.fast_acl_rules.prefix](resources--fast_acl--reference--group-001.md#canonical-0313320122213021-1012310012121113-3302332132321101-1331303011330201-2020030010012301-0203310111202103-0313123301020202-3132213032111200)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-3122332100322313-2031131132103113-3311023121020300-2310111001133321-0120011322032101-3230323102111030-0111011300102110-0311030112121100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121033220022331-1020313210132321-3123020130323132-0101331233331332-1032031113333212-2223130113122121-2000111121122231-2032120311300201"></a>

## re_acl.fast_acl_rules.action — action / 203323113021 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- re_acl.fast_acl_rules.action

<a id="canonical-3100003210032133-1300133020003332-0022121332120103-3233013230103232-2212322111212030-1123301021202133-1202132322321030-2322032000320200"></a>

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

<a id="canonical-2030211320022310-0122311011222310-2332000311032321-0222201233201200-0212222133222223-0113203301230013-1003001121203130-1202110123100123"></a>

## Direct properties — action / 203323113021 / 3

- [policer_action](resources--fast_acl--reference--group-001.md#canonical-2323002023331031-1212132201020130-1133223202221102-2301310212213112-2030231023120022-0132213330112313-0001020121221131-2002310313033032): complete subsection reference.

- [protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-3123221112132311-2320320111000120-2121123230330033-1033133300331230-0331301100211122-3013111132013123-3111021213033122-1330031322112302): complete subsection reference.

<a id="canonical-1001120102321220-1011322121321310-2111200113012021-3103221202030201-0032232332123003-3012232310032122-0333133013100223-3230310130021100"></a>

<a id="canonical-2211122022222223-2320000210300021-1102303303022312-1220121011302301-0132311031113302-3302013212102012-3210032131001111-0202233332110312"></a>

## simple_action property — action / 203323113021 / 4

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

<a id="canonical-0313231112022003-2203133333330103-2003010110111000-3232102012201323-0023011130323010-2221321321211000-2303311130201112-2120232012122033"></a>

## Next pages — action / 203323113021 / 5

- [re_acl.fast_acl_rules.action.policer_action](resources--fast_acl--reference--group-001.md#canonical-2323002023331031-1212132201020130-1133223202221102-2301310212213112-2030231023120022-0132213330112313-0001020121221131-2002310313033032)
- [re_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-3123221112132311-2320320111000120-2121123230330033-1033133300331230-0331301100211122-3013111132013123-3111021213033122-1330031322112302)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-2323002023331031-1212132201020130-1133223202221102-2301310212213112-2030231023120022-0132213330112313-0001020121221131-2002310313033032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201331211213230-0303030301022221-0330111210333321-0301002030031130-2332332200200212-0031021203112121-3232330113102311-1203231232000131"></a>

## re_acl.fast_acl_rules.action.policer_action — policer_action / 131203132120 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- [re_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-3122332100322313-2031131132103113-3311023121020300-2310111001133321-0120011322032101-3230323102111030-0111011300102110-0311030112121100)
- re_acl.fast_acl_rules.action.policer_action

<a id="canonical-2221321200021220-1021231010323222-0003223032203322-1301103222010312-0323000011012021-2222223120033131-2323321000032131-1331020003220232"></a>

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

<a id="canonical-3112100011322312-3130231302320110-2302112132332232-3221003200133230-1202320201121113-2031022013000201-2000303211312310-0013023113033312"></a>

## Direct properties — policer_action / 131203132120 / 3

- [ref](resources--fast_acl--reference--group-001.md#canonical-0021013201130213-0331013020032301-3020232110332313-0021011012033300-2213301233010011-3201312333110320-2002212023301301-0013210033103322): complete subsection reference.

<a id="canonical-0301233302310010-2231113032311323-2233223030013110-1200222202113222-3201022320101032-0311301011200332-1031203233101002-1323100300302233"></a>

## Next pages — policer_action / 131203132120 / 4

- [re_acl.fast_acl_rules.action.policer_action.ref](resources--fast_acl--reference--group-001.md#canonical-0021013201130213-0331013020032301-3020232110332313-0021011012033300-2213301233010011-3201312333110320-2002212023301301-0013210033103322)
- [re_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-3122332100322313-2031131132103113-3311023121020300-2310111001133321-0120011322032101-3230323102111030-0111011300102110-0311030112121100)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-0021013201130213-0331013020032301-3020232110332313-0021011012033300-2213301233010011-3201312333110320-2002212023301301-0013210033103322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133032020210133-2332023011101020-3312201113033333-2110012311331312-3033300012010102-0220300110010012-2023203123310322-1331131030201132"></a>

## re_acl.fast_acl_rules.action.policer_action.ref — ref / 320313103011 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- [re_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-3122332100322313-2031131132103113-3311023121020300-2310111001133321-0120011322032101-3230323102111030-0111011300102110-0311030112121100)
- [re_acl.fast_acl_rules.action.policer_action](resources--fast_acl--reference--group-001.md#canonical-2323002023331031-1212132201020130-1133223202221102-2301310212213112-2030231023120022-0132213330112313-0001020121221131-2002310313033032)
- re_acl.fast_acl_rules.action.policer_action.ref

<a id="canonical-3303301232300322-0201201333333222-3332111123121223-1311201100222110-1330311333201000-1203211322003001-0100113022120332-1233110002311323"></a>

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

<a id="canonical-1300333313303323-3201123100301330-3101223300103132-1130323021331202-3122113131113313-2000013000211032-0203013000300203-3100231300002001"></a>

## Direct properties — ref / 320313103011 / 3

<a id="canonical-2030102221303120-1330302012323120-0013121121202231-0223203203300230-3023330313110000-2202032131133102-2033112111323322-0210223230133333"></a>

<a id="canonical-0310222320331220-2001222231012110-1122023210003002-2201232213033010-0020021100002103-3003112012022220-3211333131102131-3332323123101121"></a>

## kind property — ref / 320313103011 / 4

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

<a id="canonical-2131201312023011-3020110200120121-1123133003112312-2222013230222021-2031222003313121-1220003333310023-0321202102220132-1212310132102303"></a>

<a id="canonical-3110201100131011-0210122223201322-3000311100320103-2103210233323222-3133113032322321-0121122020031013-2002123203100331-3303023112300112"></a>

## name property — ref / 320313103011 / 5

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

<a id="canonical-0212033232012031-0132220230222233-3010321012021031-2331233313130333-0031000301013121-0311111133133233-2102011031300002-0303013100100200"></a>

<a id="canonical-2130111220102312-2133133010313232-3123303130123222-0133303321200030-1300132223131301-1330132203033333-2310300213321113-0330303302122000"></a>

## namespace property — ref / 320313103011 / 6

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

<a id="canonical-2001232321123122-1210101010132331-3120011303302333-3012300203020213-0311030232321213-1100132321331013-3021000222010331-3200122110020132"></a>

<a id="canonical-2110302333020201-3311013003321303-1122032131231133-2112332203120003-1330213033321222-0221102311320232-0320223022012233-2022031001123331"></a>

## tenant property — ref / 320313103011 / 7

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

<a id="canonical-0330020210222020-3210300232202230-3301201031313123-1000033012111230-2011202300202220-3311220023231031-3212013103310222-3223200323111203"></a>

<a id="canonical-1130321310000013-0231001122221103-3322330122033320-2213213032232000-3230120010320131-0210221102303011-0233011100333033-3220330232321213"></a>

## uid property — ref / 320313103011 / 8

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

<a id="canonical-3200032322131113-3210002133101120-2101103321013312-0030112131133002-3131311121011230-2102102200111302-3233232301303312-2213202130313213"></a>

## Next pages — ref / 320313103011 / 9

- [re_acl.fast_acl_rules.action.policer_action](resources--fast_acl--reference--group-001.md#canonical-2323002023331031-1212132201020130-1133223202221102-2301310212213112-2030231023120022-0132213330112313-0001020121221131-2002310313033032)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-3123221112132311-2320320111000120-2121123230330033-1033133300331230-0331301100211122-3013111132013123-3111021213033122-1330031322112302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113211232212322-2222010300003133-1203330233002320-3332302231311100-3200323103000030-0013311122013033-0333231300213323-2333122000333012"></a>

## re_acl.fast_acl_rules.action.protocol_policer_action — protocol_policer_action / 110033202223 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- [re_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-3122332100322313-2031131132103113-3311023121020300-2310111001133321-0120011322032101-3230323102111030-0111011300102110-0311030112121100)
- re_acl.fast_acl_rules.action.protocol_policer_action

<a id="canonical-2133210311112322-2222201220111231-0222111210202131-0101312010212230-0203131023311230-3110320011320201-1003233223221013-0001123133101010"></a>

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

<a id="canonical-0021031203031112-2210010011121022-0120101021020210-2201210230101210-3130302311031300-3012000103000122-3020133110313332-1323123311100122"></a>

## Direct properties — protocol_policer_action / 110033202223 / 3

- [ref](resources--fast_acl--reference--group-001.md#canonical-0222112011031021-2000230021210100-2023000033101130-2232211000021303-3322230110001131-0001032311003122-1002310122121322-2300333012220032): complete subsection reference.

<a id="canonical-0331321003021210-1012230111133112-0312201022212012-2102310213221030-3010302003320102-3200300032333200-0102211221111200-3303110020222130"></a>

## Next pages — protocol_policer_action / 110033202223 / 4

- [re_acl.fast_acl_rules.action.protocol_policer_action.ref](resources--fast_acl--reference--group-001.md#canonical-0222112011031021-2000230021210100-2023000033101130-2232211000021303-3322230110001131-0001032311003122-1002310122121322-2300333012220032)
- [re_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-3122332100322313-2031131132103113-3311023121020300-2310111001133321-0120011322032101-3230323102111030-0111011300102110-0311030112121100)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-0222112011031021-2000230021210100-2023000033101130-2232211000021303-3322230110001131-0001032311003122-1002310122121322-2300333012220032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120203032330211-3220011301002030-0201121003132003-2100331030210233-0331012103003311-3122213203302001-2012132313230220-0100223121232230"></a>

## re_acl.fast_acl_rules.action.protocol_policer_action.ref — ref / 120133033203 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- [re_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-3122332100322313-2031131132103113-3311023121020300-2310111001133321-0120011322032101-3230323102111030-0111011300102110-0311030112121100)
- [re_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-3123221112132311-2320320111000120-2121123230330033-1033133300331230-0331301100211122-3013111132013123-3111021213033122-1330031322112302)
- re_acl.fast_acl_rules.action.protocol_policer_action.ref

<a id="canonical-1313330221332023-1321131011332000-3122301210211113-1032322032310213-0202032132013312-2010323120133031-2032233013202223-2311313013213033"></a>

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

<a id="canonical-1010122222110212-1220210203211110-0311120003312101-2130310330333003-2103233330232202-1020022200123032-2100230333102110-1303331311000232"></a>

## Direct properties — ref / 120133033203 / 3

<a id="canonical-0300310132113201-3212321121021121-1010201132020212-0120013301011321-3000010331321312-1310012110211102-1310301211212023-3101131213102020"></a>

<a id="canonical-0120031322113223-0103310103212333-1221131322130301-3203131332311301-0103030132321331-1112330132210033-1121300230023302-1101010011003222"></a>

## kind property — ref / 120133033203 / 4

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

<a id="canonical-3233302112110031-3023130001003022-0332022330113212-1323030010113200-0230030023202211-3300011100323102-2223210020121202-0101301002000322"></a>

<a id="canonical-0123132110002003-3123002331311232-3121311003332311-3021020232311122-1311300121222121-0201011103300211-0100233212113311-1101100201121002"></a>

## name property — ref / 120133033203 / 5

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

<a id="canonical-0322020220223110-3200311120222213-0113022320112210-0110111331213001-2121031122103211-3013123212312113-3321301232030313-3330320030301211"></a>

<a id="canonical-2321300122021331-0001100022000311-0310220100123203-2003033223122021-3221030200330212-1323202000210212-3321002113111000-3000110223220312"></a>

## namespace property — ref / 120133033203 / 6

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

<a id="canonical-3001023011110333-0200122312322010-2101203011312121-3130120322122231-0210002120330300-0301320300313001-3001102312020020-1311100310321200"></a>

<a id="canonical-2111303301020103-2321132122300303-2002233133102233-3120133031103110-3203322020310332-3021100022303313-3231000033132032-2330331323312030"></a>

## tenant property — ref / 120133033203 / 7

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

<a id="canonical-0310003232210223-1022322020312002-2113220122310112-3311012231122302-3003203012333002-2113023030331220-2301203011310200-3133220210122103"></a>

<a id="canonical-1310113121031312-1130203211230320-1001211221203003-2311123022333222-1002223021012111-0002220331211223-1211310330030212-3202213130010200"></a>

## uid property — ref / 120133033203 / 8

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

<a id="canonical-0300312003323121-0000222012223332-3223332322030311-0113103233331103-1031112200122333-1022120113123002-1120302332011331-3320122313320021"></a>

## Next pages — ref / 120133033203 / 9

- [re_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-3123221112132311-2320320111000120-2121123230330033-1033133300331230-0331301100211122-3013111132013123-3111021213033122-1330031322112302)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-1311020211222203-0131210113113321-1223212120122230-0001123311013200-0102302122213323-0133232111010113-1301220312311300-3210332033012101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221302130223102-2132231202023202-0303223033302200-2103232001012102-3100313210322310-1011021113001013-0222001332123000-1233201000331000"></a>

## re_acl.fast_acl_rules.ip_prefix_set — ip_prefix_set / 320013211131 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- re_acl.fast_acl_rules.ip_prefix_set

<a id="canonical-2301212323203202-0133322022203122-1020010002002110-0311011201131113-1021222003022201-2233122201303010-3021132200033332-1130121312023311"></a>

Type: `"object"`. single nested block, Optional.

List of references to ip\_prefix\_set objects.

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

Terraform syntax:

```terraform
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233013123033232-1020223222010323-0331013310021332-1001030311113010-1322201301331032-3202002030332201-2312033003203133-1201133130120021"></a>

## Direct properties — ip_prefix_set / 320013211131 / 3

- [ref](resources--fast_acl--reference--group-001.md#canonical-0312223311122231-0123110121331010-0102233012100102-0311220132021303-0113212122130101-2103110231132300-1200021102033322-1131102203313001): complete subsection reference.

<a id="canonical-3222102232323001-2202330300012113-0212231303120012-3113231003100031-3002101000100000-0000113033233100-0120202303303133-2100003112211032"></a>

## Next pages — ip_prefix_set / 320013211131 / 4

- [re_acl.fast_acl_rules.ip_prefix_set.ref](resources--fast_acl--reference--group-001.md#canonical-0312223311122231-0123110121331010-0102233012100102-0311220132021303-0113212122130101-2103110231132300-1200021102033322-1131102203313001)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-0312223311122231-0123110121331010-0102233012100102-0311220132021303-0113212122130101-2103110231132300-1200021102033322-1131102203313001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112012022112032-0111302210000213-3312003033132223-1302230012021103-0311302210101103-1210233323200231-1033111322331200-1232010202231333"></a>

## re_acl.fast_acl_rules.ip_prefix_set.ref — ref / 112000132013 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- [re_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-1311020211222203-0131210113113321-1223212120122230-0001123311013200-0102302122213323-0133232111010113-1301220312311300-3210332033012101)
- re_acl.fast_acl_rules.ip_prefix_set.ref

<a id="canonical-0131113102320202-1313021131330100-0130312331220113-0303132111101031-1102213231030312-2132202213010112-1011203303313201-1132213011010022"></a>

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

<a id="canonical-1021113301203013-3031322033012300-2220103301030020-1203000312100330-0031111033320211-1113020212213110-3323320003031222-0032033101321222"></a>

## Direct properties — ref / 112000132013 / 3

<a id="canonical-3111300002213303-2301331033022321-3210210223201200-1000221033001320-3200022023132210-0000301310212103-3100132030013313-2123232321133022"></a>

<a id="canonical-1021011200310000-1322320022132030-1131110202021211-1100333212201023-3230311202331010-2021030031001230-3033203011133100-0110020320322333"></a>

## kind property — ref / 112000132013 / 4

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

<a id="canonical-0010013321332001-2300112233102110-1030130103122223-0302123231301113-0212221032323111-1212300202311331-2113312220112133-0030313012010023"></a>

<a id="canonical-3002013212313312-0011112320202321-1032300303210200-3131213010203011-3200132301012333-0113230320021220-1101300032322013-3322213221101210"></a>

## name property — ref / 112000132013 / 5

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

<a id="canonical-1120101320310120-0031130030313101-1011221222000030-2202233231110022-2010130223230120-2131010331223032-0212231323331223-2203210110213022"></a>

<a id="canonical-0333101321023003-3320112201323203-0201123212002112-2302330101032213-3013102211132311-0330311330222121-3033102003130223-0211233033331023"></a>

## namespace property — ref / 112000132013 / 6

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

<a id="canonical-1213320112131032-2011101221133011-0320012230333203-2303130011302001-2103200232111101-0323201002001222-1310101302213302-1123023203021203"></a>

<a id="canonical-2101120030030212-2011322111120221-1011332100200132-2330213011312203-2223330201031323-3030121020321100-2002122310013003-2233312311203331"></a>

## tenant property — ref / 112000132013 / 7

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

<a id="canonical-0320310111322103-3221320030213012-3333322033312110-2123001210200300-3031213122132030-1233333132130003-2121122300222022-0301020102320031"></a>

<a id="canonical-2231200112121220-1000213320303003-1121023203231210-1133203033200033-2201013032003213-1102332013331233-0002232102203002-2103112212030103"></a>

## uid property — ref / 112000132013 / 8

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

<a id="canonical-3312323100120333-2133011331321112-1120120232330311-2220212230123103-2212313303021313-3220233023111112-2233013232311232-0011312002221131"></a>

## Next pages — ref / 112000132013 / 9

- [re_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-1311020211222203-0131210113113321-1223212120122230-0001123311013200-0102302122213323-0133232111010113-1301220312311300-3210332033012101)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-1003102013211211-0022312332203312-3322132221023230-3231230011221311-0311032111302220-0000200231022301-3100103010001033-2001300202332000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111200332322210-2322032322310233-3132210021030202-3321132232233102-0032320201122313-3123303101102221-3131212203231201-2232002032122110"></a>

## re_acl.fast_acl_rules.metadata — metadata / 132031321132 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- re_acl.fast_acl_rules.metadata

<a id="canonical-2033031123311032-2220212123122120-2100323130003020-2303202202300111-1201130132120023-2131220031311320-2312023220103312-3010312111102020"></a>

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

<a id="canonical-2013230310132102-2012233001220320-0203002300321322-2331201220200103-2021312322302132-2031231222212231-1302323012213213-0230110120112031"></a>

## Direct properties — metadata / 132031321132 / 3

<a id="canonical-2322310021321012-3312021012312223-1032321001322301-2033231010213000-3003122021020300-0313233000223232-0030101023323033-0233033111333112"></a>

<a id="canonical-3002333021131122-1330023310120203-2220132111222311-2010020232320003-2102332112003013-2220123313110133-3212021221111120-0132213103131111"></a>

## description_spec property — metadata / 132031321132 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1002131002303103-0302120012022022-0321132232110031-1202020200211313-1021100223200221-3202100011323232-2303111330033220-0230331121021003"></a>

<a id="canonical-1323112133003001-3333230220102112-1132121011323002-2231201122033002-3030223030331011-2310033303302111-3320213231033222-1212202320311233"></a>

## name property — metadata / 132031321132 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3123102333020102-1311102322010111-1211333321100030-2330202133132111-0101123221020201-1323123231223110-3110230312323322-0213200031320300"></a>

## Next pages — metadata / 132031321132 / 6

- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-0100122032000213-2012010313310033-1011023301101021-2123211122003310-0322300001013201-2323000322233123-1313222213331131-2232231202013120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110320032200112-1102000022101323-2031202310020131-2333132201120021-2232300221202001-3230310000002112-0210332201300221-3110113022202333"></a>

## re_acl.fast_acl_rules.port — port / 201113232321 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- re_acl.fast_acl_rules.port

<a id="canonical-3022310110023111-3200133200322213-2232223122230201-0330002020330223-3023132212322310-1032321110131212-2203013300012111-2313231110323002"></a>

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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
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

<a id="canonical-3223121303010103-3320030330113321-0121012102320032-1231233112022201-0113211211231000-2133103312001221-3311021311013223-2003023323200102"></a>

## Direct properties — port / 201113232321 / 3

- [all](resources--fast_acl--reference--group-001.md#canonical-2022232312110011-1211111213012111-2330013133321113-3111030330322333-0120110121121021-0002010233112330-0212120013220113-0123102220030221): complete subsection reference.

- [DNS](resources--fast_acl--reference--group-001.md#canonical-1130232133313021-3101221310020111-3301210100122123-2301132121320001-3101001001313331-2001221031221033-3212031100102121-0322221110201200): complete subsection reference.

<a id="canonical-3232212313000203-3223032223313312-0212202211122222-3322020200020010-1331001132231311-3012022033303033-0021211022331333-0121003302110130"></a>

<a id="canonical-0322233321222113-3100002001010303-0331330030211101-2212023302230230-3032021300220102-1012222020131203-2230030012212111-0130010120303020"></a>

## user_defined property — port / 201113232321 / 4

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

<a id="canonical-0101330213301102-1020311130211131-3021310230110033-0200022122030211-1222003203122111-3230233100230213-3123012130030020-1103332221232222"></a>

## Next pages — port / 201113232321 / 5

- [re_acl.fast_acl_rules.port.all](resources--fast_acl--reference--group-001.md#canonical-2022232312110011-1211111213012111-2330013133321113-3111030330322333-0120110121121021-0002010233112330-0212120013220113-0123102220030221)
- [re_acl.fast_acl_rules.port.dns](resources--fast_acl--reference--group-001.md#canonical-1130232133313021-3101221310020111-3301210100122123-2301132121320001-3101001001313331-2001221031221033-3212031100102121-0322221110201200)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-2022232312110011-1211111213012111-2330013133321113-3111030330322333-0120110121121021-0002010233112330-0212120013220113-0123102220030221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100020022230030-1312303121330303-3203213222121102-2013310021001113-0200033200101201-1230131020203001-3202302113322313-1223333013303102"></a>

## re_acl.fast_acl_rules.port.all — all / 231002230100 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- [re_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-0100122032000213-2012010313310033-1011023301101021-2123211122003310-0322300001013201-2323000322233123-1313222213331131-2232231202013120)
- re_acl.fast_acl_rules.port.all

<a id="canonical-2131230023003202-0130322001023200-0021100003102210-1201322101203202-2130231121223212-2310031132300322-3133110320222102-1222220023233013"></a>

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

<a id="canonical-2320022330103330-1110132223132013-1122110012023121-2320033030031101-1130130323100002-0031101111223302-1200132013330103-1310003230233033"></a>

## Direct properties — all / 231002230100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221211312302010-2013201120023210-0031003012232230-1231031201132102-2332101021300221-1130303200122303-2011000320331203-2132020223320011"></a>

## Next pages — all / 231002230100 / 4

- [re_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-0100122032000213-2012010313310033-1011023301101021-2123211122003310-0322300001013201-2323000322233123-1313222213331131-2232231202013120)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-1130232133313021-3101221310020111-3301210100122123-2301132121320001-3101001001313331-2001221031221033-3212031100102121-0322221110201200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100102301233232-0311030303102003-1003101221221013-3110003333301013-3202112011210211-2330310200202322-0110312102132311-0001110033222113"></a>

## re_acl.fast_acl_rules.port.DNS — DNS / 000100001333 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- [re_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-0100122032000213-2012010313310033-1011023301101021-2123211122003310-0322300001013201-2323000322233123-1313222213331131-2232231202013120)
- re_acl.fast_acl_rules.port.DNS

<a id="canonical-2032313011223231-1223223312333021-3310033132010001-2233300230121322-2313130010033300-1012021233111330-0113011132313113-2121223221331213"></a>

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

<a id="canonical-3121002213303320-0303203130301133-1303130302320323-1211312301202010-1320220002232033-3301231021010301-1020323223030022-1001333113201202"></a>

## Direct properties — DNS / 000100001333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211323233023132-1211032321001313-2020310302120012-0220122110102301-0332331231220013-0331033223322211-3012112323331202-2011012101301000"></a>

## Next pages — DNS / 000100001333 / 4

- [re_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-0100122032000213-2012010313310033-1011023301101021-2123211122003310-0322300001013201-2323000322233123-1313222213331131-2232231202013120)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-0313320122213021-1012310012121113-3302332132321101-1331303011330201-2020030010012301-0203310111202103-0313123301020202-3132213032111200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113212232201102-2101022212113312-0003313222130223-1323003333333023-3011011010300022-2021333001232021-1111000201021131-3113313013113021"></a>

## re_acl.fast_acl_rules.prefix — prefix / 233311123131 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- re_acl.fast_acl_rules.prefix

<a id="canonical-1012333012203023-1201331111213100-0131110023002222-0100223113221120-3013221023310303-0012131331222023-2001300131312310-3101003023222231"></a>

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

<a id="canonical-1330101110320100-1202322113012103-2300321013032302-1112201201200202-0331022203310322-0222202300031013-3320232023113323-2011312130110223"></a>

## Direct properties — prefix / 233311123131 / 3

<a id="canonical-2303131030001221-3111320203111332-0220033222312301-3313132230101120-1100330112331031-2100021113010212-1000212031131132-3212013211310120"></a>

<a id="canonical-1102131320022333-1220213000320032-2020112303120112-3201031032220111-0310223011302200-3233200202132320-2113322210030213-2012100021133003"></a>

## prefix property — prefix / 233311123131 / 4

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

<a id="canonical-1231031122301132-3233010233000123-1011302132033202-1130233030230023-3320033302111132-1020103320102311-0130112020011333-0212033303023133"></a>

## Next pages — prefix / 233311123131 / 5

- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-2313112310232031-0132231301303321-1222013103002123-1203001002202003-2121231321230200-2210112100123321-0011302310000002-2202122102102002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222303330330333-2213001011130003-1110102220120003-1211320302132013-2013300331110003-3330302022213230-2330201212003332-1320100011313010"></a>

## re_acl.selected_tenant_vip — selected_tenant_vip / 320232123013 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- re_acl.selected_tenant_vip

<a id="canonical-0033230212102001-0013122110231101-2021311311032200-3212000132322202-0120230311112211-3330132001223300-1302020003311012-1211110232122210"></a>

Type: `"object"`. single nested block, Optional.

Specific Tenant VIP. Select various tenant public VIP(s)

Upstream description:

Select various tenant public VIP(s)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("public_ip_refs")}
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
selected_tenant_vip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2201201130100211-2030312221230303-1013330323330311-2023221122230322-3133013222130001-1001011313033000-1110132133300300-0123312003031200"></a>

## Direct properties — selected_tenant_vip / 320232123013 / 3

<a id="canonical-1131333232223032-0011332132331201-0112001000011101-0123301213200303-1023211021120133-2320122101012033-2123111333302203-2221312212022223"></a>

<a id="canonical-1221312113203010-2202310102332103-1322323001113102-3013222012332330-3012303123023321-1303203300021222-0330201203011002-1320221133230002"></a>

## default_tenant_vip property — selected_tenant_vip / 320232123013 / 4

Type: `"bool"`. Optional.

Include tenant VIP in list of specific VIP(s).

Upstream description:

Include tenant VIP in list of specific VIP(s)

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

- [public_ip_refs](resources--fast_acl--reference--group-001.md#canonical-2100303222200103-2310111221033132-3220031130200001-1021132210220001-3013101002232022-3323303001031200-1201110302333332-0312100000011003): complete subsection reference.

<a id="canonical-1220010012002113-1113232302033131-3222333210210201-3111003002001212-2012033201320212-1120121032203323-1311000133223132-1201011010223101"></a>

## Next pages — selected_tenant_vip / 320232123013 / 5

- [re_acl.selected_tenant_vip.public_ip_refs](resources--fast_acl--reference--group-001.md#canonical-2100303222200103-2310111221033132-3220031130200001-1021132210220001-3013101002232022-3323303001031200-1201110302333332-0312100000011003)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-2100303222200103-2310111221033132-3220031130200001-1021132210220001-3013101002232022-3323303001031200-1201110302333332-0312100000011003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011213113303311-3001302212320233-3222123333022300-0212032013133032-1322030120223123-2103001303201110-1311230302302321-2121000311211033"></a>

## re_acl.selected_tenant_vip.public_ip_refs — public_ip_refs / 312320002210 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221)
- [re_acl.selected_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-2313112310232031-0132231301303321-1222013103002123-1203001002202003-2121231321230200-2210112100123321-0011302310000002-2202122102102002)
- re_acl.selected_tenant_vip.public_ip_refs

<a id="canonical-3220331011030310-1022210133001210-1300333123001300-3011220013331220-2231302121303331-0320322031000011-2212320120013301-3330220102033211"></a>

Type: `"object"`. list nested block, Optional.

Select Public VIP(s). Select additional public VIP(s)

Upstream description:

Select additional public VIP(s)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
public_ip_refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-2123100211300013-0220100000223232-3322001222011113-2013122030121133-2321311032100330-1322122113202310-2332223331031333-0311232210200103"></a>

## Direct properties — public_ip_refs / 312320002210 / 3

<a id="canonical-2221023021030002-0333223230110311-0022220302223122-0013220313000322-1301333232201021-0131223210302311-1301001012310203-3013101323001323"></a>

<a id="canonical-0021331300123211-3123332301311332-2231131230020133-3002100020200323-2023132223133310-1330222211031203-2103332320110333-0201311300110221"></a>

## name property — public_ip_refs / 312320002210 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-3301301000130003-2100322102330301-1321032301022322-3133302120023020-0330213012101322-1203111201002332-3312123233010331-2211330203312110"></a>

<a id="canonical-2032202333220003-3020302133232131-1211110110012033-2203012211230133-0012120310011111-3130023202332130-3320000203303331-3111011103030012"></a>

## namespace property — public_ip_refs / 312320002210 / 5

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
}
```

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

<a id="canonical-2022332020333202-2000012312232102-2131112322230232-2322312021201001-3030311212111213-0032313110010131-3020123323231321-0012023312320110"></a>

<a id="canonical-0221130131300031-0122310333120312-1213011133012013-2230332002300300-3113332213101001-1233100202012330-1213220102233130-0321032110103333"></a>

## tenant property — public_ip_refs / 312320002210 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-3023312113033013-1020212000112012-2022030223112102-1110231233203013-2103210133233312-2022303121300103-3112333120321321-0021110323203112"></a>

## Next pages — public_ip_refs / 312320002210 / 7

- [re_acl.selected_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-2313112310232031-0132231301303321-1222013103002123-1203001002202003-2121231321230200-2210112100123321-0011302310000002-2202122102102002)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200200133111212-0100230300122021-1210220131100013-2030022001121230-0213000012030003-1233123031322131-1133220110130332-0133312032100110"></a>

## site_acl — site_acl / 201310220012 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- site_acl

<a id="canonical-0210112023132223-2100131101110110-3332213000200131-2011231123303112-3110103303220200-2121123213022301-0002112303001221-3213202223312213"></a>

Type: `"object"`. single nested block, Optional.

Fast ACL for Site. Fast ACL definition for Site.

Upstream description:

Fast ACL definition for Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_services",
    "interface_services"),
  validators.ConflictingObjectAttributes("all_services",
    "vip_services"),
  validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("interface_services",
    "vip_services")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\"]",
  "x-ves-oneof-field-vip_choice": "[\"all_services\",\"interface_services\",\"vip_services\"]"
}
```

Terraform syntax:

```terraform
site_acl {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010211013331001-2023210222220010-2223321323030101-0221230000222221-3123202032330102-2202131331123010-0132202300222212-1201003022302033"></a>

## Direct properties — site_acl / 201310220012 / 3

- [all_services](resources--fast_acl--reference--group-001.md#canonical-3002331313121103-1120320220322131-3333123321102023-1230312112333102-2210331213111202-1111211200332330-1203223100112323-0022123012201123): complete subsection reference.

- [fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021): complete subsection reference.

- [inside_network](resources--fast_acl--reference--group-001.md#canonical-1011113211300013-0310222210321103-3120202103223210-2302031103003022-3303133300033132-3011130220111202-0023312230200233-0130212220210212): complete subsection reference.

- [interface_services](resources--fast_acl--reference--group-001.md#canonical-1321200022302131-0120001312332131-0201323001110020-0202332123131202-1111131111022301-2111230113101021-0033321301200201-1223203132133322): complete subsection reference.

- [outside_network](resources--fast_acl--reference--group-001.md#canonical-0100002120122101-2203013320202213-1010212021212123-2303103001020030-2000123211301003-0333310303223020-3223201133302113-0013132000233110): complete subsection reference.

- [vip_services](resources--fast_acl--reference--group-001.md#canonical-1121310113111120-3012323012103120-0001121322200330-3011123103320332-2101132030100233-0211202212100030-2331013021000231-2333030222201101): complete subsection reference.

<a id="canonical-2021031330331113-3210013323310113-1322200120123011-3310031031000033-3031331211133003-2320111000320131-1132331320133113-2022013210200330"></a>

## Next pages — site_acl / 201310220012 / 4

- [site_acl.all_services](resources--fast_acl--reference--group-001.md#canonical-3002331313121103-1120320220322131-3333123321102023-1230312112333102-2210331213111202-1111211200332330-1203223100112323-0022123012201123)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- [site_acl.inside_network](resources--fast_acl--reference--group-001.md#canonical-1011113211300013-0310222210321103-3120202103223210-2302031103003022-3303133300033132-3011130220111202-0023312230200233-0130212220210212)
- [site_acl.interface_services](resources--fast_acl--reference--group-001.md#canonical-1321200022302131-0120001312332131-0201323001110020-0202332123131202-1111131111022301-2111230113101021-0033321301200201-1223203132133322)
- [site_acl.outside_network](resources--fast_acl--reference--group-001.md#canonical-0100002120122101-2203013320202213-1010212021212123-2303103001020030-2000123211301003-0333310303223020-3223201133302113-0013132000233110)
- [site_acl.vip_services](resources--fast_acl--reference--group-001.md#canonical-1121310113111120-3012323012103120-0001121322200330-3011123103320332-2101132030100233-0211202212100030-2331013021000231-2333030222201101)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-3002331313121103-1120320220322131-3333123321102023-1230312112333102-2210331213111202-1111211200332330-1203223100112323-0022123012201123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210023323230213-3102212322332320-2202113211311233-1200210202113003-0323112210131212-2330133031003230-2030023123233300-2013102020010130"></a>

## site_acl.all_services — all_services / 002112310223 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- site_acl.all_services

<a id="canonical-2001202110321132-0323203012111101-1223021000320122-0002313213030130-2232321231322220-0021023313222232-0202030202332233-2301330223313113"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all services.

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
all_services = {}
```

<a id="canonical-2223113331220011-2003303013231122-0031330310320212-1113031010302100-3212020032322223-1303213132022010-0113322202200221-1120323310233322"></a>

## Direct properties — all_services / 002112310223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132100300300130-2223303011133132-1330123213200310-2233100130033310-3030211313230132-1212301330232233-3213323001201220-0322030203012330"></a>

## Next pages — all_services / 002112310223 / 4

- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123222331111312-3103202302322321-0010121103230323-3232231012313023-3012122120201000-3133002222031300-0131011330010100-1310112001101311"></a>

## site_acl.fast_acl_rules — fast_acl_rules / 023122002120 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- site_acl.fast_acl_rules

<a id="canonical-1213000133313233-3120011203322202-3030010321322320-2103231231130020-1222211223113301-2220031231333320-2312120202103030-0101221321123213"></a>

Type: `"object"`. list nested block, Optional.

Rules. Fast ACL rules to match.

Upstream description:

Fast ACL rules to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix")}
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
fast_acl_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132111322030133-2033330233033111-1122012232331033-1002022111010202-3111233232221331-3223321311033223-0232110120301220-2100120000103131"></a>

## Direct properties — fast_acl_rules / 023122002120 / 3

- [action](resources--fast_acl--reference--group-001.md#canonical-2321023022222021-0200223332121311-2013111113203100-3230132223203312-3132013201311133-3012003001001202-0231211232121310-3223023011031222): complete subsection reference.

- [ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-0323313132130031-2133201223301022-2303032133200032-2113320213211102-1323032130132221-0230112032020023-2101202323012303-2231200133113010): complete subsection reference.

- [metadata](resources--fast_acl--reference--group-001.md#canonical-2101232110302332-1120212113301231-1032120103123202-1123001211000130-0222033110322200-3330031111013121-2210101132030211-1131122301033233): complete subsection reference.

- [port](resources--fast_acl--reference--group-001.md#canonical-2111312000322020-3101211201200202-2303133201012311-0031111313033003-1130300321231133-2313013031123003-2300132113301310-1033023230021002): complete subsection reference.

- [prefix](resources--fast_acl--reference--group-001.md#canonical-2022103011003101-2303300133020322-1131111223011122-0220301311130022-1333310313231122-1312230231203133-1222001330031132-2123301002033021): complete subsection reference.

<a id="canonical-1310131210023230-0203323022201011-1323221230231302-3020203303212010-0023201033131132-1330332311023033-2031001001103222-3131013002100300"></a>

## Next pages — fast_acl_rules / 023122002120 / 4

- [site_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-2321023022222021-0200223332121311-2013111113203100-3230132223203312-3132013201311133-3012003001001202-0231211232121310-3223023011031222)
- [site_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-0323313132130031-2133201223301022-2303032133200032-2113320213211102-1323032130132221-0230112032020023-2101202323012303-2231200133113010)
- [site_acl.fast_acl_rules.metadata](resources--fast_acl--reference--group-001.md#canonical-2101232110302332-1120212113301231-1032120103123202-1123001211000130-0222033110322200-3330031111013121-2210101132030211-1131122301033233)
- [site_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-2111312000322020-3101211201200202-2303133201012311-0031111313033003-1130300321231133-2313013031123003-2300132113301310-1033023230021002)
- [site_acl.fast_acl_rules.prefix](resources--fast_acl--reference--group-001.md#canonical-2022103011003101-2303300133020322-1131111223011122-0220301311130022-1333310313231122-1312230231203133-1222001330031132-2123301002033021)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-2321023022222021-0200223332121311-2013111113203100-3230132223203312-3132013201311133-3012003001001202-0231211232121310-3223023011031222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300032330103023-2211112323212031-3011102332010321-2311313012303302-1101230231013330-2230010103222331-1130133020303033-0322102330003233"></a>

## site_acl.fast_acl_rules.action — action / 303133221321 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- site_acl.fast_acl_rules.action

<a id="canonical-3302002002122313-2301103031020011-3200320132101311-3311322130112011-2321022303120000-2313213031222000-2232130330302012-3220102332200313"></a>

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

<a id="canonical-0323131221320233-2032221021132203-2230330301201331-2120100113122332-1100333323030000-3003210032130111-3233203032212321-0313021010031011"></a>

## Direct properties — action / 303133221321 / 3

- [policer_action](resources--fast_acl--reference--group-001.md#canonical-2221203321120020-0122003030120223-1102012200130021-3323202212112111-1131302110020123-1330321231323101-2320010122110200-2201300021303001): complete subsection reference.

- [protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-3223200333033130-3333302210113232-3022221201003133-1200121200002033-0303020320120312-0210322323011100-1110022122213113-2313000130133000): complete subsection reference.

<a id="canonical-0121302222110021-3233031022213201-0002232121311121-0332300313032133-0223330022223032-0202132312111030-2323123301102223-0130322121211322"></a>

<a id="canonical-3331310030330333-0330010122332301-0212213131313300-1121333300212333-0130210101010021-1301203211030313-2312320030020033-0133132020333203"></a>

## simple_action property — action / 303133221321 / 4

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

<a id="canonical-1311202312133313-2222011023100311-1323100123121313-3001200130123303-0223131223000323-1202130301330100-0013010011111323-2033120023301221"></a>

## Next pages — action / 303133221321 / 5

- [site_acl.fast_acl_rules.action.policer_action](resources--fast_acl--reference--group-001.md#canonical-2221203321120020-0122003030120223-1102012200130021-3323202212112111-1131302110020123-1330321231323101-2320010122110200-2201300021303001)
- [site_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-3223200333033130-3333302210113232-3022221201003133-1200121200002033-0303020320120312-0210322323011100-1110022122213113-2313000130133000)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-2221203321120020-0122003030120223-1102012200130021-3323202212112111-1131302110020123-1330321231323101-2320010122110200-2201300021303001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022131330030002-3222323022313330-1000100031033230-0002231132231102-2231310203102001-2210223022013223-0312222233201233-3023020111020200"></a>

## site_acl.fast_acl_rules.action.policer_action — policer_action / 110201230211 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- [site_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-2321023022222021-0200223332121311-2013111113203100-3230132223203312-3132013201311133-3012003001001202-0231211232121310-3223023011031222)
- site_acl.fast_acl_rules.action.policer_action

<a id="canonical-3111013313030110-0111031233222000-3330113210231300-3332210132330121-3231231233022013-3020013123103203-1313122013232012-2302031213100110"></a>

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

<a id="canonical-2103101122231131-1011021232201231-0201123330111230-2303102310203300-3221002113321100-3322022120223121-1022210320112022-0232120120220201"></a>

## Direct properties — policer_action / 110201230211 / 3

- [ref](resources--fast_acl--reference--group-001.md#canonical-3201223330303202-1301332122132323-3212101210103223-1030130301010123-1111222300312321-0120220030102320-2200033123313223-3132211220103120): complete subsection reference.

<a id="canonical-2101211121103313-0012012103331330-3132203323323322-3313202011233020-2013230312302203-3013300302320123-3200331321132300-2022130202133132"></a>

## Next pages — policer_action / 110201230211 / 4

- [site_acl.fast_acl_rules.action.policer_action.ref](resources--fast_acl--reference--group-001.md#canonical-3201223330303202-1301332122132323-3212101210103223-1030130301010123-1111222300312321-0120220030102320-2200033123313223-3132211220103120)
- [site_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-2321023022222021-0200223332121311-2013111113203100-3230132223203312-3132013201311133-3012003001001202-0231211232121310-3223023011031222)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-3201223330303202-1301332122132323-3212101210103223-1030130301010123-1111222300312321-0120220030102320-2200033123313223-3132211220103120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221332320313020-2232232112200112-0022320212102111-0121133322203133-2221010120120303-2333230321311001-1012022332123033-0211123003021130"></a>

## site_acl.fast_acl_rules.action.policer_action.ref — ref / 222120010301 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- [site_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-2321023022222021-0200223332121311-2013111113203100-3230132223203312-3132013201311133-3012003001001202-0231211232121310-3223023011031222)
- [site_acl.fast_acl_rules.action.policer_action](resources--fast_acl--reference--group-001.md#canonical-2221203321120020-0122003030120223-1102012200130021-3323202212112111-1131302110020123-1330321231323101-2320010122110200-2201300021303001)
- site_acl.fast_acl_rules.action.policer_action.ref

<a id="canonical-1201013223301020-1312101233320332-0331023212112111-2211210112231210-0030322123113311-2311301111331033-3310320113220003-3321012030300020"></a>

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

<a id="canonical-1312121023320223-2112032011302300-2023101200222112-2102131101130103-2133102030100332-0203110103110301-0212303320123210-2302232302033022"></a>

## Direct properties — ref / 222120010301 / 3

<a id="canonical-0222010221330122-2002222312203000-3032023020230131-1231102123032313-1332001110232022-3233311113320313-1302033013100322-2322223110000200"></a>

<a id="canonical-2212313321133033-1003322320000221-0232031213111320-3023131101312322-1323122213211132-0130130221120232-2022311110223211-0322200331023002"></a>

## kind property — ref / 222120010301 / 4

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

<a id="canonical-2111110023001022-1010010321011231-0330211202211301-0332101201330203-1103232312010331-0001332333311230-2020112013113332-2113000233113322"></a>

<a id="canonical-0311123320222112-1023310233203122-1112200012113203-2002303031113132-1001001212123213-0310210113330100-1010000220200211-0210123203303200"></a>

## name property — ref / 222120010301 / 5

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

<a id="canonical-3210211210303003-3210113010113122-2200030031230301-3032032213032031-1003100333003133-2122132320320101-1003223220120310-2232322032121131"></a>

<a id="canonical-0210233302320302-0032120030303033-2113011302023111-1232310300023011-3102121213132111-1002321112013220-3323100011101021-1121330310222113"></a>

## namespace property — ref / 222120010301 / 6

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

<a id="canonical-2210132000102221-2001231122301220-2122013220031120-0101220310130223-3133132312001100-3101223333101011-1223123112021211-2330202013013012"></a>

<a id="canonical-2313031113103013-1010223201001223-1322103302120123-3303301231200232-3231322213111000-1302112110320210-0211101201323110-1020003112022202"></a>

## tenant property — ref / 222120010301 / 7

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

<a id="canonical-3100232000311330-1032330330321321-2000332012200102-2231031303001101-0010011331103010-0332123132222223-1123303130100113-1112302010002331"></a>

<a id="canonical-1201233000302100-2100300220321001-3123302112132102-3001201121220101-2302301223313331-2302102333212021-2300212030231300-1021010001003331"></a>

## uid property — ref / 222120010301 / 8

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

<a id="canonical-3222123121213202-1002122201120302-0233303213210032-0210123133113031-0330113130301031-0230013132121301-2000313132131233-1132002013132321"></a>

## Next pages — ref / 222120010301 / 9

- [site_acl.fast_acl_rules.action.policer_action](resources--fast_acl--reference--group-001.md#canonical-2221203321120020-0122003030120223-1102012200130021-3323202212112111-1131302110020123-1330321231323101-2320010122110200-2201300021303001)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-3223200333033130-3333302210113232-3022221201003133-1200121200002033-0303020320120312-0210322323011100-1110022122213113-2313000130133000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330020100222201-0202122300311331-1012302001011313-2033001211233310-2312303321121120-0313320213023322-3320123323232130-1213013313113012"></a>

## site_acl.fast_acl_rules.action.protocol_policer_action — protocol_policer_action / 230121222222 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- [site_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-2321023022222021-0200223332121311-2013111113203100-3230132223203312-3132013201311133-3012003001001202-0231211232121310-3223023011031222)
- site_acl.fast_acl_rules.action.protocol_policer_action

<a id="canonical-0313322010100321-1213120033203313-3200332302223001-3133033102131000-1222020201332000-3111100120020303-0213111131022110-1210312110323312"></a>

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

<a id="canonical-2230132011122130-3223203332210303-0330321322111133-1312031303332223-2302331110111121-0113022230302132-3212322213013101-1211220233310101"></a>

## Direct properties — protocol_policer_action / 230121222222 / 3

- [ref](resources--fast_acl--reference--group-001.md#canonical-1021210323021001-2002121223312213-1100303202132011-1330333133003021-3233032111131130-1302332120112021-1113123012332310-3120013000001232): complete subsection reference.

<a id="canonical-3211220013232031-1311001323122221-3213220320331230-3313010201303303-2100302330123230-2222003222221223-0011010100333233-2231300212102333"></a>

## Next pages — protocol_policer_action / 230121222222 / 4

- [site_acl.fast_acl_rules.action.protocol_policer_action.ref](resources--fast_acl--reference--group-001.md#canonical-1021210323021001-2002121223312213-1100303202132011-1330333133003021-3233032111131130-1302332120112021-1113123012332310-3120013000001232)
- [site_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-2321023022222021-0200223332121311-2013111113203100-3230132223203312-3132013201311133-3012003001001202-0231211232121310-3223023011031222)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-1021210323021001-2002121223312213-1100303202132011-1330333133003021-3233032111131130-1302332120112021-1113123012332310-3120013000001232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132330003212302-1033133331102111-1120032311320022-2321202213102111-0310101112122001-0210311101323203-3023102333112311-1210032210203002"></a>

## site_acl.fast_acl_rules.action.protocol_policer_action.ref — ref / 031220203333 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- [site_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-2321023022222021-0200223332121311-2013111113203100-3230132223203312-3132013201311133-3012003001001202-0231211232121310-3223023011031222)
- [site_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-3223200333033130-3333302210113232-3022221201003133-1200121200002033-0303020320120312-0210322323011100-1110022122213113-2313000130133000)
- site_acl.fast_acl_rules.action.protocol_policer_action.ref

<a id="canonical-0330001223012233-2000201022122200-2231130010330203-1312320003111220-3013013031133101-3203220110331033-2113212332132331-1212323122032023"></a>

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

<a id="canonical-2113111121221110-0330231020232223-3330000103122321-0100123232323233-2313233330103113-3210111010122210-0323021010330232-0122311031000131"></a>

## Direct properties — ref / 031220203333 / 3

<a id="canonical-0301101323331131-2033112101223110-1220201002212012-1030312030100010-3233311023131112-1003001121130213-3001200121213102-1100023102203002"></a>

<a id="canonical-1110301312013101-2113300333220303-3231302222330130-0020310012022010-0022231201211111-0211330112301111-1212223001210201-0020003003232000"></a>

## kind property — ref / 031220203333 / 4

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

<a id="canonical-1000211132011122-3132211131231032-3203112023002312-0201331020301312-0023331002032301-3131321023003131-0112121002210123-2220322032223221"></a>

<a id="canonical-1000211111103021-2130133021223001-0111231231310200-1113303330221223-0103112210221203-1230113022210020-0030001013323300-1112102203201310"></a>

## name property — ref / 031220203333 / 5

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

<a id="canonical-3132110133030000-2300312231123233-1111303000200021-2301032222111220-0211122022030022-1001022111203210-1000302230100023-1202223222213003"></a>

<a id="canonical-3300201320002111-0310132200020121-1011010331032303-2120003000031022-2200111212302012-3020103231330132-0202022223311011-1222232212322022"></a>

## namespace property — ref / 031220203333 / 6

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

<a id="canonical-1123122123100302-1003222021313003-1030202211121333-1231222322121223-0330031212112112-3331133132031312-0020201112002000-1011011212130123"></a>

<a id="canonical-2231102023133020-0223022031330122-2132113120222302-2122111333302030-1323332230010322-0131313112110032-0302122332102030-3113330233333232"></a>

## tenant property — ref / 031220203333 / 7

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

<a id="canonical-0332313201222132-1231123303101310-1032022221002333-3233202312211002-3222112010302102-1221323201021133-0032030322110120-0231030203133031"></a>

<a id="canonical-1230133300320213-2112232211032211-2320112303023003-2210223302313102-1132021103022210-0300103303002023-2033132302012223-3031120023123020"></a>

## uid property — ref / 031220203333 / 8

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

<a id="canonical-1130012210203021-0103212233122002-2133212023303021-2121100302100103-1301133122003311-1112321300301321-0313302311022110-2330202233212131"></a>

## Next pages — ref / 031220203333 / 9

- [site_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-3223200333033130-3333302210113232-3022221201003133-1200121200002033-0303020320120312-0210322323011100-1110022122213113-2313000130133000)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-0323313132130031-2133201223301022-2303032133200032-2113320213211102-1323032130132221-0230112032020023-2101202323012303-2231200133113010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230311221303012-2133220311312110-1131202333122121-0312133021100231-3310023001100230-3130113203033320-1102001213231223-0032322010020323"></a>

## site_acl.fast_acl_rules.ip_prefix_set — ip_prefix_set / 033301121220 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- site_acl.fast_acl_rules.ip_prefix_set

<a id="canonical-3322223021113030-3221231231223211-1001203202213311-2313013022323321-3230012332021120-3030223003020130-0232121013332323-3131000111112300"></a>

Type: `"object"`. single nested block, Optional.

List of references to ip\_prefix\_set objects.

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

Terraform syntax:

```terraform
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200202000010233-2331132330030233-0012023211022033-1331113012232033-3131030021130001-2233300103002131-1313023020003000-2033223023023313"></a>

## Direct properties — ip_prefix_set / 033301121220 / 3

- [ref](resources--fast_acl--reference--group-001.md#canonical-2313310231332020-0101333010332300-2002321300230031-2021312123103002-3203203331322331-1012201211330002-2131210303002331-1223332300231230): complete subsection reference.

<a id="canonical-3132310321130230-1113020311033220-3210033011031330-2023132011332212-2231232022003113-0232322321101202-0122003311120011-2221123320333130"></a>

## Next pages — ip_prefix_set / 033301121220 / 4

- [site_acl.fast_acl_rules.ip_prefix_set.ref](resources--fast_acl--reference--group-001.md#canonical-2313310231332020-0101333010332300-2002321300230031-2021312123103002-3203203331322331-1012201211330002-2131210303002331-1223332300231230)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-2313310231332020-0101333010332300-2002321300230031-2021312123103002-3203203331322331-1012201211330002-2131210303002331-1223332300231230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331203201220022-3210223110220031-2023310213011111-2210131030022110-1023123012131113-1210202230212220-3302201303323003-1112021323303322"></a>

## site_acl.fast_acl_rules.ip_prefix_set.ref — ref / 310203122003 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- [site_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-0323313132130031-2133201223301022-2303032133200032-2113320213211102-1323032130132221-0230112032020023-2101202323012303-2231200133113010)
- site_acl.fast_acl_rules.ip_prefix_set.ref

<a id="canonical-0010103121201203-1033221132212222-3313210132012122-0201111203210232-0223333220123132-0123012122321310-3003220301323101-1322121011113012"></a>

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

<a id="canonical-3301220233232323-0121002013321223-2033003231120002-3212031200201100-2032333001201301-2123031321121213-1322210330223220-2230302313220320"></a>

## Direct properties — ref / 310203122003 / 3

<a id="canonical-1323233213001133-0232101133032310-1001333230102300-2100130212211020-3321131213100101-1313311123221300-0133202131202102-2020323203012011"></a>

<a id="canonical-3103230322013102-2110312302220211-3301123113301000-1330321030111323-0330123120103310-1230212122011012-3213021332232001-2201000300110023"></a>

## kind property — ref / 310203122003 / 4

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

<a id="canonical-1332300001323031-1022121320322213-1223313302330202-0012101210330022-0112120023021020-3011311333221300-3203023130002331-0203332031033001"></a>

<a id="canonical-3131232102332212-1101232221201031-3203311210211210-0231300231223001-3222222210310102-0030230011120031-0322022022131333-3332302112322213"></a>

## name property — ref / 310203122003 / 5

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

<a id="canonical-2200031130302301-2133330102120230-0221333300003312-2120200311132202-2120322323012032-1013202323130121-2113232033020232-3131001132032123"></a>

<a id="canonical-2100122321112033-2131120123230311-2330030301323002-3322203221312232-0231222301130221-0331020210020112-2111011320333000-1130231312213202"></a>

## namespace property — ref / 310203122003 / 6

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

<a id="canonical-1230210312001010-3000032110221032-1103031301202202-1322213322102213-1222130231031011-0333333312113103-2321003101113112-3213233311013030"></a>

<a id="canonical-3221312023023323-1133222312132003-0101002030012232-3230311213010321-2310311203022323-0123030123103103-2232312211003210-0022123322033101"></a>

## tenant property — ref / 310203122003 / 7

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

<a id="canonical-2000331131312131-1001113220101312-0312033311220331-3120022332222222-0023020230230123-3231311220031011-1021330212131323-0011323131133230"></a>

<a id="canonical-0102012102011300-0011003230132122-2312203311202333-0322200230313230-1330230212121032-0301233000021002-3021002013323313-1331112223311110"></a>

## uid property — ref / 310203122003 / 8

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

<a id="canonical-2132132222322322-0222311030203311-0311032011200313-2211100230130031-1213012330113230-3010102021100301-2021221312223012-3121301121202121"></a>

## Next pages — ref / 310203122003 / 9

- [site_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-0323313132130031-2133201223301022-2303032133200032-2113320213211102-1323032130132221-0230112032020023-2101202323012303-2231200133113010)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-2101232110302332-1120212113301231-1032120103123202-1123001211000130-0222033110322200-3330031111013121-2210101132030211-1131122301033233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011221322101011-1310223021333200-2311003331111033-3231333333321333-1003102013310131-0230231102300001-2033320002022202-3100133323213330"></a>

## site_acl.fast_acl_rules.metadata — metadata / 032132321002 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- site_acl.fast_acl_rules.metadata

<a id="canonical-2313322301233213-0302103210100232-0111200002311223-3222232011020210-1000321030331301-2000110110122030-0003003100132010-3311202113212203"></a>

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

<a id="canonical-2131200323002301-3012232310222221-0001320032102331-0123321110110211-0211013322033103-3311132001333220-1230222212031133-0012212001202303"></a>

## Direct properties — metadata / 032132321002 / 3

<a id="canonical-2010001000033113-2010322302111313-1100000022222213-0213320001302122-3030323223012022-3332130100021320-3113311121003120-3222200030201002"></a>

<a id="canonical-3301000211213333-2011331230321203-0202321030110321-2200222022103222-1102100023001232-1023031313212123-1221033032301011-1100300322312301"></a>

## description_spec property — metadata / 032132321002 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2213011320312132-3001232021111310-0212310022123300-2332321133101320-2322233111312113-1120022120000132-3332102110120300-2031030320020203"></a>

<a id="canonical-2221221031113112-2222112113320123-2330033211102022-0300011202333323-1012200202133331-0123300232112013-3011200302321302-2211330302100012"></a>

## name property — metadata / 032132321002 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2122201313200330-1311202120231123-2303313223012222-3312313023221220-3113300211030222-3113011201011223-2303013013111322-1323033321332003"></a>

## Next pages — metadata / 032132321002 / 6

- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-2111312000322020-3101211201200202-2303133201012311-0031111313033003-1130300321231133-2313013031123003-2300132113301310-1033023230021002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121033332230322-0201020022330000-1203211313211300-0233211013111223-2320331003011103-2032020021111313-2133201313100312-1303110103021032"></a>

## site_acl.fast_acl_rules.port — port / 202131221130 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- site_acl.fast_acl_rules.port

<a id="canonical-1211210303302232-1322333111030122-3121032331200033-3020000003333130-1132213123011322-2000022231131003-1211130321331320-2313221022132001"></a>

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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
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

<a id="canonical-0123312302023201-3133233210123212-3102002000000133-1011212030310322-3021330201230113-1022100222202032-2122303230313220-2222232020302301"></a>

## Direct properties — port / 202131221130 / 3

- [all](resources--fast_acl--reference--group-001.md#canonical-0101003313322330-0112001122300313-3321303131230020-3032211010013020-2013103330200001-0311312023322210-2300001013232320-1021331200310302): complete subsection reference.

- [DNS](resources--fast_acl--reference--group-001.md#canonical-1302012222320101-2012110301330333-1122211223030100-2122232102230321-1103211330311313-3323022220223111-0112222322303230-1211003332300201): complete subsection reference.

<a id="canonical-1033013212131230-1010311212231102-2310001230113123-1011101223022013-1222023112221232-2212231211231011-2211200001120030-3210311001103101"></a>

<a id="canonical-1232030210202020-1322130031333223-1020031112333101-3323030032122110-0323231123132311-1233222120231101-0313121001220313-2202213230230233"></a>

## user_defined property — port / 202131221130 / 4

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

<a id="canonical-2001321002001323-1321103022330023-0003013210232332-2302130220213101-1333330010002032-0001003303123222-3011301300302210-0102332210203300"></a>

## Next pages — port / 202131221130 / 5

- [site_acl.fast_acl_rules.port.all](resources--fast_acl--reference--group-001.md#canonical-0101003313322330-0112001122300313-3321303131230020-3032211010013020-2013103330200001-0311312023322210-2300001013232320-1021331200310302)
- [site_acl.fast_acl_rules.port.dns](resources--fast_acl--reference--group-001.md#canonical-1302012222320101-2012110301330333-1122211223030100-2122232102230321-1103211330311313-3323022220223111-0112222322303230-1211003332300201)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-0101003313322330-0112001122300313-3321303131230020-3032211010013020-2013103330200001-0311312023322210-2300001013232320-1021331200310302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113313221213332-1022222012223031-2323212333333103-2200300210303312-0100000320233101-1202020323032323-3120120121331101-2030122001121323"></a>

## site_acl.fast_acl_rules.port.all — all / 223121302033 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- [site_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-2111312000322020-3101211201200202-2303133201012311-0031111313033003-1130300321231133-2313013031123003-2300132113301310-1033023230021002)
- site_acl.fast_acl_rules.port.all

<a id="canonical-3320013122130033-2010013312330311-2213023333210300-0130200212010333-2211311203000222-2310212030223113-0312313220020203-3302020123002220"></a>

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

<a id="canonical-2010220230203013-1312322233130201-1303112001102103-1221233230321013-2010022301131111-2313003302110030-1123003221333303-0220002302001331"></a>

## Direct properties — all / 223121302033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221232231132232-2323132220121000-3211230111230123-3323021302023230-3033322023033320-3221002103233233-0032013231000131-2223000203302210"></a>

## Next pages — all / 223121302033 / 4

- [site_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-2111312000322020-3101211201200202-2303133201012311-0031111313033003-1130300321231133-2313013031123003-2300132113301310-1033023230021002)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-1302012222320101-2012110301330333-1122211223030100-2122232102230321-1103211330311313-3323022220223111-0112222322303230-1211003332300201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333133131210022-2110322321323303-3202233011012212-1022111302300032-3113232330313300-3112302110001223-2310122211333113-1121220310022012"></a>

## site_acl.fast_acl_rules.port.DNS — DNS / 121310102220 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- [site_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-2111312000322020-3101211201200202-2303133201012311-0031111313033003-1130300321231133-2313013031123003-2300132113301310-1033023230021002)
- site_acl.fast_acl_rules.port.DNS

<a id="canonical-1200123303201100-0332222121333010-1331130230201230-2221123331312320-1101010021301322-3012233200031203-2132111303010132-2012001001100232"></a>

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

<a id="canonical-2320303301003220-1033033201232011-3213113302221332-0211133131201301-3022001132222111-2320223100020313-2132111012212331-3210213213013332"></a>

## Direct properties — DNS / 121310102220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131123102222100-1221231231301001-2322131302232210-1221012312231331-0121010311301321-1312330122223303-2123220113311131-0133201002301232"></a>

## Next pages — DNS / 121310102220 / 4

- [site_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-2111312000322020-3101211201200202-2303133201012311-0031111313033003-1130300321231133-2313013031123003-2300132113301310-1033023230021002)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-2022103011003101-2303300133020322-1131111223011122-0220301311130022-1333310313231122-1312230231203133-1222001330031132-2123301002033021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313013002002313-3022113320300031-1032320110101210-0021113023223120-0301233311101323-2123111323331231-1210323032333101-1013132332201201"></a>

## site_acl.fast_acl_rules.prefix — prefix / 213031223332 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- site_acl.fast_acl_rules.prefix

<a id="canonical-2023023130320121-0232021103113002-3222220012201110-3231300013211102-1302313100100113-2333300030023132-1002131023110020-0123120332020001"></a>

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

<a id="canonical-3102231223330322-3101203320203203-1010033011310301-0201010323102001-0321103212332312-1013113333330021-2020333223331312-1232320133320031"></a>

## Direct properties — prefix / 213031223332 / 3

<a id="canonical-0101112312101310-3222221003300130-1302003223220231-1300010033221303-2113100132312110-0021220303130210-3103101123220301-1131132310031112"></a>

<a id="canonical-1013003303112132-0303222231003101-1311321210123000-3211031300131220-0131021130202322-0200212020310213-0023331333322231-2212123201300222"></a>

## prefix property — prefix / 213031223332 / 4

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

<a id="canonical-3302211130033313-0303331331002102-3223200112210201-2021333322132013-1010110013022003-2100000321232320-0303120301112131-0111020233120302"></a>

## Next pages — prefix / 213031223332 / 5

- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-1011113211300013-0310222210321103-3120202103223210-2302031103003022-3303133300033132-3011130220111202-0023312230200233-0130212220210212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333131323010102-1301221022003201-1210233210020020-0302223002203301-2130212313301210-2132000031302020-3013100330312312-3013233301130223"></a>

## site_acl.inside_network — inside_network / 000331332212 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- site_acl.inside_network

<a id="canonical-1001013123000211-3010211112110201-2111001023021302-1321100212000233-0200031120331212-1231223221000031-0121121333320020-2133312020231321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

<a id="canonical-1203331223331331-1330033102232122-1000013212321210-1022330301131001-0311213122131232-2323133022100311-2131232310121222-2303330203201311"></a>

## Direct properties — inside_network / 000331332212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213303301223011-2123301032000010-1320313223111003-3222013333113012-3330320101122111-3000032011331110-3102303033133001-3222003001033010"></a>

## Next pages — inside_network / 000331332212 / 4

- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-1321200022302131-0120001312332131-0201323001110020-0202332123131202-1111131111022301-2111230113101021-0033321301200201-1223203132133322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130110202030211-1212320213102003-3222012303020312-2002110002202301-1332013301303022-3233232100010031-0122131300333030-2320100013321300"></a>

## site_acl.interface_services — interface_services / 302000003330 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- site_acl.interface_services

<a id="canonical-3322231330113110-2203003221011110-2120332013131202-3121302103222030-0203013001130012-3322120023122022-2221221201122120-1120213010313233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for interface services.

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
interface_services = {}
```

<a id="canonical-2302021111123132-0130011310221102-1123223331321123-1113103123101232-1323000111000332-2212013011033032-0210212032133302-2303211323031031"></a>

## Direct properties — interface_services / 302000003330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230301012333033-3120130203031232-3032313003203300-2012321300012302-0202112012222203-3200013003332122-0233100131121033-3122200233133313"></a>

## Next pages — interface_services / 302000003330 / 4

- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-0100002120122101-2203013320202213-1010212021212123-2303103001020030-2000123211301003-0333310303223020-3223201133302113-0013132000233110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122301333311310-3020001310103321-0220313232303222-3210130033232000-0312302021302233-2121012303213313-3010333330201011-2313100210321031"></a>

## site_acl.outside_network — outside_network / 111230022303 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- site_acl.outside_network

<a id="canonical-1001130101230122-3111133131222303-2232330233111011-1220030303010101-1030203100323310-0120023202202303-2201322000030322-3002233223123112"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

<a id="canonical-3220101000312131-1300032310323000-1023203222010022-0331233311323011-1213133330030003-1101332011002223-3202322210112231-2020332332212132"></a>

## Direct properties — outside_network / 111230022303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0300221313031032-2132231322130003-1130222230030220-1310012020123303-1132210011022112-1123210100103102-0332210211013300-3110210031231302"></a>

## Next pages — outside_network / 111230022303 / 4

- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-1121310113111120-3012323012103120-0001121322200330-3011123103320332-2101132030100233-0211202212100030-2331013021000231-2333030222201101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331100011112302-1222332111330223-0123300301332031-2301032301223100-3130033011310031-0100201311003132-3030010223103230-0102112320231020"></a>

## site_acl.vip_services — vip_services / 323111121100 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- site_acl.vip_services

<a id="canonical-1000222221131001-0201030101213220-2213322200112130-2230120001030211-1111132130201300-1222302121021031-1002222123033222-2133013300133021"></a>

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
vip_services = {}
```

<a id="canonical-1331233023100231-3133223133321212-1111332202012220-1032121100232120-0123200313300012-3001222123230323-0112020332012323-0112331232332031"></a>

## Direct properties — vip_services / 323111121100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300303112310200-3303301230032330-1001211330231010-3203232030201133-2201022222011033-1202203020102132-1332310202230311-3003323213132303"></a>

## Next pages — vip_services / 323111121100 / 4

- [site_acl](resources--fast_acl--reference--group-001.md#canonical-0201213330331112-1112120003222321-2001112121100120-0103311332002230-0201010113012332-0111321222020211-3211020031331102-0102201002003322)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)

<a id="canonical-2222200211213322-0210312320233133-3101222313001002-3200023112003003-2231122211030313-0030223121111320-3321112132323230-1211100021300311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121333230301221-0301202313013113-1202012223001223-2310110100012021-2012320033022012-3021321023220112-0032230310333311-1122133103200213"></a>

## timeouts — timeouts / 132212101301 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- timeouts

<a id="canonical-2113103320231323-2313313023211221-3121130102023332-2003111211203121-2220212210302132-2100311120322312-2230102101110220-1202231002210231"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3012023323332011-1123030013302300-3220133102321022-0132003213003200-0310333310101221-1002223300103220-0123010322111020-2123220023302310"></a>

## Direct properties — timeouts / 132212101301 / 3

<a id="canonical-3123113132301203-2100023011311002-0021332200320330-1313111333002100-0221123100122212-2223302221333201-3111222112333310-1002121103110221"></a>

<a id="canonical-3131010331013120-2022231133110223-0130322230232023-1001003111002210-1021020021030301-1201333300200301-0121332232100212-1231200222111031"></a>

## create property — timeouts / 132212101301 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1021112223333130-1130210232233233-2223012030321310-3013111103201312-1331022232132323-2331130002202301-1130302333101331-2332110311130222"></a>

<a id="canonical-0300313312213231-3312011313131000-1120130202012021-0333332330110310-0012331102130103-3013020023022333-3023301301303312-1332303010302100"></a>

## delete property — timeouts / 132212101301 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3323003310010333-1222202331110112-0212023320123300-0230130320130210-2213203231231132-1011103103130320-1321210301321202-0333310120021300"></a>

<a id="canonical-1102010002102001-2103212200211202-2320131313311122-2132110232303123-1011001000212210-2121010221232202-1313010313001002-2110230303331220"></a>

## read property — timeouts / 132212101301 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0112100001222103-2303331111231020-3330110301213321-0310132100010131-3220322012131130-3333200103303233-0322013321001211-2111001033311233"></a>

<a id="canonical-1233201222211000-0030123312103101-0120200122122213-2121112233113322-2202131222312032-0103332111030330-2033122033322032-0130122213201201"></a>

## update property — timeouts / 132212101301 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1002231131112001-0331210201211112-3213321011332302-3030230023233103-3133302131132103-2111230221212003-2321303013120013-3230103010211301"></a>

## Next pages — timeouts / 132212101301 / 8

- [Property reference](resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102)
