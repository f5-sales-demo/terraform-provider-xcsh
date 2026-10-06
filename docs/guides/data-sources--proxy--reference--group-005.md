---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-1101012223103130-2201131330221131-1331022102330232-3330000112313011-2223113210330001-3012032001012200-0303100022020121-1102012131121322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_local_inside_network` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- site_local_inside_network

<a id="canonical-3011333013231301-3311000112331023-1233112001130020-2301120100033013-0202330003110123-3112302112010301-0101132030321312-1221012000013202"></a>

Type: `["object", {}]`. Computed.

\[OneOf: site\_local\_inside\_network, site\_local\_network\] Enable this option

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

- [site_local_inside_network](data-sources--proxy--reference--group-005.md#canonical-3011333013231301-3311000112331023-1233112001130020-2301120100033013-0202330003110123-3112302112010301-0101132030321312-1221012000013202)
- [site_local_network](data-sources--proxy--reference--group-005.md#canonical-2120312123230111-3221222202321003-3323023300110113-2301113023020122-0303322012003223-0223312123130333-1302023103013221-2212023311032121)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313130232332212-2111301322113102-2223321320232221-1220210033111003-1130311211020203-2230312001121313-0110231202002231-1133102331220011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_local_network` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- site_local_network

<a id="canonical-2120312123230111-3221222202321003-3323023300110113-2301113023020122-0303322012003223-0223312123130333-1302023103013221-2212023311032121"></a>

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

<a id="canonical-1012201322333030-2201110212121013-2012102101130230-0200303100213333-0202312022013300-2020101021021321-1320201313130100-3300013202221330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_virtual_sites` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- site_virtual_sites

<a id="canonical-0112331330121211-0322102123023021-1032313031312220-3300101133311202-1313122111333003-0221311103130312-2122111322222100-2203321301210323"></a>

Type: `"single"`. Computed.

This defines a way to advertise a VIP on specific sites.

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

<a id="canonical-0201200111033103-3033011111023222-0202010000030210-0301230201101120-3120303133021023-0320211103001031-3112022021100030-0231233032123001"></a>

### Direct properties for `site_virtual_sites`

- [advertise_where](data-sources--proxy--reference--group-005.md#canonical-1101311121310312-0200013102222131-0303021112013023-0200203203003231-1221230311001231-0230200233121202-1303302123221200-0111202023011311): complete subsection reference.

<a id="canonical-1101311121310312-0200013102222131-0303021112013023-0200203203003231-1221230311001231-0230200233121202-1303302123221200-0111202023011311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_virtual_sites.advertise_where` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [site_virtual_sites](data-sources--proxy--reference--group-005.md#canonical-1012201322333030-2201110212121013-2012102101130230-0200303100213333-0202312022013300-2020101021021321-1320201313130100-3300013202221330)
- site_virtual_sites.advertise_where

<a id="canonical-1200222300023123-0302000223123311-2123221321233322-1010333233212212-2321202302300322-2012321212102331-3303103110300100-3303320120320112"></a>

Type: `"list"`. Computed.

Where should this load balancer be available.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1130022310333201-0122312310013312-2130101010101322-2103332300131033-2010233330211200-0321101033300132-0011223112300313-3332131303022222"></a>

### Direct properties for `site_virtual_sites.advertise_where`

<a id="canonical-1003300213011310-1131002231013312-2001022031013312-0121210110123100-1210120023012300-2113300211020220-1300003101302022-0013221322100331"></a>

#### `site_virtual_sites.advertise_where.port` property

Type: `"number"`. Computed.

Exclusive with \[use\_default\_port\] TCP port to Listen.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
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

- [site](data-sources--proxy--reference--group-005.md#canonical-1333202122120023-2021303133022133-2310020033131020-2232302312131321-3111212111303323-0022222112300022-1210100123323001-0202030210022123): complete subsection reference.

- [use_default_port](data-sources--proxy--reference--group-005.md#canonical-3323010231122112-0120132012222113-3332210013112122-3301101113111032-2002323210013011-1321302031300111-3102001300123202-2111110012033312): complete subsection reference.

- [virtual_site](data-sources--proxy--reference--group-005.md#canonical-0011000332332130-3130012231211212-3202112322222113-3111230132210223-2100313233112001-3112212110232222-2031021001200031-3310310312213001): complete subsection reference.

<a id="canonical-1333202122120023-2021303133022133-2310020033131020-2232302312131321-3111212111303323-0022222112300022-1210100123323001-0202030210022123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_virtual_sites.advertise_where.site` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [site_virtual_sites](data-sources--proxy--reference--group-005.md#canonical-1012201322333030-2201110212121013-2012102101130230-0200303100213333-0202312022013300-2020101021021321-1320201313130100-3300013202221330)
- [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-005.md#canonical-1101311121310312-0200013102222131-0303021112013023-0200203203003231-1221230311001231-0230200233121202-1303302123221200-0111202023011311)
- site_virtual_sites.advertise_where.site

<a id="canonical-2012232123210032-2130001211033133-0303112231020220-1031303102323210-1122110133122322-2220130102120022-2012222330012033-2011300322132030"></a>

Type: `"single"`. Computed.

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

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

<a id="canonical-1211220230211131-1111312101120113-2201013131232310-3030310032303311-3231001010102033-3333013131201330-1112320133103112-0132302023231020"></a>

### Direct properties for `site_virtual_sites.advertise_where.site`

<a id="canonical-1302131332333120-1102321233132020-2302312223133223-0230201112113333-2023022101000223-3120201021121003-2120112020133113-3032022223022133"></a>

#### `site_virtual_sites.advertise_where.site.ip` property

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1122332323303310-3312233030010333-2221100001312001-0222200032012321-1313232000200032-1200333331021300-1002133123002101-0021322201133333"></a>

<a id="canonical-2033100311010111-1212321100201223-2100021022310313-0223031202310100-3033212003321300-0203221332332020-1122303213221221-3030133333002223"></a>

#### `site_virtual_sites.advertise_where.site.network` property

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](data-sources--proxy--reference--group-005.md#canonical-0333023233320103-0333102001103011-0330232011323331-0300003103313121-1020313321130030-0120110010300213-3001310001222022-1022001022013210): complete subsection reference.

<a id="canonical-0333023233320103-0333102001103011-0330232011323331-0300003103313121-1020313321130030-0120110010300213-3001310001222022-1022001022013210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_virtual_sites.advertise_where.site.site` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [site_virtual_sites](data-sources--proxy--reference--group-005.md#canonical-1012201322333030-2201110212121013-2012102101130230-0200303100213333-0202312022013300-2020101021021321-1320201313130100-3300013202221330)
- [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-005.md#canonical-1101311121310312-0200013102222131-0303021112013023-0200203203003231-1221230311001231-0230200233121202-1303302123221200-0111202023011311)
- [site_virtual_sites.advertise_where.site](data-sources--proxy--reference--group-005.md#canonical-1333202122120023-2021303133022133-2310020033131020-2232302312131321-3111212111303323-0022222112300022-1210100123323001-0202030210022123)
- site_virtual_sites.advertise_where.site.site

<a id="canonical-2012113200013210-2311222310130020-1222230031123000-0012231310010203-1112223133301330-1330022320000303-1131213031302031-1300112130230223"></a>

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

<a id="canonical-0311202313102302-3011312100213132-2231103201013010-0213101123322020-1202000010033212-1311231301212221-2213112302212020-3220202312333120"></a>

### Direct properties for `site_virtual_sites.advertise_where.site.site`

<a id="canonical-0012223120102133-3213112123330320-2213332311232003-2331113031020023-3200310320320223-1113221300330310-3221103311103200-2333121012203323"></a>

#### `site_virtual_sites.advertise_where.site.site.name` property

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

<a id="canonical-2033122221300231-0112332232003212-2221120000202000-0021303332001113-2101122323322311-2213310323313321-1023312001221300-2133320313010302"></a>

<a id="canonical-0322112131000313-3222023133230013-0332001121111310-3131331222320313-0333230300122232-2132230013010011-2221231313320121-2302221232103022"></a>

#### `site_virtual_sites.advertise_where.site.site.namespace` property

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

<a id="canonical-1331132003120312-2310221101112021-3100333113333103-3313200330003332-2133322123333000-1312312121331031-2131222322020212-1001303131321030"></a>

<a id="canonical-1002313210012122-0022130103311333-1023011013023001-0112230231123333-0102300202201203-0032131031003211-0010033303311220-2000320002021120"></a>

#### `site_virtual_sites.advertise_where.site.site.tenant` property

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

<a id="canonical-3323010231122112-0120132012222113-3332210013112122-3301101113111032-2002323210013011-1321302031300111-3102001300123202-2111110012033312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_virtual_sites.advertise_where.use_default_port` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [site_virtual_sites](data-sources--proxy--reference--group-005.md#canonical-1012201322333030-2201110212121013-2012102101130230-0200303100213333-0202312022013300-2020101021021321-1320201313130100-3300013202221330)
- [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-005.md#canonical-1101311121310312-0200013102222131-0303021112013023-0200203203003231-1221230311001231-0230200233121202-1303302123221200-0111202023011311)
- site_virtual_sites.advertise_where.use_default_port

<a id="canonical-3113233331220322-3002231203202102-3112012102020323-3002010222031113-1020101312111023-3020322300132123-2232233003113231-1133211003031123"></a>

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

<a id="canonical-0011000332332130-3130012231211212-3202112322222113-3111230132210223-2100313233112001-3112212110232222-2031021001200031-3310310312213001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_virtual_sites.advertise_where.virtual_site` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [site_virtual_sites](data-sources--proxy--reference--group-005.md#canonical-1012201322333030-2201110212121013-2012102101130230-0200303100213333-0202312022013300-2020101021021321-1320201313130100-3300013202221330)
- [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-005.md#canonical-1101311121310312-0200013102222131-0303021112013023-0200203203003231-1221230311001231-0230200233121202-1303302123221200-0111202023011311)
- site_virtual_sites.advertise_where.virtual_site

<a id="canonical-0122302112031312-2233031012313202-2222333322110221-1032100112002110-1223131202102313-2003322332112302-2011203233121300-0131010111020301"></a>

Type: `"single"`. Computed.

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

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

<a id="canonical-1231230311003321-3101223212231101-3022300303300321-0012001100030203-2003231312023300-0011213131230213-3303323301122300-1230233003232110"></a>

### Direct properties for `site_virtual_sites.advertise_where.virtual_site`

<a id="canonical-0222120033110300-0222120010010000-3112100331310022-0123123022321231-1022120003330213-2001010322213310-2212220103223230-3302330231222233"></a>

#### `site_virtual_sites.advertise_where.virtual_site.network` property

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--proxy--reference--group-005.md#canonical-1302033331002032-1233232303323222-2033003231120320-2232201001323013-0321020302101332-0221333200123101-1123201021133203-1022302201220101): complete subsection reference.

<a id="canonical-1302033331002032-1233232303323222-2033003231120320-2232201001323013-0321020302101332-0221333200123101-1123201021133203-1022302201220101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_virtual_sites.advertise_where.virtual_site.virtual_site` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [site_virtual_sites](data-sources--proxy--reference--group-005.md#canonical-1012201322333030-2201110212121013-2012102101130230-0200303100213333-0202312022013300-2020101021021321-1320201313130100-3300013202221330)
- [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-005.md#canonical-1101311121310312-0200013102222131-0303021112013023-0200203203003231-1221230311001231-0230200233121202-1303302123221200-0111202023011311)
- [site_virtual_sites.advertise_where.virtual_site](data-sources--proxy--reference--group-005.md#canonical-0011000332332130-3130012231211212-3202112322222113-3111230132210223-2100313233112001-3112212110232222-2031021001200031-3310310312213001)
- site_virtual_sites.advertise_where.virtual_site.virtual_site

<a id="canonical-3201133000030101-2111222220311312-1223130313233220-0003221131111102-3001333133233332-1111313322322230-2032222020330202-3202301320120331"></a>

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

<a id="canonical-2320120130322313-3313020223222002-3003132330011111-3210320222220230-3102223300101002-3100200301321123-2100112033310323-1111023233313020"></a>

### Direct properties for `site_virtual_sites.advertise_where.virtual_site.virtual_site`

<a id="canonical-1201303031032001-1001122221023331-3320003123203113-1132120311222212-0322333312112103-0100123030300021-1300222133300112-3030230332331030"></a>

#### `site_virtual_sites.advertise_where.virtual_site.virtual_site.name` property

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

<a id="canonical-2212033210031212-3310110313112022-3023000013120231-2133003010321322-0311122100133321-3321322031031220-2011330021220100-0011310223203122"></a>

<a id="canonical-2313330221233120-3003310302211031-3213103300112033-0313301120201031-3230031212002111-3033021322122122-3111213013101110-2322013332021323"></a>

#### `site_virtual_sites.advertise_where.virtual_site.virtual_site.namespace` property

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

<a id="canonical-0332021221300200-0222232011123103-0233323231032102-1212312113123021-1233032122322330-0332003210200322-1333012322002222-3032021030131001"></a>

<a id="canonical-0013212103331330-2000210222102223-0231022110323300-2213323312100102-2212033313012311-2103303010122302-0213301312313002-0211320313320203"></a>

#### `site_virtual_sites.advertise_where.virtual_site.virtual_site.tenant` property

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

<a id="canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- tls_intercept

<a id="canonical-2121232312132030-1021002130200310-0121113113303233-3010003311301312-3031223213223011-2223320220020023-0033302022021110-1101201020202012"></a>

Type: `"single"`. Computed.

Configuration to enable TLS interception.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-interception_policy_choice": "[\"enable_for_all_domains\",\"policy\"]",
  "x-ves-oneof-field-signing_cert_choice": "[\"custom_certificate\",\"volterra_certificate\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca_url\",\"volterra_trusted_ca\"]"
}
```

<a id="canonical-3031321233011130-1322000233203311-1102003212023022-3101112132211020-3222222311032110-3200003332220300-2120223320032310-0331222311332311"></a>

### Direct properties for `tls_intercept`

- [custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220): complete subsection reference.

- [enable_for_all_domains](data-sources--proxy--reference--group-005.md#canonical-3230230222011313-3333232211102332-1213233310022012-2200313222323113-0332021201200103-0202133211030313-3132201012212010-1321323130131223): complete subsection reference.

- [policy](data-sources--proxy--reference--group-005.md#canonical-3103321231313310-3112212012231231-2021103203231010-3203131230032202-2311212331112113-2203200112111100-2213330010011221-1023110203010100): complete subsection reference.

<a id="canonical-0102133020010102-2213012022123223-3310030013031013-0203321130301312-2202121010022122-3010200322231312-1221222212020132-2033120022201111"></a>

<a id="canonical-3312032223302110-0030323133121132-2010302101022133-0121231220321233-2133132203203212-0212303130000210-2123300211000101-2113303013311030"></a>

#### `tls_intercept.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [volterra_certificate](data-sources--proxy--reference--group-005.md#canonical-3020010120123000-0301123030210101-0321130032221323-2123030030303222-2003011230132203-2230223213222003-3100213222023330-0120202213313121): complete subsection reference.

- [volterra_trusted_ca](data-sources--proxy--reference--group-005.md#canonical-0200220011202212-2321233023123120-1202023310102102-0222131001130112-0320121130100111-2103030131332202-1120332223231302-0301113013200231): complete subsection reference.

<a id="canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.custom_certificate` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- tls_intercept.custom_certificate

<a id="canonical-3232201211233031-3302012031311123-3001022011213013-3030010132310122-1012011300203110-2120201011302001-2312333002221232-2013321321332302"></a>

Type: `"single"`. Computed.

Configuration parameter for custom certificate.

Additional upstream details:

Handle to fetch certificate and key.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ocsp_stapling_choice": "[\"custom_hash_algorithms\",\"disable_ocsp_stapling\",\"use_system_defaults\"]"
}
```

<a id="canonical-2311322213113333-1000213013332332-3312010221123112-0031311330022123-1223021103213231-1013132232003220-0301123202001221-3220233011113111"></a>

### Direct properties for `tls_intercept.custom_certificate`

<a id="canonical-3032113001001113-0020202300333322-2021210112230303-0303003102320003-0302033300021212-0320103321201300-0232333012323230-3310330020032210"></a>

#### `tls_intercept.custom_certificate.certificate_url` property

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](data-sources--proxy--reference--group-005.md#canonical-1023213030332201-2312100033010103-1010203103321123-0111001103030100-0132333110030230-3313200112220113-1202302003003233-1021310321232310): complete subsection reference.

<a id="canonical-2020220001230310-1310333302301032-0112320213131013-2131311131001203-1233030201122032-0221031232311101-2331332332010021-2022033103331101"></a>

<a id="canonical-1331202331313130-3000230123312102-0211323211033223-3031010110032211-3223232220311123-3003121010102131-0022221112333333-0123320130133033"></a>

#### `tls_intercept.custom_certificate.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--proxy--reference--group-005.md#canonical-1030233231323122-3203323120201203-2300102221222310-3100102220230001-1033122103110300-0110200213132320-3133201033332133-1023222210303130): complete subsection reference.

- [private_key](data-sources--proxy--reference--group-005.md#canonical-3110120000322320-0101001302313122-1331202112230322-2202301331030002-0023312212222101-3110002202232021-3023303132133330-0232333301023013): complete subsection reference.

- [use_system_defaults](data-sources--proxy--reference--group-005.md#canonical-1302232200212121-3021213100303023-2102012001133212-3021120313030310-3122221011012111-2302202111200013-3232133330311313-3213203112122232): complete subsection reference.

<a id="canonical-1023213030332201-2312100033010103-1010203103321123-0111001103030100-0132333110030230-3313200112220113-1202302003003233-1021310321232310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.custom_certificate.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- tls_intercept.custom_certificate.custom_hash_algorithms

<a id="canonical-0130300112022033-0030312103101302-2022332303012121-2223121110330330-1012220323322231-3210232222201113-0223302013033130-0122303212333021"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

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

<a id="canonical-1302030002110213-3303122330002000-1221321232230312-3111232232311313-3100322132310000-2333131320311213-3003000122022101-3223030331211111"></a>

### Direct properties for `tls_intercept.custom_certificate.custom_hash_algorithms`

<a id="canonical-3200323301112120-3101303120100111-2021213132221023-3212121111103233-1022010200113323-2130201330103311-2202010233201012-2033110311302111"></a>

#### `tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1030233231323122-3203323120201203-2300102221222310-3100102220230001-1033122103110300-0110200213132320-3133201033332133-1023222210303130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.custom_certificate.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- tls_intercept.custom_certificate.disable_ocsp_stapling

<a id="canonical-3201322323312302-3101312131312220-2022201233311300-0021122010211331-3201230333320103-1221301010202003-1100011032121030-3221131332121201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

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

<a id="canonical-3110120000322320-0101001302313122-1331202112230322-2202301331030002-0023312212222101-3110002202232021-3023303132133330-0232333301023013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.custom_certificate.private_key` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- tls_intercept.custom_certificate.private_key

<a id="canonical-1321110132331221-3323111110320330-3201123022231032-1020320300200320-0200110113213103-1200221012112103-0011002102322123-2311311312011132"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-2223223323231300-1022213023200011-3311013113301311-3200111312102133-2331111211320231-1002313201032302-0302001230211112-3222232322022100"></a>

### Direct properties for `tls_intercept.custom_certificate.private_key`

- [blindfold_secret_info](data-sources--proxy--reference--group-005.md#canonical-0322103332321130-1213301201221100-2100011323312323-0130121211021000-0130121132113332-2011103223230322-0000212002313013-3212022300212232): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-005.md#canonical-1322212321230133-2133303132322001-3330322212211033-1232322231000102-3231132220100133-0132221001212132-3321011302213132-2200131333232202): complete subsection reference.

<a id="canonical-0322103332321130-1213301201221100-2100011323312323-0130121211021000-0130121132113332-2011103223230322-0000212002313013-3212022300212232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.custom_certificate.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- [tls_intercept.custom_certificate.private_key](data-sources--proxy--reference--group-005.md#canonical-3110120000322320-0101001302313122-1331202112230322-2202301331030002-0023312212222101-3110002202232021-3023303132133330-0232333301023013)
- tls_intercept.custom_certificate.private_key.blindfold_secret_info

<a id="canonical-2130120001211200-2023110233223223-0222210032032110-1233200000232212-2311301002222312-0120313103023031-1000020000022222-1123030132123131"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-3112032113003021-0121213211333022-2331213322200003-2023100130310331-1220300022201101-0002120032031311-2013223023102211-1100133313302320"></a>

### Direct properties for `tls_intercept.custom_certificate.private_key.blindfold_secret_info`

<a id="canonical-3133020322322213-2333332200311132-0130310312200321-2122211023321002-1202113031210112-0203031120110211-3301223211131021-1310203132122120"></a>

#### `tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-3300030203302003-3021022002320200-2212330120322031-3023010201301310-2123012002331131-2133310031211122-2231200320122121-1002313301032322"></a>

<a id="canonical-0000012120210132-0300323012032101-2130020301020223-0323311123323203-3130120233231300-3103100323120220-0131000331231020-1320031022230022"></a>

#### `tls_intercept.custom_certificate.private_key.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3323231132013021-1201331320010211-0210103100312033-3011001303330321-0132133201011233-1303003000210203-2301001312231223-2130133020133313"></a>

<a id="canonical-3132100321211030-2320312211102212-2222300230201302-1033211031202023-2002123300003221-0213233000010033-2101303303101113-3213132210320113"></a>

#### `tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-1322212321230133-2133303132322001-3330322212211033-1232322231000102-3231132220100133-0132221001212132-3321011302213132-2200131333232202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.custom_certificate.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- [tls_intercept.custom_certificate.private_key](data-sources--proxy--reference--group-005.md#canonical-3110120000322320-0101001302313122-1331202112230322-2202301331030002-0023312212222101-3110002202232021-3023303132133330-0232333301023013)
- tls_intercept.custom_certificate.private_key.clear_secret_info

<a id="canonical-1220233030332323-0010003022111130-1221031230130102-3021122123302321-2221222223133323-0133001022232301-2021221101210130-2333211301320002"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-1002011220031132-1330133330213310-3003310230112031-0020230331031101-2232231112031223-2223311111031231-3012113111330033-3012223030332131"></a>

### Direct properties for `tls_intercept.custom_certificate.private_key.clear_secret_info`

<a id="canonical-2023012032300020-3130130122203230-0200111221020000-3021113231330001-2233110010301332-0203331020321031-3300300113232133-1203013223223222"></a>

#### `tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3003332223312231-3301323201013300-1201200303001212-0030101103331331-1302011320113303-1203213132132320-2223023212111003-1321313220320020"></a>

<a id="canonical-0332310322200131-3211033320133233-0321232010203111-2332123132000123-2111201111210213-1100330332021110-0312022133121303-2131033200013200"></a>

#### `tls_intercept.custom_certificate.private_key.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1302232200212121-3021213100303023-2102012001133212-3021120313030310-3122221011012111-2302202111200013-3232133330311313-3213203112122232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.custom_certificate.use_system_defaults` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- tls_intercept.custom_certificate.use_system_defaults

<a id="canonical-3210101010210211-0003110102111231-0313232033230131-1320322121023121-0211233123211003-0103112203310321-3320111020221230-0113222113120201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

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

<a id="canonical-3230230222011313-3333232211102332-1213233310022012-2200313222323113-0332021201200103-0202133211030313-3132201012212010-1321323130131223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.enable_for_all_domains` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- tls_intercept.enable_for_all_domains

<a id="canonical-1330022123231330-2322122032123301-0312121302133213-0322313230020103-0033201221231313-1133301231300323-0012010313013202-1302211022331311"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable for all domains.

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

<a id="canonical-3103321231313310-3112212012231231-2021103203231010-3203131230032202-2311212331112113-2203200112111100-2213330010011221-1023110203010100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.policy` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- tls_intercept.policy

<a id="canonical-2101031213001101-1213003123113223-1320233222033313-3033132022011020-1230122203212100-3230122133133232-3210213110311030-2201310011021210"></a>

Type: `"single"`. Computed.

Policy to enable or disable TLS interception.

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

<a id="canonical-2001313200102311-0021032132311103-0302213323111211-2302001011122300-1232000303323223-1003200301331233-3202200313232200-1003211201300222"></a>

### Direct properties for `tls_intercept.policy`

- [interception_rules](data-sources--proxy--reference--group-005.md#canonical-1023013012011313-3030230203032111-3020023301210312-3201110221202022-1210031331033210-2233033030120113-0233111003200102-0320103010122112): complete subsection reference.

<a id="canonical-1023013012011313-3030230203032111-3020023301210312-3201110221202022-1210031331033210-2233033030120113-0233111003200102-0320103010122112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.policy.interception_rules` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-3103321231313310-3112212012231231-2021103203231010-3203131230032202-2311212331112113-2203200112111100-2213330010011221-1023110203010100)
- tls_intercept.policy.interception_rules

<a id="canonical-1113032002213112-3322213023132212-3103021213030223-0312300312212233-0303223123321130-3231332133031300-3212332320202121-1301323320211021"></a>

Type: `"list"`. Computed.

List of ordered rules to enable or disable for TLS interception.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0220113111232110-3222131202023203-1010020313221212-2030230101113133-1300032233210332-3010201013230232-0202302211221110-1330213332110102"></a>

### Direct properties for `tls_intercept.policy.interception_rules`

- [disable_interception](data-sources--proxy--reference--group-005.md#canonical-2120002331023310-2300002310011311-3023231120133323-3123331131120200-1110110202022323-3020312210233102-1302301033100013-2113011320103113): complete subsection reference.

- [domain_match](data-sources--proxy--reference--group-005.md#canonical-0301211320300332-0131231113212310-0300320322212133-2233333301210212-0130231321210120-1221333032210230-0012022300222310-0303111321201312): complete subsection reference.

- [enable_interception](data-sources--proxy--reference--group-005.md#canonical-1023223231130220-3312032212030130-3011110301100311-0220302013032313-0001011202133010-2031233331000130-3221121321223012-3030210003033123): complete subsection reference.

<a id="canonical-2120002331023310-2300002310011311-3023231120133323-3123331131120200-1110110202022323-3020312210233102-1302301033100013-2113011320103113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.policy.interception_rules.disable_interception` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-3103321231313310-3112212012231231-2021103203231010-3203131230032202-2311212331112113-2203200112111100-2213330010011221-1023110203010100)
- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-1023013012011313-3030230203032111-3020023301210312-3201110221202022-1210031331033210-2233033030120113-0233111003200102-0320103010122112)
- tls_intercept.policy.interception_rules.disable_interception

<a id="canonical-3230310311131200-2202011103231103-1133330322213212-3011112320010120-2313232303213000-1102322223010202-1222012233231310-0213032033000332"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable interception.

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

<a id="canonical-0301211320300332-0131231113212310-0300320322212133-2233333301210212-0130231321210120-1221333032210230-0012022300222310-0303111321201312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.policy.interception_rules.domain_match` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-3103321231313310-3112212012231231-2021103203231010-3203131230032202-2311212331112113-2203200112111100-2213330010011221-1023110203010100)
- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-1023013012011313-3030230203032111-3020023301210312-3201110221202022-1210031331033210-2233033030120113-0233111003200102-0320103010122112)
- tls_intercept.policy.interception_rules.domain_match

<a id="canonical-2000211212020220-0323011121211120-3022030001001331-2333103121200000-1100332222323320-2303213003233321-3100201022110311-0132232200311121"></a>

Type: `"single"`. Computed.

Configuration parameter for domain match.

Additional upstream details:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-3101111313030222-2030202103111210-2032223333000110-1221213112113303-1013023200100132-0221033331021003-2121213200301212-1221320203001011"></a>

### Direct properties for `tls_intercept.policy.interception_rules.domain_match`

<a id="canonical-3121132113013333-2011011013232313-3020113110312310-0012302303230203-0302013022233011-3210033020331303-2102200101103230-1203301202033122"></a>

#### `tls_intercept.policy.interception_rules.domain_match.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1210110210120023-1103002303201111-1103113101110233-3132303330011010-0203300113020133-3222233222001221-2230310212311113-2100312032012320"></a>

<a id="canonical-2232011233033313-3123020323021313-3210122201033330-0210011012120332-2111331002100221-3030330110032113-0220113103020100-2131001003303322"></a>

#### `tls_intercept.policy.interception_rules.domain_match.regex_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1010202121322323-3021203112223202-2113023231112103-2020231202201303-2102203201030300-1013032132130231-1300331012111300-2303021121300211"></a>

<a id="canonical-2332002113133103-2102133221101123-3300232201322332-2331201311330100-3302221003300121-2020102322230120-1110330321013133-0130011200100012"></a>

#### `tls_intercept.policy.interception_rules.domain_match.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1023223231130220-3312032212030130-3011110301100311-0220302013032313-0001011202133010-2031233331000130-3221121321223012-3030210003033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.policy.interception_rules.enable_interception` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-3103321231313310-3112212012231231-2021103203231010-3203131230032202-2311212331112113-2203200112111100-2213330010011221-1023110203010100)
- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-1023013012011313-3030230203032111-3020023301210312-3201110221202022-1210031331033210-2233033030120113-0233111003200102-0320103010122112)
- tls_intercept.policy.interception_rules.enable_interception

<a id="canonical-1100020223102300-2322222300101130-2112310310233020-1220231110200301-3020010330133230-1301022323201233-1320122113213301-2310222122330113"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable interception.

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

<a id="canonical-3020010120123000-0301123030210101-0321130032221323-2123030030303222-2003011230132203-2230223213222003-3100213222023330-0120202213313121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.volterra_certificate` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- tls_intercept.volterra_certificate

<a id="canonical-3232030101211322-0013110331131102-1232303132112113-0233211103033200-1033002221322313-2220003312130223-3301221323232332-3331032222313221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra certificate.

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

<a id="canonical-0200220011202212-2321233023123120-1202023310102102-0222131001130112-0320121130100111-2103030131332202-1120332223231302-0301113013200231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.volterra_trusted_ca` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- tls_intercept.volterra_trusted_ca

<a id="canonical-3310300332010012-0031300302033223-2003332123223232-2201023223220121-2313332013010022-0102102113300211-3023130302230021-3310013030133301"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra trusted ca.

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
