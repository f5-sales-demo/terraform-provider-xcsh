---
page_title: "xcsh_site_mesh_group reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site_mesh_group reference."
---

# xcsh_site_mesh_group reference

<a id="canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- Property reference

<a id="canonical-0313111121221011-0001100332112133-3021320121032321-2022133333122000-1302103213221320-1001221300111103-2220001103332200-2201000333223133"></a>

### Direct properties for `xcsh_site_mesh_group`

<a id="canonical-3112221023020213-1330002022203302-3001000030121132-3303303022210013-3203000201021330-0330232312112022-3031110123300313-3021200333002310"></a>

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

- [bfd_disabled](data-sources--site_mesh_group--reference--group-001.md#canonical-0120010001120102-0333310311331020-1313102312032202-0033232210203222-2131203003233011-3032223131120022-3221131301030203-3323321003210020): complete subsection reference.

- [bfd_enabled](data-sources--site_mesh_group--reference--group-001.md#canonical-3213211310202102-0211010112232012-2201332101100002-0000332311333313-2331210033000311-0220021223001321-0113001300201002-0221021220002020): complete subsection reference.

<a id="canonical-1233232110322033-0033223320123012-3332100110211012-1101023203112031-3210112110113122-3002021333212310-0023013122110111-2321110223303123"></a>

<a id="canonical-2302202201130220-3323020032021233-3022311220313003-1013220113311230-1230021001231300-1000013111210200-2013123023021023-0100100112200010"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the SiteMeshGroup.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [disable_re_fallback](data-sources--site_mesh_group--reference--group-001.md#canonical-0012313312310213-3103022222303201-1021230022300313-3021012123030303-3003233320002212-0111300110022330-0031230303230301-1003123131131222): complete subsection reference.

- [enable_re_fallback](data-sources--site_mesh_group--reference--group-001.md#canonical-1010003320223310-2111201101111132-3311321013102132-1303313201123330-3310011000101310-0123310221220033-3311210202201231-3021310232333000): complete subsection reference.

- [full_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-2011311010303311-3101333312312330-0321213121133131-0132222021121301-0211101131120320-2330312212002010-0320002102331223-2100120012300212): complete subsection reference.

- [hub_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-3310111110031031-0031312012323212-1332110322233232-3121231211210320-1320111110002130-2310110100121303-2100213103112331-3310121200102323): complete subsection reference.

<a id="canonical-1021221312133000-0200031121220101-3000203230232231-3023111033301130-3323123302031322-0031002312112300-2110212013000033-1203323220100222"></a>

<a id="canonical-3001232103021021-2213223220223020-1231223201202021-3023131030223101-2011103102131110-3320233033203023-2001101011133122-2232202320133122"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3330000023233310-2232011102313210-0320020131002033-3101110230211303-2331133002320021-3201313132313130-3221200110210101-2110210022231032"></a>

<a id="canonical-3010010230101211-3131012200323222-1233011002010222-0023313030001111-2232200131203103-2033131311020233-2010211102333301-0100113223003000"></a>

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

<a id="canonical-2031133200323330-2012133032230230-1211110232030102-3300201023130021-1300131133311102-3210313130130013-3201003232201301-3200002032010323"></a>

<a id="canonical-2333311000120021-3332321231232122-2333231323221103-0230001122003130-1220202020001331-1010300332310313-3321323231321232-1212211201022221"></a>

#### `name` property

Type: `"string"`. Required.

Name of the SiteMeshGroup.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0211310322010023-1321312023132331-0221010022132302-1002113003102323-0132223133010220-2200132111131122-3011300031330103-0330113302020023"></a>

<a id="canonical-1100031303031030-3001111102201323-1230120312111220-3002333101331113-2202032002031113-3101110020200031-3033101110210102-3322022122122213"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the SiteMeshGroup exists.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [spoke_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-3303230010212332-0123100010332211-2131111021030230-3301330103211030-1101120303232312-0201021220110323-0203132331020320-2023123001121013): complete subsection reference.

- [virtual_site](data-sources--site_mesh_group--reference--group-001.md#canonical-2221320302210102-1123132201311320-1010233213033022-3003013130331211-2310202133312102-2122000123100331-1022023011231202-3011111122002232): complete subsection reference.

<a id="canonical-3102000013001001-2121333112030133-0101323321010310-2222110031023300-3021303233321023-2031330222123131-0110131300211313-1011230000121113"></a>

### All schema paths for `xcsh_site_mesh_group`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--site_mesh_group--reference--group-001.md#canonical-3112221023020213-1330002022203302-3001000030121132-3303303022210013-3203000201021330-0330232312112022-3031110123300313-3021200333002310) |
| `bfd_disabled` | [bfd_disabled](data-sources--site_mesh_group--reference--group-001.md#canonical-1201201003100222-0213211102011012-2231031021022222-2020231321332232-2022211203221323-1133210013211023-0323333233322233-2020020300012010) |
| `bfd_enabled` | [bfd_enabled](data-sources--site_mesh_group--reference--group-001.md#canonical-3321310132013103-3202122002202013-2022032202112131-3003113310113201-2112221031100011-0020101301213131-0132120310023131-0103033221312321) |
| `bfd_enabled.multiplier` | [bfd_enabled.multiplier](data-sources--site_mesh_group--reference--group-001.md#canonical-1302101301003202-0102110222032122-0100301000113303-1313300023223131-3121020113020033-3301111021100203-2231310302321200-2001131212121221) |
| `bfd_enabled.receive_interval_milliseconds` | [bfd_enabled.receive_interval_milliseconds](data-sources--site_mesh_group--reference--group-001.md#canonical-0133221103112030-2122121233210031-1021321203211221-2010110011000001-1231230013203021-2131011312300003-1331200101331220-2113220102212332) |
| `bfd_enabled.transmit_interval_milliseconds` | [bfd_enabled.transmit_interval_milliseconds](data-sources--site_mesh_group--reference--group-001.md#canonical-0121302220310312-3222002123022320-2203300210231132-3012312111211221-2032110332030200-3323233112310112-0122020131303232-2010230330311133) |
| `description` | [description](data-sources--site_mesh_group--reference--group-001.md#canonical-1233232110322033-0033223320123012-3332100110211012-1101023203112031-3210112110113122-3002021333212310-0023013122110111-2321110223303123) |
| `disable_re_fallback` | [disable_re_fallback](data-sources--site_mesh_group--reference--group-001.md#canonical-2323220320221132-3001020102202031-3332203312012111-2331222310233123-3212303233221303-3201033102122210-3001121332121011-3100310202333022) |
| `enable_re_fallback` | [enable_re_fallback](data-sources--site_mesh_group--reference--group-001.md#canonical-1022310133321013-1123230211303132-0232321111223230-3120332301113120-1332102010100233-0310303112231223-1132213032221031-1020312203102221) |
| `full_mesh` | [full_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-3313130321120000-3333021330130130-3313100202200023-2012221313210300-2321321013212220-3301333222031302-2211310221010003-3113030100000332) |
| `full_mesh.control_and_data_plane_mesh` | [full_mesh.control_and_data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-1111311121032333-0031122022113000-0303110320322331-2033201200221212-1131220232201202-1131003100101221-1021232000332210-2231222112212313) |
| `full_mesh.data_plane_mesh` | [full_mesh.data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-1222210331200010-0323131322232202-3201132212011221-3102213033220002-1303021121131230-1330023202310020-3032023102020112-2102102320020313) |
| `hub_mesh` | [hub_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-0211032203133322-1201321110323212-3332102010113022-1312303311100101-0200332110223031-1101303121301220-1332331311303320-0112302223021330) |
| `hub_mesh.control_and_data_plane_mesh` | [hub_mesh.control_and_data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-0111231123022221-3301222102221322-3022222211321032-2231100021333121-3121302310211033-3300330012002230-2121130231022323-1020310002113032) |
| `hub_mesh.data_plane_mesh` | [hub_mesh.data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-1011100202202313-1331033103301212-1003311012022301-2131103121220033-2322011131101011-1032220202031122-3202302230303321-1031203312200121) |
| `id` | [ID](data-sources--site_mesh_group--reference--group-001.md#canonical-1021221312133000-0200031121220101-3000203230232231-3023111033301130-3323123302031322-0031002312112300-2110212013000033-1203323220100222) |
| `labels` | [labels](data-sources--site_mesh_group--reference--group-001.md#canonical-3330000023233310-2232011102313210-0320020131002033-3101110230211303-2331133002320021-3201313132313130-3221200110210101-2110210022231032) |
| `name` | [name](data-sources--site_mesh_group--reference--group-001.md#canonical-2031133200323330-2012133032230230-1211110232030102-3300201023130021-1300131133311102-3210313130130013-3201003232201301-3200002032010323) |
| `namespace` | [namespace](data-sources--site_mesh_group--reference--group-001.md#canonical-0211310322010023-1321312023132331-0221010022132302-1002113003102323-0132223133010220-2200132111131122-3011300031330103-0330113302020023) |
| `spoke_mesh` | [spoke_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-2201320211322133-2323020032020103-2011310103130100-0131122201130203-3110000132220303-3030001230310222-2223233220120003-2012002203331202) |
| `spoke_mesh.control_and_data_plane_mesh` | [spoke_mesh.control_and_data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-3122021131101022-3013201313330100-1103231311120323-0003210012012001-3001300103120003-2032212213203321-1001231003012230-2012303032233002) |
| `spoke_mesh.data_plane_mesh` | [spoke_mesh.data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-2112030302301133-2223100030123323-0111103212330323-2210122132320023-0223033021220110-2332223022202332-0213333122010233-0310033131200031) |
| `spoke_mesh.hub_mesh_group` | [spoke_mesh.hub_mesh_group](data-sources--site_mesh_group--reference--group-001.md#canonical-1122233213203010-1031330122201330-0000330133121322-0301121220123221-1122303302313332-3112330002112223-0021201301303211-0332030103311332) |
| `spoke_mesh.hub_mesh_group.name` | [spoke_mesh.hub_mesh_group.name](data-sources--site_mesh_group--reference--group-001.md#canonical-1222323200210122-0323123203002300-3111231330220331-2220012200120221-0112112010100300-1030012333311010-2112112020200323-1001203130002001) |
| `spoke_mesh.hub_mesh_group.namespace` | [spoke_mesh.hub_mesh_group.namespace](data-sources--site_mesh_group--reference--group-001.md#canonical-2221103000011231-0032300203020131-0012210320130333-2210033002003320-3003031122213210-2022022023323100-1120112211023222-2102010122330301) |
| `spoke_mesh.hub_mesh_group.tenant` | [spoke_mesh.hub_mesh_group.tenant](data-sources--site_mesh_group--reference--group-001.md#canonical-0333322003120202-2123223111223002-2312113203010201-0103033111020123-1111333033221102-1113212333231120-3103213031312213-2233230020313111) |
| `virtual_site` | [virtual_site](data-sources--site_mesh_group--reference--group-001.md#canonical-1203212022030010-1212310223011312-3133322033330000-2200022121233110-1011012303021002-3112331033312320-1020330031003032-3033322011033201) |
| `virtual_site.kind` | [virtual_site.kind](data-sources--site_mesh_group--reference--group-001.md#canonical-3212110233033321-3002301130000213-2103201013333320-0210330331301011-0123233111133001-3332102320213102-3332130000203220-0223012202320230) |
| `virtual_site.name` | [virtual_site.name](data-sources--site_mesh_group--reference--group-001.md#canonical-3011200120130123-2213302323213030-2122310022102333-1320232031102230-1002330310020102-3032212322121210-3001002323233112-2002301122231223) |
| `virtual_site.namespace` | [virtual_site.namespace](data-sources--site_mesh_group--reference--group-001.md#canonical-0201120032211301-1313201300132021-3212212331100330-3311220022000110-1001100020321222-2213230101103122-0120322123130221-1033032220002112) |
| `virtual_site.tenant` | [virtual_site.tenant](data-sources--site_mesh_group--reference--group-001.md#canonical-2302303300002300-2210021132312301-3113013232113131-2323102322033003-3212110022132300-0130122313020320-1312110231103222-1323010311203203) |
| `virtual_site.uid` | [virtual_site.uid](data-sources--site_mesh_group--reference--group-001.md#canonical-0230210330200133-0320123230300112-3302233131102333-1123220330132010-0222310100221211-2223010323103023-2101010222031302-2231303323133203) |

<a id="canonical-0120010001120102-0333310311331020-1313102312032202-0033232210203222-2131203003233011-3032223131120022-3221131301030203-3323321003210020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bfd_disabled` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011)
- bfd_disabled

<a id="canonical-1201201003100222-0213211102011012-2231031021022222-2020231321332232-2022211203221323-1133210013211023-0323333233322233-2020020300012010"></a>

Type: `["object", {}]`. Computed.

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

- [bfd_disabled](data-sources--site_mesh_group--reference--group-001.md#canonical-1201201003100222-0213211102011012-2231031021022222-2020231321332232-2022211203221323-1133210013211023-0323333233322233-2020020300012010)
- [bfd_enabled](data-sources--site_mesh_group--reference--group-001.md#canonical-3321310132013103-3202122002202013-2022032202112131-3003113310113201-2112221031100011-0020101301213131-0132120310023131-0103033221312321)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213211310202102-0211010112232012-2201332101100002-0000332311333313-2331210033000311-0220021223001321-0113001300201002-0221021220002020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bfd_enabled` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011)
- bfd_enabled

<a id="canonical-3321310132013103-3202122002202013-2022032202112131-3003113310113201-2112221031100011-0020101301213131-0132120310023131-0103033221312321"></a>

Type: `"single"`. Computed.

BFD. BFD parameters.

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

<a id="canonical-1011122302321001-3022332303121121-2300332230221230-1220220121030231-0113321223121230-2221032331121100-2211102213331212-0232223020030213"></a>

### Direct properties for `bfd_enabled`

<a id="canonical-1302101301003202-0102110222032122-0100301000113303-1313300023223131-3121020113020033-3301111021100203-2231310302321200-2001131212121221"></a>

#### `bfd_enabled.multiplier` property

Type: `"number"`. Computed.

Specify Number of missed packets to bring session down'.

Additional upstream details:

Specify Number of missed packets to bring session down"

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0133221103112030-2122121233210031-1021321203211221-2010110011000001-1231230013203021-2131011312300003-1331200101331220-2113220102212332"></a>

<a id="canonical-3030130011320330-1201312202111332-0121321122032003-0332212333320003-1112210110202121-3100033030213221-0031133100102101-0020101322320131"></a>

#### `bfd_enabled.receive_interval_milliseconds` property

Type: `"number"`. Computed.

BFD receive interval timer, in milliseconds.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0121302220310312-3222002123022320-2203300210231132-3012312111211221-2032110332030200-3323233112310112-0122020131303232-2010230330311133"></a>

<a id="canonical-1121112030331133-3111211303233323-1101032303332132-0232331022222213-3122132221012321-3033133133211233-2332010210131112-0112022122012030"></a>

#### `bfd_enabled.transmit_interval_milliseconds` property

Type: `"number"`. Computed.

BFD transmit interval timer, in milliseconds.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0012313312310213-3103022222303201-1021230022300313-3021012123030303-3003233320002212-0111300110022330-0031230303230301-1003123131131222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_re_fallback` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011)
- disable_re_fallback

<a id="canonical-2323220320221132-3001020102202031-3332203312012111-2331222310233123-3212303233221303-3201033102122210-3001121332121011-3100310202333022"></a>

Type: `["object", {}]`. Computed.

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

- [disable_re_fallback](data-sources--site_mesh_group--reference--group-001.md#canonical-2323220320221132-3001020102202031-3332203312012111-2331222310233123-3212303233221303-3201033102122210-3001121332121011-3100310202333022)
- [enable_re_fallback](data-sources--site_mesh_group--reference--group-001.md#canonical-1022310133321013-1123230211303132-0232321111223230-3120332301113120-1332102010100233-0310303112231223-1132213032221031-1020312203102221)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010003320223310-2111201101111132-3311321013102132-1303313201123330-3310011000101310-0123310221220033-3311210202201231-3021310232333000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_re_fallback` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011)
- enable_re_fallback

<a id="canonical-1022310133321013-1123230211303132-0232321111223230-3120332301113120-1332102010100233-0310303112231223-1132213032221031-1020312203102221"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011311010303311-3101333312312330-0321213121133131-0132222021121301-0211101131120320-2330312212002010-0320002102331223-2100120012300212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `full_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011)
- full_mesh

<a id="canonical-3313130321120000-3333021330130130-3313100202200023-2012221313210300-2321321013212220-3301333222031302-2211310221010003-3113030100000332"></a>

Type: `"single"`. Computed.

\[OneOf: full\_mesh, hub\_mesh, spoke\_mesh\] Full Mesh. Details of Full Mesh Group Type.

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

- [full_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-3313130321120000-3333021330130130-3313100202200023-2012221313210300-2321321013212220-3301333222031302-2211310221010003-3113030100000332)
- [hub_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-0211032203133322-1201321110323212-3332102010113022-1312303311100101-0200332110223031-1101303121301220-1332331311303320-0112302223021330)
- [spoke_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-2201320211322133-2323020032020103-2011310103130100-0131122201130203-3110000132220303-3030001230310222-2223233220120003-2012002203331202)

Select alternatives according to the provider validators above.

<a id="canonical-0313223002030333-3013222123300120-3123310230002000-1212311221101010-0321013303131210-1001023001212300-1302221101310303-3222211023021121"></a>

### Direct properties for `full_mesh`

- [control_and_data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-2301201112331101-3332111320033330-0020020203010101-3201103300023010-2110301100313122-2312102230331021-1002000010312301-1120033122013010): complete subsection reference.

- [data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-0020001230122013-0201110230133002-3211033113001120-3103303233033132-2233302202021022-3311333032033333-3010322131000002-1012022110022330): complete subsection reference.

<a id="canonical-2301201112331101-3332111320033330-0020020203010101-3201103300023010-2110301100313122-2312102230331021-1002000010312301-1120033122013010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `full_mesh.control_and_data_plane_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011)
- [full_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-2011311010303311-3101333312312330-0321213121133131-0132222021121301-0211101131120320-2330312212002010-0320002102331223-2100120012300212)
- full_mesh.control_and_data_plane_mesh

<a id="canonical-1111311121032333-0031122022113000-0303110320322331-2033201200221212-1131220232201202-1131003100101221-1021232000332210-2231222112212313"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0020001230122013-0201110230133002-3211033113001120-3103303233033132-2233302202021022-3311333032033333-3010322131000002-1012022110022330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `full_mesh.data_plane_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011)
- [full_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-2011311010303311-3101333312312330-0321213121133131-0132222021121301-0211101131120320-2330312212002010-0320002102331223-2100120012300212)
- full_mesh.data_plane_mesh

<a id="canonical-1222210331200010-0323131322232202-3201132212011221-3102213033220002-1303021121131230-1330023202310020-3032023102020112-2102102320020313"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310111110031031-0031312012323212-1332110322233232-3121231211210320-1320111110002130-2310110100121303-2100213103112331-3310121200102323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `hub_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011)
- hub_mesh

<a id="canonical-0211032203133322-1201321110323212-3332102010113022-1312303311100101-0200332110223031-1101303121301220-1332331311303320-0112302223021330"></a>

Type: `"single"`. Computed.

Hub Full Mesh. Details of Hub Full Mesh Group Type.

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

<a id="canonical-2021222030001213-3021311030200233-3220233232303011-2323120202203001-3301113203102322-0330200232012003-0010002032020033-3131313322101333"></a>

### Direct properties for `hub_mesh`

- [control_and_data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-2003221203223012-2310123232203332-3103121031010302-0130003132113332-3121222313112302-3010013013003222-1330313130010232-3321102131310021): complete subsection reference.

- [data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-3023231123300303-0131020203203002-1120131233022032-3213110223103202-3202011221131312-2110003320332322-0003233332320103-1013200230120320): complete subsection reference.

<a id="canonical-2003221203223012-2310123232203332-3103121031010302-0130003132113332-3121222313112302-3010013013003222-1330313130010232-3321102131310021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `hub_mesh.control_and_data_plane_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011)
- [hub_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-3310111110031031-0031312012323212-1332110322233232-3121231211210320-1320111110002130-2310110100121303-2100213103112331-3310121200102323)
- hub_mesh.control_and_data_plane_mesh

<a id="canonical-0111231123022221-3301222102221322-3022222211321032-2231100021333121-3121302310211033-3300330012002230-2121130231022323-1020310002113032"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023231123300303-0131020203203002-1120131233022032-3213110223103202-3202011221131312-2110003320332322-0003233332320103-1013200230120320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `hub_mesh.data_plane_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011)
- [hub_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-3310111110031031-0031312012323212-1332110322233232-3121231211210320-1320111110002130-2310110100121303-2100213103112331-3310121200102323)
- hub_mesh.data_plane_mesh

<a id="canonical-1011100202202313-1331033103301212-1003311012022301-2131103121220033-2322011131101011-1032220202031122-3202302230303321-1031203312200121"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303230010212332-0123100010332211-2131111021030230-3301330103211030-1101120303232312-0201021220110323-0203132331020320-2023123001121013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `spoke_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011)
- spoke_mesh

<a id="canonical-2201320211322133-2323020032020103-2011310103130100-0131122201130203-3110000132220303-3030001230310222-2223233220120003-2012002203331202"></a>

Type: `"single"`. Computed.

Spoke. Details of Spoke Mesh Group Type.

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

<a id="canonical-2113301133030001-2020031021101331-0332330223101230-2103312210121131-2132220212122230-3213231011032202-0312333101030102-2312211300322210"></a>

### Direct properties for `spoke_mesh`

- [control_and_data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-1030130010011202-3131331100222300-1210223133111031-2102210322100333-1321021113221200-3332012022230012-1203233033200323-3121021321030300): complete subsection reference.

- [data_plane_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-2213200121023202-1030111321221033-0211123332233133-3013322023303100-0220201113231211-3223132203131302-3120023133000030-3003232003003230): complete subsection reference.

- [hub_mesh_group](data-sources--site_mesh_group--reference--group-001.md#canonical-1332303021011210-0112003021202312-2031201100311213-1202200202112201-0233100232203020-0202210222031031-0202331111033122-1220203123200302): complete subsection reference.

<a id="canonical-1030130010011202-3131331100222300-1210223133111031-2102210322100333-1321021113221200-3332012022230012-1203233033200323-3121021321030300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `spoke_mesh.control_and_data_plane_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011)
- [spoke_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-3303230010212332-0123100010332211-2131111021030230-3301330103211030-1101120303232312-0201021220110323-0203132331020320-2023123001121013)
- spoke_mesh.control_and_data_plane_mesh

<a id="canonical-3122021131101022-3013201313330100-1103231311120323-0003210012012001-3001300103120003-2032212213203321-1001231003012230-2012303032233002"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213200121023202-1030111321221033-0211123332233133-3013322023303100-0220201113231211-3223132203131302-3120023133000030-3003232003003230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `spoke_mesh.data_plane_mesh` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011)
- [spoke_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-3303230010212332-0123100010332211-2131111021030230-3301330103211030-1101120303232312-0201021220110323-0203132331020320-2023123001121013)
- spoke_mesh.data_plane_mesh

<a id="canonical-2112030302301133-2223100030123323-0111103212330323-2210122132320023-0223033021220110-2332223022202332-0213333122010233-0310033131200031"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332303021011210-0112003021202312-2031201100311213-1202200202112201-0233100232203020-0202210222031031-0202331111033122-1220203123200302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `spoke_mesh.hub_mesh_group` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011)
- [spoke_mesh](data-sources--site_mesh_group--reference--group-001.md#canonical-3303230010212332-0123100010332211-2131111021030230-3301330103211030-1101120303232312-0201021220110323-0203132331020320-2023123001121013)
- spoke_mesh.hub_mesh_group

<a id="canonical-1122233213203010-1031330122201330-0000330133121322-0301121220123221-1122303302313332-3112330002112223-0021201301303211-0332030103311332"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-0302312332022003-2112113030300233-2021232030332013-3022333111130322-3302121030300201-3111332212101202-3103322312102210-1310221300003013"></a>

### Direct properties for `spoke_mesh.hub_mesh_group`

<a id="canonical-1222323200210122-0323123203002300-3111231330220331-2220012200120221-0112112010100300-1030012333311010-2112112020200323-1001203130002001"></a>

#### `spoke_mesh.hub_mesh_group.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2221103000011231-0032300203020131-0012210320130333-2210033002003320-3003031122213210-2022022023323100-1120112211023222-2102010122330301"></a>

<a id="canonical-3002310200211000-3133133301101320-0101103313022012-2331122030210210-0201120002111121-0331330003101103-0000212301302321-2330012102302120"></a>

#### `spoke_mesh.hub_mesh_group.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0333322003120202-2123223111223002-2312113203010201-0103033111020123-1111333033221102-1113212333231120-3103213031312213-2233230020313111"></a>

<a id="canonical-1311012001103202-3200321032201023-1033221301000113-1200112003331110-3021010032030311-3333202212231202-1112112112003000-1202210013122010"></a>

#### `spoke_mesh.hub_mesh_group.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2221320302210102-1123132201311320-1010233213033022-3003013130331211-2310202133312102-2122000123100331-1022023011231202-3011111122002232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_site` properties

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- [Property reference](data-sources--site_mesh_group--reference--group-001.md#canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011)
- virtual_site

<a id="canonical-1203212022030010-1212310223011312-3133322033330000-2200022121233110-1011012303021002-3112331033312320-1020330031003032-3033322011033201"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3122031323232221-3223121023230022-0302212311120032-2322123303230121-3232313213231300-0013322202233231-2112000120111121-3330003030033033"></a>

### Direct properties for `virtual_site`

<a id="canonical-3212110233033321-3002301130000213-2103201013333320-0210330331301011-0123233111133001-3332102320213102-3332130000203220-0223012202320230"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3011200120130123-2213302323213030-2122310022102333-1320232031102230-1002330310020102-3032212322121210-3001002323233112-2002301122231223"></a>

<a id="canonical-0332032113213030-0303223132323022-2012213222132331-1031313022132030-3332122311221010-3003321323101213-3213331130320131-0130231011100231"></a>

#### `virtual_site.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0201120032211301-1313201300132021-3212212331100330-3311220022000110-1001100020321222-2213230101103122-0120322123130221-1033032220002112"></a>

<a id="canonical-3023012221212130-3212131311230030-0332321021313132-1002112010003212-0000100301231033-1333010001211032-1212022222231302-3010132232102230"></a>

#### `virtual_site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2302303300002300-2210021132312301-3113013232113131-2323102322033003-3212110022132300-0130122313020320-1312110231103222-1323010311203203"></a>

<a id="canonical-2220331213020102-1021031133201310-3031221330201101-0312012321221233-3003011002022233-3212320110101012-2201303222101101-0233120000120022"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0230210330200133-0320123230300112-3302233131102333-1123220330132010-0222310100221211-2223010323103023-2101010222031302-2231303323133203"></a>

<a id="canonical-3223031131103301-3300303220301032-2013233110232132-1122233201332231-1231230332311112-0220013100220121-1113023310322121-3003111112030130"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
