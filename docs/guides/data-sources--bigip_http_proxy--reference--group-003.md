---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-2221023022333132-3231101113032111-2022130110100010-1202002101222030-2112001012230220-1030331133132332-3311130113031221-3201322212132100"></a>

## Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_site`

<a id="canonical-1321100312313220-0232330212120230-1313330122022330-3222112023031123-2201212121300331-2120202202132210-2001110322211321-3013102133110020"></a>

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

- [virtual_site](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1320203302132111-0231031230302100-2313202000222012-3321132002322112-0103220133022231-2322002322223133-0003213323110023-1021131300311210): complete subsection reference.

<a id="canonical-1320203302132111-0231031230302100-2313202000222012-3321132002322112-0103220133022231-2322002322223133-0003213323110023-1021131300311210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1211323222210122-1130111030022212-0020323332120321-2103130323032233-0033123323232002-3233010033330003-1033312111111321-1311010213013021)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-0203212213312200-1213301303110312-3203122112033013-1311011201332031-0330022110310021-2213302111013020-3122231002013020-3210300213122333"></a>

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

<a id="canonical-3222130323310131-2202020010313313-1123332211221031-0000232202313123-2301320211301213-0212011032030003-3202332123023312-3000011313133120"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site`

<a id="canonical-2133213131130210-0110202300300232-2213213020120332-3001330130001000-1102023111212123-3130312003003102-3320022113131001-1231122222230230"></a>

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

<a id="canonical-3213132102033000-2032330310332303-2002220111022100-3222231032013221-3102222203310220-2101303210301200-2031320212020123-1222131012100001"></a>

<a id="canonical-0101233223200330-2203202112030101-3320110002310010-3130311122310033-1302203021110303-1303311203323300-1210211211103012-3013203033020301"></a>

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

<a id="canonical-2201310332223313-1331031201211000-2113310101220020-0013203001103302-1201321302101010-0031330210322232-3120110030210102-3120030303302111"></a>

<a id="canonical-3312010220120300-2320332330210013-1012021010113031-3020312220323123-2213111022221220-2321203201130113-2122001220000202-1322213231302030"></a>

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

<a id="canonical-0301313112323223-1130020222220223-2203231102032211-2233233211022332-1020011330022232-0201311310121232-0123123020333210-0022212123000333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-1122132330121232-1132313130031103-3101320012332120-0221203002210100-2023000232031100-3103301003112101-0320222303023330-1022332223232310"></a>

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

<a id="canonical-1233312111012311-1030011232230120-1020232222320203-2221130012000301-1300002223100223-3331020330100023-2133320111130030-3222130231210223"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip`

<a id="canonical-0303222303201033-0323200223210132-1012320323021000-0223222131311223-1312303203013102-0101230122232010-2202113021212312-1330322200032111"></a>

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

<a id="canonical-2130223322111333-3111322112131022-2003330010333310-1003022022202232-2231332323312012-1021102232021212-0113323001200201-1011222232013322"></a>

<a id="canonical-0222222200302202-3311131222023323-1132013203110220-1021101310002332-1303313223111303-3101322312011330-2230223220321011-2131323233221212"></a>

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

- [virtual_site](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2330311332201012-3220103023330122-3030002132113020-3102032002322313-2112210203012332-1302000031112122-2000222113121022-3002121100323032): complete subsection reference.

<a id="canonical-2330311332201012-3220103023330122-3030002132113020-3102032002322313-2112210203012332-1302000031112122-2000222113121022-3002121100323032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0301313112323223-1130020222220223-2203231102032211-2233233211022332-1020011330022232-0201311310121232-0123123020333210-0022212123000333)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-1100303310300221-1313121023300131-2201030032333231-0310000230030031-1311123203011031-1333332011023112-0113013221000310-3311001321113330"></a>

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

<a id="canonical-0032231021002112-0112002113233303-1210212031123023-2231122231312310-3302330333110012-3211310100332321-2301220231320230-3113220000203320"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site`

<a id="canonical-1132302322102021-2211132030220312-0112201013010021-2133013122002221-0301331323113110-0000320310302223-3021220123101320-2201030320103110"></a>

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

<a id="canonical-3230113001133133-0013333203311312-2113131221232212-3322112122300132-2030322011302102-3023333011220032-2210222201231322-0321110031313000"></a>

<a id="canonical-1130232133133110-1001101312120021-3213220223033332-2321330203223101-2010322002033211-0120311230122111-0023130130112321-1201332321120203"></a>

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

<a id="canonical-0110011202131232-0123222120301122-2022202123023313-2310203021303331-1123300113322132-1123121032111223-2331020230323213-1012100033100303"></a>

<a id="canonical-1103102103203200-1322310110202301-1323023200121101-2311233302002320-1103111023133031-0321103310221220-0111202131302200-3010310020331330"></a>

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

<a id="canonical-1023212332121021-1203300333013112-0132223101030032-3311122322201033-2123023230101031-1310123320320032-3210221320220021-1232101110313030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service

<a id="canonical-3302102031323013-1031023301100102-2030020221113121-1332322010121203-2333330330001313-2223100110003332-3321112320300222-1222012101311032"></a>

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

<a id="canonical-0313330321000312-1303200021120001-2120002120202011-0321232013110011-1103320133012231-1021333330203330-0210101112230031-3102213202011122"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.vk8s_service`

- [site](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1000332312222102-1100331022002101-2312302023120101-0023112133322111-0311202131330130-2322010132231001-2123030231132130-1103030101031122): complete subsection reference.

- [virtual_site](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0333021200303212-1103101033213311-1301021220312333-0020203310331111-2013000032022332-3313300310233010-1011203231101021-1331222120112301): complete subsection reference.

<a id="canonical-1000332312222102-1100331022002101-2312302023120101-0023112133322111-0311202131330130-2322010132231001-2123030231132130-1103030101031122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1023212332121021-1203300333013112-0132223101030032-3311122322201033-2123023230101031-1310123320320032-3210221320220021-1232101110313030)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-0012200233112331-3223031211012323-2331102223220223-1322020003301010-2012103302300332-1110310023233310-1331212311133120-0303300231320202"></a>

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

<a id="canonical-0020002031130303-1123013200102322-1313220333101031-3200311110110223-2102123303210222-1023031333020022-2231100102330103-0333121323221020"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site`

<a id="canonical-1223201301030011-1232002321002123-2313000302320031-1200030032310321-1020131123101211-3322021200122220-2230303313331012-2021002111213322"></a>

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

<a id="canonical-3331311100213120-0311223332022213-0023022122101011-3111211133030201-0332011010021211-1203232000002030-1300312233101112-2322123003332203"></a>

<a id="canonical-1111031301021013-0330103032010131-3210113302211333-2111221031102220-2320302031200320-2311323332112002-1312311011212130-1102033110230211"></a>

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

<a id="canonical-1022320212003032-2021131202222131-3012120331333010-0012031330112213-3023333230332212-3201303201023220-3022333120330232-0201020213113033"></a>

<a id="canonical-3032211331021302-0222220201223331-3210230230131223-0233231133113201-0302123013232230-3010011320001323-0130021100033013-2031113122200221"></a>

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

<a id="canonical-0333021200303212-1103101033213311-1301021220312333-0020203310331111-2013000032022332-3313300310233010-1011203231101021-1331222120112301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1023212332121021-1203300333013112-0132223101030032-3311122322201033-2123023230101031-1310123320320032-3210221320220021-1232101110313030)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-2300231330123212-3103010303300101-0213100331032030-3311201123122011-3110131301020021-1313311120013120-2113123033032310-0000233300122230"></a>

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

<a id="canonical-2032021010012021-0132232212131233-0032332301133102-1232110323201303-2330101102112103-1130013221101300-0002333311110200-1013110031022033"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site`

<a id="canonical-3131132331120122-3223303301123301-2101232000130133-0110213331001212-1023113003302020-2323320003000001-1113122310320223-3113313331132212"></a>

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

<a id="canonical-2320302213133330-2003233320333010-3120322322022011-3100302220212012-0032213221203032-0321231123031230-2003331231121121-1310331311310020"></a>

<a id="canonical-0311310021022033-0200123103122222-0220003200003323-0100003203321033-3202202103023011-2120312132002321-3321112312212303-0112301023231021"></a>

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

<a id="canonical-3103322221002010-3112120013320302-0112122121211133-0303332320303230-2220021313200122-2301030113111032-3133020232120111-2212201211220030"></a>

<a id="canonical-0111231020132322-1023310111011201-2133021232121020-0101030322213310-3201131122202211-1330022032123023-0101222123213020-2210210331232031"></a>

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

<a id="canonical-3312032130021330-3212102123320232-0213011310323122-0022113030030231-3323210230320221-1030023032302123-3132212013230101-0110213332222021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.do_not_advertise` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- proxy_advertisement.do_not_advertise

<a id="canonical-3031312323231020-0113200313213210-0033130220231123-1200131232031233-2103303311133022-0312121011122032-3321311301333313-0022033223202221"></a>

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

<a id="canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- proxy_config

<a id="canonical-1231210333102133-2313202213322022-2310312222033130-2121022032332201-0122120101110200-0120012230031221-1133223213031110-3132012221102322"></a>

Type: `"single"`. Computed.

HTTP/HTTPS Load Balancer. HTTP/HTTPS Load balancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]"
}
```

<a id="canonical-3132311111123002-0110211303010120-2303113300123310-3230212312202323-2212222311120330-3322023103333321-1331011231101233-2222303220233013"></a>

### Direct properties for `proxy_config`

<a id="canonical-3230213210313020-2020303033023122-2323101130120200-1313312022031113-1112113100102300-0022132002300312-3311322232313120-3133303303320000"></a>

#### `proxy_config.domains` property

Type: `["list", "string"]`. Computed.

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*-bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*-bar.example.com\`\` will match
\`\`baz-bar.example.com\`\` but not \`\`-bar.example.com\`\`. The longest wildcards match first.

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [http](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0011021330202210-0002020013211203-3031010020133013-3110000232312130-0001313101300321-0312001331300310-1333323101313213-2101212130332011): complete subsection reference.

- [https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321): complete subsection reference.

- [https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330): complete subsection reference.

<a id="canonical-0011021330202210-0002020013211203-3031010020133013-3110000232312130-0001313101300321-0312001331300310-1333323101313213-2101212130332011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.http` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- proxy_config.http

<a id="canonical-0201322302313330-2111332301120033-2311111103100321-2223001033000122-2303030300200010-0123002001300030-1231120123232113-3331133131313001"></a>

Type: `"single"`. Computed.

HTTP Choice. Choice for selecting HTTP proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

<a id="canonical-3233100031122321-3111000023333301-0023103030131332-0223023033301310-0233020220010023-1311000112303031-1203010130210231-0003303131313211"></a>

### Direct properties for `proxy_config.http`

<a id="canonical-1210201310033332-1003311010320030-3133110110022310-0002213030132210-0202003123220012-1131033011023200-1333202133031030-0022132022202320"></a>

#### `proxy_config.http.dns_volterra_managed` property

Type: `"bool"`. Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

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

<a id="canonical-1023000321232100-3033123133210311-0033302131012302-3330311301123220-2133120132310030-0233033313132110-2221113122320312-2110022001032332"></a>

<a id="canonical-1021002122123330-1203312000222311-1131333121033030-2332111013113332-2002111232120000-2221321020121033-1103322223213011-2210303300112320"></a>

#### `proxy_config.http.port` property

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTP port to Listen.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3232003033211322-0311310003213102-2003220133301100-3120331330312313-1010033232123131-0211002112220210-0311000132030102-2222131220211202"></a>

<a id="canonical-3032103032031330-0011013321011301-0212301333201302-0330020211321112-1221230112210113-1110110210020233-3130312201232032-0101212231032102"></a>

#### `proxy_config.http.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- proxy_config.https

<a id="canonical-3131210001331111-1202123021322110-2023023311130330-0230100302301332-3210232200200233-3231011021323203-0223301322310112-2131200112121210"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

<a id="canonical-3203002313200223-2213210311200230-0010201133212121-0223031023320031-2123231312303132-1133013303102010-1113100030003200-3033002033010301"></a>

### Direct properties for `proxy_config.https`

<a id="canonical-1033332002033311-0310122232131012-3000123133001122-1133033130213120-1012202021300310-1113210103231132-3120110302221133-1230033211020032"></a>

#### `proxy_config.https.add_hsts` property

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header.

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

<a id="canonical-0223203231112032-2102010131131012-0010222220012032-2320121200110231-0201231312120210-3313312210212133-2103233013121220-2322213331032110"></a>

<a id="canonical-1312303230003233-3131233303110123-2230230232102003-1310131211222103-0131212330232320-3223033102123123-1211022032313000-3213003301120033"></a>

#### `proxy_config.https.append_server_name` property

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1010220232111020-0012112302000310-0323101031023100-1131011320103303-2213211001212310-2123110123021133-0300131020331201-1303223232212322): complete subsection reference.

<a id="canonical-2233330131102310-0213111020102010-3103212210033110-1323001302233332-3030301321022101-1100131331000023-2301102112310332-2120310102200102"></a>

<a id="canonical-1022130211002112-2231000333031021-1123221102300101-3132201222103101-1300300230211323-2203101013022302-1301302201211201-0110003033110132"></a>

#### `proxy_config.https.connection_idle_timeout` property

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0331121010100121-1223320123102002-2331110023123211-2303112313223333-2133310113300122-2002110011321130-1233020121210300-1001201113203300): complete subsection reference.

- [default_loadbalancer](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1313301030300323-0311033002233211-1100331312211011-1102320003312100-3100332020323101-0103312113202220-2231202312103223-3032221212311332): complete subsection reference.

- [disable_path_normalize](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3011231311323101-2323203002201101-2000230001120021-2313313201321132-0313030103323133-0120113203223211-0223200330233310-0301312223212022): complete subsection reference.

- [enable_path_normalize](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0222102223110101-2311000212121200-1022311003231221-3301110103020210-0322312322100322-1023011012132101-1330012032323331-1102031231122003): complete subsection reference.

- [http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010): complete subsection reference.

<a id="canonical-0200210020210213-3213310120131313-3221320311112110-2100001202031232-2311022013033230-0211323002212202-2131212311012203-1311323331322023"></a>

<a id="canonical-3123003111211330-1002230333231322-0302301022021110-2103133300112110-1213131321303011-2333211330100313-1130101003331032-3000032322233302"></a>

#### `proxy_config.https.http_redirect` property

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

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

- [non_default_loadbalancer](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0020110300033033-3031303222203120-0333011030102032-1223202132032201-0021211131323220-2101023111332012-1233131232230211-3202203330302201): complete subsection reference.

- [pass_through](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0202323330310033-1233203111130031-0320101100321323-0002103202032223-1323222010212022-1111323110203232-0232022132031303-1112203333301222): complete subsection reference.

<a id="canonical-1101220131221202-1212023323213220-0230133033002030-1330101003300333-2310032112031011-1213110221132113-1113030113321013-0102010112230020"></a>

<a id="canonical-0230110113130032-2331030033021102-3133132022312133-3312220020013301-1022201010111211-2121102313121130-2110133212231003-0223233002101113"></a>

#### `proxy_config.https.port` property

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1022101111013211-1231330120003213-2212133223010003-0230311311112201-1121323200233003-2302011111230103-0121010203203021-0301122101221311"></a>

<a id="canonical-3020231202321033-1030011322102032-2202213003002322-3313202000133211-2122011332231122-0121332211201023-3200021330323112-1122000112113212"></a>

#### `proxy_config.https.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-1332220233223222-1031003231203133-0201002030200123-0320021221312111-3331330023012203-0332310332313210-2011132312021320-2210220112111222"></a>

<a id="canonical-0232213003000213-1001332133210110-1022211200232103-2010321031331302-0120213111003101-3132111120033210-3133003003231301-1033123201001230"></a>

#### `proxy_config.https.server_name` property

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320): complete subsection reference.

- [tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100): complete subsection reference.

<a id="canonical-1010220232111020-0012112302000310-0323101031023100-1131011320103303-2213211001212310-2123110123021133-0300131020331201-1303223232212322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.coalescing_options` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.coalescing_options

<a id="canonical-0223231011323112-3000232123323222-0332233310120310-3001030300330231-3333200112321222-2333101032331223-3233012033000113-1013110120100322"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

<a id="canonical-3022210122321122-2001000232033321-1113013130201233-3101030001230002-3033113001213332-1220013220023122-0213023120202301-2002321031003121"></a>

### Direct properties for `proxy_config.https.coalescing_options`

- [default_coalescing](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0201301100311301-1230220231000123-2233210031232121-1103320120113103-3000203121201132-3030130233012232-2310100331002020-3321010322121200): complete subsection reference.

- [strict_coalescing](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3000101003000001-2010013302122010-3010031131312123-1001210321223233-2012120213223202-3220021213011122-2032131333230223-0220112231020330): complete subsection reference.

<a id="canonical-0201301100311301-1230220231000123-2233210031232121-1103320120113103-3000203121201132-3030130233012232-2310100331002020-3321010322121200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1010220232111020-0012112302000310-0323101031023100-1131011320103303-2213211001212310-2123110123021133-0300131020331201-1303223232212322)
- proxy_config.https.coalescing_options.default_coalescing

<a id="canonical-3132130121020200-1300333001210012-2311212022321233-3322322020003031-1330023133132101-0212313130211000-3131112022021112-2110313130020002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-3000101003000001-2010013302122010-3010031131312123-1001210321223233-2012120213223202-3220021213011122-2032131333230223-0220112231020330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1010220232111020-0012112302000310-0323101031023100-1131011320103303-2213211001212310-2123110123021133-0300131020331201-1303223232212322)
- proxy_config.https.coalescing_options.strict_coalescing

<a id="canonical-3102333313000021-0230033211232111-3120020123023330-3113203323133023-1311301301013100-3203302311230302-0300221320311203-0322000303303103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-0331121010100121-1223320123102002-2331110023123211-2303112313223333-2133310113300122-2002110011321130-1233020121210300-1001201113203300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.default_header` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.default_header

<a id="canonical-3311022003010330-0320020132023301-2110032022032111-0123300032111330-3101032303020332-0330110132002333-3130312030302322-3303222033330030"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

<a id="canonical-1313301030300323-0311033002233211-1100331312211011-1102320003312100-3100332020323101-0103312113202220-2231202312103223-3032221212311332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.default_loadbalancer

<a id="canonical-2303231011310021-2121212100000201-2231210333303212-0101131301121120-3300202303301302-2030222012131020-3003202301032210-0111002003101313"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default loadbalancer.

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

<a id="canonical-3011231311323101-2323203002201101-2000230001120021-2313313201321132-0313030103323133-0120113203223211-0223200330233310-0301312223212022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.disable_path_normalize

<a id="canonical-1111030201013103-3030332220030332-1303232300120201-1230223320032332-2013002313202013-1020001333213210-3333010100231210-2001123300333111"></a>

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

<a id="canonical-0222102223110101-2311000212121200-1022311003231221-3301110103020210-0322312322100322-1023011012132101-1330012032323331-1102031231122003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.enable_path_normalize

<a id="canonical-2231231313011132-3322020313011102-2121111311033203-1203133323122223-2312030121223312-3300012011023103-1100122221111221-3331121110132233"></a>

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

<a id="canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.http_protocol_options` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.http_protocol_options

<a id="canonical-2020232033203000-0122333001010013-1013233013122011-2123211103311130-0331303022202031-2321020020100120-3032203031221001-1332313132211202"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

<a id="canonical-3012113330000301-3202132202111113-1120323233232213-0332312012000310-1112130101010001-2210100110321100-1200310221200112-1332300021210223"></a>

### Direct properties for `proxy_config.https.http_protocol_options`

- [http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3233023112203213-1011001323132233-1220102230110322-2210300302000131-3011103312112021-0113211223102301-2222303003302201-0100031132101201): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0222230210022000-1200302133300220-1010303223330123-0200323010132223-3333331033303302-1331011023131010-1212131032200301-2233303112321313): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3231020003213122-0011232310032202-2031100100021022-2023321111122223-0312313330210300-2100301102131223-3013030123222221-0110333213303222): complete subsection reference.

<a id="canonical-3233023112203213-1011001323132233-1220102230110322-2210300302000131-3011103312112021-0113211223102301-2222303003302201-0100031132101201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-0312330010131033-2113021130013333-3311222321101023-3301101332031012-0223120330321002-0013222321131022-1003303300133320-1110030102223220"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for downstream connections.

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

<a id="canonical-0011033301302123-1202323101210310-0102103110110233-0023011030020101-2222201032022213-2030210122113101-2332213021302120-2332120222210221"></a>

### Direct properties for `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1331100001023220-0230220013312022-0302023012213121-1230022100210132-2201121001111331-2300313313231221-3013332233322120-0212010103100201): complete subsection reference.

<a id="canonical-1331100001023220-0230220013312022-0302023012213121-1230022100210132-2201121001111331-2300313313231221-3013332233322120-0212010103100201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3233023112203213-1011001323132233-1220102230110322-2210300302000131-3011103312112021-0113211223102301-2222303003302201-0100031132101201)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-1310032333303232-2300012323233112-2132020103112102-0122213120020001-3331112213011302-0103110211202313-0100020123221012-1031031121033320"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

<a id="canonical-2022120212111303-3230012102302101-1000301032212300-0231200321203030-1301331301100020-0213033321333112-3122003121323210-2123002012201033"></a>

### Direct properties for `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0000310123233001-2021232010200113-2232111321012130-3112323131323233-3122210010023211-2031102203113003-1023201223022102-2133113211220020): complete subsection reference.

- [preserve_case_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0223202012222003-3210230110011202-0012302020011112-2233313012231223-1021031322131211-1323300231032100-0321233030133330-1323303313331012): complete subsection reference.

- [proper_case_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0123231201002122-1311233120132130-3022220110013000-2111212311131102-2001213202222011-2230201120301321-1101130022130200-0210121121300130): complete subsection reference.

<a id="canonical-0000310123233001-2021232010200113-2232111321012130-3112323131323233-3122210010023211-2031102203113003-1023201223022102-2133113211220020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3233023112203213-1011001323132233-1220102230110322-2210300302000131-3011103312112021-0113211223102301-2222303003302201-0100031132101201)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1331100001023220-0230220013312022-0302023012213121-1230022100210132-2201121001111331-2300313313231221-3013332233322120-0212010103100201)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-3002202220301331-1023230122223303-1123331023010202-2311332132123212-0301202300121011-2312320323210203-2312122310011313-3000112322000020"></a>

Type: `["object", {}]`. Computed.

Use the platform's current default HTTP header transformation behavior.

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

<a id="canonical-0223202012222003-3210230110011202-0012302020011112-2233313012231223-1021031322131211-1323300231032100-0321233030133330-1323303313331012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3233023112203213-1011001323132233-1220102230110322-2210300302000131-3011103312112021-0113211223102301-2222303003302201-0100031132101201)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1331100001023220-0230220013312022-0302023012213121-1230022100210132-2201121001111331-2300313313231221-3013332233322120-0212010103100201)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-2132310332210102-2230333101302022-1301321113101011-0303313201231311-2232010310031303-2301212233322300-3112211303013012-3031333211021333"></a>

Type: `["object", {}]`. Computed.

Preserve HTTP header-name case when upstream case must remain unchanged.

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

<a id="canonical-0123231201002122-1311233120132130-3022220110013000-2111212311131102-2001213202222011-2230201120301321-1101130022130200-0210121121300130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3233023112203213-1011001323132233-1220102230110322-2210300302000131-3011103312112021-0113211223102301-2222303003302201-0100031132101201)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1331100001023220-0230220013312022-0302023012213121-1230022100210132-2201121001111331-2300313313231221-3013332233322120-0212010103100201)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-0022233102323321-2202130121021321-3022113221120012-2133100300302121-3312300100323020-2331301011203112-2210113103100301-3331220311022212"></a>

Type: `["object", {}]`. Computed.

Transform HTTP header names to proper case when explicit transformation is required.

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

<a id="canonical-0222230210022000-1200302133300220-1010303223330123-0200323010132223-3333331033303302-1331011023131010-1212131032200301-2233303112321313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2213131012001100-0113012110132301-2032101031231233-0300102231313021-0033312202320102-0032110130132220-0320333001100211-2312221321320313"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

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

<a id="canonical-3231020003213122-0011232310032202-2031100100021022-2023321111122223-0312313330210300-2100301102131223-3013030123222221-0110333213303222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- proxy_config.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-2232132303313112-0200000020322011-2312131010032230-1133220100302323-2313203231132131-1100312111222313-2020013221221200-3212233222133223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

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

<a id="canonical-0020110300033033-3031303222203120-0333011030102032-1223202132032201-0021211131323220-2101023111332012-1233131232230211-3202203330302201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.non_default_loadbalancer

<a id="canonical-0210223233031120-1000222222333132-2232011202202123-1323202312132210-3220221210321110-2121233222203100-1002020110113313-3320223323010312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

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

<a id="canonical-0202323330310033-1233203111130031-0320101100321323-0002103202032223-1323222010212022-1111323110203232-0232022132031303-1112203333301222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.pass_through` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.pass_through

<a id="canonical-0332022220001232-1320033131013103-3110333312203221-3201111230332103-1020312003332212-2100300020002312-2332002321120110-3312010012132321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

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

<a id="canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.tls_cert_params

<a id="canonical-3122310321332130-3023333001301322-0022230033301010-0131301212333330-1122301332230332-1223002012120301-2023012112110013-2230001312331233"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Additional upstream details:

Select TLS Parameters and Certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-3203321133030233-2132010102222010-0112331103023222-2111203330133210-1300302321303010-2002201210301120-3100323100222122-2033210021013133"></a>

### Direct properties for `proxy_config.https.tls_cert_params`

- [certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0022323232333203-1012222123002210-1310333231033311-0212213221333303-1120213321301313-0002212330132130-0133000303310211-3232133333203200): complete subsection reference.

- [no_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0120000123102023-2110310333100130-1323131321110301-3112223211301323-1113211100021123-2200111000330010-1013210122222003-0100101303223011): complete subsection reference.

- [tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023): complete subsection reference.

- [use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303): complete subsection reference.

<a id="canonical-0022323232333203-1012222123002210-1310333231033311-0212213221333303-1120213321301313-0002212330132130-0133000303310211-3232133333203200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- proxy_config.https.tls_cert_params.certificates

<a id="canonical-0012010113200032-1223311313032332-0112311213223312-3313231133311231-1011023100130310-2232130112302203-1220310022321221-3013131330301021"></a>

Type: `"list"`. Computed.

Select one or more certificates with any domain names.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1203021211333022-2023010311020013-3030000322320030-1033012213120322-0011003123132330-0231313331330120-1311110323333221-1221102113030230"></a>

### Direct properties for `proxy_config.https.tls_cert_params.certificates`

<a id="canonical-3113212031330030-0033131021032003-3121120032313332-3130033230202101-1222231220312122-2012331330023023-1003121031121310-1012133132010220"></a>

#### `proxy_config.https.tls_cert_params.certificates.name` property

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

<a id="canonical-2300202210323110-2321000122221301-2331100300223320-2032010101023033-1103011233022101-2333131112110021-1302333322130131-1111332030123303"></a>

<a id="canonical-2312033213200023-0323123211030200-3333030320211131-3333322000133102-2321203122113002-1013201023100202-0113303332210203-2310301311221232"></a>

#### `proxy_config.https.tls_cert_params.certificates.namespace` property

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

<a id="canonical-0210320223023102-3312023001031210-1030223113322213-2200022333222200-1202302303102002-2222121323122113-1100203230102033-2323120132021133"></a>

<a id="canonical-0202310212111023-2122110332331301-0301000212033212-1031010223302133-2101103321320000-1333132133333031-2321232030300120-0112131222120320"></a>

#### `proxy_config.https.tls_cert_params.certificates.tenant` property

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

<a id="canonical-0120000123102023-2110310333100130-1323131321110301-3112223211301323-1113211100021123-2200111000330010-1013210122222003-0100101303223011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.no_mtls` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- proxy_config.https.tls_cert_params.no_mtls

<a id="canonical-2102303322033200-1222032222111111-3222021312212332-1203120020123222-0203120132001203-0133032000211022-1012212201022013-2302332220120212"></a>

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

<a id="canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.tls_config` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- proxy_config.https.tls_cert_params.tls_config

<a id="canonical-3331311330100132-3321222311203013-0322132311102103-3111101333133032-3022112121202201-1133332222201330-2310300210303020-2210033103123021"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-0213013132331203-3110222133302300-3012232230011221-1211331333332332-1011303023312031-2332213011130130-3133302123203223-3231022011312300"></a>

### Direct properties for `proxy_config.https.tls_cert_params.tls_config`

- [custom_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0200200113202333-1213031102323202-1233223223020000-2211020331232110-1320131001110031-1133120131212233-0322033200310121-3031313033323210): complete subsection reference.

- [default_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2222233232033312-1112222110113100-3000132201001332-3220221331132200-3201022133222321-0221203300200212-3112322232100030-0231022321133012): complete subsection reference.

- [low_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1202300221110131-1303212030233202-0013311003011331-0022322101032012-2123221213133331-0230113231000313-3313111110102130-1313033003113301): complete subsection reference.

- [medium_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1131213330232321-0201220033010230-0203023031001232-0123131112131011-2111101132132020-3302211322031211-2320003213100202-2133312312020011): complete subsection reference.

<a id="canonical-0200200113202333-1213031102323202-1233223223020000-2211020331232110-1320131001110031-1133120131212233-0322033200310121-3031313033323210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023)
- proxy_config.https.tls_cert_params.tls_config.custom_security

<a id="canonical-0202303133113310-3000002010303332-0013312210203002-2121222322131221-2122320102211002-0022222132233030-3322310113211333-1303120103013122"></a>

Type: `"single"`. Computed.

This defines TLS protocol config including min/max versions and allowed ciphers.

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

<a id="canonical-0110121132131233-2321322331102212-3222313011220230-0212330322210333-1211233212103320-2233303030223022-1230022222100210-1031213101103203"></a>

### Direct properties for `proxy_config.https.tls_cert_params.tls_config.custom_security`

<a id="canonical-2103322232331000-1130011110013132-2210120302100021-0220231022100110-2331011130223100-3102113300330002-1312221131033221-1110230120232001"></a>

#### `proxy_config.https.tls_cert_params.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1230012312302133-3033300223120113-2233333230333120-3211212030211321-1131010311300111-0032213312220111-2033033300033311-3230113110323301"></a>

<a id="canonical-1103023310301100-2310212001111331-0231011021312223-2022022120222221-2012330300031013-0000030110101022-2312121213303133-0032013122232330"></a>

#### `proxy_config.https.tls_cert_params.tls_config.custom_security.max_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3233021320101312-3322002332302312-1103332310202200-0300002232300102-0112232333131330-2302333312133123-1032201012301110-3332200322023011"></a>

<a id="canonical-3200020030022030-2233102200323021-1012301321300230-2020012022200022-1323212102023101-1220233101230221-2031320022330121-2113330120321323"></a>

#### `proxy_config.https.tls_cert_params.tls_config.custom_security.min_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2222233232033312-1112222110113100-3000132201001332-3220221331132200-3201022133222321-0221203300200212-3112322232100030-0231022321133012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023)
- proxy_config.https.tls_cert_params.tls_config.default_security

<a id="canonical-2011131020303231-0123323201330332-0100220102332203-2121210113201232-0320230333022103-0303310022103201-1003323223010120-0022110033210301"></a>

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

<a id="canonical-1202300221110131-1303212030233202-0013311003011331-0022322101032012-2123221213133331-0230113231000313-3313111110102130-1313033003113301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023)
- proxy_config.https.tls_cert_params.tls_config.low_security

<a id="canonical-0100201211030102-1220202111320022-2130213102221113-1230132000300222-2231111201020112-1001223011203132-3321022013321111-2213020030331301"></a>

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

<a id="canonical-1131213330232321-0201220033010230-0203023031001232-0123131112131011-2111101132132020-3302211322031211-2320003213100202-2133312312020011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_config.https.tls_cert_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023)
- proxy_config.https.tls_cert_params.tls_config.medium_security

<a id="canonical-2210202220021222-0010200321103210-3100212230303022-0212233233102310-2211021233303230-0101011113203123-1213302300221002-2300113312231300"></a>

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
