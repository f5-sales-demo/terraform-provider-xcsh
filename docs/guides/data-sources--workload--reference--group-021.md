---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-2010003021032221-3230103000202223-1231022313202111-0232030203130200-0012002310002030-3000312330323221-0331322122211000-1210013223321221"></a>

## Next pages — http_loadbalancer / 122030121122 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-021.md#canonical-2030310202300112-3303300032130110-0103103322303032-0211102122210322-0023002233231012-2302312031202003-1110220232133120-3323320120132322)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http](data-sources--workload--reference--group-021.md#canonical-0112310331012222-3310301101023101-3030311303122110-1002120000131211-2132222333033211-0221132323013310-2331200022030010-1001111302103012)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-023.md#canonical-0200311133303313-1102333013020100-0231130033310130-0220330111033330-0233120333330111-0310003313220103-3012130211002322-2030012200210113)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2030310202300112-3303300032130110-0103103322303032-0211102122210322-0023002233231012-2302312031202003-1110220232133120-3323320120132322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003021330332012-3010011320312012-2213033033220302-2330111300130122-0303012223112322-1302021031210031-0111113200320301-2100213233120331"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route — default_route / 301322230323 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route

<a id="canonical-2321021000121222-3213213111123223-1021111301331300-2312222301003103-1230000233131301-1301233122221120-2203123221221202-2023013110203022"></a>

Type: `"single"`. Computed.

Configuration parameter for default route.

Upstream description:

Default route matching all APIs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

<a id="canonical-0132203222332321-2110301111010333-2231213001132213-1323002003021220-2020220213023212-3212022320300130-1321013311202230-0232230000321121"></a>

## Direct properties — default_route / 301322230323 / 3

- [auto_host_rewrite](data-sources--workload--reference--group-021.md#canonical-2312133100231010-1132023310002323-2221211231002103-2232301033231112-1230121030131330-3203001022103000-0012103030331233-3332211021001220): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-021.md#canonical-2321022330333131-1333102212121100-3211320232032223-3310311201013132-2233202021220201-1222002100310323-1311203223102201-3100031132310112): complete subsection reference.

<a id="canonical-2233221011211332-1310130333332232-0303132322121203-3233022220222121-2233310100221012-1201033130001010-1201131022013200-2212110010332332"></a>

<a id="canonical-0210300030200230-0112320311113230-1232112020203223-2311212310313232-1200200102322030-1302330213323301-0231220202222222-0310030010313232"></a>

## host_rewrite property — default_route / 301322230323 / 4

Type: `"string"`. Computed.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-3320222211220210-1011220301221001-3332203132220113-2003113213021201-1310132001120101-1120022111300032-0131000022020212-3310110201113330"></a>

## Next pages — default_route / 301322230323 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite](data-sources--workload--reference--group-021.md#canonical-2312133100231010-1132023310002323-2221211231002103-2232301033231112-1230121030131330-3203001022103000-0012103030331233-3332211021001220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite](data-sources--workload--reference--group-021.md#canonical-2321022330333131-1333102212121100-3211320232032223-3310311201013132-2233202021220201-1222002100310323-1311203223102201-3100031132310112)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2312133100231010-1132023310002323-2221211231002103-2232301033231112-1230121030131330-3203001022103000-0012103030331233-3332211021001220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030111101122212-1211210110332203-0010233012011222-2202031023322121-1023030210232233-0211130202332233-1102203002212221-2113211131132312"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite — auto_host_rewrite / 020320011321 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-021.md#canonical-2030310202300112-3303300032130110-0103103322303032-0211102122210322-0023002233231012-2302312031202003-1110220232133120-3323320120132322)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-3030110323132322-2200113011321301-1130032122031013-0102303322212233-3112013313203022-2021102322123312-2312102013303100-3022201030320133"></a>

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

<a id="canonical-2033120222313132-1310021120121130-1203203022313220-2113220113110002-1002110300113213-0233220212112030-1202002203113013-2003031203103203"></a>

## Direct properties — auto_host_rewrite / 020320011321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303221232010122-0200101310021230-0311320110013210-3212333212302013-3010101000012022-3331122201231132-0233120323231131-3121110131102030"></a>

## Next pages — auto_host_rewrite / 020320011321 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-021.md#canonical-2030310202300112-3303300032130110-0103103322303032-0211102122210322-0023002233231012-2302312031202003-1110220232133120-3323320120132322)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2321022330333131-1333102212121100-3211320232032223-3310311201013132-2233202021220201-1222002100310323-1311203223102201-3100031132310112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330020231030003-0101203220030320-2112221120321233-0020201311123022-1101311012030302-2022221212223003-0111001203100002-0103120232322333"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite — disable_host_rewrite / 202000302303 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-021.md#canonical-2030310202300112-3303300032130110-0103103322303032-0211102122210322-0023002233231012-2302312031202003-1110220232133120-3323320120132322)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-1332303113222011-0210122330122231-2200302001321222-3133033010320232-3202112122120023-2112213131302133-2323330223131103-2122111230333320"></a>

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

<a id="canonical-3133012000331230-1032332100023131-2202100101323330-2201311331202032-0010311023100101-3303331313112123-1133232123331100-3311300330211013"></a>

## Direct properties — disable_host_rewrite / 202000302303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001032031202010-3123013113131001-1103210210002032-2300310320033133-2201233320312303-3213202203221011-0032312022301321-2230121113123320"></a>

## Next pages — disable_host_rewrite / 202000302303 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-021.md#canonical-2030310202300112-3303300032130110-0103103322303032-0211102122210322-0023002233231012-2302312031202003-1110220232133120-3323320120132322)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0112310331012222-3310301101023101-3030311303122110-1002120000131211-2132222333033211-0221132323013310-2331200022030010-1001111302103012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201221012221123-0120220032220033-3203223301032021-1202121000012111-0313021303130131-3021123030032023-2221220112032132-0311333201132021"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http — http / 020123123112 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http

<a id="canonical-3130301310003333-2033002331303022-3213301223311103-1021123220112320-1330000132102231-1020000220003221-0200222021022330-0300003103130113"></a>

Type: `"single"`. Computed.

HTTP Choice. Choice for selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

<a id="canonical-0330221330123120-1323300100133001-1222012011103201-0002310303221301-2030021220200011-2221100232302232-2021311102301203-2311122232323211"></a>

## Direct properties — http / 020123123112 / 3

<a id="canonical-0001303201003022-1130301031232000-3222311021130311-0110321003332331-0222103110230321-2303220002003212-0122322033330300-1001302332100210"></a>

<a id="canonical-0002310320302101-3332122323320132-3030211200311133-3332002232203131-2323322002103233-2010313232323201-2120020212010021-0123320320232230"></a>

## dns_volterra_managed property — http / 020123123112 / 4

Type: `"bool"`. Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Upstream description:

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2323121110323320-1222201331003321-2120310201323031-2222311322313201-3203303102132220-3320203003103121-2002021101123312-1332322202231113"></a>

<a id="canonical-1123112120302000-0023223010331313-0321130321202123-0103310302310111-1032032203333001-2101011011331121-2313003122222122-0032002222203200"></a>

## port property — http / 020123123112 / 5

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1033322120122221-3330132312131032-0311131011210211-3103000001313312-0003232211100331-2130223132320223-0211323323112022-0230210322110100"></a>

<a id="canonical-0102001332323010-3021021101321212-2310231221120100-0231322133220033-2012113111102123-0302222231233110-0310333103130223-1222033132331212"></a>

## port_ranges property — http / 020123123112 / 6

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

<a id="canonical-2311033030002012-3231101203310321-3002322202312223-3231132301220322-1332002203220220-1320010032113310-0213323200023212-3232112203323210"></a>

## Next pages — http / 020123123112 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320120133301203-2230032020103133-2132220230231000-0020232102322310-2223220220333110-2103010311130311-1103233222221001-2212222133203302"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https — https / 010313211311 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https

<a id="canonical-1123030011002223-0102231300311003-2121230032211023-0010032001111100-3120331310120332-1312322120021013-2321310103223130-1121000223210103"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

<a id="canonical-2230200031003011-1230032031223211-0031001332223212-3201031202311030-2200030311333102-2000003021003021-3202333330321212-2032323002000130"></a>

## Direct properties — https / 010313211311 / 3

<a id="canonical-2311233231201300-1000303020202212-1023202130000221-0113122023120123-3022102123012130-0311102321303113-1033010321331330-1301312032212211"></a>

<a id="canonical-3100031133033322-0210230110121312-3102130331302222-1203001030231302-3021123213113232-3222200111213003-3312232310301313-3123230301201210"></a>

## add_hsts property — https / 010313211311 / 4

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

<a id="canonical-0202301330000230-0312333130312320-0023233220303302-1210001011030133-3121203132222132-1033231010003121-2131023010001310-0100030333313023"></a>

<a id="canonical-0301103231103100-0310022111010112-3013110312330113-3020132021100310-1110020011301102-1010201032110032-0010200313012330-2102130031102133"></a>

## append_server_name property — https / 010313211311 / 5

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](data-sources--workload--reference--group-021.md#canonical-0223013021320231-2210210321322111-3013223300203101-1323100321101202-1032021203311322-2030033033323022-1130010310100330-0302011323001003): complete subsection reference.

<a id="canonical-2010031033203322-0211222210300130-2321111010300113-2210121100321121-3021133122001111-1303103210200300-1121022231030103-3223221112012303"></a>

<a id="canonical-1200102110120221-0023130221310211-2233100121223132-2000222302121331-0002200203212202-0221123133113230-2122201122132220-2301120200303130"></a>

## connection_idle_timeout property — https / 010313211311 / 6

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

- [default_header](data-sources--workload--reference--group-021.md#canonical-3122021022021130-2031201222023001-0203131103201023-1223221202213312-3323012032021310-3300120210220111-3321210123213333-2303131033232212): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-021.md#canonical-3312132213003110-3220223213233203-1013321120233203-0131000211303220-2233231131201013-1011202203021310-3331211003001131-2303221001021310): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-021.md#canonical-0311110212032330-3222120112023023-3321032333013030-3103101303022011-0120222102333201-3200113323301112-2303001031323212-2120120221322321): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-021.md#canonical-0230203033333220-1103323303221313-1132312203222022-3311332312300330-0122230203313031-2330333031222120-0132110230101001-3022030003101312): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-021.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310): complete subsection reference.

<a id="canonical-0033330330222110-3132231321331320-2123321111020322-1113110032010320-0023301333221013-2212120121002303-0030121013323330-3230303011133100"></a>

<a id="canonical-2131033011030031-0220130121210320-1133033200223000-0133222122323000-0333303103230320-2321023210102322-3301200012031333-2032000011330101"></a>

## http_redirect property — https / 010313211311 / 7

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

- [non_default_loadbalancer](data-sources--workload--reference--group-021.md#canonical-0011033233320121-3333020002121100-1201121003022320-2303002220331022-3023331020321233-0320223300120033-3021203022022003-2131002203233322): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-021.md#canonical-0000223200231120-3103130233313330-2132331313003321-2203211112112303-0302010122020301-3031301310323333-3132033122212331-3022302332313302): complete subsection reference.

<a id="canonical-1222321232300231-1120312023030321-1333101233313322-1310231303130313-3312202311023032-0102221133121303-1123010110333230-1121203012220203"></a>

<a id="canonical-3302121120102003-1032223112120312-0010110021100221-0232300313033313-2200123222012102-0220331032110021-1220010300313133-3003111213303003"></a>

## port property — https / 010313211311 / 8

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3001010111013311-1103303132000023-3210012122013022-1011331010012130-1100302123301331-2113021331031310-0333120030110320-2301213133313322"></a>

<a id="canonical-3133311100120312-1000112013232113-1321321112321201-1120000232201222-1013331012210223-3221021303032100-2322013123323211-0030203321022110"></a>

## port_ranges property — https / 010313211311 / 9

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

<a id="canonical-0313312001100203-3131001232213013-0032322102321022-0112233332232300-0133121030101311-3100030211311322-3220303223322231-3310301332130222"></a>

<a id="canonical-1103210200110231-0023323233120003-1000233113310131-2011101032301221-2311203101103333-1012113120100133-0100002011033230-1122112113123033"></a>

## server_name property — https / 010313211311 / 10

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313): complete subsection reference.

- [tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303): complete subsection reference.

<a id="canonical-1110030301221003-0003121322223011-2102133311210311-1300333232012012-3203221033320021-0333322230001213-1023312330102230-3312301032032302"></a>

## Next pages — https / 010313211311 / 11

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-021.md#canonical-0223013021320231-2210210321322111-3013223300203101-1323100321101202-1032021203311322-2030033033323022-1130010310100330-0302011323001003)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header](data-sources--workload--reference--group-021.md#canonical-3122021022021130-2031201222023001-0203131103201023-1223221202213312-3323012032021310-3300120210220111-3321210123213333-2303131033232212)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer](data-sources--workload--reference--group-021.md#canonical-3312132213003110-3220223213233203-1013321120233203-0131000211303220-2233231131201013-1011202203021310-3331211003001131-2303221001021310)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize](data-sources--workload--reference--group-021.md#canonical-0311110212032330-3222120112023023-3321032333013030-3103101303022011-0120222102333201-3200113323301112-2303001031323212-2120120221322321)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize](data-sources--workload--reference--group-021.md#canonical-0230203033333220-1103323303221313-1132312203222022-3311332312300330-0122230203313031-2330333031222120-0132110230101001-3022030003101312)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-021.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer](data-sources--workload--reference--group-021.md#canonical-0011033233320121-3333020002121100-1201121003022320-2303002220331022-3023331020321233-0320223300120033-3021203022022003-2131002203233322)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through](data-sources--workload--reference--group-021.md#canonical-0000223200231120-3103130233313330-2132331313003321-2203211112112303-0302010122020301-3031301310323333-3132033122212331-3022302332313302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0223013021320231-2210210321322111-3013223300203101-1323100321101202-1032021203311322-2030033033323022-1130010310100330-0302011323001003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200132110103332-2131320210110001-0010321111113100-3200010021222331-1030211221032131-2303322212101130-1112203321231212-2120021200322130"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options — coalescing_options / 103230131201 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-1203020012302030-0001013011021012-0113002313023321-2332002311302210-2313233303021300-3203211012313303-0013312310011230-0221332120322010"></a>

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

<a id="canonical-3200132023133223-1300111221202320-2032221010231133-0021230203321203-0100012001312201-1313221202103330-0033123233020311-2120103212211001"></a>

## Direct properties — coalescing_options / 103230131201 / 3

- [default_coalescing](data-sources--workload--reference--group-021.md#canonical-1222213323120211-3133013131321230-3103230031310110-1230220133202132-1121323100033323-1022200022001013-3231132033010103-0330112230023020): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-021.md#canonical-1021011213302123-2123303232103013-3021200330320233-1031313321012003-2310000001130323-3230012231310302-1032123110132323-0131220032121320): complete subsection reference.

<a id="canonical-0320011200100221-1233023220100233-0312111013103201-1021210233223230-3000220223213020-0010332223133303-2310123223121231-0332201010022030"></a>

## Next pages — coalescing_options / 103230131201 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing](data-sources--workload--reference--group-021.md#canonical-1222213323120211-3133013131321230-3103230031310110-1230220133202132-1121323100033323-1022200022001013-3231132033010103-0330112230023020)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing](data-sources--workload--reference--group-021.md#canonical-1021011213302123-2123303232103013-3021200330320233-1031313321012003-2310000001130323-3230012231310302-1032123110132323-0131220032121320)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1222213323120211-3133013131321230-3103230031310110-1230220133202132-1121323100033323-1022200022001013-3231132033010103-0330112230023020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112023010222300-2211203210232120-0110100003003000-2113232031122032-1101223131011330-0023033022020230-0322031111120320-0003200112101110"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing — default_coalescing / 302301202011 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-021.md#canonical-0223013021320231-2210210321322111-3013223300203101-1323100321101202-1032021203311322-2030033033323022-1130010310100330-0302011323001003)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-3111201011232022-3223222131210302-1112212313230333-3201031122003313-1220213313001203-3330202203312301-1210133321132220-2321201210302213"></a>

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

<a id="canonical-0300030013302003-2033131202121111-2032213123132323-1330210031210332-2212003112313113-2321332213201023-3311331001333131-1210221301331213"></a>

## Direct properties — default_coalescing / 302301202011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301200330203323-3121121033203130-3321211122322131-3211113132203203-3111211230131130-2011122101130321-2130110303320213-2111011133222023"></a>

## Next pages — default_coalescing / 302301202011 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-021.md#canonical-0223013021320231-2210210321322111-3013223300203101-1323100321101202-1032021203311322-2030033033323022-1130010310100330-0302011323001003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1021011213302123-2123303232103013-3021200330320233-1031313321012003-2310000001130323-3230012231310302-1032123110132323-0131220032121320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321332023203121-2113322232003302-0133112203110121-2021120032333232-1320030021023003-0130310312211030-1032013122212010-3320030311311313"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing — strict_coalescing / 233110230003 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-021.md#canonical-0223013021320231-2210210321322111-3013223300203101-1323100321101202-1032021203311322-2030033033323022-1130010310100330-0302011323001003)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-3110202200031103-0303321121121022-1300021000113233-1323203020211231-2230010121111002-2103330120113321-1111222121111210-1333003100132202"></a>

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

<a id="canonical-3333202031201102-2002331213111323-3331011032102231-0311002322333003-0213221301001210-0221232312110310-0323303013101102-2110032021202200"></a>

## Direct properties — strict_coalescing / 233110230003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201021233100300-2031220022031303-2023311130013131-0331210120231011-2110122133110332-2113302113312013-0312213030211221-1120132000302001"></a>

## Next pages — strict_coalescing / 233110230003 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-021.md#canonical-0223013021320231-2210210321322111-3013223300203101-1323100321101202-1032021203311322-2030033033323022-1130010310100330-0302011323001003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3122021022021130-2031201222023001-0203131103201023-1223221202213312-3323012032021310-3300120210220111-3321210123213333-2303131033232212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131021230003120-3313230102120231-2011213313021111-3112001213132000-0020332001233310-3120130222002332-0211101231033333-0100320020133221"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header — default_header / 120213322012 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header

<a id="canonical-2311202111003331-2112022010232200-3230203111332032-2033101203233231-2011322133230323-0033323033310120-3103020312013303-2102221202121310"></a>

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

<a id="canonical-2033102032113221-1223221322221331-1031121302013011-3011131231312102-1302333120101231-3131113200022101-2123012122302200-0000111302130203"></a>

## Direct properties — default_header / 120213322012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2013013003221120-2123322123221010-2233210200313223-1311031211222232-2122330001232103-3111102112030002-3330330230130310-1232231101033332"></a>

## Next pages — default_header / 120213322012 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3312132213003110-3220223213233203-1013321120233203-0131000211303220-2233231131201013-1011202203021310-3331211003001131-2303221001021310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333322220200103-3232132130300122-0231100211231020-2312121333102322-2031102002023023-1331112021233310-0221222031300312-2003002313010303"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer — default_loadbalancer / 212212010230 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer

<a id="canonical-2011020033212012-3310030032302102-2031213101201230-3023132322023313-0323203011232230-1012001201330210-2311232333211201-3301100212110211"></a>

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

<a id="canonical-2022310220031101-0021311003123130-0113333322233223-2311133320200130-2200200021221023-2300212333123100-3021211302301112-1030003130100000"></a>

## Direct properties — default_loadbalancer / 212212010230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111200331033230-2031232031112030-2310311210323203-3111033002331223-3233330032322303-2310103000110333-3032333202212223-0133220310311031"></a>

## Next pages — default_loadbalancer / 212212010230 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0311110212032330-3222120112023023-3321032333013030-3103101303022011-0120222102333201-3200113323301112-2303001031323212-2120120221322321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003121111033123-3212032103013112-0003303200001110-1223323313000130-2231013302221011-2303200133222231-2310311230210021-2313212133113131"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize — disable_path_normalize / 201102010200 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize

<a id="canonical-3113201230223310-1101010202313122-1010010230221211-2312020310013210-3322331001110033-1330220320021001-2300333002031121-3011202101331023"></a>

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

<a id="canonical-3213100013132013-1113131100310303-1032213032223130-2020130233211313-0002021020110131-3213310221020202-0032030121130323-3000121203222002"></a>

## Direct properties — disable_path_normalize / 201102010200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031033203223313-3111021113132001-2123323323002023-0022011323232031-0021122220012101-1120211131031031-2222103302323030-0131100331013320"></a>

## Next pages — disable_path_normalize / 201102010200 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0230203033333220-1103323303221313-1132312203222022-3311332312300330-0122230203313031-2330333031222120-0132110230101001-3022030003101312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112103210313031-3112313033213132-0032221101000000-0311013222300300-0232310300032221-2303112301232203-2101231000303222-1021210222211122"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize — enable_path_normalize / 001002331113 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize

<a id="canonical-1321333120211302-1323031332311223-3230210330131210-2131300111031201-0123101302311020-1123131301023012-0212133113010013-2330022132002312"></a>

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

<a id="canonical-0330102021223333-0113233111230300-1223003320333311-0133330032313322-2132302002313220-3330230021131310-1213312113213310-2230311332301210"></a>

## Direct properties — enable_path_normalize / 001002331113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211200110230101-3123110220320033-1000010333233110-1032231310321301-0222010322202321-2222122232213023-3133113111303212-2122001013121121"></a>

## Next pages — enable_path_normalize / 001002331113 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220310033003200-1312203231313011-1132213301013021-0103013213313301-1023103113320000-0123110001101011-2022022233202313-2331220213030322"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options — http_protocol_options / 331001332211 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options

<a id="canonical-2233212223322132-2203202033301320-2300130223012101-2230122312222321-3031002102022001-1213010020121300-3303113230232320-0133031200021022"></a>

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

<a id="canonical-3000201203201103-2230311112312200-0033331000010013-3132002033331101-0121320323303123-1202232033011212-2230033113210013-3311323201233222"></a>

## Direct properties — http_protocol_options / 331001332211 / 3

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-021.md#canonical-1033131203023032-2102001223320103-0323131200122322-2301231132233112-1223111332323113-3231022031110001-1223033201000032-3310113300131012): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-021.md#canonical-2203131003213122-3331030200230003-1001023323300301-1023203312002120-2112303110321012-2131131303320221-3032033312210013-2023033121200003): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-021.md#canonical-2330131301031201-1210032120330101-1321113003123120-0103221211320010-1333112300210212-1220220331310110-2030120112022001-2103332101212322): complete subsection reference.

<a id="canonical-3100300312101230-0133223111123222-2130312331231220-2021131310223303-0032223002102230-3132112011032000-1203302210330110-3310320331123302"></a>

## Next pages — http_protocol_options / 331001332211 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-021.md#canonical-1033131203023032-2102001223320103-0323131200122322-2301231132233112-1223111332323113-3231022031110001-1223033201000032-3310113300131012)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](data-sources--workload--reference--group-021.md#canonical-2203131003213122-3331030200230003-1001023323300301-1023203312002120-2112303110321012-2131131303320221-3032033312210013-2023033121200003)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](data-sources--workload--reference--group-021.md#canonical-2330131301031201-1210032120330101-1321113003123120-0103221211320010-1333112300210212-1220220331310110-2030120112022001-2103332101212322)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1033131203023032-2102001223320103-0323131200122322-2301231132233112-1223111332323113-3231022031110001-1223033201000032-3310113300131012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110312012230003-1023123011112210-3120220302132332-2022013220101310-1302033302300002-0103112011033103-3311223031102300-3221302210101222"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 003303010313 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-021.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-1122311230210313-3100223113201021-1210201003033133-1211220033201102-0233122111221133-0131130333103323-0001120332221133-0013222201103333"></a>

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

<a id="canonical-1321031301320203-2200112021203232-3100333101103121-2201101103002012-1223221330332300-1122010323202313-0022232203212021-2032023102302311"></a>

## Direct properties — http_protocol_enable_v1_only / 003303010313 / 3

- [header_transformation](data-sources--workload--reference--group-021.md#canonical-1000301122202312-3011230331001222-1103133100213100-0100322330231321-3210311100100323-1222032203300032-0300300200303201-1102223101130200): complete subsection reference.

<a id="canonical-2001331030010022-2211232021202321-3000221303300110-0332220021323211-3130023111111320-1201100222211222-1012233013101333-2323002132321333"></a>

## Next pages — http_protocol_enable_v1_only / 003303010313 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-021.md#canonical-1000301122202312-3011230331001222-1103133100213100-0100322330231321-3210311100100323-1222032203300032-0300300200303201-1102223101130200)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-021.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1000301122202312-3011230331001222-1103133100213100-0100322330231321-3210311100100323-1222032203300032-0300300200303201-1102223101130200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300213131233303-1200332103201230-0320333300010331-0323021021202133-3100000300211002-0122021132110000-0120302002320022-1333130011310310"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 031300113311 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-021.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-021.md#canonical-1033131203023032-2102001223320103-0323131200122322-2301231132233112-1223111332323113-3231022031110001-1223033201000032-3310113300131012)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-2130030310302323-1213012112303232-1100121133013132-2001131232012211-0011012131222110-1323231303133030-3311302312001121-3302232131102000"></a>

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

<a id="canonical-0331120113003232-2312213110332030-2132221122223023-3112021233121100-1231021001222131-1233220202312100-2122311301323203-0103033210312121"></a>

## Direct properties — header_transformation / 031300113311 / 3

- [default_header_transformation](data-sources--workload--reference--group-021.md#canonical-2013002213000101-1330333123203131-3300111133110113-3333110213211322-0013302210300232-3202020310201112-3001233321300232-1100302301212021): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-021.md#canonical-1200211130220221-3133230021110123-3003231020332330-0213130301203020-2300203200121212-0301202202201310-3321211010321203-2030023311211121): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-021.md#canonical-2112201210020123-2221033111203013-3202202012112111-0123232323311222-0013103312310011-1103112023303220-0201102213220121-0220331221020210): complete subsection reference.

<a id="canonical-1233011123133010-2111011222213311-0302313210003020-2203331320032233-2031232210113203-0031101023310103-2331221030332033-1322332101333300"></a>

## Next pages — header_transformation / 031300113311 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--workload--reference--group-021.md#canonical-2013002213000101-1330333123203131-3300111133110113-3333110213211322-0013302210300232-3202020310201112-3001233321300232-1100302301212021)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--workload--reference--group-021.md#canonical-1200211130220221-3133230021110123-3003231020332330-0213130301203020-2300203200121212-0301202202201310-3321211010321203-2030023311211121)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--workload--reference--group-021.md#canonical-2112201210020123-2221033111203013-3202202012112111-0123232323311222-0013103312310011-1103112023303220-0201102213220121-0220331221020210)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-021.md#canonical-1033131203023032-2102001223320103-0323131200122322-2301231132233112-1223111332323113-3231022031110001-1223033201000032-3310113300131012)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2013002213000101-1330333123203131-3300111133110113-3333110213211322-0013302210300232-3202020310201112-3001233321300232-1100302301212021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302302323012100-3310001213103310-1100200303200023-0023320203223322-0312013131030221-0110201221332132-2122200230132121-1333031022211301"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 003001101103 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-021.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-021.md#canonical-1033131203023032-2102001223320103-0323131200122322-2301231132233112-1223111332323113-3231022031110001-1223033201000032-3310113300131012)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-021.md#canonical-1000301122202312-3011230331001222-1103133100213100-0100322330231321-3210311100100323-1222032203300032-0300300200303201-1102223101130200)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-1032131320333323-2211001000012001-2322031113022232-1200033201022102-0211133031002332-0100301201131202-3020031012223100-3031032202212222"></a>

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

<a id="canonical-1221030222330222-2110320022212003-2222131211331033-3322010111000020-0013201120231301-3312223103101210-2020032033002103-1002103131000032"></a>

## Direct properties — default_header_transformation / 003001101103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200132020233321-3110011103223002-2212002221111213-0202201301132010-3023001011313032-0123113231231333-3011233310200120-2210200300322311"></a>

## Next pages — default_header_transformation / 003001101103 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-021.md#canonical-1000301122202312-3011230331001222-1103133100213100-0100322330231321-3210311100100323-1222032203300032-0300300200303201-1102223101130200)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1200211130220221-3133230021110123-3003231020332330-0213130301203020-2300203200121212-0301202202201310-3321211010321203-2030023311211121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120320101313032-3223121311122013-0032320310202002-1213201123002023-2331223011210212-1303000222000333-1020000311022202-0331313130110032"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 230302230322 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-021.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-021.md#canonical-1033131203023032-2102001223320103-0323131200122322-2301231132233112-1223111332323113-3231022031110001-1223033201000032-3310113300131012)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-021.md#canonical-1000301122202312-3011230331001222-1103133100213100-0100322330231321-3210311100100323-1222032203300032-0300300200303201-1102223101130200)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-2113302011011022-1302022023202000-0220201130022221-0332133302203212-2221121101133330-3121333220130211-0003003002003130-3321112111201102"></a>

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

<a id="canonical-3303100112032213-1100213211010223-2213331313300010-0221120123022030-3303232320122013-1013110230132001-3311310303030001-1103023203120210"></a>

## Direct properties — preserve_case_header_transformation / 230302230322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132312112210333-0110200230330203-3100330103222311-2220021012321300-0313230130132223-3310033232020010-1200320012130221-3220002001203203"></a>

## Next pages — preserve_case_header_transformation / 230302230322 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-021.md#canonical-1000301122202312-3011230331001222-1103133100213100-0100322330231321-3210311100100323-1222032203300032-0300300200303201-1102223101130200)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2112201210020123-2221033111203013-3202202012112111-0123232323311222-0013103312310011-1103112023303220-0201102213220121-0220331221020210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213330033033133-0110320110013310-2101203223310012-3200012310300102-2201230301001101-3003303202302030-2033013313330312-0301332001212201"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 121322210320 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-021.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-021.md#canonical-1033131203023032-2102001223320103-0323131200122322-2301231132233112-1223111332323113-3231022031110001-1223033201000032-3310113300131012)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-021.md#canonical-1000301122202312-3011230331001222-1103133100213100-0100322330231321-3210311100100323-1222032203300032-0300300200303201-1102223101130200)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-2020100112310122-1313220102020333-2320311020321232-0203222031301312-2122103321322003-0011130212120212-0213020330102101-1330313001122220"></a>

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

<a id="canonical-0311130301303311-2023201022102013-0130000112323211-3322203021121000-1320323013230333-3211023303200011-3122003220111232-1322221330033120"></a>

## Direct properties — proper_case_header_transformation / 121322210320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023323022012300-2230103301200020-2000222000211323-3331312312323003-1333303213203110-2213221022010200-3013121001133010-2020232313103330"></a>

## Next pages — proper_case_header_transformation / 121322210320 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-021.md#canonical-1000301122202312-3011230331001222-1103133100213100-0100322330231321-3210311100100323-1222032203300032-0300300200303201-1102223101130200)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2203131003213122-3331030200230003-1001023323300301-1023203312002120-2112303110321012-2131131303320221-3032033312210013-2023033121200003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231321001312133-2312221032112031-3133201302311213-2111202300101130-0031123102120311-2022211012112020-2000010332112022-2011111311112020"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 002113011130 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-021.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2331123030131123-2002202100033231-3322101322212130-3001333300210022-3212212320232230-0101120131132031-0010121332233301-0232112220233011"></a>

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

<a id="canonical-0220320232322200-0323131031303300-1010001233110233-3321103012002131-1100033103123100-2002123301101021-3121102001100112-0332012202322203"></a>

## Direct properties — http_protocol_enable_v1_v2 / 002113011130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0220130330300011-3020023030301022-1010331033233000-1322132333132232-3230233132213111-2223220311033030-2330203221311020-2303032313320100"></a>

## Next pages — http_protocol_enable_v1_v2 / 002113011130 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-021.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2330131301031201-1210032120330101-1321113003123120-0103221211320010-1333112300210212-1220220331310110-2030120112022001-2103332101212322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012223002002211-2023013333300113-0222011102101111-1232311221033133-2121333033211233-1120123320200232-0233010001102220-3102200020133032"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 313010201012 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-021.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-0113333003323322-3213003110033122-2211120330332332-3200310200300310-1302203112202231-0330213030021010-2201222033322133-2131323232222300"></a>

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

<a id="canonical-3332321321201203-0012301333111231-3113223310221020-2212303003013111-1302320300210332-3332311033301032-1012201131021123-3200033210033230"></a>

## Direct properties — http_protocol_enable_v2_only / 313010201012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332333200103303-1220232311002303-1313330203202120-3123022301110212-3322133221321202-1013101033120203-1122002200121211-3132131120232213"></a>

## Next pages — http_protocol_enable_v2_only / 313010201012 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-021.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0011033233320121-3333020002121100-1201121003022320-2303002220331022-3023331020321233-0320223300120033-3021203022022003-2131002203233322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321000323022322-0233320120322131-2302123030011100-0213101332030322-2323330131033112-0030102130103100-0101322332233212-0100013232102112"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer — non_default_loadbalancer / 222020111113 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-3302103220201211-1221111323001110-3311302020010322-3023003202032332-0232113322030010-0233310112030010-3332112123210311-0322133323022312"></a>

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

<a id="canonical-0310110003213322-1231133233111221-1203233212131232-3300321133302320-1323021033330011-3032232012203333-1203100203110330-1031221112212100"></a>

## Direct properties — non_default_loadbalancer / 222020111113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120303111300120-0321320331203131-2213030131001010-0121120223220000-2301321301303232-2301001123231313-2012301231102331-1000310312100111"></a>

## Next pages — non_default_loadbalancer / 222020111113 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0000223200231120-3103130233313330-2132331313003321-2203211112112303-0302010122020301-3031301310323333-3132033122212331-3022302332313302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031223030003202-0301233132021032-3003201132203212-0120122212023200-0211321133100320-2221012110212032-2102103122330312-2001330211003322"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through — pass_through / 203222132303 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through

<a id="canonical-3213302221223101-2120012310120030-2223223110111232-3231113001113103-2300122333120032-0223232332333213-1312320130202303-0303211100222222"></a>

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

<a id="canonical-2120323131101223-3013102332212322-0121300023322220-3013002321021021-2211313231021212-1231032200311101-2320100311310231-2200221303300130"></a>

## Direct properties — pass_through / 203222132303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031301301221313-3023303131332201-1303320033103210-0101111232131022-0223230303133303-3112310033222211-1212110321233121-1301013332132331"></a>

## Next pages — pass_through / 203222132303 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211110212300031-3122200012221221-1011331123013100-0333012212210231-2113131312221122-3113110320112210-2123213130321300-1321212233330203"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params — tls_cert_params / 101330302320 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params

<a id="canonical-2302013020310322-1011222101323331-2133212202023023-0011312332233003-0321233203213330-1230232222331121-1310022311310313-3321133020330111"></a>

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

<a id="canonical-1233303100321000-1201113123301131-1313301021120130-3322133001331111-1102323033032013-1330231301330232-1332100232313203-1232230321023331"></a>

## Direct properties — tls_cert_params / 101330302320 / 3

- [certificates](data-sources--workload--reference--group-021.md#canonical-3020211132312003-2132320021113331-2130210102202222-3121203032130210-2022321103332203-2002131130120123-2232003211133002-1130333010212001): complete subsection reference.

- [no_mtls](data-sources--workload--reference--group-021.md#canonical-3223131212211020-2200001022101101-1300033222302333-0031122200112122-1201200332201032-2331221210210122-1231313302012312-1010003303232203): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-021.md#canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-021.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102): complete subsection reference.

<a id="canonical-3003200110213100-3223231130133111-0032002231102222-1101110203011021-1312011120110210-2101110323121032-1313120231121122-2120031301031101"></a>

## Next pages — tls_cert_params / 101330302320 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates](data-sources--workload--reference--group-021.md#canonical-3020211132312003-2132320021113331-2130210102202222-3121203032130210-2022321103332203-2002131130120123-2232003211133002-1130333010212001)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls](data-sources--workload--reference--group-021.md#canonical-3223131212211020-2200001022101101-1300033222302333-0031122200112122-1201200332201032-2331221210210122-1231313302012312-1010003303232203)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-021.md#canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-021.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3020211132312003-2132320021113331-2130210102202222-3121203032130210-2022321103332203-2002131130120123-2232003211133002-1130333010212001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200131303012003-2322102330130022-3203333230323120-1220102000201331-2102010003330200-3031110022230321-2302331012221022-3211321303132102"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates — certificates / 321123210021 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-1033231102230031-3003330321122313-0032323110013313-0223301033000100-1131231021213111-1113132332021301-0312113032113232-0331221113102102"></a>

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

<a id="canonical-0103002132203100-0323112121022230-0230130311103222-2120112100231231-1330012121122323-2110121211223210-3031201311010123-1100120120013031"></a>

## Direct properties — certificates / 321123210021 / 3

<a id="canonical-0000220331230122-2312302022033102-2031031301233320-0323013012100002-1012303130102202-1321102203033300-2110303103313021-2001003111311220"></a>

<a id="canonical-2230210201013012-1211120123221210-1333302332313320-1323300333111201-3312230110223020-1122111213110023-0311312121322202-0101302211021012"></a>

## name property — certificates / 321123210021 / 4

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

<a id="canonical-1222302220330333-2212130002202210-2203122312332210-1231230300022222-3032230232030102-3322311033232031-2022022300000102-1322213132331221"></a>

<a id="canonical-0313022133223300-0210022111130321-3323333012133103-2302011201232320-3133220331133021-2233023200311332-2023033101230210-1133221231120020"></a>

## namespace property — certificates / 321123210021 / 5

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

<a id="canonical-0012333213213201-0321312113333121-1023313003002310-2233313110001221-1123103033322310-2102101011322033-3113210332310123-3223130313212123"></a>

<a id="canonical-3110131001013333-0012002113302132-2120133013123011-3132122023033122-0320233111023321-2023232331121121-1021132111223100-0210213320220102"></a>

## tenant property — certificates / 321123210021 / 6

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

<a id="canonical-2012001222311230-3333103213111013-0030032112002330-2223333111000332-3010122112032131-2111231311032213-0323031302133302-2102101100013112"></a>

## Next pages — certificates / 321123210021 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3223131212211020-2200001022101101-1300033222302333-0031122200112122-1201200332201032-2331221210210122-1231313302012312-1010003303232203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332102010320031-0220223322032002-1010322231012001-3122222132333222-0011231221231203-1110323000113210-0333213313012020-0312010121103020"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls — no_mtls / 321001131210 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-2201330023113121-0103131303302033-0120230001023332-1001111323030113-1230033312212301-0320112200022312-2033201132020320-3101313130323323"></a>

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

<a id="canonical-3133021013010311-1223331122313303-1303010313131200-1010200311213122-1320001013231000-3033230101131301-0130133122210212-2123100201201030"></a>

## Direct properties — no_mtls / 321001131210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301230321323310-1232231220300131-2133032222210221-1022120302113003-0020110221213203-1133130302203322-3323322102031111-2223312322002303"></a>

## Next pages — no_mtls / 321001131210 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001113023233320-1102332213302033-2300022120131023-3130033312223122-2012311100320100-3233020002123303-3301231132003100-1103211011201312"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config — tls_config / 110033033223 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-1103202111030200-0102212233223010-0300211020230312-2313102332011332-1301222211131221-3131321002122132-3330201331203002-3112313001202233"></a>

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

<a id="canonical-3113220001323130-1100230310111001-2103302301212020-1122021201132333-0230312133221112-2101033331202011-1031220310200001-1010120321021132"></a>

## Direct properties — tls_config / 110033033223 / 3

- [custom_security](data-sources--workload--reference--group-021.md#canonical-1202101111222012-3002332031013121-2120032333220101-0100133001331210-2222120200120130-0131113102130032-2203223223311223-0300113121310002): complete subsection reference.

- [default_security](data-sources--workload--reference--group-021.md#canonical-3222102130230311-0231110001321300-0011230131323031-1020131303012223-0130201332100213-2032113133300312-0322320210210001-3332232002211021): complete subsection reference.

- [low_security](data-sources--workload--reference--group-021.md#canonical-1312002002313030-0033001230133003-3003333310321212-2201302000122233-2221200310113302-2032001032230302-1210113010131121-0022032131222010): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-021.md#canonical-0112110021323100-0222131201310330-2023013210131131-2303321032320031-3233033313023032-2101013020213200-2300300333133001-0202101221002022): complete subsection reference.

<a id="canonical-1021323320201213-3012100302021210-0203030313321333-2331231233322323-1311121011020121-3330203211001132-3123301122313030-0113031302033323"></a>

## Next pages — tls_config / 110033033223 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security](data-sources--workload--reference--group-021.md#canonical-1202101111222012-3002332031013121-2120032333220101-0100133001331210-2222120200120130-0131113102130032-2203223223311223-0300113121310002)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security](data-sources--workload--reference--group-021.md#canonical-3222102130230311-0231110001321300-0011230131323031-1020131303012223-0130201332100213-2032113133300312-0322320210210001-3332232002211021)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security](data-sources--workload--reference--group-021.md#canonical-1312002002313030-0033001230133003-3003333310321212-2201302000122233-2221200310113302-2032001032230302-1210113010131121-0022032131222010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security](data-sources--workload--reference--group-021.md#canonical-0112110021323100-0222131201310330-2023013210131131-2303321032320031-3233033313023032-2101013020213200-2300300333133001-0202101221002022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1202101111222012-3002332031013121-2120032333220101-0100133001331210-2222120200120130-0131113102130032-2203223223311223-0300113121310002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022012213000013-2313022232000212-3302032133011321-0110031031012011-1123200210203031-3231023131212100-2201333320231102-3022003012303111"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security — custom_security / 112011022202 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-021.md#canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-3121001020131300-0230230322303010-2020221332211102-3000110300200020-3203112201202132-0122332231033010-1020122201112100-2111331030302011"></a>

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

<a id="canonical-1011100021313311-3310213310213013-1310111303232303-1332012302031323-3033110123323020-0022000201101221-1211202010101223-3201212302211120"></a>

## Direct properties — custom_security / 112011022202 / 3

<a id="canonical-2322303121210022-0100231310323201-0031332112310013-1130312120030000-3010312013120300-3333133012022132-1300113103113023-1300021300032103"></a>

<a id="canonical-3302320030113211-2003231133030223-3011133322310222-1032221233010210-2130131113310113-1131220222130031-3132011123121123-0321103233112322"></a>

## cipher_suites property — custom_security / 112011022202 / 4

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

<a id="canonical-0013233223301130-0010313130323002-2102332332010231-3023221101300113-2311210122331131-0120302011001103-0000230321000011-3103222201313211"></a>

<a id="canonical-0131001321120101-3212111320303120-2130332110131310-0201210232133012-0101220021110330-1012033221120010-1231011301120123-3030111202102301"></a>

## max_version property — custom_security / 112011022202 / 5

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

<a id="canonical-1212221230103111-3131212211021302-3312120001213010-0333331201100303-2101013030233131-1020320010100221-2223133302102323-1323311320122200"></a>

<a id="canonical-3110201213330103-3013002110221013-0200101121302301-3022200303230112-1020310233222121-1010032323300233-0110200101313211-2111021301133201"></a>

## min_version property — custom_security / 112011022202 / 6

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

<a id="canonical-2100103033321123-3120101211030120-0102232322213202-0002212232033203-3120221212113232-0123321131201110-0131001301102212-0331020313312022"></a>

## Next pages — custom_security / 112011022202 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-021.md#canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3222102130230311-0231110001321300-0011230131323031-1020131303012223-0130201332100213-2032113133300312-0322320210210001-3332232002211021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020000302332010-1000020122120023-3303212121211312-2203023310330333-3122120223333331-1131213233313301-1022030232300222-0131010020023202"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security — default_security / 033230320000 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-021.md#canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-2333320333032221-3033301301213332-0322222220333022-1110332031230032-3123203203332302-0331113130112130-2033201011010301-0310211111003330"></a>

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

<a id="canonical-3130103221321132-0131013321003111-0230221331211322-1101013213003333-0002101101321133-1130101100211031-1320233112313210-2100123330133222"></a>

## Direct properties — default_security / 033230320000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301102222323100-3310132313213323-2330323201113131-0202210002301103-2302300031230103-3120232333320012-2311011000032330-0331023000302121"></a>

## Next pages — default_security / 033230320000 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-021.md#canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1312002002313030-0033001230133003-3003333310321212-2201302000122233-2221200310113302-2032001032230302-1210113010131121-0022032131222010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300230230233010-2020210200231030-1131311123222300-3132212033023213-0322012203121110-3220112213210121-2333313132001201-1111120230010120"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security — low_security / 233102222133 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-021.md#canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-0000022131332011-3111330210000222-1030113013121210-3121200212111022-0311211132213113-1222232321312201-0120021132010312-0100120122131112"></a>

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

<a id="canonical-1003220300321022-3323213202121110-0120231313320212-2111030231211201-1020000120212121-1111130020130131-2022210233132102-0003001312132303"></a>

## Direct properties — low_security / 233102222133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200210110232023-3112020131313102-3020230321031332-1131213013102123-1210331001010122-1333303310313331-2233112320220322-2003202103133302"></a>

## Next pages — low_security / 233102222133 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-021.md#canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0112110021323100-0222131201310330-2023013210131131-2303321032320031-3233033313023032-2101013020213200-2300300333133001-0202101221002022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132320313301032-3022120122230331-2100200200211221-3032330301301203-2003131200011011-3031020210102100-1102300332233013-2033121313133213"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security — medium_security / 231220021103 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-021.md#canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-3312113202131202-2212212223010303-1100020011131111-0121302321311213-2103012000312313-0201322231320020-2130230332130202-1220320203203011"></a>

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

<a id="canonical-2101312303302321-0322232111121123-3311201032211032-1112100222011103-1013202202231101-1021132230331200-2121321213012002-3321113311121121"></a>

## Direct properties — medium_security / 231220021103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230031013122030-3013013200303310-2120131100123313-3033330032301102-2110003122221313-1002311201120132-2223310101210012-2322132302211220"></a>

## Next pages — medium_security / 231220021103 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-021.md#canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331211102131330-3312133313312202-2321033103120323-3303223310200022-0211331232000221-0123232120323320-3112212002122220-1230313313022312"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls — use_mtls / 232323300131 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-2310330323202330-3013131313331233-3121002021111310-0022332221101111-0031013110301211-2232221323101330-0302320033102321-3232223300123133"></a>

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

<a id="canonical-2312320333231330-0201313322020033-1101302110013200-0332300113020323-1132223030103232-2212230320212320-2022122203110100-1110121031330022"></a>

## Direct properties — use_mtls / 232323300131 / 3

<a id="canonical-3321210203200101-3122303011023333-1003223102331212-2112303233101001-0030220233322201-3110313000131213-2203133131000323-1213032120203320"></a>

<a id="canonical-2213023012120110-0033303022300120-0100103200231333-1120210330330330-2201002221230031-1321003312333023-3003321201122201-1022001232222111"></a>

## client_certificate_optional property — use_mtls / 232323300131 / 4

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

- [crl](data-sources--workload--reference--group-021.md#canonical-1310213332120302-0103322220333223-1233111332130300-0311100132002131-0311133030130313-3311110112032331-0033021231223221-3320330103113330): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-021.md#canonical-3330111311003230-1013210213213010-3302003001221223-2022033011323212-0322102310001103-1012310011203120-1023110220010313-2110020233313100): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-021.md#canonical-2211301003212230-3232001131031222-1123233020020033-3303312111320132-0110303113321131-1210032210322331-3120003200201311-0221322022323101): complete subsection reference.

<a id="canonical-3013231120002001-0212101031210220-0113233212130212-0221112321121321-2002213133003201-1331302230023112-1330113202031201-3221010102303132"></a>

<a id="canonical-0021311110132010-1212332120111030-2031011112002023-1320220212012201-2021203011312001-2102211202012232-2120120332023030-1101232202113120"></a>

## trusted_ca_url property — use_mtls / 232323300131 / 5

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

- [xfcc_disabled](data-sources--workload--reference--group-022.md#canonical-0030111210231013-2222220121122130-2010310023010310-1130132313031002-1330222301322013-1210011310321303-2232121032200102-2310010213332002): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-022.md#canonical-3221100202303211-1330113112311332-1132323132323111-2020203132010123-0100312032302132-3010031232110212-2323021103032020-2102132331230010): complete subsection reference.

<a id="canonical-1311303003232031-2113110132012323-1221321011033200-3111032212112233-1300202120302210-3320311110130112-0032203100120000-3020320002031232"></a>

## Next pages — use_mtls / 232323300131 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl](data-sources--workload--reference--group-021.md#canonical-1310213332120302-0103322220333223-1233111332130300-0311100132002131-0311133030130313-3311110112032331-0033021231223221-3320330103113330)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl](data-sources--workload--reference--group-021.md#canonical-3330111311003230-1013210213213010-3302003001221223-2022033011323212-0322102310001103-1012310011203120-1023110220010313-2110020233313100)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca](data-sources--workload--reference--group-021.md#canonical-2211301003212230-3232001131031222-1123233020020033-3303312111320132-0110303113321131-1210032210322331-3120003200201311-0221322022323101)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled](data-sources--workload--reference--group-022.md#canonical-0030111210231013-2222220121122130-2010310023010310-1130132313031002-1330222301322013-1210011310321303-2232121032200102-2310010213332002)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options](data-sources--workload--reference--group-022.md#canonical-3221100202303211-1330113112311332-1132323132323111-2020203132010123-0100312032302132-3010031232110212-2323021103032020-2102132331230010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1310213332120302-0103322220333223-1233111332130300-0311100132002131-0311133030130313-3311110112032331-0033021231223221-3320330103113330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031230321332233-1122022111302021-0213221110031230-3132303002000030-2113223112323223-1002033133002123-2220132321330323-1023021211021100"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl — crl / 003100103130 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-021.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-0132112022122020-3331210302212100-1112313002323020-1312321303210211-2330302231012123-3132110222202012-3231030113023131-2201030210011033"></a>

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

<a id="canonical-1330102303100012-0301012111322231-2232111110221203-2222020113211321-0330213323023212-2321231111202233-1010333113110223-1031213131203113"></a>

## Direct properties — crl / 003100103130 / 3

<a id="canonical-1002330032111230-1303130211231033-3212132103312003-2302033302130022-3332332231201332-0013022312321031-3000000101222322-3312001123021110"></a>

<a id="canonical-2120222023230301-2131320123000330-0210111213030213-2333303012110232-0100022221201201-1201022022331233-3330202222303100-2120132130230321"></a>

## name property — crl / 003100103130 / 4

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

<a id="canonical-0221011030200332-2122132330003310-3332200020023310-0202022002233131-1013202220131320-3321203312001300-2012231222001331-1332031132122011"></a>

<a id="canonical-2030133302333131-1011300331212032-1201223310013221-0031132020320012-3103301331112311-3100022031323132-3211222101330131-0210133031202101"></a>

## namespace property — crl / 003100103130 / 5

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

<a id="canonical-1131320230231002-1021002133130230-1001101303030320-2030003221121023-0113100220223303-3120010203233232-3022311202322233-0131332233202023"></a>

<a id="canonical-2312300211032133-2002113011002302-2103330021022312-2202303222332123-3000000123100220-2022310020213011-1202202221312001-3122213332000101"></a>

## tenant property — crl / 003100103130 / 6

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

<a id="canonical-1001110222120233-2323232032332100-2033323333310303-1321220220122022-1122121012120331-1212033312332311-2201122100131232-0311320000131111"></a>

## Next pages — crl / 003100103130 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-021.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3330111311003230-1013210213213010-3302003001221223-2022033011323212-0322102310001103-1012310011203120-1023110220010313-2110020233313100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131120132022323-1233102121102330-3220002113302000-3032230133213012-2022122003011332-0331130232302013-2232111010231030-3330202012211131"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl — no_crl / 201211220222 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-021.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-3322331220000112-1112303031111110-3102132101120030-0310210322033131-0130002020103201-2030320223110322-3031303132113103-2322211123020112"></a>

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

<a id="canonical-1001222311310213-1100221011231223-0232112033220332-1033121130031202-0300302211220003-0112232311312102-3221102213230330-1103100132001311"></a>

## Direct properties — no_crl / 201211220222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012221123202021-2311231112332031-1212333332133033-3320021220212201-2113031030202012-2203021131300100-0331210000013030-3001301112121030"></a>

## Next pages — no_crl / 201211220222 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-021.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2211301003212230-3232001131031222-1123233020020033-3303312111320132-0110303113321131-1210032210322331-3120003200201311-0221322022323101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231211300213112-3303300120132233-1023021303230123-2111102132122002-1322232030213122-0332231202130212-1010231122212313-1011210120123232"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca — trusted_ca / 012001310310 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-021.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-1100232020213003-3312331302332222-2032320333311030-0232020232122203-2101331223200120-0223231031102022-2222013011222212-0303222332331232"></a>

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

<a id="canonical-2132102200022323-3322303103310332-0302312100203331-0210320132110112-1103112310102032-1131332223302311-2112230221321000-3121313132111220"></a>

## Direct properties — trusted_ca / 012001310310 / 3

<a id="canonical-3021002003302121-0333220030022021-1331221022212331-2200003232010331-2230013200220010-1102121201002221-1312100033231321-1303330223032021"></a>

<a id="canonical-1120130012133133-2101021322311300-0101101121010101-1001320113230210-1232022202311030-0132323302231333-0000221113120302-0023303002313000"></a>

## name property — trusted_ca / 012001310310 / 4

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

<a id="canonical-3011112333030103-2332320102300210-3131033301231220-0321123002121311-0230222313113213-1332203002121003-2330011010231120-0330102212030321"></a>
