---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-2121320001122022-1000211022330010-1232131322320333-2032023230121302-1232233311131033-1130231222131221-0203023303230233-3220013022320131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-013.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-013.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- rate_limit.rate_limiter.action_block

<a id="canonical-3031132200211133-1023311111200312-2101201003113000-3223033010030221-0122300321113023-0100301000300220-3001220320112013-0021231001133011"></a>

Type: `"object"`. single nested block, Optional.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("hours",
    "minutes"),
  validators.ConflictingObjectAttributes("hours",
    "seconds"),
  validators.ConflictingObjectAttributes("minutes",
    "seconds")}
```

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

Terraform syntax:

```terraform
action_block {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022323122112022-1133131113332201-3221032210123200-1112100212210330-1300231201110121-2013322011212312-3211330122130010-0100120001310230"></a>

### Direct properties for `rate_limit.rate_limiter.action_block`

- [hours](resources--cdn_loadbalancer--reference--group-014.md#canonical-0001220131200010-1033233213023323-3333300333212200-2133013021233210-0301302322310330-3130330031202120-2310111303100313-0313332031331111): complete subsection reference.

- [minutes](resources--cdn_loadbalancer--reference--group-014.md#canonical-2232102003211200-3310003313003312-0012200212113100-0002323313331101-1020102203000121-1302032211020003-3000302231323112-1301000013201333): complete subsection reference.

- [seconds](resources--cdn_loadbalancer--reference--group-014.md#canonical-1320212332333121-2121103102302021-2000231022122212-1221020333221331-0120202330203000-1000300113331220-0211232131110222-2311323201323002): complete subsection reference.

<a id="canonical-0001220131200010-1033233213023323-3333300333212200-2133013021233210-0301302322310330-3130330031202120-2310111303100313-0313332031331111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.hours` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-013.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-013.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-2121320001122022-1000211022330010-1232131322320333-2032023230121302-1232233311131033-1130231222131221-0203023303230233-3220013022320131)
- rate_limit.rate_limiter.action_block.hours

<a id="canonical-1020022103303203-0100213122112331-3022212321120202-0112003001222223-3033200113332023-3131300313013123-0003221133123311-2223123221123333"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330103111232101-3132233001303331-2222303213302230-3133203101103011-0002013223121020-1200133123211331-1313111122221333-3111101303113100"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.hours`

<a id="canonical-3132212132133200-1032200323323022-1102313200331323-2333002220222202-2111313231323230-0313212020210323-0233100320002111-3330002311211120"></a>

#### `rate_limit.rate_limiter.action_block.hours.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 48),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2232102003211200-3310003313003312-0012200212113100-0002323313331101-1020102203000121-1302032211020003-3000302231323112-1301000013201333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.minutes` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-013.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-013.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-2121320001122022-1000211022330010-1232131322320333-2032023230121302-1232233311131033-1130231222131221-0203023303230233-3220013022320131)
- rate_limit.rate_limiter.action_block.minutes

<a id="canonical-1232303121032132-2302122103131301-0332200103310110-0121121331333213-0110203030123330-2021310203021323-3321031321010313-3202013302022010"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320001103010013-2123312000020332-0221013030031313-0112133130111222-1221212333123221-3101310031123010-2031220232203020-1321012131321323"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.minutes`

<a id="canonical-0011330110031122-2111300321230301-2010003300013220-2321013203003023-1233200032301030-0110003120000232-0233220112012303-0023133110231233"></a>

#### `rate_limit.rate_limiter.action_block.minutes.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 60),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1320212332333121-2121103102302021-2000231022122212-1221020333221331-0120202330203000-1000300113331220-0211232131110222-2311323201323002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.action_block.seconds` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-013.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-013.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-2121320001122022-1000211022330010-1232131322320333-2032023230121302-1232233311131033-1130231222131221-0203023303230233-3220013022320131)
- rate_limit.rate_limiter.action_block.seconds

<a id="canonical-0011331330000033-2011200332122022-3130122233122122-3102131133231221-1011200333001022-0121000030102123-1031021113230122-3122122231021333"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
seconds {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112002010011011-0031211333010321-0032300003201331-1200001222200130-2213330022332331-0013212103113003-2231213200332320-2312220113312100"></a>

### Direct properties for `rate_limit.rate_limiter.action_block.seconds`

<a id="canonical-1232220201100310-3222033101201002-1113033122230001-3023003203210332-3001000320210310-1112102100100001-1012323033202123-0323303120223132"></a>

#### `rate_limit.rate_limiter.action_block.seconds.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 300),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0121113003211121-3000233231212222-0210322310333131-1232221120110102-0323331121001101-3300320021101330-3211100221101210-3303200202212010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.disabled` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-013.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-013.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- rate_limit.rate_limiter.disabled

<a id="canonical-1121312100131031-2223013120000303-3012122112011231-3000022132131122-3323321221113132-0212210312123320-0211111221020201-3202230211322330"></a>

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
disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121301302012332-0200200003121032-3002102313230233-0221300213312113-3323111113013323-2321301313301030-1201012223223023-2123010120320301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.leaky_bucket` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-013.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-013.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- rate_limit.rate_limiter.leaky_bucket

<a id="canonical-1001223310211010-2211301333033321-0122300320033221-2301010110311101-3010003232121100-0232213013303010-2200103300012030-1211133312003023"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
leaky_bucket = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211020022010202-1101201222211102-3010213311221001-2102010330320122-0120301001332101-0332320102011122-0321200121122313-3032003110020000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter.token_bucket` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [rate_limit](resources--cdn_loadbalancer--reference--group-013.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-013.md#canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211)
- rate_limit.rate_limiter.token_bucket

<a id="canonical-2012231200012321-2301010321102022-1122211011332002-1231202222131112-3130202030223103-0021133201020233-1230312222321023-2132012312022122"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
token_bucket = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101301311121220-1022212312223231-0103230132010000-1031303112000210-1003201331112110-2210011211032323-3113203213013300-1111103022320301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sensitive_data_policy` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- sensitive_data_policy

<a id="canonical-1300020121333330-1303221303112300-0023322213301102-1002303223322122-2130031202131001-1332232121231102-1300111112133313-2022220010223230"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
sensitive_data_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032323113320110-2313230222132212-3120030220032000-1220220010022010-2303000220211222-1112102222102130-1312103200122002-3031000100322233"></a>

### Direct properties for `sensitive_data_policy`

- [sensitive_data_policy_ref](resources--cdn_loadbalancer--reference--group-014.md#canonical-0201002102200313-1002330132131011-3132132012332313-3300220302033000-0332032322333323-1213010011231132-0030331210203000-3231223310010012): complete subsection reference.

<a id="canonical-0201002102200313-1002330132131011-3132132012332313-3300220302033000-0332032322333323-1213010011231132-0030331210203000-3231223310010012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sensitive_data_policy.sensitive_data_policy_ref` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [sensitive_data_policy](resources--cdn_loadbalancer--reference--group-014.md#canonical-3101301311121220-1022212312223231-0103230132010000-1031303112000210-1003201331112110-2210011211032323-3113203213013300-1111103022320301)
- sensitive_data_policy.sensitive_data_policy_ref

<a id="canonical-3210100202211022-1332030101023121-0300010231230001-1202122213333001-1232033103122211-3303322310121230-2112033112212102-1132200333031200"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
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
sensitive_data_policy_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122030223133001-2022103302033130-1233230121002121-0030020222322220-3112301320232110-1130111000232021-2303202323213330-0020121231121002"></a>

### Direct properties for `sensitive_data_policy.sensitive_data_policy_ref`

<a id="canonical-3230300102010002-0201110113031320-0022020212113030-0121221121322003-2133231213332121-1001113233222233-2301023100323131-2001011011201110"></a>

#### `sensitive_data_policy.sensitive_data_policy_ref.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-0212131002020121-1011132203320132-1000212330230121-1310111023233102-3111303021131202-3122212022313200-1021101222032031-0013210320032301"></a>

<a id="canonical-2112000210032202-1310001212303322-1233133002313133-1330013133112110-1022323100030300-0231120203033323-2112012022213013-0011101101301200"></a>

#### `sensitive_data_policy.sensitive_data_policy_ref.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-0330322312332322-0133321130012122-3213100012313221-3212213022003220-0223110100300333-3121202122202330-0003030013300301-0332102323012000"></a>

<a id="canonical-3200032001301210-3232023103003122-3211020120313101-2022131301013320-1333301011222013-1021320300011112-1302222221303232-0010313030321021"></a>

#### `sensitive_data_policy.sensitive_data_policy_ref.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-0232201322200031-0330122100331311-3221203302302022-3021301122011232-3230010232020130-3332030103113102-0223113230122010-1011332330021133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service_policies_from_namespace` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- service_policies_from_namespace

<a id="canonical-0212231232132301-3212113303020322-3222132120230030-1123313213201131-1130023023230203-2203303210021230-2103211113213120-2321022020101213"></a>

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
service_policies_from_namespace = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233003020202211-1121000220031121-0002002321202011-3200220102103002-3312211001303013-1111001323333210-3232211013001110-3302023302323112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slow_ddos_mitigation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- slow_ddos_mitigation

<a id="canonical-0112221300332233-1321331031100323-2021100310011011-3232120031232301-3333233030001223-2021312201110121-1201200310231320-3001221030310130"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: slow\_ddos\_mitigation, system\_default\_timeouts; Default: system\_default\_timeouts\]
'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Additional upstream details:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("request_headers_timeout"),
  validators.ConflictingObjectAttributes("disable_request_timeout",
    "request_timeout")}
```

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

- [slow_ddos_mitigation](resources--cdn_loadbalancer--reference--group-014.md#canonical-0112221300332233-1321331031100323-2021100310011011-3232120031232301-3333233030001223-2021312201110121-1201200310231320-3001221030310130)
- [system_default_timeouts](resources--cdn_loadbalancer--reference--group-014.md#canonical-2003220332131013-2211333310231313-0030002201303021-1102200303220023-3011112311203322-2110130310322231-2112032032222113-2133222202031012)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
slow_ddos_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122132022120031-2021300013310121-2103211111333321-2220222032031101-0111213320011131-2233002323312120-0003001310021110-1333332333001333"></a>

### Direct properties for `slow_ddos_mitigation`

- [disable_request_timeout](resources--cdn_loadbalancer--reference--group-014.md#canonical-3202203332302003-2003031212203232-1211332211112331-3321100223302000-0123021002310131-1231322130131002-0102223001003233-0223223333102113): complete subsection reference.

<a id="canonical-2322123032100230-1212230310211102-2033220310120103-2122100002033212-1023221333000200-3002013000301011-0322121113232330-1002120021000122"></a>

<a id="canonical-3323330331003121-0203313020131132-3233322231302213-0201331032132213-3221122131222332-2121321311303020-2023131230331300-2113332120120130"></a>

#### `slow_ddos_mitigation.request_headers_timeout` property

Type: `"number"`. Optional.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Additional upstream details:

The default value is 10000 milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(2000, 30000),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3321022001022222-2212122111020130-3210110301332333-2103103221230003-3320210211002133-3202311210030301-0033010210211313-2003030101212202"></a>

<a id="canonical-0122201220311230-3002002211223033-3002322121032000-0022130303323003-1203120213220113-1133323121032323-3110210312012100-2321222212023302"></a>

#### `slow_ddos_mitigation.request_timeout` property

Type: `"number"`. Optional.

Exclusive with \[disable\_request\_timeout\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(2000, 300000),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3202203332302003-2003031212203232-1211332211112331-3321100223302000-0123021002310131-1231322130131002-0102223001003233-0223223333102113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slow_ddos_mitigation.disable_request_timeout` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [slow_ddos_mitigation](resources--cdn_loadbalancer--reference--group-014.md#canonical-1233003020202211-1121000220031121-0002002321202011-3200220102103002-3312211001303013-1111001323333210-3232211013001110-3302023302323112)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-3332032001123310-1001010123211132-1310032101023223-1023003323313332-0330032123100302-1023220133021202-2311123330010300-2112113320131312"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_request_timeout = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302302300220120-1232133331030022-2202232123331001-1001230212313333-2002032030033031-0300122120322121-0023222330210101-0203103301320000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `system_default_timeouts` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- system_default_timeouts

<a id="canonical-2003220332131013-2211333310231313-0030002201303021-1102200303220023-3011112311203322-2110130310322231-2112032032222113-2133222202031012"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
system_default_timeouts = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120320111222111-0221312232103321-2100002103330133-0121321020113111-3112222032011030-0332323220222220-0312012223333230-3320123022202133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- timeouts

<a id="canonical-2023222110021223-0210202233233101-1321110100223020-1120313011212123-0223220120201112-1013311101313021-3002233013330332-1022010010310301"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3033333132112201-0003122320323232-1321112013033002-1131120120022002-1231101321303211-2020022230131023-1313131031121320-1120032220201212"></a>

### Direct properties for `timeouts`

<a id="canonical-1202010300212021-0220220031031001-1110010200022332-0111212020123101-1310200302213232-0313130023130313-3221133313213320-1002323130220123"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3200032132201311-2302121230021222-0222123300131002-1021013022311310-2020223031212312-0220310230122023-1230230123323011-2002301233101322"></a>

<a id="canonical-0221222112021312-3120021032330023-0001223200202103-2311012020121133-1312202302030321-3333122232310013-1121002201232132-1130012212120111"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3233220020230300-2103131210212301-2313232001223301-2133333100111002-3032111213101303-2011130211103030-2013311320311320-3221112002301030"></a>

<a id="canonical-3231111223100301-0213013210231131-0001303323013220-1213103130120332-1011332103233202-3020133323012222-1220223320022131-2113213233131103"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0021130100120101-2302132113100113-0330220313110200-0100231332332113-1223103101230003-0310121233210103-2112221122132211-3230223131300210"></a>

<a id="canonical-2100200303120130-2200211101003321-3102312331001303-2031200233021132-1311322301231333-3312233220232102-1103003311123220-3332310000211330"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- trusted_clients

<a id="canonical-0331013110121213-1103203320113023-3133220211231201-2333202101220103-3100110012031010-3020111011320232-0113200202211221-1322232320320102"></a>

Type: `"object"`. list nested block, Optional.

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. For clients.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("actions"),
  validators.ConflictingListObjectAttributes("as_number",
    "http_header"),
  validators.ConflictingListObjectAttributes("as_number",
    "ip_prefix"),
  validators.ConflictingListObjectAttributes("as_number",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("as_number",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("bot_skip_processing",
    "skip_processing"),
  validators.ConflictingListObjectAttributes("bot_skip_processing",
    "waf_skip_processing"),
  validators.ConflictingListObjectAttributes("http_header",
    "ip_prefix"),
  validators.ConflictingListObjectAttributes("http_header",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("http_header",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("ip_prefix",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("ip_prefix",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("ipv6_prefix",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("skip_processing",
    "waf_skip_processing")}
```

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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
trusted_clients {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203321333110330-3120020012011011-2121301302103230-2010333311132133-1023331230130203-3100202323010002-3222002300130313-3303321232303223"></a>

### Direct properties for `trusted_clients`

<a id="canonical-3112200000232130-3003122100223203-0113201130223100-0300010022020000-0211011120021011-3132100123303120-1110123122012101-0322020210232312"></a>

#### `trusted_clients.actions` property

Type: `["list", "string"]`. Optional.

\[Enum:
SKIP\_PROCESSING\_WAF|SKIP\_PROCESSING\_BOT|SKIP\_PROCESSING\_MUM|SKIP\_PROCESSING\_IP\_REPUTATION|SKIP\_PROCESSING\_API\_PROTECTION|SKIP\_PROCESSING\_OAS\_VALIDATION|SKIP\_PROCESSING\_DDOS\_PROTECTION|SKIP\_PROCESSING\_THREAT\_MESH|SKIP\_PROCESSING\_MALWARE\_PROTECTION\]
Actions that should be taken when client identifier matches the rule. Possible values are
\`SKIP\_PROCESSING\_WAF\`, \`SKIP\_PROCESSING\_BOT\`, \`SKIP\_PROCESSING\_MUM\`,
\`SKIP\_PROCESSING\_IP\_REPUTATION\`, \`SKIP\_PROCESSING\_API\_PROTECTION\`,
\`SKIP\_PROCESSING\_OAS\_VALIDATION\`, \`SKIP\_PROCESSING\_DDOS\_PROTECTION\`,
\`SKIP\_PROCESSING\_THREAT\_MESH\`, \`SKIP\_PROCESSING\_MALWARE\_PROTECTION\`. Defaults to
\`SKIP\_PROCESSING\_WAF\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(10),
}
```

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

<a id="canonical-0212200203132213-3001331122003020-1100120311121011-2011223102122132-0302001023333313-2220000122122020-2000113213010333-0111101213032303"></a>

<a id="canonical-2000212210003103-1100310233303223-2110112002302232-3110212312032332-0200310002232003-2200201333332013-0003110131233322-0110031222233010"></a>

#### `trusted_clients.as_number` property

Type: `"number"`. Optional.

Exclusive with \[http\_header ip\_prefix IPv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 401308),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [bot_skip_processing](resources--cdn_loadbalancer--reference--group-014.md#canonical-2023122312000030-3003311220212320-2301130321111110-0030022112103323-1322310121101221-1022311201233033-2333330301223133-0212311012200001): complete subsection reference.

<a id="canonical-1013332032303321-1223021223132000-1230130322032131-0222321312223220-2231332113200131-2323203212300202-1222122120010023-0202012331320020"></a>

<a id="canonical-3320002023201121-3203021322233130-3223220321112020-1122233102332030-3110003313323030-1120110101002102-3222010030001033-1103323332333232"></a>

#### `trusted_clients.expiration_timestamp` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [http_header](resources--cdn_loadbalancer--reference--group-014.md#canonical-1201221302301222-1010320030333121-0113223131202312-2330113130111003-3012011333202310-0101333000203222-3111003100213100-3013310003133311): complete subsection reference.

<a id="canonical-3333000311032033-0131020213223221-0120323022130202-2131131301110212-2111122020132003-2013321321120123-0120320033112233-1312131221102111"></a>

<a id="canonical-2331111130030312-2012131022103332-3102211103001212-3112320302121031-0002030112322220-1003133322102221-1330202310212011-3131302320033221"></a>

#### `trusted_clients.ip_prefix` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-0300100301021032-3101203201231302-1132102303301201-1133111021223302-2113233133033310-0320112302023202-3103102013322100-2201100000231210"></a>

<a id="canonical-0313130033231100-2011223101202121-2001222023133301-1231302031010020-2212101310113230-3303301331233211-3302111133122203-3102101322002202"></a>

#### `trusted_clients.ipv6_prefix` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

- [metadata](resources--cdn_loadbalancer--reference--group-014.md#canonical-3321333031300023-0102112101330300-2231023022110222-0003023011013002-0222123013100321-2220230222132333-1231133330031312-2133323113312100): complete subsection reference.

- [skip_processing](resources--cdn_loadbalancer--reference--group-014.md#canonical-1133331223201300-3010010301001020-1022001210300133-3110022001203122-3310333101330031-3301103022232211-3311232311220231-2323022320130113): complete subsection reference.

<a id="canonical-2123202023131030-3010311112220303-3330112013322203-1323302030032030-1331230311012013-1310113121033130-0311010031033211-0030320210023021"></a>

<a id="canonical-0230323302113323-0102232133212033-0101201022210233-0301020301011012-0101201130333100-3310113222302130-0013233102011320-3003213220123312"></a>

#### `trusted_clients.user_identifier` property

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix IPv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [waf_skip_processing](resources--cdn_loadbalancer--reference--group-014.md#canonical-0033212232020333-0132213332131031-2012121203120303-1333000331032120-1202131120001010-1123301101201001-2130112022103101-1312012221031000): complete subsection reference.

<a id="canonical-2023122312000030-3003311220212320-2301130321111110-0030022112103323-1322310121101221-1022311201233033-2333330301223133-0212311012200001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.bot_skip_processing` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- trusted_clients.bot_skip_processing

<a id="canonical-1022012022211010-1210121312322011-1023230031203200-2222210221220032-0120131033201333-0010311121223113-0012130302221330-1011032000131121"></a>

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
bot_skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201221302301222-1010320030333121-0113223131202312-2330113130111003-3012011333202310-0101333000203222-3111003100213100-3013310003133311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.http_header` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- trusted_clients.http_header

<a id="canonical-0322123120203221-0100300302011021-1121311023232123-3320122012131233-0323003120103312-0110013312300132-0313133021231021-2312311132030230"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Additional upstream details:

Request header name and value pairs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("headers")}
```

Receipt-pinned upstream constraints:

```json
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
http_header {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300021013131333-2023322313210003-1233011311023113-2220121200213030-2103031013133012-3011132003211122-1311333313222023-2320330012100032"></a>

### Direct properties for `trusted_clients.http_header`

- [headers](resources--cdn_loadbalancer--reference--group-014.md#canonical-2103310013311322-2231301123301132-1201212221203032-3131133002313323-2232222322322030-2331020231133012-0310131022023013-2131021111020033): complete subsection reference.

<a id="canonical-2103310013311322-2231301123301132-1201212221203032-3131133002313323-2232222322322030-2331020231133012-0310131022023013-2131021111020033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.http_header.headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- [trusted_clients.http_header](resources--cdn_loadbalancer--reference--group-014.md#canonical-1201221302301222-1010320030333121-0113223131202312-2330113130111003-3012011333202310-0101333000203222-3111003100213100-3013310003133311)
- trusted_clients.http_header.headers

<a id="canonical-2000322121230030-0022220021121112-2201233121221032-0031032311023330-0033033013130303-0330202202333213-1013310233302102-2230001012123033"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP header name and value pairs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "presence"),
  validators.ConflictingListObjectAttributes("exact",
    "regex"),
  validators.ConflictingListObjectAttributes("presence",
    "regex")}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330223120330213-1122203211113231-0100102031010001-0133013012333310-2020322231312213-1222103000121321-3103102133210001-3100301011231121"></a>

### Direct properties for `trusted_clients.http_header.headers`

<a id="canonical-3220103332030110-1213001312123313-3210230312321221-1202012303032020-3302302201033210-3032323300023303-3303101030200212-3010332120301130"></a>

#### `trusted_clients.http_header.headers.exact` property

Type: `"string"`. Optional.

Exclusive with \[presence regular expression\] Header value to match exactly.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-3313220001133133-1121130121321131-0033032031232200-0110202121223032-1222321101310311-0132131300001123-0310233103113002-0331332101013300"></a>

<a id="canonical-0301330122201323-1201303330100013-2301302020231032-3222333331031222-3232100301023330-1031113113203310-1330030232223323-2023120212110233"></a>

#### `trusted_clients.http_header.headers.invert_match` property

Type: `"bool"`. Optional.

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

<a id="canonical-2221010211133323-0323211213102111-1131310122303112-0023000120031133-0013031001122220-1030212330230011-0030103121230000-2203213123210300"></a>

<a id="canonical-2102213231101232-2201311200303223-2002212323012321-2012201330113030-0113132321022021-2123001320302212-3310022232333002-2003112111323003"></a>

#### `trusted_clients.http_header.headers.name` property

Type: `"string"`. Optional.

Name. Name of the header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-2031033203233313-3312313022123220-1111001331103320-1133222210201333-3231300230221133-3213321310203120-1300200201233320-2303230020111013"></a>

<a id="canonical-3120232323200023-1001311230101200-1023101321230333-2032221231102220-2021033132131300-0002101200333201-0202031130111123-0021220031001021"></a>

#### `trusted_clients.http_header.headers.presence` property

Type: `"bool"`. Optional.

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

<a id="canonical-3010020331200331-1321310000200122-0032230201111120-0101223100111023-2330121200203002-3210023111202221-0233110232221030-1023113131103111"></a>

<a id="canonical-2121310023230010-3012120222032111-3330112323210011-0303220131211020-2231321322131003-0032213222203200-3133231001302113-1231220013130012"></a>

#### `trusted_clients.http_header.headers.regex` property

Type: `"string"`. Optional.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3321333031300023-0102112101330300-2231023022110222-0003023011013002-0222123013100321-2220230222132333-1231133330031312-2133323113312100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- trusted_clients.metadata

<a id="canonical-1000301111301000-1312320223100112-3200300211222232-2223203202013233-1130130031130303-0221103020230232-1310130022020320-0100132320130311"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2003323233110122-2101311312331100-0103033121323310-2101213231111233-3300022001220000-3331311133022013-1302231033003210-3100033133113003"></a>

### Direct properties for `trusted_clients.metadata`

<a id="canonical-0101221020300010-2332030003001321-1132211021032213-3223321111302102-0032303300130112-2032303103033213-2212031012313133-0101000330120211"></a>

#### `trusted_clients.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2212213230022003-3300202232232023-0320131000010103-1132002310223131-0132300101302310-3101202001201002-3321023103032322-2210122212202313"></a>

<a id="canonical-0300313112000203-0112301202311113-0123013103203203-2122032212120301-1301221010212331-0201003300033232-2022313213033033-2021111322211020"></a>

#### `trusted_clients.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-1133331223201300-3010010301001020-1022001210300133-3110022001203122-3310333101330031-3301103022232211-3311232311220231-2323022320130113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.skip_processing` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- trusted_clients.skip_processing

<a id="canonical-3000003300203230-1300311212133303-1102110023002131-2322002121220233-1000012202023332-1131322103221222-0001101303123131-0003210232023011"></a>

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
skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033212232020333-0132213332131031-2012121203120303-1333000331032120-1202131120001010-1123301101201001-2130112022103101-1312012221031000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.waf_skip_processing` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- trusted_clients.waf_skip_processing

<a id="canonical-1122312021132001-2000211210010321-2322101221220113-0023320221030130-2333001313120122-3212313230120213-3013121310120001-3003311332210232"></a>

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
waf_skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302200200302032-1033011220200321-0322330033221123-1310331022111310-3130333121221113-1313030021202201-1313311122312131-2320123301120231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_id_client_ip` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- user_id_client_ip

<a id="canonical-1312331103213230-3110310233011330-2210113221300211-1033233123013102-0122311213033031-2300023231023202-0020301210013322-3331103133021000"></a>

Type: `["object", {}]`. Optional.

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

- [user_id_client_ip](resources--cdn_loadbalancer--reference--group-014.md#canonical-1312331103213230-3110310233011330-2210113221300211-1033233123013102-0122311213033031-2300023231023202-0020301210013322-3331103133021000)
- [user_identification](resources--cdn_loadbalancer--reference--group-014.md#canonical-2112002310201312-1032232321300131-0300121020020022-3132103001332012-2011021103101103-2113302011323011-0303100300100221-2032311102302321)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
user_id_client_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112230101301222-3222000100021310-1113222132033220-1211130133022022-2010313030300030-0312321030200003-0211330333102113-1230112232300132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_identification` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- user_identification

<a id="canonical-2112002310201312-1032232321300131-0300121020020022-3132103001332012-2011021103101103-2113302011323011-0303100300100221-2032311102302321"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
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
user_identification {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210332013201120-2121101113030333-1130130122023333-3123203130102102-1303320130001321-2203001122022202-2333302032100112-1201021022010101"></a>

### Direct properties for `user_identification`

<a id="canonical-0222001330202213-1301303101100023-3010112331222021-1330310122022030-0313213112103330-0212220211132231-0000101100233230-3122222223313303"></a>

#### `user_identification.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-1013111212202101-3003101031301112-1102201131121302-2320023133132131-2310210303312031-2120230303002211-1122011133012300-0213012312013113"></a>

<a id="canonical-3203322031020100-3023002333303210-2333130303122300-2332112211223101-3311320310213132-3333110222210030-0130201230000022-2201122001330303"></a>

#### `user_identification.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3021200331031030-2111010322222101-1211310300301001-1001001320233213-0201110033122131-1113333003303012-2320113132130302-2020230313330101"></a>

<a id="canonical-3320112000011200-0221112323220003-3033320221310013-1331133020203222-3300120122332003-2212233233203210-3331231111130333-3332012201110132"></a>

#### `user_identification.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- waf_exclusion

<a id="canonical-3133121121103332-1133013020232202-3130021232011203-2031210311001023-1203011213110230-3113302211313322-0321121230102331-1012313300100101"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for waf exclusion.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("waf_exclusion_inline_rules",
    "waf_exclusion_policy")}
```

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

Terraform syntax:

```terraform
waf_exclusion {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303321020120023-1201021311220032-1233102031311121-2222330321210323-1210303321303122-3112200213110312-1000123303200131-3222113101231230"></a>

### Direct properties for `waf_exclusion`

- [waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201): complete subsection reference.

- [waf_exclusion_policy](resources--cdn_loadbalancer--reference--group-015.md#canonical-1100230020012103-3113212312312201-0333111112211303-3110022123123322-2311123203013303-0310013020333123-1123010130230011-0021210221300211): complete subsection reference.

<a id="canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- waf_exclusion.waf_exclusion_inline_rules

<a id="canonical-1212001002312000-0012133303123123-2000232211122311-3002102033112112-2110022102222313-3310113010203102-2011310020212000-1223223023100130"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
waf_exclusion_inline_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0333122002020202-3113122321010023-0222201020100113-2320322102223011-0101202030011003-2222300130101000-0232021303113120-1301133210210032"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules`

- [rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213): complete subsection reference.

<a id="canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- waf_exclusion.waf_exclusion_inline_rules.rules

<a id="canonical-2003101333112322-3001223301200130-0032011223220110-0333103230321002-0220230210300333-3033300003002022-2313331312323233-0332030333122232"></a>

Type: `"object"`. list nested block, Optional.

An ordered list of WAF Exclusions specific to this Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "exact_value"),
  validators.ConflictingListObjectAttributes("any_domain",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_prefix"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_regex"),
  validators.ConflictingListObjectAttributes("app_firewall_detection_control",
    "waf_skip_processing"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("path_prefix",
    "path_regex")}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302322223323000-0302103222001023-3112101210232113-0211033322132301-1303320223121321-1220202302003202-2023200300330232-2233311133231120"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules`

- [any_domain](resources--cdn_loadbalancer--reference--group-014.md#canonical-3022320322112022-3333012112032330-3300003313032211-2323211231200303-2311023110331323-0032303033101131-3112110001223120-0303310010000223): complete subsection reference.

- [any_path](resources--cdn_loadbalancer--reference--group-014.md#canonical-0001231022023102-0110000230133130-0122331322013010-2103313031223130-3221102322032223-1230200321000030-2121233132312331-1321021200213020): complete subsection reference.

- [app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-014.md#canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021): complete subsection reference.

<a id="canonical-2323233211011212-3232322132321211-0130002201021032-3232001033020210-2310001021011303-1230113120021003-0112033033032313-2010230000203030"></a>

<a id="canonical-2213100120213233-0202313102031132-3313103313010312-3231010131123322-1300311113001211-0010331233322102-3320111313011011-2112331032012202"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-0002210212203322-0113010033300122-2212111101011130-2331031200321000-2033210233130110-0333121323322023-3300131013103122-1132031310300101"></a>

<a id="canonical-3103323002112333-3322021003220202-0321110003300003-1131201100322303-2130323113202023-1310212103013230-2313311021112233-1220131100220230"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.expiration_timestamp` property

Type: `"string"`. Optional.

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

- [metadata](resources--cdn_loadbalancer--reference--group-015.md#canonical-1100320320310223-0023211210033123-1113013301230002-3221300313013010-2113301212000110-0030310031330102-0213210202322323-3012112020212112): complete subsection reference.

<a id="canonical-0332332333223133-1223001321022120-3021022310330113-0222230123223311-2232131122000000-1232222120122112-2101333002330123-1130211002100220"></a>

<a id="canonical-3231012322330303-1020301310011011-1031201110202101-3013322213202202-0213302232103112-0303312213003013-3111311100302301-2232120230332201"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3011333123211130-1321133220120002-2011220000221302-0222021103301213-0221111331323232-3003033120331121-3313110113002310-2030203033022313"></a>

<a id="canonical-2230312302000031-2103330322221023-3312003332311303-0022120330322203-2333032303002221-2202232202012323-2033222121301221-1310100322032021"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.path_prefix` property

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0331312212031120-0211003013303102-1123021123013113-2132030210230003-3111002333032230-3122100012121111-2121132113321333-1300130211022023"></a>

<a id="canonical-3312103301330032-1003022031022111-1310031022031231-1000201231010123-0231003102311032-3122113222121201-3023332120222023-3032211321120310"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.path_regex` property

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\] Define the regular expression for the path. For example, the regular expression
^/.\*$ will match on all paths.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0110233022303222-1300302212001333-3321220222203302-2232202201020030-0131301020031020-1110330322011101-3020102011113323-3010322221130022"></a>

<a id="canonical-1201031311302131-1211011132331030-0220000120202030-2201322001303113-3333323220110123-1310113131032212-3302101332332030-0321103330123122"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

- [waf_skip_processing](resources--cdn_loadbalancer--reference--group-015.md#canonical-1112133013100123-1033321203321311-1212031110232013-2131012113311222-2100101110213002-0010002030211101-3313022003322302-1102030212210021): complete subsection reference.

<a id="canonical-3022320322112022-3333012112032330-3300003313032211-2323211231200303-2311023110331323-0032303033101131-3112110001223120-0303310010000223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_domain

<a id="canonical-2301113101301021-3123011031200031-2102123013001202-1333320301230313-0012200220031222-0320033033230301-2000112120133211-3023100100032223"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0001231022023102-0110000230133130-0122331322013010-2103313031223130-3221102322032223-1230200321000030-2121233132312331-1321021200213020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.any_path` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_path

<a id="canonical-0102132222033230-2023322100221231-3001331113212133-3222123031120231-0001131310310221-2032121001020010-1302322002302032-1320131131332322"></a>

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
any_path = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control

<a id="canonical-2020012320113230-2201321002210211-0102332322323000-1312222330132010-2123100212213122-0102012203133221-3012131021031130-2232200220300213"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
app_firewall_detection_control {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303321323321331-3122320002312322-0200103010233232-1302232222333300-1333133312032012-2320313133002211-2203313313320031-3020020000112022"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control`

- [exclude_attack_type_contexts](resources--cdn_loadbalancer--reference--group-014.md#canonical-3333111003113030-0310202220110102-3323000113120032-1120203111100230-2012120032100132-0010001302123233-2312133102313210-3000120031313022): complete subsection reference.

- [exclude_bot_name_contexts](resources--cdn_loadbalancer--reference--group-014.md#canonical-2132132020113210-2232202020110030-2223103032021030-1030133123030210-1100303210200002-1303020232023212-3332002212313210-1023120212221121): complete subsection reference.

- [exclude_signature_contexts](resources--cdn_loadbalancer--reference--group-014.md#canonical-2320222201123010-3220010212313021-0233121121320010-2023232310301102-3032322211311032-1002233201331200-1112001223000030-0022333123133120): complete subsection reference.

- [exclude_violation_contexts](resources--cdn_loadbalancer--reference--group-014.md#canonical-0323332330322210-0121333001000321-2210331102321302-2210302333313133-3302001200121103-3310033132222031-2223210102030211-1022223122131233): complete subsection reference.

<a id="canonical-3333111003113030-0310202220110102-3323000113120032-1120203111100230-2012120032100132-0010001302123233-2312133102313210-3000120031313022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-014.md#canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-1133001132012132-1021021010201032-1130101022303021-3133210221303231-0032130330000011-2330310231200313-2103121322212230-1131221002131222"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
exclude_attack_type_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103010030331130-3232112201303321-2130101223110101-0111210013303010-3120111123112200-3231330223221202-2001312113022300-3210302331032020"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts`

<a id="canonical-0313110220123233-3102022222221220-3201123131102112-0321100232213213-1211301321230302-3032132330300210-1010201232021013-3311223100333121"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts.context` property

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CONTEXT_ANY","CONTEXT_BODY","CONTEXT_COOKIE","CONTEXT_HEADER","CONTEXT_PARAMETER","CONTEXT_REQUEST","CONTEXT_RESPONSE","CONTEXT_URI","CONTEXT_URL"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

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

<a id="canonical-2302012033213210-1000322110001233-0212023001333010-3310021220302202-1313021021201110-2000023333332223-0313213001110313-2101302231203330"></a>

<a id="canonical-3222311011012300-0023012013132131-0221213333312210-2132211033032210-2311033210113133-1032302021130220-2023023012020033-1100202110311201"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts.context_name` property

Type: `"string"`. Optional.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2012000002313331-3313012223100302-1003011311031321-3003323032032300-3232030212220332-2110023022123320-3133311213001112-1031001122223301"></a>

<a id="canonical-0331222122020000-3230203202313320-3101120313002002-3101233212111131-1112011112100213-0122220201122122-0131231203221032-2131331210131031"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` property

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY","ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS","ATTACK_TYPE_BUFFER_OVERFLOW","ATTACK_TYPE_COMMAND_EXECUTION","ATTACK_TYPE_CROSS_SITE_SCRIPTING","ATTACK_TYPE_DENIAL_OF_SERVICE","ATTACK_TYPE_DETECTION_EVASION","ATTACK_TYPE_DIRECTORY_INDEXING","ATTACK_TYPE_FORCEFUL_BROWSING","ATTACK_TYPE_GRAPHQL_PARSER_ATTACK","ATTACK_TYPE_HTTP_PARSER_ATTACK","ATTACK_TYPE_HTTP_RESPONSE_SPLITTING","ATTACK_TYPE_INFORMATION_LEAKAGE","ATTACK_TYPE_LDAP_INJECTION","ATTACK_TYPE_MALICIOUS_FILE_UPLOAD","ATTACK_TYPE_NONE","ATTACK_TYPE_NON_BROWSER_CLIENT","ATTACK_TYPE_OTHER_APPLICATION_ATTACKS","ATTACK_TYPE_PATH_TRAVERSAL","ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION","ATTACK_TYPE_REMOTE_FILE_INCLUDE","ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION","ATTACK_TYPE_SESSION_HIJACKING","ATTACK_TYPE_SQL_INJECTION","ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE","ATTACK_TYPE_VULNERABILITY_SCAN","ATTACK_TYPE_XPATH_INJECTION"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ATTACK_TYPE_NONE",
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
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"),
}
```

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

<a id="canonical-2132132020113210-2232202020110030-2223103032021030-1030133123030210-1100303210200002-1303020232023212-3332002212313210-1023120212221121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-014.md#canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-3312233330320232-3310011332123212-3023221120200112-3323310322303223-1233100323102003-1203032120200132-3310221111130222-2201223301130223"></a>

Type: `"object"`. list nested block, Optional.

Bot Names to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("bot_name")}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
exclude_bot_name_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221023003203323-3311003221121202-3020331123003113-1100013123202111-2332301120002300-0020320033130212-3302001212200021-1022111220112310"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts`

<a id="canonical-0010222300330313-3212121223020132-3211213033123133-3210221121232001-2313120303112113-1022331101000201-3201212012003230-0323233003102000"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2320222201123010-3220010212313021-0233121121320010-2023232310301102-3032322211311032-1002233201331200-1112001223000030-0022333123133120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-014.md#canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-2233010310221202-3221123301320321-2213120020213031-3230032022103200-1110301030231333-3113021130030312-3120012123032111-1212131232011112"></a>

Type: `"object"`. list nested block, Optional.

Signature IDs to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("signature_id")}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
exclude_signature_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122220033121212-0302132231222302-0112230223322333-0220000122331300-3312203333010300-3120000011232202-1003312130231021-0310331221221300"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts`

<a id="canonical-3313030202332321-2332231320132212-0332303102023311-0302021333133322-1113002213120030-1001200330202330-2123021021233212-2311332032032001"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts.context` property

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CONTEXT_ANY","CONTEXT_BODY","CONTEXT_COOKIE","CONTEXT_HEADER","CONTEXT_PARAMETER","CONTEXT_REQUEST","CONTEXT_RESPONSE","CONTEXT_URI","CONTEXT_URL"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

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

<a id="canonical-0232002232021233-3300010232131103-0113223212331131-1113220113330332-1020212022333100-0211320120113023-3230101323222110-1230302022032113"></a>

<a id="canonical-3130230012131030-2023111111131222-0211100130132123-2301312011011220-2122003223022312-0111120301132333-2101113121112110-1122311210312131"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts.context_name` property

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0313310131002133-2201322213200330-1100102333233312-0000132231102031-2221033301313010-2012023302332121-1031130132203032-0033113101031322"></a>

<a id="canonical-0003021110020203-1212202103323233-1032220331013320-2030200022201233-0200112312302223-2002100012023213-1233331310231202-2330210000000021"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts.signature_id` property

Type: `"number"`. Optional.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 299999999),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0323332330322210-0121333001000321-2210331102321302-2210302333313133-3302001200121103-3310033132222031-2223210102030211-1022223122131233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-014.md#canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-1023100121212323-2112111211310031-2113300003133132-3231022310010220-2303302313203302-1321120301100333-0131222111131223-0123132133013121"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
exclude_violation_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131000202301312-1220332230030321-1301103030310132-1032013230230202-2312021202331221-3020212301310112-2200211111301100-2322201330201211"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts`

<a id="canonical-0110331312333112-3313312333321200-3030120233022232-1212121010011212-0121101303110132-3200330032112100-2201301232213310-0101322030323131"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts.context` property

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CONTEXT_ANY","CONTEXT_BODY","CONTEXT_COOKIE","CONTEXT_HEADER","CONTEXT_PARAMETER","CONTEXT_REQUEST","CONTEXT_RESPONSE","CONTEXT_URI","CONTEXT_URL"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

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

<a id="canonical-3002301230110231-1021302012112210-2113012301302110-3013020013000111-3112011101303010-0122330102103231-2211220311130303-0221120220101202"></a>

<a id="canonical-0211202233023122-1232311302210121-0121200303033201-3222323203322210-3223110002101021-0323013012200321-2233021213133231-1123223110202120"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts.context_name` property

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1101223313211322-1103122123301102-3001311200232013-0002323102032003-1003132102223203-3001220322231022-3033123022323222-1012032321302022"></a>
