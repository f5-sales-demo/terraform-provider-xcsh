---
page_title: "xcsh_virtual_network reference"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_virtual_network reference."
---

# xcsh_virtual_network reference

<a id="canonical-2130101021323011-3002330133321312-3110103001302103-2133320321012310-3010320330020032-1032031113023113-3330002130320310-0200303030321113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-2112023310021033-2012123022032122-0030112331222000-3101313021113310-0300312330333113-3110103330132201-2330310010223213-3013312303120122)
- Property reference

<a id="canonical-0000120210001033-0210223112210010-3320113101012223-1333313321230130-2012010021000300-0103331303302102-0232030210120203-1221302322131313"></a>

### Direct properties for `xcsh_virtual_network`

<a id="canonical-1201100201033131-2231213213321103-3222233311320011-3331331321130031-1331201100332030-0201311211033331-2202230133203121-1311231100203020"></a>

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

<a id="canonical-1303012121303020-3033103030330112-1013012223130202-3101321213103221-2202001112033121-3103210023230110-1133031213321032-3301102233011330"></a>

<a id="canonical-0032130322030301-2101332200003120-0123131033311101-2030310123333013-2010233301121113-1222123132332333-1300213110023212-2201302321333232"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the VirtualNetwork.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [global_network](data-sources--virtual_network--reference--group-001.md#canonical-1332013022320101-1021031123100220-2033102232003030-3201221313323023-2302312133113002-0113321030111130-3313002330222000-3033021020231011): complete subsection reference.

<a id="canonical-0033032322331033-2021312032311003-2101131200010031-1103200222213010-0320112332123211-2003122332130131-1021332102333322-0122222201101232"></a>

<a id="canonical-3110110122312123-0220011210001003-2320201202113033-2013322033011310-0213121122001330-0303201230200220-1022000123121120-3302100303231330"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2221121033103011-0331201232121333-3121303321033120-0312130121023032-1120102222100233-1300231303122233-0130111300000022-2010310021201130"></a>

<a id="canonical-0013032312302112-1021303222122221-3311223230202133-0203210233211202-1132310331100233-0102201100210202-3020030303132233-3112021110001031"></a>

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

<a id="canonical-0303130121022112-2212110202221102-1333133232010022-2022010022313203-0333203210321033-2231003230001210-0202103332121010-0213121113232331"></a>

<a id="canonical-2231010322302133-3322223020103003-1133303202201102-3222231323220113-3213213000003223-1003000230330120-2203222233111331-2310331230131310"></a>

#### `name` property

Type: `"string"`. Required.

Name of the VirtualNetwork.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3322013012132030-0003302312200210-2101232120332000-1233202002031320-3203032222123021-1301100310222331-3303112232321133-0120330110232001"></a>

<a id="canonical-1112122320230031-0313210200131223-3232031030213333-0023012223021022-2321232010311230-1021123133222003-1111331333013320-2202122131321020"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace where the VirtualNetwork exists.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [site_local_inside_network](data-sources--virtual_network--reference--group-001.md#canonical-2020011021121001-2021003222110303-1303220333033001-0002300113101021-1310201312312201-1331200230323001-1033110332330302-3033311320320022): complete subsection reference.

- [site_local_network](data-sources--virtual_network--reference--group-001.md#canonical-0021131112022130-2303120033303331-1112013031030100-1222221101001130-3221130212300323-1020002201031323-1003130311113210-0211332332033212): complete subsection reference.

- [static_routes](data-sources--virtual_network--reference--group-001.md#canonical-3021012132213221-1330322221222333-3321310120010031-1221230231213120-1100203011130122-2330113131123330-2012322323011111-1211200322300231): complete subsection reference.

<a id="canonical-3232311302313131-2111213321101302-3111210203030231-1000221222220301-2131303121130120-1031012203120231-1232211321302301-0030312310133222"></a>

### All schema paths for `xcsh_virtual_network`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--virtual_network--reference--group-001.md#canonical-1201100201033131-2231213213321103-3222233311320011-3331331321130031-1331201100332030-0201311211033331-2202230133203121-1311231100203020) |
| `description` | [description](data-sources--virtual_network--reference--group-001.md#canonical-1303012121303020-3033103030330112-1013012223130202-3101321213103221-2202001112033121-3103210023230110-1133031213321032-3301102233011330) |
| `global_network` | [global_network](data-sources--virtual_network--reference--group-001.md#canonical-0302012002030030-3313233131002030-2132121021313131-0210030300322311-2011010321113221-1022232232000312-1031020111130211-2231000022303330) |
| `id` | [ID](data-sources--virtual_network--reference--group-001.md#canonical-0033032322331033-2021312032311003-2101131200010031-1103200222213010-0320112332123211-2003122332130131-1021332102333322-0122222201101232) |
| `labels` | [labels](data-sources--virtual_network--reference--group-001.md#canonical-2221121033103011-0331201232121333-3121303321033120-0312130121023032-1120102222100233-1300231303122233-0130111300000022-2010310021201130) |
| `name` | [name](data-sources--virtual_network--reference--group-001.md#canonical-0303130121022112-2212110202221102-1333133232010022-2022010022313203-0333203210321033-2231003230001210-0202103332121010-0213121113232331) |
| `namespace` | [namespace](data-sources--virtual_network--reference--group-001.md#canonical-3322013012132030-0003302312200210-2101232120332000-1233202002031320-3203032222123021-1301100310222331-3303112232321133-0120330110232001) |
| `site_local_inside_network` | [site_local_inside_network](data-sources--virtual_network--reference--group-001.md#canonical-2212132032220032-0102312101330123-3201130101313331-3012100232003113-2201313121110213-1120022122012121-1302120300211231-1320330120110031) |
| `site_local_network` | [site_local_network](data-sources--virtual_network--reference--group-001.md#canonical-2200030030120311-1112020331312220-1101312112330332-2320332111232301-1021013330130330-3020230110010121-3030200213331003-0023130120031012) |
| `static_routes` | [static_routes](data-sources--virtual_network--reference--group-001.md#canonical-0212212330321302-1321132223330022-0123232000333010-3321103321111012-2022100000311313-3103113311213332-2330203030110013-1023132103310210) |
| `static_routes.attrs` | [static_routes.attrs](data-sources--virtual_network--reference--group-001.md#canonical-0222312103111321-2103032122023000-2013303223130302-3020022211301033-0120120000300203-0121200132330021-0020002300103312-2323231323110320) |
| `static_routes.default_gateway` | [static_routes.default_gateway](data-sources--virtual_network--reference--group-001.md#canonical-2001213212032023-3100321213310001-3110013320013122-0133001132301112-2023113020122023-2222321101123033-0212012203122111-2313031112331310) |
| `static_routes.ip_address` | [static_routes.ip_address](data-sources--virtual_network--reference--group-001.md#canonical-0131222300233220-2120220223330132-2212320033320331-3022023010322322-1310002333203321-0120303011022112-2203003222233220-1111000132001302) |
| `static_routes.ip_prefixes` | [static_routes.ip_prefixes](data-sources--virtual_network--reference--group-001.md#canonical-0120111330123332-1011123021302030-3221202333011302-1121033030001231-1222200220200001-3320011330322210-0221003123203111-1321103000030320) |
| `static_routes.node_interface` | [static_routes.node_interface](data-sources--virtual_network--reference--group-001.md#canonical-0111311023021103-1003121021212321-3102102112301222-2210031333300312-3310201123231113-1032011313031002-1110221200123110-0201322201121210) |
| `static_routes.node_interface.list` | [static_routes.node_interface.list](data-sources--virtual_network--reference--group-001.md#canonical-2303331213303132-1002002112002331-2032012232031022-3133003202213202-3232032131011112-2121102112012231-3220233020022032-1110101130033333) |
| `static_routes.node_interface.list.interface` | [static_routes.node_interface.list.interface](data-sources--virtual_network--reference--group-001.md#canonical-3030002031301133-3100233323220200-0313311130011112-2132201230201310-2312102320031203-1313003013332002-2101002312201322-0220321310013333) |
| `static_routes.node_interface.list.interface.kind` | [static_routes.node_interface.list.interface.kind](data-sources--virtual_network--reference--group-001.md#canonical-0133301320232012-2102203122231123-2111302213022211-3330102231300130-1110121023313120-0020312120213301-2122302113213332-0310200303001122) |
| `static_routes.node_interface.list.interface.name` | [static_routes.node_interface.list.interface.name](data-sources--virtual_network--reference--group-001.md#canonical-2002023201021113-1211112130012103-2320223022320003-1021301000012101-2110212012100133-3213231103202133-2133022002001111-1013301310302021) |
| `static_routes.node_interface.list.interface.namespace` | [static_routes.node_interface.list.interface.namespace](data-sources--virtual_network--reference--group-001.md#canonical-0031010111121231-0333131000230331-0123021000213033-0100223032200203-2122010203201322-1213001200100312-1111013231222033-3032030231000233) |
| `static_routes.node_interface.list.interface.tenant` | [static_routes.node_interface.list.interface.tenant](data-sources--virtual_network--reference--group-001.md#canonical-2232021011112332-2030102330321102-0030320033211002-3110102203121102-3320110222312232-0033212330121133-0100212310033200-2301121231100113) |
| `static_routes.node_interface.list.interface.uid` | [static_routes.node_interface.list.interface.uid](data-sources--virtual_network--reference--group-001.md#canonical-2120032131133212-0212221323113102-0300332002003120-2102120123002212-3010232200003001-0112122003213121-3100012022300001-3331021021310133) |
| `static_routes.node_interface.list.node` | [static_routes.node_interface.list.node](data-sources--virtual_network--reference--group-001.md#canonical-2223121332232330-2131212321200102-3333210030211201-1113220321321121-2231012203102303-2203122203200333-1332320202032202-3021210103100123) |

<a id="canonical-1332013022320101-1021031123100220-2033102232003030-3201221313323023-2302312133113002-0113321030111130-3313002330222000-3033021020231011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `global_network` properties

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-2112023310021033-2012123022032122-0030112331222000-3101313021113310-0300312330333113-3110103330132201-2330310010223213-3013312303120122)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-2130101021323011-3002330133321312-3110103001302103-2133320321012310-3010320330020032-1032031113023113-3330002130320310-0200303030321113)
- global_network

<a id="canonical-0302012002030030-3313233131002030-2132121021313131-0210030300322311-2011010321113221-1022232232000312-1031020111130211-2231000022303330"></a>

Type: `["object", {}]`. Computed.

\[OneOf: global\_network, site\_local\_inside\_network, site\_local\_network\] Select the global
virtual-network scope for connectivity across participating sites.

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

- [global_network](data-sources--virtual_network--reference--group-001.md#canonical-0302012002030030-3313233131002030-2132121021313131-0210030300322311-2011010321113221-1022232232000312-1031020111130211-2231000022303330)
- [site_local_inside_network](data-sources--virtual_network--reference--group-001.md#canonical-2212132032220032-0102312101330123-3201130101313331-3012100232003113-2201313121110213-1120022122012121-1302120300211231-1320330120110031)
- [site_local_network](data-sources--virtual_network--reference--group-001.md#canonical-2200030030120311-1112020331312220-1101312112330332-2320332111232301-1021013330130330-3020230110010121-3030200213331003-0023130120031012)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020011021121001-2021003222110303-1303220333033001-0002300113101021-1310201312312201-1331200230323001-1033110332330302-3033311320320022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_local_inside_network` properties

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-2112023310021033-2012123022032122-0030112331222000-3101313021113310-0300312330333113-3110103330132201-2330310010223213-3013312303120122)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-2130101021323011-3002330133321312-3110103001302103-2133320321012310-3010320330020032-1032031113023113-3330002130320310-0200303030321113)
- site_local_inside_network

<a id="canonical-2212132032220032-0102312101330123-3201130101313331-3012100232003113-2201313121110213-1120022122012121-1302120300211231-1320330120110031"></a>

Type: `["object", {}]`. Computed.

Select the site-local inside network for site-internal connectivity.

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

<a id="canonical-0021131112022130-2303120033303331-1112013031030100-1222221101001130-3221130212300323-1020002201031323-1003130311113210-0211332332033212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_local_network` properties

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-2112023310021033-2012123022032122-0030112331222000-3101313021113310-0300312330333113-3110103330132201-2330310010223213-3013312303120122)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-2130101021323011-3002330133321312-3110103001302103-2133320321012310-3010320330020032-1032031113023113-3330002130320310-0200303030321113)
- site_local_network

<a id="canonical-2200030030120311-1112020331312220-1101312112330332-2320332111232301-1021013330130330-3020230110010121-3030200213331003-0023130120031012"></a>

Type: `["object", {}]`. Computed.

Select a site-local virtual network when connectivity must remain within one site.

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

<a id="canonical-3021012132213221-1330322221222333-3321310120010031-1221230231213120-1100203011130122-2330113131123330-2012322323011111-1211200322300231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `static_routes` properties

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-2112023310021033-2012123022032122-0030112331222000-3101313021113310-0300312330333113-3110103330132201-2330310010223213-3013312303120122)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-2130101021323011-3002330133321312-3110103001302103-2133320321012310-3010320330020032-1032031113023113-3330002130320310-0200303030321113)
- static_routes

<a id="canonical-0212212330321302-1321132223330022-0123232000333010-3321103321111012-2022100000311313-3103113311213332-2330203030110013-1023132103310210"></a>

Type: `"list"`. Computed.

List of static routes on the virtual network.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 165,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 165,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "165",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "165",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3110333203233000-2001222011332302-3001303300301000-1032111010133003-0012113113031000-1312333330030000-0300021100301200-0320321301030130"></a>

### Direct properties for `static_routes`

<a id="canonical-0222312103111321-2103032122023000-2013303223130302-3020022211301033-0120120000300203-0121200132330021-0020002300103312-2323231323110320"></a>

#### `static_routes.attrs` property

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--virtual_network--reference--group-001.md#canonical-1320030033203322-0122333002112002-3211111203121001-1113020333313023-3010312021102220-3230103111103002-2120221002032302-3323133300013332): complete subsection reference.

<a id="canonical-0131222300233220-2120220223330132-2212320033320331-3022023010322322-1310002333203321-0120303011022112-2203003222233220-1111000132001302"></a>

<a id="canonical-0011110101303211-2120102202132133-3032112131210233-3312311121220031-0010030032203211-1023123031213220-0103122122100133-0221213330230223"></a>

#### `static_routes.ip_address` property

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-0120111330123332-1011123021302030-3221202333011302-1121033030001231-1222200220200001-3320011330322210-0221003123203111-1321103000030320"></a>

<a id="canonical-1320132211133221-0120111200000231-0222321330013031-0332030112221231-2033300202033203-0112222221313213-3001101321220020-3303121211221123"></a>

#### `static_routes.ip_prefixes` property

Type: `["list", "string"]`. Computed.

List of route prefixes that have common next hop and attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--virtual_network--reference--group-001.md#canonical-1203112303110222-2133131133220301-0313031022113033-0021122131133012-0022332222203001-3200320000113102-3132203122220232-2201200023310120): complete subsection reference.

<a id="canonical-1320030033203322-0122333002112002-3211111203121001-1113020333313023-3010312021102220-3230103111103002-2120221002032302-3323133300013332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `static_routes.default_gateway` properties

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-2112023310021033-2012123022032122-0030112331222000-3101313021113310-0300312330333113-3110103330132201-2330310010223213-3013312303120122)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-2130101021323011-3002330133321312-3110103001302103-2133320321012310-3010320330020032-1032031113023113-3330002130320310-0200303030321113)
- [static_routes](data-sources--virtual_network--reference--group-001.md#canonical-3021012132213221-1330322221222333-3321310120010031-1221230231213120-1100203011130122-2330113131123330-2012322323011111-1211200322300231)
- static_routes.default_gateway

<a id="canonical-2001213212032023-3100321213310001-3110013320013122-0133001132301112-2023113020122023-2222321101123033-0212012203122111-2313031112331310"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

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

<a id="canonical-1203112303110222-2133131133220301-0313031022113033-0021122131133012-0022332222203001-3200320000113102-3132203122220232-2201200023310120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `static_routes.node_interface` properties

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-2112023310021033-2012123022032122-0030112331222000-3101313021113310-0300312330333113-3110103330132201-2330310010223213-3013312303120122)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-2130101021323011-3002330133321312-3110103001302103-2133320321012310-3010320330020032-1032031113023113-3330002130320310-0200303030321113)
- [static_routes](data-sources--virtual_network--reference--group-001.md#canonical-3021012132213221-1330322221222333-3321310120010031-1221230231213120-1100203011130122-2330113131123330-2012322323011111-1211200322300231)
- static_routes.node_interface

<a id="canonical-0111311023021103-1003121021212321-3102102112301222-2210031333300312-3310201123231113-1032011313031002-1110221200123110-0201322201121210"></a>

Type: `"single"`. Computed.

On multinode site, this type holds the information about per node interfaces.

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

<a id="canonical-3202203223022122-3121212320203220-1032120330313020-3121130322312331-2130003332032131-0011121022300013-3322322132303020-0330322002131222"></a>

### Direct properties for `static_routes.node_interface`

- [list](data-sources--virtual_network--reference--group-001.md#canonical-1211120123103031-0301310100203102-2332232103303120-0130100220101311-0010030033211213-2011102230130022-1313232101210022-2321222333123022): complete subsection reference.

<a id="canonical-1211120123103031-0301310100203102-2332232103303120-0130100220101311-0010030033211213-2011102230130022-1313232101210022-2321222333123022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `static_routes.node_interface.list` properties

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-2112023310021033-2012123022032122-0030112331222000-3101313021113310-0300312330333113-3110103330132201-2330310010223213-3013312303120122)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-2130101021323011-3002330133321312-3110103001302103-2133320321012310-3010320330020032-1032031113023113-3330002130320310-0200303030321113)
- [static_routes](data-sources--virtual_network--reference--group-001.md#canonical-3021012132213221-1330322221222333-3321310120010031-1221230231213120-1100203011130122-2330113131123330-2012322323011111-1211200322300231)
- [static_routes.node_interface](data-sources--virtual_network--reference--group-001.md#canonical-1203112303110222-2133131133220301-0313031022113033-0021122131133012-0022332222203001-3200320000113102-3132203122220232-2201200023310120)
- static_routes.node_interface.list

<a id="canonical-2303331213303132-1002002112002331-2032012232031022-3133003202213202-3232032131011112-2121102112012231-3220233020022032-1110101130033333"></a>

Type: `"list"`. Computed.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-2212022223221332-0300221321003212-2102333121212003-3303333300120312-3203233210322330-2211123301100311-3332003221210230-1303233220313013"></a>

### Direct properties for `static_routes.node_interface.list`

- [interface](data-sources--virtual_network--reference--group-001.md#canonical-1311023100332110-1211220033202220-2233113202232000-0133100132001112-1000130220220020-2020030320012232-3003210300233010-2032301112021011): complete subsection reference.

<a id="canonical-2223121332232330-2131212321200102-3333210030211201-1113220321321121-2231012203102303-2203122203200333-1332320202032202-3021210103100123"></a>

<a id="canonical-1002033200023130-3210323111320310-2102001200030211-0113333213212313-2301012120033223-0320232000222300-3113101012031110-1223133332130112"></a>

#### `static_routes.node_interface.list.node` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1311023100332110-1211220033202220-2233113202232000-0133100132001112-1000130220220020-2020030320012232-3003210300233010-2032301112021011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `static_routes.node_interface.list.interface` properties

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-2112023310021033-2012123022032122-0030112331222000-3101313021113310-0300312330333113-3110103330132201-2330310010223213-3013312303120122)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-2130101021323011-3002330133321312-3110103001302103-2133320321012310-3010320330020032-1032031113023113-3330002130320310-0200303030321113)
- [static_routes](data-sources--virtual_network--reference--group-001.md#canonical-3021012132213221-1330322221222333-3321310120010031-1221230231213120-1100203011130122-2330113131123330-2012322323011111-1211200322300231)
- [static_routes.node_interface](data-sources--virtual_network--reference--group-001.md#canonical-1203112303110222-2133131133220301-0313031022113033-0021122131133012-0022332222203001-3200320000113102-3132203122220232-2201200023310120)
- [static_routes.node_interface.list](data-sources--virtual_network--reference--group-001.md#canonical-1211120123103031-0301310100203102-2332232103303120-0130100220101311-0010030033211213-2011102230130022-1313232101210022-2321222333123022)
- static_routes.node_interface.list.interface

<a id="canonical-3030002031301133-3100233323220200-0313311130011112-2132201230201310-2312102320031203-1313003013332002-2101002312201322-0220321310013333"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1120223010032020-0031203011000211-3032203313231001-3213120222003122-0333311222022211-0011013131303102-3301000113033213-3322013202330323"></a>

### Direct properties for `static_routes.node_interface.list.interface`

<a id="canonical-0133301320232012-2102203122231123-2111302213022211-3330102231300130-1110121023313120-0020312120213301-2122302113213332-0310200303001122"></a>

#### `static_routes.node_interface.list.interface.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2002023201021113-1211112130012103-2320223022320003-1021301000012101-2110212012100133-3213231103202133-2133022002001111-1013301310302021"></a>

<a id="canonical-1222211101320120-0222210302213100-3111233102020033-0323231122100103-0331302020023133-0300203122303010-1223030012330332-0110003100310201"></a>

#### `static_routes.node_interface.list.interface.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0031010111121231-0333131000230331-0123021000213033-0100223032200203-2122010203201322-1213001200100312-1111013231222033-3032030231000233"></a>

<a id="canonical-1011121321122102-0122001123320302-0203033232132002-2101333023100123-0221012033302000-0032220232110013-0032300121130110-3210001323222222"></a>

#### `static_routes.node_interface.list.interface.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2232021011112332-2030102330321102-0030320033211002-3110102203121102-3320110222312232-0033212330121133-0100212310033200-2301121231100113"></a>

<a id="canonical-2023101010322123-0020011133203233-3301230132311113-3213313102310323-3210323002000300-0313332103321330-1120301110011212-3021223031322131"></a>

#### `static_routes.node_interface.list.interface.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2120032131133212-0212221323113102-0300332002003120-2102120123002212-3010232200003001-0112122003213121-3100012022300001-3331021021310133"></a>

<a id="canonical-1102213212023203-2331330032322301-2032002021131311-3031300113323321-1311011320203210-2030131110223201-1312303200002310-2311012222221212"></a>

#### `static_routes.node_interface.list.interface.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
