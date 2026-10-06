---
page_title: "xcsh_site_mesh_group reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site_mesh_group reference."
---

# xcsh_site_mesh_group reference

<a id="canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- Property reference

<a id="canonical-1220110130102321-3233301021201011-1032020223022230-0202020233120332-3020111303100333-1301023001333300-3332022113222330-0033123232011103"></a>

### Direct properties for `xcsh_site_mesh_group`

<a id="canonical-1131222011310230-3003323000001233-1302023002030332-1131111200110222-2003130231201123-0332312121133211-2333231310322300-1221331313130221"></a>

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

- [bfd_disabled](resources--site_mesh_group--reference--group-001.md#canonical-1011232323113121-1231320202132111-0231022231120310-2210321103313102-0131133032201310-0313033331122131-2021033033022030-0101021032123202): complete subsection reference.

- [bfd_enabled](resources--site_mesh_group--reference--group-001.md#canonical-0030312132232112-2303110333223203-2033302223131130-2123310023033230-1313103013010103-0011023222103320-3330033300112032-0231330313022102): complete subsection reference.

<a id="canonical-1032002310300133-1023311111200210-3322103002021103-1032131321103123-2012101330030121-0131121321200313-1300003121222133-1301303202002210"></a>

<a id="canonical-0113030111302322-1321112033122013-1232221002121313-3000103131000220-2221130203010012-2022222133131102-2213001302130001-0333303311201202"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0003111130311220-3322113321221310-3222102032210330-3111131301013220-0023222231023001-0322323030213313-3223102333200210-1010011022021312"></a>

<a id="canonical-2032112121320001-1203223111111113-3002221001112111-3013003032201230-0311023211100003-1301010120220233-1031213210232131-0012220010230130"></a>

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

- [disable_re_fallback](resources--site_mesh_group--reference--group-001.md#canonical-2300102311310103-1132033112330111-0232101032131033-2010201232322220-2230101212022313-0030303231303133-2111320001222103-2322220330311130): complete subsection reference.

- [enable_re_fallback](resources--site_mesh_group--reference--group-001.md#canonical-3210020200220232-0120031130020021-3022232122303201-3202302300231332-2221213031003201-0111003122200303-1332212132102212-1001102111323201): complete subsection reference.

- [full_mesh](resources--site_mesh_group--reference--group-001.md#canonical-2031001210322023-1032211110103322-2220103113201122-2113312202301300-3130323130310120-1212022100320032-0000210211323322-1333323000112211): complete subsection reference.

- [hub_mesh](resources--site_mesh_group--reference--group-001.md#canonical-1011133001312013-2312132233110312-1332302222102230-1122130331221011-0121321303130030-0111233333133330-2112011323002320-2210320323220112): complete subsection reference.

<a id="canonical-0320111022102011-2330001121322022-3022331302311030-3230023330200120-3121131323222333-1023323200023112-3221102303023121-0302221230011012"></a>

<a id="canonical-1332020200212133-0011323120211002-1023202133022201-2112230133203332-2020001003320200-1021311133200022-3333010200203003-0131031303331032"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3330331331201030-1310031200101101-3112132002001330-3200103203121100-1023210133031203-0111110023301022-2020030212131003-2323310230111231"></a>

<a id="canonical-0123112211232122-3331320133000010-3231232132212033-2103301103122121-3201321203303210-1233330312131020-0323003220103112-1010323121120003"></a>

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

<a id="canonical-3111013230113021-3310121112132001-0133211110322003-3121003210021231-3032221002333111-2122033233233102-1131110011031223-3233200031330333"></a>

<a id="canonical-1312001023022220-1212133122021320-0130012023002310-0310021320211303-2323012131033102-1123122313022030-2110011123111001-3221330333122003"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Site Mesh Group. Must be unique within the namespace.

Additional upstream details:

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0321030010210132-0130123112013111-2103300130020021-3001231032033220-3031133012313300-2131210332000113-0310230200332131-1331300131212323"></a>

<a id="canonical-1033323000233213-1301003021302202-3312203100121001-0031100233230233-0311113323100021-0113022130312120-1131220211323123-0102032111303320"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Site Mesh Group is created.

Additional upstream details:

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [spoke_mesh](resources--site_mesh_group--reference--group-001.md#canonical-2020322201021122-0131011313222121-0101022011301302-3002331111101232-3230013331330023-0032312121112300-2322110320130302-2012302020131121): complete subsection reference.

- [timeouts](resources--site_mesh_group--reference--group-001.md#canonical-0011100231203201-0101031320333120-1133233320032200-1122201010121132-2120012103121301-1030131300330330-1202223111001323-0213302002311313): complete subsection reference.

- [virtual_site](resources--site_mesh_group--reference--group-001.md#canonical-1220313230311013-1302220301022233-1002221232313132-0312222012013021-1302112330300313-3232013033020322-3011101331113203-2312331132230000): complete subsection reference.

<a id="canonical-3121223211211113-1320130320111333-1113113023322022-2032223001020231-3110100322200201-3103220121022303-2131330120221301-2110212320210202"></a>

### All schema paths for `xcsh_site_mesh_group`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--site_mesh_group--reference--group-001.md#canonical-1131222011310230-3003323000001233-1302023002030332-1131111200110222-2003130231201123-0332312121133211-2333231310322300-1221331313130221) |
| `bfd_disabled` | [bfd_disabled](resources--site_mesh_group--reference--group-001.md#canonical-0101313111020322-2113120031302131-1332220223000312-1131211102201331-2221111101132222-1331003201020030-3122022132321201-2301032211320230) |
| `bfd_enabled` | [bfd_enabled](resources--site_mesh_group--reference--group-001.md#canonical-2321013203133110-3011233023021013-0220101011031001-1201120002111200-0023133300322211-1130312000103211-3313233103203200-1303211011012003) |
| `bfd_enabled.multiplier` | [bfd_enabled.multiplier](resources--site_mesh_group--reference--group-001.md#canonical-1011210210120003-2010200031331221-3130230210210311-2003011112003330-2233231123120310-3220102331333100-2111011031110011-3120223123312222) |
| `bfd_enabled.receive_interval_milliseconds` | [bfd_enabled.receive_interval_milliseconds](resources--site_mesh_group--reference--group-001.md#canonical-2322210323101230-0000303010023331-2321132202013001-3301110321021310-3311012300011301-0231020131311023-3210103213123113-2103022320321231) |
| `bfd_enabled.transmit_interval_milliseconds` | [bfd_enabled.transmit_interval_milliseconds](resources--site_mesh_group--reference--group-001.md#canonical-2000123332033210-2113220001230010-0212011233011112-0202202113103313-3321121023300201-3233231212020111-0231120320220010-3223100311303133) |
| `description` | [description](resources--site_mesh_group--reference--group-001.md#canonical-1032002310300133-1023311111200210-3322103002021103-1032131321103123-2012101330030121-0131121321200313-1300003121222133-1301303202002210) |
| `disable` | [disable](resources--site_mesh_group--reference--group-001.md#canonical-0003111130311220-3322113321221310-3222102032210330-3111131301013220-0023222231023001-0322323030213313-3223102333200210-1010011022021312) |
| `disable_re_fallback` | [disable_re_fallback](resources--site_mesh_group--reference--group-001.md#canonical-1211023223123110-2301031233302033-0331221310230020-1113230000221013-0012311000330321-3131020312011232-3102110311301020-3111103122320201) |
| `enable_re_fallback` | [enable_re_fallback](resources--site_mesh_group--reference--group-001.md#canonical-1131110120030302-2220303333132110-1002133212021210-3303333010122321-0132310030210303-2030310233233102-1002323332113323-0300211120333310) |
| `full_mesh` | [full_mesh](resources--site_mesh_group--reference--group-001.md#canonical-0122310323131300-1222303131332010-2033032130111011-2003113221301311-3230212121213220-2210131302032320-3212330202102001-1111201213102310) |
| `full_mesh.control_and_data_plane_mesh` | [full_mesh.control_and_data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-3002031122222220-0201203221211030-1031202103123113-1012100300133331-2211012321101130-3310103332321200-1123220032001022-2230232310333121) |
| `full_mesh.data_plane_mesh` | [full_mesh.data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-2031333110320310-3222032001301300-1222102032313131-0203312013313102-1222013033022033-2103002201202111-3033133200113100-3100202231011023) |
| `hub_mesh` | [hub_mesh](resources--site_mesh_group--reference--group-001.md#canonical-3211112213301131-2132020103310313-1103001123230200-2021232113010023-3303132200003003-2010310110012030-2230103121320323-3330201233210110) |
| `hub_mesh.control_and_data_plane_mesh` | [hub_mesh.control_and_data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-1120030130300021-1132111233003120-1203333131232132-1311120223023032-3003003231301223-0301310321130112-1311323001312121-0110212000022100) |
| `hub_mesh.data_plane_mesh` | [hub_mesh.data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-0231011202033320-2103113322232103-0122131313020233-1110020112121123-3103102101112022-0030222212133000-0311120232030313-1222131202032321) |
| `id` | [ID](resources--site_mesh_group--reference--group-001.md#canonical-0320111022102011-2330001121322022-3022331302311030-3230023330200120-3121131323222333-1023323200023112-3221102303023121-0302221230011012) |
| `labels` | [labels](resources--site_mesh_group--reference--group-001.md#canonical-3330331331201030-1310031200101101-3112132002001330-3200103203121100-1023210133031203-0111110023301022-2020030212131003-2323310230111231) |
| `name` | [name](resources--site_mesh_group--reference--group-001.md#canonical-3111013230113021-3310121112132001-0133211110322003-3121003210021231-3032221002333111-2122033233233102-1131110011031223-3233200031330333) |
| `namespace` | [namespace](resources--site_mesh_group--reference--group-001.md#canonical-0321030010210132-0130123112013111-2103300130020021-3001231032033220-3031133012313300-2131210332000113-0310230200332131-1331300131212323) |
| `spoke_mesh` | [spoke_mesh](resources--site_mesh_group--reference--group-001.md#canonical-2321033110130033-1103201210203103-0333013232313111-2332333310012302-0132123103301003-0011303113023300-1110213220112101-1100002031033022) |
| `spoke_mesh.control_and_data_plane_mesh` | [spoke_mesh.control_and_data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-3202313111202322-2200023333021010-1202321332011031-0233131030102231-0103213312211320-0001310113201033-3123120020121230-2223011233101313) |
| `spoke_mesh.data_plane_mesh` | [spoke_mesh.data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-2032032300223010-0121201213032122-0023102200320320-2333131001322112-0031232313220013-1322312222220310-0332220200330311-2310030303320013) |
| `spoke_mesh.hub_mesh_group` | [spoke_mesh.hub_mesh_group](resources--site_mesh_group--reference--group-001.md#canonical-2013133220032202-2332200332221021-0210013202030112-0113213002123231-2010133231311330-1200113112001132-2021033120322001-2022211012122003) |
| `spoke_mesh.hub_mesh_group.name` | [spoke_mesh.hub_mesh_group.name](resources--site_mesh_group--reference--group-001.md#canonical-0303312000221021-1000221120221002-3101032230233322-3010033130212003-3202011021131212-0013333201122213-1113302333103302-3000020132133202) |
| `spoke_mesh.hub_mesh_group.namespace` | [spoke_mesh.hub_mesh_group.namespace](resources--site_mesh_group--reference--group-001.md#canonical-1131303032212333-0213031133102322-3222233133011223-0330230201032312-3133110331011210-0221320302313310-0303103330201110-3022221133101010) |
| `spoke_mesh.hub_mesh_group.tenant` | [spoke_mesh.hub_mesh_group.tenant](resources--site_mesh_group--reference--group-001.md#canonical-0311302213110201-3202102303233333-2011201322033000-0302012032011012-0312133222211202-0221111021023032-1110233312012212-2033122330030220) |
| `timeouts` | [timeouts](resources--site_mesh_group--reference--group-001.md#canonical-1332313310221201-2332331210021231-3130101223313020-2003313201320021-3333213100120131-0232233323231123-0201001011202300-3013111211011010) |
| `timeouts.create` | [timeouts.create](resources--site_mesh_group--reference--group-001.md#canonical-3302032300020032-3121300131301010-1003103030211220-0033203131332011-1010200203111330-0232110312202022-2000012001100101-2301103032111001) |
| `timeouts.delete` | [timeouts.delete](resources--site_mesh_group--reference--group-001.md#canonical-0012033010303223-2221320122020002-0221102123033031-1022323133212020-1023103203012211-2001333220332020-3320332002130322-1312033232132033) |
| `timeouts.read` | [timeouts.read](resources--site_mesh_group--reference--group-001.md#canonical-1303132320023132-2212121300131111-0323322003103000-3313220312330111-3231231310300331-0110120302312221-3112001102232321-2132211202303300) |
| `timeouts.update` | [timeouts.update](resources--site_mesh_group--reference--group-001.md#canonical-1321312101012232-2230030333210203-1000222022300232-0132313332021200-3111031032010033-0333131201211332-0300022310122030-0200322110003231) |
| `virtual_site` | [virtual_site](resources--site_mesh_group--reference--group-001.md#canonical-1233323031233020-1112232000321131-0001130110220231-1132023000111121-1123203022120303-2311212121323132-3100102110120103-0323200332323023) |
| `virtual_site.kind` | [virtual_site.kind](resources--site_mesh_group--reference--group-001.md#canonical-3333331123031323-1233131233110013-1133322113130132-3232123023103203-2113330223023130-2020020000122132-0222321312312203-0132003030110003) |
| `virtual_site.name` | [virtual_site.name](resources--site_mesh_group--reference--group-001.md#canonical-0003102302133023-1322031220103230-3131110122032303-1102002203131222-1132310303322020-0031222013300012-2113223210120022-2300220021332033) |
| `virtual_site.namespace` | [virtual_site.namespace](resources--site_mesh_group--reference--group-001.md#canonical-0311013212131122-0022121202130123-3302102321113210-1030331000020102-0332113303230131-3301023330200201-1101312001210300-1322231331331200) |
| `virtual_site.tenant` | [virtual_site.tenant](resources--site_mesh_group--reference--group-001.md#canonical-3221222211001210-1323232233301210-0212303333333132-3103212112120110-3232231212032020-2312131001021330-1003321320032023-0231232101102010) |
| `virtual_site.uid` | [virtual_site.uid](resources--site_mesh_group--reference--group-001.md#canonical-0301003310012011-1000012223120233-2130220233222002-2313130302200332-1110122313302111-3303233321112021-1122001000300222-0131002110322312) |

<a id="canonical-1011232323113121-1231320202132111-0231022231120310-2210321103313102-0131133032201310-0313033331122131-2021033033022030-0101021032123202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bfd_disabled` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- bfd_disabled

<a id="canonical-0101313111020322-2113120031302131-1332220223000312-1131211102201331-2221111101132222-1331003201020030-3122022132321201-2301032211320230"></a>

Type: `["object", {}]`. Optional.

\[OneOf: bfd\_disabled, bfd\_enabled\] Enable this option

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

OneOf alternatives in this subsection:

- [bfd_disabled](resources--site_mesh_group--reference--group-001.md#canonical-0101313111020322-2113120031302131-1332220223000312-1131211102201331-2221111101132222-1331003201020030-3122022132321201-2301032211320230)
- [bfd_enabled](resources--site_mesh_group--reference--group-001.md#canonical-2321013203133110-3011233023021013-0220101011031001-1201120002111200-0023133300322211-1130312000103211-3313233103203200-1303211011012003)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bfd_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030312132232112-2303110333223203-2033302223131130-2123310023033230-1313103013010103-0011023222103320-3330033300112032-0231330313022102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bfd_enabled` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- bfd_enabled

<a id="canonical-2321013203133110-3011233023021013-0220101011031001-1201120002111200-0023133300322211-1130312000103211-3313233103203200-1303211011012003"></a>

Type: `"object"`. single nested block, Optional.

BFD. BFD parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("multiplier",
    "receive_interval_milliseconds",
    "transmit_interval_milliseconds")}
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
bfd_enabled {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330212321330132-2211122231133202-0203103002333002-2321301102210222-2032230210200120-2000312300232111-3013033102002002-1013321000001121"></a>

### Direct properties for `bfd_enabled`

<a id="canonical-1011210210120003-2010200031331221-3130230210210311-2003011112003330-2233231123120310-3220102331333100-2111011031110011-3120223123312222"></a>

#### `bfd_enabled.multiplier` property

Type: `"number"`. Optional.

Specify Number of missed packets to bring session down'.

Additional upstream details:

Specify Number of missed packets to bring session down"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(2, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-2322210323101230-0000303010023331-2321132202013001-3301110321021310-3311012300011301-0231020131311023-3210103213123113-2103022320321231"></a>

<a id="canonical-1331312212302002-1231223031320001-1300321220111120-0021331002011322-2220102010013000-2312211033233320-1232121101322102-0113233331200121"></a>

#### `bfd_enabled.receive_interval_milliseconds` property

Type: `"number"`. Optional.

BFD receive interval timer, in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(300, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-2000123332033210-2113220001230010-0212011233011112-0202202113103313-3321121023300201-3233231212020111-0231120320220010-3223100311303133"></a>

<a id="canonical-2213222201312312-3023321023331200-0110032113212132-3201213200223100-3112123223323132-0020211020132320-1330312020211300-1021331200112021"></a>

#### `bfd_enabled.transmit_interval_milliseconds` property

Type: `"number"`. Optional.

BFD transmit interval timer, in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(300, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-2300102311310103-1132033112330111-0232101032131033-2010201232322220-2230101212022313-0030303231303133-2111320001222103-2322220330311130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_re_fallback` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- disable_re_fallback

<a id="canonical-1211023223123110-2301031233302033-0331221310230020-1113230000221013-0012311000330321-3131020312011232-3102110311301020-3111103122320201"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_re\_fallback, enable\_re\_fallback; Default: disable\_re\_fallback\] Configuration
parameter for disable re fallback.

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

OneOf alternatives in this subsection:

- [disable_re_fallback](resources--site_mesh_group--reference--group-001.md#canonical-1211023223123110-2301031233302033-0331221310230020-1113230000221013-0012311000330321-3131020312011232-3102110311301020-3111103122320201)
- [enable_re_fallback](resources--site_mesh_group--reference--group-001.md#canonical-1131110120030302-2220303333132110-1002133212021210-3303333010122321-0132310030210303-2030310233233102-1002323332113323-0300211120333310)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_re_fallback = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210020200220232-0120031130020021-3022232122303201-3202302300231332-2221213031003201-0111003122200303-1332212132102212-1001102111323201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_re_fallback` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- enable_re_fallback

<a id="canonical-1131110120030302-2220303333132110-1002133212021210-3303333010122321-0132310030210303-2030310233233102-1002323332113323-0300211120333310"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable re fallback.

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
enable_re_fallback = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031001210322023-1032211110103322-2220103113201122-2113312202301300-3130323130310120-1212022100320032-0000210211323322-1333323000112211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `full_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- full_mesh

<a id="canonical-0122310323131300-1222303131332010-2033032130111011-2003113221301311-3230212121213220-2210131302032320-3212330202102001-1111201213102310"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: full\_mesh, hub\_mesh, spoke\_mesh\] Full Mesh. Details of Full Mesh Group Type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("control_and_data_plane_mesh",
    "data_plane_mesh")}
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
  "x-ves-oneof-field-full_mesh_choice": "[\"control_and_data_plane_mesh\",\"data_plane_mesh\"]"
}
```

OneOf alternatives in this subsection:

- [full_mesh](resources--site_mesh_group--reference--group-001.md#canonical-0122310323131300-1222303131332010-2033032130111011-2003113221301311-3230212121213220-2210131302032320-3212330202102001-1111201213102310)
- [hub_mesh](resources--site_mesh_group--reference--group-001.md#canonical-3211112213301131-2132020103310313-1103001123230200-2021232113010023-3303132200003003-2010310110012030-2230103121320323-3330201233210110)
- [spoke_mesh](resources--site_mesh_group--reference--group-001.md#canonical-2321033110130033-1103201210203103-0333013232313111-2332333310012302-0132123103301003-0011303113023300-1110213220112101-1100002031033022)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
full_mesh {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221130223111011-3333121020021121-3031111103203223-3331200213021322-1102113201020310-3320213121322111-3122313032210301-0213130313122212"></a>

### Direct properties for `full_mesh`

- [control_and_data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-3102203223310012-3000220213010202-2300032101300223-0211031112110121-2313102332221113-3222000021231111-1031131212203031-2210300201133230): complete subsection reference.

- [data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-2332121020021133-1213312001333122-0021013333033323-3123120211213223-1133201321011030-1211323033122032-0213102302211233-2233101022010232): complete subsection reference.

<a id="canonical-3102203223310012-3000220213010202-2300032101300223-0211031112110121-2313102332221113-3222000021231111-1031131212203031-2210300201133230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `full_mesh.control_and_data_plane_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- [full_mesh](resources--site_mesh_group--reference--group-001.md#canonical-2031001210322023-1032211110103322-2220103113201122-2113312202301300-3130323130310120-1212022100320032-0000210211323322-1333323000112211)
- full_mesh.control_and_data_plane_mesh

<a id="canonical-3002031122222220-0201203221211030-1031202103123113-1012100300133331-2211012321101130-3310103332321200-1123220032001022-2230232310333121"></a>

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
control_and_data_plane_mesh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332121020021133-1213312001333122-0021013333033323-3123120211213223-1133201321011030-1211323033122032-0213102302211233-2233101022010232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `full_mesh.data_plane_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- [full_mesh](resources--site_mesh_group--reference--group-001.md#canonical-2031001210322023-1032211110103322-2220103113201122-2113312202301300-3130323130310120-1212022100320032-0000210211323322-1333323000112211)
- full_mesh.data_plane_mesh

<a id="canonical-2031333110320310-3222032001301300-1222102032313131-0203312013313102-1222013033022033-2103002201202111-3033133200113100-3100202231011023"></a>

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
data_plane_mesh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011133001312013-2312132233110312-1332302222102230-1122130331221011-0121321303130030-0111233333133330-2112011323002320-2210320323220112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `hub_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- hub_mesh

<a id="canonical-3211112213301131-2132020103310313-1103001123230200-2021232113010023-3303132200003003-2010310110012030-2230103121320323-3330201233210110"></a>

Type: `"object"`. single nested block, Optional.

Hub Full Mesh. Details of Hub Full Mesh Group Type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("control_and_data_plane_mesh",
    "data_plane_mesh")}
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
  "x-ves-oneof-field-hub_full_mesh_choice": "[\"control_and_data_plane_mesh\",\"data_plane_mesh\"]"
}
```

Terraform syntax:

```terraform
hub_mesh {
  # Configure direct properties listed below.
}
```

<a id="canonical-3323123200330213-2030021333313332-2032030013333200-2030303132213212-1323223333222331-0212122211210323-0021101121003322-1312031110303213"></a>

### Direct properties for `hub_mesh`

- [control_and_data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-2332011320012102-0132213232111320-0002300001102211-3003211032123322-1222011323103100-0213321200113220-1300230103101010-0111313130120321): complete subsection reference.

- [data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-3031302310013033-1032002130101310-2103322102131013-2022120201030321-0130120330223130-3013222200012001-0201112220311002-2101332220222000): complete subsection reference.

<a id="canonical-2332011320012102-0132213232111320-0002300001102211-3003211032123322-1222011323103100-0213321200113220-1300230103101010-0111313130120321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `hub_mesh.control_and_data_plane_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- [hub_mesh](resources--site_mesh_group--reference--group-001.md#canonical-1011133001312013-2312132233110312-1332302222102230-1122130331221011-0121321303130030-0111233333133330-2112011323002320-2210320323220112)
- hub_mesh.control_and_data_plane_mesh

<a id="canonical-1120030130300021-1132111233003120-1203333131232132-1311120223023032-3003003231301223-0301310321130112-1311323001312121-0110212000022100"></a>

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
control_and_data_plane_mesh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031302310013033-1032002130101310-2103322102131013-2022120201030321-0130120330223130-3013222200012001-0201112220311002-2101332220222000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `hub_mesh.data_plane_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- [hub_mesh](resources--site_mesh_group--reference--group-001.md#canonical-1011133001312013-2312132233110312-1332302222102230-1122130331221011-0121321303130030-0111233333133330-2112011323002320-2210320323220112)
- hub_mesh.data_plane_mesh

<a id="canonical-0231011202033320-2103113322232103-0122131313020233-1110020112121123-3103102101112022-0030222212133000-0311120232030313-1222131202032321"></a>

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
data_plane_mesh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020322201021122-0131011313222121-0101022011301302-3002331111101232-3230013331330023-0032312121112300-2322110320130302-2012302020131121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `spoke_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- spoke_mesh

<a id="canonical-2321033110130033-1103201210203103-0333013232313111-2332333310012302-0132123103301003-0011303113023300-1110213220112101-1100002031033022"></a>

Type: `"object"`. single nested block, Optional.

Spoke. Details of Spoke Mesh Group Type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("control_and_data_plane_mesh",
    "data_plane_mesh")}
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
  "x-ves-oneof-field-spoke_hub_mesh_choice": "[\"control_and_data_plane_mesh\",\"data_plane_mesh\"]"
}
```

Terraform syntax:

```terraform
spoke_mesh {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233003122110131-0102133030010300-1012311012110132-0130333023231012-3131320221022302-2021200211012203-3300210210233011-3031101301021033"></a>

### Direct properties for `spoke_mesh`

- [control_and_data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-2312110011102211-3323211202311201-3233222223013310-0320013122201012-2101111102011132-0002103320100303-3221013002331302-2330211233300031): complete subsection reference.

- [data_plane_mesh](resources--site_mesh_group--reference--group-001.md#canonical-3313130132312313-3323213312213100-2112302012030121-2212303020012102-3300032113133100-2311213210213213-3033001231302012-0323132112303313): complete subsection reference.

- [hub_mesh_group](resources--site_mesh_group--reference--group-001.md#canonical-2213123103020333-2012132232120023-2222111203030131-1030232222213102-1111323003213030-0120131301031003-3331100003131103-0002100030030303): complete subsection reference.

<a id="canonical-2312110011102211-3323211202311201-3233222223013310-0320013122201012-2101111102011132-0002103320100303-3221013002331302-2330211233300031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `spoke_mesh.control_and_data_plane_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- [spoke_mesh](resources--site_mesh_group--reference--group-001.md#canonical-2020322201021122-0131011313222121-0101022011301302-3002331111101232-3230013331330023-0032312121112300-2322110320130302-2012302020131121)
- spoke_mesh.control_and_data_plane_mesh

<a id="canonical-3202313111202322-2200023333021010-1202321332011031-0233131030102231-0103213312211320-0001310113201033-3123120020121230-2223011233101313"></a>

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
control_and_data_plane_mesh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313130132312313-3323213312213100-2112302012030121-2212303020012102-3300032113133100-2311213210213213-3033001231302012-0323132112303313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `spoke_mesh.data_plane_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- [spoke_mesh](resources--site_mesh_group--reference--group-001.md#canonical-2020322201021122-0131011313222121-0101022011301302-3002331111101232-3230013331330023-0032312121112300-2322110320130302-2012302020131121)
- spoke_mesh.data_plane_mesh

<a id="canonical-2032032300223010-0121201213032122-0023102200320320-2333131001322112-0031232313220013-1322312222220310-0332220200330311-2310030303320013"></a>

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
data_plane_mesh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213123103020333-2012132232120023-2222111203030131-1030232222213102-1111323003213030-0120131301031003-3331100003131103-0002100030030303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `spoke_mesh.hub_mesh_group` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- [spoke_mesh](resources--site_mesh_group--reference--group-001.md#canonical-2020322201021122-0131011313222121-0101022011301302-3002331111101232-3230013331330023-0032312121112300-2322110320130302-2012302020131121)
- spoke_mesh.hub_mesh_group

<a id="canonical-2013133220032202-2332200332221021-0210013202030112-0113213002123231-2010133231311330-1200113112001132-2021033120322001-2022211012122003"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
hub_mesh_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-1123003232313000-1300213232333113-3302031020222211-2132330100132202-2330111303031121-1002223223230003-0003132001201210-0130110013023121"></a>

### Direct properties for `spoke_mesh.hub_mesh_group`

<a id="canonical-0303312000221021-1000221120221002-3101032230233322-3010033130212003-3202011021131212-0013333201122213-1113302333103302-3000020132133202"></a>

#### `spoke_mesh.hub_mesh_group.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1131303032212333-0213031133102322-3222233133011223-0330230201032312-3133110331011210-0221320302313310-0303103330201110-3022221133101010"></a>

<a id="canonical-0032331033310331-2101201202122200-2201303000120122-3212130300321313-0210020103300301-3000202013312101-3111211300210200-1111011013311220"></a>

#### `spoke_mesh.hub_mesh_group.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0311302213110201-3202102303233333-2011201322033000-0302012032011012-0312133222211202-0221111021023032-1110233312012212-2033122330030220"></a>

<a id="canonical-3311300030030013-0203111222211030-3302331003011322-1201212302011001-0213030131202003-2033203231033231-3203222023132202-0233302333111112"></a>

#### `spoke_mesh.hub_mesh_group.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0011100231203201-0101031320333120-1133233320032200-1122201010121132-2120012103121301-1030131300330330-1202223111001323-0213302002311313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- timeouts

<a id="canonical-1332313310221201-2332331210021231-3130101223313020-2003313201320021-3333213100120131-0232233323231123-0201001011202300-3013111211011010"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330303120213100-0303020331132003-1210113222211213-3212021212333313-2303100331222023-3020011232021303-0112213311110030-3031100012133013"></a>

### Direct properties for `timeouts`

<a id="canonical-3302032300020032-3121300131301010-1003103030211220-0033203131332011-1010200203111330-0232110312202022-2000012001100101-2301103032111001"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0012033010303223-2221320122020002-0221102123033031-1022323133212020-1023103203012211-2001333220332020-3320332002130322-1312033232132033"></a>

<a id="canonical-1003331223010330-3120000320002200-2010300311332001-0133020332210100-3002312010203020-2012233111023130-1330012011221222-1020210132312303"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1303132320023132-2212121300131111-0323322003103000-3313220312330111-3231231310300331-0110120302312221-3112001102232321-2132211202303300"></a>

<a id="canonical-3101233222003322-0303131220201023-1032111303201221-2130123203033102-1233012130311221-3223103122220131-1132032230220201-2122120313302333"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1321312101012232-2230030333210203-1000222022300232-0132313332021200-3111031032010033-0333131201211332-0300022310122030-0200322110003231"></a>

<a id="canonical-2230031231301321-0013213123331112-1310010221123013-2022320211123111-0001222302201321-2011011213123221-2230230103313101-3102312101211033"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1220313230311013-1302220301022233-1002221232313132-0312222012013021-1302112330300313-3232013033020322-3011101331113203-2312331132230000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_site` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Property reference](resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- virtual_site

<a id="canonical-1233323031233020-1112232000321131-0001130110220231-1132023000111121-1123203022120303-2311212121323132-3100102110120103-0323200332323023"></a>

Type: `"object"`. list nested block, Optional.

Set of sites for which this mesh group config is valid. If 'Type' is Spoke, then it gives set of
spoke sites. If 'Type' is Hub, then it gives set of hub sites. If 'Type' is Full Mesh, then it gives
set of sites that are connected in full mesh.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210133102032330-2302323000100303-0300202023112123-1313131022312010-1221303101021110-2213121310103031-1100332210002012-2112333010122223"></a>

### Direct properties for `virtual_site`

<a id="canonical-3333331123031323-1233131233110013-1133322113130132-3232123023103203-2113330223023130-2020020000122132-0222321312312203-0132003030110003"></a>

#### `virtual_site.kind` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0003102302133023-1322031220103230-3131110122032303-1102002203131222-1132310303322020-0031222013300012-2113223210120022-2300220021332033"></a>

<a id="canonical-1110102311232032-0323012301020221-3003032132012133-0011113201210003-3122212332121313-1302331220310221-3000032233333021-3203222230301121"></a>

#### `virtual_site.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0311013212131122-0022121202130123-3302102321113210-1030331000020102-0332113303230131-3301023330200201-1101312001210300-1322231331331200"></a>

<a id="canonical-0210001112120131-1030203003330112-0202121102200333-0231210112123222-0003313010111032-2002133111000312-1121033132013003-2313212022212120"></a>

#### `virtual_site.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3221222211001210-1323232233301210-0212303333333132-3103212112120110-3232231212032020-2312131001021330-1003321320032023-0231232101102010"></a>

<a id="canonical-1230230113213002-0130201133212233-1132113222002302-0232332302021133-1002332100332010-2303103310212103-2012202031333130-3201113231130023"></a>

#### `virtual_site.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0301003310012011-1000012223120233-2130220233222002-2313130302200332-1110122313302111-3303233321112021-1122001000300222-0131002110322312"></a>

<a id="canonical-0122230202003221-3202232211022201-0003303102120223-0122030330320230-1331211131001223-0021131012231230-1111110202102131-1220313332101032"></a>

#### `virtual_site.uid` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
