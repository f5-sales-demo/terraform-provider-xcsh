---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3210323323021023-2301233332012302-1132311331201003-0212201012311032-1233312303310322-1122313230211311-2013300200000020-0033323210303031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-008.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
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

<a id="canonical-3313232100230310-3021313221103300-3002212332320312-3113011333020223-3300231202311213-3003013333113102-1030033032201331-1322033002130323"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](data-sources--workload--reference--group-009.md#canonical-3331320130210331-3111131030321131-0323013231233332-2201020112111222-3232121230120132-0030023001010010-1312233131133023-3112223001300231): complete subsection reference.

<a id="canonical-3331320130210331-3111131030321131-0323013231233332-2201020112111222-3232121230120132-0030023001010010-1312233131133023-3112223001300231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-008.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
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

<a id="canonical-2002330100202131-3221300121032222-2232301313000130-3131002132202203-3302302203001110-3330322302313211-3013213303232301-2023231011213211"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](data-sources--workload--reference--group-009.md#canonical-1301032131112110-0233013120023201-0221310221301202-1200000223103010-0333103202123202-0011300212220201-2103012132221023-0002122110003011): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-009.md#canonical-0222320211203212-0331232131110201-0301310121003012-1110111331301321-3110232003131013-1013030123111002-1022121121302101-1313110320020300): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-009.md#canonical-3332301131202033-1122022103001200-2213212132310303-3001312000201000-2110032201201313-2332011200032131-0313131010133023-2331111131312003): complete subsection reference.

<a id="canonical-1301032131112110-0233013120023201-0221310221301202-1200000223103010-0333103202123202-0011300212220201-2103012132221023-0002122110003011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-008.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222320211203212-0331232131110201-0301310121003012-1110111331301321-3110232003131013-1013030123111002-1022121121302101-1313110320020300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-008.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332301131202033-1122022103001200-2213212132310303-3001312000201000-2110032201201313-2332011200032131-0313131010133023-2331111131312003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-008.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300333110331030-2300220013021003-3211133121101120-3012231311030202-0312102232102223-1020202303022201-0030301332102213-0022223231322122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-008.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-0212333312221013-0132102332023213-0320122030020002-3011223203013231-3311123321220231-1310112122303023-0202010323211110-1311003322011011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

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

<a id="canonical-1303310302200120-1022220302203012-3030003321113020-2000202021010232-0130020133221331-3210002111033101-1312210103332222-1013230332320232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-008.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-3302213132033132-1230203212313111-3011130100001223-3212023302003030-1102013333330330-0310033013233003-1012202102131003-0102230330320313"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

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

<a id="canonical-1013232211321311-1221201331033333-2201030202022021-1320003311233133-0220230311310331-2232322213312220-1302012300002130-3310213120300233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer` properties

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

<a id="canonical-2233310030023222-0131300231233113-3030232101112200-1312321321013102-3112131313133123-1322203203132333-2210331221110011-0133101213221230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through` properties

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

<a id="canonical-3322111131021231-2001130012331203-1312332223212202-1223010323223313-0002020210033213-0120200332320012-3130023000210301-2112321010002003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params` properties

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

Additional upstream details:

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

<a id="canonical-0302113121210021-3220233310112002-0002122002113332-0311313100322110-3111030132020013-1213010330013333-1131312332323210-3333331311322212"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params`

- [certificates](data-sources--workload--reference--group-009.md#canonical-1301033320033002-0103021100223102-2300222133210001-1203021323012013-2232300231322121-2130113323330213-2023123331110021-2302003003033000): complete subsection reference.

- [no_mtls](data-sources--workload--reference--group-009.md#canonical-3022220032303311-0302131212132302-1033232221302310-0132112112132131-3223130103202323-0000313011013112-2230230103220322-2020200300201311): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-009.md#canonical-1133031130201122-0201003200011212-2221030021203322-2133332333333002-3302131022132331-2303132230320130-2301120223312101-3103131200232222): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-009.md#canonical-0202033310131322-1320211131131020-0122011020002120-1022330311232030-2322222101011030-1110301022313120-1300002301113101-2111020021320003): complete subsection reference.

<a id="canonical-1301033320033002-0103021100223102-2300222133210001-1203021323012013-2232300231322121-2130113323330213-2023123331110021-2302003003033000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates` properties

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

<a id="canonical-2003212122131211-2321213012010012-1212321220130221-2102322311313231-0100333200012121-0121221210201311-3120011121211311-2231301022200300"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates`

<a id="canonical-2320202031203110-0010011012330201-0212303000312322-1331011021003023-2112322201100001-2223122203200130-3201033032220013-0312221202102001"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates.name` property

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

<a id="canonical-0210310130210120-1303310230010000-3030330013303223-3323320232322101-1103001123012300-3301330132001020-0010323233110022-1113000103320213"></a>

<a id="canonical-3231111302212122-2112312310330230-3211133001103132-3313331300032010-0120111213033132-0011300030021112-3210321322033133-0300221132132112"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates.namespace` property

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

<a id="canonical-0101213120131110-3021311331310212-0032000022211110-0310233320101331-2202132032030113-0233200213102011-3331321110211103-3323021012321030"></a>

<a id="canonical-3210001231211021-3321111330233002-3122333330100110-3102100133123023-3103000310031310-0100200300111101-2001100121302001-1222220122220312"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates.tenant` property

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

<a id="canonical-3022220032303311-0302131212132302-1033232221302310-0132112112132131-3223130103202323-0000313011013112-2230230103220322-2020200300201311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls` properties

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

<a id="canonical-1133031130201122-0201003200011212-2221030021203322-2133332333333002-3302131022132331-2303132230320130-2301120223312101-3103131200232222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config` properties

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

<a id="canonical-2212202321113113-1201002100300213-1023233033113023-3023301210002023-2212110210033010-3303021122211231-1113022332303023-2223231110023120"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config`

- [custom_security](data-sources--workload--reference--group-009.md#canonical-3101302131121110-2013010220213123-2231300332033021-1120311321332213-3001132000030001-1122111230030323-0323113332020011-1221320110221032): complete subsection reference.

- [default_security](data-sources--workload--reference--group-009.md#canonical-3330103230033201-3202102322300020-1321010222201301-1201233212031103-2211301200303031-3000111122322101-2223202021333202-1113011110110030): complete subsection reference.

- [low_security](data-sources--workload--reference--group-009.md#canonical-3311121211221310-2312002033213230-3223110321131132-2210231122122033-1120033031013011-0311011101000231-1101232303212111-1000001310013001): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-009.md#canonical-1101313220012023-2232012323321112-3230222131313112-1233002131122122-2301230333200331-3313210121130131-0233133303022112-1310312003233212): complete subsection reference.

<a id="canonical-3101302131121110-2013010220213123-2231300332033021-1120311321332213-3001132000030001-1122111230030323-0323113332020011-1221320110221032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security` properties

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

<a id="canonical-1222132210210330-1311201300022003-3311230032032003-0113100112113122-0212223233200010-2120312202223002-2311002012101232-0222002113032111"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security`

<a id="canonical-1011113103103010-3333103222001330-0102201201131220-1102232301010213-1312133222311323-1011233022330331-0012330011020221-0033023313002133"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-0102313113030230-2022201100113302-0112121222021111-1100133022203123-2303032302032330-1231220131312120-2310003320223003-1011123031311121"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.max_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-0200030311001300-1301202313010311-0223322122023212-3031310033103000-2330323033003020-0223030211210320-2321132310233123-1303001202210101"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.min_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-3330103230033201-3202102322300020-1321010222201301-1201233212031103-2211301200303031-3000111122322101-2223202021333202-1113011110110030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security` properties

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

<a id="canonical-3311121211221310-2312002033213230-3223110321131132-2210231122122033-1120033031013011-0311011101000231-1101232303212111-1000001310013001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security` properties

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

<a id="canonical-1101313220012023-2232012323321112-3230222131313112-1233002131122122-2301230333200331-3313210121130131-0233133303022112-1310312003233212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security` properties

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

<a id="canonical-0202033310131322-1320211131131020-0122011020002120-1022330311232030-2322222101011030-1110301022313120-1300002301113101-2111020021320003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls` properties

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

<a id="canonical-1233021001210301-1313322300222011-0202111231203011-1303131131210121-1212331322333231-2311133210210030-1023310033232020-3220310300102131"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls`

<a id="canonical-1103020303110200-2300002122000123-0012303023121113-0202003221122202-0003323223013233-0200030002001233-2231322222313200-1121230321132000"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.client_certificate_optional` property

Type: `"bool"`. Computed.

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

<a id="canonical-0133021000211010-3200320220010200-3333310033231213-0103333232120221-0230213321001202-0203012203130230-0221021032111023-2130110321321230"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca_url` property

Type: `"string"`. Computed.

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

<a id="canonical-2222033212130122-1001021102002121-3211312033211330-3131112203123203-3231132102101020-2000012021132331-2302320123103312-3203110121212220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl` properties

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

<a id="canonical-1223130032101001-1332013133212110-0230013123032023-1002031033331210-0030300131113102-0322032020133222-0200102002223130-1112210321201130"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl`

<a id="canonical-0010203312213112-3002110203030132-2303222221112000-2012330313313313-0203111113212302-0032002000231032-0012100313031123-2013223230302233"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl.name` property

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

<a id="canonical-0111011320332020-3332333232333011-1222132303110200-2033331132121122-1030123003200033-3310202302221321-3020103003013331-3203032103001212"></a>

<a id="canonical-0111022220030302-0032312023333303-0020303113113222-2002032311113221-0300120323333033-1333012301310103-0333000230011032-0221323200230200"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl.namespace` property

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

<a id="canonical-1002320123320012-1030020110013213-0112313330212331-2023230211132120-0222030200033101-1200103100120032-0331302101230312-1322011200230103"></a>

<a id="canonical-1013033001023210-2230032222123303-1032031203303002-1031320110303020-1303201003111211-2112220213303112-2101112200131303-2301131200133132"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl.tenant` property

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

<a id="canonical-2301231110110112-1111001101300301-2230333110011302-0001000303112222-1212022120003211-1123102310213031-0312223033232100-2121133133231033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl` properties

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

<a id="canonical-1232323020201333-0333321220001321-1123223300000303-1332112333002313-1122013030001311-1033102220300230-2020133202112331-0112332013013022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca` properties

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

<a id="canonical-0030221202131330-0321221130112112-0312123333033230-1303222201233101-2311013120031203-0212101303031021-1122101203013002-3330213120311311"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca`

<a id="canonical-3002221213203103-1032011013111000-0130310132212133-2203201320222112-0212131202231011-0001020331220231-0221031021130130-3220010333031112"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.name` property

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

<a id="canonical-2030031110101202-1312010021213100-0123303302102011-2223203023113023-3122102213331323-0032322023220011-0121213233321020-3102332103013232"></a>

<a id="canonical-1110022203311231-1133303131013313-1013323233012303-2001132321303322-0132312112323002-0101213132030300-1323101133110323-0312320233230111"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-3223012030322231-1322331221230030-0332302013221212-2310203303000100-3113330203103030-1312000133022021-2110001201322211-0211113231020222"></a>

<a id="canonical-0002113120010123-0132311222200131-3230120101001103-1032221130210330-2123112233000323-0320323122223002-3313110303020310-1301213023313310"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-1221332103023213-2210121102102321-1000330013211131-2022100133032133-1120130233320130-2122202132220112-0303333110033023-2201023022213311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled` properties

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

<a id="canonical-1000013110310033-0010112331030310-0313201011122001-1113011011003033-3113120002322303-1302311302132010-1220213333113203-3010110303100013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options` properties

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

<a id="canonical-1111331033223010-3023121331010022-0023321223103130-0122213012012332-3010032131213320-1133033111130213-2330030102223233-2210101320211310"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options`

<a id="canonical-2011122110000132-0210000001030231-3023210300221020-2321331020300021-1332221200103131-0330011001003301-0020320220111301-0300003130222002"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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

<a id="canonical-3030231203112023-2233233210333103-0332221021303202-1101310333001013-0301031200110232-3203210012020013-3333202011021213-3213221121020113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters` properties

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

Additional upstream details:

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

<a id="canonical-1310332032133113-2203323200300010-0010233312111112-0103310010113312-3312103031021111-0002331010010003-3123322031031101-1202031123103023"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters`

- [no_mtls](data-sources--workload--reference--group-009.md#canonical-1022220202023211-2112103231203311-3211313330213002-0102131100230102-2212031023301210-0012233320113301-0211330101222230-0130200313032110): complete subsection reference.

- [tls_certificates](data-sources--workload--reference--group-009.md#canonical-0033313103322112-0203330310133201-1121203002123230-1312213033300032-2111110320220201-1011112012220232-2020022322133332-0013130111131230): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-009.md#canonical-2021120001321031-2332013010231300-0320101103110022-3002023131032222-2331001221230213-3231132113321331-1002110103130201-1313302223310002): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-009.md#canonical-1221213020103013-1101032121313122-3122223200121312-3212031302101022-2110022231001123-2002020021110131-0300022001201321-2030210010213223): complete subsection reference.

<a id="canonical-1022220202023211-2112103231203311-3211313330213002-0102131100230102-2212031023301210-0012233320113301-0211330101222230-0130200313032110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls` properties

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

<a id="canonical-0033313103322112-0203330310133201-1121203002123230-1312213033300032-2111110320220201-1011112012220232-2020022322133332-0013130111131230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates` properties

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

<a id="canonical-1030022103100031-1223130233010000-3100022302311310-1110133212122020-2123230202323010-3201113112020132-1221332222101231-0012122311212323"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates`

<a id="canonical-1321103021330322-0311023010021331-0331202023012312-2233200120030013-2001012211113223-2121301103213033-3213211120123312-1230031132101113"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.certificate_url` property

Type: `"string"`. Computed.

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

- [custom_hash_algorithms](data-sources--workload--reference--group-009.md#canonical-2312312231122010-0112321210102210-0312231303231022-0321032030033220-3111012131003300-1101001121022133-3031002332131001-1312233303023232): complete subsection reference.

<a id="canonical-3332133131032012-0010312022023213-0110001130303320-1323012302023301-3222013101203233-2012112302000330-1002311232323021-0101331003223110"></a>

<a id="canonical-1221102102233013-3100131120300232-2201311011021003-2201211203003131-2202213130021112-1312311320213032-2333112123213000-0320323221211112"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--workload--reference--group-009.md#canonical-1120032103312330-1013313123310133-2201313112202130-0233022000320132-1132332303200112-1320120221103112-3211012311202300-3111022123232033): complete subsection reference.

- [private_key](data-sources--workload--reference--group-009.md#canonical-1302100033300221-2102012111111321-1023301023321300-0110012000301230-0331333033011211-2232220103030223-3220002123030011-1200300112330031): complete subsection reference.

- [use_system_defaults](data-sources--workload--reference--group-009.md#canonical-2122002210233200-1113103222322201-2011310200023302-3300211211332303-0333300103200322-3212232111032221-0331213213021101-0011032123121012): complete subsection reference.

<a id="canonical-2312312231122010-0112321210102210-0312231303231022-0321032030033220-3111012131003300-1101001121022133-3031002332131001-1312233303023232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-009.md#canonical-0033313103322112-0203330310133201-1121203002123230-1312213033300032-2111110320220201-1011112012220232-2020022322133332-0013130111131230)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-0021100032211213-2121230233320101-3232222011121212-1113023002112322-2122130112300103-2130011300010121-1322001120233300-2322010011300221"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0110103133200303-3130013230232122-1212011321302013-2231121201313121-3203120332323023-3121111131000322-0033001212203102-3122003210131311"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms`

<a id="canonical-3210230001013332-2232323002222111-3232222102223011-2210303023302320-1212110123011303-3133232323032212-1103002022301332-2020221311120020"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1120032103312330-1013313123310133-2201313112202130-0233022000320132-1132332303200112-1320120221103112-3211012311202300-3111022123232033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-009.md#canonical-0033313103322112-0203330310133201-1121203002123230-1312213033300032-2111110320220201-1011112012220232-2020022322133332-0013130111131230)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-3201302103231012-0110233003312212-0131310120211020-3101201121310310-2003333320013211-0312030101220122-1201030120111120-0013322311302010"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

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

<a id="canonical-1302100033300221-2102012111111321-1023301023321300-0110012000301230-0331333033011211-2232220103030223-3220002123030011-1200300112330031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-009.md#canonical-0033313103322112-0203330310133201-1121203002123230-1312213033300032-2111110320220201-1011112012220232-2020022322133332-0013130111131230)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-2312023321221301-2122202202120330-1220121302000020-0231111132003203-1021310213310000-1001100023023123-3013020113122102-0130300103212311"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-2122000131002223-3121221232131332-2010321233331003-2330113232130012-3233112110122210-1232011112223123-0231103210221023-1210131221131300"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--workload--reference--group-009.md#canonical-1301323120313211-0232113332313202-0200332020031030-2033112203112203-2001023221210212-3020301122222201-2230023201130320-0201332231312102): complete subsection reference.

- [clear_secret_info](data-sources--workload--reference--group-009.md#canonical-0222033232212000-0110032031132130-3333000110020302-0101111033023321-2011232322101313-1030120100333231-3213130320121321-2212023321230213): complete subsection reference.

<a id="canonical-1301323120313211-0232113332313202-0200332020031030-2033112203112203-2001023221210212-3020301122222201-2230023201130320-0201332231312102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-009.md#canonical-0033313103322112-0203330310133201-1121203002123230-1312213033300032-2111110320220201-1011112012220232-2020022322133332-0013130111131230)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-009.md#canonical-1302100033300221-2102012111111321-1023301023321300-0110012000301230-0331333033011211-2232220103030223-3220002123030011-1200300112330031)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3332323122233123-0210013123310010-0102223102031200-3101203010121020-2332312211113220-3210311222130313-2122001200032013-1021032123302331"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1031222201132333-2331312223332310-3200022103213313-2120020213330310-0011221120113110-2332311021302300-3030022331311231-2331302311230300"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-2021120302220101-0332311230022131-2323200132113230-0210320203203112-2010332321301310-3200012131332020-3020021022232223-0211031020103200"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-2002330300300320-0301022330300320-0102320302330201-3013213333010230-1320123232030120-1100033213123212-1110311201012122-1320332210300123"></a>

<a id="canonical-2333011003013223-2013112032123100-1123030303210230-3211333221121331-0301122001132100-3110030310030200-1132000122100222-2110203221132103"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2211001311321030-1320203310201200-1002312121122102-1011210221130320-1213200323201113-2211302003130223-2323132223111131-0201213130123331"></a>

<a id="canonical-1331033220321310-3013232310302111-2312013010103233-1330222132302321-2101122313132130-2303100020003221-2030232323011323-1000202232100121"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-0222033232212000-0110032031132130-3333000110020302-0101111033023321-2011232322101313-1030120100333231-3213130320121321-2212023321230213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-009.md#canonical-0033313103322112-0203330310133201-1121203002123230-1312213033300032-2111110320220201-1011112012220232-2020022322133332-0013130111131230)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-009.md#canonical-1302100033300221-2102012111111321-1023301023321300-0110012000301230-0331333033011211-2232220103030223-3220002123030011-1200300112330031)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-2321202302330223-3230022031033230-3310002233121322-0023321101132302-1000031132313220-2110311102131021-0211002103132123-2010131322003111"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1223032020303212-0310110022232101-3300003023122130-2001222331211123-2310221321101233-1001112002332031-1323330323231000-3120213311100203"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info`

<a id="canonical-2133033100103102-2112133221301032-0211220313003211-2330022211201012-0223211033031103-3321303011211303-0202000312332113-0323201321131132"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2210310302222311-3131000032332031-2231211232303201-1122121202020133-0333123021330221-3102100021331013-2232101200203211-2032021221232030"></a>

<a id="canonical-1202233111302101-1210212013030202-1010121332212231-0213321133321322-2030122310030201-1020210203322112-3110321321201112-1233113203000121"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2122002210233200-1113103222322201-2011310200023302-3300211211332303-0333300103200322-3212232111032221-0331213213021101-0011032123121012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-009.md#canonical-0033313103322112-0203330310133201-1121203002123230-1312213033300032-2111110320220201-1011112012220232-2020022322133332-0013130111131230)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-0001123100223312-1221332013111013-0311200312100131-1323130002331302-3032232031002111-3312010321130031-3031102031232013-2000332102101103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

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

<a id="canonical-2021120001321031-2332013010231300-0320101103110022-3002023131032222-2331001221230213-3231132113321331-1002110103130201-1313302223310002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config` properties

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-1101221000211203-1302200022010211-0023130222012123-3000230201131310-1020011131001021-2011020321201300-0000001210011002-1201033312113120"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2201233302332120-3120110000222303-2110333202320020-0001230130203022-0300122301223021-0120302230122321-3223032331121131-1200323011330232"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config`

- [custom_security](data-sources--workload--reference--group-009.md#canonical-0021110322311100-1233331010330203-3321321030023033-2220303312221012-2001110332020123-2311233130330321-3032201310201310-1130112313200010): complete subsection reference.

- [default_security](data-sources--workload--reference--group-009.md#canonical-0222312130302322-2112211122212322-3113000221000102-3232012333320221-3110123220233001-2303321111110021-2300223200220220-1231210120013112): complete subsection reference.

- [low_security](data-sources--workload--reference--group-009.md#canonical-2312113031333231-1211113201133220-2222301332010020-1013233222130030-3033130000310033-3011030111000233-0020132103201031-0030233101000200): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-009.md#canonical-3112221012213003-1211301233130313-0213031203222200-0203201033303020-0102013223033213-2103002310110001-2020230023130011-3232012210032110): complete subsection reference.

<a id="canonical-0021110322311100-1233331010330203-3321321030023033-2220303312221012-2001110332020123-2311233130330321-3032201310201310-1130112313200010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-009.md#canonical-2021120001321031-2332013010231300-0320101103110022-3002023131032222-2331001221230213-3231132113321331-1002110103130201-1313302223310002)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-1202312132111030-3201133220313331-3303001022012332-3210123323203122-2202321201303300-0200332231102120-0302012023333030-1032303023132110"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0100103100220002-2100301121331231-0200121300102000-0031121220320220-3021231122130213-1222130303331133-3231210202301100-1032301003032112"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security`

<a id="canonical-3013203213300203-2323203223201233-3323302002130113-0110111000112333-3101103030232201-0023021203212330-0003012103231012-0201033011312030"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-0321322132132111-1311022321323101-0102002300302130-0233331022333232-3231213330211211-0230102322133112-2101123221301132-1110003113133202"></a>

<a id="canonical-2200300220233131-1132130030123213-1001320020102001-3030203303023312-0202200332201301-3301130330020132-0012220120023222-2200212230002213"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.max_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-2323313132010223-2221313213210323-1101123331230331-0121221102221121-1232002302222322-2212202301203301-0310222313002300-2301233200100100"></a>

<a id="canonical-3312033030023103-0113303023201230-2133201132320210-3330030102113021-0022031222223130-1102330311223320-1102322020213110-0301231210100123"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.min_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-0222312130302322-2112211122212322-3113000221000102-3232012333320221-3110123220233001-2303321111110021-2300223200220220-1231210120013112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-009.md#canonical-2021120001321031-2332013010231300-0320101103110022-3002023131032222-2331001221230213-3231132113321331-1002110103130201-1313302223310002)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-3232221332013112-1030022233003231-0330131210330331-2000323022001132-1012032130333111-0000330010000231-2101113310230233-0020102131220100"></a>

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

<a id="canonical-2312113031333231-1211113201133220-2222301332010020-1013233222130030-3033130000310033-3011030111000233-0020132103201031-0030233101000200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-009.md#canonical-2021120001321031-2332013010231300-0320101103110022-3002023131032222-2331001221230213-3231132113321331-1002110103130201-1313302223310002)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-1210013230320120-3112233332220231-3320301032321323-3323030023203211-2330013120132311-2001102112231201-0120302103300113-1322323331033330"></a>

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

<a id="canonical-3112221012213003-1211301233130313-0213031203222200-0203201033303020-0102013223033213-2103002310110001-2020230023130011-3232012210032110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-009.md#canonical-2021120001321031-2332013010231300-0320101103110022-3002023131032222-2331001221230213-3231132113321331-1002110103130201-1313302223310002)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-3330331212121100-1032102103002132-0110101300303010-3120211002220122-0131313032230320-2321011121133031-3230320233103211-2130132212213301"></a>

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

<a id="canonical-1221213020103013-1101032121313122-3122223200121312-3212031302101022-2110022231001123-2002020021110131-0300022001201321-2030210010213223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls` properties

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-2100213230021303-0000233312120112-3002321122202333-2131032302330201-3130120103221223-0220320321203023-2010102122022133-2212030313230113"></a>

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

<a id="canonical-1011020201101233-0032101113121312-1331220202220013-2332102123000132-1012030030210210-2131111032300123-0321100023333132-1200201111031213"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls`

<a id="canonical-0310213112301232-0210002322020021-3310202000300311-0120232220123303-1320313203121223-2320220022331103-3333102213003013-2312003300020121"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.client_certificate_optional` property

Type: `"bool"`. Computed.

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

- [crl](data-sources--workload--reference--group-009.md#canonical-2010231112203210-2321321230031230-1031202100123203-3300331013111220-3111122131000211-2201200232223312-2212130131231033-2002131230020022): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-009.md#canonical-1312031310323333-2320230222022213-2202112330311000-0231000212131010-0211003003111312-1211002332011011-0202101130110030-2211010111201333): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-009.md#canonical-0120233122003201-1312022001222131-2303323213013223-1220332221121021-1032303201233211-3200001330103330-2122121200222111-1102133311331310): complete subsection reference.

<a id="canonical-0011201103202121-3012200231033031-2111212212302313-1101223322123232-2332310111131310-0211223331320303-1313232301030321-2003003023233231"></a>

<a id="canonical-1121200221130111-0232133323203120-3312003121333211-3203202331231110-1020130123313310-3211303230032313-0231020210313023-2101223022031001"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca_url` property

Type: `"string"`. Computed.

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

- [xfcc_disabled](data-sources--workload--reference--group-010.md#canonical-2221030012022310-3122122210023232-0010001313101112-1230130022123321-0010012111231201-1221200201133201-2011103003002200-1221013312133231): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-010.md#canonical-3112132121310121-2032333222021130-2220001312011220-2132100032300100-1022221033033022-2300110303003221-1013030300130330-2021031122101133): complete subsection reference.

<a id="canonical-2010231112203210-2321321230031230-1031202100123203-3300331013111220-3111122131000211-2201200232223312-2212130131231033-2002131230020022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-009.md#canonical-1221213020103013-1101032121313122-3122223200121312-3212031302101022-2110022231001123-2002020021110131-0300022001201321-2030210010213223)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-3023333322021303-1103301102133203-3320223223031232-0032302000111232-0102221102012123-2012120321333112-1202032003221003-1113213230120200"></a>

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

<a id="canonical-0233203203312102-1111332223122310-2001202333103212-2022203301021221-1020021130033201-2322002023102033-1031102103033001-0322302103300121"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl`

<a id="canonical-0313320131002031-3333113030222001-0020003003332132-0201212302220232-0102210010132110-1310222022213212-3101302032331002-1200011203311122"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl.name` property

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

<a id="canonical-0113103023200023-1232110111113020-2020323022212311-2131122101310100-3333333003110011-1321133101001331-1100312133202021-0311202330220111"></a>

<a id="canonical-1230332112100211-0012023222320022-1332223032123202-2030031331311212-2211332023032123-1031313013210212-3031021000122020-0010221221223102"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl.namespace` property

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

<a id="canonical-0022023113320003-3230312032123023-3311332211101220-1121112213220013-2323200110102112-0000210201220102-3303000031213110-0202301202022233"></a>

<a id="canonical-3213211200121020-2103231022103302-1110123033303030-1003201120310121-1133332112333220-3013001212010023-3223033100122220-2231032132120013"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl.tenant` property

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

<a id="canonical-1312031310323333-2320230222022213-2202112330311000-0231000212131010-0211003003111312-1211002332011011-0202101130110030-2211010111201333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-009.md#canonical-1221213020103013-1101032121313122-3122223200121312-3212031302101022-2110022231001123-2002020021110131-0300022001201321-2030210010213223)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-1323111223031232-2302220031112320-3133131001212020-0031203023201212-2210001222021213-1220302312032333-3002110131011032-3010311230030220"></a>

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

<a id="canonical-0120233122003201-1312022001222131-2303323213013223-1220332221121021-1032303201233211-3200001330103330-2122121200222111-1102133311331310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-009.md#canonical-1221213020103013-1101032121313122-3122223200121312-3212031302101022-2110022231001123-2002020021110131-0300022001201321-2030210010213223)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-1212111312031201-1303333002211022-1202202010131100-0333201103103000-2302120213312023-3001031233032030-1033113223031223-2310200232230230"></a>

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

<a id="canonical-3213133120203211-0120032000010022-1313301301120200-3223312210313323-0020113312103123-3111023320110220-2202330202130111-2020001121013213"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca`

<a id="canonical-2112113003312210-2122301231212213-2022012200101133-0203222211133000-2333222221113311-2212133011010332-2323011320231303-3302310321100032"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.name` property

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

<a id="canonical-0212303000131300-3001132320323003-3301110301322033-1301133220130302-2101001201020130-1101223111130001-3320110021310302-3203003023321132"></a>

<a id="canonical-0303132303302201-2320320111300133-3321101203010003-2103110302301233-0103301132113113-3210300300333110-3323322232301203-3010323232111310"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-2311033212221323-1102332032001221-3002033133222001-2311020111200032-1331133313311130-1002322131023333-0311122111211231-0212213033231220"></a>

<a id="canonical-2023202201103322-0000213333230001-0200111202032021-3131023221030312-2012021003211212-1230011103012110-3302303021110132-2212131311101200"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.tenant` property

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
