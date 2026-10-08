---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-2003022110023131-2020123013032313-0310123231223001-1223322211211210-0312230131030133-0311310220103133-1330133030212311-2123132120033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- origin_pools.pools.origin_servers.origin_servers.private_ip

<a id="canonical-2331310112133313-1233200310301033-1121111223333013-0111322001120320-0230101230300101-1030212320011122-3332133011103032-2031301020000210"></a>

Type: `"single"`. Computed.

Specify origin server with private or public IP address and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]",
  "x-ves-oneof-field-private_ip_choice": "[\"ip\"]"
}
```

<a id="canonical-0330013222231101-0111302313211012-0113011310030212-0131123311312120-0202313222133302-1020302210021303-2211300321131003-2313331011000131"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.private_ip`

- [inside_network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0231333023131120-3022302322332031-1130021301023232-1030022000021331-1333032321312300-1010202203211010-3222330130022212-1133100311322331): complete subsection reference.

<a id="canonical-2023303032331010-0201031130103000-2300100330010320-0233110302321003-3000023320233310-3121310121031013-3103231001301120-3212333203103303"></a>

<a id="canonical-1211002122030221-1201021210030222-1310211113031021-2221132322032033-0122203012122133-1130101133322113-1131303002001022-2111200033032220"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.ip` property

Type: `"string"`. Computed.

IP. Exclusive with \[\] Private IPv4 address.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [outside_network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1322330330301222-1221133111122022-1322020213010013-2233030211120212-3213111003301011-0211221301212231-0031003220330021-1122201023311132): complete subsection reference.

- [segment](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2103031232002221-0320222113010120-2301111030232320-3111101130030230-2302323203212300-2010010222311103-0303011133002321-2231311232112231): complete subsection reference.

- [site_locator](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2302332012023230-2102201112120031-0103222010102101-2333012320202103-2033220230000200-2023202130120103-0020022023221222-1330233003113010): complete subsection reference.

- [snat_pool](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2232002130102200-1332010101121230-2100301101002302-2111331012101002-1313332323232122-3132200222311233-0012020333100300-2103013132111312): complete subsection reference.

<a id="canonical-0231333023131120-3022302322332031-1130021301023232-1030022000021331-1333032321312300-1010202203211010-3222330130022212-1133100311322331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2003022110023131-2020123013032313-0310123231223001-1223322211211210-0312230131030133-0311310220103133-1330133030212311-2123132120033110)
- origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network

<a id="canonical-0010322203003011-1212000123103300-2332230032003330-3330122221000113-0333332013323210-1003212202111121-2201123321021011-1303022132102113"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

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

<a id="canonical-1322330330301222-1221133111122022-1322020213010013-2233030211120212-3213111003301011-0211221301212231-0031003220330021-1122201023311132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2003022110023131-2020123013032313-0310123231223001-1223322211211210-0312230131030133-0311310220103133-1330133030212311-2123132120033110)
- origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network

<a id="canonical-0121202032233212-1310133303311002-1310130233232333-0003102320123300-1131320332320223-0322312122223232-3231023202202133-0112123122103213"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

<a id="canonical-2103031232002221-0320222113010120-2301111030232320-3111101130030230-2302323203212300-2010010222311103-0303011133002321-2231311232112231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.segment` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2003022110023131-2020123013032313-0310123231223001-1223322211211210-0312230131030133-0311310220103133-1330133030212311-2123132120033110)
- origin_pools.pools.origin_servers.origin_servers.private_ip.segment

<a id="canonical-1113313212013321-2022103022132332-0032330333130230-3311132113102000-3222221321202022-2330311213033030-1113103122221320-1003210200211030"></a>

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

<a id="canonical-1003102133021231-3020101321320212-1202013121301112-3100011102303110-0133322002200231-3011231323213212-0101330213122020-1221201312123110"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.private_ip.segment`

<a id="canonical-3220032020132013-0132031230030132-2223212012011002-3122311313001123-1010131023032120-0311132223321321-0111312001003300-3130323201020300"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.name` property

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

<a id="canonical-3303033122022030-3133133213132323-2312033013100033-0300010022122202-3113000322232023-3221320032121100-0221103133211202-2002313221021020"></a>

<a id="canonical-1221133102212111-1021232220210300-2033030223123112-3001200101032201-1110331010103000-2110113333311000-3302222121111212-3311202113323203"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.namespace` property

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

<a id="canonical-0102231011123312-0012030313030201-2212232331330112-1113132203121213-0223321202110311-3203011231300321-0200010100000021-0230331323100122"></a>

<a id="canonical-0221200122303021-2012021022033210-3222032332211103-1030232321131010-0112312210033022-0313201322103332-3200003200222302-1001223003213201"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.tenant` property

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

<a id="canonical-2302332012023230-2102201112120031-0103222010102101-2333012320202103-2033220230000200-2023202130120103-0020022023221222-1330233003113010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2003022110023131-2020123013032313-0310123231223001-1223322211211210-0312230131030133-0311310220103133-1330133030212311-2123132120033110)
- origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator

<a id="canonical-3302002313311110-3230032223330121-1312310222223032-3011101321100013-0022202013211300-2100322123120212-0231022112103031-2300221213323312"></a>

Type: `"single"`. Computed.

This message defines a reference to a site or virtual site object.

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

<a id="canonical-2102133011032322-1223130033002023-3232032101003211-0331222021330302-2322111113022013-0303111113102212-2113030201013023-2102120120110000"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator`

- [site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2323230220101012-3113212203313332-3320033101033002-2313323332203321-0012321022010311-2111112311211200-3021212323133203-2300310330011100): complete subsection reference.

- [virtual_site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2032022311200333-2220310201210321-0033111103330220-1303120031013020-0330022311300121-1212223321100000-1310311123010102-0000011120220220): complete subsection reference.

<a id="canonical-2323230220101012-3113212203313332-3320033101033002-2313323332203321-0012321022010311-2111112311211200-3021212323133203-2300310330011100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2003022110023131-2020123013032313-0310123231223001-1223322211211210-0312230131030133-0311310220103133-1330133030212311-2123132120033110)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2302332012023230-2102201112120031-0103222010102101-2333012320202103-2033220230000200-2023202130120103-0020022023221222-1330233003113010)
- origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site

<a id="canonical-0000123300013210-1211120333113203-0132010312033211-0123310021003030-3033020122010021-3103132022121030-3001230301210123-0202231310311312"></a>

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

<a id="canonical-3011013302213121-0322222102120023-2002223131002131-1311230132311110-1032021312122020-2212002333301002-2011112103330203-1011231233133120"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site`

<a id="canonical-0313020003320111-0030011133102310-0130200300032333-3212313301233232-0013323130032213-2223012113323022-1122012123201223-1211223200131121"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.name` property

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

<a id="canonical-1103212311122202-3100101300231132-1213123111130221-2130200012301022-3022003203331312-0320231210122122-0133103100021113-0231002210000012"></a>

<a id="canonical-0312121323232113-1010310331132223-1002210133033002-1132211131030201-2122201112101201-3210123130300120-1113233023233031-2130003321230310"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.namespace` property

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

<a id="canonical-1320303210231232-0010311302100132-0201331030213113-1120331311230223-1322031001313313-1320112013021201-0021313322031300-1223212031320320"></a>

<a id="canonical-0132121303031320-1100110223330110-1223122030130210-3323203213031232-2221103132003212-3112303120212211-1010000022331232-0230320011033011"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.tenant` property

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

<a id="canonical-2032022311200333-2220310201210321-0033111103330220-1303120031013020-0330022311300121-1212223321100000-1310311123010102-0000011120220220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2003022110023131-2020123013032313-0310123231223001-1223322211211210-0312230131030133-0311310220103133-1330133030212311-2123132120033110)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2302332012023230-2102201112120031-0103222010102101-2333012320202103-2033220230000200-2023202130120103-0020022023221222-1330233003113010)
- origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site

<a id="canonical-2112221121330022-0013300222002200-1122221331020100-1000022112113312-1330221130121003-2111210130201323-2111001201123302-1321120130103120"></a>

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

<a id="canonical-2323230211132021-0120211233003321-0131030001000033-3130013231211320-2202100303113322-2222332231220102-3202010132300000-3033103032022122"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site`

<a id="canonical-1333300110133201-1101131001201030-2310230320222123-3000312000313121-0321203032311103-3333221320311212-3313031003333011-2233110101011230"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.name` property

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

<a id="canonical-3100120002320023-2201100013231313-0133312011212232-2123330313321021-1232201102312101-2131102320221010-3130120333003301-1031020132212333"></a>

<a id="canonical-2022200023133322-1331112120223333-2120110031003100-0300232330131330-3120211230203011-0100223111212031-2102210231200323-2111000312223210"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.namespace` property

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

<a id="canonical-0121000310011132-0031213203011330-1331222321021311-2100121222130212-1010000010003200-3222033311112322-2200321210110032-3223011012020333"></a>

<a id="canonical-1231122111332203-0021100123213100-2210332331012303-2223012002033323-1202330202233010-2320101020103302-1203121220203322-3211032203333213"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.tenant` property

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

<a id="canonical-2232002130102200-1332010101121230-2100301101002302-2111331012101002-1313332323232122-3132200222311233-0012020333100300-2103013132111312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2003022110023131-2020123013032313-0310123231223001-1223322211211210-0312230131030133-0311310220103133-1330133030212311-2123132120033110)
- origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool

<a id="canonical-2022030121202102-3010303020221303-3031332323133112-2311002033033021-1031130301330013-2320002000223111-0011123303012200-0300010003300002"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

<a id="canonical-3102323302333020-2312211311223003-0233332103111010-1021312001210333-3321312033200122-3123213202212203-0000200332101313-0031223031302332"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool`

- [no_snat_pool](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0131210331133133-1200103300033130-1312120022010100-3302310320220020-3310213211001122-2212223003120013-1121103031121303-1313031122100132): complete subsection reference.

- [snat_pool](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3120203122013111-3330203010230003-3013030203120130-3222010102131330-2221023100133031-2230202301011301-0012102113203100-0013300201233131): complete subsection reference.

<a id="canonical-0131210331133133-1200103300033130-1312120022010100-3302310320220020-3310213211001122-2212223003120013-1121103031121303-1313031122100132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2003022110023131-2020123013032313-0310123231223001-1223322211211210-0312230131030133-0311310220103133-1330133030212311-2123132120033110)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2232002130102200-1332010101121230-2100301101002302-2111331012101002-1313332323232122-3132200222311233-0012020333100300-2103013132111312)
- origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool

<a id="canonical-2213212310113002-3223303121311130-0010010332230020-3013332232332123-3010212321003100-2333301001313133-0110303321101222-2031120201231110"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

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

<a id="canonical-3120203122013111-3330203010230003-3013030203120130-3222010102131330-2221023100133031-2230202301011301-0012102113203100-0013300201233131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2003022110023131-2020123013032313-0310123231223001-1223322211211210-0312230131030133-0311310220103133-1330133030212311-2123132120033110)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2232002130102200-1332010101121230-2100301101002302-2111331012101002-1313332323232122-3132200222311233-0012020333100300-2103013132111312)
- origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool

<a id="canonical-0230133210232311-2100300232330102-0122002303213112-3302032123211221-0212332202221330-3113022301332122-3133120302233230-3323020220230311"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-0022132000122110-3312330112203301-3313310121121011-1301223020323220-2213102132221332-0130322111031201-3303012202033023-3213023100313132"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool`

<a id="canonical-2321012201211133-3031331122000113-1102222302131222-3100133210211002-1233213212322323-0032331313211112-1121000120103032-3121333001022033"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool.prefixes` property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0301302103000331-2321000012100202-2033332133332322-1332223222321202-0032201121100013-0302222333032221-3311002011230333-1233131303200331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.public_ip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- origin_pools.pools.origin_servers.origin_servers.public_ip

<a id="canonical-2113011032311101-3113133101120122-3300030131200130-1332301310320033-3032230321113110-1313130001012132-1101331012131011-3001000200022333"></a>

Type: `"single"`. Computed.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

<a id="canonical-2321010133031211-0123021203310103-2302322223100311-1103010231202010-2102010000302111-3302302112300020-0133303220032333-1310100333013312"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.public_ip`

<a id="canonical-3030103313031333-3200032100310102-1023120101222313-3332121012131323-2332301000331033-0103021131032203-3321001101202030-0010330131333000"></a>

#### `origin_pools.pools.origin_servers.origin_servers.public_ip.ip` property

Type: `"string"`. Computed.

Public IPv4. Exclusive with \[\] Public IPv4 address.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3131332003012130-3200013310120130-0300132000112002-1220331130310023-3213322212232121-2031131233110313-1200331011110213-1110323201210131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.public_name` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- origin_pools.pools.origin_servers.origin_servers.public_name

<a id="canonical-0130012011320102-1021302121303132-3000331123211113-1233222001210101-3110110331121312-3201330033101033-1021233233330221-2312201031320100"></a>

Type: `"single"`. Computed.

Specify origin server with public DNS name.

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

<a id="canonical-1101020020310112-3202320110103103-3300133101312023-0033100331212323-1011101132212012-0310013033220331-3121210332132320-1000133230131110"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.public_name`

<a id="canonical-3010022033023220-0212200300302231-2333033001203303-1101131022203213-0120111201020332-1103011121030231-3130130310212121-1331131010322112"></a>

#### `origin_pools.pools.origin_servers.origin_servers.public_name.dns_name` property

Type: `"string"`. Computed.

DNS Name. DNS Name

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2133222010200210-0321012220311021-2220013112223300-0011220100222333-0030033312013010-2323221313212203-0202100102113113-3222232121023001"></a>

<a id="canonical-2211310222301233-1021031112312121-2121201221230031-0300313230100233-2122202012312203-1330321011011221-0000223233011103-0111330022012122"></a>

#### `origin_pools.pools.origin_servers.origin_servers.public_name.refresh_interval` property

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- proxy_advertisement

<a id="canonical-0020211011012111-0120022332000031-3102023112001321-1320200010231012-1001021323010201-0103220210122213-3111221320232220-3312130330033222"></a>

Type: `"single"`. Computed.

Configuration parameter for proxy advertisement.

Additional upstream details:

Proxy Advertisement Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"do_not_advertise\"]"
}
```

<a id="canonical-3321130231221022-0230300321012200-3023010220323002-2012223123221023-3330333321013233-1032223302122102-2011020020222132-0031102202311302"></a>

### Direct properties for `proxy_advertisement`

- [advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223): complete subsection reference.

- [do_not_advertise](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3312032130021330-3212102123320232-0213011310323122-0022113030030231-3323210230320221-1030023032302123-3132212013230101-0110213332222021): complete subsection reference.

<a id="canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- proxy_advertisement.advertise_custom

<a id="canonical-3230111111231132-0222333311021112-1001120322312330-3322110123213122-1013131302001031-2210311212320310-0101132322010300-0201103231102332"></a>

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

<a id="canonical-1220322232001130-1133030111123231-2012131230023323-2201122000330022-1323203023201102-3333022100012013-2212220213220200-0301100321010011"></a>

### Direct properties for `proxy_advertisement.advertise_custom`

- [advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303): complete subsection reference.

<a id="canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- proxy_advertisement.advertise_custom.advertise_where

<a id="canonical-3002113210120210-3211302001002102-2300301120211221-0310002123213201-2303002203121313-2110103301022331-0323313211302221-0112101002130331"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3223201012203211-3321223003200023-1132110102221333-2202123323213322-3203011303111122-2311302102030201-2313103020002132-3112022020023321"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where`

- [advertise_dualstack_on_public](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1320131110322313-3101001302031311-0312111013120122-3231203213110221-3201301023112310-0113011133211110-1013211032311133-1320013221102232): complete subsection reference.

- [advertise_on_public](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3110133021213102-0102133001001302-2202002301223211-3000313300032331-0103131213013203-2022020210223312-2010130223202212-2331233310021010): complete subsection reference.

- [advertise_v6_on_public](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1223302102212331-2120122013331131-3301201221122203-1330011213220120-1222131321023231-1221122122123003-0303003102311120-2002103130322320): complete subsection reference.

<a id="canonical-3321320330323133-0021011223112110-0002313300312010-2033120313322033-3020023300110020-1122131220032213-3132333201201320-1112022213132002"></a>

<a id="canonical-3123023210032021-0300332331221012-3232211122233212-1303131002030302-0033011211002002-0011330200012001-2021210321232112-3110223020013213"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.port` property

Type: `"number"`. Computed.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2111031010330232-0320320002220222-3233300022103210-2230012203321031-0313221321132203-1320001102102100-1030102003100000-3131221112122100"></a>

<a id="canonical-1130310212013211-3200003211311010-1131222330032123-2300301232203002-0320130323100130-1001031121203233-3000102221101303-0020132200333101"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2011220001130233-1001113120333232-1132211221130203-0331333130003122-3033020323211003-3331200020130321-3001312002000012-2101132013301301): complete subsection reference.

- [use_default_port](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2311020213333231-0110311320322301-0031111233020133-2230303213202323-3110222023012120-1223130211133201-3200020322223122-3022031230102310): complete subsection reference.

- [virtual_network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0002321213331213-3031003030301000-3020103001322102-1331230021110113-1133222012103201-3330032133030231-2232211001301220-0023300203100001): complete subsection reference.

- [virtual_site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1211323222210122-1130111030022212-0020323332120321-2103130323032233-0033123323232002-3233010033330003-1033312111111321-1311010213013021): complete subsection reference.

- [virtual_site_with_vip](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0301313112323223-1130020222220223-2203231102032211-2233233211022332-1020011330022232-0201311310121232-0123123020333210-0022212123000333): complete subsection reference.

- [vk8s_service](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1023212332121021-1203300333013112-0132223101030032-3311122322201033-2123023230101031-1310123320320032-3210221320220021-1232101110313030): complete subsection reference.

<a id="canonical-1320131110322313-3101001302031311-0312111013120122-3231203213110221-3201301023112310-0113011133211110-1013211032311133-1320013221102232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-2222321231312023-3013132311311232-0123210320120020-2013123023303103-2133232133200313-2012020110111022-1031222130313220-3022331231033222"></a>

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

<a id="canonical-2023232022003300-1131332033311231-1013011020020330-3021203120022310-2312030023322322-3333101322220122-1323010211011313-3023031303012110"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public`

- [public_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3133010213231131-0122323201313202-3001232120010211-2113313333032210-3113211323220023-2003122222211123-3031322220233022-2000132130021310): complete subsection reference.

<a id="canonical-3133010213231131-0122323201313202-3001232120010211-2113313333032210-3113211323220023-2003122222211123-3031322220233022-2000132130021310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1320131110322313-3101001302031311-0312111013120122-3231203213110221-3201301023112310-0113011133211110-1013211032311133-1320013221102232)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-3033212131211122-1301202010102021-2230312312030312-1123330211301010-2011231010220133-3202012302201222-2100330021103003-0010320203222103"></a>

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

<a id="canonical-3213022331103231-3233121330112300-2201221011013221-1213002133112212-3000001110133220-3022310122211133-1203332300011222-3121303323322223"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip`

<a id="canonical-2203313303323333-1330121131232232-3322200120012322-0232201330120001-2321003101100323-2103110122323031-2021313320303123-3020023233130223"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` property

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

<a id="canonical-0311010010310101-0201103201330323-2011331210113233-0210133300233203-2022121022001133-1301110122023120-0002020233112203-3213231000101323"></a>

<a id="canonical-0323313131121303-1103232121013002-0200131202002212-2211303333101133-0000221303310111-3311021212021302-2021202332101223-3031011330131211"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` property

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

<a id="canonical-1330321132220330-1233121030122223-2313003333022020-3302323212132320-0311212301032013-1103101031110220-3021131113323302-1131013303331302"></a>

<a id="canonical-1322031103031321-3321003002231301-0102132303311310-0233102013132322-1201230112222323-3210113232003221-1313312331322301-2003232300213311"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` property

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

<a id="canonical-3110133021213102-0102133001001302-2202002301223211-3000313300032331-0103131213013203-2022020210223312-2010130223202212-2331233310021010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- proxy_advertisement.advertise_custom.advertise_where.advertise_on_public

<a id="canonical-1021311111323300-0013012010022031-2332332102301323-1202212213132322-2010310011223103-0103133221101201-1320023103220212-3200320302321300"></a>

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

<a id="canonical-3313222330033010-2212031113120022-3010211130212302-1011311022230100-2120111131330333-3121311310111030-3223300222102302-2203012232212002"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public`

- [public_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1222302222311331-1011013211121310-2012112210231032-1113031031113010-1331311032322120-0111331010000301-1322311123222130-3311121001023011): complete subsection reference.

<a id="canonical-1222302222311331-1011013211121310-2012112210231032-1113031031113010-1331311032322120-0111331010000301-1322311123222130-3311121001023011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3110133021213102-0102133001001302-2202002301223211-3000313300032331-0103131213013203-2022020210223312-2010130223202212-2331233310021010)
- proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-2311332011002103-2310100023230001-3021233202013323-0013221212303003-0133131131332300-1203011111120032-3231003301101123-2000013330133110"></a>

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

<a id="canonical-0233010220301221-1223322332301302-2322202323003213-1310121100210320-0032201023101110-3221022011021220-0303322131331100-2011223002321032"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip`

<a id="canonical-0120031020013213-1113030133110333-2112310020113310-2120103131023302-1302013321001223-3330030103302103-1212233202001130-2323302133332120"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name` property

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

<a id="canonical-1021130023233012-1210132332231110-3002213212201312-0121210011001321-3131023123200003-2210300132331332-2331231310233110-0310210001021303"></a>

<a id="canonical-1021000301121210-1333113212300031-3010210123213123-2032301031013203-2313132101213302-1330000011322320-1132333100321333-1031323132231101"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` property

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

<a id="canonical-2310211011102302-2223023121303223-2203213132303130-3202102201002102-3131211133113113-2230132012122333-0332310102330210-2101003222223310"></a>

<a id="canonical-3330200200210021-0123333102132112-3221130213302223-0301211333010023-0013321230333132-3232111233220220-0003311111101323-3223200010110322"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` property

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

<a id="canonical-1223302102212331-2120122013331131-3301201221122203-1330011213220120-1222131321023231-1221122122123003-0303003102311120-2002103130322320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-2320321012110212-0332312031122030-3013300233111331-1100130120200322-0230233101233331-2222010100313230-1103233223222013-2031102113133213"></a>

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

<a id="canonical-0231121330130323-1023031312100200-2110220232133130-1321001322122331-3121002301013100-0220013030132133-2332331032300323-2010222112200000"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public`

- [public_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3000002222312112-3210202122321030-2003033112021133-1211101010310222-0303322301203201-0333222120302100-2101003113123131-0111112332123013): complete subsection reference.

<a id="canonical-3000002222312112-3210202122321030-2003033112021133-1211101010310222-0303322301203201-0333222120302100-2101003113123131-0111112332123013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1223302102212331-2120122013331131-3301201221122203-1330011213220120-1222131321023231-1221122122123003-0303003102311120-2002103130322320)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-2221300230303120-3112123123031131-2011031031022201-2100230122012013-0111322211231323-1121330230110101-1103300012122302-1330013311233010"></a>

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

<a id="canonical-1123322310223101-2013310132332102-3021012110233231-3131323212130132-1213120113312132-0212122301213233-2001011003121322-1313312302321130"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip`

<a id="canonical-1323102020221011-3110002302203023-2110320222212200-3233000011310023-3001030213232212-0121130213111231-1213311310330223-0303031301103232"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` property

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

<a id="canonical-0201311110001120-3223101102213311-2133311000300331-3310203103200020-1232300000302122-1030211113101001-2021210120001003-3011010102032210"></a>

<a id="canonical-3332321111322013-2332013110111201-0131101021322020-3131131030312003-0202333230332310-3000010013302131-2000301112122120-0101130211222302"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` property

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

<a id="canonical-2100012120313313-0122321310030132-0131003022311213-2303123033120022-2222111210302111-1003230203001033-0021130322210331-2021030020203012"></a>

<a id="canonical-0122100013333311-0101232203031223-1331201200100123-1310001331022330-3010333130010212-1012012333131003-1133310233213100-3123012303301333"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` property

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

<a id="canonical-2011220001130233-1001113120333232-1132211221130203-0331333130003122-3033020323211003-3331200020130321-3001312002000012-2101132013301301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- proxy_advertisement.advertise_custom.advertise_where.site

<a id="canonical-1032132212302120-1033321020322003-0123333302220030-0212233323201211-2030111103230132-1120321031200220-0122031320020232-1233021133313113"></a>

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

<a id="canonical-0312020310010112-0323203200102230-0022023201230111-1102113012231330-0022320000211222-0200000111030021-3330221331230013-2233301213032203"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.site`

<a id="canonical-1311331100011200-0321010203333033-0321122233122202-3300003133032301-1232322022200022-3330202323231030-3110000331031330-2333320123130011"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.site.ip` property

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2121000013313202-0322100310132301-0003100222330132-2310232230033002-2100120313000333-2132313202320213-1020012110030303-3133031102130011"></a>

<a id="canonical-1300010021103003-2203131331211331-1300313113211230-3021210213131202-2312202333110122-0300212331102132-1322300131021300-3221121121001030"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.site.network` property

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

- [site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1311113233310100-3102210030212112-2322031111201020-0010031321200023-0133020121012131-3102322312012213-0112300332203003-3031013202111222): complete subsection reference.

<a id="canonical-1311113233310100-3102210030212112-2322031111201020-0010031321200023-0133020121012131-3102322312012213-0112300332203003-3031013202111222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.site.site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- [proxy_advertisement.advertise_custom.advertise_where.site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2011220001130233-1001113120333232-1132211221130203-0331333130003122-3033020323211003-3331200020130321-3001312002000012-2101132013301301)
- proxy_advertisement.advertise_custom.advertise_where.site.site

<a id="canonical-0321131201020030-3232000222232030-2321121301033213-2130001202231023-2322322331131210-0012003300311002-2211330101310333-0013333033331003"></a>

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

<a id="canonical-1100221311201322-3333022310020031-0000200111301121-1112133013220010-2311102120211023-3000122212100113-2100120133022313-0233001132031300"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.site.site`

<a id="canonical-3320320123200122-0230313301010001-0321002102033302-1122130112300123-3312122313211333-2210313102121212-3203300030000332-2330031021320201"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.site.site.name` property

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

<a id="canonical-3000002132313002-2111302021003322-1121213230020322-1320102010202303-0111020233331321-0200030201013010-1021223130232123-3021102200201310"></a>

<a id="canonical-3001110212103120-0220303220032012-1102232010100032-3122311130320223-2100012300330312-2021101200121113-2200333113002333-2100120200220012"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.site.site.namespace` property

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

<a id="canonical-1121010132021202-0321002022320232-2222022203012233-2212200032021123-1231231101202012-3023231313013232-2130321232323223-3323131031322300"></a>

<a id="canonical-1132301201330313-1103100030013331-1220301311022213-1121300031011200-2231200023213003-0002330233100200-2130312132333213-3122313021322112"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.site.site.tenant` property

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

<a id="canonical-2311020213333231-0110311320322301-0031111233020133-2230303213202323-3110222023012120-1223130211133201-3200020322223122-3022031230102310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.use_default_port` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- proxy_advertisement.advertise_custom.advertise_where.use_default_port

<a id="canonical-0032031112220102-1111202020000200-2003020231023031-1000211011330131-2223313023313020-3012220300231320-0232300002032201-2231303131013101"></a>

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

<a id="canonical-0002321213331213-3031003030301000-3020103001322102-1331230021110113-1133222012103201-3330032133030231-2232211001301220-0023300203100001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_network` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network

<a id="canonical-1132223020322110-3321111012121321-3130022301311123-1030221002130330-0211321003123333-0310233231213100-2300323000033200-1031003021331320"></a>

Type: `"single"`. Computed.

Parameters to advertise on a given virtual network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-v6_vip_choice": "[\"default_v6_vip\",\"specific_v6_vip\"]",
  "x-ves-oneof-field-vip_choice": "[\"default_vip\",\"specific_vip\"]"
}
```

<a id="canonical-0021223222233310-3311100021021131-3332223122110010-0331100110223030-3302311030110101-2022100031331100-2223113031000303-1122122022010233"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_network`

- [default_v6_vip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2200132221022011-2230030111331023-2122132130331321-0233032223012212-0311221010203223-3030202310113332-1221322012032101-2030021321302021): complete subsection reference.

- [default_vip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1021120111120200-3113110200301021-1310133201003203-3221123130223301-2210202031001131-2320230313023100-0103300101032332-3033223312121112): complete subsection reference.

<a id="canonical-0003311311132021-2100300331313220-2301020320220301-0103311013231123-1310032212330221-1310323212320322-1100303222022030-0010202133032212"></a>

<a id="canonical-3033103202212030-1031001003103232-3033201202222021-3211213022331303-2132131223213332-2222321002133322-0120203301203121-3203022301313103"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip` property

Type: `"string"`. Computed.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-2111303010010002-3123330330211132-1110202200320330-3320133302102230-3231133013132220-2223011102221302-2322332213131220-3130313322220230"></a>

<a id="canonical-3111130332233320-1331212230212221-2312110222123133-2202303220320200-0220322113330233-1002311100310220-1213023212203312-2132313212301231"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` property

Type: `"string"`. Computed.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [virtual_network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0122313310330230-1320021220210332-0230200013321111-2332221102031013-3111123211212010-1233123003112222-2330020132202321-0203123222231003): complete subsection reference.

<a id="canonical-2200132221022011-2230030111331023-2122132130331321-0233032223012212-0311221010203223-3030202310113332-1221322012032101-2030021321302021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0002321213331213-3031003030301000-3020103001322102-1331230021110113-1133222012103201-3330032133030231-2232211001301220-0023300203100001)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-1201120322223300-3233201301321300-1023111220232233-0310012102101332-2032133200320100-2212101323301231-3323331033120233-0021213031322223"></a>

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

<a id="canonical-1021120111120200-3113110200301021-1310133201003203-3221123130223301-2210202031001131-2320230313023100-0103300101032332-3033223312121112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0002321213331213-3031003030301000-3020103001322102-1331230021110113-1133222012103201-3330032133030231-2232211001301220-0023300203100001)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-2333110310022301-1100133133120101-2112210211131231-1333123013300030-1110210232222011-0311300133103122-2211132000310111-2022131322121202"></a>

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

<a id="canonical-0122313310330230-1320021220210332-0230200013321111-2332221102031013-3111123211212010-1233123003112222-2330020132202321-0203123222231003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0002321213331213-3031003030301000-3020103001322102-1331230021110113-1133222012103201-3330032133030231-2232211001301220-0023300203100001)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-0231231003211002-0213012112320221-1030220202313203-0221300133331021-1022120200020012-3032033233112331-3221033322100322-1100212023210122"></a>

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

<a id="canonical-2022320220302223-2231012032023322-3223232222203211-3021101203231203-2333213311320222-3110113000120301-2032112132203013-3131003312321233"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network`

<a id="canonical-0113021200132133-1313133120221011-1110210000311021-0232200213212210-0313020332210200-2000303222130023-3120300100333223-2332301202223221"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name` property

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

<a id="canonical-3011102120230122-3030100113220220-3123011233123300-1023021102312321-1210201311023010-0133301200222201-1320330101220133-3201331100010313"></a>

<a id="canonical-0330311200201222-0203223110103303-2310120232313001-1032200122121310-3022132023111000-2133213231211201-3221100130300213-0123103220330212"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace` property

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

<a id="canonical-1212120121300332-0332121320202322-0203022010131310-1313233031333330-1021102201300121-0331102232222331-3213003031222200-3113112122320123"></a>

<a id="canonical-1110221203032023-1333113100032200-2203023021221330-2111010313121123-1210222313331111-2110310300002111-3033233331110132-2333200022203013"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant` property

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

<a id="canonical-1211323222210122-1130111030022212-0020323332120321-2103130323032233-0033123323232002-3233010033330003-1033312111111321-1311010213013021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site

<a id="canonical-0323223020130111-2311111313332201-2032032012212131-0102302031313020-1230332113013331-1033001213103213-0323033000110222-3310212220022301"></a>

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
