---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-2313310002111202-1332131303132010-2111300111031203-3121312022221032-3203333230000233-0300233203011031-2310330201113301-1120323203132331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.policies` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- rate_limit.policies

<a id="canonical-2002302101102310-2312130131203003-2103203120111012-0000100230221121-3101002120331211-2021303310233010-1131121311001033-1312121022302122"></a>

Type: `"single"`. Computed.

List of rate limiter policies to be applied.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2122213033323300-3013333313111133-0231001223320030-3232000203010213-3232112133323313-1200320212103001-2312313220132032-3202320103103222"></a>

### Direct properties for `rate_limit.policies`

- [policies](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-2123030230123213-3003201113200321-2102002112113222-1221332132310001-0233022103323232-0111022033303003-2232321202011322-0212021002130321): complete subsection reference.

<a id="canonical-2123030230123213-3003201113200321-2102002112113222-1221332132310001-0233022103323232-0111022033303003-2232321202011322-0212021002130321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.policies.policies` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.policies](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-2313310002111202-1332131303132010-2111300111031203-3121312022221032-3203333230000233-0300233203011031-2310330201113301-1120323203132331)
- rate_limit.policies.policies

<a id="canonical-3312320300123301-3221022003313313-2321331010100021-1311330100023302-2333121231233230-0020202031200230-0033200232301223-3300312302132201"></a>

Type: `"list"`. Computed.

Rate Limiter Policies. Ordered list of rate limiter policies.

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2323112011010222-0012201000130232-1230003232333100-0223333012120023-2320302220331222-0131212220133131-3200122223320220-3230201331103013"></a>

### Direct properties for `rate_limit.policies.policies`

<a id="canonical-1131121121101021-3232100301213121-2232123111021101-0231221111110333-3133311002312201-1021321002302223-2231021213011231-2332130121233201"></a>

#### `rate_limit.policies.policies.name` property

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

<a id="canonical-3113101002001101-0223303202310133-2110010211311021-3012221112122202-2312002022223022-1200200002221321-1300103022003201-2221021330332323"></a>

<a id="canonical-1321232203122003-2321333132021202-0100030302333202-2032303020300033-1203310332010232-2112331201211122-1011003133023002-1010322310002111"></a>

#### `rate_limit.policies.policies.namespace` property

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

<a id="canonical-2132120312032312-1321312232202210-3212210233130310-0003131012201222-0023033211130200-1130231213001110-0332223132301101-0222030032220303"></a>

<a id="canonical-3310131232301303-2201200102310123-1201010302032221-3212011111201030-2003021001313211-2011032223210011-2033333003121010-1113000302323033"></a>

#### `rate_limit.policies.policies.tenant` property

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

<a id="canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- rate_limit.rate_limiter

<a id="canonical-0323031000020212-1020332030233022-3133010311222331-1230211303131230-3201022033111030-1330201222023320-1030132330000100-2022212301110311"></a>

Type: `"single"`. Computed.

A tuple consisting of a rate limit period unit and the total number of allowed requests for that
period.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"action_block\",\"disabled\"]",
  "x-ves-oneof-field-algorithm": "[\"leaky_bucket\",\"token_bucket\"]"
}
```

<a id="canonical-2212132100032320-3232020123033021-0131210303002121-1332013221032002-0011311311300320-3130011222020220-1013202320020123-2003002102222301"></a>

### Direct properties for `rate_limit.rate_limiter`

- [action_block](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3101021031220001-3001323321230112-2231323032031231-3102032211012101-2012121302130120-1303001033213021-3101130232220320-2122303100331133): complete subsection reference.

<a id="canonical-3323232212030120-1203201103011312-1220122323212130-1322200222013110-2103313330130022-2213021000202222-3301232133332321-1330110333132123"></a>

<a id="canonical-3022000322213113-1330120131130220-2101121321200022-1101133031210310-0202201000322333-3101222311112323-3221202132313200-2022300022320332"></a>

#### `rate_limit.rate_limiter.burst_multiplier` property

Type: `"number"`. Computed.

The maximum burst of requests to accommodate, expressed as a multiple of the rate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [disabled](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0331302311112031-3222133032210112-3302130211000130-1133123210332321-2113023112030121-3311200112202331-0110212213031013-2121312333102320): complete subsection reference.

- [leaky_bucket](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3201330333010112-1212102231030302-0330013131032310-0133101113323222-2230123323311220-1321321310323102-3030011201103331-0122022211303132): complete subsection reference.

<a id="canonical-3311031021010330-3133203130310321-0200200211302221-0211210201031210-1313313003101223-0312331202233300-3203231220221213-0303123033231122"></a>

<a id="canonical-1231032002033322-2023233111020222-1320000110221022-3331232220223032-0330330310320211-3322122223033032-0130111021322230-1322113222210220"></a>

#### `rate_limit.rate_limiter.period_multiplier` property

Type: `"number"`. Computed.

Setting, combined with Per Period units, provides a duration. Server applies default when omitted.

Additional upstream details:

This setting, combined with Per Period units, provides a duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  }
}
```

- [token_bucket](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-2100233220213310-3331000322232012-1332130230010221-0311210210033233-2323222231303200-1333133201102311-1001122011310233-0321213032103003): complete subsection reference.

<a id="canonical-3012102131022310-2102100031000121-0203030022332112-0012130120231112-3102212220320121-3230020010213300-3011000111113231-0310102103032121"></a>

<a id="canonical-3231111003222023-3212012110212031-2002012011202321-3100113333123330-3033010210120012-2031121123003112-3121312333220311-0301231211103000"></a>

#### `rate_limit.rate_limiter.total_number` property

Type: `"number"`. Computed.

The total number of allowed requests per rate-limiting period.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-1213202012010123-1020212013303230-2033132102130322-3121301212312022-2212203101321131-1211333332310310-3211333200113121-0320020323230232"></a>

<a id="canonical-0303001030133201-3023222322310201-3012210010202030-1221120201100131-0011102221113233-0103002131013123-3003323213233223-2003010012303200"></a>

#### `rate_limit.rate_limiter.unit` property

Type: `"string"`. Computed.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3101021031220001-3001323321230112-2231323032031231-3102032211012101-2012121302130120-1303001033213021-3101130232220320-2122303100331133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- rate_limit.rate_limiter.action_block

<a id="canonical-3003333123001223-1001231201202333-0002212202323201-2020312231033021-2113330202001022-2331021312103202-2223002101220330-2010223030202033"></a>

Type: `"single"`. Computed.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

<a id="canonical-3203311320010120-1320312302001211-2103303210030030-1030132312122030-2033230030012212-1121102310011332-1222133031132010-3323201120300321"></a>

### Direct properties for `rate_limit.rate_limiter.action_block`

- [hours](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3211123323310320-1011311213122201-0203313032010331-3113132222223002-3032123221132011-1313100231111312-2020301103032321-1332011112011111): complete subsection reference.

- [minutes](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-2220112301220021-0220033021122302-0201203030311023-3330000301302111-1012133210213233-2121320201330333-2233220312100030-1130330132100023): complete subsection reference.

- [seconds](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3231033012222010-2120232321333112-0312012110302313-2013223121230200-3132113103010121-3220000003010221-1221021310321010-2213212111200310): complete subsection reference.

<a id="canonical-3211123323310320-1011311213122201-0203313032010331-3113132222223002-3032123221132011-1313100231111312-2020301103032321-1332011112011111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.hours` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3101021031220001-3001323321230112-2231323032031231-3102032211012101-2012121302130120-1303001033213021-3101130232220320-2122303100331133)
- rate_limit.rate_limiter.action_block.hours

<a id="canonical-2202031212331312-1310122211201230-3001313302302132-2130013122111221-0332131012222313-1311121232210300-3030122210233211-2012001232112201"></a>

Type: `"single"`. Computed.

Hours. Input Duration Hours.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3312013113303010-1113001322123211-3321101133310321-1320323312111112-3132010213221010-3101202223300003-1220200013303023-0122020222100031"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.hours`

<a id="canonical-0311001032230122-1231321322233333-3112321100133312-0302130132102101-2301023311113102-1232223020103213-3113313223131223-2333201021220233"></a>

#### `rate_limit.rate_limiter.action_block.hours.duration` property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 48,
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  }
}
```

<a id="canonical-2220112301220021-0220033021122302-0201203030311023-3330000301302111-1012133210213233-2121320201330333-2233220312100030-1130330132100023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.minutes` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3101021031220001-3001323321230112-2231323032031231-3102032211012101-2012121302130120-1303001033213021-3101130232220320-2122303100331133)
- rate_limit.rate_limiter.action_block.minutes

<a id="canonical-2202301202301330-1010021011122321-3232302032013101-2310322213233012-0303030230002032-2122220012110210-0333012212110112-1303011232010103"></a>

Type: `"single"`. Computed.

Minutes. Input Duration Minutes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1311320210010122-0100102132213000-2101301233131321-1032030211103312-3100332020110330-3230132330201330-2121023222213112-1021322230123312"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.minutes`

<a id="canonical-1031100223211130-2020212312232323-1231100310320312-2232302102310113-2312333303212312-1032300233203003-1131331233332111-0212311122220000"></a>

#### `rate_limit.rate_limiter.action_block.minutes.duration` property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60,
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  }
}
```

<a id="canonical-3231033012222010-2120232321333112-0312012110302313-2013223121230200-3132113103010121-3220000003010221-1221021310321010-2213212111200310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.seconds` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3101021031220001-3001323321230112-2231323032031231-3102032211012101-2012121302130120-1303001033213021-3101130232220320-2122303100331133)
- rate_limit.rate_limiter.action_block.seconds

<a id="canonical-2312113213323230-1332103132221020-0012133313200220-1010302211210232-3100133312013311-1211203223120032-3131312023220333-2101200031111233"></a>

Type: `"single"`. Computed.

Seconds. Input Duration Seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0313122001123200-3322200302102333-2311311031302231-3233133330310133-1300311123203210-0222301132030013-0233200021123332-3223012013011313"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.seconds`

<a id="canonical-1302120110211202-3120111030031011-1210103033003300-1221023000322130-3300022010213230-3021313312212220-3001020023212331-3112313013123232"></a>

#### `rate_limit.rate_limiter.action_block.seconds.duration` property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-0331302311112031-3222133032210112-3302130211000130-1133123210332321-2113023112030121-3311200112202331-0110212213031013-2121312333102320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.disabled` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- rate_limit.rate_limiter.disabled

<a id="canonical-3320112322221201-0020020311231301-1313221021301323-0222220213103031-3220323103013010-0130300011201201-2210210013323010-3033203131230032"></a>

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

<a id="canonical-3201330333010112-1212102231030302-0330013131032310-0133101113323222-2230123323311220-1321321310323102-3030011201103331-0122022211303132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.leaky_bucket` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- rate_limit.rate_limiter.leaky_bucket

<a id="canonical-3321030103331122-2103312033303231-1130111200311102-0333220232132302-2100122101103000-0333311300102021-1001331121030232-0203031320312321"></a>

Type: `["object", {}]`. Computed.

Leaky-Bucket is the default rate limiter algorithm for F5.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-2100233220213310-3331000322232012-1332130230010221-0311210210033233-2323222231303200-1333133201102311-1001122011310233-0321213032103003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.token_bucket` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- rate_limit.rate_limiter.token_bucket

<a id="canonical-2121010302310131-2112203311222031-3133302230101321-2022013321131011-0230123133332301-1302011203202122-3031133221101132-2313321133322220"></a>

Type: `["object", {}]`. Computed.

Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-2323021210132303-1002202011101212-1233211310323332-0231102230303132-1000132101121121-1023313113211031-3200220110120300-3202302103311313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sensitive_data_policy` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- sensitive_data_policy

<a id="canonical-2122233223233301-0013212123322233-3031133033203232-2102201333013333-1030211211110231-1012300301301312-3233203101222202-1212331200200122"></a>

Type: `"single"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Settings for data type policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2022321300200121-0232222331022232-3003331120312033-1223001300130112-0323300121123000-2332202102303212-0311020010002332-3211301003321322"></a>

### Direct properties for `sensitive_data_policy`

- [sensitive_data_policy_ref](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0122021200123201-2333022220103031-1003130123300110-2103010222210022-3331133212310133-1210012021332011-0032031301021233-3012221222302111): complete subsection reference.

<a id="canonical-0122021200123201-2333022220103031-1003130123300110-2103010222210022-3331133212310133-1210012021332011-0032031301021233-3012221222302111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sensitive_data_policy.sensitive_data_policy_ref` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [sensitive_data_policy](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-2323021210132303-1002202011101212-1233211310323332-0231102230303132-1000132101121121-1023313113211031-3200220110120300-3202302103311313)
- sensitive_data_policy.sensitive_data_policy_ref

<a id="canonical-2231132022303230-0121331112312222-1022210213113132-0333323120223110-2233033323330000-1113210210130010-1231210331321103-3220113011113233"></a>

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

<a id="canonical-2011111013223121-3102020030003002-1332203333133210-1300133333232011-3211310021313212-3301332110013231-3200023001302331-3112031011131200"></a>

### Direct properties for `sensitive_data_policy.sensitive_data_policy_ref`

<a id="canonical-1332020033303331-1313120133321112-3100120000310211-1323013103023310-0310302132333122-3110331233133001-3022212120032101-3212023010130221"></a>

#### `sensitive_data_policy.sensitive_data_policy_ref.name` property

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

<a id="canonical-2100302322021300-2102102322033121-2233312023012011-2201213301211232-3032230102100322-1210131201311133-3301021233003312-1100031023112332"></a>

<a id="canonical-2200121233100300-1302130302103211-1323321020213320-1102103023232111-0331211133100103-2110312210220231-0111130300332011-1121331023213231"></a>

#### `sensitive_data_policy.sensitive_data_policy_ref.namespace` property

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

<a id="canonical-3101220003101333-1101010010313001-2330113323330010-2021220133222230-2000121331002032-2101112102012220-2112322231031000-1003122203101031"></a>

<a id="canonical-2123122001122122-1201302221023130-1123120312001100-3301011121232220-3030003203232000-2211113002332213-2303003212211033-0202221013300300"></a>

#### `sensitive_data_policy.sensitive_data_policy_ref.tenant` property

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

<a id="canonical-2013200020130301-1113102330210232-2012033221211111-1333232332201202-0131030031333211-0001020213120320-0310201210113323-0331122232203010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service_policies_from_namespace` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- service_policies_from_namespace

<a id="canonical-1331300022003002-1003301212201010-2332302332323033-1313300013202011-1100101111101023-0231310011000031-3302000103312122-2320001102213020"></a>

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

<a id="canonical-0001301213230111-3103221002232321-2201322121023120-2011320222222231-0223121210332333-2311331313122203-3000301231032130-2101022223112121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slow_ddos_mitigation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- slow_ddos_mitigation

<a id="canonical-0023310130101023-0230201330111310-3031213012121012-3103313113100121-1221222010012130-0111200013330032-3031010222212132-3121322220130212"></a>

Type: `"single"`. Computed.

\[OneOf: slow\_ddos\_mitigation, system\_default\_timeouts; Default: system\_default\_timeouts\]
'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Additional upstream details:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-request_timeout_choice": "[\"disable_request_timeout\",\"request_timeout\"]"
}
```

OneOf alternatives in this subsection:

- [slow_ddos_mitigation](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0023310130101023-0230201330111310-3031213012121012-3103313113100121-1221222010012130-0111200013330032-3031010222212132-3121322220130212)
- [system_default_timeouts](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3001200310333022-1331031220203200-0013302003210231-2102121022303323-2101333130010312-0023302113231120-3031313301021300-2120031030013100)

Select alternatives according to the provider validators above.

<a id="canonical-1130220321321010-2313102021123121-0231122302111023-1121230233120231-1210122113023311-1020033112113322-3033110021213300-2003021132011201"></a>

### Direct properties for `slow_ddos_mitigation`

- [disable_request_timeout](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1220220212022233-3120201131002122-2132032112200101-0212011221332003-0001313022200102-0013302320021130-3322132230001110-2320330203021223): complete subsection reference.

<a id="canonical-2320000122020013-0210303313220232-2220033122001120-1331123032022023-0332232210100001-3203213030321012-1112011212033001-0322003020021013"></a>

<a id="canonical-2010123112311213-1033232033232330-0012001323223210-2102100130111002-1201103101323022-0211322312012320-3230212130133132-0001012020132333"></a>

#### `slow_ddos_mitigation.request_headers_timeout` property

Type: `"number"`. Computed.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Additional upstream details:

The default value is 10000 milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-0223123000120113-1011122122130300-1220021000012131-3310132122302311-2103302032303201-3322232122210201-0112213020101202-2102233313103332"></a>

<a id="canonical-0022210033010313-1112002301330223-1012332210302203-2310120113203323-0211001332020221-1002101012303102-3313223002001122-2320300311202110"></a>

#### `slow_ddos_mitigation.request_timeout` property

Type: `"number"`. Computed.

Exclusive with \[disable\_request\_timeout\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  }
}
```

<a id="canonical-1220220212022233-3120201131002122-2132032112200101-0212011221332003-0001313022200102-0013302320021130-3322132230001110-2320330203021223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slow_ddos_mitigation.disable_request_timeout` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [slow_ddos_mitigation](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0001301213230111-3103221002232321-2201322121023120-2011320222222231-0223121210332333-2311331313122203-3000301231032130-2101022223112121)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-2322213132011011-0331103323322123-1000020210313102-3323023222032013-0130023322000023-0110130311321310-3101330333130023-3230013323322310"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable request timeout.

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

<a id="canonical-3322032032102002-0320311003220131-1102031321312113-0210213301323200-2120310230130103-3321101121233321-0212132023110130-0313002323330033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `system_default_timeouts` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- system_default_timeouts

<a id="canonical-3001200310333022-1331031220203200-0013302003210231-2102121022303323-2101333130010312-0023302113231120-3031313301021300-2120031030013100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for system default timeouts.

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

<a id="canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- trusted_clients

<a id="canonical-1231002223203112-2102323012011101-2120210301321000-1212120202121220-1221020100301101-0321313233000031-3032023110221030-3230033232302203"></a>

Type: `"list"`. Computed.

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. For clients.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-2011013103330302-1010132332202310-2230001332023221-0120133211112213-0022232323213110-1312131213102133-2331323121331020-3123002123321130"></a>

### Direct properties for `trusted_clients`

<a id="canonical-1103130113231113-2333112300011113-3311010310112210-0211223213022102-3203231223300223-1031303221222231-3120213300323101-0110133020221121"></a>

#### `trusted_clients.actions` property

Type: `["list", "string"]`. Computed.

\[Enum:
SKIP\_PROCESSING\_WAF|SKIP\_PROCESSING\_BOT|SKIP\_PROCESSING\_MUM|SKIP\_PROCESSING\_IP\_REPUTATION|SKIP\_PROCESSING\_API\_PROTECTION|SKIP\_PROCESSING\_OAS\_VALIDATION|SKIP\_PROCESSING\_DDOS\_PROTECTION|SKIP\_PROCESSING\_THREAT\_MESH|SKIP\_PROCESSING\_MALWARE\_PROTECTION\]
Actions that should be taken when client identifier matches the rule. Possible values are
\`SKIP\_PROCESSING\_WAF\`, \`SKIP\_PROCESSING\_BOT\`, \`SKIP\_PROCESSING\_MUM\`,
\`SKIP\_PROCESSING\_IP\_REPUTATION\`, \`SKIP\_PROCESSING\_API\_PROTECTION\`,
\`SKIP\_PROCESSING\_OAS\_VALIDATION\`, \`SKIP\_PROCESSING\_DDOS\_PROTECTION\`,
\`SKIP\_PROCESSING\_THREAT\_MESH\`, \`SKIP\_PROCESSING\_MALWARE\_PROTECTION\`. Defaults to
\`SKIP\_PROCESSING\_WAF\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3323210032130133-2102021100011201-1031130311121132-2130110121222130-2032232312320131-1100000200201000-1220203123331330-0210013033021030"></a>

<a id="canonical-3030131311223203-2321003323021201-3130211213032221-1230200220203213-3220130022213133-0311101201223232-2203323300203300-2221221000112312"></a>

#### `trusted_clients.as_number` property

Type: `"number"`. Computed.

Exclusive with \[http\_header ip\_prefix IPv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 401308,
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
    "ves.io.schema.rules.uint32.lte": "401308"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  }
}
```

- [bot_skip_processing](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1002221023312331-2122202022123303-3203010332010121-1333203310022212-3123210030132320-2032131300232232-2211222212021011-1023102010122303): complete subsection reference.

<a id="canonical-3230320020132232-3120233111333120-0123022133303002-2320331313032212-0003213303123001-0010001212013331-0003312303200233-2202203131022232"></a>

<a id="canonical-0310113110230023-1221330121223213-0032020313030023-3203311201012020-3031332130120321-1022202203113311-0320222203322010-1030313132113012"></a>

#### `trusted_clients.expiration_timestamp` property

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Additional upstream details:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [http_header](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3331313011113302-2133211203211001-3331033213021103-3121012200312032-0111302102020100-2110033001021302-3231323300322103-3311322122012121): complete subsection reference.

<a id="canonical-3023211101203210-1123323211202221-0212320321123223-2101132020201111-2300310120133202-3120021312333222-1203322223033103-2013233223023113"></a>

<a id="canonical-1033210023010220-1203121201303103-0331303220013222-2030130312130233-3131231202330001-2020321311222012-3012020100130110-0101230330323110"></a>

#### `trusted_clients.ip_prefix` property

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header IPv6\_prefix user\_identifier\] IPv4 prefix string.

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

<a id="canonical-2012330230123301-1212221200112112-3120310013001200-3121201102112113-2103020133332301-0002222200331003-2333331121301021-2302002222322223"></a>

<a id="canonical-2022110310303123-2010211210232131-0330202123111331-2000000322011232-0131221323023123-0103033221222132-3102211022201033-2231021322011320"></a>

#### `trusted_clients.ipv6_prefix` property

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

- [metadata](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3223231021000022-1323110230023303-2001333120313120-2313222332030222-3132312200330212-2013103300312311-0321130232003030-3331121232230020): complete subsection reference.

- [skip_processing](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3121002301000111-3032111303202230-0123013210102222-0022322133120202-2110121100211130-0132320310300120-2321303012320220-1103131131231233): complete subsection reference.

<a id="canonical-1102112132002313-1303333030221120-0020322231011032-3331310122300300-1123121002221021-0322102031022030-0321003220220131-1212000031030200"></a>

<a id="canonical-2233201111013110-2230331001101030-1320311202110011-2103233011213003-3033121100102213-1320332320121112-3200130001323033-1331201022303233"></a>

#### `trusted_clients.user_identifier` property

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header ip\_prefix IPv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

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

- [waf_skip_processing](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3302120220301322-0303132200032323-0103130112023221-1022001202212200-0313022213103220-2013200021311220-1002322311331131-1302331323230003): complete subsection reference.

<a id="canonical-1002221023312331-2122202022123303-3203010332010121-1333203310022212-3123210030132320-2032131300232232-2211222212021011-1023102010122303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.bot_skip_processing` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- trusted_clients.bot_skip_processing

<a id="canonical-1300022011211111-2013333013333230-2101222302130303-2233022201331223-3103312223333030-2121120212122211-3003113100100101-1232320130330003"></a>

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

<a id="canonical-3331313011113302-2133211203211001-3331033213021103-3121012200312032-0111302102020100-2110033001021302-3231323300322103-3311322122012121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.http_header` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- trusted_clients.http_header

<a id="canonical-1033112300123121-2231301220122033-2120223131312202-0021011200120001-3300133130131023-2131223030323223-3323013322200100-0110320123332020"></a>

Type: `"single"`. Computed.

Configuration parameter for http header.

Additional upstream details:

Request header name and value pairs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0032032311111011-0013123130312103-0222233123322123-3112230211330110-0133131110310110-1220211230323133-1222101000012103-2030021320213312"></a>

### Direct properties for `trusted_clients.http_header`

- [headers](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3203211213223211-1213100101012313-3103031023200023-0330023211213211-2311233023100013-3232132132220313-1003200120032131-1110023100132320): complete subsection reference.

<a id="canonical-3203211213223211-1213100101012313-3103031023200023-0330023211213211-2311233023100013-3232132132220313-1003200120032131-1110023100132320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.http_header.headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- [trusted_clients.http_header](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3331313011113302-2133211203211001-3331033213021103-3121012200312032-0111302102020100-2110033001021302-3231323300322103-3311322122012121)
- trusted_clients.http_header.headers

<a id="canonical-1111321020000232-2001322220322101-2031322330032332-3002031312121322-0222313323122201-2003302130220203-2312303300021032-1122133302122310"></a>

Type: `"list"`. Computed.

List of HTTP header name and value pairs.

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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-0212000013112212-0120212133120111-0303122232000113-0022333012220130-2003001310123112-1211300332031131-0223020030101231-3121203333220222"></a>

### Direct properties for `trusted_clients.http_header.headers`

<a id="canonical-0032211222003320-2021201213110322-0333333202230233-0103023130302003-2132302021313202-3122030122132130-2302201210122032-0003131322202001"></a>

#### `trusted_clients.http_header.headers.exact` property

Type: `"string"`. Computed.

Exclusive with \[presence regular expression\] Header value to match exactly.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-1332311313210200-3210031130301112-1232321123301202-3002211333331200-0302233210100221-2330220231122122-0133133301121102-3210222203332323"></a>

<a id="canonical-2012120002203313-2302100323003223-1122103121201311-0012013122322210-3322220002222213-2232211100331312-1101013133203113-0223101011203311"></a>

#### `trusted_clients.http_header.headers.invert_match` property

Type: `"bool"`. Computed.

Invert the result of the match to detect missing header or non-matching value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0230300301213333-3323021302232030-0202313221222211-2332210321201030-1000230322022210-0323223011211330-3100031332201320-2003122231121223"></a>

<a id="canonical-3200212230122130-1110002111112303-3302131202323301-2211123212213332-0123101220202231-0210213230211210-1222333223100130-0001300121312332"></a>

#### `trusted_clients.http_header.headers.name` property

Type: `"string"`. Computed.

Name. Name of the header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1322332010222223-2210100110130303-3333033331312311-1332100201332133-1003131311321031-2013102132311310-2012212122301232-0033000230020300"></a>

<a id="canonical-3110213130222331-0012101111330022-2133033033323221-0323133222333023-3021211101123212-2332231233300023-0213222103032201-1010122322100010"></a>

#### `trusted_clients.http_header.headers.presence` property

Type: `"bool"`. Computed.

Exclusive with \[exact regular expression\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1023220201111123-0211333301213000-0222202121012220-0001300030101120-3320113320233102-1332322032012330-0222102100330130-3100331013021233"></a>

<a id="canonical-0332033023232201-3020313002013112-0200100230033312-3202330133013031-0110131201310310-1131011330023320-0102332302033000-2212113001211032"></a>

#### `trusted_clients.http_header.headers.regex` property

Type: `"string"`. Computed.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3223231021000022-1323110230023303-2001333120313120-2313222332030222-3132312200330212-2013103300312311-0321130232003030-3331121232230020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- trusted_clients.metadata

<a id="canonical-0011310032132021-3120112111333121-2311311133330201-2132002122331100-2100020011000103-1120001301222033-2320022030202302-3201031223221021"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0333333200213302-1331220220303111-2131120200121233-0010313200220030-2201023233311302-2210121000031032-0020232112301310-3033022330033131"></a>

### Direct properties for `trusted_clients.metadata`

<a id="canonical-1300331013303003-0333303003012113-2200112003122302-3102010000110330-2001221222233231-0112230301203131-0302233302312101-1102222232111112"></a>

#### `trusted_clients.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0122110002222310-1103202023302132-1110032313230323-0022213120132310-2220112011300032-2330213203112120-2332030330030230-1330301310002312"></a>

<a id="canonical-2212133333112312-3310213213110302-1012112031233000-3002113130311323-3012221322113202-2002233302231232-2210320313312122-3113021322033211"></a>

#### `trusted_clients.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-3121002301000111-3032111303202230-0123013210102222-0022322133120202-2110121100211130-0132320310300120-2321303012320220-1103131131231233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.skip_processing` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- trusted_clients.skip_processing

<a id="canonical-2321322132001300-0303103233131111-0321033123120102-2030111330121302-3313013110203133-3313322102211310-0223121012231012-0101333231331332"></a>

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

<a id="canonical-3302120220301322-0303132200032323-0103130112023221-1022001202212200-0313022213103220-2013200021311220-1002322311331131-1302331323230003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.waf_skip_processing` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- trusted_clients.waf_skip_processing

<a id="canonical-3011013220020030-2133013130301101-3020312311223320-3221330002010120-3003001030122130-3221220202200313-1232020333021013-3000310001331330"></a>

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

<a id="canonical-1101232232111221-2200012223301301-3121031112233231-2313000312132330-2230301033331100-2213101132232022-2002003032201231-2132233222201102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_id_client_ip` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- user_id_client_ip

<a id="canonical-1103212001033132-1131132103332211-3110210103311010-2133102010102320-0331222110230320-0110210021332111-1203332023133231-3032231000330123"></a>

Type: `["object", {}]`. Computed.

\[OneOf: user\_id\_client\_ip, user\_identification\] Enable this option

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

- [user_id_client_ip](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1103212001033132-1131132103332211-3110210103311010-2133102010102320-0331222110230320-0110210021332111-1203332023133231-3032231000330123)
- [user_identification](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3023223112130033-0030213022332131-3112101302132332-0010202322122012-0033013021310002-2333323313013112-0330023220021223-0120320020303321)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030211022123120-0133111110122030-3220101023132303-2132132233213000-3210220030203103-2301322311111110-2301202312112311-2202100000202220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_identification` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- user_identification

<a id="canonical-3023223112130033-0030213022332131-3112101302132332-0010202322122012-0033013021310002-2333323313013112-0330023220021223-0120320020303321"></a>

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

<a id="canonical-0013102103113010-3000133200312020-3301101320121011-2211100302131333-3322202023011321-3123022321331031-0220313002002303-0123003031012121"></a>

### Direct properties for `user_identification`

<a id="canonical-1203000121123121-2003120000320013-3021233213033300-1013232122303212-0111330103100020-2132223322332030-2121132303023223-3022322332030222"></a>

#### `user_identification.name` property

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

<a id="canonical-2313102201031103-3220222022031223-3130302320322232-2230320103003021-1323301021313112-1223003313310012-0113220200313232-1032331323002332"></a>

<a id="canonical-2320121322033010-3130302222330201-2020023112211232-1102300012322301-2130303011120222-0301230300322231-1331133220333222-0111001021112220"></a>

#### `user_identification.namespace` property

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

<a id="canonical-1112103033033030-1122313210000230-2232121313322230-3102213211112031-1130310123321233-0010323102010332-1100201121022200-1011000130123230"></a>

<a id="canonical-1222031123313313-3033230230323122-1113130201010130-0123013032022301-2001232313002022-0023231023232033-2122012130321000-1020011023323331"></a>

#### `user_identification.tenant` property

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

<a id="canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- waf_exclusion

<a id="canonical-1323231122322113-0110133233002221-2233201022202133-0020023133332330-1021321132112132-3230311133323020-1111032013201301-0030010131101202"></a>

Type: `"single"`. Computed.

Configuration parameter for waf exclusion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-waf_exclusion_choice": "[\"waf_exclusion_inline_rules\",\"waf_exclusion_policy\"]"
}
```

<a id="canonical-2012023231221011-1211012110323012-3121032003001110-2232032220122330-3101202033011331-1320300120230120-0020302030211013-3031331111113111"></a>

### Direct properties for `waf_exclusion`

- [waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022): complete subsection reference.

- [waf_exclusion_policy](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0321010221131133-3112232201230130-1002303120223013-3111331130200211-2222333313020223-0221201123100030-2000031120023210-0330322333212330): complete subsection reference.

<a id="canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- waf_exclusion.waf_exclusion_inline_rules

<a id="canonical-0001133202332131-0222200331122130-1301233223001210-1132211132002002-1200002313032112-3200313003101013-1231311302133033-1222223031320123"></a>

Type: `"single"`. Computed.

A list of WAF exclusion rules that will be applied inline.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2131310000022202-1212201333003331-0023232313321120-3201203231010001-2002000012133300-0032131312231121-2320100101121330-2331010031211030"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules`

- [rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301): complete subsection reference.

<a id="canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- waf_exclusion.waf_exclusion_inline_rules.rules

<a id="canonical-2233201300313312-2201030313200332-0002310331323323-0302313112320311-3331210110030310-1100121221222031-1101113011111133-3010002111313021"></a>

Type: `"list"`. Computed.

An ordered list of WAF Exclusions specific to this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-1021333001303322-1112311131212022-1102232010310002-0101130000121020-1220302132301000-1310101000032302-1300202130312011-3222312123212102"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules`

- [any_domain](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3132002120003311-1331211223030211-2012300231100311-0211230102033330-2222002201013123-2001102322312111-2311300203301320-2233011121021303): complete subsection reference.

- [any_path](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3011001323222321-3133223031103001-2112320103312230-1313030102230103-0330110330321030-0022020300333113-3320123012121202-3331333000312211): complete subsection reference.

- [app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200): complete subsection reference.

<a id="canonical-2300330100121302-3001130111101122-2102111023330102-3233011002021310-0102231300330320-3123120222300021-3012000311232210-3200221223102323"></a>

<a id="canonical-1000220311231331-2101111032003211-0202101132102112-1310310211131312-1311332033213311-3021312322011031-1332211020212002-1103010121031222"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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

<a id="canonical-0031232100112320-0331031013023202-2300001101213222-2202221022112123-1130333110301300-1213333100211333-0121031113012313-1112221111332330"></a>

<a id="canonical-1211103320123301-1102103231031300-1102211212111230-2130220323130122-3203220331303323-0330021322300332-1032121321001132-1112002333332032"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.expiration_timestamp` property

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Additional upstream details:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

- [metadata](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1132002310133120-2021321221110121-0021011222300122-1110320101123002-3122011332033011-1133200310300121-2333230123010330-1301000200133122): complete subsection reference.

<a id="canonical-3011102033200223-1003101133031223-2321110000020311-0313230000001031-2021302132330212-3212123013101222-2020101021313233-3130000220031110"></a>

<a id="canonical-0133300331033223-0201000331230221-1033100131301221-2122011331020133-3332302031230201-1002102332022302-3113020131020111-3312313211302133"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.methods` property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3123222003201303-1320033333030111-3203321302233033-1232313323232302-3320121112003303-1213032222120013-3222011223232103-2010212021320021"></a>

<a id="canonical-1232110332312102-2123023303310103-2012011301223103-1212031013223003-3010103102011023-1232213000232023-0122131213202122-3112200103231213"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.path_prefix` property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0031310112211133-2323113111113311-0222031003301323-2130121233320011-0320123313203222-0300001133030331-1322023301323102-2301230311332302"></a>

<a id="canonical-3103213331321120-2030001133332030-0200110001231101-3132101232113231-0100200003332310-3201310323001322-3012212112011331-1200113030022133"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.path_regex` property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\] Define the regular expression for the path. For example, the regular expression
^/.\*$ will match on all paths.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2122012221212130-2110302233201203-2222320331320201-2133212100312211-0330020023211221-3030302000231321-2113320322312113-2320320132012232"></a>

<a id="canonical-1122133023001211-0202300120202102-2122021123232202-1011010030223300-1310002331333331-2230220023313323-2212221001032313-3101210210021330"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
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

- [waf_skip_processing](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1032010213111131-0112203220133133-1023020203013112-3001302122333112-3332030132000301-3330222033303230-3313221321211113-1023220000033122): complete subsection reference.

<a id="canonical-3132002120003311-1331211223030211-2012300231100311-0211230102033330-2222002201013123-2001102322312111-2311300203301320-2233011121021303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_domain

<a id="canonical-3132301232012110-3130301331303131-1032033230303200-3021010202203100-1003130123330103-0011020231212310-2132330122320331-0222131130013131"></a>

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

<a id="canonical-3011001323222321-3133223031103001-2112320103312230-1313030102230103-0330110330321030-0022020300333113-3320123012121202-3331333000312211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.any_path` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_path

<a id="canonical-1130020210123220-3323310212023102-2310213022323101-0201200331332103-2120001110211332-3110203011312221-0323103122321110-3101101230321110"></a>

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

<a id="canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control

<a id="canonical-3323021110332122-0032211020313131-3013102333232221-3222122131121012-0002311131010020-2303131310122000-0201322312323200-1120233320010321"></a>

Type: `"single"`. Computed.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1031313012122132-1312312031023023-0011023222111322-3301133022321020-1232031110103212-1100323010102232-3300000303122113-3331303012233022"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control`

- [exclude_attack_type_contexts](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-2111021103111010-2200232001233022-1313121320012230-2001203300032101-1131321023110033-2113101020222323-1323120233312112-1312211020333331): complete subsection reference.

- [exclude_bot_name_contexts](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0311123311100230-3102011311332200-2103010021121131-3320123011302201-2031213133123331-1301212330011221-1100130300223300-3201301012320211): complete subsection reference.

- [exclude_signature_contexts](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3123002100101323-3100130021102213-2332233320213022-1120102222231300-2012132220113223-0311001130122230-3120231221010102-3300013312330112): complete subsection reference.

- [exclude_violation_contexts](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-2011222110302220-2303223101032203-3312222312001320-1110120000023210-0220032212211332-0100000003301221-1001233103030133-0213213231031311): complete subsection reference.

<a id="canonical-2111021103111010-2200232001233022-1313121320012230-2001203300032101-1131321023110033-2113101020222323-1323120233312112-1312211020333331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-0322023303120120-0213122113311232-2133023310330033-2222021301002322-0231212032301101-0033000102000322-1133332323100233-1102322211031233"></a>

Type: `"list"`. Computed.

Exclude an entire attack type only in the named context. For migrated per-parameter exceptions,
prefer this over signature-ID exclusions because one payload can trigger several signatures;
unrelated parameters and attack types remain protected.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3211021000313033-2301120330330013-0123230221022103-1103122133030021-1203232223021321-3300321031321010-3132132301222001-1301132211120120"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts`

<a id="canonical-0021013303033312-0222021321102001-0111131221001002-0030221002310002-0312312030213200-2123122011132312-1320222010330322-3000001233133201"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts.context` property

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Additional upstream details:

Exclusion scope. Use CONTEXT\_PARAMETER with context\_name for one parameter, CONTEXT\_COOKIE for
one cookie, or CONTEXT\_ANY only for an intentionally global scope.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1211011320020101-0320221310103310-3102222331300300-2311221121011121-3200331021330033-3313321310022310-1012330010333321-3110203232020133"></a>

<a id="canonical-0321133003222113-1331312000111023-3231131331300201-0110011123220120-2221303032221330-0010021212103110-1000102132020211-3130301223121121"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts.context_name` property

Type: `"string"`. Computed.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2323021011020210-3221032332312002-2310312313222002-0130310023223131-3120223012310011-3130112201201033-1213000010013221-2012100200220011"></a>

<a id="canonical-2213110202113110-3002123230132023-2000100202212233-2222010111123323-2131113210022200-3210012012333113-3000311131122332-3102201103102310"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` property

Type: `"string"`. Computed.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of all Attack Types ATTACK\_TYPE\_NONE ATTACK\_TYPE\_NON\_BROWSER\_CLIENT
ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE
ATTACK\_TYPE\_DETECTION\_EVASION ATTACK\_TYPE\_VULNERABILITY\_SCAN
ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS..
Possible values are \`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Additional upstream details:

Attack-type enum excluded in this context, for example ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING. Other
attack types remain enforced.

Receipt-pinned upstream constraints:

```json
{
  "default": "ATTACK_TYPE_NONE",
  "enum": [
    "ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0311123311100230-3102011311332200-2103010021121131-3320123011302201-2031213133123331-1301212330011221-1100130300223300-3201301012320211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-3020213201322123-1231332101222223-2202000021012131-2101120033211130-1001213121223010-1302333022021123-1331220232122332-1022232212222311"></a>

Type: `"list"`. Computed.

Bot Names to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0302311122213322-3003120302303221-1023231101002100-1330232030323103-0321023203020331-2223330102212220-1233113211330013-2223323221211103"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts`

<a id="canonical-0212203200131123-1102211311100233-3313220323312010-3222101211313030-1112110203023313-3122110312232100-2302200230203222-2120032012010220"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` property

Type: `"string"`. Computed.

Bot Name. Human-readable name for the resource

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

<a id="canonical-3123002100101323-3100130021102213-2332233320213022-1120102222231300-2012132220113223-0311001130122230-3120231221010102-3300013312330112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-0010210221303301-2232031113102322-2302223031301333-2310021010232120-3301021333132003-0321020321122033-0321133322313131-1121301200132000"></a>

Type: `"list"`. Computed.

Signature IDs to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0031130211201321-2022322313310033-3012113333312122-0210213023202122-2333320112123123-0211002131010233-1122320233133103-0331203100020000"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts`

<a id="canonical-0112022113112112-2211221031320132-2223323022302211-2312202232130132-2102210030111021-1111110320233230-3101001021231120-0113332120201110"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts.context` property

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Additional upstream details:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1331323212000121-0013330031220033-2330011001233032-3010211310332011-1033330120333303-0210323133330032-0320012330223022-3233023031110221"></a>

<a id="canonical-2012322302123013-0133333101110020-0323300000032223-3211332110211113-3320110322001000-1222003002100102-1210031323212222-0300012213221132"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts.context_name` property

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0101200310012012-1330111233203311-3002030200011021-1000220333300110-1011030030322012-3110122312331213-3232120322222031-0111213003323031"></a>

<a id="canonical-1002213333100112-2223110023300230-0033032233232032-0103023113220331-0103131333221101-2210002301101210-3223122122312023-2032213310333032"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts.signature_id` property

Type: `"number"`. Computed.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 299999999,
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
    "ves.io.schema.rules.uint32.lte": "299999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  }
}
```

<a id="canonical-2011222110302220-2303223101032203-3312222312001320-1110120000023210-0220032212211332-0100000003301221-1001233103030133-0213213231031311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-0033311111112212-2021000202001303-3321132031201001-3123031313000122-2111333130013100-2201301321311313-2022021231020132-3031122113221323"></a>

Type: `"list"`. Computed.

Violations to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1032232231101030-2213012221300011-2110000333123033-0021001033013023-1021033200012222-0233222301010000-0120202110221123-0211222331302000"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts`

<a id="canonical-1331130131310233-1021320033110113-0100223012032303-2330232233230230-1100013002233220-0112332120033311-0000313103202302-2320031301131032"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts.context` property

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Additional upstream details:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2001031100201002-2121021123232012-2222031331030120-0103102202010233-1001103230313020-0220213100332201-2332030100013323-2311110021212203"></a>

<a id="canonical-3020323120101330-0120232222111102-1003323031212103-3220312001030302-1323310120101000-0013320121322303-1003102000122333-1202033012111222"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts.context_name` property

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1101031011311112-2203022211320311-3322211133230212-0122120233331121-2312112222131303-2021013202131300-0132220102312222-1130230101000231"></a>

<a id="canonical-3121102312300201-0032313133133030-3232302111120121-2301301130101310-3103123032112322-1100302222233321-0112022320022321-1101311213313101"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` property

Type: `"string"`. Computed.

\[Enum:
VIOL\_NONE|VIOL\_FILETYPE|VIOL\_METHOD|VIOL\_MANDATORY\_HEADER|VIOL\_HTTP\_RESPONSE\_STATUS|VIOL\_REQUEST\_MAX\_LENGTH|VIOL\_FILE\_UPLOAD|VIOL\_FILE\_UPLOAD\_IN\_BODY|VIOL\_XML\_MALFORMED|VIOL\_JSON\_MALFORMED|VIOL\_ASM\_COOKIE\_MODIFIED|VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS|VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE|VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT|VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST|VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION|VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS|VIOL\_EVASION\_DIRECTORY\_TRAVERSALS|VIOL\_MALFORMED\_REQUEST|VIOL\_EVASION\_MULTIPLE\_DECODING|VIOL\_DATA\_GUARD|VIOL\_EVASION\_APACHE\_WHITESPACE|VIOL\_COOKIE\_MODIFIED|VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS|VIOL\_EVASION\_IIS\_BACKSLASHES|VIOL\_EVASION\_PERCENT\_U\_DECODING|VIOL\_EVASION\_BARE\_BYTE\_DECODING|VIOL\_EVASION\_BAD\_UNESCAPE|VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST|VIOL\_ENCODING|VIOL\_COOKIE\_MALFORMED|VIOL\_GRAPHQL\_FORMAT|VIOL\_GRAPHQL\_MALFORMED|VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\]
List of all supported Violation Types VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER
VIOL\_HTTP\_RESPONSE\_STATUS VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD
VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED
VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS.. Possible values are \`VIOL\_NONE\`,
\`VIOL\_FILETYPE\`, \`VIOL\_METHOD\`, \`VIOL\_MANDATORY\_HEADER\`, \`VIOL\_HTTP\_RESPONSE\_STATUS\`,
\`VIOL\_REQUEST\_MAX\_LENGTH\`, \`VIOL\_FILE\_UPLOAD\`, \`VIOL\_FILE\_UPLOAD\_IN\_BODY\`,
\`VIOL\_XML\_MALFORMED\`, \`VIOL\_JSON\_MALFORMED\`, \`VIOL\_ASM\_COOKIE\_MODIFIED\`,
\`VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE\`,
\`VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT\`, \`VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION\`,
\`VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS\`,
\`VIOL\_EVASION\_DIRECTORY\_TRAVERSALS\`, \`VIOL\_MALFORMED\_REQUEST\`,
\`VIOL\_EVASION\_MULTIPLE\_DECODING\`, \`VIOL\_DATA\_GUARD\`, \`VIOL\_EVASION\_APACHE\_WHITESPACE\`,
\`VIOL\_COOKIE\_MODIFIED\`, \`VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS\`,
\`VIOL\_EVASION\_IIS\_BACKSLASHES\`, \`VIOL\_EVASION\_PERCENT\_U\_DECODING\`,
\`VIOL\_EVASION\_BARE\_BYTE\_DECODING\`, \`VIOL\_EVASION\_BAD\_UNESCAPE\`,
\`VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST\`, \`VIOL\_ENCODING\`,
\`VIOL\_COOKIE\_MALFORMED\`, \`VIOL\_GRAPHQL\_FORMAT\`, \`VIOL\_GRAPHQL\_MALFORMED\`,
\`VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\`. Defaults to \`VIOL\_NONE\`.

Additional upstream details:

List of all supported Violation Types

VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER VIOL\_HTTP\_RESPONSE\_STATUS
VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED
VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS
VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT
VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION
VIOL\_HTTP\_PROTOCOL\_CRLF\_CHARACTERS\_BEFORE\_REQUEST\_START
VIOL\_HTTP\_PROTOCOL\_NO\_HOST\_HEADER\_IN\_HTTP\_1\_1\_REQUEST
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_PARAMETERS\_PARSING
VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS
VIOL\_HTTP\_PROTOCOL\_CONTENT\_LENGTH\_SHOULD\_BE\_A\_POSITIVE\_NUMBER
VIOL\_EVASION\_DIRECTORY\_TRAVERSALS VIOL\_MALFORMED\_REQUEST VIOL\_EVASION\_MULTIPLE\_DECODING
VIOL\_DATA\_GUARD VIOL\_EVASION\_APACHE\_WHITESPACE VIOL\_COOKIE\_MODIFIED
VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS VIOL\_EVASION\_IIS\_BACKSLASHES
VIOL\_EVASION\_PERCENT\_U\_DECODING VIOL\_EVASION\_BARE\_BYTE\_DECODING VIOL\_EVASION\_BAD\_UNESCAPE
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_FORMDATA\_REQUEST\_PARSING
VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST
VIOL\_HTTP\_PROTOCOL\_HIGH\_ASCII\_CHARACTERS\_IN\_HEADERS VIOL\_ENCODING VIOL\_COOKIE\_MALFORMED
VIOL\_GRAPHQL\_FORMAT VIOL\_GRAPHQL\_MALFORMED VIOL\_GRAPHQL\_INTROSPECTION\_QUERY.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIOL_NONE",
  "enum": [
    "VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START",
    "VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1132002310133120-2021321221110121-0021011222300122-1110320101123002-3122011332033011-1133200310300121-2333230123010330-1301000200133122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- waf_exclusion.waf_exclusion_inline_rules.rules.metadata

<a id="canonical-1112313103131012-2112101110312021-0321020020031132-1200111000132031-1210201222121232-0021031103221321-3200003201232303-0021102021222032"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0323322103201013-2330121130013321-3221013221330122-0031012133203123-2321203002023033-2232131311221213-3302121201220230-0000130213121333"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.metadata`

<a id="canonical-0133210121121331-3013130210221303-1112032001023310-2303033220100311-0312123301200111-1212201310103200-0311332112213021-3220213332030120"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1223331011121000-3111120222022021-1130031122101212-0222021323311103-3333322313033211-2013032321201321-3132102032033121-3333313012323301"></a>

<a id="canonical-0032313113031010-2132120012111112-0111120231231322-2013312013330313-3112100020333021-2010233203310120-2313230031201230-3332303010203220"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-1032010213111131-0112203220133133-1023020203013112-3001302122333112-3332030132000301-3330222033303230-3313221321211113-1023220000033122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing

<a id="canonical-3210100213303332-3313031222303332-3310223203100312-1023230030213122-3210202320201131-2230331321231112-2332303131112313-2013002030203021"></a>

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

<a id="canonical-0321010221131133-3112232201230130-1002303120223013-3111331130200211-2222333313020223-0221201123100030-2000031120023210-0330322333212330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_policy` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- waf_exclusion.waf_exclusion_policy

<a id="canonical-2121321300001002-1213202131231302-2131220321001302-3130122001220200-2122301223312011-3022130111233110-3303110320100010-3301002120212311"></a>

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

<a id="canonical-0312010112022031-0032313023023211-1303000020012231-2310000223233130-1303322332302310-3221210122133113-3031323103230023-0332032002213322"></a>

### Direct properties for `waf_exclusion.waf_exclusion_policy`

<a id="canonical-3331030213310010-1231030210010012-0031012033321011-1030320101202313-3001012303332300-3033013100300011-0200131301122221-2210130230023003"></a>

#### `waf_exclusion.waf_exclusion_policy.name` property

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

<a id="canonical-1332221010021312-0222232233031033-0131123321312032-2122303022221110-3131012123123211-2211103030111332-1102311301203233-2231031230331212"></a>

<a id="canonical-3221232010110323-1331101001322310-1130313113323003-1113002201302101-3311001313031111-0031203111330120-2200011301211022-1112302022112131"></a>

#### `waf_exclusion.waf_exclusion_policy.namespace` property

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

<a id="canonical-3131130330100212-2202311003120320-0000110223120001-3222022130320120-2122102212303310-2132303013012032-1011001321210223-0313112100200133"></a>

<a id="canonical-1312022010203231-1211222310031203-1011011212211020-1031210103132230-2123331303111222-0132323220322032-1223000221312212-0113022322113301"></a>

#### `waf_exclusion.waf_exclusion_policy.tenant` property

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
