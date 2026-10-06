---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1023120122302200-2011233320230320-1102233120232223-2013001132122323-1223332131001032-3021210021131100-2203100132021113-3312331033302103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options

<a id="canonical-2203323301213223-3231023200122010-0203013223011010-1203101230213020-1332312123201321-3022301311200330-0110002301123311-2130222212110020"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

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

<a id="canonical-2310220020013310-3230302303221223-2233121331033021-0201001102201213-3223011312300033-0222010012230121-0203022023031110-3031010213122131"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options`

- [default_coalescing](data-sources--workload--reference--group-024.md#canonical-2103211223112000-1203311230111230-1221320100113101-0132330220101102-0013033232113021-3310121201133210-2031320011000311-3010132331303100): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-024.md#canonical-3011320110330222-2030112311201003-2112133010201030-0222301212130320-0103330022012101-3332200131230123-0322122330323231-0200323013130132): complete subsection reference.

<a id="canonical-2103211223112000-1203311230111230-1221320100113101-0132330220101102-0013033232113021-3310121201133210-2031320011000311-3010132331303100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-024.md#canonical-1023120122302200-2011233320230320-1102233120232223-2013001132122323-1223332131001032-3021210021131100-2203100132021113-3312331033302103)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-0332300103311012-3001223203123001-3003031013333003-2303031100123233-2031323200302030-2211123300320102-1213123310222031-3321022323003211"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-3011320110330222-2030112311201003-2112133010201030-0222301212130320-0103330022012101-3332200131230123-0322122330323231-0200323013130132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-024.md#canonical-1023120122302200-2011233320230320-1102233120232223-2013001132122323-1223332131001032-3021210021131100-2203100132021113-3312331033302103)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-1303221103313030-2330002231312200-0200021333101231-0220111122021313-3202231212110113-0230323233133103-1110133320232000-1031003221313211"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-3212030333331111-1323213222123123-3113032102033012-1200212310332211-0113033032320311-1220000132023221-2113121320322131-2232332232333031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header

<a id="canonical-2120032322102101-0131120102310322-0021032210033303-2111303001201312-1113303100023120-3112212213123231-2100022011112333-0001320000222101"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

<a id="canonical-0223223101131313-1300310221112102-0321210132001022-3112212300000023-0132301001002132-3302030003012101-3320110120100223-1110101022012123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer

<a id="canonical-3011333321320210-0020320200302030-3100221212332001-1101133220001133-1220132321212222-1321121022331123-1101333022032300-2130233233022221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default loadbalancer.

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

<a id="canonical-2333223311220130-2111013331020103-1221120102101213-3220002012231222-2021122321032332-1333013010022023-0131031320013100-0120221203102300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize

<a id="canonical-0012133022203313-0002301230330313-2130112010230333-3230132023203113-3223333232030302-3223100123322230-2012321301201030-0331310110231213"></a>

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

<a id="canonical-0223132321010103-3022011311232223-3212232213012101-1320003202112021-3211202013132331-3121121312002301-1330203100230032-2221331332130333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize

<a id="canonical-2130223123320021-0033132311231022-3220133232103221-1312301000020130-2321103113221103-0201101131132003-2311202313003133-1031213122032010"></a>

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

<a id="canonical-3333101033030012-1032120202302303-2102001301123323-2131030113312320-3213223211120231-2113300031233322-0220322000113332-0333121113331310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options

<a id="canonical-3012323131331231-3210131102231231-3210232003130213-2233231032030210-1313321121332022-2031033310100120-2123012022221000-1012102033313121"></a>

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

<a id="canonical-2211200222203012-2320201210100332-3101231112022310-2113002301301222-2323201333223130-0200311033123101-2301120313322301-2032212021310133"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options`

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-024.md#canonical-3300320011203011-1210133323120233-2131113302030100-1230323010213120-0203303020111003-2013130012011210-2102210122232110-2112230323021100): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-024.md#canonical-0220131310311122-3120003032210212-2001333311320302-3123100130023201-1020032013313313-1232310322030222-3110103122332311-0310201020323021): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-024.md#canonical-1230003231101131-0102300131032130-2022231311330100-1021022132012311-1121102112112022-0002330130333303-3011301112221000-1002232020213000): complete subsection reference.

<a id="canonical-3300320011203011-1210133323120233-2131113302030100-1230323010213120-0203303020111003-2013130012011210-2102210122232110-2112230323021100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-024.md#canonical-3333101033030012-1032120202302303-2102001301123323-2131030113312320-3213223211120231-2113300031233322-0220322000113332-0333121113331310)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-3203110303300131-2022121200003033-3231030210012320-3101222031230022-0212102231303312-2131220230121232-1211223333331101-0022200022023333"></a>

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

<a id="canonical-0131212012022320-0003021003001112-3211121030232223-1300120112111323-1033330030011121-1323010222300310-2122320330001211-3313312310233033"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](data-sources--workload--reference--group-024.md#canonical-3122222212030101-1203221120031210-3333023032021123-0223230011222133-3233112301020012-2011220311220011-1122123312122111-0133320312023110): complete subsection reference.

<a id="canonical-3122222212030101-1203221120031210-3333023032021123-0223230011222133-3233112301020012-2011220311220011-1122123312122111-0133320312023110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-024.md#canonical-3333101033030012-1032120202302303-2102001301123323-2131030113312320-3213223211120231-2113300031233322-0220322000113332-0333121113331310)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-024.md#canonical-3300320011203011-1210133323120233-2131113302030100-1230323010213120-0203303020111003-2013130012011210-2102210122232110-2112230323021100)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-3230110020032111-2003011223211232-1301010122202330-1310022011000001-2103110013121223-3120220211003312-3001102133200212-2303220211323120"></a>

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

<a id="canonical-0030121333133203-0132132110022221-0333100313232013-2101021221031110-3001011112102222-2121322001013300-2010100003302101-1122233220310313"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](data-sources--workload--reference--group-024.md#canonical-0031323121102300-0333321012113223-0011020330023110-0023320100323302-2233001223221013-1112130133013220-3120222213021313-0213203101321111): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-024.md#canonical-2101222223223301-2010302221021123-3103213310223003-3221232220011220-0012133332002131-0121301010130302-0323032003213322-0231033023001123): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-024.md#canonical-1022122133111101-1332003303032202-1301322031220110-2332232111311102-3200000213102003-3031002202333202-2133022321322313-2003323102002120): complete subsection reference.

<a id="canonical-0031323121102300-0333321012113223-0011020330023110-0023320100323302-2233001223221013-1112130133013220-3120222213021313-0213203101321111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-024.md#canonical-3333101033030012-1032120202302303-2102001301123323-2131030113312320-3213223211120231-2113300031233322-0220322000113332-0333121113331310)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-024.md#canonical-3300320011203011-1210133323120233-2131113302030100-1230323010213120-0203303020111003-2013130012011210-2102210122232110-2112230323021100)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-024.md#canonical-3122222212030101-1203221120031210-3333023032021123-0223230011222133-3233112301020012-2011220311220011-1122123312122111-0133320312023110)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-2233033030101013-0201323221030333-3230002321301312-1122221120303321-3122132202313012-3230031211321011-0230231030021100-1322203111203133"></a>

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

<a id="canonical-2101222223223301-2010302221021123-3103213310223003-3221232220011220-0012133332002131-0121301010130302-0323032003213322-0231033023001123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-024.md#canonical-3333101033030012-1032120202302303-2102001301123323-2131030113312320-3213223211120231-2113300031233322-0220322000113332-0333121113331310)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-024.md#canonical-3300320011203011-1210133323120233-2131113302030100-1230323010213120-0203303020111003-2013130012011210-2102210122232110-2112230323021100)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-024.md#canonical-3122222212030101-1203221120031210-3333023032021123-0223230011222133-3233112301020012-2011220311220011-1122123312122111-0133320312023110)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-2211333311220222-3113010321013123-1012233003311321-3303212011231231-1310302223210112-3221120320023201-2203020202310221-0321013200222333"></a>

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

<a id="canonical-1022122133111101-1332003303032202-1301322031220110-2332232111311102-3200000213102003-3031002202333202-2133022321322313-2003323102002120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-024.md#canonical-3333101033030012-1032120202302303-2102001301123323-2131030113312320-3213223211120231-2113300031233322-0220322000113332-0333121113331310)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-024.md#canonical-3300320011203011-1210133323120233-2131113302030100-1230323010213120-0203303020111003-2013130012011210-2102210122232110-2112230323021100)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-024.md#canonical-3122222212030101-1203221120031210-3333023032021123-0223230011222133-3233112301020012-2011220311220011-1122123312122111-0133320312023110)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-2030331220203113-1033001311112020-1233201221322301-1233221132300122-1230332322110210-3301302213000311-0313331132020103-3210032032301210"></a>

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

<a id="canonical-0220131310311122-3120003032210212-2001333311320302-3123100130023201-1020032013313313-1232310322030222-3110103122332311-0310201020323021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-024.md#canonical-3333101033030012-1032120202302303-2102001301123323-2131030113312320-3213223211120231-2113300031233322-0220322000113332-0333121113331310)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-0101112321113123-3131122001200301-0203320133102232-2101321010220003-0310112221311202-2120121323122203-3301111001123202-1130020230332311"></a>

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

<a id="canonical-1230003231101131-0102300131032130-2022231311330100-1021022132012311-1121102112112022-0002330130333303-3011301112221000-1002232020213000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-024.md#canonical-3333101033030012-1032120202302303-2102001301123323-2131030113312320-3213223211120231-2113300031233322-0220322000113332-0333121113331310)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-3113102210312010-2020323030103332-3322100110211132-3220010130113123-1012003201022020-0213102032220322-1020303332230213-1021212032001102"></a>

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

<a id="canonical-0301011213122133-1300033113213030-3120002233210213-1112230010011111-0020221320132203-0321110303123233-2332103123233312-1221221101223010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-3012330323012332-1120120111311001-0110032002132120-0001003232313311-2011303132013302-3323000101231212-0032313211130233-2111223220323131"></a>

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

<a id="canonical-0122012011231231-1200131133311312-1030012131023030-2131222022303011-3123012333133233-0323201300031012-2231023102333330-1101231222212323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through

<a id="canonical-2232321122113312-2300011300233000-2323300211203130-0231212201212131-3301120133331010-3302130012011333-3101223220203312-3132213332330130"></a>

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

<a id="canonical-3322122302333010-1112232022110101-3332210121013303-0203013103132310-3121123232001030-0330332323311120-3221321223111332-3210202001303013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params

<a id="canonical-2210303120223311-0110232332231012-0122211022032001-2222321022322202-1323230321222122-2231121112021033-2211122111030013-1103211211202031"></a>

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

<a id="canonical-0001223113120330-3001221032313230-3213310000123131-0133220031122130-2232213002031020-1000030133131030-0021112322100030-0111332030130120"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params`

- [certificates](data-sources--workload--reference--group-024.md#canonical-0230221313321213-3221233322101333-1223003312031031-1213110020320012-2210321031103322-2103221020110201-1203311213001003-1103102310132100): complete subsection reference.

- [no_mtls](data-sources--workload--reference--group-024.md#canonical-2201031030113022-2230123200132220-1321103313222301-3233013212030033-1310233223020232-3123330222012031-2320021130013102-1332023202200023): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-024.md#canonical-2332113310012103-1233213011002233-0333303323312301-2120002031202112-0321232302020133-2013331200133202-3213321020311202-2033132113000201): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-024.md#canonical-3311323300302010-3101202300203330-3130113330211032-0321122322321200-0001203110102011-3311311300223331-0013303330212203-1222121012000312): complete subsection reference.

<a id="canonical-0230221313321213-3221233322101333-1223003312031031-1213110020320012-2210321031103322-2103221020110201-1203311213001003-1103102310132100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-024.md#canonical-3322122302333010-1112232022110101-3332210121013303-0203013103132310-3121123232001030-0330332323311120-3221321223111332-3210202001303013)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-2111302202202021-1320102330022122-2323103021001101-0021011321100330-2200022131313031-2202211321103020-1003131101021030-2231013312123232"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0210221332033003-1130311311202002-2330200021110233-0323121302033032-2130322102133220-0012310203132031-1120101313102010-1302212333310132"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates`

<a id="canonical-3331033010231020-2121013212223121-2310301023001032-2123213202011203-1233223123112300-3203100013111320-0203120030302033-0203132012112322"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2230210303232111-2112010002031101-0323202133220100-2002020131223001-3313033103121322-1033032112332011-0221031033032322-0332221003220130"></a>

<a id="canonical-0102332121201011-0010332123023032-2212031112233003-2112202100320121-2301132303032331-3130233120100303-2032011301003132-2333201032220020"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0211212223131031-1202020313332112-1121032033322130-3111233302133123-0113221112220213-2103013020230302-0122221100122323-2333332030103031"></a>

<a id="canonical-0033213220211300-2122220130223331-2211030022101200-1322120131000002-1220212031322300-0002121132231303-3220332202033003-0313211102102022"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2201031030113022-2230123200132220-1321103313222301-3233013212030033-1310233223020232-3123330222012031-2320021130013102-1332023202200023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-024.md#canonical-3322122302333010-1112232022110101-3332210121013303-0203013103132310-3121123232001030-0330332323311120-3221321223111332-3210202001303013)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-2310300201302021-2210332221003012-2231223233300101-1033223022313101-2100330110132130-3013322103300032-3031213113112021-1013201233023330"></a>

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

<a id="canonical-2332113310012103-1233213011002233-0333303323312301-2120002031202112-0321232302020133-2013331200133202-3213321020311202-2033132113000201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-024.md#canonical-3322122302333010-1112232022110101-3332210121013303-0203013103132310-3121123232001030-0330332323311120-3221321223111332-3210202001303013)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-2130302002230111-1001132112302130-1211321113233331-0122313323133012-1231200320133110-1331230000313203-1111020011023213-0030331131011302"></a>

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

<a id="canonical-3302302310100331-3003122131000322-2030323130031332-1001131320003213-3220332202011103-1001113002303303-2303101012120301-3313102202333113"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config`

- [custom_security](data-sources--workload--reference--group-024.md#canonical-0231020022032213-2333011101233000-1030001332200000-3100011210012102-3011202223010023-1101032332301200-2010201013200320-2120301232112112): complete subsection reference.

- [default_security](data-sources--workload--reference--group-024.md#canonical-2211102222202013-0101332012312320-0112232010111332-0102012211112011-2302001030321003-0221001031100113-0213232112022030-2101111333130111): complete subsection reference.

- [low_security](data-sources--workload--reference--group-024.md#canonical-0133323332032102-0222033300032232-0033233332011221-1022302110001011-2323203033031313-3230003321121210-0020322130003110-2332002201010030): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-024.md#canonical-3022122303300322-3200032122032220-3231202231230231-2003302203330003-1231032310302020-0113222131330113-1213033101120232-3202213031321323): complete subsection reference.

<a id="canonical-0231020022032213-2333011101233000-1030001332200000-3100011210012102-3011202223010023-1101032332301200-2010201013200320-2120301232112112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-024.md#canonical-3322122302333010-1112232022110101-3332210121013303-0203013103132310-3121123232001030-0330332323311120-3221321223111332-3210202001303013)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-024.md#canonical-2332113310012103-1233213011002233-0333303323312301-2120002031202112-0321232302020133-2013331200133202-3213321020311202-2033132113000201)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-1312221311332333-0112233130212312-2232131001021003-0113032011221211-2220001333302032-1123222230021122-0201011200033012-2103120101033020"></a>

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

<a id="canonical-0212033023110132-1032232323213201-3232233111301033-3133100210130210-0132320300212111-2200113131301030-1222033300212331-1333022323221312"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security`

<a id="canonical-3023002123122033-2021122003223330-0033212023030232-0100220022002003-3010203122012232-3230001010322312-0213000213010202-2223311120013323"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1133303133202103-0213013303121132-2211200331201032-2203000011232312-0210121300231202-1303023003132001-0011011021212223-0320301122321312"></a>

<a id="canonical-1011001002332301-1212023233101033-0002033323222320-2212320322233132-1011301213131133-1231120132010123-0300111030330330-3003231332333133"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.max_version` property

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

<a id="canonical-0200332100202200-3200011031113230-1300103001213223-2221121333220031-3012220023220002-3101201032332133-2011311132201323-0322311223322112"></a>

<a id="canonical-1023030213000033-1031333033212322-2013020220313131-1232111111310303-2030013003120113-1321313000023133-1221332002023233-1012231311022003"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.min_version` property

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

<a id="canonical-2211102222202013-0101332012312320-0112232010111332-0102012211112011-2302001030321003-0221001031100113-0213232112022030-2101111333130111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-024.md#canonical-3322122302333010-1112232022110101-3332210121013303-0203013103132310-3121123232001030-0330332323311120-3221321223111332-3210202001303013)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-024.md#canonical-2332113310012103-1233213011002233-0333303323312301-2120002031202112-0321232302020133-2013331200133202-3213321020311202-2033132113000201)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-1230322203130133-0320123301333220-2301332033310000-2022130101200323-2133221021120020-0211301233210001-0120203202021211-3022121301001130"></a>

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

<a id="canonical-0133323332032102-0222033300032232-0033233332011221-1022302110001011-2323203033031313-3230003321121210-0020322130003110-2332002201010030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-024.md#canonical-3322122302333010-1112232022110101-3332210121013303-0203013103132310-3121123232001030-0330332323311120-3221321223111332-3210202001303013)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-024.md#canonical-2332113310012103-1233213011002233-0333303323312301-2120002031202112-0321232302020133-2013331200133202-3213321020311202-2033132113000201)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-1122330020131202-3122302210303121-3223202301011200-3122232133312013-0130131113330003-2133032321001120-0101111221110012-1122100333301210"></a>

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

<a id="canonical-3022122303300322-3200032122032220-3231202231230231-2003302203330003-1231032310302020-0113222131330113-1213033101120232-3202213031321323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-024.md#canonical-3322122302333010-1112232022110101-3332210121013303-0203013103132310-3121123232001030-0330332323311120-3221321223111332-3210202001303013)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-024.md#canonical-2332113310012103-1233213011002233-0333303323312301-2120002031202112-0321232302020133-2013331200133202-3213321020311202-2033132113000201)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-3111331321100322-3021031102203233-0113302031021132-1133333020101000-3303333112102213-3221033102112300-1331101010110120-1233310102001311"></a>

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

<a id="canonical-3311323300302010-3101202300203330-3130113330211032-0321122322321200-0001203110102011-3311311300223331-0013303330212203-1222121012000312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-024.md#canonical-3322122302333010-1112232022110101-3332210121013303-0203013103132310-3121123232001030-0330332323311120-3221321223111332-3210202001303013)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-0033221100122110-0022323313133312-1023213131030123-3132210312112223-1103121112033322-1023322322112230-1323222300221300-0113332203020312"></a>

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

<a id="canonical-1101101202211131-0001222232000323-2301123033301231-2120302232030200-2231010000320132-1031123302223012-0010323332231031-3030232320303220"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls`

<a id="canonical-2300333311313332-3101210323203113-3110202102013201-0210212310222212-0333132201330322-2130023131120101-3032023103121233-1212101000300232"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.client_certificate_optional` property

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

- [crl](data-sources--workload--reference--group-024.md#canonical-1322210212311300-0120303320312100-2010321311203220-0011020321230311-0300022101033101-2001021120313103-0003320320231212-3231132133321102): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-024.md#canonical-1111332031120210-1210322313131321-2301111013331022-1133202031210002-3223330200213233-1112222221002323-1103021232023032-0220110221333000): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-024.md#canonical-1122020101313322-0320231133130112-2212312330101323-2212210322220000-3101003123211303-3233232211131103-2031231331032312-1333222213213003): complete subsection reference.

<a id="canonical-1302222012202332-0223112333131302-0101021011122300-0123320022101023-2122112032023303-0022332331231000-3021331333203021-3221010310112113"></a>

<a id="canonical-1133310210312100-2100102320110101-3013231313012010-2233301222213121-2333313121303122-1223033013212033-1011203120321023-2200322110003233"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca_url` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [xfcc_disabled](data-sources--workload--reference--group-024.md#canonical-3020002101110313-1113012332301310-1203102321333300-2310111013300232-1221033213031333-2222122110123012-2232230330010010-1121320331120210): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-024.md#canonical-2212233010100203-3310012333011213-2333231230331322-1312223221321102-1131312202113121-3333202121212220-0330331031211131-1213202223233230): complete subsection reference.

<a id="canonical-1322210212311300-0120303320312100-2010321311203220-0011020321230311-0300022101033101-2001021120313103-0003320320231212-3231132133321102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-024.md#canonical-3322122302333010-1112232022110101-3332210121013303-0203013103132310-3121123232001030-0330332323311120-3221321223111332-3210202001303013)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-024.md#canonical-3311323300302010-3101202300203330-3130113330211032-0321122322321200-0001203110102011-3311311300223331-0013303330212203-1222121012000312)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-3121200322110203-1222022223210220-1123111031313032-0321223230012301-0101302033012232-3021131333232132-3300311302000310-2100001333313222"></a>

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

<a id="canonical-3033232103333132-0002202202103133-0323233000032212-1111331221001213-2330210003112103-0012300220001123-3330211110231323-0131221030003000"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl`

<a id="canonical-3001330120221211-1313031022031102-1233302000202223-3102010320013101-1132000331332313-2103030320123323-3210210332030330-3333003132013311"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2331203003001220-3211320003210212-3133133000303100-2033030022331122-3331021212120312-1110321321202123-2021321122130300-0002313202110110"></a>

<a id="canonical-1323011310223130-1212033201130113-3210122032200101-3031013003130032-1212101000311130-3312213023033003-3133001130323011-0101211301311312"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2133302231232210-2021331013002033-0322032310332001-3122031001131231-0210331221032123-3300333112012332-1110213033013213-2101212132022232"></a>

<a id="canonical-0233222311013001-3330002301103231-3310302132213213-1323022220100103-0321031030102002-1123313202112012-2232001203331103-3122123003331132"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1111332031120210-1210322313131321-2301111013331022-1133202031210002-3223330200213233-1112222221002323-1103021232023032-0220110221333000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-024.md#canonical-3322122302333010-1112232022110101-3332210121013303-0203013103132310-3121123232001030-0330332323311120-3221321223111332-3210202001303013)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-024.md#canonical-3311323300302010-3101202300203330-3130113330211032-0321122322321200-0001203110102011-3311311300223331-0013303330212203-1222121012000312)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-2213301210211322-3122120130300102-3120303030003031-2220033200102121-2332123221103220-2013122132032011-3000130010133322-3230021011031302"></a>

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

<a id="canonical-1122020101313322-0320231133130112-2212312330101323-2212210322220000-3101003123211303-3233232211131103-2031231331032312-1333222213213003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-024.md#canonical-3322122302333010-1112232022110101-3332210121013303-0203013103132310-3121123232001030-0330332323311120-3221321223111332-3210202001303013)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-024.md#canonical-3311323300302010-3101202300203330-3130113330211032-0321122322321200-0001203110102011-3311311300223331-0013303330212203-1222121012000312)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-2030002320122003-3012000031120020-0001031310132022-2201201030130131-1011310110103210-3231012111022010-2211122320101103-2220130013013110"></a>

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

<a id="canonical-2121231202302231-0112011233200210-2010113112133201-3302300131330022-2323202303130000-3003331122322130-1312113223010120-1200301301321320"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca`

<a id="canonical-0311301201001102-2123210203122300-0022333132100300-1020333000023211-3111123013323121-2332110320203211-3321222100322111-1131131032110232"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3313013011201103-2222103330211302-0332322221212311-1021233032233300-1113212133213011-2221210132230321-2010212023100113-0302130302000233"></a>

<a id="canonical-2003231112110220-1102322021220202-3221312002323102-3131313210303113-1010223022020011-1021023033033020-1011302330010231-0033220110311321"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1112101020311332-1021201200103103-3012331000210310-3232202132123110-2322322133213332-3112000110310303-0130021032021130-1231201301331212"></a>

<a id="canonical-1010301100321110-3002322032132322-0201033032020023-0020310101103001-0030333032333320-2320122233333022-0003020033232202-3322012131311130"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3020002101110313-1113012332301310-1203102321333300-2310111013300232-1221033213031333-2222122110123012-2232230330010010-1121320331120210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-024.md#canonical-3322122302333010-1112232022110101-3332210121013303-0203013103132310-3121123232001030-0330332323311120-3221321223111332-3210202001303013)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-024.md#canonical-3311323300302010-3101202300203330-3130113330211032-0321122322321200-0001203110102011-3311311300223331-0013303330212203-1222121012000312)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-3302330330002333-1102031013020020-1222020302032210-2312120000331213-2221022033120332-0202332120221032-3200201213001213-2301112320210202"></a>

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

<a id="canonical-2212233010100203-3310012333011213-2333231230331322-1312223221321102-1131312202113121-3333202121212220-0330331031211131-1213202223233230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-024.md#canonical-3322122302333010-1112232022110101-3332210121013303-0203013103132310-3121123232001030-0330332323311120-3221321223111332-3210202001303013)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-024.md#canonical-3311323300302010-3101202300203330-3130113330211032-0321122322321200-0001203110102011-3311311300223331-0013303330212203-1222121012000312)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-0203102222122303-1210210312020031-3311011310310111-0322203122330210-3320220232323330-2023211121012113-0102330023131121-0301131131031131"></a>

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

<a id="canonical-1132102033330312-1320223323121320-1231322000301212-0320322010230031-3213300313010013-0302301130111003-1023131232102312-1323103203333033"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options`

<a id="canonical-0331123332311230-0031311100130211-1020031110210332-2121220120211113-1221212301311313-3011033301212101-3203130122012023-1201312300322323"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters

<a id="canonical-1131131123101310-3211220221112101-3303120231201300-3333222032231120-1302321000011333-3121122122102332-1210201002131200-1232031003233021"></a>

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

<a id="canonical-0013300330100322-2122130130032122-0122230200220130-2121323230132020-2013013122202233-1303220132033302-0311102031212020-0133030010020303"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters`

- [no_mtls](data-sources--workload--reference--group-024.md#canonical-1101201013233002-3121112201112023-2012100223011132-1133123311223320-2320002322310213-0122123303102203-0332223231203212-3210300133233302): complete subsection reference.

- [tls_certificates](data-sources--workload--reference--group-024.md#canonical-2301301320022131-3030321332300120-2133112211332222-1133202321312320-2003223211103012-1100102101302222-2230220032020100-0103112132001130): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-024.md#canonical-0002112221203130-3222001102200220-2032222020121321-3103132000033101-2032102220132111-0220231110310120-3113223131220200-2222002011303230): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-024.md#canonical-0231230331021312-3333101230213213-0310012020001323-2320030300101223-3001310131113102-0200031101113321-3330320112132230-2313300010123231): complete subsection reference.

<a id="canonical-1101201013233002-3121112201112023-2012100223011132-1133123311223320-2320002322310213-0122123303102203-0332223231203212-3210300133233302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-3011000231230110-3020100120002102-3012123313202202-0310312110331200-3222022120122231-1313223212301210-3001320003233331-0313221021012023"></a>

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

<a id="canonical-2301301320022131-3030321332300120-2133112211332222-1133202321312320-2003223211103012-1100102101302222-2230220032020100-0103112132001130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-0310130000122302-0313223023301013-0113111212333321-2013122332021330-2030123122103113-0013203333003302-0030213133223131-3330001210011133"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1222101301313312-0021120201001002-1102101030012323-3333021320330020-1310130030211021-1031132021021112-3321120102200002-2002301323313333"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates`

<a id="canonical-1231233222201111-1023010221331003-3112010103121311-1322222213102322-3011020302302231-2003031202102111-1032300111110300-0130120121003121"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.certificate_url` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [custom_hash_algorithms](data-sources--workload--reference--group-024.md#canonical-0002102113331302-3323200330021012-3333212022001322-1310023010211010-3322033021311122-2202333020013333-3002020011203331-2232313211332321): complete subsection reference.

<a id="canonical-2311232100021230-3123023103022222-1111203300230011-2323022313230131-0330232301221110-0023302331013303-3223210100311000-1231213330022211"></a>

<a id="canonical-0220200031120232-2200110103233031-3213111010303201-0023113111032201-1121322311113123-1201033011133200-2131201302303020-0102031322331023"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--workload--reference--group-024.md#canonical-0101303012023122-0213203233001223-3300133311102033-3301302020201132-3310001213223123-0300320331221313-0122133203212322-2320132230230101): complete subsection reference.

- [private_key](data-sources--workload--reference--group-024.md#canonical-1320032000001102-3003021300010230-2222021100012323-2220211330100010-3222220022332030-0233002322302332-2101010033222210-0212223203120101): complete subsection reference.

- [use_system_defaults](data-sources--workload--reference--group-024.md#canonical-0231033013322102-0100022313333332-3003032011012221-1212113032001213-1331113310202231-0223300022031020-1222301000101200-0222001003220111): complete subsection reference.

<a id="canonical-0002102113331302-3323200330021012-3333212022001322-1310023010211010-3322033021311122-2202333020013333-3002020011203331-2232313211332321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-024.md#canonical-2301301320022131-3030321332300120-2133112211332222-1133202321312320-2003223211103012-1100102101302222-2230220032020100-0103112132001130)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-0112013120000130-0002330223201322-2013100020032303-3102200100130333-3101212003001331-1132303001011303-2023002113333311-2323221202200223"></a>

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

<a id="canonical-0132231303112011-0102331131302231-3003112203122213-0031123023133223-3010330122330300-0223333132013130-1033101000203021-1321013011021322"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms`

<a id="canonical-0213112131100022-1012112130220302-2110333212311123-1331331021030123-3302013200203112-1231122020120103-2033322310332121-1131312320111100"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0101303012023122-0213203233001223-3300133311102033-3301302020201132-3310001213223123-0300320331221313-0122133203212322-2320132230230101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-024.md#canonical-2301301320022131-3030321332300120-2133112211332222-1133202321312320-2003223211103012-1100102101302222-2230220032020100-0103112132001130)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-3210212202113020-1320212120100101-0330013031110232-0203030202132020-3023111313312002-3331032030300230-3110313211200213-2313300020023122"></a>

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

<a id="canonical-1320032000001102-3003021300010230-2222021100012323-2220211330100010-3222220022332030-0233002322302332-2101010033222210-0212223203120101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-024.md#canonical-2301301320022131-3030321332300120-2133112211332222-1133202321312320-2003223211103012-1100102101302222-2230220032020100-0103112132001130)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-3020031121103122-2312100322233100-2300321233321322-3313212030231021-3322232333102321-2233212121331132-0233321202233323-3103113011233030"></a>

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

<a id="canonical-2121301003102232-3002200002310333-1103232331031200-0213201301121021-3013010012302313-2010333000212101-0122032100002201-3020031220232013"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--workload--reference--group-024.md#canonical-2302320331110300-3112230130313232-2001233032100130-0013322313110311-1013113012131022-3220322321230322-3101313011211010-2300123033130121): complete subsection reference.

- [clear_secret_info](data-sources--workload--reference--group-024.md#canonical-1321111001031033-0030300023222231-1120020202330112-2023122103220313-2002003212332112-3031212313320002-0220020031202123-0333123312133203): complete subsection reference.

<a id="canonical-2302320331110300-3112230130313232-2001233032100130-0013322313110311-1013113012131022-3220322321230322-3101313011211010-2300123033130121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-024.md#canonical-2301301320022131-3030321332300120-2133112211332222-1133202321312320-2003223211103012-1100102101302222-2230220032020100-0103112132001130)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-024.md#canonical-1320032000001102-3003021300010230-2222021100012323-2220211330100010-3222220022332030-0233002322302332-2101010033222210-0212223203120101)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-1121212201031302-2030022322223232-1031011303311130-1223202122323012-1100100203321013-0102230200110333-1131031011220021-2321213212100011"></a>

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

<a id="canonical-2233023233013010-3031133320113133-3113233232232130-0003131300022312-3001133303022202-0021221221131033-0012001223022123-3121211311210303"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-3023231002133030-2213112131212220-3120030002333332-1113113021213300-3333021312031220-1133300320230003-3303013313213202-1101103132133000"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0120232321002230-0020332331203103-2232020310232231-1011000131100231-3212023120001000-0203132130023200-3331331302232003-1212332202121300"></a>

<a id="canonical-2232223010201233-3200003022003220-2213232113331113-3023322320333122-2200211003311102-3212021022031330-0223313102032013-0201030131131121"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3133020021113231-1312301212321021-3032320220000332-0011122003001201-2111032110030100-0302201122201122-2330132221300010-1132112120320211"></a>

<a id="canonical-2022121330211013-0013120223210022-2123202322212200-1110232212202200-2320003003212333-3330021003233131-3131123210321200-2002033030200323"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1321111001031033-0030300023222231-1120020202330112-2023122103220313-2002003212332112-3031212313320002-0220020031202123-0333123312133203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-024.md#canonical-2301301320022131-3030321332300120-2133112211332222-1133202321312320-2003223211103012-1100102101302222-2230220032020100-0103112132001130)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-024.md#canonical-1320032000001102-3003021300010230-2222021100012323-2220211330100010-3222220022332030-0233002322302332-2101010033222210-0212223203120101)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-2332130231103203-2110211013311200-2121211100201332-0210031220003320-0102311023221012-1203121320322010-3312102003223310-1312023322130011"></a>

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

<a id="canonical-1201212011130200-3321221033102220-1131101131003023-1033132002310123-2123232023332022-1001331122300223-2002022301203231-3323320220012112"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info`

<a id="canonical-3201301322023102-3001133300023001-2113213022120031-0300332332222111-1221321000032112-3113202202101313-1110013110120221-0101023001331301"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0222231103032330-0233120112321313-3110300010332222-1020221323332031-3311011210110332-2111330210300230-1123231300320110-0212301101330313"></a>

<a id="canonical-1121122320121210-0011103121200130-0033113320313213-1112120102202313-2333200223023001-2020113112032203-0021233101013333-0010301111000030"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0231033013322102-0100022313333332-3003032011012221-1212113032001213-1331113310202231-0223300022031020-1222301000101200-0222001003220111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-024.md#canonical-2301301320022131-3030321332300120-2133112211332222-1133202321312320-2003223211103012-1100102101302222-2230220032020100-0103112132001130)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-3123231112032010-3313001210111210-3202131333203133-2313302102232312-0221020303030033-3230031213113011-3012321032221320-2200031130023112"></a>

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

<a id="canonical-0002112221203130-3222001102200220-2032222020121321-3103132000033101-2032102220132111-0220231110310120-3113223131220200-2222002011303230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-1003123102331113-3031133013200023-2023300223102111-3002210121201100-2321103021133033-1303130323303131-2110120133100010-0000320301101310"></a>

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

<a id="canonical-1111022011311222-3233311203001010-2203200030002232-3033312222010323-1202322023132213-3101013333133233-0000021110212303-2121333301111133"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config`

- [custom_security](data-sources--workload--reference--group-024.md#canonical-3312111303233330-1300311333002311-3331023113221232-1003001230301210-3201333202032330-0333330301012121-1210203312100000-3121103030310200): complete subsection reference.

- [default_security](data-sources--workload--reference--group-024.md#canonical-2121203120130231-0231200103311200-2132330033310320-2013213030232030-2023101222233310-2320110122231030-3100322002313123-2201301122333131): complete subsection reference.

- [low_security](data-sources--workload--reference--group-024.md#canonical-0300301323000203-2301201113230120-0230010131010220-3333011300302111-2210230201011312-2233003320123013-1102302101231323-2130020321033332): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-024.md#canonical-1230030003203322-3130312230220123-1021212203311011-1032230202120133-3332231122212321-2133222201310300-1320100222133231-3112321310313002): complete subsection reference.

<a id="canonical-3312111303233330-1300311333002311-3331023113221232-1003001230301210-3201333202032330-0333330301012121-1210203312100000-3121103030310200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-024.md#canonical-0002112221203130-3222001102200220-2032222020121321-3103132000033101-2032102220132111-0220231110310120-3113223131220200-2222002011303230)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-0102223113013113-1111321133233123-2332103331223203-0232130311323211-1301232222011033-3100000103112310-3201320220332300-0112313120301102"></a>

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

<a id="canonical-3212031310312133-3110202300221113-0331131330323001-3003001133033133-3112332002303021-2220333113101333-1100023333022103-0130032300022210"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security`

<a id="canonical-3212130002120031-1200312301132211-1031131223300023-1331113120132100-2200121033332212-3312113231033202-2102030110113030-0033333033201220"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0131203130011320-3100131120210330-2022222312322111-3000031210310031-1211202201220031-3210302131011032-3322000000212102-3102031013222102"></a>

<a id="canonical-0203113132111232-3330130021100311-2320232300311110-3202113101102020-0131002211113313-3110213230303323-1120012102322300-3322203333222311"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security.max_version` property

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

<a id="canonical-0201010311313002-1322121011300312-0222002211320121-0311232220010310-0033310022313001-0333320332320013-1132221223031333-1132122212000302"></a>

<a id="canonical-1103321333033331-1103110231200311-3021011333132001-1230211310112003-2011331333012111-3121222301023020-3113213312122023-3222331002300000"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security.min_version` property

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

<a id="canonical-2121203120130231-0231200103311200-2132330033310320-2013213030232030-2023101222233310-2320110122231030-3100322002313123-2201301122333131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-024.md#canonical-0002112221203130-3222001102200220-2032222020121321-3103132000033101-2032102220132111-0220231110310120-3113223131220200-2222002011303230)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-0110221122320033-3303001213320330-0322321001203131-0310303211302033-0302002313121102-0203323313023013-2031032311102320-0321300131213011"></a>

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

<a id="canonical-0300301323000203-2301201113230120-0230010131010220-3333011300302111-2210230201011312-2233003320123013-1102302101231323-2130020321033332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-024.md#canonical-0002112221203130-3222001102200220-2032222020121321-3103132000033101-2032102220132111-0220231110310120-3113223131220200-2222002011303230)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-0121023123013030-0132011021222002-0231301122233333-0310232011031331-3002113213021022-2133012111031332-1132233230220203-3032321230021300"></a>

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

<a id="canonical-1230030003203322-3130312230220123-1021212203311011-1032230202120133-3332231122212321-2133222201310300-1320100222133231-3112321310313002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-024.md#canonical-0002112221203130-3222001102200220-2032222020121321-3103132000033101-2032102220132111-0220231110310120-3113223131220200-2222002011303230)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-3332313312230000-3310022313321202-0331333232000311-3031131033120122-3322230123101021-2212331102323213-3200130203212011-3332023003300223"></a>

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

<a id="canonical-0231230331021312-3333101230213213-0310012020001323-2320030300101223-3001310131113102-0200031101113321-3330320112132230-2313300010123231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-0003033031203332-3220203013310120-3231202312100331-2221030003330213-3221032001201021-3030033133002103-0210313312003201-0233302013200120"></a>

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

<a id="canonical-2212110301231120-0230330031101321-0220221010330113-0022233120031303-0112020013013220-2133120000133132-3003310102202223-0231030201023321"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls`

<a id="canonical-3000211220333312-0000212113332130-1330211113231111-1033200111102201-0130322212203312-1230113113330112-2020202000202300-1232322002032333"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.client_certificate_optional` property

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

- [crl](data-sources--workload--reference--group-024.md#canonical-3213011110330110-2203232013310222-1200101202231300-0103113311200313-1312032330230020-0321123312333301-0013231033102000-3012333221023130): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-025.md#canonical-0032320022012020-1333113003320330-2300102003230231-3000311133020020-1032310222312100-2013232033301300-2332001210210001-0300331022301100): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-025.md#canonical-1010221000110112-2101231031301222-3212121303112313-0300111201100101-0123113113211302-2121121331212012-0023032203123212-0131123000310002): complete subsection reference.

<a id="canonical-0202131220022032-1300211332030010-1203100223021000-0222021332330211-0212333231122002-3213133020312210-1010102303023023-3131330122033113"></a>

<a id="canonical-0301032203331300-0313122121131113-3331203323102113-1003100031023202-3231211012213322-0100030133312303-2123231112332112-3102022133021112"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca_url` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [xfcc_disabled](data-sources--workload--reference--group-025.md#canonical-0331200310110320-1302010120001200-1032220122300213-1303212000232222-3203101003221210-3212111332132200-1023120122331301-1203330313012132): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-025.md#canonical-1322220012202300-2332001221123030-2321023200031200-0000311110311323-0232032000010231-1022032232003030-1001112033221010-1303211233212100): complete subsection reference.

<a id="canonical-3213011110330110-2203232013310222-1200101202231300-0103113311200313-1312032330230020-0321123312333301-0013231033102000-3012333221023130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-024.md#canonical-0231230331021312-3333101230213213-0310012020001323-2320030300101223-3001310131113102-0200031101113321-3330320112132230-2313300010123231)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-3103113212023001-3113233303213131-3231031223033330-1330122033231333-2133233212011233-3301203331112332-0112202022001331-1323013333112122"></a>

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

<a id="canonical-3120113232312212-1013202303101112-1030332203311301-1320103202131231-3001101302132322-1023133123313300-1201000123100313-0230220101002001"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl`

<a id="canonical-0201232121003311-0210312311333111-2310003130130203-2130200122332130-2033333310120313-3132103101133130-2311221200321022-1003233102021031"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2032233300032123-2002013133120110-1211012313333131-2123000203313330-2322012032201013-3220101221203102-0102323130113002-2210030203003110"></a>
