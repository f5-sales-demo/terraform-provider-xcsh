---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3103111000311001-2123100122110201-1031202311021322-1220202000211210-3310020232300023-2310003222321322-0113032232110210-3003312100312112"></a>

## add_hsts property — https_auto_cert / 023302212101 / 4

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0131231133130102-0222323300032000-1022233013023223-2112123000130330-0111330211121013-2233130012220032-3120131332321322-0121120023131323"></a>

<a id="canonical-2101113203013002-3031333201203202-2323233303311321-3311323133131320-2033332003212001-1333230300232133-1302232211022103-2301121302030023"></a>

## append_server_name property — https_auto_cert / 023302212101 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](data-sources--workload--reference--group-007.md#canonical-2003301130333322-2101101031011320-1100101132221102-1230230213222110-3112133132333232-3111013121312102-1310020020113021-3233311312133103): complete subsection reference.

<a id="canonical-2011231330120010-3321200003121110-3121110230233232-0133103023010112-0302112132023120-2212302000013012-3222331113333230-1113101222100103"></a>

<a id="canonical-0131312313230013-1023200233023100-1303233211201222-2120000313322110-2303020220203032-0233022031300012-0100011110020111-2322011311322110"></a>

## connection_idle_timeout property — https_auto_cert / 023302212101 / 6

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](data-sources--workload--reference--group-007.md#canonical-0032013200101331-3121101033132032-0201111023102312-3310122312023202-1011012000331210-3022033220220302-0200212023113213-0311301132011033): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-007.md#canonical-2303033023003313-0320312301010320-1232130123120130-1031233120230310-3121203332301102-3100303300110230-2031223030231013-2131312113012223): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-007.md#canonical-2202303122103033-0101112331121031-2302302322233103-1211312111211132-3001231303232133-2132213010032312-3012203131322300-1130300102322033): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-007.md#canonical-2100001101020213-1220100210103220-1202213122130010-3121202002331001-1122131123302303-1211021303321321-1020132302133130-1012012113002110): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133): complete subsection reference.

<a id="canonical-0132022311223331-1310331213233231-1220213000013313-1233231110123300-3302023002211130-3221122120331012-0110220330122113-0023110331213323"></a>

<a id="canonical-1320321221210011-2112100313022100-3020301133310011-1100120132322202-1013233230322201-2211331010011122-1302101033323331-2321033213233221"></a>

## http_redirect property — https_auto_cert / 023302212101 / 7

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_mtls](data-sources--workload--reference--group-007.md#canonical-0231223001232221-2201332131123020-1313300302020312-0221201321000320-0033233210023223-2110213013103111-0102220330311200-2331112323002321): complete subsection reference.

- [non_default_loadbalancer](data-sources--workload--reference--group-007.md#canonical-3031130303320312-3011210220300230-1112101210130120-3233211133223031-3232211103310222-3322323003212131-0300113203000232-1001321332302222): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-007.md#canonical-2213213201020321-0303321203211020-2202012103010022-2302121320212312-1023313101332232-3230002111032132-2311201002200103-3213110102330302): complete subsection reference.

<a id="canonical-3003311001113331-2012032101133002-0102020123011321-2022101023023322-1023221232111333-2210202033020200-1222333022103211-2232130100012113"></a>

<a id="canonical-3300121032313320-2332221120232000-2132320132010200-1320202202110131-2021332211202122-0302203331333103-0132000220211333-1103203121213220"></a>

## port property — https_auto_cert / 023302212101 / 8

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3331110120222322-0210303322322302-1331132213031331-1211213210103330-0113000021331103-1102221021100110-2231120333320010-1222222330021303"></a>

<a id="canonical-2221332012301013-2332203000010021-1211232330111113-1001111301202031-2302021130030321-2212322203103120-0132322302310133-1032102331323002"></a>

## port_ranges property — https_auto_cert / 023302212101 / 9

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-1022320132333323-0312210223033001-1132123011011111-0021111200112312-3333033311201122-1101302330312033-2101221011212010-2133103302003133"></a>

<a id="canonical-1222211223031001-2231303021133302-0320232133211231-1000331232311213-2210322221032221-0331330120201113-0011021010022011-3103230101110330"></a>

## server_name property — https_auto_cert / 023302212101 / 10

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_config](data-sources--workload--reference--group-007.md#canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221): complete subsection reference.

<a id="canonical-2021332100130001-2103002223201301-1321131312121022-0210000233122332-3220300030001331-1231322321122330-1212232030033100-3202203320303330"></a>

## Next pages — https_auto_cert / 023302212101 / 11

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-007.md#canonical-2003301130333322-2101101031011320-1100101132221102-1230230213222110-3112133132333232-3111013121312102-1310020020113021-3233311312133103)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header](data-sources--workload--reference--group-007.md#canonical-0032013200101331-3121101033132032-0201111023102312-3310122312023202-1011012000331210-3022033220220302-0200212023113213-0311301132011033)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer](data-sources--workload--reference--group-007.md#canonical-2303033023003313-0320312301010320-1232130123120130-1031233120230310-3121203332301102-3100303300110230-2031223030231013-2131312113012223)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize](data-sources--workload--reference--group-007.md#canonical-2202303122103033-0101112331121031-2302302322233103-1211312111211132-3001231303232133-2132213010032312-3012203131322300-1130300102322033)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize](data-sources--workload--reference--group-007.md#canonical-2100001101020213-1220100210103220-1202213122130010-3121202002331001-1122131123302303-1211021303321321-1020132302133130-1012012113002110)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls](data-sources--workload--reference--group-007.md#canonical-0231223001232221-2201332131123020-1313300302020312-0221201321000320-0033233210023223-2110213013103111-0102220330311200-2331112323002321)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer](data-sources--workload--reference--group-007.md#canonical-3031130303320312-3011210220300230-1112101210130120-3233211133223031-3232211103310222-3322323003212131-0300113203000232-1001321332302222)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through](data-sources--workload--reference--group-007.md#canonical-2213213201020321-0303321203211020-2202012103010022-2302121320212312-1023313101332232-3230002111032132-2311201002200103-3213110102330302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-007.md#canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2003301130333322-2101101031011320-1100101132221102-1230230213222110-3112133132333232-3111013121312102-1310020020113021-3233311312133103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333030101022212-0012103210331333-3000322301120210-1001323232301121-0003012002332003-3100022331332200-0213233213013102-2132310123130211"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options — coalescing_options / 303023130322 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-3130013303313213-3302322223320022-1323001030121102-3212321332123211-1332020301003102-0311030012332233-2331130302332201-3121301313323121"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

<a id="canonical-1131000001123121-1230120203301332-0302332223213231-3203320313200320-0302023021231023-3032110130101313-2200330121123003-3011010233121333"></a>

## Direct properties — coalescing_options / 303023130322 / 3

- [default_coalescing](data-sources--workload--reference--group-007.md#canonical-2020120202212302-0203103023210320-2111111222030002-2212211322121022-0202020233223311-1201103112121023-1000233112231110-1313230233220121): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-007.md#canonical-0302313111113032-1332101201121110-2132322021011032-2102211001003303-3011233030221220-0130010201300210-3201230000213021-3310320223021033): complete subsection reference.

<a id="canonical-3003332113100110-2210111210000331-1033223320202330-2222303230122311-0122211120133002-0010010223233110-3300113030121111-0021221011130303"></a>

## Next pages — coalescing_options / 303023130322 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing](data-sources--workload--reference--group-007.md#canonical-2020120202212302-0203103023210320-2111111222030002-2212211322121022-0202020233223311-1201103112121023-1000233112231110-1313230233220121)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing](data-sources--workload--reference--group-007.md#canonical-0302313111113032-1332101201121110-2132322021011032-2102211001003303-3011233030221220-0130010201300210-3201230000213021-3310320223021033)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2020120202212302-0203103023210320-2111111222030002-2212211322121022-0202020233223311-1201103112121023-1000233112231110-1313230233220121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330320233031112-3320200030132231-1001220213333120-3113013012012030-2310120120133122-2312132011230010-3223310331101203-0113303322120131"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing — default_coalescing / 030232231000 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-007.md#canonical-2003301130333322-2101101031011320-1100101132221102-1230230213222110-3112133132333232-3111013121312102-1310020020113021-3233311312133103)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-2100003222032312-0331332301000102-0232023011031201-3331333213122110-3332022333033130-3001203123221203-3031333123332320-2333100031022030"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-3331311121001032-3321230121002202-0212122111332133-3322320012101020-3023331101322333-1021312322311233-0322322110211130-2230203032131220"></a>

## Direct properties — default_coalescing / 030232231000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303022302210330-2103121222201111-1131330221333302-1200232303001100-3311121133313030-0212321312332120-0011320332133103-0202012310211033"></a>

## Next pages — default_coalescing / 030232231000 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-007.md#canonical-2003301130333322-2101101031011320-1100101132221102-1230230213222110-3112133132333232-3111013121312102-1310020020113021-3233311312133103)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0302313111113032-1332101201121110-2132322021011032-2102211001003303-3011233030221220-0130010201300210-3201230000213021-3310320223021033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101213231221221-1001213012130333-3302011112332333-3101320333100003-1113221320203312-1203331031101133-1103232123101200-0023200001212122"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing — strict_coalescing / 212120330012 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-007.md#canonical-2003301130333322-2101101031011320-1100101132221102-1230230213222110-3112133132333232-3111013121312102-1310020020113021-3233311312133103)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-0200021310311221-3022223301122020-3000020202032313-3102312213202312-1301112111030302-0120023323231022-1020100300111111-2101322220220302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-1211131330022110-1321312331020003-2023333001122322-0223110122311323-3212110222000300-3120231322111213-3113103000333121-3320132231203000"></a>

## Direct properties — strict_coalescing / 212120330012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222333111010031-1233202111130021-2303112222022020-3302300010020210-0301022232033300-1010321133333031-3013021103013102-2021233323302113"></a>

## Next pages — strict_coalescing / 212120330012 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-007.md#canonical-2003301130333322-2101101031011320-1100101132221102-1230230213222110-3112133132333232-3111013121312102-1310020020113021-3233311312133103)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0032013200101331-3121101033132032-0201111023102312-3310122312023202-1011012000331210-3022033220220302-0200212023113213-0311301132011033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321112011103131-2033213112120000-0331333111001300-1200301330123013-0220231021121330-3122130233122222-0003031322211213-1233200303100000"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header — default_header / 231211300331 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-1213203211230211-3222310303202300-0112201021312230-2320000013033323-2331203131011000-1021020220302100-0310302112303031-0333030232221221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

<a id="canonical-0310230103003323-3330101222331021-0132230201231221-3003102330201331-0021131231213210-0010111002022202-1231233023121111-1330013022110222"></a>

## Direct properties — default_header / 231211300331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100133123231101-2113330300032310-0020030133331001-3320312302200030-1021212101002003-1222302120010113-3110030231222002-1033132313323233"></a>

## Next pages — default_header / 231211300331 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2303033023003313-0320312301010320-1232130123120130-1031233120230310-3121203332301102-3100303300110230-2031223030231013-2131312113012223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230022001232223-0312113112203313-3312222030133303-1100332212300133-2230100010112032-3100230002021031-3202233012023303-1330030021223300"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer — default_loadbalancer / 033330232202 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-2032322313312233-0322013121000201-1233132312212212-3332212322230231-2123001121122203-2102200233310330-2210031033303233-2232330233333030"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default loadbalancer.

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

<a id="canonical-3012331131323121-1021203200101201-2132333301123100-3333202313312120-3221321033031121-3201011101330332-2320321213332022-2100021222230002"></a>

## Direct properties — default_loadbalancer / 033330232202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323010032012031-1211133002323110-0022120331330033-1202210112221011-0023110130233323-0322313300233101-1023310113201230-2233032223212102"></a>

## Next pages — default_loadbalancer / 033330232202 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2202303122103033-0101112331121031-2302302322233103-1211312111211132-3001231303232133-2132213010032312-3012203131322300-1130300102322033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110221022121200-3331011301112220-1200231112012231-0322221001112021-0112020320230012-1022310323312220-3200103111030123-0010221233301132"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize — disable_path_normalize / 111301311202 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-0210303030313003-2030132311201230-2230021321113113-3332022102001000-3010230113223023-1323320100020013-3300300201202232-1031011333231101"></a>

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

<a id="canonical-1001100032122212-0301203320332330-3111301313230300-2102333103332120-2101120022100130-1122110113223003-3121200211032020-1113010202031220"></a>

## Direct properties — disable_path_normalize / 111301311202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231002013321033-3313012332011313-2033312010001113-3022032101232202-0130000011113323-2212212120003100-0010313223210030-0033101301311010"></a>

## Next pages — disable_path_normalize / 111301311202 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2100001101020213-1220100210103220-1202213122130010-3121202002331001-1122131123302303-1211021303321321-1020132302133130-1012012113002110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033230030021302-1231302303213012-2300303121321321-3222201310103130-0001311112020220-0330301230001221-3033311012103221-1122002113330220"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize — enable_path_normalize / 330333233213 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-1313031321203132-3333222313302030-1312113032133013-0332202033121103-3201210233021231-1131011003000212-3012200112021213-3330112110300111"></a>

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

<a id="canonical-1201031331303231-0003223032021133-2110220202121121-0330023303000033-2013030331201133-1312232101300101-2330221210221030-3123222123132223"></a>

## Direct properties — enable_path_normalize / 330333233213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232001010321212-3121330113110302-1100130212220210-3202011222033131-2103111223132022-1311200002103111-2232011322033010-1131331301212331"></a>

## Next pages — enable_path_normalize / 330333233213 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231222212133323-1331123303231210-0002033313110232-2332213323232032-0201100300021213-2102313303332333-2221131222203302-0230211012102001"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options — http_protocol_options / 112122221020 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-2100111033000101-0122211330030000-2110030330103100-3111031213320112-0011021232231013-0130302232100032-3133323120012003-1002122332312131"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

<a id="canonical-3222030323330203-2303022302130123-2103131133221022-3001001100032122-2213332233313130-3031012102211010-2021111233023003-3330023031302012"></a>

## Direct properties — http_protocol_options / 112122221020 / 3

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-007.md#canonical-0313330233330203-2210131200320301-0112311232031030-0030310033001000-1111311123002300-1120301203122022-2023103123330113-1210113200121113): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-007.md#canonical-1031212221332330-2102230133301102-3232001300110013-1222113100011112-1303112001102330-2000111220132012-2103131121223001-3302311113122332): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-007.md#canonical-3320031010211320-3011101111303120-1233002033320120-2011222132200301-3123232201023323-2123111102131312-2033100130113132-2003203113012202): complete subsection reference.

<a id="canonical-2331121300232123-2223012232000100-2230203003022133-3310211231322202-2233322300303111-2110122333122331-2321022203331032-2012132301300300"></a>

## Next pages — http_protocol_options / 112122221020 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-007.md#canonical-0313330233330203-2210131200320301-0112311232031030-0030310033001000-1111311123002300-1120301203122022-2023103123330113-1210113200121113)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](data-sources--workload--reference--group-007.md#canonical-1031212221332330-2102230133301102-3232001300110013-1222113100011112-1303112001102330-2000111220132012-2103131121223001-3302311113122332)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](data-sources--workload--reference--group-007.md#canonical-3320031010211320-3011101111303120-1233002033320120-2011222132200301-3123232201023323-2123111102131312-2033100130113132-2003203113012202)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0313330233330203-2210131200320301-0112311232031030-0030310033001000-1111311123002300-1120301203122022-2023103123330113-1210113200121113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120022313330023-3333303313032332-0301013313330313-2011033300212313-2132322301303212-1033103313113021-3323201212321000-3213012102100321"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 121123132011 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-0130322020100113-0021302023333123-0211000313302110-2131232200122001-2133330210210033-1023223212030210-3222233332231000-1033330302210202"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0012001211120130-2221101210311231-0231300312133132-1031223302103121-3103220032312311-2212301320031122-0321300022003110-1200111221020330"></a>

## Direct properties — http_protocol_enable_v1_only / 121123132011 / 3

- [header_transformation](data-sources--workload--reference--group-007.md#canonical-1203331331220112-2131023002033023-0301200012303210-2132100321103323-1011330221333033-2213123000133220-3101102220110123-1320212113020321): complete subsection reference.

<a id="canonical-0203303111333133-0133101221231003-0310121220023023-1232230013212030-1120201200021220-3002302300333003-3333301310030301-3001011013302322"></a>

## Next pages — http_protocol_enable_v1_only / 121123132011 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-007.md#canonical-1203331331220112-2131023002033023-0301200012303210-2132100321103323-1011330221333033-2213123000133220-3101102220110123-1320212113020321)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1203331331220112-2131023002033023-0301200012303210-2132100321103323-1011330221333033-2213123000133220-3101102220110123-1320212113020321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133233322230103-3110233003311311-3303301323230232-3111233200223203-1320310221111100-2233210222102230-3231032012102312-0103301201312112"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 101120002100 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-007.md#canonical-0313330233330203-2210131200320301-0112311232031030-0030310033001000-1111311123002300-1120301203122022-2023103123330113-1210113200121113)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0022211303323132-1222002133111232-2331320232120231-2031010122110020-1333012221333220-1022133210130300-2120320332211100-1123013131302210"></a>

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

<a id="canonical-2023002212002201-3333320332301213-0100330333032320-1131220010222301-3123000030101323-1021032300120302-2302022222011010-1130112002230323"></a>

## Direct properties — header_transformation / 101120002100 / 3

- [default_header_transformation](data-sources--workload--reference--group-007.md#canonical-2323123103323022-0232220111220332-3100302222001000-3230202321222102-1202302031020232-1332022102101223-2121112001332111-2202132232200332): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-007.md#canonical-0012200202002333-2311131131332011-1312223033312002-3301223023332033-2120233112302231-2130320201012022-2221221311130003-0023013023223032): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-007.md#canonical-2213131103121312-1022113221103330-3321130300100332-1322010130302300-1310020232332330-1213122321110011-2332022013001021-0212121311112212): complete subsection reference.

<a id="canonical-3111133300010113-2203321212223203-0231132323131131-3130103311031301-3321033032012320-1321210103201023-3202001302221211-0023201210101310"></a>

## Next pages — header_transformation / 101120002100 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--workload--reference--group-007.md#canonical-2323123103323022-0232220111220332-3100302222001000-3230202321222102-1202302031020232-1332022102101223-2121112001332111-2202132232200332)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--workload--reference--group-007.md#canonical-0012200202002333-2311131131332011-1312223033312002-3301223023332033-2120233112302231-2130320201012022-2221221311130003-0023013023223032)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--workload--reference--group-007.md#canonical-2213131103121312-1022113221103330-3321130300100332-1322010130302300-1310020232332330-1213122321110011-2332022013001021-0212121311112212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-007.md#canonical-0313330233330203-2210131200320301-0112311232031030-0030310033001000-1111311123002300-1120301203122022-2023103123330113-1210113200121113)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2323123103323022-0232220111220332-3100302222001000-3230202321222102-1202302031020232-1332022102101223-2121112001332111-2202132232200332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332131032130120-2001333310120323-3311123212122233-3120233021211230-0210330323312233-0223322111322320-3020032100322032-0010202222320102"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 210002220111 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-007.md#canonical-0313330233330203-2210131200320301-0112311232031030-0030310033001000-1111311123002300-1120301203122022-2023103123330113-1210113200121113)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-007.md#canonical-1203331331220112-2131023002033023-0301200012303210-2132100321103323-1011330221333033-2213123000133220-3101102220110123-1320212113020321)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-1301311113011003-2133320121330213-2013320322010013-1031320110132103-3222022122213010-2133120032003013-2311333100001310-2200310031231122"></a>

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

<a id="canonical-1323200130133012-0233231233302113-1112232220300212-0210212333131021-3210213022212121-2031313101230230-0233032323312022-1212132222001103"></a>

## Direct properties — default_header_transformation / 210002220111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122033301213002-0322333033131130-3102110123301200-2201220121123232-2310322003133001-0303111320311230-2110301322103300-1112100320233023"></a>

## Next pages — default_header_transformation / 210002220111 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-007.md#canonical-1203331331220112-2131023002033023-0301200012303210-2132100321103323-1011330221333033-2213123000133220-3101102220110123-1320212113020321)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0012200202002333-2311131131332011-1312223033312002-3301223023332033-2120233112302231-2130320201012022-2221221311130003-0023013023223032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133120001033332-2222233303321300-3300200332230332-2101103223323103-1312000322212310-2331323132310110-3230320003332310-1123130112300013"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 302211011130 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-007.md#canonical-0313330233330203-2210131200320301-0112311232031030-0030310033001000-1111311123002300-1120301203122022-2023103123330113-1210113200121113)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-007.md#canonical-1203331331220112-2131023002033023-0301200012303210-2132100321103323-1011330221333033-2213123000133220-3101102220110123-1320212113020321)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-1123211312233133-3112223321202203-1002200332210223-3001032210333203-1231320133220030-2103300210110003-3210130132123033-3322310103010131"></a>

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

<a id="canonical-2113002121212333-3200233312100032-1113232221213230-2132023332320022-0212311233023200-2122320133113202-0200210231331013-1233120333123032"></a>

## Direct properties — preserve_case_header_transformation / 302211011130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221001032022113-3202012131102300-3022021002310123-0320223103102320-2121023211121121-1003301211123331-3013202212012311-1030233231012313"></a>

## Next pages — preserve_case_header_transformation / 302211011130 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-007.md#canonical-1203331331220112-2131023002033023-0301200012303210-2132100321103323-1011330221333033-2213123000133220-3101102220110123-1320212113020321)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2213131103121312-1022113221103330-3321130300100332-1322010130302300-1310020232332330-1213122321110011-2332022013001021-0212121311112212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301232312201211-2213203123221010-2000333301032000-1222330222001312-1233012132011011-2331010021310233-2301203112120031-3300202131230332"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 300331031211 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-007.md#canonical-0313330233330203-2210131200320301-0112311232031030-0030310033001000-1111311123002300-1120301203122022-2023103123330113-1210113200121113)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-007.md#canonical-1203331331220112-2131023002033023-0301200012303210-2132100321103323-1011330221333033-2213123000133220-3101102220110123-1320212113020321)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-1230333133100031-3312012213121101-3110101013021003-1110122131132301-3123223301101013-3202203232022221-0001333001211031-0212313331112310"></a>

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

<a id="canonical-1321111101210000-3100200021130212-3201000122010132-2122021133132130-1220130030203332-3313002202020000-0133332311123202-3322022021312131"></a>

## Direct properties — proper_case_header_transformation / 300331031211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2310012200131103-1333133220003103-2220212322132133-1110302123210222-2333313233201300-2211030101130121-3113133123313301-1220113012023012"></a>

## Next pages — proper_case_header_transformation / 300331031211 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-007.md#canonical-1203331331220112-2131023002033023-0301200012303210-2132100321103323-1011330221333033-2213123000133220-3101102220110123-1320212113020321)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1031212221332330-2102230133301102-3232001300110013-1222113100011112-1303112001102330-2000111220132012-2103131121223001-3302311113122332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031002221330123-3101220211111221-3230022111030222-0232201021113002-0131000222111032-3313032122223031-1211030210201221-0113032321123003"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 210213233331 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-3113101330200330-1232203333212313-3122012131132003-3331321230201213-2100313111112201-2012303122311001-2001332130312202-2221100010202331"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

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

<a id="canonical-0023311121232330-3233111221220013-0211312221013001-2312132102303031-2220323023110333-3323112333231130-0020102022220300-0113120021220000"></a>

## Direct properties — http_protocol_enable_v1_v2 / 210213233331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120001133012011-3102013333212323-3323113322022323-3303010301223130-2102222312213201-2320132223013201-1331000022113232-3111302010001200"></a>

## Next pages — http_protocol_enable_v1_v2 / 210213233331 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3320031010211320-3011101111303120-1233002033320120-2011222132200301-3123232201023323-2123111102131312-2033100130113132-2003203113012202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332313112003101-2000302100322323-3202100012331212-1203100103201000-0203000031333230-0232201033200132-2032323020030230-0313113013212332"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 322223100233 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-3023231221013111-2213111001203120-0232010001122112-3302031212331200-2311200310112013-1232120030023321-1232323133110221-3030321203001213"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

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

<a id="canonical-3111203001011132-1113031121030113-3303000030113111-0330012023032121-1003121213233201-2102220231330223-2032031203021301-0113122013002110"></a>

## Direct properties — http_protocol_enable_v2_only / 322223100233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122013101333122-3332312223203200-1232313311202212-3032013013303220-2231020120231233-1302111002101200-2123201033132310-0322113012031230"></a>

## Next pages — http_protocol_enable_v2_only / 322223100233 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-007.md#canonical-0323232133231230-3022303100131011-2320322220320221-3331202202011020-0203323311110200-0212313030201321-1011333211321030-2132121133003133)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0231223001232221-2201332131123020-1313300302020312-0221201321000320-0033233210023223-2110213013103111-0102220330311200-2331112323002321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321301113223332-3221122011133002-2003133123133021-3032132311122300-3131302102210313-2222210320301023-2001100032312122-2001133232100112"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls — no_mtls / 102232013213 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls

<a id="canonical-3311322301021120-0310223201221103-1311202122202221-0021301331302202-2130332110222220-2330302012123201-1000300332231130-1320322210332313"></a>

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

<a id="canonical-1012330330332301-2132100213002122-0210021203312223-2123332032112032-0313020302022012-1302031220102300-0212122133312113-0202312020101222"></a>

## Direct properties — no_mtls / 102232013213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212023221132003-3113211032133133-1210212203011313-0220320033020133-0221321030312010-2030322303133111-3310230003003200-2100132320100012"></a>

## Next pages — no_mtls / 102232013213 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3031130303320312-3011210220300230-1112101210130120-3233211133223031-3232211103310222-3322323003212131-0300113203000232-1001321332302222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301130203210000-3221302320001113-2112101013213013-1220000101323212-0011012323110221-3101220302122001-0030010102200312-2010223333313031"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer — non_default_loadbalancer / 132122130312 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer

<a id="canonical-0233011301010011-2201203111202133-3102020022301303-2321333002102212-3322311002321213-3312020232211001-1200013330210330-0200313302223100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

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

<a id="canonical-0002003210033202-2203020231013210-2131221121023331-0133123001310112-1020323322012100-0323210030101123-0320023002100333-3133303111211030"></a>

## Direct properties — non_default_loadbalancer / 132122130312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233313230212021-2022123130311010-3000212200202033-2333132112211111-2301233030230123-0032111021110332-0113303200123123-0010320231133133"></a>

## Next pages — non_default_loadbalancer / 132122130312 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2213213201020321-0303321203211020-2202012103010022-2302121320212312-1023313101332232-3230002111032132-2311201002200103-3213110102330302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201023331202012-1011223023111013-0103002312021130-3333320022031223-1031301001232222-0232332320212020-3121303031200333-3130132333132303"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through — pass_through / 120212310323 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through

<a id="canonical-2010212131111300-0112030223010112-0203001303303010-2301030230002003-2103101123121222-3322330120132303-0131021022322110-0010303232022010"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

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

<a id="canonical-0021233301001233-0101001101021131-2213103100112233-1013212300322111-0013013121302230-1103312100102100-2320002231320000-2020310100210331"></a>

## Direct properties — pass_through / 120212310323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021132023011212-3332113021230101-3313331233121002-3311323313121322-2212020302312320-2113233003130100-0213130221310330-3331211303303100"></a>

## Next pages — pass_through / 120212310323 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033130131330030-1310110020301231-3100021031212211-2303322021002013-2022013301011022-2203200023212232-2112302203311212-0222131110310312"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config — tls_config / 210033232233 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config

<a id="canonical-0031221303330331-2120132321333323-2130101022302100-2332130101101010-3331212021120222-2300232100211100-3000321103002032-1302110202101203"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

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

<a id="canonical-2322202032301130-2001130003210120-1330332103123122-0011210102221211-0012012330110010-1302310120321200-0230010113030221-1211201323132120"></a>

## Direct properties — tls_config / 210033232233 / 3

- [custom_security](data-sources--workload--reference--group-007.md#canonical-1122123112031113-1311322231112010-1311313021221333-1332011303002300-0033333233222000-2220321100231003-0001301112320010-1330313001100032): complete subsection reference.

- [default_security](data-sources--workload--reference--group-007.md#canonical-3110301022202322-3120003223122202-0210213132332120-1122122112330322-1022112221330300-1323121311122332-1113213001001211-3330333200000031): complete subsection reference.

- [low_security](data-sources--workload--reference--group-007.md#canonical-1222030000033211-0230203211002333-3000321123130130-0120202320022023-1210033033100201-2331311023003130-0133111221033110-0123231022112023): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-007.md#canonical-3010321212113333-1303302230102330-1213023330002121-1130231200013033-3300232122111101-0221100211110312-2123031031021310-3223010111221222): complete subsection reference.

<a id="canonical-3101003020223103-0112333213311222-0323211021003202-0222232113032031-2230110313020012-1333110132111132-0000321220230002-1023102100222302"></a>

## Next pages — tls_config / 210033232233 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security](data-sources--workload--reference--group-007.md#canonical-1122123112031113-1311322231112010-1311313021221333-1332011303002300-0033333233222000-2220321100231003-0001301112320010-1330313001100032)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security](data-sources--workload--reference--group-007.md#canonical-3110301022202322-3120003223122202-0210213132332120-1122122112330322-1022112221330300-1323121311122332-1113213001001211-3330333200000031)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security](data-sources--workload--reference--group-007.md#canonical-1222030000033211-0230203211002333-3000321123130130-0120202320022023-1210033033100201-2331311023003130-0133111221033110-0123231022112023)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security](data-sources--workload--reference--group-007.md#canonical-3010321212113333-1303302230102330-1213023330002121-1130231200013033-3300232122111101-0221100211110312-2123031031021310-3223010111221222)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1122123112031113-1311322231112010-1311313021221333-1332011303002300-0033333233222000-2220321100231003-0001301112320010-1330313001100032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200220223323020-3121303200323332-1030231023010013-0113003302023021-1333100023320132-2133202113133230-1331201002121322-0000210030112121"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security — custom_security / 310311021102 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-007.md#canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security

<a id="canonical-3120031211203232-2010320231212022-3010000013013002-1312133223332010-2333213003303010-2230201101322300-2232110100021032-1031033222213102"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

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

<a id="canonical-1110132100023312-1322112120212110-2030110132313323-3310300203312233-2230203021130212-2011231133210222-3132110100012021-1303000113122322"></a>

## Direct properties — custom_security / 310311021102 / 3

<a id="canonical-2223130123313210-0020200321033032-3120130223320313-2223322322103221-2223010202011231-0003211203132130-0332103032121011-0133313300312102"></a>

<a id="canonical-0202100210113211-0130310212211003-2232122010101110-2331330213130020-3221123012011123-0203000122200030-1031211323212201-0222030311221200"></a>

## cipher_suites property — custom_security / 310311021102 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1110030212013200-2022203023311131-2103110311023232-1031330201231320-3033233302101301-2203210200312012-3300211322311231-3021301130023121"></a>

<a id="canonical-1212200313231313-0220131030220212-0313223321300023-0123232201311223-1312302313223023-1230321301332201-3003303313032031-1133203121021310"></a>

## max_version property — custom_security / 310311021102 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-3303222103022001-3031123032202330-0311001320232300-1133230231211113-0010221323211313-2112013013130311-1232203121322232-1010232033030200"></a>

<a id="canonical-2033303331131020-2300311203011001-1211113201013331-2110021301022331-1203113112312002-2300021131202303-3231322112221312-3311222002232233"></a>

## min_version property — custom_security / 310311021102 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-0322121332110211-1122121020103022-2021210213213120-0203122003232200-1332210021122001-0313210002131113-1033331223200110-0003133202302212"></a>

## Next pages — custom_security / 310311021102 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-007.md#canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3110301022202322-3120003223122202-0210213132332120-1122122112330322-1022112221330300-1323121311122332-1113213001001211-3330333200000031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123121010022303-2100200331301102-0121132101332101-0321221211232122-1111322321023231-1003122221232200-2221110012111020-3102001322102232"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security — default_security / 232213013123 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-007.md#canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security

<a id="canonical-3112103320132031-0103133202203211-0312303320022031-3222103210202111-2321333032023032-1233233103002303-2111221011032201-2122022231321210"></a>

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

<a id="canonical-0022323303303310-0111133312333030-0100023032021233-0133120223133002-1122022312301121-1231203320000321-3001012320030333-0113103302012310"></a>

## Direct properties — default_security / 232213013123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311222031033231-2230022311330023-0230222323101003-0232201210203223-0031023100203223-0032130020231201-2101033201230203-2033310220332200"></a>

## Next pages — default_security / 232213013123 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-007.md#canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1222030000033211-0230203211002333-3000321123130130-0120202320022023-1210033033100201-2331311023003130-0133111221033110-0123231022112023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202110021233211-1001110221201112-1220130101121030-2111301233121002-1123001203233312-2120233201301330-2201321113210231-1112323031220311"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security — low_security / 101010301121 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-007.md#canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security

<a id="canonical-0010130220213121-1223133030201130-2212330102010101-3303013003331210-2231321230103200-2311331210220012-0031213200012101-2202201122002333"></a>

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

<a id="canonical-2330132322021123-1003322110321023-1302120121223332-0201320301102113-2231323223303233-1333013100332013-3011010213201303-1303131111222221"></a>

## Direct properties — low_security / 101010301121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300002322132033-1331010103223202-1211113032310032-1120300323223310-2111331333313212-2332101202333232-0031001301332210-1001200103301123"></a>

## Next pages — low_security / 101010301121 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-007.md#canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3010321212113333-1303302230102330-1213023330002121-1130231200013033-3300232122111101-0221100211110312-2123031031021310-3223010111221222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012200023121012-0102030133301330-0120233330230023-2021221113012311-3012221202213302-0010210123113022-3222133213330201-0211300130010333"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security — medium_security / 221131202302 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-007.md#canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security

<a id="canonical-0222322111133123-0032222223222303-2321032220213233-3332033131123310-3232213231022020-0013230002233123-0030112320033302-1233103131012013"></a>

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

<a id="canonical-1213002211133203-1100311311120200-0313123200313330-0003131313011111-1112210021303323-0000230310022000-3322222003032022-0020023130210301"></a>

## Direct properties — medium_security / 221131202302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033200123000221-2003111322032003-1333100310332210-0110310010122233-0320311212313013-0003221313032221-3010330113312013-3220000211001312"></a>

## Next pages — medium_security / 221131202302 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-007.md#canonical-3011311310023010-1302103332310101-0313000333303201-3123210003312103-0032023211103202-3221230031111121-0002120321110302-2310202032322030)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030032232300210-3003222221322131-0313120101011123-3132011320220130-3101302001000210-0030020133312213-0321231133202110-1001103331110302"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls — use_mtls / 000033211310 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls

<a id="canonical-3210223312013303-3213002100203233-1310321002102031-1132230203011023-3013032310122030-2011313022130220-1122113212333133-1100223022213203"></a>

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

<a id="canonical-3323333210121221-1000131310102302-3301202030300011-3011132132221332-2111123122301223-2022333133000031-0101101222111102-0010112100102131"></a>

## Direct properties — use_mtls / 000033211310 / 3

<a id="canonical-1132110321310320-0011220103211020-1313200231233020-0301223201321320-3332123233001110-1320211221312001-3223103221103020-2320313333211300"></a>

<a id="canonical-0301300102003121-1001223030230221-0310131322220110-2012021321130332-2001133233301201-2122201223132103-2323010233022322-2003121312032001"></a>

## client_certificate_optional property — use_mtls / 000033211310 / 4

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

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

- [crl](data-sources--workload--reference--group-007.md#canonical-2321000000031102-3203222333122132-2203200210333132-2000101211323221-1100101121002222-1303302132031121-2201001000232031-1231122122132000): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-007.md#canonical-1213223112013210-0233221103323322-2101133230011202-1111203101300231-2021331221302330-3133100001330003-1032200101012223-2212313120122303): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-007.md#canonical-3322301003313000-0331132313112123-3301002020332202-2022220001000011-3032101221331121-2000011120023012-0023023311320330-2123103003232213): complete subsection reference.

<a id="canonical-0230330131022200-1201133323101323-1230132210120130-1123020223001322-1321002232231301-2003231103230220-2332133010021110-3210110122310310"></a>

<a id="canonical-0013010331031210-2030312131011311-0213031123300023-0310013230221113-1233012210302101-3111130212130111-0213301112203223-0133003213031222"></a>

## trusted_ca_url property — use_mtls / 000033211310 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

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

- [xfcc_disabled](data-sources--workload--reference--group-007.md#canonical-3331323222331113-1111323202020100-1330010100013220-0123312320332130-2321310322332231-0311101220122021-3300002302031313-0021132000310301): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-007.md#canonical-3311120233030132-3103210132011111-3331020301021232-0201010233332213-1111233322120100-0311320320303323-1230032102001101-0121220011121313): complete subsection reference.

<a id="canonical-2322210013203320-1121111313313101-1313231123313130-0111110010020001-1103213233120100-1311203220031123-0321002201021212-3031021121300002"></a>

## Next pages — use_mtls / 000033211310 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl](data-sources--workload--reference--group-007.md#canonical-2321000000031102-3203222333122132-2203200210333132-2000101211323221-1100101121002222-1303302132031121-2201001000232031-1231122122132000)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl](data-sources--workload--reference--group-007.md#canonical-1213223112013210-0233221103323322-2101133230011202-1111203101300231-2021331221302330-3133100001330003-1032200101012223-2212313120122303)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca](data-sources--workload--reference--group-007.md#canonical-3322301003313000-0331132313112123-3301002020332202-2022220001000011-3032101221331121-2000011120023012-0023023311320330-2123103003232213)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled](data-sources--workload--reference--group-007.md#canonical-3331323222331113-1111323202020100-1330010100013220-0123312320332130-2321310322332231-0311101220122021-3300002302031313-0021132000310301)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options](data-sources--workload--reference--group-007.md#canonical-3311120233030132-3103210132011111-3331020301021232-0201010233332213-1111233322120100-0311320320303323-1230032102001101-0121220011121313)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2321000000031102-3203222333122132-2203200210333132-2000101211323221-1100101121002222-1303302132031121-2201001000232031-1231122122132000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200223103133200-1301013223123103-3121323111210021-3111023333221101-1310013322130011-0301021122131001-1100210320110330-1002233213022202"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl — crl / 131013202222 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl

<a id="canonical-0122002311033323-1000220103123122-0200033210200033-1132211320232333-1310012221002011-2032230223332100-3222312133032101-3002001033022003"></a>

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

<a id="canonical-3033201121203100-3331202212313021-3112302323332122-3230300111111312-2332033230000232-2333121133131001-2312022100122130-0103210210230030"></a>

## Direct properties — crl / 131013202222 / 3

<a id="canonical-0332032223322023-1211100232002032-0231301302120200-2301213111112003-2220100101121122-0123032211030030-0311122012022122-3112033121201122"></a>

<a id="canonical-3312020313202322-1133212303100101-2203112122112020-3111311133333233-0020131130012323-2332010010312322-0310312132320321-0223102232301002"></a>

## name property — crl / 131013202222 / 4

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

<a id="canonical-2022200122011223-2302012303232012-1313003122301330-2300323220212021-1110332320013212-3122203320200321-3111222331132302-1200331200213231"></a>

<a id="canonical-0130012230121101-3221212103130302-0110023020203210-0230200301220110-1020022303020110-3133032000320011-2333300122031030-3323321200233222"></a>

## namespace property — crl / 131013202222 / 5

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

<a id="canonical-0222112222103212-0302300133022021-2200030031213101-3031013310222203-2031230312113313-3230210322022312-2202223200122030-2331012121020321"></a>

<a id="canonical-0010101133320100-1102212210230313-2231013123100322-2103131011203323-3231010200122230-0333300323303012-1231113321300202-3002132203233331"></a>

## tenant property — crl / 131013202222 / 6

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

<a id="canonical-3031020130312312-3132322110311302-1310112211300322-0332000121133030-0321110311303000-0310231102231012-2111301210011333-2213331212000110"></a>

## Next pages — crl / 131013202222 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1213223112013210-0233221103323322-2101133230011202-1111203101300231-2021331221302330-3133100001330003-1032200101012223-2212313120122303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003202010320002-2331211201210230-1010201102103011-0231110330102000-2210212332002322-3213312302233103-1322311203332203-2222113030002022"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl — no_crl / 033013211321 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl

<a id="canonical-2300312203223212-1300103231113020-3302331110211031-2010123031031032-1121003300321221-3331130202220113-0212330321230231-0123132103331101"></a>

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

<a id="canonical-3010110303222113-1203320332032212-3302003131033021-3131111311332022-2100203331111002-0133232122333322-2331112033222003-0221120003122210"></a>

## Direct properties — no_crl / 033013211321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100010102312110-3302310333000103-1111101202212111-0002233203213212-0130310013102210-2300131103100333-3033223010022030-2000013202031130"></a>

## Next pages — no_crl / 033013211321 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3322301003313000-0331132313112123-3301002020332202-2022220001000011-3032101221331121-2000011120023012-0023023311320330-2123103003232213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331113121022200-3300030300122330-0310022303132012-1002200000030102-0112122002113230-2110000122203330-0031322022321003-3203222230013332"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca — trusted_ca / 210110013113 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-1303000121113201-3321230212202302-0003200100103032-0302321032312003-3220332302120333-2011300202200002-0220030311030222-2322321113222321"></a>

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

<a id="canonical-1230303102011312-0213212203102010-1231101010201123-0012232033300222-2132001322210230-3203003302113023-0310000131301200-1331321023121112"></a>

## Direct properties — trusted_ca / 210110013113 / 3

<a id="canonical-2333312313332133-0330310100202003-2032231213011123-1231301131213312-2023212302011013-2222203333113022-2212301202011022-2120223302221110"></a>

<a id="canonical-1123102111010311-1113011110300000-0330002030011133-1212231122323130-3300022102133010-3021101203120011-1130223231321121-1003202033322320"></a>

## name property — trusted_ca / 210110013113 / 4

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

<a id="canonical-2301121132132113-3223102133302000-1110301330010013-1100122223303001-1320321032320201-2032213213301211-0021012112320132-1203030102211220"></a>

<a id="canonical-3213132320123232-3203221033320032-3130222113010313-2021020133032112-0230122320010021-1103121010003100-0230010222212223-1133321103313111"></a>

## namespace property — trusted_ca / 210110013113 / 5

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

<a id="canonical-3000133322120022-1032221210032321-0210122302210130-3020000323013213-2022120213301130-2300112200113123-2121100200222110-2212331103223322"></a>

<a id="canonical-1132131022133132-3223112012021322-2033132032110102-2213202201223010-0000301331230002-3110213202021002-1131003033112133-0230120211120310"></a>

## tenant property — trusted_ca / 210110013113 / 6

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

<a id="canonical-2003110312113200-2021103123211232-1131130101033320-2210232113010301-2331312333233303-1311323333231312-3312202311202100-1011100110302003"></a>

## Next pages — trusted_ca / 210110013113 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3331323222331113-1111323202020100-1330010100013220-0123312320332130-2321310322332231-0311101220122021-3300002302031313-0021132000310301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312112022222010-3233231210020012-0110310011012212-0212020131303113-0210221032110233-0333022332302003-3130212122010121-1331022002211103"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled — xfcc_disabled / 031220012021 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-1000322312111210-0211323022022332-2100331110230333-3011023101332121-2130233013032312-1133213312112130-1130203110220202-3112011232212122"></a>

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

<a id="canonical-0302332010123201-0313102013202320-1332133110103101-0333321313013102-0311311201003233-1311103021000232-3301211232232131-0333322030212111"></a>

## Direct properties — xfcc_disabled / 031220012021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333120211132222-3031111302021321-0030222230210013-2031100122133021-1023010300111330-1233303133102200-3132010330312002-2310003200012321"></a>

## Next pages — xfcc_disabled / 031220012021 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3311120233030132-3103210132011111-3331020301021232-0201010233332213-1111233322120100-0311320320303323-1230032102001101-0121220011121313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011001020002102-1202020302220233-1323212220022310-3230202132002223-3003111200320001-1232122100002300-3022030212121230-1033230313322322"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options — xfcc_options / 032230022310 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-006.md#canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-1131232031011203-0202313231003002-2321322231031312-0033233322202133-3012322003313330-2020020233022023-0121020330030323-0220100021110220"></a>

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

<a id="canonical-2021200300020110-3011120201231320-1120200102221112-1112230111200033-0013111030000312-2002222133313033-1101122211133112-3022231110122033"></a>

## Direct properties — xfcc_options / 032230022310 / 3

<a id="canonical-3320101012301111-2023111220123312-3210002000101200-0313322320121201-0313130032223113-0110001303010022-1011232211013322-0201333110212313"></a>

<a id="canonical-3002031203121312-2303023011033322-3011021311210223-0033100113232330-3330120332131001-0020133333210232-1303101312021203-1021122211220301"></a>

## xfcc_header_elements property — xfcc_options / 032230022310 / 4

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="canonical-0212312112030300-0223131002220202-2012323120303120-0220020200122203-3121211122112121-0003132020132102-2131231010022113-0212312231223321"></a>

## Next pages — xfcc_options / 032230022310 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-007.md#canonical-0213220033202012-3011300310133230-3033211333110333-1202213000021301-0201020110303021-1300113132302223-3213033232020010-3212230023230221)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320333002230013-1302101012111221-2100023202313133-0303233300303232-1300321230011023-1223232211223303-1101031000211103-3002332330020200"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes — specific_routes / 102112320131 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes

<a id="canonical-3300032310020202-3222222111323200-3232230300021332-0020030313020100-1032203000002030-0121223123121011-2330233221121002-0001000011032100"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to define a route.

Upstream description:

This defines various OPTIONS to define a route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1223121213013210-1130333010321332-2312102200113333-0130012121010320-2121032221121102-2233312021020223-1112323302002011-1021120012210210"></a>

## Direct properties — specific_routes / 102112320131 / 3

- [routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212): complete subsection reference.

<a id="canonical-3011322332330331-0320003210313313-0302031222132021-1113120323202323-0120202120013210-3133110330232121-2212012231201302-3020020211133102"></a>

## Next pages — specific_routes / 102112320131 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103100202223123-1313220120013203-2232230001303231-0110030001233202-0203221211032302-1302310113233203-2120313312312033-1233030222320301"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes — routes / 200131102011 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes

<a id="canonical-1103310133211320-1200303022033030-0301101133212303-1330231003331120-2112110012020222-1003013320302212-0011021212332320-2111120003303031"></a>

Type: `"list"`. Computed.

Routes. Routes for this loadbalancer.

Upstream description:

Routes for this loadbalancer.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-1202310323322212-0231323002322232-1212102003013233-1322230333131201-3310002311011031-3000311112203011-1222302232312021-2211023220111230"></a>

## Direct properties — routes / 200131102011 / 3

- [custom_route_object](data-sources--workload--reference--group-007.md#canonical-3310301321330002-3202101101112322-1121213312013012-2033010312102123-2030330103220221-3223121123231031-1203300131321221-0012023103032321): complete subsection reference.

- [direct_response_route](data-sources--workload--reference--group-007.md#canonical-3113321211103300-3121123312121102-2103230121222320-0023030100232313-0030331001233112-3103020012203031-1211333322302002-3323010310210312): complete subsection reference.

- [redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132): complete subsection reference.

- [simple_route](data-sources--workload--reference--group-008.md#canonical-0310112122113210-2110103333320012-3100102321331121-2303102223302303-0033320103233333-1221123312232313-2132013023311033-1120032010201221): complete subsection reference.

<a id="canonical-3320121203223033-2233302120111132-3312300232021121-1231233223011021-3313112202120030-1211111130111221-2200210331321220-2310210023230212"></a>

## Next pages — routes / 200131102011 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-007.md#canonical-3310301321330002-3202101101112322-1121213312013012-2033010312102123-2030330103220221-3223121123231031-1203300131321221-0012023103032321)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-007.md#canonical-3113321211103300-3121123312121102-2103230121222320-0023030100232313-0030331001233112-3103020012203031-1211333322302002-3323010310210312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-008.md#canonical-0310112122113210-2110103333320012-3100102321331121-2303102223302303-0033320103233333-1221123312232313-2132013023311033-1120032010201221)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3310301321330002-3202101101112322-1121213312013012-2033010312102123-2030330103220221-3223121123231031-1203300131321221-0012023103032321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311320212020121-1132313231212001-0103203232203331-1331131022021021-2103110322201332-0032333133321310-0212033201102102-1031330230101302"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object — custom_route_object / 212213321120 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object

<a id="canonical-0303230120203313-1101131022103103-1212010331323221-3123232013113302-3130311122132120-0200230113031202-2102022023323212-0233301313012113"></a>

Type: `"single"`. Computed.

Custom route uses a route object created outside of this view.

Upstream description:

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

<a id="canonical-0012133200121330-3100221030331113-3120220230111012-3031310023010312-0001030312333311-3130031333321200-1011222220123312-3312111302101033"></a>

## Direct properties — custom_route_object / 212213321120 / 3

- [caching_disable](data-sources--workload--reference--group-007.md#canonical-3033221110010120-0000213123333022-2212203332220022-0212200211130030-3033123113222311-1302020323102320-3320211300022231-0212023321202201): complete subsection reference.

- [caching_inherit](data-sources--workload--reference--group-007.md#canonical-1310213032101203-3202230001230030-2333033331321112-1100211203120120-0221021123003333-2112112210132320-2332012020103010-1003032032131313): complete subsection reference.

- [route_ref](data-sources--workload--reference--group-007.md#canonical-1111320223332200-0311233203030221-1010332113310300-2311002021001232-0212212120201020-0120103313011213-2322330113202223-1222023101231301): complete subsection reference.

<a id="canonical-3231100320101013-2211001020303231-0100130021101220-0031303023333222-0101321013303103-3002121103233033-3330100033303113-2021212030121213"></a>

## Next pages — custom_route_object / 212213321120 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable](data-sources--workload--reference--group-007.md#canonical-3033221110010120-0000213123333022-2212203332220022-0212200211130030-3033123113222311-1302020323102320-3320211300022231-0212023321202201)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit](data-sources--workload--reference--group-007.md#canonical-1310213032101203-3202230001230030-2333033331321112-1100211203120120-0221021123003333-2112112210132320-2332012020103010-1003032032131313)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref](data-sources--workload--reference--group-007.md#canonical-1111320223332200-0311233203030221-1010332113310300-2311002021001232-0212212120201020-0120103313011213-2322330113202223-1222023101231301)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3033221110010120-0000213123333022-2212203332220022-0212200211130030-3033123113222311-1302020323102320-3320211300022231-0212023321202201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313312300332311-2323103212022112-0100222122202330-0321211001222021-1220201323030331-1301211023333030-3233302300211001-0113012323123030"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable — caching_disable / 210000011011 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-007.md#canonical-3310301321330002-3202101101112322-1121213312013012-2033010312102123-2030330103220221-3223121123231031-1203300131321221-0012023103032321)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

<a id="canonical-2022123120232312-1133301120201131-0321323020203122-1011033113021100-0232012100021021-0311120030030132-2220113231201300-3221012210121230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching disable.

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

<a id="canonical-3101201131300303-3312000121301320-3202131020211202-3303031123301203-3301022002312133-3023331110200233-0113313111310032-1102200033133030"></a>

## Direct properties — caching_disable / 210000011011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000232010022021-1131012030302202-0120111013133322-3033211030122312-1333131200102001-1313202033100122-2003001322030222-0121201231300203"></a>

## Next pages — caching_disable / 210000011011 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-007.md#canonical-3310301321330002-3202101101112322-1121213312013012-2033010312102123-2030330103220221-3223121123231031-1203300131321221-0012023103032321)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1310213032101203-3202230001230030-2333033331321112-1100211203120120-0221021123003333-2112112210132320-2332012020103010-1003032032131313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012123312230023-3223012203020020-3123032212021210-2321131201011202-2300332021113210-1133201120102011-3000223213122203-0012000330112223"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit — caching_inherit / 120000010221 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-007.md#canonical-3310301321330002-3202101101112322-1121213312013012-2033010312102123-2030330103220221-3223121123231031-1203300131321221-0012023103032321)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit

<a id="canonical-3323113122202130-2133000001032120-1012113310331022-3303011012001030-2230111030331330-0001313302230300-1320330130103130-1101123303101132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching inherit.

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

<a id="canonical-2233333312130230-0010201133121310-2213300202230112-3312101202231130-3330022123311333-1113323000212030-0032223331331010-0130022121200220"></a>

## Direct properties — caching_inherit / 120000010221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112311311202313-2232200302333123-0131230102301110-3203012321232222-0223313103113232-1330030232012022-1201212111030021-2022030101020003"></a>

## Next pages — caching_inherit / 120000010221 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-007.md#canonical-3310301321330002-3202101101112322-1121213312013012-2033010312102123-2030330103220221-3223121123231031-1203300131321221-0012023103032321)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1111320223332200-0311233203030221-1010332113310300-2311002021001232-0212212120201020-0120103313011213-2322330113202223-1222023101231301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111132111113320-0023012301230303-1031023113230022-1221120233201311-1313000101232031-1221202231322333-3003113033033231-2012133033032000"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref — route_ref / 211331201023 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-007.md#canonical-3310301321330002-3202101101112322-1121213312013012-2033010312102123-2030330103220221-3223121123231031-1203300131321221-0012023103032321)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref

<a id="canonical-1110100212223133-0000033001301201-3232300020200233-2101030222033010-1022323203003103-3103310102131113-2013322321110323-1200332011001002"></a>

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

<a id="canonical-2123220312313123-3323200010110333-2331102311330111-1133013022020232-2030021313311303-0322220023321101-2030200103132003-3113003312031130"></a>

## Direct properties — route_ref / 211331201023 / 3

<a id="canonical-0322122002211100-0013031333310302-2101002112002130-2100023331203200-2202220310310222-2030111020101012-0123133130133122-1310332002322123"></a>

<a id="canonical-1211110301203032-3020111332121003-1233012211130213-3121333213323130-3220310121103131-0110033211312320-1220213110102000-2123032032021010"></a>

## name property — route_ref / 211331201023 / 4

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

<a id="canonical-1123120321121132-1130100030113310-1002032331013330-0223033303223230-2030323001210321-0311113101000020-0130033231122002-2101210102031120"></a>

<a id="canonical-0112000312230312-2322203012103030-3303113003101303-2202312002303212-1322103013033112-2311021132232233-0121312012231322-2121231300121111"></a>

## namespace property — route_ref / 211331201023 / 5

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

<a id="canonical-1233200031213312-0212202001202101-2023330021321121-1032002221311333-1100323303331101-3321303212210223-0021022132123322-2102111330011223"></a>

<a id="canonical-2022222031332231-1321131121122001-1133313012130230-0323300031113110-3320232100233230-0202111003201000-1300323012100212-1330002113130022"></a>

## tenant property — route_ref / 211331201023 / 6

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

<a id="canonical-3320120200000223-3032123310131212-1230102203200020-1211311101030032-1310123121223330-2312030210230221-2011131102232131-1231230133001122"></a>

## Next pages — route_ref / 211331201023 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-007.md#canonical-3310301321330002-3202101101112322-1121213312013012-2033010312102123-2030330103220221-3223121123231031-1203300131321221-0012023103032321)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3113321211103300-3121123312121102-2103230121222320-0023030100232313-0030331001233112-3103020012203031-1211333322302002-3323010310210312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002331120302120-0112031231330002-0203030303322100-1201130100231022-1120210331300302-3002201230230113-0323000012110200-0013111311311123"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route — direct_response_route / 103211001231 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route

<a id="canonical-0311012010021310-1120121100032011-2332303021212100-1321310112321013-3231210213202033-2332232312030133-3033032033123231-1132111131333120"></a>

Type: `"single"`. Computed.

Direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

Upstream description:

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

<a id="canonical-1221103231212003-3322120231120202-2302312321132111-2223310022320203-0023121211033133-2313023030303310-1330232313033031-2223121333213111"></a>

## Direct properties — direct_response_route / 103211001231 / 3

- [headers](data-sources--workload--reference--group-007.md#canonical-3213031321232203-0200213113100310-1332113110320200-0322013221032123-1312302112121331-1122331203010123-2333231121233310-3100113321320200): complete subsection reference.

<a id="canonical-2202003133120203-2302231311312030-2322102101031022-0103301213110223-3310002330210211-1200213311300231-3030230123331002-1322020020110031"></a>

<a id="canonical-0333300203203012-1222232222123101-1331132320122130-3203030300210231-2112120322121110-0022001100300113-0321300213132031-2123001122320012"></a>

## http_method property — direct_response_route / 103211001231 / 4

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

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

- [incoming_port](data-sources--workload--reference--group-008.md#canonical-1220223331102112-3311030110330320-3312022010122030-0323002023012210-2211120031010111-3202221103032312-3113313102103213-0311102310032123): complete subsection reference.

- [path](data-sources--workload--reference--group-008.md#canonical-3112303221220112-3220200323222331-3031032321232031-3231213031300132-0222021122010321-3010200221120330-1303033103030000-0212303111301332): complete subsection reference.

- [route_direct_response](data-sources--workload--reference--group-008.md#canonical-1233012303212313-3101123331213120-2210130313232300-3010333233222310-0331231001203323-1000212221033112-1323002011222232-1030300300323101): complete subsection reference.

<a id="canonical-0103300203310021-0332201103211221-2102130112200203-1020122101003021-0232223120311100-1023031130211020-2111021121223332-3111002123131000"></a>

## Next pages — direct_response_route / 103211001231 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers](data-sources--workload--reference--group-007.md#canonical-3213031321232203-0200213113100310-1332113110320200-0322013221032123-1312302112121331-1122331203010123-2333231121233310-3100113321320200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](data-sources--workload--reference--group-008.md#canonical-1220223331102112-3311030110330320-3312022010122030-0323002023012210-2211120031010111-3202221103032312-3113313102103213-0311102310032123)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path](data-sources--workload--reference--group-008.md#canonical-3112303221220112-3220200323222331-3031032321232031-3231213031300132-0222021122010321-3010200221120330-1303033103030000-0212303111301332)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response](data-sources--workload--reference--group-008.md#canonical-1233012303212313-3101123331213120-2210130313232300-3010333233222310-0331231001203323-1000212221033112-1323002011222232-1030300300323101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3213031321232203-0200213113100310-1332113110320200-0322013221032123-1312302112121331-1122331203010123-2333231121233310-3100113321320200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101001322231013-2331032101130103-1031210012321320-1122202112322011-0031111212002113-1202220012002030-0233122301313011-0003012013301322"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers — headers / 012220312021 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-007.md#canonical-3113321211103300-3121123312121102-2103230121222320-0023030100232313-0030331001233112-3103020012203031-1211333322302002-3323010310210312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers

<a id="canonical-2233312221220102-0212223222100300-3202032102211132-2331320222312022-2311332012133102-1323013233101122-1302230322213021-3223101322211130"></a>

Type: `"list"`. Computed.

Headers. List of (key, value) headers.

Upstream description:

List of (key, value) headers.

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

<a id="canonical-1202203112121232-0310301301313100-0321220202123233-3200002202223000-1211200220303221-2133320202013133-0221101231021003-2123213332122331"></a>

## Direct properties — headers / 012220312021 / 3

<a id="canonical-1121303313231201-3202210212211310-1001101200310232-2130001122002202-0100203300330033-1030133003230030-2322310312112203-1102300223230001"></a>

<a id="canonical-3000330133221330-2202122312121031-3133210222312233-1010021302121101-0223310222332002-0133032233201331-0220130322021110-0321213113231131"></a>

## exact property — headers / 012220312021 / 4

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

<a id="canonical-1213232323012010-2300023230331123-1122003222202003-2201101003033003-3231110102101203-2033332022222120-3010213333202330-3003022002113233"></a>

<a id="canonical-1301130111030110-2103201112020330-0023120311232321-3021120133220020-3311222121020110-3202313312212110-0321233020020011-1121100012222103"></a>

## invert_match property — headers / 012220312021 / 5

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

<a id="canonical-1100113112200212-3200310312030000-3112023011010303-0321312023321223-3203221101202311-3212223032021033-2001202312020330-0032331120110212"></a>
