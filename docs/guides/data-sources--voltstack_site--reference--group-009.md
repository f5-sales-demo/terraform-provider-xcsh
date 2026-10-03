---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-3331121021332202-3202102211212020-3120012102332132-3221133220202212-3230023231211301-3111003231010330-1201221210133231-2310331132033002"></a>

## Next pages — enable_vm / 212222322202 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1132230233133203-3131223100302000-2120313313123100-3220310332231001-0231212333120223-3101313210021001-1002131331120000-1311031032021230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320102130021002-0231332312122000-3301312200323021-0323302222033102-2133331321032030-3223332232030133-0010233233210303-1222021121211231"></a>

## k8s_cluster — k8s_cluster / 023210323211 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- k8s_cluster

<a id="canonical-2311100022221013-1222221000021030-1011010103103301-1213133132103320-2222320100220012-2232023100002321-1102310223031122-1100212010322121"></a>

Type: `"single"`. Computed.

\[OneOf: k8s\_cluster, no\_k8s\_cluster; Default: no\_k8s\_cluster\] Type establishes a direct
reference from one object(the referrer) to another(the referred). Such a reference is in form of
tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [k8s_cluster](data-sources--voltstack_site--reference--group-009.md#canonical-2311100022221013-1222221000021030-1011010103103301-1213133132103320-2222320100220012-2232023100002321-1102310223031122-1100212010322121)
- [no_k8s_cluster](data-sources--voltstack_site--reference--group-009.md#canonical-1013201012231322-1122033102111202-3123031303112213-3001013311103230-3000012313111132-3211203112221102-3332233221001323-0030300122300212)

Select alternatives according to the provider validators above.

<a id="canonical-3111121031313312-1300321000103230-1302133202133211-2330123201232023-3000323123010122-2001123212230320-0320220102032132-2000021101100213"></a>

## Direct properties — k8s_cluster / 023210323211 / 3

<a id="canonical-1121300331321002-2321103130333332-3320310110220000-3303120322323300-0112201012223211-3230332022210201-1312231032120121-0220011111100010"></a>

<a id="canonical-0033200022311122-1323022213232013-2322132303202330-1331000313212103-2013310101203312-1313101332011312-2312212331233030-2000332120001032"></a>

## name property — k8s_cluster / 023210323211 / 4

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

<a id="canonical-1023122021113322-2302231321123203-3312331321331103-3302021331223023-0022131311313333-3121132321111100-1002201012312333-0032100120302310"></a>

<a id="canonical-1200023320331010-1123000302310133-2303303120302202-0112322313221230-2010133233301331-2023230221113213-2213021232202311-3131212110010110"></a>

## namespace property — k8s_cluster / 023210323211 / 5

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

<a id="canonical-3033223332302031-1210101210312233-0000210303003233-1333101011231110-3011123221200213-0102203321223223-1101031112233120-1021022322011201"></a>

<a id="canonical-0300133333010021-3103111313200312-0113230230000331-3332322232223302-2123311030300333-1233302132100001-0231113313220203-1311332322023302"></a>

## tenant property — k8s_cluster / 023210323211 / 6

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

<a id="canonical-2122321121232011-1233203012310030-0333201230223003-0131223231020003-2120113313101303-2022201233111002-3332003202221021-3210113312020010"></a>

## Next pages — k8s_cluster / 023210323211 / 7

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3320131330201111-2130101200102000-2332212122311310-2211301113232033-3130213303132011-2010111113113320-3322320132110020-1012320321211130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223112230120310-1203213131230022-2331320312213330-0311110311331000-1312201220031132-3200021220000030-2300102021222021-3101133110322222"></a>

## kubernetes_upgrade_drain — kubernetes_upgrade_drain / 131320020031 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- kubernetes_upgrade_drain

<a id="canonical-3220310013233131-2033032033210030-3200231302112233-1020312221202230-3123222312203212-0332321333211312-3111110012220210-2002121300332233"></a>

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

<a id="canonical-3120323232030200-3321100323102023-3230122013133200-3230330232213223-0132200122320102-0033220020330310-0221203331300011-0310032021323230"></a>

## Direct properties — kubernetes_upgrade_drain / 131320020031 / 3

- [disable_upgrade_drain](data-sources--voltstack_site--reference--group-009.md#canonical-3202220213021313-1210331303010123-0002003222230311-3222323110032311-1022002203201103-2232102331221222-1100330230201232-2222212120323312): complete subsection reference.

- [enable_upgrade_drain](data-sources--voltstack_site--reference--group-009.md#canonical-2201032123122320-1112111213321322-2202220303331033-0021102101122331-2210202201331321-2020322203313132-2000332220310222-2222202210100102): complete subsection reference.

<a id="canonical-0233313202231123-2200230113112022-3300121123221011-3332031233010132-0023200100123212-0122333101202111-3221100011030010-1320022120132212"></a>

## Next pages — kubernetes_upgrade_drain / 131320020031 / 4

- [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--voltstack_site--reference--group-009.md#canonical-3202220213021313-1210331303010123-0002003222230311-3222323110032311-1022002203201103-2232102331221222-1100330230201232-2222212120323312)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--voltstack_site--reference--group-009.md#canonical-2201032123122320-1112111213321322-2202220303331033-0021102101122331-2210202201331321-2020322203313132-2000332220310222-2222202210100102)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3202220213021313-1210331303010123-0002003222230311-3222323110032311-1022002203201103-2232102331221222-1100330230201232-2222212120323312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231202123333113-0313332321022330-0133001003003023-1213111122012023-3311110333330110-3030111013230101-2210012133030003-0030120212221113"></a>

## kubernetes_upgrade_drain.disable_upgrade_drain — disable_upgrade_drain / 133123000003 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [kubernetes_upgrade_drain](data-sources--voltstack_site--reference--group-009.md#canonical-3320131330201111-2130101200102000-2332212122311310-2211301113232033-3130213303132011-2010111113113320-3322320132110020-1012320321211130)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-2131000111020031-3033113012121322-1231100321301131-1102300210332111-3212331002112233-0211333301133232-0220312303312312-0221110031111103"></a>

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

<a id="canonical-2203213203300221-3320012022013112-1202230312222133-0020203120312300-3303100001122011-1012022300000301-2230132123230333-0221321202201330"></a>

## Direct properties — disable_upgrade_drain / 133123000003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101020222002103-2323220111330330-1003033122320133-3011130012213210-2323313000322123-0023330013101022-3223112231331021-1101012302233022"></a>

## Next pages — disable_upgrade_drain / 133123000003 / 4

- [kubernetes_upgrade_drain](data-sources--voltstack_site--reference--group-009.md#canonical-3320131330201111-2130101200102000-2332212122311310-2211301113232033-3130213303132011-2010111113113320-3322320132110020-1012320321211130)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2201032123122320-1112111213321322-2202220303331033-0021102101122331-2210202201331321-2020322203313132-2000332220310222-2222202210100102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213000102202333-1001322122013203-1003301001131323-2332100231030302-0020221101232113-3001331132330112-3301302101331031-0030000213211130"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain — enable_upgrade_drain / 021000012333 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [kubernetes_upgrade_drain](data-sources--voltstack_site--reference--group-009.md#canonical-3320131330201111-2130101200102000-2332212122311310-2211301113232033-3130213303132011-2010111113113320-3322320132110020-1012320321211130)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-2202322220010100-1033113033100021-0111011103120110-1033223213200021-0320320103210321-1013003310001113-2230101102300120-0232131210100000"></a>

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

<a id="canonical-2313133000032300-1202202110310112-0222202031222012-3121013030100013-0023003222002003-2321321102213212-1333011310102030-2320111303023233"></a>

## Direct properties — enable_upgrade_drain / 021000012333 / 3

- [disable_vega_upgrade_mode](data-sources--voltstack_site--reference--group-009.md#canonical-0012003213230202-1300310202231110-0121010321001003-1313301100112300-0212201220112011-3202232332203012-1221121011001113-3323131203232221): complete subsection reference.

<a id="canonical-2233021232211000-2313231113211123-1003130133121011-2322230102121022-3321023303201223-1200131220102121-1010112101100011-3103131112103000"></a>

<a id="canonical-1121103123213031-3300120123301012-3310020021332222-3302130111212001-2223022112101320-2012020201011230-0110012320231033-3223100112200220"></a>

## drain_max_unavailable_node_count property — enable_upgrade_drain / 021000012333 / 4

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

<a id="canonical-0103012313301000-3321321020002200-1010303320232100-2110320301120201-2331130111000323-2133313000101101-1103322130210321-3303222033000220"></a>

<a id="canonical-1211211231003213-0321103302030001-0102012232020033-2200000121001102-1002013223112201-1200102003123311-2223120003121212-1303102031220130"></a>

## drain_max_unavailable_node_percentage property — enable_upgrade_drain / 021000012333 / 5

Type: `"number"`. Computed.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-2203313020331100-0003312313102232-3203333323030032-3213111101302322-0222213000212021-2003021302302203-3213003011303013-3210110003110022"></a>

<a id="canonical-3102030121201120-2202212103000133-1233130301023310-3030123113012132-1200121221303100-0330310331211233-3100003330333322-0202022021210302"></a>

## drain_node_timeout property — enable_upgrade_drain / 021000012333 / 6

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

- [enable_vega_upgrade_mode](data-sources--voltstack_site--reference--group-009.md#canonical-3003003221302111-0211121220333332-0031002121122021-2231131303223231-2321033232012103-3202103022123230-1121320330321020-1000001221033021): complete subsection reference.

<a id="canonical-1133202020300113-3122221322021100-1003101221010003-3220300231302003-2031100003122212-1012012301133123-2211132203121233-1332123223331130"></a>

## Next pages — enable_upgrade_drain / 021000012333 / 7

- [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--voltstack_site--reference--group-009.md#canonical-0012003213230202-1300310202231110-0121010321001003-1313301100112300-0212201220112011-3202232332203012-1221121011001113-3323131203232221)
- [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--voltstack_site--reference--group-009.md#canonical-3003003221302111-0211121220333332-0031002121122021-2231131303223231-2321033232012103-3202103022123230-1121320330321020-1000001221033021)
- [kubernetes_upgrade_drain](data-sources--voltstack_site--reference--group-009.md#canonical-3320131330201111-2130101200102000-2332212122311310-2211301113232033-3130213303132011-2010111113113320-3322320132110020-1012320321211130)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0012003213230202-1300310202231110-0121010321001003-1313301100112300-0212201220112011-3202232332203012-1221121011001113-3323131203232221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111031202103233-3002103003101321-2010120303101213-1330031213233302-0331030112013133-0333210230233321-0100020011033010-0032313302031101"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — disable_vega_upgrade_mode / 322010032131 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [kubernetes_upgrade_drain](data-sources--voltstack_site--reference--group-009.md#canonical-3320131330201111-2130101200102000-2332212122311310-2211301113232033-3130213303132011-2010111113113320-3322320132110020-1012320321211130)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--voltstack_site--reference--group-009.md#canonical-2201032123122320-1112111213321322-2202220303331033-0021102101122331-2210202201331321-2020322203313132-2000332220310222-2222202210100102)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-2121202231320023-1111232023232101-3320330020003121-0130103012033211-3312230303300232-0302321001332021-0100230203220031-3321003202232210"></a>

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

<a id="canonical-1103231130322321-2203111122020011-3100122012003230-0110122030322030-1002003231333012-2101101322311212-1221211223301121-0023321123002212"></a>

## Direct properties — disable_vega_upgrade_mode / 322010032131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103320211130000-1320121221002113-2310130032000002-2312230121013012-0122132313300103-1210320031322330-0232230020003000-1323230313211220"></a>

## Next pages — disable_vega_upgrade_mode / 322010032131 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--voltstack_site--reference--group-009.md#canonical-2201032123122320-1112111213321322-2202220303331033-0021102101122331-2210202201331321-2020322203313132-2000332220310222-2222202210100102)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3003003221302111-0211121220333332-0031002121122021-2231131303223231-2321033232012103-3202103022123230-1121320330321020-1000001221033021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101232031000330-2210112210210022-2023322203311210-0133003220010331-0012111213131102-1213103231031310-3021322212332311-0020313131112001"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — enable_vega_upgrade_mode / 123311201110 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [kubernetes_upgrade_drain](data-sources--voltstack_site--reference--group-009.md#canonical-3320131330201111-2130101200102000-2332212122311310-2211301113232033-3130213303132011-2010111113113320-3322320132110020-1012320321211130)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--voltstack_site--reference--group-009.md#canonical-2201032123122320-1112111213321322-2202220303331033-0021102101122331-2210202201331321-2020322203313132-2000332220310222-2222202210100102)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-3222032132011001-3111121202313333-2203202100221110-3020330130210030-3210301012333020-3020213013221003-3010102310003100-3010330222133000"></a>

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

<a id="canonical-3303033001210302-0131000130323320-2012130020330031-0010132202300103-3301311131113311-2212113313012202-0301220333031202-3022212333020220"></a>

## Direct properties — enable_vega_upgrade_mode / 123311201110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121301233021222-2300300312011001-1010021223101031-3200133310133322-1113330310210003-3133111023101103-2033321201333210-0120001320121003"></a>

## Next pages — enable_vega_upgrade_mode / 123311201110 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--voltstack_site--reference--group-009.md#canonical-2201032123122320-1112111213321322-2202220303331033-0021102101122331-2210202201331321-2020322203313132-2000332220310222-2222202210100102)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331130130110133-2212320102202102-3103220202223002-0022210323000100-2311033311310110-0320132312200303-0310010312111211-3320110132222301"></a>

## local_control_plane — local_control_plane / 122122000003 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- local_control_plane

<a id="canonical-0330033002033320-3101210200232331-2102133133011002-2223121111202013-2220213221131111-2010000010323221-1123001101312110-0312131303301320"></a>

Type: `"single"`. Computed.

\[OneOf: local\_control\_plane, no\_local\_control\_plane; Default: no\_local\_control\_plane\]
Enable local control plane for L3VPN, SRV6, EVPN etc.

Upstream description:

Enable local control plane for L3VPN, SRV6, EVPN etc.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_vn\",\"outside_vn\"]"
}
```

OneOf alternatives in this subsection:

- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0330033002033320-3101210200232331-2102133133011002-2223121111202013-2220213221131111-2010000010323221-1123001101312110-0312131303301320)
- [no_local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-1002001200332030-1120100012133311-3302033110222333-1033022120110300-3020111201310130-2303203020323123-0121331333222122-1200321120212102)

Select alternatives according to the provider validators above.

<a id="canonical-2020021203112022-3300001003320002-3011201313102032-0131032320012223-2031112323030231-3323132310332012-1130000103200130-1211121030302120"></a>

## Direct properties — local_control_plane / 122122000003 / 3

- [bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100): complete subsection reference.

- [inside_vn](data-sources--voltstack_site--reference--group-009.md#canonical-0130201121003322-1033330032323001-0032001100333232-3213223220333223-1300030112323300-2202310100232112-0222103313212322-0110022010301010): complete subsection reference.

- [outside_vn](data-sources--voltstack_site--reference--group-009.md#canonical-3231111032012121-2322313113231303-1010332200001131-2302202320230012-3132121121003131-1233002013011133-3330003222212231-2102020101131222): complete subsection reference.

<a id="canonical-2112313021313111-3020001001003010-2312201312301332-2020012310132122-3102211323232203-0221113320011202-1001210230033333-2123023321300103"></a>

## Next pages — local_control_plane / 122122000003 / 4

- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.inside_vn](data-sources--voltstack_site--reference--group-009.md#canonical-0130201121003322-1033330032323001-0032001100333232-3213223220333223-1300030112323300-2202310100232112-0222103313212322-0110022010301010)
- [local_control_plane.outside_vn](data-sources--voltstack_site--reference--group-009.md#canonical-3231111032012121-2322313113231303-1010332200001131-2302202320230012-3132121121003131-1233002013011133-3330003222212231-2102020101131222)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211033022030301-3133021123233000-1300003203000023-2032312321331222-0020232231121013-3100100333333012-3012313101132030-0331011333131201"></a>

## local_control_plane.bgp_config — bgp_config / 223333322313 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- local_control_plane.bgp_config

<a id="canonical-2211323101210230-0302032102013203-2001013100222011-2202101210212033-1122311220332233-0202221211202200-0001120111201130-0230321212003033"></a>

Type: `"single"`. Computed.

BGP Configuration. BGP configuration parameters.

Upstream description:

BGP configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0231213113021301-2013122323032223-3320120321301222-2101122222233010-3000201220220231-1233231031003233-0110013031211300-2313230203222013"></a>

## Direct properties — bgp_config / 223333322313 / 3

<a id="canonical-1202000210132021-3002300021301330-3103311111332200-2012311333323331-2002223202112030-3121012203013221-1011330113313132-2001311000011223"></a>

<a id="canonical-3202001102323201-2232321321120100-3000330111030012-1230121012233332-0002213330302312-0213023001103130-2302003110002331-0103100013023031"></a>

## asn property — bgp_config / 223333322313 / 4

Type: `"number"`. Computed.

ASN. Autonomous System Number.

Upstream description:

Autonomous System Number.

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
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300): complete subsection reference.

<a id="canonical-0020121202010312-3133213000001212-0130113200220101-0031010022203011-1030211201211320-0232111101111322-2012112013203132-2121331023103202"></a>

## Next pages — bgp_config / 223333322313 / 5

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031322300233313-3130013301213113-1120230201210330-1133231122303131-3213002310313202-1111211112233231-0222312300231102-1310032332130332"></a>

## local_control_plane.bgp_config.peers — peers / 111230312320 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- local_control_plane.bgp_config.peers

<a id="canonical-2221321000032032-3311010002031203-3131111312112023-1031201320323022-0130200111113203-2031332331103333-2133201110032223-2211231121023301"></a>

Type: `"list"`. Computed.

Peers. BGP parameters for peer.

Upstream description:

BGP parameters for peer.

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-3103103220323100-3011312012312322-2003211112103112-1200303232011022-3103201112222222-0122111112322123-2122302301231223-0333313003311323"></a>

## Direct properties — peers / 111230312320 / 3

- [bfd_disabled](data-sources--voltstack_site--reference--group-009.md#canonical-2030030223112303-1330003220013303-3230230311303321-0322302102331103-1222303231123300-3320130003331233-0031322223210021-1333100022112201): complete subsection reference.

- [bfd_enabled](data-sources--voltstack_site--reference--group-009.md#canonical-3103202303233312-0130333322022003-3332201201110002-3230331312312022-2002022320230210-2100120222110300-1333122020032222-2022123330003020): complete subsection reference.

- [disable_spec](data-sources--voltstack_site--reference--group-009.md#canonical-1121132012200303-3002212112032202-2031330221132330-2103011113200033-0000303110023300-3302330230230031-2231103010221033-1222022223210310): complete subsection reference.

- [ebgp_multihop_disabled](data-sources--voltstack_site--reference--group-009.md#canonical-0302123211022230-0030120332220022-1110120120202230-0223010231032123-0130302103213210-1301003223221221-1031311021200230-3100322212030223): complete subsection reference.

- [ebgp_multihop_enabled](data-sources--voltstack_site--reference--group-009.md#canonical-2203031313131132-3330232311210023-1032321110003101-1201313233130132-1133102110111030-0113323002311332-0102032321111113-0002312011032310): complete subsection reference.

- [external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032): complete subsection reference.

<a id="canonical-3000231130110210-0121122123310202-2003203322311200-1321000300102303-2010221210220332-3130213020303323-3233300112303113-0323003131203112"></a>

<a id="canonical-0111000132020002-3313101231102120-1030200111233322-1013001313123013-3302020132333211-1332311320110301-3013002330122322-3323021333201303"></a>

## label property — peers / 111230312320 / 4

Type: `"string"`. Computed.

Label. Specify whether this peer should be.

Upstream description:

Specify whether this peer should be.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "labeling",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
  }
}
```

- [metadata](data-sources--voltstack_site--reference--group-009.md#canonical-1210200222013002-2333330303222230-3300230230131300-1232322120231233-3201210010102332-1132101220313021-0312330321333302-1133300320121213): complete subsection reference.

- [passive_mode_disabled](data-sources--voltstack_site--reference--group-009.md#canonical-3030031322003002-3310331113031030-1333030211320001-0011320003312210-1133203313103101-2203021020313121-3321113323232230-3002111200320020): complete subsection reference.

- [passive_mode_enabled](data-sources--voltstack_site--reference--group-009.md#canonical-1013002103211301-2302230310013122-0113201300210220-2300111003331321-0220020303021323-2322203102212012-2310110221101000-1010132131310313): complete subsection reference.

- [routing_policies](data-sources--voltstack_site--reference--group-009.md#canonical-2001023103302213-2103112012320003-2200301222330212-0131031010310000-0223223132103300-1232301311310232-2131320333110112-3233212131323232): complete subsection reference.

<a id="canonical-3123230113120023-2010020100103012-3330223323332020-0030033123223233-0330302232112312-3010210123031311-0331030320221103-3300202303302233"></a>

## Next pages — peers / 111230312320 / 5

- [local_control_plane.bgp_config.peers.bfd_disabled](data-sources--voltstack_site--reference--group-009.md#canonical-2030030223112303-1330003220013303-3230230311303321-0322302102331103-1222303231123300-3320130003331233-0031322223210021-1333100022112201)
- [local_control_plane.bgp_config.peers.bfd_enabled](data-sources--voltstack_site--reference--group-009.md#canonical-3103202303233312-0130333322022003-3332201201110002-3230331312312022-2002022320230210-2100120222110300-1333122020032222-2022123330003020)
- [local_control_plane.bgp_config.peers.disable_spec](data-sources--voltstack_site--reference--group-009.md#canonical-1121132012200303-3002212112032202-2031330221132330-2103011113200033-0000303110023300-3302330230230031-2231103010221033-1222022223210310)
- [local_control_plane.bgp_config.peers.ebgp_multihop_disabled](data-sources--voltstack_site--reference--group-009.md#canonical-0302123211022230-0030120332220022-1110120120202230-0223010231032123-0130302103213210-1301003223221221-1031311021200230-3100322212030223)
- [local_control_plane.bgp_config.peers.ebgp_multihop_enabled](data-sources--voltstack_site--reference--group-009.md#canonical-2203031313131132-3330232311210023-1032321110003101-1201313233130132-1133102110111030-0113323002311332-0102032321111113-0002312011032310)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [local_control_plane.bgp_config.peers.metadata](data-sources--voltstack_site--reference--group-009.md#canonical-1210200222013002-2333330303222230-3300230230131300-1232322120231233-3201210010102332-1132101220313021-0312330321333302-1133300320121213)
- [local_control_plane.bgp_config.peers.passive_mode_disabled](data-sources--voltstack_site--reference--group-009.md#canonical-3030031322003002-3310331113031030-1333030211320001-0011320003312210-1133203313103101-2203021020313121-3321113323232230-3002111200320020)
- [local_control_plane.bgp_config.peers.passive_mode_enabled](data-sources--voltstack_site--reference--group-009.md#canonical-1013002103211301-2302230310013122-0113201300210220-2300111003331321-0220020303021323-2322203102212012-2310110221101000-1010132131310313)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--reference--group-009.md#canonical-2001023103302213-2103112012320003-2200301222330212-0131031010310000-0223223132103300-1232301311310232-2131320333110112-3233212131323232)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2030030223112303-1330003220013303-3230230311303321-0322302102331103-1222303231123300-3320130003331233-0031322223210021-1333100022112201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102302330000332-3010232332303020-0232101233210233-0013030133222211-2203011203111223-0103322231000120-2100013120310031-0231232013210200"></a>

## local_control_plane.bgp_config.peers.bfd_disabled — bfd_disabled / 221020003132 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- local_control_plane.bgp_config.peers.bfd_disabled

<a id="canonical-3133320120300031-2331312202211130-2313301312003103-0021233210001310-3032223102231311-1301200103331102-0113231301030023-1233132301211331"></a>

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

<a id="canonical-0231133130032022-0232322232333103-0211011020033300-2110233023332210-1022101212110313-0132213200232310-2233332103013232-0002133310032032"></a>

## Direct properties — bfd_disabled / 221020003132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000222002000032-3103033020202132-1110102303203201-3000311103133333-1311131111332021-0322130020322310-1012013212111100-3310233310211000"></a>

## Next pages — bfd_disabled / 221020003132 / 4

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3103202303233312-0130333322022003-3332201201110002-3230331312312022-2002022320230210-2100120222110300-1333122020032222-2022123330003020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203220322010200-3021031121213011-3233033103111113-2103210303001202-3110223322322220-0320111230130300-0132112223111130-2301102300232010"></a>

## local_control_plane.bgp_config.peers.bfd_enabled — bfd_enabled / 021111303122 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- local_control_plane.bgp_config.peers.bfd_enabled

<a id="canonical-0030032211222001-2211031220203302-3233020300312220-2110333331130010-1232322331010231-2122000203111210-2113113302230322-0320202302031203"></a>

Type: `"single"`. Computed.

BFD. BFD parameters.

Upstream description:

BFD parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1321031223102330-1302232012123133-2202202112001020-0212033222230002-1100330132122232-2301132223301322-3213303003123001-0032022313231032"></a>

## Direct properties — bfd_enabled / 021111303122 / 3

<a id="canonical-2010200232332221-2013003030011202-1333033212322313-2120000230033022-3230002330333220-1312223113233331-2003001111030312-0101020301110320"></a>

<a id="canonical-2331003033032310-3303112122010133-2212310311323033-0201321310202330-1020131221000320-0121300113230233-2230132223231202-2131122220300011"></a>

## multiplier property — bfd_enabled / 021111303122 / 4

Type: `"number"`. Computed.

Specify Number of missed packets to bring session down'.

Upstream description:

Specify Number of missed packets to bring session down"

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
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-1230121103013001-2021102223031012-3311320232122200-3311300120030011-3131213101200131-3212220111031023-3333111313123111-2021110122301311"></a>

<a id="canonical-3231303201123223-2020022202000212-0130123201333022-0211320012013120-2201222121101231-1210211033132110-0223112332203211-1113310121310111"></a>

## receive_interval_milliseconds property — bfd_enabled / 021111303122 / 5

Type: `"number"`. Computed.

BFD receive interval timer, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-0322331323000200-2330022122132031-1032002013130100-1311311211123111-1301102101302011-0123100203001120-2200130010212011-2223203312200132"></a>

<a id="canonical-2122221232112131-1331233003233332-2030313232013032-0232223330320320-0113023330312011-0201202102110133-1322203201213101-2302101220320110"></a>

## transmit_interval_milliseconds property — bfd_enabled / 021111303122 / 6

Type: `"number"`. Computed.

BFD transmit interval timer, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-3213302211033033-3002230232202100-3200321100201213-2202020230103300-3131013201000032-3112211312310212-1110131232031232-2321023213103020"></a>

## Next pages — bfd_enabled / 021111303122 / 7

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1121132012200303-3002212112032202-2031330221132330-2103011113200033-0000303110023300-3302330230230031-2231103010221033-1222022223210310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123321020132313-2331331301132300-3221003212133122-1322200233213221-0122023320000013-2120023033030300-0323013011333111-2203312130130323"></a>

## local_control_plane.bgp_config.peers.disable_spec — disable_spec / 102232120222 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- local_control_plane.bgp_config.peers.disable_spec

<a id="canonical-3030033120020130-3332013300130011-2330012030033303-0010101331101121-0333322103133230-3033033222023200-2221133103023003-0131101201103031"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-1223101111223001-2213031030003311-0223032133020023-0221203220203323-0003213022121130-1033000320211212-3020323303311130-0002113320010312"></a>

## Direct properties — disable_spec / 102232120222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132212232321321-2011000202220103-2122111202011221-3332011002032231-0113003103333331-1233203033222330-2233001130200202-2200002031202122"></a>

## Next pages — disable_spec / 102232120222 / 4

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0302123211022230-0030120332220022-1110120120202230-0223010231032123-0130302103213210-1301003223221221-1031311021200230-3100322212030223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000221113122112-3003311102113021-3022232211133001-1001303121033330-3032022133032333-3010231103311231-1010202202131232-3330011122103221"></a>

## local_control_plane.bgp_config.peers.ebgp_multihop_disabled — ebgp_multihop_disabled / 132232322013 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- local_control_plane.bgp_config.peers.ebgp_multihop_disabled

<a id="canonical-0213023202101121-3211003202123031-1232233311333030-2021010310011033-0211311231200000-3312102230323222-3033231030131111-1133230010131321"></a>

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

<a id="canonical-0200301112321012-0232011102111203-1331201223202020-2123113322111020-2222022100133302-3313203010312200-2231313000223321-0121032200320030"></a>

## Direct properties — ebgp_multihop_disabled / 132232322013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120112112113033-2020013313303220-3323032101331023-2213011110013320-0122331020222022-0232130130301333-2233033300122001-0120333203130012"></a>

## Next pages — ebgp_multihop_disabled / 132232322013 / 4

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2203031313131132-3330232311210023-1032321110003101-1201313233130132-1133102110111030-0113323002311332-0102032321111113-0002312011032310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002111222212012-3030233311202222-1103111112320011-1123330001013101-0221211303101201-1230133212103123-0332302213123112-1311213101100033"></a>

## local_control_plane.bgp_config.peers.ebgp_multihop_enabled — ebgp_multihop_enabled / 001203213131 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- local_control_plane.bgp_config.peers.ebgp_multihop_enabled

<a id="canonical-3323303012002100-1323321001033031-0231130130302233-2002301100220310-2310102001210211-1200031030122201-3302031133013132-1012332033102303"></a>

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

<a id="canonical-3232112320321123-2200233212102313-2130112311033232-3122102200332123-0122103330103312-3032331211130103-1231200101100130-3220032000233130"></a>

## Direct properties — ebgp_multihop_enabled / 001203213131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1102221000232011-0333023001301332-0232020312100122-1330000300221302-0302213033023300-2223201020323100-3221121330302023-2112233130222203"></a>

## Next pages — ebgp_multihop_enabled / 001203213131 / 4

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010331212031210-0320220310133113-2213220232213132-0102123202132231-1233030203200133-0202120233022331-2013320232300032-2133222001000231"></a>

## local_control_plane.bgp_config.peers.external — external / 032121222121 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- local_control_plane.bgp_config.peers.external

<a id="canonical-2033332022322311-1222233200033231-3323113100020202-0101100111222030-2323310111212333-1003303331321110-0221000300002213-3103303132001010"></a>

Type: `"single"`. Computed.

External BGP Peer. External BGP Peer parameters.

Upstream description:

External BGP Peer parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_choice": "[\"address\",\"default_gateway\",\"disable\",\"external_connector\",\"from_site\",\"subnet_begin_offset\",\"subnet_end_offset\"]",
  "x-ves-oneof-field-address_choice_v6": "[\"address_ipv6\",\"default_gateway_v6\",\"disable_v6\",\"from_site_v6\",\"subnet_begin_offset_v6\",\"subnet_end_offset_v6\"]",
  "x-ves-oneof-field-auth_choice": "[\"md5_auth_key\",\"no_authentication\"]",
  "x-ves-oneof-field-interface_choice": "[\"interface\",\"interface_list\"]"
}
```

<a id="canonical-1122100303132330-0030101322033221-1113110211132213-0211211313000233-0020301223013330-3201001213111032-0130212303032002-0122211221123110"></a>

## Direct properties — external / 032121222121 / 3

<a id="canonical-3130113232120231-2221311322331222-1313032001220032-0033033331130133-0032010220131013-1030010030100030-2222112111330231-1111221100100321"></a>

<a id="canonical-3203323020302332-2323012303333211-1103210011322213-1320302031300022-0332101331031031-0300320213121122-2201201113121330-2231311132133323"></a>

## address property — external / 032121222121 / 4

Type: `"string"`. Computed.

Exclusive with \[default\_gateway disable external\_connector from\_site subnet\_begin\_offset
subnet\_end\_offset\] Specify IPv4 peer address.

Upstream description:

Exclusive with \[default\_gateway disable external\_connector from\_site subnet\_begin\_offset
subnet\_end\_offset\] Specify IPv4 peer address.

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

<a id="canonical-1213233133103122-2311202311320320-3311310133100002-2331230112120313-2323311120303331-0131303023303212-2130033123020233-2230022013100020"></a>

<a id="canonical-1303003011232221-3011130321101032-0331313320322200-1131030330121212-1201221203212112-2012123310201312-3120102100020310-0010210020121000"></a>

## address_ipv6 property — external / 032121222121 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway\_v6 disable\_v6 from\_site\_v6 subnet\_begin\_offset\_v6
subnet\_end\_offset\_v6\] Specify peer IPv6 address.

Upstream description:

Exclusive with \[default\_gateway\_v6 disable\_v6 from\_site\_v6 subnet\_begin\_offset\_v6
subnet\_end\_offset\_v6\] Specify peer IPv6 address.

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-1320222200203122-1020010210311102-3002010203310301-0231112213010023-2100213213111313-1032003132202020-1333310120102211-2310212311101323"></a>

<a id="canonical-0103221131031132-0212030021303022-0100232211003121-0102012133223231-1032323221203303-3230012112223330-0113031122310120-0122200320031300"></a>

## asn property — external / 032121222121 / 6

Type: `"number"`. Computed.

ASN. Autonomous System Number for BGP peer.

Upstream description:

Autonomous System Number for BGP peer.

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
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [default_gateway](data-sources--voltstack_site--reference--group-009.md#canonical-2323230131030103-1213321231011310-3121001013203111-2102121211102210-0010012003220311-0321033330021231-0030122012201330-3322312332222000): complete subsection reference.

- [default_gateway_v6](data-sources--voltstack_site--reference--group-009.md#canonical-3223111011121331-1112013003320231-0300303132031220-3001231102212202-3113031230232100-1330122300112001-0012111101100131-3111113213002120): complete subsection reference.

- [disable_spec](data-sources--voltstack_site--reference--group-009.md#canonical-1211210210203322-1002100303003102-1133002021312023-2301232010233020-2131310302322130-0333323012330121-1222312302110210-0000012302210230): complete subsection reference.

- [disable_v6](data-sources--voltstack_site--reference--group-009.md#canonical-2222110030303223-3323030323113022-2213223311213013-1013111121122232-1032103232210211-2100330100121233-3111011223021232-1110101002011013): complete subsection reference.

- [external_connector](data-sources--voltstack_site--reference--group-009.md#canonical-0023220103203200-0030011020231131-1311003223230303-0022020333030013-1101013221211021-3333010102230122-2000001101112021-2000301313010222): complete subsection reference.

- [family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-3221003333002022-3202311222300331-3333013012011232-3310323022112233-1131131123010033-3010210203333033-1010102022312212-0012320221130233): complete subsection reference.

- [from_site](data-sources--voltstack_site--reference--group-009.md#canonical-1201201010122201-2022231201323000-2032221200102133-0220122012320231-1032202200330003-1132300033213023-1332122232301011-2203031330330330): complete subsection reference.

- [from_site_v6](data-sources--voltstack_site--reference--group-009.md#canonical-2113323132220010-1220300310331131-3323013211010133-2121220033003321-2020201323230211-1130011103323002-0200220031312033-2323300101323123): complete subsection reference.

- [interface](data-sources--voltstack_site--reference--group-009.md#canonical-0332101120132301-2303030130200020-0022330323103120-2323303031021112-0233020322013000-0113301322333310-0333023111223012-2313002222223110): complete subsection reference.

- [interface_list](data-sources--voltstack_site--reference--group-009.md#canonical-0101123232112131-3200301120023312-1212313300313020-0113321213022331-1100111312221132-0000011020002230-2330231001003001-2032122210112203): complete subsection reference.

<a id="canonical-3210220022022311-2332212120222201-2133113013212133-3120003101231020-0010203231103101-3322101002113303-0002023001022210-0101001322111313"></a>

<a id="canonical-1312120023213121-0320202103233332-3211212221300100-3120330222323221-2303001112003131-1131300202201131-0030300323101131-1311123302220003"></a>

## md5_auth_key property — external / 032121222121 / 7

Type: `"string"`. Computed.

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385).

Upstream description:

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385)

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

- [no_authentication](data-sources--voltstack_site--reference--group-009.md#canonical-2320022202022210-0332101100020023-0020122201230130-1113332221311231-2123020002313213-2201313311132333-3001130123320220-0320002003103000): complete subsection reference.

<a id="canonical-2211113310121300-3210012013321013-2120120321330322-0023102200201201-2230101033320231-0022032232133003-0032033320201113-1100102212312030"></a>

<a id="canonical-1001201000311130-2202313333121223-1121230300130013-1020202203233021-2222320003330112-0313203200212313-2332130021330202-3122332200001211"></a>

## port property — external / 032121222121 / 8

Type: `"number"`. Computed.

Peer Port. Peer TCP port number.

Upstream description:

Peer TCP port number.

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
    "create": true,
    "minimum_config": true,
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

<a id="canonical-3020122103013033-3322300121001320-2001200112101123-0320232233323111-1320010010232011-2001001323022103-2113111122230222-0201033312002022"></a>

<a id="canonical-2132022111223002-3133203103213300-1021301300221310-1122200201323212-3011002210001003-3333132333032022-1103103010303011-1322120232032121"></a>

## subnet_begin_offset property — external / 032121222121 / 9

Type: `"number"`. Computed.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_end\_offset\] Calculate peer address using offset from the beginning of the subnet.

Upstream description:

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_end\_offset\] Calculate peer address using offset from the beginning of the subnet.

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
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-1221213301202031-0222020003123122-1133311312102030-0012233033000110-3020232110333201-3211111121001032-0000003011321130-1012021133121103"></a>

<a id="canonical-1102030311222333-3120131322211202-2333230320323110-3230312231203313-1300313131101220-1321221113021032-2123111311110322-2320122313023130"></a>

## subnet_begin_offset_v6 property — external / 032121222121 / 10

Type: `"number"`. Computed.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_end\_offset\_v6\] Calculate peer address using offset from the beginning of the subnet.

Upstream description:

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_end\_offset\_v6\] Calculate peer address using offset from the beginning of the subnet.

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
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-0100100323313221-0111103102312012-1330220213012103-1232130330322122-1010122013012213-0122211003112322-2010003101002303-0000122331300023"></a>

<a id="canonical-0010232201330331-3323203332121112-1123113123333110-0322313321310031-0302323112131121-2311211313112232-3002111202033010-3221012131220230"></a>

## subnet_end_offset property — external / 032121222121 / 11

Type: `"number"`. Computed.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_begin\_offset\] Calculate peer address using offset from the end of the subnet.

Upstream description:

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_begin\_offset\] Calculate peer address using offset from the end of the subnet.

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
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-1020323122211032-1213111133231021-2131011323210310-2122321021301233-0113213321020021-3012131203002013-1132221012020113-1212231130313010"></a>

<a id="canonical-0001120033310130-1022131121112030-3200331322103210-2121113333321113-1033233331321030-2022122331122312-2310011222013233-1311032201110331"></a>

## subnet_end_offset_v6 property — external / 032121222121 / 12

Type: `"number"`. Computed.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_begin\_offset\_v6\] Calculate peer address using offset from the end of the subnet.

Upstream description:

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_begin\_offset\_v6\] Calculate peer address using offset from the end of the subnet.

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
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-3010030100022201-0133131121231012-3013001103020011-1300220003101031-0113213322022112-2300233220130012-3013002210200322-0021321300232033"></a>

## Next pages — external / 032121222121 / 13

- [local_control_plane.bgp_config.peers.external.default_gateway](data-sources--voltstack_site--reference--group-009.md#canonical-2323230131030103-1213321231011310-3121001013203111-2102121211102210-0010012003220311-0321033330021231-0030122012201330-3322312332222000)
- [local_control_plane.bgp_config.peers.external.default_gateway_v6](data-sources--voltstack_site--reference--group-009.md#canonical-3223111011121331-1112013003320231-0300303132031220-3001231102212202-3113031230232100-1330122300112001-0012111101100131-3111113213002120)
- [local_control_plane.bgp_config.peers.external.disable_spec](data-sources--voltstack_site--reference--group-009.md#canonical-1211210210203322-1002100303003102-1133002021312023-2301232010233020-2131310302322130-0333323012330121-1222312302110210-0000012302210230)
- [local_control_plane.bgp_config.peers.external.disable_v6](data-sources--voltstack_site--reference--group-009.md#canonical-2222110030303223-3323030323113022-2213223311213013-1013111121122232-1032103232210211-2100330100121233-3111011223021232-1110101002011013)
- [local_control_plane.bgp_config.peers.external.external_connector](data-sources--voltstack_site--reference--group-009.md#canonical-0023220103203200-0030011020231131-1311003223230303-0022020333030013-1101013221211021-3333010102230122-2000001101112021-2000301313010222)
- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-3221003333002022-3202311222300331-3333013012011232-3310323022112233-1131131123010033-3010210203333033-1010102022312212-0012320221130233)
- [local_control_plane.bgp_config.peers.external.from_site](data-sources--voltstack_site--reference--group-009.md#canonical-1201201010122201-2022231201323000-2032221200102133-0220122012320231-1032202200330003-1132300033213023-1332122232301011-2203031330330330)
- [local_control_plane.bgp_config.peers.external.from_site_v6](data-sources--voltstack_site--reference--group-009.md#canonical-2113323132220010-1220300310331131-3323013211010133-2121220033003321-2020201323230211-1130011103323002-0200220031312033-2323300101323123)
- [local_control_plane.bgp_config.peers.external.interface](data-sources--voltstack_site--reference--group-009.md#canonical-0332101120132301-2303030130200020-0022330323103120-2323303031021112-0233020322013000-0113301322333310-0333023111223012-2313002222223110)
- [local_control_plane.bgp_config.peers.external.interface_list](data-sources--voltstack_site--reference--group-009.md#canonical-0101123232112131-3200301120023312-1212313300313020-0113321213022331-1100111312221132-0000011020002230-2330231001003001-2032122210112203)
- [local_control_plane.bgp_config.peers.external.no_authentication](data-sources--voltstack_site--reference--group-009.md#canonical-2320022202022210-0332101100020023-0020122201230130-1113332221311231-2123020002313213-2201313311132333-3001130123320220-0320002003103000)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2323230131030103-1213321231011310-3121001013203111-2102121211102210-0010012003220311-0321033330021231-0030122012201330-3322312332222000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223021230233230-1100210000003130-0003000131212303-2102203311210022-1333220103120012-2311201210000031-0033310013031310-2202323011113230"></a>

## local_control_plane.bgp_config.peers.external.default_gateway — default_gateway / 313211213010 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- local_control_plane.bgp_config.peers.external.default_gateway

<a id="canonical-2021201110222121-2222210022121222-0302200010220233-1121303222122031-2102130132100201-0313120100113232-2021001001213103-3020013230200111"></a>

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

<a id="canonical-3132001003130200-0322101102300320-3331230301231300-2011310212102232-3213030131331020-0311333210132013-1110120201232120-1201330030002133"></a>

## Direct properties — default_gateway / 313211213010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1230022233312303-3120211133100301-1033022313001123-0110213300311132-1112201220123130-1001010023210030-0211301102110130-3123100232211130"></a>

## Next pages — default_gateway / 313211213010 / 4

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3223111011121331-1112013003320231-0300303132031220-3001231102212202-3113031230232100-1330122300112001-0012111101100131-3111113213002120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102212101313131-2111131230301113-0323223010322221-3021210032113122-1230012002330021-2101301100332130-2203320201312110-3303010321330201"></a>

## local_control_plane.bgp_config.peers.external.default_gateway_v6 — default_gateway_v6 / 031110320330 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- local_control_plane.bgp_config.peers.external.default_gateway_v6

<a id="canonical-3112221211112201-2102303313313021-3331220122011231-2322122313121000-2110110202011110-3010212321300012-2203131201200012-0331332033100013"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway v6.

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

<a id="canonical-3202313331002022-3232311222000322-3002220213100321-1320012022013301-2211110133022221-0223220030323303-1211000032303120-3220122012002113"></a>

## Direct properties — default_gateway_v6 / 031110320330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110331332111010-2101213230123322-1322012030010101-2213020212022010-1022310213301101-1112200310210100-2012223121030220-3111002102110112"></a>

## Next pages — default_gateway_v6 / 031110320330 / 4

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1211210210203322-1002100303003102-1133002021312023-2301232010233020-2131310302322130-0333323012330121-1222312302110210-0000012302210230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003232231230333-2031031310001102-1202103131232302-1130233002013301-0330001221012330-1100302022001322-0023113313220122-0023311103230130"></a>

## local_control_plane.bgp_config.peers.external.disable_spec — disable_spec / 202312132101 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- local_control_plane.bgp_config.peers.external.disable_spec

<a id="canonical-2020220320023111-0320022023020113-2230020113013011-3122323110032030-3231302031031121-3132121233100211-2120330020331002-3131320220110312"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-0312312310110011-2022000010001211-2232012200330332-2210233313322322-2021032000311322-1322002330300010-1101202200033100-2131122312310123"></a>

## Direct properties — disable_spec / 202312132101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213310233310113-1120020120310130-2321321210201132-2303002300002333-3231211130021023-3022220133223312-1223113331231321-1331110320233013"></a>

## Next pages — disable_spec / 202312132101 / 4

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2222110030303223-3323030323113022-2213223311213013-1013111121122232-1032103232210211-2100330100121233-3111011223021232-1110101002011013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032322123023113-3302123212013103-2100023023330030-1330002031230322-3001301201323232-0212313312200203-2111113221100113-1130000100001101"></a>

## local_control_plane.bgp_config.peers.external.disable_v6 — disable_v6 / 111222022121 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- local_control_plane.bgp_config.peers.external.disable_v6

<a id="canonical-1113212100102303-1033010220201300-1033321003303022-1212232013113120-2300301223201302-2322122022230033-2102220011111222-1130123013112121"></a>

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

<a id="canonical-3113301033313120-0231212003310033-1113211303301033-1221121230230031-0203110302100032-3333303211301023-1132033203030232-2323101103230211"></a>

## Direct properties — disable_v6 / 111222022121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200130121132230-0233313013112132-2010130020221003-1111321231302003-0033100313200233-1010130312301322-0130211311322222-2013322012002202"></a>

## Next pages — disable_v6 / 111222022121 / 4

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0023220103203200-0030011020231131-1311003223230303-0022020333030013-1101013221211021-3333010102230122-2000001101112021-2000301313010222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330011221000003-1200021030220201-3100101200100301-1213323310303221-3111333301123001-3120020300333222-1330010330230102-2130011131101120"></a>

## local_control_plane.bgp_config.peers.external.external_connector — external_connector / 213313321022 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- local_control_plane.bgp_config.peers.external.external_connector

<a id="canonical-3001211031201231-2002012312102223-3121331101121323-1322003233102002-0203003232321131-2301030032212212-2010021122130003-1002133030111100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for external connector.

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

<a id="canonical-2121332202121200-0122110021231121-0110001022333200-3020331020110301-3230220222330311-2031012013120323-2333201210223222-1020023122113010"></a>

## Direct properties — external_connector / 213313321022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111301130033303-0211331011101031-0031122101321322-2130013331302030-2110101130133312-3102110012223003-3302301222122203-3033203210012023"></a>

## Next pages — external_connector / 213313321022 / 4

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3221003333002022-3202311222300331-3333013012011232-3310323022112233-1131131123010033-3010210203333033-1010102022312212-0012320221130233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112031111000223-2111212031221102-2311123133023201-1232221023330311-0130231121333223-3010123030212333-2001230002131313-2011030211021103"></a>

## local_control_plane.bgp_config.peers.external.family_inet — family_inet / 213121121212 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- local_control_plane.bgp_config.peers.external.family_inet

<a id="canonical-2321033332012101-3232312211202112-0103312032323003-3100332110003230-0101002020233033-0011122301332322-3122321231333312-1133303000121332"></a>

Type: `"single"`. Computed.

Configuration parameter for family inet.

Upstream description:

Parameters for inet family.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-enable_choice": "[\"disable\",\"enable\"]"
}
```

<a id="canonical-0012131113313000-0333023220231301-0302000021312310-1102311102220302-2013121332130011-1323333013003133-3203003331212103-0112121231310113"></a>

## Direct properties — family_inet / 213121121212 / 3

- [disable_spec](data-sources--voltstack_site--reference--group-009.md#canonical-1101033203020131-2210301332201000-3201132223213312-1233211100112033-2302313321120001-2323111130330311-0131311013311313-1112331121122213): complete subsection reference.

- [enable](data-sources--voltstack_site--reference--group-009.md#canonical-2302232033220033-1320110312123221-0323113130000231-2203011201231320-2233113333222222-0221131013300330-1322031110001200-0100133300030032): complete subsection reference.

<a id="canonical-2221212010313120-1013312020203122-1213320122222323-0002121100031213-0032023023010313-2002310112211013-1033133222323332-0323120000322202"></a>

## Next pages — family_inet / 213121121212 / 4

- [local_control_plane.bgp_config.peers.external.family_inet.disable_spec](data-sources--voltstack_site--reference--group-009.md#canonical-1101033203020131-2210301332201000-3201132223213312-1233211100112033-2302313321120001-2323111130330311-0131311013311313-1112331121122213)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](data-sources--voltstack_site--reference--group-009.md#canonical-2302232033220033-1320110312123221-0323113130000231-2203011201231320-2233113333222222-0221131013300330-1322031110001200-0100133300030032)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1101033203020131-2210301332201000-3201132223213312-1233211100112033-2302313321120001-2323111130330311-0131311013311313-1112331121122213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313311120333003-3211103303322230-0213321222222112-1011320111003301-0031313120110001-0231030332103222-2323333012032023-1300122222221331"></a>

## local_control_plane.bgp_config.peers.external.family_inet.disable_spec — disable_spec / 121313321132 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-3221003333002022-3202311222300331-3333013012011232-3310323022112233-1131131123010033-3010210203333033-1010102022312212-0012320221130233)
- local_control_plane.bgp_config.peers.external.family_inet.disable_spec

<a id="canonical-1300303200321001-0120332230313200-1000000032100212-0323101312322131-2202023033300202-3313012312111012-2120023320122032-1300002020300222"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-1211113030032203-3231213101111310-1330030330112100-0333111301213300-2110220001113322-2103330103232310-2123023210223310-1021313210031130"></a>

## Direct properties — disable_spec / 121313321132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3012212233011022-2003210213232133-1201231123002012-1100331122223031-1010023222010332-0222103210132230-0310322301012330-1220000221120133"></a>

## Next pages — disable_spec / 121313321132 / 4

- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-3221003333002022-3202311222300331-3333013012011232-3310323022112233-1131131123010033-3010210203333033-1010102022312212-0012320221130233)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2302232033220033-1320110312123221-0323113130000231-2203011201231320-2233113333222222-0221131013300330-1322031110001200-0100133300030032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221102220320111-0102202221301120-1112322210321021-1323011011333111-3100233320212022-3301032320323231-1131211023201102-0033102231030331"></a>

## local_control_plane.bgp_config.peers.external.family_inet.enable — enable / 011031010033 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-3221003333002022-3202311222300331-3333013012011232-3310323022112233-1131131123010033-3010210203333033-1010102022312212-0012320221130233)
- local_control_plane.bgp_config.peers.external.family_inet.enable

<a id="canonical-2112011113113001-0113320020302312-1333321232311233-3203022300311222-0301311301301202-3013231311102101-3120000130300102-1030310223102030"></a>

Type: `"single"`. Computed.

Unicast IPv4. IPv4 Unicast.

Upstream description:

IPv4 Unicast.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3232022202211111-1100133210310103-2220100021003113-1301220010201103-1333221203123101-1323013131201012-1010133301100032-2333332301013112"></a>

## Direct properties — enable / 011031010033 / 3

- [aggregation](data-sources--voltstack_site--reference--group-009.md#canonical-0231010302233033-0033032011233330-1201223333032310-3232013232130210-1233330333333232-2122132020130320-2202203121033001-3033102311013122): complete subsection reference.

<a id="canonical-1031121013031210-3313223332230033-3120210002030223-3130123133210211-1313202031100211-1232213103020301-2221221121202023-2320010002133302"></a>

## Next pages — enable / 011031010033 / 4

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](data-sources--voltstack_site--reference--group-009.md#canonical-0231010302233033-0033032011233330-1201223333032310-3232013232130210-1233330333333232-2122132020130320-2202203121033001-3033102311013122)
- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-3221003333002022-3202311222300331-3333013012011232-3310323022112233-1131131123010033-3010210203333033-1010102022312212-0012320221130233)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0231010302233033-0033032011233330-1201223333032310-3232013232130210-1233330333333232-2122132020130320-2202203121033001-3033102311013122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112032002033203-2211212130110212-3103031012103311-1121003200313012-2301211333120011-3020113023302210-1120230211030013-2013011212133332"></a>

## local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation — aggregation / 331232300132 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-3221003333002022-3202311222300331-3333013012011232-3310323022112233-1131131123010033-3010210203333033-1010102022312212-0012320221130233)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](data-sources--voltstack_site--reference--group-009.md#canonical-2302232033220033-1320110312123221-0323113130000231-2203011201231320-2233113333222222-0221131013300330-1322031110001200-0100133300030032)
- local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation

<a id="canonical-3023222210330112-1010332013212331-2101000012332010-1301332230302013-2313122310132012-2002010003220322-1023212002020133-2011011112221033"></a>

Type: `"list"`. Computed.

BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take
effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing
table and applies to outbound advertisements.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3333302131131121-2033101122010001-2332010313020220-1020022013100302-2302310133213232-3131030201022331-1103220120020332-1213002231231301"></a>

## Direct properties — aggregation / 331232300132 / 3

<a id="canonical-1111211123230130-3102313013001122-2113312330011302-1110122233132011-1321322031002332-2111303301320033-2132020221003022-3020031322111030"></a>

<a id="canonical-2202122031321312-1312022123201300-2332233302223300-2110331120313023-1122323311013222-1113203323201211-1331311002211323-0013202321011313"></a>

## ip_prefix property — aggregation / 331232300132 / 4

Type: `"string"`. Computed.

IP Prefix. Specify IPv4 subnet for aggregation.

Upstream description:

Specify IPv4 subnet for aggregation.

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

- [options](data-sources--voltstack_site--reference--group-009.md#canonical-0322213212021212-1103311301130001-2021312223201203-0320303023030302-0220002031310033-0100213230310203-3330322211023202-2001203300220333): complete subsection reference.

<a id="canonical-0033012031111221-3132330233123210-2320310023213322-3003231211113123-0312023133322200-2023121311322032-0332120222222320-3111311201021111"></a>

## Next pages — aggregation / 331232300132 / 5

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options](data-sources--voltstack_site--reference--group-009.md#canonical-0322213212021212-1103311301130001-2021312223201203-0320303023030302-0220002031310033-0100213230310203-3330322211023202-2001203300220333)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](data-sources--voltstack_site--reference--group-009.md#canonical-2302232033220033-1320110312123221-0323113130000231-2203011201231320-2233113333222222-0221131013300330-1322031110001200-0100133300030032)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0322213212021212-1103311301130001-2021312223201203-0320303023030302-0220002031310033-0100213230310203-3330322211023202-2001203300220333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031032033311213-2031123312103033-2013322031010331-3123133000011130-2202112230331133-3210220011023312-0030323301213033-2331100310200102"></a>

## local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options — options / 031321121020 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-3221003333002022-3202311222300331-3333013012011232-3310323022112233-1131131123010033-3010210203333033-1010102022312212-0012320221130233)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](data-sources--voltstack_site--reference--group-009.md#canonical-2302232033220033-1320110312123221-0323113130000231-2203011201231320-2233113333222222-0221131013300330-1322031110001200-0100133300030032)
- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](data-sources--voltstack_site--reference--group-009.md#canonical-0231010302233033-0033032011233330-1201223333032310-3232013232130210-1233330333333232-2122132020130320-2202203121033001-3033102311013122)
- local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options

<a id="canonical-1121003013011112-0321302103103232-2102223310331302-0102102123231133-0303312100123011-1211313031033313-2320300223200330-2001203012313332"></a>

Type: `"list"`. Computed.

Aggregation OPTIONS. Configuration parameter for options

Upstream description:

Configuration parameter for options

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0131212210201012-2333021223222102-1322123321320021-1000312010120231-2313021232133331-0112230312123101-2200122010321211-2000233032013322"></a>

## Direct properties — options / 031321121020 / 3

- [summary_only](data-sources--voltstack_site--reference--group-009.md#canonical-1000313032103201-0012313023121122-1123121333133110-2311121111131012-2131212012010223-2031333233031022-2001121221231120-3121213033332332): complete subsection reference.

<a id="canonical-3333203202211023-1233313021223131-0331113322212001-2332201222233133-1200130313130202-3200033320301033-0230110030232323-1210332021010122"></a>

## Next pages — options / 031321121020 / 4

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options.summary_only](data-sources--voltstack_site--reference--group-009.md#canonical-1000313032103201-0012313023121122-1123121333133110-2311121111131012-2131212012010223-2031333233031022-2001121221231120-3121213033332332)
- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](data-sources--voltstack_site--reference--group-009.md#canonical-0231010302233033-0033032011233330-1201223333032310-3232013232130210-1233330333333232-2122132020130320-2202203121033001-3033102311013122)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1000313032103201-0012313023121122-1123121333133110-2311121111131012-2131212012010223-2031333233031022-2001121221231120-3121213033332332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320210012121330-2013222030212322-1302132201312110-3321302322000320-2011221013102133-3010112232220202-0102213312021320-1002012313123023"></a>

## local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options.summary_only — summary_only / 203112201010 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [local_control_plane.bgp_config.peers.external.family_inet](data-sources--voltstack_site--reference--group-009.md#canonical-3221003333002022-3202311222300331-3333013012011232-3310323022112233-1131131123010033-3010210203333033-1010102022312212-0012320221130233)
- [local_control_plane.bgp_config.peers.external.family_inet.enable](data-sources--voltstack_site--reference--group-009.md#canonical-2302232033220033-1320110312123221-0323113130000231-2203011201231320-2233113333222222-0221131013300330-1322031110001200-0100133300030032)
- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation](data-sources--voltstack_site--reference--group-009.md#canonical-0231010302233033-0033032011233330-1201223333032310-3232013232130210-1233330333333232-2122132020130320-2202203121033001-3033102311013122)
- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options](data-sources--voltstack_site--reference--group-009.md#canonical-0322213212021212-1103311301130001-2021312223201203-0320303023030302-0220002031310033-0100213230310203-3330322211023202-2001203300220333)
- local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options.summary_only

<a id="canonical-0033001121323122-3212100210022223-0320211002011002-3101332022020013-0221222303012230-0102302112033020-3302312120302303-1332113311300000"></a>

Type: `"single"`. Computed.

Configuration parameter for summary only.

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

<a id="canonical-2320010022323012-2132120211222022-0213221202030011-0332130032312012-2321311031023210-1030110100001003-1311231001320023-2013211032023301"></a>

## Direct properties — summary_only / 203112201010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033231022001230-0333202203133122-0222333212300301-2103122230100011-2323330113221302-2031123211102023-0121301201220331-2230300133200331"></a>

## Next pages — summary_only / 203112201010 / 4

- [local_control_plane.bgp_config.peers.external.family_inet.enable.aggregation.options](data-sources--voltstack_site--reference--group-009.md#canonical-0322213212021212-1103311301130001-2021312223201203-0320303023030302-0220002031310033-0100213230310203-3330322211023202-2001203300220333)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1201201010122201-2022231201323000-2032221200102133-0220122012320231-1032202200330003-1132300033213023-1332122232301011-2203031330330330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132210200212331-0031002201313212-1222222231022220-0312231032030102-3112033223130032-3230203023320120-3033331311313311-1301231021103302"></a>

## local_control_plane.bgp_config.peers.external.from_site — from_site / 111222033211 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- local_control_plane.bgp_config.peers.external.from_site

<a id="canonical-0332113231231331-0321121210131313-0231012131220312-3113030231323101-0110311211101030-2020302011230123-3232212311020110-0320320132222331"></a>

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

<a id="canonical-1220312202011020-2200110131312321-0202130330130012-2033131300013302-3221310221033221-3303331331121303-3232210303013332-2032220020200033"></a>

## Direct properties — from_site / 111222033211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113221021112303-1121310033203001-3332222221221301-2002323330113313-3333200003322311-1100313303222300-1333220022001011-1312100132220000"></a>

## Next pages — from_site / 111222033211 / 4

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2113323132220010-1220300310331131-3323013211010133-2121220033003321-2020201323230211-1130011103323002-0200220031312033-2323300101323123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133201033100113-1101233003102212-2223130122011302-2230032223111103-2303131301130120-2313130202220121-2000313132223030-1212000321332100"></a>

## local_control_plane.bgp_config.peers.external.from_site_v6 — from_site_v6 / 300303211302 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- local_control_plane.bgp_config.peers.external.from_site_v6

<a id="canonical-0111111200013100-2010121233102210-2212011032100010-3312301331333102-0123022322201233-1231223022032332-1313030221223323-0323211230100103"></a>

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

<a id="canonical-2023012013210012-1133000133232302-2312301213122011-0020313113132220-1032130113022112-3302013222110310-0012133233331120-3011222321213033"></a>

## Direct properties — from_site_v6 / 300303211302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200133130131110-3323013313013020-0330121220301203-3002030102321330-0033220300033122-3121321302101322-0213321311233333-2330300003013211"></a>

## Next pages — from_site_v6 / 300303211302 / 4

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0332101120132301-2303030130200020-0022330323103120-2323303031021112-0233020322013000-0113301322333310-0333023111223012-2313002222223110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033200301323231-1013031022303100-1013303122310322-3001033220320120-0300201121112103-0022301122022003-1212222113200321-2223031013220020"></a>

## local_control_plane.bgp_config.peers.external.interface — interface / 330220113211 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- local_control_plane.bgp_config.peers.external.interface

<a id="canonical-3310233333013313-3131111313301133-2201310231033031-1231320302103201-3301232203213313-1313301330011212-2011131230301120-2320120212131333"></a>

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

<a id="canonical-0010010310022023-0133000001002030-0230022200221023-2130320101033202-3313000021231021-0203212121333312-1011023211202020-1210102022331130"></a>

## Direct properties — interface / 330220113211 / 3

<a id="canonical-0303330033220030-1103011301233103-1102102311022201-3221331321020001-2221300001200210-2302100313333331-1020033120222122-2322113303110202"></a>

<a id="canonical-2210023101311203-0133213100110013-0013220233201100-0232223320331023-3213121011211200-3200133131101101-1000202003010200-1130200012002101"></a>

## name property — interface / 330220113211 / 4

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

<a id="canonical-0113332012313211-3022002032002033-2201223111231330-2331210001203221-3110332121020002-1230321030322110-1102100011010110-2021201223330020"></a>

<a id="canonical-2211010221102121-0011213103300322-3030012121233013-3321102223333231-2331230133231012-1123122002030011-1223310331022112-1223133033202223"></a>

## namespace property — interface / 330220113211 / 5

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

<a id="canonical-0301121100201223-1111123331233300-3332031231201122-3123320112320021-1101230302212123-2033210203102133-3110021132000212-0113213303302333"></a>

<a id="canonical-2321201013233233-2123100100122320-0010210110321202-1020200211321122-2110123030031330-3332232101010121-0323323310312232-2111113221032302"></a>

## tenant property — interface / 330220113211 / 6

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

<a id="canonical-0312332121003213-1011020110133002-3110330031001203-1113002112232323-1211223123220021-0002213202212101-3121211200313021-3222132120132202"></a>

## Next pages — interface / 330220113211 / 7

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0101123232112131-3200301120023312-1212313300313020-0113321213022331-1100111312221132-0000011020002230-2330231001003001-2032122210112203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323123202332201-2203003212322321-3232223023300103-0332120121221020-1000113020132023-0231220233103102-1122330223011121-0113101200102113"></a>

## local_control_plane.bgp_config.peers.external.interface_list — interface_list / 011022030031 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- local_control_plane.bgp_config.peers.external.interface_list

<a id="canonical-1023020000022033-2220111121313301-0020032232000210-2032221223003020-3333201120221110-0000012120000101-1333033111232102-3102332232222310"></a>

Type: `"single"`. Computed.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1020100300321113-3010323122031122-1103032223323222-1300010000201130-1121230332120102-0001121212210130-3033301102101122-2132301330221002"></a>

## Direct properties — interface_list / 011022030031 / 3

- [interfaces](data-sources--voltstack_site--reference--group-009.md#canonical-3002321111223310-0003032010322213-1013313301220322-3232020210112233-1302223101010231-0302033300211120-2122023313113213-0101210003131222): complete subsection reference.

<a id="canonical-1220021122023331-3121200013221200-1011123300000032-3133311332011121-2301120221202010-1113223301110330-3121120321323212-2121222033110300"></a>

## Next pages — interface_list / 011022030031 / 4

- [local_control_plane.bgp_config.peers.external.interface_list.interfaces](data-sources--voltstack_site--reference--group-009.md#canonical-3002321111223310-0003032010322213-1013313301220322-3232020210112233-1302223101010231-0302033300211120-2122023313113213-0101210003131222)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3002321111223310-0003032010322213-1013313301220322-3232020210112233-1302223101010231-0302033300211120-2122023313113213-0101210003131222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333022101200130-3001313003231003-3223330213302313-1122031122332320-0313020011211310-3201130022223001-1301103232011300-1102103012230303"></a>

## local_control_plane.bgp_config.peers.external.interface_list.interfaces — interfaces / 030021013132 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [local_control_plane.bgp_config.peers.external.interface_list](data-sources--voltstack_site--reference--group-009.md#canonical-0101123232112131-3200301120023312-1212313300313020-0113321213022331-1100111312221132-0000011020002230-2330231001003001-2032122210112203)
- local_control_plane.bgp_config.peers.external.interface_list.interfaces

<a id="canonical-3123001201020230-2210113330212223-2132221002231233-0032102013310202-3033200333321030-0331020201221002-0023111302222212-1301331101322012"></a>

Type: `"list"`. Computed.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

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

<a id="canonical-0323120013110031-2130022030130111-2213122001122000-3132112231220123-3131201210012133-0223132020311100-1031230332111112-0321223212020231"></a>

## Direct properties — interfaces / 030021013132 / 3

<a id="canonical-1230122312310101-3323223003111313-2132232001121011-3132211032032212-3112223230010220-0010312212002232-3231200321013220-3211030013320211"></a>

<a id="canonical-0333021330121033-2020132203011331-2222023200101113-0233001022102013-3202021300031122-0322230021022002-2030100311223223-0112213333103202"></a>

## name property — interfaces / 030021013132 / 4

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

<a id="canonical-2200111333332103-0033000201201220-0031100100133200-0230300002222122-3300110201100222-3221030003033320-0111003202032331-2030321330131212"></a>

<a id="canonical-3331231232032211-2002110303230131-3320312330303022-0130303301132100-1201001022133120-1101310131213012-2020002121302302-1100301301110020"></a>

## namespace property — interfaces / 030021013132 / 5

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

<a id="canonical-2302301131120101-3312233321032032-2223311103312321-0233331102103221-2020311123213031-2321023221231032-0311010211203301-1230231203113011"></a>

<a id="canonical-2313002321002212-1302311032323033-0101321311330321-1033133213330220-3233310003203332-1220222212112020-2331223201313032-0013022033021000"></a>

## tenant property — interfaces / 030021013132 / 6

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

<a id="canonical-3303222122312030-3021013002323030-0101311331332302-2310300100002112-3202231121302110-2122301201011302-3131023133122221-2021301220033101"></a>

## Next pages — interfaces / 030021013132 / 7

- [local_control_plane.bgp_config.peers.external.interface_list](data-sources--voltstack_site--reference--group-009.md#canonical-0101123232112131-3200301120023312-1212313300313020-0113321213022331-1100111312221132-0000011020002230-2330231001003001-2032122210112203)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2320022202022210-0332101100020023-0020122201230130-1113332221311231-2123020002313213-2201313311132333-3001130123320220-0320002003103000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003220113112002-1130120220332230-0030221121013232-2012133210000333-2122320302203000-1312113032112021-0012322000013311-0333322113333100"></a>

## local_control_plane.bgp_config.peers.external.no_authentication — no_authentication / 122311102303 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- local_control_plane.bgp_config.peers.external.no_authentication

<a id="canonical-2322130332211212-2013330010020221-2233120300121211-0021330213133120-0232313213021211-3213322030300330-1222121120301031-0123333031113312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no authentication.

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

<a id="canonical-1200322201210303-1113102301012223-1232211003301303-0033122122221322-3020032311131200-1322313013303330-3222132303222102-3300030230023131"></a>

## Direct properties — no_authentication / 122311102303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302232031012222-0033311122330222-1030002312300313-1210211131222132-1300000012021111-1330233000300033-0330303002120311-1031203021312112"></a>

## Next pages — no_authentication / 122311102303 / 4

- [local_control_plane.bgp_config.peers.external](data-sources--voltstack_site--reference--group-009.md#canonical-0321100202101203-3110113011001130-1130223112300301-0313311001013200-2313300130311310-3012122221113022-0013322311322002-3300131200203032)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1210200222013002-2333330303222230-3300230230131300-1232322120231233-3201210010102332-1132101220313021-0312330321333302-1133300320121213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232312000301020-1233211211311133-3120023311303220-3032323310110010-3101020130300000-0301102003312000-1201203023121210-2323332310021102"></a>

## local_control_plane.bgp_config.peers.metadata — metadata / 213321103113 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- local_control_plane.bgp_config.peers.metadata

<a id="canonical-1011332331123332-3101000103333322-1132322301201013-1231113030010302-2020320200301012-2322123111302022-3101330130103010-0101002203002112"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

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

<a id="canonical-0200100203033131-0021123231300232-2122320323311332-1112103233310102-3323330302331100-1302112331331112-0213303003022221-2000021323232211"></a>

## Direct properties — metadata / 213321103113 / 3

<a id="canonical-3311132320222313-2020211120132022-1010323032122311-2133201113133232-2211320001301201-3030002212222103-2021010311211021-3020311303300211"></a>

<a id="canonical-0030312110132300-3121133112101222-2100032300200312-3123300130133311-3111010020320223-2022222101112322-2011030121220011-2222022302330120"></a>

## description_spec property — metadata / 213321103113 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3223033021111132-3220121123013210-3221111000232001-1102220003030103-3113230332013232-2333231303112031-0131303011101100-3330033331310011"></a>

<a id="canonical-3120221110232322-0030322221031002-0013213100003003-1131203203202331-2101332202223103-2100213132302102-0000320002000030-3003001302001003"></a>

## name property — metadata / 213321103113 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

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

<a id="canonical-1033000102313010-0230321222310111-0012303112122030-0013322121013030-1312102310003220-3133201313111200-0101103100132110-2002331100333332"></a>

## Next pages — metadata / 213321103113 / 6

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3030031322003002-3310331113031030-1333030211320001-0011320003312210-1133203313103101-2203021020313121-3321113323232230-3002111200320020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231221013101100-1321132310220023-1310000312020312-1332101302011301-1022130312231200-3333120101103010-3023211300033201-0331110112112312"></a>

## local_control_plane.bgp_config.peers.passive_mode_disabled — passive_mode_disabled / 120331303020 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- local_control_plane.bgp_config.peers.passive_mode_disabled

<a id="canonical-0132331220300321-0231230202211133-2323030313103221-0030323311313212-3311020032001313-0323320102233303-1220320011323001-0003022103312112"></a>

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

<a id="canonical-0333331120121022-1031321021112302-2022101303122213-1130013203120111-2113202201102111-0200302031200131-3021331003103020-0322021230131211"></a>

## Direct properties — passive_mode_disabled / 120331303020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220301212011200-3332330332133230-3231100100220133-0303300232210220-1111210331132203-1221132320312211-0222233321321200-0302130120111133"></a>

## Next pages — passive_mode_disabled / 120331303020 / 4

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1013002103211301-2302230310013122-0113201300210220-2300111003331321-0220020303021323-2322203102212012-2310110221101000-1010132131310313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211120122023103-3003223323122303-2333211312202230-3212311302333112-2233120303031223-2001333001301323-0330032002330320-3121120133322213"></a>

## local_control_plane.bgp_config.peers.passive_mode_enabled — passive_mode_enabled / 232120320022 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- local_control_plane.bgp_config.peers.passive_mode_enabled

<a id="canonical-2322223230322221-2312233333201233-0033331310001012-3202323302001201-2323330310032020-2202322310133222-1113031000120013-0321303003022013"></a>

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

<a id="canonical-3110300112211132-1022000320013100-0322131022333103-2320310111110123-2302020122121300-2133333101321122-3300111202013100-3011211223311222"></a>

## Direct properties — passive_mode_enabled / 232120320022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230121100102012-2320001003202003-3101231122033313-1220010113101232-0311033333020002-3230101131000200-3000113132313323-2323033203200100"></a>

## Next pages — passive_mode_enabled / 232120320022 / 4

- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2001023103302213-2103112012320003-2200301222330212-0131031010310000-0223223132103300-1232301311310232-2131320333110112-3233212131323232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210213323111131-1321332310023303-1122222132020302-2332020200120201-3110300212101102-1211303202001301-2333110013200320-3023210201013312"></a>

## local_control_plane.bgp_config.peers.routing_policies — routing_policies / 303120002032 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- local_control_plane.bgp_config.peers.routing_policies

<a id="canonical-0222022303022200-1221120313301000-2311022302202012-0232221301211033-2121010132123010-1303033133103131-2330201103222332-0122310002201200"></a>

Type: `"single"`. Computed.

List of rules which can be applied on all or particular nodes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1031133032130231-2133022011332312-1222310200133220-3331033320121022-2322101322223031-2131130220133202-1330110211020300-0013122131022220"></a>

## Direct properties — routing_policies / 303120002032 / 3

- [route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-3032323230032201-2333123200302210-2103002023033311-2133211101201032-3110310023102330-1111210233031121-1023220002122211-2113320002132323): complete subsection reference.

<a id="canonical-0202231003120133-0330222233321100-2211120021112203-1032123322222021-0331122012211221-1100010310003021-3033301201211010-1010130332123212"></a>

## Next pages — routing_policies / 303120002032 / 4

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-3032323230032201-2333123200302210-2103002023033311-2133211101201032-3110310023102330-1111210233031121-1023220002122211-2113320002132323)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3032323230032201-2333123200302210-2103002023033311-2133211101201032-3110310023102330-1111210233031121-1023220002122211-2113320002132323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300023312332130-2321011211321121-1203213320033032-0231201221300010-2030110220102330-1232002321100312-0320023121022021-2010003113231103"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy — route_policy / 101022113033 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--reference--group-009.md#canonical-2001023103302213-2103112012320003-2200301222330212-0131031010310000-0223223132103300-1232301311310232-2131320333110112-3233212131323232)
- local_control_plane.bgp_config.peers.routing_policies.route_policy

<a id="canonical-0110022123113210-2101012313232233-0023221112333021-0132103022031233-0102333032220220-2310330002301102-2322022033200222-2231300121311100"></a>

Type: `"list"`. Computed.

Policy configuration for this feature.

Upstream description:

Route policy to be applied.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-0221022301313001-2031113202001123-0100221120001322-3203322020332321-3330321103003123-2320123203300202-3111213120031120-0120221012023333"></a>

## Direct properties — route_policy / 101022113033 / 3

- [all_nodes](data-sources--voltstack_site--reference--group-009.md#canonical-0002310233203123-0103012121302133-3311022211101212-2100102201302322-2313000023323010-2202113100002131-2121223001123330-0100101132232122): complete subsection reference.

- [inbound](data-sources--voltstack_site--reference--group-009.md#canonical-0110332133020101-1211310033331033-2211133200001123-0300123313133231-0323133002003302-1301013212131232-0333131201021123-2130100233221122): complete subsection reference.

- [node_name](data-sources--voltstack_site--reference--group-009.md#canonical-3122001120111212-1131100132120202-0032000302223102-1300033121230003-1113030323021013-0231332123122100-1123230322230112-1301223330131333): complete subsection reference.

- [object_refs](data-sources--voltstack_site--reference--group-009.md#canonical-3001133133110010-0303123001111211-3230311222123002-0221122202133031-3002103023230102-3200222332312321-0302322001033320-3211232133023001): complete subsection reference.

- [outbound](data-sources--voltstack_site--reference--group-009.md#canonical-0123230330313110-0310231220311100-1020122022201021-0300101320223002-2223011311110332-3202310033100133-0102130010113010-2131102320010012): complete subsection reference.

<a id="canonical-0210332310113103-2111021100003321-0020003231301100-1231210303310132-1110032130031011-1133230012323313-2000122331200233-0311103102021100"></a>

## Next pages — route_policy / 101022113033 / 4

- [local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes](data-sources--voltstack_site--reference--group-009.md#canonical-0002310233203123-0103012121302133-3311022211101212-2100102201302322-2313000023323010-2202113100002131-2121223001123330-0100101132232122)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound](data-sources--voltstack_site--reference--group-009.md#canonical-0110332133020101-1211310033331033-2211133200001123-0300123313133231-0323133002003302-1301013212131232-0333131201021123-2130100233221122)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name](data-sources--voltstack_site--reference--group-009.md#canonical-3122001120111212-1131100132120202-0032000302223102-1300033121230003-1113030323021013-0231332123122100-1123230322230112-1301223330131333)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs](data-sources--voltstack_site--reference--group-009.md#canonical-3001133133110010-0303123001111211-3230311222123002-0221122202133031-3002103023230102-3200222332312321-0302322001033320-3211232133023001)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound](data-sources--voltstack_site--reference--group-009.md#canonical-0123230330313110-0310231220311100-1020122022201021-0300101320223002-2223011311110332-3202310033100133-0102130010113010-2131102320010012)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--reference--group-009.md#canonical-2001023103302213-2103112012320003-2200301222330212-0131031010310000-0223223132103300-1232301311310232-2131320333110112-3233212131323232)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0002310233203123-0103012121302133-3311022211101212-2100102201302322-2313000023323010-2202113100002131-2121223001123330-0100101132232122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300320313031130-3231001120311321-0031321303302222-0333310303022310-1323003020033300-2311022311313012-0111320133021320-2122033230321223"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes — all_nodes / 112101130203 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--reference--group-009.md#canonical-2001023103302213-2103112012320003-2200301222330212-0131031010310000-0223223132103300-1232301311310232-2131320333110112-3233212131323232)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-3032323230032201-2333123200302210-2103002023033311-2133211101201032-3110310023102330-1111210233031121-1023220002122211-2113320002132323)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.all_nodes

<a id="canonical-3101313301032031-2301222020003100-3133333103111113-2013012212222223-1311212211100011-3303233300132310-3101302102121210-0301301112102230"></a>

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

<a id="canonical-2012311121131033-0310320213113211-0330311002312123-0321113302030121-2031312200202113-0211202312012023-3203203133201110-0103202313011202"></a>

## Direct properties — all_nodes / 112101130203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320031121031201-1321332123203123-3202023120111301-3221221302100321-2220312300331103-2102213323032321-0103000133310003-3113332133302200"></a>

## Next pages — all_nodes / 112101130203 / 4

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-3032323230032201-2333123200302210-2103002023033311-2133211101201032-3110310023102330-1111210233031121-1023220002122211-2113320002132323)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0110332133020101-1211310033331033-2211133200001123-0300123313133231-0323133002003302-1301013212131232-0333131201021123-2130100233221122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312030320232331-1030111320300111-1123111231201102-2312313231320230-1030231031101221-2333133221000022-3231030102120010-1001210313100320"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound — inbound / 011103013123 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--reference--group-009.md#canonical-2001023103302213-2103112012320003-2200301222330212-0131031010310000-0223223132103300-1232301311310232-2131320333110112-3233212131323232)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-3032323230032201-2333123200302210-2103002023033311-2133211101201032-3110310023102330-1111210233031121-1023220002122211-2113320002132323)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.inbound

<a id="canonical-2331212301133203-1302221302111223-1113031310020200-1300311310021001-1323213303013011-0231222302123202-0312002101312001-0123222320113222"></a>

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

<a id="canonical-1311003333200323-1201033121303102-1000133303211320-3113100233021300-2030032331212010-0220122201330311-3323020103020103-3200202311231313"></a>

## Direct properties — inbound / 011103013123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221021300112223-1220302233311131-0102223010310030-0103032033122311-0212013222013102-3023211021323121-0210330311110320-0313001100302212"></a>

## Next pages — inbound / 011103013123 / 4

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-3032323230032201-2333123200302210-2103002023033311-2133211101201032-3110310023102330-1111210233031121-1023220002122211-2113320002132323)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3122001120111212-1131100132120202-0032000302223102-1300033121230003-1113030323021013-0231332123122100-1123230322230112-1301223330131333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313232322311122-1232113031030231-1021011110002133-2120330301030122-3032323010011132-3002023222032033-1113210023323021-2333333222112020"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name — node_name / 033320031231 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--reference--group-009.md#canonical-2001023103302213-2103112012320003-2200301222330212-0131031010310000-0223223132103300-1232301311310232-2131320333110112-3233212131323232)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-3032323230032201-2333123200302210-2103002023033311-2133211101201032-3110310023102330-1111210233031121-1023220002122211-2113320002132323)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name

<a id="canonical-1000022330122301-0030111303000230-1332030121333102-0033132000013113-3201333121322231-3200011333331022-3132031221113203-0213030030102030"></a>

Type: `"single"`. Computed.

List of nodes on which BGP routing policy has to be applied.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1133300231200331-0221012313331330-1012021002100130-1310111320013123-1331322212103203-2203001201213230-3321130132011303-3303210203222010"></a>

## Direct properties — node_name / 033320031231 / 3

<a id="canonical-3100300003221222-1133103232232023-1221112120022313-2232102320130230-0210102212020310-1100323321032113-2021010301302310-0321212113213232"></a>

<a id="canonical-1011230131100323-3311101133301203-3221000313231212-0232030311312112-3131213310020133-0021011213033113-3130231202301000-1322221202111020"></a>

## node property — node_name / 033320031231 / 4

Type: `["list", "string"]`. Computed.

Select BGP Session on which policy will be applied.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1001013021302133-2212233332023103-2220221330202132-0113230121312001-2333003132221221-1001032013001333-0010211332210223-0110032001003012"></a>

## Next pages — node_name / 033320031231 / 5

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-3032323230032201-2333123200302210-2103002023033311-2133211101201032-3110310023102330-1111210233031121-1023220002122211-2113320002132323)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3001133133110010-0303123001111211-3230311222123002-0221122202133031-3002103023230102-3200222332312321-0302322001033320-3211232133023001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310110110212030-0032302311033031-1133133233013331-0221032110333232-0022322300202332-0101132000032302-0103230220031113-3212231121000131"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs — object_refs / 113133332212 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--reference--group-009.md#canonical-2001023103302213-2103112012320003-2200301222330212-0131031010310000-0223223132103300-1232301311310232-2131320333110112-3233212131323232)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-3032323230032201-2333123200302210-2103002023033311-2133211101201032-3110310023102330-1111210233031121-1023220002122211-2113320002132323)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs

<a id="canonical-3201132022133133-2123303101210222-2210231221003110-0030101133233312-0113021020123331-3332201123212330-3120130302210130-1030110130333201"></a>

Type: `"list"`. Computed.

BGP routing policy. Select route policy to apply.

Upstream description:

Select route policy to apply.

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

<a id="canonical-0223232331120203-0201323103230123-2001111011011301-1303020103211001-2010120010132030-0011021233022010-2302330010032020-3022011100232001"></a>

## Direct properties — object_refs / 113133332212 / 3

<a id="canonical-1312020202222121-2131323102231001-1103311031322033-3000010223201213-2232011323323033-1001303132211111-0322212030132023-3103123213111003"></a>

<a id="canonical-3303223221220100-1113321112301331-0223333103132031-1222031320300320-2011300001223300-0020213320201122-1203301303311100-3203200031320110"></a>

## kind property — object_refs / 113133332212 / 4

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

<a id="canonical-2303012231210110-0220030003310130-3313302201012121-0020300003203203-1301332202031330-0322020323220233-0131332322332120-0230300002210102"></a>

<a id="canonical-1101203300313023-0212013103101222-2123202233130223-0023313300233030-3313311300033232-3233313123200230-1133222102100332-2321232310301330"></a>

## name property — object_refs / 113133332212 / 5

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

<a id="canonical-0210303321311213-0330233020211320-0033331003022201-0111313102011332-3012113021011301-2232200203001331-1313220212232002-3323122030122022"></a>

<a id="canonical-0231033130133033-3322112001303101-3033230203300311-3013213113001120-0000233133312330-2121321300133023-2012210131222232-2012330313013232"></a>

## namespace property — object_refs / 113133332212 / 6

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

<a id="canonical-0000002331000022-2023003033311120-0211303011003211-1130031111300223-1313011232101302-1131201103031131-1213302312233100-0112000323311121"></a>

<a id="canonical-1303322111020021-2331031231020102-1102032011333002-2131031212311311-3211333100231003-0230032230011311-3010110331303113-3102230301331202"></a>

## tenant property — object_refs / 113133332212 / 7

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

<a id="canonical-1133102211302330-1003102312001310-3033231132303102-2310311121223220-3010320320011000-2101000130302032-0023100330302323-2032310201103101"></a>

<a id="canonical-0010111211311302-1132020201321031-3301123112330302-0130201303030000-2012022302102022-1301013320330331-1113113201102231-0223112031330221"></a>

## uid property — object_refs / 113133332212 / 8

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

<a id="canonical-0311212331223231-0031021102100121-0213122303121000-3320130110322132-2300320210200232-2313103020211302-2203031010202022-3001311021112020"></a>

## Next pages — object_refs / 113133332212 / 9

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-3032323230032201-2333123200302210-2103002023033311-2133211101201032-3110310023102330-1111210233031121-1023220002122211-2113320002132323)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0123230330313110-0310231220311100-1020122022201021-0300101320223002-2223011311110332-3202310033100133-0102130010113010-2131102320010012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030001123120211-2012003200023200-0332201201212220-2330013011100323-2212322313102310-3232233203322303-3331103102102111-2213330132111101"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound — outbound / 311102123111 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [local_control_plane.bgp_config](data-sources--voltstack_site--reference--group-009.md#canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100)
- [local_control_plane.bgp_config.peers](data-sources--voltstack_site--reference--group-009.md#canonical-0123123202030010-1332222013232130-2203012003011301-2000322333332233-3101011130303110-2321323210311030-2102312311332311-1332112101232300)
- [local_control_plane.bgp_config.peers.routing_policies](data-sources--voltstack_site--reference--group-009.md#canonical-2001023103302213-2103112012320003-2200301222330212-0131031010310000-0223223132103300-1232301311310232-2131320333110112-3233212131323232)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-3032323230032201-2333123200302210-2103002023033311-2133211101201032-3110310023102330-1111210233031121-1023220002122211-2113320002132323)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound

<a id="canonical-0200320311301022-1210222010203303-1023332221031202-0033201132032323-3200301031011230-3133123200022222-2210312033030120-0310010323200323"></a>

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

<a id="canonical-0111001012230033-0103122202202011-3123101302322133-2322131231010011-0302111200200220-2130313122321013-1232212321033130-1312002213023333"></a>

## Direct properties — outbound / 311102123111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133330020022222-3220002230321100-3003300032111201-0011030012303200-3100120101302031-3201210031322023-3010212310103033-3310010323323203"></a>

## Next pages — outbound / 311102123111 / 4

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](data-sources--voltstack_site--reference--group-009.md#canonical-3032323230032201-2333123200302210-2103002023033311-2133211101201032-3110310023102330-1111210233031121-1023220002122211-2113320002132323)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0130201121003322-1033330032323001-0032001100333232-3213223220333223-1300030112323300-2202310100232112-0222103313212322-0110022010301010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130203011012231-2120031223210222-0310212020330112-0302021022203200-1131133113002301-2032301221313123-2100323222132232-1011310032031033"></a>

## local_control_plane.inside_vn — inside_vn / 021102201100 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- local_control_plane.inside_vn

<a id="canonical-0311112021130231-2113313100223020-1003330100120002-1233332123100332-2232303322111102-0020231311022220-2103020102220101-2200021201001323"></a>

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

<a id="canonical-0312213000211031-2301230120122201-3121032122001100-2000121223133333-1032303023302002-2122031311323303-2223110030132320-0220301210222333"></a>

## Direct properties — inside_vn / 021102201100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113110200313012-2223303203102213-3201031312231120-3012331212233011-3113221312202001-3313310011333311-2020022200233033-0103320010231232"></a>

## Next pages — inside_vn / 021102201100 / 4

- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3231111032012121-2322313113231303-1010332200001131-2302202320230012-3132121121003131-1233002013011133-3330003222212231-2102020101131222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022221210021001-3033211213110102-0030232202112101-1303121111203031-2323031323113112-2000001303322012-1301323111203333-1011221232213120"></a>

## local_control_plane.outside_vn — outside_vn / 022302200213 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- local_control_plane.outside_vn

<a id="canonical-2220133330130232-3122113203003002-1313102000122300-2112221112103310-3131220130130103-0313031001121031-0001203220210221-2132013131133221"></a>

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

<a id="canonical-0001232220310030-0030300123211311-0111233100330011-2133100102311120-1002330233222233-3003022023213200-0303121321010123-3300221203130020"></a>

## Direct properties — outside_vn / 022302200213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323320213100301-1330131322002232-1011110210222331-1303110003320320-0100132232230000-1223221200212033-2012332030212111-1213212313122120"></a>

## Next pages — outside_vn / 022302200213 / 4

- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1100301300011210-1122120223131022-0032330322112300-1102133312212303-2031101220113112-1233113010003231-0222131002303000-0111231120132313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301232221213202-2113021301231132-0100233102233101-1133101313323312-2100220131230022-3303110302221313-1333132222210023-2021013320330112"></a>

## log_receiver — log_receiver / 022230012230 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- log_receiver

<a id="canonical-2033313120203100-3100321013232211-1332113112103102-3332202121230001-2320322003111230-1220123232000113-3313011121022201-0103321003101113"></a>

Type: `"single"`. Computed.

\[OneOf: log\_receiver, logs\_streaming\_disabled\] Type establishes a direct reference from one
object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [log_receiver](data-sources--voltstack_site--reference--group-009.md#canonical-2033313120203100-3100321013232211-1332113112103102-3332202121230001-2320322003111230-1220123232000113-3313011121022201-0103321003101113)
- [logs_streaming_disabled](data-sources--voltstack_site--reference--group-009.md#canonical-0300230333133110-3123121323033302-1112102133332100-1032023003230231-0200110333130101-1113112221100032-2323010203131102-2202010313011302)

Select alternatives according to the provider validators above.

<a id="canonical-2233113311111112-0031203220301000-1010213111312323-0010232312001230-1323111230213312-1012222212332233-2210130222123222-2301121230120121"></a>

## Direct properties — log_receiver / 022230012230 / 3

<a id="canonical-1213310100312011-0232212213032031-1302332200031121-0023211032012030-1121120000002103-0213031212220021-2100201032321113-3300031302320031"></a>

<a id="canonical-0002101033033101-3033221002133233-2132000132033333-1130212232012133-0332121022021323-3111033122031210-3101022013203113-0130331213212210"></a>

## name property — log_receiver / 022230012230 / 4

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

<a id="canonical-1020020003132110-1213120003312202-2012010132223230-2203320011330113-3323220331300210-1001011331200322-3013012211211011-0223131302231122"></a>

<a id="canonical-0313223020120032-3130213202313112-0003112313113012-3202210331131321-0202312032032112-0200231232200002-0220102232120101-3221313010313121"></a>

## namespace property — log_receiver / 022230012230 / 5

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

<a id="canonical-3210123001222030-2212223131221332-2020200123031202-3302311113131130-1330132101113330-0231120120123020-2222001110301320-3222233331321220"></a>

<a id="canonical-1331132332031212-0121132000020013-0221113131103203-2120002220103220-3103113211331333-2202020330133001-2231012213000322-2220212322111210"></a>

## tenant property — log_receiver / 022230012230 / 6

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

<a id="canonical-0231300023011313-0020221120033323-2202300022013010-0131332233033030-3132130000321110-3301313213312302-2100101201013132-0132013122220200"></a>

## Next pages — log_receiver / 022230012230 / 7

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1100220121332220-0132031112133000-1301301112301011-2112212331211203-0122110000030122-2312103221213312-2022032120012312-1133010222332101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031101030330321-0332031030022132-0110100031022303-0122110020100030-0033222323330001-2310223120322333-1222112203002010-2230033202100123"></a>

## logs_streaming_disabled — logs_streaming_disabled / 113333301023 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- logs_streaming_disabled

<a id="canonical-0300230333133110-3123121323033302-1112102133332100-1032023003230231-0200110333130101-1113112221100032-2323010203131102-2202010313011302"></a>

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

<a id="canonical-3300222333210210-3102302313320013-3333112022322330-2120211022123022-1321301122001121-2131222322202222-2120130300113022-2100233310133311"></a>

## Direct properties — logs_streaming_disabled / 113333301023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220112332212310-1210130303222220-1133201310233013-1111202210232212-2301311123010123-2301300032000103-2333322322123233-0013002110332121"></a>

## Next pages — logs_streaming_disabled / 113333301023 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1003223011333012-3011003222112223-1001002232011322-1312002320233201-0300202321101310-1133031313320121-1320013310302222-1103131023001123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233101213020103-3003120012323033-2233210021313220-3031310102230123-3120222321101221-2111231232120333-0330011221212311-1313010100223322"></a>

## master_node_configuration — master_node_configuration / 222312211320 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- master_node_configuration

<a id="canonical-2022321123330010-2112200020231102-0021101100230220-2110220032121111-0033232132121011-0133303331221031-1122201211221313-3222033231111032"></a>

Type: `"list"`. Computed.

Master Nodes. Configuration of master nodes.

Upstream description:

Configuration of master nodes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3"
  }
}
```

<a id="canonical-3313133301030221-3100300310312103-3211300332312010-0232110023313330-3032332122320130-3212220331032232-1211321130222123-3113311021331322"></a>

## Direct properties — master_node_configuration / 222312211320 / 3

<a id="canonical-1311303001122233-3010201012012200-0231311222211221-2121000030032011-0322111302130233-3020311133131230-3312110102102330-3231103021012121"></a>

<a id="canonical-1112322003212121-2101101330322000-1132302333230121-3003033111102211-3120201312331030-2221301010202030-3021002230031023-2310002101023203"></a>

## name property — master_node_configuration / 222312211320 / 4

Type: `"string"`. Computed.

Name. Names of master node.

Upstream description:

Names of master node.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1301131213123102-3030202301230210-3020031133213031-3133313212022033-1131323323213103-2113110113222022-2210300213203102-3201030202120211"></a>

<a id="canonical-2222320333321200-2210102300033133-0233102002220020-1103112310001323-2210331120033003-2212223112233031-0102132213101301-3021010213322333"></a>

## public_ip property — master_node_configuration / 222312211320 / 5

Type: `"string"`. Computed.

IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.

Upstream description:

IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.

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

<a id="canonical-2203232133113303-2210001302133111-3301002222131200-0320213022001300-3220132021201300-2230202200303333-0012222030002302-2232211210012203"></a>

## Next pages — master_node_configuration / 222312211320 / 6

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0120020123011223-1311000213300223-2220032023023232-0032303201302330-0123302133032113-3033112202120311-0320232003212212-3311033312103210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332332111032110-3231232321001321-3133003011131023-0333120200013300-0113211201021303-1300021010022001-1233322121000223-2212032023231101"></a>

## no_bond_devices — no_bond_devices / 010212112300 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- no_bond_devices

<a id="canonical-3103232323001023-3201121010200223-2110312103200202-1132133321000201-3033000201333120-1022211100022310-0012310123300123-3211131333000221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no bond devices.

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

<a id="canonical-1023023323021323-2013010330132300-0210330321012010-1210001011212322-2322223213220221-2002132113131211-3303122201121123-2232031332221020"></a>

## Direct properties — no_bond_devices / 010212112300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011133330223333-3101320000122021-2213321311221022-3233033131132232-1310233203202120-0011312112131311-3011312301013211-3103121133001002"></a>

## Next pages — no_bond_devices / 010212112300 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2111131212322100-1100210202033210-0300330213220310-1223033122312231-3201223203323210-0320122322220222-1211023320230222-1310012211130330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113323130021201-0101322131012322-1201022031210100-1203223301303130-3203123113001321-3220122113130203-3130221203000313-0110223222100010"></a>

## no_k8s_cluster — no_k8s_cluster / 000122013302 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- no_k8s_cluster

<a id="canonical-1013201012231322-1122033102111202-3123031303112213-3001013311103230-3000012313111132-3211203112221102-3332233221001323-0030300122300212"></a>

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

<a id="canonical-3330031132200220-2111002012320222-2013232230200201-1112203132103132-3303230223022021-2010323230332202-3020213220010100-0302022230001002"></a>

## Direct properties — no_k8s_cluster / 000122013302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111321133112210-2322020133013201-2203322021113010-2031323113331210-3212202130231001-2010203031122011-1201001201110303-3102322321101200"></a>

## Next pages — no_k8s_cluster / 000122013302 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1321303301011331-3111230333030121-1023310203120102-1032331222030111-1112320330012000-1133103002222310-2000121010031030-0122132302233222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220333032113010-0010032103130230-0013230300012010-0213331221222313-3002131013200200-0100333232002211-2120223202130020-2213231331302131"></a>

## no_local_control_plane — no_local_control_plane / 022123130211 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- no_local_control_plane

<a id="canonical-1002001200332030-1120100012133311-3302033110222333-1033022120110300-3020111201310130-2303203020323123-0121331333222122-1200321120212102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no local control plane.

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

<a id="canonical-1030230321102302-3330132210103321-1232321332302101-3230032233021002-2221100302102113-2201222113130311-2132000020111012-3210112022213213"></a>

## Direct properties — no_local_control_plane / 022123130211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233101210001111-0133322323101122-1200210200322220-0122113201203100-0320301310222111-1333010213220012-2333230300123320-0311331212133321"></a>

## Next pages — no_local_control_plane / 022123130211 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3012201203013310-2202301221032113-3211212303232023-3001230312211120-0223122330211122-3321300212212230-1123030231222122-0310112232022022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112213013022130-1123312323322231-2322100301112123-2012332112121200-1203202302132131-3000031102213010-2111303232322302-2122233322220011"></a>

## offline_survivability_mode — offline_survivability_mode / 113030230113 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- offline_survivability_mode

<a id="canonical-3010131003101323-1023121113300011-0002011012031333-3020112220120333-1033121023320101-0311130032113001-0330323000001033-2100200212220003"></a>

Type: `"single"`. Computed.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7..

Upstream description:

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-offline_survivability_mode_choice": "[\"enable_offline_survivability_mode\",\"no_offline_survivability_mode\"]"
}
```

<a id="canonical-3032302001201201-3211310333033101-3112331230311102-1201211011113023-1300330232333030-3021122323100121-0331010323133111-1333321103232223"></a>

## Direct properties — offline_survivability_mode / 113030230113 / 3

- [enable_offline_survivability_mode](data-sources--voltstack_site--reference--group-009.md#canonical-1331030231322120-0212323100303100-2112223232321001-1110003311110112-3122001210210001-3223101222110003-1321132123012221-1023012013213012): complete subsection reference.

- [no_offline_survivability_mode](data-sources--voltstack_site--reference--group-010.md#canonical-2213313110020223-0011332022223020-3310200013233110-0012221031220221-1023011112213300-1203030003333122-0330220112113303-3330213112123223): complete subsection reference.

<a id="canonical-1212001303202122-3133101131301031-0130312101120333-0100030133221102-2021221321030221-2313022201221010-3313301013131112-1332213030323100"></a>

## Next pages — offline_survivability_mode / 113030230113 / 4

- [offline_survivability_mode.enable_offline_survivability_mode](data-sources--voltstack_site--reference--group-009.md#canonical-1331030231322120-0212323100303100-2112223232321001-1110003311110112-3122001210210001-3223101222110003-1321132123012221-1023012013213012)
- [offline_survivability_mode.no_offline_survivability_mode](data-sources--voltstack_site--reference--group-010.md#canonical-2213313110020223-0011332022223020-3310200013233110-0012221031220221-1023011112213300-1203030003333122-0330220112113303-3330213112123223)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1331030231322120-0212323100303100-2112223232321001-1110003311110112-3122001210210001-3223101222110003-1321132123012221-1023012013213012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033322001012001-2101312311033303-0332011323120123-1023033021322203-2331030033322222-0011223322231101-1203101033321311-1210230310121111"></a>

## offline_survivability_mode.enable_offline_survivability_mode — enable_offline_survivability_mode / 123302101233 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [offline_survivability_mode](data-sources--voltstack_site--reference--group-009.md#canonical-3012201203013310-2202301221032113-3211212303232023-3001230312211120-0223122330211122-3321300212212230-1123030231222122-0310112232022022)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-2202022310201203-0200133233103321-2321032232220213-3133322120133301-1212311320033121-0023203021210013-2202213031221020-1101232123213220"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable offline survivability mode.

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

<a id="canonical-2310330020121310-1323332010221100-3121123020003021-0320033021002020-0113310201302323-3220312031323303-1200203321301220-3121310023330112"></a>

## Direct properties — enable_offline_survivability_mode / 123302101233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320020323332222-2231123020133312-1101103223113123-1332223331033010-0330120010231012-0322002320202331-0100111222322212-2002001120330333"></a>

## Next pages — enable_offline_survivability_mode / 123302101233 / 4

- [offline_survivability_mode](data-sources--voltstack_site--reference--group-009.md#canonical-3012201203013310-2202301221032113-3211212303232023-3001230312211120-0223122330211122-3321300212212230-1123030231222122-0310112232022022)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
