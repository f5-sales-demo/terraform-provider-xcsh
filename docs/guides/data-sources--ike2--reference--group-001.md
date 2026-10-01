---
page_title: "xcsh_ike2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike2 reference."
---

# xcsh_ike2 reference

<a id="canonical-2210220022220321-0032112331013302-1322232000331110-0332122122320233-1212003130332202-2130200012212213-0001200123130221-1102033222321323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323231013331021-2132001232003333-2021120002113010-2233001202121000-2233233121300132-3321001113313100-3011011323110101-2020030001332232"></a>

## Property reference — Property reference / 212112310003 / 2

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md#canonical-1332033132212110-1332133210322000-0321030302123031-3010310020103003-2122221201132213-2201311321232022-0331322230010032-2233221102111032)
- Property reference

<a id="canonical-1330110120332003-1001210232230231-3100113121113213-3133320022030211-3211203223023330-3100002022011000-0002330223213103-2001023021332201"></a>

## Direct properties — Property reference / 212112310003 / 3

<a id="canonical-3000230003322203-3132030203101033-0200213332012303-0101202121223131-0331300221233323-1203230310220331-2233313300220110-0222011222313201"></a>

<a id="canonical-3203203223120100-2031003100022202-3310222110033023-3230130212302220-1303222023111333-3002111330111301-2232311310003020-2300010302231103"></a>

## annotations property — Property reference / 212112310003 / 4

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

<a id="canonical-1120331131222331-1301333102323200-0332200321023123-2120311122100030-1032210132100310-3103033002302210-2103201120300332-2232100013301022"></a>

<a id="canonical-2131120123133121-1120101202202111-3133222331100032-2022331011002030-1020031201133003-3312112022132322-2002210220232223-1330020000232212"></a>

## description property — Property reference / 212112310003 / 5

Type: `"string"`. Computed.

Description of the Ike2.

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

- [dh_group_set](data-sources--ike2--reference--group-001.md#canonical-2321302020002112-0310322221000012-2331102322022133-1132023021030113-2123220320303032-3322223132200331-2110031120023331-3313012221120203): complete subsection reference.

- [disable_pfs](data-sources--ike2--reference--group-001.md#canonical-1312221221200230-3200223001010023-3010122220033323-0221202201231312-0021311201030213-1002331312323030-2031223131210032-2313331111003023): complete subsection reference.

<a id="canonical-2300222132132130-2032230011321021-0132131123312032-0323312001330220-3110023212113201-0301001202332120-0202132133121331-2323113233201011"></a>

<a id="canonical-2303332101312201-2310200102332210-1102013223321213-3022030032122322-3122130332001302-1101011220223233-3003113021012012-1011302022233021"></a>

## ID property — Property reference / 212112310003 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](data-sources--ike2--reference--group-001.md#canonical-0120220320102313-2020033132300233-1201001110323033-0131000220102200-1213302223130331-2312110121023222-3122313002230322-1232021113223112): complete subsection reference.

- [ike_keylifetime_minutes](data-sources--ike2--reference--group-001.md#canonical-2003210323203303-0200300230323221-0022102111032231-0230131102331120-1130330001211122-3111022233310330-2221231230000322-3023330222232103): complete subsection reference.

<a id="canonical-0100100020331001-1012132213213130-1310230123222333-1232203313311220-1130321231212233-3103201203031032-1332123002131212-2013232320202223"></a>

<a id="canonical-3021222003103131-2313013323003012-0011133103232233-1101313013033320-2322313002001111-3213311132322300-2130130321110013-3132113321320001"></a>

## labels property — Property reference / 212112310003 / 7

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

<a id="canonical-0132120200030202-0020221021112231-1103221021020232-2033021113213321-1231020221101110-3211120111231032-0332331310120010-3311223130001331"></a>

<a id="canonical-3000132011323313-2230131022210122-2300011020302312-3211033201233000-2032221021223210-1123121030322300-1122132023001200-2133332021121130"></a>

## name property — Property reference / 212112310003 / 8

Type: `"string"`. Required.

Name of the Ike2.

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

<a id="canonical-1200232301003223-3323323132322321-3301301213112113-2311112200313313-3320311111322120-2222231321213332-3321231111302331-0122001032230133"></a>

<a id="canonical-1110223233032222-2013332112010313-3210132012232002-0333002101133301-2231312330302021-0132230223202232-2210332322310130-3233013221012031"></a>

## namespace property — Property reference / 212112310003 / 9

Type: `"string"`. Required.

Namespace where the Ike2 exists.

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

- [use_default_keylifetime](data-sources--ike2--reference--group-001.md#canonical-2101213102331132-2011013013211220-0230123122131310-2111331103302100-3210100221233301-1020123312002000-2332110321121230-3121310012010332): complete subsection reference.

<a id="canonical-1320111102202030-2032312012010301-3232130033122313-1312003030300032-3010112302003030-1313202233120221-3323113210330321-3212001010021102"></a>

## All schema paths — Property reference / 212112310003 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--ike2--reference--group-001.md#canonical-3000230003322203-3132030203101033-0200213332012303-0101202121223131-0331300221233323-1203230310220331-2233313300220110-0222011222313201) |
| `description` | [description](data-sources--ike2--reference--group-001.md#canonical-1120331131222331-1301333102323200-0332200321023123-2120311122100030-1032210132100310-3103033002302210-2103201120300332-2232100013301022) |
| `dh_group_set` | [dh_group_set](data-sources--ike2--reference--group-001.md#canonical-3300222300132212-1220030123211210-1220122132100331-0301113203002131-2221310101231012-0233003023121103-1101001231120022-3013011200031131) |
| `dh_group_set.dh_groups` | [dh_group_set.dh_groups](data-sources--ike2--reference--group-001.md#canonical-0210123000131300-1321200301323131-3000023302120121-2212130131321202-0320002112130121-1023313203120103-0300201020213003-0102300132321131) |
| `disable_pfs` | [disable_pfs](data-sources--ike2--reference--group-001.md#canonical-3110111100332303-1001023323331233-3311001100131033-0113231023123322-3223322120303322-3030132000300232-2122110131312120-2232122023122301) |
| `id` | [ID](data-sources--ike2--reference--group-001.md#canonical-2300222132132130-2032230011321021-0132131123312032-0323312001330220-3110023212113201-0301001202332120-0202132133121331-2323113233201011) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](data-sources--ike2--reference--group-001.md#canonical-3012000233311330-1031112333001132-2121013100311122-0303330212330000-0021200110303013-2230111202323221-3022200221121233-0220131130033120) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](data-sources--ike2--reference--group-001.md#canonical-1310320332211103-1231123010133223-1133201333312311-1103120210020312-1031223032123022-1001321130312331-3310023322333020-1120003230233030) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](data-sources--ike2--reference--group-001.md#canonical-1310213202221322-1130003032310011-2101110312012212-2301213223122320-0302200013233131-1033311230031012-3110013202011000-1230132133301023) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](data-sources--ike2--reference--group-001.md#canonical-1130311131003001-2202202132130303-2233222310320311-2223020130302013-3322322011301010-2302320312232203-1211012003223010-2111203221313302) |
| `labels` | [labels](data-sources--ike2--reference--group-001.md#canonical-0100100020331001-1012132213213130-1310230123222333-1232203313311220-1130321231212233-3103201203031032-1332123002131212-2013232320202223) |
| `name` | [name](data-sources--ike2--reference--group-001.md#canonical-0132120200030202-0020221021112231-1103221021020232-2033021113213321-1231020221101110-3211120111231032-0332331310120010-3311223130001331) |
| `namespace` | [namespace](data-sources--ike2--reference--group-001.md#canonical-1200232301003223-3323323132322321-3301301213112113-2311112200313313-3320311111322120-2222231321213332-3321231111302331-0122001032230133) |
| `use_default_keylifetime` | [use_default_keylifetime](data-sources--ike2--reference--group-001.md#canonical-0102210300122110-3233033220032231-1101113333302032-0111113211200313-2321123232331312-3102212302332331-2212233203110022-1331131231031231) |

<a id="canonical-1113221120012131-3100032031022310-2223002331111113-0210011022003120-2230030222021320-2313132023303130-3321032011112021-2110023313123110"></a>

## Next pages — Property reference / 212112310003 / 11

- [dh_group_set](data-sources--ike2--reference--group-001.md#canonical-2321302020002112-0310322221000012-2331102322022133-1132023021030113-2123220320303032-3322223132200331-2110031120023331-3313012221120203)
- [disable_pfs](data-sources--ike2--reference--group-001.md#canonical-1312221221200230-3200223001010023-3010122220033323-0221202201231312-0021311201030213-1002331312323030-2031223131210032-2313331111003023)
- [ike_keylifetime_hours](data-sources--ike2--reference--group-001.md#canonical-0120220320102313-2020033132300233-1201001110323033-0131000220102200-1213302223130331-2312110121023222-3122313002230322-1232021113223112)
- [ike_keylifetime_minutes](data-sources--ike2--reference--group-001.md#canonical-2003210323203303-0200300230323221-0022102111032231-0230131102331120-1130330001211122-3111022233310330-2221231230000322-3023330222232103)
- [use_default_keylifetime](data-sources--ike2--reference--group-001.md#canonical-2101213102331132-2011013013211220-0230123122131310-2111331103302100-3210100221233301-1020123312002000-2332110321121230-3121310012010332)
- [xcsh_ike2](../data-sources/ike2.md#canonical-1332033132212110-1332133210322000-0321030302123031-3010310020103003-2122221201132213-2201311321232022-0331322230010032-2233221102111032)

<a id="canonical-2321302020002112-0310322221000012-2331102322022133-1132023021030113-2123220320303032-3322223132200331-2110031120023331-3313012221120203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013023333101003-3100101310313301-0322232322320302-0122020311213300-1200323221311212-3110212310102333-1223132013002302-0100032300221312"></a>

## dh_group_set — dh_group_set / 002003223030 / 2

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md#canonical-1332033132212110-1332133210322000-0321030302123031-3010310020103003-2122221201132213-2201311321232022-0331322230010032-2233221102111032)
- [Property reference](data-sources--ike2--reference--group-001.md#canonical-2210220022220321-0032112331013302-1322232000331110-0332122122320233-1212003130332202-2130200012212213-0001200123130221-1102033222321323)
- dh_group_set

<a id="canonical-3300222300132212-1220030123211210-1220122132100331-0301113203002131-2221310101231012-0233003023121103-1101001231120022-3013011200031131"></a>

Type: `"single"`. Computed.

\[OneOf: dh\_group\_set, disable\_pfs; Default: disable\_pfs\] Choose the acceptable Diffie
Hellman(DH) Group or Groups that you are willing to accept as part of this profile.

Upstream description:

Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of
this profile.

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

- [dh_group_set](data-sources--ike2--reference--group-001.md#canonical-3300222300132212-1220030123211210-1220122132100331-0301113203002131-2221310101231012-0233003023121103-1101001231120022-3013011200031131)
- [disable_pfs](data-sources--ike2--reference--group-001.md#canonical-3110111100332303-1001023323331233-3311001100131033-0113231023123322-3223322120303322-3030132000300232-2122110131312120-2232122023122301)

Select alternatives according to the provider validators above.

<a id="canonical-2001201311020333-3103213130112033-0310110001122011-2213030211201203-1113212300212010-1221001020110211-2300213133222221-0230303033022212"></a>

## Direct properties — dh_group_set / 002003223030 / 3

<a id="canonical-0210123000131300-1321200301323131-3000023302120121-2212130131321202-0320002112130121-1023313203120103-0300201020213003-0102300132321131"></a>

<a id="canonical-0022231033232131-3103022103010013-2310300303330233-0323202301111110-0120232120011310-3003133122310201-2031210311101302-1120010331221302"></a>

## dh_groups property — dh_group_set / 002003223030 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
DH\_GROUP\_DEFAULT|DH\_GROUP\_14|DH\_GROUP\_15|DH\_GROUP\_16|DH\_GROUP\_17|DH\_GROUP\_18|DH\_GROUP\_19|DH\_GROUP\_20|DH\_GROUP\_21|DH\_GROUP\_26\]
Diffie Hellman Groups. Group or collection configuration. Possible values are
\`DH\_GROUP\_DEFAULT\`, \`DH\_GROUP\_14\`, \`DH\_GROUP\_15\`, \`DH\_GROUP\_16\`, \`DH\_GROUP\_17\`,
\`DH\_GROUP\_18\`, \`DH\_GROUP\_19\`, \`DH\_GROUP\_20\`, \`DH\_GROUP\_21\`, \`DH\_GROUP\_26\`.
Defaults to \`DH\_GROUP\_DEFAULT\`.

Upstream description:

Group or collection configuration

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

<a id="canonical-0332103132123031-1130221001323310-0301101303032012-3123101111023303-2232311333301031-3120313121030132-2321020231303111-2031003030120203"></a>

## Next pages — dh_group_set / 002003223030 / 5

- [Property reference](data-sources--ike2--reference--group-001.md#canonical-2210220022220321-0032112331013302-1322232000331110-0332122122320233-1212003130332202-2130200012212213-0001200123130221-1102033222321323)
- [xcsh_ike2](../data-sources/ike2.md#canonical-1332033132212110-1332133210322000-0321030302123031-3010310020103003-2122221201132213-2201311321232022-0331322230010032-2233221102111032)

<a id="canonical-1312221221200230-3200223001010023-3010122220033323-0221202201231312-0021311201030213-1002331312323030-2031223131210032-2313331111003023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312103100132121-0202211210203313-2033013332300200-2333232210202020-1331313120330222-2121021020020010-3220301210302221-3202302200113010"></a>

## disable_pfs — disable_pfs / 222231220130 / 2

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md#canonical-1332033132212110-1332133210322000-0321030302123031-3010310020103003-2122221201132213-2201311321232022-0331322230010032-2233221102111032)
- [Property reference](data-sources--ike2--reference--group-001.md#canonical-2210220022220321-0032112331013302-1322232000331110-0332122122320233-1212003130332202-2130200012212213-0001200123130221-1102033222321323)
- disable_pfs

<a id="canonical-3110111100332303-1001023323331233-3311001100131033-0113231023123322-3223322120303322-3030132000300232-2122110131312120-2232122023122301"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable pfs.

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

<a id="canonical-2021112332121230-1110221121332112-3000302202230210-1310100322330312-2211120200221323-1110102100032102-3310020112201332-3130322223102123"></a>

## Direct properties — disable_pfs / 222231220130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0121123122223023-2233132232222112-1331113220020212-2220331301232202-0100233030331312-2001011210331320-1323011011033030-3203021303003130"></a>

## Next pages — disable_pfs / 222231220130 / 4

- [Property reference](data-sources--ike2--reference--group-001.md#canonical-2210220022220321-0032112331013302-1322232000331110-0332122122320233-1212003130332202-2130200012212213-0001200123130221-1102033222321323)
- [xcsh_ike2](../data-sources/ike2.md#canonical-1332033132212110-1332133210322000-0321030302123031-3010310020103003-2122221201132213-2201311321232022-0331322230010032-2233221102111032)

<a id="canonical-0120220320102313-2020033132300233-1201001110323033-0131000220102200-1213302223130331-2312110121023222-3122313002230322-1232021113223112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111113302200031-0011101110233211-2222323321313032-0302000333203330-2202211201210121-1331002112332212-0233013333322011-2112013101001031"></a>

## ike_keylifetime_hours — ike_keylifetime_hours / 322302121020 / 2

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md#canonical-1332033132212110-1332133210322000-0321030302123031-3010310020103003-2122221201132213-2201311321232022-0331322230010032-2233221102111032)
- [Property reference](data-sources--ike2--reference--group-001.md#canonical-2210220022220321-0032112331013302-1322232000331110-0332122122320233-1212003130332202-2130200012212213-0001200123130221-1102033222321323)
- ike_keylifetime_hours

<a id="canonical-3012000233311330-1031112333001132-2121013100311122-0303330212330000-0021200110303013-2230111202323221-3022200221121233-0220131130033120"></a>

Type: `"single"`. Computed.

\[OneOf: ike\_keylifetime\_hours, ike\_keylifetime\_minutes, use\_default\_keylifetime; Default:
use\_default\_keylifetime\] Configuration parameter for ike keylifetime hours.

Upstream description:

Input Hours.

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

- [ike_keylifetime_hours](data-sources--ike2--reference--group-001.md#canonical-3012000233311330-1031112333001132-2121013100311122-0303330212330000-0021200110303013-2230111202323221-3022200221121233-0220131130033120)
- [ike_keylifetime_minutes](data-sources--ike2--reference--group-001.md#canonical-1310213202221322-1130003032310011-2101110312012212-2301213223122320-0302200013233131-1033311230031012-3110013202011000-1230132133301023)
- [use_default_keylifetime](data-sources--ike2--reference--group-001.md#canonical-0102210300122110-3233033220032231-1101113333302032-0111113211200313-2321123232331312-3102212302332331-2212233203110022-1331131231031231)

Select alternatives according to the provider validators above.

<a id="canonical-1002200230000123-3210133201002000-2113121210201020-2102030201000201-3102301131203202-0313103103102101-2221211221201131-3110320100303213"></a>

## Direct properties — ike_keylifetime_hours / 322302121020 / 3

<a id="canonical-1310320332211103-1231123010133223-1133201333312311-1103120210020312-1031223032123022-1001321130312331-3310023322333020-1120003230233030"></a>

<a id="canonical-3210301320233312-0131122110103000-0012132311113103-3023331221332011-0133323110330131-2030112331013030-2120003333031002-0213220022331122"></a>

## duration property — ike_keylifetime_hours / 322302121020 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-0103202122331123-3032120231032021-2130103201002032-0000333011031220-0002132031221023-2223232002110023-2232223211023110-3110003023010012"></a>

## Next pages — ike_keylifetime_hours / 322302121020 / 5

- [Property reference](data-sources--ike2--reference--group-001.md#canonical-2210220022220321-0032112331013302-1322232000331110-0332122122320233-1212003130332202-2130200012212213-0001200123130221-1102033222321323)
- [xcsh_ike2](../data-sources/ike2.md#canonical-1332033132212110-1332133210322000-0321030302123031-3010310020103003-2122221201132213-2201311321232022-0331322230010032-2233221102111032)

<a id="canonical-2003210323203303-0200300230323221-0022102111032231-0230131102331120-1130330001211122-3111022233310330-2221231230000322-3023330222232103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133232111003112-1302200212101122-1120132211021003-2022013231123232-3003210001223013-1103020012233223-0322223032121210-3100212312230313"></a>

## ike_keylifetime_minutes — ike_keylifetime_minutes / 320130002031 / 2

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md#canonical-1332033132212110-1332133210322000-0321030302123031-3010310020103003-2122221201132213-2201311321232022-0331322230010032-2233221102111032)
- [Property reference](data-sources--ike2--reference--group-001.md#canonical-2210220022220321-0032112331013302-1322232000331110-0332122122320233-1212003130332202-2130200012212213-0001200123130221-1102033222321323)
- ike_keylifetime_minutes

<a id="canonical-1310213202221322-1130003032310011-2101110312012212-2301213223122320-0302200013233131-1033311230031012-3110013202011000-1230132133301023"></a>

Type: `"single"`. Computed.

Configuration parameter for ike keylifetime minutes.

Upstream description:

Set IKE Key Lifetime in minutes.

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

<a id="canonical-1302212031221101-3001331232323112-3301132331033200-0331101112122210-2003032302222133-3302311203111212-3121133310000112-1220212013323123"></a>

## Direct properties — ike_keylifetime_minutes / 320130002031 / 3

<a id="canonical-1130311131003001-2202202132130303-2233222310320311-2223020130302013-3322322011301010-2302320312232203-1211012003223010-2111203221313302"></a>

<a id="canonical-0013320230111032-1233032010001310-1203031222331113-2032001202230003-0202130022332330-2220102113230320-2010323022131011-3023101203212131"></a>

## duration property — ike_keylifetime_minutes / 320130002031 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 10
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-0002220321311332-1003212220321200-2121232023331333-3202230003310311-0333320002201110-2120301032212122-1122102020032133-1132013313230330"></a>

## Next pages — ike_keylifetime_minutes / 320130002031 / 5

- [Property reference](data-sources--ike2--reference--group-001.md#canonical-2210220022220321-0032112331013302-1322232000331110-0332122122320233-1212003130332202-2130200012212213-0001200123130221-1102033222321323)
- [xcsh_ike2](../data-sources/ike2.md#canonical-1332033132212110-1332133210322000-0321030302123031-3010310020103003-2122221201132213-2201311321232022-0331322230010032-2233221102111032)

<a id="canonical-2101213102331132-2011013013211220-0230123122131310-2111331103302100-3210100221233301-1020123312002000-2332110321121230-3121310012010332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101320011023122-2212001212113130-1133203331302022-3200130222323303-1123200310200221-1022320133123122-3123002020320210-1301022220111210"></a>

## use_default_keylifetime — use_default_keylifetime / 302111022032 / 2

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md#canonical-1332033132212110-1332133210322000-0321030302123031-3010310020103003-2122221201132213-2201311321232022-0331322230010032-2233221102111032)
- [Property reference](data-sources--ike2--reference--group-001.md#canonical-2210220022220321-0032112331013302-1322232000331110-0332122122320233-1212003130332202-2130200012212213-0001200123130221-1102033222321323)
- use_default_keylifetime

<a id="canonical-0102210300122110-3233033220032231-1101113333302032-0111113211200313-2321123232331312-3102212302332331-2212233203110022-1331131231031231"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use default keylifetime.

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

<a id="canonical-0232203131001121-3212030220103302-1021331030233333-0231133122130132-3323210000021101-0022200112031122-0223212312022312-3320210100002202"></a>

## Direct properties — use_default_keylifetime / 302111022032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100300132132130-1230001022203220-2301032321032202-1212000301013001-3311113002132020-1302002303231021-1023221332010130-3112230312303113"></a>

## Next pages — use_default_keylifetime / 302111022032 / 4

- [Property reference](data-sources--ike2--reference--group-001.md#canonical-2210220022220321-0032112331013302-1322232000331110-0332122122320233-1212003130332202-2130200012212213-0001200123130221-1102033222321323)
- [xcsh_ike2](../data-sources/ike2.md#canonical-1332033132212110-1332133210322000-0321030302123031-3010310020103003-2122221201132213-2201311321232022-0331322230010032-2233221102111032)
