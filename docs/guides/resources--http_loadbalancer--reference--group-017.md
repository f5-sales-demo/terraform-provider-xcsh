---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1120023221213120-2123303131313120-2330120122131213-0300330020133120-0212112313233313-0103101002230133-3220031310222302-1320121220000230"></a>

## Direct properties for `default_pool.origin_servers.private_name`

<a id="canonical-2321002022221332-1113231110230323-3330033033002333-3203312303313032-0312000101312312-2023012301203123-0212021222200030-0312110100322311"></a>

### `default_pool.origin_servers.private_name.dns_name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [inside_network](resources--http_loadbalancer--reference--group-017.md#canonical-1012330102231300-3300211120313123-1033000300003311-2033220022003310-3110013230300333-3111000023130032-1032303200312332-0223310111113132): complete subsection reference.

- [outside_network](resources--http_loadbalancer--reference--group-017.md#canonical-0311301122230130-3202322330221100-3010223302023213-3000100231023333-1000112032103002-1132120213013212-3131232033220233-0012222212112021): complete subsection reference.

<a id="canonical-1221000301200322-3330011231323302-0231220220131132-2213333022101012-1301322203313333-2333200033123022-1032031303110323-3101231123320210"></a>

<a id="canonical-1011133100002120-2332011130303320-0322122322130331-2101003201203023-2001223023223120-2220220023031333-0023103203030321-2230303203231000"></a>

### `default_pool.origin_servers.private_name.refresh_interval` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [segment](resources--http_loadbalancer--reference--group-017.md#canonical-3322321310112123-3331203210113301-2130130132133122-2032103130113112-2102312033030203-0213230232130021-0013232233102032-3321020322030233): complete subsection reference.

- [site_locator](resources--http_loadbalancer--reference--group-017.md#canonical-0211232012223111-1001231003113023-1212212001033231-0031131311031121-2101212231112023-0332221013322020-3210130213310212-0001122020323330): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-017.md#canonical-0312130022121113-1033003212331120-0023101102200212-1312013113201020-0203311020213123-0223031101003330-3232310110011201-2312012230001330): complete subsection reference.

<a id="canonical-1012330102231300-3300211120313123-1033000300003311-2033220022003310-3110013230300333-3111000023130032-1032303200312332-0223310111113132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.inside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- default_pool.origin_servers.private_name.inside_network

<a id="canonical-0201300111332021-1302021030321323-3333021213202233-1200021003223030-1022212320212310-3313023211131022-0133111023222130-3313012200323102"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
inside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311301122230130-3202322330221100-3010223302023213-3000100231023333-1000112032103002-1132120213013212-3131232033220233-0012222212112021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.outside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- default_pool.origin_servers.private_name.outside_network

<a id="canonical-1303232330031210-1303120031000032-3121101321212101-3011030103331201-0300021021232131-1323122223113303-0302002200012013-2130332303030222"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
outside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322321310112123-3331203210113301-2130130132133122-2032103130113112-2102312033030203-0213230232130021-0013232233102032-3321020322030233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.segment` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- default_pool.origin_servers.private_name.segment

<a id="canonical-3022021303003031-2123233030302002-2011220010333131-3022300113220132-0221301023133002-3320202022222133-2032023213100133-1133213011213000"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-2211333233120030-2211100212213312-3332210233112313-3131322331213002-0003303012312310-0020121310121332-3013013332323330-1223120222220121"></a>

### Direct properties for `default_pool.origin_servers.private_name.segment`

<a id="canonical-3133220022030122-2001112013033120-0301332133201113-0221312333132323-3003011133211330-2201313321201230-0030200011220110-2201102232303312"></a>

#### `default_pool.origin_servers.private_name.segment.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3131032023101000-1101313122023232-3202002233300101-3212221023331011-0323111231002113-0301232010020211-2223302303302233-2022212302102102"></a>

<a id="canonical-1101302301003120-2223131202302022-1321121221232021-1313100303320312-3003202101303221-0012112131330203-1113100133030032-1031311033033121"></a>

#### `default_pool.origin_servers.private_name.segment.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3113201210120312-0120312220002321-3103203002233112-0010313123233212-3321122300201121-2031201120003100-2303300222102011-1233211012301031"></a>

<a id="canonical-0210030033230132-0320212110023020-3231131002210033-0220211331011113-2031223312333322-3321031323030230-3103131233001213-2003131130303302"></a>

#### `default_pool.origin_servers.private_name.segment.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0211232012223111-1001231003113023-1212212001033231-0031131311031121-2101212231112023-0332221013322020-3210130213310212-0001122020323330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.site_locator` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- default_pool.origin_servers.private_name.site_locator

<a id="canonical-0302202221122101-1020001003021233-0032023122222011-1221020112122103-2132010320021102-2120111200010321-0300002230031221-2230323332022001"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121100023202223-3112023213031201-1100112121113300-2230320101001233-0032230233231231-2203221123030003-1222202121303021-0213021322321011"></a>

### Direct properties for `default_pool.origin_servers.private_name.site_locator`

- [site](resources--http_loadbalancer--reference--group-017.md#canonical-1332313011120202-1010002020031311-2303320320310102-1231133210103131-2100011110331232-2120331112022210-3101220212213003-1023322200303033): complete subsection reference.

- [virtual_site](resources--http_loadbalancer--reference--group-017.md#canonical-3130200003012212-2200212123310111-3112101032311330-2012320313331221-2010131132322310-0221123031111332-1110002013020130-0321020223010200): complete subsection reference.

<a id="canonical-1332313011120202-1010002020031311-2303320320310102-1231133210103131-2100011110331232-2120331112022210-3101220212213003-1023322200303033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.site_locator.site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [default_pool.origin_servers.private_name.site_locator](resources--http_loadbalancer--reference--group-017.md#canonical-0211232012223111-1001231003113023-1212212001033231-0031131311031121-2101212231112023-0332221013322020-3210130213310212-0001122020323330)
- default_pool.origin_servers.private_name.site_locator.site

<a id="canonical-3011232212001132-0012120012320321-3203323001312230-1332203313101030-0122122310023210-1113102111312203-1013001220102211-1123030312100101"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212303123122001-1023332133231033-2300333001022031-3111000211232303-1200231012031331-1130102201120013-2230132022303333-3223201012001201"></a>

### Direct properties for `default_pool.origin_servers.private_name.site_locator.site`

<a id="canonical-1112010021112123-3201102202302321-0220032131233133-0310202121331023-0100030103201033-0203230023013110-0122301130111133-2332022121003203"></a>

#### `default_pool.origin_servers.private_name.site_locator.site.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1202300301222100-2012022221333013-0130102030301210-3211110031202011-1112130231303231-3301233000303100-3021130001013120-0101123122323123"></a>

<a id="canonical-1212202213030213-3313100131131111-1122202323000010-3210121332000221-1300331302200220-1100100300300230-0333002012222133-3201203231003103"></a>

#### `default_pool.origin_servers.private_name.site_locator.site.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3321233211033221-1112031222131010-0111133332112121-1020200133103212-0010332012112320-0010120322322131-1312102121302022-0321131212103032"></a>

<a id="canonical-1101211011012322-1311333112031322-0102123032221300-2203333012000131-1002230111211332-1300103213202231-2030303320210301-1301102311101322"></a>

#### `default_pool.origin_servers.private_name.site_locator.site.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3130200003012212-2200212123310111-3112101032311330-2012320313331221-2010131132322310-0221123031111332-1110002013020130-0321020223010200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [default_pool.origin_servers.private_name.site_locator](resources--http_loadbalancer--reference--group-017.md#canonical-0211232012223111-1001231003113023-1212212001033231-0031131311031121-2101212231112023-0332221013322020-3210130213310212-0001122020323330)
- default_pool.origin_servers.private_name.site_locator.virtual_site

<a id="canonical-0302210233131221-3232120302303310-1133033103111201-3300011212221102-0101010023112321-2321222002132013-0223002103201020-2213012313131001"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223221322210212-3111322032302000-2121113112333322-0132022222200010-0213301012301120-2113011120303210-3302201210311012-3310102023110102"></a>

### Direct properties for `default_pool.origin_servers.private_name.site_locator.virtual_site`

<a id="canonical-3032210012121211-0020022121323101-3002112200130113-3321113021123233-0032313032222222-1311120011003331-2323100232302320-3122123332213031"></a>

#### `default_pool.origin_servers.private_name.site_locator.virtual_site.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3200302122000310-3220202232022232-0021023220310201-3133232302030220-1132131001032020-1221023113200231-2133233212113203-2322133101011120"></a>

<a id="canonical-3101303321313233-1332132103232211-2020231332130212-3222100133131020-1123001013303132-1003210112323001-3120322212112013-0123321000123321"></a>

#### `default_pool.origin_servers.private_name.site_locator.virtual_site.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2112200130122133-3333330303023123-2323223000030303-3321121222100313-3331232330003120-1300230212232230-1103221010002100-3031003233322001"></a>

<a id="canonical-0023221331103231-0221002200302031-3112321310313101-1332310112001310-0333200331123113-3132133033100020-3113313120332223-2300031200310101"></a>

#### `default_pool.origin_servers.private_name.site_locator.virtual_site.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0312130022121113-1033003212331120-0023101102200212-1312013113201020-0203311020213123-0223031101003330-3232310110011201-2312012230001330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- default_pool.origin_servers.private_name.snat_pool

<a id="canonical-2320301120230101-1321101213212110-0231221300332001-3023123200110103-1221133312100032-3322323102022310-0230013131131311-3113211013332331"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-2201320020130103-2220032303002212-3303002110102321-1112300021103311-2112232130102210-3300020310113320-3222323021311201-1233023333320310"></a>

### Direct properties for `default_pool.origin_servers.private_name.snat_pool`

- [no_snat_pool](resources--http_loadbalancer--reference--group-017.md#canonical-1221311033021110-0022231221001322-3333312321200100-2220310323023132-2323323222132122-3101233013033220-0133113303032131-3110200300020302): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-017.md#canonical-0023201330313033-3012311220001023-3020110330231121-1123002133300333-0230303300110332-2203000201021020-1110113003100210-2301230331311013): complete subsection reference.

<a id="canonical-1221311033021110-0022231221001322-3333312321200100-2220310323023132-2323323222132122-3101233013033220-0133113303032131-3110200300020302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [default_pool.origin_servers.private_name.snat_pool](resources--http_loadbalancer--reference--group-017.md#canonical-0312130022121113-1033003212331120-0023101102200212-1312013113201020-0203311020213123-0223031101003330-3232310110011201-2312012230001330)
- default_pool.origin_servers.private_name.snat_pool.no_snat_pool

<a id="canonical-3010011311112200-0330331122233210-0203123023122111-3231232220230010-1333302213032132-3013220010023111-1320001113233330-1301233110031332"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_snat_pool = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023201330313033-3012311220001023-3020110330231121-1123002133300333-0230303300110332-2203000201021020-1110113003100210-2301230331311013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [default_pool.origin_servers.private_name.snat_pool](resources--http_loadbalancer--reference--group-017.md#canonical-0312130022121113-1033003212331120-0023101102200212-1312013113201020-0203311020213123-0223031101003330-3232310110011201-2312012230001330)
- default_pool.origin_servers.private_name.snat_pool.snat_pool

<a id="canonical-3103122222102210-1012332212123003-1330010023111300-1221023101111210-0011332012303200-1113323033323333-0303222033220103-2302113302130333"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232102210211131-0030103333320023-2110231230311201-1300310202020011-0330120131132113-3302131020301023-3220233333011021-2220002123003111"></a>

### Direct properties for `default_pool.origin_servers.private_name.snat_pool.snat_pool`

<a id="canonical-1012311200321331-1320031233130221-2302010222201100-0020311210310333-3002011300322112-0010103000322223-0323312202323330-0123310223012013"></a>

#### `default_pool.origin_servers.private_name.snat_pool.snat_pool.prefixes` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1123323200112132-2011301210333320-0230333321233220-3120200223212323-0221012302113321-1102011330332121-0221103113230111-0011130112303131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.public_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.public_ip

<a id="canonical-2233020311003121-3121333321021131-3001122003132113-1301222300022020-1120113331012303-0312322233233120-0102201110013213-2020312300030233"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320121003222032-3310311231032211-2212031120013103-1010112113121012-0202231230232310-1231023111303210-2212313121302320-3313002301322323"></a>

### Direct properties for `default_pool.origin_servers.public_ip`

<a id="canonical-1000013300233110-0012120220113212-2010132312303303-3231031102201032-2131001003003211-1012030320130230-1011300130213021-1200312020010213"></a>

#### `default_pool.origin_servers.public_ip.ip` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3302320302223102-1023033333230010-2222202133303121-2101111202221230-2113030322031221-1332101010000231-1202031111300230-1003313332002300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.public_name` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.public_name

<a id="canonical-1001132011011000-0202223032003200-2201310223322201-2301022130002302-2310300302010222-3101213112032003-2323032031122031-2013322101233301"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
public_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302123012033012-3230112200032102-1222323103130102-1002102230210231-1301002131301222-2033203301323333-1103323003201233-1100223202233321"></a>

### Direct properties for `default_pool.origin_servers.public_name`

<a id="canonical-0333302200003211-1210112202300202-0220113030021222-2000123111031031-3110333222010230-3000302321313102-2001003200333210-3202302130002112"></a>

#### `default_pool.origin_servers.public_name.dns_name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0023123303332312-3122011222213000-0211231320220230-2331322222311001-1232100123220010-3223120300221100-2221002333213110-0212123331231301"></a>

<a id="canonical-1320331003032222-1030022331322123-1013130302320223-3130122222212223-3232002013020003-0311202330031201-1110201231220202-3032103030211312"></a>

#### `default_pool.origin_servers.public_name.refresh_interval` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0022310013130030-2111303220210213-2020201222032233-1120103131300012-1111203033323012-1021113330333233-3121232112020030-0200301203202122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.vn_private_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.vn_private_ip

<a id="canonical-2210121123312130-3103200230220021-0013001213201023-0331001313021131-0203202312200302-2012112310012022-1230221032133132-3123200303232133"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
vn_private_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203231103120212-1232312001313221-2223333132133133-2303213133032222-0300130320322121-0221221001112112-2220123220301330-2221120101331320"></a>

### Direct properties for `default_pool.origin_servers.vn_private_ip`

<a id="canonical-1101003013130331-3001011200313212-0300323212213213-2311110210321110-3100023320302121-0303100212220002-3112330322123123-0310113311002322"></a>

#### `default_pool.origin_servers.vn_private_ip.ip` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [virtual_network](resources--http_loadbalancer--reference--group-017.md#canonical-0022111032132322-0332003132203223-2021223302101220-2331130303203021-3023123313013130-3330101201022111-1223020213011322-1133103323103121): complete subsection reference.

<a id="canonical-0022111032132322-0332003132203223-2021223302101220-2331130303203021-3023123313013130-3330101201022111-1223020213011322-1133103323103121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.vn_private_ip.virtual_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.vn_private_ip](resources--http_loadbalancer--reference--group-017.md#canonical-0022310013130030-2111303220210213-2020201222032233-1120103131300012-1111203033323012-1021113330333233-3121232112020030-0200301203202122)
- default_pool.origin_servers.vn_private_ip.virtual_network

<a id="canonical-2222011203123233-1231300123133010-1302300210111000-2103002203232133-2031123112002200-3321113323303113-0321013120333312-0303331013203100"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-3031020010100311-3021333303030010-0211302300122020-0020232100111200-3233111202131312-0102002210003132-3132231212300130-2313212300331323"></a>

### Direct properties for `default_pool.origin_servers.vn_private_ip.virtual_network`

<a id="canonical-0111132101010223-0301022232122322-1311111222032122-1031021203103310-1302233303013320-3100120111011302-1323033103102332-0030131231201332"></a>

#### `default_pool.origin_servers.vn_private_ip.virtual_network.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3321123122101032-3011112312021012-3121100022223300-3000220113012023-2223003010233133-0012211113000110-2101303132201031-2012321022013003"></a>

<a id="canonical-2231330321103330-3222021321030102-1031333231321010-0213132030313201-0012322010000322-3213133123030021-2100331330221210-0333212213212213"></a>

#### `default_pool.origin_servers.vn_private_ip.virtual_network.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3113310013103302-1211100331111300-2013123333022101-0202332000302211-3333112003322321-0131111103312020-3021011210211103-3130003323100213"></a>

<a id="canonical-1022211103320111-1101210012120010-2013130133302110-3323121231031320-3131300111221203-2213220101023020-2322023222200012-3312312101032232"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2101033313001230-3100221301111002-2011003233203011-1100130312221121-1002231322001133-0002023210303033-3312032213131120-2332031232122312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.vn_private_name` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.vn_private_name

<a id="canonical-2220132331000203-3110331213121213-1230131221203200-3213321231111321-1213102122203002-2020001233110101-1102020200321030-2220133333332113"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
vn_private_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223011300100011-2120100033322020-0001033103022102-0201100331322113-2113030033113200-2212021322300120-2332020021230023-2131100310313012"></a>

### Direct properties for `default_pool.origin_servers.vn_private_name`

<a id="canonical-1311303101122130-3300123300122221-3301210122222203-1012132302312301-2112320111001300-2230213032023133-0132233212221123-1131023230213030"></a>

#### `default_pool.origin_servers.vn_private_name.dns_name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [private_network](resources--http_loadbalancer--reference--group-017.md#canonical-0203020320000112-0102231033232010-1223102323000102-2100120021102020-0320320120202210-2013203021122133-2011000000301030-2011300203301111): complete subsection reference.

<a id="canonical-0203020320000112-0102231033232010-1223102323000102-2100120021102020-0320320120202210-2013203021122133-2011000000301030-2011300203301111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.vn_private_name.private_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.vn_private_name](resources--http_loadbalancer--reference--group-017.md#canonical-2101033313001230-3100221301111002-2011003233203011-1100130312221121-1002231322001133-0002023210303033-3312032213131120-2332031232122312)
- default_pool.origin_servers.vn_private_name.private_network

<a id="canonical-2331213320121203-1220232222333230-1322222102013220-1112010332020110-1020313232110022-0133002120120112-0112031123001302-2113033022000322"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
private_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023200301022111-0203023133002333-2011021330101000-2021131020301130-1112032312131203-3013023122103020-1012320322113233-1231122120110233"></a>

### Direct properties for `default_pool.origin_servers.vn_private_name.private_network`

<a id="canonical-0011022111002210-1100033313003213-0132120333103120-2203231211201231-3233131011300303-1022121012131113-3101000020303213-0111213023203223"></a>

#### `default_pool.origin_servers.vn_private_name.private_network.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1001300112001202-2021132333113100-0200010200333013-3113203220203232-3013330013000200-2221122112103213-3100220033332133-3312211111203200"></a>

<a id="canonical-1123033330210120-1101010110111212-1021210130110330-0302210113130130-0003312031222022-1101021300312200-3221222121213200-2130003330022120"></a>

#### `default_pool.origin_servers.vn_private_name.private_network.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3223331230003012-0130313330232130-0230233311212200-3022210100300310-1320323311020300-0120030112030103-2030321023330030-1033123110000302"></a>

<a id="canonical-3311033330330213-0100132332301200-2322122221133201-3320022301101332-3022210133122233-2101033233323100-3022102030201123-2332322302000102"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1232100333033333-0221010121032103-2010033203301010-1223212021203320-2001102032013210-0333301222330003-1231200013113223-2010211132110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.same_as_endpoint_port` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.same_as_endpoint_port

<a id="canonical-2033213221230001-2133023101322313-1230200120100220-1012323122020200-2222200231023032-1012202202131121-3102330121032032-1030231322320113"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
same_as_endpoint_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112003010300102-0001210211031011-0021210310211031-2120013323123301-1203112022230211-2311310013202331-2323031211322202-1321003211023230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.upstream_conn_pool_reuse_type` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.upstream_conn_pool_reuse_type

<a id="canonical-0123322202101000-1321100330232001-2300223133032122-2100000323122112-0020213022331320-3213210102203003-0120120302223030-2210302332110313"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
upstream_conn_pool_reuse_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302310033303032-3323330200011311-3000120333200003-0010033311321113-1221202201023033-0110230231300033-0121123233023220-3122102133100002"></a>

### Direct properties for `default_pool.upstream_conn_pool_reuse_type`

- [disable_conn_pool_reuse](resources--http_loadbalancer--reference--group-017.md#canonical-3022003332112203-3313121000200322-1020012101323212-1011010321303313-0220213231201133-3123213123103002-0213302001100112-0013020212002000): complete subsection reference.

- [enable_conn_pool_reuse](resources--http_loadbalancer--reference--group-017.md#canonical-0122233311301230-2033102130212202-1303001033113312-2303130331113333-0121303320101230-1213010322331113-0123121210000310-3310332303220133): complete subsection reference.

<a id="canonical-3022003332112203-3313121000200322-1020012101323212-1011010321303313-0220213231201133-3123213123103002-0213302001100112-0013020212002000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.upstream_conn_pool_reuse_type](resources--http_loadbalancer--reference--group-017.md#canonical-1112003010300102-0001210211031011-0021210310211031-2120013323123301-1203112022230211-2311310013202331-2323031211322202-1321003211023230)
- default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-1033210212033123-0322132311231131-2003001201100033-1222031301130021-3020120221133033-0133021221203001-1300033210001320-3322032323211021"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_conn_pool_reuse = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122233311301230-2033102130212202-1303001033113312-2303130331113333-0121303320101230-1213010322331113-0123121210000310-3310332303220133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.upstream_conn_pool_reuse_type](resources--http_loadbalancer--reference--group-017.md#canonical-1112003010300102-0001210211031011-0021210310211031-2120013323123301-1203112022230211-2311310013202331-2323031211322202-1321003211023230)
- default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-3113313020022133-0330300231330231-3032212110201123-0231101332023132-3121312031113003-3020230103323212-3303122210032002-3213330300132320"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enable_conn_pool_reuse = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.use_tls

<a id="canonical-2332213023031332-1013323323031210-0100123112302011-2330013230103010-0102132311330320-2010003302003033-2313120303330331-0312102122123221"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103122312033133-0232020122221200-3322322232313322-3012123330332032-0000121120031111-1021333002233122-0311211200123330-1203031110312330"></a>

### Direct properties for `default_pool.use_tls`

- [default_session_key_caching](resources--http_loadbalancer--reference--group-017.md#canonical-2202223033321103-1023313022220032-1100110331023310-2000233222211323-0101221310312021-3102332013202020-0101210122012110-2103210131213301): complete subsection reference.

- [disable_session_key_caching](resources--http_loadbalancer--reference--group-017.md#canonical-1013331020231333-0202203320300021-1232330310100320-2300031020020203-1123330312131323-3001113200312303-0122030002201222-3233033231132121): complete subsection reference.

- [disable_sni](resources--http_loadbalancer--reference--group-017.md#canonical-2300213333200132-0123123213203223-0200321201003332-1133131212130322-0120022102301322-3132311102333002-1200211200102323-3233332031012120): complete subsection reference.

<a id="canonical-1213330013021023-0213323002023222-2310113003121231-0010330232301332-1211121220311200-0331022332003023-0330221202302201-1302013110331311"></a>

<a id="canonical-2333032221123011-1312320102032030-0222130012112100-0300313100213112-3310300033310112-3300102003222320-0201120203223113-2323001120313230"></a>

#### `default_pool.use_tls.max_session_keys` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [no_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-0303331032002210-0312010001012120-1330112001013310-0010320010222322-2132130013000332-1231131311012023-3021102123223331-2330222230010032): complete subsection reference.

- [skip_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-0303332132110311-1223303331210202-2030002213310231-2101021031103002-3222212230031331-2213230130330002-3303131031031310-2010310111301030): complete subsection reference.

<a id="canonical-3203103222000233-2001032322130211-2232220032202333-3332301033030323-2013113331033102-3033013031113121-0111203103331320-1122023110231123"></a>

<a id="canonical-0001021102223102-1002303213321120-2321132020021113-3001122012301201-3312032111232212-2323012221313122-2233031230223221-1213112022113001"></a>

#### `default_pool.use_tls.sni` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112): complete subsection reference.

- [use_host_header_as_sni](resources--http_loadbalancer--reference--group-017.md#canonical-1322020121333230-3003002212033100-2332300001021121-1102021010223121-3131123312322320-1020132011323230-0202022322031312-3313322321233013): complete subsection reference.

- [use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031): complete subsection reference.

- [use_mtls_obj](resources--http_loadbalancer--reference--group-017.md#canonical-2011310303210021-1222202221103110-1130033232102230-1131030100322101-0032220031011132-0202012301223101-3013101110032002-0102022212323012): complete subsection reference.

- [use_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-3302122303122133-0323233203310323-3012112030002300-2003211111010032-2031330012231010-1011303111010020-0113222311123003-1002223100220030): complete subsection reference.

- [volterra_trusted_ca](resources--http_loadbalancer--reference--group-017.md#canonical-3321010323221300-3221232231012333-1000122333111322-3013010233100302-2311021323120212-1121301113331022-3222031002320022-3101011201003023): complete subsection reference.

<a id="canonical-2202223033321103-1023313022220032-1100110331023310-2000233222211323-0101221310312021-3102332013202020-0101210122012110-2103210131213301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.default_session_key_caching` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.default_session_key_caching

<a id="canonical-0121200021230013-1210110220011333-2010002013003130-0332331033121331-3210220321103133-0003331123200131-0031332131222032-3333001233022131"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
default_session_key_caching = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013331020231333-0202203320300021-1232330310100320-2300031020020203-1123330312131323-3001113200312303-0122030002201222-3233033231132121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.disable_session_key_caching` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.disable_session_key_caching

<a id="canonical-3133021332013103-1320013333102132-1221130111322301-2103212230210211-0003102122023023-0211011210211020-1221123121131000-3233100211221103"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_session_key_caching = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300213333200132-0123123213203223-0200321201003332-1133131212130322-0120022102301322-3132311102333002-1200211200102323-3233332031012120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.disable_sni` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.disable_sni

<a id="canonical-2220212132201313-2030112331313111-2010033223000013-0212133231032022-3210020311230301-3211012122013333-3113202200333123-2122133012213030"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_sni = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303331032002210-0312010001012120-1330112001013310-0010320010222322-2132130013000332-1231131311012023-3021102123223331-2330222230010032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.no_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.no_mtls

<a id="canonical-1332101130102002-2013003030331203-2321313313211121-1313313112103031-3132102123231203-1113032031122220-1011112032010133-3230233321301131"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303332132110311-1223303331210202-2030002213310231-2101021031103002-3222212230031331-2213230130330002-3303131031031310-2010310111301030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.skip_server_verification` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.skip_server_verification

<a id="canonical-1212002212011012-3301133312303113-1013031220110323-3303223000111220-3012313232202033-2132132223120222-0322311011312202-0002032201302312"></a>

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
skip_server_verification = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.tls_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.tls_config

<a id="canonical-3200322132320010-0012131232103001-1322321033121031-0031033101302232-1313330021103112-3102301211222010-1001032231012121-2320123131013313"></a>

Type: `"object"`. single nested block, Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231013032332023-2323012321201032-0113233032310002-0121100300001020-2201301333031211-0032310301110012-0233111121030003-2131203211113222"></a>

### Direct properties for `default_pool.use_tls.tls_config`

- [custom_security](resources--http_loadbalancer--reference--group-017.md#canonical-2200333031323122-3012011201002322-3130330223013331-1233310103131031-2030012001303132-0303212231120001-1201102011200321-1010332223231233): complete subsection reference.

- [default_security](resources--http_loadbalancer--reference--group-017.md#canonical-0100203223333331-3222113310331031-3023321131313312-3310033221020010-1313130223033303-1002200231311002-3020333203131231-3220021002103110): complete subsection reference.

- [low_security](resources--http_loadbalancer--reference--group-017.md#canonical-0303100332303222-3233213233330122-3131223311022000-1321203230112312-0031221232100122-2031312301300013-2212123210100030-0103022302311131): complete subsection reference.

- [medium_security](resources--http_loadbalancer--reference--group-017.md#canonical-0222121100303332-0023323120332332-0330001011121000-1210020031030130-2220031231123120-0120333002033001-2221322033313322-0131003300131300): complete subsection reference.

<a id="canonical-2200333031323122-3012011201002322-3130330223013331-1233310103131031-2030012001303132-0303212231120001-1201102011200321-1010332223231233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- default_pool.use_tls.tls_config.custom_security

<a id="canonical-1220133112231103-3133202012022230-0331123031020012-3112010330311112-0202101003110102-1213221301210002-0300333021103032-1121322030113003"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102220012202123-3321223213022213-3301212212010123-0333200120320230-1033011213022010-3321101213333321-1332021300021123-1310232010223020"></a>

### Direct properties for `default_pool.use_tls.tls_config.custom_security`

<a id="canonical-2302131311321310-2313111213002113-0100103000220313-2320303332300231-2303022003032033-3132133303200303-2222210003100101-0321130323100122"></a>

#### `default_pool.use_tls.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3021112011223320-0021132033313301-1023310221110301-3311031233211110-1301132310331000-2010122332103110-2133232211213330-1203020030100311"></a>

<a id="canonical-3112003212202120-3021113011313011-0233233221011133-1011020231013310-1231200131020013-2220221123130033-0303332321101210-0210131310221023"></a>

#### `default_pool.use_tls.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

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

<a id="canonical-1301111312203121-3301313013110313-3101201213000210-3230133310320030-0031223213103012-1130123011223330-1010032130303022-3332012300313323"></a>

<a id="canonical-2002032131210120-1031100130223013-3020001330000232-0222311122301100-0323212221310101-3002323113221332-3102230320032000-2210233221322203"></a>

#### `default_pool.use_tls.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

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

<a id="canonical-0100203223333331-3222113310331031-3023321131313312-3310033221020010-1313130223033303-1002200231311002-3020333203131231-3220021002103110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- default_pool.use_tls.tls_config.default_security

<a id="canonical-2001121231113323-2202302022312013-1331012232331221-2323103111030031-1130222031133131-0113313212323213-0300232212322333-3320000103001121"></a>

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
default_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303100332303222-3233213233330122-3131223311022000-1321203230112312-0031221232100122-2031312301300013-2212123210100030-0103022302311131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- default_pool.use_tls.tls_config.low_security

<a id="canonical-0130332222331002-0011221102013332-0301322031321302-3022120020233011-2302221311111002-3302121031031122-3230332233003002-0003201330112311"></a>

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
low_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222121100303332-0023323120332332-0330001011121000-1210020031030130-2220031231123120-0120333002033001-2221322033313322-0131003300131300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- default_pool.use_tls.tls_config.medium_security

<a id="canonical-1000302303231022-3322010013221311-0123300200102313-2202220003120212-3111311222313132-2222122030223012-1312020130221101-3112232313213203"></a>

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
medium_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322020121333230-3003002212033100-2332300001021121-1102021010223121-3131123312322320-1020132011323230-0202022322031312-3313322321233013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_host_header_as_sni` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.use_host_header_as_sni

<a id="canonical-0313031012231211-1002132313113203-0121133132332010-3323033303133001-2320223110111303-0231130011133120-1311031200103022-3301103201002302"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
use_host_header_as_sni = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.use_mtls

<a id="canonical-0012033332103102-2000112211213221-2202123011211022-1112211212300213-3323011221300011-1203333113303113-3102003221010000-1301333232022310"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302122203211220-2121332323211221-0102313012221101-0033133321201112-1003001100120130-1000100220210213-2313103302202011-2133103030320110"></a>

### Direct properties for `default_pool.use_tls.use_mtls`

- [tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023): complete subsection reference.

<a id="canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- default_pool.use_tls.use_mtls.tls_certificates

<a id="canonical-2212133130313131-0212312110312103-1001211323113313-0122031120232010-2021312001202331-3111320320020322-2211331321031122-1112301000111222"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322332012030213-2101321010323011-0323230010300033-0311312032020030-1220103001011021-3023011030213022-2013210133320030-3112131223100023"></a>

### Direct properties for `default_pool.use_tls.use_mtls.tls_certificates`

- [blindfold](resources--http_loadbalancer--reference--group-017.md#canonical-1231203312231012-2130222312112113-3003223202011032-0301001133222102-3232033311312312-1112332102012033-3302330222313122-2221203330030021): complete subsection reference.

<a id="canonical-0020200123311232-3102103202333331-0002123021121203-3031120233301302-2100323232320100-2020032022301301-3100213212010200-3211120101203113"></a>

<a id="canonical-0103022201121202-1302132303010331-0023323221331221-0322223330213331-3011130023133202-0113120122130133-0233013302312332-3333021001221030"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.certificate_url` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [custom_hash_algorithms](resources--http_loadbalancer--reference--group-017.md#canonical-2020213113300220-2333110233301203-3300222130030231-2130303232320030-2032111112212201-2300222323013202-0110222032000202-1320111112111031): complete subsection reference.

<a id="canonical-2221103223010020-3100323022012230-3032130320011322-2113230031211331-1222220132311223-0211100112030330-0302212101133010-0201131020300123"></a>

<a id="canonical-2033010230112102-3130213100112022-0220203003032213-0310023023321023-0103311120133130-2300033321031122-3303320110231211-0300102213322313"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--http_loadbalancer--reference--group-017.md#canonical-1322333013020002-3320101320322330-1230010223210130-0332001200121133-0020320123030210-0012032023300221-0031101222021121-2210211232313232): complete subsection reference.

- [private_key](resources--http_loadbalancer--reference--group-017.md#canonical-2131202211301100-0331121201212013-2133312320113023-3033313101313002-3322302102012100-0123222110310223-0321031132101002-3021002201003223): complete subsection reference.

- [use_system_defaults](resources--http_loadbalancer--reference--group-017.md#canonical-1112030331010001-1223303203031330-0333010110213321-3202130323302212-0212221311313102-1030203333133001-1312232202322201-0223303001130330): complete subsection reference.

<a id="canonical-1231203312231012-2130222312112113-3003223202011032-0301001133222102-3232033311312312-1112332102012033-3302330222313122-2221203330030021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.blindfold` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- default_pool.use_tls.use_mtls.tls_certificates.blindfold

<a id="canonical-0002002131210020-3320231131002222-1102112030221013-3030001103033022-0010210232000110-2022111212301011-0020300132122233-2031203302113213"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

<a id="canonical-2000132022123301-1133032023032301-2132023132200233-2213033211113203-1303313222011230-3031002330323023-3233023101212101-1330220000130030"></a>

### Direct properties for `default_pool.use_tls.use_mtls.tls_certificates.blindfold`

<a id="canonical-0211103003321210-2312033011111013-1231331300110130-0012121132101011-3203013200033110-2012030213130020-2133222002331112-0122003022012202"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.algorithm` property

Type: `"string"`. Computed.

<a id="canonical-0100321033333233-2301123202223303-0332010211022022-2110032222130321-0020011032012220-2330330231222212-1323101222211000-2310210232210312"></a>

<a id="canonical-2321221210112221-3030203033000201-1330002313310020-2320201033310123-3102022123000321-1102020122300133-1302312322123330-2213013222330330"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.certificate_file` property

Type: `"string"`. Optional.

<a id="canonical-1103211112321103-3123231210311323-0020233022100212-3203021113330211-1131122232111102-0112200020320031-0222123012011212-1100223030320113"></a>

<a id="canonical-2010201223120220-2233212320330021-3212333220011031-1231212332331332-2122031303011002-3101222001122123-3110302322211120-2230203122112303"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.certificate_pem` property

Type: `"string"`. Optional.

<a id="canonical-1320312321101313-2013030032031321-3033003221221010-3020021020011301-0320123032010112-3131211313213013-2000102302210012-2222132130133312"></a>

<a id="canonical-0013011232332102-3232130201032231-0010010111013113-0031003032133123-1322020022112210-3132210012111323-2120301132010203-2113013023211012"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.chain_identity` property

Type: `"string"`. Computed.

<a id="canonical-3133111201102113-3120032331212100-0133323213000222-1322211002333121-0220322021332121-1313022022202013-3213232020303031-2322013302302031"></a>

<a id="canonical-3220313131230230-3201123320210220-3130101332001231-3031302021320332-1210232020001230-0303223232032030-2123203331133031-2101112123020031"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.context_digest` property

Type: `"string"`. Computed.

<a id="canonical-3203323132033000-3330223210121332-2200000321313100-3310110233313212-2122211333000203-0332132103202030-0012113302233133-0322203333032120"></a>

<a id="canonical-3203301103012032-1033133332231100-1013030201122003-3012230011321110-1032121121122132-3300301323120313-2033123232011213-3113032210323311"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.encrypted_location` property

Type: `"string"`. Computed, Sensitive.

<a id="canonical-0212300211020320-1330113032332031-3010013222030211-3300030011122033-0000200220303233-1000013102003103-1310121113201311-0023221211122032"></a>

<a id="canonical-0102330003131322-3001213300122031-1232202222221210-2202221033200301-3110210001101330-1310020323313031-2213113133102212-3012001211323011"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.expires_at` property

Type: `"string"`. Computed.

<a id="canonical-3031322212120022-2302201333021133-2232023332130023-1032330001130032-0303021003220110-0312313320022310-2211021032123300-3103303001033231"></a>

<a id="canonical-3120203232000013-3313333330300033-3210101320213103-2213002330311223-0002201123002131-1002011320000303-1303231020123021-3102110232321201"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.fingerprint` property

Type: `"string"`. Computed.

<a id="canonical-0210103102233320-0331130010013001-3130303102223230-2030120100313312-3110012031312332-3132112202203300-0303121021021301-2312102000312000"></a>

<a id="canonical-1000301013023321-1113132110232013-0222100210302033-0103303132102313-3320220332210102-1331020121030233-0220023231232120-3313211301322000"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.id` property

Type: `"string"`. Optional.

<a id="canonical-1210200320322033-2131022101112131-2033303102200110-1003132033110122-1123022321020231-0211232310320320-1023000002120123-2012121213023002"></a>

<a id="canonical-2201203122233323-1202200020120202-1001223301333020-2310102131101023-1203200002321320-0201330300020031-0331020230121210-1203033012130201"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.material_version` property

Type: `"string"`. Optional.

<a id="canonical-3003210302313030-1213211212132000-0021203220232311-3131312322113021-3030232310012111-2301001121130210-3202031321033302-3221121113033021"></a>

<a id="canonical-0013220030030232-1230020331202033-2132002033323211-1003321000130011-1022012012221322-2132101033221002-3100030200110231-1332111210232220"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.passphrase_env` property

Type: `"string"`. Optional.

<a id="canonical-1131323210021111-3331030033233111-2122022132330312-1120130221011311-0020213112223231-1011010123113311-3033113112231202-1033102302123230"></a>

<a id="canonical-2321120330113220-0200000212023123-3332332002221113-1131023231132311-3133331332203301-1031112011213301-1230120033132330-2321022333312033"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.passphrase_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-0232200332210113-2122220331311031-0101332010033323-1331200123100022-0120001332333331-2012202003131003-0331010020012112-3222221320332113"></a>

<a id="canonical-1221123332003302-3213031300010112-2332003122112110-3110213032113320-1321100132030123-0132102201323320-0013323330331310-1120203211113112"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.pkcs12_file` property

Type: `"string"`. Optional.

<a id="canonical-2300123231200300-0321330112103233-0203332023201323-3022301102003112-2211121121311313-0001012330230120-1221313130013132-3123211032020110"></a>

<a id="canonical-1033203313300220-2232322033303313-3112312211110310-0333231110200102-1013020021021201-1020322000333202-2320013212211020-2323303211031111"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.pkcs12_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-3310131032333222-2122311102223002-1232031022021012-1031321110131110-0100222122312302-1111310222012233-2320303112002032-2112113330312121"></a>

<a id="canonical-1013333211103123-2102221023033300-3302332033022131-1322311002202002-2033230101131200-1120011132010010-1230113231111111-3023203010030133"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.policy` property

Type: `"string"`. Optional.

<a id="canonical-0110111331100323-3222112332013310-2032012300321220-3223321132310220-3001312323223220-1010010031332223-2323322121303132-3313202112213130"></a>

<a id="canonical-2001221331003211-0212102203102220-0333122023020031-1330031213323011-0203302032133030-1121322310332031-3023323300122111-3203123301131021"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.prepared_identity` property

Type: `"string"`. Computed.

<a id="canonical-0311311113022001-1323013323102011-1202103133231113-2333232221333211-1230002111210301-3220111230003221-3223232133212203-0023211021220013"></a>

<a id="canonical-3322332003303121-1313023032010212-1013013130003303-2321013231120303-2130023203331221-2123131302310002-0231311222003331-0011211321020321"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.private_key_file` property

Type: `"string"`. Optional.

<a id="canonical-0200223111300120-3322322202200011-1100320310222102-2033222031011011-1112120222012312-1233103021100002-0231001302113202-3031011011002111"></a>

<a id="canonical-2303331323112332-3103100112311010-2031320310321131-1213212300331313-0122120121221030-2131301230223012-1303212110302132-0100031122303231"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.private_key_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-0232102302213201-1030300331222210-0202110200301221-0330332221230333-1121303110232300-0202133031012131-0121122301230333-0111213210000123"></a>

<a id="canonical-3132220310033101-1212122132112210-2112213103000131-3102131233123202-2322313331023300-3300111213020113-3313020312012111-3010213011202200"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.blindfold.spki_identity` property

Type: `"string"`. Computed.

<a id="canonical-2020213113300220-2333110233301203-3300222130030231-2130303232320030-2032111112212201-2300222323013202-0110222032000202-1320111112111031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms

<a id="canonical-1030021013012233-1132023103020003-0322031310212302-1102323132200031-1330231033232303-1101231202333203-0200121200223332-0323333130022223"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223031031211033-0320311121020102-1310011323203221-2211230320301032-1332022133211022-3133010133211020-1321002202133110-2002301230121103"></a>

### Direct properties for `default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms`

<a id="canonical-3331002031032101-2001130333223101-0302021313121000-3033323313301110-2233313222022113-0310102131332223-3010311101032113-2133332103300131"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1322333013020002-3320101320322330-1230010223210130-0332001200121133-0020320123030210-0012032023300221-0031101222021121-2210211232313232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="canonical-1300032021202110-0011031032313012-0300212012032212-3333302112232001-0331202031331003-3330010131203303-1201000113210300-3211001103103331"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_ocsp_stapling = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131202211301100-0331121201212013-2133312320113023-3033313101313002-3322302102012100-0123222110310223-0321031132101002-3021002201003223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- default_pool.use_tls.use_mtls.tls_certificates.private_key

<a id="canonical-2203210231332111-3031132312132123-1130302233133303-0110113131033132-2310332200311320-0330113002033032-3000202300132332-3111130320320313"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000020210331122-1132010321001303-2131303013231130-2312030102130212-2113301020003331-2331220300303232-2023111000331011-1122212223201201"></a>

### Direct properties for `default_pool.use_tls.use_mtls.tls_certificates.private_key`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-3202131212220022-3310331222003333-0132103303101313-3122213000000300-2332013301033012-3213003113213002-3322023211002302-2102301302313033): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-3303001112302110-2223022023002212-0233233102211101-2011031233223201-2222213211201220-2121132300103222-3002120121011222-3323020013000102): complete subsection reference.

<a id="canonical-3202131212220022-3310331222003333-0132103303101313-3122213000000300-2332013301033012-3213003113213002-3322023211002302-2102301302313033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key](resources--http_loadbalancer--reference--group-017.md#canonical-2131202211301100-0331121201212013-2133312320113023-3033313101313002-3322302102012100-0123222110310223-0321031132101002-3021002201003223)
- default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0022023210023202-2110310333003320-1012230102222210-2233113231311001-0133022031032111-3201010222100011-1032010111103100-3001101003123221"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022123013132032-1211200012123101-3203133100311113-1330210213232022-1131312032133021-3230003113201313-2023110023022120-1131222231323102"></a>

### Direct properties for `default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-2333112213321311-2001302311302120-3231303003201322-3022311002323320-3030210223320032-1001332012312303-0112023031001221-0011200221322313"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2111132120203332-1311130030030123-3213013212012201-2031333230020311-3233301023312330-3113330012020303-2132001123033300-2113231333020013"></a>

<a id="canonical-3231112100121132-2030200201021300-2003030123133322-2003200330323323-2003312223220220-0103112113321223-2312121112000310-0333120312220303"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0200303010213013-1132000102112230-2003002331021311-3213033131202221-3310031302201130-2213111100122120-0022023220201031-2111201322113232"></a>

<a id="canonical-0032031131221303-0132210003012202-3333101212203313-0300302322331020-0130222203322333-0130102020222133-2303100032333132-1101022210223210"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3303001112302110-2223022023002212-0233233102211101-2011031233223201-2222213211201220-2121132300103222-3002120121011222-3323020013000102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key](resources--http_loadbalancer--reference--group-017.md#canonical-2131202211301100-0331121201212013-2133312320113023-3033313101313002-3322302102012100-0123222110310223-0321031132101002-3021002201003223)
- default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

<a id="canonical-3300000121331033-2030010003221332-0033231020032330-2023223021123323-1001021121030230-1033001102331233-2321102003300023-0012201032120330"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1120230331230010-1032221231322231-2230320202013233-3123322312010201-1003002323310101-1310011332310131-1303020203311212-2003033002230001"></a>

### Direct properties for `default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info`

<a id="canonical-0200323133202323-3011021311113011-1303110212331210-1112101223331200-2031320312033333-0221222200323121-1013130021123101-0100323201211333"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3202120300223332-3022201000103233-1010230210303111-2022212302032133-3133120013120033-1032031221023220-1302011331020303-1232333130333211"></a>

<a id="canonical-2032120031030031-3103133010001123-2131200301121033-1203103231032202-2003133303003033-0232030202032332-3332331321300132-3022221221202222"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1112030331010001-1223303203031330-0333010110213321-3202130323302212-0212221311313102-1030203333133001-1312232202322201-0223303001130330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults

<a id="canonical-0133000112321230-1110222120112232-0333021301203130-3323020310312330-0023121000131003-1102101110023131-3330122333222230-0223021310123001"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_system_defaults = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011310303210021-1222202221103110-1130033232102230-1131030100322101-0032220031011132-0202012301223101-3013101110032002-0102022212323012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls_obj` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.use_mtls_obj

<a id="canonical-0022112112130111-3011023102110131-0333031030310303-2030022132112002-3133330123223031-1132202201000033-0130031010113333-0033302321210101"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
use_mtls_obj {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231122113302332-2002303233330021-0332111332303312-3023103303132330-2332103233301201-3220121103302022-1120330231012032-0320120231010122"></a>

### Direct properties for `default_pool.use_tls.use_mtls_obj`

<a id="canonical-0121211103102032-0330313001322120-1112233233332320-0130021120223001-0203302301023132-1303301312221132-1221121130032311-2110223112313130"></a>

#### `default_pool.use_tls.use_mtls_obj.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0123332020332101-0232133120232030-2000000233320132-1320110222321033-0220020120231321-0321213213001001-3211301002323303-1330020100320110"></a>

<a id="canonical-0302202001310110-1101211010303021-1032211223322000-1232121312203003-3332330010311011-0321132021331220-0003133013313231-2013033123223330"></a>

#### `default_pool.use_tls.use_mtls_obj.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2020133233123002-2322012220031112-1200111211102032-2113130302302200-2010102122302000-0310102312220002-2030202200222210-3221003330003323"></a>

<a id="canonical-3102322103321232-0122203111103223-2102232232331100-2032233303011301-2022033303310230-3002023012311031-3011112002202211-1301201333322133"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3302122303122133-0323233203310323-3012112030002300-2003211111010032-2031330012231010-1011303111010020-0113222311123003-1002223100220030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_server_verification` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.use_server_verification

<a id="canonical-2130103020232302-1031121000213331-1200111111332320-1123033212132121-1232002101101333-3211110303301331-3210133303122301-3102331023033332"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
use_server_verification {
  # Configure direct properties listed below.
}
```

<a id="canonical-2332002310031003-3001331101220003-3212223023322323-0121333111101322-0000211211232121-2311233003220121-1032333303122010-2031313303131031"></a>

### Direct properties for `default_pool.use_tls.use_server_verification`

- [trusted_ca](resources--http_loadbalancer--reference--group-017.md#canonical-3231101200310233-1023123032133030-2222223213120303-3302312032100301-0101231330003333-1230030230003100-2231003210132301-2012033101130000): complete subsection reference.

<a id="canonical-3302121201002223-1132313201012120-1122021012131202-2321102221122231-2311231033221310-1033030303311121-2313113220020103-2230223112333331"></a>

<a id="canonical-3321330102033021-2122301031113312-0311012031222230-2121230320311030-1222222133230230-1300113301122133-0311122210322212-2203322332333311"></a>

#### `default_pool.use_tls.use_server_verification.trusted_ca_url` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3231101200310233-1023123032133030-2222223213120303-3302312032100301-0101231330003333-1230030230003100-2231003210132301-2012033101130000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_server_verification.trusted_ca` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-3302122303122133-0323233203310323-3012112030002300-2003211111010032-2031330012231010-1011303111010020-0113222311123003-1002223100220030)
- default_pool.use_tls.use_server_verification.trusted_ca

<a id="canonical-1231200013223113-3313130001203221-0300003313303101-2321020102203333-3002301211023301-0111230320300103-1111010320123121-3333222302202002"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231001132223201-3000000230221201-3230301222103130-0111320221302321-0030221130133001-3003331233013102-2011232310231012-3302310102230223"></a>

### Direct properties for `default_pool.use_tls.use_server_verification.trusted_ca`

<a id="canonical-3213210320020231-0003132323301200-3232213232232323-3302123312201310-2032212230132002-1333320233001021-1313033203123313-3203312333330102"></a>

#### `default_pool.use_tls.use_server_verification.trusted_ca.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3202102011303222-0022021021120200-3313011233333120-0022133020230332-3023001112110013-3123131200213303-0211321023130233-1331322321231010"></a>

<a id="canonical-1222321332201113-1102203230332110-0233103200220002-0001321230302221-1032233212332323-3021332200322330-1000210222110302-3200312033033032"></a>

#### `default_pool.use_tls.use_server_verification.trusted_ca.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1211213032201123-1101110211320311-3002013002111020-2023201011122120-3122022312332313-0231013113112212-0332301133201321-2203322031021012"></a>

<a id="canonical-0133100203011322-1032013323121132-2012222331222131-1332113030212203-1100112231213000-3020003322203031-2223200131311022-3122323211010102"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3321010323221300-3221232231012333-1000122333111322-3013010233100302-2311021323120212-1121301113331022-3222031002320022-3101011201003023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.volterra_trusted_ca` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.volterra_trusted_ca

<a id="canonical-1123131102231031-3303002311113322-2223231122012003-1221030332331211-3223111220030330-1032310301010110-3231233101113000-2110012102333313"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
volterra_trusted_ca = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210031322021130-2010122303312010-3232211331331322-0132231002001122-1111001310002033-1200111313232223-3023311310020010-2112111310030332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.view_internal` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.view_internal

<a id="canonical-3111311300001101-1221223112021201-0221010001311121-2322133311013113-0202010001021132-0303200203011113-2102003220222010-2213223113000030"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
view_internal {
  # Configure direct properties listed below.
}
```

<a id="canonical-3332302222233002-1022201332120301-1130021200330033-2321122302112121-2312313112303302-2001301003102032-1032000212211130-3031303000121202"></a>

### Direct properties for `default_pool.view_internal`

<a id="canonical-2330201021132320-2202103302033221-1223110301030212-1210130012230213-3210132102120002-1001323103002030-2220033103133223-2131111310322330"></a>

#### `default_pool.view_internal.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1021013023031201-3020331300312210-1023213331112123-0032310012033303-0212302203133210-0123200110130131-2001122112230101-1023220200023222"></a>

<a id="canonical-0321012230120312-3030331101102322-0213222012012323-3121111131332033-0211023011120020-1101001302000322-3030132201202212-0132023333312133"></a>

#### `default_pool.view_internal.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3300212223020031-1202212011022000-2031210322333130-2033121132021000-3230132232123100-2211220301323322-0321113101221012-1130213312120220"></a>

<a id="canonical-2133223222112030-1320112311020112-1322022131230120-1221330231310223-1330010330210302-1133120002210301-3120230310123131-2110031003121202"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- default_pool_list

<a id="canonical-0130032320003112-3013231322132011-2010113111012220-3312013033110232-0201300301333010-2200230132012313-2122033102033333-3101112132302330"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
default_pool_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333223003212333-2032110013333332-0132202330213022-0021102200222000-1122122323130323-2112303013332332-0001213202032131-0210212101121222"></a>

### Direct properties for `default_pool_list`

- [pools](resources--http_loadbalancer--reference--group-017.md#canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112): complete subsection reference.

<a id="canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool_list.pools` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311)
- default_pool_list.pools

<a id="canonical-1130133122303033-3113103332212302-2223111002322122-3312003023012013-1321033110333130-2213233123210322-2330020110210033-0103212032210123"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-2320223320331321-3221011123300323-3302000000033320-3113130111031300-0323033222220323-1321310012301331-3231033022120221-3220032030010132"></a>

### Direct properties for `default_pool_list.pools`

- [cluster](resources--http_loadbalancer--reference--group-017.md#canonical-0203211033012202-0330031110233023-3003123220203232-0332013233023211-3220312232333201-2131113120302313-1301222332311031-0011201322300131): complete subsection reference.

- [endpoint_subsets](resources--http_loadbalancer--reference--group-017.md#canonical-3210332310321130-0113222213001001-2113013221322330-3212101123033102-3001002230002303-2132202333322300-1201201130002212-1120123101133323): complete subsection reference.

- [pool](resources--http_loadbalancer--reference--group-017.md#canonical-3013121320100123-0230112023312222-2313223300220001-0112310021033000-1001121323133302-2112022210113221-1313103310012302-3212310030321130): complete subsection reference.

<a id="canonical-3220013120332110-2330332100201203-3201012203133303-2311130120312132-1233130333001130-1120103003303110-1120202112311201-0132333330222330"></a>

<a id="canonical-0113203302233112-1001122031303022-2313321023320320-1032313021123112-2012022312020033-1310112020012032-3232321102213121-0001122123222010"></a>

#### `default_pool_list.pools.priority` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0213013110313133-0003022203030213-3033000010320023-2113231020313113-0020330133323230-3203123222100322-3031010223100220-0020323202001132"></a>

<a id="canonical-3300133133230030-1101000330222113-2132032031211202-3110133033312000-2120232030002232-0201312300133321-2303201000032001-1231230012322112"></a>

#### `default_pool_list.pools.weight` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0203211033012202-0330031110233023-3003123220203232-0332013233023211-3220312232333201-2131113120302313-1301222332311031-0011201322300131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool_list.pools.cluster` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311)
- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112)
- default_pool_list.pools.cluster

<a id="canonical-2022032001300210-0100113302101222-3222032123312312-2103201330210120-2101032102101110-2020300132231333-1131210122301102-1320322013203231"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032300223120013-0121112210213021-1031132233222331-2200211321102301-0102313223032300-2022010200130031-2322220132210302-0102100311111312"></a>

### Direct properties for `default_pool_list.pools.cluster`

<a id="canonical-0311122122232003-0001330212231032-0323030213332321-0012221211302122-2003033001030210-3023310203222202-1121310000101322-0002111102232213"></a>

#### `default_pool_list.pools.cluster.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0031320031332122-3030010133313031-0033331201001223-2330020223303233-3211232012020303-1100302133331102-2221200102013223-1333202223010221"></a>

<a id="canonical-0010331000011002-2223303223033332-1233332311032113-2323323113320301-1303111002111311-2301302300012322-2221203212303231-2000331110020212"></a>

#### `default_pool_list.pools.cluster.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2203132230103200-3030001133021013-0313332320330322-3221131102121132-2203131102233131-3211013012332223-2010111123133012-0311030233332021"></a>

<a id="canonical-1310302033320232-0021130020221332-2222331200333301-0323033233332122-2100000030123033-1012012110201121-1103233123033122-3320211021012331"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3210332310321130-0113222213001001-2113013221322330-3212101123033102-3001002230002303-2132202333322300-1201201130002212-1120123101133323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool_list.pools.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311)
- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112)
- default_pool_list.pools.endpoint_subsets

<a id="canonical-1323010032123123-0013103233303002-1223231222330131-1111223120220111-2332113311112011-3202211000212220-3321213312023021-0302111200223122"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
endpoint_subsets {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013121320100123-0230112023312222-2313223300220001-0112310021033000-1001121323133302-2112022210113221-1313103310012302-3212310030321130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool_list.pools.pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311)
- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112)
- default_pool_list.pools.pool

<a id="canonical-1321233303001203-2111130000110322-0211222220122211-2123023121132032-1032320012321220-0131303212011332-1101301023112013-0003013300110320"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302123202312111-3301030131020023-3202233131312000-2222221101000120-1322230133012023-1002033133020302-1320031011320232-0301210013321231"></a>

### Direct properties for `default_pool_list.pools.pool`

<a id="canonical-1103222010120231-0100200233110002-1220013113330203-3333210122000023-1221312202221212-3032031003201223-3302103010103000-3223001323303211"></a>

#### `default_pool_list.pools.pool.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2100233002120313-3201030132212123-0133331233101330-3022112132212221-0231332211031121-2323002013110211-0113131301203220-0330011120102133"></a>

<a id="canonical-1133000313130031-0212102100203203-3120023023113011-2023121032133012-3022021222032220-3333020030002311-0120222221330223-1132232031010011"></a>

#### `default_pool_list.pools.pool.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3310123313033112-2332030321211130-2022230010133023-2313000120231121-2332002201223320-0113030131130032-2310231032122020-3222230320203322"></a>

<a id="canonical-3130310303233111-1113032000331331-3331100000103113-1002121331300021-0123210033021010-3033100122310201-1303132121000203-0021211101322100"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_route_pools` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- default_route_pools

<a id="canonical-2030002000113130-1013213300122022-0332132120103002-3113212002131130-0011121312031211-0300001031000233-1333033103010131-3022331121021300"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Terraform syntax:

```terraform
default_route_pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-2111211333211322-2011031300210312-2200120222113021-3221311301201202-1320123221120321-1300131112132100-1012310023303323-1021123303322101"></a>

### Direct properties for `default_route_pools`

- [cluster](resources--http_loadbalancer--reference--group-017.md#canonical-1022310213023222-3210202333110033-3221111331321321-0023021200321202-3231012111100012-3231303220210100-0203203303300122-2202110232103332): complete subsection reference.

- [endpoint_subsets](resources--http_loadbalancer--reference--group-017.md#canonical-2122023231002110-2012202011030013-1132311102333232-0002323211230012-0321001232122003-1133213213111110-1221003221320031-2001301333112321): complete subsection reference.

- [pool](resources--http_loadbalancer--reference--group-017.md#canonical-2023320030321121-2112102000220202-0120020303101031-3101321002131113-3302203333030331-0031013322012221-1013012120003122-0122223312101000): complete subsection reference.

<a id="canonical-0330313201303120-1120121031310101-2210120002330330-1110232011202121-2330200231230010-1130003330010010-3120322130131301-3033000121220021"></a>

<a id="canonical-2323132220313230-2322132032303220-1033113033001003-3121011010333303-0231332223302213-1011112021203130-2132313321122023-1113221203202302"></a>

#### `default_route_pools.priority` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3133330301132011-3110200033223102-3003102020133323-0222120231200131-3033032131111222-3032200323331021-2100103232101121-0022232200012203"></a>

<a id="canonical-3200132213031302-1111031322331322-2013231313021023-0313312213302002-2210222033311202-1310000121302312-1130033202230230-3131313330010313"></a>

#### `default_route_pools.weight` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1022310213023222-3210202333110033-3221111331321321-0023021200321202-3231012111100012-3231303220210100-0203203303300122-2202110232103332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_route_pools.cluster` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_route_pools](resources--http_loadbalancer--reference--group-017.md#canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133)
- default_route_pools.cluster

<a id="canonical-2002221002211212-1223131201203133-2220103133110301-1121220022322222-2023301320102130-2132122333112203-0113111212121101-2311200332220212"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022133022132112-0332300221313020-3010002331322323-3200212330312322-0321223003220202-2021202131300321-2300022010221231-2322223101012130"></a>

### Direct properties for `default_route_pools.cluster`

<a id="canonical-2321130310233212-3011120101131101-0122311321202322-1021012232310321-1313330103231211-2212121312003322-1331100331302010-3111232302131313"></a>

#### `default_route_pools.cluster.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2132021200000231-0203102312233100-3021313333113112-2202332213211223-3101300232322222-0332332222011122-2000003003323323-0303033001301231"></a>

<a id="canonical-2001222232323231-2133210032020233-0023010132120010-3230333223102321-1023113322211122-3301201121012313-2311103023030331-0322311201031202"></a>

#### `default_route_pools.cluster.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2102211032310201-1130032233102003-1021203102210223-1112333023310123-1133012101230322-1323311102023132-1023102312213322-3133110103122021"></a>

<a id="canonical-3123223223301132-0120001023122331-0320112213011320-2231022223300100-0300323311311232-3022032202021312-2330312021311023-3130011331020102"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2122023231002110-2012202011030013-1132311102333232-0002323211230012-0321001232122003-1133213213111110-1221003221320031-2001301333112321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_route_pools.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_route_pools](resources--http_loadbalancer--reference--group-017.md#canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133)
- default_route_pools.endpoint_subsets

<a id="canonical-0112120313012032-0001303222003221-0211201103201130-2122120211000011-0323110300131202-0213103210100230-1120231223120312-1102301031011133"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
endpoint_subsets {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023320030321121-2112102000220202-0120020303101031-3101321002131113-3302203333030331-0031013322012221-1013012120003122-0122223312101000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_route_pools.pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_route_pools](resources--http_loadbalancer--reference--group-017.md#canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133)
- default_route_pools.pool

<a id="canonical-1203032021222111-2113022103331021-1300131313111112-1232202130003023-0110330312210131-0133132312322110-0333230121100002-2101201013033011"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103203221110330-3232222320030333-1032132220311011-1322222231230311-3110003113321132-3330333133133332-1201120132203133-2021103200002021"></a>

### Direct properties for `default_route_pools.pool`

<a id="canonical-3122123130000212-3202202332332100-0012310000222010-0231000023303120-1110023001033323-3221023310100320-1222311003133300-3201200231001333"></a>

#### `default_route_pools.pool.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1311111200130213-0111212202112123-3003121023302121-1220103230113311-2121002110301013-1320222032032311-1123111310200212-0021023223231212"></a>

<a id="canonical-3210301233203132-3220201221301321-2002111111022233-0013313332123000-2013103312202312-2122231312023020-0133332110302232-0001320100003312"></a>

#### `default_route_pools.pool.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1322332320222033-3223230221233020-0112021211033100-2100203221010300-3103122211123110-3312121002123322-1222131123201132-2213003001211233"></a>

<a id="canonical-1333200201100110-3223023202112330-1011111322210113-0101130301102003-1203322112211122-2323323223213030-3222301132333102-1021333002331002"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2012113030000132-3012100211231221-3130003303031233-0223013133101100-2030232213320012-3201201203311111-0130303112131112-1022213320131303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_sensitive_data_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- default_sensitive_data_policy

<a id="canonical-2033022123213123-2323113031001012-2201021021002231-3221313211000130-1202101233120223-2020021110211031-2103021121302100-2001220200021113"></a>

Type: `["object", {}]`. Optional, Computed.

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

- [default_sensitive_data_policy](resources--http_loadbalancer--reference--group-017.md#canonical-2033022123213123-2323113031001012-2201021021002231-3221313211000130-1202101233120223-2020021110211031-2103021121302100-2001220200021113)
- [sensitive_data_policy](resources--http_loadbalancer--reference--group-027.md#canonical-2331322013011313-3111202121321023-1113002312200000-0132100333303032-2311313300200313-3003330023023213-0133322032323223-3212012023320222)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_sensitive_data_policy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123111133030011-2123003110002011-1031032130100110-1022210200133330-3303330101000301-0203011300213011-1020303111133311-2132323123000212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_api_definition` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_api_definition

<a id="canonical-0121212322222130-3103021101033321-1123000003031210-3103120000300132-1002112302102003-3020032021001211-1330203332020133-1300330021233313"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
disable_api_definition = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232001202112132-3122220002210323-2123112210120123-2111123031002213-0132120133332022-2111330310202011-2220013001013012-0203213300003333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_api_discovery` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_api_discovery

<a id="canonical-3000030203212331-0031103221201012-0231200113003012-3001100132332313-3331212332021012-3322030122330202-3001012130221221-0001222023213222"></a>

Type: `["object", {}]`. Optional, Computed.

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

- [disable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3000030203212331-0031103221201012-0231200113003012-3001100132332313-3331212332021012-3322030122330202-3001012130221221-0001222023213222)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3332100020132100-1000303313230033-0200301222233203-1122131021321120-2033312022300121-1323331001230322-1320111223102220-3211202332310210)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_api_discovery = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313202103132030-0013213023113203-2231222203010002-3132330201211211-1031002112330010-2320302022023023-1021321113132320-2221311103102213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_api_testing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_api_testing

<a id="canonical-2102333133332200-1121021111010220-1122302223121000-2100012203322001-3000221313331130-1113021300210303-2123021303330312-2201022000213330"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
disable_api_testing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123003300002010-0033323131033110-3101003330221223-1232122203203021-0032123103302012-3321133121110332-3110131110231100-2002011203003323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_bot_defense` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_bot_defense

<a id="canonical-0031322221201123-0210020223332001-1212312133213033-3030023312222333-1332200121023202-1033110032330120-1010322121131033-3331133222011133"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
disable_bot_defense = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321101032330312-2021001302001033-0303012101301113-3221010231232001-2031231033003121-0102030231321121-2103303220300112-2032023320022021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_caching` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_caching

<a id="canonical-2310010330303022-3212102030230223-3311321133320300-3300013001030002-1102021010321320-0321310020220321-0011203222103312-1122212300121322"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_caching = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122321323123120-0131330211013033-0223101332002110-1133012200030110-0300010032111010-3133032130220300-0032001323100111-3300101020203300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_client_side_defense` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_client_side_defense

<a id="canonical-1313000133203032-2310331201223313-3311201202200022-2303113031201211-1010211132031221-2313112333213303-2001211021211300-1211030023323302"></a>

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
disable_client_side_defense = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202122330212221-3201331122203123-3033330230333010-1333110133101010-1102330300133111-1223133210101012-2101013113121223-0330312222221002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_ip_reputation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_ip_reputation

<a id="canonical-1300101232320303-3120302303031113-3023013133003230-2333310323023013-1321321023310323-0310123021012021-2232302023102200-1011120203321300"></a>

Type: `["object", {}]`. Optional, Computed.

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

- [disable_ip_reputation](resources--http_loadbalancer--reference--group-017.md#canonical-1300101232320303-3120302303031113-3023013133003230-2333310323023013-1321321023310323-0310123021012021-2232302023102200-1011120203321300)
- [enable_ip_reputation](resources--http_loadbalancer--reference--group-018.md#canonical-1331031122011033-2302131121012122-3313303233203321-2313300032330123-2331111321311122-1221001122320101-3033133302202020-0230023201131132)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ip_reputation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3012231223303232-2102332203003323-3102212131101310-2012202210031310-2221201202003333-1132132220302101-0321000202112303-1321313331130122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_malicious_user_detection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_malicious_user_detection

<a id="canonical-0131101221211313-1111300023003030-2102333222020213-1220113103132330-2110321310303202-1230232021202133-2200013231320023-2210202210021231"></a>

Type: `["object", {}]`. Optional, Computed.

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

- [disable_malicious_user_detection](resources--http_loadbalancer--reference--group-017.md#canonical-0131101221211313-1111300023003030-2102333222020213-1220113103132330-2110321310303202-1230232021202133-2200013231320023-2210202210021231)
- [enable_malicious_user_detection](resources--http_loadbalancer--reference--group-018.md#canonical-1200132303110020-2331223220030122-3300330122310330-2033201303222103-2311130312023020-2001011122320233-2033303101211233-0313110100132203)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_malicious_user_detection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303333110313121-3231031113100003-0101230011321330-0011203322101330-0102021301032021-1333303102122201-3123112221102102-3010103301030101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_malware_protection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_malware_protection

<a id="canonical-2102310133222323-3303112320121111-1102223222210312-0130011112231230-2320133120213321-1320112231313022-3302310213131033-2003201033332322"></a>

Type: `["object", {}]`. Optional, Computed.

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

- [disable_malware_protection](resources--http_loadbalancer--reference--group-017.md#canonical-2102310133222323-3303112320121111-1102223222210312-0130011112231230-2320133120213321-1320112231313022-3302310213131033-2003201033332322)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-021.md#canonical-2131132110220032-0231001231223002-1030113021221223-1220220323132330-0320302100310030-1022231121013021-1003021302003300-2223310300300221)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_malware_protection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010203201310323-0211110203000202-0211202211033123-0200222331221023-2010100321330113-1113001031021000-2100021313301232-0222321320030113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_rate_limit` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_rate_limit

<a id="canonical-1013223133201022-1313332221032232-1331200200123112-3000301121301112-0033001302203101-2312333301123320-2111320003320121-2311310003323022"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
disable_rate_limit = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201101022220023-1103210213332120-0012201103310011-2012220223033011-2311212133223002-2010030131033033-1301101200102210-0132000311201023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_threat_mesh` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_threat_mesh

<a id="canonical-2020310210321000-1320011300000131-3121003230133311-3120332120300122-0323030233301302-1132133023102100-3031111331113233-0231102311232332"></a>

Type: `["object", {}]`. Optional, Computed.

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

- [disable_threat_mesh](resources--http_loadbalancer--reference--group-017.md#canonical-2020310210321000-1320011300000131-3121003230133311-3120332120300122-0323030233301302-1132133023102100-3031111331113233-0231102311232332)
- [enable_threat_mesh](resources--http_loadbalancer--reference--group-018.md#canonical-2100022020020011-0200313113330200-3010002022110312-1310112023231102-1110222331121122-3303232030030220-1330000212210230-0030303210201210)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_threat_mesh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010130233211201-0102210200231211-2331021322013011-3201112123233301-1203313132231022-1130220320000222-0131212013022231-0330001200132223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_trust_client_ip_headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_trust_client_ip_headers

<a id="canonical-1023002121132310-3320330110101320-2303202132230113-1310320101111023-0012122130303312-2130100033130133-1323012232230013-1203210220020211"></a>

Type: `["object", {}]`. Optional, Computed.

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

- [disable_trust_client_ip_headers](resources--http_loadbalancer--reference--group-017.md#canonical-1023002121132310-3320330110101320-2303202132230113-1310320101111023-0012122130303312-2130100033130133-1323012232230013-1203210220020211)
- [enable_trust_client_ip_headers](resources--http_loadbalancer--reference--group-018.md#canonical-3331120100110120-1303333032232002-0301220313322023-3200013000131023-3322323222012333-3110100300213120-3312203133332021-1101320122300121)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_trust_client_ip_headers = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231033130322231-2203223313010330-1200320322113021-1333023133103333-3331123032011122-3320211321333012-0333310010101210-0312032310201021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_waf` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_waf

<a id="canonical-0021311132330031-3013103020211321-1223133002212101-2010122201101202-3101313331110313-2013021300302031-2213220033313203-3000113332031001"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
disable_waf = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302213303120331-0322020022021133-2103031201103012-0013213333030332-3331330132313320-2112113133220233-3320130323210022-3120123311111331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `do_not_advertise` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- do_not_advertise

<a id="canonical-3021311322031332-3001130120301333-0101131021301003-2310200320010222-0133031223311213-1001132011030231-1000103013131203-3110003031031311"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
do_not_advertise = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- enable_api_discovery

<a id="canonical-3332100020132100-1000303313230033-0200301222233203-1122131021321120-2033312022300121-1323331001230322-1320111223102220-3211202332310210"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
enable_api_discovery {
  # Configure direct properties listed below.
}
```

<a id="canonical-3030002131303001-0023303102231103-2212131100113301-3101321331002112-3303200131002010-3333200332003322-3013033322133100-2200231130120112"></a>

### Direct properties for `enable_api_discovery`

- [api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023): complete subsection reference.

- [api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-018.md#canonical-2022211120020101-1303213211101122-3331012220102332-2002132023213223-2013303102010122-0111321021212300-2033123331121313-1123111003011221): complete subsection reference.

- [custom_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-0000012321113332-1333302321200013-0013010310122302-2330031222212000-0110030002020101-3212323100133110-1003113113210123-3333201032123220): complete subsection reference.

- [default_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-2111032232010102-2001332210110221-3230210002330121-0101133121102002-0323332000113323-1033002020031111-0300232133222113-3321010221300221): complete subsection reference.

- [disable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-018.md#canonical-3303311120333123-3232301013211201-0312120323310102-2211133020013121-3002223311123231-0120220320333311-2120022120333300-3201333130101313): complete subsection reference.

- [discovered_api_settings](resources--http_loadbalancer--reference--group-018.md#canonical-3211321230303003-0331201322102020-1311001123210021-3120000101132231-1300020013122112-1222332100212100-3131201021231312-1022133323122023): complete subsection reference.

- [enable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-018.md#canonical-0330033332332131-3300000123331232-3002221001230011-1010202232113010-0222112121000100-1313311123011300-3110331031201133-0301111301201030): complete subsection reference.

<a id="canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- enable_api_discovery.api_crawler

<a id="canonical-3313213101132122-1202132221130300-1302030331330003-0311102333000311-0023322113031310-0303110020123220-2331222223222032-1200233000212320"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
api_crawler {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210212201020102-0220031023312332-0122121322322100-2133122102321131-2031230221030112-3132233110231131-0313131333010023-0223331032322322"></a>

### Direct properties for `enable_api_discovery.api_crawler`

- [api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310): complete subsection reference.

- [disable_api_crawler](resources--http_loadbalancer--reference--group-018.md#canonical-2130122330020020-0031312021130033-1110002033303211-2002113130032011-2312312132210322-0332022003213222-0313300212332231-1033111032302210): complete subsection reference.

<a id="canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- enable_api_discovery.api_crawler.api_crawler_config

<a id="canonical-0011013001010323-0102201332220003-3312322223210210-1000003121102120-2110103300130133-0222112132001333-3313031321030031-3123020010111021"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
api_crawler_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103321121120320-0223203021001320-0032031100023110-1201102311323111-2201202003322010-2031010302303311-1222003120010312-1300002220023001"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config`

- [domains](resources--http_loadbalancer--reference--group-017.md#canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002): complete subsection reference.

<a id="canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310)
- enable_api_discovery.api_crawler.api_crawler_config.domains

<a id="canonical-1000021102003321-3312202002202311-1232002300103200-0303230203313302-1222132301110311-3203011300032320-2232110133033022-3100220102020121"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

<a id="canonical-1223212023130233-1030031111233223-1101320131330101-0222231323330033-3102113321031221-3303133333010110-2230211133001233-0332231233102301"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains`

<a id="canonical-3231210223231332-2322303133210301-2113102310231310-1010033112023331-2310102001302003-0110102010333310-1011322123333203-3020211010113302"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.domain` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [simple_login](resources--http_loadbalancer--reference--group-017.md#canonical-1122331133203032-1120013303232001-1000132231113220-1022232221100203-2100330222212113-3021221231121331-0200032200310230-0330312322321302): complete subsection reference.

<a id="canonical-1122331133203032-1120013303232001-1000132231113220-1022232221100203-2100330222212113-3021221231121331-0200032200310230-0330312322321302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-017.md#canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="canonical-2231030121113112-0333331210332332-2033232111011202-3221331312310032-0221133131211003-0201231321332311-2303130001132121-0331103222231223"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
simple_login {
  # Configure direct properties listed below.
}
```

<a id="canonical-1300103322200223-3300132020200032-0302323210320033-0301321231311112-3322001101311302-2003123300112101-3020320001310301-0203113112322312"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login`

- [password](resources--http_loadbalancer--reference--group-017.md#canonical-0012010023323012-2110330100100131-2300112220200111-2330121011231202-1230012000332310-3223331231032021-0233102333030123-3001010310120021): complete subsection reference.

<a id="canonical-2211221111221131-3022220213010230-2121111111123201-3310220122311230-1232121000100223-1320032010100013-1233021331203020-3002211202131223"></a>

<a id="canonical-2303101311223013-1202320110213302-0311112123003011-3102323130000010-3021130311020111-0333220330103201-0113320332301321-0231333002211233"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.user` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0012010023323012-2110330100100131-2300112220200111-2330121011231202-1230012000332310-3223331231032021-0233102333030123-3001010310120021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-017.md#canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-017.md#canonical-1122331133203032-1120013303232001-1000132231113220-1022232221100203-2100330222212113-3021221231121331-0200032200310230-0330312322321302)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="canonical-0013221312120121-0121102003001202-0310012321221120-3320101300202223-0030113212200113-0330321012013030-0331302223133203-1201101320032231"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-1031233220310211-2321023021233100-3230123111310031-1000113030232231-0131321202132002-0003100301332133-3112010100331100-0213102102120313"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-3323233012331102-0210001301121113-1121332230211211-2111000000202033-2313313231002121-0110230011132333-0312110220111301-3012133330122133): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-018.md#canonical-1200020222210303-3300121022123201-1022330201232230-3212111103021323-2111332321032021-0130222012011322-1231220123003301-0320102102202021): complete subsection reference.

<a id="canonical-3323233012331102-0210001301121113-1121332230211211-2111000000202033-2313313231002121-0110230011132333-0312110220111301-3012133330122133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-017.md#canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-017.md#canonical-1122331133203032-1120013303232001-1000132231113220-1022232221100203-2100330222212113-3021221231121331-0200032200310230-0330312322321302)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-017.md#canonical-0012010023323012-2110330100100131-2300112220200111-2330121011231202-1230012000332310-3223331231032021-0233102333030123-3001010310120021)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

<a id="canonical-1131032213202212-3233101101203313-2201300100101322-3323113312122311-3021123003223220-2120222122121201-0012032131012321-0311200132222132"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001321030313203-1011000331232220-0030323302103230-0122231123230220-1001001201313333-3121122230022311-2002211221112212-2322203203031321"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info`

<a id="canonical-3222202303113312-2003211322313021-1213323210211310-2021211030312112-0031121313200000-3030313202313133-0013312130301313-3202013313001010"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0303130123333311-1122031200201332-2002211102311220-1321301032212100-2130112023032011-0200321112132021-2131010331232300-0001302311212020"></a>

<a id="canonical-3012302311322313-2120100011121200-0121330231213310-2021003033002333-1020232130133313-0010212213020120-2102133212020232-3002010023331023"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3300313000211221-1202020320112203-0022312132111110-0212012031332300-2013212102230320-0332300313311300-3011221231222311-1301210323230113"></a>
