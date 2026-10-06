---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-0230212322001230-3331102300103113-2323232333133232-3112023113220320-1103210000000301-0010022332202310-3010021110132210-2133102000122010"></a>

## `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.node` property

Type: `"string"`. Computed.

Node. Node name on this site.

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

<a id="canonical-0023300312002232-3313011201013002-3012111310212102-2230130201211100-2213212022200312-2330002021113101-1332223300323200-3222322012002211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- [segment_vrf.segment_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1303120210133032-0020301133000323-3311321220022302-3113222021233210-2213313002331302-1100333233332222-0121333220032113-2133332330221312)
- [segment_vrf.segment_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1013123331120122-0033112130321011-1333300111133331-1321221010013020-1000023032023001-2330032010300003-1210003001133331-2321100233112220)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1231220311131020-0222213111202120-3211201112032233-0310210332003311-1012332122003020-1212302300220032-3311102022222031-2332011132202103)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2111011003301202-3133233121210301-3322233332303203-2223312211310330-0200330203023310-3221102013202221-3321212012023211-1012312333112021)
- segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-3022302131021320-2201221321101200-1131213033032122-1010201012303013-1211330331021221-3120310001002023-3102302031231333-1002112301031010"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

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

<a id="canonical-3113223322213111-0223231022013230-0321203212211312-2323101311301101-3000031111203202-0111221001213200-3101032020112200-3203323131032001"></a>

### Direct properties for `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface`

<a id="canonical-2332010323001233-0033212230320312-0012320032131232-1323201212012311-0132121312301131-3220120131300211-1122012212333312-1210222332333301"></a>

#### `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface.kind` property

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

<a id="canonical-0223123221203302-1200310301132012-1122010132010120-3011130233100230-2023133032030120-1113023001030132-1111113111013330-1300113232233220"></a>

<a id="canonical-1223221323000233-3330033313221231-1302321132210330-2012133311113210-2020120212101133-2200223000222212-3332112220100033-2123320213100132"></a>

#### `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface.name` property

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

<a id="canonical-2111222213231023-3332311201223001-3021323333231030-2012120322000130-1221022333003031-3031210222122202-1211102121210221-0011230131220020"></a>

<a id="canonical-3332133213101003-1013223230112130-1100013231033213-0212133312203033-2031022322321232-3223300222232233-3320231021033010-1211010223013203"></a>

#### `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` property

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

<a id="canonical-0103313331021203-1213003020120022-0101013121111103-0330101332131022-3001333123023003-1330032300130212-2021112203301313-3313022313202221"></a>

<a id="canonical-3030231103003111-3020312333311200-1313023122213222-3321313010212231-0312023333310013-2003132213321111-3312020001310303-0002011212111223"></a>

#### `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` property

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

<a id="canonical-3321220133202023-2313120330201233-0023133220122003-1011133031320022-2102223002222333-0021320103130012-3002313122101100-2133202121113333"></a>

<a id="canonical-3130322311303223-1123121031233131-1110302102020303-2302011103133013-2323103321131202-2111103223222103-0111220132031320-0220231131330130"></a>

#### `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface.uid` property

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

<a id="canonical-0021113223132201-3022121331332032-3312123112022020-1203310012023130-2303101212300123-0001032222032213-1023013332311022-2231232033102112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_vrf.segment_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- segment_vrf.segment_network

<a id="canonical-0300320221222203-1033020102013121-2021120223120123-1123013320023122-2132101033100012-1222122212100313-0131233032002321-2213322222031102"></a>

Type: `"single"`. Computed.

This type establishes a 'direct reference' from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name for public API and Uid for private API This
type of reference is called direct because the relation is explicit and concrete (as opposed to
selector reference which builds a group based on labels of selectee objects)

<a id="canonical-3311000111020320-3133101310110200-2111011220103001-3110111032303203-3300111320223303-3101033200312322-2333220233003221-0021200223232102"></a>

### Direct properties for `segment_vrf.segment_network`

<a id="canonical-0113133220101213-2013133323020021-3232131101311221-3020010011232223-0101130333232211-3203122003123330-3012320320011301-2122110033331312"></a>

#### `segment_vrf.segment_network.kind` property

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

<a id="canonical-1100232212302112-2230021121122212-0203001110103233-0101310322300111-2332233030321031-0132013212130030-1121322010230121-0100033133100212"></a>

<a id="canonical-2202003022211221-2213322021300002-2100310232313011-2001223011210012-3311132302102023-0321003230113302-2111300311300020-0100302210321203"></a>

#### `segment_vrf.segment_network.name` property

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

<a id="canonical-1303311203120033-2231012200031110-2332021312001320-1202301032013212-3212230222113033-3330211113100202-0130300033022122-1212212000212010"></a>

<a id="canonical-1020320020223100-2022321333111210-1133131311123020-0200221132031311-0032200033100003-1212230011001031-1103100311121231-1030221001300313"></a>

#### `segment_vrf.segment_network.namespace` property

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

<a id="canonical-3303122030023312-3130010003201020-2231233011310123-1022020130323133-1033012312203000-2122101100312112-3000313321100112-1330221001123303"></a>

<a id="canonical-0122222332120123-0332233130220123-0231332032323213-1313301132003230-1232231113131200-2202311001013131-2010131232001320-1103210231312322"></a>

#### `segment_vrf.segment_network.tenant` property

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

<a id="canonical-1313120100332311-3120221010233230-0321332113212122-1212332212111201-2221301232132222-1321001022332131-1302002013231301-1010000131002102"></a>

<a id="canonical-1013110112303230-0023111232233332-0301311030012221-2111013121131332-2332130302132031-0031332003212033-3121200100231102-2203321202303010"></a>

#### `segment_vrf.segment_network.uid` property

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

<a id="canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_mesh_group_on_slo` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- site_mesh_group_on_slo

<a id="canonical-3221121221012221-1120312011123103-2232231111100313-0130001022010320-2112303310110222-0231322011230011-1013121213110102-1122313133201011"></a>

Type: `"single"`. Computed.

Select how the site mesh group will be connected. By default, public IPs of the control nodes of the
site will be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-site_mesh_group_choice": "[\"no_site_mesh_group\",\"site_mesh_group\"]",
  "x-ves-oneof-field-site_mesh_group_ip_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

<a id="canonical-0000033030322131-1123330122213033-2003103303102222-0200202303032031-1112122110013331-2321301132332200-1102311311201223-3322201021001233"></a>

### Direct properties for `site_mesh_group_on_slo`

- [no_site_mesh_group](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1331023023112310-1033013130131111-3120331103012113-0132333331321323-2320011313200210-0232023003301123-1202131331013033-3023122020220302): complete subsection reference.

- [site_mesh_group](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0310123133313310-3101121232031020-1302233301002121-1311111131332031-0220212001000223-0130032130110330-3121332202001300-1311033113113121): complete subsection reference.

- [sm_connection_public_ip](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0021212330003021-0231020110233323-1113133321030020-3002133131130110-3022123222213311-1302223331212110-1302131120032033-1302010200230222): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1222032311200220-2311221211201222-0221011112203212-3300123011013123-1123122011112211-2032010011311300-0033332111023213-2002023320101000): complete subsection reference.

<a id="canonical-1331023023112310-1033013130131111-3120331103012113-0132333331321323-2320011313200210-0232023003301123-1202131331013033-3023122020220302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_mesh_group_on_slo.no_site_mesh_group` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031)
- site_mesh_group_on_slo.no_site_mesh_group

<a id="canonical-3323330023302300-2312013002310001-0132301322100012-0100033310122203-2332310220213010-0302000120331313-2100003211301101-2331223200220003"></a>

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

<a id="canonical-0310123133313310-3101121232031020-1302233301002121-1311111131332031-0220212001000223-0130032130110330-3121332202001300-1311033113113121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_mesh_group_on_slo.site_mesh_group` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031)
- site_mesh_group_on_slo.site_mesh_group

<a id="canonical-1321102112031133-3302103200313300-2210220030221232-0033030310120132-1330331003032320-3221120333013130-3311031023133032-2220022223220233"></a>

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

<a id="canonical-2330013102030133-2312110332133211-2231103013100123-0223321333000301-0323022120222112-1010111023231031-3232110332023111-3301033022011133"></a>

### Direct properties for `site_mesh_group_on_slo.site_mesh_group`

<a id="canonical-0211002112203302-3300021232020332-3310031323000002-3201122232221132-0233231321220110-0033232022300032-3130321231220133-2133013103112022"></a>

#### `site_mesh_group_on_slo.site_mesh_group.name` property

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

<a id="canonical-2211210230111122-3303102031032312-2302021010101110-1323102233102122-2332102203110333-2020032131113113-0132000122101203-0012113211323020"></a>

<a id="canonical-3300323322022232-2303131130021031-0122120302020321-3010012033130133-1313011333133100-2121330303033220-3213111120133132-3211012221233011"></a>

#### `site_mesh_group_on_slo.site_mesh_group.namespace` property

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

<a id="canonical-0030223333011111-3132131001110000-0132203313321220-0130302133022031-2121102313121320-0133321001301211-3301131102013100-2000322333321120"></a>

<a id="canonical-1203133201100332-2033230112331302-0020213333012211-3103201131333323-3221000213233020-1111312213323333-0132123002100002-2222221230030332"></a>

#### `site_mesh_group_on_slo.site_mesh_group.tenant` property

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

<a id="canonical-0021212330003021-0231020110233323-1113133321030020-3002133131130110-3022123222213311-1302223331212110-1302131120032033-1302010200230222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_mesh_group_on_slo.sm_connection_public_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031)
- site_mesh_group_on_slo.sm_connection_public_ip

<a id="canonical-0111322032223002-1110313232312231-1210230113110231-0100123311113302-0032122111030213-3020201012133322-2113200321010330-0302113011100121"></a>

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

<a id="canonical-1222032311200220-2311221211201222-0221011112203212-3300123011013123-1123122011112211-2032010011311300-0033332111023213-2002023320101000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_mesh_group_on_slo.sm_connection_pvt_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031)
- site_mesh_group_on_slo.sm_connection_pvt_ip

<a id="canonical-0321003000331112-0200312122110200-1033010120202310-0032303120310330-0211012333331322-3132003010010201-3330123131313120-2021102030132323"></a>

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

<a id="canonical-0030100133010212-1200030111021312-3123032332031021-0010103012023120-2202203012321113-0113233110320102-2113301111200001-0103012002023131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upgrade_settings` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- upgrade_settings

<a id="canonical-3000003010233110-3122233323302221-3330021003012222-2202323100231222-2331122302332101-0210223033131103-2012230122031022-3013001013111221"></a>

Type: `"single"`. Computed.

Configuration parameter for upgrade settings.

Additional upstream details:

Specify how a site will be upgraded.

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

<a id="canonical-1011320210120000-0321202330123003-2012323301232312-3030023230033332-2230102100300321-0211002210023223-3230302301310130-2120032010202120"></a>

### Direct properties for `upgrade_settings`

- [kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1311312312123132-2132122001121023-2310232230000031-3331033300123111-2112333121031222-2121330032031221-0220303303333022-3300013222010111): complete subsection reference.

<a id="canonical-1311312312123132-2132122001121023-2310232230000031-3331033300123111-2112333121031222-2121330032031221-0220303303333022-3300013222010111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upgrade_settings.kubernetes_upgrade_drain` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [upgrade_settings](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0030100133010212-1200030111021312-3123032332031021-0010103012023120-2202203012321113-0113233110320102-2113301111200001-0103012002023131)
- upgrade_settings.kubernetes_upgrade_drain

<a id="canonical-3131120002001132-2202330000310201-3032300003223200-0313232213321102-2010233303320121-0022133231122000-2103330100310323-1121212100202111"></a>

Type: `"single"`. Computed.

Specify how worker nodes within a site will be upgraded.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

<a id="canonical-3331000032132302-0232223100331202-3312303313212121-2031300313123333-2321133133320021-2200111003113201-2212212330210010-0330112013322221"></a>

### Direct properties for `upgrade_settings.kubernetes_upgrade_drain`

- [disable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1101002301102122-2031301321121333-2111011110222102-2010003033323032-0102303100133321-0020123323032330-0333110320203310-0011221132211102): complete subsection reference.

- [enable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0021210223223221-1202110301313103-0001033232221120-3133022310111011-1230301200223222-3001001331031032-3021203121012132-2203311323331123): complete subsection reference.

<a id="canonical-1101002301102122-2031301321121333-2111011110222102-2010003033323032-0102303100133321-0020123323032330-0333110320203310-0011221132211102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [upgrade_settings](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0030100133010212-1200030111021312-3123032332031021-0010103012023120-2202203012321113-0113233110320102-2113301111200001-0103012002023131)
- [upgrade_settings.kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1311312312123132-2132122001121023-2310232230000031-3331033300123111-2112333121031222-2121330032031221-0220303303333022-3300013222010111)
- upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-3010223302103231-0121212223111333-0033200230113113-2312000210203033-1013031101231311-0300020320223111-0221320200233103-0233111000310333"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable upgrade drain.

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

<a id="canonical-0021210223223221-1202110301313103-0001033232221120-3133022310111011-1230301200223222-3001001331031032-3021203121012132-2203311323331123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [upgrade_settings](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0030100133010212-1200030111021312-3123032332031021-0010103012023120-2202203012321113-0113233110320102-2113301111200001-0103012002023131)
- [upgrade_settings.kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1311312312123132-2132122001121023-2310232230000031-3331033300123111-2112333121031222-2121330032031221-0220303303333022-3300013222010111)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-0300003311032023-3302312013333111-1013300221333331-1132000231211330-0131310303121002-3303100323302022-0221203131201212-1310113221312212"></a>

Type: `"single"`. Computed.

Specify batch upgrade settings for worker nodes within a site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-drain_max_unavailable_choice": "[\"drain_max_unavailable_node_count\", \"drain_max_unavailable_node_percentage\"]",
  "x-ves-oneof-field-vega_upgrade_mode_toggle_choice": "[\"disable_vega_upgrade_mode\",\"enable_vega_upgrade_mode\"]"
}
```

<a id="canonical-2001002323211133-2010100021023302-2202232201133203-0323112112110323-0310102010033120-1320321131100003-3131100221223230-3112030232010313"></a>

### Direct properties for `upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain`

- [disable_vega_upgrade_mode](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1000023132331110-0302231221120301-0201322113131332-0032111123321231-0233302023230311-0121012200211000-0201320301232210-2200300300312222): complete subsection reference.

<a id="canonical-0311322102003321-2301212210200030-0210212211333222-0323000203303023-1012321021101233-0300101132033131-1100013022323313-0310230211031113"></a>

<a id="canonical-3031012333310031-0312132030203002-0322313203232000-0012233210100122-3111232022200310-3112321303211302-0001313101203132-1301122323121201"></a>

#### `upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` property

Type: `"number"`. Computed.

Node Batch Size Count. Exclusive with \[\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
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
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-3333010022123302-3013231132331010-3332032200232101-3320222032120110-2003023013112232-2221020322210110-3330212223131313-2301212201221310"></a>

<a id="canonical-2203220230121202-1302312011033313-2110230333101032-0211312300201100-0310223010312230-2111320302010312-0313213210111032-1033203310200211"></a>

#### `upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` property

Type: `"number"`. Computed.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-1303010013002002-2321012312121223-2211203313030312-2123020130012330-1300021033102121-1201312132110022-0020332333123221-1121122321311301"></a>

<a id="canonical-3301323103333113-1333331111212121-1020011313023122-1012022132232003-0022103010103231-3210111022013312-0031022102103301-0330301112131030"></a>

#### `upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` property

Type: `"number"`. Computed.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 900,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  }
}
```

- [enable_vega_upgrade_mode](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2300333333220033-2221032002323131-1300130033000020-1312322233132110-1321210332101211-3032200330111121-0101011011010333-0132101013322210): complete subsection reference.

<a id="canonical-1000023132331110-0302231221120301-0201322113131332-0032111123321231-0233302023230311-0121012200211000-0201320301232210-2200300300312222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [upgrade_settings](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0030100133010212-1200030111021312-3123032332031021-0010103012023120-2202203012321113-0113233110320102-2113301111200001-0103012002023131)
- [upgrade_settings.kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1311312312123132-2132122001121023-2310232230000031-3331033300123111-2112333121031222-2121330032031221-0220303303333022-3300013222010111)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0021210223223221-1202110301313103-0001033232221120-3133022310111011-1230301200223222-3001001331031032-3021203121012132-2203311323331123)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-1330203022203012-1033220033123111-3000323122013033-0302123231303221-2000013012122033-0223310311003123-2000313133031133-3323231021211231"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable vega upgrade mode.

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

<a id="canonical-2300333333220033-2221032002323131-1300130033000020-1312322233132110-1321210332101211-3032200330111121-0101011011010333-0132101013322210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [upgrade_settings](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0030100133010212-1200030111021312-3123032332031021-0010103012023120-2202203012321113-0113233110320102-2113301111200001-0103012002023131)
- [upgrade_settings.kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1311312312123132-2132122001121023-2310232230000031-3331033300123111-2112333121031222-2121330032031221-0220303303333022-3300013222010111)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0021210223223221-1202110301313103-0001033232221120-3133022310111011-1230301200223222-3001001331031032-3021203121012132-2203311323331123)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-2121011223011123-1302123220211003-0332120032330212-2001113222130320-3110310330010032-3132212213013230-1330331013203313-3120012133321123"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable vega upgrade mode.

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

<a id="canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- vmware

<a id="canonical-1000002233221021-3103002001232013-0203230211301022-0002322012112022-0103232111201010-2202332322301210-2103102202330030-0012103321010332"></a>

Type: `"single"`. Computed.

VMware Provider Type. VMware Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-orchestration_choice": "[\"not_managed\"]"
}
```

<a id="canonical-1223003300002323-0122220231023121-2001311103221300-1302203011122312-2210122230211231-3023112121110301-3322101322021313-0203111022133332"></a>

### Direct properties for `vmware`

- [not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212): complete subsection reference.

<a id="canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- vmware.not_managed

<a id="canonical-1213021313222310-3121323213021023-3100100011030113-2203123332213102-3002212031223001-0332033323310330-2031023000300211-3233020032202103"></a>

Type: `"single"`. Computed.

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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

<a id="canonical-0223201211333321-1303031113222032-3113232320303000-3222002222131001-1021012330112330-2120333132210333-0211323111001020-1311022120330210"></a>

### Direct properties for `vmware.not_managed`

- [node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101): complete subsection reference.

<a id="canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- vmware.not_managed.node_list

<a id="canonical-3323211213321023-0300012331322202-3113132330121223-0202231132030323-3110203121001203-0030000122201002-3131331231202022-0322131213303333"></a>

Type: `"list"`. Computed.

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3112222222102022-2332330131012312-3202200231002212-0000312113301003-0011230101210013-0020301000030032-0112020001002111-2230013131320122"></a>

### Direct properties for `vmware.not_managed.node_list`

<a id="canonical-1330320300132102-3003201133213233-3033200223322102-0210033321200302-1301001323031110-3200020101022131-1122200023320130-2220120213330011"></a>

#### `vmware.not_managed.node_list.hostname` property

Type: `"string"`. Computed.

Hostname. Hostname for this Node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313): complete subsection reference.

<a id="canonical-3112312130031131-3113231233213033-3130130000301210-3020123212000211-2002331212010203-2021320211102210-1300123221132101-2220012210310310"></a>

<a id="canonical-2231332332323222-1121231032210003-0311131131101230-0003103110122212-1203332310222123-2102102323023221-2311303013202233-0331203200322220"></a>

#### `vmware.not_managed.node_list.public_ip` property

Type: `"string"`. Computed.

Public IP. Public IP for this Node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3331232010122312-0303012232331222-1321001323220011-2021101333103201-0122331230301022-0322200330012301-2133002122213310-2320100122322002"></a>

<a id="canonical-2322202313322310-2333202222112131-3331101111111032-0133022221121101-0210302323120230-0233223333223321-0231103002212212-0121320311121013"></a>

#### `vmware.not_managed.node_list.type` property

Type: `"string"`. Computed.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "Control",
    "Worker"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- vmware.not_managed.node_list.interface_list

<a id="canonical-2303201333320202-2033221000121101-0101033200102132-0202022023021013-3200231131132103-2310012001220221-3320133020212202-2130102201120311"></a>

Type: `"list"`. Computed.

Manage interfaces belonging to this node.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0322212323120203-1030311000110100-0103121133133201-1300003011102123-2333220133033330-2300012231013032-0102221122200303-3122201300012332"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list`

- [bond_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3113233331323212-0201311132302221-0323311320033121-1221020221221020-3312022121220131-1310230333121213-3312122212120123-3133323202010303): complete subsection reference.

<a id="canonical-2113032123123233-2223013301330211-0311210331211323-0202023223300112-2301003313210031-2102321100101301-0203300331203330-2032120210300102"></a>

<a id="canonical-1111032320203000-1030130210010223-2111000200311122-1302002012003300-0322232101101212-0010221121120202-1232111130322130-2013010020311022"></a>

#### `vmware.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2002000010203020-1313101303100132-3221022203001132-2310200303013001-2231130113003003-2000023321030001-0130012302212201-1131013323110103): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--reference--group-018.md#canonical-2131211010200313-3110301321130222-1102123200113032-2221200010222312-3311330101103223-1332332132131213-0022022112111000-0312023302203233): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322): complete subsection reference.

<a id="canonical-2121201231002132-1031200301000122-2000000202332123-1110132222130313-3020200103123120-2022213210133032-3332333132021302-3212033302231121"></a>

<a id="canonical-0200102013121121-1001033221332321-0130032233221120-0033113021200323-0021103311203232-0130132302201102-2120223123001321-2033233312311021"></a>

#### `vmware.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-2311010321323012-1013333123112010-3333211220231113-2311123003312212-0203100303101301-2311120321031123-1002001020331121-3003202100300110"></a>

<a id="canonical-3220310002312222-0313133110212213-2000233331303201-0003310321110321-2120131122113102-3020210033201320-0111020012202210-2331223213031203"></a>

#### `vmware.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-1003201320200111-0312010010101233-0323101202003131-3210012010113032-0003223010320132-0221133330020033-0110331123112003-3221111201201201"></a>

<a id="canonical-0120121123333130-0313023330232200-1022301021303321-2332031213301313-0310330101310122-3232102221100032-2030132022001131-1302110030103310"></a>

#### `vmware.not_managed.node_list.interface_list.labels` property

Type: `["map", "string"]`. Computed.

Add Labels for this Interface, these labels can be used in firewall policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
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
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "64",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 64,
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
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [monitor](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3000311103023222-2111102210333112-0020030303122131-2213331012001330-2122023121013123-1302301120133213-1000011210233330-1002330303330200): complete subsection reference.

- [monitor_disabled](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3111220120002011-0130022213201132-0230331301200221-3221122101201330-2230002300332002-2030010032220330-2113133003302200-1121030221220201): complete subsection reference.

<a id="canonical-0221211210113233-0021203212212332-0023320201021110-2110012010211200-3211300220010212-3332120122131020-0003032210021313-3302031011231133"></a>

<a id="canonical-0003012020113133-3133031020332101-2113122010101132-2202201001311232-2203210330030123-0112230202000330-2302301000003313-3322220220003221"></a>

#### `vmware.not_managed.node_list.interface_list.mtu` property

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-2002203023213212-2221112003011030-3323323320031131-2311333033323020-1212202212001333-3112031302111201-1002020123133102-0030110310031220"></a>

<a id="canonical-3011311231211103-2232300320121311-1303310003001310-1110313201100130-1303033010131023-3322320222112321-2011230012031021-2223221022230023"></a>

#### `vmware.not_managed.node_list.interface_list.name` property

Type: `"string"`. Computed.

Interface Name. Name of this Interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [network_option](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3320130231323321-3022122100223033-1022132021022321-2011100212203002-0033031330120212-0212210003112230-2202320210023220-3322202133132320): complete subsection reference.

- [no_ipv4_address](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1311301011230121-2103012100303202-2231012122020031-2032112300013302-0112321213131123-3032030000103020-1100121311123230-2033303223003012): complete subsection reference.

- [no_ipv6_address](data-sources--securemesh_site_v2--reference--group-018.md#canonical-2221230121110112-0000022310030300-3331023302031310-2301013022330330-3222121321032112-3330122131020130-1322003213222330-3223020113112010): complete subsection reference.

<a id="canonical-0220101000022202-3012021030213223-3233210310120311-3121331321202320-3033301333230320-2202121130320023-2102210312213303-1132301130330121"></a>

<a id="canonical-0300331122322331-1120232301112131-3023123120313302-2233302120022023-1121202320331213-1222331300020231-3003010323102003-0020012102303103"></a>

#### `vmware.not_managed.node_list.interface_list.priority` property

Type: `"number"`. Computed.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-018.md#canonical-2121211122013120-3103203013202130-0310310221301133-3300123122110130-2321201321123323-0200313022223020-0323030000022030-0322311230123100): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-018.md#canonical-2030023132322101-2011112010030122-0311133322003030-1133300113311203-2103001023133200-0123113110313232-0323002030100222-0231102301022003): complete subsection reference.

- [static_ip](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0322013101212310-2222330010320203-0121011121312102-0210101001022231-1220313221113220-1031320022212303-0320003022130210-3333003000030110): complete subsection reference.

- [static_ipv6_address](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0101212013130321-0131302111333332-2312131200130323-3111003310322033-0033230100220132-3322223222230132-0011121000101030-0332100211010113): complete subsection reference.

- [vlan_interface](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0230030312120110-2020013303020100-2023331011232031-0302302103032030-2321322301231232-2223121203332121-2210203321013231-3032330030011313): complete subsection reference.

<a id="canonical-3113233331323212-0201311132302221-0323311320033121-1221020221221020-3312022121220131-1310230333121213-3312122212120123-3133323202010303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.bond_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.bond_interface

<a id="canonical-2012002022300012-1022323203003101-2230130102103113-3203123302112203-3313301113200232-1021103331110112-2222000102222022-2000203032102230"></a>

Type: `"single"`. Computed.

Configuration parameter for bond interface.

Additional upstream details:

Bond devices configuration for fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

<a id="canonical-0030110122010013-3010211212020201-0323200301111001-3210303023313231-1203311031131122-2133221031332110-3123110301130323-0120233120132311"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.bond_interface`

- [active_backup](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2031330002323032-2123111311232010-1120303302202311-3121022123032202-1003330232223031-2330103300122002-1222220222211303-0302010012132232): complete subsection reference.

<a id="canonical-2321300120021010-1223111133312223-1203103131320121-1233012300200121-0100203321200011-0123131212301001-1001021121231302-3302113232231131"></a>

<a id="canonical-0023013301233111-3023322223030221-3233211320301010-1003311131222303-2023110002122321-2302111203310232-1010011222123333-3002120131332122"></a>

#### `vmware.not_managed.node_list.interface_list.bond_interface.devices` property

Type: `["list", "string"]`. Computed.

Ethernet devices that will make up this bond.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [lacp](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2130120320303322-3321320202312321-2021131333120113-2332323013211201-0132230023220011-0232221021323110-1010011312210232-3012030023222333): complete subsection reference.

<a id="canonical-3311312023213303-2323300012021110-0300321303303131-1210201312233331-0232312313230020-2231211123333011-2032000003023113-0302213233112113"></a>

<a id="canonical-1223020321132133-1131332011232223-2310010130332032-1033313110233120-1013110032002212-0221023201131312-1221331102300333-1223300111202111"></a>

#### `vmware.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

Type: `"number"`. Computed.

Link Polling Interval. Link polling interval in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 500
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-2012222002320333-3022023221212212-3220311120111220-2023121232030302-2121210110132001-2012113031031113-0333002123100203-2331111030332231"></a>

<a id="canonical-3031330201313120-0101332332130210-0332302223032102-0120010230211202-1321312112321133-1002323212002233-3123101130333230-0233100213011220"></a>

#### `vmware.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

Type: `"number"`. Computed.

Milliseconds wait before link is declared up.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  }
}
```

<a id="canonical-0130002211331232-2031303020101030-2022132332311202-0110133120011030-1320002033101211-3323210113010110-0333223030103202-3200031011221221"></a>

<a id="canonical-0031020102321022-3301333301321022-3322222022101130-1022130212311122-3331022332213130-2102100132220132-1132323200331100-2213230000333223"></a>

#### `vmware.not_managed.node_list.interface_list.bond_interface.name` property

Type: `"string"`. Computed.

Bond Device Name. Name for the Bond. Ex 'bond0'

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-2031330002323032-2123111311232010-1120303302202311-3121022123032202-1003330232223031-2330103300122002-1222220222211303-0302010012132232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3113233331323212-0201311132302221-0323311320033121-1221020221221020-3312022121220131-1310230333121213-3312122212120123-3133323202010303)
- vmware.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-0103023233200111-2333211312120133-3200111111300111-0001111302021021-0330303203213032-1311012330231231-2021011033202113-0310103020000223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for active backup.

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

<a id="canonical-2130120320303322-3321320202312321-2021131333120113-2332323013211201-0132230023220011-0232221021323110-1010011312210232-3012030023222333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3113233331323212-0201311132302221-0323311320033121-1221020221221020-3312022121220131-1310230333121213-3312122212120123-3133323202010303)
- vmware.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-1113101212311011-3232300320032032-1302330033313021-2121130002031313-1203030211131022-1133010232123332-1012031110100313-1012010221113003"></a>

Type: `"single"`. Computed.

LACP parameters. LACP parameters for the bond device.

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

<a id="canonical-1323100330333313-1213302131022301-1132213102301012-3103302110231031-3123332010212200-2301012032212302-3033133300200330-3131230223012330"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-1200230302122301-2232120123323111-3033223220312333-1331330311310230-3103211112030100-0233220200130001-0100002202330000-1120033223331311"></a>

#### `vmware.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

Type: `"number"`. Computed.

Interval in seconds to transmit LACP packets.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-2002000010203020-1313101303100132-3221022203001132-2310200303013001-2231130113003003-2000023321030001-0130012302212201-1131013323110103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-0232132301021011-3131023000002323-1101101013002211-2130023331222021-2120001210021033-0311233211302012-3003323133023122-3001213322123032"></a>

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

<a id="canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-3133210230132030-0301122322021130-3030203312030010-3130320221300101-3210111320223031-0201112212331031-3211131312111202-0022012322001033"></a>

Type: `"single"`. Computed.

DHCPServerParametersType.

Additional upstream details:

DHCP server configuration for this interface.

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

<a id="canonical-3031302030220313-0310313330212130-3103122321332011-2130203102122130-1310212233302210-1113121023033300-0010112123030002-3030123212331111"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0320110132001232-1212113133132310-1032110023222132-3203003230233032-1111202032101132-1233313000331011-2101322222213322-2021003032002031): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1303101120331011-1120100003103333-1002213010101100-3230221311032330-0312301003032301-3013232020203303-0010220000322023-2132001030010320): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3312332001302203-1100203030010230-1023021103011201-1132103011022313-3211122221323002-1211022311330013-2003133301021121-3110130232123100): complete subsection reference.

<a id="canonical-1211333103231232-3320301112130301-3010100101330231-2210112023011030-2021223120210002-1212101331331311-3333033030221101-2133100022113020"></a>

<a id="canonical-3201010300223321-3301021303231000-2021211132221320-3113022132112220-0302000111101311-3332131211321020-1231122122330211-2233011022213230"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-1003311022212130-2032233003023000-1103111311121133-0111123113101332-0231031220320303-1001300211130122-1103100203031013-0213120022300310"></a>

<a id="canonical-2223221021331021-3100331312113020-1022211133323201-3211333321313231-3231111001131320-3310013133122211-2312002112330100-0002231103001311"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

Type: `["map", "string"]`. Computed.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
    },
    "category": "discovery",
    "constraintType": "map",
    "crossEntry": {
      "uniqueValues": true
    },
    "deterministic": true,
    "keys": {
      "format": "mac-address",
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.mac": "true",
      "ves.io.schema.rules.map.max_pairs": "128",
      "ves.io.schema.rules.map.unique_values": "true",
      "ves.io.schema.rules.map.values.string.ipv4": "true"
    },
    "values": {
      "format": "ipv4",
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
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2021031130232120-3200220000131200-1001020120200301-0023113313303313-0030333112223332-0111302212212100-2133112222330323-1311023021103331): complete subsection reference.

<a id="canonical-0320110132001232-1212113133132310-1032110023222132-3203003230233032-1111202032101132-1233313000331011-2101322222213322-2021003032002031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130)
- vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-2321330131002131-3022203130211323-3300332102222311-0021311321012111-3103210022310201-3110013300023020-2011223133211313-3000133211201102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from end.

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

<a id="canonical-1303101120331011-1120100003103333-1002213010101100-3230221311032330-0312301003032301-3013232020203303-0010220000322023-2132001030010320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130)
- vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-2112120133131110-2231210030023232-1023022320102012-1122233000303222-2201332002030203-2122212331122223-0111333100333213-0130311121302200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from start.

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

<a id="canonical-3312332001302203-1100203030010230-1023021103011201-1132103011022313-3211122221323002-1211022311330013-2003133301021121-3110130232123100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-1312112131012232-1231211331203200-0130001311201300-2133022002232021-0112321102313212-2310210121103000-0120033213303211-0322120310330022"></a>

Type: `"list"`. Computed.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0010212331123133-0212100033201331-3321331210203002-0122323103100212-0013023133311321-3223010113012310-0030302023223223-3333100100331121"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-1210300111112022-0210011002123011-3331112010021131-0022303133330122-3011223233212010-2223033323322311-0232210310232032-0230010223230121"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1000231312231213-1022001033223012-1202130321002031-3010303202011213-0231223322111023-1303022320203232-2021000103313003-1110333003013320"></a>

<a id="canonical-1220202300111230-2312023103121010-3111233233112232-3201231102320331-0123110013322213-1103133331121221-1130210330301021-2103031130110220"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

Type: `"string"`. Computed.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [first_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0101101323233310-2231003203210003-3023100300130303-1130122031133212-2321200331300001-2021222221231000-3230001203330120-1122201221110103): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1323110321213113-2203323132221100-3122210220311030-1100010211103331-0010200031122321-0333111133302011-1023320030221133-2111233321210023): complete subsection reference.

<a id="canonical-0001120032311222-0301013013221303-1223003311303202-1322202213211220-1111311110111102-2331212223113312-2320113030110300-2312213332013000"></a>

<a id="canonical-3313220022101112-1322333233010201-2231013220132312-1200021332321332-2122010002110312-0323030013331012-1032220111303231-2221222120311121"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

Type: `"string"`. Computed.

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-0212230110300312-1020321110231000-2322313203010011-3111013122200130-1231211313313331-0231211102012301-1222230332211020-2302212331001222"></a>

<a id="canonical-0311322013213312-0202012220121313-1113322210030301-3330211333000123-1230212023101330-1320331202021131-3310311132331211-3013200130302002"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

Type: `"string"`. Computed.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1102133112010213-0320030302302113-2212101002231010-2312003311100000-0033333103323102-2332211110210133-2121222220223213-1110110111230133): complete subsection reference.

- [same_as_dgw](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3233112100230330-3202012031002311-2302300131131010-2320310011303020-2000320302121032-1222302011233320-0021300110003031-2122103332103031): complete subsection reference.

<a id="canonical-0101101323233310-2231003203210003-3023100300130303-1130122031133212-2321200331300001-2021222221231000-3230001203330120-1122201221110103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3312332001302203-1100203030010230-1023021103011201-1132103011022313-3211122221323002-1211022311330013-2003133301021121-3110130232123100)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-1202331221333112-2110003023210011-1123231030111103-0010332013321122-0320232230211230-2233210101031321-2302230301132003-3132213131032133"></a>

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

<a id="canonical-1323110321213113-2203323132221100-3122210220311030-1100010211103331-0010200031122321-0333111133302011-1023320030221133-2111233321210023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3312332001302203-1100203030010230-1023021103011201-1132103011022313-3211122221323002-1211022311330013-2003133301021121-3110130232123100)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-0010321303033311-0323130113313320-3231312101002130-3131131000023001-1121320321002013-1210310030011112-1111310101123032-0313132010202132"></a>

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

<a id="canonical-1102133112010213-0320030302302113-2212101002231010-2312003311100000-0033333103323102-2332211110210133-2121222220223213-1110110111230133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3312332001302203-1100203030010230-1023021103011201-1132103011022313-3211122221323002-1211022311330013-2003133301021121-3110130232123100)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-2032313300120012-0131120023213111-2211131011212102-0331310010132320-2033023223003223-1011013233032020-2123131301003031-1303301132022300"></a>

Type: `"list"`. Computed.

List of non overlapping IP address ranges.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
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

<a id="canonical-1313122312320130-0113221123120102-3111002031330231-0203020013032232-1320333320030330-3232332301110212-1232312323003011-2011210213221220"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools`

<a id="canonical-3212032203121033-3121033233103022-3313323022201132-2331021301023311-1203030302103222-0331321230010003-1333100132213213-1031333031231311"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.end_ip` property

Type: `"string"`. Computed.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2323132213323203-1113021200313030-0003131232112303-2110013323110233-1233213323103332-2110031002200033-0010122101300020-2000013220012111"></a>

<a id="canonical-3112201030133110-2112210320032123-0223212101313012-2010021131313332-3312032113130312-2220103230001313-3333021221333212-3221221220312100"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.exclude` property

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-1321230122111333-2203011111303320-2211303301020000-1130213312202031-2033123303010311-0330220200200031-2113020033301330-2102320332123021"></a>

<a id="canonical-1032011001201312-2332133222021120-3003310032012211-2131303220132301-1011221010321321-0000110020312231-0221123131010023-3312002233312123"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.start_ip` property

Type: `"string"`. Computed.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3233112100230330-3202012031002311-2302300131131010-2320310011303020-2000320302121032-1222302011233320-0021300110003031-2122103332103031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3312332001302203-1100203030010230-1023021103011201-1132103011022313-3211122221323002-1211022311330013-2003133301021121-3110130232123100)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-2023222211032222-2201212103300013-3032211013110001-1030123103200232-3010231132200023-0301013332031312-3023021310322331-1133331332233011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for same as dgw.

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

<a id="canonical-2021031130232120-3200220000131200-1001020120200301-0023113313303313-0030333112223332-0111302212212100-2133112222330323-1311023021103331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130)
- vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-2222201311331210-0302001332100030-2301320011102222-1331230223022311-3221311010133212-2033003332203221-2230302220320330-2111202232033131"></a>

Type: `"single"`. Computed.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

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

<a id="canonical-1133123212311311-3220220213022132-2022212312012132-3212111201333122-1210333121222222-2113230212201323-2300230110030303-3120220031021233"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map`

<a id="canonical-3333010210102310-0032333021033000-3022032133103123-1033123000010021-1020032123203223-2010201332311030-2213130123223301-3031301010031111"></a>

#### `vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Computed.

Specify static IPv4 addresses per site:node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.ipv4": "true"
    },
    "values": {
      "format": "ipv4",
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```
