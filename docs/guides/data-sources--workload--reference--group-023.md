---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-0020332313030311-2112212112321203-3211212300311000-2010320223013330-2201030130233320-1321010003011120-2003311030320100-0132001333021303"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 310302000330 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-3103012111223102-2323200023212331-0010133311323102-1302201311002220-0233003131123331-1001102120301011-3332012213301122-1001000212330320)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-022.md#canonical-3211312311101322-0011233123001133-2222321232121121-1130131311030122-1331023130321112-0202213032010323-2000222320222013-2310033313303222)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-1212223032101101-3331311033003230-3210021213303101-1330233132213112-2132321303222211-0321013330202222-2031301121033130-3102303033023111"></a>

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

<a id="canonical-2100012030320122-2201310323112120-3131312200011023-1112313123022213-2321120113221002-0210233312213310-0321112311321222-1230222301030332"></a>

## Direct properties — header_transformation / 310302000330 / 3

- [default_header_transformation](data-sources--workload--reference--group-023.md#canonical-2202313201323121-2130021110112101-1330102212032310-3322112310003210-0112321131331013-3230022201031000-3210031033222321-1202131302013300): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-023.md#canonical-1302200023021231-2033302333221013-0121111211121311-0113031203332202-1113003201012020-0121133201003122-1203021233300321-3133000111022011): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-023.md#canonical-2323111132010002-0002231230303010-0322213211323220-0321100203210001-1223231332230312-0321111232103020-3322030110322131-2021331201221310): complete subsection reference.

<a id="canonical-3233310011100102-0001303221100313-0132233230313032-2032123012003130-3103210201112211-3131010101012121-2312023032011103-3231101030120101"></a>

## Next pages — header_transformation / 310302000330 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--workload--reference--group-023.md#canonical-2202313201323121-2130021110112101-1330102212032310-3322112310003210-0112321131331013-3230022201031000-3210031033222321-1202131302013300)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--workload--reference--group-023.md#canonical-1302200023021231-2033302333221013-0121111211121311-0113031203332202-1113003201012020-0121133201003122-1203021233300321-3133000111022011)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--workload--reference--group-023.md#canonical-2323111132010002-0002231230303010-0322213211323220-0321100203210001-1223231332230312-0321111232103020-3322030110322131-2021331201221310)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-022.md#canonical-3211312311101322-0011233123001133-2222321232121121-1130131311030122-1331023130321112-0202213032010323-2000222320222013-2310033313303222)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2202313201323121-2130021110112101-1330102212032310-3322112310003210-0112321131331013-3230022201031000-3210031033222321-1202131302013300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120013130320310-1021031020310313-1312221001031110-1203023031221023-3022112312330232-1330112001123300-0232212201310213-1230133131020302"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 030302121003 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-3103012111223102-2323200023212331-0010133311323102-1302201311002220-0233003131123331-1001102120301011-3332012213301122-1001000212330320)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-022.md#canonical-3211312311101322-0011233123001133-2222321232121121-1130131311030122-1331023130321112-0202213032010323-2000222320222013-2310033313303222)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-022.md#canonical-0313000213023230-1103132333313023-2122323321233320-3113020132003202-0112333312000310-2012012221020031-1121331233130323-0313023300320013)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-3010123231001130-1331303020200022-3000132300320010-2033231011331100-2130021223113301-2301020011202112-1223023210132212-3103102100020130"></a>

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

<a id="canonical-0212023203123221-3121033203233320-1122223213202221-3111333200310123-1333000020213032-1002030132331011-3010113310210203-3312101321112323"></a>

## Direct properties — default_header_transformation / 030302121003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122210312220212-3021213010211302-0013131103030233-3010213121330333-2012102000221112-0302203321220011-2312032102032231-1121322032220301"></a>

## Next pages — default_header_transformation / 030302121003 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-022.md#canonical-0313000213023230-1103132333313023-2122323321233320-3113020132003202-0112333312000310-2012012221020031-1121331233130323-0313023300320013)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1302200023021231-2033302333221013-0121111211121311-0113031203332202-1113003201012020-0121133201003122-1203021233300321-3133000111022011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101322113211001-1300301110022232-1212232333123333-0010121112133332-0321230131031033-2203032132001300-0312033102003222-3301223032020322"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 023303120331 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-3103012111223102-2323200023212331-0010133311323102-1302201311002220-0233003131123331-1001102120301011-3332012213301122-1001000212330320)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-022.md#canonical-3211312311101322-0011233123001133-2222321232121121-1130131311030122-1331023130321112-0202213032010323-2000222320222013-2310033313303222)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-022.md#canonical-0313000213023230-1103132333313023-2122323321233320-3113020132003202-0112333312000310-2012012221020031-1121331233130323-0313023300320013)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-0002330312120112-1030203003133003-3212220131030031-3120323111112321-3002023020102230-2300323231013002-1330200032221221-0332330311132100"></a>

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

<a id="canonical-2330110320013031-3311332211110022-0103232110332113-0322200123131222-2100200221310311-3000221120130101-3300122120103132-2120203021103210"></a>

## Direct properties — preserve_case_header_transformation / 023303120331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000333310113222-1331101020111322-2001202100303011-2013300301102013-2303321010301101-3021122233020123-1130030333103223-1003202010012320"></a>

## Next pages — preserve_case_header_transformation / 023303120331 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-022.md#canonical-0313000213023230-1103132333313023-2122323321233320-3113020132003202-0112333312000310-2012012221020031-1121331233130323-0313023300320013)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2323111132010002-0002231230303010-0322213211323220-0321100203210001-1223231332230312-0321111232103020-3322030110322131-2021331201221310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223001333112103-0012302333011202-2301300101002133-1311303200333332-3323210132021012-1321003030202312-0212113333121303-1020312012133020"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 302211021323 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-3103012111223102-2323200023212331-0010133311323102-1302201311002220-0233003131123331-1001102120301011-3332012213301122-1001000212330320)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-022.md#canonical-3211312311101322-0011233123001133-2222321232121121-1130131311030122-1331023130321112-0202213032010323-2000222320222013-2310033313303222)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-022.md#canonical-0313000213023230-1103132333313023-2122323321233320-3113020132003202-0112333312000310-2012012221020031-1121331233130323-0313023300320013)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-0013100120011322-0121103033223200-1111102321223012-3001220213022323-3313111131021331-2223013233220333-1100322322313303-0331121111232110"></a>

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

<a id="canonical-3210222310031313-2331333120130003-2220330001021332-3211222312210202-2132222330021313-2213300102021303-0232332001233021-0211031300232022"></a>

## Direct properties — proper_case_header_transformation / 302211021323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303213001312301-2201113232212333-1320233223211213-1200131231121110-1011111000030320-3122321100000030-2113332331232001-2220113132132222"></a>

## Next pages — proper_case_header_transformation / 302211021323 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-022.md#canonical-0313000213023230-1103132333313023-2122323321233320-3113020132003202-0112333312000310-2012012221020031-1121331233130323-0313023300320013)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1002130312012131-2130013003020321-3020012310312330-0003001123130301-3113012102022320-0031211011022012-0211201202322331-1102010222133300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022201110200123-2103332221221011-2330021300001211-3100322321002232-1210103021313123-1223332300113033-0103130302330202-0001220232221213"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 111323303222 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-3103012111223102-2323200023212331-0010133311323102-1302201311002220-0233003131123331-1001102120301011-3332012213301122-1001000212330320)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-3220100031200001-0111003123111100-1311330231021032-2221120013201223-3232111220323020-1030220132203321-1010112303301303-3313333233010313"></a>

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

<a id="canonical-3020221122033113-2021110113130003-2121310313233021-3002311333221332-0303311233233211-2101002211010232-3330301303223332-1333211102222101"></a>

## Direct properties — http_protocol_enable_v1_v2 / 111323303222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323320033121200-2303323023223102-2022110021332203-2001130331311001-2221121322310213-2221232312202131-2230010311303320-3023033312233303"></a>

## Next pages — http_protocol_enable_v1_v2 / 111323303222 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-3103012111223102-2323200023212331-0010133311323102-1302201311002220-0233003131123331-1001102120301011-3332012213301122-1001000212330320)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3203112021002200-2003130103200020-3102221112202010-1001201203131130-1332310133132030-3132031202003220-2310330323112100-0220010200131331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221000323223200-1330013033221122-2223001031332300-3233312001323101-2102313000011130-3020213111100201-3121232002300033-0023030000131200"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 331213121200 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-3103012111223102-2323200023212331-0010133311323102-1302201311002220-0233003131123331-1001102120301011-3332012213301122-1001000212330320)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-2102301302310011-3012111021312230-1210001031013021-1333210102012330-2013300102002311-1130210011133123-2101300112013322-2301313233331203"></a>

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

<a id="canonical-3033033213312003-1000012303322003-0133302022300003-0100232003033213-0321200123020301-0220132322023303-0302302031313320-1112233130320330"></a>

## Direct properties — http_protocol_enable_v2_only / 331213121200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002003001203131-1103102030032231-2110233223023200-0203323220321310-2112302200133010-2300213131302303-3012213113300313-2022110111323133"></a>

## Next pages — http_protocol_enable_v2_only / 331213121200 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-3103012111223102-2323200023212331-0010133311323102-1302201311002220-0233003131123331-1001102120301011-3332012213301122-1001000212330320)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0102000212101103-1332202100200201-0020022323300210-0303223303003330-1231312230131200-0133331031320223-2133030210013201-3021133112100221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233222103100233-0203321230311231-3003200310131133-0111313313013313-0330231320231012-0111300232211020-0012223020320321-3200201231111211"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.no_mtls — no_mtls / 112110130013 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.no_mtls

<a id="canonical-0111003121222312-3020313123232020-1112000113212021-2213323131310122-2311203211100100-3000130211211333-3023311323311233-0013100232221323"></a>

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

<a id="canonical-0203231233231012-0212312223320333-3310202113313033-2333300310322230-3200021302122303-1003300121301101-3122020000022200-1001130000211010"></a>

## Direct properties — no_mtls / 112110130013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132233011101320-3222321201311010-3333001100010003-0221012313113301-0002303203033111-3333210203000100-1033221022211302-0012131110230100"></a>

## Next pages — no_mtls / 112110130013 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3031212332220231-0321113012022312-0200221332210221-3132010220210300-3031230023332113-2003022322002010-1323013012111002-0230210102222131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333322011223222-2322231211331322-3010101002100303-2331120131131312-2330112303323023-0120123303132110-3110200113112010-3021022021300230"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer — non_default_loadbalancer / 012002310331 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer

<a id="canonical-3132210213023333-1130100300301011-3100233202202130-2120320202222310-2303313331201012-2132322103102310-1123103033031330-0123231121033302"></a>

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

<a id="canonical-0030331022102103-2121111201131331-1130210230311232-3100000131012220-1112102011201330-0120022101102132-3323121122023231-1233013300220010"></a>

## Direct properties — non_default_loadbalancer / 012002310331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112320120001200-2312230011333311-3332332201030112-3130011022310023-2000111223201333-1333112011233100-1123322331330123-1022301132313121"></a>

## Next pages — non_default_loadbalancer / 012002310331 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3121300110323012-2311211111332231-1331033331133131-2013310211001020-0200001010132331-2300201012013200-0012030313110330-2201201212000113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122000022001321-2310121013131121-1233100203112022-0132003113300301-1120201303133123-1202113221100203-2003100232011312-2000321122133210"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.pass_through — pass_through / 200231020330 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.pass_through

<a id="canonical-0113132333100033-1332113011132001-1320300303233231-3303221130132323-3131110212100230-3123102023332013-3103123302012222-0013030210301013"></a>

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

<a id="canonical-3010003003110233-1303021010213220-3210303321120010-0130312033021020-1120310233332303-0013002031003013-3301011102130122-1012221102303110"></a>

## Direct properties — pass_through / 200231020330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311132321122001-1312221211130023-1303003002313011-2023322111021221-1220311223321211-0302301032311311-3110112230323232-1213102032320032"></a>

## Next pages — pass_through / 200231020330 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2323123020222323-0201333323111230-1231221203230030-3330100031210233-0222100010111322-3313102010313113-3112021331102313-2133301202102100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103201013331321-0331103313312013-3222012002112220-1111332111230121-1131230331013123-2313213222222232-1210011331332333-0230200130220333"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config — tls_config / 033133023121 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config

<a id="canonical-3332031000331130-3121111120333010-1102101223001020-1133020122001210-1300230330302302-3020012102122301-0221201131200010-0130112221222212"></a>

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

<a id="canonical-2033322031122211-2322330232301121-1113102332002013-2330223310113113-1011023002003302-0010320232013123-1032112220232121-0330033203031320"></a>

## Direct properties — tls_config / 033133023121 / 3

- [custom_security](data-sources--workload--reference--group-023.md#canonical-3321110203012120-3000011120110010-3320030333133031-0212132032322011-1302031101032010-2131312302111321-3332323110230122-1102300212113022): complete subsection reference.

- [default_security](data-sources--workload--reference--group-023.md#canonical-2210101032122201-3032301102103102-0302313223310313-2323120323222332-2020313330203203-3201223131102103-3222001133031210-0321333130123113): complete subsection reference.

- [low_security](data-sources--workload--reference--group-023.md#canonical-0013101230002333-3003222322321331-3321101302000101-3022120330213133-1121322232002220-2110003203011020-0303100113130313-0313310011021322): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-023.md#canonical-1132221310301033-2020203122100013-1132033323313312-1100211222313023-1132132110231233-2022123300000232-1013333100222200-2122220213002232): complete subsection reference.

<a id="canonical-3203312100322120-2202232223321302-3301313120020312-3111122213311103-0133312111101300-1012233223213000-2301011301022013-0221102213000330"></a>

## Next pages — tls_config / 033133023121 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security](data-sources--workload--reference--group-023.md#canonical-3321110203012120-3000011120110010-3320030333133031-0212132032322011-1302031101032010-2131312302111321-3332323110230122-1102300212113022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.default_security](data-sources--workload--reference--group-023.md#canonical-2210101032122201-3032301102103102-0302313223310313-2323120323222332-2020313330203203-3201223131102103-3222001133031210-0321333130123113)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.low_security](data-sources--workload--reference--group-023.md#canonical-0013101230002333-3003222322321331-3321101302000101-3022120330213133-1121322232002220-2110003203011020-0303100113130313-0313310011021322)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security](data-sources--workload--reference--group-023.md#canonical-1132221310301033-2020203122100013-1132033323313312-1100211222313023-1132132110231233-2022123300000232-1013333100222200-2122220213002232)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3321110203012120-3000011120110010-3320030333133031-0212132032322011-1302031101032010-2131312302111321-3332323110230122-1102300212113022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103332212132322-3002210302130101-1223321311320201-0121212023231001-1103201310131030-2013000221301003-3313221320311021-2202213032220012"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security — custom_security / 120110000201 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-023.md#canonical-2323123020222323-0201333323111230-1231221203230030-3330100031210233-0222100010111322-3313102010313113-3112021331102313-2133301202102100)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security

<a id="canonical-3022002111333323-3032003011001313-2031002200132312-1333230033112331-2130303130113203-2000330121030131-1032313312301120-2003002332020101"></a>

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

<a id="canonical-2021111032330202-2023331023000311-2310232113132120-0313211111231332-0130033301332202-0303331312123022-3203023011132313-1103111130231003"></a>

## Direct properties — custom_security / 120110000201 / 3

<a id="canonical-0120130302320213-3022031332212121-0320231203113300-2220201100022321-1000230212312231-2000320201122210-0213200203213312-1002111321003231"></a>

<a id="canonical-3100110132203332-1323201111312033-1022233032221221-0200013011230211-3120213312033320-1311302222013020-1112121132011102-3211331323211001"></a>

## cipher_suites property — custom_security / 120110000201 / 4

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

<a id="canonical-1203013013103311-3222131101331010-0201312211100012-2001202311212130-1112320031212220-0002123031232030-1311110002133233-0132002030120032"></a>

<a id="canonical-0202301003222011-2000233111321222-0001301231132333-1320031201223202-2010112203200032-2202111302210022-1120132022122103-3022122132003123"></a>

## max_version property — custom_security / 120110000201 / 5

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

<a id="canonical-3012321332221330-3301300301032333-1023021101120213-2201121123131223-2201023333020101-1232120012302030-1032302333302210-2113303100220203"></a>

<a id="canonical-1321112113210231-0101002201030223-3312123200002201-0223310331022131-0100210312301112-2121230312132301-1201001223131210-3232232312002333"></a>

## min_version property — custom_security / 120110000201 / 6

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

<a id="canonical-3013112101013103-1333112213302103-2300322210101301-0110232201022121-1200020013031110-0030201231323120-2311132221203320-0012213233102303"></a>

## Next pages — custom_security / 120110000201 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-023.md#canonical-2323123020222323-0201333323111230-1231221203230030-3330100031210233-0222100010111322-3313102010313113-3112021331102313-2133301202102100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2210101032122201-3032301102103102-0302313223310313-2323120323222332-2020313330203203-3201223131102103-3222001133031210-0321333130123113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123333200220032-2232200132211112-3100211220100111-2213013102112102-3113200200021101-0212030100311022-1133002113111131-0012021222010232"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.default_security — default_security / 123003033203 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-023.md#canonical-2323123020222323-0201333323111230-1231221203230030-3330100031210233-0222100010111322-3313102010313113-3112021331102313-2133301202102100)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.default_security

<a id="canonical-0002310313133230-0130321111012030-0321330230131200-0113021220321322-3232010220210211-2231022210021231-0100022003121123-0113321320102320"></a>

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

<a id="canonical-0322002313121130-2112333323223303-1323311002133100-1221233003230020-1033203232212203-3120220000222203-1101230103320213-3320102102020200"></a>

## Direct properties — default_security / 123003033203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003322102313310-0023131221100121-3211330332303133-0030111300200130-0313113213212033-1320323303220322-1033023223012011-0121333221030202"></a>

## Next pages — default_security / 123003033203 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-023.md#canonical-2323123020222323-0201333323111230-1231221203230030-3330100031210233-0222100010111322-3313102010313113-3112021331102313-2133301202102100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0013101230002333-3003222322321331-3321101302000101-3022120330213133-1121322232002220-2110003203011020-0303100113130313-0313310011021322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233033202020211-0301233021113132-1031130310020000-0213222230132102-3112130223203220-3120202203323301-0320130213012133-0323110122132003"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.low_security — low_security / 331132011101 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-023.md#canonical-2323123020222323-0201333323111230-1231221203230030-3330100031210233-0222100010111322-3313102010313113-3112021331102313-2133301202102100)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.low_security

<a id="canonical-3222321300312100-1120000323330020-2223320310223023-1133322031002303-3101120132030131-0301002311013021-1221311320002313-2130233112000203"></a>

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

<a id="canonical-1312201312013022-0002222323232232-2111203113311200-0211221133010310-2312301200220232-2120113323002111-2201332021333013-1003313130302101"></a>

## Direct properties — low_security / 331132011101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102303213223023-1220103203332013-1020033032232112-1102032331030221-2300000322222122-3121000221033221-1012231130122033-3033121311002033"></a>

## Next pages — low_security / 331132011101 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-023.md#canonical-2323123020222323-0201333323111230-1231221203230030-3330100031210233-0222100010111322-3313102010313113-3112021331102313-2133301202102100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1132221310301033-2020203122100013-1132033323313312-1100211222313023-1132132110231233-2022123300000232-1013333100222200-2122220213002232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023103100232011-0210221002313303-2112301122211220-2131123130021202-0311212333211121-2232013031021032-3301330222231100-3010321003333121"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security — medium_security / 103001230300 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-023.md#canonical-2323123020222323-0201333323111230-1231221203230030-3330100031210233-0222100010111322-3313102010313113-3112021331102313-2133301202102100)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security

<a id="canonical-2322303232033110-0331133123331113-0002310003200122-3300001001321323-1003333001003031-1223021103101223-0200323003213103-1233022201110230"></a>

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

<a id="canonical-1233100131023012-1200132203310200-1333003332323110-1210330032103231-3112120103302130-3201230111103012-1333302310001102-3112120221221011"></a>

## Direct properties — medium_security / 103001230300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011112103233012-2223223330003103-0210312322002221-0200310123131022-1231000223203020-1302311330221121-0300121022021021-3031201322300310"></a>

## Next pages — medium_security / 103001230300 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-023.md#canonical-2323123020222323-0201333323111230-1231221203230030-3330100031210233-0222100010111322-3313102010313113-3112021331102313-2133301202102100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3332223322300330-0300332101012221-2000211222102303-3100312212012321-3332100132031130-1303221032333032-2111132330120012-0313030221023031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020322101310122-3131301321101011-3032101110001213-0023002330213013-3232030201232010-3103123102333012-0101313111202300-2301000322131023"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls — use_mtls / 003130121022 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls

<a id="canonical-2013203011233321-3112100231032021-2021302021001030-1131100012330022-0031012200000333-3113101022312020-3313013311302310-2033130102222312"></a>

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

<a id="canonical-1213100200312323-3030213022213202-0031123301103212-2022033032232331-2101022000333121-2203133020200021-0003222023021202-3211131133003320"></a>

## Direct properties — use_mtls / 003130121022 / 3

<a id="canonical-2113121210111313-2020022330100221-0201011302300123-2133113330021122-3300222331112011-2011223332102322-2021102011113121-3032332202012002"></a>

<a id="canonical-2212220300202213-1111333313010112-1211202101132010-3223000102020211-0011112030303001-2310300101132002-3233123000012101-2303210220222202"></a>

## client_certificate_optional property — use_mtls / 003130121022 / 4

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

- [crl](data-sources--workload--reference--group-023.md#canonical-1302302122322011-3012311330013021-1010012001113310-0220131221013133-0132023202212010-3100322120110023-2301001232313023-2203020303002103): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-023.md#canonical-1211331323202123-0202212232111213-0312020303322120-0331123232201022-0332221021010321-1202010101113031-1033300011021232-0321220331102113): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-023.md#canonical-3002113312132130-3301030011033002-0132201013202331-1301110110111011-0110120123200030-1200111132232201-2121132103321323-1000112101132303): complete subsection reference.

<a id="canonical-2200323012211320-2220333101103221-3332033200310022-2222312032013011-0033123133302002-0222133111220211-0313020320232310-3030120313120013"></a>

<a id="canonical-2101232222313311-0223022022211310-2133331121111302-2210323103300010-2102300002101100-0332312002111112-3320200012220220-1003113030230003"></a>

## trusted_ca_url property — use_mtls / 003130121022 / 5

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

- [xfcc_disabled](data-sources--workload--reference--group-023.md#canonical-0201023222022233-1211300301021112-3111211033001321-2322322310310201-2123020020232302-3220232322230330-2130212001213303-2202100301302203): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-023.md#canonical-1203033213002212-2010333113032100-2302131313333221-1120231213112300-3102012221222330-0332133232022231-1222020321302103-3303302102101302): complete subsection reference.

<a id="canonical-1113211103023221-2000003311003313-0322202332322213-1300202210332322-2021230001313231-1310022131313113-0220012030231132-1301213221131103"></a>

## Next pages — use_mtls / 003130121022 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.crl](data-sources--workload--reference--group-023.md#canonical-1302302122322011-3012311330013021-1010012001113310-0220131221013133-0132023202212010-3100322120110023-2301001232313023-2203020303002103)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl](data-sources--workload--reference--group-023.md#canonical-1211331323202123-0202212232111213-0312020303322120-0331123232201022-0332221021010321-1202010101113031-1033300011021232-0321220331102113)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca](data-sources--workload--reference--group-023.md#canonical-3002113312132130-3301030011033002-0132201013202331-1301110110111011-0110120123200030-1200111132232201-2121132103321323-1000112101132303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled](data-sources--workload--reference--group-023.md#canonical-0201023222022233-1211300301021112-3111211033001321-2322322310310201-2123020020232302-3220232322230330-2130212001213303-2202100301302203)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options](data-sources--workload--reference--group-023.md#canonical-1203033213002212-2010333113032100-2302131313333221-1120231213112300-3102012221222330-0332133232022231-1222020321302103-3303302102101302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1302302122322011-3012311330013021-1010012001113310-0220131221013133-0132023202212010-3100322120110023-2301001232313023-2203020303002103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112003322321132-1010023212301302-2302030132000220-2031030310323011-3302133232100322-3322303131033301-1202113320121330-1131121321320301"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.crl — crl / 111300023230 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-023.md#canonical-3332223322300330-0300332101012221-2000211222102303-3100312212012321-3332100132031130-1303221032333032-2111132330120012-0313030221023031)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.crl

<a id="canonical-1112121033312310-1011011222332021-3222012300311322-3031202011220313-1111123332023330-2301201011112210-2220112003100213-2300323203010102"></a>

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

<a id="canonical-3131011031332222-0302301033012012-2000201233102021-1302301000023220-0030010232213320-3231111203223312-0111320102000323-2033201011332331"></a>

## Direct properties — crl / 111300023230 / 3

<a id="canonical-1100203301230313-0022200132222120-0130212230113013-1101100013333310-2211213312332333-1101131200333332-0102010101010100-0023332012023101"></a>

<a id="canonical-3003313310121100-2103333112123033-1321011022220323-0011012030122131-0100223000312102-1023200321312011-3003320011103331-3223122020210300"></a>

## name property — crl / 111300023230 / 4

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

<a id="canonical-0230112130130332-1102013331302113-1132012121001300-1103001031311221-1003120313222110-3223023000211201-1230230300201033-3001300112211121"></a>

<a id="canonical-1111222003332232-2202032001213221-1302200302232203-1022333312013303-1021303332022121-1331023302112032-2003010300120100-2223003001031110"></a>

## namespace property — crl / 111300023230 / 5

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

<a id="canonical-1033110102232002-3231111003100110-1302011020303320-1021211113000223-0103212213331122-2020013110213203-2101111132110301-2123210201222003"></a>

<a id="canonical-3011301110003222-3310010102102301-0113221003203013-3321312210220121-1310330032100212-2300212113032310-2001213133112333-1222223232302230"></a>

## tenant property — crl / 111300023230 / 6

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

<a id="canonical-0303102200321232-3202230202110023-0122110132123213-3211210112313330-1200303102131312-3123301020022030-3313112301223111-2220233231121320"></a>

## Next pages — crl / 111300023230 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-023.md#canonical-3332223322300330-0300332101012221-2000211222102303-3100312212012321-3332100132031130-1303221032333032-2111132330120012-0313030221023031)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1211331323202123-0202212232111213-0312020303322120-0331123232201022-0332221021010321-1202010101113031-1033300011021232-0321220331102113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210000210131313-2333132333210031-1102032100011013-1000130101300301-3003210033121003-2323110321002230-3011102003321331-1332022112330201"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl — no_crl / 321103122233 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-023.md#canonical-3332223322300330-0300332101012221-2000211222102303-3100312212012321-3332100132031130-1303221032333032-2111132330120012-0313030221023031)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl

<a id="canonical-1111122031212202-3310001231123310-0203201003132332-0032112111031123-3220311213033130-0310203330131012-0332011121220301-1210012302102221"></a>

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

<a id="canonical-2000113302101303-3130103231213122-3000232031333211-1313012102010310-2311103221210103-3233132232112223-3233021031332032-3102321311220000"></a>

## Direct properties — no_crl / 321103122233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0020320332210003-2332032233220321-3330320112233022-3211230332233323-0311300101232231-3013010310100120-2022021332133322-3131113333213331"></a>

## Next pages — no_crl / 321103122233 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-023.md#canonical-3332223322300330-0300332101012221-2000211222102303-3100312212012321-3332100132031130-1303221032333032-2111132330120012-0313030221023031)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3002113312132130-3301030011033002-0132201013202331-1301110110111011-0110120123200030-1200111132232201-2121132103321323-1000112101132303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033222111010111-3120132112221111-0010100000222222-2202103130010311-0012013311023002-0210100133013100-0201120203231131-1211123132313111"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca — trusted_ca / 032202330311 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-023.md#canonical-3332223322300330-0300332101012221-2000211222102303-3100312212012321-3332100132031130-1303221032333032-2111132330120012-0313030221023031)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-0111033111133021-1301231212020310-1002100033122323-1323212131013002-0312320213233222-1201331210231000-3321213122003001-2131102121203101"></a>

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

<a id="canonical-2232001330113012-3213313200113222-0021330312122131-1120113223232110-0213203201200133-1320001120322102-1010330032210330-2012111231302103"></a>

## Direct properties — trusted_ca / 032202330311 / 3

<a id="canonical-1331100332321000-3100210213001131-1033022100332333-2320320232021133-0013211003023201-0103310320112131-3323022133023202-3021203020231230"></a>

<a id="canonical-1223310001330013-2101211001310130-3320212210310103-1323201331001212-3310320233201102-3332220303333113-1210111331121300-3220331332023313"></a>

## name property — trusted_ca / 032202330311 / 4

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

<a id="canonical-1323210122023113-0120300231331301-0331013333002212-0302010330333320-3120302322303110-0032330012131111-3012330233233313-3121310031322130"></a>

<a id="canonical-0233331301211213-2111122333110130-0200303103213010-2310200101011211-0310102002210302-0122333331133301-3232132033103301-0231013131233312"></a>

## namespace property — trusted_ca / 032202330311 / 5

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

<a id="canonical-3001203220101220-0111223122013130-0013101103113031-0321301203001223-0300032200012221-3022203133100201-3211132101320013-0131310221333101"></a>

<a id="canonical-3232131320312223-0301021012302000-3301131013023002-2300113001013301-0031200102022212-0231220033113312-3310100321131230-3122002333300233"></a>

## tenant property — trusted_ca / 032202330311 / 6

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

<a id="canonical-0332320203323202-3101232213230000-1310130311132103-0011123121111201-3211221321112013-0100322213031113-3020000232011212-3113132013300021"></a>

## Next pages — trusted_ca / 032202330311 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-023.md#canonical-3332223322300330-0300332101012221-2000211222102303-3100312212012321-3332100132031130-1303221032333032-2111132330120012-0313030221023031)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0201023222022233-1211300301021112-3111211033001321-2322322310310201-2123020020232302-3220232322230330-2130212001213303-2202100301302203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011223311322023-1332210201011222-2321110023113131-2210233123201011-1010333320122101-2311311021201301-3233120111031312-1210133333022210"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled — xfcc_disabled / 022200310132 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-023.md#canonical-3332223322300330-0300332101012221-2000211222102303-3100312212012321-3332100132031130-1303221032333032-2111132330120012-0313030221023031)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-0220233323103122-2133000130201112-0201230022033012-3221122110110231-0013300133000330-0233212202102210-1101120132021301-2110301130310313"></a>

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

<a id="canonical-2203111303232302-3313030013011203-1220210120123332-3120120311103030-2221030100300310-1211132302310310-1200023203121021-1131331123002301"></a>

## Direct properties — xfcc_disabled / 022200310132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001023221310032-1032113023121111-0303220020311020-1320311102312230-3233300111310100-0010302220123322-1131102020031300-0133121131212010"></a>

## Next pages — xfcc_disabled / 022200310132 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-023.md#canonical-3332223322300330-0300332101012221-2000211222102303-3100312212012321-3332100132031130-1303221032333032-2111132330120012-0313030221023031)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1203033213002212-2010333113032100-2302131313333221-1120231213112300-3102012221222330-0332133232022231-1222020321302103-3303302102101302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301210232321011-0211103310231100-1030220221203213-3310232220022103-1002130112221332-2133121002220333-1131220021220302-0322303202012010"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options — xfcc_options / 102132002302 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-023.md#canonical-3332223322300330-0300332101012221-2000211222102303-3100312212012321-3332100132031130-1303221032333032-2111132330120012-0313030221023031)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-0233303331230230-2131310020303012-2212233122013021-1312003320301220-3002023233322121-0201301003132323-1131321311222301-3100023103031311"></a>

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

<a id="canonical-2231323031230302-3023010323003320-0323010320313233-2203132303023133-2323022333320323-1211220323320211-3310102013031112-3111002301111110"></a>

## Direct properties — xfcc_options / 102132002302 / 3

<a id="canonical-1101010230301310-2302003030303200-0103130123210313-1320111103020322-1221100023200001-1221032003120022-1130111322012003-0120223302120131"></a>

<a id="canonical-2130120123130032-2322013221200031-0200033323131022-2122320033203203-1033221031231133-3202323000311201-0102031111202132-1121313021200313"></a>

## xfcc_header_elements property — xfcc_options / 102132002302 / 4

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

<a id="canonical-0203330023101221-3002211332210311-0212031221103101-3331210323113031-0123221002312330-0101303012212232-1112332113002310-2302313102130230"></a>

## Next pages — xfcc_options / 102132002302 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-023.md#canonical-3332223322300330-0300332101012221-2000211222102303-3100312212012321-3332100132031130-1303221032333032-2111132330120012-0313030221023031)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0200311133303313-1102333013020100-0231130033310130-0220330111033330-0233120333330111-0310003313220103-3012130211002322-2030012200210113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003331113310133-3102323033333303-1330111030220333-3302332303113322-3003113033210112-3220102323212130-2311020120210122-2322322122120202"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes — specific_routes / 113031130132 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes

<a id="canonical-2311121110310130-2310313003100321-3012021331111220-3223231021032102-3203301000232310-1220103111033231-3020210103302330-2131012101101111"></a>

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

<a id="canonical-0211300013133221-1202210301013022-1110310212333131-0111031233112102-2330213010333211-0333132320101201-2121220303330012-2111203321203103"></a>

## Direct properties — specific_routes / 113031130132 / 3

- [routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230): complete subsection reference.

<a id="canonical-0330103313103311-1123103010332002-1303132323103031-1203020323210103-0133030122200320-1301110111003212-1121133001220311-0031210023120032"></a>

## Next pages — specific_routes / 113031130132 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312201321320010-2110121301213101-0321030033303023-2022002100023312-2030032221223321-1001133233023000-1222123011202113-0033210130323000"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes — routes / 323013312133 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-0200311133303313-1102333013020100-0231130033310130-0220330111033330-0233120333330111-0310003313220103-3012130211002322-2030012200210113)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes

<a id="canonical-3300022001022001-2201133302301332-0130121213122130-0033012100000332-2233023133233211-2301122202102033-3011111221111231-1230230032011032"></a>

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

<a id="canonical-1323223030120020-2210021111103200-1130321323023020-3203212303313333-3101300103320121-1201302100201222-0131202102232033-3032021310001311"></a>

## Direct properties — routes / 323013312133 / 3

- [custom_route_object](data-sources--workload--reference--group-023.md#canonical-0123133211222112-3001221211131002-1032120001233233-2013213213112012-2201200011313023-3110233211212233-3232033300203013-0023201303021113): complete subsection reference.

- [direct_response_route](data-sources--workload--reference--group-023.md#canonical-2101110022312210-0030310322121123-1033101310330332-1312310032311100-3100032320110000-0021230210123302-0001022012103121-3022213313133301): complete subsection reference.

- [redirect_route](data-sources--workload--reference--group-023.md#canonical-1132000213131301-3210120212301310-3301022312202023-2233211123100020-0010000313311330-3133001202320010-3022300313000021-0233300001331101): complete subsection reference.

- [simple_route](data-sources--workload--reference--group-024.md#canonical-2222012210320200-3020020103010020-1310330110032302-3031321112003031-2302313332132112-0211011323231201-0210100133220232-0030303122321111): complete subsection reference.

<a id="canonical-2332323332122111-3101123221320303-0222022333200200-0131031012131112-3203111100310213-2013211311020030-1220133232203222-1003003333210312"></a>

## Next pages — routes / 323013312133 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-023.md#canonical-0123133211222112-3001221211131002-1032120001233233-2013213213112012-2201200011313023-3110233211212233-3232033300203013-0023201303021113)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-2101110022312210-0030310322121123-1033101310330332-1312310032311100-3100032320110000-0021230210123302-0001022012103121-3022213313133301)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-1132000213131301-3210120212301310-3301022312202023-2233211123100020-0010000313311330-3133001202320010-3022300313000021-0233300001331101)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-024.md#canonical-2222012210320200-3020020103010020-1310330110032302-3031321112003031-2302313332132112-0211011323231201-0210100133220232-0030303122321111)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-0200311133303313-1102333013020100-0231130033310130-0220330111033330-0233120333330111-0310003313220103-3012130211002322-2030012200210113)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0123133211222112-3001221211131002-1032120001233233-2013213213112012-2201200011313023-3110233211212233-3232033300203013-0023201303021113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010223012212223-1233221312321113-1332332230011033-1323213213003233-2311130121131112-0033032023221130-3232310312210002-3110000301210113"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object — custom_route_object / 232313020132 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-0200311133303313-1102333013020100-0231130033310130-0220330111033330-0233120333330111-0310003313220103-3012130211002322-2030012200210113)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object

<a id="canonical-0320000302101321-1213210222121211-0300031020033013-2120330112333100-1321233233231130-3120313033321001-1132302332333213-3011001132103131"></a>

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

<a id="canonical-2301312211133000-0300231333110323-1020323122132202-3030133013332310-3213031221123312-3203120201301110-1033220321112301-2331233023101303"></a>

## Direct properties — custom_route_object / 232313020132 / 3

- [caching_disable](data-sources--workload--reference--group-023.md#canonical-2230202300012122-1221311312323000-1133010301331112-3203002001303322-0023311103221211-2223122211112311-3022001132022123-2330230001322232): complete subsection reference.

- [caching_inherit](data-sources--workload--reference--group-023.md#canonical-1221011031033112-0101332131132022-3133200211101323-2011122323123302-3311201310113020-1111100301203001-0211223111321001-2000033010113021): complete subsection reference.

- [route_ref](data-sources--workload--reference--group-023.md#canonical-3102133132211010-1011033110013230-0013203132231213-0132311032331201-1332202121200303-1212012130102301-0021303223213022-3232211302000231): complete subsection reference.

<a id="canonical-0000103311033112-2122200310121333-3113213230202132-1300230022020331-0301100121011123-2213321303302231-1011201113310131-1322011231000121"></a>

## Next pages — custom_route_object / 232313020132 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable](data-sources--workload--reference--group-023.md#canonical-2230202300012122-1221311312323000-1133010301331112-3203002001303322-0023311103221211-2223122211112311-3022001132022123-2330230001322232)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit](data-sources--workload--reference--group-023.md#canonical-1221011031033112-0101332131132022-3133200211101323-2011122323123302-3311201310113020-1111100301203001-0211223111321001-2000033010113021)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref](data-sources--workload--reference--group-023.md#canonical-3102133132211010-1011033110013230-0013203132231213-0132311032331201-1332202121200303-1212012130102301-0021303223213022-3232211302000231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2230202300012122-1221311312323000-1133010301331112-3203002001303322-0023311103221211-2223122211112311-3022001132022123-2330230001322232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210012022102122-3012130310030311-0201331001012233-0121000221320220-0212311332102212-0012330000020211-3301220322101113-1200112030203033"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable — caching_disable / 002230231202 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-0200311133303313-1102333013020100-0231130033310130-0220330111033330-0233120333330111-0310003313220103-3012130211002322-2030012200210113)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-023.md#canonical-0123133211222112-3001221211131002-1032120001233233-2013213213112012-2201200011313023-3110233211212233-3232033300203013-0023201303021113)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

<a id="canonical-0221000113111101-0312110202323332-1113121311013131-3313001203031112-1302112233021233-0023301302111332-1002221000111201-2203212022202203"></a>

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

<a id="canonical-0230003122132033-1000113210111131-1020321332012232-0331211101120130-0132222300121103-2330122001331000-3212130112321003-0210030330002121"></a>

## Direct properties — caching_disable / 002230231202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1132210302201300-3302011033333130-1302000333312121-3032031323200031-1201233232202202-2021211013111110-2010221231302010-0230112120212030"></a>

## Next pages — caching_disable / 002230231202 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-023.md#canonical-0123133211222112-3001221211131002-1032120001233233-2013213213112012-2201200011313023-3110233211212233-3232033300203013-0023201303021113)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1221011031033112-0101332131132022-3133200211101323-2011122323123302-3311201310113020-1111100301203001-0211223111321001-2000033010113021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010031032203011-3130332110122030-2201030001012200-2031330231110000-2123132021121323-2230100202322001-3302120222113020-0223113230211033"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit — caching_inherit / 332212130321 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-0200311133303313-1102333013020100-0231130033310130-0220330111033330-0233120333330111-0310003313220103-3012130211002322-2030012200210113)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-023.md#canonical-0123133211222112-3001221211131002-1032120001233233-2013213213112012-2201200011313023-3110233211212233-3232033300203013-0023201303021113)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit

<a id="canonical-1330121220100122-3302311020303110-2221302103330211-3112031023232021-3032231320100233-3031310212100032-2200011113312121-1101213010312102"></a>

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

<a id="canonical-1013223200201100-1112213313113222-3111132121310132-3233210302203131-3012301120231231-2320222033111310-1000102131022310-0230032122311131"></a>

## Direct properties — caching_inherit / 332212130321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101332111233101-0103022323233233-0201321202031223-2111310233333311-2032233113221112-1201333030230101-2003230312201301-0103120212132031"></a>

## Next pages — caching_inherit / 332212130321 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-023.md#canonical-0123133211222112-3001221211131002-1032120001233233-2013213213112012-2201200011313023-3110233211212233-3232033300203013-0023201303021113)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3102133132211010-1011033110013230-0013203132231213-0132311032331201-1332202121200303-1212012130102301-0021303223213022-3232211302000231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011232322220331-2102322232310220-1122133132212311-1313221002233013-1113120210131123-0303212120100011-2213330112022321-3110220200220022"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref — route_ref / 010322131101 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-0200311133303313-1102333013020100-0231130033310130-0220330111033330-0233120333330111-0310003313220103-3012130211002322-2030012200210113)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-023.md#canonical-0123133211222112-3001221211131002-1032120001233233-2013213213112012-2201200011313023-3110233211212233-3232033300203013-0023201303021113)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref

<a id="canonical-2201213223123102-1313113031220001-3113111020333303-2302323010011233-0200110102231230-0111221330012202-3222322112021031-3311303323122130"></a>

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

<a id="canonical-1023331101311222-1211333120010100-3230121310131202-0212212202203132-2103122010123223-2122020101323330-3210231000022110-1012221320321221"></a>

## Direct properties — route_ref / 010322131101 / 3

<a id="canonical-0323032210022133-2020312033111202-1022031101030201-3232333130011310-0132201303100122-1232103030322032-0122231003011130-2120023332010323"></a>

<a id="canonical-1001311301101120-3012231033310022-0302303122200212-3122332331321112-3232333020220203-1022302021212100-0011033013001030-0011201120130212"></a>

## name property — route_ref / 010322131101 / 4

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

<a id="canonical-1310301121311300-0211311130122002-2200000330202301-3122011111131030-1032110101303033-2130331033023131-1321302223022120-2333023310000000"></a>

<a id="canonical-1302232311002103-1021301230023132-0103031210302202-1221202210320002-1003201332332030-2333033212012300-1111330302111022-3333110230103003"></a>

## namespace property — route_ref / 010322131101 / 5

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

<a id="canonical-1112203200011212-1122110332103323-0312223033303210-0200222133220120-2210221303113133-1012302113301221-2203321310132233-0122133223103222"></a>

<a id="canonical-1313200013213100-1011333300322112-3332300111110021-1000110113010111-1111232233133230-2013323032023302-0210310312301321-0231210112200003"></a>

## tenant property — route_ref / 010322131101 / 6

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

<a id="canonical-3201031230310322-3323011000200322-2332223020033201-2122213102010113-2221130202112123-1020022121230330-2001021321311001-1102330130232011"></a>

## Next pages — route_ref / 010322131101 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-023.md#canonical-0123133211222112-3001221211131002-1032120001233233-2013213213112012-2201200011313023-3110233211212233-3232033300203013-0023201303021113)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2101110022312210-0030310322121123-1033101310330332-1312310032311100-3100032320110000-0021230210123302-0001022012103121-3022213313133301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201100200330302-1100321211303013-3121201112223132-1233220023120130-3330320203121023-3022133323333233-3332223331220322-1313222101233321"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route — direct_response_route / 122100123030 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-0200311133303313-1102333013020100-0231130033310130-0220330111033330-0233120333330111-0310003313220103-3012130211002322-2030012200210113)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route

<a id="canonical-3032221030200133-0332201100032101-1033012201002101-3323212323010030-1030132213110132-3023221222021313-1112333022311001-3223200213202021"></a>

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

<a id="canonical-2112133313123203-1311101120232212-3310132131101332-2313022332113301-3203012313000112-1213132321333231-3303230100031112-3221320221113123"></a>

## Direct properties — direct_response_route / 122100123030 / 3

- [headers](data-sources--workload--reference--group-023.md#canonical-1221221112303030-3212001322131323-1220013303203113-2310312330130302-2322033333012230-3203120230230131-3310130002212022-2302201313100132): complete subsection reference.

<a id="canonical-2023013322330013-1012312302111100-0222210031320122-1222210111110233-0131232003313133-3202220101031200-0013301132221313-1322131233003011"></a>

<a id="canonical-0031013312100312-2100111032033123-3123201321020030-0112300021131332-0022002100020023-1232030022033312-2120300013120123-2233300303333123"></a>

## http_method property — direct_response_route / 122100123030 / 4

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

- [incoming_port](data-sources--workload--reference--group-023.md#canonical-2100013333130113-0302302131312031-1313121322130020-2200221320020231-2303300330312123-2330211132322331-0033321103202203-1320230013113232): complete subsection reference.

- [path](data-sources--workload--reference--group-023.md#canonical-2000221121000303-0331022212122032-1123230321120222-2132303323103102-3012112300102233-1013001221202303-3133202001313121-1212123200131121): complete subsection reference.

- [route_direct_response](data-sources--workload--reference--group-023.md#canonical-2023311233113121-0202031002222230-0223312311220002-1231321111301330-2110211030223003-3001303230220121-1112123300022010-2233033331031230): complete subsection reference.

<a id="canonical-1203121232001212-1030032202301312-2330332033013310-0112123230301100-2203022232023110-1220003333010103-1201232012132123-3200221132333130"></a>

## Next pages — direct_response_route / 122100123030 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers](data-sources--workload--reference--group-023.md#canonical-1221221112303030-3212001322131323-1220013303203113-2310312330130302-2322033333012230-3203120230230131-3310130002212022-2302201313100132)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](data-sources--workload--reference--group-023.md#canonical-2100013333130113-0302302131312031-1313121322130020-2200221320020231-2303300330312123-2330211132322331-0033321103202203-1320230013113232)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path](data-sources--workload--reference--group-023.md#canonical-2000221121000303-0331022212122032-1123230321120222-2132303323103102-3012112300102233-1013001221202303-3133202001313121-1212123200131121)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response](data-sources--workload--reference--group-023.md#canonical-2023311233113121-0202031002222230-0223312311220002-1231321111301330-2110211030223003-3001303230220121-1112123300022010-2233033331031230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1221221112303030-3212001322131323-1220013303203113-2310312330130302-2322033333012230-3203120230230131-3310130002212022-2302201313100132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030222231110320-2121112303032331-0211213100311312-1011302300012121-2333212321321230-2020211112313231-3301232131310122-2321100321021330"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers — headers / 322332120120 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-0200311133303313-1102333013020100-0231130033310130-0220330111033330-0233120333330111-0310003313220103-3012130211002322-2030012200210113)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-2101110022312210-0030310322121123-1033101310330332-1312310032311100-3100032320110000-0021230210123302-0001022012103121-3022213313133301)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers

<a id="canonical-2031121130120303-3003000103212230-3112100220200010-3030132001303112-3323233030030310-0202231020031301-0210110000003110-0020010021030313"></a>

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

<a id="canonical-3203100022311102-3132033013202103-3332130332131121-0311001231222231-3001022211201231-0110120200101320-2013003333001311-0013212023120331"></a>

## Direct properties — headers / 322332120120 / 3

<a id="canonical-1110312212123010-0310200030331020-3213220021202012-1331200131130111-0330233222320220-0202231300320321-2333123203322001-2110110012300132"></a>

<a id="canonical-2203102320011110-2101323220222221-2020202113013000-1221031333103222-2333001112100212-0121200033233320-0220313013223011-0101010313203123"></a>

## exact property — headers / 322332120120 / 4

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

<a id="canonical-0203122212221221-3021013000232300-0032001301013033-3133032102212300-0223203222112300-3130122030311130-3100011220112313-1101133310102213"></a>

<a id="canonical-2211122101211131-3232211002221233-1002233201122211-2303002312211223-1211111032313023-0332210121311102-1313331312123031-3113020232320313"></a>

## invert_match property — headers / 322332120120 / 5

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

<a id="canonical-0113313321030332-2231213201302322-3313013231231323-1011011002031331-2103000230213302-2030330220220033-3020230002122121-2233210023201013"></a>

<a id="canonical-0320022300133103-0212113121130321-1211120220132321-2310033323211011-2322112302133323-0212221100103220-1032212303210120-1312310223111122"></a>

## name property — headers / 322332120120 / 6

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

<a id="canonical-2320202332130011-0222023220110313-2231210002222133-0123023212313131-2331012012022313-3212120203100022-1200310103201302-3020111030300002"></a>

<a id="canonical-0232321012232222-2102021002100223-3013213202222020-0212102222233211-2230132230022001-2100221132331013-0320322230120210-2323130212201033"></a>

## presence property — headers / 322332120120 / 7

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

<a id="canonical-1111323110313223-3030003020102213-3231321332110103-1313100223230103-2130100020122030-2113133212230103-3022201313331121-3131010313202223"></a>

<a id="canonical-1223201113013101-3020120330330022-0033332121000221-1111000312023230-2131123122331303-2322102100301012-2113112231003003-2332030212022301"></a>

## regular expression property — headers / 322332120120 / 8

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

<a id="canonical-1203333200022302-0013011202112123-1130121321330022-3310331220213331-0012120022320232-1022122131230001-0312221131223310-1223302030213011"></a>

## Next pages — headers / 322332120120 / 9

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-2101110022312210-0030310322121123-1033101310330332-1312310032311100-3100032320110000-0021230210123302-0001022012103121-3022213313133301)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2100013333130113-0302302131312031-1313121322130020-2200221320020231-2303300330312123-2330211132322331-0033321103202203-1320230013113232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000223310132203-0032221021123010-1110212311003332-0302322111202000-1313330220322231-2332223212311223-3021023222230230-0112222211000021"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port — incoming_port / 203103102002 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-0200311133303313-1102333013020100-0231130033310130-0220330111033330-0233120333330111-0310003313220103-3012130211002322-2030012200210113)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-2101110022312210-0030310322121123-1033101310330332-1312310032311100-3100032320110000-0021230210123302-0001022012103121-3022213313133301)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port

<a id="canonical-0231011230132121-0003032223030031-1322102300112020-1301031310210311-2112020000323301-3222101210002011-3303031120123321-1212231111120322"></a>

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

<a id="canonical-2201033310123233-3332031001102111-3223120331323132-2101320013332230-3110330011113200-1013212112221023-1322303020200120-1330110230203123"></a>

## Direct properties — incoming_port / 203103102002 / 3

- [no_port_match](data-sources--workload--reference--group-023.md#canonical-2203132122332223-1331101320201211-3311323012301022-3202333103323232-3300220132320120-3023201310102113-0222011030211023-0131020132230100): complete subsection reference.

<a id="canonical-3333001112320121-1203021121000020-3120322030033331-1121122213110221-2121033333332321-3233201232330102-2132110202320221-2313200120330103"></a>

<a id="canonical-2311200100131030-1100002220322020-3132332221202212-2100110321302313-0010310102010222-0003011200200023-2002100030323030-3131003333030102"></a>

## port property — incoming_port / 203103102002 / 4

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

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

<a id="canonical-1121112003300320-3101123303321230-1201030021321212-2122112202202110-0330311021001202-2330030220102113-1122201221011133-0300230022323232"></a>

<a id="canonical-2230322010031330-2221322122332002-0102330303000211-3112101003331130-2210123103210321-3123310322322122-2223302212300310-1320211002002103"></a>

## port_ranges property — incoming_port / 203103102002 / 5

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

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

<a id="canonical-3330312101313211-3000112330110130-1303113001330232-3013103103101113-2132313103232303-3202331121032212-0320031322201203-0003122030231320"></a>

## Next pages — incoming_port / 203103102002 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match](data-sources--workload--reference--group-023.md#canonical-2203132122332223-1331101320201211-3311323012301022-3202333103323232-3300220132320120-3023201310102113-0222011030211023-0131020132230100)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-2101110022312210-0030310322121123-1033101310330332-1312310032311100-3100032320110000-0021230210123302-0001022012103121-3022213313133301)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2203132122332223-1331101320201211-3311323012301022-3202333103323232-3300220132320120-3023201310102113-0222011030211023-0131020132230100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300231213222230-0312030031300232-0111103022322300-2302300113232022-2231133212233310-0010220222233231-3102300332313203-0110111311323002"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match — no_port_match / 123321331003 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-0200311133303313-1102333013020100-0231130033310130-0220330111033330-0233120333330111-0310003313220103-3012130211002322-2030012200210113)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-2101110022312210-0030310322121123-1033101310330332-1312310032311100-3100032320110000-0021230210123302-0001022012103121-3022213313133301)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](data-sources--workload--reference--group-023.md#canonical-2100013333130113-0302302131312031-1313121322130020-2200221320020231-2303300330312123-2330211132322331-0033321103202203-1320230013113232)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match

<a id="canonical-2111010303100102-1211311111113033-2101313232100211-3211013212333230-1012012213102323-3311031013101123-2020230222010031-0322101323130211"></a>

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

<a id="canonical-1211200132333201-3000303000122101-3113333111203321-1030222113310010-2121103122130100-0113323122200123-1103203031221012-1212201022010211"></a>

## Direct properties — no_port_match / 123321331003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302121301322312-1322003333321210-2110220300102101-3323203331212132-2211222120310202-3033130203312300-0013011231120112-0120300212121111"></a>

## Next pages — no_port_match / 123321331003 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](data-sources--workload--reference--group-023.md#canonical-2100013333130113-0302302131312031-1313121322130020-2200221320020231-2303300330312123-2330211132322331-0033321103202203-1320230013113232)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2000221121000303-0331022212122032-1123230321120222-2132303323103102-3012112300102233-1013001221202303-3133202001313121-1212123200131121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303030220110313-2020302121033033-2002022231132012-0032320331320010-0010311002202112-3301031212033233-3103230303231211-1101300311031033"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path — path / 110103320311 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-0200311133303313-1102333013020100-0231130033310130-0220330111033330-0233120333330111-0310003313220103-3012130211002322-2030012200210113)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-2101110022312210-0030310322121123-1033101310330332-1312310032311100-3100032320110000-0021230210123302-0001022012103121-3022213313133301)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path

<a id="canonical-3231222100301212-1331102223200212-0021013220212113-1331030211000003-0101002302320210-0102323221213212-0111332131321022-2032120232313302"></a>

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

<a id="canonical-3100223132102200-2300211100013310-2003133323312112-0012123333133220-3322032323033031-0300313231301013-3212020022311123-0030333130333022"></a>

## Direct properties — path / 110103320311 / 3

<a id="canonical-1320132112121112-1321113030230330-3331310230222000-1111223020300010-3032121223013013-0300311301123102-3332012133030213-0310131112220112"></a>

<a id="canonical-1101310103222011-3301133220123112-1021031113302132-2020212120033301-0133202013110003-0000030023022100-3312121121023012-1221012102132113"></a>

## path property — path / 110103320311 / 4

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

<a id="canonical-2110202133021123-2012302210312002-2223002321311033-0123120102312200-2323012011331310-3000011320221011-3322032213002103-1331222313330201"></a>

<a id="canonical-1213330100133210-0202132212210333-2122200222313011-2031013202212121-3222332300210110-2010310120110113-1320300212333020-0002101121032112"></a>

## prefix property — path / 110103320311 / 5

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

<a id="canonical-0013010202111213-1211330001310233-2123001001103323-3230333110122123-1332100111202112-2031133312033223-2100321021303231-3223103220023303"></a>

<a id="canonical-0030033223003323-1012330003003120-0300031301231130-0103320332302310-1310202030123012-0110301311021111-1303332023230123-3322303102310120"></a>

## regular expression property — path / 110103320311 / 6

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

<a id="canonical-3101031101333213-1222003231320132-2022201223232103-3303021302231230-2323313321001000-1233313121313301-0122113213031303-0110203000131120"></a>

## Next pages — path / 110103320311 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-2101110022312210-0030310322121123-1033101310330332-1312310032311100-3100032320110000-0021230210123302-0001022012103121-3022213313133301)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2023311233113121-0202031002222230-0223312311220002-1231321111301330-2110211030223003-3001303230220121-1112123300022010-2233033331031230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232330332233332-3302203030322302-3213331002103031-1312301331322320-0001112221221300-2323131331133320-3010230131313103-2111302013023010"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response — route_direct_response / 210031023310 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-0200311133303313-1102333013020100-0231130033310130-0220330111033330-0233120333330111-0310003313220103-3012130211002322-2030012200210113)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-2101110022312210-0030310322121123-1033101310330332-1312310032311100-3100032320110000-0021230210123302-0001022012103121-3022213313133301)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

<a id="canonical-2330110330232012-3202232100031333-1301202310332033-3001112130333203-0020303203033110-2123121120233113-2332112102012000-1102130312112123"></a>

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

<a id="canonical-2322311302100221-2203100331033101-0312221311100321-3100302001022322-1233031002312032-2232310303101210-2233333133211032-0220331131330330"></a>

## Direct properties — route_direct_response / 210031023310 / 3

<a id="canonical-2310322012313312-1112231132233232-2322331322233102-2213331123001331-1121300132012031-3233312112332013-2300222112211000-1321202330221112"></a>

<a id="canonical-3031220130103233-1030033310223121-3201120323200020-1002122230310331-1022010030110220-1123210120212332-3132121223211210-2012211222320000"></a>

## response_body_encoded property — route_direct_response / 210031023310 / 4

Type: `"string"`. Computed.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in base64 format. The message can be either plain text or HTML.

Upstream description:

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
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0212323103001123-0112301013031201-0300200020311200-0321311033122001-1000033000300023-0111132133231203-1102203032011113-0130113112233202"></a>

<a id="canonical-3103122223200303-2100301113221102-0231300331221013-2231313203330123-3033212100103311-3200021103333300-0012112203112303-1311133300223331"></a>

## response_code property — route_direct_response / 210031023310 / 5

Type: `"number"`. Computed.

Response Code. Response code to send.

Upstream description:

Response code to send.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1102101111302302-1132213201332101-0131212100310022-1032020102202233-2223330222332221-1132030110303230-2030231133012220-1203013313102210"></a>

## Next pages — route_direct_response / 210031023310 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-023.md#canonical-2101110022312210-0030310322121123-1033101310330332-1312310032311100-3100032320110000-0021230210123302-0001022012103121-3022213313133301)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1132000213131301-3210120212301310-3301022312202023-2233211123100020-0010000313311330-3133001202320010-3022300313000021-0233300001331101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320211303232100-2020022332212101-0320320331111301-3221213103023032-3323212300203322-1222132011113202-1100210133200202-1103011133312101"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route — redirect_route / 203002230103 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-0200311133303313-1102333013020100-0231130033310130-0220330111033330-0233120333330111-0310003313220103-3012130211002322-2030012200210113)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-1110111020332302-2230121011011001-1300123223200023-2112032332001023-1121232020112022-1031111123333000-0230333322233101-1200001320131033"></a>

Type: `"single"`. Computed.

Redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the
matching traffic to a different URL.

Upstream description:

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

<a id="canonical-0133100011300001-3110133030110300-3033232232003013-2121323312312020-1133130002320333-1233200323221102-1021020233231312-3103020332330201"></a>

## Direct properties — redirect_route / 203002230103 / 3

- [headers](data-sources--workload--reference--group-023.md#canonical-1221022002020202-2130300033032210-2300231033031221-0033023211112000-0023200030311332-2312103332123221-3122103102023002-0011020203212020): complete subsection reference.

<a id="canonical-2113110223011132-2031211001300302-1121023121322212-1132300003113013-2301330023331101-1231212112033000-3301201311330200-0003201133301221"></a>

<a id="canonical-3121203221332300-1130201232331330-2322333110102203-1113113123211031-0033233310332221-2012121322023223-3213200230203030-1310301332020111"></a>

## http_method property — redirect_route / 203002230103 / 4

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

- [incoming_port](data-sources--workload--reference--group-024.md#canonical-3111310222120011-1312331330223212-3133120333322312-2233123210220020-0032013011232223-1333032020012112-1132100033233330-2222202321201012): complete subsection reference.

- [path](data-sources--workload--reference--group-024.md#canonical-2030021223211200-0220022322131110-3000001103231221-1102213111213203-0301100030020313-0022322101021100-3031031012212133-0023103012332120): complete subsection reference.

- [route_redirect](data-sources--workload--reference--group-024.md#canonical-2323211032212033-2302322032020303-1333300100001201-2220011220303100-3333113102013320-3013212233000332-3122202001112233-2233313303121232): complete subsection reference.

<a id="canonical-3010310003000133-2032003321121003-3330312032111003-3013122002333210-0011321102103203-3323202313120131-1332133301213311-1011302223130013"></a>

## Next pages — redirect_route / 203002230103 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers](data-sources--workload--reference--group-023.md#canonical-1221022002020202-2130300033032210-2300231033031221-0033023211112000-0023200030311332-2312103332123221-3122103102023002-0011020203212020)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-024.md#canonical-3111310222120011-1312331330223212-3133120333322312-2233123210220020-0032013011232223-1333032020012112-1132100033233330-2222202321201012)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path](data-sources--workload--reference--group-024.md#canonical-2030021223211200-0220022322131110-3000001103231221-1102213111213203-0301100030020313-0022322101021100-3031031012212133-0023103012332120)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-024.md#canonical-2323211032212033-2302322032020303-1333300100001201-2220011220303100-3333113102013320-3013212233000332-3122202001112233-2233313303121232)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1221022002020202-2130300033032210-2300231033031221-0033023211112000-0023200030311332-2312103332123221-3122103102023002-0011020203212020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
