---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-0100013013213323-3130321013002331-1030102012333312-0031031112003311-3231312030013310-3022231023112022-2201311333120100-3303322111133133"></a>

## secondary_nameserver property — segment_config / 132013300310 / 5

Type: `"string"`. Computed.

Optional Secondary IPv4 DNS server to be used for name resolution.

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

- [static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3013013210212323-2130222123331122-3102220321233132-3211131233310133-1323001231230321-0311212023230121-1333121113123323-1203333212113331): complete subsection reference.

- [static_v6_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1303120210133032-0020301133000323-3311321220022302-3113222021233210-2213313002331302-1100333233332222-0121333220032113-2133332330221312): complete subsection reference.

<a id="canonical-0203321230202323-0130322121010202-0102301320223102-1332001132112321-1131322300113030-0303302223003222-3312213022011203-2230221320022112"></a>

## Next pages — segment_config / 132013300310 / 6

- [segment_vrf.segment_config.no_static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2213212113001023-3222200101023002-3003233112232032-1121202333120003-3121311012003003-1100013010011131-1000330210221100-0331203103033130)
- [segment_vrf.segment_config.no_v6_static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1313023123020211-0120123112112300-3221032032302320-3031213110010011-3001133223321001-3102021223330110-0301020220220231-2001111231133132)
- [segment_vrf.segment_config.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3013013210212323-2130222123331122-3102220321233132-3211131233310133-1323001231230321-0311212023230121-1333121113123323-1203333212113331)
- [segment_vrf.segment_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1303120210133032-0020301133000323-3311321220022302-3113222021233210-2213313002331302-1100333233332222-0121333220032113-2133332330221312)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2213212113001023-3222200101023002-3003233112232032-1121202333120003-3121311012003003-1100013010011131-1000330210221100-0331203103033130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021012312300122-3300033111323131-3220213231023001-3132100003221333-0032030310030002-0213300232231133-2013312102102010-3103113132123023"></a>

## segment_vrf.segment_config.no_static_routes — no_static_routes / 030131102310 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- segment_vrf.segment_config.no_static_routes

<a id="canonical-2033323313123131-1320001203232102-0333203222230131-1103003133210330-1100021311331202-3033003013330320-3020010301111313-3110010132201100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no static routes.

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

<a id="canonical-2331100021013133-3302302230231030-0031211303212203-3233200233010230-2122211321231111-1332332322231031-1023331231222100-1111202002231303"></a>

## Direct properties — no_static_routes / 030131102310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221313320312000-3320323002012110-0033023111231111-2102332123123313-3130131202333010-2203232303213322-2121022301011013-3210301200330111"></a>

## Next pages — no_static_routes / 030131102310 / 4

- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1313023123020211-0120123112112300-3221032032302320-3031213110010011-3001133223321001-3102021223330110-0301020220220231-2001111231133132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123322102322311-2323223200313212-0321202222221212-2333103111301201-3120113230032221-0233311213323111-1221321221012312-2321212010110212"></a>

## segment_vrf.segment_config.no_v6_static_routes — no_v6_static_routes / 013020320222 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- segment_vrf.segment_config.no_v6_static_routes

<a id="canonical-1120232021133000-2311302200023323-0322002123322003-0323330023022200-3122131032230230-2013023112200020-0200113001210122-3220101223132100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no v6 static routes.

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

<a id="canonical-3313333221132310-3120132102311123-0131202020303112-3312121212221021-2001233331303200-3020010221200300-3203101131220002-0133200312120210"></a>

## Direct properties — no_v6_static_routes / 013020320222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103330330301110-1333122331312022-2120200202202221-2200332133102203-2022231002003203-0032312031110020-0113313210120330-0021222220111002"></a>

## Next pages — no_v6_static_routes / 013020320222 / 4

- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3013013210212323-2130222123331122-3102220321233132-3211131233310133-1323001231230321-0311212023230121-1333121113123323-1203333212113331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100132322311030-2012131020202120-2102103031032332-1310310233203033-3003020130213110-0020313021031202-0022002311130020-3113122320122113"></a>

## segment_vrf.segment_config.static_routes — static_routes / 311100110211 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- segment_vrf.segment_config.static_routes

<a id="canonical-2001331030002323-1201301113111300-0010122200310311-3001200033211113-2212023123302022-2133000101023333-1311302331213222-2020203202323232"></a>

Type: `"single"`. Computed.

Configuration parameter for static routes.

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

<a id="canonical-3132103121321332-3023202332312333-1310232323313030-3120301000311113-3331312220210201-0031013200333322-0130200201332000-2232000320031110"></a>

## Direct properties — static_routes / 311100110211 / 3

- [static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0310112223023333-0331032222330012-0023123003302222-3100030000102221-3333323013021300-3203311323301121-3323203213212332-0332221210020010): complete subsection reference.

<a id="canonical-1210333212313330-0122023001203320-1302221213003333-3012232032001232-2213332130002213-1011213103013203-3002131312030200-2100223202203022"></a>

## Next pages — static_routes / 311100110211 / 4

- [segment_vrf.segment_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0310112223023333-0331032222330012-0023123003302222-3100030000102221-3333323013021300-3203311323301121-3323203213212332-0332221210020010)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0310112223023333-0331032222330012-0023123003302222-3100030000102221-3333323013021300-3203311323301121-3323203213212332-0332221210020010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220230001100023-3113331101001310-0101013311021232-2013013322020230-0331310110311101-1220201220112302-1323033231223201-0130221332012303"></a>

## segment_vrf.segment_config.static_routes.static_routes — static_routes / 011233020103 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- [segment_vrf.segment_config.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3013013210212323-2130222123331122-3102220321233132-3211131233310133-1323001231230321-0311212023230121-1333121113123323-1203333212113331)
- segment_vrf.segment_config.static_routes.static_routes

<a id="canonical-0120121232013133-1232211231101021-1132032213012121-1021013302321121-3032111111230120-2030000231322123-2132311321103311-1131311111311020"></a>

Type: `"list"`. Computed.

Configuration parameter for static routes.

Upstream description:

Configuration parameter for static routes

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0302212302310022-3301202113311031-2003223332033133-2013330211111021-2201322230322223-1133133113322323-2103232312220232-1130211113112333"></a>

## Direct properties — static_routes / 011233020103 / 3

<a id="canonical-2322021221323322-2303103321203012-2032020320222330-2221133310131023-2010303221130230-3322012302120030-0121121200122021-2211313222330121"></a>

<a id="canonical-0212220123132030-1320111113312021-2102203001103201-1323023312313000-0103100323301301-1132110211120301-0112131222010322-2321311113120232"></a>

## attrs property — static_routes / 011233020103 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3003113202331230-2203201010200220-2222020320210320-0313121332011110-1322321122100133-1123002123031312-2030001320333103-0020210023332311): complete subsection reference.

<a id="canonical-1030232232102301-1102213323330233-0222102121313200-1232303001120313-3111332303031103-3013213211213313-0222103102130202-3312100010101130"></a>

<a id="canonical-3113010133313211-0300210103230102-1033133323011221-1312113302322201-1130302311112203-1320320031201230-3232313310202113-3001011110023123"></a>

## ip_address property — static_routes / 011233020103 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1331022231123221-0100003022103013-1133031221030300-1321233011021010-3130112021120021-2333322112002300-3223232130311300-1120101131112000"></a>

<a id="canonical-1323121231233010-2301213212013021-1002303120023313-2300000302232030-0110312213332022-3213300303320121-1122030130313130-0020110203222011"></a>

## ip_prefixes property — static_routes / 011233020103 / 6

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

- [node_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3121131131312033-3102100012000023-3020022122121233-2220301022320012-2131202021030211-1023002230213012-1210100032202030-3212320101223111): complete subsection reference.

<a id="canonical-0222010102111211-3201213230302223-2223130333332003-0330003100320310-2111321310132321-2101122112222311-0200011131200012-3212000303203110"></a>

## Next pages — static_routes / 011233020103 / 7

- [segment_vrf.segment_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3003113202331230-2203201010200220-2222020320210320-0313121332011110-1322321122100133-1123002123031312-2030001320333103-0020210023332311)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3121131131312033-3102100012000023-3020022122121233-2220301022320012-2131202021030211-1023002230213012-1210100032202030-3212320101223111)
- [segment_vrf.segment_config.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3013013210212323-2130222123331122-3102220321233132-3211131233310133-1323001231230321-0311212023230121-1333121113123323-1203333212113331)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3003113202331230-2203201010200220-2222020320210320-0313121332011110-1322321122100133-1123002123031312-2030001320333103-0020210023332311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023033121212003-2221232230002113-2002020212212103-1131111302200312-1323110100210301-3211010201000123-0121012310021031-0133000130333121"></a>

## segment_vrf.segment_config.static_routes.static_routes.default_gateway — default_gateway / 302220300330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- [segment_vrf.segment_config.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3013013210212323-2130222123331122-3102220321233132-3211131233310133-1323001231230321-0311212023230121-1333121113123323-1203333212113331)
- [segment_vrf.segment_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0310112223023333-0331032222330012-0023123003302222-3100030000102221-3333323013021300-3203311323301121-3323203213212332-0332221210020010)
- segment_vrf.segment_config.static_routes.static_routes.default_gateway

<a id="canonical-0233001213011113-3331323102232230-3110320021031233-2003202231000030-3211312211002030-0131313033322000-3121102002221130-3133232012212320"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

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

<a id="canonical-1102033111130132-3200002312010303-3031133212222221-1322213222111232-0202303303130123-3201031010002212-3012023100232101-3133023031031211"></a>

## Direct properties — default_gateway / 302220300330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123222200333023-0013110212011023-0233102013232103-2233012131002213-0130013300000311-1113132212001332-0122121032020020-0003213210011111"></a>

## Next pages — default_gateway / 302220300330 / 4

- [segment_vrf.segment_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0310112223023333-0331032222330012-0023123003302222-3100030000102221-3333323013021300-3203311323301121-3323203213212332-0332221210020010)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3121131131312033-3102100012000023-3020022122121233-2220301022320012-2131202021030211-1023002230213012-1210100032202030-3212320101223111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330113300231232-3321003321132221-2132311121323211-0122202233202122-1313331100321333-0001320103331310-3101000030303111-0200031101100312"></a>

## segment_vrf.segment_config.static_routes.static_routes.node_interface — node_interface / 122030301111 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- [segment_vrf.segment_config.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3013013210212323-2130222123331122-3102220321233132-3211131233310133-1323001231230321-0311212023230121-1333121113123323-1203333212113331)
- [segment_vrf.segment_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0310112223023333-0331032222330012-0023123003302222-3100030000102221-3333323013021300-3203311323301121-3323203213212332-0332221210020010)
- segment_vrf.segment_config.static_routes.static_routes.node_interface

<a id="canonical-3012330320302021-1000033203012303-1312320020300201-0120122121001221-2022301002200300-3020322303021300-2213232100112020-3102113020202222"></a>

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

<a id="canonical-1231030223122231-2111021231310232-1132200123030200-2222113003211201-3323012110022102-1321322201221030-2310001323020113-3232202203200000"></a>

## Direct properties — node_interface / 122030301111 / 3

- [list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2113202010113230-3032001331112030-3120300302011222-3020201303330131-0212311302213032-0113001101100232-0322210002022111-3322002001103121): complete subsection reference.

<a id="canonical-0320001311233322-3230101101003201-1220311033110323-2221100100313102-2313223000202301-3031330223203212-1212121030321233-1000002312320021"></a>

## Next pages — node_interface / 122030301111 / 4

- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2113202010113230-3032001331112030-3120300302011222-3020201303330131-0212311302213032-0113001101100232-0322210002022111-3322002001103121)
- [segment_vrf.segment_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0310112223023333-0331032222330012-0023123003302222-3100030000102221-3333323013021300-3203311323301121-3323203213212332-0332221210020010)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2113202010113230-3032001331112030-3120300302011222-3020201303330131-0212311302213032-0113001101100232-0322210002022111-3322002001103121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221323312011200-1222022201232200-3030230122031001-0033012210310201-3012233013110201-1033331000032122-1110220033011132-2302112222221121"></a>

## segment_vrf.segment_config.static_routes.static_routes.node_interface.list — list / 002112102013 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- [segment_vrf.segment_config.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3013013210212323-2130222123331122-3102220321233132-3211131233310133-1323001231230321-0311212023230121-1333121113123323-1203333212113331)
- [segment_vrf.segment_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0310112223023333-0331032222330012-0023123003302222-3100030000102221-3333323013021300-3203311323301121-3323203213212332-0332221210020010)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3121131131312033-3102100012000023-3020022122121233-2220301022320012-2131202021030211-1023002230213012-1210100032202030-3212320101223111)
- segment_vrf.segment_config.static_routes.static_routes.node_interface.list

<a id="canonical-2033112010303020-3000221210120012-1122202001113233-2123020320122300-1302310301031210-1231321232023013-3212200100103003-2332221130022102"></a>

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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-2132300100332110-1010212120002102-0023322210131021-1002133322320221-3020010022103332-1210120301112210-2111232100023013-2313331002310302"></a>

## Direct properties — list / 002112102013 / 3

- [interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2213211200210120-2102100323223003-2201333130313213-1321122021103201-1331023233110213-3320132132021333-1113122121230301-0301021323200313): complete subsection reference.

<a id="canonical-1132131022020222-0112231312223001-1200002122122000-3033301003000213-1100311001120021-0021213100310300-0213321123323020-3331211123011330"></a>

<a id="canonical-1022131302233110-3112201122322113-1022023111023303-0212303310213302-3101013100133020-2201220120323310-2010031220303110-2332313021033302"></a>

## node property — list / 002112102013 / 4

Type: `"string"`. Computed.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-2330003222103012-3201213211232321-2013221123101320-1020122302333302-0201131323312231-0130201222100332-1112300200322233-3031020111222203"></a>

## Next pages — list / 002112102013 / 5

- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2213211200210120-2102100323223003-2201333130313213-1321122021103201-1331023233110213-3320132132021333-1113122121230301-0301021323200313)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3121131131312033-3102100012000023-3020022122121233-2220301022320012-2131202021030211-1023002230213012-1210100032202030-3212320101223111)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2213211200210120-2102100323223003-2201333130313213-1321122021103201-1331023233110213-3320132132021333-1113122121230301-0301021323200313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230303231021131-0130121201022213-1313211023212110-3002023232112311-0012131323220033-3203333222211002-1021330003210202-2001303230330212"></a>

## segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface — interface / 213032222130 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- [segment_vrf.segment_config.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3013013210212323-2130222123331122-3102220321233132-3211131233310133-1323001231230321-0311212023230121-1333121113123323-1203333212113331)
- [segment_vrf.segment_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0310112223023333-0331032222330012-0023123003302222-3100030000102221-3333323013021300-3203311323301121-3323203213212332-0332221210020010)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3121131131312033-3102100012000023-3020022122121233-2220301022320012-2131202021030211-1023002230213012-1210100032202030-3212320101223111)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2113202010113230-3032001331112030-3120300302011222-3020201303330131-0212311302213032-0113001101100232-0322210002022111-3322002001103121)
- segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-0003100202123023-3311220311031310-0131321030120311-1021132323222121-1311232103103003-2022121131201300-1220003321303233-3001312022232110"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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

<a id="canonical-0203133321012002-2023122211012332-2003321233013300-1011112130002223-2202301301323012-2221020112303301-3121231131213122-1311232221121001"></a>

## Direct properties — interface / 213032222130 / 3

<a id="canonical-2001332021013223-1030311103320102-2123320032031212-2221111112313312-2321002320011113-2102101103323113-0202032121121131-1312032100100201"></a>

<a id="canonical-1203232130031013-1111102311233013-1112032322111233-3102133330230011-3203033101300323-0310003312213230-3222302131231322-1100020303202001"></a>

## kind property — interface / 213032222130 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-0112323123332001-0100322202013220-2320030030221332-0300211132131120-3012303122023302-0312211301310130-0312213230132211-0000130122321320"></a>

<a id="canonical-3020221100201130-3023332111011202-2330101223100323-2202133222011220-2113331112332030-1110223032333210-2202022332111313-2113321123113213"></a>

## name property — interface / 213032222130 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-0220001030122011-2123032003121330-2100123330111110-2220313300032323-2330211112013020-0023210132210023-0320113331000011-0320031110322102"></a>

<a id="canonical-2102223111331013-3130201231302322-0203022332230021-1223320103000100-3013032000100331-2321303122022100-1202020112100322-0323302301020311"></a>

## namespace property — interface / 213032222130 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3000331210120211-1200302032312003-0321013000310130-1111012212201121-3202323130103310-1311123000022303-1112230320212020-2011321132103012"></a>

<a id="canonical-2233233013023211-0223212222121302-3012300002021221-3301213232020022-3310133210303021-2121000321320002-0222222202211013-2212103032203110"></a>

## tenant property — interface / 213032222130 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3301013221312102-1212020333211131-0003001230210101-0113112213120020-2230002223100133-2233111123111130-2213010233122002-2313221133220311"></a>

<a id="canonical-2032322213001031-1330312330023210-0320330220211113-2230010131110032-3001300321310133-1211310030212231-2133033032020032-1003212202101301"></a>

## uid property — interface / 213032222130 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-2113313233203301-3311333112133133-2201032122103222-1001110231313132-3222200133220320-1210021233013122-0001323121211210-3300013230101110"></a>

## Next pages — interface / 213032222130 / 9

- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2113202010113230-3032001331112030-3120300302011222-3020201303330131-0212311302213032-0113001101100232-0322210002022111-3322002001103121)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1303120210133032-0020301133000323-3311321220022302-3113222021233210-2213313002331302-1100333233332222-0121333220032113-2133332330221312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212101323102221-1223302012131031-0231022222213210-3001231301331022-1103113222003122-0233200121111112-1223301301233323-1233120211312331"></a>

## segment_vrf.segment_config.static_v6_routes — static_v6_routes / 023032303303 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- segment_vrf.segment_config.static_v6_routes

<a id="canonical-2321210112013323-0200121000100001-2030323320102231-3010023203322200-3233011121013132-2002012001102123-1132003322330131-3322303222233230"></a>

Type: `"single"`. Computed.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

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

<a id="canonical-3212103030031000-0000320221331202-0122000222211223-2303313332120230-0001222033312121-2003003230302303-3232010313311121-2301121031132232"></a>

## Direct properties — static_v6_routes / 023032303303 / 3

- [static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1013123331120122-0033112130321011-1333300111133331-1321221010013020-1000023032023001-2330032010300003-1210003001133331-2321100233112220): complete subsection reference.

<a id="canonical-2102332210220231-0012022101201011-3113201322110110-3013030313321203-2233110103213102-0020133332120112-0321320100221031-2100010321202220"></a>

## Next pages — static_v6_routes / 023032303303 / 4

- [segment_vrf.segment_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1013123331120122-0033112130321011-1333300111133331-1321221010013020-1000023032023001-2330032010300003-1210003001133331-2321100233112220)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1013123331120122-0033112130321011-1333300111133331-1321221010013020-1000023032023001-2330032010300003-1210003001133331-2321100233112220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303011001133310-0213301223101122-3303110122133321-3332131320330323-2103000100330113-1111020233213333-0311233311013303-1203201021012231"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes — static_routes / 122001113001 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- [segment_vrf.segment_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1303120210133032-0020301133000323-3311321220022302-3113222021233210-2213313002331302-1100333233332222-0121333220032113-2133332330221312)
- segment_vrf.segment_config.static_v6_routes.static_routes

<a id="canonical-3020102102020121-1120013321312220-2012333223312020-2232302223203102-2222312021110303-0012032300123022-0220010231323133-0120230233023001"></a>

Type: `"list"`. Computed.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2020333102301002-1023301321220223-3233210001120211-0032331221322301-0213303222312322-3121010033201203-1001131203202012-0221110233000221"></a>

## Direct properties — static_routes / 122001113001 / 3

<a id="canonical-3230231020310322-2100021111113131-0332113011323221-1001322301211221-2200220103122333-2302013201032313-1201122231123113-2211012120332230"></a>

<a id="canonical-2232223212030133-3102300230010313-1200032031120112-3333101000100111-3000020212223121-1002222323011010-3323201331201332-3120021000100222"></a>

## attrs property — static_routes / 122001113001 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0021111133011010-0121301021321223-2020311330132131-0102103101030102-0103232233203020-3023020111011331-0123313200002020-0023100031323033): complete subsection reference.

<a id="canonical-0000323213212212-1201210000321322-1303300201112231-3311130203001130-3120230200023033-1013102330003113-1313110323321322-2213302321021333"></a>

<a id="canonical-1032230032111233-0200313220212001-3132112230203202-0303302001200210-3110031333312230-0031120111112113-3230023213323212-1130223001320130"></a>

## ip_address property — static_routes / 122001113001 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3203311223300213-3312303320221323-0302011130302222-0320112333313213-2331113232011222-2111030322023320-1203130321230013-2320102031221310"></a>

<a id="canonical-2003002211022122-0230020101220123-1302021200031103-2101332110300200-2002222220232113-1331120220131221-0111123210333111-2303031021333032"></a>

## ip_prefixes property — static_routes / 122001113001 / 6

Type: `["list", "string"]`. Computed.

List of IPv6 route prefixes that have common next hop and attributes.

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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1231220311131020-0222213111202120-3211201112032233-0310210332003311-1012332122003020-1212302300220032-3311102022222031-2332011132202103): complete subsection reference.

<a id="canonical-2302130103210033-3320022213120100-3110202231030100-3031111231322121-1020323003022230-2110020023231313-0001312212123301-3323123303211230"></a>

## Next pages — static_routes / 122001113001 / 7

- [segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0021111133011010-0121301021321223-2020311330132131-0102103101030102-0103232233203020-3023020111011331-0123313200002020-0023100031323033)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1231220311131020-0222213111202120-3211201112032233-0310210332003311-1012332122003020-1212302300220032-3311102022222031-2332011132202103)
- [segment_vrf.segment_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1303120210133032-0020301133000323-3311321220022302-3113222021233210-2213313002331302-1100333233332222-0121333220032113-2133332330221312)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0021111133011010-0121301021321223-2020311330132131-0102103101030102-0103232233203020-3023020111011331-0123313200002020-0023100031323033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131023111230032-3332333333123233-1121201330100323-2113122022201210-3123302031331110-0321103301310322-1333321002132333-2123210013102123"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway — default_gateway / 202311003330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- [segment_vrf.segment_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1303120210133032-0020301133000323-3311321220022302-3113222021233210-2213313002331302-1100333233332222-0121333220032113-2133332330221312)
- [segment_vrf.segment_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1013123331120122-0033112130321011-1333300111133331-1321221010013020-1000023032023001-2330032010300003-1210003001133331-2321100233112220)
- segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-1003321003320000-1330302021000032-0202233021220011-2312233333121001-1203020020211203-0012232130001112-0311020113213031-1320202101311122"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

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

<a id="canonical-2333111021323330-0102221112212001-3113210131001323-3200120201032021-1332332211213322-2223301131131113-0103323320012003-1210320110000330"></a>

## Direct properties — default_gateway / 202311003330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331122113131033-0311001013232332-0032112320203232-2101302133102323-3320100111112233-1001022320203020-0322030212211330-3331110220120110"></a>

## Next pages — default_gateway / 202311003330 / 4

- [segment_vrf.segment_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1013123331120122-0033112130321011-1333300111133331-1321221010013020-1000023032023001-2330032010300003-1210003001133331-2321100233112220)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1231220311131020-0222213111202120-3211201112032233-0310210332003311-1012332122003020-1212302300220032-3311102022222031-2332011132202103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102211123222302-2121313011120120-3111023230023230-1312220121110102-2210202022132123-1220201021132130-1320232310221332-0123121012321301"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes.node_interface — node_interface / 000000013210 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- [segment_vrf.segment_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1303120210133032-0020301133000323-3311321220022302-3113222021233210-2213313002331302-1100333233332222-0121333220032113-2133332330221312)
- [segment_vrf.segment_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1013123331120122-0033112130321011-1333300111133331-1321221010013020-1000023032023001-2330032010300003-1210003001133331-2321100233112220)
- segment_vrf.segment_config.static_v6_routes.static_routes.node_interface

<a id="canonical-3120311032131001-1033110212200312-3122000310233020-3232012321103221-3233101203112330-1201303120111023-2021021000022332-2132302212111330"></a>

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

<a id="canonical-0323301203203130-1131123023201120-0030231013330032-3231001023121132-2333023322201032-0133223102113300-0232322331232202-0321212211313323"></a>

## Direct properties — node_interface / 000000013210 / 3

- [list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2111011003301202-3133233121210301-3322233332303203-2223312211310330-0200330203023310-3221102013202221-3321212012023211-1012312333112021): complete subsection reference.

<a id="canonical-0221302013320310-2110213212301312-3201232220013020-2233131012323111-0312231033222233-0213123023021211-2222131032002102-3022023023131103"></a>

## Next pages — node_interface / 000000013210 / 4

- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2111011003301202-3133233121210301-3322233332303203-2223312211310330-0200330203023310-3221102013202221-3321212012023211-1012312333112021)
- [segment_vrf.segment_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1013123331120122-0033112130321011-1333300111133331-1321221010013020-1000023032023001-2330032010300003-1210003001133331-2321100233112220)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2111011003301202-3133233121210301-3322233332303203-2223312211310330-0200330203023310-3221102013202221-3321212012023211-1012312333112021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303033230322112-3023130210020222-1001002311300303-1330103222112231-1033200321301302-3023210002012023-1000121312231020-3300232132321102"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list — list / 000322303103 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- [segment_vrf.segment_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1303120210133032-0020301133000323-3311321220022302-3113222021233210-2213313002331302-1100333233332222-0121333220032113-2133332330221312)
- [segment_vrf.segment_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1013123331120122-0033112130321011-1333300111133331-1321221010013020-1000023032023001-2330032010300003-1210003001133331-2321100233112220)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1231220311131020-0222213111202120-3211201112032233-0310210332003311-1012332122003020-1212302300220032-3311102022222031-2332011132202103)
- segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-3231201210012330-2031112300020123-2111112020133311-1013130302121333-1232010021321210-1131330230213013-2001130301001013-3302223002021230"></a>

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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-0230212322001230-3331102300103113-2323232333133232-3112023113220320-1103210000000301-0010022332202310-3010021110132210-2133102000122010"></a>

## Direct properties — list / 000322303103 / 3

- [interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0023300312002232-3313011201013002-3012111310212102-2230130201211100-2213212022200312-2330002021113101-1332223300323200-3222322012002211): complete subsection reference.

<a id="canonical-1031213210131020-3310231323013220-1001221012000022-3202222123121113-1322030101330031-0332103111322000-2312222123323100-0322330030301110"></a>

<a id="canonical-3130032132203220-2311300020302131-1101111002132311-1200333122203320-1221330230100312-3200311332012110-1011023103323002-1301232131123321"></a>

## node property — list / 000322303103 / 4

Type: `"string"`. Computed.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-1321033202333020-3022310101121013-3111212230231121-1210112002032300-1312330232331001-0121011032323231-2001130333022311-1231332012223111"></a>

## Next pages — list / 000322303103 / 5

- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0023300312002232-3313011201013002-3012111310212102-2230130201211100-2213212022200312-2330002021113101-1332223300323200-3222322012002211)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1231220311131020-0222213111202120-3211201112032233-0310210332003311-1012332122003020-1212302300220032-3311102022222031-2332011132202103)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0023300312002232-3313011201013002-3012111310212102-2230130201211100-2213212022200312-2330002021113101-1332223300323200-3222322012002211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113223322213111-0223231022013230-0321203212211312-2323101311301101-3000031111203202-0111221001213200-3101032020112200-3203323131032001"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface — interface / 300022201101 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0012020300313222-3111301312332300-1312100300000223-0332113232310300-2130132120201331-3001012312210132-3203200130212031-2113221223310320)
- [segment_vrf.segment_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1303120210133032-0020301133000323-3311321220022302-3113222021233210-2213313002331302-1100333233332222-0121333220032113-2133332330221312)
- [segment_vrf.segment_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1013123331120122-0033112130321011-1333300111133331-1321221010013020-1000023032023001-2330032010300003-1210003001133331-2321100233112220)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1231220311131020-0222213111202120-3211201112032233-0310210332003311-1012332122003020-1212302300220032-3311102022222031-2332011132202103)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2111011003301202-3133233121210301-3322233332303203-2223312211310330-0200330203023310-3221102013202221-3321212012023211-1012312333112021)
- segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-3022302131021320-2201221321101200-1131213033032122-1010201012303013-1211330331021221-3120310001002023-3102302031231333-1002112301031010"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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

<a id="canonical-1223221323000233-3330033313221231-1302321132210330-2012133311113210-2020120212101133-2200223000222212-3332112220100033-2123320213100132"></a>

## Direct properties — interface / 300022201101 / 3

<a id="canonical-2332010323001233-0033212230320312-0012320032131232-1323201212012311-0132121312301131-3220120131300211-1122012212333312-1210222332333301"></a>

<a id="canonical-3332133213101003-1013223230112130-1100013231033213-0212133312203033-2031022322321232-3223300222232233-3320231021033010-1211010223013203"></a>

## kind property — interface / 300022201101 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-3030231103003111-3020312333311200-1313023122213222-3321313010212231-0312023333310013-2003132213321111-3312020001310303-0002011212111223"></a>

## name property — interface / 300022201101 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3130322311303223-1123121031233131-1110302102020303-2302011103133013-2323103321131202-2111103223222103-0111220132031320-0220231131330130"></a>

## namespace property — interface / 300022201101 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0322121013200102-1023220110002323-3221100000322300-1330031330322131-2131133010033213-1300212301032100-1223003111233112-1332213312301330"></a>

## tenant property — interface / 300022201101 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-0300330221323310-0211211131311330-2031233303203313-0223100012232001-1003220100013332-3131321101001321-2131320132300122-0012133020033303"></a>

## uid property — interface / 300022201101 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-3211013122313103-3230212322100132-2021322200201020-2032313023323221-2003032310212233-0313323123113032-2313023133111301-0123332121310331"></a>

## Next pages — interface / 300022201101 / 9

- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2111011003301202-3133233121210301-3322233332303203-2223312211310330-0200330203023310-3221102013202221-3321212012023211-1012312333112021)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0021113223132201-3022121331332032-3312123112022020-1203310012023130-2303101212300123-0001032222032213-1023013332311022-2231232033102112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311000111020320-3133101310110200-2111011220103001-3110111032303203-3300111320223303-3101033200312322-2333220233003221-0021200223232102"></a>

## segment_vrf.segment_network — segment_network / 222313011113 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- segment_vrf.segment_network

<a id="canonical-0300320221222203-1033020102013121-2021120223120123-1123013320023122-2132101033100012-1222122212100313-0131233032002321-2213322222031102"></a>

Type: `"single"`. Computed.

Type establishes a 'direct reference' from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name for public API and Uid for private API This type of
reference is called direct because the relation is explicit and concrete (as opposed to selector..

Upstream description:

This type establishes a 'direct reference' from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name for public API and Uid for private API This
type of reference is called direct because the relation is explicit and concrete (as opposed to
selector reference which builds a group based on labels of selectee objects)

<a id="canonical-2202003022211221-2213322021300002-2100310232313011-2001223011210012-3311132302102023-0321003230113302-2111300311300020-0100302210321203"></a>

## Direct properties — segment_network / 222313011113 / 3

<a id="canonical-0113133220101213-2013133323020021-3232131101311221-3020010011232223-0101130333232211-3203122003123330-3012320320011301-2122110033331312"></a>

<a id="canonical-1020320020223100-2022321333111210-1133131311123020-0200221132031311-0032200033100003-1212230011001031-1103100311121231-1030221001300313"></a>

## kind property — segment_network / 222313011113 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-0122222332120123-0332233130220123-0231332032323213-1313301132003230-1232231113131200-2202311001013131-2010131232001320-1103210231312322"></a>

## name property — segment_network / 222313011113 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-1013110112303230-0023111232233332-0301311030012221-2111013121131332-2332130302132031-0031332003212033-3121200100231102-2203321202303010"></a>

## namespace property — segment_network / 222313011113 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0212220333211023-2302313021113110-2103020131230121-3123101330301132-2102022313211022-0310132301022101-1210102231232202-2213302020221212"></a>

## tenant property — segment_network / 222313011113 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3310200012110011-1233102300032301-1203032033300032-1011331221200123-3231101103101322-0223122002213230-2222230021112120-0132203123320032"></a>

## uid property — segment_network / 222313011113 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-1032321003203233-2212212220013233-0320333201123201-1313030301233312-0031303100331001-3123231201012013-2103112001030121-1032221132101303"></a>

## Next pages — segment_network / 222313011113 / 9

- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000033030322131-1123330122213033-2003103303102222-0200202303032031-1112122110013331-2321301132332200-1102311311201223-3322201021001233"></a>

## site_mesh_group_on_slo — site_mesh_group_on_slo / 230033110201 / 2

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

<a id="canonical-2310220201101023-0122233021300212-0301131203310113-2301302330322132-2212311123213220-2322230331332133-0213302103102003-1102111321132233"></a>

## Direct properties — site_mesh_group_on_slo / 230033110201 / 3

- [no_site_mesh_group](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1331023023112310-1033013130131111-3120331103012113-0132333331321323-2320011313200210-0232023003301123-1202131331013033-3023122020220302): complete subsection reference.

- [site_mesh_group](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0310123133313310-3101121232031020-1302233301002121-1311111131332031-0220212001000223-0130032130110330-3121332202001300-1311033113113121): complete subsection reference.

- [sm_connection_public_ip](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0021212330003021-0231020110233323-1113133321030020-3002133131130110-3022123222213311-1302223331212110-1302131120032033-1302010200230222): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1222032311200220-2311221211201222-0221011112203212-3300123011013123-1123122011112211-2032010011311300-0033332111023213-2002023320101000): complete subsection reference.

<a id="canonical-0210300230311321-2021231311332321-1000233332021031-2301002112330121-3110021223323303-3210231323031211-3011202233022103-3311211022003111"></a>

## Next pages — site_mesh_group_on_slo / 230033110201 / 4

- [site_mesh_group_on_slo.no_site_mesh_group](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1331023023112310-1033013130131111-3120331103012113-0132333331321323-2320011313200210-0232023003301123-1202131331013033-3023122020220302)
- [site_mesh_group_on_slo.site_mesh_group](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0310123133313310-3101121232031020-1302233301002121-1311111131332031-0220212001000223-0130032130110330-3121332202001300-1311033113113121)
- [site_mesh_group_on_slo.sm_connection_public_ip](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0021212330003021-0231020110233323-1113133321030020-3002133131130110-3022123222213311-1302223331212110-1302131120032033-1302010200230222)
- [site_mesh_group_on_slo.sm_connection_pvt_ip](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1222032311200220-2311221211201222-0221011112203212-3300123011013123-1123122011112211-2032010011311300-0033332111023213-2002023320101000)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1331023023112310-1033013130131111-3120331103012113-0132333331321323-2320011313200210-0232023003301123-1202131331013033-3023122020220302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202330100103032-0332210111020001-3331100222330333-2203213131031020-1333333010312202-1123013200111321-2033303113221313-1111022000211311"></a>

## site_mesh_group_on_slo.no_site_mesh_group — no_site_mesh_group / 323212000032 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031)
- site_mesh_group_on_slo.no_site_mesh_group

<a id="canonical-3323330023302300-2312013002310001-0132301322100012-0100033310122203-2332310220213010-0302000120331313-2100003211301101-2331223200220003"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-0031222310230230-3321131201210002-0321103113322010-2331031023013131-3100302120023021-2033212213202122-0013230201102032-3101030112113013"></a>

## Direct properties — no_site_mesh_group / 323212000032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001211113002030-2102212131023112-0003121310200112-0120122300113031-3231013303221111-2013223232020132-1230020013230132-1023321111101312"></a>

## Next pages — no_site_mesh_group / 323212000032 / 4

- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0310123133313310-3101121232031020-1302233301002121-1311111131332031-0220212001000223-0130032130110330-3121332202001300-1311033113113121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330013102030133-2312110332133211-2231103013100123-0223321333000301-0323022120222112-1010111023231031-3232110332023111-3301033022011133"></a>

## site_mesh_group_on_slo.site_mesh_group — site_mesh_group / 303223222022 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031)
- site_mesh_group_on_slo.site_mesh_group

<a id="canonical-1321102112031133-3302103200313300-2210220030221232-0033030310120132-1330331003032320-3221120333013130-3311031023133032-2220022223220233"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-3300323322022232-2303131130021031-0122120302020321-3010012033130133-1313011333133100-2121330303033220-3213111120133132-3211012221233011"></a>

## Direct properties — site_mesh_group / 303223222022 / 3

<a id="canonical-0211002112203302-3300021232020332-3310031323000002-3201122232221132-0233231321220110-0033232022300032-3130321231220133-2133013103112022"></a>

<a id="canonical-1203133201100332-2033230112331302-0020213333012211-3103201131333323-3221000213233020-1111312213323333-0132123002100002-2222221230030332"></a>

## name property — site_mesh_group / 303223222022 / 4

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

<a id="canonical-1203232303322120-2001002313121323-2300132002131020-3301333332122032-2133230312103110-2320012201010322-0323033230203123-3333221022200232"></a>

## namespace property — site_mesh_group / 303223222022 / 5

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

<a id="canonical-2212201220033001-1301022010323233-1110323211101002-2010030331302003-2232332230212033-2300010302033133-2111232313030333-1202101020101303"></a>

## tenant property — site_mesh_group / 303223222022 / 6

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

<a id="canonical-2232300300000110-1032100210302233-1122300021213021-1012310112212120-0100010331213013-1122010112201102-0103122302021012-1131311100113131"></a>

## Next pages — site_mesh_group / 303223222022 / 7

- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0021212330003021-0231020110233323-1113133321030020-3002133131130110-3022123222213311-1302223331212110-1302131120032033-1302010200230222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030021233021300-3320133023320112-2032223323121213-2122233013310002-1011322130131010-0331211000101132-2311321333130021-2310222122332333"></a>

## site_mesh_group_on_slo.sm_connection_public_ip — sm_connection_public_ip / 223303103002 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031)
- site_mesh_group_on_slo.sm_connection_public_ip

<a id="canonical-0111322032223002-1110313232312231-1210230113110231-0100123311113302-0032122111030213-3020201012133322-2113200321010330-0302113011100121"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-3021222222302300-3322123233032233-0023322322111231-0300023010031030-0303101323231011-3303203330021222-2003131211311133-2121330201011231"></a>

## Direct properties — sm_connection_public_ip / 223303103002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000221011300003-0022000312132311-2310023201022221-2112102112131120-2112103331210021-3110123000123000-3200133333111221-2322123103202211"></a>

## Next pages — sm_connection_public_ip / 223303103002 / 4

- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1222032311200220-2311221211201222-0221011112203212-3300123011013123-1123122011112211-2032010011311300-0033332111023213-2002023320101000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111131033130220-0333122120321110-3003232001103012-1111121011223103-2310030122021220-2023022320102003-3302321232001000-3220302320221230"></a>

## site_mesh_group_on_slo.sm_connection_pvt_ip — sm_connection_pvt_ip / 332111230203 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031)
- site_mesh_group_on_slo.sm_connection_pvt_ip

<a id="canonical-0321003000331112-0200312122110200-1033010120202310-0032303120310330-0211012333331322-3132003010010201-3330123131313120-2021102030132323"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-0030200210211200-3310001230122132-2022310202221313-1010031012222001-1222123322333033-0102231233303323-3001102220113322-3110133110320100"></a>

## Direct properties — sm_connection_pvt_ip / 332111230203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032200133123001-1112211120130100-3111211113123120-1001123301313202-2133302323301131-2213333030302312-2300101303012211-0211313201223201"></a>

## Next pages — sm_connection_pvt_ip / 332111230203 / 4

- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0030100133010212-1200030111021312-3123032332031021-0010103012023120-2202203012321113-0113233110320102-2113301111200001-0103012002023131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011320210120000-0321202330123003-2012323301232312-3030023230033332-2230102100300321-0211002210023223-3230302301310130-2120032010202120"></a>

## upgrade_settings — upgrade_settings / 331002213020 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- upgrade_settings

<a id="canonical-3000003010233110-3122233323302221-3330021003012222-2202323100231222-2331122302332101-0210223033131103-2012230122031022-3013001013111221"></a>

Type: `"single"`. Computed.

Configuration parameter for upgrade settings.

Upstream description:

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

<a id="canonical-0023120231301103-3100230323331222-1122012110103300-2200301112311012-1031202131322010-0231333320103012-0302031112233031-0101201002030131"></a>

## Direct properties — upgrade_settings / 331002213020 / 3

- [kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1311312312123132-2132122001121023-2310232230000031-3331033300123111-2112333121031222-2121330032031221-0220303303333022-3300013222010111): complete subsection reference.

<a id="canonical-1033211013011223-2111122131232223-3001333130222023-0120222301133331-2013101203133131-1302033021220320-3131102110013333-2001103203010310"></a>

## Next pages — upgrade_settings / 331002213020 / 4

- [upgrade_settings.kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1311312312123132-2132122001121023-2310232230000031-3331033300123111-2112333121031222-2121330032031221-0220303303333022-3300013222010111)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1311312312123132-2132122001121023-2310232230000031-3331033300123111-2112333121031222-2121330032031221-0220303303333022-3300013222010111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331000032132302-0232223100331202-3312303313212121-2031300313123333-2321133133320021-2200111003113201-2212212330210010-0330112013322221"></a>

## upgrade_settings.kubernetes_upgrade_drain — kubernetes_upgrade_drain / 020133003023 / 2

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

<a id="canonical-0213113330202011-0100300202130132-1112031210133021-0232312112320203-3220130201231011-1312103003133231-3200322311311031-0233012233003010"></a>

## Direct properties — kubernetes_upgrade_drain / 020133003023 / 3

- [disable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1101002301102122-2031301321121333-2111011110222102-2010003033323032-0102303100133321-0020123323032330-0333110320203310-0011221132211102): complete subsection reference.

- [enable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0021210223223221-1202110301313103-0001033232221120-3133022310111011-1230301200223222-3001001331031032-3021203121012132-2203311323331123): complete subsection reference.

<a id="canonical-3210313230011002-2021102301221133-2230210201230132-1022202010301301-2100320221201212-0012231222000232-0032110231110010-3013031323021223"></a>

## Next pages — kubernetes_upgrade_drain / 020133003023 / 4

- [upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1101002301102122-2031301321121333-2111011110222102-2010003033323032-0102303100133321-0020123323032330-0333110320203310-0011221132211102)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0021210223223221-1202110301313103-0001033232221120-3133022310111011-1230301200223222-3001001331031032-3021203121012132-2203311323331123)
- [upgrade_settings](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0030100133010212-1200030111021312-3123032332031021-0010103012023120-2202203012321113-0113233110320102-2113301111200001-0103012002023131)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1101002301102122-2031301321121333-2111011110222102-2010003033323032-0102303100133321-0020123323032330-0333110320203310-0011221132211102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031331220221012-1031023023310322-1233213213120200-3210222222003003-3211221122100122-1303233023320212-2133303103232232-3232020330222132"></a>

## upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain — disable_upgrade_drain / 211231023223 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [upgrade_settings](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0030100133010212-1200030111021312-3123032332031021-0010103012023120-2202203012321113-0113233110320102-2113301111200001-0103012002023131)
- [upgrade_settings.kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1311312312123132-2132122001121023-2310232230000031-3331033300123111-2112333121031222-2121330032031221-0220303303333022-3300013222010111)
- upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-3010223302103231-0121212223111333-0033200230113113-2312000210203033-1013031101231311-0300020320223111-0221320200233103-0233111000310333"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable upgrade drain.

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

<a id="canonical-3300122330101030-3210220300001322-2220310333213301-0200212321331100-2312231002232033-3222303121102222-1310023032103111-1303020012201312"></a>

## Direct properties — disable_upgrade_drain / 211231023223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200302202013202-0101130222232201-3123213231301003-2321103112023120-3030311300301311-2020233321220332-2100003333103002-0232210203021000"></a>

## Next pages — disable_upgrade_drain / 211231023223 / 4

- [upgrade_settings.kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1311312312123132-2132122001121023-2310232230000031-3331033300123111-2112333121031222-2121330032031221-0220303303333022-3300013222010111)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0021210223223221-1202110301313103-0001033232221120-3133022310111011-1230301200223222-3001001331031032-3021203121012132-2203311323331123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001002323211133-2010100021023302-2202232201133203-0323112112110323-0310102010033120-1320321131100003-3131100221223230-3112030232010313"></a>

## upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain — enable_upgrade_drain / 300231331331 / 2

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

<a id="canonical-3031012333310031-0312132030203002-0322313203232000-0012233210100122-3111232022200310-3112321303211302-0001313101203132-1301122323121201"></a>

## Direct properties — enable_upgrade_drain / 300231331331 / 3

- [disable_vega_upgrade_mode](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1000023132331110-0302231221120301-0201322113131332-0032111123321231-0233302023230311-0121012200211000-0201320301232210-2200300300312222): complete subsection reference.

<a id="canonical-0311322102003321-2301212210200030-0210212211333222-0323000203303023-1012321021101233-0300101132033131-1100013022323313-0310230211031113"></a>

<a id="canonical-2203220230121202-1302312011033313-2110230333101032-0211312300201100-0310223010312230-2111320302010312-0313213210111032-1033203310200211"></a>

## drain_max_unavailable_node_count property — enable_upgrade_drain / 300231331331 / 4

Type: `"number"`. Computed.

Node Batch Size Count. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

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

<a id="canonical-3301323103333113-1333331111212121-1020011313023122-1012022132232003-0022103010103231-3210111022013312-0031022102103301-0330301112131030"></a>

## drain_max_unavailable_node_percentage property — enable_upgrade_drain / 300231331331 / 5

Type: `"number"`. Computed.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-1303010013002002-2321012312121223-2211203313030312-2123020130012330-1300021033102121-1201312132110022-0020332333123221-1121122321311301"></a>

<a id="canonical-1323103003212232-0013122223010021-2313203200030111-2100310133333100-0212301013032301-3302321002231212-3220111010201021-0010121032100320"></a>

## drain_node_timeout property — enable_upgrade_drain / 300231331331 / 6

Type: `"number"`. Computed.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It
is..

Upstream description:

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

<a id="canonical-3101311012122313-2310122132111323-2322232211212002-2311011231122310-1310303213323303-2300013300333232-3221000213010333-0001002221021210"></a>

## Next pages — enable_upgrade_drain / 300231331331 / 7

- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1000023132331110-0302231221120301-0201322113131332-0032111123321231-0233302023230311-0121012200211000-0201320301232210-2200300300312222)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2300333333220033-2221032002323131-1300130033000020-1312322233132110-1321210332101211-3032200330111121-0101011011010333-0132101013322210)
- [upgrade_settings.kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1311312312123132-2132122001121023-2310232230000031-3331033300123111-2112333121031222-2121330032031221-0220303303333022-3300013222010111)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1000023132331110-0302231221120301-0201322113131332-0032111123321231-0233302023230311-0121012200211000-0201320301232210-2200300300312222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131020233102202-0330223021102111-2301320202033032-3133111233120300-3222321100220211-2310301030313231-0211032133111031-2301332111222131"></a>

## upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — disable_vega_upgrade_mode / 201321121131 / 2

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

<a id="canonical-2331311211300122-1132312330022203-2312313233133123-0123003213111023-2200021203121123-2003230033331333-2101000131302223-1220332322121030"></a>

## Direct properties — disable_vega_upgrade_mode / 201321121131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113232203103200-0102121301102012-0313130033120113-1332030310233122-2211033130331330-1012213131011233-0020300120233223-2320232213033210"></a>

## Next pages — disable_vega_upgrade_mode / 201321121131 / 4

- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0021210223223221-1202110301313103-0001033232221120-3133022310111011-1230301200223222-3001001331031032-3021203121012132-2203311323331123)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2300333333220033-2221032002323131-1300130033000020-1312322233132110-1321210332101211-3032200330111121-0101011011010333-0132101013322210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212313230321322-1001101123032012-1013330002131030-3231000202021110-0323122032113323-1110133012330120-2110303321303120-0230323331322222"></a>

## upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — enable_vega_upgrade_mode / 231202133030 / 2

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

<a id="canonical-0220312131130322-3010203121203311-1312031333231211-2023301113222300-3230020201020230-3203213331233220-3301213112221221-3231303322001330"></a>

## Direct properties — enable_vega_upgrade_mode / 231202133030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211233022221102-3220220113322133-0011332220321030-1020330022312202-3313223032212313-0320121102200102-0023232310223023-3311300123321103"></a>

## Next pages — enable_vega_upgrade_mode / 231202133030 / 4

- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0021210223223221-1202110301313103-0001033232221120-3133022310111011-1230301200223222-3001001331031032-3021203121012132-2203311323331123)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223003300002323-0122220231023121-2001311103221300-1302203011122312-2210122230211231-3023112121110301-3322101322021313-0203111022133332"></a>

## vmware — vmware / 023203332131 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- vmware

<a id="canonical-1000002233221021-3103002001232013-0203230211301022-0002322012112022-0103232111201010-2202332322301210-2103102202330030-0012103321010332"></a>

Type: `"single"`. Computed.

VMware Provider Type. VMware Provider Type.

Upstream description:

VMware Provider Type.

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

<a id="canonical-0133201002330312-3102312203232312-0302232120131022-3100021101022003-0310122201020223-2022120120323131-1221121011322111-0031031131212003"></a>

## Direct properties — vmware / 023203332131 / 3

- [not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212): complete subsection reference.

<a id="canonical-3211123130220013-2301121303021300-3323322010302331-0102013033112300-0220223311130113-3310210330211123-2032103222321133-1032210211111220"></a>

## Next pages — vmware / 023203332131 / 4

- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223201211333321-1303031113222032-3113232320303000-3222002222131001-1021012330112330-2120333132210333-0211323111001020-1311022120330210"></a>

## vmware.not_managed — not_managed / 013113220213 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- vmware.not_managed

<a id="canonical-1213021313222310-3121323213021023-3100100011030113-2203123332213102-3002212031223001-0332033323310330-2031023000300211-3233020032202103"></a>

Type: `"single"`. Computed.

Section will show nodes associated with this site.

Upstream description:

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

<a id="canonical-0021320310210333-1011032213330031-0121303310203213-0003112020203100-3322032333122310-3110332110131310-3333333023031001-0203221100333221"></a>

## Direct properties — not_managed / 013113220213 / 3

- [node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101): complete subsection reference.

<a id="canonical-3232033313010123-0133211110332300-2121101023110032-2031001223000101-2220032133120121-0331313311330230-3010123100313001-3012101321002113"></a>

## Next pages — not_managed / 013113220213 / 4

- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112222222102022-2332330131012312-3202200231002212-0000312113301003-0011230101210013-0020301000030032-0112020001002111-2230013131320122"></a>

## vmware.not_managed.node_list — node_list / 221302200230 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- vmware.not_managed.node_list

<a id="canonical-3323211213321023-0300012331322202-3113132330121223-0202231132030323-3110203121001203-0030000122201002-3131331231202022-0322131213303333"></a>

Type: `"list"`. Computed.

Section will show nodes associated with this site.

Upstream description:

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

<a id="canonical-2231332332323222-1121231032210003-0311131131101230-0003103110122212-1203332310222123-2102102323023221-2311303013202233-0331203200322220"></a>

## Direct properties — node_list / 221302200230 / 3

<a id="canonical-1330320300132102-3003201133213233-3033200223322102-0210033321200302-1301001323031110-3200020101022131-1122200023320130-2220120213330011"></a>

<a id="canonical-2322202313322310-2333202222112131-3331101111111032-0133022221121101-0210302323120230-0233223333223321-0231103002212212-0121320311121013"></a>

## hostname property — node_list / 221302200230 / 4

Type: `"string"`. Computed.

Hostname. Hostname for this Node.

Upstream description:

Hostname for this Node.

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

<a id="canonical-3303320303000212-2013133212020101-2230013120201323-2310001220210331-3103312320103023-1120333200210300-1101323001103301-3203023103312031"></a>

## public_ip property — node_list / 221302200230 / 5

Type: `"string"`. Computed.

Public IP. Public IP for this Node.

Upstream description:

Public IP for this Node.

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

<a id="canonical-3011303320223012-1033131100100131-1203201030122101-2230232303111112-3130020030101211-1100003103120231-0030121321020311-3301030203002110"></a>

## type property — node_list / 221302200230 / 6

Type: `"string"`. Computed.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Upstream description:

Type for this Node, can be Control or Worker.

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

<a id="canonical-2102033202003011-1122001010232301-3122310102023122-1133031002032031-0011002212133113-2032021322320232-3013222233031103-0133131032201021"></a>

## Next pages — node_list / 221302200230 / 7

- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322212323120203-1030311000110100-0103121133133201-1300003011102123-2333220133033330-2300012231013032-0102221122200303-3122201300012332"></a>

## vmware.not_managed.node_list.interface_list — interface_list / 203220132002 / 2

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

<a id="canonical-1111032320203000-1030130210010223-2111000200311122-1302002012003300-0322232101101212-0010221121120202-1232111130322130-2013010020311022"></a>

## Direct properties — interface_list / 203220132002 / 3

- [bond_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3113233331323212-0201311132302221-0323311320033121-1221020221221020-3312022121220131-1310230333121213-3312122212120123-3133323202010303): complete subsection reference.

<a id="canonical-2113032123123233-2223013301330211-0311210331211323-0202023223300112-2301003313210031-2102321100101301-0203300331203330-2032120210300102"></a>

<a id="canonical-0200102013121121-1001033221332321-0130032233221120-0033113021200323-0021103311203232-0130132302201102-2120223123001321-2033233312311021"></a>

## description_spec property — interface_list / 203220132002 / 4

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2002000010203020-1313101303100132-3221022203001132-2310200303013001-2231130113003003-2000023321030001-0130012302212201-1131013323110103): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2131211010200313-3110301321130222-1102123200113032-2221200010222312-3311330101103223-1332332132131213-0022022112111000-0312023302203233): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322): complete subsection reference.

<a id="canonical-2121201231002132-1031200301000122-2000000202332123-1110132222130313-3020200103123120-2022213210133032-3332333132021302-3212033302231121"></a>

<a id="canonical-3220310002312222-0313133110212213-2000233331303201-0003310321110321-2120131122113102-3020210033201320-0111020012202210-2331223213031203"></a>

## is_management property — interface_list / 203220132002 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-2311010321323012-1013333123112010-3333211220231113-2311123003312212-0203100303101301-2311120321031123-1002001020331121-3003202100300110"></a>

<a id="canonical-0120121123333130-0313023330232200-1022301021303321-2332031213301313-0310330101310122-3232102221100032-2030132022001131-1302110030103310"></a>

## is_primary property — interface_list / 203220132002 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-1003201320200111-0312010010101233-0323101202003131-3210012010113032-0003223010320132-0221133330020033-0110331123112003-3221111201201201"></a>

<a id="canonical-0003012020113133-3133031020332101-2113122010101132-2202201001311232-2203210330030123-0112230202000330-2302301000003313-3322220220003221"></a>

## labels property — interface_list / 203220132002 / 7

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

<a id="canonical-3011311231211103-2232300320121311-1303310003001310-1110313201100130-1303033010131023-3322320222112321-2011230012031021-2223221022230023"></a>

## mtu property — interface_list / 203220132002 / 8

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

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

<a id="canonical-0300331122322331-1120232301112131-3023123120313302-2233302120022023-1121202320331213-1222331300020231-3003010323102003-0020012102303103"></a>

## name property — interface_list / 203220132002 / 9

Type: `"string"`. Computed.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

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

<a id="canonical-3001333130330311-2213000321212222-2012321223011130-0301230020210023-2212001312330020-1331230023003213-1333010002301203-2133012022232330"></a>

## priority property — interface_list / 203220132002 / 10

Type: `"number"`. Computed.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

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

<a id="canonical-1012031332330122-2210002333223303-2011030000121032-1123032200130100-0031010113110122-1130031331032312-2120313130010313-1121332331100020"></a>

## Next pages — interface_list / 203220132002 / 11

- [vmware.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3113233331323212-0201311132302221-0323311320033121-1221020221221020-3312022121220131-1310230333121213-3312122212120123-3133323202010303)
- [vmware.not_managed.node_list.interface_list.dhcp_client](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2002000010203020-1313101303100132-3221022203001132-2310200303013001-2231130113003003-2000023321030001-0130012302212201-1131013323110103)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130)
- [vmware.not_managed.node_list.interface_list.ethernet_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2131211010200313-3110301321130222-1102123200113032-2221200010222312-3311330101103223-1332332132131213-0022022112111000-0312023302203233)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322)
- [vmware.not_managed.node_list.interface_list.monitor](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3000311103023222-2111102210333112-0020030303122131-2213331012001330-2122023121013123-1302301120133213-1000011210233330-1002330303330200)
- [vmware.not_managed.node_list.interface_list.monitor_disabled](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3111220120002011-0130022213201132-0230331301200221-3221122101201330-2230002300332002-2030010032220330-2113133003302200-1121030221220201)
- [vmware.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3320130231323321-3022122100223033-1022132021022321-2011100212203002-0033031330120212-0212210003112230-2202320210023220-3322202133132320)
- [vmware.not_managed.node_list.interface_list.no_ipv4_address](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1311301011230121-2103012100303202-2231012122020031-2032112300013302-0112321213131123-3032030000103020-1100121311123230-2033303223003012)
- [vmware.not_managed.node_list.interface_list.no_ipv6_address](data-sources--securemesh_site_v2--reference--group-018.md#canonical-2221230121110112-0000022310030300-3331023302031310-2301013022330330-3222121321032112-3330122131020130-1322003213222330-3223020113112010)
- [vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-018.md#canonical-2121211122013120-3103203013202130-0310310221301133-3300123122110130-2321201321123323-0200313022223020-0323030000022030-0322311230123100)
- [vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-018.md#canonical-2030023132322101-2011112010030122-0311133322003030-1133300113311203-2103001023133200-0123113110313232-0323002030100222-0231102301022003)
- [vmware.not_managed.node_list.interface_list.static_ip](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0322013101212310-2222330010320203-0121011121312102-0210101001022231-1220313221113220-1031320022212303-0320003022130210-3333003000030110)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0101212013130321-0131302111333332-2312131200130323-3111003310322033-0033230100220132-3322223222230132-0011121000101030-0332100211010113)
- [vmware.not_managed.node_list.interface_list.vlan_interface](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0230030312120110-2020013303020100-2023331011232031-0302302103032030-2321322301231232-2223121203332121-2210203321013231-3032330030011313)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3113233331323212-0201311132302221-0323311320033121-1221020221221020-3312022121220131-1310230333121213-3312122212120123-3133323202010303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030110122010013-3010211212020201-0323200301111001-3210303023313231-1203311031131122-2133221031332110-3123110301130323-0120233120132311"></a>

## vmware.not_managed.node_list.interface_list.bond_interface — bond_interface / 222203212101 / 2

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

Upstream description:

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

<a id="canonical-0023013301233111-3023322223030221-3233211320301010-1003311131222303-2023110002122321-2302111203310232-1010011222123333-3002120131332122"></a>

## Direct properties — bond_interface / 222203212101 / 3

- [active_backup](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2031330002323032-2123111311232010-1120303302202311-3121022123032202-1003330232223031-2330103300122002-1222220222211303-0302010012132232): complete subsection reference.

<a id="canonical-2321300120021010-1223111133312223-1203103131320121-1233012300200121-0100203321200011-0123131212301001-1001021121231302-3302113232231131"></a>

<a id="canonical-1223020321132133-1131332011232223-2310010130332032-1033313110233120-1013110032002212-0221023201131312-1221331102300333-1223300111202111"></a>

## devices property — bond_interface / 222203212101 / 4

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

<a id="canonical-3031330201313120-0101332332130210-0332302223032102-0120010230211202-1321312112321133-1002323212002233-3123101130333230-0233100213011220"></a>

## link_polling_interval property — bond_interface / 222203212101 / 5

Type: `"number"`. Computed.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

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

<a id="canonical-0031020102321022-3301333301321022-3322222022101130-1022130212311122-3331022332213130-2102100132220132-1132323200331100-2213230000333223"></a>

## link_up_delay property — bond_interface / 222203212101 / 6

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

<a id="canonical-3300112010323113-0101330332100301-2010132301031020-2112303311131310-3102222011131123-2222121110113202-3211302133003202-3120302212002121"></a>

## name property — bond_interface / 222203212101 / 7

Type: `"string"`. Computed.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

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

<a id="canonical-0330020032010311-0220113023020333-1003033332200123-0220211133313232-2033101320022310-0030232110323131-0320333322003313-0022212301230110"></a>

## Next pages — bond_interface / 222203212101 / 8

- [vmware.not_managed.node_list.interface_list.bond_interface.active_backup](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2031330002323032-2123111311232010-1120303302202311-3121022123032202-1003330232223031-2330103300122002-1222220222211303-0302010012132232)
- [vmware.not_managed.node_list.interface_list.bond_interface.lacp](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2130120320303322-3321320202312321-2021131333120113-2332323013211201-0132230023220011-0232221021323110-1010011312210232-3012030023222333)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2031330002323032-2123111311232010-1120303302202311-3121022123032202-1003330232223031-2330103300122002-1222220222211303-0302010012132232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200131022310003-3122112200133231-3203332120130020-1222301301110201-3000310012013212-1210033101101233-1001300213300300-0320010313023100"></a>

## vmware.not_managed.node_list.interface_list.bond_interface.active_backup — active_backup / 233112321323 / 2

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

<a id="canonical-0232023212013110-2033123031211001-3100213122320032-0032133230020011-3121330231302213-1223011120002322-2303121123103332-3132211330202003"></a>

## Direct properties — active_backup / 233112321323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331223023223212-3023223223320320-0300010301130100-0113131223231211-0222331112323001-3212223110033111-3003023313213013-3230321332213130"></a>

## Next pages — active_backup / 233112321323 / 4

- [vmware.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3113233331323212-0201311132302221-0323311320033121-1221020221221020-3312022121220131-1310230333121213-3312122212120123-3133323202010303)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2130120320303322-3321320202312321-2021131333120113-2332323013211201-0132230023220011-0232221021323110-1010011312210232-3012030023222333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323100330333313-1213302131022301-1132213102301012-3103302110231031-3123332010212200-2301012032212302-3033133300200330-3131230223012330"></a>

## vmware.not_managed.node_list.interface_list.bond_interface.lacp — lacp / 323222010211 / 2

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

Upstream description:

LACP parameters for the bond device.

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

<a id="canonical-0013102232313113-0122002313300023-2002031200211313-2230213013202101-3222220202103111-0300132201223301-0120303131312213-1102203231303221"></a>

## Direct properties — lacp / 323222010211 / 3

<a id="canonical-1200230302122301-2232120123323111-3033223220312333-1331330311310230-3103211112030100-0233220200130001-0100002202330000-1120033223331311"></a>

<a id="canonical-2221201222231330-2033001223213021-1130211031020101-0132020133201132-3122023321001210-0030332220302123-0000222010122313-1301103301211302"></a>

## rate property — lacp / 323222010211 / 4

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

<a id="canonical-1013222112103031-3010201233323320-3202122120122122-2320220032010330-0220122112223013-1210012321200312-0003222112122323-2010201022311300"></a>

## Next pages — lacp / 323222010211 / 5

- [vmware.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3113233331323212-0201311132302221-0323311320033121-1221020221221020-3312022121220131-1310230333121213-3312122212120123-3133323202010303)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2002000010203020-1313101303100132-3221022203001132-2310200303013001-2231130113003003-2000023321030001-0130012302212201-1131013323110103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311110123303012-3312230220323101-1012310103003012-0123313232210131-3330030233301010-3030302002233133-2303130002220033-0002211213030310"></a>

## vmware.not_managed.node_list.interface_list.dhcp_client — dhcp_client / 132001321322 / 2

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

<a id="canonical-1201020200112213-2222202103210321-1232332000103202-2213003123231301-0223133213333001-3102323233003202-0221232310133210-0211223013301213"></a>

## Direct properties — dhcp_client / 132001321322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333323203322302-2123212203310020-3020012112310333-0003211311333103-1113231011201012-0222200030212203-1232033310320203-1223101302321031"></a>

## Next pages — dhcp_client / 132001321322 / 4

- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031302030220313-0310313330212130-3103122321332011-2130203102122130-1310212233302210-1113121023033300-0010112123030002-3030123212331111"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server — dhcp_server / 112312003111 / 2

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

Upstream description:

DHCP server configuration for this interface.

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

<a id="canonical-3201010300223321-3301021303231000-2021211132221320-3113022132112220-0302000111101311-3332131211321020-1231122122330211-2233011022213230"></a>

## Direct properties — dhcp_server / 112312003111 / 3

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0320110132001232-1212113133132310-1032110023222132-3203003230233032-1111202032101132-1233313000331011-2101322222213322-2021003032002031): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1303101120331011-1120100003103333-1002213010101100-3230221311032330-0312301003032301-3013232020203303-0010220000322023-2132001030010320): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3312332001302203-1100203030010230-1023021103011201-1132103011022313-3211122221323002-1211022311330013-2003133301021121-3110130232123100): complete subsection reference.

<a id="canonical-1211333103231232-3320301112130301-3010100101330231-2210112023011030-2021223120210002-1212101331331311-3333033030221101-2133100022113020"></a>

<a id="canonical-2223221021331021-3100331312113020-1022211133323201-3211333321313231-3231111001131320-3310013133122211-2312002112330100-0002231103001311"></a>

## dhcp_option82_tag property — dhcp_server / 112312003111 / 4

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-1003311022212130-2032233003023000-1103111311121133-0111123113101332-0231031220320303-1001300211130122-1103100203031013-0213120022300310"></a>

<a id="canonical-3013201220123120-1013233100020312-2321312331102033-1300333131023233-1132110321023210-1130011332012132-1203322202121210-2223101103020102"></a>

## fixed_ip_map property — dhcp_server / 112312003111 / 5

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

<a id="canonical-3320222211021013-2021020310221300-1100002001221033-3210010302120322-2300223201312330-0223211111213313-3303101003323011-0211013222203012"></a>

## Next pages — dhcp_server / 112312003111 / 6

- [vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0320110132001232-1212113133132310-1032110023222132-3203003230233032-1111202032101132-1233313000331011-2101322222213322-2021003032002031)
- [vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1303101120331011-1120100003103333-1002213010101100-3230221311032330-0312301003032301-3013232020203303-0010220000322023-2132001030010320)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3312332001302203-1100203030010230-1023021103011201-1132103011022313-3211122221323002-1211022311330013-2003133301021121-3110130232123100)
- [vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2021031130232120-3200220000131200-1001020120200301-0023113313303313-0030333112223332-0111302212212100-2133112222330323-1311023021103331)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0320110132001232-1212113133132310-1032110023222132-3203003230233032-1111202032101132-1233313000331011-2101322222213322-2021003032002031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101101332022210-0331202222102210-2103212020221020-1120100300311023-0021303212002101-3111302102011020-1120003101131201-0110030111221101"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — automatic_from_end / 000013133011 / 2

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

<a id="canonical-3111222101122223-2111102332300201-2000121102131130-1022300301002303-2113210102312012-2031232310023033-1302123332030321-0102012233221310"></a>

## Direct properties — automatic_from_end / 000013133011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120333010130203-3000120313231001-1120033303012220-3300120221102320-2103302002322330-3010031103001331-3323111323111003-2132323031332303"></a>

## Next pages — automatic_from_end / 000013133011 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1303101120331011-1120100003103333-1002213010101100-3230221311032330-0312301003032301-3013232020203303-0010220000322023-2132001030010320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001332003031100-3121010012232000-0033112023322001-0231202211323011-2213113311223020-2212232102202010-0101132031203031-2032112000301303"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — automatic_from_start / 021203001021 / 2

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

<a id="canonical-0201332122113032-1032211002122003-1013202010010123-0330023010120033-1300300123313300-3213203313322130-0032221301001313-0321010131010301"></a>

## Direct properties — automatic_from_start / 021203001021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110111100201303-0313121213001132-1231201202233111-2120002123001022-1203200330101012-3031222321313330-0313223011222110-2111032031121103"></a>

## Next pages — automatic_from_start / 021203001021 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3312332001302203-1100203030010230-1023021103011201-1132103011022313-3211122221323002-1211022311330013-2003133301021121-3110130232123100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010212331123133-0212100033201331-3321331210203002-0122323103100212-0013023133311321-3223010113012310-0030302023223223-3333100100331121"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — dhcp_networks / 101101123300 / 2

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

<a id="canonical-1220202300111230-2312023103121010-3111233233112232-3201231102320331-0123110013322213-1103133331121221-1130210330301021-2103031130110220"></a>

## Direct properties — dhcp_networks / 101101123300 / 3

<a id="canonical-1210300111112022-0210011002123011-3331112010021131-0022303133330122-3011223233212010-2223033323322311-0232210310232032-0230010223230121"></a>

<a id="canonical-3313220022101112-1322333233010201-2231013220132312-1200021332321332-2122010002110312-0323030013331012-1032220111303231-2221222120311121"></a>

## dgw_address property — dhcp_networks / 101101123300 / 4

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

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

<a id="canonical-0311322013213312-0202012220121313-1113322210030301-3330211333000123-1230212023101330-1320331202021131-3310311132331211-3013200130302002"></a>

## dns_address property — dhcp_networks / 101101123300 / 5

Type: `"string"`. Computed.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

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

<a id="canonical-0310023030312200-2133133023202101-3023231032333100-0221121123012322-1220122221213012-1102101303201023-2222113312323020-2233310011103231"></a>

## network_prefix property — dhcp_networks / 101101123300 / 6

Type: `"string"`. Computed.

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

Upstream description:

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

<a id="canonical-2111003211200322-0210131030333333-3223201301332310-3321200121230130-1032022203212202-1201102230033132-1030312301003232-1132012221023133"></a>

## pool_settings property — dhcp_networks / 101101123300 / 7

Type: `"string"`. Computed.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

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

<a id="canonical-3223313011110101-3133023123211302-1223112211120030-1311111301132033-1121023333230202-3101320333003311-3211022213131123-0132210011331131"></a>

## Next pages — dhcp_networks / 101101123300 / 8

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0101101323233310-2231003203210003-3023100300130303-1130122031133212-2321200331300001-2021222221231000-3230001203330120-1122201221110103)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1323110321213113-2203323132221100-3122210220311030-1100010211103331-0010200031122321-0333111133302011-1023320030221133-2111233321210023)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](data-sources--securemesh_site_v2--reference--group-017.md#canonical-1102133112010213-0320030302302113-2212101002231010-2312003311100000-0033333103323102-2332211110210133-2121222220223213-1110110111230133)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3233112100230330-3202012031002311-2302300131131010-2320310011303020-2000320302121032-1222302011233320-0021300110003031-2122103332103031)
- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0101101323233310-2231003203210003-3023100300130303-1130122031133212-2321200331300001-2021222221231000-3230001203330120-1122201221110103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200232212233132-2310222002111121-1133233313333112-2131002213312332-3332000132233231-0230233000323113-1020110020332013-0223310031320310"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — first_address / 222031000010 / 2

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

<a id="canonical-0213322102323331-1120012310111223-0112122120212332-2330011201302122-2112101200132301-1312020221010023-0032301013132322-2002311322211030"></a>

## Direct properties — first_address / 222031000010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020001123303313-3022103112200113-0132131332133113-2323201102131113-1001001112133132-0020001013233213-1201003202111302-2203203010303011"></a>

## Next pages — first_address / 222031000010 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3312332001302203-1100203030010230-1023021103011201-1132103011022313-3211122221323002-1211022311330013-2003133301021121-3110130232123100)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1323110321213113-2203323132221100-3122210220311030-1100010211103331-0010200031122321-0333111133302011-1023320030221133-2111233321210023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220011313122031-3113022310032201-2301233101022203-1212223323020002-0331010120211033-2312013002210233-2321333331123323-0031110300111103"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — last_address / 312120012131 / 2

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

<a id="canonical-3111132300213323-2011031202130110-2102320223321000-1222003223331103-1120231201111320-2112203212230221-1302200302200022-2313213332232012"></a>

## Direct properties — last_address / 312120012131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1102010032310120-3330022131012130-3120211212133021-3231013000213001-1333212100121121-1333232322120323-1011121332031201-1220100311111231"></a>

## Next pages — last_address / 312120012131 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3312332001302203-1100203030010230-1023021103011201-1132103011022313-3211122221323002-1211022311330013-2003133301021121-3110130232123100)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1102133112010213-0320030302302113-2212101002231010-2312003311100000-0033333103323102-2332211110210133-2121222220223213-1110110111230133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313122312320130-0113221123120102-3111002031330231-0203020013032232-1320333320030330-3232332301110212-1232312323003011-2011210213221220"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — pools / 212221031023 / 2

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

<a id="canonical-3112201030133110-2112210320032123-0223212101313012-2010021131313332-3312032113130312-2220103230001313-3333021221333212-3221221220312100"></a>

## Direct properties — pools / 212221031023 / 3

<a id="canonical-3212032203121033-3121033233103022-3313323022201132-2331021301023311-1203030302103222-0331321230010003-1333100132213213-1031333031231311"></a>

<a id="canonical-1032011001201312-2332133222021120-3003310032012211-2131303220132301-1011221010321321-0000110020312231-0221123131010023-3312002233312123"></a>

## end_ip property — pools / 212221031023 / 4

Type: `"string"`. Computed.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

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

<a id="canonical-2132202111000212-0023223102233311-1233000022230010-2211211120120032-1313031010221003-3122102000320220-3131312030012301-0202221020030031"></a>

## exclude property — pools / 212221031023 / 5

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-1321230122111333-2203011111303320-2211303301020000-1130213312202031-2033123303010311-0330220200200031-2113020033301330-2102320332123021"></a>

<a id="canonical-3213132121232111-3210101023132331-3011202231010330-2313022032222231-2102221213030311-0330303011101001-1033202232121100-2220110120332133"></a>

## start_ip property — pools / 212221031023 / 6

Type: `"string"`. Computed.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

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

<a id="canonical-1123130222122011-2223002332121010-3223020333211333-3133001033120222-2220002110101303-3233110003312003-3003233200203323-2023231001322122"></a>

## Next pages — pools / 212221031023 / 7

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3312332001302203-1100203030010230-1023021103011201-1132103011022313-3211122221323002-1211022311330013-2003133301021121-3110130232123100)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3233112100230330-3202012031002311-2302300131131010-2320310011303020-2000320302121032-1222302011233320-0021300110003031-2122103332103031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102122233223102-2011232011122112-1312102112102320-3302102130312300-1022332213311222-1023302103000222-1301010112121000-0121313213030301"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — same_as_dgw / 212100303011 / 2

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

<a id="canonical-2300033002301232-2321320131111233-0001121333120310-3323102131113003-2111103002231013-0032223001032222-2121202232223023-3201121203032031"></a>

## Direct properties — same_as_dgw / 212100303011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101202133110310-1222321123130112-1332231233232200-2221120130321321-2302022103312023-0201033332010232-3012020312200010-2302213302311111"></a>

## Next pages — same_as_dgw / 212100303011 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-017.md#canonical-3312332001302203-1100203030010230-1023021103011201-1132103011022313-3211122221323002-1211022311330013-2003133301021121-3110130232123100)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2021031130232120-3200220000131200-1001020120200301-0023113313303313-0030333112223332-0111302212212100-2133112222330323-1311023021103331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133123212311311-3220220213022132-2022212312012132-3212111201333122-1210333121222222-2113230212201323-2300230110030303-3120220031021233"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — interface_ip_map / 331030010010 / 2

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

Upstream description:

Specify static IPv4 addresses per node.

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

<a id="canonical-0033100003023120-3323303000321102-0320231310010111-2212301301312011-2210312031023320-3200111100010131-3203020033223131-2112312221000232"></a>

## Direct properties — interface_ip_map / 331030010010 / 3

<a id="canonical-3333010210102310-0032333021033000-3022032133103123-1033123000010021-1020032123203223-2010201332311030-2213130123223301-3031301010031111"></a>

<a id="canonical-3233323100113311-0200011303230010-3222230232123232-3033032120033112-1323033022213020-1233331020201112-0231020333013112-0330000303110021"></a>

## interface_ip_map property — interface_ip_map / 331030010010 / 4

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

<a id="canonical-1232322301030300-2021213302103123-1111210022323132-3021032213302013-0133313322000320-3331211312230121-2123212212220203-3311230220323122"></a>

## Next pages — interface_ip_map / 331030010010 / 5

- [vmware.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2100312320002010-0031121021022313-1113100101312300-2021213023201121-0123330230233203-1131112032130212-1032132120200111-3331013023201130)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2131211010200313-3110301321130222-1102123200113032-2221200010222312-3311330101103223-1332332132131213-0022022112111000-0312023302203233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110221131232023-3322021231310331-3133110322122111-0031001232030012-3231211212203320-0103020003313122-0120233132131122-3222313000300322"></a>

## vmware.not_managed.node_list.interface_list.ethernet_interface — ethernet_interface / 321333131230 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-0211222221122032-0202311232031001-0001331010000021-3313132333320000-1302200010033333-1033320023320020-3210101221330230-1000010321323301"></a>

Type: `"single"`. Computed.

Configuration parameter for ethernet interface.

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

<a id="canonical-3020202032100000-0220001103103213-2223312013303122-0103112012230222-0213022232132311-3323310031132110-0202001211321111-1113332020220211"></a>

## Direct properties — ethernet_interface / 321333131230 / 3

<a id="canonical-0002310331000132-1110331332230133-0120232100012112-3223020011013200-3301323012310123-2230211012001120-2030221203033113-3232212030302301"></a>

<a id="canonical-2212023223213312-1331123121202131-2230230111002002-1333102000021233-0221131101211310-0123233022013111-3011220100130130-3221313031200223"></a>

## device property — ethernet_interface / 321333131230 / 4

Type: `"string"`. Computed.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Upstream description:

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2213230032010212-2000201201321230-0220322123232003-2111002231102232-3020211312201201-2333210200102002-1202111302311322-3013221201302201"></a>

<a id="canonical-0211320130331330-2222333123302203-1223230320221110-2332112101000100-3000112302203102-3122012300300022-3301021221012300-3312121312223021"></a>

## mac property — ethernet_interface / 321333131230 / 5

Type: `"string"`. Computed.

MAC Address. Configuration parameter for mac

Upstream description:

Configuration parameter for mac

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  }
}
```

<a id="canonical-2233133000033100-3001213313113203-0332203200302132-2101213332221223-2331032021023102-0012001333211330-3322100200321321-1102130003012201"></a>

## Next pages — ethernet_interface / 321333131230 / 6

- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231112021321003-3330332012301213-3113031201203033-3022221130120110-3131013231032230-0023000320132230-0311322311311301-3221221310313231"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config — ipv6_auto_config / 002303112310 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-1031111302020110-2120302023321000-3321302133332202-0122033312131020-1221132300102023-1300033001311100-2303102100133203-0332321311010113"></a>

Type: `"single"`. Computed.

IPV6AutoConfigType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

<a id="canonical-3102331313302302-2013121010020203-1223102321002331-0102020323220231-1032213300202320-2031313113202031-1322021022002103-0120122221032313"></a>

## Direct properties — ipv6_auto_config / 002303112310 / 3

- [host](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1132013312112011-1022000231113230-2333103132010021-1113013220021002-1313101023330120-1133232113001311-3313013011223212-0122102311222033): complete subsection reference.

- [router](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0013223033023011-3101013310332311-0020120001101222-1022211232333002-2121101202122232-1223120302232000-0000131101220133-0212203332220310): complete subsection reference.

<a id="canonical-3112100203000312-2312031111320212-3203030030300203-2221311113012231-0132023231200130-2012313033202200-0033320123333303-3132210212311003"></a>

## Next pages — ipv6_auto_config / 002303112310 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.host](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1132013312112011-1022000231113230-2333103132010021-1113013220021002-1313101023330120-1133232113001311-3313013011223212-0122102311222033)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0013223033023011-3101013310332311-0020120001101222-1022211232333002-2121101202122232-1223120302232000-0000131101220133-0212203332220310)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
