---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2002113022131110-2321231021211020-2332032331320331-3002032212123200-1222131233033331-3302101222200303-1011312200031331-1322203111233132"></a>

## data_guard_rules.path — path / 020222321021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [data_guard_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-3102323311020130-1122103001002133-3000002211222110-3333021323121303-3331301010012011-3020222021303103-3231322003001201-3132013303031003)
- data_guard_rules.path

<a id="canonical-3000203000322121-0000031031101010-2112013002033303-0130232030133322-1202110011311010-2231231313013111-2012232233232210-3030113101210230"></a>

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

<a id="canonical-2300303212331332-1321303033321233-1103221133113031-0023210100122032-0132003333012331-1222033122003101-3302310123220223-3021110101221210"></a>

## Direct properties — path / 020222321021 / 3

<a id="canonical-3231333121123031-3232131010320022-0230103210323122-2103100332201310-2102321221001021-0121100023330123-2013300012330332-2300131230121201"></a>

<a id="canonical-2131302322222000-3021001210301211-0221212202203203-2211310030130301-2031201032132311-3300030023011221-0331311131310332-3323110121313131"></a>

## path property — path / 020222321021 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2223020313313013-3031212302132220-0201101031002333-3010012232332113-0112302102302232-2320330021030122-0202330320010201-0301001320023002"></a>

<a id="canonical-1301020031010013-0013022033311010-0313213120121331-3122331102123030-0323223012130110-2310113022332121-2123112103332210-3130032312132112"></a>

## prefix property — path / 020222321021 / 5

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

<a id="canonical-1211230011131211-1210112202330331-0132133011300013-2301320001202123-0323313302010103-2033001103123222-3220230032230211-1120110211332003"></a>

<a id="canonical-0312233002012232-2230021123023231-2310321302133012-3012310000120100-0311013320310011-1001330301022133-1001133212110203-0323203023330023"></a>

## regular expression property — path / 020222321021 / 6

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

<a id="canonical-0132301303233003-0201333133001011-3122001111111002-1320010010113133-2132310122220330-3103123011111200-0231332201033030-3320223101113130"></a>

## Next pages — path / 020222321021 / 7

- [data_guard_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-3102323311020130-1122103001002133-3000002211222110-3333021323121303-3331301010012011-3020222021303103-3231322003001201-3132013303031003)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2111200221101132-1102030310103202-0123131133013021-2000001130233001-2321323033332102-1202013010233021-0002303120110202-1322310203021212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221122201221313-0200123002030221-2110300212231303-3120322103210003-0120102030202120-3210033323122132-1331113012331331-2000231230021023"></a>

## data_guard_rules.skip_data_guard — skip_data_guard / 121122313132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [data_guard_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-3102323311020130-1122103001002133-3000002211222110-3333021323121303-3331301010012011-3020222021303103-3231322003001201-3132013303031003)
- data_guard_rules.skip_data_guard

<a id="canonical-2112022130230301-0101311030002021-1032221133312010-2230032200210311-2020311310313113-1222131212020123-0202231202001333-2132323333323102"></a>

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

<a id="canonical-1201000211320321-0231300231111022-2223010311313312-3023112001220102-3020000231102003-3133123132223022-1200202102311021-0203013131010331"></a>

## Direct properties — skip_data_guard / 121122313132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123303211011122-0131201300221122-3010220131131220-2232001223311033-1212211111200330-2032321313300232-2331322001233332-2301311023222212"></a>

## Next pages — skip_data_guard / 121122313132 / 4

- [data_guard_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-3102323311020130-1122103001002133-3000002211222110-3333021323121303-3331301010012011-3020222021303103-3231322003001201-3132013303031003)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122102303110000-1311220122212220-3010103102322121-2100200302031130-2202002202303000-0122210101121212-2301231011230030-0112123330330322"></a>

## ddos_mitigation_rules — ddos_mitigation_rules / 311112102130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- ddos_mitigation_rules

<a id="canonical-0222330120211123-0011202332200013-2112301203110122-1110303322031332-1103131020323122-2323100131123103-3202211033221113-1010303020023112"></a>

Type: `"list"`. Computed.

Define manual mitigation rules to block L7 DDoS attacks.

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

<a id="canonical-1323203003010022-1322230211320233-0123230010311313-2000220131233020-2011312012132031-2011333333331303-1000210223203122-1002320222233311"></a>

## Direct properties — ddos_mitigation_rules / 311112102130 / 3

- [block](data-sources--http_loadbalancer--reference--group-015.md#canonical-3123311102233033-2133220002012003-1200022032312311-2033132203131033-0302133003210102-0211133322321231-1023020222113331-2200302213010302): complete subsection reference.

- [ddos_client_source](data-sources--http_loadbalancer--reference--group-015.md#canonical-2010132332231223-1321201300322332-1003012200330023-2322321131102132-2330033003320010-2133300030311102-1211110222320210-1201200010002203): complete subsection reference.

<a id="canonical-2203100123132033-1122223233202320-2032112011201002-3332121103001221-1330203133213012-3011022002221322-3233301023113233-0011123302003330"></a>

<a id="canonical-1000233213310223-3133302011333013-2233103130110300-3301232123011123-2103302222000021-0001121203022123-0012010301220022-3203120121111202"></a>

## expiration_timestamp property — ddos_mitigation_rules / 311112102130 / 4

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

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

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-015.md#canonical-2211231022121230-1011323212121001-2133133122021303-3332030210202013-2210323321121202-3323133220200310-1123101311301033-1102111103021310): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-015.md#canonical-3211223222113323-1020030013300003-0322233310230102-1333301120222330-2330300300132111-2030313102210023-3301233101112230-0011300100122012): complete subsection reference.

<a id="canonical-2322011012032300-3001120021110232-0013000110002203-3131330111300201-3233010320233312-1310210211212021-0033332320012201-0032010001331213"></a>

## Next pages — ddos_mitigation_rules / 311112102130 / 5

- [ddos_mitigation_rules.block](data-sources--http_loadbalancer--reference--group-015.md#canonical-3123311102233033-2133220002012003-1200022032312311-2033132203131033-0302133003210102-0211133322321231-1023020222113331-2200302213010302)
- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-015.md#canonical-2010132332231223-1321201300322332-1003012200330023-2322321131102132-2330033003320010-2133300030311102-1211110222320210-1201200010002203)
- [ddos_mitigation_rules.ip_prefix_list](data-sources--http_loadbalancer--reference--group-015.md#canonical-2211231022121230-1011323212121001-2133133122021303-3332030210202013-2210323321121202-3323133220200310-1123101311301033-1102111103021310)
- [ddos_mitigation_rules.metadata](data-sources--http_loadbalancer--reference--group-015.md#canonical-3211223222113323-1020030013300003-0322233310230102-1333301120222330-2330300300132111-2030313102210023-3301233101112230-0011300100122012)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3123311102233033-2133220002012003-1200022032312311-2033132203131033-0302133003210102-0211133322321231-1023020222113331-2200302213010302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211011102132121-1230031110021332-2010323000321023-0311212300003302-0021300212123331-1010031212032232-2200011113021122-0233012232030010"></a>

## ddos_mitigation_rules.block — block / 223020330002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- ddos_mitigation_rules.block

<a id="canonical-1210012321002200-3333111301020223-2212103132320332-0321130112320022-0312332321031123-1330132230122313-1333100201321203-2030000200321323"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1003202121202320-0303311332003310-2023233021230323-0013112222330320-1133011112202222-1111022003012123-3012111222113000-0310232011101033"></a>

## Direct properties — block / 223020330002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132212311031211-1331021101003023-3233000103012322-3131032213131313-0203123323002301-0302030233001033-0102001001330102-3233103012211313"></a>

## Next pages — block / 223020330002 / 4

- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2010132332231223-1321201300322332-1003012200330023-2322321131102132-2330033003320010-2133300030311102-1211110222320210-1201200010002203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013213121011330-3212301332131133-2001233313102221-3300221031003130-2320013102200330-0223103123102121-1301213012311112-3213113122003311"></a>

## ddos_mitigation_rules.ddos_client_source — ddos_client_source / 301332110033 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- ddos_mitigation_rules.ddos_client_source

<a id="canonical-2032003231013220-1023210320302003-2231130230313201-3231322112330332-2021331312003021-1111331330303120-1023123233012013-1033231132332323"></a>

Type: `"single"`. Computed.

DDoS Client Source Choice. DDoS Mitigation sources to be blocked.

Upstream description:

DDoS Mitigation sources to be blocked.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0233213202130221-2112011311122202-1011031100122230-3320021022212123-2012202002131321-0223212120220212-3021131322221210-2233022333123032"></a>

## Direct properties — ddos_client_source / 301332110033 / 3

- [asn_list](data-sources--http_loadbalancer--reference--group-015.md#canonical-3330032330013000-1211133133122010-2130122100120230-0100010212032131-2332210311132233-0101302230121001-1213033123113100-2332311202002232): complete subsection reference.

<a id="canonical-0122330112132233-0210130302032203-2032023000213302-2321212112010300-3331132132022321-1103310013113213-0121111201020120-0103003310011321"></a>

<a id="canonical-2012103323312220-0231221321023011-2010312130203012-2331130322001103-0100220311012132-2231210110132021-2012003003220010-0132333112221131"></a>

## country_list property — ddos_client_source / 301332110033 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Sources that are located in one of the countries in the given list. Possible values are
\`COUNTRY\_NONE\`, \`COUNTRY\_AD\`, \`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`,
\`COUNTRY\_AI\`, \`COUNTRY\_AL\`, \`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`,
\`COUNTRY\_AQ\`, \`COUNTRY\_AR\`, \`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`,
\`COUNTRY\_AW\`, \`COUNTRY\_AX\`, \`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`,
\`COUNTRY\_BD\`, \`COUNTRY\_BE\`, \`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`,
\`COUNTRY\_BI\`, \`COUNTRY\_BJ\`, \`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`,
\`COUNTRY\_BO\`, \`COUNTRY\_BQ\`, \`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`,
\`COUNTRY\_BV\`, \`COUNTRY\_BW\`, \`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`,
\`COUNTRY\_CC\`, \`COUNTRY\_CD\`, \`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`,
\`COUNTRY\_CI\`, \`COUNTRY\_CK\`, \`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`,
\`COUNTRY\_CO\`, \`COUNTRY\_CR\`, \`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`,
\`COUNTRY\_CW\`, \`COUNTRY\_CX\`, \`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`,
\`COUNTRY\_DJ\`, \`COUNTRY\_DK\`, \`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`,
\`COUNTRY\_EC\`, \`COUNTRY\_EE\`, \`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`,
\`COUNTRY\_ES\`, \`COUNTRY\_ET\`, \`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`,
\`COUNTRY\_FM\`, \`COUNTRY\_FO\`, \`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`,
\`COUNTRY\_GD\`, \`COUNTRY\_GE\`, \`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`,
\`COUNTRY\_GI\`, \`COUNTRY\_GL\`, \`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`,
\`COUNTRY\_GQ\`, \`COUNTRY\_GR\`, \`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`,
\`COUNTRY\_GW\`, \`COUNTRY\_GY\`, \`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`,
\`COUNTRY\_HR\`, \`COUNTRY\_HT\`, \`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`,
\`COUNTRY\_IL\`, \`COUNTRY\_IM\`, \`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`,
\`COUNTRY\_IR\`, \`COUNTRY\_IS\`, \`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`,
\`COUNTRY\_JO\`, \`COUNTRY\_JP\`, \`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`,
\`COUNTRY\_KI\`, \`COUNTRY\_KM\`, \`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`,
\`COUNTRY\_KW\`, \`COUNTRY\_KY\`, \`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`,
\`COUNTRY\_LC\`, \`COUNTRY\_LI\`, \`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`,
\`COUNTRY\_LT\`, \`COUNTRY\_LU\`, \`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`,
\`COUNTRY\_MC\`, \`COUNTRY\_MD\`, \`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`,
\`COUNTRY\_MH\`, \`COUNTRY\_MK\`, \`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`,
\`COUNTRY\_MO\`, \`COUNTRY\_MP\`, \`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`,
\`COUNTRY\_MT\`, \`COUNTRY\_MU\`, \`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`,
\`COUNTRY\_MY\`, \`COUNTRY\_MZ\`, \`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`,
\`COUNTRY\_NF\`, \`COUNTRY\_NG\`, \`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`,
\`COUNTRY\_NP\`, \`COUNTRY\_NR\`, \`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`,
\`COUNTRY\_PA\`, \`COUNTRY\_PE\`, \`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`,
\`COUNTRY\_PK\`, \`COUNTRY\_PL\`, \`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`,
\`COUNTRY\_PS\`, \`COUNTRY\_PT\`, \`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`,
\`COUNTRY\_RE\`, \`COUNTRY\_RO\`, \`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`,
\`COUNTRY\_SA\`, \`COUNTRY\_SB\`, \`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`,
\`COUNTRY\_SG\`, \`COUNTRY\_SH\`, \`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`,
\`COUNTRY\_SL\`, \`COUNTRY\_SM\`, \`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`,
\`COUNTRY\_SS\`, \`COUNTRY\_ST\`, \`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`,
\`COUNTRY\_SZ\`, \`COUNTRY\_TC\`, \`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`,
\`COUNTRY\_TH\`, \`COUNTRY\_TJ\`, \`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`,
\`COUNTRY\_TN\`, \`COUNTRY\_TO\`, \`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`,
\`COUNTRY\_TW\`, \`COUNTRY\_TZ\`, \`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`,
\`COUNTRY\_US\`, \`COUNTRY\_UY\`, \`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`,
\`COUNTRY\_VE\`, \`COUNTRY\_VG\`, \`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`,
\`COUNTRY\_WF\`, \`COUNTRY\_WS\`, \`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`,
\`COUNTRY\_YT\`, \`COUNTRY\_ZA\`, \`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Upstream description:

Sources that are located in one of the countries in the given list.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [ja4_tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-015.md#canonical-2022101130232011-2200203221023030-2232112210033120-2231121332031203-0320103110231212-0303033002020202-0213132313131332-0103023320030032): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-015.md#canonical-3221220000322331-1103032020303023-3303013302320311-3111100110122312-3021233331202213-0102203132223012-2320202201010211-0132302112322330): complete subsection reference.

<a id="canonical-2110333122133312-3311221103020013-3231113211010022-2331212013233320-1103230301021000-2331301202122321-2032330131101100-2321112210231000"></a>

## Next pages — ddos_client_source / 301332110033 / 5

- [ddos_mitigation_rules.ddos_client_source.asn_list](data-sources--http_loadbalancer--reference--group-015.md#canonical-3330032330013000-1211133133122010-2130122100120230-0100010212032131-2332210311132233-0101302230121001-1213033123113100-2332311202002232)
- [ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-015.md#canonical-2022101130232011-2200203221023030-2232112210033120-2231121332031203-0320103110231212-0303033002020202-0213132313131332-0103023320030032)
- [ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-015.md#canonical-3221220000322331-1103032020303023-3303013302320311-3111100110122312-3021233331202213-0102203132223012-2320202201010211-0132302112322330)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3330032330013000-1211133133122010-2130122100120230-0100010212032131-2332210311132233-0101302230121001-1213033123113100-2332311202002232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123101310123232-3220010020330321-3201221220311312-1203032231221111-2122221030120212-1123112013223133-0333220112133112-1030102203313013"></a>

## ddos_mitigation_rules.ddos_client_source.asn_list — asn_list / 333123312103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-015.md#canonical-2010132332231223-1321201300322332-1003012200330023-2322321131102132-2330033003320010-2133300030311102-1211110222320210-1201200010002203)
- ddos_mitigation_rules.ddos_client_source.asn_list

<a id="canonical-3030120101310201-3131323021220120-3002102113002201-0013130130220233-0130210121123012-1211312231321012-1301211121301000-2232130111211332"></a>

Type: `"single"`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3220222230022123-2222222201201010-1211111020331132-3232011002301220-3001020032213222-3330030231312232-3010201132002032-3203323030330130"></a>

## Direct properties — asn_list / 333123312103 / 3

<a id="canonical-1121022201203112-3300122201331132-1331003130111210-3211231332321222-0312202220333331-1133200331020111-2332332301303023-3333210120102220"></a>

<a id="canonical-3130130230100221-2100211320313313-1302201202321132-3110000021122310-1023001302003001-1003323200012101-3230133233101122-0011121201221102"></a>

## as_numbers property — asn_list / 333123312103 / 4

Type: `["list", "number"]`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-0220331220033220-1323332303103333-2111230311000100-2322013302123013-2132012220120203-2223031201102303-3113122332321202-3320022222130030"></a>

## Next pages — asn_list / 333123312103 / 5

- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-015.md#canonical-2010132332231223-1321201300322332-1003012200330023-2322321131102132-2330033003320010-2133300030311102-1211110222320210-1201200010002203)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2022101130232011-2200203221023030-2232112210033120-2231121332031203-0320103110231212-0303033002020202-0213132313131332-0103023320030032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233013322001311-0232020232330333-2003010212223223-3233231123013230-2212330130212321-3200222202031201-0322202122100110-1200302312100133"></a>

## ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher — ja4_tls_fingerprint_matcher / 231330303232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-015.md#canonical-2010132332231223-1321201300322332-1003012200330023-2322321131102132-2330033003320010-2133300030311102-1211110222320210-1201200010002203)
- ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher

<a id="canonical-1003033321212022-2320022013112220-3123033231332312-2010313110330210-0032313013233233-1102311232021301-3010000131303103-3110310312023311"></a>

Type: `"single"`. Computed.

Extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

Upstream description:

An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1113331003312213-3123300311103312-0120313311223110-2330331213232212-1203133303032233-0221031131221122-1130320320333000-0130233212320121"></a>

## Direct properties — ja4_tls_fingerprint_matcher / 231330303232 / 3

<a id="canonical-0110123013032300-0212020312223201-0311212211331220-1031103023220211-3011231222232213-3111301100133012-3121012222103202-1122000013033020"></a>

<a id="canonical-1101301012221011-1212001312001000-0320222231012011-1133013021012310-1123033023333232-3202032333133312-3111223322212232-1213003301131333"></a>

## exact_values property — ja4_tls_fingerprint_matcher / 231330303232 / 4

Type: `["list", "string"]`. Computed.

List of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Upstream description:

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0202012000313322-1033111130031231-0330130203003102-1133111002312113-0001221111110233-2133131212021212-3013023222332003-3102202321010111"></a>

## Next pages — ja4_tls_fingerprint_matcher / 231330303232 / 5

- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-015.md#canonical-2010132332231223-1321201300322332-1003012200330023-2322321131102132-2330033003320010-2133300030311102-1211110222320210-1201200010002203)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3221220000322331-1103032020303023-3303013302320311-3111100110122312-3021233331202213-0102203132223012-2320202201010211-0132302112322330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020312333301333-0100303023222310-0001232320223221-0122010232312221-1333021313023230-3012303030131022-3223102332230030-0322013310213330"></a>

## ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher — tls_fingerprint_matcher / 311313002023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-015.md#canonical-2010132332231223-1321201300322332-1003012200330023-2322321131102132-2330033003320010-2133300030311102-1211110222320210-1201200010002203)
- ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher

<a id="canonical-2201001123302320-2201213010210332-3220002202202121-2333202003320130-3001230102003121-2023231300103231-0213211333213020-2203303331200010"></a>

Type: `"single"`. Computed.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0302101110221203-1112002013201300-2212030113321311-0330222202203030-2000130301222312-1311333221330303-3221313300222303-2032201211232230"></a>

## Direct properties — tls_fingerprint_matcher / 311313002023 / 3

<a id="canonical-2123212230122013-0011230111323131-2102001210120213-0110230110231020-2313022021331012-0331013120110300-0212213323101131-0201310231131222"></a>

<a id="canonical-0213223011012331-0210012332311221-1032022322010301-1011300000301331-2212010032030213-2023133300222332-2012102302200221-3003311033213001"></a>

## classes property — tls_fingerprint_matcher / 311313002023 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1113221201032101-3121131303312211-3132102112031311-3213030303300021-1030330332102323-0013333103311021-3313203333232120-2031130120212311"></a>

<a id="canonical-0212030211122323-1023101031311023-3010003113310221-0310311131330000-0031000210023133-3022112332331223-0101222322302103-0313330323010211"></a>

## exact_values property — tls_fingerprint_matcher / 311313002023 / 5

Type: `["list", "string"]`. Computed.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0201300110310020-0021120302031133-1323111331131122-2101203031102332-0311120221222212-2133020323200303-2013301113000310-3321120211112130"></a>

<a id="canonical-0321231023203010-3201100231330303-2011332122301010-0232231231033032-1002311022331023-3123021020321123-2120300001020030-3231012021101231"></a>

## excluded_values property — tls_fingerprint_matcher / 311313002023 / 6

Type: `["list", "string"]`. Computed.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3321321023011213-3132320223332303-3332023303031121-2133102330001002-1232310232311222-3200302311110322-3222112332330122-2122100200330022"></a>

## Next pages — tls_fingerprint_matcher / 311313002023 / 7

- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-015.md#canonical-2010132332231223-1321201300322332-1003012200330023-2322321131102132-2330033003320010-2133300030311102-1211110222320210-1201200010002203)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2211231022121230-1011323212121001-2133133122021303-3332030210202013-2210323321121202-3323133220200310-1123101311301033-1102111103021310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103310032102032-0312302032320311-2221031133021120-0311212310100111-3111321313101030-3113210213131323-2122112323302103-2123222102032120"></a>

## ddos_mitigation_rules.ip_prefix_list — ip_prefix_list / 012131232031 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- ddos_mitigation_rules.ip_prefix_list

<a id="canonical-3201210333000031-1121021211121013-3302123100100212-3021231202223130-2103302202022131-2223203122303123-3011120202230210-0200221100302231"></a>

Type: `"single"`. Computed.

List of IP Prefix strings to match against.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1203313102130010-1332223211233023-2222201330111211-3012121332232110-1121223132001310-0002002322032101-3113003002110131-2133020320020202"></a>

## Direct properties — ip_prefix_list / 012131232031 / 3

<a id="canonical-1030010202323302-2323000311201012-3032102022000302-3301321013011101-2030000030020010-2313211011231130-1100013202131223-1313222120220200"></a>

<a id="canonical-2021322231330012-0220232313013000-2303311233332012-0123100203012102-2102312221000313-3101101222013102-0011130011200002-2321213232110003"></a>

## invert_match property — ip_prefix_list / 012131232031 / 4

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2031231122103033-0322320332231110-3203003212000333-0300023302123212-1231210323200203-3211120210033123-0023010122020331-2312302222232000"></a>

<a id="canonical-1213312222330101-2331022203110313-3312012102023211-1300200200113231-0222012122302210-1221133220013011-3232032000131212-3002032122231313"></a>

## ip_prefixes property — ip_prefix_list / 012131232031 / 5

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3320311113212130-3020310232121022-3213200103010222-3103201332322123-0030123221313321-0110132211130023-1011230302131312-1221101112321312"></a>

## Next pages — ip_prefix_list / 012131232031 / 6

- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3211223222113323-1020030013300003-0322233310230102-1333301120222330-2330300300132111-2030313102210023-3301233101112230-0011300100122012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203110000200300-1323130211032113-2000023203101221-0310231110031102-0021021222323033-1112123030313321-1031230211231233-3331011323110300"></a>

## ddos_mitigation_rules.metadata — metadata / 333113112311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- ddos_mitigation_rules.metadata

<a id="canonical-3301003113312330-2102210322322323-0302133112133201-2303333111222011-1302111133110310-0320201212312311-3303313302233111-3222032210101222"></a>

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

<a id="canonical-1331003211002222-2320311011012330-3112201200332122-1232333013030131-1321023211003322-3120132320132233-1211333312213312-1220011223011321"></a>

## Direct properties — metadata / 333113112311 / 3

<a id="canonical-1011232300100310-3030023011313311-1332312030300031-1330110202303322-1000311203100011-2323310301222023-1113311121213321-2323202210321323"></a>

<a id="canonical-2302013130130302-2100120021030312-0133320212101132-2200133110322110-3323332111213213-1311220113332222-0210100212122002-1132132122332210"></a>

## description_spec property — metadata / 333113112311 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0200102200220211-3121003332113211-0220212233230103-2121203103232231-1013113011123032-0002022112330330-2233113023231301-2031303020233213"></a>

<a id="canonical-3101032311032220-0111302033213322-3013323122122002-0131332120032221-0120220030020321-3011013013202203-0110011330031330-1111331133121000"></a>

## name property — metadata / 333113112311 / 5

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

<a id="canonical-1020023102323333-3032230132300111-0131211103230030-1123010001030310-0021013011113011-3102311020301003-0331031000231031-1301322202203000"></a>

## Next pages — metadata / 333113112311 / 6

- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320102310020121-1003113131120100-0201333101311300-1233103331122112-3002320003302310-2123202330323222-0012131322232332-0001311020013322"></a>

## default_pool — default_pool / 331000113103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- default_pool

<a id="canonical-2113122000131113-2331303012231120-1313102301020223-2110331202012112-1022100313100132-0300120312032331-3001102320000100-1333210220033320"></a>

Type: `"single"`. Computed.

\[OneOf: default\_pool, default\_pool\_list; Default: default\_pool\] Configuration parameter for
default pool.

Upstream description:

Shape of the origin pool specification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_port_choice": "[\"health_check_port\",\"same_as_endpoint_port\"]",
  "x-ves-oneof-field-port_choice": "[\"automatic_port\",\"lb_port\",\"port\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

OneOf alternatives in this subsection:

- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-2113122000131113-2331303012231120-1313102301020223-2110331202012112-1022100313100132-0300120312032331-3001102320000100-1333210220033320)
- [default_pool_list](data-sources--http_loadbalancer--reference--group-017.md#canonical-1130033132313333-3221230023302310-2121330233203033-0200002030033001-3110123103121011-1230303201033313-3011113212231322-1133312312033333)

Select alternatives according to the provider validators above.

<a id="canonical-0320310022202021-3201030130022211-3320102010231231-2103333312003320-3302202011333311-1202232300330012-0202010200311301-0302333330100112"></a>

## Direct properties — default_pool / 331000113103 / 3

- [advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013): complete subsection reference.

- [automatic_port](data-sources--http_loadbalancer--reference--group-015.md#canonical-3333201113312220-0211103023020231-1231022023021132-3130330021322021-1011333013330100-1221301202121312-1123312313021112-0032303012213003): complete subsection reference.

<a id="canonical-0202020300030231-3010110031021002-1021333103322333-2031332121030231-1012222012330000-0133002321230300-3131001032302122-3332300302303012"></a>

<a id="canonical-3222121322103120-3302200030310232-0011123313010203-3112113122203333-0212312003332112-3022230312110013-2030212111331221-1133000310133031"></a>

## endpoint_selection property — default_pool / 331000113103 / 4

Type: `"string"`. Computed.

\[Enum: DISTRIBUTED|LOCAL\_ONLY|LOCAL\_PREFERRED\] Policy for selection of endpoints from local
site/remote site/both Consider both remote and local endpoints for load balancing LOCAL\_ONLY:
Consider only local endpoints for load balancing Enable this policy to load balance ONLY among
locally discovered endpoints Prefer the local endpoints for.. Possible values are \`DISTRIBUTED\`,
\`LOCAL\_ONLY\`, \`LOCAL\_PREFERRED\`. Defaults to \`DISTRIBUTED\`. Server applies default when
omitted.

Upstream description:

Policy for selection of endpoints from local site/remote site/both

Consider both remote and local endpoints for load balancing LOCAL\_ONLY: Consider only local
endpoints for load balancing Enable this policy to load balance ONLY among locally discovered
endpoints Prefer the local endpoints for load balancing. If local endpoints are not present remote
endpoints will be considered.

Receipt-pinned upstream constraints:

```json
{
  "default": "DISTRIBUTED",
  "enum": [
    "DISTRIBUTED",
    "LOCAL_ONLY",
    "LOCAL_PREFERRED"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1103101202101131-2010321102031001-1203203131303113-1221132101310313-0110303222121002-2222301003103220-2200132132330123-2131100233300133"></a>

<a id="canonical-3131110110111002-1232332301213121-1021130131010221-0021201233213311-3020102331310203-0132203223300220-0011012211310323-0202110220330312"></a>

## health_check_port property — default_pool / 331000113103 / 5

Type: `"number"`. Computed.

Exclusive with \[same\_as\_endpoint\_port\] Port used for performing health check.

Upstream description:

Exclusive with \[same\_as\_endpoint\_port\] Port used for performing health check.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "category": "networking",
      "confidence": 0.99,
      "note": "Asymmetry: port enforces [1,65535], health_check_port allows 0",
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [healthcheck](data-sources--http_loadbalancer--reference--group-015.md#canonical-1202331312130223-3323033003031000-1300032201301212-2133102012302220-1012021120110103-2231131032011311-0213332212020202-2021031122333020): complete subsection reference.

- [lb_port](data-sources--http_loadbalancer--reference--group-015.md#canonical-2122312022231123-2122221001303110-1111302333031202-2020023322122331-0020010213103110-0012000101322211-1200121210030020-1320232013101313): complete subsection reference.

<a id="canonical-3201120301112012-3101322101210201-3103000300300123-2201300330100000-1031203121212222-2133232010120212-2020111233030003-1332020332311322"></a>

<a id="canonical-3131020313302000-3020010200101022-1221303303212021-0120313223121122-2303103132333303-2313310021102112-2212322211310321-2020213020123012"></a>

## loadbalancer_algorithm property — default_pool / 331000113103 / 6

Type: `"string"`. Computed.

\[Enum: ROUND\_ROBIN|LEAST\_REQUEST|RING\_HASH|RANDOM|LB\_OVERRIDE\] Different load balancing
algorithms supported When a connection to a endpoint in an upstream cluster is required, the load
balancer uses loadbalancer\_algorithm to determine which host is selected. - ROUND\_ROBIN:
ROUND\_ROBIN Policy in which each healthy/available upstream endpoint is selected in.. Possible
values are \`ROUND\_ROBIN\`, \`LEAST\_REQUEST\`, \`RING\_HASH\`, \`RANDOM\`, \`LB\_OVERRIDE\`.
Defaults to \`ROUND\_ROBIN\`. Server applies default when omitted.

Upstream description:

Different load balancing algorithms supported When a connection to a endpoint in an upstream cluster
is required, the load balancer uses loadbalancer\_algorithm to determine which host is selected.

&#8203;- ROUND\_ROBIN: ROUND\_ROBIN

Policy in which each healthy/available upstream endpoint is selected in round robin order. &#8203;-
LEAST\_REQUEST: LEAST\_REQUEST

Policy in which loadbalancer picks the upstream endpoint which has the fewest active requests
&#8203;- RING\_HASH: RING\_HASH

Policy implements consistent hashing to upstream endpoints using ring hash of endpoint names Hash of
the incoming request is calculated using request hash policy. The ring/modulo hash load balancer
implements consistent hashing to upstream hosts. The algorithm is based on mapping all hosts onto a
circle such that the addition or removal of a host from the host set changes only affect 1/N
requests. This technique is also commonly known as “ketama” hashing. A consistent hashing load
balancer is only effective when protocol routing is used that specifies a value to hash on. The
minimum ring size governs the replication factor for each host in the ring. For example, if the
minimum ring size is 1024 and there are 16 hosts, each host will be replicated 64 times. &#8203;-
RANDOM: RANDOM

Policy in which each available upstream endpoint is selected in random order. The random load
balancer selects a random healthy host. The random load balancer generally performs better than
round robin if no health checking policy is configured. Random selection avoids bias towards the
host in the set that comes after a failed host. &#8203;- LB\_OVERRIDE: Load Balancer Override

Hash policy is taken from from the load balancer which is using this origin pool.

Receipt-pinned upstream constraints:

```json
{
  "default": "ROUND_ROBIN",
  "enum": [
    "ROUND_ROBIN",
    "LEAST_REQUEST",
    "RING_HASH",
    "RANDOM",
    "LB_OVERRIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_tls](data-sources--http_loadbalancer--reference--group-015.md#canonical-1102212210131000-1000101230031101-0023123120321223-0123033000022131-0222321112012111-0001230102130023-2110102101300223-1203122202213132): complete subsection reference.

- [origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000): complete subsection reference.

<a id="canonical-2022100303132311-1033010221001301-0310233200232222-0002003122000012-1211022211021322-3002002021301230-0220302310311031-3010201002213020"></a>

<a id="canonical-1002203312121210-1311101001000221-2323200031220200-1112330212000313-3323213022012111-0213010020012212-1223313012132023-0303202000323031"></a>

## port property — default_pool / 331000113103 / 7

Type: `"number"`. Computed.

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port. Recommended:
\`443\`.

Upstream description:

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port.

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
    "create": false,
    "minimum_config": false,
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

- [same_as_endpoint_port](data-sources--http_loadbalancer--reference--group-016.md#canonical-1220231310023223-2021212032323232-0022231222211011-1010322033032201-0313321021233313-2333033123131031-3012210101010020-1022012223113311): complete subsection reference.

- [upstream_conn_pool_reuse_type](data-sources--http_loadbalancer--reference--group-016.md#canonical-0112303232111020-1021002100101202-3000002001230023-3321200302003202-1220310122032130-2003211001121303-1230131031213110-1111332002330012): complete subsection reference.

- [use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031): complete subsection reference.

- [view_internal](data-sources--http_loadbalancer--reference--group-017.md#canonical-2322220323320321-1023031031123201-0303333212100213-0313203031010132-0222033300002102-3131312330012013-2002323331211102-0113131333012333): complete subsection reference.

<a id="canonical-1230132132013310-1123231133133131-2301210310123110-0113130121232322-1301233310312301-1030223113212103-1300000131011031-0303000113111101"></a>

## Next pages — default_pool / 331000113103 / 8

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.automatic_port](data-sources--http_loadbalancer--reference--group-015.md#canonical-3333201113312220-0211103023020231-1231022023021132-3130330021322021-1011333013330100-1221301202121312-1123312313021112-0032303012213003)
- [default_pool.healthcheck](data-sources--http_loadbalancer--reference--group-015.md#canonical-1202331312130223-3323033003031000-1300032201301212-2133102012302220-1012021120110103-2231131032011311-0213332212020202-2021031122333020)
- [default_pool.lb_port](data-sources--http_loadbalancer--reference--group-015.md#canonical-2122312022231123-2122221001303110-1111302333031202-2020023322122331-0020010213103110-0012000101322211-1200121210030020-1320232013101313)
- [default_pool.no_tls](data-sources--http_loadbalancer--reference--group-015.md#canonical-1102212210131000-1000101230031101-0023123120321223-0123033000022131-0222321112012111-0001230102130023-2110102101300223-1203122202213132)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.same_as_endpoint_port](data-sources--http_loadbalancer--reference--group-016.md#canonical-1220231310023223-2021212032323232-0022231222211011-1010322033032201-0313321021233313-2333033123131031-3012210101010020-1022012223113311)
- [default_pool.upstream_conn_pool_reuse_type](data-sources--http_loadbalancer--reference--group-016.md#canonical-0112303232111020-1021002100101202-3000002001230023-3321200302003202-1220310122032130-2003211001121303-1230131031213110-1111332002330012)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031)
- [default_pool.view_internal](data-sources--http_loadbalancer--reference--group-017.md#canonical-2322220323320321-1023031031123201-0303333212100213-0313203031010132-0222033300002102-3131312330012013-2002323331211102-0113131333012333)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300201212313221-1003000321321132-0120022230001130-2022013331212333-3331302332102121-2011233212220113-1331013102201302-2033310231003120"></a>

## default_pool.advanced_options — advanced_options / 313312011300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- default_pool.advanced_options

<a id="canonical-1213011323322132-3010010030122001-3022303000113300-2213101320030312-1221223103103220-2311330332110122-0302030120203210-0310033332231120"></a>

Type: `"single"`. Computed.

Configure Advanced OPTIONS for origin pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-circuit_breaker_choice": "[\"circuit_breaker\",\"default_circuit_breaker\",\"disable_circuit_breaker\"]",
  "x-ves-oneof-field-http_protocol_type": "[\"auto_http_config\",\"http1_config\",\"http2_options\"]",
  "x-ves-oneof-field-lb_source_ip_persistence_choice": "[\"disable_lb_source_ip_persistence\",\"enable_lb_source_ip_persistence\"]",
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-outlier_detection_choice": "[\"disable_outlier_detection\",\"outlier_detection\"]",
  "x-ves-oneof-field-panic_threshold_type": "[\"no_panic_threshold\",\"panic_threshold\"]",
  "x-ves-oneof-field-proxy_protocol_choice": "[\"disable_proxy_protocol\",\"proxy_protocol_v1\",\"proxy_protocol_v2\"]",
  "x-ves-oneof-field-subset_choice": "[\"disable_subsets\",\"enable_subsets\"]"
}
```

<a id="canonical-1330203002131101-2332103232333201-3230103133232323-1002023311102101-1233332221332020-0110111123003101-3002001101200113-0331102301311003"></a>

## Direct properties — advanced_options / 313312011300 / 3

- [auto_http_config](data-sources--http_loadbalancer--reference--group-015.md#canonical-1130221102311203-0231221321000220-1010101302131120-3122131031311211-1022130022230031-1203101122023121-2310200000010220-1133033110333012): complete subsection reference.

- [circuit_breaker](data-sources--http_loadbalancer--reference--group-015.md#canonical-0232020321002221-1233132132022200-0013031023102223-3210100000312120-0101013022202003-0033322120303102-3303002223033210-3201111133320202): complete subsection reference.

<a id="canonical-0132330233100200-0211310020333001-3302031320000121-2211022212210102-3022321330300300-3312000212113220-0022001203002230-3211011112121111"></a>

<a id="canonical-0130101123322211-0200121112320032-3011113023003311-2202222010222223-2130200121121132-0202330211133112-1022301330233033-3220333123201203"></a>

## connection_timeout property — advanced_options / 313312011300 / 4

Type: `"number"`. Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds. Server applies default when omitted. Recommended:
\`2000\`.

Upstream description:

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1800000,
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
    "ves.io.schema.rules.uint32.lte": "1800000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  }
}
```

- [default_circuit_breaker](data-sources--http_loadbalancer--reference--group-015.md#canonical-1132203130230100-0021230222302023-1012032232112121-3112301022311233-1010303330331103-0121010203330113-2100020332123203-1302311223302003): complete subsection reference.

- [disable_circuit_breaker](data-sources--http_loadbalancer--reference--group-015.md#canonical-1103122131020223-0233032003331303-1203212013001231-0131222300130203-3021303201233003-1211300321313311-2102120320200231-0201002130102103): complete subsection reference.

- [disable_lb_source_ip_persistence](data-sources--http_loadbalancer--reference--group-015.md#canonical-3020030101330100-1312220330202101-0111001313103013-0331013030011111-3112230002222211-1112210013113123-3220132020200020-2031212222123123): complete subsection reference.

- [disable_outlier_detection](data-sources--http_loadbalancer--reference--group-015.md#canonical-1332012311033103-1321022311010313-0101213221232232-0132003113002213-3021102132022102-0220110203312233-3301313112203210-1011220103310130): complete subsection reference.

- [disable_proxy_protocol](data-sources--http_loadbalancer--reference--group-015.md#canonical-1030313332320200-1303011110223202-3320200013010123-1213002201221302-0322320302210223-3011323333221012-1011213301321111-1030131230021230): complete subsection reference.

- [disable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-0321332300211123-3230221211223301-3331112313102200-1103202311312112-3110110211112113-1131120230022300-2112221130003120-3310200211021222): complete subsection reference.

- [enable_lb_source_ip_persistence](data-sources--http_loadbalancer--reference--group-015.md#canonical-3020213032033211-0032113103233022-0223003002111333-3021113232030003-2300200121010101-0213221210001011-3032132003022122-2110211110222232): complete subsection reference.

- [enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003): complete subsection reference.

- [http1_config](data-sources--http_loadbalancer--reference--group-015.md#canonical-2012200011203112-1330213223231102-1210101031311213-0110120103020111-2112000010131232-3013022312210133-1322001232311220-1330102303301023): complete subsection reference.

- [http2_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-0302103210203123-2300200322030323-3303020103012020-2031120223010323-2313200232200231-3203333211320133-2033203220101011-1023222100033000): complete subsection reference.

<a id="canonical-2000112112220300-3312013130222213-2003213001031320-2021331312022232-0020032120210012-2013201303203033-2210013021201122-2211121030301311"></a>

<a id="canonical-1003112301113222-1232113321031002-2221120220230023-1020003032202332-1332233003032033-2123312113303000-0303100033213111-3210201330120102"></a>

## http_idle_timeout property — advanced_options / 313312011300 / 5

Type: `"number"`. Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Server applies default when omitted. Recommended: \`300000\`.

Upstream description:

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.
This is specified in milliseconds. The default value is 5 minutes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-0112210022031110-3312312021113132-2030002010031131-3201013000230020-3023022231223312-1032133123013111-0012220011130020-1033221122133311"></a>

<a id="canonical-0301101213212233-2312111232121222-1123332001130131-3232132021031123-1231101121132331-1300121012213331-0310332001001312-2300313012310103"></a>

## max_requests_per_connection property — advanced_options / 313312011300 / 6

Type: `"number"`. Computed.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [no_panic_threshold](data-sources--http_loadbalancer--reference--group-015.md#canonical-0212013320121110-3033210311131000-3201212002231122-1110302332211220-0112213030302123-2320021233211303-0110130210300021-1323000111111223): complete subsection reference.

- [no_request_limit_per_connection](data-sources--http_loadbalancer--reference--group-015.md#canonical-1321012101122211-2102000320220313-1331310322231032-1203123320313111-1311130321320303-1020331231312103-1330303100211131-0032021020020300): complete subsection reference.

- [outlier_detection](data-sources--http_loadbalancer--reference--group-015.md#canonical-0233100000211130-1311321230320220-1220101003130100-1202010131212223-3301032010002020-3231201110022231-3002222332232320-2213210303031120): complete subsection reference.

<a id="canonical-0330333010122110-1021103333321221-1023201000303013-2112002023222000-3001030321111022-1213023030031231-2230232331211023-1223102123322000"></a>

<a id="canonical-1200323230120011-1101110312232031-1002130121223020-0012333231032312-3101212310311330-2312302032323022-2222302301020100-2021202013331311"></a>

## panic_threshold property — advanced_options / 313312011300 / 7

Type: `"number"`. Computed.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for load balancing ignoring its health status.

Upstream description:

Exclusive with \[no\_panic\_threshold\]

Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be
considered for load balancing ignoring its health status.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [proxy_protocol_v1](data-sources--http_loadbalancer--reference--group-015.md#canonical-1320312022011223-0112101033132110-2322303210023201-1320121013000330-1331321233310011-3220111102322200-3021312012310331-0301233032003001): complete subsection reference.

- [proxy_protocol_v2](data-sources--http_loadbalancer--reference--group-015.md#canonical-0211103313321032-0201122021031121-2311101000312300-2230110230320310-3213220020320023-2010000121013233-2311011010210301-3323120323303331): complete subsection reference.

<a id="canonical-0031212222302010-0333120202003012-3331031231323102-1302310100102110-1003211332123212-3121103103221132-0101210201231233-2020312120312112"></a>

## Next pages — advanced_options / 313312011300 / 8

- [default_pool.advanced_options.auto_http_config](data-sources--http_loadbalancer--reference--group-015.md#canonical-1130221102311203-0231221321000220-1010101302131120-3122131031311211-1022130022230031-1203101122023121-2310200000010220-1133033110333012)
- [default_pool.advanced_options.circuit_breaker](data-sources--http_loadbalancer--reference--group-015.md#canonical-0232020321002221-1233132132022200-0013031023102223-3210100000312120-0101013022202003-0033322120303102-3303002223033210-3201111133320202)
- [default_pool.advanced_options.default_circuit_breaker](data-sources--http_loadbalancer--reference--group-015.md#canonical-1132203130230100-0021230222302023-1012032232112121-3112301022311233-1010303330331103-0121010203330113-2100020332123203-1302311223302003)
- [default_pool.advanced_options.disable_circuit_breaker](data-sources--http_loadbalancer--reference--group-015.md#canonical-1103122131020223-0233032003331303-1203212013001231-0131222300130203-3021303201233003-1211300321313311-2102120320200231-0201002130102103)
- [default_pool.advanced_options.disable_lb_source_ip_persistence](data-sources--http_loadbalancer--reference--group-015.md#canonical-3020030101330100-1312220330202101-0111001313103013-0331013030011111-3112230002222211-1112210013113123-3220132020200020-2031212222123123)
- [default_pool.advanced_options.disable_outlier_detection](data-sources--http_loadbalancer--reference--group-015.md#canonical-1332012311033103-1321022311010313-0101213221232232-0132003113002213-3021102132022102-0220110203312233-3301313112203210-1011220103310130)
- [default_pool.advanced_options.disable_proxy_protocol](data-sources--http_loadbalancer--reference--group-015.md#canonical-1030313332320200-1303011110223202-3320200013010123-1213002201221302-0322320302210223-3011323333221012-1011213301321111-1030131230021230)
- [default_pool.advanced_options.disable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-0321332300211123-3230221211223301-3331112313102200-1103202311312112-3110110211112113-1131120230022300-2112221130003120-3310200211021222)
- [default_pool.advanced_options.enable_lb_source_ip_persistence](data-sources--http_loadbalancer--reference--group-015.md#canonical-3020213032033211-0032113103233022-0223003002111333-3021113232030003-2300200121010101-0213221210001011-3032132003022122-2110211110222232)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003)
- [default_pool.advanced_options.http1_config](data-sources--http_loadbalancer--reference--group-015.md#canonical-2012200011203112-1330213223231102-1210101031311213-0110120103020111-2112000010131232-3013022312210133-1322001232311220-1330102303301023)
- [default_pool.advanced_options.http2_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-0302103210203123-2300200322030323-3303020103012020-2031120223010323-2313200232200231-3203333211320133-2033203220101011-1023222100033000)
- [default_pool.advanced_options.no_panic_threshold](data-sources--http_loadbalancer--reference--group-015.md#canonical-0212013320121110-3033210311131000-3201212002231122-1110302332211220-0112213030302123-2320021233211303-0110130210300021-1323000111111223)
- [default_pool.advanced_options.no_request_limit_per_connection](data-sources--http_loadbalancer--reference--group-015.md#canonical-1321012101122211-2102000320220313-1331310322231032-1203123320313111-1311130321320303-1020331231312103-1330303100211131-0032021020020300)
- [default_pool.advanced_options.outlier_detection](data-sources--http_loadbalancer--reference--group-015.md#canonical-0233100000211130-1311321230320220-1220101003130100-1202010131212223-3301032010002020-3231201110022231-3002222332232320-2213210303031120)
- [default_pool.advanced_options.proxy_protocol_v1](data-sources--http_loadbalancer--reference--group-015.md#canonical-1320312022011223-0112101033132110-2322303210023201-1320121013000330-1331321233310011-3220111102322200-3021312012310331-0301233032003001)
- [default_pool.advanced_options.proxy_protocol_v2](data-sources--http_loadbalancer--reference--group-015.md#canonical-0211103313321032-0201122021031121-2311101000312300-2230110230320310-3213220020320023-2010000121013233-2311011010210301-3323120323303331)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1130221102311203-0231221321000220-1010101302131120-3122131031311211-1022130022230031-1203101122023121-2310200000010220-1133033110333012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332103003112103-2012003120220320-2313310201120122-2202321212310330-3210303131323311-2122100132131102-2022301332302001-1103022030003303"></a>

## default_pool.advanced_options.auto_http_config — auto_http_config / 332202310303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.auto_http_config

<a id="canonical-2303133030223031-0313032012122100-1231110001311200-0000001303113020-2222022002213112-3232300113033033-1301131010133312-0012202300023130"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-2310001123120002-0032121223211230-0030131212313003-0000301200121303-3110111111013230-3200100200221132-1220202310021031-2011233211210333"></a>

## Direct properties — auto_http_config / 332202310303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101202323202310-0300313210100102-2311300223331313-0211110022132311-1132301032033132-0101211211110012-3101103221011232-0302222113232010"></a>

## Next pages — auto_http_config / 332202310303 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0232020321002221-1233132132022200-0013031023102223-3210100000312120-0101013022202003-0033322120303102-3303002223033210-3201111133320202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030121300032332-3230332202211013-1301030232220033-0031333113013122-3121230321113331-0213303200113103-3333001203001021-0022013332310021"></a>

## default_pool.advanced_options.circuit_breaker — circuit_breaker / 102313200330 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.circuit_breaker

<a id="canonical-0121223110230212-3121203013333032-1332033013002311-2331112200111322-1211332213231312-2332312321132130-2300010011023322-2233023320003310"></a>

Type: `"single"`. Computed.

CircuitBreaker provides a mechanism for watching failures in upstream connections or requests and if
the failures reach a certain threshold, automatically fail subsequent requests which allows to apply
back pressure on downstream quickly.

Upstream description:

CircuitBreaker provides a mechanism for watching failures in upstream connections or requests and if
the failures reach a certain threshold, automatically fail subsequent requests which allows to apply
back pressure on downstream quickly.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2310231333321022-2023312212310230-1133200221212121-1200012100313321-3120023130212300-3110132122301203-2322113200233022-2121331033222302"></a>

## Direct properties — circuit_breaker / 102313200330 / 3

<a id="canonical-0323330021323031-2122112203011211-2012120120021210-0300011020122120-3110221322121010-2333222133221303-3232301102202113-0210121203120221"></a>

<a id="canonical-3302032131133313-3211311100211233-2300121020133320-1310320132322300-0003001310131132-3210031003301030-0330220201200033-2012000313003112"></a>

## connection_limit property — circuit_breaker / 102313200330 / 4

Type: `"number"`. Computed.

The maximum number of connections that loadbalancer will establish to all hosts in an upstream
cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single
connection to each host. Remove endpoint out of load balancing decision, if number of connections..

Upstream description:

The maximum number of connections that loadbalancer will establish to all hosts in an upstream
cluster. In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single
connection to each host. Remove endpoint out of load balancing decision, if number of connections
reach connection limit.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-1311203223221131-1213203021321233-0210011133223123-2102002123311320-2132213110311003-2303211212032233-3330133230113111-1233122131303110"></a>

<a id="canonical-1132301300212303-2220022031000221-0030300123030123-2202123122232111-3202100000103100-1013313223101200-2331203333121233-1123120313131332"></a>

## max_requests property — circuit_breaker / 102313200330 / 5

Type: `"number"`. Computed.

The maximum number of requests that can be outstanding to all hosts in a cluster at any given time.
In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the
maximum connections (connection\_limit). Remove endpoint out of load balancing decision, if
requests..

Upstream description:

The maximum number of requests that can be outstanding to all hosts in a cluster at any given time.
In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the
maximum connections (connection\_limit). Remove endpoint out of load balancing decision, if requests
exceed this count.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-1003030011012210-2320001222001202-1213121213102020-2200002222010211-0330210212303012-0323310201213030-3320321010202102-0031113000100321"></a>

<a id="canonical-1010333032132010-1013323302103032-0033133313302100-0032020002020323-1100210212101023-3030102122103300-2331331011210101-2301303120302130"></a>

## pending_requests property — circuit_breaker / 102313200330 / 6

Type: `"number"`. Computed.

The maximum number of requests that will be queued while waiting for a ready connection pool
connection. Since HTTP/2 requests are sent over a single connection, this circuit breaker only comes
into play as the initial connection is created, as requests will be multiplexed immediately..

Upstream description:

The maximum number of requests that will be queued while waiting for a ready connection pool
connection. Since HTTP/2 requests are sent over a single connection, this circuit breaker only comes
into play as the initial connection is created, as requests will be multiplexed immediately
afterwards. For HTTP/1.1, requests are added to the list of pending requests whenever there aren’t
enough upstream connections available to immediately dispatch the request, so this circuit breaker
will remain in play for the lifetime of the process. Remove endpoint out of load balancing decision,
if pending request reach pending\_request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-0210321322231000-3302132111022212-3013100033130330-0333102223330112-1031000202313323-2300311211101132-2321112113302313-1012311012231213"></a>

<a id="canonical-1122231101020330-0100122221030103-0233002213312223-3130110012101112-0212223001013031-0233022211312312-0303031330020322-1112132033121033"></a>

## priority property — circuit_breaker / 102313200330 / 7

Type: `"string"`. Computed.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Upstream description:

Priority routing for each request. Different connection pools are used based on the priority
selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based
on selected priority.

Default routing mechanism High-Priority routing mechanism.

Receipt-pinned upstream constraints:

```json
{
  "default": "DEFAULT",
  "enum": [
    "DEFAULT",
    "HIGH"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0021330111130221-3011231023023132-1221103110100320-3003011111101211-0002103313332311-0132313030031203-0132112212200230-3230110311000301"></a>

<a id="canonical-3132200311110323-0200132100301010-1332001112022213-1020020203012233-3220203101122212-2021020230012013-1333011201131001-1331022022023000"></a>

## retries property — circuit_breaker / 102313200330 / 8

Type: `"number"`. Computed.

The maximum number of retries that can be outstanding to all hosts in a cluster at any given time.
Remove endpoint out of load balancing decision, if retries for request exceed this count.

Upstream description:

The maximum number of retries that can be outstanding to all hosts in a cluster at any given time.
Remove endpoint out of load balancing decision, if retries for request exceed this count.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

<a id="canonical-0031221200110203-1232002300002111-2013232032333032-0222033322311123-1122230232002021-0112031300233233-1220322012203010-2301213110310322"></a>

## Next pages — circuit_breaker / 102313200330 / 9

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1132203130230100-0021230222302023-1012032232112121-3112301022311233-1010303330331103-0121010203330113-2100020332123203-1302311223302003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122221130332231-2000030012033113-0021320210111030-0210222123000200-0223222132003131-1230222103300210-2313003223220313-2020032132023202"></a>

## default_pool.advanced_options.default_circuit_breaker — default_circuit_breaker / 310330001202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.default_circuit_breaker

<a id="canonical-3323333213022201-3233113132221222-2212302201312103-0001322321332010-1130211102211013-3021311121110300-2011321012311203-2022020332320002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default circuit breaker. Defaults to \`map\[\]\`. Server applies default
when omitted.

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

<a id="canonical-3212312113011001-3000022300121322-1213022121230213-3022312203300323-0220013302121232-3331302022330102-3102003202231320-3312121033211133"></a>

## Direct properties — default_circuit_breaker / 310330001202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103332213001323-1001101330313301-3222222222201302-1113233130113210-2021123212022300-1133032222202102-0031112101001331-3003213011311210"></a>

## Next pages — default_circuit_breaker / 310330001202 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1103122131020223-0233032003331303-1203212013001231-0131222300130203-3021303201233003-1211300321313311-2102120320200231-0201002130102103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132200330232322-2202012232223010-1100321011002122-2210312331203103-3123330223132223-1310112030132223-0200022123122101-2020030032100310"></a>

## default_pool.advanced_options.disable_circuit_breaker — disable_circuit_breaker / 023323013032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.disable_circuit_breaker

<a id="canonical-0133311100122000-2110323231233301-3310232123110330-2113213201100121-3022330322021003-1033232311103210-3302113000333020-0320100102120002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable circuit breaker.

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

<a id="canonical-1002322222233133-2312332212031232-0010000312000330-2103221311012221-2313033113003221-3321132033103223-2002303213121010-1203012303110302"></a>

## Direct properties — disable_circuit_breaker / 023323013032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303310033033023-1222022103003122-2023222231001000-0210232110222212-2001301320130300-2332102222033310-0323010231000002-3211113302323223"></a>

## Next pages — disable_circuit_breaker / 023323013032 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3020030101330100-1312220330202101-0111001313103013-0331013030011111-3112230002222211-1112210013113123-3220132020200020-2031212222123123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013200202201202-2022112012233302-0310302120310122-0132301021231002-3122233200313030-0331103011102022-3231230122010331-1000220010212001"></a>

## default_pool.advanced_options.disable_lb_source_ip_persistence — disable_lb_source_ip_persistence / 222131300110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.disable_lb_source_ip_persistence

<a id="canonical-1111123310003113-3230003100000131-3023333111311210-1222201302212032-0111320222212121-3111101223302221-2023023033021331-3233201020101022"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

IP address configuration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1323211100002113-0021312003232200-0022232113300332-0030003323000123-1001301201310220-2013123311021210-1111133102312132-0231000322133112"></a>

## Direct properties — disable_lb_source_ip_persistence / 222131300110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032103201123210-3120110102231313-0011221312202200-1231302102230211-3130122000203321-0130030221033300-0320032310322223-3210122230232211"></a>

## Next pages — disable_lb_source_ip_persistence / 222131300110 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1332012311033103-1321022311010313-0101213221232232-0132003113002213-3021102132022102-0220110203312233-3301313112203210-1011220103310130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132303133110112-3321322331212113-1111120111123103-3032001322033301-0020200110013202-1321322330122303-1301220132001022-0320231211031303"></a>

## default_pool.advanced_options.disable_outlier_detection — disable_outlier_detection / 230301322220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.disable_outlier_detection

<a id="canonical-0111323303223211-0301331211031012-0331100132212102-2321333223310231-1030132023201032-3333100123331321-3123020333110322-1231230323112222"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable outlier detection. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

<a id="canonical-0230031132023100-0133121020333323-2021220110013203-1313023131211202-1001013103131222-1232123230230023-2113331312231232-1312000001230013"></a>

## Direct properties — disable_outlier_detection / 230301322220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133323323320132-1213312202320123-2202332210330001-3131310131312010-0222201030210110-1021031302210021-2331012312321112-3303300313333323"></a>

## Next pages — disable_outlier_detection / 230301322220 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1030313332320200-1303011110223202-3320200013010123-1213002201221302-0322320302210223-3011323333221012-1011213301321111-1030131230021230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222210212103201-2001002310112113-2233113122322033-1203220030101010-1131232133010112-3031101313211130-3220203320021020-0130000033223102"></a>

## default_pool.advanced_options.disable_proxy_protocol — disable_proxy_protocol / 223322132130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.disable_proxy_protocol

<a id="canonical-2112322110223113-1231230330100332-0222010013223310-0311302311301013-1221231321103232-3113303221011303-2203301322320033-3301121013201302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable proxy protocol.

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

<a id="canonical-1200203300112312-3133223210211131-1102032323001333-0212113231223100-1113310020313302-0031312120123102-1003301313320010-3000312003331020"></a>

## Direct properties — disable_proxy_protocol / 223322132130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300013101033022-1202302111200230-3033010302101133-3030120032121202-2011333030101321-0021333000003221-3111321311111033-2233032030120122"></a>

## Next pages — disable_proxy_protocol / 223322132130 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0321332300211123-3230221211223301-3331112313102200-1103202311312112-3110110211112113-1131120230022300-2112221130003120-3310200211021222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022330210331110-2200003221032013-0113230302022020-2230211231003322-0031313132303303-0103201030132202-2211111130202221-1023002210121233"></a>

## default_pool.advanced_options.disable_subsets — disable_subsets / 021320003212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.disable_subsets

<a id="canonical-2222333131110113-0130133330011002-2331330103022231-1011322200100000-2210133113300022-3031030030322023-3110212221013310-1020201221232103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable subsets. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-3132022010023123-2213322320313122-0033121002123322-0003102112212011-0033210220011310-0000321330223302-0231302113000000-1302012033101122"></a>

## Direct properties — disable_subsets / 021320003212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221323010103022-1032131332113102-1323111122223212-1022331333133200-1101211010121233-1123012132323123-2313303110211200-2123102303323111"></a>

## Next pages — disable_subsets / 021320003212 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3020213032033211-0032113103233022-0223003002111333-3021113232030003-2300200121010101-0213221210001011-3032132003022122-2110211110222232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013201203230311-3100333131032130-1210022213021003-3013322313311333-3020220111320013-1022013302031311-0213121230110222-2110131033320300"></a>

## default_pool.advanced_options.enable_lb_source_ip_persistence — enable_lb_source_ip_persistence / 232122320020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.enable_lb_source_ip_persistence

<a id="canonical-2110330023312300-2113130223210221-0023220220113303-2021103302133212-3010322101201221-3100301333103321-1100020022203310-2201010103302003"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

IP address configuration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0321221312322113-2022221102133011-1232332301201110-1121312120110330-0111222133321122-0332303330320123-1221002330121332-0220001121220010"></a>

## Direct properties — enable_lb_source_ip_persistence / 232122320020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301011300033112-2232000023010300-2331010122330032-2002033311013000-2321202220010123-3110311203113120-0213021231013230-0002123101310000"></a>

## Next pages — enable_lb_source_ip_persistence / 232122320020 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032033231111003-1103010133110130-2200213211230301-2332230321333232-0120130023220113-1121131320233210-2022110012130020-1320103003312001"></a>

## default_pool.advanced_options.enable_subsets — enable_subsets / 201233332003 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.enable_subsets

<a id="canonical-0022320222203203-2202021010102323-3113301101302003-2222130322332132-0231311331202012-1020220022010203-2002120312330120-3110223100223201"></a>

Type: `"single"`. Computed.

Configure subset OPTIONS for origin pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fallback_policy_choice": "[\"any_endpoint\",\"default_subset\",\"fail_request\"]"
}
```

<a id="canonical-1232231310212113-1312310300120311-1031123032301111-0230023312323000-3003122231320013-3110201232300023-2331033220031101-0112032133021120"></a>

## Direct properties — enable_subsets / 201233332003 / 3

- [any_endpoint](data-sources--http_loadbalancer--reference--group-015.md#canonical-0033000210210213-1033200310020203-0323331100133030-2000220212322321-3103221110220030-0120323132332102-1302311330200321-2003320122230112): complete subsection reference.

- [default_subset](data-sources--http_loadbalancer--reference--group-015.md#canonical-3221030202032311-3023220331221201-0010011333222003-3130103002003322-1120311323130221-2112030023023112-0331303202211111-0202213322011212): complete subsection reference.

- [endpoint_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-2121322022211212-1333302100133030-0301201230333113-3032010200012231-1121031001010220-0233210000022211-1200323002312122-2132002133332223): complete subsection reference.

- [fail_request](data-sources--http_loadbalancer--reference--group-015.md#canonical-1021103000030322-0131231321303110-3133120333303230-1013320313021013-1320103003030210-0130323110332113-0022001201131110-3003322320132213): complete subsection reference.

<a id="canonical-0230211333022020-0130010020031021-0211221333100233-3221112103022013-3230220010013210-3312231221233202-0213220101233203-3320120212331033"></a>

## Next pages — enable_subsets / 201233332003 / 4

- [default_pool.advanced_options.enable_subsets.any_endpoint](data-sources--http_loadbalancer--reference--group-015.md#canonical-0033000210210213-1033200310020203-0323331100133030-2000220212322321-3103221110220030-0120323132332102-1302311330200321-2003320122230112)
- [default_pool.advanced_options.enable_subsets.default_subset](data-sources--http_loadbalancer--reference--group-015.md#canonical-3221030202032311-3023220331221201-0010011333222003-3130103002003322-1120311323130221-2112030023023112-0331303202211111-0202213322011212)
- [default_pool.advanced_options.enable_subsets.endpoint_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-2121322022211212-1333302100133030-0301201230333113-3032010200012231-1121031001010220-0233210000022211-1200323002312122-2132002133332223)
- [default_pool.advanced_options.enable_subsets.fail_request](data-sources--http_loadbalancer--reference--group-015.md#canonical-1021103000030322-0131231321303110-3133120333303230-1013320313021013-1320103003030210-0130323110332113-0022001201131110-3003322320132213)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0033000210210213-1033200310020203-0323331100133030-2000220212322321-3103221110220030-0120323132332102-1302311330200321-2003320122230112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300301322302320-0122110112200231-1132113201131032-2131132031222320-0312232310032133-3212302301201100-3202212001131220-2123130113300230"></a>

## default_pool.advanced_options.enable_subsets.any_endpoint — any_endpoint / 221330002210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003)
- default_pool.advanced_options.enable_subsets.any_endpoint

<a id="canonical-0311300213223111-2201202121002221-0122000013231000-0303221003130133-2130123231122221-3220201101202122-0122121112213103-2321300321003022"></a>

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

<a id="canonical-2330122231321210-1203310003113111-0300201333333233-0133300321011101-0101023002030300-3102103222102220-1113123310320232-0110213322020322"></a>

## Direct properties — any_endpoint / 221330002210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322201322302133-0112012123012300-0323103312221132-0112200332232303-0321000220230111-2002122101103232-0223010200202100-0313032221210102"></a>

## Next pages — any_endpoint / 221330002210 / 4

- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3221030202032311-3023220331221201-0010011333222003-3130103002003322-1120311323130221-2112030023023112-0331303202211111-0202213322011212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000120121133102-3301110212121103-1203302322121002-1003001323213013-2320303032022000-0122333120132110-2222011121213232-3312223310102310"></a>

## default_pool.advanced_options.enable_subsets.default_subset — default_subset / 022003002022 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003)
- default_pool.advanced_options.enable_subsets.default_subset

<a id="canonical-2220211111101212-2133122301022030-3320221133101133-2312200212330211-3130123022203321-1112301220030203-2230323121103100-3211002303222203"></a>

Type: `"single"`. Computed.

Configuration parameter for default subset.

Upstream description:

Default Subset definition.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0120311031113022-3131101310112003-3232221302222232-0021312030332301-0023212001303112-2222123330101231-2230012110032221-2021103212122223"></a>

## Direct properties — default_subset / 022003002022 / 3

- [default_subset](data-sources--http_loadbalancer--reference--group-015.md#canonical-0120210320232100-2212123202222221-3213201133300210-1121313230211211-3331230033212003-2321021021210032-3023030110313013-1120310032212120): complete subsection reference.

<a id="canonical-3220002030302200-3220002032223120-1031310302313201-1023302201312301-2233332322202230-1033113211200322-1001210300102201-2131300122323022"></a>

## Next pages — default_subset / 022003002022 / 4

- [default_pool.advanced_options.enable_subsets.default_subset.default_subset](data-sources--http_loadbalancer--reference--group-015.md#canonical-0120210320232100-2212123202222221-3213201133300210-1121313230211211-3331230033212003-2321021021210032-3023030110313013-1120310032212120)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0120210320232100-2212123202222221-3213201133300210-1121313230211211-3331230033212003-2321021021210032-3023030110313013-1120310032212120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212032211101132-1202331002322132-2133003010113211-2003112313313100-3222102333212332-2103002031332113-3323302113230203-3332301013113022"></a>

## default_pool.advanced_options.enable_subsets.default_subset.default_subset — default_subset / 322012222002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003)
- [default_pool.advanced_options.enable_subsets.default_subset](data-sources--http_loadbalancer--reference--group-015.md#canonical-3221030202032311-3023220331221201-0010011333222003-3130103002003322-1120311323130221-2112030023023112-0331303202211111-0202213322011212)
- default_pool.advanced_options.enable_subsets.default_subset.default_subset

<a id="canonical-1301120023233123-0110021332212222-1322121213131100-2321322300122233-1012010110101021-2012230201020301-3302002330022231-3221111223032333"></a>

Type: `"single"`. Computed.

List of key-value pairs that define default subset. Which gets used when route specifies no metadata
or no subset matching the metadata exists.

Upstream description:

List of key-value pairs that define default subset. Which gets used when route specifies no metadata
or no subset matching the metadata exists.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 32
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "32"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "32"
  }
}
```

<a id="canonical-2202211131120300-1030311011100020-0221000223012023-2201321221012020-3110122110131002-1333122100311223-1210113330120332-1320002301202332"></a>

## Direct properties — default_subset / 322012222002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003020012222002-1312231023101132-1110000001203302-0310033232013102-2212310212312321-2313011312231110-0121312022320230-3122120101311320"></a>

## Next pages — default_subset / 322012222002 / 4

- [default_pool.advanced_options.enable_subsets.default_subset](data-sources--http_loadbalancer--reference--group-015.md#canonical-3221030202032311-3023220331221201-0010011333222003-3130103002003322-1120311323130221-2112030023023112-0331303202211111-0202213322011212)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2121322022211212-1333302100133030-0301201230333113-3032010200012231-1121031001010220-0233210000022211-1200323002312122-2132002133332223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131300101120113-0322232122302010-0311302301010312-2322101020222012-0003332131321332-0132231321010300-2232121012123320-0131113033231032"></a>

## default_pool.advanced_options.enable_subsets.endpoint_subsets — endpoint_subsets / 221302030212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003)
- default_pool.advanced_options.enable_subsets.endpoint_subsets

<a id="canonical-3201033213033012-0332321133203303-3230001303231313-3331013211332013-3010212232121031-2122031012332021-1133132220023310-0122200023210001"></a>

Type: `"list"`. Computed.

List of subset class. Subsets class is defined using list of keys. Every unique combination of
values of these keys form a subset within the class.

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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-1322311321233131-3302233113233013-3002023330010210-0323113332031110-0203020131210122-2321012232131231-1010120320300310-2311103220122201"></a>

## Direct properties — endpoint_subsets / 221302030212 / 3

<a id="canonical-1122303030032221-2022330332030310-1102310103331303-0013013320203310-3313221030310031-2102121010212112-3231110200002130-2110012323210111"></a>

<a id="canonical-3232000021101322-3103100231200231-3312002030322230-2331031232030320-3212232222033200-0300000130132002-1322000013123001-1033301032203132"></a>

## keys property — endpoint_subsets / 221302030212 / 4

Type: `["list", "string"]`. Computed.

List of keys that define a cluster subset class.

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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-1303100101232123-2101123232011333-0313120212232120-1220230331021323-3232332300012011-3313113100122011-2033223132230202-2223010313002111"></a>

## Next pages — endpoint_subsets / 221302030212 / 5

- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1021103000030322-0131231321303110-3133120333303230-1013320313021013-1320103003030210-0130323110332113-0022001201131110-3003322320132213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222201131221200-1300301321133020-3302313131313130-0222220231202020-3111113011213001-0132030020220210-0310230202003032-2230013010230332"></a>

## default_pool.advanced_options.enable_subsets.fail_request — fail_request / 313122203323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003)
- default_pool.advanced_options.enable_subsets.fail_request

<a id="canonical-2313303003321112-0000210310032302-2211132133222033-1132100100302020-3131012101102313-1211313222310123-2210132123222300-0122122202210100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for fail request.

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

<a id="canonical-3003323222212120-3232310230301033-1031020333201110-2203132212012010-0000003322032231-2323000001110310-2202100201011233-0032223000330310"></a>

## Direct properties — fail_request / 313122203323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232300021121112-3001310020032110-1113310032020231-2201233002203032-2102031031031120-2100320202003033-2310011010333220-0200211203123230"></a>

## Next pages — fail_request / 313122203323 / 4

- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210022300212221-3211123332013211-1231023200002203-1212303112323113-3033333111231100-3021012002201332-1222200100230231-2012021103302003)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2012200011203112-1330213223231102-1210101031311213-0110120103020111-2112000010131232-3013022312210133-1322001232311220-1330102303301023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221131010001301-3013103201312302-0123211310323303-0003211211101003-1100311121130303-3102100231221130-2003303131331203-3333030123123330"></a>

## default_pool.advanced_options.http1_config — http1_config / 013012202012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.http1_config

<a id="canonical-3121213230322101-3320303123221221-3003231201320123-2113331231122033-3033203001323001-0211012003001210-0300001013330303-0302221100100122"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for upstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0111313222201201-3122311113013333-2030331303032321-2113330302213211-1320231320111312-3023211022013323-3030123233212011-3003222231001033"></a>

## Direct properties — http1_config / 013012202012 / 3

- [header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210232213221001-1003012011230031-0223332332021110-1200121313023211-3331133113002110-0210120312210132-1233232332323013-2312321030211220): complete subsection reference.

<a id="canonical-0021031103203321-3313233001022203-1112231230132233-0323103032020303-3302213110323013-2300103302331222-2210013032001121-1121120212102221"></a>

## Next pages — http1_config / 013012202012 / 4

- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210232213221001-1003012011230031-0223332332021110-1200121313023211-3331133113002110-0210120312210132-1233232332323013-2312321030211220)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2210232213221001-1003012011230031-0223332332021110-1200121313023211-3331133113002110-0210120312210132-1233232332323013-2312321030211220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021022012211113-3213033302120013-0332102033201211-3312303021003330-0201221112321202-2123200311303000-3110222030303113-0331323210102333"></a>

## default_pool.advanced_options.http1_config.header_transformation — header_transformation / 311212310113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.http1_config](data-sources--http_loadbalancer--reference--group-015.md#canonical-2012200011203112-1330213223231102-1210101031311213-0110120103020111-2112000010131232-3013022312210133-1322001232311220-1330102303301023)
- default_pool.advanced_options.http1_config.header_transformation

<a id="canonical-3212331102111021-1102132100213031-3331312211201303-1111201302333210-3302111331303103-2233033113113030-1131322132320022-2002110131313100"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

<a id="canonical-2112213131033133-1231000203230222-0103021130312133-0112300231212032-2321223221233101-0002321132202030-0213212313113312-2130101323001133"></a>

## Direct properties — header_transformation / 311212310113 / 3

- [default_header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-1031301133331213-2300012111223000-2130200222220110-0203311230222302-2132131323010111-1032200131122133-3332211210231010-1333323220323000): complete subsection reference.

- [preserve_case_header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-2322110023300022-1112030202333011-1220031323001120-0023100110112310-1331110010231311-2231310233230220-3310132123233210-0012132303311201): complete subsection reference.

- [proper_case_header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-1323033233133112-2221211333320013-3022331323213222-0002321132302020-3131022033110130-0232310130213020-0331010110220112-1221033122230311): complete subsection reference.

<a id="canonical-1210111021220033-3001202111331112-1020213111001123-0210222210213020-3321133122012232-3313301132131130-1100022220220212-3010110100202031"></a>

## Next pages — header_transformation / 311212310113 / 4

- [default_pool.advanced_options.http1_config.header_transformation.default_header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-1031301133331213-2300012111223000-2130200222220110-0203311230222302-2132131323010111-1032200131122133-3332211210231010-1333323220323000)
- [default_pool.advanced_options.http1_config.header_transformation.preserve_case_header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-2322110023300022-1112030202333011-1220031323001120-0023100110112310-1331110010231311-2231310233230220-3310132123233210-0012132303311201)
- [default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-1323033233133112-2221211333320013-3022331323213222-0002321132302020-3131022033110130-0232310130213020-0331010110220112-1221033122230311)
- [default_pool.advanced_options.http1_config](data-sources--http_loadbalancer--reference--group-015.md#canonical-2012200011203112-1330213223231102-1210101031311213-0110120103020111-2112000010131232-3013022312210133-1322001232311220-1330102303301023)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1031301133331213-2300012111223000-2130200222220110-0203311230222302-2132131323010111-1032200131122133-3332211210231010-1333323220323000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023231010332003-2300323122301200-0232323231122030-1331110333000101-3010021101220133-3313203123232132-1001322013303303-1222002011122013"></a>

## default_pool.advanced_options.http1_config.header_transformation.default_header_transformation — default_header_transformation / 201311113103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.http1_config](data-sources--http_loadbalancer--reference--group-015.md#canonical-2012200011203112-1330213223231102-1210101031311213-0110120103020111-2112000010131232-3013022312210133-1322001232311220-1330102303301023)
- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210232213221001-1003012011230031-0223332332021110-1200121313023211-3331133113002110-0210120312210132-1233232332323013-2312321030211220)
- default_pool.advanced_options.http1_config.header_transformation.default_header_transformation

<a id="canonical-1230103222331211-0033223020323100-0213012330200223-1202122223311022-3210133122213112-1131013013301320-2231023331010333-1111030103211030"></a>

Type: `["object", {}]`. Computed.

Use the platform's current default HTTP header transformation behavior.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1130021023332213-2300202032331210-0110200201331130-1111301322113211-3113131010210223-0203221012011222-0301311203100310-3103230332322121"></a>

## Direct properties — default_header_transformation / 201311113103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301233132102002-2121033120033023-0333222323230022-3022302221120333-3232233323232312-3210232003133030-3230131030033230-0231130232103201"></a>

## Next pages — default_header_transformation / 201311113103 / 4

- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210232213221001-1003012011230031-0223332332021110-1200121313023211-3331133113002110-0210120312210132-1233232332323013-2312321030211220)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2322110023300022-1112030202333011-1220031323001120-0023100110112310-1331110010231311-2231310233230220-3310132123233210-0012132303311201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223313032200333-0031230121331123-3300323000230200-3320003033131111-2103303101003023-1303300110112311-2123311210330203-2133111120323211"></a>

## default_pool.advanced_options.http1_config.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 020023012333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.http1_config](data-sources--http_loadbalancer--reference--group-015.md#canonical-2012200011203112-1330213223231102-1210101031311213-0110120103020111-2112000010131232-3013022312210133-1322001232311220-1330102303301023)
- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210232213221001-1003012011230031-0223332332021110-1200121313023211-3331133113002110-0210120312210132-1233232332323013-2312321030211220)
- default_pool.advanced_options.http1_config.header_transformation.preserve_case_header_transformation

<a id="canonical-2313122323023020-2221202012000010-0100231220120310-0312231122210002-0231013033332023-2231321210113010-1203202010223101-3021122302123110"></a>

Type: `["object", {}]`. Computed.

Preserve HTTP header-name case when upstream case must remain unchanged.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3233200210332022-3023110300131320-0102001201103113-1132332030122322-2311222013123200-1230333132133000-2212212131000113-3202230232011333"></a>

## Direct properties — preserve_case_header_transformation / 020023012333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210223330001301-1303001200021031-1110010211023133-3110300032311333-3013313122232230-0232321310221010-0320202131110303-1333031210231323"></a>

## Next pages — preserve_case_header_transformation / 020023012333 / 4

- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210232213221001-1003012011230031-0223332332021110-1200121313023211-3331133113002110-0210120312210132-1233232332323013-2312321030211220)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1323033233133112-2221211333320013-3022331323213222-0002321132302020-3131022033110130-0232310130213020-0331010110220112-1221033122230311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113121332220031-2002301100103123-1101302101313101-3333323011001312-2331210332121311-3022320322210222-0120100210232220-3003102312332011"></a>

## default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 013033003013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [default_pool.advanced_options.http1_config](data-sources--http_loadbalancer--reference--group-015.md#canonical-2012200011203112-1330213223231102-1210101031311213-0110120103020111-2112000010131232-3013022312210133-1322001232311220-1330102303301023)
- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210232213221001-1003012011230031-0223332332021110-1200121313023211-3331133113002110-0210120312210132-1233232332323013-2312321030211220)
- default_pool.advanced_options.http1_config.header_transformation.proper_case_header_transformation

<a id="canonical-0301010220313301-0000200200011023-0032313211231001-0322221220232302-3230101331201313-2033003103310002-3102123130322210-1001230203132313"></a>

Type: `["object", {}]`. Computed.

Transform HTTP header names to proper case when explicit transformation is required.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2233332313300130-2010212300213130-3323211231233001-2110210132121202-1103333123033001-0303102021002222-0012101013221230-2110233120002213"></a>

## Direct properties — proper_case_header_transformation / 013033003013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220213000211133-2130003030311323-3230130300100132-1310023113321213-3120023213133210-0130211112111022-0030132211310100-2301303021220111"></a>

## Next pages — proper_case_header_transformation / 013033003013 / 4

- [default_pool.advanced_options.http1_config.header_transformation](data-sources--http_loadbalancer--reference--group-015.md#canonical-2210232213221001-1003012011230031-0223332332021110-1200121313023211-3331133113002110-0210120312210132-1233232332323013-2312321030211220)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0302103210203123-2300200322030323-3303020103012020-2031120223010323-2313200232200231-3203333211320133-2033203220101011-1023222100033000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210020201233232-2001333123230232-3013230032233323-0312031321011211-1222020323232310-2121330302303202-1211320322121101-2033300011030131"></a>

## default_pool.advanced_options.http2_options — http2_options / 210222131201 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.http2_options

<a id="canonical-1031310232203301-1103211232010213-1030121201031122-1002133103113210-0132111323003320-2213332302122311-2123332201312010-3330323330230003"></a>

Type: `"single"`. Computed.

Http2 Protocol OPTIONS for upstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3331330302020310-2101103303101322-3232201132122212-2101013120223133-0300201313330303-1111202021000010-0013221133131032-2133330011132203"></a>

## Direct properties — http2_options / 210222131201 / 3

<a id="canonical-3033132012123301-2301031330001031-3033230032203230-2120310100222232-1013103330300012-3000111123110221-0221033111223131-2223022321020123"></a>

<a id="canonical-2102201211221230-1010131203113033-3213023312001130-3013223233201310-1013301220031200-2113020020020010-1102110333112112-1032131320332120"></a>

## enabled property — http2_options / 210222131201 / 4

Type: `"bool"`. Computed.

Enable/disable HTTP2 Protocol for upstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1031302322123123-2102332202230211-3032130213011212-3300320322133012-3001323221322330-3113230110121100-0323130011020321-3103102300112213"></a>

## Next pages — http2_options / 210222131201 / 5

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0212013320121110-3033210311131000-3201212002231122-1110302332211220-0112213030302123-2320021233211303-0110130210300021-1323000111111223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022130030302303-3021100031301111-3001313311332320-3203101003313202-2032133001012320-0121103000022310-1002011031033001-3133220033203203"></a>

## default_pool.advanced_options.no_panic_threshold — no_panic_threshold / 330132003213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.no_panic_threshold

<a id="canonical-2102001132321320-2131332001300320-2221023030201300-3301330333303210-3312220002301102-2122310021031132-3220113301121300-1303312120203001"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no panic threshold. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-0202001221201212-0120333230001232-2313003031122303-3112201111310321-0030110121332121-1333003020322023-3111120012332122-0210222003002212"></a>

## Direct properties — no_panic_threshold / 330132003213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011101312332233-3311302103220233-3011012010131110-2313010220220031-3132322102211232-2323130111232323-1110300311301030-2220320133012003"></a>

## Next pages — no_panic_threshold / 330132003213 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1321012101122211-2102000320220313-1331310322231032-1203123320313111-1311130321320303-1020331231312103-1330303100211131-0032021020020300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223013313132131-2111321221300123-3212122322131020-3023221313311111-2230201223003312-0301121021030110-0211231300012020-2323021103203033"></a>

## default_pool.advanced_options.no_request_limit_per_connection — no_request_limit_per_connection / 232133320002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.no_request_limit_per_connection

<a id="canonical-3213313013103333-1010210020121031-3333032203130212-0012021030212213-1333233031223100-2201011212000112-3212332102200331-3221210232302212"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no request limit per connection. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

<a id="canonical-1111121011201320-2120031211302101-3032320023022202-2323022033000130-3231310302231031-0011003322333230-1222111322112332-1013103013220031"></a>

## Direct properties — no_request_limit_per_connection / 232133320002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001120331130013-2130331033133333-0010311300200321-3003223233032302-2013230313010122-3010311011220013-0212330221000201-2311011130223333"></a>

## Next pages — no_request_limit_per_connection / 232133320002 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0233100000211130-1311321230320220-1220101003130100-1202010131212223-3301032010002020-3231201110022231-3002222332232320-2213210303031120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121120333102223-3030310013032113-3120000110201022-1123101003220230-1121102313001211-3101113310303001-1102220013123003-2333230013021312"></a>

## default_pool.advanced_options.outlier_detection — outlier_detection / 210332031220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.outlier_detection

<a id="canonical-2021121330101131-0023221102321201-1102111202100320-2213210223231213-3303201011110310-2021233211300103-1301213331003112-0322130123323202"></a>

Type: `"single"`. Computed.

Outlier detection and ejection is the process of dynamically determining whether some number of
hosts in an upstream cluster are performing unlike the others and removing them from the healthy
load balancing set. Outlier detection is a form of passive health checking. Algorithm 1.

Upstream description:

Outlier detection and ejection is the process of dynamically determining whether some number of
hosts in an upstream cluster are performing unlike the others and removing them from the healthy
load balancing set. Outlier detection is a form of passive health checking.

Algorithm

&#8203;1. A endpoint is determined to be an outlier (based on configured number of consecutive\_5xx
or consecutive\_gateway\_failures) . &#8203;2. If no endpoints have been ejected, loadbalancer will
eject the host immediately. Otherwise, it checks to make sure the number of ejected hosts is below
the allowed threshold (specified via max\_ejection\_percent setting). If the number of ejected hosts
is above the threshold, the host is not ejected. &#8203;3. The endpoint is ejected for some number
of milliseconds. Ejection means that the endpoint is marked unhealthy and will not be used during
load balancing. The number of milliseconds is equal to the base\_ejection\_time value multiplied by
the number of times the host has been ejected. &#8203;4. An ejected endpoint will automatically be
brought back into service after the ejection time has been satisfied.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3111321122132102-1033223000022110-2030011321300120-1131231133012332-3031331201123223-2201030103221122-1133311211223032-3100011330232030"></a>

## Direct properties — outlier_detection / 210332031220 / 3

<a id="canonical-0321123221032321-1121212112332200-1103312113133212-2211203220020331-2212110023120201-0123313302102011-0221130102311013-0230132333133310"></a>

<a id="canonical-1311233110222110-3110230120302013-3011030303000023-2202121010123020-2030221030300101-2203301020100311-0032211111202331-3223120130031333"></a>

## base_ejection_time property — outlier_detection / 210332031220 / 4

Type: `"number"`. Computed.

The base time that a host is ejected for. The real time is equal to the base time multiplied by the
number of times the host has been ejected. This causes hosts to GET ejected for longer periods if
they continue to fail.

Upstream description:

The base time that a host is ejected for. The real time is equal to the base time multiplied by the
number of times the host has been ejected. This causes hosts to GET ejected for longer periods if
they continue to fail. Defaults to 30000ms or 30s. Specified in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1800000,
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
    "ves.io.schema.rules.uint32.lte": "1800000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  }
}
```

<a id="canonical-2102233030131311-1311020320200111-1020313003302232-1210301320013112-0122122021120033-0122012320302332-1121213221123303-0031123123101213"></a>

<a id="canonical-1121201023111100-0302300023103110-2031010022333312-2320310130132013-0202321202020213-2120131112302003-3320321023101201-2210312131233213"></a>

## consecutive_5xx property — outlier_detection / 210332031220 / 5

Type: `"number"`. Computed.

If an upstream endpoint returns some number of consecutive 5xx, it will be ejected. Note that in
this case a 5xx means an actual 5xx respond code, or an event that would cause the HTTP router to
return one on the upstream’s behalf(reset, connection failure, etc.) consecutive\_5xx indicates
the..

Upstream description:

If an upstream endpoint returns some number of consecutive 5xx, it will be ejected. Note that in
this case a 5xx means an actual 5xx respond code, or an event that would cause the HTTP router to
return one on the upstream’s behalf(reset, connection failure, etc.) consecutive\_5xx indicates the
number of consecutive 5xx responses required before a consecutive 5xx ejection occurs. Defaults to
5.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

<a id="canonical-2322210233320332-3230213220110320-1120033111132103-1111123312002330-0001203033123231-0300230312200212-2022023310102311-2210011320031202"></a>

<a id="canonical-0021320303113310-2331001212210220-2231101201302121-1220012213332002-3230132132303002-2120020231113302-3101113222230021-2210030030023323"></a>

## consecutive_gateway_failure property — outlier_detection / 210332031220 / 6

Type: `"number"`. Computed.

If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status
code), it will be ejected. Note that this includes events that would cause the HTTP router to return
one of these status codes on the upstream’s behalf (reset, connection failure, etc.)..

Upstream description:

If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status
code), it will be ejected. Note that this includes events that would cause the HTTP router to return
one of these status codes on the upstream’s behalf (reset, connection failure, etc.).
Consecutive\_gateway\_failure indicates the number of consecutive gateway failures before a
consecutive gateway failure ejection occurs. Defaults to 5.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

<a id="canonical-0022200230312302-3323200320110333-0330031113103121-2212231021111323-1333100010320202-1331300332012020-1322232012132311-1310101233220102"></a>

<a id="canonical-1331030233101213-0200320230113011-0211203303311222-3313300002233033-0333031021021030-2031120023101113-3330211111301132-1103331222230321"></a>

## interval property — outlier_detection / 210332031220 / 7

Type: `"number"`. Computed.

The time interval between ejection analysis sweeps. This can result in both new ejections as well as
endpoints being returned to service. Defaults to \`10000ms\`.

Upstream description:

The time interval between ejection analysis sweeps. This can result in both new ejections as well as
endpoints being returned to service. Defaults to 10000ms or 10s. Specified in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-2111013133301020-1031100000120202-2210030012232122-1002023222233111-0332013021131330-1312310133311213-2300321113210131-3223333321020120"></a>

<a id="canonical-0111233212000023-2231012200211030-3130102021330110-0113312003301023-3233103330312113-0010000110332233-3010323032123103-2021211122203303"></a>

## max_ejection_percent property — outlier_detection / 210332031220 / 8

Type: `"number"`. Computed.

The maximum % of an upstream cluster that can be ejected due to outlier detection. but will eject at
least one host regardless of the value. Defaults to \`10%\`.

Upstream description:

The maximum % of an upstream cluster that can be ejected due to outlier detection. Defaults to 10%
but will eject at least one host regardless of the value.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-3011203311202320-1211101331212000-1002103203102231-2131221102321013-0223023331030013-3222200333210311-2030303212310313-2002110203232321"></a>

## Next pages — outlier_detection / 210332031220 / 9

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1320312022011223-0112101033132110-2322303210023201-1320121013000330-1331321233310011-3220111102322200-3021312012310331-0301233032003001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121330322311210-2331201003132233-3230102023132100-1111222332012120-1012102033223123-2021201222212332-2022110020332323-1111211111201322"></a>

## default_pool.advanced_options.proxy_protocol_v1 — proxy_protocol_v1 / 210330202202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.proxy_protocol_v1

<a id="canonical-1022300310301110-1013222230001303-1321223121122031-3302210203201232-3202011021200000-3002030130101030-2301333203102020-0221103110320020"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for proxy protocol v1.

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

<a id="canonical-0303201033122120-0223013011222232-3123003221133303-3020102303220100-3310021223020331-2000012212231033-3203212003013213-2020031013301232"></a>

## Direct properties — proxy_protocol_v1 / 210330202202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211221001332221-1001222022103131-3203202220032113-2330133000000203-3320202230112233-3131032131123231-1332301133121011-0201120000131022"></a>

## Next pages — proxy_protocol_v1 / 210330202202 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0211103313321032-0201122021031121-2311101000312300-2230110230320310-3213220020320023-2010000121013233-2311011010210301-3323120323303331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100221321033231-0113031000320312-2020223133001110-1222011012301011-2113033201020112-0010332023232132-2321232010113031-1230232113131002"></a>

## default_pool.advanced_options.proxy_protocol_v2 — proxy_protocol_v2 / 110133310332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- default_pool.advanced_options.proxy_protocol_v2

<a id="canonical-2221233201202201-3321310323030200-2331030221001122-1011021121021120-2123313101110221-1113020303123332-1032202100322312-3322032302111101"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for proxy protocol v2.

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

<a id="canonical-0303212010313101-0121012103313011-3012301022122310-1010323210032012-3100212112111200-0131313102110033-1021230321023101-2010212023231330"></a>

## Direct properties — proxy_protocol_v2 / 110133310332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002233133302330-2322320330113012-0321100100010022-0211130120112103-3231102022232123-3012312123133012-3011310220210203-3001002102132120"></a>

## Next pages — proxy_protocol_v2 / 110133310332 / 4

- [default_pool.advanced_options](data-sources--http_loadbalancer--reference--group-015.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3333201113312220-0211103023020231-1231022023021132-3130330021322021-1011333013330100-1221301202121312-1123312313021112-0032303012213003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110030020233212-1121001230003002-2113112012033112-1120213301011313-0232312011100000-1223012122211322-1323100132232131-3220100121113010"></a>

## default_pool.automatic_port — automatic_port / 203131011001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- default_pool.automatic_port

<a id="canonical-2000202230023223-2221130111130032-0330112333113212-2012022301223222-3023331313033032-1321120333233210-1033330330300323-3001110330311302"></a>

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

<a id="canonical-2121010033031330-2110022013331301-2113033030033131-2013333300220130-2302030110211103-1312323212212110-0120201230013311-2032032013033332"></a>

## Direct properties — automatic_port / 203131011001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010303331003022-3330010220323210-1201133002223300-2031020232103032-3312202031212302-1003323101003221-2120230100101120-3033031203112210"></a>

## Next pages — automatic_port / 203131011001 / 4

- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1202331312130223-3323033003031000-1300032201301212-2133102012302220-1012021120110103-2231131032011311-0213332212020202-2021031122333020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133020203000330-2300333101123121-1121313300030233-3331222232031302-3222033131311032-1331033131032210-3033121211110222-0003232320010131"></a>

## default_pool.healthcheck — healthcheck / 123321321323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- default_pool.healthcheck

<a id="canonical-0222331313333111-1100300002202300-1310203123203011-3211013210000120-0301320100310221-3200112003012230-0132301121212120-1122003300103131"></a>

Type: `"list"`. Computed.

Reference to healthcheck configuration objects. Defaults to \`\[\]\`. Server applies default when
omitted.

Upstream description:

Reference to healthcheck configuration objects.

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

<a id="canonical-3323132303323130-0013000111030120-1112300102101131-3213300221223221-2300300323312030-3200001120020121-1103321130312330-2100333312032000"></a>

## Direct properties — healthcheck / 123321321323 / 3

<a id="canonical-2133012220020301-3333113212231020-2231102300202120-1120323023331122-2200232123101133-1221033021321111-2331032112120321-2130123113112003"></a>

<a id="canonical-0323210112320201-0202001210332032-1101120112003211-1233121121011312-1132133302301320-3030210100223210-1321021310321103-0113132002320310"></a>

## name property — healthcheck / 123321321323 / 4

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

<a id="canonical-0311101012013130-3303321230202133-2113233312230310-3120002323122213-2330220110321230-3222201010103230-2221201110131202-0131111330301001"></a>

<a id="canonical-2310023212332200-0310120220320212-1230321010123212-3022111123213210-1020323223231020-2123302111132302-0213122203233221-3021122221202301"></a>

## namespace property — healthcheck / 123321321323 / 5

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

<a id="canonical-3110113022031032-0321313003300313-2333103213331203-3021200222302030-1220120001322231-2130131312032230-0133213221222000-1211112331230310"></a>

<a id="canonical-0133123212212021-3230112312303131-1331233010232033-2303201203112302-1021122233123303-2113030111102322-2333333113230123-1200030100112231"></a>

## tenant property — healthcheck / 123321321323 / 6

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

<a id="canonical-3132012312330122-3021302310120120-3202210122021103-1023321113221123-1233230101022112-3331013033302322-3102033022111122-2203120323113010"></a>

## Next pages — healthcheck / 123321321323 / 7

- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2122312022231123-2122221001303110-1111302333031202-2020023322122331-0020010213103110-0012000101322211-1200121210030020-1320232013101313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022220323333022-0331320311120000-0022332220223301-2213332313022201-3012103310033200-1023122331320010-1201130232233030-3122231232111221"></a>

## default_pool.lb_port — lb_port / 101202101331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- default_pool.lb_port

<a id="canonical-1323032232231033-1120102212003320-3222332333022333-2203011002113032-0231033220031303-2322212202133333-1221313001031302-1130210123313201"></a>

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

<a id="canonical-0130310301030023-1211330321313202-0022111121121333-0123221010313132-2312023010210331-0100321102100020-2310002030011030-2321020221220020"></a>

## Direct properties — lb_port / 101202101331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321200201013310-3332023312102321-3113302120113113-2222320223233130-1320302032003121-3003131002130122-0321121103201131-0020202230210022"></a>

## Next pages — lb_port / 101202101331 / 4

- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1102212210131000-1000101230031101-0023123120321223-0123033000022131-0222321112012111-0001230102130023-2110102101300223-1203122202213132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023102011033203-0201210231111132-0332101320132131-0111321130331313-0310122312211113-1000201233032322-0031001211231130-1032003312302200"></a>

## default_pool.no_tls — no_tls / 222200222102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- default_pool.no_tls

<a id="canonical-2010010321203031-1132003130232303-3311120001013322-3032312231122132-3213321031332211-2123300121002303-1112132113123100-3221131331300303"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-3133311221121023-3022131022033010-1033231233120213-3312101232231232-3123313011022030-1213022302222331-2231201132200002-1202323332332012"></a>

## Direct properties — no_tls / 222200222102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202223103331313-0023331321323032-1321103102233002-1032103303100202-1001130330310113-3220113331302121-3202031220100203-3302212000232132"></a>

## Next pages — no_tls / 222200222102 / 4

- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232230331222022-3231010233203031-1020123032202131-1012311221103101-3211002020231203-0203210030230320-0102202011313213-0323310200223001"></a>

## default_pool.origin_servers — origin_servers / 012200032220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- default_pool.origin_servers

<a id="canonical-3200221210222012-1011330212322330-1011303021020311-1130231023312210-3032131323013121-1020022313021301-2000100030321332-1203110230010022"></a>

Type: `"list"`. Computed.

Origin Servers. List of origin servers in this pool.

Upstream description:

List of origin servers in this pool.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3020020233022102-0230023310031021-0201232103230111-2303101131010222-0302222322132132-0212300013023121-0023102123333321-0323033303331131"></a>

## Direct properties — origin_servers / 012200032220 / 3

- [cbip_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-1100110213210332-0331202020032201-1112330111231031-3202003310322001-0123113030213003-1233321303111302-0212313031020331-3112211233012012): complete subsection reference.

- [consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012): complete subsection reference.

- [custom_endpoint_object](data-sources--http_loadbalancer--reference--group-016.md#canonical-0011112221011021-1222300212032111-2121332331230003-0310301311130002-2230221332312323-2332003230101112-2222223001102001-0201111233222103): complete subsection reference.

- [k8s_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-2200121001022120-1312020022302121-2020021130031332-3321020003031001-0313000101213021-1313031221121233-0203213331003200-3032133202102321): complete subsection reference.

<a id="canonical-3220012212232102-0122011121020002-2221120121012223-2301313121101320-3221121310023233-1131223123132232-2232311212030300-2000322013203300"></a>

<a id="canonical-2303231213113200-3010320302032013-2201302132031233-0311332301200221-0113003030022010-1220302131103112-3213110010021011-1213131123332230"></a>

## labels property — origin_servers / 012200032220 / 4

Type: `["map", "string"]`. Computed.

Add Labels for this origin server, these labels can be used to form subset.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-0011033200022011-2330321200020232-3033333103222033-2212313221112010-3303213020232130-1303311022310010-0112020221103100-3223101011022033): complete subsection reference.

- [private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-3031132322222002-1300113131221032-0020330231123313-2012010120103023-2212101132230332-1211331333233331-2131020011332033-1120131100222222): complete subsection reference.

- [public_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-3230232123101031-2001223303123020-1213102222300331-0202333300210101-2132110133113330-1212330031022321-1002123133032130-2211123112002303): complete subsection reference.

- [public_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-3213313232020122-3010103102320133-2122200332231332-1300302012030303-1213021123010331-2023230013331002-2031230031231102-3020232021202232): complete subsection reference.

- [vn_private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-0211323132213300-2233131133130022-1221132031330120-3031032002223301-0212313121002022-3330103333233300-1120331120002031-0013013203231112): complete subsection reference.

- [vn_private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-0010223201122332-2321021313212022-0223002110212131-1011310123201121-3013320132013212-3232112322022332-1201100103213213-2331200010213312): complete subsection reference.

<a id="canonical-2012321333323023-0133120230331011-1012103111323032-0023111132130221-2330103000211201-1123323122313002-1321130321203323-3113202330033211"></a>

## Next pages — origin_servers / 012200032220 / 5

- [default_pool.origin_servers.cbip_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-1100110213210332-0331202020032201-1112330111231031-3202003310322001-0123113030213003-1233321303111302-0212313031020331-3112211233012012)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- [default_pool.origin_servers.custom_endpoint_object](data-sources--http_loadbalancer--reference--group-016.md#canonical-0011112221011021-1222300212032111-2121332331230003-0310301311130002-2230221332312323-2332003230101112-2222223001102001-0201111233222103)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--reference--group-016.md#canonical-2200121001022120-1312020022302121-2020021130031332-3321020003031001-0313000101213021-1313031221121233-0203213331003200-3032133202102321)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-0011033200022011-2330321200020232-3033333103222033-2212313221112010-3303213020232130-1303311022310010-0112020221103100-3223101011022033)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-3031132322222002-1300113131221032-0020330231123313-2012010120103023-2212101132230332-1211331333233331-2131020011332033-1120131100222222)
- [default_pool.origin_servers.public_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-3230232123101031-2001223303123020-1213102222300331-0202333300210101-2132110133113330-1212330031022321-1002123133032130-2211123112002303)
- [default_pool.origin_servers.public_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-3213313232020122-3010103102320133-2122200332231332-1300302012030303-1213021123010331-2023230013331002-2031230031231102-3020232021202232)
- [default_pool.origin_servers.vn_private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-0211323132213300-2233131133130022-1221132031330120-3031032002223301-0212313121002022-3330103333233300-1120331120002031-0013013203231112)
- [default_pool.origin_servers.vn_private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-0010223201122332-2321021313212022-0223002110212131-1011310123201121-3013320132013212-3232112322022332-1201100103213213-2331200010213312)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1100110213210332-0331202020032201-1112330111231031-3202003310322001-0123113030213003-1233321303111302-0212313031020331-3112211233012012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230332012223123-2322222211131102-1111332020233222-2112003301322032-1131303121110013-1312133122010331-0332220322332223-2131110231131211"></a>

## default_pool.origin_servers.cbip_service — cbip_service / 220232302310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- default_pool.origin_servers.cbip_service

<a id="canonical-1023032123012012-1223301032311301-3103011301200020-0213200010003201-3201203123331103-2301311121122002-0010232201220022-3221110221003032"></a>

Type: `"single"`. Computed.

Specify origin server with Classic BIG-IP Service (Virtual Server).

Upstream description:

Specify origin server with Classic BIG-IP Service (Virtual Server)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1110022232003222-1121203300011223-2320131100321103-2212110113003002-1300101010121321-2023311023302021-2001233122322022-3301302313001223"></a>

## Direct properties — cbip_service / 220232302310 / 3

<a id="canonical-2301032302323030-1023111100022130-1331003333031131-1302012101003321-0231023333311130-2131032301013220-2300033130000222-1100302201303101"></a>

<a id="canonical-3313323311032212-0211121230103310-3003112120232132-1231012131103021-0311123120031231-1010013303002300-0102023022010203-2220213030333202"></a>

## service_name property — cbip_service / 220232302310 / 4

Type: `"string"`. Computed.

Name of the discovered Classic BIG-IP virtual server to be used as origin.

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

<a id="canonical-3320122103000001-0013010200121311-0302121032303130-1302012123331302-2113100230333010-1131201111303301-1102012122303110-1003330002312112"></a>

## Next pages — cbip_service / 220232302310 / 5

- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303110231203201-2130130311003332-2130033132220311-3311321111023302-1321201323131311-0132322221120121-2010021322232322-0310120120033120"></a>

## default_pool.origin_servers.consul_service — consul_service / 121203321022 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- default_pool.origin_servers.consul_service

<a id="canonical-1011102303010031-3200003132020001-2032010332311102-1001022101210323-2130200220122201-1101130310103103-1122332232300133-1313311102112321"></a>

Type: `"single"`. Computed.

Specify origin server with HashiCorp Consul service name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\"]"
}
```

<a id="canonical-0000311321120203-0001323302132130-2023332120010131-0101130022001103-2223020010102101-3230002130302300-2330032301131301-0312003132332100"></a>

## Direct properties — consul_service / 121203321022 / 3

- [inside_network](data-sources--http_loadbalancer--reference--group-015.md#canonical-2303011000103121-1202321131031130-1312032030123001-0100322010011312-1103002003010023-2203233111322212-1332012232203120-1123020331123103): complete subsection reference.

- [outside_network](data-sources--http_loadbalancer--reference--group-015.md#canonical-2123003131002121-2101100002331012-0231302322200311-0131202231323330-3210011131000003-0220031210113203-1003133130103120-2331231203300310): complete subsection reference.

<a id="canonical-1231130201330020-3320201331200011-0203322212230100-3202023201202310-2220100300110000-3300201213020023-2323311020322303-0003303222330003"></a>

<a id="canonical-0223233302230012-2303130321010220-3112010031202120-1100110023122303-1130323033102313-1230323301021020-3232301301212203-0323123223022020"></a>

## service_name property — consul_service / 121203321022 / 4

Type: `"string"`. Computed.

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

Upstream description:

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

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

- [site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-2211302003222203-2132112033110303-1130013133302330-2212121220213032-2023201120022331-2300321022331111-3203122221333323-0021100301112301): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1020110202101103-1103030232200131-1032032323332003-1131212310011122-2120300021033002-3103222220102110-2211231210132213-2331112221023131): complete subsection reference.

<a id="canonical-0102110213210100-3330103212002303-2112203303101120-2023302212022321-0133300100310012-0333323311222132-1333213020031303-1033023301103130"></a>

## Next pages — consul_service / 121203321022 / 5

- [default_pool.origin_servers.consul_service.inside_network](data-sources--http_loadbalancer--reference--group-015.md#canonical-2303011000103121-1202321131031130-1312032030123001-0100322010011312-1103002003010023-2203233111322212-1332012232203120-1123020331123103)
- [default_pool.origin_servers.consul_service.outside_network](data-sources--http_loadbalancer--reference--group-015.md#canonical-2123003131002121-2101100002331012-0231302322200311-0131202231323330-3210011131000003-0220031210113203-1003133130103120-2331231203300310)
- [default_pool.origin_servers.consul_service.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-2211302003222203-2132112033110303-1130013133302330-2212121220213032-2023201120022331-2300321022331111-3203122221333323-0021100301112301)
- [default_pool.origin_servers.consul_service.snat_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1020110202101103-1103030232200131-1032032323332003-1131212310011122-2120300021033002-3103222220102110-2211231210132213-2331112221023131)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2303011000103121-1202321131031130-1312032030123001-0100322010011312-1103002003010023-2203233111322212-1332012232203120-1123020331123103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331131211302113-3322031321311013-0302211310333132-3200032232231311-2333303302232331-2012130113202110-3111112132110001-0111331310212010"></a>

## default_pool.origin_servers.consul_service.inside_network — inside_network / 212122110223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- default_pool.origin_servers.consul_service.inside_network

<a id="canonical-2303213312213120-1202023200323020-2213112323201103-2132000031312322-1100000201232230-2013103313011112-0001020321033002-1021330323020201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

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

<a id="canonical-3132120301231313-2103332310003100-1013211023123033-0020121003223021-3311003323022020-2113132002312232-2320322323131100-3102311000220102"></a>

## Direct properties — inside_network / 212122110223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011003010211220-2011122220223321-1031233330212322-2133012123301301-0013000132310223-3311201031201300-1232220113301022-1301021212132202"></a>

## Next pages — inside_network / 212122110223 / 4

- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2123003131002121-2101100002331012-0231302322200311-0131202231323330-3210011131000003-0220031210113203-1003133130103120-2331231203300310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000112223030022-3301320122011211-3000310123232001-2000020113103233-0203302011222232-3020220120311010-1103321130112031-1112300313310211"></a>

## default_pool.origin_servers.consul_service.outside_network — outside_network / 220010022122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- default_pool.origin_servers.consul_service.outside_network

<a id="canonical-1231310033033000-0221220230321230-0313111002213120-2211111010110310-2333122113313001-3223310330032230-1200233220320220-3333313312031200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

<a id="canonical-1223002301002022-3102021013001110-2310323303122132-3210032203331230-1111023020020011-1132333231220230-3020022203031202-2031302111310202"></a>

## Direct properties — outside_network / 220010022122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130130122030210-3011133122002332-2133200001111033-3113021303123211-2032133323200330-2312031133033323-2212111211021121-3301111010302031"></a>

## Next pages — outside_network / 220010022122 / 4

- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2211302003222203-2132112033110303-1130013133302330-2212121220213032-2023201120022331-2300321022331111-3203122221333323-0021100301112301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232201331222211-1223013301121330-2332030003023300-1300003032111312-0221231003312100-3230230012231202-0023130101022231-1321031203121110"></a>

## default_pool.origin_servers.consul_service.site_locator — site_locator / 030101110102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- default_pool.origin_servers.consul_service.site_locator

<a id="canonical-3311002110203303-3221311131012203-2000030203233202-2102320103101003-3021323131012102-0233032331113323-2132000230123112-3320023231233001"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

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

<a id="canonical-0310003322301222-0020132311320012-3330303002103313-3000011133320101-3222222020232133-1120032301131210-2301310320102110-0201202231013131"></a>

## Direct properties — site_locator / 030101110102 / 3

- [site](data-sources--http_loadbalancer--reference--group-015.md#canonical-0113133213010031-0221212130213112-1101012102300010-3020102310301112-1102021210300300-2220312222302102-1123231320022231-2210313003012232): complete subsection reference.

- [virtual_site](data-sources--http_loadbalancer--reference--group-015.md#canonical-0211211332011303-3233123201231311-3302120303322113-3031011333200011-1233213031111300-3020313112312331-2023000033332212-0132111131313211): complete subsection reference.

<a id="canonical-1101032303033110-1302301210321210-1033032001332330-1232202001212230-2113103202322331-0121111213300111-3221220300200100-2021303030020311"></a>

## Next pages — site_locator / 030101110102 / 4

- [default_pool.origin_servers.consul_service.site_locator.site](data-sources--http_loadbalancer--reference--group-015.md#canonical-0113133213010031-0221212130213112-1101012102300010-3020102310301112-1102021210300300-2220312222302102-1123231320022231-2210313003012232)
- [default_pool.origin_servers.consul_service.site_locator.virtual_site](data-sources--http_loadbalancer--reference--group-015.md#canonical-0211211332011303-3233123201231311-3302120303322113-3031011333200011-1233213031111300-3020313112312331-2023000033332212-0132111131313211)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0113133213010031-0221212130213112-1101012102300010-3020102310301112-1102021210300300-2220312222302102-1123231320022231-2210313003012232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131133330211210-1033303011110322-1023003322132111-1101323323103200-2232111331202111-1301233211210311-3311203002322203-0032001111130320"></a>

## default_pool.origin_servers.consul_service.site_locator.site — site / 300032102302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- [default_pool.origin_servers.consul_service.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-2211302003222203-2132112033110303-1130013133302330-2212121220213032-2023201120022331-2300321022331111-3203122221333323-0021100301112301)
- default_pool.origin_servers.consul_service.site_locator.site

<a id="canonical-3030001010123121-2100323101030311-2121121322133021-2122332330103001-2320323233123012-1300011230322010-3200222031200101-2003121101322301"></a>

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

<a id="canonical-3233031301003222-1023033322110000-1100030020113210-1233323123321210-2023020122313020-1032120222001210-1010010323000131-3231012123102010"></a>

## Direct properties — site / 300032102302 / 3

<a id="canonical-0303031302132022-2122200122203010-0031102112212303-0103322321033102-0033212202221021-3012222202111212-3112212002303220-0330121010322300"></a>

<a id="canonical-3332131123303210-3311122321023031-0203003120031220-3333130220322131-1023130203200333-2310201101223233-3330213123012311-3210201110033230"></a>

## name property — site / 300032102302 / 4

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

<a id="canonical-3223113320211321-3102313112122100-3002221312133032-1212233311301203-0132201133303130-1202101320313033-2333331012311011-2200222301232322"></a>

<a id="canonical-2033113002223312-2212032223030223-1330020131312003-3311230002320102-1011233322321120-1000111100120031-3023322222133123-0102332023323301"></a>

## namespace property — site / 300032102302 / 5

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

<a id="canonical-0310303223130012-0012022310012113-1312212131130211-0301222133033012-2203020000102131-3210023012332011-1010321300233332-1010233103022130"></a>

<a id="canonical-2210010303113003-2110110201221211-1320302221122313-0000023033311230-3321213030112330-3020302000130101-1113213310102030-1123032311102210"></a>

## tenant property — site / 300032102302 / 6

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

<a id="canonical-3330132122012022-0230020202133301-2323102011103212-3003223133010123-3113210302300030-2212030111323113-0232131131210301-0303012300111313"></a>

## Next pages — site / 300032102302 / 7

- [default_pool.origin_servers.consul_service.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-2211302003222203-2132112033110303-1130013133302330-2212121220213032-2023201120022331-2300321022331111-3203122221333323-0021100301112301)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0211211332011303-3233123201231311-3302120303322113-3031011333200011-1233213031111300-3020313112312331-2023000033332212-0132111131313211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001122310011330-3113110021022220-3232311211213322-2300203201032313-0320221002230332-3313103023321020-1211311020110030-2100002200220232"></a>

## default_pool.origin_servers.consul_service.site_locator.virtual_site — virtual_site / 212333312122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000)
- [default_pool.origin_servers.consul_service](data-sources--http_loadbalancer--reference--group-015.md#canonical-1112103230131002-1200210223233200-0200101321102012-1223021132110233-3111100121021212-2220131300023032-1220002021113223-0223302333001012)
- [default_pool.origin_servers.consul_service.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-2211302003222203-2132112033110303-1130013133302330-2212121220213032-2023201120022331-2300321022331111-3203122221333323-0021100301112301)
- default_pool.origin_servers.consul_service.site_locator.virtual_site

<a id="canonical-1022002303110200-0201012003302300-1200331232123321-0033313122102300-1012310211302203-0131232030111320-0210320221012123-2031301011330033"></a>

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

<a id="canonical-1130311202103121-0303201213311033-2023210112302022-3333311033323120-0011113033113201-2113112122120222-0231333010122103-0112112102321312"></a>

## Direct properties — virtual_site / 212333312122 / 3

<a id="canonical-2013331321002110-3320233022020002-2123111221023213-3333031132012213-0223120211301021-3032102132013212-2203231210030300-3132111232122001"></a>

<a id="canonical-0300011130312223-0232230311302011-0302123332112102-3030202021113220-2300302320332030-1313000210132031-1331323230303003-0011333331212023"></a>

## name property — virtual_site / 212333312122 / 4

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

<a id="canonical-1131322131332303-0230032100320113-2201310302320010-3001211200311200-1312201323332200-1331032200113201-0013223321133232-0321012133201331"></a>

<a id="canonical-0233301102301110-1312023201330313-1312000322220002-1012333003133210-0213020132310010-0222011021222020-2231110222002013-2122020311213231"></a>

## namespace property — virtual_site / 212333312122 / 5

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

<a id="canonical-3031112031222100-3121303012212312-1202201103122322-0001103202021131-2333030122100123-2302321300020201-2212312122310000-3201302021131000"></a>

<a id="canonical-0210332111230331-1011001133310323-3333012020032233-3010313210023032-2221111212312212-2220013202331122-1211002032330102-2132212323333020"></a>

## tenant property — virtual_site / 212333312122 / 6

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

<a id="canonical-1132010031031030-1201221010033032-3132100322010200-3113301333021302-3233021022200211-1331320223310232-2020220021112323-3010301111130212"></a>

## Next pages — virtual_site / 212333312122 / 7

- [default_pool.origin_servers.consul_service.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-2211302003222203-2132112033110303-1130013133302330-2212121220213032-2023201120022331-2300321022331111-3203122221333323-0021100301112301)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1020110202101103-1103030232200131-1032032323332003-1131212310011122-2120300021033002-3103222220102110-2211231210132213-2331112221023131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
