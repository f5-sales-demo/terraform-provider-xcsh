---
page_title: "xcsh_network_firewall reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_firewall reference."
---

# xcsh_network_firewall reference

<a id="canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003232020321302-0321221210232232-2310031332121032-3122300202320231-3230202312001213-2020033132000301-0110320230231223-3303011203222201"></a>

## Property reference — Property reference / 302301021120 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- Property reference

<a id="canonical-3011123213023232-1320322231202133-3121320320021123-0310303320202123-1210132120000002-2132021303202133-2311321003202121-3110101132022103"></a>

## Direct properties — Property reference / 302301021120 / 3

- [active_enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-3001312303222121-0312302310023012-1101130213031230-0333300122321021-0200212131012113-2223020201000110-1033221221230031-1120130013233030): complete subsection reference.

- [active_fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-3110031110222221-1202323102011100-3331003003330103-0303323021130010-3102212020312231-0021130321211312-1222212111102211-3101302030000102): complete subsection reference.

- [active_forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-2233311113113112-1012303201230000-3131021231333032-0301110132312132-0000211131030031-0310103000022311-3312020012101120-3122100311021210): complete subsection reference.

- [active_network_policies](data-sources--network_firewall--reference--group-001.md#canonical-0300113221000300-3323031321112301-2303200000031221-3220101113022012-2132110033002231-3323003110211123-0002220301010033-3322101121111123): complete subsection reference.

<a id="canonical-1323230202213112-1100020301302102-0012302231303012-2020310213223223-0022113300130230-0222012312020330-0313310130212210-0331013021233012"></a>

<a id="canonical-0310231110013010-2101230131021230-3103320131210221-2201100333100331-0133033201201223-0000312212101031-2220120323202113-0222031321103302"></a>

## annotations property — Property reference / 302301021120 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

<a id="canonical-1200303010330021-2020100002120023-2333221131110312-3011011322131021-2003202312322123-3012322123033002-2033131022222301-3310233300133322"></a>

<a id="canonical-2010202321123230-2313121111120333-1311011000303133-2123031100123113-0332020222030321-0010300223302111-2230110201003230-0120301223212113"></a>

## description property — Property reference / 302301021120 / 5

Type: `"string"`. Computed.

Description of the NetworkFirewall.

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

- [disable_fast_acl](data-sources--network_firewall--reference--group-001.md#canonical-1112103131113333-1121212021221301-3111232301200132-0233211310020303-0203211310320133-1023300102123003-1030120300323032-1132232021223232): complete subsection reference.

- [disable_forward_proxy_policy](data-sources--network_firewall--reference--group-001.md#canonical-0100002122021020-1303013222121200-1101200333333311-2301022320231012-1200022022121321-1102131103022223-0200320323012013-2031001123113313): complete subsection reference.

- [disable_network_policy](data-sources--network_firewall--reference--group-001.md#canonical-0103010021311202-2013232200312121-3032223321111100-1223032112300103-2020000111230302-0012202131030220-3122212103110212-0113122131321023): complete subsection reference.

<a id="canonical-3000113322102212-1321020203220123-2130012121001232-1330022032113012-3123230200111130-3301020210230101-3302220023210220-1020013322213301"></a>

<a id="canonical-0212132023131030-0220012001230031-1232003300313032-1111302032233130-0321002100332233-3031213231033021-2312323330033220-3022101321030211"></a>

## ID property — Property reference / 302301021120 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0203111130201313-3223232131322112-0211331311030000-3222303023102000-1000323110222233-1321102110321220-0032133133022212-0232031200000011"></a>

<a id="canonical-0032003023032120-2000000311232313-2331203003123300-0100022320113010-1300003331100321-1120003000133212-3310201212031203-3000231022102011"></a>

## labels property — Property reference / 302301021120 / 7

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

<a id="canonical-3310313130312031-2232000300023001-3132122301233110-2131203202302021-3033021123303313-0120310103233000-0000032112032111-2012331102123132"></a>

<a id="canonical-2331133012212320-1301001201330321-1012310103202003-1130021030220232-3030310013111331-3123122023233322-0220232013200021-3322233120000001"></a>

## name property — Property reference / 302301021120 / 8

Type: `"string"`. Required.

Name of the NetworkFirewall.

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

<a id="canonical-1100133023222232-3110011021200310-0011333111002033-2220030201320320-3011230130221023-3001220121320113-1013323302203302-0220031312232122"></a>

<a id="canonical-0031313002031013-2201322211210031-1320112032200220-3313002032213020-2300133330101202-1012202103331213-3331103220211100-1023303332100331"></a>

## namespace property — Property reference / 302301021120 / 9

Type: `"string"`. Optional, Computed.

Namespace where the NetworkFirewall exists.

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

<a id="canonical-1111312021321000-3312211312023110-2203111020232131-1333232320112320-1203012103113003-3203211102220020-1302221321232022-0122102212121130"></a>

## All schema paths — Property reference / 302301021120 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_enhanced_firewall_policies` | [active_enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-3010001212200013-3103033011100002-2000333033221323-0123000122301112-2213013122213213-3221100320333100-0312300110102221-1110022212213203) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies` | [active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-2030122330300332-1013312030022001-3011121110311132-1013222310030012-1131032211212012-2232102301102120-2031321103000310-1120310132132033) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--network_firewall--reference--group-001.md#canonical-3311000101010123-3111211232101300-2313030112322102-3200322002333201-0133133010122013-1330233330333011-0212130120202012-2230000312221330) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--network_firewall--reference--group-001.md#canonical-0310100231102122-0100202131022002-3323110330013023-1313102211002212-1220310130223301-3230220320130110-0103001001331302-0233202312011112) |
| `active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--network_firewall--reference--group-001.md#canonical-3223012312003313-1200310113232230-3210020322032100-0031102213222233-2003200321120031-1201122232030100-2031301233131220-0002222200101321) |
| `active_fast_acls` | [active_fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-1103332321100131-3213313220332330-1333303120101011-3303123101221111-2220002110002020-3322102102323023-2332222310121300-2211233210302321) |
| `active_fast_acls.fast_acls` | [active_fast_acls.fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-1200232010132232-2312321331120203-3102033320231322-2113332330131322-3311101300101230-3212300110120333-2012333231130011-0310112021223332) |
| `active_fast_acls.fast_acls.name` | [active_fast_acls.fast_acls.name](data-sources--network_firewall--reference--group-001.md#canonical-1210103123232011-3232003233321311-1033103111133031-0212010323203231-1210220101311020-0012000223010100-1323131133202202-3323220123310303) |
| `active_fast_acls.fast_acls.namespace` | [active_fast_acls.fast_acls.namespace](data-sources--network_firewall--reference--group-001.md#canonical-1221001122310132-2230002130310002-3313231210022322-2010020320333030-1001013321011003-0110201103011331-1232200303020120-3220130212210112) |
| `active_fast_acls.fast_acls.tenant` | [active_fast_acls.fast_acls.tenant](data-sources--network_firewall--reference--group-001.md#canonical-0322011102011123-3111012112021002-2023030223123332-2111000223033321-1303031300003023-2110320312120232-1032133011132212-3330203202230122) |
| `active_forward_proxy_policies` | [active_forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-0021303122111230-1013000133332130-2332303332303323-0032131021331133-1033112010221033-1310233303321110-3022212132121331-0201120321023311) |
| `active_forward_proxy_policies.forward_proxy_policies` | [active_forward_proxy_policies.forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-0032322211223110-1000022110210021-0012030231303203-3003130023131103-0211101102322000-2322121010101311-0233120122022030-2111121303030002) |
| `active_forward_proxy_policies.forward_proxy_policies.name` | [active_forward_proxy_policies.forward_proxy_policies.name](data-sources--network_firewall--reference--group-001.md#canonical-3113022330230020-0023030122200012-3320313000103312-3201301013130032-2121211101323320-2110023211201301-1211210112022001-2311130130303233) |
| `active_forward_proxy_policies.forward_proxy_policies.namespace` | [active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--network_firewall--reference--group-001.md#canonical-0112201311213120-2213031313121311-1131010230213021-0302133322303130-1200103102311223-0200100100203300-2133103233331331-2311022123303131) |
| `active_forward_proxy_policies.forward_proxy_policies.tenant` | [active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--network_firewall--reference--group-001.md#canonical-3301222101221320-3300113132020210-1212031031200002-1023321232220200-0032001002132313-2022301021320332-1203002203300112-2020111303121032) |
| `active_network_policies` | [active_network_policies](data-sources--network_firewall--reference--group-001.md#canonical-2312022030130130-1313313112011221-0322000121202111-3223331132101013-0110233220001103-0222121110100132-3231333220133311-2222122111200120) |
| `active_network_policies.network_policies` | [active_network_policies.network_policies](data-sources--network_firewall--reference--group-001.md#canonical-3030213121113222-0011122101133131-3202020120012113-1022210131003322-1333032210121313-3033031010132233-0212320131332132-2021102211333323) |
| `active_network_policies.network_policies.name` | [active_network_policies.network_policies.name](data-sources--network_firewall--reference--group-001.md#canonical-2312301220200032-1003222313311302-0101101233100030-1122213223231332-3001211202130212-0103333211100312-2332122310320301-1200210101031232) |
| `active_network_policies.network_policies.namespace` | [active_network_policies.network_policies.namespace](data-sources--network_firewall--reference--group-001.md#canonical-3313121101020020-2313003301010222-2100320222100222-1130000212110110-1013021012311220-3202330300000121-2132130101002330-0233303113300130) |
| `active_network_policies.network_policies.tenant` | [active_network_policies.network_policies.tenant](data-sources--network_firewall--reference--group-001.md#canonical-3333301110223323-3001101032020032-2121000230001212-1100223130210021-1123221130303111-2102301002211031-1322320310333203-1211010021232210) |
| `annotations` | [annotations](data-sources--network_firewall--reference--group-001.md#canonical-1323230202213112-1100020301302102-0012302231303012-2020310213223223-0022113300130230-0222012312020330-0313310130212210-0331013021233012) |
| `description` | [description](data-sources--network_firewall--reference--group-001.md#canonical-1200303010330021-2020100002120023-2333221131110312-3011011322131021-2003202312322123-3012322123033002-2033131022222301-3310233300133322) |
| `disable_fast_acl` | [disable_fast_acl](data-sources--network_firewall--reference--group-001.md#canonical-0030330022021222-2220030333321200-2102310032230000-0031332233021110-3313213232201220-1320232103113102-0123133032310213-2202213012112330) |
| `disable_forward_proxy_policy` | [disable_forward_proxy_policy](data-sources--network_firewall--reference--group-001.md#canonical-3031232303002032-0120110310101132-2300001001200200-3010210033012120-0122203110223120-3232320032111012-0010213200020322-0101003230022311) |
| `disable_network_policy` | [disable_network_policy](data-sources--network_firewall--reference--group-001.md#canonical-3302010102221131-0122321003223010-2013230201020210-3202300102122221-2131302313232332-3220223110311301-1323030202322123-1031212320200102) |
| `id` | [ID](data-sources--network_firewall--reference--group-001.md#canonical-3000113322102212-1321020203220123-2130012121001232-1330022032113012-3123230200111130-3301020210230101-3302220023210220-1020013322213301) |
| `labels` | [labels](data-sources--network_firewall--reference--group-001.md#canonical-0203111130201313-3223232131322112-0211331311030000-3222303023102000-1000323110222233-1321102110321220-0032133133022212-0232031200000011) |
| `name` | [name](data-sources--network_firewall--reference--group-001.md#canonical-3310313130312031-2232000300023001-3132122301233110-2131203202302021-3033021123303313-0120310103233000-0000032112032111-2012331102123132) |
| `namespace` | [namespace](data-sources--network_firewall--reference--group-001.md#canonical-1100133023222232-3110011021200310-0011333111002033-2220030201320320-3011230130221023-3001220121320113-1013323302203302-0220031312232122) |

<a id="canonical-3100030230102103-1011231002013133-1301303010133012-1103223303110031-1311312321013001-2123213333010233-2313301010300211-3002333021122111"></a>

## Next pages — Property reference / 302301021120 / 11

- [active_enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-3001312303222121-0312302310023012-1101130213031230-0333300122321021-0200212131012113-2223020201000110-1033221221230031-1120130013233030)
- [active_fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-3110031110222221-1202323102011100-3331003003330103-0303323021130010-3102212020312231-0021130321211312-1222212111102211-3101302030000102)
- [active_forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-2233311113113112-1012303201230000-3131021231333032-0301110132312132-0000211131030031-0310103000022311-3312020012101120-3122100311021210)
- [active_network_policies](data-sources--network_firewall--reference--group-001.md#canonical-0300113221000300-3323031321112301-2303200000031221-3220101113022012-2132110033002231-3323003110211123-0002220301010033-3322101121111123)
- [disable_fast_acl](data-sources--network_firewall--reference--group-001.md#canonical-1112103131113333-1121212021221301-3111232301200132-0233211310020303-0203211310320133-1023300102123003-1030120300323032-1132232021223232)
- [disable_forward_proxy_policy](data-sources--network_firewall--reference--group-001.md#canonical-0100002122021020-1303013222121200-1101200333333311-2301022320231012-1200022022121321-1102131103022223-0200320323012013-2031001123113313)
- [disable_network_policy](data-sources--network_firewall--reference--group-001.md#canonical-0103010021311202-2013232200312121-3032223321111100-1223032112300103-2020000111230302-0012202131030220-3122212103110212-0113122131321023)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)

<a id="canonical-3001312303222121-0312302310023012-1101130213031230-0333300122321021-0200212131012113-2223020201000110-1033221221230031-1120130013233030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323320013301010-0013331020131201-1132322130332301-2012132023021100-0021110302302023-2201332323002222-3031303000213031-3232232100311322"></a>

## active_enhanced_firewall_policies — active_enhanced_firewall_policies / 301203122201 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- active_enhanced_firewall_policies

<a id="canonical-3010001212200013-3103033011100002-2000333033221323-0123000122301112-2213013122213213-3221100320333100-0312300110102221-1110022212213203"></a>

Type: `"single"`. Computed.

\[OneOf: active\_enhanced\_firewall\_policies, active\_network\_policies, disable\_network\_policy;
Default: disable\_network\_policy\] List of Enhanced Firewall Policies These policies use
session-based rules and provide all OPTIONS available under firewall policies with an additional
option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

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

- [active_enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-3010001212200013-3103033011100002-2000333033221323-0123000122301112-2213013122213213-3221100320333100-0312300110102221-1110022212213203)
- [active_network_policies](data-sources--network_firewall--reference--group-001.md#canonical-2312022030130130-1313313112011221-0322000121202111-3223331132101013-0110233220001103-0222121110100132-3231333220133311-2222122111200120)
- [disable_network_policy](data-sources--network_firewall--reference--group-001.md#canonical-3302010102221131-0122321003223010-2013230201020210-3202300102122221-2131302313232332-3220223110311301-1323030202322123-1031212320200102)

Select alternatives according to the provider validators above.

<a id="canonical-1022310201100322-2031310322223030-3313021231102223-1332020133021213-2001101220020032-0002322300002231-1101010220123323-1311021132111101"></a>

## Direct properties — active_enhanced_firewall_policies / 301203122201 / 3

- [enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-3232120202110122-0103303332221311-3133322012313022-2001320132310002-0103000322302203-0111110023300203-3120100000131102-2201023111032323): complete subsection reference.

<a id="canonical-2233132310302030-3231211202321010-1303123212023302-2301000221200221-3213200330233311-1021223132303103-2013223331122311-3133232003303031"></a>

## Next pages — active_enhanced_firewall_policies / 301203122201 / 4

- [active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-3232120202110122-0103303332221311-3133322012313022-2001320132310002-0103000322302203-0111110023300203-3120100000131102-2201023111032323)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)

<a id="canonical-3232120202110122-0103303332221311-3133322012313022-2001320132310002-0103000322302203-0111110023300203-3120100000131102-2201023111032323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132012012230230-2322021000023033-2310312323101200-2303001300032231-0212222201202131-1310103211321331-2321133112031131-2011031001233220"></a>

## active_enhanced_firewall_policies.enhanced_firewall_policies — enhanced_firewall_policies / 313113200032 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- [active_enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-3001312303222121-0312302310023012-1101130213031230-0333300122321021-0200212131012113-2223020201000110-1033221221230031-1120130013233030)
- active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-2030122330300332-1013312030022001-3011121110311132-1013222310030012-1131032211212012-2232102301102120-2031321103000310-1120310132132033"></a>

Type: `"list"`. Computed.

Ordered List of Enhanced Firewall Policies active.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2121311232330000-0020131033202120-3323323102333122-3021300002000010-3011203132100313-2121232211213003-2333303313321102-3213231023230311"></a>

## Direct properties — enhanced_firewall_policies / 313113200032 / 3

<a id="canonical-3311000101010123-3111211232101300-2313030112322102-3200322002333201-0133133010122013-1330233330333011-0212130120202012-2230000312221330"></a>

<a id="canonical-3013213032011310-3330120320111001-3213230133213033-3120313123322112-3102103110123030-2331330133210001-0112101020011301-3330123020301003"></a>

## name property — enhanced_firewall_policies / 313113200032 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0310100231102122-0100202131022002-3323110330013023-1313102211002212-1220310130223301-3230220320130110-0103001001331302-0233202312011112"></a>

<a id="canonical-2303002022103012-3120332323101012-3112132310123320-1220302002201320-3310332032033301-0023310030203301-1032133020032032-3003300331230003"></a>

## namespace property — enhanced_firewall_policies / 313113200032 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3223012312003313-1200310113232230-3210020322032100-0031102213222233-2003200321120031-1201122232030100-2031301233131220-0002222200101321"></a>

<a id="canonical-3201101203133322-1212030311222222-3320313102202201-1021212333111330-3310333113111312-2130223120123320-0010113301310200-3100200223020133"></a>

## tenant property — enhanced_firewall_policies / 313113200032 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0222020221323132-0013222322111001-3111032332321301-2332130200121230-3300203201012330-2012232321121030-3320111123222301-1332211012122111"></a>

## Next pages — enhanced_firewall_policies / 313113200032 / 7

- [active_enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-3001312303222121-0312302310023012-1101130213031230-0333300122321021-0200212131012113-2223020201000110-1033221221230031-1120130013233030)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)

<a id="canonical-3110031110222221-1202323102011100-3331003003330103-0303323021130010-3102212020312231-0021130321211312-1222212111102211-3101302030000102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032213221122213-3311313122003310-2032333101221312-1232100311030311-1110321102220121-1312113001200311-0321332231301123-0323002223000203"></a>

## active_fast_acls — active_fast_acls / 002110032132 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- active_fast_acls

<a id="canonical-1103332321100131-3213313220332330-1333303120101011-3303123101221111-2220002110002020-3322102102323023-2332222310121300-2211233210302321"></a>

Type: `"single"`. Computed.

\[OneOf: active\_fast\_acls, disable\_fast\_acl; Default: disable\_fast\_acl\] Configuration
parameter for active fast acls.

Upstream description:

List of Fast ACL(s).

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

- [active_fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-1103332321100131-3213313220332330-1333303120101011-3303123101221111-2220002110002020-3322102102323023-2332222310121300-2211233210302321)
- [disable_fast_acl](data-sources--network_firewall--reference--group-001.md#canonical-0030330022021222-2220030333321200-2102310032230000-0031332233021110-3313213232201220-1320232103113102-0123133032310213-2202213012112330)

Select alternatives according to the provider validators above.

<a id="canonical-2022313121032332-1131203022003201-3220203120210213-1103200221013111-0123100110232303-3010301030203000-1033120123031331-0321131303002010"></a>

## Direct properties — active_fast_acls / 002110032132 / 3

- [fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-2032023122200023-0312101130113022-0001301200120023-2132310321112211-2201113103122213-1320232201331222-1213313203313020-3012100313003031): complete subsection reference.

<a id="canonical-0120103000010211-2332210120011300-1320310130231231-1310133333110320-0112321300203223-3220310001033103-1333120301310122-2210200003000032"></a>

## Next pages — active_fast_acls / 002110032132 / 4

- [active_fast_acls.fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-2032023122200023-0312101130113022-0001301200120023-2132310321112211-2201113103122213-1320232201331222-1213313203313020-3012100313003031)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)

<a id="canonical-2032023122200023-0312101130113022-0001301200120023-2132310321112211-2201113103122213-1320232201331222-1213313203313020-3012100313003031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101303202131102-3030132032122022-2100021211212020-2123103330303220-2221100213131133-1123132130310000-2301031113023232-1021331110330131"></a>

## active_fast_acls.fast_acls — fast_acls / 221000132132 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- [active_fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-3110031110222221-1202323102011100-3331003003330103-0303323021130010-3102212020312231-0021130321211312-1222212111102211-3101302030000102)
- active_fast_acls.fast_acls

<a id="canonical-1200232010132232-2312321331120203-3102033320231322-2113332330131322-3311101300101230-3212300110120333-2012333231130011-0310112021223332"></a>

Type: `"list"`. Computed.

Ordered List of Fast ACL(s) active for this network firewall.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2132221330023003-3310303123330321-1103201011330333-0333132222211211-1132313011223103-0311200120302111-1000303002311102-2101010201100323"></a>

## Direct properties — fast_acls / 221000132132 / 3

<a id="canonical-1210103123232011-3232003233321311-1033103111133031-0212010323203231-1210220101311020-0012000223010100-1323131133202202-3323220123310303"></a>

<a id="canonical-1131303123030313-0210120222022231-0132020323231011-1110130303133122-1131313321033011-0222133032103101-0203120210022322-2232031121320213"></a>

## name property — fast_acls / 221000132132 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1221001122310132-2230002130310002-3313231210022322-2010020320333030-1001013321011003-0110201103011331-1232200303020120-3220130212210112"></a>

<a id="canonical-0131311202202123-1102121130131231-2312013021323112-1032021311332333-1231123110211110-1130000332302213-0020132321003133-2220110302013102"></a>

## namespace property — fast_acls / 221000132132 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0322011102011123-3111012112021002-2023030223123332-2111000223033321-1303031300003023-2110320312120232-1032133011132212-3330203202230122"></a>

<a id="canonical-1231120223210021-3303132220321032-3021102003333111-1202313300211211-2322310230300022-3301011120031232-1001330112120103-1202313123112133"></a>

## tenant property — fast_acls / 221000132132 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1111110301213312-2311010303022001-3232212023113031-0122110311212231-3210202332330103-3321320003131002-0310120302230200-0023123333322010"></a>

## Next pages — fast_acls / 221000132132 / 7

- [active_fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-3110031110222221-1202323102011100-3331003003330103-0303323021130010-3102212020312231-0021130321211312-1222212111102211-3101302030000102)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)

<a id="canonical-2233311113113112-1012303201230000-3131021231333032-0301110132312132-0000211131030031-0310103000022311-3312020012101120-3122100311021210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030120113023032-2001223001031332-0013131223310010-0103300201220013-1312230322213310-3023122120000313-2321120310232223-0002200201130230"></a>

## active_forward_proxy_policies — active_forward_proxy_policies / 223311311100 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- active_forward_proxy_policies

<a id="canonical-0021303122111230-1013000133332130-2332303332303323-0032131021331133-1033112010221033-1310233303321110-3022212132121331-0201120321023311"></a>

Type: `"single"`. Computed.

\[OneOf: active\_forward\_proxy\_policies, disable\_forward\_proxy\_policy; Default:
disable\_forward\_proxy\_policy\] Ordered List of Forward Proxy Policies active.

Upstream description:

Ordered List of Forward Proxy Policies active.

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

- [active_forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-0021303122111230-1013000133332130-2332303332303323-0032131021331133-1033112010221033-1310233303321110-3022212132121331-0201120321023311)
- [disable_forward_proxy_policy](data-sources--network_firewall--reference--group-001.md#canonical-3031232303002032-0120110310101132-2300001001200200-3010210033012120-0122203110223120-3232320032111012-0010213200020322-0101003230022311)

Select alternatives according to the provider validators above.

<a id="canonical-0301320023331031-3021210222130032-0322303121012113-0101300002111202-1013301012003220-0331221333133322-1132100210300202-3013310333102202"></a>

## Direct properties — active_forward_proxy_policies / 223311311100 / 3

- [forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-1031303212131310-3122313002101032-2313103003301312-0322332322221300-2100113032331323-2133031111131322-3030011132102010-1201010223100303): complete subsection reference.

<a id="canonical-0120023313110210-0121021121333223-3113220010222203-3133203021101030-1332111033021102-2312002032303231-1120031121231132-0021121223131323"></a>

## Next pages — active_forward_proxy_policies / 223311311100 / 4

- [active_forward_proxy_policies.forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-1031303212131310-3122313002101032-2313103003301312-0322332322221300-2100113032331323-2133031111131322-3030011132102010-1201010223100303)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)

<a id="canonical-1031303212131310-3122313002101032-2313103003301312-0322332322221300-2100113032331323-2133031111131322-3030011132102010-1201010223100303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212210303331021-2020322220222130-2230100032222332-0131220213200212-0300102102030211-3033030231031033-2120330001012030-3311003103013200"></a>

## active_forward_proxy_policies.forward_proxy_policies — forward_proxy_policies / 313002211210 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- [active_forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-2233311113113112-1012303201230000-3131021231333032-0301110132312132-0000211131030031-0310103000022311-3312020012101120-3122100311021210)
- active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-0032322211223110-1000022110210021-0012030231303203-3003130023131103-0211101102322000-2322121010101311-0233120122022030-2111121303030002"></a>

Type: `"list"`. Computed.

Ordered List of Forward Proxy Policies active.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0301130021021222-0202311313230000-2110203113222320-0311211202321301-1011220201001303-0002031031303102-2330310201201000-3001032232332330"></a>

## Direct properties — forward_proxy_policies / 313002211210 / 3

<a id="canonical-3113022330230020-0023030122200012-3320313000103312-3201301013130032-2121211101323320-2110023211201301-1211210112022001-2311130130303233"></a>

<a id="canonical-2321213033002310-0220113331232133-1023212113200201-0101003011021301-1021002332122000-0331012213021033-2320323003030203-3023000101232220"></a>

## name property — forward_proxy_policies / 313002211210 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0112201311213120-2213031313121311-1131010230213021-0302133322303130-1200103102311223-0200100100203300-2133103233331331-2311022123303131"></a>

<a id="canonical-1033233321301123-3022010210021320-2300103101230103-2232103320331122-0221200211300102-3201102001100102-2323110012303102-1021001202003033"></a>

## namespace property — forward_proxy_policies / 313002211210 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3301222101221320-3300113132020210-1212031031200002-1023321232220200-0032001002132313-2022301021320332-1203002203300112-2020111303121032"></a>

<a id="canonical-3123233113100223-3231111213200321-1312213220310320-0300100320033210-1120031300001201-1032000300123212-1032213023322132-2310030102032010"></a>

## tenant property — forward_proxy_policies / 313002211210 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2303021030210310-0211122123320132-2103320020321332-0222223321131201-3203302113021222-0210232222211112-3212131303000330-0133001333120101"></a>

## Next pages — forward_proxy_policies / 313002211210 / 7

- [active_forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-2233311113113112-1012303201230000-3131021231333032-0301110132312132-0000211131030031-0310103000022311-3312020012101120-3122100311021210)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)

<a id="canonical-0300113221000300-3323031321112301-2303200000031221-3220101113022012-2132110033002231-3323003110211123-0002220301010033-3322101121111123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102133331022223-0313213202331123-0333021031213002-3010312322111120-0123102022330232-2033131331133100-2111312212023030-3232311110200220"></a>

## active_network_policies — active_network_policies / 033213221032 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- active_network_policies

<a id="canonical-2312022030130130-1313313112011221-0322000121202111-3223331132101013-0110233220001103-0222121110100132-3231333220133311-2222122111200120"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

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

<a id="canonical-1331122132110011-0221323132321123-0201331213222001-3220113200032320-2212032333022122-0012020212302120-1033001221011310-2332323133210222"></a>

## Direct properties — active_network_policies / 033213221032 / 3

- [network_policies](data-sources--network_firewall--reference--group-001.md#canonical-2033112001120312-0311000031201000-1231213311321222-0303111013233010-2102012301121220-3120022013303003-3332201020313030-0102221203032121): complete subsection reference.

<a id="canonical-2001331001130232-3330112001311212-2101133231221210-3011321022232200-2311333313031023-1002302113233210-0120030321212113-3013311002100211"></a>

## Next pages — active_network_policies / 033213221032 / 4

- [active_network_policies.network_policies](data-sources--network_firewall--reference--group-001.md#canonical-2033112001120312-0311000031201000-1231213311321222-0303111013233010-2102012301121220-3120022013303003-3332201020313030-0102221203032121)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)

<a id="canonical-2033112001120312-0311000031201000-1231213311321222-0303111013233010-2102012301121220-3120022013303003-3332201020313030-0102221203032121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000322030213310-2000300010011311-2221033013003033-3230223302212000-0230110022102223-2020022230203110-3321312312202000-3301003233312100"></a>

## active_network_policies.network_policies — network_policies / 000232022102 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- [active_network_policies](data-sources--network_firewall--reference--group-001.md#canonical-0300113221000300-3323031321112301-2303200000031221-3220101113022012-2132110033002231-3323003110211123-0002220301010033-3322101121111123)
- active_network_policies.network_policies

<a id="canonical-3030213121113222-0011122101133131-3202020120012113-1022210131003322-1333032210121313-3033031010132233-0212320131332132-2021102211333323"></a>

Type: `"list"`. Computed.

Ordered List of Firewall Policies active for this network firewall.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1033120020021200-3230333201333333-3032112313312001-3120133300030031-2000313002320032-2310331111221033-3101321121233312-2011221120011223"></a>

## Direct properties — network_policies / 000232022102 / 3

<a id="canonical-2312301220200032-1003222313311302-0101101233100030-1122213223231332-3001211202130212-0103333211100312-2332122310320301-1200210101031232"></a>

<a id="canonical-2021000130011222-0023331103332010-2203331313232100-0310130012220200-1213120322310221-1212101102230303-0002302301020030-1002010301222323"></a>

## name property — network_policies / 000232022102 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3313121101020020-2313003301010222-2100320222100222-1130000212110110-1013021012311220-3202330300000121-2132130101002330-0233303113300130"></a>

<a id="canonical-1101002323330310-2322233111211302-0322000332021321-0011223311112210-2020022122031003-3003310013301003-3003211110003311-2313333022332222"></a>

## namespace property — network_policies / 000232022102 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3333301110223323-3001101032020032-2121000230001212-1100223130210021-1123221130303111-2102301002211031-1322320310333203-1211010021232210"></a>

<a id="canonical-2003200110322010-1022322012103331-0022213222123113-3133210131212230-2201302013021000-0122200013003301-3122010232023011-0102312212320002"></a>

## tenant property — network_policies / 000232022102 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3012132310033013-1113212301212321-2011113020020211-3131112200001023-3023113130022010-0123210301210222-3003032221122101-0102331200330301"></a>

## Next pages — network_policies / 000232022102 / 7

- [active_network_policies](data-sources--network_firewall--reference--group-001.md#canonical-0300113221000300-3323031321112301-2303200000031221-3220101113022012-2132110033002231-3323003110211123-0002220301010033-3322101121111123)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)

<a id="canonical-1112103131113333-1121212021221301-3111232301200132-0233211310020303-0203211310320133-1023300102123003-1030120300323032-1132232021223232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201301011022111-1210002330303313-2021220031130120-1213322023330200-3211033031012113-0303013020133101-2312222030322213-1312020131122010"></a>

## disable_fast_acl — disable_fast_acl / 313302332020 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- disable_fast_acl

<a id="canonical-0030330022021222-2220030333321200-2102310032230000-0031332233021110-3313213232201220-1320232103113102-0123133032310213-2202213012112330"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable fast acl. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-3221332112002132-2210011220123220-2123100101321231-1331132312003032-1103220223011312-3031222102200321-3131010133130311-1001132230232303"></a>

## Direct properties — disable_fast_acl / 313302332020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133232221012230-1033301232132300-3131001203320332-3110230023313030-1220121233211332-3320201310031322-0201003002023301-3213131202132103"></a>

## Next pages — disable_fast_acl / 313302332020 / 4

- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)

<a id="canonical-0100002122021020-1303013222121200-1101200333333311-2301022320231012-1200022022121321-1102131103022223-0200320323012013-2031001123113313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220033102203122-3103311101101312-2003130302032213-1131310121023213-0231233030311233-2130103330312132-3103221112202223-2231103120201133"></a>

## disable_forward_proxy_policy — disable_forward_proxy_policy / 320100022100 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- disable_forward_proxy_policy

<a id="canonical-3031232303002032-0120110310101132-2300001001200200-3010210033012120-0122203110223120-3232320032111012-0010213200020322-0101003230022311"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-3010231312202013-2032232003012300-2233232021120203-3312200330303212-1132011132113302-2103030103010122-2030210321031120-3220133200323103"></a>

## Direct properties — disable_forward_proxy_policy / 320100022100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101213313213121-1212100032212023-3330221220020300-3013313233210323-3131312131302302-3201021220223330-0131333120113030-3213323123001310"></a>

## Next pages — disable_forward_proxy_policy / 320100022100 / 4

- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)

<a id="canonical-0103010021311202-2013232200312121-3032223321111100-1223032112300103-2020000111230302-0012202131030220-3122212103110212-0113122131321023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132113000033020-1001000321032313-2111211003312312-0212022232021033-2312001112301221-0103300023223122-3100122122031322-0321321033223200"></a>

## disable_network_policy — disable_network_policy / 212321332102 / 2

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- disable_network_policy

<a id="canonical-3302010102221131-0122321003223010-2013230201020210-3202300102122221-2131302313232332-3220223110311301-1323030202322123-1031212320200102"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-2202310232123023-2003022233223320-1012321300022001-3013210130300200-3133121312011022-2003312323211310-0123103221200131-2210232113133220"></a>

## Direct properties — disable_network_policy / 212321332102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331310023131121-3033202123232333-2011023023113311-3213030210222311-1233231102102112-1313303113133111-3323221000311112-3112203013031210"></a>

## Next pages — disable_network_policy / 212321332102 / 4

- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
