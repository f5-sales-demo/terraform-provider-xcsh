---
page_title: "xcsh_dns_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy reference."
---

# xcsh_dns_proxy reference

<a id="canonical-2211122331313122-2210032001000130-1233112330000231-3332132220302101-2130212230003003-0132203130333103-2201210033301231-3323123103323301"></a>

## Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_site`

<a id="canonical-0332203300131313-0233111331002231-3030103103333230-3301031331232012-0110031302332022-1122001011312131-0222331321203132-3213132101002323"></a>

### `proxy_advertisement.advertise_custom.advertise_where.virtual_site.network` property

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

- [virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-2012022001232103-2333011021301010-1022323013200032-2132200313111213-1210231313331331-1101121323103202-3111221320333330-1122310023121203): complete subsection reference.

<a id="canonical-2012022001232103-2333011021301010-1022323013200032-2132200313111213-1210231313331331-1101121323103202-3111221320333330-1122310023121203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` properties

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](data-sources--dns_proxy--reference--group-001.md#canonical-0032203302302013-3111220213021010-0303003102222033-2220212203222302-1212122000221101-3333223131032300-0332333012122110-1013233212322320)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-1230302220330002-1201213102013312-2333020322031000-2303231223301033-0013101310221231-1101133203012323-3230020110101223-0100203131023232"></a>

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

<a id="canonical-1232100003322213-0310121110032122-0011301321112133-2033232011203201-3300100032333223-2012310210200010-0112320332320023-3310221032301002"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site`

<a id="canonical-2032030110020013-0012111102233320-0011233121202322-2122220311313123-3130121333000231-0312013310231300-0322312213221201-3330311113212023"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0222020302101131-2211101131110232-2011221000010132-2010313321230320-3002133310331212-1231320312202221-2322200203033110-3002310210312122"></a>

<a id="canonical-0220212111111221-2232033223032310-2112112221110032-3223321030231323-0030231021222202-1203003323123303-1231332221201133-3202100332312002"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2032333332030030-3202030323230123-3321221002101010-0222003313202013-0333031223032002-3123332021313110-1032111230203200-3203303323322201"></a>

<a id="canonical-0012000131333130-3321130133202213-1233230211321031-3013320003011313-2032320320231231-2331003001312313-1233222021002121-2112301112102331"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0033131110232301-2002233322320212-3111233001121322-2023001311103132-3232133223101330-2202030133010212-2111220111013323-0233220313313133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-1332100020122100-2213112023010000-2322331010010130-1113001101313131-1132321130101011-2103010110333013-2011103123233303-1312022022102233"></a>

Type: `"single"`. Computed.

This defines a reference to a customer site virtual site along with network type and IP where a load
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

<a id="canonical-2000013213200133-2220101231332203-3032113330331211-0322123000002220-0220123123221220-0210103031102010-0032131023220110-1332300330302022"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip`

<a id="canonical-0003331001023312-1213212033132123-0232221022122222-0030231331013012-3011001033100033-0203223231300133-3000130212310332-1001002213132201"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3021000131331310-2113113231220321-2210201320223000-1232321320122223-1233200320332133-0021021221323331-1033302111310302-1003013012333332"></a>

<a id="canonical-1212320101303112-0001323333333121-2103311211201201-1232323222323021-2023110100221232-0332232030323330-2022213221222110-0323122100202213"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network` property

Type: `"string"`. Computed.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on virtual-site with specified VIP

All outside networks.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
  "enum": [
    "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-2113023311201301-3102123023111311-0101233331301221-2012003322331012-3323013312013221-1311303011302101-2132201230221112-1321203031122231): complete subsection reference.

<a id="canonical-2113023311201301-3102123023111311-0101233331301221-2012003322331012-3323013312013221-1311303011302101-2132201230221112-1321203031122231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` properties

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](data-sources--dns_proxy--reference--group-002.md#canonical-0033131110232301-2002233322320212-3111233001121322-2023001311103132-3232133223101330-2202030133010212-2111220111013323-0233220313313133)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-0321113323330103-1222322133010120-0233132221011111-1321122110102210-1112231321310303-2203102111213301-1113310000132302-0330231123301021"></a>

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

<a id="canonical-2101113311013022-0000223202320231-2222201122333013-3302132100331010-0020301111330210-3102211112113323-3120131132302112-1100331022120233"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site`

<a id="canonical-2212102111221101-3100300312221123-1100020321303032-1200231221013030-1321303222003303-2333111201022010-0333320001000010-0032101301212023"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2030310102002022-3102331323011030-2132221222201003-1131330011201313-2132011213023202-2021120300220101-1220022133311221-0101033321231013"></a>

<a id="canonical-1110212202121012-3132122312211211-2012210021303331-3332332010302300-1022033022330123-1103300120222303-2200323032302031-1322301030131220"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1101023220003033-2313312100323330-3210321223312300-2220203003302132-1322130313323033-2310101211311121-0033313022201100-3201003212012233"></a>

<a id="canonical-0032202213103032-3202213221220111-1120300202123022-0203132103003223-3000032001322010-2330320301231113-0021221223302133-1303002321332013"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2033011233111310-3303121203100330-1200131323233001-0011132112303132-2021223312002302-0232310330222110-2120123123021030-3030010023303130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` properties

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service

<a id="canonical-1033112100123332-2013010212202232-1333211020033213-0120032111021023-0233033021033101-1203002012112211-2110220212023321-1230300000122231"></a>

Type: `"single"`. Computed.

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-1223200003011000-3030323000030023-0013320021311211-3113300323220122-2033331232211213-2233031122033330-0023013023222110-0120220303213103"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.vk8s_service`

- [site](data-sources--dns_proxy--reference--group-002.md#canonical-1333333123130202-1310021311112302-3233033330332110-2002110010032233-1031113301212123-3033003233023030-3221321013303232-0000131332301200): complete subsection reference.

- [virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-3311332103031002-3020231101022322-1133023003222201-3101010220323231-1023320312310130-3003302013311030-0221110133012303-2000220023100331): complete subsection reference.

<a id="canonical-1333333123130202-1310021311112302-3233033330332110-2002110010032233-1031113301212123-3033003233023030-3221321013303232-0000131332301200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` properties

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--dns_proxy--reference--group-002.md#canonical-2033011233111310-3303121203100330-1200131323233001-0011132112303132-2021223312002302-0232310330222110-2120123123021030-3030010023303130)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-3102211002111302-2222211312331122-0031322132221331-3012231000123103-1223231203330221-1102003221333021-3333310232323222-2000302312212130"></a>

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

<a id="canonical-3210300021310221-0120212321302120-3323121101313021-0201031010323102-2210303322033211-1002133200301230-0001320003013000-0220010312121313"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site`

<a id="canonical-3200323101031203-3211102110021310-3323303132003123-3330210011000220-0033302310103312-1303323001113212-3003132232323003-0202010103203123"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3100230232202330-1213233121303210-2223212101222022-1121321000111203-3331303222203323-0213332021210010-1011010212003012-1011031203100111"></a>

<a id="canonical-2230331111123231-3000223332311230-2111230101331010-2111201233320212-1111222322132000-1113323212331023-0212200332303313-2312212131310302"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3211311003012203-3300013122331133-3101231222011223-3011321212301330-3232313200113220-2313223323230203-3133130301001221-1201320200223222"></a>

<a id="canonical-3031223210220002-0232203000332210-2310113303313022-1221303132113223-0320033100322313-1121221023220321-3120321200222100-3133101332102111"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3311332103031002-3020231101022322-1133023003222201-3101010220323231-1023320312310130-3003302013311030-0221110133012303-2000220023100331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` properties

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--dns_proxy--reference--group-002.md#canonical-2033011233111310-3303121203100330-1200131323233001-0011132112303132-2021223312002302-0232310330222110-2120123123021030-3030010023303130)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-1020330201012223-1013003233110010-2120313231101200-1033232133013211-2312320323231303-2213320102131223-2111313033313313-1310100101202113"></a>

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

<a id="canonical-0321222131201122-1330233313033330-3013333203223130-0133110013212210-1101021021310100-0320000003311332-2022203000100303-1230313111311323"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site`

<a id="canonical-1020210112011102-0032103032012201-2110310110332023-0130203302120322-2032231222110020-2231132203132311-2303301320302331-2132130011300122"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2022123211100103-2301211113323202-1321001223021030-0201333100112322-1313311001231231-2030202002000312-0213023132231232-2031212113302010"></a>

<a id="canonical-2332103012333012-2231220222200311-1323112010102302-2330222202311030-0222232011012333-3302130133231022-0333122220101200-0221301013031203"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1333222213132301-0301212022320032-3101201223022103-3122312203300303-3323130120131210-3233001122120021-3221202230311121-3222313030303110"></a>

<a id="canonical-1330230233001032-3231231122311331-0103010131111133-1310001121121112-0020203110223223-1030301313112332-2032030201132112-2001332323311102"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0201333130332013-0112020301131201-2021022313321201-3013233130311223-3022300202201021-2131301231003023-3121323331203223-2002102000332121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_dualstack_on_public` properties

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- proxy_advertisement.advertise_dualstack_on_public

<a id="canonical-0321120322221213-3230322211221020-3033020113111230-3330130321232023-3212302102201312-3202323102023302-1111220132323122-2323200203010022"></a>

Type: `"single"`. Computed.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

<a id="canonical-0002331320302033-3202303313220012-1103302131210031-2010023322330233-3032033130001120-1011030013221210-3013011333231013-2213222002203303"></a>

### Direct properties for `proxy_advertisement.advertise_dualstack_on_public`

- [public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-1010032223032131-2222112020111232-2030210133123303-2023212111030102-3303322102211213-0212330211213013-2101323101230100-1013202012100331): complete subsection reference.

<a id="canonical-1010032223032131-2222112020111232-2030210133123303-2023212111030102-3303322102211213-0212330211213013-2101323101230100-1013202012100331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_dualstack_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-0201333130332013-0112020301131201-2021022313321201-3013233130311223-3022300202201021-2131301231003023-3121323331203223-2002102000332121)
- proxy_advertisement.advertise_dualstack_on_public.public_ip

<a id="canonical-3132031011321222-1222130130003221-1011323212132333-1200311333232201-3033211212101232-2321002132302213-3200121312022003-0132113311131320"></a>

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

<a id="canonical-0130321010031311-3310202021102202-1322222123322330-1331123301032110-0122212130221011-1031232210102312-2202301122202223-2123320103033123"></a>

### Direct properties for `proxy_advertisement.advertise_dualstack_on_public.public_ip`

<a id="canonical-0330211001301212-3011302010302213-0211330220211020-2232022321302301-0022201023032033-0011301201223232-2102133012102013-2331002302100123"></a>

#### `proxy_advertisement.advertise_dualstack_on_public.public_ip.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0101313111332023-1323012133113303-2222132230032020-3233000220233120-3200232111321012-0301111221323132-2211112331313203-3203303110132323"></a>

<a id="canonical-2303020322000031-2111231321032212-3223131132232003-2312332003023323-0012102320223133-2113023002100210-1332010310120312-2331210131312210"></a>

#### `proxy_advertisement.advertise_dualstack_on_public.public_ip.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2300201110000333-1212300211222212-0112132022012221-0131003301103013-0002300202130132-0323311133210122-0200221131201102-2220111021322302"></a>

<a id="canonical-3121132001110020-2011222100001002-1122132012202111-0311333022222102-1123300221022311-1301032121203300-2213013000233300-1202020221312000"></a>

#### `proxy_advertisement.advertise_dualstack_on_public.public_ip.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3310110220212000-2133311323002011-3330302201310131-3311331023300100-3031000200023101-0321033223132322-1202213132212101-2303002101010130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_on_public` properties

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- proxy_advertisement.advertise_on_public

<a id="canonical-2201111330223031-3100333012303300-1322103101012023-3113003022102223-2113331212030300-2101002211303002-2312100311212232-3220203201203203"></a>

Type: `"single"`. Computed.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

<a id="canonical-3233102100000012-1120110002213123-2112320302230022-3100123030131223-2320000232223312-1102020132100121-3212311131102210-3200133122023032"></a>

### Direct properties for `proxy_advertisement.advertise_on_public`

- [public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-0303001021330002-0331320211123221-2320033130031302-3230331111330111-0012312332123311-3112122330203003-1222112033022233-1103301022002213): complete subsection reference.

<a id="canonical-0303001021330002-0331320211123221-2320033130031302-3230331111330111-0012312332123311-3112122330203003-1222112033022233-1103301022002213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3310110220212000-2133311323002011-3330302201310131-3311331023300100-3031000200023101-0321033223132322-1202213132212101-2303002101010130)
- proxy_advertisement.advertise_on_public.public_ip

<a id="canonical-0311110300000112-2122110131010333-2302313301012321-1233101230312222-2131300313320323-3223230201323122-0132320200003303-1112203021210312"></a>

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

<a id="canonical-0131202012313113-2020320202230232-0232220300030312-2123131223200012-3020301113100222-2213322033223300-3103010330031323-2012223111213233"></a>

### Direct properties for `proxy_advertisement.advertise_on_public.public_ip`

<a id="canonical-0313301030330320-3212022003311130-0111130200010033-1332003012230031-1020333003200220-1300010113332202-0300200322201132-2002103000132133"></a>

#### `proxy_advertisement.advertise_on_public.public_ip.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2100010122300000-1011033203333301-1010010023022330-2031030321120320-3120031123031132-1221130123230321-3110112301221131-2301322220130300"></a>

<a id="canonical-1321230331200023-3030332312231002-1101010210123213-2300210210221301-0023221102210113-1212023022233301-0321033100032302-2011310222330012"></a>

#### `proxy_advertisement.advertise_on_public.public_ip.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3010310230103100-1331012023321322-2300303011213312-2113121322022121-3103013300132103-3033310200021301-3230022310013123-1301311321113030"></a>

<a id="canonical-1223133033211332-2221031311302211-3232011233101032-1101200020312013-0201000131122011-1223203102131233-2002233200021111-1220113030203302"></a>

#### `proxy_advertisement.advertise_on_public.public_ip.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1131222002311300-2222212030001300-0012300103101232-0031023232101211-3302302310331123-2301120201231323-1120203233013132-2310220201330210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_on_public_default_dualstack_vip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- proxy_advertisement.advertise_on_public_default_dualstack_vip

<a id="canonical-0002131110333021-0323213313130330-0303222033120113-3121123012310001-0310022123203211-3132333211233222-0021222122221202-0311333213020313"></a>

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

<a id="canonical-0010111020311300-1202110002311312-3111100031020231-3233312101103231-2321101300233120-2111220013030021-0203011312322031-2032301213311120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_on_public_default_ipv6_vip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- proxy_advertisement.advertise_on_public_default_ipv6_vip

<a id="canonical-2100231102122202-1003100230223330-0201110101011120-0310333333020133-0210311033230211-0033222310022021-0001122310313100-0202211021231331"></a>

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

<a id="canonical-1022322020332313-1201203221223320-0201230131212313-1100002123112013-2313010311202123-1200121102113312-1212301113321232-0123230022132320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_on_public_default_vip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- proxy_advertisement.advertise_on_public_default_vip

<a id="canonical-0302120000223013-1311100103032323-1130112221333032-2022213100113122-3233203233211300-2000133011000311-1310023331233323-2121130232011303"></a>

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

<a id="canonical-3023010033002312-1211102200100233-0013330212203322-0033130332302222-1020020000213233-0021211213221212-1221003001332210-3122311210303213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_v6_on_public` properties

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- proxy_advertisement.advertise_v6_on_public

<a id="canonical-1101300330101121-0301312132001130-1100120010001300-3010321212103232-1010000210322330-2323121332332221-0032001133121212-2233330130210210"></a>

Type: `"single"`. Computed.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

<a id="canonical-1303332002033131-0112310103013222-3020123330203331-2011212311120211-2132100313021321-3312333133233332-1221331333113032-3111022033221213"></a>

### Direct properties for `proxy_advertisement.advertise_v6_on_public`

- [public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-0000322003033311-1021323203112330-0200122313332003-3020332323121223-3311020323221022-3203121320010213-2323130323121130-0230011313330231): complete subsection reference.

<a id="canonical-0000322003033311-1021323203112330-0200122313332003-3020332323121223-3311020323221022-3203121320010213-2323130323121130-0230011313330231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_v6_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3023010033002312-1211102200100233-0013330212203322-0033130332302222-1020020000213233-0021211213221212-1221003001332210-3122311210303213)
- proxy_advertisement.advertise_v6_on_public.public_ip

<a id="canonical-3100220323023013-2123101123301222-3300213013002132-2211313010132301-0100201000131223-0332010232210323-3031111311321013-0110011322011300"></a>

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

<a id="canonical-0221233331210222-2020021113202033-2333230003203203-2110012230031001-2221232022033221-1322210023210010-3302131030121130-3333331123333110"></a>

### Direct properties for `proxy_advertisement.advertise_v6_on_public.public_ip`

<a id="canonical-0031200223103212-2302232323030101-2332323010002310-0323212023001321-1321231202110001-1133313121312031-3331022002321302-2001033210120233"></a>

#### `proxy_advertisement.advertise_v6_on_public.public_ip.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3323103212220131-1332012031332222-2001313132200200-0300213113223102-3323201103321333-1320112020033002-1020212203333020-2013133310021331"></a>

<a id="canonical-0311211210131023-3210202001313130-1000113003323112-2001000001031302-1001203011313321-1301002231102213-1121012020000323-1113101123233101"></a>

#### `proxy_advertisement.advertise_v6_on_public.public_ip.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2201022200321133-0211213110211113-1230210323023132-2302223133020011-0233112121221232-3220120032022220-0000211013202012-2313232110222222"></a>

<a id="canonical-0223000010031131-3013102200030023-1100122212201110-3232131130010022-3312222332031121-0012331133333031-3130322212000301-0121130102011123"></a>

#### `proxy_advertisement.advertise_v6_on_public.public_ip.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1012133330232223-3320102120220023-3002110212112031-0330112322100132-3202033312322102-3301100130033303-3200110031011233-3223231020123322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.do_not_advertise` properties

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- proxy_advertisement.do_not_advertise

<a id="canonical-1200102101022133-0203120300303311-2321030000103133-0230002201020300-1210232213020300-2023300023002312-3221333213021123-1111131110311202"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for do not advertise.

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
