---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1022320030333303-0030313102323203-2023213133322331-1001110102321313-1031220330111122-2020312112202202-0023230032031312-1300300133212303"></a>

## connection_idle_timeout property — https / 330302311213 / 6

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

- [default_header](data-sources--workload--reference--group-009.md#canonical-1310111300201130-3300312200323211-1201232033332103-3222013101021223-0201122120303221-1310323210223030-1233300300012011-1003022011133120): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-009.md#canonical-3001133332232322-1200321121321301-3200302113300201-3132003123223110-3211030323222301-0020231321013113-2123223023131023-3023220001323230): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-009.md#canonical-3202103311120012-2030332330132132-1320131200311223-1131012021212131-3031223333232020-0102201200031203-3031323220032133-2011203301101132): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-009.md#canonical-0302222110022333-1201233200130332-0000013212100201-1310010102120323-3312231232120302-3122110200301111-3212021331031120-0300033111013020): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-009.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021): complete subsection reference.

<a id="canonical-3301230313221100-3100023123331103-1103023321220021-2112000223212102-0323003122012032-2232033000003020-2012001331010220-3011032111231030"></a>

<a id="canonical-0331003313233211-2003023000302030-1012211211223333-2112301001320123-1002230200023223-0320220102131020-3121203333000333-0230021231310011"></a>

## http_redirect property — https / 330302311213 / 7

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

- [non_default_loadbalancer](data-sources--workload--reference--group-009.md#canonical-1013232211321311-1221201331033333-2201030202022021-1320003311233133-0220230311310331-2232322213312220-1302012300002130-3310213120300233): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-009.md#canonical-2233310030023222-0131300231233113-3030232101112200-1312321321013102-3112131313133123-1322203203132333-2210331221110011-0133101213221230): complete subsection reference.

<a id="canonical-1012330032033232-1102211002102200-3122020320123132-0213223111000133-2030113202211031-1311021323303132-3031310013310010-1220022011123223"></a>

<a id="canonical-0133210012103031-0121330302300010-3012201303121300-3202311212300112-2031131032013303-3321203102120133-2323130102232230-0030202102112011"></a>

## port property — https / 330302311213 / 8

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

<a id="canonical-2331002332120122-0022001122030011-1212130031113321-3112102133302213-0001330322033130-2322020310001302-3132112320301122-2211102200203332"></a>

<a id="canonical-3130020332002203-3002011012022110-2310132223001122-2132011303023310-3032331300112001-0010132230313312-3002210023333302-3213330313012333"></a>

## port_ranges property — https / 330302311213 / 9

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

<a id="canonical-3233122231200201-1131201123200132-2121302130312303-2321122321111003-3213012110000223-0231123310100323-0101023112312311-3310221233330221"></a>

<a id="canonical-3320231001123231-2132323102332101-2331130001133100-2233013230021232-0012001213033333-1321100002213320-1312203121001310-3113333201123100"></a>

## server_name property — https / 330302311213 / 10

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

- [tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003): complete subsection reference.

- [tls_parameters](data-sources--workload--reference--group-009.md#canonical-3030231203112023-2233233210333103-0332221021303202-1101310333001013-0301031200110232-3203210012020013-3333202011021213-3213221121020113): complete subsection reference.

<a id="canonical-1311321112331313-0313320130300121-1302231103301002-2101321033111130-2100303201032121-0031230121120110-3201000330310110-1132230100001333"></a>

## Next pages — https / 330302311213 / 11

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-009.md#canonical-1021033321200220-2122123222010113-3330010201103233-2103020220030330-2021200120310310-1302033032023022-1321103303100000-3212020230312013)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header](data-sources--workload--reference--group-009.md#canonical-1310111300201130-3300312200323211-1201232033332103-3222013101021223-0201122120303221-1310323210223030-1233300300012011-1003022011133120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer](data-sources--workload--reference--group-009.md#canonical-3001133332232322-1200321121321301-3200302113300201-3132003123223110-3211030323222301-0020231321013113-2123223023131023-3023220001323230)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize](data-sources--workload--reference--group-009.md#canonical-3202103311120012-2030332330132132-1320131200311223-1131012021212131-3031223333232020-0102201200031203-3031323220032133-2011203301101132)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize](data-sources--workload--reference--group-009.md#canonical-0302222110022333-1201233200130332-0000013212100201-1310010102120323-3312231232120302-3122110200301111-3212021331031120-0300033111013020)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-009.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer](data-sources--workload--reference--group-009.md#canonical-1013232211321311-1221201331033333-2201030202022021-1320003311233133-0220230311310331-2232322213312220-1302012300002130-3310213120300233)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through](data-sources--workload--reference--group-009.md#canonical-2233310030023222-0131300231233113-3030232101112200-1312321321013102-3112131313133123-1322203203132333-2210331221110011-0133101213221230)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-3030231203112023-2233233210333103-0332221021303202-1101310333001013-0301031200110232-3203210012020013-3333202011021213-3213221121020113)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1021033321200220-2122123222010113-3330010201103233-2103020220030330-2021200120310310-1302033032023022-1321103303100000-3212020230312013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000010001310223-1322322102203232-1031103232112201-2130210322003330-0100330331202131-3021213112303222-1211213300103123-3020233000102021"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options — coalescing_options / 023122003310 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-1012131322003202-1332022332030202-3032033100321312-2022321331133202-1322031203031232-3021203131313313-3230230033100312-2103133032133130"></a>

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

<a id="canonical-2210011023332031-1011102032011131-1021133233321230-1021022222113033-2322200111003302-3201100200032221-3003303301312201-2021312320300132"></a>

## Direct properties — coalescing_options / 023122003310 / 3

- [default_coalescing](data-sources--workload--reference--group-009.md#canonical-3200123202312121-2130211222213001-2030033011000112-3333201223200023-3010011103121031-0323231112232001-3020200220333112-3330331230311000): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-009.md#canonical-1122112303323203-3030011120030002-3312320202001030-3001223221203301-1222133322102300-2131130301122202-3102130201321233-1212112213300210): complete subsection reference.

<a id="canonical-1120202130123012-1032211213211230-2331311133233031-0133203112211223-2133313210000330-2320110300221311-2021231103111213-2021103331023221"></a>

## Next pages — coalescing_options / 023122003310 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing](data-sources--workload--reference--group-009.md#canonical-3200123202312121-2130211222213001-2030033011000112-3333201223200023-3010011103121031-0323231112232001-3020200220333112-3330331230311000)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing](data-sources--workload--reference--group-009.md#canonical-1122112303323203-3030011120030002-3312320202001030-3001223221203301-1222133322102300-2131130301122202-3102130201321233-1212112213300210)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3200123202312121-2130211222213001-2030033011000112-3333201223200023-3010011103121031-0323231112232001-3020200220333112-3330331230311000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111032003101133-3111133231203020-0011223121132200-3213223121130100-0213211113120333-3211331122011212-3130130112211031-2312011101130013"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing — default_coalescing / 312321220031 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-009.md#canonical-1021033321200220-2122123222010113-3330010201103233-2103020220030330-2021200120310310-1302033032023022-1321103303100000-3212020230312013)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-2202133203302123-2232332023103200-3201312013032213-2332313233002203-3003110213202201-2212033232011303-2011121221011222-1312112110120101"></a>

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

<a id="canonical-0013321321101110-3301132103313100-2100000023131222-3113030303012322-1023030130211313-3011212022223102-2120001031311010-0313023310001201"></a>

## Direct properties — default_coalescing / 312321220031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002223333133010-3112011031230211-1001131220310333-2030213113001203-2113012100310130-0020213313120320-2031113211020012-0300311021333110"></a>

## Next pages — default_coalescing / 312321220031 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-009.md#canonical-1021033321200220-2122123222010113-3330010201103233-2103020220030330-2021200120310310-1302033032023022-1321103303100000-3212020230312013)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1122112303323203-3030011120030002-3312320202001030-3001223221203301-1222133322102300-2131130301122202-3102130201321233-1212112213300210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223331312230203-2113321331121113-2012302003210022-2332031110201233-3330332300332200-1303120213021013-3031201312100223-2023013010111123"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing — strict_coalescing / 201001303110 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-009.md#canonical-1021033321200220-2122123222010113-3330010201103233-2103020220030330-2021200120310310-1302033032023022-1321103303100000-3212020230312013)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-2020113321312021-3221220311210321-0300123331302312-2230222021002011-2132132030102030-0200030220220032-1011223321332210-3110112300132233"></a>

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

<a id="canonical-2023112020012020-1133311021333201-3323022333022331-1031021102001200-2303023320103033-1123102310121332-0100232130231312-0312231101321023"></a>

## Direct properties — strict_coalescing / 201001303110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321321332022232-0201320302011301-2303321221213300-0112302122012022-3111303123131322-3303102323013003-2222112031200111-1210211033002211"></a>

## Next pages — strict_coalescing / 201001303110 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-009.md#canonical-1021033321200220-2122123222010113-3330010201103233-2103020220030330-2021200120310310-1302033032023022-1321103303100000-3212020230312013)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1310111300201130-3300312200323211-1201232033332103-3222013101021223-0201122120303221-1310323210223030-1233300300012011-1003022011133120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221321303313000-3332301120230213-3021031301311113-2321030232101010-2012211221202321-1120301020213011-2230211130003213-2022303100020122"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header — default_header / 310002021223 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header

<a id="canonical-1132220302121320-3111021222002031-2321102332220320-2010001213200220-1213003200023112-1320213321100200-3312120103303133-2220332320223300"></a>

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

<a id="canonical-3223003221112333-0033123220023000-3232112203220300-3113121303301020-1030233131113130-2023030133122212-1030333330011133-1111332020103231"></a>

## Direct properties — default_header / 310002021223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100032121333111-3032200231103000-0330301100012102-3023212231030313-2020010211213102-1331301003130332-1331111110303201-0302220032222210"></a>

## Next pages — default_header / 310002021223 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3001133332232322-1200321121321301-3200302113300201-3132003123223110-3211030323222301-0020231321013113-2123223023131023-3023220001323230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321212300102312-3302211201022113-1033011323313201-0311232122322312-0013322020322210-1300313002130011-1212101132111122-3110100300121123"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer — default_loadbalancer / 032121100200 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer

<a id="canonical-0301130101203120-0221230132223303-3322311200233331-3121123021303223-2312133003032032-3002101223132231-1133100031002302-1102013022023220"></a>

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

<a id="canonical-3301203111332123-3203331323121202-1301330332211011-0232101320211200-2101010032212223-2211122100000130-1223201333102101-0010133213223123"></a>

## Direct properties — default_loadbalancer / 032121100200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133303332212100-0333213033011203-0320131120030130-0320102020312010-0320332131033301-0323211301010031-0212101200200112-2120220230203312"></a>

## Next pages — default_loadbalancer / 032121100200 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3202103311120012-2030332330132132-1320131200311223-1131012021212131-3031223333232020-0102201200031203-3031323220032133-2011203301101132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201002323212321-0312301303023123-1011303211300232-3011200221120332-0211121331122300-1022000320203132-2001112001213211-3003330301333122"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize — disable_path_normalize / 103211003023 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize

<a id="canonical-3232020033023311-1012331213311032-3021111331130113-0233030310231121-0202310303232101-2333032111323201-0300301110333210-0132130330310121"></a>

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

<a id="canonical-0020021121321112-3103001011212032-2312211231312312-3110021100201301-0311033111333033-2030320233221303-2313223132033323-2022232310332313"></a>

## Direct properties — disable_path_normalize / 103211003023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212332230001000-1023221103202212-0013200122111133-1201300122121312-3102201000212130-2212231133331201-1323013300120220-3201013322020332"></a>

## Next pages — disable_path_normalize / 103211003023 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0302222110022333-1201233200130332-0000013212100201-1310010102120323-3312231232120302-3122110200301111-3212021331031120-0300033111013020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311302112100102-3001231122002311-2220131313133321-1230012320232012-2313202301023313-2312131220123311-0330002032101300-2323133010311202"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize — enable_path_normalize / 001131121332 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize

<a id="canonical-1001301111231101-3023223113222330-0031300001110000-2212103130323321-2012222002323003-1302112211021010-3301231133101223-0202101330323212"></a>

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

<a id="canonical-2001221301332220-3032101320333203-2321232030120213-1113011101130001-2310220200321100-1113033313333313-0233103333312100-3021320223333121"></a>

## Direct properties — enable_path_normalize / 001131121332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032010333131111-0131030233031230-0202330102200133-3313033332230212-1120313033121313-0032200123032112-0301321200123333-0301203133123213"></a>

## Next pages — enable_path_normalize / 001131121332 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012213301013320-2032202213010031-3000213221111101-3313200230030233-1322030300011011-0111312130310311-3310122210133102-2033123321321121"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options — http_protocol_options / 102231220000 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options

<a id="canonical-0310100112033202-3131032211103320-0331110320102020-1130112220323233-3230321000023222-0330223010323322-0300113230121110-0120220212010021"></a>

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

<a id="canonical-3331101133013200-3001313100223022-0233011113000131-3221320000101203-2132023033220001-1203033032230023-0212200010213200-0223322133223031"></a>

## Direct properties — http_protocol_options / 102231220000 / 3

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-009.md#canonical-3210323323021023-2301233332012302-1132311331201003-0212201012311032-1233312303310322-1122313230211311-2013300200000020-0033323210303031): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-009.md#canonical-3300333110331030-2300220013021003-3211133121101120-3012231311030202-0312102232102223-1020202303022201-0030301332102213-0022223231322122): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-009.md#canonical-1303310302200120-1022220302203012-3030003321113020-2000202021010232-0130020133221331-3210002111033101-1312210103332222-1013230332320232): complete subsection reference.

<a id="canonical-3331112131112323-1112331000231212-2032222232320330-3023211110213001-0302323333231112-3123122023132021-0332003030301023-2312231322323030"></a>

## Next pages — http_protocol_options / 102231220000 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-009.md#canonical-3210323323021023-2301233332012302-1132311331201003-0212201012311032-1233312303310322-1122313230211311-2013300200000020-0033323210303031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](data-sources--workload--reference--group-009.md#canonical-3300333110331030-2300220013021003-3211133121101120-3012231311030202-0312102232102223-1020202303022201-0030301332102213-0022223231322122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](data-sources--workload--reference--group-009.md#canonical-1303310302200120-1022220302203012-3030003321113020-2000202021010232-0130020133221331-3210002111033101-1312210103332222-1013230332320232)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3210323323021023-2301233332012302-1132311331201003-0212201012311032-1233312303310322-1122313230211311-2013300200000020-0033323210303031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313232100230310-3021313221103300-3002212332320312-3113011333020223-3300231202311213-3003013333113102-1030033032201331-1322033002130323"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 030001002031 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-009.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-1023210103211111-0303133201033330-1230332332121213-2313000022313111-3021131133012300-3120302322010033-3020232311013010-3300303312301313"></a>

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

<a id="canonical-1131101102033210-2033012332330202-0322021301113030-3330211320030110-3021212010311023-2000100233030202-3222320230102010-0221013000233123"></a>

## Direct properties — http_protocol_enable_v1_only / 030001002031 / 3

- [header_transformation](data-sources--workload--reference--group-009.md#canonical-3331320130210331-3111131030321131-0323013231233332-2201020112111222-3232121230120132-0030023001010010-1312233131133023-3112223001300231): complete subsection reference.

<a id="canonical-0031321330122232-0032323032200013-1020112131330132-3221331000220231-3022121123012302-0222001113313331-3112011032022320-1322333213321003"></a>

## Next pages — http_protocol_enable_v1_only / 030001002031 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-009.md#canonical-3331320130210331-3111131030321131-0323013231233332-2201020112111222-3232121230120132-0030023001010010-1312233131133023-3112223001300231)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-009.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3331320130210331-3111131030321131-0323013231233332-2201020112111222-3232121230120132-0030023001010010-1312233131133023-3112223001300231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002330100202131-3221300121032222-2232301313000130-3131002132202203-3302302203001110-3330322302313211-3013213303232301-2023231011213211"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 212320032131 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-009.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-009.md#canonical-3210323323021023-2301233332012302-1132311331201003-0212201012311032-1233312303310322-1122313230211311-2013300200000020-0033323210303031)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-2310310313213012-2302200132213231-3033102110133203-3313113301221020-0012001013113121-1311330133213331-1230313322321011-0031322331120013"></a>

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

<a id="canonical-1301330111320112-1311130323311320-3031131221003103-2210103102111221-1000033000303102-3200301101010110-2210011110333300-3031131222112322"></a>

## Direct properties — header_transformation / 212320032131 / 3

- [default_header_transformation](data-sources--workload--reference--group-009.md#canonical-1301032131112110-0233013120023201-0221310221301202-1200000223103010-0333103202123202-0011300212220201-2103012132221023-0002122110003011): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-009.md#canonical-0222320211203212-0331232131110201-0301310121003012-1110111331301321-3110232003131013-1013030123111002-1022121121302101-1313110320020300): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-009.md#canonical-3332301131202033-1122022103001200-2213212132310303-3001312000201000-2110032201201313-2332011200032131-0313131010133023-2331111131312003): complete subsection reference.

<a id="canonical-3121210322110013-3302322121223002-2323321031330100-1100133011131023-0222312131321121-2201212231033000-1311001301311200-1302103113232120"></a>

## Next pages — header_transformation / 212320032131 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--workload--reference--group-009.md#canonical-1301032131112110-0233013120023201-0221310221301202-1200000223103010-0333103202123202-0011300212220201-2103012132221023-0002122110003011)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--workload--reference--group-009.md#canonical-0222320211203212-0331232131110201-0301310121003012-1110111331301321-3110232003131013-1013030123111002-1022121121302101-1313110320020300)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--workload--reference--group-009.md#canonical-3332301131202033-1122022103001200-2213212132310303-3001312000201000-2110032201201313-2332011200032131-0313131010133023-2331111131312003)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-009.md#canonical-3210323323021023-2301233332012302-1132311331201003-0212201012311032-1233312303310322-1122313230211311-2013300200000020-0033323210303031)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1301032131112110-0233013120023201-0221310221301202-1200000223103010-0333103202123202-0011300212220201-2103012132221023-0002122110003011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322131102101211-1310201113221311-1232031312013200-2032233022012220-0011223332030311-2320311201010202-1222013132230202-0203303231201020"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 311130303330 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-009.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-009.md#canonical-3210323323021023-2301233332012302-1132311331201003-0212201012311032-1233312303310322-1122313230211311-2013300200000020-0033323210303031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-009.md#canonical-3331320130210331-3111131030321131-0323013231233332-2201020112111222-3232121230120132-0030023001010010-1312233131133023-3112223001300231)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-3120331213123131-1331220302222113-1333102222010032-2122133322101303-1113213233100103-0231031022220100-1101033131023230-2100113000002003"></a>

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

<a id="canonical-0212031000112030-0231310113210312-3313113231121233-3010220133212101-1023002321223232-1230310103323210-0033022330021213-0111102011202122"></a>

## Direct properties — default_header_transformation / 311130303330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302003223300311-3132333023132201-3113133302030211-1020000111111212-3202210303203113-3011201103030301-3300131003002312-3001303300033023"></a>

## Next pages — default_header_transformation / 311130303330 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-009.md#canonical-3331320130210331-3111131030321131-0323013231233332-2201020112111222-3232121230120132-0030023001010010-1312233131133023-3112223001300231)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0222320211203212-0331232131110201-0301310121003012-1110111331301321-3110232003131013-1013030123111002-1022121121302101-1313110320020300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003113100032210-0111233300220122-2032213311231111-2300212311102300-2310232100232221-0301001232012022-3231022233012120-0111130222232120"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 213033132231 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-009.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-009.md#canonical-3210323323021023-2301233332012302-1132311331201003-0212201012311032-1233312303310322-1122313230211311-2013300200000020-0033323210303031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-009.md#canonical-3331320130210331-3111131030321131-0323013231233332-2201020112111222-3232121230120132-0030023001010010-1312233131133023-3112223001300231)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-0323232222020030-0322023323012311-1120122133201122-2231330222001320-1313210213321020-1023302313013010-1212213321323313-2322321120211321"></a>

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

<a id="canonical-2100022312203030-3002110212311310-0232022103101011-1231230120322320-2002232200033131-0320212310102020-0311332303331030-3010030311011111"></a>

## Direct properties — preserve_case_header_transformation / 213033132231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113031233101111-3311200220331332-1300031031212021-1231110123112212-0333232313310113-3220232210210111-0000103003012011-1320013002011012"></a>

## Next pages — preserve_case_header_transformation / 213033132231 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-009.md#canonical-3331320130210331-3111131030321131-0323013231233332-2201020112111222-3232121230120132-0030023001010010-1312233131133023-3112223001300231)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3332301131202033-1122022103001200-2213212132310303-3001312000201000-2110032201201313-2332011200032131-0313131010133023-2331111131312003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201000111131220-0010002001211011-1001200101200220-0322223023033302-1211033210021312-0021123312103111-0302030011121030-3330023011232110"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 120301211010 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-009.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-009.md#canonical-3210323323021023-2301233332012302-1132311331201003-0212201012311032-1233312303310322-1122313230211311-2013300200000020-0033323210303031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-009.md#canonical-3331320130210331-3111131030321131-0323013231233332-2201020112111222-3232121230120132-0030023001010010-1312233131133023-3112223001300231)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-3013032300001223-2001233122313300-3002002233230002-0311201111113120-3110222300010321-0121123012003000-0133300033332231-1301001320101122"></a>

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

<a id="canonical-2301102323120112-0012331210100222-0213332230100103-0020100322221133-2010223313200020-3132221333101230-3132320123113201-0203220031012212"></a>

## Direct properties — proper_case_header_transformation / 120301211010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232301112103031-2312312002310313-1120323320232111-0131321033201210-1013121221031021-1110323210311330-1000211131303130-3101003200312110"></a>

## Next pages — proper_case_header_transformation / 120301211010 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-009.md#canonical-3331320130210331-3111131030321131-0323013231233332-2201020112111222-3232121230120132-0030023001010010-1312233131133023-3112223001300231)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3300333110331030-2300220013021003-3211133121101120-3012231311030202-0312102232102223-1020202303022201-0030301332102213-0022223231322122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323003313113232-2011110210203233-3311203000311013-1123113221000211-0312323100130231-0033030011311212-3322320320033122-3103102321121230"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 311100031232 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-009.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-0212333312221013-0132102332023213-0320122030020002-3011223203013231-3311123321220231-1310112122303023-0202010323211110-1311003322011011"></a>

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

<a id="canonical-2122000001120232-0300013011030203-0002013320213303-0003200012320302-1000031001133032-3100312100103311-3212312110311232-1213030122300112"></a>

## Direct properties — http_protocol_enable_v1_v2 / 311100031232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203123232121232-0021322102323333-3311332310333111-1122010321332030-0002003020031023-1031012001213202-0320201233110211-2213303130232133"></a>

## Next pages — http_protocol_enable_v1_v2 / 311100031232 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-009.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1303310302200120-1022220302203012-3030003321113020-2000202021010232-0130020133221331-3210002111033101-1312210103332222-1013230332320232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301100223312232-0110321233211023-0122012202203111-3230030103313023-3211011232203023-0032102103231133-3012302000102012-2221322202322011"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 121320133310 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-009.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-3302213132033132-1230203212313111-3011130100001223-3212023302003030-1102013333330330-0310033013233003-1012202102131003-0102230330320313"></a>

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

<a id="canonical-1133310320131113-0303121212221323-3330202130222030-0211203113330301-2302003012313213-3110102011220033-3200221233301211-2313221032030131"></a>

## Direct properties — http_protocol_enable_v2_only / 121320133310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030230002321222-3113010012003321-0102331102020322-0031130211101222-2022110213022300-0231031323331002-3133301213030222-2331102231011233"></a>

## Next pages — http_protocol_enable_v2_only / 121320133310 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-009.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1013232211321311-1221201331033333-2201030202022021-1320003311233133-0220230311310331-2232322213312220-1302012300002130-3310213120300233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021322321211331-3300322233202013-0021210301033011-1320131330301320-3100230003333220-0012230020211310-2103223120213232-0322122233103201"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer — non_default_loadbalancer / 200113111020 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-1102310200112100-0202332030132000-1021113222323210-2033033233012313-3230011031200231-3030322102202132-2102221222121310-0311101031021001"></a>

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

<a id="canonical-0303121030112321-2323113102001323-2032000013120021-3123022032200233-3233031311123210-0032310231213133-1033023211031231-0303233101020032"></a>

## Direct properties — non_default_loadbalancer / 200113111020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130320112213320-3030221122222312-2101223111300311-1102031003200223-1312111232130302-2210130113013021-2002212323233110-1323223031221331"></a>

## Next pages — non_default_loadbalancer / 200113111020 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2233310030023222-0131300231233113-3030232101112200-1312321321013102-3112131313133123-1322203203132333-2210331221110011-0133101213221230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330111020123032-2222200002121230-0123130330123002-2313232220031101-1100300003003313-2103130320320213-3122003112213022-3322302200312331"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through — pass_through / 110010203013 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through

<a id="canonical-1231221103002033-2033222023323200-0132111123223120-3022331322112011-3001320121101202-0112102311213211-2322002330131203-3033132030111130"></a>

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

<a id="canonical-1211210031322300-0102322211322311-1311112030010332-2120021023222120-1213012331212100-2133310013302000-0211330302232120-0212302310202111"></a>

## Direct properties — pass_through / 110010203013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102120013213022-0321230321132003-3001101033102200-3321330113031030-2022232102111011-0200233013102113-0303221201310310-0200133012200312"></a>

## Next pages — pass_through / 110010203013 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302113121210021-3220233310112002-0002122002113332-0311313100322110-3111030132020013-1213010330013333-1131312332323210-3333331311322212"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params — tls_cert_params / 111011123103 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params

<a id="canonical-0312203210210030-0322310111113210-0003023222322332-0231020013210001-3233031022311013-2221122003001202-3213033213301312-2020010312300231"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

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

<a id="canonical-2302021030211312-0203110123012211-2031130312103001-3321113022223112-2233103021300011-3201232333221022-1331300202311011-2122013202132131"></a>

## Direct properties — tls_cert_params / 111011123103 / 3

- [certificates](data-sources--workload--reference--group-009.md#canonical-1301033320033002-0103021100223102-2300222133210001-1203021323012013-2232300231322121-2130113323330213-2023123331110021-2302003003033000): complete subsection reference.

- [no_mtls](data-sources--workload--reference--group-009.md#canonical-3022220032303311-0302131212132302-1033232221302310-0132112112132131-3223130103202323-0000313011013112-2230230103220322-2020200300201311): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-009.md#canonical-1133031130201122-0201003200011212-2221030021203322-2133332333333002-3302131022132331-2303132230320130-2301120223312101-3103131200232222): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-009.md#canonical-0202033310131322-1320211131131020-0122011020002120-1022330311232030-2322222101011030-1110301022313120-1300002301113101-2111020021320003): complete subsection reference.

<a id="canonical-2312101211211233-1330030010100332-3120120212121213-0023033210301110-3301231011211010-1031222123013000-2102321210032202-1132010200200202"></a>

## Next pages — tls_cert_params / 111011123103 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates](data-sources--workload--reference--group-009.md#canonical-1301033320033002-0103021100223102-2300222133210001-1203021323012013-2232300231322121-2130113323330213-2023123331110021-2302003003033000)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls](data-sources--workload--reference--group-009.md#canonical-3022220032303311-0302131212132302-1033232221302310-0132112112132131-3223130103202323-0000313011013112-2230230103220322-2020200300201311)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-009.md#canonical-1133031130201122-0201003200011212-2221030021203322-2133332333333002-3302131022132331-2303132230320130-2301120223312101-3103131200232222)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-009.md#canonical-0202033310131322-1320211131131020-0122011020002120-1022330311232030-2322222101011030-1110301022313120-1300002301113101-2111020021320003)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1301033320033002-0103021100223102-2300222133210001-1203021323012013-2232300231322121-2130113323330213-2023123331110021-2302003003033000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003212122131211-2321213012010012-1212321220130221-2102322311313231-0100333200012121-0121221210201311-3120011121211311-2231301022200300"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates — certificates / 032213031020 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-3333231323323220-2302031001000323-0012200030311133-2000202201021031-3321333232301220-1022303233211311-1103000121022111-3110103221033232"></a>

Type: `"list"`. Computed.

Select one or more certificates with any domain names.

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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3231111302212122-2112312310330230-3211133001103132-3313331300032010-0120111213033132-0011300030021112-3210321322033133-0300221132132112"></a>

## Direct properties — certificates / 032213031020 / 3

<a id="canonical-2320202031203110-0010011012330201-0212303000312322-1331011021003023-2112322201100001-2223122203200130-3201033032220013-0312221202102001"></a>

<a id="canonical-3210001231211021-3321111330233002-3122333330100110-3102100133123023-3103000310031310-0100200300111101-2001100121302001-1222220122220312"></a>

## name property — certificates / 032213031020 / 4

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

<a id="canonical-0210310130210120-1303310230010000-3030330013303223-3323320232322101-1103001123012300-3301330132001020-0010323233110022-1113000103320213"></a>

<a id="canonical-2203200203331220-3221110321013322-1030031101302303-1201131311223123-0231013210121112-3120321133232211-3303302013320332-1103332022332312"></a>

## namespace property — certificates / 032213031020 / 5

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

<a id="canonical-0101213120131110-3021311331310212-0032000022211110-0310233320101331-2202132032030113-0233200213102011-3331321110211103-3323021012321030"></a>

<a id="canonical-0303211220233133-2300132233300311-1213000121030201-1112213330122033-1113230132213301-1033220300122212-3222132300002012-2331003120301033"></a>

## tenant property — certificates / 032213031020 / 6

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

<a id="canonical-2302101310231312-2131222002133312-2321111213113303-2121203320132112-2220030310313101-2303201200103322-1112320323110231-0123121123302133"></a>

## Next pages — certificates / 032213031020 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3022220032303311-0302131212132302-1033232221302310-0132112112132131-3223130103202323-0000313011013112-2230230103220322-2020200300201311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233313113231210-1221322011022103-1122001130110031-2100132201011233-0313213322110022-1030032013231133-1202110120012011-0021133222113111"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls — no_mtls / 130202333023 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-0102112331131003-0112323300103231-1232221231200230-3001320233330121-1030222320120303-0013112123103333-2002102001020002-1033010301220032"></a>

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

<a id="canonical-1201010303123321-2220310322221012-1320312131230023-2001311300103323-3013120121201213-3003013012101323-1311002220320031-1221331321132230"></a>

## Direct properties — no_mtls / 130202333023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331212123010130-0200103212123213-2210003132012110-1222031302313221-3303030011233211-0010002333313200-0232210020011103-0213102200132102"></a>

## Next pages — no_mtls / 130202333023 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1133031130201122-0201003200011212-2221030021203322-2133332333333002-3302131022132331-2303132230320130-2301120223312101-3103131200232222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212202321113113-1201002100300213-1023233033113023-3023301210002023-2212110210033010-3303021122211231-1113022332303023-2223231110023120"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config — tls_config / 233321212003 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-3011202011020303-2122113323323030-3013332030310033-0132011303323303-3003333102013331-2111200031310022-1103033131332001-2223210330023300"></a>

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

<a id="canonical-1222023021231022-0300020032100310-0111112310002022-0212003113223030-3223110200001330-0021311312331131-3113133332201020-3301103003213211"></a>

## Direct properties — tls_config / 233321212003 / 3

- [custom_security](data-sources--workload--reference--group-009.md#canonical-3101302131121110-2013010220213123-2231300332033021-1120311321332213-3001132000030001-1122111230030323-0323113332020011-1221320110221032): complete subsection reference.

- [default_security](data-sources--workload--reference--group-009.md#canonical-3330103230033201-3202102322300020-1321010222201301-1201233212031103-2211301200303031-3000111122322101-2223202021333202-1113011110110030): complete subsection reference.

- [low_security](data-sources--workload--reference--group-009.md#canonical-3311121211221310-2312002033213230-3223110321131132-2210231122122033-1120033031013011-0311011101000231-1101232303212111-1000001310013001): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-009.md#canonical-1101313220012023-2232012323321112-3230222131313112-1233002131122122-2301230333200331-3313210121130131-0233133303022112-1310312003233212): complete subsection reference.

<a id="canonical-2221113010121131-2230323320003301-2312300222110212-1101231132313211-0010320300201233-3021013030303212-0311231201310100-1013322011311231"></a>

## Next pages — tls_config / 233321212003 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security](data-sources--workload--reference--group-009.md#canonical-3101302131121110-2013010220213123-2231300332033021-1120311321332213-3001132000030001-1122111230030323-0323113332020011-1221320110221032)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security](data-sources--workload--reference--group-009.md#canonical-3330103230033201-3202102322300020-1321010222201301-1201233212031103-2211301200303031-3000111122322101-2223202021333202-1113011110110030)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security](data-sources--workload--reference--group-009.md#canonical-3311121211221310-2312002033213230-3223110321131132-2210231122122033-1120033031013011-0311011101000231-1101232303212111-1000001310013001)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security](data-sources--workload--reference--group-009.md#canonical-1101313220012023-2232012323321112-3230222131313112-1233002131122122-2301230333200331-3313210121130131-0233133303022112-1310312003233212)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3101302131121110-2013010220213123-2231300332033021-1120311321332213-3001132000030001-1122111230030323-0323113332020011-1221320110221032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222132210210330-1311201300022003-3311230032032003-0113100112113122-0212223233200010-2120312202223002-2311002012101232-0222002113032111"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security — custom_security / 331200113101 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-009.md#canonical-1133031130201122-0201003200011212-2221030021203322-2133332333333002-3302131022132331-2303132230320130-2301120223312101-3103131200232222)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-2123233021033332-2301223310213001-3000112323130211-3002030213120130-3020122023332302-2232130113233020-0021313222202301-2332230011311203"></a>

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

<a id="canonical-0102313113030230-2022201100113302-0112121222021111-1100133022203123-2303032302032330-1231220131312120-2310003320223003-1011123031311121"></a>

## Direct properties — custom_security / 331200113101 / 3

<a id="canonical-1011113103103010-3333103222001330-0102201201131220-1102232301010213-1312133222311323-1011233022330331-0012330011020221-0033023313002133"></a>

<a id="canonical-0200030311001300-1301202313010311-0223322122023212-3031310033103000-2330323033003020-0223030211210320-2321132310233123-1303001202210101"></a>

## cipher_suites property — custom_security / 331200113101 / 4

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

<a id="canonical-0331332101020003-2112030122201210-2313120112301131-3022222000211300-3112030030131002-1010220323120320-2232021301331030-0332231320220312"></a>

<a id="canonical-0221033230210233-3201110021033232-0230002200011011-0233201023010103-0232100033131013-2000331021230120-3331033313333013-1330021301130320"></a>

## max_version property — custom_security / 331200113101 / 5

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

<a id="canonical-2303233122230323-2102100221011131-3323000123213022-3122031232222330-2030323312231211-0021233033033120-2130221220332233-2020301303130212"></a>

<a id="canonical-2110001100221311-1222012133133221-1113331123022300-1333210223000200-1103231300311021-1002320133233330-2101213230112003-1023031133033102"></a>

## min_version property — custom_security / 331200113101 / 6

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

<a id="canonical-2030223031301300-2010311111012103-1110223023231330-0322322313103022-0130000100213032-2202233112303323-3021233120123131-0100021320321203"></a>

## Next pages — custom_security / 331200113101 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-009.md#canonical-1133031130201122-0201003200011212-2221030021203322-2133332333333002-3302131022132331-2303132230320130-2301120223312101-3103131200232222)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3330103230033201-3202102322300020-1321010222201301-1201233212031103-2211301200303031-3000111122322101-2223202021333202-1113011110110030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213100112012221-2003131302233213-1200001213030003-3321210030111121-1230102123230330-3032132100321333-3333310302021131-2132332231322230"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security — default_security / 122303101221 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-009.md#canonical-1133031130201122-0201003200011212-2221030021203322-2133332333333002-3302131022132331-2303132230320130-2301120223312101-3103131200232222)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-3311013021012212-3010312023311311-3201133330023320-1101321010021211-3322133123132213-3223113013111021-0123303123102303-1320313332201101"></a>

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

<a id="canonical-2030230113122320-0021222112200121-2133333013202001-0312303233032200-3332010111002333-2321331313103022-2300103112323210-1203113200113120"></a>

## Direct properties — default_security / 122303101221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123101033212220-2310130000111031-2322031103103032-0112121122331321-1032012123232013-0210210323232301-0203123113002113-3002010001332031"></a>

## Next pages — default_security / 122303101221 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-009.md#canonical-1133031130201122-0201003200011212-2221030021203322-2133332333333002-3302131022132331-2303132230320130-2301120223312101-3103131200232222)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3311121211221310-2312002033213230-3223110321131132-2210231122122033-1120033031013011-0311011101000231-1101232303212111-1000001310013001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002010121030221-1203113323122011-2330230033030312-3323031320120313-2103110332133011-0120001213112301-1210332030220122-1112210200013302"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security — low_security / 013203130121 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-009.md#canonical-1133031130201122-0201003200011212-2221030021203322-2133332333333002-3302131022132331-2303132230320130-2301120223312101-3103131200232222)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-3331032030131312-2100300313200001-0203232230030032-0330130131222312-0111210130100102-1000121210223102-0211032022100320-3102322002333100"></a>

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

<a id="canonical-2112031013212013-0221321330200300-2022200201310131-0023000010201111-1211333221131003-3303321201101002-1312121222100130-2120030101032211"></a>

## Direct properties — low_security / 013203130121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022120010013101-3010130010003001-2132131101113010-3211112022112231-1331030010131312-1012322010033123-0003123012022021-0122120133022321"></a>

## Next pages — low_security / 013203130121 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-009.md#canonical-1133031130201122-0201003200011212-2221030021203322-2133332333333002-3302131022132331-2303132230320130-2301120223312101-3103131200232222)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1101313220012023-2232012323321112-3230222131313112-1233002131122122-2301230333200331-3313210121130131-0233133303022112-1310312003233212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102330221012202-2330022033323112-3110100301003112-2111313330232031-0031323113223012-3022000111200202-2012223323310023-0131313023312112"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security — medium_security / 231100320212 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-009.md#canonical-1133031130201122-0201003200011212-2221030021203322-2133332333333002-3302131022132331-2303132230320130-2301120223312101-3103131200232222)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-2110310010311112-1100122213132100-1213211122103312-1132123311221212-2221031200222330-1203131313133020-0232211012030201-3120130000303133"></a>

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

<a id="canonical-3201330322312331-3222303100131223-1322213301323313-0023131001031103-1333212000220111-1022221110323102-1222031123201100-3221210221002232"></a>

## Direct properties — medium_security / 231100320212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112100101132012-0013003122133313-0223200231132010-1002033113003320-3130133333132201-0230332100112032-1230300333211210-2202301033322123"></a>

## Next pages — medium_security / 231100320212 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-009.md#canonical-1133031130201122-0201003200011212-2221030021203322-2133332333333002-3302131022132331-2303132230320130-2301120223312101-3103131200232222)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0202033310131322-1320211131131020-0122011020002120-1022330311232030-2322222101011030-1110301022313120-1300002301113101-2111020021320003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233021001210301-1313322300222011-0202111231203011-1303131131210121-1212331322333231-2311133210210030-1023310033232020-3220310300102131"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls — use_mtls / 201310020303 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-2010312330212032-1321031211003001-1203321330022022-1120303132330222-3203231332330300-0203200032103221-0200021233301100-0133331210111311"></a>

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

<a id="canonical-0133021000211010-3200320220010200-3333310033231213-0103333232120221-0230213321001202-0203012203130230-0221021032111023-2130110321321230"></a>

## Direct properties — use_mtls / 201310020303 / 3

<a id="canonical-1103020303110200-2300002122000123-0012303023121113-0202003221122202-0003323223013233-0200030002001233-2231322222313200-1121230321132000"></a>

<a id="canonical-2313200010101333-3031321300111120-2330022100300022-3300222100323322-1011101213203321-0130331221222101-1030001003021230-1323310131222012"></a>

## client_certificate_optional property — use_mtls / 201310020303 / 4

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

- [crl](data-sources--workload--reference--group-009.md#canonical-2222033212130122-1001021102002121-3211312033211330-3131112203123203-3231132102101020-2000012021132331-2302320123103312-3203110121212220): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-009.md#canonical-2301231110110112-1111001101300301-2230333110011302-0001000303112222-1212022120003211-1123102310213031-0312223033232100-2121133133231033): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-009.md#canonical-1232323020201333-0333321220001321-1123223300000303-1332112333002313-1122013030001311-1033102220300230-2020133202112331-0112332013013022): complete subsection reference.

<a id="canonical-0102133130322122-3032123231202322-0222121121331012-2020122131212332-0203121130132112-2113033101320113-0130112211220103-2220000231023121"></a>

<a id="canonical-3311303200230133-2232021233301311-3223111002030220-2201322123000102-0310102122020320-0323003000221102-1100222110011300-1300012011230120"></a>

## trusted_ca_url property — use_mtls / 201310020303 / 5

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

- [xfcc_disabled](data-sources--workload--reference--group-009.md#canonical-1221332103023213-2210121102102321-1000330013211131-2022100133032133-1120130233320130-2122202132220112-0303333110033023-2201023022213311): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-009.md#canonical-1000013110310033-0010112331030310-0313201011122001-1113011011003033-3113120002322303-1302311302132010-1220213333113203-3010110303100013): complete subsection reference.

<a id="canonical-3101013303213232-0011003211212021-3222112202333021-2230221333333130-1113103231121330-3010212132230113-2313322330103020-1130201333031211"></a>

## Next pages — use_mtls / 201310020303 / 6

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl](data-sources--workload--reference--group-009.md#canonical-2222033212130122-1001021102002121-3211312033211330-3131112203123203-3231132102101020-2000012021132331-2302320123103312-3203110121212220)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl](data-sources--workload--reference--group-009.md#canonical-2301231110110112-1111001101300301-2230333110011302-0001000303112222-1212022120003211-1123102310213031-0312223033232100-2121133133231033)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca](data-sources--workload--reference--group-009.md#canonical-1232323020201333-0333321220001321-1123223300000303-1332112333002313-1122013030001311-1033102220300230-2020133202112331-0112332013013022)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled](data-sources--workload--reference--group-009.md#canonical-1221332103023213-2210121102102321-1000330013211131-2022100133032133-1120130233320130-2122202132220112-0303333110033023-2201023022213311)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options](data-sources--workload--reference--group-009.md#canonical-1000013110310033-0010112331030310-0313201011122001-1113011011003033-3113120002322303-1302311302132010-1220213333113203-3010110303100013)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2222033212130122-1001021102002121-3211312033211330-3131112203123203-3231132102101020-2000012021132331-2302320123103312-3203110121212220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223130032101001-1332013133212110-0230013123032023-1002031033331210-0030300131113102-0322032020133222-0200102002223130-1112210321201130"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl — crl / 111332013013 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-009.md#canonical-0202033310131322-1320211131131020-0122011020002120-1022330311232030-2322222101011030-1110301022313120-1300002301113101-2111020021320003)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-3131021223303033-1312222030100201-0010211121121132-3030331012320121-0011131210031002-2302101110330211-0311200210211210-2111000211012033"></a>

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

<a id="canonical-0111022220030302-0032312023333303-0020303113113222-2002032311113221-0300120323333033-1333012301310103-0333000230011032-0221323200230200"></a>

## Direct properties — crl / 111332013013 / 3

<a id="canonical-0010203312213112-3002110203030132-2303222221112000-2012330313313313-0203111113212302-0032002000231032-0012100313031123-2013223230302233"></a>

<a id="canonical-1013033001023210-2230032222123303-1032031203303002-1031320110303020-1303201003111211-2112220213303112-2101112200131303-2301131200133132"></a>

## name property — crl / 111332013013 / 4

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

<a id="canonical-0111011320332020-3332333232333011-1222132303110200-2033331132121122-1030123003200033-3310202302221321-3020103003013331-3203032103001212"></a>

<a id="canonical-3311120000031031-2133012112310200-3311211002213310-3032221213232031-1210003022220323-3132020200212122-2103331100002332-2121220300102113"></a>

## namespace property — crl / 111332013013 / 5

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

<a id="canonical-1002320123320012-1030020110013213-0112313330212331-2023230211132120-0222030200033101-1200103100120032-0331302101230312-1322011200230103"></a>

<a id="canonical-3121333210233230-0323001130213033-0201011121331121-1211110033203310-2120321021222330-3321333121120122-3203022210233232-1200203210011111"></a>

## tenant property — crl / 111332013013 / 6

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

<a id="canonical-0013122312130230-2233022111221301-2020213123211120-0123120021110201-1130102323220202-0311222222110230-0131220132002202-2320130322031311"></a>

## Next pages — crl / 111332013013 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-009.md#canonical-0202033310131322-1320211131131020-0122011020002120-1022330311232030-2322222101011030-1110301022313120-1300002301113101-2111020021320003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2301231110110112-1111001101300301-2230333110011302-0001000303112222-1212022120003211-1123102310213031-0312223033232100-2121133133231033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023000313110111-1023113322001230-3230331103302012-2200020320122212-0113323100003331-3132000221012010-3311300203113102-2310121010312011"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl — no_crl / 231232131023 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-009.md#canonical-0202033310131322-1320211131131020-0122011020002120-1022330311232030-2322222101011030-1110301022313120-1300002301113101-2111020021320003)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-1032301201320311-2110130233010331-3103002131120001-3221001312300121-1003223200131003-3200111331000101-1231003311203021-1210331021032230"></a>

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

<a id="canonical-0123101232021001-2120301133210203-0212311112313121-0003003323132233-0010003222233231-1123130031021303-2030022330212030-2121333123331101"></a>

## Direct properties — no_crl / 231232131023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311300221131030-0112021202132213-3033331002120200-0121232033123232-0111231201123020-2112022033010230-2312022101113201-0200131333300323"></a>

## Next pages — no_crl / 231232131023 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-009.md#canonical-0202033310131322-1320211131131020-0122011020002120-1022330311232030-2322222101011030-1110301022313120-1300002301113101-2111020021320003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1232323020201333-0333321220001321-1123223300000303-1332112333002313-1122013030001311-1033102220300230-2020133202112331-0112332013013022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030221202131330-0321221130112112-0312123333033230-1303222201233101-2311013120031203-0212101303031021-1122101203013002-3330213120311311"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca — trusted_ca / 331130232033 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-009.md#canonical-0202033310131322-1320211131131020-0122011020002120-1022330311232030-2322222101011030-1110301022313120-1300002301113101-2111020021320003)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-0033301221332211-3222112000031103-1012011020022013-2021202121230302-1122112301302312-2131112110003100-3220210002101222-3121232202200032"></a>

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

<a id="canonical-1110022203311231-1133303131013313-1013323233012303-2001132321303322-0132312112323002-0101213132030300-1323101133110323-0312320233230111"></a>

## Direct properties — trusted_ca / 331130232033 / 3

<a id="canonical-3002221213203103-1032011013111000-0130310132212133-2203201320222112-0212131202231011-0001020331220231-0221031021130130-3220010333031112"></a>

<a id="canonical-0002113120010123-0132311222200131-3230120101001103-1032221130210330-2123112233000323-0320323122223002-3313110303020310-1301213023313310"></a>

## name property — trusted_ca / 331130232033 / 4

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

<a id="canonical-2030031110101202-1312010021213100-0123303302102011-2223203023113023-3122102213331323-0032322023220011-0121213233321020-3102332103013232"></a>

<a id="canonical-3022322333102330-0031121112110122-1111122121312123-2011000102102311-0112002323010001-2032302023320202-2230300031322011-3220231133332102"></a>

## namespace property — trusted_ca / 331130232033 / 5

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

<a id="canonical-3223012030322231-1322331221230030-0332302013221212-2310203303000100-3113330203103030-1312000133022021-2110001201322211-0211113231020222"></a>

<a id="canonical-1301102220330002-0221022033331200-1222332310120331-0002003020212302-3000200223122301-1111230301323102-3132310102000311-0130122032130113"></a>

## tenant property — trusted_ca / 331130232033 / 6

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

<a id="canonical-1313301312300211-1220012001302022-0031230023223002-3220100203211222-0223202001331100-3233323201130102-2232312023021301-0231032101032002"></a>

## Next pages — trusted_ca / 331130232033 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-009.md#canonical-0202033310131322-1320211131131020-0122011020002120-1022330311232030-2322222101011030-1110301022313120-1300002301113101-2111020021320003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1221332103023213-2210121102102321-1000330013211131-2022100133032133-1120130233320130-2122202132220112-0303333110033023-2201023022213311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121001033213331-1200110032131201-0131021313222123-1210003212101333-3312123330023313-0122003030320123-3212020033000022-2103023301320322"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled — xfcc_disabled / 021203201122 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-009.md#canonical-0202033310131322-1320211131131020-0122011020002120-1022330311232030-2322222101011030-1110301022313120-1300002301113101-2111020021320003)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-0200101223010202-2301201213001012-3020020333221030-3102323130030102-0000203031320213-1000313220231303-1221232320011331-3230120022010223"></a>

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

<a id="canonical-3101301003213210-2303121122130331-1322033112311223-1132302221102310-3001231323221011-2012220223201203-3103000323110213-2002001320033212"></a>

## Direct properties — xfcc_disabled / 021203201122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032223012012130-0002011201231100-1302110230010123-1311211232011320-2302331110103232-2031132213130101-3132023322023223-1132230130130133"></a>

## Next pages — xfcc_disabled / 021203201122 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-009.md#canonical-0202033310131322-1320211131131020-0122011020002120-1022330311232030-2322222101011030-1110301022313120-1300002301113101-2111020021320003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1000013110310033-0010112331030310-0313201011122001-1113011011003033-3113120002322303-1302311302132010-1220213333113203-3010110303100013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111331033223010-3023121331010022-0023321223103130-0122213012012332-3010032131213320-1133033111130213-2330030102223233-2210101320211310"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options — xfcc_options / 302030122230 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-009.md#canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-009.md#canonical-0202033310131322-1320211131131020-0122011020002120-1022330311232030-2322222101011030-1110301022313120-1300002301113101-2111020021320003)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-0221213121321010-0111301212111200-0332032000010133-1313302223011002-1332311131321221-2120313112103301-0033110012001010-1200300303133020"></a>

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

<a id="canonical-3012230331121302-1313230021001100-2012003202133130-1100111202323002-3022102103213032-0210202213101003-2003210012301310-2203313012102202"></a>

## Direct properties — xfcc_options / 302030122230 / 3

<a id="canonical-2011122110000132-0210000001030231-3023210300221020-2321331020300021-1332221200103131-0330011001003301-0020320220111301-0300003130222002"></a>

<a id="canonical-1210333110101201-3211102302300221-1203213123203320-0213331202212230-2130222033131023-1012022113321013-3212012002231001-2313112232200001"></a>

## xfcc_header_elements property — xfcc_options / 302030122230 / 4

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

<a id="canonical-2110032330021120-0202033113233302-0012333121232301-3230023231222010-3231333330222133-2331210032102310-3201300211301230-3201201033030310"></a>

## Next pages — xfcc_options / 302030122230 / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-009.md#canonical-0202033310131322-1320211131131020-0122011020002120-1022330311232030-2322222101011030-1110301022313120-1300002301113101-2111020021320003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3030231203112023-2233233210333103-0332221021303202-1101310333001013-0301031200110232-3203210012020013-3333202011021213-3213221121020113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310332032133113-2203323200300010-0010233312111112-0103310010113312-3312103031021111-0002331010010003-3123322031031101-1202031123103023"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters — tls_parameters / 103013003333 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters

<a id="canonical-2100302312223011-3122232001231021-1202202100223303-0132220002300012-2103113230113231-0113111013303132-3121030213202003-2203210113322111"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

Upstream description:

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

<a id="canonical-3322131221100303-2023232201012031-1211221011301020-3031203213321022-2110013011131003-2300320010132130-0211101212213130-2312221031131332"></a>

## Direct properties — tls_parameters / 103013003333 / 3

- [no_mtls](data-sources--workload--reference--group-009.md#canonical-1022220202023211-2112103231203311-3211313330213002-0102131100230102-2212031023301210-0012233320113301-0211330101222230-0130200313032110): complete subsection reference.

- [tls_certificates](data-sources--workload--reference--group-009.md#canonical-0033313103322112-0203330310133201-1121203002123230-1312213033300032-2111110320220201-1011112012220232-2020022322133332-0013130111131230): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-010.md#canonical-2021120001321031-2332013010231300-0320101103110022-3002023131032222-2331001221230213-3231132113321331-1002110103130201-1313302223310002): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-010.md#canonical-1221213020103013-1101032121313122-3122223200121312-3212031302101022-2110022231001123-2002020021110131-0300022001201321-2030210010213223): complete subsection reference.

<a id="canonical-3111102213111032-3131130122201101-3313031320200132-1111120320323301-2230101333111111-2113113010313303-2011112331100232-3103030331320120"></a>

## Next pages — tls_parameters / 103013003333 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls](data-sources--workload--reference--group-009.md#canonical-1022220202023211-2112103231203311-3211313330213002-0102131100230102-2212031023301210-0012233320113301-0211330101222230-0130200313032110)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-009.md#canonical-0033313103322112-0203330310133201-1121203002123230-1312213033300032-2111110320220201-1011112012220232-2020022322133332-0013130111131230)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-010.md#canonical-2021120001321031-2332013010231300-0320101103110022-3002023131032222-2331001221230213-3231132113321331-1002110103130201-1313302223310002)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-010.md#canonical-1221213020103013-1101032121313122-3122223200121312-3212031302101022-2110022231001123-2002020021110131-0300022001201321-2030210010213223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1022220202023211-2112103231203311-3211313330213002-0102131100230102-2212031023301210-0012233320113301-0211330101222230-0130200313032110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100122233111131-0003202331033010-2032333321131032-1010111020203201-2110130300120233-0013111001312223-1031200012221031-0311111302010023"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls — no_mtls / 112100210301 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-3030231203112023-2233233210333103-0332221021303202-1101310333001013-0301031200110232-3203210012020013-3333202011021213-3213221121020113)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-3330330312101113-2002112223203011-2322331000320323-2110213310211311-2313230231021211-0103031113301132-2113111221312230-1111132111001300"></a>

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

<a id="canonical-1321322001201120-0302113230302200-2003031021013331-2031202000131003-0222000010023232-1110203000020110-3323211001123320-0301312032021123"></a>

## Direct properties — no_mtls / 112100210301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303102130313130-2210220300023330-0113320112202212-2332313233021210-1330030101033131-1111311020303201-3202120333200222-2110230202013033"></a>

## Next pages — no_mtls / 112100210301 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-3030231203112023-2233233210333103-0332221021303202-1101310333001013-0301031200110232-3203210012020013-3333202011021213-3213221121020113)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0033313103322112-0203330310133201-1121203002123230-1312213033300032-2111110320220201-1011112012220232-2020022322133332-0013130111131230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030022103100031-1223130233010000-3100022302311310-1110133212122020-2123230202323010-3201113112020132-1221332222101231-0012122311212323"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates — tls_certificates / 032212232332 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-009.md#canonical-3030231203112023-2233233210333103-0332221021303202-1101310333001013-0301031200110232-3203210012020013-3333202011021213-3213221121020113)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-0111123132223230-1311312121333130-2211323321123002-1233002201201300-3313201332131320-1033310320223031-1330111130113002-3320323002320101"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1221102102233013-3100131120300232-2201311011021003-2201211203003131-2202213130021112-1312311320213032-2333112123213000-0320323221211112"></a>

## Direct properties — tls_certificates / 032212232332 / 3

<a id="canonical-1321103021330322-0311023010021331-0331202023012312-2233200120030013-2001012211113223-2121301103213033-3213211120123312-1230031132101113"></a>

<a id="canonical-3301013122210100-1233201111103100-1113300023213123-2133003021003120-0032032320300113-3133231002320020-3321031102300233-2101322221031131"></a>

## certificate_url property — tls_certificates / 032212232332 / 4

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

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

- [custom_hash_algorithms](data-sources--workload--reference--group-010.md#canonical-2312312231122010-0112321210102210-0312231303231022-0321032030033220-3111012131003300-1101001121022133-3031002332131001-1312233303023232): complete subsection reference.

<a id="canonical-3332133131032012-0010312022023213-0110001130303320-1323012302023301-3222013101203233-2012112302000330-1002311232323021-0101331003223110"></a>

<a id="canonical-0312331100332211-3013131030022123-2320332023201333-1033101221301000-2003220200222233-3232230302103022-1201001213132120-3230203112221230"></a>

## description_spec property — tls_certificates / 032212232332 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--workload--reference--group-010.md#canonical-1120032103312330-1013313123310133-2201313112202130-0233022000320132-1132332303200112-1320120221103112-3211012311202300-3111022123232033): complete subsection reference.

- [private_key](data-sources--workload--reference--group-010.md#canonical-1302100033300221-2102012111111321-1023301023321300-0110012000301230-0331333033011211-2232220103030223-3220002123030011-1200300112330031): complete subsection reference.

- [use_system_defaults](data-sources--workload--reference--group-010.md#canonical-2122002210233200-1113103222322201-2011310200023302-3300211211332303-0333300103200322-3212232111032221-0331213213021101-0011032123121012): complete subsection reference.
