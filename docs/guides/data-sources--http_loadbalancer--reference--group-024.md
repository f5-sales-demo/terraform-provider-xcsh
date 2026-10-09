---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1200211123003010-1233100203111302-0210123300133003-2122300031101021-1201313133330132-2231222222013231-2311202021020223-0023311303233311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-023.md#canonical-2001202332221232-3333020122211202-2233223223033313-1230201322322310-0013102301231100-3030011221013313-3003132022332331-1003223103230131)
- rate_limit.rate_limiter.action_block

<a id="canonical-1301120202122223-3301230230013331-2203012323303003-2231022111013230-3230010111030130-2012222111121303-0001000012122012-2330320221202003"></a>

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

<a id="canonical-1312013330013333-2210031230202222-1323121212022332-0010133203303102-3303110322323332-0300022202023212-2000023231102122-3132321222133301"></a>

### Direct properties for `rate_limit.rate_limiter.action_block`

- [hours](data-sources--http_loadbalancer--reference--group-024.md#canonical-2223130102331122-1332301030101033-1212201121231300-3223021022100021-1032222203131222-2332203332322013-2131010302302310-3321232203312013): complete subsection reference.

- [minutes](data-sources--http_loadbalancer--reference--group-024.md#canonical-0203101011221301-0203010102021300-1123313100010010-2232332331323111-3331021300331332-1223020231332321-1211023111100013-3310331220032020): complete subsection reference.

- [seconds](data-sources--http_loadbalancer--reference--group-024.md#canonical-3212030121100303-3320023121033103-2100032013011113-2221112200011220-2330232010102122-2033123003221110-0113021011233303-3113100031130322): complete subsection reference.

<a id="canonical-2223130102331122-1332301030101033-1212201121231300-3223021022100021-1032222203131222-2332203332322013-2131010302302310-3321232203312013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.hours` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-023.md#canonical-2001202332221232-3333020122211202-2233223223033313-1230201322322310-0013102301231100-3030011221013313-3003132022332331-1003223103230131)
- [rate_limit.rate_limiter.action_block](data-sources--http_loadbalancer--reference--group-024.md#canonical-1200211123003010-1233100203111302-0210123300133003-2122300031101021-1201313133330132-2231222222013231-2311202021020223-0023311303233311)
- rate_limit.rate_limiter.action_block.hours

<a id="canonical-1000033002130203-1021030233001100-1323211322012313-2020120131221033-2133222201113031-2010021221012021-0222132030333332-3133003031003321"></a>

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

<a id="canonical-3233123112033120-1223332202330212-0030300020002330-3312330211121113-0133132332223302-2111210323221311-0121122021110333-3122213323112211"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.hours`

<a id="canonical-3300001321220103-0303123310131031-3202121102011003-3311120322030033-3330212221302020-1123303220211200-0200001303211200-0122322011231121"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0203101011221301-0203010102021300-1123313100010010-2232332331323111-3331021300331332-1223020231332321-1211023111100013-3310331220032020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.minutes` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-023.md#canonical-2001202332221232-3333020122211202-2233223223033313-1230201322322310-0013102301231100-3030011221013313-3003132022332331-1003223103230131)
- [rate_limit.rate_limiter.action_block](data-sources--http_loadbalancer--reference--group-024.md#canonical-1200211123003010-1233100203111302-0210123300133003-2122300031101021-1201313133330132-2231222222013231-2311202021020223-0023311303233311)
- rate_limit.rate_limiter.action_block.minutes

<a id="canonical-0302310302113313-2100030110320320-2023213320301303-3030132130310330-2020332031302313-2330230332332000-3130320320131102-2211203312232121"></a>

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

<a id="canonical-2332133203300322-0300123303211222-2101122330123013-3301122301321333-1102000331203012-2301130003003031-3310330120312321-0123210012103002"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.minutes`

<a id="canonical-1323220020103331-0131313320323321-1133100310312331-0312111323203120-3012303030223323-0121320111130202-0121330002121210-1003013322312003"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3212030121100303-3320023121033103-2100032013011113-2221112200011220-2330232010102122-2033123003221110-0113021011233303-3113100031130322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.seconds` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-023.md#canonical-2001202332221232-3333020122211202-2233223223033313-1230201322322310-0013102301231100-3030011221013313-3003132022332331-1003223103230131)
- [rate_limit.rate_limiter.action_block](data-sources--http_loadbalancer--reference--group-024.md#canonical-1200211123003010-1233100203111302-0210123300133003-2122300031101021-1201313133330132-2231222222013231-2311202021020223-0023311303233311)
- rate_limit.rate_limiter.action_block.seconds

<a id="canonical-0321023001200032-1033003131302101-2031230033112313-3130233133102022-0010332201210011-1122310013223232-0231022322201132-1302101301223230"></a>

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

<a id="canonical-3020311103100113-1333100203212023-2113231120021312-1212310121120222-1130210312003031-1113301300023201-1332233122231302-3011111110303230"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.seconds`

<a id="canonical-2202311213200213-1322013132212120-3330321212211033-2031330202113103-0200102301303001-2211222222323023-3110213223103200-2232302001323332"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3000320012303032-0010200332232323-1213021202222230-2302102123202102-0111031331020030-0022220102102200-1000221213100020-2102210110132210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-023.md#canonical-2001202332221232-3333020122211202-2233223223033313-1230201322322310-0013102301231100-3030011221013313-3003132022332331-1003223103230131)
- rate_limit.rate_limiter.disabled

<a id="canonical-0201210232323213-3210112310022221-0130020111320103-3131102221020322-0231210211100320-1003101030320232-1232233111031221-3023001200202310"></a>

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

<a id="canonical-3102303111103122-1313202011123133-1200301113011002-1220133132121012-0003013002301132-1010100032120132-3002323202122132-3030102232323133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.leaky_bucket` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-023.md#canonical-2001202332221232-3333020122211202-2233223223033313-1230201322322310-0013102301231100-3030011221013313-3003132022332331-1003223103230131)
- rate_limit.rate_limiter.leaky_bucket

<a id="canonical-0110033302120101-3323030022322313-3231023120020123-0231331222030001-0031013201033003-3113022023223223-2231233232131230-2200123022000321"></a>

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

<a id="canonical-3303231201200320-2100220323123301-3112332201232200-1311113100210010-3131102021100010-1233103202133031-1031230200300002-0010311031033232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.token_bucket` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--reference--group-023.md#canonical-2001202332221232-3333020122211202-2233223223033313-1230201322322310-0013102301231100-3030011221013313-3003132022332331-1003223103230131)
- rate_limit.rate_limiter.token_bucket

<a id="canonical-3001013133231003-3301213311121101-1110013332332332-0333033131133123-3223120210120020-0200101202111233-1330311303103110-1313333001300131"></a>

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

<a id="canonical-3223003312110002-0221122213310322-3103023023321131-1323121100223120-2100103010223110-1231210111013100-0101202220221310-2300203220330323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- ring_hash

<a id="canonical-0323322103101000-0200310331030220-3311002233001310-1210203133131012-0201131201133211-2020032232103330-0213201232312002-3301321221002222"></a>

Type: `"single"`. Computed.

Hash Policy List. List of hash policy rules.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3233232202213010-2311122110112221-2111202301101331-1022031220002013-0132002233033230-2321231220302121-2102211223233022-3202321330132113"></a>

### Direct properties for `ring_hash`

- [hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-3100323310012210-0321302221233031-2333102213101122-0102010120300121-1121012300222112-2310201300213122-1023032330011310-3332013312113120): complete subsection reference.

<a id="canonical-3100323310012210-0321302221233031-2333102213101122-0102010120300121-1121012300222112-2310201300213122-1023032330011310-3332013312113120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ring_hash](data-sources--http_loadbalancer--reference--group-024.md#canonical-3223003312110002-0221122213310322-3103023023321131-1323121100223120-2100103010223110-1231210111013100-0101202220221310-2300203220330323)
- ring_hash.hash_policy

<a id="canonical-0203101032012321-3223201123333003-0130301333323012-3020201223312212-0033011023330212-0033023012313022-2113003133010301-0323333023200220"></a>

Type: `"list"`. Computed.

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0113232011211030-2101220103113032-3101013312131003-3312232320303133-3110230230000332-2200222010101103-2320330221021331-0023313232302311"></a>

### Direct properties for `ring_hash.hash_policy`

- [cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-3123121022030112-0020313323322102-2002103303311120-1111230331223131-3222200023223133-2310102220210203-0211201100030221-0300221122200332): complete subsection reference.

<a id="canonical-2013312211230121-2033311301212303-0333002222002110-3112211001300220-3333132123030120-2130020033200203-1121232001032132-2012302331333101"></a>

<a id="canonical-3121112003302203-1132120022312322-3331233332233130-1330232112301023-0221210211302232-1300321023231300-1201233123231312-0123120103332333"></a>

#### `ring_hash.hash_policy.header_name` property

Type: `"string"`. Computed.

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2133023110303200-0211201232230200-0320031323331120-1000131021212221-2021132033122003-0202210213003103-2110101310222202-1022130200222031"></a>

<a id="canonical-3333122120203113-2321301230023102-0301110132113120-2232023113223121-2220013110200102-2132303203212133-0311130330022113-1230010221031032"></a>

#### `ring_hash.hash_policy.source_ip` property

Type: `"bool"`. Computed.

Exclusive with \[cookie header\_name\] Hash based on source IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0100133322033230-1331123112133102-3100222122002221-3200321023323220-3110022003032302-3213302333211013-2102012330003320-2320213222000032"></a>

<a id="canonical-0103223012320301-0313202223303120-2323010020100323-2303203030102230-1323022310121031-3013013100021210-1213312300130232-0321230301033022"></a>

#### `ring_hash.hash_policy.terminal` property

Type: `"bool"`. Computed.

Terminal. Specify if its a terminal policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3123121022030112-0020313323322102-2002103303311120-1111230331223131-3222200023223133-2310102220210203-0211201100030221-0300221122200332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ring_hash](data-sources--http_loadbalancer--reference--group-024.md#canonical-3223003312110002-0221122213310322-3103023023321131-1323121100223120-2100103010223110-1231210111013100-0101202220221310-2300203220330323)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-3100323310012210-0321302221233031-2333102213101122-0102010120300121-1121012300222112-2310201300213122-1023032330011310-3332013312113120)
- ring_hash.hash_policy.cookie

<a id="canonical-3313023213310230-2021020003032021-0203111022000033-3123222102003313-1312120313102210-3020200313300200-0133323132231303-0100220033202321"></a>

Type: `"single"`. Computed.

Two types of cookie affinity:

&#8203;1. Passive. Takes a cookie that's present in the cookies header and hashes on its value.

&#8203;2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from
the client in its response to the client, based on the endpoint the request gets sent to. The client
then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure
these requests GET sent to the same endpoint. The cookie is generated by hashing the source and
destination ports and addresses so that multiple independent HTTP2 streams on the same connection
will independently receive the same cookie, even if they arrive simultaneously.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-httponly": "[\"add_httponly\",\"ignore_httponly\"]",
  "x-ves-oneof-field-samesite": "[\"ignore_samesite\",\"samesite_lax\",\"samesite_none\",\"samesite_strict\"]",
  "x-ves-oneof-field-secure": "[\"add_secure\",\"ignore_secure\"]"
}
```

<a id="canonical-1212012002031112-1223203003300003-0022231102230032-2121200133330213-0121311002033300-0020313020322011-1033323333300020-1231131202311011"></a>

### Direct properties for `ring_hash.hash_policy.cookie`

- [add_httponly](data-sources--http_loadbalancer--reference--group-024.md#canonical-2012331031213103-0032132322110121-1113002113222021-0100130332232033-3020020121201320-2233232230332322-1112310222211223-3012021112232223): complete subsection reference.

- [add_secure](data-sources--http_loadbalancer--reference--group-024.md#canonical-0013203223120330-0332201011030003-2013233220331133-0222033001023230-1023102213220213-0200123331200031-2210121102223313-1000320131002301): complete subsection reference.

- [ignore_httponly](data-sources--http_loadbalancer--reference--group-024.md#canonical-0133023231312203-1332203210310021-2003232030131220-1010213212332220-1312202221032030-3311303123121221-2300003302323023-1012002002032230): complete subsection reference.

- [ignore_samesite](data-sources--http_loadbalancer--reference--group-024.md#canonical-0301100203322320-2103332121320321-2312203200102032-0312122331121301-0022220112202311-0220212212101033-2002022113001123-3113131331320011): complete subsection reference.

- [ignore_secure](data-sources--http_loadbalancer--reference--group-024.md#canonical-1233200302110232-1230300130300233-2002311033320021-1300200210331311-3100233120220121-1231303322211321-2120103002232231-3322133110220210): complete subsection reference.

<a id="canonical-0132230322212122-3002302021101211-2300121201310302-3331110022111120-3303331231123131-0113230002201333-1103220033300201-1320001203302331"></a>

<a id="canonical-2123320120302132-3132000110013101-0100001110102331-1021132211130322-1013130032201330-1211212231032332-0011211121123203-2030102202220131"></a>

#### `ring_hash.hash_policy.cookie.name` property

Type: `"string"`. Computed.

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1100010102021330-0132311333130011-3120000111222001-1120002313313030-0003212023201212-3122231222200133-3123212013311313-2002030233032203"></a>

<a id="canonical-3213100320233002-0313012103011331-3221202130300313-3013020031323003-0232203020211302-2022223121221313-0333013033230203-3111101013331212"></a>

#### `ring_hash.hash_policy.cookie.path` property

Type: `"string"`. Computed.

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](data-sources--http_loadbalancer--reference--group-024.md#canonical-0321201013120231-3312013322112011-2201012022120312-0221202232330122-2322002002223032-3032222122101331-1222023300012120-0112113032201011): complete subsection reference.

- [samesite_none](data-sources--http_loadbalancer--reference--group-024.md#canonical-2001121322033001-3123222013030103-2032331130312113-1100212110001020-3132023123001210-2120231023033220-1320321113312133-3123130110200112): complete subsection reference.

- [samesite_strict](data-sources--http_loadbalancer--reference--group-024.md#canonical-3032202301332133-1220303223100123-2100213233322111-1230103012111132-1013303010310103-2031310133332332-1103221033022300-2103331103020202): complete subsection reference.

<a id="canonical-3100303030013123-3200231111210012-2321222102321031-0233101303231023-0000323222020002-2121131033033102-1313123121301001-2312233312030202"></a>

<a id="canonical-0310303120121131-3002113000201222-1033113233303223-3003123332030103-0220131222331122-2313322311222332-3013000012013133-1012133330002323"></a>

#### `ring_hash.hash_policy.cookie.ttl` property

Type: `"number"`. Computed.

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2012331031213103-0032132322110121-1113002113222021-0100130332232033-3020020121201320-2233232230332322-1112310222211223-3012021112232223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.add_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ring_hash](data-sources--http_loadbalancer--reference--group-024.md#canonical-3223003312110002-0221122213310322-3103023023321131-1323121100223120-2100103010223110-1231210111013100-0101202220221310-2300203220330323)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-3100323310012210-0321302221233031-2333102213101122-0102010120300121-1121012300222112-2310201300213122-1023032330011310-3332013312113120)
- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-3123121022030112-0020313323322102-2002103303311120-1111230331223131-3222200023223133-2310102220210203-0211201100030221-0300221122200332)
- ring_hash.hash_policy.cookie.add_httponly

<a id="canonical-1302130023303203-1322310330330100-1311021120022003-3230331032030101-3022211130203320-3223000203122333-3100213131111303-3102310010012330"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add httponly.

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

<a id="canonical-0013203223120330-0332201011030003-2013233220331133-0222033001023230-1023102213220213-0200123331200031-2210121102223313-1000320131002301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.add_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ring_hash](data-sources--http_loadbalancer--reference--group-024.md#canonical-3223003312110002-0221122213310322-3103023023321131-1323121100223120-2100103010223110-1231210111013100-0101202220221310-2300203220330323)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-3100323310012210-0321302221233031-2333102213101122-0102010120300121-1121012300222112-2310201300213122-1023032330011310-3332013312113120)
- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-3123121022030112-0020313323322102-2002103303311120-1111230331223131-3222200023223133-2310102220210203-0211201100030221-0300221122200332)
- ring_hash.hash_policy.cookie.add_secure

<a id="canonical-3322220322203213-2311303201220032-2131013310221102-2022013122323320-1102122233021311-3121020110230212-0032120111303030-1110203010130101"></a>

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

<a id="canonical-0133023231312203-1332203210310021-2003232030131220-1010213212332220-1312202221032030-3311303123121221-2300003302323023-1012002002032230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.ignore_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ring_hash](data-sources--http_loadbalancer--reference--group-024.md#canonical-3223003312110002-0221122213310322-3103023023321131-1323121100223120-2100103010223110-1231210111013100-0101202220221310-2300203220330323)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-3100323310012210-0321302221233031-2333102213101122-0102010120300121-1121012300222112-2310201300213122-1023032330011310-3332013312113120)
- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-3123121022030112-0020313323322102-2002103303311120-1111230331223131-3222200023223133-2310102220210203-0211201100030221-0300221122200332)
- ring_hash.hash_policy.cookie.ignore_httponly

<a id="canonical-2300200220203130-1133132330022203-2033011122000023-1310003133233011-2313232203033231-3122102331333201-0231321113023112-2012003313302330"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore httponly.

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

<a id="canonical-0301100203322320-2103332121320321-2312203200102032-0312122331121301-0022220112202311-0220212212101033-2002022113001123-3113131331320011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.ignore_samesite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ring_hash](data-sources--http_loadbalancer--reference--group-024.md#canonical-3223003312110002-0221122213310322-3103023023321131-1323121100223120-2100103010223110-1231210111013100-0101202220221310-2300203220330323)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-3100323310012210-0321302221233031-2333102213101122-0102010120300121-1121012300222112-2310201300213122-1023032330011310-3332013312113120)
- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-3123121022030112-0020313323322102-2002103303311120-1111230331223131-3222200023223133-2310102220210203-0211201100030221-0300221122200332)
- ring_hash.hash_policy.cookie.ignore_samesite

<a id="canonical-2231203203311122-2313301100113201-2112221200000200-1212001111103022-1031323312311212-0201130003132221-3000002102002332-1130133033320313"></a>

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

<a id="canonical-1233200302110232-1230300130300233-2002311033320021-1300200210331311-3100233120220121-1231303322211321-2120103002232231-3322133110220210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.ignore_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ring_hash](data-sources--http_loadbalancer--reference--group-024.md#canonical-3223003312110002-0221122213310322-3103023023321131-1323121100223120-2100103010223110-1231210111013100-0101202220221310-2300203220330323)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-3100323310012210-0321302221233031-2333102213101122-0102010120300121-1121012300222112-2310201300213122-1023032330011310-3332013312113120)
- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-3123121022030112-0020313323322102-2002103303311120-1111230331223131-3222200023223133-2310102220210203-0211201100030221-0300221122200332)
- ring_hash.hash_policy.cookie.ignore_secure

<a id="canonical-0010223011323212-3033100300313223-2203021223232220-3231220331011323-1030332000122222-0321012212012102-2303203223003022-3202332232101230"></a>

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

<a id="canonical-0321201013120231-3312013322112011-2201012022120312-0221202232330122-2322002002223032-3032222122101331-1222023300012120-0112113032201011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.samesite_lax` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ring_hash](data-sources--http_loadbalancer--reference--group-024.md#canonical-3223003312110002-0221122213310322-3103023023321131-1323121100223120-2100103010223110-1231210111013100-0101202220221310-2300203220330323)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-3100323310012210-0321302221233031-2333102213101122-0102010120300121-1121012300222112-2310201300213122-1023032330011310-3332013312113120)
- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-3123121022030112-0020313323322102-2002103303311120-1111230331223131-3222200023223133-2310102220210203-0211201100030221-0300221122200332)
- ring_hash.hash_policy.cookie.samesite_lax

<a id="canonical-3222211320023231-2322121223110013-1300303012010133-2323000203111323-1120001100001103-2003303222103102-0032201200133303-1111101133121001"></a>

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

<a id="canonical-2001121322033001-3123222013030103-2032331130312113-1100212110001020-3132023123001210-2120231023033220-1320321113312133-3123130110200112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.samesite_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ring_hash](data-sources--http_loadbalancer--reference--group-024.md#canonical-3223003312110002-0221122213310322-3103023023321131-1323121100223120-2100103010223110-1231210111013100-0101202220221310-2300203220330323)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-3100323310012210-0321302221233031-2333102213101122-0102010120300121-1121012300222112-2310201300213122-1023032330011310-3332013312113120)
- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-3123121022030112-0020313323322102-2002103303311120-1111230331223131-3222200023223133-2310102220210203-0211201100030221-0300221122200332)
- ring_hash.hash_policy.cookie.samesite_none

<a id="canonical-3203233200032202-1020230202212000-1033031113202121-1212022300330123-3232003120000323-0223232132031110-0301222120110030-0003300301030231"></a>

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

<a id="canonical-3032202301332133-1220303223100123-2100213233322111-1230103012111132-1013303010310103-2031310133332332-1103221033022300-2103331103020202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.samesite_strict` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ring_hash](data-sources--http_loadbalancer--reference--group-024.md#canonical-3223003312110002-0221122213310322-3103023023321131-1323121100223120-2100103010223110-1231210111013100-0101202220221310-2300203220330323)
- [ring_hash.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-3100323310012210-0321302221233031-2333102213101122-0102010120300121-1121012300222112-2310201300213122-1023032330011310-3332013312113120)
- [ring_hash.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-3123121022030112-0020313323322102-2002103303311120-1111230331223131-3222200023223133-2310102220210203-0211201100030221-0300221122200332)
- ring_hash.hash_policy.cookie.samesite_strict

<a id="canonical-0110303210132021-0020312320212031-2031331100331320-3202230213301333-0122203010211330-2113120323113330-2133103301000333-2120312222321313"></a>

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

<a id="canonical-2100021222003130-0100011231321302-3110002003123302-1203123023020200-3232130013121212-2002101032122031-3302132222233331-3002022203023111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `round_robin` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- round_robin

<a id="canonical-0113223033112001-3033310213310132-3331011031030032-1122222000211010-1301100002000310-2102331021103223-3122121232331301-2211310020000133"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for round robin. Defaults to \`map\[\]\`. Server applies default when
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

<a id="canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- routes

<a id="canonical-1322330112031301-1003000103123112-1230022230200011-2131232122303103-1032300033222313-1202311110311111-3112202313111313-2212303100013003"></a>

Type: `"list"`. Computed.

Routes allow users to define match condition on a path and/or HTTP method to either forward matching
traffic to origin pool or redirect matching traffic to a different URL or respond directly to
matching traffic.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0111312211123103-0022313002311131-3222020100221320-0101221323101303-3312103232110230-2032333031302123-1120302022023133-3021101020133100"></a>

### Direct properties for `routes`

- [custom_route_object](data-sources--http_loadbalancer--reference--group-024.md#canonical-2111330031230112-2100030222001233-0032133121122121-0332123032310212-0211320223220312-2120000231001112-0112012320313200-3022311310331230): complete subsection reference.

- [direct_response_route](data-sources--http_loadbalancer--reference--group-024.md#canonical-0210011221102330-1323020232130122-0032102300113122-0021323302222233-0301002133020010-2320213030311223-0231003110231332-2310320220131232): complete subsection reference.

- [redirect_route](data-sources--http_loadbalancer--reference--group-024.md#canonical-3031130012330222-2232013233232232-2002103331033200-2211001300312033-3312120221121332-1323201212001023-1020112222113103-2303212300333332): complete subsection reference.

- [route_state_disabled](data-sources--http_loadbalancer--reference--group-025.md#canonical-0212220310112030-0001203011201320-1321132003202101-0030012323230322-2230131230233030-0331131330200131-0210031030010303-1022222000131200): complete subsection reference.

- [route_state_enabled](data-sources--http_loadbalancer--reference--group-025.md#canonical-3010003023310013-1332311212102210-3312231203102211-2303031101111132-1232020321022120-0112033112132211-2100030122110111-3111232221322303): complete subsection reference.

- [simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103): complete subsection reference.

<a id="canonical-2111330031230112-2100030222001233-0032133121122121-0332123032310212-0211320223220312-2120000231001112-0112012320313200-3022311310331230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom_route_object` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- routes.custom_route_object

<a id="canonical-0231310302223312-1211010102001332-1333322231100020-2100323312233310-3031333010110300-3033310022001201-2003033210103101-3131001231232232"></a>

Type: `"single"`. Computed.

A custom route uses a route object created outside of this view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

<a id="canonical-0021022122233032-3001313222021120-0130220031220330-3002230220112311-1210332001010232-2213011033203010-1312311300321232-3133123100310031"></a>

### Direct properties for `routes.custom_route_object`

- [caching_disable](data-sources--http_loadbalancer--reference--group-024.md#canonical-1233032302330113-1313220303300313-2232233310110310-2230033220202223-3203021213210330-0000203221300113-3323322331011132-3123123323321011): complete subsection reference.

- [caching_inherit](data-sources--http_loadbalancer--reference--group-024.md#canonical-2220313001123202-1003202322103322-0330020303202232-0100032220113021-3112102121310310-1020233202202100-0320222001120301-2100013200332200): complete subsection reference.

- [route_ref](data-sources--http_loadbalancer--reference--group-024.md#canonical-0111321123220120-1031123133323313-1100223210013102-1130111101323202-2222203110120120-3321113123330312-0002013223013132-3013310122312011): complete subsection reference.

<a id="canonical-1233032302330113-1313220303300313-2232233310110310-2230033220202223-3203021213210330-0000203221300113-3323322331011132-3123123323321011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom_route_object.caching_disable` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.custom_route_object](data-sources--http_loadbalancer--reference--group-024.md#canonical-2111330031230112-2100030222001233-0032133121122121-0332123032310212-0211320223220312-2120000231001112-0112012320313200-3022311310331230)
- routes.custom_route_object.caching_disable

<a id="canonical-2000012023320323-0210311020133323-3120132111201230-0321321133210103-0021230120122013-0203111020022312-1120213320333131-1201313222102121"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching disable.

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

<a id="canonical-2220313001123202-1003202322103322-0330020303202232-0100032220113021-3112102121310310-1020233202202100-0320222001120301-2100013200332200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom_route_object.caching_inherit` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.custom_route_object](data-sources--http_loadbalancer--reference--group-024.md#canonical-2111330031230112-2100030222001233-0032133121122121-0332123032310212-0211320223220312-2120000231001112-0112012320313200-3022311310331230)
- routes.custom_route_object.caching_inherit

<a id="canonical-0012213021213110-2032211303303233-1313030233002203-0212102130221331-0111113222202203-3123301301320313-3023102030112023-2031023200212302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching inherit.

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

<a id="canonical-0111321123220120-1031123133323313-1100223210013102-1130111101323202-2222203110120120-3321113123330312-0002013223013132-3013310122312011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom_route_object.route_ref` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.custom_route_object](data-sources--http_loadbalancer--reference--group-024.md#canonical-2111330031230112-2100030222001233-0032133121122121-0332123032310212-0211320223220312-2120000231001112-0112012320313200-3022311310331230)
- routes.custom_route_object.route_ref

<a id="canonical-1121122030203330-1033200300201233-1333212321332013-1212212330320013-0003001031100113-3032101233210330-3231212332122132-3201000233112323"></a>

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

<a id="canonical-0223321231131011-0321232223201123-2121102031130310-2002323031201102-1110020301331222-2321030020031012-0100103003203132-2221011300303122"></a>

### Direct properties for `routes.custom_route_object.route_ref`

<a id="canonical-3123123112130031-0203131001012003-0313201120122112-0321112011023110-0131103312221330-2011232232313211-0101103311312212-2320113200020012"></a>

#### `routes.custom_route_object.route_ref.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0211033011201211-2102020210332122-1011101001320120-3131022021122102-3112323311202113-1103122233222011-3033210233330003-3112123321223102"></a>

<a id="canonical-0033023110110322-1132000320323330-0320002300011223-2013203200013223-1011123100320300-1300211321002230-2202022130132102-2111223010110331"></a>

#### `routes.custom_route_object.route_ref.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3020333231201203-2222232232122000-2113031021122130-3320021201002101-2100013102310011-2002032210312001-0111333221120133-1223120312022211"></a>

<a id="canonical-2320130322313131-0032023110021322-2122233003330231-3001213101302002-0122223302121200-0002231313213033-3313201223201201-2321103321022022"></a>

#### `routes.custom_route_object.route_ref.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0210011221102330-1323020232130122-0032102300113122-0021323302222233-0301002133020010-2320213030311223-0231003110231332-2310320220131232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.direct_response_route` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- routes.direct_response_route

<a id="canonical-1321321000333110-1103311331130133-1013231313302132-1100300031011011-3101211033020020-1131100300032332-2223020130220200-0022010122300030"></a>

Type: `"single"`. Computed.

A direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3321302123200013-1023000313313221-3220332033301022-0030203303321030-1131320033211033-1213132111000100-3233300112012023-2221211011331023"></a>

### Direct properties for `routes.direct_response_route`

- [headers](data-sources--http_loadbalancer--reference--group-024.md#canonical-1312003333310320-0131033011312212-3332203311320023-1310003020010333-3031223213011201-3123033231122132-3222103001020231-0130211211103030): complete subsection reference.

<a id="canonical-1323332000011230-3113333312020002-3230101023213221-0033002111111320-2300023131023123-2113110131231311-3210022011120030-0001123030011012"></a>

<a id="canonical-2211232222331033-0131011312103033-0211133320300210-0210003331013030-2322003222333210-1023022333110322-2313103112001010-3333101331310312"></a>

#### `routes.direct_response_route.http_method` property

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](data-sources--http_loadbalancer--reference--group-024.md#canonical-1123212110221010-3110101320002001-3010221233120213-2303302112311110-3022220131231011-3120311113001120-1012003033032323-1331223222133033): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-024.md#canonical-0201233220200032-1302120031321311-1303113301011322-3310203313023103-1322331203232301-3030321032323131-3002003111302133-0011322000012000): complete subsection reference.

- [route_direct_response](data-sources--http_loadbalancer--reference--group-024.md#canonical-2331322212101113-0300120012002100-1021210210302331-3012003202132110-0311320022310330-3332012312231230-3332321122031331-1103300221333021): complete subsection reference.

<a id="canonical-1312003333310320-0131033011312212-3332203311320023-1310003020010333-3031223213011201-3123033231122132-3222103001020231-0130211211103030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.direct_response_route.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.direct_response_route](data-sources--http_loadbalancer--reference--group-024.md#canonical-0210011221102330-1323020232130122-0032102300113122-0021323302222233-0301002133020010-2320213030311223-0231003110231332-2310320220131232)
- routes.direct_response_route.headers

<a id="canonical-2333331303033033-1321231123213310-1000303222333013-2130123100032322-1330220032012133-2300010030100231-1332020102033102-0231110030013110"></a>

Type: `"list"`. Computed.

Headers. List of (key, value) headers.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minItems": 0,
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

<a id="canonical-1023333010020232-0022030230231221-2011003312021132-3022132013121303-0100010113100330-3332220311312033-0303002200333133-1331330332121202"></a>

### Direct properties for `routes.direct_response_route.headers`

<a id="canonical-3223113223311130-1310220301031112-2120010132330332-1320000123110100-0111300013232312-2000332210101201-3200021031110213-1311110112120120"></a>

#### `routes.direct_response_route.headers.exact` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1220020013033123-2300133001102211-2320122002010223-0213201033100121-1012232212032221-3231231001031301-3302003211022123-2233321330102113"></a>

<a id="canonical-1100032032222212-0112311131232021-2031202002231030-1122102033320113-2133211230032301-2202300330103211-2020202002332323-1033311311310010"></a>

#### `routes.direct_response_route.headers.invert_match` property

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

<a id="canonical-3131220132211003-0232211031310213-3122120112311303-1231100222103101-3100301203213233-2002320121013033-0223332033123330-1010002312302211"></a>

<a id="canonical-0323201202313330-0310000332200123-3111320222211132-0021233120002013-1320211133121220-0321120222322010-0202133233210202-0100212221000111"></a>

#### `routes.direct_response_route.headers.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3033332000213031-2321130113100310-2233221122211023-2002311223211233-1330300310101102-2312301132231012-2101210220113201-3322032122101201"></a>

<a id="canonical-3303202332210033-0230321113330130-3311023111213212-3111013102102030-2320211123011223-2313030222212310-3303200010203233-1122023032221212"></a>

#### `routes.direct_response_route.headers.presence` property

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

<a id="canonical-1022111302011221-3123320311103103-3021102302301213-3133303000033331-2222330130120221-2211220032223111-2233331032020003-3333101032233200"></a>

<a id="canonical-1022030112323122-3121001230121322-2122103233202001-0121131021320011-2112022133211312-0130332122333333-1011222120303030-0130322300023210"></a>

#### `routes.direct_response_route.headers.regex` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1123212110221010-3110101320002001-3010221233120213-2303302112311110-3022220131231011-3120311113001120-1012003033032323-1331223222133033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.direct_response_route.incoming_port` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.direct_response_route](data-sources--http_loadbalancer--reference--group-024.md#canonical-0210011221102330-1323020232130122-0032102300113122-0021323302222233-0301002133020010-2320213030311223-0231003110231332-2310320220131232)
- routes.direct_response_route.incoming_port

<a id="canonical-1200210100223202-1310102112120123-3312000102230020-1223003103220122-1110032000222312-0101120010303230-0330110130122120-0222233320113130"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-0321122103011320-1010103002000122-1012102311002233-3320320330120311-2300313232102103-3330303210110333-3331313103321222-2200021022101111"></a>

### Direct properties for `routes.direct_response_route.incoming_port`

- [no_port_match](data-sources--http_loadbalancer--reference--group-024.md#canonical-0320121333110200-2200032000230212-1301312201000100-1020310330332121-0112330322120302-1120000330122003-0020202223313230-1322003132021112): complete subsection reference.

<a id="canonical-1331111012013022-1233223100200312-0333222303111222-0200020321221130-0123033001000310-2100312001133011-0003022332030013-3122301120202201"></a>

<a id="canonical-0332331031102100-1201100220123132-0301331321320200-0212223011110110-3013230010121220-3201133322132030-2202211222201013-0300213111300012"></a>

#### `routes.direct_response_route.incoming_port.port` property

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2101112323101101-2113311330200223-3230232011201020-2322022220323030-1023200233002331-0131000100333032-3023102331203322-1100200331030112"></a>

<a id="canonical-1302111013223212-3313203312113022-1230110330101221-3023320012320332-0122103200311221-2333123311323310-1230323311201100-3130231122300021"></a>

#### `routes.direct_response_route.incoming_port.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-0320121333110200-2200032000230212-1301312201000100-1020310330332121-0112330322120302-1120000330122003-0020202223313230-1322003132021112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.direct_response_route.incoming_port.no_port_match` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.direct_response_route](data-sources--http_loadbalancer--reference--group-024.md#canonical-0210011221102330-1323020232130122-0032102300113122-0021323302222233-0301002133020010-2320213030311223-0231003110231332-2310320220131232)
- [routes.direct_response_route.incoming_port](data-sources--http_loadbalancer--reference--group-024.md#canonical-1123212110221010-3110101320002001-3010221233120213-2303302112311110-3022220131231011-3120311113001120-1012003033032323-1331223222133033)
- routes.direct_response_route.incoming_port.no_port_match

<a id="canonical-1230203211312203-0221223220221113-2100230330000022-0201322313121202-2302103233133022-0102303200322300-3302113133003201-0230110301122333"></a>

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

<a id="canonical-0201233220200032-1302120031321311-1303113301011322-3310203313023103-1322331203232301-3030321032323131-3002003111302133-0011322000012000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.direct_response_route.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.direct_response_route](data-sources--http_loadbalancer--reference--group-024.md#canonical-0210011221102330-1323020232130122-0032102300113122-0021323302222233-0301002133020010-2320213030311223-0231003110231332-2310320220131232)
- routes.direct_response_route.path

<a id="canonical-3100000220303120-0003210011010200-0131033011032010-3031233022333001-0202302212113023-2301013112122313-1213223131311003-3323301113100303"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-1330232210213122-0012031301302310-0113131301222320-0030010200310303-1311120331001331-2022311200333132-2200303133202220-1032112211031110"></a>

### Direct properties for `routes.direct_response_route.path`

<a id="canonical-2033011011220232-3333100322100023-2131300020231233-1303100023230102-1000321112030333-0321200010210111-2311012003330230-2232330321011233"></a>

#### `routes.direct_response_route.path.path` property

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-1210131312113022-2222001013200302-3121322033202030-1101202203303023-1330200131303123-1133202332022133-2123012122322322-1120221332100332"></a>

<a id="canonical-1200121210133122-1103301003010303-0230323110011012-1113212300003031-3223013231023300-1200322302330333-1022032011002212-3010310011000232"></a>

#### `routes.direct_response_route.path.prefix` property

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1221321123302021-1011113231103301-0332320101221112-0110200203311220-2020013220110200-3303001230020002-2332023132331123-3312133013130222"></a>

<a id="canonical-1323022130202132-2001000101132032-3211301220121330-2112302133012201-0120001232022311-0131330130200331-2100222300313300-1033111331011022"></a>

#### `routes.direct_response_route.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2331322212101113-0300120012002100-1021210210302331-3012003202132110-0311320022310330-3332012312231230-3332321122031331-1103300221333021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.direct_response_route.route_direct_response` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.direct_response_route](data-sources--http_loadbalancer--reference--group-024.md#canonical-0210011221102330-1323020232130122-0032102300113122-0021323302222233-0301002133020010-2320213030311223-0231003110231332-2310320220131232)
- routes.direct_response_route.route_direct_response

<a id="canonical-3333112220101012-3020122210031311-3201020130332011-3232320013211121-3023101111232332-3302133130202222-2030133103032101-1221220033230223"></a>

Type: `"single"`. Computed.

Send this direct response in case of route match action is direct response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3002313032223223-3102113232023222-3131313330132032-2200031303103300-1002303223130111-3003221332122230-1331001133002211-1323331020220201"></a>

### Direct properties for `routes.direct_response_route.route_direct_response`

<a id="canonical-0032232212133300-2220221113102033-0131213100222301-0121001120000221-2003310120220320-0030331303200311-1312121002030321-0232321230012121"></a>

#### `routes.direct_response_route.route_direct_response.response_body_encoded` property

Type: `"string"`. Computed.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in base64 format. The message can be either plain text or HTML. E.g. "&lt;p&gt; Access
Denied &lt;/p&gt;". base64 encoded string URL for this is
string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 65536
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2123031000001100-3121032230331211-2003200302111300-3320331012211310-2023023231233110-0101213102330213-1120301123112213-2100303302000213"></a>

<a id="canonical-0121331220120110-0033003131321320-2111112132320300-0233111202110023-2112201022213023-3323033211000313-1011301033102200-3201100233100121"></a>

#### `routes.direct_response_route.route_direct_response.response_code` property

Type: `"number"`. Computed.

Response Code. Response code to send.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 100
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

<a id="canonical-3031130012330222-2232013233232232-2002103331033200-2211001300312033-3312120221121332-1323201212001023-1020112222113103-2303212300333332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.redirect_route` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- routes.redirect_route

<a id="canonical-1110222022231022-3330310022303302-3312201300120303-0100231013102012-3013313333110221-0132211312123113-3213123312111331-3010311321332100"></a>

Type: `"single"`. Computed.

A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects
the matching traffic to a different URL.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3212312032323120-1123220321111110-1222022221312002-0120221323131011-0312211111113311-3032300031300203-1011230202123121-1130012313122331"></a>

### Direct properties for `routes.redirect_route`

- [headers](data-sources--http_loadbalancer--reference--group-024.md#canonical-1223101330120300-0132100302130130-2210221302103231-2102211101322213-2323033030232003-3012132221012210-0122030120100112-0033002130131022): complete subsection reference.

<a id="canonical-1001020202102102-1331123320222130-0131331013123102-1332202003013111-0233132333330300-1013020303033331-3012103100332322-2121300003012003"></a>

<a id="canonical-2222313021223333-3113131231321132-1223223311221000-1000222232203212-3213132123000123-0003331011001330-1033133330012300-0013310112322203"></a>

#### `routes.redirect_route.http_method` property

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](data-sources--http_loadbalancer--reference--group-024.md#canonical-0111320122003131-0002103122020030-2233231202013221-1103233233033110-2001021310233231-2333311301301323-3032022213102110-1320103011323112): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-024.md#canonical-3220122132100311-0223000001001221-3310131012113020-3312033022131012-0330331203313223-0300133120232310-3100322232100313-3000312211321120): complete subsection reference.

- [route_redirect](data-sources--http_loadbalancer--reference--group-024.md#canonical-3110312013020332-0323110212313021-3313003231012220-0312332112101122-0122103201031233-3003001222133102-1311200122000230-3110313102321223): complete subsection reference.

<a id="canonical-1223101330120300-0132100302130130-2210221302103231-2102211101322213-2323033030232003-3012132221012210-0122030120100112-0033002130131022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.redirect_route.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-024.md#canonical-3031130012330222-2232013233232232-2002103331033200-2211001300312033-3312120221121332-1323201212001023-1020112222113103-2303212300333332)
- routes.redirect_route.headers

<a id="canonical-1121113002130201-2232302301132213-0022011133330201-2030211332201032-1100022313121310-1211020101320332-3201100021312203-0301221312003101"></a>

Type: `"list"`. Computed.

Headers. List of (key, value) headers.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minItems": 0,
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

<a id="canonical-1212030320201032-1130321013113030-0230003120233113-2311220202200103-1301232231330110-0200221112000222-0330022123131102-0323010123022113"></a>

### Direct properties for `routes.redirect_route.headers`

<a id="canonical-1032120312332111-0101113113333113-0203132022123200-0331022312133023-1131003321233330-0123013312010000-0123131033320222-0101103332001330"></a>

#### `routes.redirect_route.headers.exact` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0100130112100013-2200301010302012-0301332122030233-3102213303130103-1133310223030210-1003013032032233-2000203010233101-1000313221001233"></a>

<a id="canonical-1020023021333231-3321303212111132-3033012003003200-0001311020332220-2112123211130223-1201123133230203-3310212201103112-0003023210112132"></a>

#### `routes.redirect_route.headers.invert_match` property

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

<a id="canonical-2211203300333103-1320111001113330-1331312213313102-3033303213131312-0330000103303310-3203033131133111-2123332013332320-1122133001133113"></a>

<a id="canonical-0222320113103321-2301001010223210-3000103311010011-0021022032112021-1303112220200133-2030012102201332-1300033213000120-1323312031020210"></a>

#### `routes.redirect_route.headers.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3220020130002300-2112210210113220-3113223222303313-2010231203332033-0200332032221031-1202021203013313-3002231230131221-0131223123202113"></a>

<a id="canonical-3003331021213102-0310303310001212-1210103230201010-2313213310230331-1231201330222030-3132332222131100-3100022202331312-2233200333211302"></a>

#### `routes.redirect_route.headers.presence` property

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

<a id="canonical-3222333202230001-1201333202022021-1011213300112023-2331000111232322-0121231300201110-0113300232322112-1021323111013221-3331302123120120"></a>

<a id="canonical-2102010102112111-2330010032103101-2330021111201313-0001112301230233-3321330311320022-1120302102221310-0311113322302010-2211033320300232"></a>

#### `routes.redirect_route.headers.regex` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0111320122003131-0002103122020030-2233231202013221-1103233233033110-2001021310233231-2333311301301323-3032022213102110-1320103011323112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.redirect_route.incoming_port` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-024.md#canonical-3031130012330222-2232013233232232-2002103331033200-2211001300312033-3312120221121332-1323201212001023-1020112222113103-2303212300333332)
- routes.redirect_route.incoming_port

<a id="canonical-3321310033000122-0300121213232323-0123220102233231-3333311333030313-3233121311103012-0010211122211232-3331132312032022-0122133313100203"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-3022131213233320-2011311312223211-2322212130122311-0121312233211320-2322320203322222-1111320121113211-3331323102213303-0130133310310122"></a>

### Direct properties for `routes.redirect_route.incoming_port`

- [no_port_match](data-sources--http_loadbalancer--reference--group-024.md#canonical-3301011013321223-3312301012113032-2011011300322033-3322133230312130-2301210230031010-3312303130211011-0110201200033123-3101031101103022): complete subsection reference.

<a id="canonical-0233121312111310-2110332203012132-2032330110231032-0303100231130230-0113120323231001-0013132023300033-1321121333321223-3312003320032331"></a>

<a id="canonical-2011210310233203-1223102132201100-1312332120010130-0102303022012231-2210120120020110-2130320121010033-3331000313202221-3320133012313022"></a>

#### `routes.redirect_route.incoming_port.port` property

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0002100100022300-3202122001313122-2312300213301010-1032021221310313-3102300201132130-1112131230101220-0110223311300222-2012202112203120"></a>

<a id="canonical-0303303220212101-2131301333331133-2232110300212300-2123000332323223-0322112123203100-0023232033223120-1201132112230311-1011033312330032"></a>

#### `routes.redirect_route.incoming_port.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-3301011013321223-3312301012113032-2011011300322033-3322133230312130-2301210230031010-3312303130211011-0110201200033123-3101031101103022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.redirect_route.incoming_port.no_port_match` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-024.md#canonical-3031130012330222-2232013233232232-2002103331033200-2211001300312033-3312120221121332-1323201212001023-1020112222113103-2303212300333332)
- [routes.redirect_route.incoming_port](data-sources--http_loadbalancer--reference--group-024.md#canonical-0111320122003131-0002103122020030-2233231202013221-1103233233033110-2001021310233231-2333311301301323-3032022213102110-1320103011323112)
- routes.redirect_route.incoming_port.no_port_match

<a id="canonical-2222321031033101-3211320230303331-3222112113233231-3213001020121212-2101332001003113-2113320333013201-3132001231021320-3213123321121123"></a>

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

<a id="canonical-3220122132100311-0223000001001221-3310131012113020-3312033022131012-0330331203313223-0300133120232310-3100322232100313-3000312211321120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.redirect_route.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-024.md#canonical-3031130012330222-2232013233232232-2002103331033200-2211001300312033-3312120221121332-1323201212001023-1020112222113103-2303212300333332)
- routes.redirect_route.path

<a id="canonical-2210233332330032-0200311230320032-0210101030231131-3302130010300101-0323220122200200-2213100201122310-2031033322321010-1000212233010332"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-0010220132123102-0121303220113332-3223030132023203-2311203232212222-2113020121322323-2322310331223032-2121130103101113-3210231312122202"></a>

### Direct properties for `routes.redirect_route.path`

<a id="canonical-3110110333211330-2030010030203312-1103231100010023-1022311313132110-0202103133232332-1032330111022102-3300002303030010-3320303322023123"></a>

#### `routes.redirect_route.path.path` property

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-0001022110003313-0111223233113023-1010331303221212-2012200103212031-2300020133113323-0231230201233210-1211313032311133-1010030120312001"></a>

<a id="canonical-0101001101033113-0010233223131302-2200231012002212-1212100030112103-1332221102111303-0320110210033103-0020212120103310-1321213132220333"></a>

#### `routes.redirect_route.path.prefix` property

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3120312320011322-2331013010233320-1230221032000302-3321302003331223-0112012112111312-0211111103213222-1222222030230003-0220112113022321"></a>

<a id="canonical-2232131212333010-3331233022021211-2312220110232320-3300023132010030-0233022322202310-1312231301000112-0012233030330333-2023200222032010"></a>

#### `routes.redirect_route.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3110312013020332-0323110212313021-3313003231012220-0312332112101122-0122103201031233-3003001222133102-1311200122000230-3110313102321223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.redirect_route.route_redirect` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-024.md#canonical-3031130012330222-2232013233232232-2002103331033200-2211001300312033-3312120221121332-1323201212001023-1020112222113103-2303212300333332)
- routes.redirect_route.route_redirect

<a id="canonical-3012132113231133-2131203321213020-3201001033212321-3233221202023311-0312101011231011-1320222201101111-1300120203103303-2322033011012212"></a>

Type: `"single"`. Computed.

Route redirect parameters when match action is redirect.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]",
  "x-ves-oneof-field-redirect_path_choice": "[\"path_redirect\",\"prefix_rewrite\"]"
}
```

<a id="canonical-3030221203121203-3320002022202203-3010012100013212-2102331312123332-1102112011023111-2120131323100121-3332110230202013-3031111112113231"></a>

### Direct properties for `routes.redirect_route.route_redirect`

<a id="canonical-0110230012311202-2330033330120232-1003301001013131-3201131022000021-2220221222232001-3322220321020102-3331333130223213-1120220330212030"></a>

#### `routes.redirect_route.route_redirect.host_redirect` property

Type: `"string"`. Computed.

Swap host part of incoming URL in redirect URL.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1202001022303312-0231113132021102-2333203322133210-0202201331213103-2010123332110121-2103120131233010-2230210210032020-1300331130011210"></a>

<a id="canonical-0110001120203213-3001022003022020-3213303331133300-0133302213300100-0302221000331230-0222332021323111-1011022231233202-0223311210023301"></a>

#### `routes.redirect_route.route_redirect.path_redirect` property

Type: `"string"`. Computed.

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1122003013130120-1001012201002011-1023120313000012-1020231033332123-0323200201233232-3103321022003311-0300121100330023-0003221211223230"></a>

<a id="canonical-1230000232111103-0112310003231000-2220001222010332-3003310230212302-2223020133023133-1103003303110232-1323221110303113-2110001120303302"></a>

#### `routes.redirect_route.route_redirect.prefix_rewrite` property

Type: `"string"`. Computed.

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3332110303003333-1011020213103312-2333013302002001-2303313133211311-0122223102110232-1012122130330310-3220200213121101-3303320102311201"></a>

<a id="canonical-3200022231123122-0103011212012212-2210102200311212-1020213003330320-0030210123320222-0203222110200112-1312011002202322-1130200222322010"></a>

#### `routes.redirect_route.route_redirect.proto_redirect` property

Type: `"string"`. Computed.

\[Enum: incoming-proto|http|https\] Swap protocol part of incoming URL in redirect URL The protocol
can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of
protocol is not done. Possible values are \`incoming-proto\`, \`http\`, \`https\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "incoming-proto",
    "http",
    "https"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  }
}
```

- [remove_all_params](data-sources--http_loadbalancer--reference--group-024.md#canonical-2033023032323312-2011123302223003-2322013001200302-1011311023121103-2132010230031322-3303103120011332-0023101001111212-3022011112220201): complete subsection reference.

<a id="canonical-2210331211223132-0312301132003103-0303110132032230-1112111132031110-2010303033231122-2133112222210310-1122003112330123-2002213031131232"></a>

<a id="canonical-0311113220323332-1010322201003233-2231210302111010-2120013311023012-3003121330200202-1021120333210332-3033202202100300-1323110312103211"></a>

#### `routes.redirect_route.route_redirect.replace_params` property

Type: `"string"`. Computed.

Exclusive with \[remove\_all\_params retain\_all\_params\].

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0122001330032003-1330213221121203-0031122330312012-0032222033001002-1100000302212030-3120122033110002-2231120133321202-1020222111213001"></a>

<a id="canonical-3331201301123331-0321001131132021-1213031232333202-1031310223130303-2133022032222223-1302322310031221-1211033121211012-3322021331102223"></a>

#### `routes.redirect_route.route_redirect.response_code` property

Type: `"number"`. Computed.

The HTTP status code to use in the redirect response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

- [retain_all_params](data-sources--http_loadbalancer--reference--group-025.md#canonical-1232030113031202-1211313001200222-1000120202221323-0300112022323122-2112110333301113-1101030233221021-2033201202203202-1323030332200230): complete subsection reference.

<a id="canonical-2033023032323312-2011123302223003-2322013001200302-1011311023121103-2132010230031322-3303103120011332-0023101001111212-3022011112220201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.redirect_route.route_redirect.remove_all_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-024.md#canonical-3031130012330222-2232013233232232-2002103331033200-2211001300312033-3312120221121332-1323201212001023-1020112222113103-2303212300333332)
- [routes.redirect_route.route_redirect](data-sources--http_loadbalancer--reference--group-024.md#canonical-3110312013020332-0323110212313021-3313003231012220-0312332112101122-0122103201031233-3003001222133102-1311200122000230-3110313102321223)
- routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-1231212100000110-3122310002321202-3003233131113003-1021210123320110-2230123321002131-2231211303120302-2213333331101301-3113232313012001"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for remove all params.

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
