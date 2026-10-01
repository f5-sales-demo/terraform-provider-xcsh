---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1011232301111322-2230220332010200-1132221312322332-2112233101022203-0020133331222133-2313001000302031-3333131002222323-3210202213212113"></a>

## user_identifier property — blocked_clients / 333103331133 / 9

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix IPv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix IPv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [waf_skip_processing](resources--http_loadbalancer--reference--group-011.md#canonical-3101023120211000-3213013112130202-0320210320110301-0232020131122102-0122130233002131-3213203302322003-2100201130323002-0231223230232131): complete subsection reference.

<a id="canonical-1003132022111223-2200200300012201-2232123121103101-1013311022231033-1203021131312003-2033220012102210-0113332112323202-2023121311212202"></a>

## Next pages — blocked_clients / 333103331133 / 10

- [blocked_clients.bot_skip_processing](resources--http_loadbalancer--reference--group-011.md#canonical-2111320002233121-1332032232333101-0220121330111113-0102332012121311-0012310101321021-3313202101231323-0023331030210211-0123313012123213)
- [blocked_clients.http_header](resources--http_loadbalancer--reference--group-011.md#canonical-0230221313302212-2300233200220021-0311223130100202-3122001333022211-1330120103113220-2311323302123200-2002021232133220-1003100000033000)
- [blocked_clients.metadata](resources--http_loadbalancer--reference--group-011.md#canonical-3110213322201011-0103123011100223-3112113313330113-1023002030300111-0033330301202031-3332132022000213-2222021202322010-1231013331301002)
- [blocked_clients.skip_processing](resources--http_loadbalancer--reference--group-011.md#canonical-0130313120300133-3201120201321012-2121223223133003-2030221323001131-1003233221121103-3330032100110222-3012322122302023-2012002122023312)
- [blocked_clients.waf_skip_processing](resources--http_loadbalancer--reference--group-011.md#canonical-3101023120211000-3213013112130202-0320210320110301-0232020131122102-0122130233002131-3213203302322003-2100201130323002-0231223230232131)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2111320002233121-1332032232333101-0220121330111113-0102332012121311-0012310101321021-3313202101231323-0023331030210211-0123313012123213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022103000121200-1203003103102220-3213332202003330-1002021333231323-0020333323023131-3100330300312321-3121211133113012-2231023112031300"></a>

## blocked_clients.bot_skip_processing — bot_skip_processing / 011230121213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- blocked_clients.bot_skip_processing

<a id="canonical-0112123000020201-3102131223033300-1102221113123201-3101001231320322-2330300301120322-0123313112231033-0212133202301011-2221221312212221"></a>

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
bot_skip_processing = {}
```

<a id="canonical-3200130002223000-2300023320301132-1010200013111131-0021010322310010-1201011122000331-3233211322200011-0111022330323113-0012032010321330"></a>

## Direct properties — bot_skip_processing / 011230121213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101321333211231-2322212031322103-1332013000203123-1010311331110313-1321122223033212-3202330022001331-3330330101322210-1330330000303111"></a>

## Next pages — bot_skip_processing / 011230121213 / 4

- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0230221313302212-2300233200220021-0311223130100202-3122001333022211-1330120103113220-2311323302123200-2002021232133220-1003100000033000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323030031203202-3103102333013122-2312123201231112-3102230100310002-0321303331000212-1031200122212100-0222331333022322-3011120102303023"></a>

## blocked_clients.http_header — http_header / 110203012301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- blocked_clients.http_header

<a id="canonical-0001333302002210-0321313332331200-2131103011012020-2231330302110030-2130130312320132-0002111200321013-3230310022300021-0121022310200021"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0122223311121313-2332031032310023-2201021310020002-3121112302303122-0330320323332331-0330213201031310-1130322323023131-1331222021110332"></a>

## Direct properties — http_header / 110203012301 / 3

- [headers](resources--http_loadbalancer--reference--group-011.md#canonical-3223333312331330-0133131100312023-1222202302023323-1020102321223011-3013101320311223-1210230303100011-2023201132302000-2100112133012232): complete subsection reference.

<a id="canonical-1311102131312300-0123310200311322-3111002010110001-2301212111223130-0012012230300121-1023221003332031-0120031311230121-3033200312101333"></a>

## Next pages — http_header / 110203012301 / 4

- [blocked_clients.http_header.headers](resources--http_loadbalancer--reference--group-011.md#canonical-3223333312331330-0133131100312023-1222202302023323-1020102321223011-3013101320311223-1210230303100011-2023201132302000-2100112133012232)
- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3223333312331330-0133131100312023-1222202302023323-1020102321223011-3013101320311223-1210230303100011-2023201132302000-2100112133012232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033212103022000-1120221032312321-0313213333020202-1303030201003323-2210133331033003-1102231221333323-2132110331311001-1001033011230230"></a>

## blocked_clients.http_header.headers — headers / 331220210132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- [blocked_clients.http_header](resources--http_loadbalancer--reference--group-011.md#canonical-0230221313302212-2300233200220021-0311223130100202-3122001333022211-1330120103113220-2311323302123200-2002021232133220-1003100000033000)
- blocked_clients.http_header.headers

<a id="canonical-1033231112201131-3023130010111222-3033230301100322-1102310203002221-1012123100010111-2233021302121332-1312201302313112-3202233231022203"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP header name and value pairs.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1201112303321033-0112110230322001-2323332202101333-1001123013001032-0102300330001003-3031031231031300-0103113330132102-0113230223203301"></a>

## Direct properties — headers / 331220210132 / 3

<a id="canonical-0300131010210022-2010131123001330-2012021020311333-1102213213122203-2000233320120203-1313010111322311-1102233333123013-3301200233333100"></a>

<a id="canonical-1301203133211032-3012220101012030-2311013230331003-1210022231122222-0200110323132020-3133113331112203-2323000032313232-1303223003202230"></a>

## exact property — headers / 331220210132 / 4

Type: `"string"`. Optional.

Exclusive with \[presence regular expression\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regular expression\] Header value to match exactly.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-2020333103333221-0123330211231200-3203110330312313-2003310133033030-2321110220012000-2120130102011033-3003312003203310-1221021203000001"></a>

<a id="canonical-3100100331120121-3022200330022023-2231230212122102-0121211320021233-1131001100233321-0113031130110222-2022201300003200-1213230121102301"></a>

## invert_match property — headers / 331220210132 / 5

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

<a id="canonical-0331020132023011-2222010131203121-1122102012030230-0101310023032231-3211332213220320-0223223232113103-2011223232132323-3120023130113101"></a>

<a id="canonical-2232202002210033-1130002313311330-1331233022301130-3302300103121023-2012002000212132-1203010113232221-3321221302311332-0111211222120202"></a>

## name property — headers / 331220210132 / 6

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2313231023220131-0200012201101321-2110032120121211-3031200221211322-1131333122012112-0002033222120123-2222313333020123-3033330321030202"></a>

<a id="canonical-0012101202033313-3222021121031220-2302111002110232-0131102222201002-2133231012101100-3303123302112210-2130200111213233-3301001320130313"></a>

## presence property — headers / 331220210132 / 7

Type: `"bool"`. Optional.

Exclusive with \[exact regular expression\] If true, check for presence of header.

Upstream description:

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

<a id="canonical-2132323130130023-0302031110030033-1130200233002132-0203100321022203-0321003030122210-1230312112120123-1203302123011230-1012323312220011"></a>

<a id="canonical-0110011122002330-2230300310100211-2323101031201123-2102133013210330-3321221230232022-2330232211331003-0311003313032321-2131102100322233"></a>

## regular expression property — headers / 331220210132 / 8

Type: `"string"`. Optional.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0222130322210312-0310312231202101-1122210210213223-1221130320323121-2011132311012102-0110110332302222-3203031323123213-2210022010120220"></a>

## Next pages — headers / 331220210132 / 9

- [blocked_clients.http_header](resources--http_loadbalancer--reference--group-011.md#canonical-0230221313302212-2300233200220021-0311223130100202-3122001333022211-1330120103113220-2311323302123200-2002021232133220-1003100000033000)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3110213322201011-0103123011100223-3112113313330113-1023002030300111-0033330301202031-3332132022000213-2222021202322010-1231013331301002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203021111103121-3211322001302101-3122110333212030-0111202123300012-3110333320122130-2023303102123013-0230230320003001-1120011333003133"></a>

## blocked_clients.metadata — metadata / 302333020333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- blocked_clients.metadata

<a id="canonical-2313331102302312-2122131303312022-1010122223131202-3011102203102201-1101202033331220-1212213020101300-3033321021120021-0332100220010332"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1300032231301020-2131120111021000-1131030300012122-1113301333101131-2200013302311102-1023313122011331-2022320033203233-2300301200033220"></a>

## Direct properties — metadata / 302333020333 / 3

<a id="canonical-0220033212100310-0232203020102000-0110000131133103-3212323133130032-3320310130300330-0203322210333021-2222333112022232-3101332212103111"></a>

<a id="canonical-2221102110121231-2201222020221311-1130011003321023-3232010221023333-1300220222131323-3021233323201111-2302030310110302-2102031320311132"></a>

## description_spec property — metadata / 302333020333 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2103221223323010-0132112020001010-0131111200221332-0112123302112331-1000010112113321-2313210203333110-3330212132331332-1202003022002000"></a>

<a id="canonical-2322210311330010-1020132102122022-1330123330332322-3120112122002212-3101322103201230-3221203230302100-1103121202330113-1123012220333123"></a>

## name property — metadata / 302333020333 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0300220111020003-2211033130000233-0302031010033200-0303203100011213-1000032313130022-3000230033223322-0232111000102301-1331003122131112"></a>

## Next pages — metadata / 302333020333 / 6

- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0130313120300133-3201120201321012-2121223223133003-2030221323001131-1003233221121103-3330032100110222-3012322122302023-2012002122023312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320022003010101-2201322023222333-0211233200032032-3023320210221211-1331012131332033-1100133221012222-0100112121220121-1200010031330300"></a>

## blocked_clients.skip_processing — skip_processing / 331221013031 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- blocked_clients.skip_processing

<a id="canonical-2021111323123203-2213312311102223-1113123022031013-0230022030121122-1033320103203010-1121032220222311-2213210300333102-1033211001210323"></a>

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
skip_processing = {}
```

<a id="canonical-0101223112202100-1023213021300220-3302320332133033-3032102301232123-2032223110133022-3311033123133121-2223113202212302-0002010103112102"></a>

## Direct properties — skip_processing / 331221013031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232201321311101-0300123212323311-3003023311001010-1123100313011210-3213032310002022-0230203310311103-0123333031302222-0022102313321120"></a>

## Next pages — skip_processing / 331221013031 / 4

- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3101023120211000-3213013112130202-0320210320110301-0232020131122102-0122130233002131-3213203302322003-2100201130323002-0231223230232131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223021201213133-1022021002122322-3312233311201203-3302222103011122-2301122233100333-0112223312300120-3012110210001122-3313201230011222"></a>

## blocked_clients.waf_skip_processing — waf_skip_processing / 122113110232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- blocked_clients.waf_skip_processing

<a id="canonical-0120122023220132-1130132311332022-0322100113100322-0030002213331320-2231232322222013-0232023323323011-1203100100032232-3100120233032300"></a>

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
waf_skip_processing = {}
```

<a id="canonical-0100323131211233-1310020111011132-3132000023112003-0302131010311000-1223113222213111-3312031311221130-2133223120121022-0302210323011231"></a>

## Direct properties — waf_skip_processing / 122113110232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112123330311323-3112320012233300-1030020203032010-1121201133220120-2213002111302121-1122110313002313-2310001110032023-3233233123202130"></a>

## Next pages — waf_skip_processing / 122113110232 / 4

- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120032000222200-0010021300110020-2222231300231020-1303010302232303-2221233310311100-3010231221103130-2020300133301311-1031111113333330"></a>

## bot_defense — bot_defense / 033221230101 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- bot_defense

<a id="canonical-0313032320202312-3130032212302312-3120330303330321-1012133211302100-3231033302020103-2202000220123011-2231131012233312-2310002330302100"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: bot\_defense, bot\_defense\_advanced\_protection, disable\_bot\_defense; Default:
disable\_bot\_defense\] Defines various configuration OPTIONS for Bot Defense Policy.

Upstream description:

This defines various configuration OPTIONS for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_cors_support",
    "enable_cors_support")}
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
  "x-ves-oneof-field-cors_support_choice": "[\"disable_cors_support\",\"enable_cors_support\"]"
}
```

OneOf alternatives in this subsection:

- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0313032320202312-3130032212302312-3120330303330321-1012133211302100-3231033302020103-2202000220123011-2231131012233312-2310002330302100)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-1331023323322330-0120310010111332-1123202023223302-3231330202031222-2112233201102033-2111320023203321-1120030130231113-3212030313203022)
- [disable_bot_defense](resources--http_loadbalancer--reference--group-018.md#canonical-0031322221201123-0210020223332001-1212312133213033-3030023312222333-1332200121023202-1033110032330120-1010322121131033-3331133222011133)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bot_defense {
  # Configure direct properties listed below.
}
```

<a id="canonical-1331321031323210-3103221200002103-1003021100022311-0010331021131100-1303101211231221-1330121302133332-3211323313001111-2330010302203233"></a>

## Direct properties — bot_defense / 033221230101 / 3

- [disable_cors_support](resources--http_loadbalancer--reference--group-011.md#canonical-1020123332231113-2020103222100113-2122232221332121-1211111212021302-3002223112023123-2130122130110110-3233211000012303-1000032320312103): complete subsection reference.

- [enable_cors_support](resources--http_loadbalancer--reference--group-011.md#canonical-2000023030332210-2100002021230202-3310100211301003-0000113131311231-0130023331301033-3223323212221100-3032133113102211-2211111313333003): complete subsection reference.

- [policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013): complete subsection reference.

<a id="canonical-0203323233332000-0321101112201023-3000100113032122-0003310023133130-3233231002310032-2100202033311131-0133332112223103-3110010120131320"></a>

<a id="canonical-1132323332111010-1302231301011312-0110313332103002-1332033001213001-0312101210132013-0032323210120111-3113111100320133-0231231000200100"></a>

## regional_endpoint property — bot_defense / 033221230101 / 4

Type: `"string"`. Optional.

\[Enum: AUTO|US|EU|ASIA\] Defines a selection for Bot Defense region - AUTO: AUTO Automatic
selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA
Asia region. Possible values are \`AUTO\`, \`US\`, \`EU\`, \`ASIA\`. Defaults to \`AUTO\`.

Upstream description:

Defines a selection for Bot Defense region

&#8203;- AUTO: AUTO

Automatic selection based on client IP address &#8203;- US: US

US region &#8203;- EU: EU

European Union region &#8203;- ASIA: ASIA

Asia region.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AUTO",
    "US",
    "EU",
    "ASIA"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AUTO",
  "enum": [
    "AUTO",
    "US",
    "EU",
    "ASIA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2320230333301021-0232013112223132-0103103201012023-1130313001100210-3231111030031332-2131030200101323-1132122123011223-3320232102122221"></a>

<a id="canonical-1132010331222132-0103310110331100-2201123030202201-2111212002202311-3321012223101303-2332323202003012-0232321012113033-2331031010111023"></a>

## timeout property — bot_defense / 033221230101 / 5

Type: `"number"`. Optional.

The timeout for the inference check, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 60000),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-1321213111202122-2201323211023111-2021310300310213-0000121133222021-2201323330013112-1310213330321330-1313230211303202-2313011231010320"></a>

## Next pages — bot_defense / 033221230101 / 6

- [bot_defense.disable_cors_support](resources--http_loadbalancer--reference--group-011.md#canonical-1020123332231113-2020103222100113-2122232221332121-1211111212021302-3002223112023123-2130122130110110-3233211000012303-1000032320312103)
- [bot_defense.enable_cors_support](resources--http_loadbalancer--reference--group-011.md#canonical-2000023030332210-2100002021230202-3310100211301003-0000113131311231-0130023331301033-3223323212221100-3032133113102211-2211111313333003)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1020123332231113-2020103222100113-2122232221332121-1211111212021302-3002223112023123-2130122130110110-3233211000012303-1000032320312103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302320312023202-0033103101203310-2121212311300321-2302200132331212-1033213023330003-2103001110111003-0031332221101000-1313021211232133"></a>

## bot_defense.disable_cors_support — disable_cors_support / 210210003021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- bot_defense.disable_cors_support

<a id="canonical-1020012310020232-3331230311031220-2223113120230001-2113320023302220-0200302113323310-2131222132211230-1321131011020000-2200032001311310"></a>

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
disable_cors_support = {}
```

<a id="canonical-2021121010212110-1121013032101323-1203211000010221-3033331020212111-0111121031000310-3112130132203230-0000333213103001-2223212030031011"></a>

## Direct properties — disable_cors_support / 210210003021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322121120333232-0233130102211230-3312223333130210-3303112100201323-1110322231101300-2321323322323010-2310032220013223-0011320222323330"></a>

## Next pages — disable_cors_support / 210210003021 / 4

- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2000023030332210-2100002021230202-3310100211301003-0000113131311231-0130023331301033-3223323212221100-3032133113102211-2211111313333003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100012022302002-0020203131211032-2230323031112222-2130332103210101-2220023103103013-2302232302021113-2023303231331312-1213303201332110"></a>

## bot_defense.enable_cors_support — enable_cors_support / 310321110213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- bot_defense.enable_cors_support

<a id="canonical-2122011110300220-3121100330333311-0102112231122120-2312301000013210-2313310000132122-3002220001232101-1221203332110231-1123320202223133"></a>

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
enable_cors_support = {}
```

<a id="canonical-0321230101302210-1021222111001313-0233123311301323-3333112210311032-1210133300020333-3302121033311102-0103211232022112-0133320023010033"></a>

## Direct properties — enable_cors_support / 310321110213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111132301102310-1112132021223103-2123320113232323-0013122231223231-3100310123122100-1301102102012121-3012100212130331-2202300112001302"></a>

## Next pages — enable_cors_support / 310321110213 / 4

- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001033102330210-3032320032202312-0101123110010100-3111122002032003-1111301100220132-1220220200333023-2320332223031021-1202001321022223"></a>

## bot_defense.policy — policy / 211233102222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- bot_defense.policy

<a id="canonical-2222131100233100-2220002021300311-0313312030210120-3100202301102120-3010031021220201-0332113213232312-2022121122101212-1212001320211323"></a>

Type: `"object"`. single nested block, Optional.

Defines various configuration OPTIONS for Bot Defense policy.

Upstream description:

This defines various configuration OPTIONS for Bot Defense policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("protected_app_endpoints"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("disable_mobile_sdk",
    "mobile_sdk_config"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages_except",
    "js_insertion_rules")}
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

Terraform syntax:

```terraform
policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220013011230330-2023100102302120-1010120332302122-0213100001221001-2023001301001112-0112130133021320-0213001113113112-1011202213102101"></a>

## Direct properties — policy / 211233102222 / 3

- [disable_js_insert](resources--http_loadbalancer--reference--group-011.md#canonical-1023120000231211-1120322321201221-1310010000102023-0310221102300332-0010113223211330-1021003232220330-2222333021111203-1003303112012220): complete subsection reference.

- [disable_mobile_sdk](resources--http_loadbalancer--reference--group-011.md#canonical-2230330011120300-3101212101220222-1220232201312210-1112101223121321-0200330122010231-1323221201321200-0331012223032322-3200102123011303): complete subsection reference.

<a id="canonical-1010212102203023-0013311013313220-0233300311131010-1133132301300230-1102121221202231-0111110133331133-0333110320303130-2311021202311012"></a>

<a id="canonical-2201322302000112-0000333310303100-0200022312020233-1112013033202210-1200233102220030-1233203012112230-0210212213121211-3110330231002302"></a>

## javascript_mode property — policy / 211233102222 / 4

Type: `"string"`. Optional.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Upstream description:

Web Client JavaScript Mode.

Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable
Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is non-cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is cacheable.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ASYNC_JS_NO_CACHING",
  "enum": [
    "ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1122030331022310-0210132211023230-2030303021110302-2111103131213323-2033012333021012-2311123000001212-0133011133011312-3132031230021201"></a>

<a id="canonical-2320203321111302-2001330312212121-0301133003231232-0101303201323222-1002002320321321-2022012333213100-3221303330203110-0300203122012032"></a>

## js_download_path property — policy / 211233102222 / 5

Type: `"string"`. Optional.

Customize Bot Defense Client JavaScript path. If not specified, default

Upstream description:

Customize Bot Defense Client JavaScript path. If not specified, default \`/common.js\`

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
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [js_insert_all_pages](resources--http_loadbalancer--reference--group-011.md#canonical-3012330211022301-3011201012333021-2322122100300313-0123223010221231-1200011130111223-2330022223013320-3221323030122131-3213132322333022): complete subsection reference.

- [js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-0003112301221022-3203021200330313-3102213123032103-2122020001032201-2313120131131020-1031130310021132-3233031303221021-1201302320113211): complete subsection reference.

- [js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323): complete subsection reference.

- [mobile_sdk_config](resources--http_loadbalancer--reference--group-011.md#canonical-2123101112110001-2230223232113311-3200210313122201-3023210012031221-2033032323201332-0233120301211233-3302003023321332-3031131301202030): complete subsection reference.

- [protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021): complete subsection reference.

<a id="canonical-2102322222210023-1133320211221033-2020333211132333-3300223020231211-3012102121220231-1333102330102223-3001220032332230-2033303103110013"></a>

## Next pages — policy / 211233102222 / 6

- [bot_defense.policy.disable_js_insert](resources--http_loadbalancer--reference--group-011.md#canonical-1023120000231211-1120322321201221-1310010000102023-0310221102300332-0010113223211330-1021003232220330-2222333021111203-1003303112012220)
- [bot_defense.policy.disable_mobile_sdk](resources--http_loadbalancer--reference--group-011.md#canonical-2230330011120300-3101212101220222-1220232201312210-1112101223121321-0200330122010231-1323221201321200-0331012223032322-3200102123011303)
- [bot_defense.policy.js_insert_all_pages](resources--http_loadbalancer--reference--group-011.md#canonical-3012330211022301-3011201012333021-2322122100300313-0123223010221231-1200011130111223-2330022223013320-3221323030122131-3213132322333022)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-0003112301221022-3203021200330313-3102213123032103-2122020001032201-2313120131131020-1031130310021132-3233031303221021-1201302320113211)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-011.md#canonical-2123101112110001-2230223232113311-3200210313122201-3023210012031221-2033032323201332-0233120301211233-3302003023321332-3031131301202030)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1023120000231211-1120322321201221-1310010000102023-0310221102300332-0010113223211330-1021003232220330-2222333021111203-1003303112012220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303213320123210-0330320032132321-0012033122131010-1232110223101203-0001001103233311-0033210121110121-1311133111100221-0331232133312001"></a>

## bot_defense.policy.disable_js_insert — disable_js_insert / 131000002202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- bot_defense.policy.disable_js_insert

<a id="canonical-2201222023013111-2022320122113322-0103123122033301-2302031210111222-2310331330201102-1321233301201213-0332021112320223-1320110103000020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

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
disable_js_insert = {}
```

<a id="canonical-2210212232130300-3130011331010232-2023130021330010-3120120132011300-2310223001333121-3003110032302300-0331132121033121-0200103131111221"></a>

## Direct properties — disable_js_insert / 131000002202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333202021132313-2012131102111012-3002212000001120-0011020321032212-3311202222331111-2030232100122321-3232302211002123-0312002101311311"></a>

## Next pages — disable_js_insert / 131000002202 / 4

- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2230330011120300-3101212101220222-1220232201312210-1112101223121321-0200330122010231-1323221201321200-0331012223032322-3200102123011303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220233133333103-3213332221130312-0222123011021022-1023202100101301-0301230033230213-2133200333003031-2221100230332323-1320001121221133"></a>

## bot_defense.policy.disable_mobile_sdk — disable_mobile_sdk / 311231022110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- bot_defense.policy.disable_mobile_sdk

<a id="canonical-0020211222310220-2300231103123101-3312301120101133-3330222101330110-1020332021101123-2310002201201221-0323021013332312-3112120222013030"></a>

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
disable_mobile_sdk = {}
```

<a id="canonical-2230322031201133-0233001013302313-2220000230023220-3212320130112022-1112113301310033-1122023001300322-0002220321103113-1323303223200302"></a>

## Direct properties — disable_mobile_sdk / 311231022110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120123121012300-3113002322313200-1101301011301030-0311132210211211-2131230331122203-1302102030123120-0202032301200221-2021333132232321"></a>

## Next pages — disable_mobile_sdk / 311231022110 / 4

- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3012330211022301-3011201012333021-2322122100300313-0123223010221231-1200011130111223-2330022223013320-3221323030122131-3213132322333022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102131011022123-0001331101030211-2102133212033232-2232211121020033-3111013010221101-2121013300003200-1002101122211310-2312103202222031"></a>

## bot_defense.policy.js_insert_all_pages — js_insert_all_pages / 002013300133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- bot_defense.policy.js_insert_all_pages

<a id="canonical-3010223110121330-2211012121113201-1232233233103003-1021310112023321-2300120321123023-2133303300101013-2112322030010012-3221213020331111"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages.

Receipt-pinned upstream constraints:

```json
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
js_insert_all_pages {
  # Configure direct properties listed below.
}
```

<a id="canonical-1023101210210030-0011012102222131-2212212233132231-3121332102302321-1130313222220122-2033101312003011-0133211333302231-0231202113133230"></a>

## Direct properties — js_insert_all_pages / 002013300133 / 3

<a id="canonical-0301223223131311-0210322021030100-3022133133101103-2321222132232133-0023132221001210-0003132011103011-3331122022310130-0223120102223100"></a>

<a id="canonical-0121313213321121-0002233300200001-3121100110231331-0113003030123113-0132332331320101-0232330001212100-3102312223001211-0233132022002231"></a>

## javascript_location property — js_insert_all_pages / 002013300133 / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2301111022130111-2330101322313000-0231013311213010-0201102030311113-0030102001000123-3012122033223122-3212031200211121-1302022122202120"></a>

## Next pages — js_insert_all_pages / 002013300133 / 5

- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0003112301221022-3203021200330313-3102213123032103-2122020001032201-2313120131131020-1031130310021132-3233031303221021-1201302320113211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202313123113310-1033203112020131-2210100322312020-2213002320203330-1123200101310202-1232210101012212-1301320320223301-0120131323202123"></a>

## bot_defense.policy.js_insert_all_pages_except — js_insert_all_pages_except / 122221202332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- bot_defense.policy.js_insert_all_pages_except

<a id="canonical-0130201302011113-3300313233220221-0312202130112001-0033302221211032-1110113211123303-0211323010120203-0001012000120310-0200203031330112"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages with the exceptions.

Receipt-pinned upstream constraints:

```json
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
js_insert_all_pages_except {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121331021020203-3112313000000203-2110110211222202-0223311231031100-2020002020020002-3230322211030020-1322320333010323-1021111033331212"></a>

## Direct properties — js_insert_all_pages_except / 122221202332 / 3

- [exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330): complete subsection reference.

<a id="canonical-3023001301210112-0102331321212123-1132010121330303-2132113010301220-3313231311112203-2200211021130130-1130301100013311-2323222111033332"></a>

<a id="canonical-3312002322222321-0003032222201210-0023121110201331-0112230222203331-0221101033010100-2002130013233031-2111021303111332-3033231210213020"></a>

## javascript_location property — js_insert_all_pages_except / 122221202332 / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3032112110132331-1202202300312211-1333311032231130-3310022232333113-2203223223320022-3331323112332321-0310122013232301-1311331001132111"></a>

## Next pages — js_insert_all_pages_except / 122221202332 / 5

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122012023031032-1300220320331320-3201021023331300-2213211012101111-1200021033102113-0110002322333320-0313110020232010-1321211302300001"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list — exclude_list / 202330011033 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-0003112301221022-3203021200330313-3102213123032103-2122020001032201-2313120131131020-1031130310021132-3233031303221021-1201302320113211)
- bot_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-1322332203033330-0231332230212320-0322300021003213-0213020100110312-2200133203300120-3111120322210310-0330302021221303-1032021110203012"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130032010322132-2112032213303010-1333202120101212-2001013010202110-0100301233322021-3230030122021233-1133311023233312-1032011001020302"></a>

## Direct properties — exclude_list / 202330011033 / 3

- [any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-0013330210300103-1203112322213232-2003211320101132-1020013012003221-1213232023331020-0120332103113030-1031122233121332-2230011301211022): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-011.md#canonical-1312232001131331-1121013032013220-1221201221330323-2303101011333113-2212002001133202-1121002332302122-3123122002223111-2122223031221311): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-011.md#canonical-2311013103031000-3033100120030313-3300312013130000-1001203302010231-0102201032321131-0321022130232121-3221210322132101-3333033023221113): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-011.md#canonical-0233102133110232-3030221023310103-3202123303203321-1020133201330112-0310200211221303-1121210230121212-1220020032112313-0213123130213311): complete subsection reference.

<a id="canonical-1132003010333033-1332321202010202-1330322133033302-0210303013003000-1331111121301323-0320131320032011-0120111120323021-0320222121022130"></a>

## Next pages — exclude_list / 202330011033 / 4

- [bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-0013330210300103-1203112322213232-2003211320101132-1020013012003221-1213232023331020-0120332103113030-1031122233121332-2230011301211022)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.domain](resources--http_loadbalancer--reference--group-011.md#canonical-1312232001131331-1121013032013220-1221201221330323-2303101011333113-2212002001133202-1121002332302122-3123122002223111-2122223031221311)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata](resources--http_loadbalancer--reference--group-011.md#canonical-2311013103031000-3033100120030313-3300312013130000-1001203302010231-0102201032321131-0321022130232121-3221210322132101-3333033023221113)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.path](resources--http_loadbalancer--reference--group-011.md#canonical-0233102133110232-3030221023310103-3202123303203321-1020133201330112-0310200211221303-1121210230121212-1220020032112313-0213123130213311)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-0003112301221022-3203021200330313-3102213123032103-2122020001032201-2313120131131020-1031130310021132-3233031303221021-1201302320113211)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0013330210300103-1203112322213232-2003211320101132-1020013012003221-1213232023331020-0120332103113030-1031122233121332-2230011301211022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313033000001203-2203133322023002-3202310202101122-2322300022132303-3103331111301130-3333010233321232-2111201030012010-2232010012301121"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain — any_domain / 022212302103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-0003112301221022-3203021200330313-3102213123032103-2122020001032201-2313120131131020-1031130310021132-3233031303221021-1201302320113211)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-0001203132213032-3200033312031022-1331132322321232-1223232333003203-3030000101121333-2132320323102111-3103012013332021-2020031112020130"></a>

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
any_domain = {}
```

<a id="canonical-2121301013201130-2003300233010211-3213223032302233-2131311003001132-0211131313322001-2022320103033010-0013023321000201-1120033012132130"></a>

## Direct properties — any_domain / 022212302103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100333122023330-3012131312322001-2231313010020000-0020202203231321-3033100023212033-2320111202123022-0013213210113223-3310323203122300"></a>

## Next pages — any_domain / 022212302103 / 4

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1312232001131331-1121013032013220-1221201221330323-2303101011333113-2212002001133202-1121002332302122-3123122002223111-2122223031221311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220301132132001-0033300123131001-0302301022220331-0003212213110323-3123103113112212-0010123012230002-3213130213020011-0120210221300330"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.domain — domain / 221213221310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-0003112301221022-3203021200330313-3102213123032103-2122020001032201-2313120131131020-1031130310021132-3233031303221021-1201302320113211)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-0211122001222230-2200031303311020-3201020033223112-2013033213300131-0110111003030110-0001232311323022-2031210012311113-0232103231103122"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221010130300122-0101302113022030-1303333320021131-1321103010302100-1213213112020122-1100203210100012-1010323303111111-1013222031112123"></a>

## Direct properties — domain / 221213221310 / 3

<a id="canonical-2302231333110321-1313022022023321-0033113301312112-3203123100021100-0002100202330330-1331021203031322-3033310021010313-2300211311203120"></a>

<a id="canonical-3212003023230020-2012132231123131-0100100022323103-0131202232002313-0320113010311322-0120233332133101-0200030031212101-1322300223032033"></a>

## exact_value property — domain / 221213221310 / 4

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1013212031322133-2221102002100223-0001022201222132-3030011033220021-3233231301121213-0331110202220313-3223011111201112-2211331013313020"></a>

<a id="canonical-0233032022303223-2130212032110211-3223330032322213-2020032312300022-0130303312310121-3012000311102122-0001103010020022-2102301030332010"></a>

## regex_value property — domain / 221213221310 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1213011212022322-0212202321321311-1000011101001312-0030022233012322-1110201032221102-1112232001221020-3112010233010323-0313020023203023"></a>

<a id="canonical-2033033010013222-3002000011033111-1131213002001202-3020023313320110-1311332121232231-0231222120331113-0012120303002020-3333011000012301"></a>

## suffix_value property — domain / 221213221310 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2210220233031120-2100220113011323-2302313331111302-2303312031032201-3201230303231311-2123020220222010-3222301221000023-0033130030111310"></a>

## Next pages — domain / 221213221310 / 7

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2311013103031000-3033100120030313-3300312013130000-1001203302010231-0102201032321131-0321022130232121-3221210322132101-3333033023221113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211210322301010-3110210331133011-3311020001033302-0223230131113332-1330111131133122-2030000202101302-3030131003213202-1211201332203002"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata — metadata / 101102202121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-0003112301221022-3203021200330313-3102213123032103-2122020001032201-2313120131131020-1031130310021132-3233031303221021-1201302320113211)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-2323130231301302-1100113331012030-0220323210331031-1113333310012132-1323203011110132-0103211012022031-1321332133032211-1101332033310212"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220202102201101-3002122003020201-2313130313320032-0000202301133222-1322331220102122-0321031210220032-1303021302222121-1211200002231131"></a>

## Direct properties — metadata / 101102202121 / 3

<a id="canonical-2101032102102012-2231222312231310-3011222200230110-3111322300022131-2222302331002003-1233003003003221-0123031231000010-2320231032333320"></a>

<a id="canonical-2310002103012113-2313033202133033-3113022220013330-2320321032321202-0233321121331110-2332130231023210-2311213132211021-2311312113210021"></a>

## description_spec property — metadata / 101102202121 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2021022012011331-2322232011021132-2201231032221333-0213001321310230-1311000131331201-2313331010010233-3120100113301123-3322033101000322"></a>

<a id="canonical-2111121231033312-3320100212333231-0331301102121031-2331222001312120-2100212233223122-0320031011221100-1100033032222313-1033012333103302"></a>

## name property — metadata / 101102202121 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1022232223213133-2111030023220020-0112323300233302-1202331130201001-0022201103010000-0311313113320000-2221321102011233-3302303330333001"></a>

## Next pages — metadata / 101102202121 / 6

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0233102133110232-3030221023310103-3202123303203321-1020133201330112-0310200211221303-1121210230121212-1220020032112313-0213123130213311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310233331011210-3001233202021111-1110333313110322-2133131121312000-1220232121131202-2113200013021212-2210310233231233-3103320112303102"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.path — path / 003202131233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-0003112301221022-3203021200330313-3102213123032103-2122020001032201-2313120131131020-1031130310021132-3233031303221021-1201302320113211)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-0032033202321002-0210212031111202-1331033003033211-1330023302332010-2122010231101000-3122012130321010-2330201301231103-3210323231212033"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2201111322121212-2211210100321031-0323331032030301-0210212013213232-1322010320000132-0012310332002232-1010323022333033-2210320213313002"></a>

## Direct properties — path / 003202131233 / 3

<a id="canonical-0113231110320131-3321302222322112-2301130133322203-2310333132113022-1310221012302111-1132232230223222-0330203312120333-2111010211233322"></a>

<a id="canonical-3131222201201022-2330112000303010-3000122003123233-2222033211110020-2002220211233222-2232002021113122-1323300020302103-0122332110330120"></a>

## path property — path / 003202131233 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0331202130013030-1000313032300001-3110110303210121-3130012311310003-3132312231203123-1312312030203211-0232000212023010-0321313131332121"></a>

<a id="canonical-1020010013223100-1113213233031303-0312012233130300-0010023331310223-2302130111320000-3010003303102203-3121223332331232-3330103321300300"></a>

## prefix property — path / 003202131233 / 5

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0213331132102113-3002210232300033-3301302100102103-3320121102012112-1321210131120203-1031222223210132-1233233211103111-0010201312312320"></a>

<a id="canonical-0013300230030133-1320000330122032-1203331012300301-1000121220333120-1113301023013030-2212033321120332-2202032330231330-2001012323230302"></a>

## regular expression property — path / 003202131233 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1031021332233210-0203021203201221-3011132301030321-3132131333320320-1003211223022101-2122102022120202-0020333101123100-0310300032313001"></a>

## Next pages — path / 003202131233 / 7

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320113312102032-1321212300331211-1222322102001031-3003333112032213-2012030322213300-1321230330122001-3012120003031223-1111300322202211"></a>

## bot_defense.policy.js_insertion_rules — js_insertion_rules / 012001200211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- bot_defense.policy.js_insertion_rules

<a id="canonical-1103211231213033-2223301031103132-1031330333130321-3301020300113332-1201332231120320-0133103202220311-3332022102021102-2023130221210131"></a>

Type: `"object"`. single nested block, Optional.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122321021100231-0100200012020122-3020000200013101-0122200200222331-1023220101003320-1300303313330313-1230202001113222-0011033132112101"></a>

## Direct properties — js_insertion_rules / 012001200211 / 3

- [exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132): complete subsection reference.

- [rules](resources--http_loadbalancer--reference--group-011.md#canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031): complete subsection reference.

<a id="canonical-1101222133302131-3013323130310120-0003132210101103-1222313210231212-2222331123011301-2213321320022102-2223222130110132-3313322110012310"></a>

## Next pages — js_insertion_rules / 012001200211 / 4

- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132)
- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121010220313011-0000230010132311-2020331301022201-0012333320302102-2112202132230113-0302011031011010-3032121213222232-1123231003132120"></a>

## bot_defense.policy.js_insertion_rules.exclude_list — exclude_list / 211033331110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- bot_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-1013221201111120-3013220222301201-1221300111001323-0123032011012332-3320202313231113-2003333323131230-2220301213320130-1102313213200212"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2331020111321301-0222301111011203-3212122000032121-3312331332033302-1112010132211100-2133231332011320-1221321122203323-1330322103233121"></a>

## Direct properties — exclude_list / 211033331110 / 3

- [any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-1221220223333201-0010122311110311-2231131121203202-1021031031023130-1203320002102002-3020112220130021-0302013010012232-2122212020133221): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-011.md#canonical-1301021021033303-3230122321212212-3201323201222312-1111132010310211-0023201303210130-3320213120010002-1222032302112321-2221201133123132): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-011.md#canonical-0222121320012233-3330310002200133-1301202100032031-0101210210020130-1211303233220213-0133131320100230-0013010000113030-1321000302230022): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-011.md#canonical-0113132333232330-3231013001310232-0233321111200331-0311220003200130-2233301330231221-0202003211332213-0300113013322033-3010203100112221): complete subsection reference.

<a id="canonical-0102312230020210-1202033223331010-1133331103030102-2022220210200033-1220312231311212-1033211121201233-1103111230203200-0200213000230000"></a>

## Next pages — exclude_list / 211033331110 / 4

- [bot_defense.policy.js_insertion_rules.exclude_list.any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-1221220223333201-0010122311110311-2231131121203202-1021031031023130-1203320002102002-3020112220130021-0302013010012232-2122212020133221)
- [bot_defense.policy.js_insertion_rules.exclude_list.domain](resources--http_loadbalancer--reference--group-011.md#canonical-1301021021033303-3230122321212212-3201323201222312-1111132010310211-0023201303210130-3320213120010002-1222032302112321-2221201133123132)
- [bot_defense.policy.js_insertion_rules.exclude_list.metadata](resources--http_loadbalancer--reference--group-011.md#canonical-0222121320012233-3330310002200133-1301202100032031-0101210210020130-1211303233220213-0133131320100230-0013010000113030-1321000302230022)
- [bot_defense.policy.js_insertion_rules.exclude_list.path](resources--http_loadbalancer--reference--group-011.md#canonical-0113132333232330-3231013001310232-0233321111200331-0311220003200130-2233301330231221-0202003211332213-0300113013322033-3010203100112221)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1221220223333201-0010122311110311-2231131121203202-1021031031023130-1203320002102002-3020112220130021-0302013010012232-2122212020133221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331232030010120-2022223010310303-1112311323013302-3113333023033211-1233033113023310-0213122020232201-2323032232220110-0103230322102132"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.any_domain — any_domain / 230303003301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132)
- bot_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-2100203121003221-0310010233213133-2022023120302212-2023023310221032-2112101311231122-3302211211033332-3230133101221012-2222121231323223"></a>

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
any_domain = {}
```

<a id="canonical-0021221303330030-1300111301313330-1323323323212121-1303221112012232-0103213231303120-3012231311202311-2221120330110012-2220300301102233"></a>

## Direct properties — any_domain / 230303003301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202100031212312-1221010102323233-3113113213000032-2333001121311230-1133133223000322-2230212122333002-3012220233330300-0231121220330303"></a>

## Next pages — any_domain / 230303003301 / 4

- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1301021021033303-3230122321212212-3201323201222312-1111132010310211-0023201303210130-3320213120010002-1222032302112321-2221201133123132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222221120102201-3323100110111213-2101310221123120-2002311201233013-3300122333200322-1210312031222132-3332112213011020-1002101003223301"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.domain — domain / 333303301231 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132)
- bot_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-1130230202222103-1330030323201020-2102221223000111-0120101133022023-1003100232313020-2202301212213320-1013313010211021-2012033010303323"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011313101332222-2130232132021312-1310301003232330-1312313232123033-3001031100120220-2011122312130012-2012020103113223-1313333233333132"></a>

## Direct properties — domain / 333303301231 / 3

<a id="canonical-2013030002322001-1133223103010233-1120230323300103-2123133000003120-0032332112031010-1101300122033313-1213210003011012-2330313222221213"></a>

<a id="canonical-1121002133010101-2300013221031300-3130321231102202-3221300330221212-0210232002222300-3210201222333021-1113203131220222-0223320033022002"></a>

## exact_value property — domain / 333303301231 / 4

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2222001330323212-2311111003300301-1120332011103320-2021322123201001-3022111323030221-2330322010102133-3332102330222020-0211010023201030"></a>

<a id="canonical-1330123112200103-1130321131223231-3100020312220300-0021003031001033-0100301233213322-0231131021020033-1333123330133312-0320200133111123"></a>

## regex_value property — domain / 333303301231 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0301101021021230-3030210210301132-3021320201031020-3321033110213323-2221130211122130-3303113221231313-3301101131231110-3121213321101002"></a>

<a id="canonical-1012120131332101-0100120121123303-3003133010301210-3332313312212011-2203123122022132-1223021201122103-1311320312001122-1202102332320310"></a>

## suffix_value property — domain / 333303301231 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0302122121002211-2333120010002311-3002223211103310-2211022232333203-0322022030021033-2000213211132000-0113222331233321-2020303330300211"></a>

## Next pages — domain / 333303301231 / 7

- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0222121320012233-3330310002200133-1301202100032031-0101210210020130-1211303233220213-0133131320100230-0013010000113030-1321000302230022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321313012322030-3310030231322131-0221120333102323-1113211023231111-3110200001232113-1130131220212011-0032300000321332-3022302222221031"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.metadata — metadata / 311210231300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132)
- bot_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-1222020333010313-2222102301120120-1331220131211023-0311131201333021-2230311030331233-2203333033302113-1002302323120010-1233121120213320"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022001221333320-1203120210300103-2212122223312020-1203000300301232-2032323020212313-0321103330212210-2320100323231011-0013132021323233"></a>

## Direct properties — metadata / 311210231300 / 3

<a id="canonical-1232002133113020-3320223033323113-0020013132332110-0201031330103211-0132231130222131-3123203121232011-1312302211201220-1331131113022013"></a>

<a id="canonical-3021302301303100-3102210332112313-2000323312121312-3001132223323110-0202212133302210-1010303312101100-3001120033122212-1122220100111310"></a>

## description_spec property — metadata / 311210231300 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3330212231102011-0130323212313303-2331321131000131-3101111211301310-3330020320331212-1112132111310023-1113033103233200-0120211123300322"></a>

<a id="canonical-1012001210331022-1111210330122020-0203331330200020-3332031100220213-0120021033130220-0012121020111221-1112201022003131-2020300102301200"></a>

## name property — metadata / 311210231300 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3202332301320132-2332021002010110-1103103012011203-3203011221032221-3302303111223212-0210322202203122-0110231033023222-1120310210122231"></a>

## Next pages — metadata / 311210231300 / 6

- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0113132333232330-3231013001310232-0233321111200331-0311220003200130-2233301330231221-0202003211332213-0300113013322033-3010203100112221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130030231133231-2232120222023313-0221110032311023-2321023232122312-0130302121030131-3021033022033222-2200021012332020-3313003213112013"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.path — path / 222000321121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132)
- bot_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-3102010230332230-0112221300321110-0232130111120300-3001302322002123-1003131313311020-1231320130320001-0311032223120131-0230102222212230"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201130000210320-1001133030322001-1032031232333021-3313302132101010-3332001032031013-2321211300012031-3030203101103321-0120332213231132"></a>

## Direct properties — path / 222000321121 / 3

<a id="canonical-0032002131013102-1012022332010222-2301202022330220-3213330031311310-0100120322333132-3322331122022100-1001232013303231-2300031302311013"></a>

<a id="canonical-2212111030331303-0112301113022200-1122111122213221-2220332220123300-2301132330223302-1310100220030313-2011231132022233-1012011231321020"></a>

## path property — path / 222000321121 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0212102102330030-2103001110101202-1312232223300001-1232012033301323-2123202201121320-2302031110030002-3322011031203303-3112022103313211"></a>

<a id="canonical-2032303001223001-3310011001002102-0102203201302013-0032203123110211-1112212233133302-0001110301133123-3310222023111130-2231111122122332"></a>

## prefix property — path / 222000321121 / 5

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2303012013211102-2101322302223222-0032230302103111-2112203013100311-2033123303131121-1120033112330321-0330112002130011-3230233032010011"></a>

<a id="canonical-2202022220220122-3332003101303322-2200013233033012-1322121122022000-1102201030300033-0001112220101211-1301113132230121-2102000122310003"></a>

## regular expression property — path / 222000321121 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0100300103221201-1100111222302132-0032233020033132-0233023123120313-1211103302302330-3113330311103220-2030201023031323-1320323033221123"></a>

## Next pages — path / 222000321121 / 7

- [bot_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-1010212333222220-0130321030121332-1302201032300020-2312032100120123-3310011122121022-0221032110101300-1312102131331323-3331112031311132)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303300012201023-3231212131120232-1132332032220321-0111121102130303-2300003323320330-1103133110313200-3000223320222220-2320223113221233"></a>

## bot_defense.policy.js_insertion_rules.rules — rules / 320121312113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- bot_defense.policy.js_insertion_rules.rules

<a id="canonical-2100002131133131-1031332133313121-0330303132011333-2313110111110102-2200220113202221-3332133322321313-3031232120130002-2000202122233221"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Bot Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322332230033112-1032033211220022-1211133000322100-0121301100011031-2313222100331313-3020310100120001-2030011121033013-2220301011211020"></a>

## Direct properties — rules / 320121312113 / 3

- [any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-1001002323002012-3110312013223130-2112013322032231-2121302330331200-3002011200120122-3200112210221223-0232130030100310-0100021123330310): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-011.md#canonical-2000321122000302-3110033333020130-1111013102331231-3013331112111232-3111201332210213-3130000122022333-2331111321011221-2030120200221031): complete subsection reference.

<a id="canonical-2203033303010310-2313012001212322-3003012123031133-1013023313022121-0203122132330103-3310003123021003-2022122302022211-3303101003332212"></a>

<a id="canonical-3301310212313100-2103122030030110-0122301322102303-0110230011020102-3210202001132132-1310002121132301-0022132011132000-3010311333210001"></a>

## javascript_location property — rules / 320121312113 / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-011.md#canonical-2333131102203011-1321321100300330-3323312330133033-3302313013310323-2122100130203302-3022020003311213-1011222232221301-1110303123111312): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-011.md#canonical-3111321310231223-2020000000111022-2010000011123120-1113323320130110-3012112311323323-1301302331103020-0003311010013012-0222003210220021): complete subsection reference.

<a id="canonical-3221212103201032-1030233102111212-2012320321012332-2032222021031321-0303320121021321-1013231331112112-3112232122132221-0032100023031030"></a>

## Next pages — rules / 320121312113 / 5

- [bot_defense.policy.js_insertion_rules.rules.any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-1001002323002012-3110312013223130-2112013322032231-2121302330331200-3002011200120122-3200112210221223-0232130030100310-0100021123330310)
- [bot_defense.policy.js_insertion_rules.rules.domain](resources--http_loadbalancer--reference--group-011.md#canonical-2000321122000302-3110033333020130-1111013102331231-3013331112111232-3111201332210213-3130000122022333-2331111321011221-2030120200221031)
- [bot_defense.policy.js_insertion_rules.rules.metadata](resources--http_loadbalancer--reference--group-011.md#canonical-2333131102203011-1321321100300330-3323312330133033-3302313013310323-2122100130203302-3022020003311213-1011222232221301-1110303123111312)
- [bot_defense.policy.js_insertion_rules.rules.path](resources--http_loadbalancer--reference--group-011.md#canonical-3111321310231223-2020000000111022-2010000011123120-1113323320130110-3012112311323323-1301302331103020-0003311010013012-0222003210220021)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1001002323002012-3110312013223130-2112013322032231-2121302330331200-3002011200120122-3200112210221223-0232130030100310-0100021123330310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330220211033100-1131123102133123-3122011310301300-2311331020020130-0213323023210333-2131302010002113-1212032132301120-3000031303002322"></a>

## bot_defense.policy.js_insertion_rules.rules.any_domain — any_domain / 023110220330 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031)
- bot_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-2203330121121303-1323131203030300-0021322221303030-3233120220302120-2210101323300120-2021323211102330-3110103333101220-3123121103120320"></a>

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
any_domain = {}
```

<a id="canonical-3220321102031210-2221230133022121-2033010001321201-3331101122313211-1220311133230300-2213033311220301-2313200021303200-1312031012030133"></a>

## Direct properties — any_domain / 023110220330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031310132203213-3102032211011122-0220200012232232-3013220010213132-0001232010120100-2032110112021101-3010122133012133-2111131113211312"></a>

## Next pages — any_domain / 023110220330 / 4

- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2000321122000302-3110033333020130-1111013102331231-3013331112111232-3111201332210213-3130000122022333-2331111321011221-2030120200221031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201101313122123-1112131121130232-2023012112332223-2132301112012332-1332321131312103-3223201002321103-0331000001310320-1022301202033101"></a>

## bot_defense.policy.js_insertion_rules.rules.domain — domain / 213331300323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031)
- bot_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-3223002013132311-2102320003330110-1001033122112321-2112121030002211-0130223321323133-2212131103311103-2000113202102330-3112013000032320"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333121120033310-1100300303023100-2012111301300322-2013011211103102-2003013030323003-2123100012211032-2231232130333312-2323202323200201"></a>

## Direct properties — domain / 213331300323 / 3

<a id="canonical-2202333231120313-3312103000320023-3122323221111222-2002000122112223-2011300001211202-1030322030010211-2111323331331033-3032010322120111"></a>

<a id="canonical-2321330203031322-1220112032332132-3221112033113302-1031112010133303-2230332230001103-3212013111331213-2020031212022200-2233020331230302"></a>

## exact_value property — domain / 213331300323 / 4

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1022313002003232-1101312221233013-3121230020220310-1102333301201300-1011120123300122-0003011213021032-1233213331323200-1230130102123033"></a>

<a id="canonical-1333310002313310-2312123011101210-0302003001233123-0332103110020303-0220132332223001-3101022301313330-3112202122131303-2212202002310132"></a>

## regex_value property — domain / 213331300323 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3302311130231011-1232332122030021-0123112222210011-0211321322031113-1010220032200002-2211000320121130-0121002103002132-2233310231130111"></a>

<a id="canonical-0330022013133021-3001002122321221-3110203312130200-1330200312010200-3102011323012201-1202020313030010-0010113111100223-2222123100103302"></a>

## suffix_value property — domain / 213331300323 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0022012210012020-3321131311200303-1121110302312300-2201222121111312-3231333022133222-1023103020203332-3122003311212203-3120110313011013"></a>

## Next pages — domain / 213331300323 / 7

- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2333131102203011-1321321100300330-3323312330133033-3302313013310323-2122100130203302-3022020003311213-1011222232221301-1110303123111312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033330013211231-2333210033001123-0021110233032230-0231211310310133-0001023300001313-3101003120301123-2110112132232003-3221123223022011"></a>

## bot_defense.policy.js_insertion_rules.rules.metadata — metadata / 030313302112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031)
- bot_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-2123300000320021-3221033021002013-0202122101023301-0121120121321003-2113102201302102-1112220033121231-0211112010001310-3132102010313031"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310011301113020-0322203223202211-3223003303231302-2031030133331101-2112310112202001-0130203202331313-2333222021020330-1010332221132100"></a>

## Direct properties — metadata / 030313302112 / 3

<a id="canonical-0302331203212100-0331003101331332-1000110022302032-0223021331222323-0030211230311000-0313303202330123-3322021200233200-2112232201321302"></a>

<a id="canonical-2030320330103122-2333012002323303-2232221323012203-0020012313223100-3102111313011312-1220111202303012-2023202001100133-2130132031313301"></a>

## description_spec property — metadata / 030313302112 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2331230212213112-2020303232132133-3203003321111310-1231210022010031-1133200213121121-3231002131121020-0010030032203330-2001031231031222"></a>

<a id="canonical-2132122032120020-3221222020001123-3030332231023110-0121333203100310-3212310313023212-0023020103012221-3122012121001012-0031001121121010"></a>

## name property — metadata / 030313302112 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0002222323131110-1103032331321321-1121010131310012-1322223122011212-2030230200033002-3010000100221110-1122320020122100-0221311031203223"></a>

## Next pages — metadata / 030313302112 / 6

- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3111321310231223-2020000000111022-2010000011123120-1113323320130110-3012112311323323-1301302331103020-0003311010013012-0222003210220021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213201032311310-2320000302211232-0031130323021012-2121022133122020-0223222313313332-2302113232302122-0223133012132332-1313300301233001"></a>

## bot_defense.policy.js_insertion_rules.rules.path — path / 032213133032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323)
- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031)
- bot_defense.policy.js_insertion_rules.rules.path

<a id="canonical-2310110003003322-3312100300203213-2221110320303001-2210000000320011-1313211103033120-1121100132032112-1323203012321200-0103201113102203"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232120100231100-3102123321030200-0112130031120122-2231112230320323-1321303222300133-1111221113010221-0210102120213230-0032211312321300"></a>

## Direct properties — path / 032213133032 / 3

<a id="canonical-0313231013122002-3333202211130033-2111133222023012-0121233032130230-2131223002321202-1023200120023210-1002333211012133-2103031123112320"></a>

<a id="canonical-1013201113310213-2123010100220101-1103223212302312-0330130013102210-2013321321300212-0202330120133231-2033030302301321-1121232231221012"></a>

## path property — path / 032213133032 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2203330103300022-3031330130020100-1112122100133320-1010211210033320-2310302021031301-0021010123120010-2202032232222033-3232320300300123"></a>

<a id="canonical-0231112032323032-3320110020103100-0000221232032031-1131101302120001-3202101221011223-2332223230211003-2132211031332331-0100101232211233"></a>

## prefix property — path / 032213133032 / 5

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3311312210001230-0102231210322312-2132223013221300-2231211112032002-1301233100311323-1321112001203203-3211230113321230-3103322333022102"></a>

<a id="canonical-0201303010102231-0230233330311230-3001331223102000-1010012233310201-3320103133030130-1323110320123132-0021233333202230-3202133311022101"></a>

## regular expression property — path / 032213133032 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1113031311032311-1313201221032213-2202011302113323-0030223211103312-2231110100303120-0203030213233300-2303113212121320-1012030003330231"></a>

## Next pages — path / 032213133032 / 7

- [bot_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-011.md#canonical-1130301220221103-0112000003000111-0100230331130213-2332312001030111-3320033112311331-0033100303133232-2021320233013331-1312220123320031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2123101112110001-2230223232113311-3200210313122201-3023210012031221-2033032323201332-0233120301211233-3302003023321332-3031131301202030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303132233331322-1003112010212222-0031212002301231-3221122310121012-3003030333102223-2120233300011133-0222103312331000-1300121331113111"></a>

## bot_defense.policy.mobile_sdk_config — mobile_sdk_config / 300203122212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- bot_defense.policy.mobile_sdk_config

<a id="canonical-0130221312123033-0313332210323112-0033311110202223-0302312211013330-3002202223300203-3321101321022011-3330221212111300-0201013100213013"></a>

Type: `"object"`. single nested block, Optional.

Mobile SDK Configuration. Mobile SDK configuration.

Upstream description:

Mobile SDK configuration.

Receipt-pinned upstream constraints:

```json
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
mobile_sdk_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3321220033221020-1012320330223102-2210213121023033-0312210031031001-0031001030212102-1202013223200101-0331003033332103-0012221000333031"></a>

## Direct properties — mobile_sdk_config / 300203122212 / 3

- [mobile_identifier](resources--http_loadbalancer--reference--group-011.md#canonical-2020001123330232-1331033320103030-2100212231130230-2211020212323023-2231323333310011-1130221220222202-1100231030232103-0020022130322213): complete subsection reference.

<a id="canonical-1112023000022102-2222111000112232-0302333002303330-2003233220203330-1233301131001203-1210311001330331-1300011023212023-1230312121122301"></a>

## Next pages — mobile_sdk_config / 300203122212 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-011.md#canonical-2020001123330232-1331033320103030-2100212231130230-2211020212323023-2231323333310011-1130221220222202-1100231030232103-0020022130322213)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2020001123330232-1331033320103030-2100212231130230-2211020212323023-2231323333310011-1130221220222202-1100231030232103-0020022130322213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213032122222103-1330020110030132-0023111022203221-1212021031021103-3003001320202011-2303020222320303-1023003213200113-1110120331030103"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier — mobile_identifier / 332312003133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-011.md#canonical-2123101112110001-2230223232113311-3200210313122201-3023210012031221-2033032323201332-0233120301211233-3302003023321332-3031131301202030)
- bot_defense.policy.mobile_sdk_config.mobile_identifier

<a id="canonical-1023033232003103-3320133232023332-3032002001222220-3300211133111232-2300021323011221-3121220221302032-2012021211110023-3011133122233013"></a>

Type: `"object"`. single nested block, Optional.

Mobile Traffic Identifier. Mobile traffic identifier type.

Upstream description:

Mobile traffic identifier type.

Receipt-pinned upstream constraints:

```json
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
mobile_identifier {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003132001313210-2131122013100302-3001030222301211-0103232020130332-0223031101321220-3020002010020131-0021012330211210-0100020301321212"></a>

## Direct properties — mobile_identifier / 332312003133 / 3

- [headers](resources--http_loadbalancer--reference--group-011.md#canonical-0130220100220230-3121123301113020-2302321022220303-0123123132122233-0220021333320201-3322321310102330-0320223123003030-1103200022300321): complete subsection reference.

<a id="canonical-1232202110303102-3331203302332232-3331211211120013-2112033123033311-0001302012020020-3122330032202230-1100331322112231-0021213101333331"></a>

## Next pages — mobile_identifier / 332312003133 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-011.md#canonical-0130220100220230-3121123301113020-2302321022220303-0123123132122233-0220021333320201-3322321310102330-0320223123003030-1103200022300321)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-011.md#canonical-2123101112110001-2230223232113311-3200210313122201-3023210012031221-2033032323201332-0233120301211233-3302003023321332-3031131301202030)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0130220100220230-3121123301113020-2302321022220303-0123123132122233-0220021333320201-3322321310102330-0320223123003030-1103200022300321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333200333203300-3111220031332232-2333211222112012-2302232031031122-1331203201311202-3211110332330203-1030122310103203-2323033300303003"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers — headers / 203211301331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-011.md#canonical-2123101112110001-2230223232113311-3200210313122201-3023210012031221-2033032323201332-0233120301211233-3302003023321332-3031131301202030)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-011.md#canonical-2020001123330232-1331033320103030-2100212231130230-2211020212323023-2231323333310011-1130221220222202-1100231030232103-0020022130322213)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-0123212033121002-1302313111011001-0112020301121321-3230233301130003-0230212030303312-3320310101130032-2312020001100001-2013311023131311"></a>

Type: `"object"`. list nested block, Optional.

Headers that can be used to identify mobile traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220112223021333-2212230203310030-3133331122223003-2013002312213032-1010110101030010-0032213021300113-2011231302312110-1213001200232213"></a>

## Direct properties — headers / 203211301331 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-011.md#canonical-0011101013202231-3320022332210222-0221111210232201-0200302133323223-3000010030130102-0013200013301212-1332301001223303-3331102021320221): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-011.md#canonical-1210333020301111-3030330011120122-3133013022002203-0203213232003310-2201220000020222-0021021010011320-1032212332122310-0011311100231301): complete subsection reference.

- [item](resources--http_loadbalancer--reference--group-011.md#canonical-2222322013313020-0312211232300320-3103122221102233-2230133301322103-0200120220011002-2013021033213100-0233101301323221-3031123110012032): complete subsection reference.

<a id="canonical-3223301332113330-1321021313210130-1223332130323320-0300331310011013-1130120200131013-1231310320013122-0332022223322323-2111112301211012"></a>

<a id="canonical-3032111300211332-2110222131003030-2121220223102010-2113102333001333-0303011200300010-1121031002213322-3100122010032333-2003313022202010"></a>

## name property — headers / 203211301331 / 4

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0322012021011233-2022120310312211-1203030231233132-2301101210200003-3020111231223022-2020330110020111-3313122132311131-2133001012311112"></a>

## Next pages — headers / 203211301331 / 5

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present](resources--http_loadbalancer--reference--group-011.md#canonical-0011101013202231-3320022332210222-0221111210232201-0200302133323223-3000010030130102-0013200013301212-1332301001223303-3331102021320221)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present](resources--http_loadbalancer--reference--group-011.md#canonical-1210333020301111-3030330011120122-3133013022002203-0203213232003310-2201220000020222-0021021010011320-1032212332122310-0011311100231301)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item](resources--http_loadbalancer--reference--group-011.md#canonical-2222322013313020-0312211232300320-3103122221102233-2230133301322103-0200120220011002-2013021033213100-0233101301323221-3031123110012032)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-011.md#canonical-2020001123330232-1331033320103030-2100212231130230-2211020212323023-2231323333310011-1130221220222202-1100231030232103-0020022130322213)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0011101013202231-3320022332210222-0221111210232201-0200302133323223-3000010030130102-0013200013301212-1332301001223303-3331102021320221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203332212131131-1031221113123202-3302201122300002-1213123100220013-3131310113021010-1333020232233231-1121032032022112-2313131313312020"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present — check_not_present / 133012011333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-011.md#canonical-2123101112110001-2230223232113311-3200210313122201-3023210012031221-2033032323201332-0233120301211233-3302003023321332-3031131301202030)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-011.md#canonical-2020001123330232-1331033320103030-2100212231130230-2211020212323023-2231323333310011-1130221220222202-1100231030232103-0020022130322213)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-011.md#canonical-0130220100220230-3121123301113020-2302321022220303-0123123132122233-0220021333320201-3322321310102330-0320223123003030-1103200022300321)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present

<a id="canonical-1133112122223310-2021103221121312-2310030100013333-0010102233301031-3201133022220000-2202113312311022-1220030022122021-0203210001101323"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-2001102301312301-0102022231113131-1033213230101333-0102320231221030-3221301300233030-2320301020023001-1321220232121032-3320212100032130"></a>

## Direct properties — check_not_present / 133012011333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200001112010130-1311210223133201-2000102123032212-0200033101002021-3133020021023130-1233313030131213-1133003113111221-0103211222013223"></a>

## Next pages — check_not_present / 133012011333 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-011.md#canonical-0130220100220230-3121123301113020-2302321022220303-0123123132122233-0220021333320201-3322321310102330-0320223123003030-1103200022300321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1210333020301111-3030330011120122-3133013022002203-0203213232003310-2201220000020222-0021021010011320-1032212332122310-0011311100231301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321213123132121-3333303000130203-0121200303132333-3323120102201102-0221122111322031-1303312223232132-2011231233122101-1202322310220302"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present — check_present / 211333220101 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-011.md#canonical-2123101112110001-2230223232113311-3200210313122201-3023210012031221-2033032323201332-0233120301211233-3302003023321332-3031131301202030)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-011.md#canonical-2020001123330232-1331033320103030-2100212231130230-2211020212323023-2231323333310011-1130221220222202-1100231030232103-0020022130322213)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-011.md#canonical-0130220100220230-3121123301113020-2302321022220303-0123123132122233-0220021333320201-3322321310102330-0320223123003030-1103200022300321)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present

<a id="canonical-3103001233332302-2321000103032021-3332230123301020-1330022220011232-0303313000213321-3111231033200112-1011023132200230-1320331102223013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-1120200320013201-2122121033100300-1033110233302121-3321310232112333-1332103231321230-3230331133130212-3101012322133113-1010213203022122"></a>

## Direct properties — check_present / 211333220101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330301331020210-1210123023203131-3231031311320020-0212303211220223-3003012302003303-3222030010220020-3303223130221313-1323312322233203"></a>

## Next pages — check_present / 211333220101 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-011.md#canonical-0130220100220230-3121123301113020-2302321022220303-0123123132122233-0220021333320201-3322321310102330-0320223123003030-1103200022300321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2222322013313020-0312211232300320-3103122221102233-2230133301322103-0200120220011002-2013021033213100-0233101301323221-3031123110012032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220300033123133-0330101102103003-1333221122002000-0021300201130210-3201320012102131-0322030122133003-2301132021223310-2232310201331230"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item — item / 312100101020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--reference--group-011.md#canonical-2123101112110001-2230223232113311-3200210313122201-3023210012031221-2033032323201332-0233120301211233-3302003023321332-3031131301202030)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--reference--group-011.md#canonical-2020001123330232-1331033320103030-2100212231130230-2211020212323023-2231323333310011-1130221220222202-1100231030232103-0020022130322213)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-011.md#canonical-0130220100220230-3121123301113020-2302321022220303-0123123132122233-0220021333320201-3322321310102330-0320223123003030-1103200022300321)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item

<a id="canonical-0202122300121202-2230201332310123-1321313033301203-2032011323231221-3301132111232211-0213010231202102-1201131201013023-0123100320110311"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231323133201120-2322000212002332-2030333111100103-3223123202213310-1211120123010011-2132113320133220-1102322100110221-3332113321213232"></a>

## Direct properties — item / 312100101020 / 3

<a id="canonical-2331011311311110-3123111222121010-2323332122032313-2000303013003033-1032132031322223-1200213212103330-2221202133110022-1300030022000312"></a>

<a id="canonical-2033100022231322-1323023113132030-1120211103032220-3322033123210300-3133132222321300-1300110203120031-1133110323202012-3232133120223211"></a>

## exact_values property — item / 312100101020 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2202302110321130-0321023030111230-2030331220003212-1103203330000021-1332010020333332-2113010332012111-1010310323310130-2011121200130230"></a>

<a id="canonical-2000303001302032-0203102213331311-0213030011303203-3032131022311330-3212121113103313-2130203123022102-2011303123033233-3020113033203233"></a>

## regex_values property — item / 312100101020 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3213123331200210-1211001201203303-1233210001230103-1122321120030222-3103310112232133-2232221031131222-0302000123030211-3110113231332020"></a>

<a id="canonical-1111330112133032-0223223120333102-2032133333331302-1123303213021301-1021301133220102-0001010133002321-2302321012023200-2011013133033303"></a>

## transformers property — item / 312100101020 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0332332022233120-1210101303033130-1102231102033330-1130331233301132-3312220320112011-0013120002001000-1330012210011203-3223323023100332"></a>

## Next pages — item / 312100101020 / 7

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--reference--group-011.md#canonical-0130220100220230-3121123301113020-2302321022220303-0123123132122233-0220021333320201-3322321310102330-0320223123003030-1103200022300321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322331020002122-0012221131111321-2322122020220311-3321033200310221-1232000120330033-2112321312131120-0120322222220103-1011333121310132"></a>

## bot_defense.policy.protected_app_endpoints — protected_app_endpoints / 011012201113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- bot_defense.policy.protected_app_endpoints

<a id="canonical-3122012111022201-0013313132031322-3323201303202211-2210110202222223-3112213100302320-3032210100003330-0322020232303102-1230001230323133"></a>

Type: `"object"`. list nested block, Optional.

List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32
endpoints per LB' after 4 LBs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("http_methods"),
  validators.ConflictingListObjectAttributes("allow_good_bots",
    "mitigate_good_bots"),
  validators.ConflictingListObjectAttributes("any_domain",
    "domain"),
  validators.ConflictingListObjectAttributes("flow_label",
    "undefined_flow_label"),
  validators.ConflictingListObjectAttributes("mobile",
    "web"),
  validators.ConflictingListObjectAttributes("mobile",
    "web_mobile"),
  validators.ConflictingListObjectAttributes("web",
    "web_mobile")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
protected_app_endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301332111013302-0332011203331013-0303220311123010-3313100002020210-0300002122200131-1120031032020032-2311230333003102-1212210021232000"></a>

## Direct properties — protected_app_endpoints / 011012201113 / 3

- [allow_good_bots](resources--http_loadbalancer--reference--group-011.md#canonical-1020312120302132-3022030103322213-3100322110233213-2122003033310021-0012313201123030-3033300203101301-0123132101123101-0221013130110003): complete subsection reference.

- [any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-1201322113311202-0002112110020122-2313113021202221-0023233220231122-2331333202131322-1121122310223233-1013231131113203-1200322213302013): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-011.md#canonical-1312130211202132-3212312301322212-1033322302133003-2031220203131211-0303200310203120-1302331210300113-1020132221310131-0012203132213203): complete subsection reference.

- [flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-012.md#canonical-3011003033310101-0122020002121002-3031202212302210-0021210121111322-1202003203203333-3130233012020232-2110123113302102-1102220303022203): complete subsection reference.

<a id="canonical-1033110211130302-1303120301301220-1021310301121301-1222003332033113-0322130203010112-3030331103333221-3012222020213000-1213333131010303"></a>

<a id="canonical-0021222313303012-2120021301130330-0212033210033231-2110332110302111-0232333112003102-3033221121103011-2311323201312221-0030133333100230"></a>

## http_methods property — protected_app_endpoints / 011012201113 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Upstream description:

List of HTTP methods.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-012.md#canonical-2023212323122011-1101323330132021-1310011202223312-0131122130110010-0101330133311212-3003232332003321-2322011030222200-1132132132002122): complete subsection reference.

- [mitigate_good_bots](resources--http_loadbalancer--reference--group-012.md#canonical-3200332103111311-3331103321320021-3300022321112003-2220210032301110-3121012100030020-3110320100033110-2213130032312123-2312232022333022): complete subsection reference.

- [mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331): complete subsection reference.

- [mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0200310013012011-1313310020131023-0203321100301133-2101212302102010-3320231321202232-3031131303020130-0120030303011112-1220331131110320): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-013.md#canonical-3132020021201300-2213301200112303-1020320300322230-3032200333333311-3311101002332123-1010201023012022-0122220101130023-1320121003331033): complete subsection reference.

<a id="canonical-2231221210102333-1021330211133211-3120201111332003-0102000303303320-2110311002112010-3321031001032120-3312122010123203-2011311323123012"></a>

<a id="canonical-2123023222333133-3332123302110312-1110133311213331-0100011202121213-0003213120330020-2310120213201300-1202012200330113-1101021201032021"></a>

## protocol property — protected_app_endpoints / 011012201113 / 5

Type: `"string"`. Optional.

\[Enum: BOTH|HTTP|HTTPS\] SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for
HTTPS:// or HTTP://. - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only.
Possible values are \`BOTH\`, \`HTTP\`, \`HTTPS\`. Defaults to \`BOTH\`.

Upstream description:

SchemeType is used to indicate URL scheme.

&#8203;- BOTH: BOTH

URL scheme for HTTPS:// or HTTP://. &#8203;- HTTP: HTTP

URL scheme HTTP:// only. &#8203;- HTTPS: HTTPS

URL scheme HTTPS:// only.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BOTH",
    "HTTP",
    "HTTPS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BOTH",
  "enum": [
    "BOTH",
    "HTTP",
    "HTTPS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [query_params](resources--http_loadbalancer--reference--group-013.md#canonical-1013232023221202-0101200322012222-1110023122333121-1203113113021223-0233213112333213-0331320110123211-3111131201203220-3313320112231132): complete subsection reference.

- [undefined_flow_label](resources--http_loadbalancer--reference--group-013.md#canonical-1203210321222120-0033120203013313-2131212022012032-1233321301313330-3120311213033232-3302302000221200-3023302011011311-0232031130132223): complete subsection reference.

- [web](resources--http_loadbalancer--reference--group-013.md#canonical-0130020123113200-0230312201110222-1333301202301130-3302211302301012-2231221131222131-3321323201300333-0101100233311213-2020002023103320): complete subsection reference.

- [web_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-3101313121221202-1301032013003023-1220301222330001-2201122130031032-2313210123222031-3023312033232202-2101213232211113-0202330221311302): complete subsection reference.

<a id="canonical-3322110211022210-1211003330321020-0212113031033333-2032233210311332-1211103011303101-2002120220011200-2032003300033111-0132231201132333"></a>

## Next pages — protected_app_endpoints / 011012201113 / 6

- [bot_defense.policy.protected_app_endpoints.allow_good_bots](resources--http_loadbalancer--reference--group-011.md#canonical-1020312120302132-3022030103322213-3100322110233213-2122003033310021-0012313201123030-3033300203101301-0123132101123101-0221013130110003)
- [bot_defense.policy.protected_app_endpoints.any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-1201322113311202-0002112110020122-2313113021202221-0023233220231122-2331333202131322-1121122310223233-1013231131113203-1200322213302013)
- [bot_defense.policy.protected_app_endpoints.domain](resources--http_loadbalancer--reference--group-011.md#canonical-1312130211202132-3212312301322212-1033322302133003-2031220203131211-0303200310203120-1302331210300113-1020132221310131-0012203132213203)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-012.md#canonical-3011003033310101-0122020002121002-3031202212302210-0021210121111322-1202003203203333-3130233012020232-2110123113302102-1102220303022203)
- [bot_defense.policy.protected_app_endpoints.metadata](resources--http_loadbalancer--reference--group-012.md#canonical-2023212323122011-1101323330132021-1310011202223312-0131122130110010-0101330133311212-3003232332003321-2322011030222200-1132132132002122)
- [bot_defense.policy.protected_app_endpoints.mitigate_good_bots](resources--http_loadbalancer--reference--group-012.md#canonical-3200332103111311-3331103321320021-3300022321112003-2220210032301110-3121012100030020-3110320100033110-2213130032312123-2312232022333022)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331)
- [bot_defense.policy.protected_app_endpoints.mobile](resources--http_loadbalancer--reference--group-013.md#canonical-0200310013012011-1313310020131023-0203321100301133-2101212302102010-3320231321202232-3031131303020130-0120030303011112-1220331131110320)
- [bot_defense.policy.protected_app_endpoints.path](resources--http_loadbalancer--reference--group-013.md#canonical-3132020021201300-2213301200112303-1020320300322230-3032200333333311-3311101002332123-1010201023012022-0122220101130023-1320121003331033)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-013.md#canonical-1013232023221202-0101200322012222-1110023122333121-1203113113021223-0233213112333213-0331320110123211-3111131201203220-3313320112231132)
- [bot_defense.policy.protected_app_endpoints.undefined_flow_label](resources--http_loadbalancer--reference--group-013.md#canonical-1203210321222120-0033120203013313-2131212022012032-1233321301313330-3120311213033232-3302302000221200-3023302011011311-0232031130132223)
- [bot_defense.policy.protected_app_endpoints.web](resources--http_loadbalancer--reference--group-013.md#canonical-0130020123113200-0230312201110222-1333301202301130-3302211302301012-2231221131222131-3321323201300333-0101100233311213-2020002023103320)
- [bot_defense.policy.protected_app_endpoints.web_mobile](resources--http_loadbalancer--reference--group-013.md#canonical-3101313121221202-1301032013003023-1220301222330001-2201122130031032-2313210123222031-3023312033232202-2101213232211113-0202330221311302)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1020312120302132-3022030103322213-3100322110233213-2122003033310021-0012313201123030-3033300203101301-0123132101123101-0221013130110003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202132030112110-2130301033012201-3302210313111012-0321322323233031-3323330321102120-1021100133033213-0213333300103201-0200010330202300"></a>

## bot_defense.policy.protected_app_endpoints.allow_good_bots — allow_good_bots / 303022233210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.allow_good_bots

<a id="canonical-1310120223301222-2132112010203210-0000022322003002-0310121100202123-1111201131203120-3113123023320233-2312130121123322-0000331121003232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for allow good bots.

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
allow_good_bots = {}
```

<a id="canonical-1332321230303303-1030020032303103-1320011122021001-2210313230220102-1100302001323001-3012032132212332-1132330000113210-3000200230112032"></a>

## Direct properties — allow_good_bots / 303022233210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132133222303002-0302330012230033-0223231202333310-3323102110333000-2320302333101031-2031331202313320-0121101000100213-3310002120201132"></a>

## Next pages — allow_good_bots / 303022233210 / 4

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1201322113311202-0002112110020122-2313113021202221-0023233220231122-2331333202131322-1121122310223233-1013231131113203-1200322213302013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220001201013133-1132113122230002-3232031002200311-0013320312213000-2212113203003223-1003133331122313-0131320213232033-1031312210203213"></a>

## bot_defense.policy.protected_app_endpoints.any_domain — any_domain / 112002120223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.any_domain

<a id="canonical-1033312211011131-1122032112000010-3132233332010133-3030303020313022-1220331203222332-0332320131230312-3123031001300201-2011121232302013"></a>

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
any_domain = {}
```

<a id="canonical-1233320102022000-1310100003320302-0000110320110332-2111302311303310-1011222031231013-3320001011112011-0023323211011331-3313122132032331"></a>

## Direct properties — any_domain / 112002120223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031233210011133-1213131021113213-0100332222131001-1013103002033220-1301203223133002-1000001331203122-0201133123021230-2013303013100133"></a>

## Next pages — any_domain / 112002120223 / 4

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1312130211202132-3212312301322212-1033322302133003-2031220203131211-0303200310203120-1302331210300113-1020132221310131-0012203132213203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232103010301233-3023103220102332-2310012112020233-3123003131022101-2221113300001321-0212322003132003-2302023123220212-0002003001300202"></a>

## bot_defense.policy.protected_app_endpoints.domain — domain / 030212222021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.domain

<a id="canonical-3312320033120230-2330033230020223-2121210233101233-2310310102311203-2201022002103030-0301022010121302-2000201232233320-1130112003003103"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2111322232310321-1122331322321221-0002332130332210-3031300313132101-3103012000102021-0102110030201231-0131231102320311-3010203230210310"></a>

## Direct properties — domain / 030212222021 / 3

<a id="canonical-1123023123133302-3202220202120302-3102221320131023-2232333303030202-2010310220331212-3031312112133321-2212011202210121-1202000011202130"></a>

<a id="canonical-0012210022300121-1221232331333031-0021130021010312-3122101201203023-0201132010121030-0330020022012112-2202022022303231-1320310122102201"></a>

## exact_value property — domain / 030212222021 / 4

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2302131001103212-1203330233113222-3313011312212010-0333213200221000-2103130211102303-0112301003301213-1330303131313212-1312122232002031"></a>

<a id="canonical-0112211200132300-2113221003201111-3330232022220012-3210312111222232-2332321213022020-0233131313110202-2333023031022132-0022133311103020"></a>

## regex_value property — domain / 030212222021 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1323003112100211-2323201222310023-3321021103320112-2230303331003112-0221022032312330-0230203323001012-3012012312310202-3011110112301202"></a>

<a id="canonical-0312123230310000-1202100020220003-3233222022031122-2321010202110323-1010230302113332-3100132223120030-2300222103123030-2221010321020022"></a>

## suffix_value property — domain / 030212222021 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3223203030303111-3112113120021012-3011322012133022-2230220333100022-3122132113113222-3332000011231200-1330200121211213-2211232113211201"></a>

## Next pages — domain / 030212222021 / 7

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231223323231111-3233130003111301-1233230202102022-2011323113011213-2221000200232232-0312130022203200-2311001023232122-0233031130322001"></a>

## bot_defense.policy.protected_app_endpoints.flow_label — flow_label / 332012320020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.flow_label

<a id="canonical-0323032312113133-2022032203013122-0130102223023310-0022103013200001-0110122332201023-1131022203203023-2031222113223310-1213120003112300"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("account_management",
    "authentication"),
  validators.ConflictingObjectAttributes("account_management",
    "financial_services"),
  validators.ConflictingObjectAttributes("account_management",
    "flight"),
  validators.ConflictingObjectAttributes("account_management",
    "profile_management"),
  validators.ConflictingObjectAttributes("account_management",
    "search"),
  validators.ConflictingObjectAttributes("account_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("authentication",
    "financial_services"),
  validators.ConflictingObjectAttributes("authentication",
    "flight"),
  validators.ConflictingObjectAttributes("authentication",
    "profile_management"),
  validators.ConflictingObjectAttributes("authentication",
    "search"),
  validators.ConflictingObjectAttributes("authentication",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("financial_services",
    "flight"),
  validators.ConflictingObjectAttributes("financial_services",
    "profile_management"),
  validators.ConflictingObjectAttributes("financial_services",
    "search"),
  validators.ConflictingObjectAttributes("financial_services",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("flight",
    "profile_management"),
  validators.ConflictingObjectAttributes("flight",
    "search"),
  validators.ConflictingObjectAttributes("flight",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("profile_management",
    "search"),
  validators.ConflictingObjectAttributes("profile_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("search",
    "shopping_gift_cards")}
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
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

Terraform syntax:

```terraform
flow_label {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200022001231102-2203002332313201-3101131002203300-1212222130022321-2020233033130110-3130231013223213-2002031232131223-1321212102331021"></a>

## Direct properties — flow_label / 332012320020 / 3

- [account_management](resources--http_loadbalancer--reference--group-011.md#canonical-0310101322123011-2311232002323331-2011130102302101-2300023323232113-3011011022101212-2300230213011101-1221102223323312-0123320212021332): complete subsection reference.

- [authentication](resources--http_loadbalancer--reference--group-011.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302): complete subsection reference.

- [financial_services](resources--http_loadbalancer--reference--group-012.md#canonical-2313330222120110-2222031320011333-2210303331220110-1210112000113221-0120310301320110-3320203021001300-3323133332212333-0032021010221112): complete subsection reference.

- [flight](resources--http_loadbalancer--reference--group-012.md#canonical-2131020221010102-2012122000002123-3132321003203233-1323312223200203-3033213210222121-1123012101220330-2201203221233110-3221101122221132): complete subsection reference.

- [profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021): complete subsection reference.

- [search](resources--http_loadbalancer--reference--group-012.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121): complete subsection reference.

- [shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220): complete subsection reference.

<a id="canonical-3001002210112023-1213013322302203-0331213113021102-1332233312000301-3203022333221131-0211021321310201-0331003010101331-2301022200221033"></a>

## Next pages — flow_label / 332012320020 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--http_loadbalancer--reference--group-011.md#canonical-0310101322123011-2311232002323331-2011130102302101-2300023323232113-3011011022101212-2300230213011101-1221102223323312-0123320212021332)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--http_loadbalancer--reference--group-012.md#canonical-2313330222120110-2222031320011333-2210303331220110-1210112000113221-0120310301320110-3320203021001300-3323133332212333-0032021010221112)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](resources--http_loadbalancer--reference--group-012.md#canonical-2131020221010102-2012122000002123-3132321003203233-1323312223200203-3033213210222121-1123012101220330-2201203221233110-3221101122221132)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0310101322123011-2311232002323331-2011130102302101-2300023323232113-3011011022101212-2300230213011101-1221102223323312-0123320212021332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222020202133123-3211221132300023-3102311023333002-2012033032311112-0010321030330030-2330003132201012-1030312122123001-1122210031230003"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management — account_management / 213000110021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management

<a id="canonical-1001303332202133-1301130232113330-1332122011133222-0020212203111102-3110121221112122-3312131231002233-3131302122133233-1002300200303310"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Account Management Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("create",
    "password_reset")}
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
  "x-ves-oneof-field-label_choice": "[\"create\",\"password_reset\"]"
}
```

Terraform syntax:

```terraform
account_management {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201022133023100-0021311333123023-2211331331310013-3301232213001331-1113313012111200-2112201323011222-3020213110110310-3320023210002301"></a>

## Direct properties — account_management / 213000110021 / 3

- [create](resources--http_loadbalancer--reference--group-011.md#canonical-2301020011313203-3212210013121110-1231032200222233-0103330212020133-1001030121111321-0223030213011232-0102201021123103-3330230323022223): complete subsection reference.

- [password_reset](resources--http_loadbalancer--reference--group-011.md#canonical-0122001002033033-0220122331200103-1032031030030033-3120113132220222-1000323220323300-2011023031301020-2123232322222330-0101231220210021): complete subsection reference.

<a id="canonical-2200012330331232-3221331221301133-3021222220323331-1330331312011020-2200120012101223-2313330022132221-0022211121020202-3030310310233032"></a>

## Next pages — account_management / 213000110021 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.create](resources--http_loadbalancer--reference--group-011.md#canonical-2301020011313203-3212210013121110-1231032200222233-0103330212020133-1001030121111321-0223030213011232-0102201021123103-3330230323022223)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset](resources--http_loadbalancer--reference--group-011.md#canonical-0122001002033033-0220122331200103-1032031030030033-3120113132220222-1000323220323300-2011023031301020-2123232322222330-0101231220210021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2301020011313203-3212210013121110-1231032200222233-0103330212020133-1001030121111321-0223030213011232-0102201021123103-3330230323022223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030133100320011-0032102300121130-3022103103112302-2103230102331133-2102221100100013-0221111303030033-2112003101010302-1222003011001132"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management.create — create / 202001110313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--http_loadbalancer--reference--group-011.md#canonical-0310101322123011-2311232002323331-2011130102302101-2300023323232113-3011011022101212-2300230213011101-1221102223323312-0123320212021332)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.create

<a id="canonical-0030323002010212-1132010201210130-2233110020212023-0220032302023122-0332131131021323-2011203123312301-0223121033032311-0002230131110331"></a>

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
create = {}
```

<a id="canonical-2320102103303110-2210110013210213-2200221222300013-1321200231103120-1120010311033302-0110222232000333-0100002031032223-0333032132331200"></a>

## Direct properties — create / 202001110313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122231221333013-3232200131123201-1220202232201310-3331002122233210-0033202333111130-1331311123313303-3001010212030120-2302120130021100"></a>

## Next pages — create / 202001110313 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--http_loadbalancer--reference--group-011.md#canonical-0310101322123011-2311232002323331-2011130102302101-2300023323232113-3011011022101212-2300230213011101-1221102223323312-0123320212021332)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0122001002033033-0220122331200103-1032031030030033-3120113132220222-1000323220323300-2011023031301020-2123232322222330-0101231220210021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322121231230102-1303330033322323-3032120330010331-3103211322322000-3303123123332032-0232120121121202-0301223132102231-1102103123012201"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset — password_reset / 202332301211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--http_loadbalancer--reference--group-011.md#canonical-0310101322123011-2311232002323331-2011130102302101-2300023323232113-3011011022101212-2300230213011101-1221102223323312-0123320212021332)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset

<a id="canonical-1011032000313301-1322303132013110-1232332133333013-3220132322302132-3013132332233011-2123113222130201-3330130003021213-1212030120130021"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for password reset.

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
password_reset = {}
```

<a id="canonical-3133030333313301-2123031010320320-3022013220020133-0013213102002020-0311310112031302-2130112303222002-3202312322310221-3333000030011111"></a>

## Direct properties — password_reset / 202332301211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120021322023101-3232332130102202-2232031131202113-1110202032113221-1123013112203322-3312032333223312-3023201210321132-2021132222001102"></a>

## Next pages — password_reset / 202332301211 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--http_loadbalancer--reference--group-011.md#canonical-0310101322123011-2311232002323331-2011130102302101-2300023323232113-3011011022101212-2300230213011101-1221102223323312-0123320212021332)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
