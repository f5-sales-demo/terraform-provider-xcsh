---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2031001210323100-0101013031033003-0132033030133213-0103323323203033-3313010122003210-0022130111012322-0131313102102113-1220113232322020"></a>

## Direct properties for `default_pool.origin_servers.private_name.site_locator.virtual_site`

<a id="canonical-0211001210021300-3132222332312303-3331311021330111-0221332231101300-0101011030012020-0023132200232201-1230330031331211-3333101300313323"></a>

### `default_pool.origin_servers.private_name.site_locator.virtual_site.name` property

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

<a id="canonical-1302123312023010-2220220111120132-2211021202121212-1330330210121310-0001113232102111-3230131133020212-1000011110321012-2131212322020233"></a>

<a id="canonical-3200003033311120-3333303211100301-3011322132301123-3311221003201030-2303132020103230-1103010230323203-3332120103132101-0002020122102101"></a>

### `default_pool.origin_servers.private_name.site_locator.virtual_site.namespace` property

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

<a id="canonical-1112230111211210-0121230320123003-3201021013013130-1320132302033002-0201213031221001-3033032220201022-3132111302120111-2300303322201132"></a>

<a id="canonical-0000223311200320-0202130233203232-0121013302232103-3111333130302303-3033021132101023-0010222202202301-0301301101302101-1000312122221211"></a>

### `default_pool.origin_servers.private_name.site_locator.virtual_site.tenant` property

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

<a id="canonical-3310022110103212-0302203330000030-3203123300323012-3232222312021232-3012311130300313-0231003000303220-1303031211212313-2321333311033331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-3031132322222002-1300113131221032-0020330231123313-2012010120103023-2212101132230332-1211331333233331-2131020011332033-1120131100222222)
- default_pool.origin_servers.private_name.snat_pool

<a id="canonical-0020301331012233-1112232200230211-3303122201223302-2103003210201000-3321001021232202-0013113130132300-2211130120323022-3131301333203120"></a>

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

<a id="canonical-2122102313221023-0131010333311021-3003021123132201-2230133223123122-2311103332231212-0301213320220301-2333023131302000-3210323131322222"></a>

### Direct properties for `default_pool.origin_servers.private_name.snat_pool`

- [no_snat_pool](data-sources--http_loadbalancer--reference--group-017.md#canonical-0322103220320121-3032122232231332-1211332221021103-2201001010130211-3233233102032200-1333011121002200-3301132032213022-1203230322213012): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-017.md#canonical-0202300232330010-3302330011233121-2110032013101320-2303331320300120-2103020001322312-2210113321201322-0213001012120011-0112033023020331): complete subsection reference.

<a id="canonical-0322103220320121-3032122232231332-1211332221021103-2201001010130211-3233233102032200-1333011121002200-3301132032213022-1203230322213012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-3031132322222002-1300113131221032-0020330231123313-2012010120103023-2212101132230332-1211331333233331-2131020011332033-1120131100222222)
- [default_pool.origin_servers.private_name.snat_pool](data-sources--http_loadbalancer--reference--group-017.md#canonical-3310022110103212-0302203330000030-3203123300323012-3232222312021232-3012311130300313-0231003000303220-1303031211212313-2321333311033331)
- default_pool.origin_servers.private_name.snat_pool.no_snat_pool

<a id="canonical-3130311221233203-0202332310013330-0101021033323120-0103130031231203-2223201033133230-0003201023032030-1112230122203023-2310110203001133"></a>

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

<a id="canonical-0202300232330010-3302330011233121-2110032013101320-2303331320300120-2103020001322312-2210113321201322-0213001012120011-0112033023020331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-3031132322222002-1300113131221032-0020330231123313-2012010120103023-2212101132230332-1211331333233331-2131020011332033-1120131100222222)
- [default_pool.origin_servers.private_name.snat_pool](data-sources--http_loadbalancer--reference--group-017.md#canonical-3310022110103212-0302203330000030-3203123300323012-3232222312021232-3012311130300313-0231003000303220-1303031211212313-2321333311033331)
- default_pool.origin_servers.private_name.snat_pool.snat_pool

<a id="canonical-1300033000123212-1212222023031311-3303011303233112-0302131233022023-1300232200201331-3221300321233001-0222121320300311-0311200312310131"></a>

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

<a id="canonical-0120122333203220-1133023322302100-2020100213202200-3312233300001311-3002331032310123-0230111100300110-2312111002102200-0212202222111103"></a>

### Direct properties for `default_pool.origin_servers.private_name.snat_pool.snat_pool`

<a id="canonical-0332031012121203-0320112033101232-2300200033103133-3320101021233301-0033102200311322-2101021211202011-3100033303202210-1032300212212130"></a>

#### `default_pool.origin_servers.private_name.snat_pool.snat_pool.prefixes` property

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

<a id="canonical-3230232123101031-2001223303123020-1213102222300331-0202333300210101-2132110133113330-1212330031022321-1002123133032130-2211123112002303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.public_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- default_pool.origin_servers.public_ip

<a id="canonical-0301222011221300-2231302232300021-1120223101212021-1003323030313323-1233000211223212-0201311012223002-1033020120100120-0310332231320302"></a>

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

<a id="canonical-0123222203003203-1303010220132011-0000331001103221-0222321030222311-3203233223012302-1323033021123302-0213302303311003-3132202030021032"></a>

### Direct properties for `default_pool.origin_servers.public_ip`

<a id="canonical-2221313322231232-1131033223021003-2332303332030303-0010321003010013-2322012102201231-0102222332011132-0213222301200202-0100313220321213"></a>

#### `default_pool.origin_servers.public_ip.ip` property

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

<a id="canonical-3213313232020122-3010103102320133-2122200332231332-1300302012030303-1213021123010331-2023230013331002-2031230031231102-3020232021202232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.public_name` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- default_pool.origin_servers.public_name

<a id="canonical-0022201210202020-1131011010000010-1020312200303300-2201310021131333-2021020211311031-3220323011303202-1231112112012103-0122012330113113"></a>

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

<a id="canonical-1122001200323101-1203222022012133-3003030203033200-2301221312302011-3033210122310103-2200103322223132-0131132220121023-3301230212211102"></a>

### Direct properties for `default_pool.origin_servers.public_name`

<a id="canonical-2313111201222213-3000102200220010-1233230331030210-0212131212232300-0233321233032230-2202213320122020-3001002031011331-2330012001303121"></a>

#### `default_pool.origin_servers.public_name.dns_name` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3323210020220210-1231122131313303-2211022011102020-2331113033121111-3033020131123122-0212122232001133-1322133031120133-2100133203100323"></a>

<a id="canonical-2100101213121331-1013131212001103-0223303000123212-0033331103213032-0211221323202102-0211100030101030-3202333030000003-3223300331310023"></a>

#### `default_pool.origin_servers.public_name.refresh_interval` property

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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-0211323132213300-2233131133130022-1221132031330120-3031032002223301-0212313121002022-3330103333233300-1120331120002031-0013013203231112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.vn_private_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- default_pool.origin_servers.vn_private_ip

<a id="canonical-3200303231101122-3122220210120100-3022000122210012-3311313200102210-2332103312202323-0233332311330023-2211311112101202-2211010023222121"></a>

Type: `"single"`. Computed.

Specify origin server with IP on Virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-virtual_network_ip_choice": "[\"ip\"]"
}
```

<a id="canonical-3021332002112311-3002030202223023-3233230132303330-0101101123223000-3031121133330330-3012200033301010-2223221033123303-1010111330213311"></a>

### Direct properties for `default_pool.origin_servers.vn_private_ip`

<a id="canonical-3211301322101020-1132222023002022-0231121331132021-2221202033201332-0031230011310233-2002111020202232-1100300013111010-0111212133021322"></a>

#### `default_pool.origin_servers.vn_private_ip.ip` property

Type: `"string"`. Computed.

IPv4. Exclusive with \[\] IPv4 address.

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

- [virtual_network](data-sources--http_loadbalancer--reference--group-017.md#canonical-0303200131030130-1112230111121222-0230322201201100-1230211221233020-0130200201210031-3221000010002011-3032011313101131-1331331001102220): complete subsection reference.

<a id="canonical-0303200131030130-1112230111121222-0230322201201100-1230211221233020-0130200201210031-3221000010002011-3032011313101131-1331331001102220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.vn_private_ip.virtual_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.vn_private_ip](data-sources--http_loadbalancer--reference--group-017.md#canonical-0211323132213300-2233131133130022-1221132031330120-3031032002223301-0212313121002022-3330103333233300-1120331120002031-0013013203231112)
- default_pool.origin_servers.vn_private_ip.virtual_network

<a id="canonical-0321200101332011-3202203121123022-0102300203003223-0312031310311313-1011200232031332-2011130031121223-1131312021213033-2133013113332011"></a>

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

<a id="canonical-0101313313110202-2113113202121002-1112020111001130-1310320023302033-0232322132111302-2221003122210332-0132022322202301-3320110123200221"></a>

### Direct properties for `default_pool.origin_servers.vn_private_ip.virtual_network`

<a id="canonical-3301120100032213-3002103010211012-3010332330121202-0121232333131133-1012000133122320-2113101212120232-1211110001333231-0002122110313331"></a>

#### `default_pool.origin_servers.vn_private_ip.virtual_network.name` property

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

<a id="canonical-0030132322111310-1113210123113311-2012123230122102-0000123120002303-2110221023130022-2300322200022332-3210312012211321-2223021232322130"></a>

<a id="canonical-3220202031300101-1011103322321202-2201312302322210-1033132031030113-0233032212301330-1001210032033011-2101101232231220-0223122321112202"></a>

#### `default_pool.origin_servers.vn_private_ip.virtual_network.namespace` property

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

<a id="canonical-1113231301011123-1302012002300231-2233012333111033-0311132101301323-1010102322213323-2131023122211311-0120021111100220-3032220202000203"></a>

<a id="canonical-3330300310011220-0331111220310302-1002023022000221-1201031201232333-0332130300022013-0231211003000211-2232311323010232-1310101221321310"></a>

#### `default_pool.origin_servers.vn_private_ip.virtual_network.tenant` property

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

<a id="canonical-0010223201122332-2321021313212022-0223002110212131-1011310123201121-3013320132013212-3232112322022332-1201100103213213-2331200010213312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.vn_private_name` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- default_pool.origin_servers.vn_private_name

<a id="canonical-3000002103113031-1032230200001022-0000032300002211-3001113131130312-3311013011331013-3232222203021121-2122210033030032-0331002320010322"></a>

Type: `"single"`. Computed.

Specify origin server with DNS name on Virtual Network.

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

<a id="canonical-3203230132133102-1310203112301213-3111000302021130-3111303223321232-2212302013231200-1113110132023133-3133223111313110-0200221301011212"></a>

### Direct properties for `default_pool.origin_servers.vn_private_name`

<a id="canonical-1300022022332131-1022212132210232-2321213311003230-0112301123230131-2311320110230230-0033002032331023-2011133231231311-2110001011331230"></a>

#### `default_pool.origin_servers.vn_private_name.dns_name` property

Type: `"string"`. Computed.

DNS Name. DNS Name

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [private_network](data-sources--http_loadbalancer--reference--group-017.md#canonical-3132233201312221-0030012100222331-2321131100302222-0030320023120023-3023321033310000-3030032120022330-0102221222130201-0232202332302233): complete subsection reference.

<a id="canonical-3132233201312221-0030012100222331-2321131100302222-0030320023120023-3023321033310000-3030032120022330-0102221222130201-0232202332302233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.vn_private_name.private_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-016.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.vn_private_name](data-sources--http_loadbalancer--reference--group-017.md#canonical-0010223201122332-2321021313212022-0223002110212131-1011310123201121-3013320132013212-3232112322022332-1201100103213213-2331200010213312)
- default_pool.origin_servers.vn_private_name.private_network

<a id="canonical-0231320113211011-0212330113202031-1000320121003333-0220223322101310-2123000100333110-1330122333111033-3101223130203111-1123100221110332"></a>

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

<a id="canonical-3211002130002322-3211302223203123-3101201003013221-0101330323300301-0013312213133010-2230000302113103-0032103332203312-3302001221120023"></a>

### Direct properties for `default_pool.origin_servers.vn_private_name.private_network`

<a id="canonical-0111033212323311-2100310013103322-1023101322130003-3200130123003033-0300331302122212-1211210123332012-0310221112020200-0002323030331320"></a>

#### `default_pool.origin_servers.vn_private_name.private_network.name` property

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

<a id="canonical-0133101312203131-2222100313313130-1121130230133311-2123232202022211-3001002031102210-3230032033112102-0021032101020113-3313133322033122"></a>

<a id="canonical-3223321223100002-2131131211120132-3123222133222123-0123313033132310-1220013201203023-3130000030110211-0301300201320000-2003023321033211"></a>

#### `default_pool.origin_servers.vn_private_name.private_network.namespace` property

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

<a id="canonical-3223101201331131-0031320220122333-3111313210333212-2013213002013030-1311022031002101-3020322300221130-0133120112220003-3102032220120010"></a>

<a id="canonical-1132123123111130-2023203211020131-3322102302303103-3130032213312312-0103000021031321-2033013333322012-2330313312033033-0323221112203122"></a>

#### `default_pool.origin_servers.vn_private_name.private_network.tenant` property

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

<a id="canonical-1220231310023223-2021212032323232-0022231222211011-1010322033032201-0313321021233313-2333033123131031-3012210101010020-1022012223113311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.same_as_endpoint_port` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- default_pool.same_as_endpoint_port

<a id="canonical-2232123001211120-0322201011213133-2002122222021133-1133203200220112-0132330021302110-1301010033220333-0132301313313133-2220101301033213"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-0112303232111020-1021002100101202-3000002001230023-3321200302003202-1220310122032130-2003211001121303-1230131031213110-1111332002330012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.upstream_conn_pool_reuse_type` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- default_pool.upstream_conn_pool_reuse_type

<a id="canonical-1203031202202101-1013001320110112-2302113310303111-1130021131200102-0021122303120132-2311332101023022-1210121113231130-1202333203102301"></a>

Type: `"single"`. Computed.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

<a id="canonical-1013020003313112-0320303020222312-1102033133131322-1113111331200233-2112221033221321-2310113103122331-3021010302212111-1220103030223313"></a>

### Direct properties for `default_pool.upstream_conn_pool_reuse_type`

- [disable_conn_pool_reuse](data-sources--http_loadbalancer--reference--group-017.md#canonical-0112131320231001-1202130131120330-3330102201113102-1210021122112301-1231233330102121-3023100213002113-2210100110322003-0322010032130120): complete subsection reference.

- [enable_conn_pool_reuse](data-sources--http_loadbalancer--reference--group-017.md#canonical-0233012302001110-1112210111133011-0110130111131302-0002103320132032-3232121320010301-3302003230131013-1302011032202023-1000111331003213): complete subsection reference.

<a id="canonical-0112131320231001-1202130131120330-3330102201113102-1210021122112301-1231233330102121-3023100213002113-2210100110322003-0322010032130120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.upstream_conn_pool_reuse_type](data-sources--http_loadbalancer--reference--group-017.md#canonical-0112303232111020-1021002100101202-3000002001230023-3321200302003202-1220310122032130-2003211001121303-1230131031213110-1111332002330012)
- default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-1021222130321213-3230123310301203-1031211002111300-3200122203211210-2132002023222303-0131213100000003-3111200233101021-1020212212112111"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable conn pool reuse.

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

<a id="canonical-0233012302001110-1112210111133011-0110130111131302-0002103320132032-3232121320010301-3302003230131013-1302011032202023-1000111331003213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.upstream_conn_pool_reuse_type](data-sources--http_loadbalancer--reference--group-017.md#canonical-0112303232111020-1021002100101202-3000002001230023-3321200302003202-1220310122032130-2003211001121303-1230131031213110-1111332002330012)
- default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-3220221131031223-2120312302231300-1122111313103111-1100011021020200-2202022210321023-1300000033313011-2201022023011012-3311133301111000"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable conn pool reuse.

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

<a id="canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- default_pool.use_tls

<a id="canonical-3211330103203302-1301322330222023-0002230201103100-0332201001323111-3132111021232032-0100221232231201-2331310103131010-0122211012000033"></a>

Type: `"single"`. Computed.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

<a id="canonical-1311031000121312-1132301110023213-1201213023030233-3311131222033032-2202323311112322-0011022331223012-0313213011310130-3011211321311021"></a>

### Direct properties for `default_pool.use_tls`

- [default_session_key_caching](data-sources--http_loadbalancer--reference--group-017.md#canonical-1201021320102230-2220130003032012-1210313322113021-0110010302311331-3212310302112122-3302022302231110-1311122021111123-2311112202013331): complete subsection reference.

- [disable_session_key_caching](data-sources--http_loadbalancer--reference--group-017.md#canonical-1133211223302203-3012222131331312-0211033222201130-1313030132222023-0332212132000330-3030120221213113-0112232222320122-3301310310031203): complete subsection reference.

- [disable_sni](data-sources--http_loadbalancer--reference--group-017.md#canonical-3223131121313133-3323113101101013-3123120302313101-2012330313203102-0012003202333110-3212022230230213-2103112332031102-2320113232323122): complete subsection reference.

<a id="canonical-0202000203003230-0010330320000010-3003020300101030-0233210002303122-2102202123231311-0330211000223132-3020311210123313-1020203330021310"></a>

<a id="canonical-1003201320111113-2323313320220003-3200033112100021-1101123302011020-1123203331333100-0101133322011120-3131201111313200-0233032013111233"></a>

#### `default_pool.use_tls.max_session_keys` property

Type: `"number"`. Computed.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

- [no_mtls](data-sources--http_loadbalancer--reference--group-017.md#canonical-0333131133101230-2023302030121221-0111032230031100-1322203201311132-0312200020030201-3201330103102002-1311213102003022-3231112331333003): complete subsection reference.

- [skip_server_verification](data-sources--http_loadbalancer--reference--group-017.md#canonical-2033203010112122-0220112232021000-1103032100113222-2011021132211023-3302123312203021-1223202213023030-2123020013323232-0103030110001003): complete subsection reference.

<a id="canonical-3302112132310011-3112310330130001-2211223313111103-1333031132031313-0223332022233023-3000101223021210-3220220231220112-1223011123110311"></a>

<a id="canonical-2320210010113111-2000213311221231-0210220021301301-0013220033311031-3000022021303132-3130022301332300-3321332302303321-1223300322001022"></a>

#### `default_pool.use_tls.sni` property

Type: `"string"`. Computed.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [tls_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-3110203310123302-0202113020220013-1211311220003232-0331212333201303-1103030110011120-3311010103322100-2211213231212323-1320100233231320): complete subsection reference.

- [use_host_header_as_sni](data-sources--http_loadbalancer--reference--group-017.md#canonical-1323030030121230-3230211332013231-2230022232310111-2200101321022323-0032010220000213-0221320212222030-2322132300202100-2133223231103310): complete subsection reference.

- [use_mtls](data-sources--http_loadbalancer--reference--group-017.md#canonical-0332011221232202-1100300300002220-1001021322222103-2011020232003102-0020313332201220-2112010320123302-1330011122002110-2311031331230202): complete subsection reference.

- [use_mtls_obj](data-sources--http_loadbalancer--reference--group-017.md#canonical-2223101000300310-1031222123301310-1003202331233300-3320220210113320-1031122020320101-2330231103223322-3133030010031023-3322221330310122): complete subsection reference.

- [use_server_verification](data-sources--http_loadbalancer--reference--group-017.md#canonical-3301132323330031-0110213303333113-0312001021332222-3201220113230012-1200332333301100-0201221312311313-0230300312123102-1023110221331202): complete subsection reference.

- [volterra_trusted_ca](data-sources--http_loadbalancer--reference--group-017.md#canonical-3311233201003203-3010222110230322-1110221302003332-3313221200001222-0231221202223223-2231000130203321-0001212203120010-3301132200212303): complete subsection reference.

<a id="canonical-1201021320102230-2220130003032012-1210313322113021-0110010302311331-3212310302112122-3302022302231110-1311122021111123-2311112202013331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.default_session_key_caching` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- default_pool.use_tls.default_session_key_caching

<a id="canonical-1332122230313111-0332321111230303-0132313203112332-1312231202022133-0322020220030102-3203331223133103-1212012311232223-3313333311021032"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

<a id="canonical-1133211223302203-3012222131331312-0211033222201130-1313030132222023-0332212132000330-3030120221213113-0112232222320122-3301310310031203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.disable_session_key_caching` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- default_pool.use_tls.disable_session_key_caching

<a id="canonical-2013312032301130-0222211002011223-2023003112311033-3123111132322131-1203102232112212-3033203123001320-2322132020103111-3332000220000031"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable session key caching.

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

<a id="canonical-3223131121313133-3323113101101013-3123120302313101-2012330313203102-0012003202333110-3212022230230213-2103112332031102-2320113232323122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.disable_sni` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- default_pool.use_tls.disable_sni

<a id="canonical-2330310012313210-3023302321020320-0210120011330320-3012320112120203-2312003002212213-1002302312212002-0313121330021322-1113300011303031"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable sni.

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

<a id="canonical-0333131133101230-2023302030121221-0111032230031100-1322203201311132-0312200020030201-3201330103102002-1311213102003022-3231112331333003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.no_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- default_pool.use_tls.no_mtls

<a id="canonical-1031101222332223-3331333322322103-0323320302321312-2311232201121021-2131123212121001-2222321331332100-3213120232323202-2103323022323133"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-2033203010112122-0220112232021000-1103032100113222-2011021132211023-3302123312203021-1223202213023030-2123020013323232-0103030110001003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.skip_server_verification` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- default_pool.use_tls.skip_server_verification

<a id="canonical-3322101112012301-0032022020321111-2223011013312112-1022221100031220-1312021100033122-0022210132300212-1123332003002022-1221203102033330"></a>

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

<a id="canonical-3110203310123302-0202113020220013-1211311220003232-0331212333201303-1103030110011120-3311010103322100-2211213231212323-1320100233231320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.tls_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- default_pool.use_tls.tls_config

<a id="canonical-2100303313020113-0012112110223202-0013001233001220-0331001102201020-1323221201113002-1310000002011210-0322013200301302-3020100131012010"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-0021220230323311-1020002113300001-1101210021120331-3122123231213221-2321210212322223-3323110313012023-2332322332003023-2133023023132003"></a>

### Direct properties for `default_pool.use_tls.tls_config`

- [custom_security](data-sources--http_loadbalancer--reference--group-017.md#canonical-2113200320100010-1203231201113220-0101201320213000-0310321132012013-0000112022101012-2330000311000121-0030202300300200-0032022131322330): complete subsection reference.

- [default_security](data-sources--http_loadbalancer--reference--group-017.md#canonical-1333233302332033-2332223302132012-2233213300030331-0303330111321210-0302332231232310-1320022002120202-3010021030232220-0300300023200323): complete subsection reference.

- [low_security](data-sources--http_loadbalancer--reference--group-017.md#canonical-0013032331122023-0122221231133103-2130310013131131-0213131100303211-2330222121221210-0302313022221331-3312022100220333-3111211130310200): complete subsection reference.

- [medium_security](data-sources--http_loadbalancer--reference--group-017.md#canonical-1032323103222301-2110101010201101-1202303200222320-3102102313310033-2111012003133021-0112103221001321-1132012011100223-1330230131113201): complete subsection reference.

<a id="canonical-2113200320100010-1203231201113220-0101201320213000-0310321132012013-0000112022101012-2330000311000121-0030202300300200-0032022131322330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- [default_pool.use_tls.tls_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-3110203310123302-0202113020220013-1211311220003232-0331212333201303-1103030110011120-3311010103322100-2211213231212323-1320100233231320)
- default_pool.use_tls.tls_config.custom_security

<a id="canonical-1213013313202030-2133203102321223-0203120102310303-3020031230220333-0110231301012300-0332033011331032-0102020133002301-3121010020200202"></a>

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

<a id="canonical-2311320300312110-1001231302010022-3002100021122331-1332303103201321-3012301213323020-2011221111303211-2223123201032223-1110310121112302"></a>

### Direct properties for `default_pool.use_tls.tls_config.custom_security`

<a id="canonical-3230123031211110-0012002023223110-2233330200121101-3103322020131002-3111233201322221-1100112331331002-0100033003313003-0133300130220313"></a>

#### `default_pool.use_tls.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-1113220012010311-3310112030000311-3101232121113112-1333321230032222-1201232113102210-2333301301213203-1123230030222010-2212312320013032"></a>

<a id="canonical-2232332023002233-2232223333121301-3333300201001203-3302010132111210-1200313232332323-0031133021023333-1320202011020310-0321330031003211"></a>

#### `default_pool.use_tls.tls_config.custom_security.max_version` property

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

<a id="canonical-3002113023310211-1210213111012310-0331220100101221-3130111112203211-1122012231123203-2011001330121311-1232331301333230-0213230032312120"></a>

<a id="canonical-2100012221010032-0032332110330202-1100132322003103-3032011032203331-1220211233331023-2301310102223011-1122200013112220-1102033120101200"></a>

#### `default_pool.use_tls.tls_config.custom_security.min_version` property

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

<a id="canonical-1333233302332033-2332223302132012-2233213300030331-0303330111321210-0302332231232310-1320022002120202-3010021030232220-0300300023200323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- [default_pool.use_tls.tls_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-3110203310123302-0202113020220013-1211311220003232-0331212333201303-1103030110011120-3311010103322100-2211213231212323-1320100233231320)
- default_pool.use_tls.tls_config.default_security

<a id="canonical-3201312202331311-0102221123330030-3011201123220300-1311113220122232-3110021103220103-3302113102110131-1123132120223320-0201231301331010"></a>

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

<a id="canonical-0013032331122023-0122221231133103-2130310013131131-0213131100303211-2330222121221210-0302313022221331-3312022100220333-3111211130310200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- [default_pool.use_tls.tls_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-3110203310123302-0202113020220013-1211311220003232-0331212333201303-1103030110011120-3311010103322100-2211213231212323-1320100233231320)
- default_pool.use_tls.tls_config.low_security

<a id="canonical-1122222122001130-0223330213132210-1312022213201323-1311001132332202-1020231221001331-3120031000200310-0102120202302002-2331213312332332"></a>

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

<a id="canonical-1032323103222301-2110101010201101-1202303200222320-3102102313310033-2111012003133021-0112103221001321-1132012011100223-1330230131113201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- [default_pool.use_tls.tls_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-3110203310123302-0202113020220013-1211311220003232-0331212333201303-1103030110011120-3311010103322100-2211213231212323-1320100233231320)
- default_pool.use_tls.tls_config.medium_security

<a id="canonical-1311003321321021-3320131101313112-3310322021020023-3113311033000313-2223100111212221-3231321100330220-1310001111213303-0131323130133331"></a>

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

<a id="canonical-1323030030121230-3230211332013231-2230022232310111-2200101321022323-0032010220000213-0221320212222030-2322132300202100-2133223231103310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_host_header_as_sni` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- default_pool.use_tls.use_host_header_as_sni

<a id="canonical-0231203320220133-0211123013322023-1032031110213332-0212100223013331-3220013122123103-2033332121203033-2233230100022310-1032111113302030"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-0332011221232202-1100300300002220-1001021322222103-2011020232003102-0020313332201220-2112010320123302-1330011122002110-2311031331230202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- default_pool.use_tls.use_mtls

<a id="canonical-0032313111213301-0131320233222313-1031021001231222-1110231130202101-3330033211011111-2103020320301231-0013031003131121-1323120322011322"></a>

Type: `"single"`. Computed.

MTLS Certificate. MTLS Client Certificate.

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

<a id="canonical-0212210203301023-0321032232231320-3311231213233111-0000130000103313-2320332003333121-0320132121303323-2213310221032222-2202112131221320"></a>

### Direct properties for `default_pool.use_tls.use_mtls`

- [tls_certificates](data-sources--http_loadbalancer--reference--group-017.md#canonical-2022100200102332-1231300023200311-0121033000000222-3220021232303100-3321113131033033-0022130220232133-0023303003021031-1320110123313021): complete subsection reference.

<a id="canonical-2022100200102332-1231300023200311-0121033000000222-3220021232303100-3321113131033033-0022130220232133-0023303003021031-1320110123313021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--reference--group-017.md#canonical-0332011221232202-1100300300002220-1001021322222103-2011020232003102-0020313332201220-2112010320123302-1330011122002110-2311031331230202)
- default_pool.use_tls.use_mtls.tls_certificates

<a id="canonical-3302310321323120-3013132221022103-1001033133022203-0033113230110011-2110331020331132-3010302123231112-2033320011030010-0202033233202110"></a>

Type: `"list"`. Computed.

MTLS Client Certificate. MTLS Client Certificate.

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
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0021112131300320-0202001023321210-0211212320113300-1200221203120130-3221230213033023-3110023333303120-1012213100200011-2313020013032131"></a>

### Direct properties for `default_pool.use_tls.use_mtls.tls_certificates`

<a id="canonical-2312002211002031-0131231201213100-0311033333222211-1302023100000133-1101133121012220-2223321211001232-0232001220122303-2101301000312310"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.certificate_url` property

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

- [custom_hash_algorithms](data-sources--http_loadbalancer--reference--group-017.md#canonical-2013030223032103-2301203021102112-1032321230121001-0021122023123121-3101303313201213-2333201120003020-3201201213311300-0332220312212003): complete subsection reference.

<a id="canonical-2332233330012321-2033000331133231-3333322102130003-0013200010210231-0212300112211221-0030103032112300-0320221012020213-1221002332222231"></a>

<a id="canonical-0132033313110333-1133121120012330-3022212131103233-3210102223313232-1331122202332110-3011021231132101-3333131313232120-3030231220010331"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--http_loadbalancer--reference--group-017.md#canonical-1022301313130121-3021222313033312-0233211222101000-2010321330130113-1223321121211100-2020200300211121-2231033130331233-1132313120002230): complete subsection reference.

- [private_key](data-sources--http_loadbalancer--reference--group-017.md#canonical-0321232103221021-1310121330201110-1030330102110301-0001112210000020-1221021101123233-3030223032220200-0310321201002221-1122013201301230): complete subsection reference.

- [use_system_defaults](data-sources--http_loadbalancer--reference--group-017.md#canonical-3231100120211221-3232103131013320-0011200032301213-0303103011220203-1130110030220120-3000100332032100-0103130312303203-2210031322122020): complete subsection reference.

<a id="canonical-2013030223032103-2301203021102112-1032321230121001-0021122023123121-3101303313201213-2333201120003020-3201201213311300-0332220312212003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--reference--group-017.md#canonical-0332011221232202-1100300300002220-1001021322222103-2011020232003102-0020313332201220-2112010320123302-1330011122002110-2311031331230202)
- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-017.md#canonical-2022100200102332-1231300023200311-0121033000000222-3220021232303100-3321113131033033-0022130220232133-0023303003021031-1320110123313021)
- default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms

<a id="canonical-2221222030232300-3221232013210110-3032230113032132-2320200323203212-0121232110210302-1000111111311003-0120200112100231-2033202003212022"></a>

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

<a id="canonical-1223223022213011-2333211323133021-3201202202122002-3222021221012121-3331111300001223-0332300320130300-3002220110030232-3322302013022300"></a>

### Direct properties for `default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms`

<a id="canonical-0123001323101111-1321321330211001-0331300032103010-3001030013103220-0030102331231220-1331101011310320-3313310023031312-3331212031122001"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms` property

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

<a id="canonical-1022301313130121-3021222313033312-0233211222101000-2010321330130113-1223321121211100-2020200300211121-2231033130331233-1132313120002230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--reference--group-017.md#canonical-0332011221232202-1100300300002220-1001021322222103-2011020232003102-0020313332201220-2112010320123302-1330011122002110-2311031331230202)
- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-017.md#canonical-2022100200102332-1231300023200311-0121033000000222-3220021232303100-3321113131033033-0022130220232133-0023303003021031-1320110123313021)
- default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="canonical-3012123100133210-2023301110213033-1001210032313210-2110311301222313-3112032030203233-0223001033102003-0232301311313310-0032220001112011"></a>

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

<a id="canonical-0321232103221021-1310121330201110-1030330102110301-0001112210000020-1221021101123233-3030223032220200-0310321201002221-1122013201301230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--reference--group-017.md#canonical-0332011221232202-1100300300002220-1001021322222103-2011020232003102-0020313332201220-2112010320123302-1330011122002110-2311031331230202)
- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-017.md#canonical-2022100200102332-1231300023200311-0121033000000222-3220021232303100-3321113131033033-0022130220232133-0023303003021031-1320110123313021)
- default_pool.use_tls.use_mtls.tls_certificates.private_key

<a id="canonical-0003012300113200-2023213030323212-0233102021202021-2211022100023010-3333110312333001-3313302100122202-1013201232301131-0230310302110203"></a>

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

<a id="canonical-0132211323133212-3120213303131110-3131102212111003-2301333023001223-3012312321120130-1020210320021213-0030302200330112-2311200210222301"></a>

### Direct properties for `default_pool.use_tls.use_mtls.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-017.md#canonical-1313231032232333-3010233333232001-3322122321311223-0013122231211113-2033123332022303-1011222010301222-0200333221302310-2102023102202323): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-017.md#canonical-3030023122112131-0011033301011123-0221123320010321-2302303333231133-2123332312320212-3220112303022313-3233103010030233-3330310123303003): complete subsection reference.

<a id="canonical-1313231032232333-3010233333232001-3322122321311223-0013122231211113-2033123332022303-1011222010301222-0200333221302310-2102023102202323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--reference--group-017.md#canonical-0332011221232202-1100300300002220-1001021322222103-2011020232003102-0020313332201220-2112010320123302-1330011122002110-2311031331230202)
- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-017.md#canonical-2022100200102332-1231300023200311-0121033000000222-3220021232303100-3321113131033033-0022130220232133-0023303003021031-1320110123313021)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-017.md#canonical-0321232103221021-1310121330201110-1030330102110301-0001112210000020-1221021101123233-3030223032220200-0310321201002221-1122013201301230)
- default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-1031102002133111-3321331321102313-3323121003333113-2003201333033232-3001200213130033-0233123100101011-3012132331123211-0200311331222231"></a>

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

<a id="canonical-0322313322311031-3333011223313011-0332112333001103-2130332321231110-2232231003211313-1221301201033121-1101000203333123-3023023032013231"></a>

### Direct properties for `default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-3030202121212121-1110313130310123-2103130120032331-1013211311103011-2310103101220203-1111302321113132-2220121210333320-1021031113131101"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3231101102313322-3010013301002202-0110210031221012-1100111213300210-1210223201310110-2331203231211100-3030232013013322-1012211101331113"></a>

<a id="canonical-1320000232200210-0231233122131301-0030022002322100-1011331031113013-2133332022212332-1320120212021211-0100031103311001-1100013203201101"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0333123110202111-2303233111331102-3112113013200012-0002203220031011-3320113111300023-0330100313012110-0213220301102030-2130012130012003"></a>

<a id="canonical-2322032220331012-1031322303112111-1032221303301103-0022020310000101-3100211112312312-0212203220330131-2320213330123200-1021320230001301"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-3030023122112131-0011033301011123-0221123320010321-2302303333231133-2123332312320212-3220112303022313-3233103010030233-3330310123303003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--reference--group-017.md#canonical-0332011221232202-1100300300002220-1001021322222103-2011020232003102-0020313332201220-2112010320123302-1330011122002110-2311031331230202)
- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-017.md#canonical-2022100200102332-1231300023200311-0121033000000222-3220021232303100-3321113131033033-0022130220232133-0023303003021031-1320110123313021)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-017.md#canonical-0321232103221021-1310121330201110-1030330102110301-0001112210000020-1221021101123233-3030223032220200-0310321201002221-1122013201301230)
- default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

<a id="canonical-0221030000021312-0000323133002111-0030113213023210-3300321030030201-3003333102210133-2022020322232010-2223323303031230-2002011023233200"></a>

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

<a id="canonical-3230333112111330-3132300030122012-1212323021133121-1001011100203013-2223000332012133-1201010303201323-0002010331002322-3233201022213210"></a>

### Direct properties for `default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info`

<a id="canonical-2232332210101200-0202203311200232-0231230000102130-2333313321102002-2020211010231223-0333102211311123-3030223302021033-0101112113112100"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1002032111131301-2102302331033302-2021023030000033-2223302121303121-2323330103202312-0333330002211202-1223102112013330-2221312310311201"></a>

<a id="canonical-3232101331120003-0110233223003103-1110031222100200-0303203313020310-2202333021210201-2001000030000303-1132103132113230-0203201033301122"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3231100120211221-3232103131013320-0011200032301213-0303103011220203-1130110030220120-3000100332032100-0103130312303203-2210031322122020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--reference--group-017.md#canonical-0332011221232202-1100300300002220-1001021322222103-2011020232003102-0020313332201220-2112010320123302-1330011122002110-2311031331230202)
- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-017.md#canonical-2022100200102332-1231300023200311-0121033000000222-3220021232303100-3321113131033033-0022130220232133-0023303003021031-1320110123313021)
- default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults

<a id="canonical-3001311030323321-2032233030203102-3321001101020312-2123123323000221-0201030321232321-1222011200130221-3011221013132030-2331321001221111"></a>

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

<a id="canonical-2223101000300310-1031222123301310-1003202331233300-3320220210113320-1031122020320101-2330231103223322-3133030010031023-3322221330310122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls_obj` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- default_pool.use_tls.use_mtls_obj

<a id="canonical-3000132021032220-1101221220122101-0103010022100310-3031002211233112-0312232301012221-0121223330020200-1222012113033221-1012120013122120"></a>

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

<a id="canonical-2132130221221131-1332320203232010-0330213320111023-2212210223132032-3101031022322200-2333233102313021-1132110001301122-2112023110111011"></a>

### Direct properties for `default_pool.use_tls.use_mtls_obj`

<a id="canonical-3033201132222001-2321121203202110-1323231002131323-2300121201100313-0110333200110303-0030131323021312-3030222011002233-0213121131032300"></a>

#### `default_pool.use_tls.use_mtls_obj.name` property

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

<a id="canonical-3330020110210032-0130010210332301-1213213331300030-1202323210003133-1321131320313120-3200111110233130-2221312233030301-0103013221000220"></a>

<a id="canonical-1313021122312120-0030133200220022-3211110131010033-0013232221301113-1003021131100033-3323110202103320-1323011300212330-2202012033111313"></a>

#### `default_pool.use_tls.use_mtls_obj.namespace` property

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

<a id="canonical-0321130013132223-2312333321123233-1123121102032012-3100332000120311-3302113001330011-0232123030100120-3312122313232312-2201033212322011"></a>

<a id="canonical-0003102033103121-0323230310220131-1203010030132010-1032201312131231-1000303213102113-1001200233320110-1132330301123301-0300211031312101"></a>

#### `default_pool.use_tls.use_mtls_obj.tenant` property

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

<a id="canonical-3301132323330031-0110213303333113-0312001021332222-3201220113230012-1200332333301100-0201221312311313-0230300312123102-1023110221331202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_server_verification` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- default_pool.use_tls.use_server_verification

<a id="canonical-1102120321230011-0132022023003333-3321011132312321-0202313323323203-1321211112120131-1200320230020113-3210003311111003-2312310011122233"></a>

Type: `"single"`. Computed.

Configuration parameter for use server verification.

Additional upstream details:

Upstream TLS Validation Context.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

<a id="canonical-2211233310001203-3303312113221222-2303103010032212-2311121011030120-2001321133012002-1110313331011003-3020032100123210-3030320301321300"></a>

### Direct properties for `default_pool.use_tls.use_server_verification`

- [trusted_ca](data-sources--http_loadbalancer--reference--group-017.md#canonical-1021022201011333-0213210013001322-1212110120102201-0011133102222230-3000211312100222-1011102200213103-0010032210203310-3022101113102302): complete subsection reference.

<a id="canonical-3230103133312021-2230033133010000-0211113300003113-3130222311110312-0001013012330210-3320233003300222-1203031201312320-0233111233031103"></a>

<a id="canonical-3212022200233310-1220232222202330-1233233311320302-3323003323020103-2033023222231332-1031133320230010-0010011230011010-0332202030033211"></a>

#### `default_pool.use_tls.use_server_verification.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-1021022201011333-0213210013001322-1212110120102201-0011133102222230-3000211312100222-1011102200213103-0010032210203310-3022101113102302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_server_verification.trusted_ca` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- [default_pool.use_tls.use_server_verification](data-sources--http_loadbalancer--reference--group-017.md#canonical-3301132323330031-0110213303333113-0312001021332222-3201220113230012-1200332333301100-0201221312311313-0230300312123102-1023110221331202)
- default_pool.use_tls.use_server_verification.trusted_ca

<a id="canonical-2113101213301022-1022220302320302-1121101000203023-1120332002013220-2131223130211012-1213212312031100-2020122030203133-3123302031323023"></a>

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

<a id="canonical-1103030022323331-2220103032331332-3210012110032003-3123333301031020-1312301121002202-2233332231020123-0223133300021002-1012000323003310"></a>

### Direct properties for `default_pool.use_tls.use_server_verification.trusted_ca`

<a id="canonical-3232030110233231-3312312130330201-1123233121203330-3033113202220303-0103200203332003-2312331222102311-0301113210211033-0201213101331101"></a>

#### `default_pool.use_tls.use_server_verification.trusted_ca.name` property

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

<a id="canonical-2202220323130233-3200323113131120-0023110022122023-3323332312322212-0232330101023233-0103212101203230-1032033100220031-3011003201303303"></a>

<a id="canonical-1220031303033011-2103230321031320-3202112032031320-1331322113313310-1211202231113200-1213013320003222-2020232332001030-3013210302220023"></a>

#### `default_pool.use_tls.use_server_verification.trusted_ca.namespace` property

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

<a id="canonical-3232002032033003-1112201031222111-3002212023100301-1012223312220102-1330112333102012-1301011210303022-0320031300321132-2311202310032320"></a>

<a id="canonical-0231323201212201-1100010030223300-1231012102033032-2200113220132111-3131313220210223-3330333323010120-0011312231321033-0310302323213213"></a>

#### `default_pool.use_tls.use_server_verification.trusted_ca.tenant` property

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

<a id="canonical-3311233201003203-3010222110230322-1110221302003332-3313221200001222-0231221202223223-2231000130203321-0001212203120010-3301132200212303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.volterra_trusted_ca` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-017.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- default_pool.use_tls.volterra_trusted_ca

<a id="canonical-0120213222013002-0013001303312002-3102200331220330-0111103233310023-3220011103110113-2031110233101110-3020111122121122-3121303310003023"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra trusted ca. Defaults to \`map\[\]\`. Server applies default
when omitted.

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

<a id="canonical-2322220323320321-1023031031123201-0303333212100213-0313203031010132-0222033300002102-3131312330012013-2002323331211102-0113131333012333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.view_internal` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- default_pool.view_internal

<a id="canonical-1230020213031113-0213232211011031-1112123300233300-3022222002333310-0030120131200202-3013003312213122-0111003300011223-1203010231113233"></a>

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

<a id="canonical-2102010223132203-3300101130121000-1223322022320220-0210331320303122-3331300133333332-1301123303032312-2101102333230122-1102313111021130"></a>

### Direct properties for `default_pool.view_internal`

<a id="canonical-2220003331000220-0201010001021301-0102220300323022-2311021020310120-1112320111002321-0321232002032310-2331300111001110-1312200111012012"></a>

#### `default_pool.view_internal.name` property

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

<a id="canonical-3232220210131111-0123213311011200-3130301203202311-3112331122233032-0011300203321100-3213323221023022-1011210023030132-0130213102021023"></a>

<a id="canonical-0101022313130333-0112332321122032-3012321103210111-0302120201011230-0301213021110300-3230011111132122-3030200202223113-3021022330321210"></a>

#### `default_pool.view_internal.namespace` property

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

<a id="canonical-2101330131000020-2232000232202322-2100302030012333-1103132121100122-0300121232301302-1231013313110032-1303022311232330-0002033020311033"></a>

<a id="canonical-0233330200201003-0023011131022002-1130231203313033-0202220100012312-0320131320320001-1011031102120202-1220022300322333-0112102331313321"></a>

#### `default_pool.view_internal.tenant` property

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

<a id="canonical-2120021320301122-2203211020120011-3331000133301003-2312333222002022-3232210233220330-1113220333233220-2323210311122033-0230203130321021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- default_pool_list

<a id="canonical-1130033132313333-3221230023302310-2121330233203033-0200002030033001-3110123103121011-1230303201033313-3011113212231322-1133312312033333"></a>

Type: `"single"`. Computed.

Origin Pool List Type. List of Origin Pools.

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

<a id="canonical-1211213330233032-1123033013323013-0303102100011330-2112201221002232-3220100323313030-1132032023133221-2003312211112222-1310210101003311"></a>

### Direct properties for `default_pool_list`

- [pools](data-sources--http_loadbalancer--reference--group-017.md#canonical-1010331020233221-0011133123131221-1211333012301211-1222312031321221-1011001112223022-3200330030200312-2012223000113213-3002220223201130): complete subsection reference.

<a id="canonical-1010331020233221-0011133123131221-1211333012301211-1222312031321221-1011001112223022-3200330030200312-2012223000113213-3002220223201130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool_list.pools` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool_list](data-sources--http_loadbalancer--reference--group-017.md#canonical-2120021320301122-2203211020120011-3331000133301003-2312333222002022-3232210233220330-1113220333233220-2323210311122033-0230203130321021)
- default_pool_list.pools

<a id="canonical-1323033001131111-1332102010311333-2033033213212001-1130333210113012-1320213231121130-1000311222313302-2022131112321212-1122201322210331"></a>

Type: `"list"`. Computed.

Origin Pools. List of Origin Pools.

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2310111103130022-1010102110202203-3112021223030211-2202003200310030-2333223121331302-2202011132102200-2101110310022023-2100032000020031"></a>

### Direct properties for `default_pool_list.pools`

- [cluster](data-sources--http_loadbalancer--reference--group-017.md#canonical-2323103010210212-2000221213313300-0200130303203213-0010100230111311-3111130000231220-1322013111010300-2030000310212300-1022220111000110): complete subsection reference.

- [endpoint_subsets](data-sources--http_loadbalancer--reference--group-017.md#canonical-0003100122012101-3330123103020021-0313303313011132-2333111112013021-1120011101212211-0312201013222332-1323203101212300-1323230312030320): complete subsection reference.

- [pool](data-sources--http_loadbalancer--reference--group-017.md#canonical-2133202011120310-0333011021320300-0211203020221032-3230012213333322-1230201333233031-2033332100232002-2211312100311110-3113310313122112): complete subsection reference.

<a id="canonical-2230021110301033-3023203330110233-2331200210110230-3102030322101010-3112202232332310-3131110002311330-3103213310001000-0120231022313320"></a>

<a id="canonical-1232210212312233-1031212032101111-1010301200131330-3200333313301311-2222012313312232-2032013132112333-2012202122032000-1122020102020301"></a>

#### `default_pool_list.pools.priority` property

Type: `"number"`. Computed.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-2121120202011331-3032313313300121-2331200111032102-2002030212300023-2122300030320120-3331012111322303-0120020112102303-1323030031212103"></a>

<a id="canonical-3332101303020233-1331132302212210-3322203223201121-2330233301320333-1112012302120103-3212332233221032-0011011121332023-0312331133312122"></a>

#### `default_pool_list.pools.weight` property

Type: `"number"`. Computed.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
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
  }
}
```

<a id="canonical-2323103010210212-2000221213313300-0200130303203213-0010100230111311-3111130000231220-1322013111010300-2030000310212300-1022220111000110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool_list.pools.cluster` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool_list](data-sources--http_loadbalancer--reference--group-017.md#canonical-2120021320301122-2203211020120011-3331000133301003-2312333222002022-3232210233220330-1113220333233220-2323210311122033-0230203130321021)
- [default_pool_list.pools](data-sources--http_loadbalancer--reference--group-017.md#canonical-1010331020233221-0011133123131221-1211333012301211-1222312031321221-1011001112223022-3200330030200312-2012223000113213-3002220223201130)
- default_pool_list.pools.cluster

<a id="canonical-0231021231001323-1231300012121200-1333221332321332-2311223133200212-3212032003113021-0320101210003111-1130021020321022-0021201000321132"></a>

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

<a id="canonical-0201231011333203-0323302121020302-2202021230120121-3020101031311303-0310201021033001-0110020112013313-1030020323123002-1033010033332220"></a>

### Direct properties for `default_pool_list.pools.cluster`

<a id="canonical-0111321001303020-2003101210030201-3022211012322320-2221032112301230-2033030101011311-2311322312223111-0123013113000210-1202300210113010"></a>

#### `default_pool_list.pools.cluster.name` property

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

<a id="canonical-1222311011123002-2202100302022221-3130301333133221-0320031030301023-2223023000301321-2021320332200033-3131123110331310-0120033201213133"></a>

<a id="canonical-1120033022010121-2100033232212201-1111202123002100-2330321001031121-0021010020302021-0000020330323112-3021120220322321-2110212111201121"></a>

#### `default_pool_list.pools.cluster.namespace` property

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

<a id="canonical-3203012311201111-3100013330203110-1330322033202221-2021312200210200-3132220313033102-2130223002122212-3120103003021101-2201312100011032"></a>

<a id="canonical-3232203101111302-1131231322230032-3123232013022233-0133131122120100-2113022113002020-0333023322330123-0113211330101200-1133022301313320"></a>

#### `default_pool_list.pools.cluster.tenant` property

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

<a id="canonical-0003100122012101-3330123103020021-0313303313011132-2333111112013021-1120011101212211-0312201013222332-1323203101212300-1323230312030320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool_list.pools.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool_list](data-sources--http_loadbalancer--reference--group-017.md#canonical-2120021320301122-2203211020120011-3331000133301003-2312333222002022-3232210233220330-1113220333233220-2323210311122033-0230203130321021)
- [default_pool_list.pools](data-sources--http_loadbalancer--reference--group-017.md#canonical-1010331020233221-0011133123131221-1211333012301211-1222312031321221-1011001112223022-3200330030200312-2012223000113213-3002220223201130)
- default_pool_list.pools.endpoint_subsets

<a id="canonical-3333033030210231-3322322310201310-1211122120211022-1330202231332031-0013333321133011-3032203223003222-3032011301101110-2011022132310030"></a>

Type: `"single"`. Computed.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

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
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "16"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133202011120310-0333011021320300-0211203020221032-3230012213333322-1230201333233031-2033332100232002-2211312100311110-3113310313122112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool_list.pools.pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool_list](data-sources--http_loadbalancer--reference--group-017.md#canonical-2120021320301122-2203211020120011-3331000133301003-2312333222002022-3232210233220330-1113220333233220-2323210311122033-0230203130321021)
- [default_pool_list.pools](data-sources--http_loadbalancer--reference--group-017.md#canonical-1010331020233221-0011133123131221-1211333012301211-1222312031321221-1011001112223022-3200330030200312-2012223000113213-3002220223201130)
- default_pool_list.pools.pool

<a id="canonical-0013210213202131-2223131032201120-1330322301301322-0102323031200332-1123222111322102-0311121011033013-2023001133231330-0201232201110031"></a>

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

<a id="canonical-1101002113031202-3022333031020110-1331231332213221-3201132132101312-3323102321301332-3123033012301113-0203203220221021-3110030211101031"></a>

### Direct properties for `default_pool_list.pools.pool`

<a id="canonical-2312310330003020-0213003310100130-0023101310202031-1102030302222202-2312131022302333-3121022031311312-0320132231022101-0100330200103123"></a>

#### `default_pool_list.pools.pool.name` property

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

<a id="canonical-1210110230220131-3032010321030201-2031131110122201-0302000003311031-3302222220233312-3033220101111030-3300010132003112-2231012123002221"></a>

<a id="canonical-2002332101303233-2302223011210311-0033102303210203-1122210012133230-1213333020020121-0300103012301013-1320123231301003-2010023313331033"></a>

#### `default_pool_list.pools.pool.namespace` property

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

<a id="canonical-3211232322313130-3023333100201122-0230331221112123-3232100321130301-2000111313201032-1213113212010221-1013333033111022-0212110211103000"></a>

<a id="canonical-1303211101202310-1000032022303220-0332023021232111-1022210312211200-2111230323210131-1002322310213311-1110101031300303-1111101130010230"></a>

#### `default_pool_list.pools.pool.tenant` property

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

<a id="canonical-2123320333302212-3312111002110010-3032303202101023-0302100013110231-3132000001122200-0133332310312030-0321211001120321-3111232322003312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_route_pools` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- default_route_pools

<a id="canonical-1320031330333310-3321310301133231-1221200111031301-1323031202011202-3120000203131301-2210321013332301-2120111112210130-1121030020021223"></a>

Type: `"list"`. Computed.

Origin Pools used when no route is specified (default route).

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2112301300102011-0103310033130013-3321201112321331-2303231331210312-3302211033033001-2031223033310233-2023033300320030-0233331310113013"></a>

### Direct properties for `default_route_pools`

- [cluster](data-sources--http_loadbalancer--reference--group-017.md#canonical-1202210300323330-2110233103013303-2212030030313201-2233223100100032-2033000003331213-3330302330103320-3013211023200320-3313331332301210): complete subsection reference.

- [endpoint_subsets](data-sources--http_loadbalancer--reference--group-017.md#canonical-1130213210101302-3030322332011301-1102321311220021-2230232113000031-3111212011230023-1032220221200011-1013102112233121-3302301220130310): complete subsection reference.

- [pool](data-sources--http_loadbalancer--reference--group-017.md#canonical-0222131133032203-1220120132131222-0311313032000032-1120102133321321-3001221001000112-2302011122113003-2322101303110331-3111302332022213): complete subsection reference.

<a id="canonical-1101220300211212-3030013033300023-1331301313132110-3021302221201021-0031213310202132-0130300333101101-3122213000221111-2331223123303202"></a>

<a id="canonical-3310020101301223-0301033013232023-3323312330232030-3313212002233112-3020121201232323-1310310101232130-0300333003233122-0300122011300302"></a>

#### `default_route_pools.priority` property

Type: `"number"`. Computed.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-1332100221120230-2330101002300020-0310110121320121-1320132010032322-1033130102302321-3112220102321301-2021323001103111-3123010111030102"></a>

<a id="canonical-1333033130202310-3032232230321120-0132123130101311-0310010303013013-2310202122122103-0012111332311020-2103333030301021-2031131220102001"></a>

#### `default_route_pools.weight` property

Type: `"number"`. Computed.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
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
  }
}
```

<a id="canonical-1202210300323330-2110233103013303-2212030030313201-2233223100100032-2033000003331213-3330302330103320-3013211023200320-3313331332301210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_route_pools.cluster` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_route_pools](data-sources--http_loadbalancer--reference--group-017.md#canonical-2123320333302212-3312111002110010-3032303202101023-0302100013110231-3132000001122200-0133332310312030-0321211001120321-3111232322003312)
- default_route_pools.cluster

<a id="canonical-1113223333020233-3032010313121310-3030310313132021-0220023320101213-3302230101333220-3320033131233222-3110231230123230-2332202110312302"></a>

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

<a id="canonical-1320320220112012-3333302202103122-0231220103012000-1103312312331120-0010123123103300-1321221200210021-3122030302232011-0031020013023313"></a>

### Direct properties for `default_route_pools.cluster`

<a id="canonical-0130233322021113-1202001021230001-3201221021123133-2001023003332002-0001220313112133-1211213213222330-0222121031030130-3132210210303232"></a>

#### `default_route_pools.cluster.name` property

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

<a id="canonical-2230001003211232-0102301222111013-2122331212232101-1132002212123121-3032012022010113-1121212300010222-2132110332031310-3132121022133233"></a>

<a id="canonical-1310203110101303-1123100202323121-1022333133302030-2102111321121231-0122332321311321-0230111201333331-0033022030021121-0113100332332213"></a>

#### `default_route_pools.cluster.namespace` property

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

<a id="canonical-0302210303311220-2323232113220203-3021101110232110-3312212021203122-1313011201213200-0102003213332231-3300213133011013-3033021132310021"></a>

<a id="canonical-2230033310010111-3232112212321300-1112033003201013-1133201331001022-2131321021320131-0023221310111110-2202323012331322-3320322101312033"></a>

#### `default_route_pools.cluster.tenant` property

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

<a id="canonical-1130213210101302-3030322332011301-1102321311220021-2230232113000031-3111212011230023-1032220221200011-1013102112233121-3302301220130310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_route_pools.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_route_pools](data-sources--http_loadbalancer--reference--group-017.md#canonical-2123320333302212-3312111002110010-3032303202101023-0302100013110231-3132000001122200-0133332310312030-0321211001120321-3111232322003312)
- default_route_pools.endpoint_subsets

<a id="canonical-1112220000233133-2102201101230331-2110213212223121-0000131202102112-0021102132321313-1301223200320010-0100222323203113-1233122201233131"></a>

Type: `"single"`. Computed.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

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
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "16"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222131133032203-1220120132131222-0311313032000032-1120102133321321-3001221001000112-2302011122113003-2322101303110331-3111302332022213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_route_pools.pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_route_pools](data-sources--http_loadbalancer--reference--group-017.md#canonical-2123320333302212-3312111002110010-3032303202101023-0302100013110231-3132000001122200-0133332310312030-0321211001120321-3111232322003312)
- default_route_pools.pool

<a id="canonical-0012220132331111-3001220212120310-0120130003111123-3323323111321102-2332123002212132-1030321322121111-3021001332213033-3111303033120122"></a>

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

<a id="canonical-3223010302123003-0003031102332001-2023132233130001-1133022021123323-0222312302201200-2313122230232031-1200302322021010-3103221313002320"></a>

### Direct properties for `default_route_pools.pool`

<a id="canonical-2210231122213302-3322222221310011-0131110111010332-0003020312001023-1121203112211120-1121223100322030-2123200211133212-2323332113330223"></a>

#### `default_route_pools.pool.name` property

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

<a id="canonical-3130100322210100-3212332102132232-3303103222131301-1313223131110231-1213301212333002-1130103200022300-3030202103002002-3001112021111120"></a>

<a id="canonical-0313131020010110-2111320130323310-0132113003031332-3001030122213201-3332302132223101-2213230103301130-3212303200113000-1122031031013032"></a>

#### `default_route_pools.pool.namespace` property

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

<a id="canonical-0233110032330331-1332103123333033-2031312012323100-0131103330303100-0103331213333301-2001233202100101-1021201030221210-0002033113213021"></a>

<a id="canonical-1230112223311113-3003321031110113-2002010021012201-3211022311023232-0130311012223101-0313302012113111-3001023110222112-2033313030221023"></a>

#### `default_route_pools.pool.tenant` property

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

<a id="canonical-0111232021032210-2203300033211333-3203221201122011-1312201110201102-2011000032031020-1230113100330322-3230101122211000-2101102002000223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_sensitive_data_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- default_sensitive_data_policy

<a id="canonical-0102302031102013-3203123231132111-0231200322301112-3320023102000113-0020102112101331-1203230320030212-3100230302211130-3223223022121210"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_sensitive\_data\_policy, sensitive\_data\_policy; Default:
default\_sensitive\_data\_policy\] Policy configuration for this feature. Defaults to \`map\[\]\`.
Server applies default when omitted.

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

- [default_sensitive_data_policy](data-sources--http_loadbalancer--reference--group-017.md#canonical-0102302031102013-3203123231132111-0231200322301112-3320023102000113-0020102112101331-1203230320030212-3100230302211130-3223223022121210)
- [sensitive_data_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-1233012233330121-3030022033230201-2333203001311122-0122203010231303-0110030001013320-3113122201121323-2222321313230113-3223013123232202)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103320133122210-1101320300230032-2103233020111021-3331023123322321-0022223323012102-2001013112130132-3013002230133211-2332223011332022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_api_definition` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- disable_api_definition

<a id="canonical-0310031002133033-3133323313302002-1023201212303220-0020212011021133-1133233311232010-3211233231213212-0132233203113201-3302223221232231"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-1302330231202101-0223003210100011-1103001000322103-2130013013301113-0031323333122311-1013313202011020-2010303010122300-2321321122232201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_api_discovery` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- disable_api_discovery

<a id="canonical-3301101303223122-1000211023332232-2102031310300001-2130312222002211-1211300300301211-1203322131321231-2303330310203133-1222223321111221"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_api\_discovery, enable\_api\_discovery; Default: disable\_api\_discovery\] Enable
this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [disable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3301101303223122-1000211023332232-2102031310300001-2130312222002211-1211300300301211-1203322131321231-2303330310203133-1222223321111221)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-1332310000223132-0300312232102030-2113102020202113-3131230231231122-2032032331222321-3211111300111301-1001021113320210-1322233103230012)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333200003312023-1221031330131101-0030223302230000-0200202023223123-3230111313110123-3030133101330021-2320023222010202-1332313230012121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_api_testing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- disable_api_testing

<a id="canonical-2330222032220201-0302311201021200-1203330011032320-1022202130300202-3111133002321333-3123200212110131-0003013132330113-2020330021120133"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-0002113012001322-1331231230330100-0313110022323200-2010333332303100-2323110232011312-3300211312030310-0133233122112310-0233100023212113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_bot_defense` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- disable_bot_defense

<a id="canonical-1130001201233333-3003230020322233-2213300103132002-0200323000130110-2010032013230101-1022120223232333-0000200323112101-3310030100011121"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable bot defense. Defaults to \`map\[\]\`. Server applies default
when omitted.

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

<a id="canonical-0002302111031031-1130133033030120-1300022330230010-2000210133200112-3322203230333132-3220320102203131-0311010231031001-2120213121211103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_caching` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- disable_caching

<a id="canonical-1312123312223031-3132131302203112-0321223312021102-2111003030332211-1303112101133122-3303033010123320-1323111123300010-0231330312020120"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable caching.

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

<a id="canonical-2123332133130320-0331211322111000-2330232110321202-2330201211130101-3011112020212323-1221111213133101-3032301031112132-3223101322300031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_client_side_defense` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- disable_client_side_defense

<a id="canonical-3313223321312022-3132121202023311-2200012303220320-3120033023232211-0103033210021322-3121211210300330-1103022020313021-0322333013120023"></a>

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

<a id="canonical-0112131023222001-3013102233003223-0212212313132321-3213111323012313-3103122301032310-1210111233031132-3121100100233010-0031322010212232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_ip_reputation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- disable_ip_reputation

<a id="canonical-1333110131232221-2013320132121202-0222200220111222-2233011322201031-0130313211123122-1130020301221012-2110003220113131-0210200232333302"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_ip\_reputation, enable\_ip\_reputation; Default: disable\_ip\_reputation\] Enable
this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [disable_ip_reputation](data-sources--http_loadbalancer--reference--group-017.md#canonical-1333110131232221-2013320132121202-0222200220111222-2233011322201031-0130313211123122-1130020301221012-2110003220113131-0210200232333302)
- [enable_ip_reputation](data-sources--http_loadbalancer--reference--group-018.md#canonical-1303123101122023-3103220103322001-2320130011231101-2030202013121330-2011333132210133-2222302130220303-2222103112211003-3201100223033010)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330323130321022-1011312022300121-0130323013103312-0203031311203302-2120100101023303-2211110102033331-2312123223013030-3010231133332323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_malicious_user_detection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- disable_malicious_user_detection

<a id="canonical-0211120230230313-2232002111020022-0310300303003112-2220212012112012-1100001110323331-2031222300111333-2211222013231322-3233223220220013"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_malicious\_user\_detection, enable\_malicious\_user\_detection; Default:
disable\_malicious\_user\_detection\] Configuration parameter for disable malicious user detection.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [disable_malicious_user_detection](data-sources--http_loadbalancer--reference--group-017.md#canonical-0211120230230313-2232002111020022-0310300303003112-2220212012112012-1100001110323331-2031222300111333-2211222013231322-3233223220220013)
- [enable_malicious_user_detection](data-sources--http_loadbalancer--reference--group-018.md#canonical-2110332331011132-0112220203021031-2310200300312113-0000113031331323-0022111100132113-1010312111223221-0011212101030212-2221110322223112)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311133202211221-1301333010001102-3133200113211313-3133322202132313-0311131220320100-2120210322310223-1301000011201332-0001021230203033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_malware_protection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- disable_malware_protection

<a id="canonical-0110111201020031-2132201201331211-2300321031202012-1220303101131323-1121013011003213-1023200210000133-3212320223100311-3202210232331013"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_malware\_protection, malware\_protection\_settings; Default:
disable\_malware\_protection\] Configuration parameter for disable malware protection. Defaults to
\`map\[\]\`. Server applies default when omitted.

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

- [disable_malware_protection](data-sources--http_loadbalancer--reference--group-017.md#canonical-0110111201020031-2132201201331211-2300321031202012-1220303101131323-1121013011003213-1023200210000133-3212320223100311-3202210232331013)
- [malware_protection_settings](data-sources--http_loadbalancer--reference--group-020.md#canonical-2310131001002223-0110011100012011-0122222211112013-2012303223103012-0223003310301012-3210222301130330-2331303223010332-2213231331220230)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0220321111232132-0223122312102103-3110030000112122-2301210200110112-0032122230032323-2030101000021220-2131032321113013-2012300202230031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_rate_limit` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- disable_rate_limit

<a id="canonical-2302102000331233-3112332122210222-1012101130112301-1003310221300013-3103112033310210-3213311023132331-2133321121033023-2003221013233103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable rate limit. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-1320232313112132-3030232230021301-1013212120230233-2020312000033331-1112112213013012-2121100232113103-2130221232312032-0300003230220011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_threat_mesh` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- disable_threat_mesh

<a id="canonical-3001133133030023-2023130020022313-0130012022020101-0300023133201122-3113113333022001-3321303132222330-3121202233022020-0222222033010212"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_threat\_mesh, enable\_threat\_mesh; Default: disable\_threat\_mesh\] Enable this
option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [disable_threat_mesh](data-sources--http_loadbalancer--reference--group-017.md#canonical-3001133133030023-2023130020022313-0130012022020101-0300023133201122-3113113333022001-3321303132222330-3121202233022020-0222222033010212)
- [enable_threat_mesh](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202002102313212-2121133022323302-2332320333313223-1312303022323030-3030133001022000-0320122330221310-0003032232310211-2110113111030212)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013212331311311-1200011021032001-3302132230121201-1322113202300110-3312032230332123-2310111203203000-0113022033313023-0023311222003010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_trust_client_ip_headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- disable_trust_client_ip_headers

<a id="canonical-2300232031100033-0133033320231310-1220022233130331-0310113203313012-0011110010013210-1123331022102110-3033013311232131-1012311131312033"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_trust\_client\_ip\_headers, enable\_trust\_client\_ip\_headers; Default:
disable\_trust\_client\_ip\_headers\] Enable this option. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

- [disable_trust_client_ip_headers](data-sources--http_loadbalancer--reference--group-017.md#canonical-2300232031100033-0133033320231310-1220022233130331-0310113203313012-0011110010013210-1123331022102110-3033013311232131-1012311131312033)
- [enable_trust_client_ip_headers](data-sources--http_loadbalancer--reference--group-018.md#canonical-2200022210121310-3030222023303302-1130102212110230-2323302120331110-2331233111002022-1020110033203002-0231333331120233-3013121212303110)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213210303223130-0120021132231012-2321001110211123-3311211203301010-3202011221331212-0333122303302032-2211232123002030-1222233220133003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_waf` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- disable_waf

<a id="canonical-2303313130312323-1302023203002133-3221130001122233-2313031223322333-3220233303102000-3110223110032023-1222130011330312-2333001030333113"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable waf. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-0312023203322213-3333031332332121-1301313200123121-1211123221033212-0103310121032331-0222122033112101-0230103313330001-3300110200032222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `do_not_advertise` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- do_not_advertise

<a id="canonical-0021302010203112-0300133233023002-1233200231022231-0002221120200231-1313312000123101-1111103332203311-2120010322013122-2021313130023120"></a>

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

<a id="canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- enable_api_discovery

<a id="canonical-1332310000223132-0300312232102030-2113102020202113-3131230231231122-2032032331222321-3211111300111301-1001021113320210-1322233103230012"></a>

Type: `"single"`. Computed.

Specifies the settings used for API discovery.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_discovery_settings_choice": "[\"custom_api_auth_discovery\",\"default_api_auth_discovery\"]",
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable_learn_from_redirect_traffic\",\"enable_learn_from_redirect_traffic\"]"
}
```

<a id="canonical-1320010130132022-3111012230130220-0133223321021220-0303221122110001-0310030003103212-0321321110013110-2100011200323230-1112011123310101"></a>

### Direct properties for `enable_api_discovery`

- [api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-2120033033320302-3321331220022230-1013201202023011-3310311210033200-2131020333021030-0310030101031120-1330332123123333-3311232313301031): complete subsection reference.

- [api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-017.md#canonical-2203231110333303-0020330220102233-1212031302200323-2010132310300200-2323303123313032-3023312023010120-1132113021000231-1003103022033213): complete subsection reference.

- [custom_api_auth_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3001200131212000-3321100200121300-1113221233313122-2013110303220031-1120101002332231-3020123311233011-0203030310301322-1223331233311201): complete subsection reference.

- [default_api_auth_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-2332023220311231-2022101012211013-2010123201123233-1232120333012130-3013001110130232-3300111322230100-2212222320203322-0333122112300110): complete subsection reference.

- [disable_learn_from_redirect_traffic](data-sources--http_loadbalancer--reference--group-017.md#canonical-3012103000100232-1333022032021133-3322020200231210-3020203332310212-3211011202000020-2113003311123103-2223221301000003-3321311331100011): complete subsection reference.

- [discovered_api_settings](data-sources--http_loadbalancer--reference--group-017.md#canonical-2123111312022000-2333133110203330-0220223302200110-2202032222101322-1202332001000200-3311323320022011-1211213100322320-1233320302231323): complete subsection reference.

- [enable_learn_from_redirect_traffic](data-sources--http_loadbalancer--reference--group-018.md#canonical-3132323100013201-3210103000323202-0120303213201211-3200221003122313-0313121001120132-1231133311121230-1021201002210111-3031132030331113): complete subsection reference.

<a id="canonical-2120033033320302-3321331220022230-1013201202023011-3310311210033200-2131020333021030-0310030101031120-1330332123123333-3311232313301031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- enable_api_discovery.api_crawler

<a id="canonical-1322000111102031-1333001120230130-1130300110230112-0101022232213011-1110031011122210-3013323330232012-2021022232202031-1311110302230310"></a>

Type: `"single"`. Computed.

API Crawling. API Crawler message.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_crawler": "[\"api_crawler_config\",\"disable_api_crawler\"]"
}
```

<a id="canonical-1212111030322123-3331030000120032-1032011030121120-0333103333001030-3111131333031100-1203100112303033-1323120212313123-0312110032303333"></a>

### Direct properties for `enable_api_discovery.api_crawler`

- [api_crawler_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-3302012213221322-2213232321020013-2311022021102331-3213331003302311-0111312333111331-0201303113011111-1203101013301010-2333010213320330): complete subsection reference.

- [disable_api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-0020030233331233-0311023322201313-0210323332213013-1020031010101000-1103231033212321-3330202022010112-1011220003103020-0021210313223203): complete subsection reference.

<a id="canonical-3302012213221322-2213232321020013-2311022021102331-3213331003302311-0111312333111331-0201303113011111-1203101013301010-2333010213320330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-2120033033320302-3321331220022230-1013201202023011-3310311210033200-2131020333021030-0310030101031120-1330332123123333-3311232313301031)
- enable_api_discovery.api_crawler.api_crawler_config

<a id="canonical-2310203123111231-1200301212312213-0121101330121331-2222212211032323-1001010311332001-2213011311210030-3100232030110232-3333303102310203"></a>

Type: `"single"`. Computed.

Crawler Configure.

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

<a id="canonical-0132001300220033-3210000130201003-2031313003132121-3223202203011011-2231333002213120-3012123120110000-0022112022111301-1010001332320321"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config`

- [domains](data-sources--http_loadbalancer--reference--group-017.md#canonical-3220113202030100-2131230032122103-2300111203123330-3010310303302300-3030132301121032-0121220020013222-0310213223311223-3100212120333333): complete subsection reference.

<a id="canonical-3220113202030100-2131230032122103-2300111203123330-3010310303302300-3030132301121032-0121220020013222-0310213223311223-3100212120333333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-2120033033320302-3321331220022230-1013201202023011-3310311210033200-2131020333021030-0310030101031120-1330332123123333-3311232313301031)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-3302012213221322-2213232321020013-2311022021102331-3213331003302311-0111312333111331-0201303113011111-1203101013301010-2333010213320330)
- enable_api_discovery.api_crawler.api_crawler_config.domains

<a id="canonical-1302100002121213-3002032212030320-0213110123133111-1332223213003033-2203320201213211-3022302223133000-2200331220030233-1223023123310011"></a>

Type: `"list"`. Computed.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-2302011303303012-1210323030232103-3032001113133030-3311213111132031-2322332312213303-1211110211213001-0013001223221201-3201230212130303"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains`

<a id="canonical-2213011313010313-1102301230321031-2310301100231112-0031233113310001-2320221123211311-2201330020003312-3130001123120331-3022231303020231"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.domain` property

Type: `"string"`. Computed.

Select the domain to execute API Crawling with given credentials.

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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](data-sources--http_loadbalancer--reference--group-017.md#canonical-1220321030133321-2111321021231121-3122312233113211-1303110020310302-2333202113132200-0212322232113011-2333230330333300-0001231102321300): complete subsection reference.

<a id="canonical-1220321030133321-2111321021231121-3122312233113211-1303110020310302-2333202113132200-0212322232113011-2333230330333300-0001231102321300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-2120033033320302-3321331220022230-1013201202023011-3310311210033200-2131020333021030-0310030101031120-1330332123123333-3311232313301031)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-3302012213221322-2213232321020013-2311022021102331-3213331003302311-0111312333111331-0201303113011111-1203101013301010-2333010213320330)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-017.md#canonical-3220113202030100-2131230032122103-2300111203123330-3010310303302300-3030132301121032-0121220020013222-0310213223311223-3100212120333333)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="canonical-0303230201312102-3120330203112332-1302010331101221-1332033301133303-2131110002101023-2311023212132203-3100311212232230-1310132222000002"></a>

Type: `"single"`. Computed.

Configuration parameter for simple login.

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

<a id="canonical-2300010100111022-1332131302112321-1220123200311001-2300000322210221-0131211312333033-3300322333311011-3330223023101201-2100233301311212"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login`

- [password](data-sources--http_loadbalancer--reference--group-017.md#canonical-0022312300202303-0230312323013100-2120001013300320-3103322220032130-3021121210013310-2023330321312111-0220132133023111-1310223122222111): complete subsection reference.

<a id="canonical-3203031330033233-0021123010303231-1333110012031013-2212211102030123-1322030300002210-1000020210010203-2302101113321033-2303010313113122"></a>

<a id="canonical-0300123101023233-0120020012201030-1031012023211013-0101131301310020-2011223333000033-3211022300221330-2331120002033100-0033330211303132"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.user` property

Type: `"string"`. Computed.

Enter the username to assign credentials for the selected domain to crawl.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0022312300202303-0230312323013100-2120001013300320-3103322220032130-3021121210013310-2023330321312111-0220132133023111-1310223122222111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-2120033033320302-3321331220022230-1013201202023011-3310311210033200-2131020333021030-0310030101031120-1330332123123333-3311232313301031)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-3302012213221322-2213232321020013-2311022021102331-3213331003302311-0111312333111331-0201303113011111-1203101013301010-2333010213320330)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-017.md#canonical-3220113202030100-2131230032122103-2300111203123330-3010310303302300-3030132301121032-0121220020013222-0310213223311223-3100212120333333)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--reference--group-017.md#canonical-1220321030133321-2111321021231121-3122312233113211-1303110020310302-2333202113132200-0212322232113011-2333230330333300-0001231102321300)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="canonical-1211230102011110-0100211312132312-1313001200203110-2122101201331220-2202011020223203-0321300301211310-3233221030303022-2230332321210131"></a>

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

<a id="canonical-0313211300321133-1322113330320030-2333030030120013-1002103222031223-0033133100111330-1030212333112322-2020321001010103-1323102023031013"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password`

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-017.md#canonical-3233213120131020-1123231222303010-2011122231031131-3303031331002313-0000131233222202-0010310211233212-3312111130212301-2203200203233300): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-017.md#canonical-1133010120121330-2231312313201313-3023131010323130-3233223211031323-0313110303033011-2312303203301112-3111222030131211-3320031023101131): complete subsection reference.

<a id="canonical-3233213120131020-1123231222303010-2011122231031131-3303031331002313-0000131233222202-0010310211233212-3312111130212301-2203200203233300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-2120033033320302-3321331220022230-1013201202023011-3310311210033200-2131020333021030-0310030101031120-1330332123123333-3311232313301031)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-3302012213221322-2213232321020013-2311022021102331-3213331003302311-0111312333111331-0201303113011111-1203101013301010-2333010213320330)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-017.md#canonical-3220113202030100-2131230032122103-2300111203123330-3010310303302300-3030132301121032-0121220020013222-0310213223311223-3100212120333333)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--reference--group-017.md#canonical-1220321030133321-2111321021231121-3122312233113211-1303110020310302-2333202113132200-0212322232113011-2333230330333300-0001231102321300)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--http_loadbalancer--reference--group-017.md#canonical-0022312300202303-0230312323013100-2120001013300320-3103322220032130-3021121210013310-2023330321312111-0220132133023111-1310223122222111)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

<a id="canonical-2222132031333130-2122222223320212-0231212021333003-2321203231013231-2320231232133331-3223231010333011-0303332003232321-3330322221213201"></a>

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

<a id="canonical-3330300001002013-0113111121232233-0321031002112103-3320133311230120-2332032013010022-1020231120233122-1333121322101333-0321111123112300"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info`

<a id="canonical-1300013021012322-0130113333320333-3101200302303132-2033030202221031-0211310031230301-2333322112320220-1103320332303302-2200213220113111"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-0122123011111121-3022003011222010-1310321220212110-3331222122011002-0122101200011311-0113013013231323-0002311300122112-1233330201021212"></a>

<a id="canonical-1100230110320122-0021323002323001-3333022011220022-0011130203022031-2101021232230102-3301302310230111-2301201300102302-0110333113332110"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2303112020030322-0001231020222233-1232030102233022-0121313330131322-3032233002123301-0301303000113223-2230200022021232-2132322111203113"></a>

<a id="canonical-1201311210221232-2330012331010210-0311013231323020-3303223033323322-1332101012022212-1223201201033310-2331232131103201-3011110231002210"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.store_provider` property

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

<a id="canonical-1133010120121330-2231312313201313-3023131010323130-3233223211031323-0313110303033011-2312303203301112-3111222030131211-3320031023101131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-2120033033320302-3321331220022230-1013201202023011-3310311210033200-2131020333021030-0310030101031120-1330332123123333-3311232313301031)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-3302012213221322-2213232321020013-2311022021102331-3213331003302311-0111312333111331-0201303113011111-1203101013301010-2333010213320330)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-017.md#canonical-3220113202030100-2131230032122103-2300111203123330-3010310303302300-3030132301121032-0121220020013222-0310213223311223-3100212120333333)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--reference--group-017.md#canonical-1220321030133321-2111321021231121-3122312233113211-1303110020310302-2333202113132200-0212322232113011-2333230330333300-0001231102321300)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--http_loadbalancer--reference--group-017.md#canonical-0022312300202303-0230312323013100-2120001013300320-3103322220032130-3021121210013310-2023330321312111-0220132133023111-1310223122222111)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="canonical-2112232301202213-1300222300313320-0210201322202000-0122211010031221-3100233200030321-3303212132000110-2111112330112333-3011002332310232"></a>

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

<a id="canonical-0033032312232231-2313311302210132-1102323332002023-2213303310332012-3110200010020201-3213302030130202-0112213203002331-3322032013013102"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info`

<a id="canonical-0323213301020322-2120333133222021-1033120330312333-0002300010012203-2302012003311032-0023200031132111-1312003012133230-1102300332301121"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2023212302011113-2010120330232122-1123103123030303-3223301120130011-0302230002201220-3220302021301220-3001023112212113-2322202100110120"></a>

<a id="canonical-1102323131031002-0001321122001110-3201021013130032-2002213323130312-3332133003102123-2101032010033020-3123033230023233-3211023013132012"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info.url` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0020030233331233-0311023322201313-0210323332213013-1020031010101000-1103231033212321-3330202022010112-1011220003103020-0021210313223203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.disable_api_crawler` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-2120033033320302-3321331220022230-1013201202023011-3310311210033200-2131020333021030-0310030101031120-1330332123123333-3311232313301031)
- enable_api_discovery.api_crawler.disable_api_crawler

<a id="canonical-0203222113130200-0231131212201332-1020331311231002-3101212202032133-0023230112013301-3312002313023210-1211110033121333-0023303231310133"></a>

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

<a id="canonical-2203231110333303-0020330220102233-1212031302200323-2010132310300200-2323303123313032-3023312023010120-1132113021000231-1003103022033213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- enable_api_discovery.api_discovery_from_code_scan

<a id="canonical-2101020210000021-2333102302201123-1122330003010112-1101032113022033-0010220113030302-1101111102221002-2333323123030012-3300311121332223"></a>

Type: `"single"`. Computed.

Select codebase and Repositories.

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

<a id="canonical-2310233210333312-0202233103102132-3122320231230023-3313311230332213-0333320020020223-2200222133131311-1002311232231122-1001323223002211"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan`

- [code_base_integrations](data-sources--http_loadbalancer--reference--group-017.md#canonical-1030211010301012-1013000322300110-2302132223232131-1222201210323202-0102230030303233-2031113111010121-1002332233020312-1302102010200200): complete subsection reference.

<a id="canonical-1030211010301012-1013000322300110-2302132223232131-1222201210323202-0102230030303233-2031113111010121-1002332233020312-1302102010200200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-017.md#canonical-2203231110333303-0020330220102233-1212031302200323-2010132310300200-2323303123313032-3023312023010120-1132113021000231-1003103022033213)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="canonical-0012223221312200-2311011130311123-2112313120032033-1121321303102302-0101322231331113-1313321323333121-3012300212201030-2221232012123302"></a>

Type: `"list"`. Computed.

Configuration parameter for codebase integrations.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0333022100112330-2223231203232300-0201311111330131-0222001031311212-2101210002131000-1010310313333211-0121011212002020-3130200232223321"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations`

- [all_repos](data-sources--http_loadbalancer--reference--group-017.md#canonical-3013112300131212-1212333112010120-1221130311112321-0233230010033211-1121021222020003-0121233123020013-1222012102131231-1003323223330000): complete subsection reference.

- [code_base_integration](data-sources--http_loadbalancer--reference--group-017.md#canonical-1332213102132331-1201231312011223-2033132100001111-1332203112102311-3020012020023013-1201000221022001-1132033121211322-0300023022121330): complete subsection reference.

- [selected_repos](data-sources--http_loadbalancer--reference--group-017.md#canonical-2321021311110321-2231002202230230-3200232012320302-2111311121123201-2233322300121121-0110123311132030-1011312220023000-1303333200312121): complete subsection reference.

<a id="canonical-3013112300131212-1212333112010120-1221130311112321-0233230010033211-1121021222020003-0121233123020013-1222012102131231-1003323223330000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-017.md#canonical-2203231110333303-0020330220102233-1212031302200323-2010132310300200-2323303123313032-3023312023010120-1132113021000231-1003103022033213)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-017.md#canonical-1030211010301012-1013000322300110-2302132223232131-1222201210323202-0102230030303233-2031113111010121-1002332233020312-1302102010200200)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos

<a id="canonical-3131130020311120-0112333110302323-1311123200311230-2231201012023222-3213002112001113-0230000020100033-1301301023333323-1113332201203020"></a>

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

<a id="canonical-1332213102132331-1201231312011223-2033132100001111-1332203112102311-3020012020023013-1201000221022001-1132033121211322-0300023022121330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-017.md#canonical-2203231110333303-0020330220102233-1212031302200323-2010132310300200-2323303123313032-3023312023010120-1132113021000231-1003103022033213)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-017.md#canonical-1030211010301012-1013000322300110-2302132223232131-1222201210323202-0102230030303233-2031113111010121-1002332233020312-1302102010200200)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration

<a id="canonical-1103212003223111-1313001021110013-3212012030203113-0331303210110231-1313233310310112-3232002000232131-2231111323001331-3001131332120201"></a>

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

<a id="canonical-1202330300321331-1303133131102321-1220321333113020-1230200103322100-1122203001223201-0133312101003321-0212222101000111-3002222321223130"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration`

<a id="canonical-0030113110200003-1310203211002221-3230122323133001-3013212111322300-2101313211322121-3211222111010312-3023333102202121-2303123030212310"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.name` property

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

<a id="canonical-1230212220302331-3211102033021302-2111103132223121-0012023030230030-1213110330120310-3032020121310001-2020103112120002-1312303312111322"></a>

<a id="canonical-0123113130131010-3033020112311213-2311201211331103-2011001001312132-2020111021111230-1233000200022130-2323113331220332-3030130032203222"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.namespace` property

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

<a id="canonical-2322032111233332-3211010121100030-1012110202231331-0313012013311313-2222020221330121-3232023232101011-3300132103230221-1331323201133113"></a>

<a id="canonical-1012033023023230-3031232300323331-2122100302121033-3232013132021222-2331133023001030-3322203121010330-2021323030213000-2213033011203200"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.tenant` property

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

<a id="canonical-2321021311110321-2231002202230230-3200232012320302-2111311121123201-2233322300121121-0110123311132030-1011312220023000-1303333200312121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-017.md#canonical-2203231110333303-0020330220102233-1212031302200323-2010132310300200-2323303123313032-3023312023010120-1132113021000231-1003103022033213)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-017.md#canonical-1030211010301012-1013000322300110-2302132223232131-1222201210323202-0102230030303233-2031113111010121-1002332233020312-1302102010200200)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-2320032312022010-3331132210103323-0232120313033333-0113312230212321-2032121003100020-2123003310303010-1133202023032210-2102112302001211"></a>

Type: `"single"`. Computed.

Select which API repositories represent the LB applications.

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

<a id="canonical-1132323323033033-3033003200130312-0100322011000201-0202030102303311-3211113023231333-2011130023030332-2300300322023023-0030330021022331"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos`

<a id="canonical-0011323012203331-2032311000310331-3202122012223300-2333110330021011-2322033322231203-0233333030011203-1211303131133132-0033323031012010"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos.api_code_repo` property

Type: `["list", "string"]`. Computed.

Code repository which contain API endpoints.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3001200131212000-3321100200121300-1113221233313122-2013110303220031-1120101002332231-3020123311233011-0203030310301322-1223331233311201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.custom_api_auth_discovery` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- enable_api_discovery.custom_api_auth_discovery

<a id="canonical-3030011312103210-3330320131011033-0030222112001221-2133302103331323-2113202033123220-0003321312101102-2331200210311003-3000133212012032"></a>

Type: `"single"`. Computed.

API Discovery Advanced Settings. API Discovery Advanced settings.

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

<a id="canonical-2232333031213130-1331102210303103-1303011031202011-2011332013030203-1233020331111022-3131123103311302-1330020320211321-0320122111313203"></a>

### Direct properties for `enable_api_discovery.custom_api_auth_discovery`

- [api_discovery_ref](data-sources--http_loadbalancer--reference--group-017.md#canonical-2012222213213103-0323031013101033-0220233220333033-3221313001311021-1221033213132101-0031210111310323-1021232100131311-3002130331232323): complete subsection reference.

<a id="canonical-2012222213213103-0323031013101033-0220233220333033-3221313001311021-1221033213132101-0031210111310323-1021232100131311-3002130331232323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [enable_api_discovery.custom_api_auth_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3001200131212000-3321100200121300-1113221233313122-2013110303220031-1120101002332231-3020123311233011-0203030310301322-1223331233311201)
- enable_api_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-2013001112102300-3331322223312130-1023303111203033-3112013233000302-3100221022232203-2121010000323103-3123310210332030-3300330232223300"></a>

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

<a id="canonical-2001111213213021-3300010000330223-2233301312101002-3230011022233331-0220331112220201-1322300220222200-1112313112021112-0112013033311310"></a>

### Direct properties for `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref`

<a id="canonical-0231222100201010-0213003203021221-2001233231133320-0203221011121100-1323030133132013-0332121222001113-2332313210211132-2033133000303303"></a>

#### `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.name` property

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

<a id="canonical-0011211300330002-1230010223113333-3332310202330221-2000131122213210-2110100003231312-3132310230022022-3012301010321221-0022010232022000"></a>

<a id="canonical-3323300031311022-2223022130233023-1323331010330023-3223322223303031-1103233001232233-3023230311112101-0020012032320000-1010020211333301"></a>

#### `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.namespace` property

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

<a id="canonical-2211113201132212-2003220302132120-1131012330200220-3302011211010230-3021210033200003-3022212313212230-0311100013013213-1321111122201323"></a>

<a id="canonical-3230311221030200-0232313302332303-3101120101200031-1103333213200323-1110011132200031-2120210101322003-0322033103331311-1211231232133231"></a>

#### `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.tenant` property

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

<a id="canonical-2332023220311231-2022101012211013-2010123201123233-1232120333012130-3013001110130232-3300111322230100-2212222320203322-0333122112300110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.default_api_auth_discovery` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- enable_api_discovery.default_api_auth_discovery

<a id="canonical-1103332303031223-2033311233200011-3323121101020100-3220211333231103-3020032221203201-1112333320211211-2313301311222310-2322120213231000"></a>

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

<a id="canonical-3012103000100232-1333022032021133-3322020200231210-3020203332310212-3211011202000020-2113003311123103-2223221301000003-3321311331100011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.disable_learn_from_redirect_traffic` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- enable_api_discovery.disable_learn_from_redirect_traffic

<a id="canonical-0031233031223010-1132233020221232-3011202122123101-0000002212013121-1323312010013102-2122223132103200-1230302003000001-0301311221020033"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable learn from redirect traffic.

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

<a id="canonical-2123111312022000-2333133110203330-0220223302200110-2202032222101322-1202332001000200-3311323320022011-1211213100322320-1233320302231323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.discovered_api_settings` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- enable_api_discovery.discovered_api_settings

<a id="canonical-1311201022323300-3033102231303223-3011232200001302-1323133312200011-0000123220321013-3311223331201212-0321310113111123-0331220230230223"></a>

Type: `"single"`. Computed.

Discovered API Settings. Configure Discovered API Settings.

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

<a id="canonical-3010101003230312-1221301233333132-2131012002102233-1120301301033311-2012131303331111-3113022120132101-1203012130113223-0030233333001111"></a>

### Direct properties for `enable_api_discovery.discovered_api_settings`

<a id="canonical-3010130013211332-3330120210313101-2202233022303113-0110031322102213-1110231112020111-2211201000200323-1332213222130312-3123111232031310"></a>

#### `enable_api_discovery.discovered_api_settings.purge_duration_for_inactive_discovered_apis` property

Type: `"number"`. Computed.

Inactive discovered API will be deleted after configured duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
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
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```
