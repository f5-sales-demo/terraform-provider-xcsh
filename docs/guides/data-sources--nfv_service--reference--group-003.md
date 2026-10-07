---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-2023003310212210-0333201133031110-1002202332200232-0222222010323012-0030320111222123-3332300202131113-0102011031023033-0203211211323021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-2133311313230323-2312031102220113-0231032111301301-2101330223210021-1220120303003222-0301102232300332-3233200231301001-0002100012111322)
- https_management.advertise_on_slo_internet_vip.use_mtls.no_crl

<a id="canonical-3333013112033210-3002032201322122-2322121003103110-1032212133202000-0133113221031330-0203031313121321-2001220300333333-3123221001322323"></a>

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

<a id="canonical-2031221030012123-1321032110300131-3211202312233202-1302222030301332-1201222310331232-1221312202031221-0012002022330313-2320211120102002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-2133311313230323-2312031102220113-0231032111301301-2101330223210021-1220120303003222-0301102232300332-3233200231301001-0002100012111322)
- https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca

<a id="canonical-3310200000001312-0022330101103213-0303021002213203-2121211321011120-3203010132122123-0103001113132022-3010212210212113-3101113111232010"></a>

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

<a id="canonical-3001223303102300-3002022123020023-2122110321223021-3210333302302321-1233101312330203-3310231320220332-2100323130312120-0310210313120003"></a>

### Direct properties for `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca`

<a id="canonical-3020203000230302-0223300123232002-2203133300000323-1200313032102111-3331212230022332-1101221033113220-1010310212032200-2331311302302322"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.name` property

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

<a id="canonical-1133300222301111-1010023010333102-0333201333311301-0021221102202321-3013001102011312-0331332033103030-2113321100103112-2303120320112223"></a>

<a id="canonical-0323222022311210-2001113333231333-0001320200110110-2202210132120233-1220012332220101-3213221223222011-2303313332100230-1110113002113221"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-0030302221012003-2033232303003110-2210312220121313-3113132022331022-2333121200122321-1131201021232012-1011210323103103-2111303221023111"></a>

<a id="canonical-0012300220202112-3223112010120121-3322232230220010-1330003033300300-0013122010000332-2102033312310002-1122131013030200-2122330012003013"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-1122022200012120-2120301112202011-3300222101201102-2213231323302120-1323121000233121-0331231101131322-0302312323113031-2231030212302013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-2133311313230323-2312031102220113-0231032111301301-2101330223210021-1220120303003222-0301102232300332-3233200231301001-0002100012111322)
- https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled

<a id="canonical-3223123030001322-1031301010313102-0001110332031221-3113330211233311-2032011301320130-0202200213120013-0000322022221233-1222331202301322"></a>

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

<a id="canonical-3133220123300002-0032012031333210-1302212011031201-2301012310203002-0120210222030320-2013132212330100-3032333231330332-1313001023320203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-2133311313230323-2312031102220113-0231032111301301-2101330223210021-1220120303003222-0301102232300332-3233200231301001-0002100012111322)
- https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options

<a id="canonical-0212332210313031-0032001310031302-1113322132113321-0121200130233013-0333220210322331-2020333113033223-0320213233100031-1332233013120303"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3312201232222031-0313110003202330-2120303303020110-1230101202211110-0231222001103032-1030031010301132-1233112011211021-1032230232331212"></a>

### Direct properties for `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options`

<a id="canonical-2011010030320310-0101110311022121-1203120331220103-0313110322223331-3021110313131103-0333231301121101-2203330103303000-3001113320313312"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- https_management.advertise_on_slo_sli

<a id="canonical-3131211132010013-0223123101133212-0231212131233312-1000123022221333-1001103213230001-2301212201113200-2033231301231311-2103100233111011"></a>

Type: `"single"`. Computed.

Configuration parameter for advertise on slo sli.

Additional upstream details:

Inline TLS parameters.

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

<a id="canonical-1233002110320322-0102010102333101-2003033013313110-2003303300203122-2103003221010122-1131001121123130-2123132120321010-1331020002100021"></a>

### Direct properties for `https_management.advertise_on_slo_sli`

- [no_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2012022332300101-3323303201302013-1323301222200232-3100302230023130-1330212212330031-2013012103320020-3023022320131130-3221321301033110): complete subsection reference.

- [tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222): complete subsection reference.

- [tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102): complete subsection reference.

- [use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001): complete subsection reference.

<a id="canonical-2012022332300101-3323303201302013-1323301222200232-3100302230023130-1330212212330031-2013012103320020-3023022320131130-3221321301033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.no_mtls` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- https_management.advertise_on_slo_sli.no_mtls

<a id="canonical-1321122231003130-1300322123110300-1020220201321121-2333132000103302-2013131323023332-2231220132132223-2320100200322300-0331220031323033"></a>

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

<a id="canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_certificates` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- https_management.advertise_on_slo_sli.tls_certificates

<a id="canonical-3222023331233302-0313302322333022-0221212310233102-3200222230312110-2123022333032330-3113301231331322-1332033212110313-3232022021012231"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3231222332022020-1310013002313132-1313233310330232-2022123210221023-0202011020020233-1320211103121223-1133232103312300-0333010322321220"></a>

### Direct properties for `https_management.advertise_on_slo_sli.tls_certificates`

<a id="canonical-0010332003012210-0331200330210321-2133011201123012-1330310203330220-1303223200031231-1101013022001201-2120122232212023-2103320003312233"></a>

#### `https_management.advertise_on_slo_sli.tls_certificates.certificate_url` property

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

- [custom_hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-2222100222232310-2013003031331332-0313213033020011-1133230222212303-2333223201103113-2013212312101133-3211310311133022-0331220301223231): complete subsection reference.

<a id="canonical-3233331210221210-1201302113031023-2011220312131102-0011113211323020-0012002222111310-3021113311101022-1300300231120113-3112021202213123"></a>

<a id="canonical-3033323203113233-1001023220103313-3123020313231120-0231310302011100-0133031121001333-1020123101202100-1101001312132132-1213033110303132"></a>

#### `https_management.advertise_on_slo_sli.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--nfv_service--reference--group-003.md#canonical-1212032211010213-1300113010212023-1211323101023211-2231123120212113-0303000101220112-0221113031322132-0303210023110130-3303113031110302): complete subsection reference.

- [private_key](data-sources--nfv_service--reference--group-003.md#canonical-2332101223013010-3001210002232222-3212231032130311-3030102210332202-1020201122000223-1130131231123133-1333223103212323-1023011200331221): complete subsection reference.

- [use_system_defaults](data-sources--nfv_service--reference--group-003.md#canonical-0133101210102013-0312301022020002-2113101030100112-2020212012103230-0020300230032231-3101002300120132-1333020300122310-3020301232020201): complete subsection reference.

<a id="canonical-2222100222232310-2013003031331332-0313213033020011-1133230222212303-2333223201103113-2013212312101133-3211310311133022-0331220301223231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms

<a id="canonical-1321301323022332-0030100002131133-0223033322200000-0122311100311112-3023232321332103-3001020333313011-2211002022332312-3202213321030212"></a>

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

<a id="canonical-3210031130021202-1332133011220110-0031010023033312-3202231102231231-1222202301332111-3231203012123122-3032300011233021-0100120300113122"></a>

### Direct properties for `https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms`

<a id="canonical-1121330301113311-2130213323201030-0002200303001022-3222110230301233-0132221022001023-2023301220303132-0110002303232101-2333210313200310"></a>

#### `https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms.hash_algorithms` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1212032211010213-1300113010212023-1211323101023211-2231123120212113-0303000101220112-0221113031322132-0303210023110130-3303113031110302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling

<a id="canonical-0021311210231103-2310002322101333-1130131301313130-1313223220220312-1001320000200121-1332133220111122-1210221332301210-0002003223230022"></a>

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

<a id="canonical-2332101223013010-3001210002232222-3212231032130311-3030102210332202-1020201122000223-1130131231123133-1333223103212323-1023011200331221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- https_management.advertise_on_slo_sli.tls_certificates.private_key

<a id="canonical-2010232010033330-2133223331322131-1002033103313321-1123230100033321-0122223001030033-2200202013012020-3323321202130202-1212030312122302"></a>

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

<a id="canonical-3332200212032010-2022011032021310-3332222021200203-2101100323232332-2321023201001102-1332200221120012-2113311132222221-1212221131010023"></a>

### Direct properties for `https_management.advertise_on_slo_sli.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-0102020033231312-3211201210113123-2303221302220301-1233233220321202-3103131303032200-1122313020233023-3132330211330300-2031230302213321): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-1231030232233310-1212233131111122-2222133203002230-3300200312020232-3023200320003111-3021232002312233-3210333323101223-0210313223112022): complete subsection reference.

<a id="canonical-0102020033231312-3211201210113123-2303221302220301-1233233220321202-3103131303032200-1122313020233023-3132330211330300-2031230302213321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-2332101223013010-3001210002232222-3212231032130311-3030102210332202-1020201122000223-1130131231123133-1333223103212323-1023011200331221)
- https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3013111330211010-1310323113231010-0020212013333233-0332213130210110-1223122312212233-2102300100220321-0022003310213020-2113022201003013"></a>

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

<a id="canonical-1002132300301031-1012033222100022-2101331020021121-2231100002132102-3303001220010120-2320231020202010-3231332223031201-1333213020322111"></a>

### Direct properties for `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-3112321331322210-0013020320113211-2033103220133202-1313110010121310-1213223023112010-2210202010233110-3301023021211121-3210200330003103"></a>

#### `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1122212111011122-0213031131013010-2201000331031110-1102010000012323-3222222100030012-0200132132102021-2331123203020011-3233131333033131"></a>

<a id="canonical-0002200221122002-2013223321232301-1031031113133011-0201011000221221-1113003333011131-3022110003001200-3323333213013112-3012321103103210"></a>

#### `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0212023201302100-3310001221033300-0002330220110011-2300232031213101-0230110331230130-3202131313013112-1102303100303210-1000331211332323"></a>

<a id="canonical-0223102113100200-0013231013012021-3130310003120313-2212123223303231-0313301001130233-0001332333112102-3333321030121320-1321203131113220"></a>

#### `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1231030232233310-1212233131111122-2222133203002230-3300200312020232-3023200320003111-3021232002312233-3210333323101223-0210313223112022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-2332101223013010-3001210002232222-3212231032130311-3030102210332202-1020201122000223-1130131231123133-1333223103212323-1023011200331221)
- https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info

<a id="canonical-3102200103003321-1303323013231331-1122013110333002-1210301123122030-3033313101202132-1311301021302112-1301022033310023-1103302103201223"></a>

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

<a id="canonical-0032011223222211-3030220210210232-2230213012313011-2001000013213211-3201121201111330-0013033133123020-3013131010310031-1132211312131232"></a>

### Direct properties for `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info`

<a id="canonical-2022120033333001-2123301212310101-3002111223030100-1322023022301333-2200123231203100-2320233223002123-1111110002313230-0332102320130020"></a>

#### `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3221220200233112-0200333223112331-2211201121120200-1120321332203232-2131233202221200-3031212102000303-3023331030323300-3232112200333201"></a>

<a id="canonical-1322031330301101-2212322320101332-1121123003232131-1121312003101132-2222231113202332-2033313003120011-1111221113200102-1313033130000330"></a>

#### `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.url` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0133101210102013-0312301022020002-2113101030100112-2020212012103230-0020300230032231-3101002300120132-1333020300122310-3020301232020201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-1121121023103032-2333031020023100-2020002313301110-3102001021210322-1330001210203301-0223223303222000-1301321001201022-0033210323233222)
- https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults

<a id="canonical-1011300131103331-1101222232230232-1112012020002203-2121322200203100-1331102312310223-0320112330212201-2223113220331301-1232301233020200"></a>

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

<a id="canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_config` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- https_management.advertise_on_slo_sli.tls_config

<a id="canonical-0203331200121210-1010331113302133-1223110132233332-3101312231211102-3203121012112221-2220302010011030-0222212022112201-2332333300213223"></a>

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

<a id="canonical-1130110131113030-1020332123023310-2132312012301023-1013123111002212-3320302012301310-2022323000003102-1013121313113232-2300023022321102"></a>

### Direct properties for `https_management.advertise_on_slo_sli.tls_config`

- [custom_security](data-sources--nfv_service--reference--group-003.md#canonical-3100231302011200-3012332322212010-0201121023003011-0131101100100332-0110221002330031-0110202001231020-2111020113221310-1232221131020020): complete subsection reference.

- [default_security](data-sources--nfv_service--reference--group-003.md#canonical-2230121303013002-1110032200323102-0102100310120112-0333321133002023-2221120220001303-1323100121112210-2123022030020020-1113320123022223): complete subsection reference.

- [low_security](data-sources--nfv_service--reference--group-003.md#canonical-0321132112313203-1121123332132013-1111113012300033-0010030031332231-3221001031220130-0321330010100231-3322032122321013-1321202222300232): complete subsection reference.

- [medium_security](data-sources--nfv_service--reference--group-003.md#canonical-3210123103110023-2233011033122122-2221001102310010-2221001111300221-3211012001310232-0313110101211103-3120310322302100-3031310132123211): complete subsection reference.

<a id="canonical-3100231302011200-3012332322212010-0201121023003011-0131101100100332-0110221002330031-0110202001231020-2111020113221310-1232221131020020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102)
- https_management.advertise_on_slo_sli.tls_config.custom_security

<a id="canonical-2003310123302022-3031022321010231-1333130130333103-2232003202202030-1310013131300333-1200303020223010-0013332312011010-1220033210010011"></a>

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

<a id="canonical-3020312023123023-3311022031330023-3232123000223233-3330022221132000-1303101233033030-3223031211103132-1012031220302131-2103013301123312"></a>

### Direct properties for `https_management.advertise_on_slo_sli.tls_config.custom_security`

<a id="canonical-3131101130000202-1333202011001023-2213000332123230-3321210310321211-0120112321312202-1031011320133203-3210023220320021-3013131113001013"></a>

#### `https_management.advertise_on_slo_sli.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2011213120131012-3003130120232020-0301101201220211-3320111022223203-1231220321103033-0001210213001233-2202331012332113-2033122020311232"></a>

<a id="canonical-1010001131223323-3131330233120033-0021020013223223-1210011333133021-0010003302111020-2222031032101120-2230120001123321-3021012313222201"></a>

#### `https_management.advertise_on_slo_sli.tls_config.custom_security.max_version` property

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

<a id="canonical-3212120022230003-2200332320323032-0201133302320001-0101233322032113-2101231010130023-3212001032220013-2131311022131322-0213131101133122"></a>

<a id="canonical-2031321231110333-1103331020021322-3333321112031203-0312233122200123-1220232330102323-2333013130120020-3002002230121133-1211320122020210"></a>

#### `https_management.advertise_on_slo_sli.tls_config.custom_security.min_version` property

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

<a id="canonical-2230121303013002-1110032200323102-0102100310120112-0333321133002023-2221120220001303-1323100121112210-2123022030020020-1113320123022223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102)
- https_management.advertise_on_slo_sli.tls_config.default_security

<a id="canonical-0121021131322000-2220310131010230-0212113211102310-3322020111010211-3103232102321111-1101211203003332-2000103012223301-0223133231001233"></a>

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

<a id="canonical-0321132112313203-1121123332132013-1111113012300033-0010030031332231-3221001031220130-0321330010100231-3322032122321013-1321202222300232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102)
- https_management.advertise_on_slo_sli.tls_config.low_security

<a id="canonical-3221103332010201-3130003221220300-0020020213203002-0303311312032012-2033100312303223-2220333303122320-3321012200331030-0330211212100212"></a>

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

<a id="canonical-3210123103110023-2233011033122122-2221001102310010-2221001111300221-3211012001310232-0313110101211103-3120310322302100-3031310132123211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102)
- https_management.advertise_on_slo_sli.tls_config.medium_security

<a id="canonical-3300332300100133-3110222201010132-0030300131320200-0213033221101010-1133210000223322-2303213302330013-0021322331230331-2201021323230302"></a>

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

<a id="canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.use_mtls` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- https_management.advertise_on_slo_sli.use_mtls

<a id="canonical-0220301210013231-2031133121322131-2321211303313123-2321200230333121-2131103031021331-3332000122011320-2322321100210003-3102211223333212"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-2332103300300300-2112322203220002-0122310130111002-1333313232230200-2131101223130232-0232203133220031-0102130301103033-2231110320031202"></a>

### Direct properties for `https_management.advertise_on_slo_sli.use_mtls`

<a id="canonical-1123210220103213-0002301223201301-1303320332033302-0021312232213330-0233022103102312-1011211113003232-0302303323010010-3231032001331011"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.client_certificate_optional` property

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](data-sources--nfv_service--reference--group-003.md#canonical-2231012030033122-3100121100332200-2111010231133322-1323303201000022-2010023120122201-2110211333230232-1123110131133233-1233012211033203): complete subsection reference.

- [no_crl](data-sources--nfv_service--reference--group-003.md#canonical-1310000021221121-2032333013130313-0110330133211200-2332221132231321-0300111223331310-2220102221030320-0103112020122322-0020300031232320): complete subsection reference.

- [trusted_ca](data-sources--nfv_service--reference--group-003.md#canonical-1202012222213033-1111202132312301-2203231331030310-2322322120332330-1330031100231220-2220012132310323-2033322100100032-1331222212323332): complete subsection reference.

<a id="canonical-3332101211031221-3302201201320010-2300211121103020-3311210223133013-3013231103121313-3321233213233231-0003001010132331-1122330310010301"></a>

<a id="canonical-0323201230233322-1120213020000101-0101311313202202-0121022123010011-2202200312300000-3123032003001211-0221201313300013-0012331300112122"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [xfcc_disabled](data-sources--nfv_service--reference--group-003.md#canonical-2001120122302233-3113012010021131-1023130131302032-2330332233131220-1022100121110221-0003220211022012-2221120213331102-2132130111101002): complete subsection reference.

- [xfcc_options](data-sources--nfv_service--reference--group-003.md#canonical-0033021311113333-3302323211310210-2332213113112032-2000223203331002-3313031020032003-0330322322203202-0012130312220323-0101013112213022): complete subsection reference.

<a id="canonical-2231012030033122-3100121100332200-2111010231133322-1323303201000022-2010023120122201-2110211333230232-1123110131133233-1233012211033203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001)
- https_management.advertise_on_slo_sli.use_mtls.crl

<a id="canonical-1030211031122330-3300021113200032-1201113121300013-0320000010120221-0011312021003132-0321030112211320-2221313231201121-1022103031202123"></a>

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

<a id="canonical-0010333320223201-2311231033120222-2201223331110020-2220233013103323-0303111202333313-2303022122110122-1300122233130101-2133200111230022"></a>

### Direct properties for `https_management.advertise_on_slo_sli.use_mtls.crl`

<a id="canonical-3312103323122202-3101213120330233-1010033020031001-2331030331112122-2333320112020201-3112131333200211-0002132031223322-0110203220202131"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.crl.name` property

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

<a id="canonical-2113331133210320-1212322022132230-1310002023132313-1133202322220032-1103311031131212-3332031101013323-2030003000300101-1230022011122011"></a>

<a id="canonical-0010001222323311-3310220220111201-1032011312313110-1330233311011020-1320013230221331-3230223113211001-0231223103020332-2230233033022322"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.crl.namespace` property

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

<a id="canonical-2010303300211231-3130302310322031-0310302211202010-0300023323222232-2032232221323020-1121320202132131-1203330203132303-2011333123332000"></a>

<a id="canonical-0131201001220123-3133312332123110-3023111202020202-2003011331210301-2113132200200100-3311113002031101-3133021330102322-0122203103100123"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.crl.tenant` property

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

<a id="canonical-1310000021221121-2032333013130313-0110330133211200-2332221132231321-0300111223331310-2220102221030320-0103112020122322-0020300031232320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001)
- https_management.advertise_on_slo_sli.use_mtls.no_crl

<a id="canonical-1103022002012221-1011003023310022-3101111122130111-0223010003023020-1131320030220203-2213002111033012-3022033003111112-3210211330010030"></a>

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

<a id="canonical-1202012222213033-1111202132312301-2203231331030310-2322322120332330-1330031100231220-2220012132310323-2033322100100032-1331222212323332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001)
- https_management.advertise_on_slo_sli.use_mtls.trusted_ca

<a id="canonical-2323320112132010-2030320110022023-1023110313032123-1031010112113332-0030011211211310-0202033011120201-3203030032303120-0011231232111112"></a>

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

<a id="canonical-1311300100133123-1012222331300310-2213330303010312-1313330033110200-1011022032221001-1311333332211010-2213332123311232-2223002112102330"></a>

### Direct properties for `https_management.advertise_on_slo_sli.use_mtls.trusted_ca`

<a id="canonical-1032102323222210-1021132002133311-2332331221320313-3112230213220211-0231211103021231-0221022210230000-0201223010231330-3023132223313022"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.name` property

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

<a id="canonical-0331311121231131-3031113302131210-0002203003122023-3010010311021013-3232000030100113-2133120100001003-0021023230013000-2203213121133011"></a>

<a id="canonical-0231021211133111-2302133011323213-2131122133313220-0112201001113330-2133210111331131-3220300230303221-3210311310221303-1223332301123210"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-3232133223312131-3330132031300022-1022212102011202-0112131330013300-2013020010023302-3112012212231022-1002213020032122-1033311111233100"></a>

<a id="canonical-0000221300011113-3111231032010002-3133331032102302-1133212231000130-2203033110213221-1003312303303302-2213003213210210-2322333310023231"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-2001120122302233-3113012010021131-1023130131302032-2330332233131220-1022100121110221-0003220211022012-2221120213331102-2132130111101002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001)
- https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled

<a id="canonical-3103311232022002-2312013011333210-3003200010031311-0203032221101331-1233201002213213-0300130200123013-2132030112233220-0013212223310122"></a>

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

<a id="canonical-0033021311113333-3302323211310210-2332213113112032-2000223203331002-3313031020032003-0330322322203202-0012130312220323-0101013112213022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-0020123120200102-2033113211132012-2200032011031312-1223023002121021-3000223312011321-0233001211302123-3033221230312323-0221203331201133)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-2220012321032010-0113022132323030-2200231112233301-2111220120023102-1131120200323311-3323112313233102-3232233320103230-3211223003120001)
- https_management.advertise_on_slo_sli.use_mtls.xfcc_options

<a id="canonical-1302302203011030-0000223002133131-3223123331103311-3202223313033231-1130130330313121-2311313031100212-2111122111101000-1001121202230323"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1323213312313121-1213030211320022-2032020110303030-3313101232323123-1321131202110311-3321102313011330-1032312000103301-3011130120323011"></a>

### Direct properties for `https_management.advertise_on_slo_sli.use_mtls.xfcc_options`

<a id="canonical-3123223231030223-0000100000012111-0132202320002301-1322331223213123-3110221033012013-2122101312300132-2013202111000111-3012123101202332"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- https_management.advertise_on_slo_vip

<a id="canonical-1201300311232210-3120333123332212-2021000011210110-3212320212312103-0200203032021102-1102331330302300-0230103102031032-3001223111313130"></a>

Type: `"single"`. Computed.

Inline TLS Parameters. Inline TLS parameters.

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

<a id="canonical-1021032320223210-1213112002110302-0222321300032130-0232000211132113-0030001232231213-0201031301130312-3213021233200211-3210303113003120"></a>

### Direct properties for `https_management.advertise_on_slo_vip`

- [no_mtls](data-sources--nfv_service--reference--group-003.md#canonical-3220211322023012-0001232100122300-2230133210110122-3120020203121220-2031310222201120-1200032123320032-1300012333132002-0010320013000121): complete subsection reference.

- [tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031): complete subsection reference.

- [tls_config](data-sources--nfv_service--reference--group-003.md#canonical-3130313110212202-3233203131203031-1300221202131002-3112121030130200-3311320033012120-0333301110000020-2023133110202300-1230200211303322): complete subsection reference.

- [use_mtls](data-sources--nfv_service--reference--group-004.md#canonical-0103133122313112-0200103330302132-3220110112310121-3311103021212012-3232011230223002-1001210001011302-0010023021332133-0311310321210300): complete subsection reference.

<a id="canonical-3220211322023012-0001232100122300-2230133210110122-3120020203121220-2031310222201120-1200032123320032-1300012333132002-0010320013000121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip.no_mtls` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- https_management.advertise_on_slo_vip.no_mtls

<a id="canonical-1220200323331332-3032011231113212-1312000202230132-1011301100021333-0033311332023202-0111021031311213-1130311120121212-0300120223131223"></a>

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

<a id="canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip.tls_certificates` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- https_management.advertise_on_slo_vip.tls_certificates

<a id="canonical-0121122103200222-1012011112210020-3110030201112113-1012023330221331-3211121123032211-3233211021301202-2123020032213103-2022312131000303"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0130031130303032-0330231201200200-1033333333330223-2020122203131132-2030212103002022-3333010113020312-0003030330133123-0033123221030023"></a>

### Direct properties for `https_management.advertise_on_slo_vip.tls_certificates`

<a id="canonical-0322323131220131-0002320112021220-3232331022302301-3103332210210312-2301321223012031-0010200132003211-0123213200211031-2003221121202331"></a>

#### `https_management.advertise_on_slo_vip.tls_certificates.certificate_url` property

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

- [custom_hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-0221000212202130-3210313010022010-2232113210121100-0011002030223301-1232233133331213-1032132202102110-2321002220230203-3120032222320220): complete subsection reference.

<a id="canonical-3123200033312303-2110122033222221-2331313212332013-3221231103322023-0022203331020131-0132232310003130-0232203022021221-3312230311310233"></a>

<a id="canonical-3223321020301103-2303030113031311-0230300320232013-3202200303212300-2033213321112111-0213132322202323-0031021101000100-0200212113211310"></a>

#### `https_management.advertise_on_slo_vip.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--nfv_service--reference--group-003.md#canonical-3103033100313002-1001200330133000-3123021313000313-3232312133120101-0203013302111030-1103330102002202-3201320312131321-3221102010023233): complete subsection reference.

- [private_key](data-sources--nfv_service--reference--group-003.md#canonical-1022230000233122-0013300323020303-2031230212033333-2112001010210000-2311031113331101-2131231221301132-2303333321320302-2033233200133201): complete subsection reference.

- [use_system_defaults](data-sources--nfv_service--reference--group-003.md#canonical-3120001133113131-1121113132221211-2220223333332210-1133330221220002-2332133022211320-3110123332103223-1012132320111232-1232332131231111): complete subsection reference.

<a id="canonical-0221000212202130-3210313010022010-2232113210121100-0011002030223301-1232233133331213-1032132202102110-2321002220230203-3120032222320220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms

<a id="canonical-1221210211321123-0213132013332310-0012011303330123-2232330212002132-3212320012133301-3022110111320211-1131031333300023-0210010213323010"></a>

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

<a id="canonical-2300313233202033-3322020103102110-2300111331332112-3231020323212223-3123132333002222-2323012203130201-0020331111301320-2031332200221303"></a>

### Direct properties for `https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms`

<a id="canonical-2302211003133230-3221322033323131-3003230303233231-2003232320002003-2011002333101230-1212200322011312-0210003220303101-2321132113312333"></a>

#### `https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3103033100313002-1001200330133000-3123021313000313-3232312133120101-0203013302111030-1103330102002202-3201320312131321-3221102010023233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling

<a id="canonical-2023321032211332-0103323222331103-2030223023200111-0231130223102001-3230000021032000-3122211131230232-0111220220031111-2122010233012110"></a>

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

<a id="canonical-1022230000233122-0013300323020303-2031230212033333-2112001010210000-2311031113331101-2131231221301132-2303333321320302-2033233200133201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- https_management.advertise_on_slo_vip.tls_certificates.private_key

<a id="canonical-3123231300111133-3110002301033002-3123233302201310-0310111320110032-2002320013002103-1032022011231322-2010301212332331-3211212321122010"></a>

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

<a id="canonical-2213223221202111-2021300010121130-0202321003022302-2310231133212032-0323220030133320-3221312230231323-3223000133220203-2032003102330331"></a>

### Direct properties for `https_management.advertise_on_slo_vip.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-0031321200102323-3033032000210210-3120313102002310-0230031222323222-2233211332203121-3032011233023200-2323220132023030-0022310322222210): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-1120022112132112-0123020212030331-1311122312300330-3123320100221133-1320031221221133-1010003212000020-0312322003022132-2003130122310131): complete subsection reference.

<a id="canonical-0031321200102323-3033032000210210-3120313102002310-0230031222323222-2233211332203121-3032011233023200-2323220132023030-0022310322222210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-1022230000233122-0013300323020303-2031230212033333-2112001010210000-2311031113331101-2131231221301132-2303333321320302-2033233200133201)
- https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3011320311322121-1221001133201211-0330102213303330-1230130220032130-2221130002331221-1111201303222201-3302310010021133-0211023002331111"></a>

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

<a id="canonical-3333103311132320-1101310212330000-3023200003111123-0302000210130210-0213031332002002-3320201301301332-0120132112110331-0133223023113312"></a>

### Direct properties for `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-3122211320123100-0333111203023112-2022331012133010-1030202301011023-2230330131321032-3211010302120120-3113213011231130-0212110001211212"></a>

#### `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2201302300211210-1001313133331333-0222223200112033-2030133101323223-0021101303313032-3211132022032203-2331022010123131-1030111310232011"></a>

<a id="canonical-1323323131111112-2312211133303022-0330202230203130-1223213013002302-3313001221311223-1000010123200122-3022222302122202-1311000320013011"></a>

#### `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1332200202330211-2033111200022200-2121333221212013-3212321212221231-1202130312133021-2210032201203233-0032221212011220-1200133112331210"></a>

<a id="canonical-0030323321301102-3303100333032022-1122023230000130-1133122003331321-3203213111130131-3020031330223012-1121201323230102-1212110110112123"></a>

#### `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1120022112132112-0123020212030331-1311122312300330-3123320100221133-1320031221221133-1010003212000020-0312322003022132-2003130122310131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-1022230000233122-0013300323020303-2031230212033333-2112001010210000-2311031113331101-2131231221301132-2303333321320302-2033233200133201)
- https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info

<a id="canonical-2321112223130110-1300120031131131-1330232121012122-3330303311000200-0131011223223121-2130122313300320-2232112032103203-1322233313303312"></a>

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

<a id="canonical-1210222110231320-2102112211230211-3203112311001113-1020321010223121-3203322123323301-1000200023033031-2103031311121212-2220301121312302"></a>

### Direct properties for `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info`

<a id="canonical-3202321232020211-1123033003113021-1333000330131303-3322112112132202-3303130122320112-0030300132300122-2112222303210023-0111021001211013"></a>

#### `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0121211120313132-1121222130130003-3223311113021333-2221102312001112-0033300113003032-0210022111223333-1212111120011322-1331021011132120"></a>

<a id="canonical-2001001322130210-3111301001111320-0333202232233130-2232233320032023-0112001333122011-3033120000222322-3021201003033112-0121011010022311"></a>

#### `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.url` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3120001133113131-1121113132221211-2220223333332210-1133330221220002-2332133022211320-3110123332103223-1012132320111232-1232332131231111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0200001102301211-1033021201203233-1030232003331233-0111221203201202-0122203321001021-3223233311021210-3322200011230332-1030333312311031)
- https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults

<a id="canonical-2113020101012132-1303202132222230-1102010311001030-2033300013033310-2101111202131100-1232223233133230-2320013211022331-0123322321000103"></a>

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

<a id="canonical-3130313110212202-3233203131203031-1300221202131002-3112121030130200-3311320033012120-0333301110000020-2023133110202300-1230200211303322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip.tls_config` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- https_management.advertise_on_slo_vip.tls_config

<a id="canonical-0331311322030123-3123133111123223-1321020012313323-3133130002311013-1233121302202011-0213001010011132-2011231203312111-1113211312013310"></a>

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

<a id="canonical-1002013220100213-0012200311122032-1120320321203213-0130022012003012-2113212211202300-1220211023023330-3330200031323103-3332100303210131"></a>

### Direct properties for `https_management.advertise_on_slo_vip.tls_config`

- [custom_security](data-sources--nfv_service--reference--group-003.md#canonical-0103313302212330-0232203133213031-1003020033121220-3102223121103031-0132120121110003-0121003220222031-1310033332203001-1033222120220333): complete subsection reference.

- [default_security](data-sources--nfv_service--reference--group-003.md#canonical-0233013023232122-2002020013231313-2022201012303003-0222102212221232-2310131321230300-0133032311222301-0000002200101300-0123132333110211): complete subsection reference.

- [low_security](data-sources--nfv_service--reference--group-004.md#canonical-2200001211220003-2202201113331331-1020230032121033-3001123000131233-0311231211131321-2010221332323310-1332322311011030-2003003121012001): complete subsection reference.

- [medium_security](data-sources--nfv_service--reference--group-004.md#canonical-1212101031202310-0021203331222312-3012300120123120-0223232312230320-3102201023221033-0233031213111021-0010032010110131-3211202131303121): complete subsection reference.

<a id="canonical-0103313302212330-0232203133213031-1003020033121220-3102223121103031-0132120121110003-0121003220222031-1310033332203001-1033222120220333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-3130313110212202-3233203131203031-1300221202131002-3112121030130200-3311320033012120-0333301110000020-2023133110202300-1230200211303322)
- https_management.advertise_on_slo_vip.tls_config.custom_security

<a id="canonical-3302111021202022-3232013023000120-1321033123102103-3002320312322001-3003112130302021-3202310211000131-2011022332212301-0211232211010110"></a>

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

<a id="canonical-1022330121321220-1231221330212111-3330130001333332-3220023121120213-1010323303203310-0102213132202003-2121001220112230-0120233120232302"></a>

### Direct properties for `https_management.advertise_on_slo_vip.tls_config.custom_security`

<a id="canonical-0120031320113330-1320031212132112-1303331130121003-3300022211100111-2000131300102320-3112301022100211-3111010033201332-2010123310330023"></a>

#### `https_management.advertise_on_slo_vip.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2311111223013322-2322021130101320-0210000303302121-0233011213331320-2020221233213003-3131111001022032-3210231113331221-0130312220310333"></a>

<a id="canonical-1022323312100021-0323001210000001-1301232333233313-3031123110302031-1100033311013111-0223223213030012-2330332020111023-2002302123313323"></a>

#### `https_management.advertise_on_slo_vip.tls_config.custom_security.max_version` property

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

<a id="canonical-2102113320230302-2212022233111301-1220300111302322-0221023133011023-2012033122312100-1330313131210110-2332323222011130-3111000130030010"></a>

<a id="canonical-2313232021311030-1311320003100220-3333233023033300-2101013023003001-0210321123302000-0200302223213220-3222122203032032-0111113033230300"></a>

#### `https_management.advertise_on_slo_vip.tls_config.custom_security.min_version` property

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

<a id="canonical-0233013023232122-2002020013231313-2022201012303003-0222102212221232-2310131321230300-0133032311222301-0000002200101300-0123132333110211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-2130231331130310-1010031103103213-2033131333110003-1133202011322121-0231021221030031-3212101323211200-2310221122201110-0231001033022213)
- [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-3130313110212202-3233203131203031-1300221202131002-3112121030130200-3311320033012120-0333301110000020-2023133110202300-1230200211303322)
- https_management.advertise_on_slo_vip.tls_config.default_security

<a id="canonical-2200010202331020-1200013310313322-2313030002130232-0121322330210222-2312122231202323-0233333320220120-3120103220121320-3200122331130231"></a>

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
