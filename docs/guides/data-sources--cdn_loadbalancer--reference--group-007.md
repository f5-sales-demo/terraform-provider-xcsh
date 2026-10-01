---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-3230300011333032-1212321203220323-3203123000120311-3223100100231320-1323100300101130-0001233030011201-0122023113201010-3100103230000103"></a>

## Next pages — bot_skip_processing / 333233313322 / 4

- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3222111102020121-2200312331132301-2313123020323212-2210321013123113-3130320020110303-2000213233000120-3333101322102201-2033103200133302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233330002321310-1010010133221302-2221123200333120-1013021123001101-0212220110331133-3110111123000121-3023312102303120-2223020011102212"></a>

## blocked_clients.http_header — http_header / 030311121020 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212)
- blocked_clients.http_header

<a id="canonical-3310002130213130-3232022220101001-1233311300313111-1300231232020021-2133200010012030-0131300122203311-3331112003301313-2200003333313000"></a>

Type: `"single"`. Computed.

Configuration parameter for http header.

Upstream description:

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

<a id="canonical-2110111111013102-2201123123310110-1321232230000320-3123021203002222-2003130130002232-3132212013011320-1302312123303130-0131133011102130"></a>

## Direct properties — http_header / 030311121020 / 3

- [headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3323103231233331-0320102122313133-0302330312020003-3200013223121301-1333323210211202-0000213201213002-3011120310102003-3100131023303111): complete subsection reference.

<a id="canonical-2120332032323001-1122210020221230-2003300130322121-3023311332103220-2330022000123130-3333000030121221-3313032110300101-0010203311122300"></a>

## Next pages — http_header / 030311121020 / 4

- [blocked_clients.http_header.headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3323103231233331-0320102122313133-0302330312020003-3200013223121301-1333323210211202-0000213201213002-3011120310102003-3100131023303111)
- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3323103231233331-0320102122313133-0302330312020003-3200013223121301-1333323210211202-0000213201213002-3011120310102003-3100131023303111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122321003032233-3012223303322333-0013031031001303-0031133312322021-3220301213121303-3100102013301321-1333023312021230-2210331121100212"></a>

## blocked_clients.http_header.headers — headers / 230000301223 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212)
- [blocked_clients.http_header](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3222111102020121-2200312331132301-2313123020323212-2210321013123113-3130320020110303-2000213233000120-3333101322102201-2033103200133302)
- blocked_clients.http_header.headers

<a id="canonical-0010311232300321-1000303023113311-2300130110000310-2101222030033311-3233020001033021-2122203022113100-3223120012230212-3021033211302311"></a>

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

<a id="canonical-1332330310011110-2203310132210132-1031002233133003-3203111020321010-1110100312202300-2121323310101021-3301020122220100-0102122232011200"></a>

## Direct properties — headers / 230000301223 / 3

<a id="canonical-3312111322200333-2120323331102220-2133023032310231-2120102100222210-2020210113213131-2200121011030300-0313122000321213-0212323303103303"></a>

<a id="canonical-1110100222112123-1021130130213121-1320202313120030-0331132320230230-3130213012232233-2022022213112301-0212231110211112-2113120213132010"></a>

## exact property — headers / 230000301223 / 4

Type: `"string"`. Computed.

Exclusive with \[presence regular expression\] Header value to match exactly.

Upstream description:

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

<a id="canonical-0131132110131200-1032202231012030-0001020202000321-1100330323120202-2023031213201223-1320122010211321-0213333111331110-1202210201232311"></a>

<a id="canonical-2200230131303321-2121020103333002-3233010122213132-3313312323121221-1123120213101303-2100013231321123-0310302032002322-2030113123301112"></a>

## invert_match property — headers / 230000301223 / 5

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

<a id="canonical-0322210211320211-0223130232120301-3331223223021020-0232103223121320-2312001120220220-2101322300322302-3233210123300322-1311201330021021"></a>

<a id="canonical-2323111213212133-3013123320033110-2101333121113033-3322301202000233-2220012011330100-0312132322230101-3233123320131320-0221001121232213"></a>

## name property — headers / 230000301223 / 6

Type: `"string"`. Computed.

Name. Name of the header.

Upstream description:

Name of the header.

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

<a id="canonical-2210020001000303-0010321320331330-1323312202013011-2001223302022320-0102112230213032-3020113011103222-2213011110113000-1230122332021311"></a>

<a id="canonical-1123313030033330-0113202303220100-3231202220230231-1110200311233210-3220222013330030-0211220200201113-3213102022030121-1032002232322223"></a>

## presence property — headers / 230000301223 / 7

Type: `"bool"`. Computed.

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

<a id="canonical-1031031033332000-1321203212200133-1031313311122323-0213201110223333-0333022122101001-0011130233202222-3300022301313301-1021233233033111"></a>

<a id="canonical-1301332112123020-3132110212030111-0131022003020310-2130233123320302-2300232011021003-2312222321023202-1000233333230213-1221310312010123"></a>

## regular expression property — headers / 230000301223 / 8

Type: `"string"`. Computed.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Upstream description:

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

<a id="canonical-1223212130323212-2303332230010111-3200101332301130-0210330332102233-3120021321230213-0211322211111112-2012221130320020-0123123313130122"></a>

## Next pages — headers / 230000301223 / 9

- [blocked_clients.http_header](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3222111102020121-2200312331132301-2313123020323212-2210321013123113-3130320020110303-2000213233000120-3333101322102201-2033103200133302)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0332122220010132-3003212310132001-3200233131130203-3002000320030333-3223100311203221-2013202331301022-2002302000000323-1210201223011100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031120113213233-2322322100133302-3320220023311230-0221000313021120-0313002210113012-2321132211232000-1200121003120223-2011233000231033"></a>

## blocked_clients.metadata — metadata / 323331321310 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212)
- blocked_clients.metadata

<a id="canonical-3011112210032213-0200001301201002-0002120300233200-2101101323110203-1313222300203022-0102031021022223-1303002010121332-3200232230111221"></a>

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

<a id="canonical-2021321221232311-3111311230011200-3033222221032311-3321303210133203-1132120200230233-2213202331032101-1033322212000200-3321101133233020"></a>

## Direct properties — metadata / 323331321310 / 3

<a id="canonical-3221303203300311-3310222013231211-2113210022020112-3333003232301330-0023231332112121-0232002003130203-2112303312112320-1202013111202022"></a>

<a id="canonical-2322013320103330-3302003101110302-3133331123120320-2203020013202203-1310303301130322-0010112213021022-1031200031332002-3123023233031012"></a>

## description_spec property — metadata / 323331321310 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3021103133232030-3120332120330301-1330333300231212-3012100102022322-0233320031231103-2113020032033112-0110233112120322-1111300001211002"></a>

<a id="canonical-0000322011211322-2130223120120120-2233003332012202-1322233121320022-0333233202112212-2031203123020313-2203333131220123-2133133013312132"></a>

## name property — metadata / 323331321310 / 5

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

<a id="canonical-2222221303020032-0113301000232122-2102102302001323-2023220021301121-1022002100103032-1000100001213302-1202012311210321-3003113232021203"></a>

## Next pages — metadata / 323331321310 / 6

- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2213100101223122-0130122203032003-2313212221101303-3022313132331032-3300322100123003-3312130112333203-2123201210320333-3001302303003303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221213232000301-2032013311201013-0121101312323002-3001010333010332-2300100021223313-3112011031030020-1310033023033323-2033111333311202"></a>

## blocked_clients.skip_processing — skip_processing / 200110212321 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212)
- blocked_clients.skip_processing

<a id="canonical-1222231102023020-2000131101313203-0200231020121332-2333121002331222-3200002310031233-0320331232322301-1210210101330221-3210233202133321"></a>

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

<a id="canonical-3131013030022221-0220211310203101-0221321302110313-0222210323320102-2233220011021112-1012000133210212-3232013310321003-0201222223222222"></a>

## Direct properties — skip_processing / 200110212321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123300332003201-3111011223002223-3300312330210211-1003111032322001-3332010213221110-1312003320210013-3012010120030313-3010303101320032"></a>

## Next pages — skip_processing / 200110212321 / 4

- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0130110100100120-2303231110321211-1013122300233201-1201111230323020-3033111300331133-3210003230300330-2000232003002333-2130012312022332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010321322033101-1333300112302310-3312203030320201-0232213211101132-3011003331022220-2313013111310000-2322011003330323-3113011013001013"></a>

## blocked_clients.waf_skip_processing — waf_skip_processing / 111213122003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212)
- blocked_clients.waf_skip_processing

<a id="canonical-2033030030301221-3122113203233233-3011013120202333-1010122012203011-0302031102232322-3323123130333010-2103010102213232-1120322110313200"></a>

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

<a id="canonical-3321002000011102-1200311113020022-0013310230210212-0303001212103100-0023032231130331-1131002102100331-0023112310022333-3320113322033033"></a>

## Direct properties — waf_skip_processing / 111213122003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023302301133221-3203331330113133-1210210123302013-2311132103120222-1200201211320002-2213122233013322-2023100130121112-3022311330033132"></a>

## Next pages — waf_skip_processing / 111213122003 / 4

- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203230021003310-2212322011211122-0213321110313201-3123231032302201-3122333110021302-3123113202300033-3322320210120032-0322313300011332"></a>

## bot_defense — bot_defense / 010123032230 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- bot_defense

<a id="canonical-2200110012102030-1101223103010130-1302102213101302-3003213131112323-1312333021000202-3112120003121220-1020101311333133-3010002001320212"></a>

Type: `"single"`. Computed.

Defines various configuration OPTIONS for Bot Defense Policy.

Upstream description:

This defines various configuration OPTIONS for Bot Defense Policy.

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

<a id="canonical-2130133220330211-3300331313310200-2002311111230112-2231030032231122-1130130012031121-1233321322212230-1033011210231121-1331123100200013"></a>

## Direct properties — bot_defense / 010123032230 / 3

- [disable_cors_support](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0220321332121111-1132233033210203-0120113230232023-3020203103011323-1012212203021000-2131322303313333-0023201201123030-0203330031230112): complete subsection reference.

- [enable_cors_support](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1101323212321033-3331011001021301-0113112011232131-3121022131023122-1202202022310203-2021121130021010-1330331331102312-1203023012203233): complete subsection reference.

- [policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111): complete subsection reference.

<a id="canonical-3101012110121222-2133102131333111-1313301000210213-2031310011221222-3312300003133300-3131003212101021-0223131023213032-2232311322222011"></a>

<a id="canonical-2022323302002320-2322112203200232-0112021033230330-3101002001311323-2232000111130210-0011320203102001-3330332323031033-1323310210223233"></a>

## regional_endpoint property — bot_defense / 010123032230 / 4

Type: `"string"`. Computed.

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

<a id="canonical-1131031120203320-2220001032201102-0320230111023111-3023111233320022-0113030302323121-1113300323000211-2202203211310123-1031303111122122"></a>

<a id="canonical-2312010103321121-1131001100233221-3131021313020322-1333320122011010-0323123111113231-3312300330032300-1101001012302313-1310110202221103"></a>

## timeout property — bot_defense / 010123032230 / 5

Type: `"number"`. Computed.

The timeout for the inference check, in milliseconds.

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

<a id="canonical-2221133023100123-1303220021032322-3220120303131131-2311120023111211-0030011021332122-1101312033221330-3111101223021002-0301022303030020"></a>

## Next pages — bot_defense / 010123032230 / 6

- [bot_defense.disable_cors_support](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0220321332121111-1132233033210203-0120113230232023-3020203103011323-1012212203021000-2131322303313333-0023201201123030-0203330031230112)
- [bot_defense.enable_cors_support](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1101323212321033-3331011001021301-0113112011232131-3121022131023122-1202202022310203-2021121130021010-1330331331102312-1203023012203233)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0220321332121111-1132233033210203-0120113230232023-3020203103011323-1012212203021000-2131322303313333-0023201201123030-0203330031230112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221001111221320-0020022003130221-3003201000102133-2003003201103132-2233101330322122-2000120220200310-0322220222121202-3312331213131202"></a>

## bot_defense.disable_cors_support — disable_cors_support / 213131120310 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- bot_defense.disable_cors_support

<a id="canonical-3122123232113000-2213122221203131-0100101002120202-1131233330301132-2331110101233001-2322233301222022-2012100320302312-1130000022121333"></a>

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

<a id="canonical-3231203131031023-2220131010033331-0301233321012200-1323222313000222-0102132333300303-2120311210111201-1233011122313020-2222031232320100"></a>

## Direct properties — disable_cors_support / 213131120310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201010120012012-0213003121123132-2002221331112133-2322033313330223-0311012013120300-1210122021200302-3220100230021033-3101002112331312"></a>

## Next pages — disable_cors_support / 213131120310 / 4

- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1101323212321033-3331011001021301-0113112011232131-3121022131023122-1202202022310203-2021121130021010-1330331331102312-1203023012203233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210323212211230-1020112121210213-0330301230121100-2133203232022133-1322033013032030-0331022312301010-1110311022220003-2302100223302223"></a>

## bot_defense.enable_cors_support — enable_cors_support / 203230310312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- bot_defense.enable_cors_support

<a id="canonical-3103132222232232-0233333210203321-1330011122322112-1322303231132201-2320133220102001-0103121112323222-2330132321223203-0011001010220320"></a>

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

<a id="canonical-3121001113321113-0000101330112133-0202132230021003-0332223110003211-2132211110212221-1011310123031101-0031111000203332-0000330002232211"></a>

## Direct properties — enable_cors_support / 203230310312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020211130030111-3011213321111002-1213332302032112-2231230313210303-2131211303221233-3011030312312320-0211211221313211-0113313300222011"></a>

## Next pages — enable_cors_support / 203230310312 / 4

- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103302110311230-1310330013321130-3330032312302331-0210322310112013-2121230221000232-1200320233031333-3000220022223110-1333200112123100"></a>

## bot_defense.policy — policy / 302022122013 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- bot_defense.policy

<a id="canonical-2320030302020113-3201031010311320-1020022233332120-0003312032201220-3111003020001202-0130112213122223-3130002133010013-1220010032231220"></a>

Type: `"single"`. Computed.

Defines various configuration OPTIONS for Bot Defense policy.

Upstream description:

This defines various configuration OPTIONS for Bot Defense policy.

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

<a id="canonical-1330121303011200-2320131011231132-1312323132011112-2323233331310130-1131033203223000-2311200210330122-2320021001233022-0002001013301211"></a>

## Direct properties — policy / 302022122013 / 3

- [disable_js_insert](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2331222130233333-3233120003332120-1203210331221010-1202002200131010-0103310113310101-1023311102332330-2000310123000223-3012211331111131): complete subsection reference.

- [disable_mobile_sdk](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2210230332322302-2113322223033313-1030112123331232-0010020011013110-0222030330103332-2010230221300103-1032031001022201-0331011123330322): complete subsection reference.

<a id="canonical-1011123212022030-3031132001202030-3121202221003002-2011130113213012-2210111003221321-2022301123001121-2001023010312231-2320303120321112"></a>

<a id="canonical-2230022123030012-1203011013132122-0232103223221010-2000100200321320-3020103020201231-2101301021211031-0312332330020122-2012220111331111"></a>

## javascript_mode property — policy / 302022122013 / 4

Type: `"string"`. Computed.

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

<a id="canonical-1301220130213313-0032111222102223-3111212333332321-0203102111110101-3132321030132230-3003201111003311-3102210100031323-2130010123102331"></a>

<a id="canonical-3011302033320313-2031023000002212-0032212030301111-3320033133032211-0030023233130030-0221120231133133-1022322120323013-3030202002203131"></a>

## js_download_path property — policy / 302022122013 / 5

Type: `"string"`. Computed.

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

- [js_insert_all_pages](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0101120322313112-2121032301302102-2131103332212003-0130020330220130-1121200132333113-0323311122012212-0131023132122330-1232021023111111): complete subsection reference.

- [js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2022211013032221-0322203320020120-2233313130120312-1220111100112121-1330312300220022-0033222022020111-2110001111121102-0111001301120220): complete subsection reference.

- [js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331): complete subsection reference.

- [mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303): complete subsection reference.

- [protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123): complete subsection reference.

<a id="canonical-1222033130000022-3011300000112223-2131121023223223-0211333112303032-0302111010311220-2111030012223121-0002002031112003-0030221112331332"></a>

## Next pages — policy / 302022122013 / 6

- [bot_defense.policy.disable_js_insert](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2331222130233333-3233120003332120-1203210331221010-1202002200131010-0103310113310101-1023311102332330-2000310123000223-3012211331111131)
- [bot_defense.policy.disable_mobile_sdk](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2210230332322302-2113322223033313-1030112123331232-0010020011013110-0222030330103332-2010230221300103-1032031001022201-0331011123330322)
- [bot_defense.policy.js_insert_all_pages](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0101120322313112-2121032301302102-2131103332212003-0130020330220130-1121200132333113-0323311122012212-0131023132122330-1232021023111111)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2022211013032221-0322203320020120-2233313130120312-1220111100112121-1330312300220022-0033222022020111-2110001111121102-0111001301120220)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2331222130233333-3233120003332120-1203210331221010-1202002200131010-0103310113310101-1023311102332330-2000310123000223-3012211331111131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123311100133120-0030123010231322-3102132003030120-1021201033112221-0133103301302123-1032222202213313-0132221011213321-1222321033121220"></a>

## bot_defense.policy.disable_js_insert — disable_js_insert / 310230130312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- bot_defense.policy.disable_js_insert

<a id="canonical-3300003122021331-2230023312201001-3201213321231130-2232222102032002-1123122023002012-1222123230031103-3203030022213131-2010030322120132"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2202300213200102-1021130031122322-2221110230212333-1202120202110132-1031201202312032-2023121033331123-2223333310203010-0302303030232121"></a>

## Direct properties — disable_js_insert / 310230130312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210322200230210-2133021112110111-3311200002211031-2131003030211320-0021233100211201-3220132222011101-2003201201003011-0100020113321333"></a>

## Next pages — disable_js_insert / 310230130312 / 4

- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2210230332322302-2113322223033313-1030112123331232-0010020011013110-0222030330103332-2010230221300103-1032031001022201-0331011123330322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133033202121121-2331131123310020-2320131113303230-2230031011331312-3213223323123012-0102230120032232-1012201230021221-1221223122331131"></a>

## bot_defense.policy.disable_mobile_sdk — disable_mobile_sdk / 023133202033 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- bot_defense.policy.disable_mobile_sdk

<a id="canonical-1323002213001003-3033102011100110-3232113213200212-0031200132112021-2313212220130103-1230203232221323-3131301033230000-0023333101323011"></a>

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

<a id="canonical-3302132320102311-0203231001121122-1103003022113130-1311301233123132-2333223310233112-2302201020233112-0300320121230110-3223200320121112"></a>

## Direct properties — disable_mobile_sdk / 023133202033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323102200133332-0330230232203211-2033203133300232-1022311001103023-3300022132223030-2022331220103223-1031311000003211-0021130101002221"></a>

## Next pages — disable_mobile_sdk / 023133202033 / 4

- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0101120322313112-2121032301302102-2131103332212003-0130020330220130-1121200132333113-0323311122012212-0131023132122330-1232021023111111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131001311012302-3020212321221202-1300031003002302-0003010003120200-3300212131010032-2023202301320212-2203321332103213-1102332110021110"></a>

## bot_defense.policy.js_insert_all_pages — js_insert_all_pages / 001013002210 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- bot_defense.policy.js_insert_all_pages

<a id="canonical-3232120300031120-1300233230000320-1000230131120223-2322120301313002-1020232231302123-2133221322131322-0123222313021011-0223000023223213"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3310313231032203-1131132312110323-3001302231032022-0231030031223132-1310333220212013-2110301102031111-2321022323232103-2223220103122022"></a>

## Direct properties — js_insert_all_pages / 001013002210 / 3

<a id="canonical-0011103323002103-2031230102232033-1010220130302223-0002102103100010-3033320320013131-1122231022313333-3303111302023212-3321303013321202"></a>

<a id="canonical-3211112210200120-0222121322300102-1103032333022203-2221223032230301-3211331010120201-2301112313131333-2130201001221331-3201100130120200"></a>

## javascript_location property — js_insert_all_pages / 001013002210 / 4

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

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

<a id="canonical-3032133021200000-2030212020333033-0022020322333012-2212311320232110-0221010130033112-1030123131121220-1023031322003103-0123313012202310"></a>

## Next pages — js_insert_all_pages / 001013002210 / 5

- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2022211013032221-0322203320020120-2233313130120312-1220111100112121-1330312300220022-0033222022020111-2110001111121102-0111001301120220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012221032332130-0212120210111303-2100321323323201-2122200101200222-0213310333123022-3312023111212021-1321231031323312-1032233122201230"></a>

## bot_defense.policy.js_insert_all_pages_except — js_insert_all_pages_except / 200120102113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- bot_defense.policy.js_insert_all_pages_except

<a id="canonical-1202100222221023-1003210212121212-0321210222120111-0312330313132030-1231332311030312-1122322223223302-3331322100202322-2002001112302322"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0302120322311130-2332332322013201-3211110201030102-2002222000031333-1110122221132220-0031033100223201-2311233302023002-1000232320212201"></a>

## Direct properties — js_insert_all_pages_except / 200120102113 / 3

- [exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310): complete subsection reference.

<a id="canonical-2013300112311010-0122100023211202-3230322130002301-3020210002212030-1210213203103121-1123330001000223-2101231320212102-0021110002300210"></a>

<a id="canonical-1133313112330003-1333101222322023-0323231231003012-0111123322002010-2123012231131100-1000200312200030-2211213303121123-3022303301033301"></a>

## javascript_location property — js_insert_all_pages_except / 200120102113 / 4

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

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

<a id="canonical-1121131312121021-0213120102223033-3210021131003233-2232330212232300-0232233123232001-3132313031032110-0013331330231031-0110203321221110"></a>

## Next pages — js_insert_all_pages_except / 200120102113 / 5

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332123021132101-0133023113321220-3100001010023312-2031103102310101-0211001010021003-0020131002203220-2332002210101020-1103211013300122"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list — exclude_list / 022003233212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2022211013032221-0322203320020120-2233313130120312-1220111100112121-1330312300220022-0033222022020111-2110001111121102-0111001301120220)
- bot_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-1312222123313202-0100030220120332-0012312333302311-3313320322031030-3002101320212312-0201332203032313-3310330010111312-2013131003231023"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

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

<a id="canonical-1120002222312131-3133122101133032-2023212122212013-2001100103032200-2001202220022122-0112320201123211-2122230101021110-0030022203023011"></a>

## Direct properties — exclude_list / 022003233212 / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3221022013311210-1303333312110100-3231202033012211-1212032101120010-3012033110122123-2112021112133103-3031121010220331-2300033211120132): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3302131133301200-2312333313230321-0231022102310322-2030333233012300-0330222233233023-1030020001221232-2130212320321310-1200012022223021): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3002001232030111-2330311122330111-1001301203133100-0311232213323021-1222202301002021-2123322010233023-1200033010220110-1013223312300122): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0303132031033301-0102301230131313-1233011001011330-3221110132312121-1303013211012200-1331301112222131-2300311201113312-3332201113113303): complete subsection reference.

<a id="canonical-1122020222033230-3220131012311013-2001220231303320-3111203310312103-2211220211100000-3321031120132013-3011223331330212-1233220311223102"></a>

## Next pages — exclude_list / 022003233212 / 4

- [bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3221022013311210-1303333312110100-3231202033012211-1212032101120010-3012033110122123-2112021112133103-3031121010220331-2300033211120132)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3302131133301200-2312333313230321-0231022102310322-2030333233012300-0330222233233023-1030020001221232-2130212320321310-1200012022223021)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3002001232030111-2330311122330111-1001301203133100-0311232213323021-1222202301002021-2123322010233023-1200033010220110-1013223312300122)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.path](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0303132031033301-0102301230131313-1233011001011330-3221110132312121-1303013211012200-1331301112222131-2300311201113312-3332201113113303)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2022211013032221-0322203320020120-2233313130120312-1220111100112121-1330312300220022-0033222022020111-2110001111121102-0111001301120220)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3221022013311210-1303333312110100-3231202033012211-1212032101120010-3012033110122123-2112021112133103-3031121010220331-2300033211120132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113110020202000-1000133200322121-3231113001220212-3131200102110310-0022132120022201-1120331300002020-0012232102312011-0312331003130023"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain — any_domain / 333021313203 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2022211013032221-0322203320020120-2233313130120312-1220111100112121-1330312300220022-0033222022020111-2110001111121102-0111001301120220)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-0230031210110011-3222000202231331-3302213031220022-0321100121311020-3001121012111331-2322101031112303-1210021212202123-2122213113020220"></a>

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

<a id="canonical-3202101132123123-3222221013021130-2220113200030230-2033211210012211-0232332003113132-0203213111201301-0221202132133321-3023001231010321"></a>

## Direct properties — any_domain / 333021313203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211223220022311-1120010210313203-3032301131322331-1033223020213133-1121222200311321-3013121322022201-3002101003103310-0121311102323002"></a>

## Next pages — any_domain / 333021313203 / 4

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3302131133301200-2312333313230321-0231022102310322-2030333233012300-0330222233233023-1030020001221232-2130212320321310-1200012022223021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121103122313132-0000030011002021-3301020123003101-0321130212210332-0121031323200132-1330321322220320-3321220220310031-1100311033033020"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.domain — domain / 231131301113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2022211013032221-0322203320020120-2233313130120312-1220111100112121-1330312300220022-0033222022020111-2110001111121102-0111001301120220)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-2303131020031210-3232303331031133-2302011203331021-1020210302102120-0121133000013113-3002100303220010-3223213022322111-1313111313021032"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

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

<a id="canonical-3223312201103312-3213300201321111-2120023002000230-0303220310110302-0021230210122013-2323320330301130-0101020000300012-2322300210333312"></a>

## Direct properties — domain / 231131301113 / 3

<a id="canonical-2330130332222131-0233312222212031-1031212322311131-1222323021103020-1122003310030200-3200223201233010-2131320230322032-1232133112200130"></a>

<a id="canonical-0201113113300031-2020200001023032-3130202300123002-2001021130320202-3333230220202023-0011131132210123-3232233031230122-1331021313312311"></a>

## exact_value property — domain / 231131301113 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

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

<a id="canonical-1312102301220300-2222002113221030-0131230101120312-2110121201022230-0222311320023222-1311232111300233-2010232312302112-1332211303101332"></a>

<a id="canonical-0102221301112333-1330000012011202-1013331003323202-0202221010100312-1230113133323321-3022310130211130-3301310232011033-0210102103131311"></a>

## regex_value property — domain / 231131301113 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-2120231013122203-2213000233213133-2101003010133323-3010303203110311-3010032222011322-3331031100022032-2102320013221212-0011000033200111"></a>

<a id="canonical-3210330021213031-3311031113112001-1320031210001010-0020121132322312-3133002120312011-2122333131221002-2032131110132200-0213132120130331"></a>

## suffix_value property — domain / 231131301113 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
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

<a id="canonical-3133213223012202-2021020211110230-3200302330033221-1001323023321001-0101302031211211-1301110030131101-0023212033202020-2131313113130103"></a>

## Next pages — domain / 231131301113 / 7

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3002001232030111-2330311122330111-1001301203133100-0311232213323021-1222202301002021-2123322010233023-1200033010220110-1013223312300122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313200112232121-3300122331020022-2221313200103131-2033020123232210-1010133010002320-0113210100300122-2112030301311323-1230020003003213"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata — metadata / 200321320333 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2022211013032221-0322203320020120-2233313130120312-1220111100112121-1330312300220022-0033222022020111-2110001111121102-0111001301120220)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-1020213201003302-0320011123322032-0310021233120132-3313331100310013-1133211311311220-1120011211102233-3221123110331023-3033333112312000"></a>

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

<a id="canonical-3132331312022333-0302010301113002-2132230021302130-3321003112102100-2323131011021101-3011322011031012-1032022121221201-2201101320011230"></a>

## Direct properties — metadata / 200321320333 / 3

<a id="canonical-1002221332113000-1330332111212132-3010030203101212-0233220013000130-0303320202023003-0000022031222223-0221122120201123-2220213123032000"></a>

<a id="canonical-1021331323310213-0112332131133321-1223331220121302-2012021320221132-1021121112303012-0233010301223010-3031212231021001-2003213012320030"></a>

## description_spec property — metadata / 200321320333 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3133331222312303-3231011103031112-3223222110102100-1022323020332021-2021121222131231-1003013122021211-0330113130330332-3101211202003113"></a>

<a id="canonical-3121110111333033-2112133330221021-1232332022103123-1100023310203100-3222302110003330-1111000212022211-2213130121110230-0010332223121122"></a>

## name property — metadata / 200321320333 / 5

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

<a id="canonical-0203120012330133-0011320012211213-0010032212231230-2221023200030321-0132112203000213-2023023013000030-0100112031033203-2110022030221323"></a>

## Next pages — metadata / 200321320333 / 6

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0303132031033301-0102301230131313-1233011001011330-3221110132312121-1303013211012200-1331301112222131-2300311201113312-3332201113113303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332333020121221-0211223210202330-0313122230133030-1012103311221321-3332001201203113-3022133013031311-0312101033102203-1333123131320103"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.path — path / 302303330223 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2022211013032221-0322203320020120-2233313130120312-1220111100112121-1330312300220022-0033222022020111-2110001111121102-0111001301120220)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-1332122200323000-2322212020132000-1321122111101230-1102002332113111-0231133121012212-1200321132221333-2312232233032122-1012320110230020"></a>

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

<a id="canonical-3322311332102230-1203112132023102-2212133321300130-2013111010022310-1300221001100021-1031312311130100-0030132331100131-2303323223022211"></a>

## Direct properties — path / 302303330223 / 3

<a id="canonical-1311003112001103-2001211122330020-3030331231112223-3112020323320313-0131323013321133-1332303133132130-0330323130010332-0321221120012323"></a>

<a id="canonical-2213123013011303-3201322110000230-1103023332210120-0221103101100131-3100013121010200-2330031103121212-3233221303120202-2230033112310220"></a>

## path property — path / 302303330223 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

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

<a id="canonical-2300023200332301-3323023220330033-2101320313310301-1210122103121300-1130112130110111-2302212101211131-2212032133133223-2210012002012302"></a>

<a id="canonical-3123103333013230-1212313102302222-3013103021022100-2323323012231113-2333121310132020-1201131311203223-3300033021231221-1222232112010120"></a>

## prefix property — path / 302303330223 / 5

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

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

<a id="canonical-1202202321203121-3032300333220332-1021210303212100-1131010303333100-0122230332103210-3010300321230233-0101221230110202-3331213002323303"></a>

<a id="canonical-0211311213032112-2120001120013223-3210312111010203-3303220023122022-2222033002231030-0132103331110203-1001321303200102-1010231102320302"></a>

## regular expression property — path / 302303330223 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-1013003210133302-1232121131231230-3201100301302302-1213003320200313-0010020203113002-1102110022111003-2231021030320232-1221330122300300"></a>

## Next pages — path / 302303330223 / 7

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301021302112330-2203011121200120-2211332130310032-0010333003131133-1232102102123102-1101001113312010-1222121111002103-1000133022223322"></a>

## bot_defense.policy.js_insertion_rules — js_insertion_rules / 022220113112 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- bot_defense.policy.js_insertion_rules

<a id="canonical-2031020211001333-0101031112333100-0102331001111233-1321001020213023-2212000222310122-3201110003023322-0102231122202110-0100130012330320"></a>

Type: `"single"`. Computed.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0012132123223332-3122123313102013-1301201013013220-0231132030121202-2211023200023013-2010130312130322-2112020233200330-1311301021312200"></a>

## Direct properties — js_insertion_rules / 022220113112 / 3

- [exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321): complete subsection reference.

- [rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203): complete subsection reference.

<a id="canonical-1332331013320110-1303121223031232-3323301213211001-1213021221213100-1103221320211331-2321300133312110-1011003203212103-2000211222013231"></a>

## Next pages — js_insertion_rules / 022220113112 / 4

- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012132012203101-0200133012212313-1221320303212232-2103230121131300-0032330330302202-3121031211320102-3130013232313122-2001220313201223"></a>

## bot_defense.policy.js_insertion_rules.exclude_list — exclude_list / 130303022101 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- bot_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-0321310300010320-2103013132212003-0013201333301131-1212222113312330-2313330311112001-1211110112003330-2131110133011201-2021103200001212"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

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

<a id="canonical-0123202320321322-0120232302231212-2223223223220122-1311032211210320-3300303120210203-0002331213231321-0132212033001332-2133301323221023"></a>

## Direct properties — exclude_list / 130303022101 / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1303010300323100-3202022113212012-0312201120103020-2003303323331302-0213232232213122-3330022022312320-3121120030112303-1032311302113203): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1201303333113203-2122223100332002-3231202110121231-0211123222003013-1111122133133223-0132303023010211-0220122303201231-0301223330331032): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2010123033223331-2020322103213330-2313232233301230-2010032011231200-2320012023222313-0132222230311120-1002211213031112-2001221121200300): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2221131121310021-2121220312021222-2103003303232102-1023032032313022-0321311030010023-1000332302000010-0331223111332311-2101101322211010): complete subsection reference.

<a id="canonical-1133110321202333-2002001130031032-0130312020203332-0233123322222333-2120232002211021-3203003121000032-1030111301200321-3002003211011132"></a>

## Next pages — exclude_list / 130303022101 / 4

- [bot_defense.policy.js_insertion_rules.exclude_list.any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1303010300323100-3202022113212012-0312201120103020-2003303323331302-0213232232213122-3330022022312320-3121120030112303-1032311302113203)
- [bot_defense.policy.js_insertion_rules.exclude_list.domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1201303333113203-2122223100332002-3231202110121231-0211123222003013-1111122133133223-0132303023010211-0220122303201231-0301223330331032)
- [bot_defense.policy.js_insertion_rules.exclude_list.metadata](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2010123033223331-2020322103213330-2313232233301230-2010032011231200-2320012023222313-0132222230311120-1002211213031112-2001221121200300)
- [bot_defense.policy.js_insertion_rules.exclude_list.path](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2221131121310021-2121220312021222-2103003303232102-1023032032313022-0321311030010023-1000332302000010-0331223111332311-2101101322211010)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1303010300323100-3202022113212012-0312201120103020-2003303323331302-0213232232213122-3330022022312320-3121120030112303-1032311302113203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323002211012021-0010130232130333-1021000302020002-3003202320112010-1003300030120323-0130111320132333-2000032033310132-3330023310323130"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.any_domain — any_domain / 020202003233 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321)
- bot_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-0103301030001003-1112103211031321-2110130211333231-1323211013211222-2122003323210331-3031133231123112-1112100301233233-2020032110232120"></a>

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

<a id="canonical-3213020021322301-3300231331110331-1213100230303302-0123232020300121-1333331102230223-1301203012202323-3132213210213113-2112203311310203"></a>

## Direct properties — any_domain / 020202003233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030032131310212-3133102313102021-3333322213231220-1222331112103120-1112002111033111-0021303222101012-3113133321303301-0123010300211102"></a>

## Next pages — any_domain / 020202003233 / 4

- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1201303333113203-2122223100332002-3231202110121231-0211123222003013-1111122133133223-0132303023010211-0220122303201231-0301223330331032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011202321233201-1221233130313122-1012223323101031-1011313112102200-2011102031323013-1112211232300233-3100212030010120-0201033203002202"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.domain — domain / 211333103113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321)
- bot_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-1213101213202120-2021233320231001-1001120312131133-2003213301301100-2313123331101233-3012103102123123-2222323023302022-3313122201031202"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

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

<a id="canonical-0212022211000233-1310130301312011-1112120210310212-1333003021021231-3033023122130112-1102013211101212-1012233121002313-3103021101122002"></a>

## Direct properties — domain / 211333103113 / 3

<a id="canonical-3301312232013001-2221222121103313-3233313231322300-3010003222001123-3322211122320210-0112120300221200-0331332131131313-0213111332201111"></a>

<a id="canonical-3103003311132322-2323223101320332-2003101330212321-3112211220102330-3331033331033113-1231300032201022-1222000313210031-3003202221131310"></a>

## exact_value property — domain / 211333103113 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

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

<a id="canonical-3010132213021321-3220000320322101-0112131330313232-1112133003213323-2021313221111313-0200202333333130-0122212003031110-2311323010030020"></a>

<a id="canonical-2202132211113332-0323000111202311-0012022122222002-3212010123210030-2033033103010211-3132232121112022-3001213232331003-2320221121120332"></a>

## regex_value property — domain / 211333103113 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-2301102131022332-0313333012323321-2123203200003201-2223331112310132-3313022130031130-1111132133312201-0333102202333012-1320222330101110"></a>

<a id="canonical-1011303030303022-1033203213012123-2022113231231303-3333001120213003-2222323110133222-0112221320223132-3320321210003330-3120003230001230"></a>

## suffix_value property — domain / 211333103113 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
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

<a id="canonical-0110330011333203-0031010110311012-0132021103023212-1301203212132133-0100112102201333-1210032012323201-2331323133230212-3211122013100321"></a>

## Next pages — domain / 211333103113 / 7

- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2010123033223331-2020322103213330-2313232233301230-2010032011231200-2320012023222313-0132222230311120-1002211213031112-2001221121200300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012300231320010-1003021220133302-1022210213232231-1220021000323112-2200203002102031-3033332300013023-2323202310310211-1320002311322010"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.metadata — metadata / 133023232313 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321)
- bot_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-1102313311101201-1111102331010002-0123230102220200-1020021220322330-3002322120030022-1301313223302130-3333320223102100-3231012030022112"></a>

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

<a id="canonical-2013312010310121-0222203132211203-0200231003110113-0333322133223130-2120201220321031-0230210120302012-2013310211213021-1033000221113303"></a>

## Direct properties — metadata / 133023232313 / 3

<a id="canonical-0202200130133231-0130303310313321-3001112132111303-1330102013011021-0120232030313000-1013331201111130-3313303302220203-3221112213002023"></a>

<a id="canonical-1321012111210213-3123121333021113-1213002200032302-1302302303312130-3330203222101132-2200303301303103-2220200013021023-3013203131011030"></a>

## description_spec property — metadata / 133023232313 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1202001003131203-1231012233321011-2303223231233203-1100022110122321-1120310002201202-1332100202300113-2031203001022010-2111123033303130"></a>

<a id="canonical-3202202000110032-3131220032101300-0301332020232103-0212133311302111-3002200213333210-3031231000302013-1300123303032010-0033323310302112"></a>

## name property — metadata / 133023232313 / 5

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

<a id="canonical-0012130201321131-1110001312211303-1112323313032201-3020110300120031-1020012002222023-2020330012200120-2301331102232222-1121000021231323"></a>

## Next pages — metadata / 133023232313 / 6

- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2221131121310021-2121220312021222-2103003303232102-1023032032313022-0321311030010023-1000332302000010-0331223111332311-2101101322211010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302003222002133-3200321003010131-2323122203323221-3002322100101323-3000300311332202-3132013233213100-2113300211210212-1112323022231230"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.path — path / 312323332330 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321)
- bot_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-3121112323033300-2102021020110010-0210100313213310-1131233213001001-2011111302303222-2101021132133010-2212120121102202-0300301002123021"></a>

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

<a id="canonical-1100300121201323-1110123003023112-2102303110010331-1201022113010021-1033312203000223-3200110231233013-3211221101233303-3320322220213210"></a>

## Direct properties — path / 312323332330 / 3

<a id="canonical-0021230210033310-3120301232022223-1123202010312120-3133301300132231-3212210301012121-3003123121200112-3330321021233221-2120320222001200"></a>

<a id="canonical-0120130221230120-3331103222123320-2020303223033001-2012032322303233-2310230113030102-3012122233321333-0201011231312020-2130333233022032"></a>

## path property — path / 312323332330 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

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

<a id="canonical-1021320132113112-0333100302210103-1331230301313020-2203233112232013-1201221032103112-3133132110122121-3312321313001202-3303022000233023"></a>

<a id="canonical-2102113110331300-0122122131222332-1021210323021332-2301031110131221-0222023021003012-1133012202323210-2222100101011110-0111033230300133"></a>

## prefix property — path / 312323332330 / 5

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

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

<a id="canonical-2123013221310130-2322233231201033-0003032302233203-0002132330102021-0023132302212033-1001131212332123-1200013030232312-0321101310122021"></a>

<a id="canonical-1331223321320223-0021212233323320-3003111332231330-3311120000232201-2110202101310222-2122332133200021-2332213021133020-1130313101211203"></a>

## regular expression property — path / 312323332330 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-2131203211020011-1330103201033033-1033223102013032-0220110022321123-3211111200130312-0110110210203311-3022120312323011-2101101011003232"></a>

## Next pages — path / 312323332330 / 7

- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0031310300120220-3232123032310120-3322233201223331-0103131110221111-3003233121100330-3121313333321330-3003310101020231-0003213330313321)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233132321100012-1211113131200131-3002303033220203-3231102302021131-0233101230022022-3311310112013031-3012311202100220-3232202201222223"></a>

## bot_defense.policy.js_insertion_rules.rules — rules / 032013023220 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- bot_defense.policy.js_insertion_rules.rules

<a id="canonical-0111001210212010-0103012321320001-0212033232231230-1230033323300103-0101000130333210-0031123002313201-2230213202001223-3220131203320300"></a>

Type: `"list"`. Computed.

Required list of pages to insert Bot Defense client JavaScript.

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

<a id="canonical-1002223202132020-0223132110033212-0230333311003322-1011201003131123-3223302201032310-2223022233133301-3013132023323201-0312221022001331"></a>

## Direct properties — rules / 032013023220 / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3033221120230023-2003032032110220-2121133021231222-2103312231132303-0300120302211022-3003202011203201-1011322302131033-3113331013013133): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3120310332302230-2022101100122101-1023223311311132-3231011323202100-2330133330200310-3322010020211011-0002322102330300-0322330010022223): complete subsection reference.

<a id="canonical-2232202132010011-0123311013333112-1320103001312331-3132220231110301-1310101310302320-3130303120202001-0030012002121100-1333102111201100"></a>

<a id="canonical-3122131310003210-3333021201302011-3030210211120320-1300023002012002-2022213120132233-1002311013223220-3311203321301233-1132222112212132"></a>

## javascript_location property — rules / 032013023220 / 4

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

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

- [metadata](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1320002021202303-1300330312202300-2213101032300021-2330111323130110-3220102210313203-2033213200201201-0010313221302112-3030032012032012): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001110231301133-3233002312320121-1323121031211213-1131301311230233-2220323210211320-0002232222312020-1120203313102001-2002212010321200): complete subsection reference.

<a id="canonical-3003331330323032-3022323300233020-1023203231120333-1330121221212220-0021321330300120-0031332300002023-2310133111202033-0332332000113313"></a>

## Next pages — rules / 032013023220 / 5

- [bot_defense.policy.js_insertion_rules.rules.any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3033221120230023-2003032032110220-2121133021231222-2103312231132303-0300120302211022-3003202011203201-1011322302131033-3113331013013133)
- [bot_defense.policy.js_insertion_rules.rules.domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3120310332302230-2022101100122101-1023223311311132-3231011323202100-2330133330200310-3322010020211011-0002322102330300-0322330010022223)
- [bot_defense.policy.js_insertion_rules.rules.metadata](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1320002021202303-1300330312202300-2213101032300021-2330111323130110-3220102210313203-2033213200201201-0010313221302112-3030032012032012)
- [bot_defense.policy.js_insertion_rules.rules.path](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001110231301133-3233002312320121-1323121031211213-1131301311230233-2220323210211320-0002232222312020-1120203313102001-2002212010321200)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3033221120230023-2003032032110220-2121133021231222-2103312231132303-0300120302211022-3003202011203201-1011322302131033-3113331013013133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121132323313231-1223201010223202-2000003323212022-0223111033333302-1101011011111032-0021022120323330-1220211023222122-1232212120230030"></a>

## bot_defense.policy.js_insertion_rules.rules.any_domain — any_domain / 012302022331 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203)
- bot_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-2131213112202030-3033113303202220-3212130232003330-0032013113022013-0022022221121321-2320332211333122-1210023322132203-3310030030331103"></a>

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

<a id="canonical-0012331020331231-1030302332101232-2233223121201001-3030013233002331-2221321310010301-3232032230323112-1000110120223330-0021132023002120"></a>

## Direct properties — any_domain / 012302022331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212220330323130-2100131000003211-3110301122210010-2323013013313030-0112202102233230-0100000131301222-2012312011002222-3330302232002203"></a>

## Next pages — any_domain / 012302022331 / 4

- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3120310332302230-2022101100122101-1023223311311132-3231011323202100-2330133330200310-3322010020211011-0002322102330300-0322330010022223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110130213111332-0100303203102302-3331200203003133-1103213233211131-3033003003221300-1001222003030100-1221333012102113-0221030332102303"></a>

## bot_defense.policy.js_insertion_rules.rules.domain — domain / 120123203312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203)
- bot_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-2222111010130012-2031111222312132-2331320322310303-2222221312032231-0123311110322102-0220323321200300-0133211103311012-3101213321321121"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

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

<a id="canonical-0300211321310311-3000000322212133-3103211233121033-3001013032202231-1231032101113303-2123310303302030-0222032201213303-0013310331301210"></a>

## Direct properties — domain / 120123203312 / 3

<a id="canonical-3102120033213113-0320333033201222-0213213333112312-1121303020100020-1020100133211202-2300033310302003-2331033311220231-0231022123100211"></a>

<a id="canonical-2102201310010102-2002231232332132-0220013211002332-0002203201333110-2301303331201012-2123200102000120-1200102231002011-3120301021120123"></a>

## exact_value property — domain / 120123203312 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

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

<a id="canonical-1131002121132012-1013011220121012-0311331032201331-0131300201320233-0221330113032003-1000211310120000-1000322110233212-1300033113210131"></a>

<a id="canonical-0011020200011101-0220023222011112-2121000032301332-1032010123011323-0112020113123031-1021013232222122-2132313112033021-0102100330323331"></a>

## regex_value property — domain / 120123203312 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-2230120010113122-1122203301211221-3311132200333003-3331331011130300-1132032111202012-1211211203011023-1202120113313102-3312031320033012"></a>

<a id="canonical-1023231132230010-3001320001221032-2010302120220222-1331032122111332-1030301123231230-1220302112213022-0033312210302020-3100332303032221"></a>

## suffix_value property — domain / 120123203312 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
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

<a id="canonical-1312130111210320-2012310322122320-0301002320013011-2033132202103233-0130123000002320-3120333223231211-1010202020022101-0310001010111203"></a>

## Next pages — domain / 120123203312 / 7

- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1320002021202303-1300330312202300-2213101032300021-2330111323130110-3220102210313203-2033213200201201-0010313221302112-3030032012032012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011220131300323-0000003332321333-2112130230131210-0111322033120030-0233311322023311-2223223223210333-3010310100020003-1120312302301113"></a>

## bot_defense.policy.js_insertion_rules.rules.metadata — metadata / 001222212222 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203)
- bot_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-0110323131220222-3131311131002312-1023102110320130-0200311112223030-1003032231212322-0110211130100320-2201033212112212-3022002211322023"></a>

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

<a id="canonical-2332330212233223-2123112113220312-0031231212130200-1230232130010212-0122033022312211-0020302030022132-1123121203313010-1013110201230123"></a>

## Direct properties — metadata / 001222212222 / 3

<a id="canonical-0030100211300001-1203211203213023-0002023320331021-0000013102310002-1031133222333230-2223331303310311-0011022000302331-3113233112211210"></a>

<a id="canonical-3020330333122031-0133123220312103-1001320331102300-2110302313110010-1121000021113103-1200321310201213-1003322302120210-3003030210012301"></a>

## description_spec property — metadata / 001222212222 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1112212231303301-0211132032221130-2313020212101300-0031131001110131-0312123023130111-2231000331202321-0112213223001020-3100102133112012"></a>

<a id="canonical-3312020120330123-2201312312220321-0300320100202321-2100030020330330-3203302233001011-1112131320111100-2300003332230001-2322221333220202"></a>

## name property — metadata / 001222212222 / 5

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

<a id="canonical-2321002230203221-2232320020011031-3103233103320133-0130220021311203-1303000022223102-0220301020322010-3223200210020130-1022302122122320"></a>

## Next pages — metadata / 001222212222 / 6

- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2001110231301133-3233002312320121-1323121031211213-1131301311230233-2220323210211320-0002232222312020-1120203313102001-2002212010321200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013202321121332-3313110300333010-3000313122112000-0200322000220113-3020031130231332-2331011313223302-3110200101210303-3022020332020233"></a>

## bot_defense.policy.js_insertion_rules.rules.path — path / 220031300303 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203)
- bot_defense.policy.js_insertion_rules.rules.path

<a id="canonical-0002102211012300-3331010101003121-1130122200130232-3122323022301311-3322120033303003-1021213132302033-2311123311033122-2311003103323201"></a>

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

<a id="canonical-2012233010311030-0332131300003112-2203200123331110-3232212112323212-2220112102131102-3032023120333302-1112201220030132-3031032211033003"></a>

## Direct properties — path / 220031300303 / 3

<a id="canonical-2212311222122110-2023220130103300-0120121303112102-1132002310332202-0100213231132023-0301200113330232-3020210210321103-2131302022320123"></a>

<a id="canonical-1323031001220102-1030331023030031-1231033310333032-0112120033201323-1233213121033000-1200200233311323-2120210230021222-3123310332102332"></a>

## path property — path / 220031300303 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

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

<a id="canonical-2332030113213010-1203131210030002-1230211131211132-0210011302333032-3333032302023133-2100300312302011-2202132012213020-0222301211121230"></a>

<a id="canonical-1213101030300110-2201100323123300-1001022333020003-0211320230313212-3220211331003031-0112202101221130-0020020210312010-2102302130230320"></a>

## prefix property — path / 220031300303 / 5

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

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

<a id="canonical-1131031100131223-3003113030230031-2303203021310011-3102132002102323-0220203101023000-1322002031013122-0323220022312102-3011111213122112"></a>

<a id="canonical-1030121223221203-3320132102232120-1203131103002000-3112323033222330-2110232221033211-2110133233102102-3110100122321023-1310211311303131"></a>

## regular expression property — path / 220031300303 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-2020122131030001-0230311012102011-3330012321013001-0303202012222101-0001000020312322-0213003333200011-2012113031220231-3323232102230102"></a>

## Next pages — path / 220031300303 / 7

- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0212032011201113-3013220200333112-0020211103232320-2200031302222032-2003022101011120-2310021110121233-3000301200111212-0220003230230203)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103100210100322-2120202100021000-1123030032221211-3131012121132311-1220111120021331-1021023102000230-2133100103001231-3131233221332102"></a>

## bot_defense.policy.mobile_sdk_config — mobile_sdk_config / 103133120220 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- bot_defense.policy.mobile_sdk_config

<a id="canonical-1021131320020210-2322012312332212-3333013322013133-2222101121013212-1013112132321302-3323333313320111-2322323012023301-1000022002301032"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3133333311332302-3111001122213003-3303103100003003-2310022121131123-1010211132212121-1123131012000321-2230320202300102-3013333012333113"></a>

## Direct properties — mobile_sdk_config / 103133120220 / 3

- [mobile_identifier](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0322101222233330-1023310231223322-3201103301212032-0020112201202223-0111021001113033-0023222330301312-1301222130120203-2010130211031122): complete subsection reference.

<a id="canonical-1010211103103033-2232223032003032-3202310223122110-0011121223133010-0311300003233233-0233022221130223-2313110233331213-2112311230322323"></a>

## Next pages — mobile_sdk_config / 103133120220 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0322101222233330-1023310231223322-3201103301212032-0020112201202223-0111021001113033-0023222330301312-1301222130120203-2010130211031122)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0322101222233330-1023310231223322-3201103301212032-0020112201202223-0111021001113033-0023222330301312-1301222130120203-2010130211031122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201220011102232-1100023131002313-0213130022033222-2111312102133021-1330121013220113-3013011212210331-0210221122220013-0200132122233203"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier — mobile_identifier / 320330123021 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303)
- bot_defense.policy.mobile_sdk_config.mobile_identifier

<a id="canonical-3123030323130011-3322222301301003-2103012023010031-3211203133011230-1131303023201103-0321220000223321-2300103212113202-3131303001020202"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3112103120110033-0321121213313003-2330331310132310-2232001021220203-0203212320233320-2030333301023300-2011200113332310-0131130033203223"></a>

## Direct properties — mobile_identifier / 320330123021 / 3

- [headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0120221020112300-0010013020111001-2322000000132103-1021000113332212-0321131110121023-1000231111222313-0203003203021130-3202102322110122): complete subsection reference.

<a id="canonical-2021023023201222-1033123010321103-1232232311133010-2012022033011301-0110033212122232-3021023322120001-1120031201011112-3312001233223022"></a>

## Next pages — mobile_identifier / 320330123021 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0120221020112300-0010013020111001-2322000000132103-1021000113332212-0321131110121023-1000231111222313-0203003203021130-3202102322110122)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0120221020112300-0010013020111001-2322000000132103-1021000113332212-0321131110121023-1000231111222313-0203003203021130-3202102322110122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310203211320121-0011232312102001-0012013123021321-1200232010103103-3211001132220022-2320011102310222-2300032201222112-1232201301021302"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers — headers / 013213003021 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0322101222233330-1023310231223322-3201103301212032-0020112201202223-0111021001113033-0023222330301312-1301222130120203-2010130211031122)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-1333103103300013-0031113132013020-3323202121320201-0020303022331312-3022003121323311-1000100010212002-3213032321103220-2312313001013022"></a>

Type: `"list"`. Computed.

Headers that can be used to identify mobile traffic.

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

<a id="canonical-3231132332030202-1302313021233103-3110220311231231-1310033132303201-0223233023130021-0030233032321321-3311222010120110-2002200310323300"></a>

## Direct properties — headers / 013213003021 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2230123331100012-2112313012300233-1103321222232130-1120002222022332-0102112232013231-0120202013301132-0220121000032311-0000323100322130): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1013313220210201-2300232330310012-1332212123231230-2313022021101001-1110231331222100-2300002030323231-3330323110020300-1031332131200231): complete subsection reference.

- [item](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0210331112101033-0310233332310321-1313223210301212-2031111013102032-1302121221100301-0110011211032013-0012312330130013-0333021111103202): complete subsection reference.

<a id="canonical-3002010313302102-0103211221033123-3221011103033111-1302103321011232-3232102012100220-3101322110320112-3331022111221002-0013123223010120"></a>

<a id="canonical-3123001100331123-3333233223212103-1012021000033313-1312201103020123-0002223030000131-0112033320013231-0001230211023113-3010230320310010"></a>

## name property — headers / 013213003021 / 4

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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

<a id="canonical-1232011130122032-2023303230011121-2323002102332012-3103023110112133-3113132330302122-2113132133032011-2230210333102230-2332030321023131"></a>

## Next pages — headers / 013213003021 / 5

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2230123331100012-2112313012300233-1103321222232130-1120002222022332-0102112232013231-0120202013301132-0220121000032311-0000323100322130)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1013313220210201-2300232330310012-1332212123231230-2313022021101001-1110231331222100-2300002030323231-3330323110020300-1031332131200231)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0210331112101033-0310233332310321-1313223210301212-2031111013102032-1302121221100301-0110011211032013-0012312330130013-0333021111103202)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0322101222233330-1023310231223322-3201103301212032-0020112201202223-0111021001113033-0023222330301312-1301222130120203-2010130211031122)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2230123331100012-2112313012300233-1103321222232130-1120002222022332-0102112232013231-0120202013301132-0220121000032311-0000323100322130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200023232033310-3023323223133011-0210213233021101-1003011121231210-1030001131110201-0230330121000333-0220233220101320-0021200110210223"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present — check_not_present / 011010331312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0322101222233330-1023310231223322-3201103301212032-0020112201202223-0111021001113033-0023222330301312-1301222130120203-2010130211031122)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0120221020112300-0010013020111001-2322000000132103-1021000113332212-0321131110121023-1000231111222313-0203003203021130-3202102322110122)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present

<a id="canonical-3012330323131123-0002302131313200-3231100003212123-3011313110133103-0021200133121232-1112001210212110-0212301100133330-2032002211120221"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3011221131320311-0303221312012131-2122123313012332-0310031330102133-3223221022330210-3231233213222132-1030312103313120-2203133012120312"></a>

## Direct properties — check_not_present / 011010331312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121230223100123-3232303230002213-3032122112010001-2322322303130021-1333132313000012-0230320002122013-2310332210221001-1333232220113202"></a>

## Next pages — check_not_present / 011010331312 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0120221020112300-0010013020111001-2322000000132103-1021000113332212-0321131110121023-1000231111222313-0203003203021130-3202102322110122)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1013313220210201-2300232330310012-1332212123231230-2313022021101001-1110231331222100-2300002030323231-3330323110020300-1031332131200231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121233233303222-1031210002102121-2113311200013030-1121002230331112-0223303212320223-3320311011332210-0310123202111103-0201001022300320"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present — check_present / 011301102003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0322101222233330-1023310231223322-3201103301212032-0020112201202223-0111021001113033-0023222330301312-1301222130120203-2010130211031122)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0120221020112300-0010013020111001-2322000000132103-1021000113332212-0321131110121023-1000231111222313-0203003203021130-3202102322110122)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present

<a id="canonical-3302032202220312-0003032200310220-0022030000202012-3033310031030011-2301210122331232-3321023000033030-1332230002100131-3330111200122021"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1210112101203022-1220313330221302-1130323230130132-2121011200103301-1113113312333213-3230213112313023-3223221110030010-0313311320123003"></a>

## Direct properties — check_present / 011301102003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131020233300010-2333211120133020-1113132302021323-1330132222332201-0030211000210203-3001330232313203-2131221201312332-0021330011302322"></a>

## Next pages — check_present / 011301102003 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0120221020112300-0010013020111001-2322000000132103-1021000113332212-0321131110121023-1000231111222313-0203003203021130-3202102322110122)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0210331112101033-0310233332310321-1313223210301212-2031111013102032-1302121221100301-0110011211032013-0012312330130013-0333021111103202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103012303302002-0030222220123133-3302123301100213-1033113101032003-3221300310001311-0102120121322022-1021203002312230-1030311111303313"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item — item / 003323312323 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0322101222233330-1023310231223322-3201103301212032-0020112201202223-0111021001113033-0023222330301312-1301222130120203-2010130211031122)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0120221020112300-0010013020111001-2322000000132103-1021000113332212-0321131110121023-1000231111222313-0203003203021130-3202102322110122)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item

<a id="canonical-3021310220012122-3333313022231113-1312313131310010-2021232200302021-0313233131020203-1323212111022313-0313332332123020-1101210313231331"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2032133013310010-1022030020010331-3101132110131002-1020211032123230-1332310003213303-2231003020321202-0133312231120003-1121312331030030"></a>

## Direct properties — item / 003323312323 / 3

<a id="canonical-1122322110201123-3230231000303001-2113300313312313-1001032231021003-2332010320322313-2211230321232212-2312231120130132-1022313310332213"></a>

<a id="canonical-3330121111013101-0213030222333211-0122111313210030-0212333222231112-0203220233103010-1220311320312333-0311100021320311-1232103132030121"></a>

## exact_values property — item / 003323312323 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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

<a id="canonical-3320130131221311-1100331130003013-3213203220220331-3021002220111203-2300032133002212-2310201033120311-3310032223100131-2110020112311220"></a>

<a id="canonical-0311301223200112-1301022112323322-2001331112131203-2031201121010010-1213203131222321-0202202103012003-2123210210312103-3103233121302012"></a>

## regex_values property — item / 003323312323 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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

<a id="canonical-2030110203302110-1120233303033012-0210121200223211-1131320302210330-1122333330202322-1320132030233211-3333130232001231-0213333113101132"></a>

<a id="canonical-2100110003010023-2323031010320122-0212123110010030-2012303300320323-3310333000312200-2000101233000233-3213010231330121-2131232213201311"></a>

## transformers property — item / 003323312323 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

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

<a id="canonical-2302112100330232-0312322000210132-2133121100213220-1332112202212220-3201213021111111-1203013001023301-2230121102220001-0322303130320001"></a>

## Next pages — item / 003323312323 / 7

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0120221020112300-0010013020111001-2322000000132103-1021000113332212-0321131110121023-1000231111222313-0203003203021130-3202102322110122)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102232130221100-1103303211323111-1322021233123211-3322020201321212-0031313301330211-2002301033302230-0330102220332103-3100210001001031"></a>

## bot_defense.policy.protected_app_endpoints — protected_app_endpoints / 221203012031 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- bot_defense.policy.protected_app_endpoints

<a id="canonical-3102110122102031-2222231231231322-0103210011120322-3023333223201112-1002103001131330-0302311333020020-2223103200323320-2302132220230322"></a>

Type: `"list"`. Computed.

List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32
endpoints per LB' after 4 LBs.

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

<a id="canonical-0030011200130303-2112111333331131-0102300202022013-0300022200222221-3223303332110011-1121013330212003-2021330133121131-1010321332212333"></a>

## Direct properties — protected_app_endpoints / 221203012031 / 3

- [allow_good_bots](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1320332211131210-1021321312023222-3123210020310132-0013210002312322-2301331012131331-3131232120002211-1123011300223200-1201120201211202): complete subsection reference.

- [any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0323031113221202-0100021020320233-1112333032313203-0303232011320202-0200230311321232-2230123010331133-2203121232101213-0223011210332121): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2021123002221002-2000312022300103-0310231302200312-0302021212223332-2321010211102020-0222122100230102-1232133203021120-0013300000200020): complete subsection reference.

- [flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123): complete subsection reference.

- [headers](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0131301311233211-1033013031310110-2112103222212000-3320331133030301-3233122330302113-0100023002323130-0221021031221210-0013231120031111): complete subsection reference.

<a id="canonical-3311121100201102-1101021211103221-2222330322221303-3321200102011322-3323220123031221-1120122330302200-2321300332311033-1122122231111303"></a>

<a id="canonical-0203201133031220-1213112120033010-3023030120211133-2123010303022333-0311220231212011-0312002033223201-0321120231021310-1033220103211110"></a>

## http_methods property — protected_app_endpoints / 221203012031 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Upstream description:

List of HTTP methods.

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

- [metadata](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0310323201120212-2211300010002233-0131322301103120-1313223022102130-3200302012100210-2312020222000331-1321211312313202-2130133321123302): complete subsection reference.

- [mitigate_good_bots](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3220231313001133-3232133123323110-2113200211332101-1220120031022303-3033110013001233-3303211003002222-1223032222003112-1321230220100320): complete subsection reference.

- [mitigation](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0200202233003202-0102330323232112-1032303100002320-0221321031020031-1300031023211332-3313011300222213-3311231221322001-1313131230321112): complete subsection reference.

- [mobile](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2131310300213221-3110023200313012-3033330132222210-2012123330230132-2313132313030233-2003133202020011-0032020300330200-2221013113321022): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3013311333023121-2031310033103232-2110030121212030-0223102001022133-1331321311100101-2101002012113311-3331321003101023-2231120301233202): complete subsection reference.

<a id="canonical-2130102213121103-1122310332212333-3011033310303101-2311102101011000-2010032023203332-1102230003100021-2310031032321313-0203322212131000"></a>

<a id="canonical-2011102310300121-1210012310102013-2211203111321031-1332021133233233-1020323102023233-2323311330231200-0321311201111132-2012201313212031"></a>

## protocol property — protected_app_endpoints / 221203012031 / 5

Type: `"string"`. Computed.

\[Enum: BOTH|HTTP|HTTPS\] SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for
HTTPS:// or HTTP://. - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only.
Possible values are \`BOTH\`, \`HTTP\`, \`HTTPS\`. Defaults to \`BOTH\`.

Upstream description:

SchemeType is used to indicate URL scheme.

&#8203;- BOTH: BOTH

URL scheme for HTTPS:// or HTTP://. &#8203;- HTTP: HTTP

URL scheme HTTP:// only. &#8203;- HTTPS: HTTPS

URL scheme HTTPS:// only.

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

- [query_params](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2033213231111310-1201133200120003-2000203101330212-0313030322132332-3033031231110033-3232201331213211-1323002113020311-3120110033330200): complete subsection reference.

- [undefined_flow_label](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1003003211020113-3103312331332221-2103313103132131-1113322131300002-1031212331122122-3012103220120231-3032222232112132-2303111022032323): complete subsection reference.

- [web](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2222302100002000-1121300110111033-2303003223331311-0113322310023022-1123233232023213-3103232323330012-0002320230103130-1111220303320033): complete subsection reference.

- [web_mobile](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2212233111032332-2322130100100201-1312110231313203-3103222223203300-1323133131321310-3310021133322023-0121033021113312-1231221220103101): complete subsection reference.

<a id="canonical-3231302233311200-0321222132111033-3033330333301333-3023311320212110-3210033023321210-2111132203110003-0120202130111220-1212121033320002"></a>

## Next pages — protected_app_endpoints / 221203012031 / 6

- [bot_defense.policy.protected_app_endpoints.allow_good_bots](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1320332211131210-1021321312023222-3123210020310132-0013210002312322-2301331012131331-3131232120002211-1123011300223200-1201120201211202)
- [bot_defense.policy.protected_app_endpoints.any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0323031113221202-0100021020320233-1112333032313203-0303232011320202-0200230311321232-2230123010331133-2203121232101213-0223011210332121)
- [bot_defense.policy.protected_app_endpoints.domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2021123002221002-2000312022300103-0310231302200312-0302021212223332-2321010211102020-0222122100230102-1232133203021120-0013300000200020)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.headers](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0131301311233211-1033013031310110-2112103222212000-3320331133030301-3233122330302113-0100023002323130-0221021031221210-0013231120031111)
- [bot_defense.policy.protected_app_endpoints.metadata](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0310323201120212-2211300010002233-0131322301103120-1313223022102130-3200302012100210-2312020222000331-1321211312313202-2130133321123302)
- [bot_defense.policy.protected_app_endpoints.mitigate_good_bots](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3220231313001133-3232133123323110-2113200211332101-1220120031022303-3033110013001233-3303211003002222-1223032222003112-1321230220100320)
- [bot_defense.policy.protected_app_endpoints.mitigation](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0200202233003202-0102330323232112-1032303100002320-0221321031020031-1300031023211332-3313011300222213-3311231221322001-1313131230321112)
- [bot_defense.policy.protected_app_endpoints.mobile](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2131310300213221-3110023200313012-3033330132222210-2012123330230132-2313132313030233-2003133202020011-0032020300330200-2221013113321022)
- [bot_defense.policy.protected_app_endpoints.path](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3013311333023121-2031310033103232-2110030121212030-0223102001022133-1331321311100101-2101002012113311-3331321003101023-2231120301233202)
- [bot_defense.policy.protected_app_endpoints.query_params](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2033213231111310-1201133200120003-2000203101330212-0313030322132332-3033031231110033-3232201331213211-1323002113020311-3120110033330200)
- [bot_defense.policy.protected_app_endpoints.undefined_flow_label](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1003003211020113-3103312331332221-2103313103132131-1113322131300002-1031212331122122-3012103220120231-3032222232112132-2303111022032323)
- [bot_defense.policy.protected_app_endpoints.web](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2222302100002000-1121300110111033-2303003223331311-0113322310023022-1123233232023213-3103232323330012-0002320230103130-1111220303320033)
- [bot_defense.policy.protected_app_endpoints.web_mobile](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2212233111032332-2322130100100201-1312110231313203-3103222223203300-1323133131321310-3310021133322023-0121033021113312-1231221220103101)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1320332211131210-1021321312023222-3123210020310132-0013210002312322-2301331012131331-3131232120002211-1123011300223200-1201120201211202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131220302312313-0121123121302011-2320320020112203-0230103221132310-1313032101122022-1101110131120103-1203003201000210-2112010023212010"></a>

## bot_defense.policy.protected_app_endpoints.allow_good_bots — allow_good_bots / 312122321010 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- bot_defense.policy.protected_app_endpoints.allow_good_bots

<a id="canonical-1201122311212111-3303132203303022-0321322123023030-2230012131030002-2130002003101333-1033310303123033-3330133000311002-3210202132111112"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0211232110321111-2312103211220321-3021020032013313-3231233213310303-3113222212122130-3030213013220103-2102122202221320-0202020111201021"></a>

## Direct properties — allow_good_bots / 312122321010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200021301000001-3130033132201030-2023132331232012-0213100231220120-1323122200300113-3111000320313212-1213230001011330-2333300310032033"></a>

## Next pages — allow_good_bots / 312122321010 / 4

- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0323031113221202-0100021020320233-1112333032313203-0303232011320202-0200230311321232-2230123010331133-2203121232101213-0223011210332121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211323021003131-1003000333300210-0302112003322020-1120001213323200-3320111101103232-3200233301303301-2331100233221032-2311300301331203"></a>

## bot_defense.policy.protected_app_endpoints.any_domain — any_domain / 002123111012 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- bot_defense.policy.protected_app_endpoints.any_domain

<a id="canonical-0221230232002202-0121031113223000-3102332331012212-1033123321310121-2002101201021210-2202331030022011-1312302102233210-3102020010323232"></a>

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

<a id="canonical-0221300321333320-3322010130223132-3021310000302321-2123332033110031-2031322113330211-0223323011323111-1100020032133203-1121133110120322"></a>

## Direct properties — any_domain / 002123111012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133023133323121-2021323202133020-3100031121233303-1032231212211031-1113103101020113-1020010223223011-1102312101222132-3030102211122200"></a>

## Next pages — any_domain / 002123111012 / 4

- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2021123002221002-2000312022300103-0310231302200312-0302021212223332-2321010211102020-0222122100230102-1232133203021120-0013300000200020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210021200132202-2123123022111231-0330100131110100-2103032211231021-0110132021002030-0311121002222020-0302020132213312-3233333222232233"></a>

## bot_defense.policy.protected_app_endpoints.domain — domain / 032123010311 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- bot_defense.policy.protected_app_endpoints.domain

<a id="canonical-2200311010131130-1300310021220000-1302331023102310-3332221033010130-0200133033323112-1301111213311003-1103330211230033-1023203301021313"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

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

<a id="canonical-0003203223230120-3332201211331013-1313131300232321-0133203130030022-1220123211112202-3221111022131021-0331302223111212-3032010303202210"></a>

## Direct properties — domain / 032123010311 / 3

<a id="canonical-0330002131302100-2032313030211330-2130300322210021-0020013022310123-1111233030033322-1011011223323031-1300221332221331-3321210201010303"></a>

<a id="canonical-2311303131120133-0311231312002320-3013002233301003-0033302013100321-0020121332202000-1012331001023102-1000021300110331-3220030112301022"></a>

## exact_value property — domain / 032123010311 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

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

<a id="canonical-2210222301110113-0032233102000031-3333233230321003-1033222311023321-0031111301232100-1333200120200313-2310231101312021-3010032221021230"></a>

<a id="canonical-3231113213321300-3031123022112132-3332313320021301-2312330320123202-3203302320230233-3012111100121111-0211222221202222-3231212312010101"></a>

## regex_value property — domain / 032123010311 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-2231211101131323-0132313313023330-2321222303130013-0233320202032321-0233212330333130-0112322021323011-1322012000300133-3312123032121301"></a>

<a id="canonical-3311203020102302-0002123011233122-2321021221111110-0120321231022310-0311033132221102-1203013313220113-3031200323001121-1000210223002100"></a>

## suffix_value property — domain / 032123010311 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
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

<a id="canonical-1032331230321230-3210020131020230-3121212030221121-3320022011332021-2003103102332013-1223130331011313-0300020203313231-3001313023303101"></a>

## Next pages — domain / 032123010311 / 7

- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011332132221230-1202010223200332-1212311012233300-3210023102233030-0201223230233301-1121131300131001-3001102103033302-3332201323301201"></a>

## bot_defense.policy.protected_app_endpoints.flow_label — flow_label / 012101323001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- bot_defense.policy.protected_app_endpoints.flow_label

<a id="canonical-3112330311201102-0010022301203030-0210332022033003-0132121311113131-1302032033231031-1100032320201032-0313110133233223-3221211123233320"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Category allows to associate traffic with selected category.

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

<a id="canonical-0223133233302323-0011103010010102-3121122023001110-3201313313303112-3032300232122020-2113100331121103-1220233231332211-3110213321200202"></a>

## Direct properties — flow_label / 012101323001 / 3

- [account_management](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3012232111033320-2231113022320000-3013203130211113-3211022000030112-2123320300103221-2123022311030010-2012200000232313-2033302023012203): complete subsection reference.

- [authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023): complete subsection reference.

- [financial_services](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3330221233222103-1313020130322211-3111000122310002-1232332223103332-3131131033011301-2302301113113112-2232323013001101-2322220003110133): complete subsection reference.

- [flight](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1121213000022101-2113233123203300-3203013121221330-1213013330213330-1201312301011221-0132003011203223-2333122213232121-2110111331301203): complete subsection reference.

- [profile_management](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0233133110211223-2113223011223221-2330102323022103-3022000031110221-3132213012031233-1123211113212302-3130110303223323-0100131321322310): complete subsection reference.

- [search](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1213100300231123-1222123321303201-0132133230023303-2331311203211323-0120200120301330-0002030202132133-1033003312303112-2222033303032022): complete subsection reference.

- [shopping_gift_cards](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1032320300202132-1310332133332021-2003001332331222-1222230122213120-3120011100320133-0220113123101222-1001120010133321-1102230112210131): complete subsection reference.

<a id="canonical-0101132230010212-0010102330032012-1113212223131302-1032233031100333-1200102102111130-2031310011102001-2303310331003212-1311201101030132"></a>

## Next pages — flow_label / 012101323001 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3012232111033320-2231113022320000-3013203130211113-3211022000030112-2123320300103221-2123022311030010-2012200000232313-2033302023012203)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3330221233222103-1313020130322211-3111000122310002-1232332223103332-3131131033011301-2302301113113112-2232323013001101-2322220003110133)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1121213000022101-2113233123203300-3203013121221330-1213013330213330-1201312301011221-0132003011203223-2333122213232121-2110111331301203)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0233133110211223-2113223011223221-2330102323022103-3022000031110221-3132213012031233-1123211113212302-3130110303223323-0100131321322310)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1213100300231123-1222123321303201-0132133230023303-2331311203211323-0120200120301330-0002030202132133-1033003312303112-2222033303032022)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1032320300202132-1310332133332021-2003001332331222-1222230122213120-3120011100320133-0220113123101222-1001120010133321-1102230112210131)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3012232111033320-2231113022320000-3013203130211113-3211022000030112-2123320300103221-2123022311030010-2012200000232313-2033302023012203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123220302101031-3321212322300123-0132320313122230-2122202030101331-1210221133303031-2011003330221211-2003022311213120-2020323131100222"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management — account_management / 110011113333 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management

<a id="canonical-2020100332200013-3200211313312301-3003302010223211-2111302320002301-0213000002022131-0032010230322012-0230103310222210-1223022213222230"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Account Management Category.

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

<a id="canonical-3303200322232132-3101033331101231-2103323013210020-2223033101102002-0022210021301232-2220120232320011-3313200002031102-1222231202302312"></a>

## Direct properties — account_management / 110011113333 / 3

- [create](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0031033110322310-2330323301202210-3332222220211301-1120030021320203-0013210222300311-2103011222332200-3312123300323020-2212110211202103): complete subsection reference.

- [password_reset](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3232030013331112-3001310332222322-0231131002323203-2220203001330222-1031300122223013-3100321312111310-3033321323222033-2310011333132222): complete subsection reference.

<a id="canonical-3321033301200102-0010103012323132-2311200312032113-1330003113202122-1231100313011131-1332330311300021-2201212302000020-3210303031113120"></a>

## Next pages — account_management / 110011113333 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.create](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0031033110322310-2330323301202210-3332222220211301-1120030021320203-0013210222300311-2103011222332200-3312123300323020-2212110211202103)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3232030013331112-3001310332222322-0231131002323203-2220203001330222-1031300122223013-3100321312111310-3033321323222033-2310011333132222)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0031033110322310-2330323301202210-3332222220211301-1120030021320203-0013210222300311-2103011222332200-3312123300323020-2212110211202103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133310103022211-0012001320220121-2223203120300001-0230113001203301-2210010322303010-3333021033122202-3032022130211331-2200322103132320"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management.create — create / 103030311100 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3012232111033320-2231113022320000-3013203130211113-3211022000030112-2123320300103221-2123022311030010-2012200000232313-2033302023012203)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.create

<a id="canonical-2221002311112031-3233002303331230-3121013020303122-1102100331220332-1031132131020130-0033201101100323-2232020020102131-0201221211123312"></a>

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

<a id="canonical-0211033103221302-1003002321223002-2212031001310202-1230233002021121-1032221100212313-3001323301233232-0213210012223321-0220001111112020"></a>

## Direct properties — create / 103030311100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302230113323211-3322001012123220-0312332222311021-1311213103220230-1103113012303002-3203003021220202-3303212321131133-2033133011111232"></a>

## Next pages — create / 103030311100 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3012232111033320-2231113022320000-3013203130211113-3211022000030112-2123320300103221-2123022311030010-2012200000232313-2033302023012203)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3232030013331112-3001310332222322-0231131002323203-2220203001330222-1031300122223013-3100321312111310-3033321323222033-2310011333132222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113303101320303-1322010001031213-2213333232022213-2211301120012233-3312022320203020-1203100202032010-2320311202311211-2331223233030003"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset — password_reset / 310133323121 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3012232111033320-2231113022320000-3013203130211113-3211022000030112-2123320300103221-2123022311030010-2012200000232313-2033302023012203)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset

<a id="canonical-0003201131100133-1012131001000121-2132233001202310-2231321211310210-0231131320220233-3111011111023101-1331311021110112-0330211302031220"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3031111232133332-1302123100030333-1223331101121003-0001312222233110-1232230302310022-1123223020030121-2010032331102023-1130131001012331"></a>

## Direct properties — password_reset / 310133323121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112133011032110-2030002012203132-1111003103002333-1230312221233120-1232333222212220-2032123320222011-1002102011111313-1311010211332331"></a>

## Next pages — password_reset / 310133323121 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3012232111033320-2231113022320000-3013203130211113-3211022000030112-2123320300103221-2123022311030010-2012200000232313-2033302023012203)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331332333220113-0231212112021133-2321232132230031-0303230210130230-0000321012032201-1332213021221211-2103310211222003-3012331201103132"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication — authentication / 102100131332 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication

<a id="canonical-0201123103200010-2210202301212311-2312331101322033-3330302203212033-2133313013022001-0231232032323011-3110230211112131-3300000023120232"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Authentication Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

<a id="canonical-2100302232211332-2013302210330323-1331120133212022-0133203220322111-0001003200233210-1103231131203132-0323111113123301-0013222302202230"></a>

## Direct properties — authentication / 102100131332 / 3

- [login](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3131113030220122-0202331010323001-2122322201031222-2313130113302010-2320111320101310-1113132210221330-1210023001313110-1031131310202220): complete subsection reference.

- [login_mfa](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2322222331301102-3001311123013310-0233301000321302-3301322311311123-1013323220233020-3333330021301001-1013130121302231-0220002321210223): complete subsection reference.

- [login_partner](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3321320212303103-2130110313220111-0323210221221213-0003221223112033-2321222101333231-3032101012310230-3232110132233212-3010013102123021): complete subsection reference.

- [logout](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2313232313010221-1120111022102132-2122233321130312-0110030020330301-1332311320000223-3311021210322103-1000112101020100-0220212132313020): complete subsection reference.

- [token_refresh](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2312332120303010-0220113223322203-1233122231300322-0012222223302112-1330302230132121-0031110300330302-0231322311230003-3020031230001103): complete subsection reference.

<a id="canonical-1313330011102103-2002031111133303-0011110012011132-2221332202102330-0320000103001222-0300011221202021-2010230211203021-2102313232302203"></a>

## Next pages — authentication / 102100131332 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3131113030220122-0202331010323001-2122322201031222-2313130113302010-2320111320101310-1113132210221330-1210023001313110-1031131310202220)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2322222331301102-3001311123013310-0233301000321302-3301322311311123-1013323220233020-3333330021301001-1013130121302231-0220002321210223)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3321320212303103-2130110313220111-0323210221221213-0003221223112033-2321222101333231-3032101012310230-3232110132233212-3010013102123021)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2313232313010221-1120111022102132-2122233321130312-0110030020330301-1332311320000223-3311021210322103-1000112101020100-0220212132313020)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2312332120303010-0220113223322203-1233122231300322-0012222223302112-1330302230132121-0031110300330302-0231322311230003-3020031230001103)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3131113030220122-0202331010323001-2122322201031222-2313130113302010-2320111320101310-1113132210221330-1210023001313110-1031131310202220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320122321101111-3313113333231321-3123102032013100-2332031002023302-1023120022222200-2321021322203211-3300102331013211-2112330021212132"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login — login / 201321001201 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login

<a id="canonical-0011033201321023-0313020223032020-3023331123121022-2320311111200100-2213213111021210-3031011023331032-2202331101320122-2302333013322203"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Upstream description:

Bot Defense Transaction Result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

<a id="canonical-3112201033310123-1323231221313001-0113022010030112-3201301011312212-0321303232121130-2201321122121230-1030123301202231-1002222213322321"></a>

## Direct properties — login / 201321001201 / 3

- [disable_transaction_result](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3233102101112130-2221002133001201-3120023200110232-2310201331023211-0131321332222323-2221033123311123-3122332120132302-0011220111132031): complete subsection reference.

- [transaction_result](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0000221300122333-3101003110122012-2331311331003032-3333033101012222-3202132202131211-1231202120000322-2302211323310213-1221210312023313): complete subsection reference.

<a id="canonical-0220023200000301-0020313332131031-3132013322222103-3320332210211123-1323311211203110-0103010031120222-0120333012321222-2011232213301132"></a>

## Next pages — login / 201321001201 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3233102101112130-2221002133001201-3120023200110232-2310201331023211-0131321332222323-2221033123311123-3122332120132302-0011220111132031)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0000221300122333-3101003110122012-2331311331003032-3333033101012222-3202132202131211-1231202120000322-2302211323310213-1221210312023313)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3233102101112130-2221002133001201-3120023200110232-2310201331023211-0131321332222323-2221033123311123-3122332120132302-0011220111132031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122022333113020-1322322001321033-3200222130212320-2221333010002101-0232302331302330-2123021330010320-3121222213130031-3131002202121233"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result — disable_transaction_result / 001311132021 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1020000200021110-1000312131300302-0331001113212120-0010211333011103-2032320210201310-3323312010112221-1110330133020021-3203203102121123)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0133130221211133-1300102212120332-3201303213303223-3031222030133020-2030121323210032-0133120201032121-0102101220323121-2113120320011023)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3131113030220122-0202331010323001-2122322201031222-2313130113302010-2320111320101310-1113132210221330-1210023001313110-1031131310202220)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result

<a id="canonical-2301221232331022-1032321131133122-1201313102222021-3230310002232130-0310231113231321-0021131120010133-1123302013311201-1210301323002020"></a>

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

<a id="canonical-3122300000131210-2203133223331031-1200212003220212-0012123311320322-2012003211133131-0011021021113133-1102202113223120-3103020002330032"></a>

## Direct properties — disable_transaction_result / 001311132021 / 3

This is an empty object or choice marker. It has no direct properties.
