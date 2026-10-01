---
page_title: "xcsh_udp_loadbalancer reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_udp_loadbalancer reference."
---

# xcsh_udp_loadbalancer reference

<a id="canonical-0322213321213202-3202110231110133-2321001131201301-2332103312001133-1123013211012223-2302321130233000-0320333322330213-0013310000232022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223301323303222-3121130130211031-0312211131131332-1302213121230321-3000230222313032-0233322300323013-1123300233031331-0323223200320311"></a>

## origin_pools_weights.pool — pool / 112013321001 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-0032213313100301-3212011112101330-1011333321321000-2333033310333031-2321102233211130-2001102232222211-1021030023313221-2023031320303303)
- origin_pools_weights.pool

<a id="canonical-3212031321111032-1323010231033032-1030221012003033-3030032003312110-2132103312113033-0210320112321133-0311102213230112-2323003313233003"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
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
pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132110122203231-1330202011112210-0231110030303322-1010013232023001-0131130033132300-0301102123030202-2132131010120010-0103100201000332"></a>

## Direct properties — pool / 112013321001 / 3

<a id="canonical-0102123012121013-1312311300310222-2201101212302230-2000012233220312-1012021020230223-3223202121020223-3311103200021203-1123222000232000"></a>

<a id="canonical-1101132312202312-1200311310002212-2313100122031200-1321002230133332-1010202001211010-0021211332120302-1101013231131023-1320230023112102"></a>

## name property — pool / 112013321001 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2213010011232331-1032132002333120-0312211201103020-2221010323111100-0310112200102201-2123202201310021-1001113131020112-3212011213032030"></a>

<a id="canonical-2003131113201313-3211201033131303-3100030002333301-1010000023311011-3232001131012031-0001333223323032-2300221021200022-1030200201011121"></a>

## namespace property — pool / 112013321001 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3210122332303233-0022311202213122-0313012133221200-0120133023020231-1130021230031130-0323301233303120-2111230221030302-0010311202132321"></a>

<a id="canonical-3311203010222200-0011022013121203-2232022201220311-3303222111003023-0212230122320102-3211302223132131-2231011202112011-1103101011123033"></a>

## tenant property — pool / 112013321001 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3302223202031230-3122322231113221-2202011231312102-3001002303133211-0322130332203002-2310100011032122-3323102330120013-0002122230123213"></a>

## Next pages — pool / 112013321001 / 7

- [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-0032213313100301-3212011112101330-1011333321321000-2333033310333031-2321102233211130-2001102232222211-1021030023313221-2023031320303303)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-2221100322323133-3222100232310313-2133010130310031-3231202023101032-3313112200110100-2130133221200013-2211012323333211-0003213013303230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130220123320013-0222320112010112-0201012313330102-3332203320220211-2333002312312031-2320203012010210-1121022233013333-1321001033111130"></a>

## service_policies_from_namespace — service_policies_from_namespace / 002332311131 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- service_policies_from_namespace

<a id="canonical-2233110233201220-2203330010011033-2213301013122032-1020103020330201-1321121230130132-1133300233332010-3021001121303331-1211000301132011"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
service_policies_from_namespace = {}
```

<a id="canonical-3333032011100220-2322323321102132-3030101302323110-1323101131021102-1021131001013032-3323100323302102-3120111132203131-2112022220233323"></a>

## Direct properties — service_policies_from_namespace / 002332311131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200123332213313-2220330031002133-3033130103002233-1122023320300323-0121021000112123-0331123233000030-2023003110130030-2301320303230133"></a>

## Next pages — service_policies_from_namespace / 002332311131 / 4

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-1222013031123323-0132302211031313-1113201031020033-3330303023103321-1132201210330120-3022032110001231-2010201311132220-1032103313302111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333002123312010-1323332102202323-0032302131121210-1102300032033110-0301313322032032-0133221101223231-3201201301022312-1122112221231110"></a>

## timeouts — timeouts / 110313003121 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- timeouts

<a id="canonical-2021022131222223-2333212001001212-2130332200003131-0020302000212332-1330323131131320-0311001103200222-0222203302031213-0330330211100201"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122200211300013-2322002330113132-3103112010332021-1221110232303232-2012333231313032-2103130002011221-0310320100110200-0321322123031331"></a>

## Direct properties — timeouts / 110313003121 / 3

<a id="canonical-3000102312123121-2222013031010322-3233121032202233-1110011132322112-3213302012103102-3230033012213102-1332113000131022-3112200030011201"></a>

<a id="canonical-3030202331201302-3110230022011220-3131113033122202-0110220130313113-0201200313102231-3332023313302200-3121101112333032-0101023101311233"></a>

## create property — timeouts / 110313003121 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2222311100010023-0101330313131303-1220221021112101-3222320033333210-0323133022321001-2221223332210312-3302101123231000-3320321103010213"></a>

<a id="canonical-1101121333000123-2330213201010010-3132310122011232-0231330111220331-1212012032033110-2001022020300233-3031130203331033-3103200231303031"></a>

## delete property — timeouts / 110313003121 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2200002032133122-1131032330110203-0012313220212113-1301020202121323-2232030303201001-0210110202200033-2232003300231203-3221323230010321"></a>

<a id="canonical-0132123232010212-1033223013221313-3323230033200110-3120320330033110-0130320320210222-3222001131032132-2001312200002010-0303121120330002"></a>

## read property — timeouts / 110313003121 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1301133113131132-0323331102033003-3132210323110023-0303000122122233-3221031302110132-2210132033112112-2220301102231110-3310221213103310"></a>

<a id="canonical-1331330032333032-1023230220323023-1333211331330021-3101123101012103-2012100220212220-3230021030223302-2322013330012121-1222221300020013"></a>

## update property — timeouts / 110313003121 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3311220123302300-0211203310323222-0020322223221113-0110100302233200-3210233020220100-1231003302302213-2100133002003001-0100030211120011"></a>

## Next pages — timeouts / 110313003121 / 8

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-3121113233300112-1011301020012201-3133332123031100-1313012230000123-2323010011202022-1122033233122001-3001313032003221-2033202233101030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100301311132121-0230202023130021-1030102222023221-3132123001123321-3120210213102320-1312330012313323-1103010032213013-0321232113312223"></a>

## udp — udp / 100031132011 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- udp

<a id="canonical-1213213303011020-2212100322320001-3012203133023012-0202200110022330-1033233132020222-1011021100233133-2203233223110313-3323232000021300"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
udp {}
```

<a id="canonical-2211111302000200-2000330202123311-0332300200232131-1131230011013300-3030101220230010-1112010113301012-3132130032331133-2000102322213003"></a>

## Direct properties — udp / 100031132011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333002330232331-1120120201013000-2031300012102102-3012112323313312-1001313303200301-3310101031320322-0213222222013003-0220233311130003"></a>

## Next pages — udp / 100031132011 / 4

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
