---
page_title: "xcsh_network_firewall reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_firewall reference."
---

# xcsh_network_firewall reference

<a id="canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- Property reference

<a id="canonical-1003232020321302-0321221210232232-2310031332121032-3122300202320231-3230202312001213-2020033132000301-0110320230231223-3303011203222201"></a>

### Direct properties for `xcsh_network_firewall`

- [active_enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-3001312303222121-0312302310023012-1101130213031230-0333300122321021-0200212131012113-2223020201000110-1033221221230031-1120130013233030): complete subsection reference.

- [active_fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-3110031110222221-1202323102011100-3331003003330103-0303323021130010-3102212020312231-0021130321211312-1222212111102211-3101302030000102): complete subsection reference.

- [active_forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-2233311113113112-1012303201230000-3131021231333032-0301110132312132-0000211131030031-0310103000022311-3312020012101120-3122100311021210): complete subsection reference.

- [active_network_policies](data-sources--network_firewall--reference--group-001.md#canonical-0300113221000300-3323031321112301-2303200000031221-3220101113022012-2132110033002231-3323003110211123-0002220301010033-3322101121111123): complete subsection reference.

<a id="canonical-1323230202213112-1100020301302102-0012302231303012-2020310213223223-0022113300130230-0222012312020330-0313310130212210-0331013021233012"></a>

<a id="canonical-3011123213023232-1320322231202133-3121320320021123-0310303320202123-1210132120000002-2132021303202133-2311321003202121-3110101132022103"></a>

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

<a id="canonical-1200303010330021-2020100002120023-2333221131110312-3011011322131021-2003202312322123-3012322123033002-2033131022222301-3310233300133322"></a>

<a id="canonical-0310231110013010-2101230131021230-3103320131210221-2201100333100331-0133033201201223-0000312212101031-2220120323202113-0222031321103302"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the NetworkFirewall.

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

- [disable_fast_acl](data-sources--network_firewall--reference--group-001.md#canonical-1112103131113333-1121212021221301-3111232301200132-0233211310020303-0203211310320133-1023300102123003-1030120300323032-1132232021223232): complete subsection reference.

- [disable_forward_proxy_policy](data-sources--network_firewall--reference--group-001.md#canonical-0100002122021020-1303013222121200-1101200333333311-2301022320231012-1200022022121321-1102131103022223-0200320323012013-2031001123113313): complete subsection reference.

- [disable_network_policy](data-sources--network_firewall--reference--group-001.md#canonical-0103010021311202-2013232200312121-3032223321111100-1223032112300103-2020000111230302-0012202131030220-3122212103110212-0113122131321023): complete subsection reference.

<a id="canonical-3000113322102212-1321020203220123-2130012121001232-1330022032113012-3123230200111130-3301020210230101-3302220023210220-1020013322213301"></a>

<a id="canonical-2010202321123230-2313121111120333-1311011000303133-2123031100123113-0332020222030321-0010300223302111-2230110201003230-0120301223212113"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0203111130201313-3223232131322112-0211331311030000-3222303023102000-1000323110222233-1321102110321220-0032133133022212-0232031200000011"></a>

<a id="canonical-0212132023131030-0220012001230031-1232003300313032-1111302032233130-0321002100332233-3031213231033021-2312323330033220-3022101321030211"></a>

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

<a id="canonical-3310313130312031-2232000300023001-3132122301233110-2131203202302021-3033021123303313-0120310103233000-0000032112032111-2012331102123132"></a>

<a id="canonical-0032003023032120-2000000311232313-2331203003123300-0100022320113010-1300003331100321-1120003000133212-3310201212031203-3000231022102011"></a>

#### `name` property

Type: `"string"`. Required.

Name of the NetworkFirewall.

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

<a id="canonical-1100133023222232-3110011021200310-0011333111002033-2220030201320320-3011230130221023-3001220121320113-1013323302203302-0220031312232122"></a>

<a id="canonical-2331133012212320-1301001201330321-1012310103202003-1130021030220232-3030310013111331-3123122023233322-0220232013200021-3322233120000001"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace where the NetworkFirewall exists.

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

<a id="canonical-0031313002031013-2201322211210031-1320112032200220-3313002032213020-2300133330101202-1012202103331213-3331103220211100-1023303332100331"></a>

### All schema paths for `xcsh_network_firewall`

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

<a id="canonical-3001312303222121-0312302310023012-1101130213031230-0333300122321021-0200212131012113-2223020201000110-1033221221230031-1120130013233030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_enhanced_firewall_policies` properties

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

<a id="canonical-2323320013301010-0013331020131201-1132322130332301-2012132023021100-0021110302302023-2201332323002222-3031303000213031-3232232100311322"></a>

### Direct properties for `active_enhanced_firewall_policies`

- [enhanced_firewall_policies](data-sources--network_firewall--reference--group-001.md#canonical-3232120202110122-0103303332221311-3133322012313022-2001320132310002-0103000322302203-0111110023300203-3120100000131102-2201023111032323): complete subsection reference.

<a id="canonical-3232120202110122-0103303332221311-3133322012313022-2001320132310002-0103000322302203-0111110023300203-3120100000131102-2201023111032323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_enhanced_firewall_policies.enhanced_firewall_policies` properties

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

<a id="canonical-1132012012230230-2322021000023033-2310312323101200-2303001300032231-0212222201202131-1310103211321331-2321133112031131-2011031001233220"></a>

### Direct properties for `active_enhanced_firewall_policies.enhanced_firewall_policies`

<a id="canonical-3311000101010123-3111211232101300-2313030112322102-3200322002333201-0133133010122013-1330233330333011-0212130120202012-2230000312221330"></a>

#### `active_enhanced_firewall_policies.enhanced_firewall_policies.name` property

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

<a id="canonical-0310100231102122-0100202131022002-3323110330013023-1313102211002212-1220310130223301-3230220320130110-0103001001331302-0233202312011112"></a>

<a id="canonical-2121311232330000-0020131033202120-3323323102333122-3021300002000010-3011203132100313-2121232211213003-2333303313321102-3213231023230311"></a>

#### `active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` property

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

<a id="canonical-3223012312003313-1200310113232230-3210020322032100-0031102213222233-2003200321120031-1201122232030100-2031301233131220-0002222200101321"></a>

<a id="canonical-3013213032011310-3330120320111001-3213230133213033-3120313123322112-3102103110123030-2331330133210001-0112101020011301-3330123020301003"></a>

#### `active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` property

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

<a id="canonical-3110031110222221-1202323102011100-3331003003330103-0303323021130010-3102212020312231-0021130321211312-1222212111102211-3101302030000102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_fast_acls` properties

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- active_fast_acls

<a id="canonical-1103332321100131-3213313220332330-1333303120101011-3303123101221111-2220002110002020-3322102102323023-2332222310121300-2211233210302321"></a>

Type: `"single"`. Computed.

\[OneOf: active\_fast\_acls, disable\_fast\_acl; Default: disable\_fast\_acl\] Configuration
parameter for active fast acls.

Additional upstream details:

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

<a id="canonical-0032213221122213-3311313122003310-2032333101221312-1232100311030311-1110321102220121-1312113001200311-0321332231301123-0323002223000203"></a>

### Direct properties for `active_fast_acls`

- [fast_acls](data-sources--network_firewall--reference--group-001.md#canonical-2032023122200023-0312101130113022-0001301200120023-2132310321112211-2201113103122213-1320232201331222-1213313203313020-3012100313003031): complete subsection reference.

<a id="canonical-2032023122200023-0312101130113022-0001301200120023-2132310321112211-2201113103122213-1320232201331222-1213313203313020-3012100313003031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_fast_acls.fast_acls` properties

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

<a id="canonical-0101303202131102-3030132032122022-2100021211212020-2123103330303220-2221100213131133-1123132130310000-2301031113023232-1021331110330131"></a>

### Direct properties for `active_fast_acls.fast_acls`

<a id="canonical-1210103123232011-3232003233321311-1033103111133031-0212010323203231-1210220101311020-0012000223010100-1323131133202202-3323220123310303"></a>

#### `active_fast_acls.fast_acls.name` property

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

<a id="canonical-1221001122310132-2230002130310002-3313231210022322-2010020320333030-1001013321011003-0110201103011331-1232200303020120-3220130212210112"></a>

<a id="canonical-2132221330023003-3310303123330321-1103201011330333-0333132222211211-1132313011223103-0311200120302111-1000303002311102-2101010201100323"></a>

#### `active_fast_acls.fast_acls.namespace` property

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

<a id="canonical-0322011102011123-3111012112021002-2023030223123332-2111000223033321-1303031300003023-2110320312120232-1032133011132212-3330203202230122"></a>

<a id="canonical-1131303123030313-0210120222022231-0132020323231011-1110130303133122-1131313321033011-0222133032103101-0203120210022322-2232031121320213"></a>

#### `active_fast_acls.fast_acls.tenant` property

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

<a id="canonical-2233311113113112-1012303201230000-3131021231333032-0301110132312132-0000211131030031-0310103000022311-3312020012101120-3122100311021210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_forward_proxy_policies` properties

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- active_forward_proxy_policies

<a id="canonical-0021303122111230-1013000133332130-2332303332303323-0032131021331133-1033112010221033-1310233303321110-3022212132121331-0201120321023311"></a>

Type: `"single"`. Computed.

\[OneOf: active\_forward\_proxy\_policies, disable\_forward\_proxy\_policy; Default:
disable\_forward\_proxy\_policy\] Ordered List of Forward Proxy Policies active.

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

<a id="canonical-0030120113023032-2001223001031332-0013131223310010-0103300201220013-1312230322213310-3023122120000313-2321120310232223-0002200201130230"></a>

### Direct properties for `active_forward_proxy_policies`

- [forward_proxy_policies](data-sources--network_firewall--reference--group-001.md#canonical-1031303212131310-3122313002101032-2313103003301312-0322332322221300-2100113032331323-2133031111131322-3030011132102010-1201010223100303): complete subsection reference.

<a id="canonical-1031303212131310-3122313002101032-2313103003301312-0322332322221300-2100113032331323-2133031111131322-3030011132102010-1201010223100303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_forward_proxy_policies.forward_proxy_policies` properties

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

<a id="canonical-0212210303331021-2020322220222130-2230100032222332-0131220213200212-0300102102030211-3033030231031033-2120330001012030-3311003103013200"></a>

### Direct properties for `active_forward_proxy_policies.forward_proxy_policies`

<a id="canonical-3113022330230020-0023030122200012-3320313000103312-3201301013130032-2121211101323320-2110023211201301-1211210112022001-2311130130303233"></a>

#### `active_forward_proxy_policies.forward_proxy_policies.name` property

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

<a id="canonical-0112201311213120-2213031313121311-1131010230213021-0302133322303130-1200103102311223-0200100100203300-2133103233331331-2311022123303131"></a>

<a id="canonical-0301130021021222-0202311313230000-2110203113222320-0311211202321301-1011220201001303-0002031031303102-2330310201201000-3001032232332330"></a>

#### `active_forward_proxy_policies.forward_proxy_policies.namespace` property

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

<a id="canonical-3301222101221320-3300113132020210-1212031031200002-1023321232220200-0032001002132313-2022301021320332-1203002203300112-2020111303121032"></a>

<a id="canonical-2321213033002310-0220113331232133-1023212113200201-0101003011021301-1021002332122000-0331012213021033-2320323003030203-3023000101232220"></a>

#### `active_forward_proxy_policies.forward_proxy_policies.tenant` property

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

<a id="canonical-0300113221000300-3323031321112301-2303200000031221-3220101113022012-2132110033002231-3323003110211123-0002220301010033-3322101121111123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_network_policies` properties

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- active_network_policies

<a id="canonical-2312022030130130-1313313112011221-0322000121202111-3223331132101013-0110233220001103-0222121110100132-3231333220133311-2222122111200120"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Additional upstream details:

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

<a id="canonical-2102133331022223-0313213202331123-0333021031213002-3010312322111120-0123102022330232-2033131331133100-2111312212023030-3232311110200220"></a>

### Direct properties for `active_network_policies`

- [network_policies](data-sources--network_firewall--reference--group-001.md#canonical-2033112001120312-0311000031201000-1231213311321222-0303111013233010-2102012301121220-3120022013303003-3332201020313030-0102221203032121): complete subsection reference.

<a id="canonical-2033112001120312-0311000031201000-1231213311321222-0303111013233010-2102012301121220-3120022013303003-3332201020313030-0102221203032121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_network_policies.network_policies` properties

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

<a id="canonical-0000322030213310-2000300010011311-2221033013003033-3230223302212000-0230110022102223-2020022230203110-3321312312202000-3301003233312100"></a>

### Direct properties for `active_network_policies.network_policies`

<a id="canonical-2312301220200032-1003222313311302-0101101233100030-1122213223231332-3001211202130212-0103333211100312-2332122310320301-1200210101031232"></a>

#### `active_network_policies.network_policies.name` property

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

<a id="canonical-3313121101020020-2313003301010222-2100320222100222-1130000212110110-1013021012311220-3202330300000121-2132130101002330-0233303113300130"></a>

<a id="canonical-1033120020021200-3230333201333333-3032112313312001-3120133300030031-2000313002320032-2310331111221033-3101321121233312-2011221120011223"></a>

#### `active_network_policies.network_policies.namespace` property

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

<a id="canonical-3333301110223323-3001101032020032-2121000230001212-1100223130210021-1123221130303111-2102301002211031-1322320310333203-1211010021232210"></a>

<a id="canonical-2021000130011222-0023331103332010-2203331313232100-0310130012220200-1213120322310221-1212101102230303-0002302301020030-1002010301222323"></a>

#### `active_network_policies.network_policies.tenant` property

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

<a id="canonical-1112103131113333-1121212021221301-3111232301200132-0233211310020303-0203211310320133-1023300102123003-1030120300323032-1132232021223232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_fast_acl` properties

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- disable_fast_acl

<a id="canonical-0030330022021222-2220030333321200-2102310032230000-0031332233021110-3313213232201220-1320232103113102-0123133032310213-2202213012112330"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable fast acl. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-0100002122021020-1303013222121200-1101200333333311-2301022320231012-1200022022121321-1102131103022223-0200320323012013-2031001123113313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_forward_proxy_policy` properties

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- disable_forward_proxy_policy

<a id="canonical-3031232303002032-0120110310101132-2300001001200200-3010210033012120-0122203110223120-3232320032111012-0010213200020322-0101003230022311"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-0103010021311202-2013232200312121-3032223321111100-1223032112300103-2020000111230302-0012202131030220-3122212103110212-0113122131321023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_network_policy` properties

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Property reference](data-sources--network_firewall--reference--group-001.md#canonical-0313202030232210-1123113202121211-0313303312022021-2323032110030123-0111210231122331-2321110233331103-3022333201232103-2230011301013103)
- disable_network_policy

<a id="canonical-3302010102221131-0122321003223010-2013230201020210-3202300102122221-2131302313232332-3220223110311301-1323030202322123-1031212320200102"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature. Defaults to \`map\[\]\`. Server applies default when omitted.

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
