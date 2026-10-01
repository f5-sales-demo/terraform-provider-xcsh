---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-2200202330032001-2201313330332121-0320032021113312-3123230123222320-2023012133220010-0312031201333211-0112200303120113-0033021210033122"></a>

## name property — headers / 012220312021 / 6

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

<a id="canonical-3321302233231302-1302113203202010-2230033301221030-3330130003300222-3121220101021022-2030123213001033-1330101231330313-0203323232303203"></a>

<a id="canonical-2032332012021020-2011320300101111-0102013301212202-3000030031022003-0302032113121202-1332102322321211-1322113133223031-1130210031130003"></a>

## presence property — headers / 012220312021 / 7

Type: `"bool"`. Computed.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2003201321213322-2320012223030011-0213010301123211-2201000302023123-0200000201210011-2100010023103333-1101132011202211-2002311230210310"></a>

<a id="canonical-2032332103001332-2210301103022200-2131000322032321-1310313210220013-1023012200330232-1200103133030301-3022012000100300-2033120311112200"></a>

## regex property — headers / 012220312021 / 8

Type: `"string"`. Computed.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

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

<a id="canonical-0220201000111230-2201123202320032-3320200021033031-3020311201131133-3212322012200322-3111003223311333-2100233211011323-3332310123011300"></a>

## Next pages — headers / 012220312021 / 9

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-007.md#canonical-3113321211103300-3121123312121102-2103230121222320-0023030100232313-0030331001233112-3103020012203031-1211333322302002-3323010310210312)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1220223331102112-3311030110330320-3312022010122030-0323002023012210-2211120031010111-3202221103032312-3113313102103213-0311102310032123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100230113130233-1132300111201122-1000220312230213-0233032013222010-0000000231221130-3011130013133231-0232311122023031-3301231033100130"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port — incoming_port / 210201111232 / 2

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
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port

<a id="canonical-2331302323211003-3110222303103001-2231231002023301-0210101121322323-1332121031033323-0300300301003001-2023132123231301-0310330233132011"></a>

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

<a id="canonical-2313312120321032-2311323331123010-0001111232102203-3023321301303031-3301022011122012-1013332033130020-3003130203033031-1021112310300023"></a>

## Direct properties — incoming_port / 210201111232 / 3

- [no_port_match](data-sources--workload--reference--group-008.md#canonical-2001211111311033-1222220210301102-0100200031231121-1222030230303223-1013230103331323-3303013212300223-3011212010131132-1220332003033010): complete subsection reference.

<a id="canonical-3203203021230303-2232112201331010-3202100102030130-3323131310032212-1030033022002000-0113223002303022-3133112132010013-3012202030122301"></a>

<a id="canonical-0300213232031300-2201112312301101-1122220003100222-0101311201210213-2023232202030300-0210323231100111-1000031230112210-2000020030220031"></a>

## port property — incoming_port / 210201111232 / 4

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

<a id="canonical-0221202323102323-2123011130303001-1223311231002001-0333022132211012-0313013003301300-3221220200130033-0012123111230133-1113012322211122"></a>

<a id="canonical-0201131313120332-0123123121230033-1321210122023201-3213311112120020-1030031302021122-1222010120000011-1030233310331322-0310220001330020"></a>

## port_ranges property — incoming_port / 210201111232 / 5

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

<a id="canonical-0202010301012122-2031101323010111-3132021023013130-0300200213012310-2011300231331301-1301221033012101-2203111333133211-2121232121011103"></a>

## Next pages — incoming_port / 210201111232 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match](data-sources--workload--reference--group-008.md#canonical-2001211111311033-1222220210301102-0100200031231121-1222030230303223-1013230103331323-3303013212300223-3011212010131132-1220332003033010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-007.md#canonical-3113321211103300-3121123312121102-2103230121222320-0023030100232313-0030331001233112-3103020012203031-1211333322302002-3323010310210312)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2001211111311033-1222220210301102-0100200031231121-1222030230303223-1013230103331323-3303013212300223-3011212010131132-1220332003033010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012231321213222-1021201101100030-2011202311312131-3200011200202030-3010220102031311-0200030222030310-3133102203200301-1123133122323331"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match — no_port_match / 201332030322 / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](data-sources--workload--reference--group-008.md#canonical-1220223331102112-3311030110330320-3312022010122030-0323002023012210-2211120031010111-3202221103032312-3113313102103213-0311102310032123)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match

<a id="canonical-0321203022211013-3012130100111211-2323333322212111-2000231110023030-0201212033303222-0002322302013301-0032220022320033-0321101312231012"></a>

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

<a id="canonical-0303310330033013-2020032231200301-1320322013201233-2013033200212122-2003130021111133-3312201103131330-2201220012021033-0303023212011003"></a>

## Direct properties — no_port_match / 201332030322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101003102211121-2021033322203102-2123021211200032-1202201303133110-0312010301212132-2021000210020023-2022233022000213-1230320222310222"></a>

## Next pages — no_port_match / 201332030322 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](data-sources--workload--reference--group-008.md#canonical-1220223331102112-3311030110330320-3312022010122030-0323002023012210-2211120031010111-3202221103032312-3113313102103213-0311102310032123)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3112303221220112-3220200323222331-3031032321232031-3231213031300132-0222021122010321-3010200221120330-1303033103030000-0212303111301332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223212303213202-3100223103230323-0010122313012200-2332230132032330-0321020203020303-0110211201210210-3212031221230200-3003230122231331"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path — path / 131100300230 / 2

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
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path

<a id="canonical-2311031322232310-0210033330301313-0102302212022230-3133022323013033-3213002110010302-0003333320222330-2210320001131120-0011321133132101"></a>

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

<a id="canonical-0231131202120002-0133130321231203-3312121020101012-0100121233000123-1102000132313121-2110231000110103-2002112330133000-3220223313212222"></a>

## Direct properties — path / 131100300230 / 3

<a id="canonical-1313303201332102-0213230133121101-0132300023031223-3231120230013212-3230122131303203-3312320002103133-3111111211320312-0312223130000331"></a>

<a id="canonical-2102232121013100-2330211200103201-2021330013220322-1322330311220221-0303222102221203-2323332100122120-1201321100310333-0220200110203021"></a>

## path property — path / 131100300230 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-2303013023110032-2102232010231123-0012130123032310-1301312013203130-1322112100003100-3203323113223311-3011000300030220-2320230202103220"></a>

<a id="canonical-1302213303222121-2223200303101033-0300320333020200-3131110310030003-2311003302100301-0213210100320221-2003211023101101-2120102231311103"></a>

## prefix property — path / 131100300230 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-2303201322102013-1032122223001310-0300032030013230-2302323030133102-2102100210221032-2010332220111201-0332133212021222-3130323213312322"></a>

<a id="canonical-2101023332213212-0011100300222211-3101222101321323-3131320302112131-3210323111031232-3102202333013200-3231321031123200-2202131331030011"></a>

## regex property — path / 131100300230 / 6

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

<a id="canonical-1000031120212301-1012112103123033-1102132031323010-2131013103211330-0110313103302120-1003312010023011-1022122012323020-0200010311102003"></a>

## Next pages — path / 131100300230 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-007.md#canonical-3113321211103300-3121123312121102-2103230121222320-0023030100232313-0030331001233112-3103020012203031-1211333322302002-3323010310210312)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1233012303212313-3101123331213120-2210130313232300-3010333233222310-0331231001203323-1000212221033112-1323002011222232-1030300300323101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320102101222212-2302022102121133-2131023120230313-1132000131312230-1321332112120203-1330232302103033-3102031221301003-1003323120023113"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response — route_direct_response / 331332311200 / 2

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
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

<a id="canonical-2023022330200333-1111223103032122-1112121321230002-3120301110033232-3122203123232031-1020120313003232-2113222302211230-3102113100000300"></a>

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

<a id="canonical-2012332202300301-0130011120330302-1301101333332231-2010131111121113-2203113131112101-1202301111111330-0001133301011210-2232330110203120"></a>

## Direct properties — route_direct_response / 331332311200 / 3

<a id="canonical-3021012312102031-1132012203123121-1321221223101223-1111200022212220-2301233322031303-3331200131232331-2202322301321013-3003003212032110"></a>

<a id="canonical-0322103100232103-2230011200030212-1132032033000103-0323132030212121-2311300313302032-0310013233233302-0321201002211331-0211001130322100"></a>

## response_body_encoded property — route_direct_response / 331332311200 / 4

Type: `"string"`. Computed.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in Base64 format. The message can be either plain text or HTML.

Upstream description:

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in Base64 format. The message can be either plain text or HTML. E.g. "&lt;p&gt; Access
Denied &lt;/p&gt;". Base64 encoded string URL for this is
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

<a id="canonical-3201301031310032-1021222132332031-2320311012232001-1012211112203023-1320303003212121-1111132221322100-3202113211200123-0223230102032312"></a>

<a id="canonical-2033130302002113-1220130011111212-0103233113200310-3120332320010310-1321213102211202-1001131101303310-1202231330003202-0031122310222031"></a>

## response_code property — route_direct_response / 331332311200 / 5

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

<a id="canonical-3300003203202311-1012013330333012-1021101322022120-2112020230101100-1000310323223311-2100010311322121-0331010303212002-2031311201312200"></a>

## Next pages — route_direct_response / 331332311200 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-007.md#canonical-3113321211103300-3121123312121102-2103230121222320-0023030100232313-0030331001233112-3103020012203031-1211333322302002-3323010310210312)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112001111222131-0122020121312032-2212203120100010-1331010111101233-2301221331310233-0123301000232323-1333031010030202-3230123003310301"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route — redirect_route / 311223101122 / 2

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
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-3003323202321120-0230323212223020-0101102231001321-1032012232002121-3311233032110000-3200221013203101-0300010231121211-2023211302121113"></a>

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

<a id="canonical-1231321232233222-1221323323123010-3033200323330002-2213331301330132-3312323301120302-3213322033300023-2113011312021002-1200101323100231"></a>

## Direct properties — redirect_route / 311223101122 / 3

- [headers](data-sources--workload--reference--group-008.md#canonical-3303100213023330-1230312202133020-1001031231121212-1112011220021122-3010122020112100-3123312320220200-1302201130011331-3221230230323130): complete subsection reference.

<a id="canonical-0020121311120310-0130300113133231-3020301312202110-0101030231000310-1233011110300103-3302203201310233-2333023321311131-0133310302133023"></a>

<a id="canonical-1223102110110031-1111233210310322-3200112331003202-1003210133332321-3121013222020033-2320033102200002-2102222000001111-3121020103132123"></a>

## http_method property — redirect_route / 311223101122 / 4

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

- [incoming_port](data-sources--workload--reference--group-008.md#canonical-0203230113012003-3121132002233002-2001233201131311-3331123331010211-0210013332003122-2210120321310031-2301031011310030-3332203103302302): complete subsection reference.

- [path](data-sources--workload--reference--group-008.md#canonical-1310113111301010-2210003323310121-2301120032013120-0131032330221011-0133210022001300-1000331030120230-0001032300130033-3332232201201223): complete subsection reference.

- [route_redirect](data-sources--workload--reference--group-008.md#canonical-1313133332010100-0222322321110231-0031303330213222-1332232222223031-1101322302012120-3321000333301121-1321131021333010-2320221231230120): complete subsection reference.

<a id="canonical-3310120310003030-0212211030201031-0221011312132032-0032231012321031-1100331120332310-0302020311220030-0113023311321202-0000300022202313"></a>

## Next pages — redirect_route / 311223101122 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers](data-sources--workload--reference--group-008.md#canonical-3303100213023330-1230312202133020-1001031231121212-1112011220021122-3010122020112100-3123312320220200-1302201130011331-3221230230323130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-008.md#canonical-0203230113012003-3121132002233002-2001233201131311-3331123331010211-0210013332003122-2210120321310031-2301031011310030-3332203103302302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path](data-sources--workload--reference--group-008.md#canonical-1310113111301010-2210003323310121-2301120032013120-0131032330221011-0133210022001300-1000331030120230-0001032300130033-3332232201201223)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-008.md#canonical-1313133332010100-0222322321110231-0031303330213222-1332232222223031-1101322302012120-3321000333301121-1321131021333010-2320221231230120)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3303100213023330-1230312202133020-1001031231121212-1112011220021122-3010122020112100-3123312320220200-1302201130011331-3221230230323130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303120100302123-0302131012200120-0023000211332332-2121001002301101-0232001333113112-2331333133213012-0302111212030131-1302221233102130"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers — headers / 113001220310 / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-2300123323331132-0220210011123233-0231221302303011-0133202101310301-2323003130311031-0002233200212102-1220012122321030-0310000202201330"></a>

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

<a id="canonical-0222310322103031-3131333312232023-3202210201231231-0120113122113031-2011013121330030-3213003301112223-1203022022013331-2200311312002003"></a>

## Direct properties — headers / 113001220310 / 3

<a id="canonical-1333203021032023-1221110132202132-0232032023332112-3133120003213212-1310203311313222-1203001230122133-3013211201200313-1311322332311310"></a>

<a id="canonical-1133033000211100-2131031220102100-2222001222313021-1102233211313311-0101003110331313-0331021223030220-2021001010103101-0212301033110222"></a>

## exact property — headers / 113001220310 / 4

Type: `"string"`. Computed.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

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

<a id="canonical-2230303223013023-3331023220000001-3231221223030200-2130131331301010-0232333231222032-3023223002002131-0230020201133101-2332112203033220"></a>

<a id="canonical-3303212210332320-3313021202102130-0100122302302210-0212220123021010-2201013013020321-3131212100032020-1331211200332120-0302030301000212"></a>

## invert_match property — headers / 113001220310 / 5

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

<a id="canonical-2030112111203003-3212112302233101-2212130210212233-2120013002132230-1232010012331010-3113103311332123-2322022202110301-2211030322311303"></a>

<a id="canonical-3322030221132231-0301111103212301-3201021010211233-3230100012310333-1303033312113131-1320232302322333-0323132211233130-3131133112020012"></a>

## name property — headers / 113001220310 / 6

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

<a id="canonical-0111310221300102-2120321312300320-1001022202233010-3021331231302221-0322222120223300-1322330033213030-1020011030202202-1032103001012221"></a>

<a id="canonical-0221301311133131-0001212333301123-3022311232220203-3012020231322303-0322010102200302-1302132220233110-1320010201333202-1020220201133121"></a>

## presence property — headers / 113001220310 / 7

Type: `"bool"`. Computed.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1012200301110001-2012310320233200-1021232013221230-3110320201033311-0120013231103311-0133030313203231-1000022012231111-0032221320212120"></a>

<a id="canonical-1213002013320203-1010201113112030-2310321330022232-0100023120232202-2213121032202112-1222312203232033-1003012012331013-1301031121223233"></a>

## regex property — headers / 113001220310 / 8

Type: `"string"`. Computed.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

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

<a id="canonical-1113331323330103-0231220003321100-0332131111211101-3220102301020000-2233111000003002-0300023110013000-3122102021311300-2133113230120111"></a>

## Next pages — headers / 113001220310 / 9

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0203230113012003-3121132002233002-2001233201131311-3331123331010211-0210013332003122-2210120321310031-2301031011310030-3332203103302302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103113110201030-0300231131200003-2002021132123301-3120012233013103-1221003022233013-2330233313231301-1322302301101031-0120001123231321"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port — incoming_port / 110100002200 / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port

<a id="canonical-3032121101012233-2121233320230221-0131030101331302-0220220313223331-1031210130301001-3233130222222312-1110321210031110-2331101112021123"></a>

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

<a id="canonical-1302102020223322-3013233232220130-1330123110120300-1011303303203201-1031103001231232-2022122121310313-1302032333212123-0323113321132232"></a>

## Direct properties — incoming_port / 110100002200 / 3

- [no_port_match](data-sources--workload--reference--group-008.md#canonical-2210233331230201-3202222231103000-3111203200330120-1000010101133223-3001132311233120-3300330013100001-3011121302311210-2013010103233120): complete subsection reference.

<a id="canonical-3200221230112201-1011123030023213-1331220132030103-0321202103130012-2122223230123023-1233213302231222-3302122203313132-2230003231112220"></a>

<a id="canonical-1123111121102310-1332220013231121-2221200200112032-0331012020323011-1101303103103020-0021031112113031-3222223321010230-0022233303203000"></a>

## port property — incoming_port / 110100002200 / 4

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

<a id="canonical-0020302113103330-3223321212101012-1331211100221111-2313231222112033-0003002213002220-1222203020302101-0320230232222310-0112100321202012"></a>

<a id="canonical-1323222333330122-1333212323103000-3030130111200020-2030202332122222-0223132233201103-1232330313032201-1212112022100023-1131100310302110"></a>

## port_ranges property — incoming_port / 110100002200 / 5

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

<a id="canonical-3302132100222223-1110230330331231-2320211011320023-1010032001101122-3320212113213302-0023311133031133-1031120133123321-2210010030312033"></a>

## Next pages — incoming_port / 110100002200 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match](data-sources--workload--reference--group-008.md#canonical-2210233331230201-3202222231103000-3111203200330120-1000010101133223-3001132311233120-3300330013100001-3011121302311210-2013010103233120)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2210233331230201-3202222231103000-3111203200330120-1000010101133223-3001132311233120-3300330013100001-3011121302311210-2013010103233120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210321122301212-2101303232333333-1220133103010133-3330130101321310-1000121323223133-3322131000011103-0120202020302211-3022221123322122"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match — no_port_match / 003232313232 / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-008.md#canonical-0203230113012003-3121132002233002-2001233201131311-3331123331010211-0210013332003122-2210120321310031-2301031011310030-3332203103302302)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-0231023231131103-2010112002030113-2121331003202330-1300332112232012-0213033103033132-1232012020001121-3022103230201302-3232130311312120"></a>

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

<a id="canonical-3113123133132021-2201133210312123-3322001013120132-1220313110112131-2303222230331021-2222312103000111-2122110310203003-0231333012303230"></a>

## Direct properties — no_port_match / 003232313232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112031033122032-0233023230230033-2322221202323102-1000012123232212-0311203310121010-3123333101010220-1023033123213000-3111310301121222"></a>

## Next pages — no_port_match / 003232313232 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-008.md#canonical-0203230113012003-3121132002233002-2001233201131311-3331123331010211-0210013332003122-2210120321310031-2301031011310030-3332203103302302)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1310113111301010-2210003323310121-2301120032013120-0131032330221011-0133210022001300-1000331030120230-0001032300130033-3332232201201223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210313211221303-1123113231033102-3223001133121100-1233300022333011-3120010122221320-1333300122210100-0201332300101022-0131012023110231"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path — path / 210202312202 / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path

<a id="canonical-2322013300312230-0213321201032023-1332001202213312-0010222311321312-2301002322130012-2131021333333302-0322132031200220-3312003301332311"></a>

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

<a id="canonical-3013100130011121-2223012003303230-0211123121203100-1030003120321213-1211302120233230-1002200320333333-1031223021102332-3213303213003000"></a>

## Direct properties — path / 210202312202 / 3

<a id="canonical-3330122221232110-1230103100233031-0320120200013331-2130112203211131-1013210030323223-2233213213133222-3220013303130221-3221322211310331"></a>

<a id="canonical-0032032022123013-1030001220220012-1121220203120331-3110200312133222-3310032300131033-2100231033202332-1101222312132023-1101333123232010"></a>

## path property — path / 210202312202 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-3132301030201023-2012113100332120-0213323100303232-1312110310201230-3310302003103001-1303333101002002-2122003303301332-3121030301110300"></a>

<a id="canonical-0111131330030001-0123012220031331-3231123122121313-0320023230213021-3312100101312200-2202001132230032-2320110312233332-2012110212222131"></a>

## prefix property — path / 210202312202 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-3201001113031231-2111003123000001-2312231321321212-1003010112330023-2210011311000132-0222100020002303-1002213012311111-1101013002133112"></a>

<a id="canonical-3033103001203310-3303121301312202-2202212110330222-2020001001222030-1322010201132232-1210012020302232-1322032022003301-3320220230212001"></a>

## regex property — path / 210202312202 / 6

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

<a id="canonical-1112210313212323-3122331210131113-0232313210322020-0103221210210203-0023131101220000-2313222331121110-3111102121132020-1220031101103221"></a>

## Next pages — path / 210202312202 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1313133332010100-0222322321110231-0031303330213222-1332232222223031-1101322302012120-3321000333301121-1321131021333010-2320221231230120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331330221110313-2311010300103220-2013031113212332-2223120133121111-1300302331312022-3200122031110130-0032220211003130-3100311132213111"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect — route_redirect / 122303312330 / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect

<a id="canonical-2301323301311313-2331231302032030-2300112230223102-1023111132233211-0011012032033010-2202333111121330-2313002012220210-1232330130130301"></a>

Type: `"single"`. Computed.

Route redirect parameters when match action is redirect.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]",
  "x-ves-oneof-field-redirect_path_choice": "[\"path_redirect\",\"prefix_rewrite\"]"
}
```

<a id="canonical-2102321210113031-2111001021203002-2120200113003120-0301231012232320-0331122003310310-3102112132032211-3321100200100211-1110331112033031"></a>

## Direct properties — route_redirect / 122303312330 / 3

<a id="canonical-2013312312232011-1200132203332001-1222210312002233-1300132221320313-0222200333222321-2333031312330213-0033202013101120-2101123312120110"></a>

<a id="canonical-3231132331101303-1322133203033120-3101120200032122-1011212110111333-2121201000102322-3330210323313003-1211000220123012-2210031101020300"></a>

## host_redirect property — route_redirect / 122303312330 / 4

Type: `"string"`. Computed.

Swap host part of incoming URL in redirect URL.

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
  }
}
```

<a id="canonical-3131223310300130-2023330020310303-3232112230302212-2123313302300213-0023002022223203-2203223122220122-1120033300011013-2323031320122130"></a>

<a id="canonical-1110212120311011-2212303213312112-3003131112313321-0310022311013032-2003212321210222-3232032100232233-2031103122331212-3321020331011031"></a>

## path_redirect property — route_redirect / 122303312330 / 5

Type: `"string"`. Computed.

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

Upstream description:

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

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

<a id="canonical-3001312300203000-1211133111233232-0221300202222123-2301113333111203-0031000003112132-1132123300123321-1030321111103102-1232301302313100"></a>

<a id="canonical-0111311122120012-1211033010222220-0023010123131123-1000300310122321-0323313200021012-2210203031012000-3202110120220120-2301013003333211"></a>

## prefix_rewrite property — route_redirect / 122303312330 / 6

Type: `"string"`. Computed.

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

Upstream description:

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

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

<a id="canonical-0201123100110203-2213111322113310-3130002102013313-3003330200222122-0013200333100101-0110002220001232-3220013111321033-1013000013302203"></a>

<a id="canonical-3301033120133103-2222302012012120-1210103112230000-2003012213200222-3032202330022022-3130122021200033-1002111302301321-2011020112310001"></a>

## proto_redirect property — route_redirect / 122303312330 / 7

Type: `"string"`. Computed.

\[Enum: incoming-proto|http|https\] Swap protocol part of incoming URL in redirect URL The protocol
can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of
protocol is not done. Possible values are \`incoming-proto\`, \`http\`, \`https\`.

Upstream description:

Swap protocol part of incoming URL in redirect URL The protocol can be swapped with either HTTP or
HTTPS When incoming-proto option is specified, swapping of protocol is not done.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "incoming-proto",
    "http",
    "https"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  }
}
```

- [remove_all_params](data-sources--workload--reference--group-008.md#canonical-1231223201210110-2300312012213203-2111211220311001-3111231133301332-3123123302131303-2130311310101323-1222220330221211-0102332232032301): complete subsection reference.

<a id="canonical-1213032032000231-1033233102310003-2000220233133130-1110100230311322-2313013211010000-2132001211030323-2122320203021113-0213302022021222"></a>

<a id="canonical-1301303112301302-1330001002312000-2232203232232013-1201232320123121-0103010013310231-2003033230302221-1200331311010301-3102303220332211"></a>

## replace_params property — route_redirect / 122303312330 / 8

Type: `"string"`. Computed.

Exclusive with \[remove\_all\_params retain\_all\_params\].

Upstream description:

Exclusive with \[remove\_all\_params retain\_all\_params\]

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1221302312130322-2312031102323221-0011222111232012-1110320112101320-1333103223333320-0122331032320213-0220132203313012-2032220213221011"></a>

<a id="canonical-0223123231030312-2023302003303032-2103220223233312-1011331013313310-1100013313213033-2101130210202001-1200033212113321-1321120002032312"></a>

## response_code property — route_redirect / 122303312330 / 9

Type: `"number"`. Computed.

The HTTP status code to use in the redirect response.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

- [retain_all_params](data-sources--workload--reference--group-008.md#canonical-2013033210103222-1320000312030300-1030201120230023-1201110333323030-3323310312300332-1211202023111330-1212320303322030-2132022101001030): complete subsection reference.

<a id="canonical-0322113220303112-0102330103031023-3311103211103120-0003330223122031-2023202322010321-1310121011012221-0010303130002332-2122130013132120"></a>

## Next pages — route_redirect / 122303312330 / 10

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params](data-sources--workload--reference--group-008.md#canonical-1231223201210110-2300312012213203-2111211220311001-3111231133301332-3123123302131303-2130311310101323-1222220330221211-0102332232032301)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params](data-sources--workload--reference--group-008.md#canonical-2013033210103222-1320000312030300-1030201120230023-1201110333323030-3323310312300332-1211202023111330-1212320303322030-2132022101001030)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1231223201210110-2300312012213203-2111211220311001-3111231133301332-3123123302131303-2130311310101323-1222220330221211-0102332232032301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032002001333111-1333110120131111-1133223002101323-2022121122000210-2311312120331330-1223231212220301-0221012121331101-0311232222120231"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params — remove_all_params / 022103312333 / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-008.md#canonical-1313133332010100-0222322321110231-0031303330213222-1332232222223031-1101322302012120-3321000333301121-1321131021333010-2320221231230120)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-1032100023030233-0013012003030303-3302101230011323-3031011330312313-3323002223232222-1331223003131111-1323113333013232-1320213032200310"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for remove all params.

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

<a id="canonical-0021331300310022-2130203332310312-3101120001122231-1002123100031301-0222332130101132-2102313022232323-3222100011333020-2132031101130000"></a>

## Direct properties — remove_all_params / 022103312333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211323123002333-0320123133312013-2103120101203200-0002113021301302-3100311103031210-0032031133001231-3010023101313222-1021320301330300"></a>

## Next pages — remove_all_params / 022103312333 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-008.md#canonical-1313133332010100-0222322321110231-0031303330213222-1332232222223031-1101322302012120-3321000333301121-1321131021333010-2320221231230120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2013033210103222-1320000312030300-1030201120230023-1201110333323030-3323310312300332-1211202023111330-1212320303322030-2132022101001030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111121302221023-2111203201023101-0302130113231203-3133000031020000-0023202320001211-1132330111102021-3033303221333133-3320001111310312"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params — retain_all_params / 222300313103 / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-008.md#canonical-1313133332010100-0222322321110231-0031303330213222-1332232222223031-1101322302012120-3321000333301121-1321131021333010-2320221231230120)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-3332303011323101-2100302310301333-3132123110221332-2013101002301231-0311333001011203-2132111312100222-0111002010211302-1013031321110101"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for retain all params.

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

<a id="canonical-2011313011331113-2100220303000120-3022031132313003-3122332012031312-0230032321110220-2301131202330030-1231322100001103-1023213032010122"></a>

## Direct properties — retain_all_params / 222300313103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123301123232031-0301210300332213-2202231210232230-3101203121213001-1202012210201320-2112112212001001-2000201220021100-2103013013312132"></a>

## Next pages — retain_all_params / 222300313103 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-008.md#canonical-1313133332010100-0222322321110231-0031303330213222-1332232222223031-1101322302012120-3321000333301121-1321131021333010-2320221231230120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0310112122113210-2110103333320012-3100102321331121-2303102223302303-0033320103233333-1221123312232313-2132013023311033-1120032010201221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001102012320333-3123030223123112-0113103232231211-1331331212220103-0021001333032212-3003330103022020-3203112211131203-3123022111230001"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route — simple_route / 012120003000 / 2

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
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-0330201022113320-0323132220101330-2032312220323130-2301020233110231-2312102032133023-1320003233211201-1132303330101312-3233011221021322"></a>

Type: `"single"`. Computed.

Simple route matches on path and/or HTTP method and forwards the matching traffic to the default
origin pool specified outside.

Upstream description:

A simple route matches on path and/or HTTP method and forwards the matching traffic to the default
origin pool specified outside.

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

<a id="canonical-1132012131213003-0113001132313113-3213111022301222-2312112203302301-3111133113332021-0322031121310000-0233102320023303-1022330003000200"></a>

## Direct properties — simple_route / 012120003000 / 3

- [auto_host_rewrite](data-sources--workload--reference--group-008.md#canonical-3111112032111131-0201023202011221-3021010031211231-3031012110222322-2223032001201031-3331110213133032-1202332132112212-3313122313032122): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-008.md#canonical-0010313130021121-3113302100330202-3300213012332313-1133220113330323-1300322030033322-2203221132112213-2200302123222212-1020323223013031): complete subsection reference.

<a id="canonical-3003103113300323-2302303313222102-1210222101032213-0001033111212320-3200331101011223-3223010321002323-1132322022001020-1200331001130312"></a>

<a id="canonical-3001032322020310-2012211001000211-2331002323231030-0121020012010032-2131121010201203-2011211223231030-0322323212212223-2200212213323213"></a>

## host_rewrite property — simple_route / 012120003000 / 4

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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-0110212121221031-2101332313303121-3033013120023013-2123232122112303-1101013311000031-2210131310313231-0112303102123001-2220120012333003"></a>

<a id="canonical-0011101020230032-3223221133011112-3303130221332211-1332323300131012-1120100230301213-2023330031112200-1220031232321201-3021202312232201"></a>

## http_method property — simple_route / 012120003000 / 5

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

- [path](data-sources--workload--reference--group-008.md#canonical-1103333221333300-3332320302022022-2130011200321223-2022010220200112-1313121333320301-3230012303211132-0021110133131330-2323223100203030): complete subsection reference.

<a id="canonical-2101301313320322-3320133121011033-1323203010231111-2222203002223032-2120133100021322-2321012102011123-2202311133120210-3113332213330133"></a>

## Next pages — simple_route / 012120003000 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite](data-sources--workload--reference--group-008.md#canonical-3111112032111131-0201023202011221-3021010031211231-3031012110222322-2223032001201031-3331110213133032-1202332132112212-3313122313032122)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite](data-sources--workload--reference--group-008.md#canonical-0010313130021121-3113302100330202-3300213012332313-1133220113330323-1300322030033322-2203221132112213-2200302123222212-1020323223013031)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path](data-sources--workload--reference--group-008.md#canonical-1103333221333300-3332320302022022-2130011200321223-2022010220200112-1313121333320301-3230012303211132-0021110133131330-2323223100203030)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3111112032111131-0201023202011221-3021010031211231-3031012110222322-2223032001201031-3331110213133032-1202332132112212-3313122313032122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133322033132130-0311300230210022-2220323321113222-3120012120132303-1002012323231330-1323002001232011-2211231111203232-0211010230221013"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite — auto_host_rewrite / 111330300313 / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-008.md#canonical-0310112122113210-2110103333320012-3100102321331121-2303102223302303-0033320103233333-1221123312232313-2132013023311033-1120032010201221)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-2100120130130201-2023123221002110-3330001131021233-1230021232003203-3330330310312201-0103121010203202-0210023200110212-3320113301101223"></a>

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

<a id="canonical-3313200313230332-1220212132313030-2223001331320322-1022032300110312-1211021110023330-0102213132221013-0013020220211011-2222221022333123"></a>

## Direct properties — auto_host_rewrite / 111330300313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201210131120322-0012221322212300-2022331110112101-1132012023223113-0102121202120102-3013230202321133-1210113031322002-3230220033000312"></a>

## Next pages — auto_host_rewrite / 111330300313 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-008.md#canonical-0310112122113210-2110103333320012-3100102321331121-2303102223302303-0033320103233333-1221123312232313-2132013023311033-1120032010201221)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0010313130021121-3113302100330202-3300213012332313-1133220113330323-1300322030033322-2203221132112213-2200302123222212-1020323223013031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332330200313231-0100312020013132-1231202100330010-1301211112322123-3300103102201002-1130010020012121-0313021213202102-2333320203231133"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite — disable_host_rewrite / 012313002010 / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-008.md#canonical-0310112122113210-2110103333320012-3100102321331121-2303102223302303-0033320103233333-1221123312232313-2132013023311033-1120032010201221)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-0003113333102010-3030031202112031-2323001011221223-2020213212110211-1012103023212102-1222211021000133-1020302130130120-0020133113133203"></a>

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

<a id="canonical-3111321312331031-2010021333313132-2332333020102221-0233203120211002-2233332201310011-1213303110231231-2202312220020113-0201332322222221"></a>

## Direct properties — disable_host_rewrite / 012313002010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233311100023020-3311322103022303-3333022231211321-3002220220121120-0201330331110200-0300310212132212-2323330003033100-1312321311213021"></a>

## Next pages — disable_host_rewrite / 012313002010 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-008.md#canonical-0310112122113210-2110103333320012-3100102321331121-2303102223302303-0033320103233333-1221123312232313-2132013023311033-1120032010201221)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1103333221333300-3332320302022022-2130011200321223-2022010220200112-1313121333320301-3230012303211132-0021110133131330-2323223100203030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121313320033332-3221131101010300-2133030333320223-0123332231332022-3123313332212322-3222321120132231-2103102113302330-1012311021301233"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path — path / 312312001012 / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-008.md#canonical-0310112122113210-2110103333320012-3100102321331121-2303102223302303-0033320103233333-1221123312232313-2132013023311033-1120032010201221)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-2100111313032331-3103210221100111-1330230013102111-2313203133022102-2013013202032213-3220312103202023-3102110111110213-3213311131300202"></a>

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

<a id="canonical-1303223312130213-1022210202233230-2012122110312223-2023233210201333-2100103020012100-0220331201131233-1230110111231110-3230133310321323"></a>

## Direct properties — path / 312312001012 / 3

<a id="canonical-3233300331021111-1112322113002212-1112113103033011-0203230032211012-1101021003010312-1100030130331013-3310221123132201-1331000232301310"></a>

<a id="canonical-3302222302133011-0303332013102112-1312210332301320-0030210122332201-1013301230103233-1010303310033331-0202230112031003-2202101100012202"></a>

## path property — path / 312312001012 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-3133213213200121-0321132101330200-1202230111011210-3003213000231102-3011130001333211-0230310031221002-1101131003321231-3300202202121323"></a>

<a id="canonical-0201120113231302-2101231231210230-0210013332032200-0322221200122002-0213023233122133-3301003300302012-2001133230032221-3231300110102310"></a>

## prefix property — path / 312312001012 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-0323213300130002-1113230322320310-2223300212000023-1323222111101121-0200032301130113-3311211000101113-2022121303113201-0311131311332302"></a>

<a id="canonical-0322212023301211-2232321030033231-2023221300133132-0122333132301131-1302101310322022-2303200201131230-1230001321103301-2131211101023130"></a>

## regex property — path / 312312001012 / 6

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

<a id="canonical-3211133212222303-0320330310232300-3203001032131030-2121132030002101-2332133030000022-3031131331220203-3112000130100003-2132212212232311"></a>

## Next pages — path / 312312001012 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-008.md#canonical-0310112122113210-2110103333320012-3100102321331121-2303102223302303-0033320103233333-1221123312232313-2132013023311033-1120032010201221)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1330120103311032-3323002022332313-1230010032111002-0301100332330310-1013300322113000-3020110122120300-1230010001312321-1311000201122222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202203313223131-3030000101123110-2331121211203333-1231023100132122-2121300333303330-2101232100200131-3120230310200233-2332031102021302"></a>

## service.advertise_options.advertise_custom.ports.port — port / 332300103010 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- service.advertise_options.advertise_custom.ports.port

<a id="canonical-0212102122101321-0101022310120302-3320221011210112-1212200300103123-2311003102102002-3200021323130221-0013003322311123-3200133130232201"></a>

Type: `"single"`. Computed.

Port. Port of the workload.

Upstream description:

Port of the workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1011303032331010-0023101010223101-1311211030310131-2311002020200313-2100212110131230-2031133233202300-2020221213032210-3102131201122321"></a>

## Direct properties — port / 332300103010 / 3

- [info](data-sources--workload--reference--group-008.md#canonical-0110122133233020-0013313330132130-1321111101003321-0211310020013220-0110220200320011-1003211320330303-1033031311131333-2211112312023000): complete subsection reference.

<a id="canonical-1322201112212023-0332000202310201-0032232311303322-2010300122201011-1223120313323221-0010002031112123-1133230010022012-0331012101003220"></a>

<a id="canonical-0323132031030023-1220110322210133-2231101002222231-2020213232210323-2312033132300012-1320031233022032-2212130133113033-1023023202222211"></a>

## name property — port / 332300103010 / 4

Type: `"string"`. Computed.

Name. Name of the Port.

Upstream description:

Name of the Port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-1300010013110010-3300101322222230-2101003121323132-2020232103303002-0112230221312123-0213231023212322-2203311012302232-3003212333131230"></a>

## Next pages — port / 332300103010 / 5

- [service.advertise_options.advertise_custom.ports.port.info](data-sources--workload--reference--group-008.md#canonical-0110122133233020-0013313330132130-1321111101003321-0211310020013220-0110220200320011-1003211320330303-1033031311131333-2211112312023000)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0110122133233020-0013313330132130-1321111101003321-0211310020013220-0110220200320011-1003211320330303-1033031311131333-2211112312023000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033201001301332-1111003111022323-1302022123003322-2311303020333213-0202321222230312-1000112130022313-1121001320331230-1320331113123203"></a>

## service.advertise_options.advertise_custom.ports.port.info — info / 113031303133 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.port](data-sources--workload--reference--group-008.md#canonical-1330120103311032-3323002022332313-1230010032111002-0301100332330310-1013300322113000-3020110122120300-1230010001312321-1311000201122222)
- service.advertise_options.advertise_custom.ports.port.info

<a id="canonical-2201301010121010-1011300020231210-2031003001213330-3000121031133210-0103230223011030-1010033320322210-0121122003203111-0331121200302330"></a>

Type: `"single"`. Computed.

Port Information. Port information.

Upstream description:

Port information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

<a id="canonical-0030231132103322-1033211023032301-1232311300321123-1201110220012012-1131100203330323-3230303331333010-1313010133100102-1213300023330323"></a>

## Direct properties — info / 113031303133 / 3

<a id="canonical-3302201320013322-2030331201100300-0211300301110321-1131312030100311-2121120320012223-2310221001213221-2012113223312112-3320331231012031"></a>

<a id="canonical-1312132120220322-1221110331000010-3211223100322333-0201120210331000-1312120023302033-3231112113131003-1021130213222132-3013030122231111"></a>

## port property — info / 113031303133 / 4

Type: `"number"`. Computed.

Port. Port the workload can be reached on.

Upstream description:

Port the workload can be reached on.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3023313222212202-2330100111311132-0102302312132312-0231130111311113-0223233002122013-3113010122321320-0330101101230131-3330223332311231"></a>

<a id="canonical-3031011131112222-3021130002003233-0310132333330331-2111332011321212-3012200231221201-1121001223300120-1122010013223313-3331310112212233"></a>

## protocol property — info / 113031303133 / 5

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

TCP &#8203;- PROTOCOL\_HTTP: HTTP

HTTP &#8203;- PROTOCOL\_HTTP2: HTTP2

HTTP2 &#8203;- PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI

TLS with SNI &#8203;- PROTOCOL\_UDP: UDP

UDP.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [same_as_port](data-sources--workload--reference--group-008.md#canonical-1210011231223112-0112103212001003-2023321021321303-2232320101010331-0013332100311120-0332013310112322-1120132310101300-2120031112211310): complete subsection reference.

<a id="canonical-3323332032320001-0212100233002230-1103331323032130-0303201120111021-1002133122110030-3201123233301130-1102203202232003-2202332303132230"></a>

<a id="canonical-0222311321333130-2103333320121332-2312230203323031-2003323023100310-2030130121201331-3010221301321331-1130233210002332-3332220013311200"></a>

## target_port property — info / 113031303133 / 6

Type: `"number"`. Computed.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Upstream description:

Exclusive with \[same\_as\_port\] Port the workload is listening on.

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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1230032320110203-3230113311020333-3311102231011232-3232120121130212-1101100302301110-0110310220332103-0012212332222233-2023103300203331"></a>

## Next pages — info / 113031303133 / 7

- [service.advertise_options.advertise_custom.ports.port.info.same_as_port](data-sources--workload--reference--group-008.md#canonical-1210011231223112-0112103212001003-2023321021321303-2232320101010331-0013332100311120-0332013310112322-1120132310101300-2120031112211310)
- [service.advertise_options.advertise_custom.ports.port](data-sources--workload--reference--group-008.md#canonical-1330120103311032-3323002022332313-1230010032111002-0301100332330310-1013300322113000-3020110122120300-1230010001312321-1311000201122222)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1210011231223112-0112103212001003-2023321021321303-2232320101010331-0013332100311120-0332013310112322-1120132310101300-2120031112211310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232121013332201-3020231300121133-1032201030221130-0121011210120020-2032121132103022-0211103301130232-2333133322113311-3002013313233330"></a>

## service.advertise_options.advertise_custom.ports.port.info.same_as_port — same_as_port / 231020313213 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.port](data-sources--workload--reference--group-008.md#canonical-1330120103311032-3323002022332313-1230010032111002-0301100332330310-1013300322113000-3020110122120300-1230010001312321-1311000201122222)
- [service.advertise_options.advertise_custom.ports.port.info](data-sources--workload--reference--group-008.md#canonical-0110122133233020-0013313330132130-1321111101003321-0211310020013220-0110220200320011-1003211320330303-1033031311131333-2211112312023000)
- service.advertise_options.advertise_custom.ports.port.info.same_as_port

<a id="canonical-2210221320120312-3321022021333223-0323113332103322-1233000122302001-2121300220310021-0002203123333001-1323200333210210-1131201203023212"></a>

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

<a id="canonical-0030312321031313-3333122000023121-2312230011312022-3311201212000312-2210200103032100-1102033103113333-1101213023311131-1103010210101020"></a>

## Direct properties — same_as_port / 231020313213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012321233330303-1313331203032202-2212302110033131-0220132220320121-1022222022132010-0012230010201001-2021012213020223-0022303003230333"></a>

## Next pages — same_as_port / 231020313213 / 4

- [service.advertise_options.advertise_custom.ports.port.info](data-sources--workload--reference--group-008.md#canonical-0110122133233020-0013313330132130-1321111101003321-0211310020013220-0110220200320011-1003211320330303-1033031311131333-2211112312023000)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1213013231333121-3300100100220230-1131111132003220-3022122220020203-0200132131130103-3220333332311300-2220212322331102-1331120231121131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321011133200303-1123213210023002-1322302000212310-3200021001203102-2131200002010300-1032313322113133-0321003203301313-3132002211310213"></a>

## service.advertise_options.advertise_custom.ports.tcp_loadbalancer — tcp_loadbalancer / 032010013222 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- service.advertise_options.advertise_custom.ports.tcp_loadbalancer

<a id="canonical-2211133012131001-2203212223030313-0131102131302203-3010322330230031-1102211232103122-1131110001212311-2123300311000321-1321223213100021"></a>

Type: `"single"`. Computed.

Configuration parameter for tcp loadbalancer.

Upstream description:

TCP loadbalancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2001120033222022-0223103231210100-3103330022203220-3232333201223333-3131113113213210-0311001201320203-1330232322022300-2003313202002121"></a>

## Direct properties — tcp_loadbalancer / 032010013222 / 3

<a id="canonical-0232330130313333-0302113210312332-0102003133012312-0102310200210112-2303212021323222-0232033020001130-2123132313210323-0033220000213201"></a>

<a id="canonical-2311233333001201-0111112220203130-3322310100001222-3113001220001132-2231113123331321-3233230300112300-0213213012322323-2233001003321132"></a>

## domains property — tcp_loadbalancer / 032010013222 / 4

Type: `["list", "string"]`. Computed.

List of additional domains (host/authority header) that will be matched to this loadbalancer.
Domains are also used for SNI matching if the is true Domains also indicate the list of names for
which DNS resolution will be done by VER.

Upstream description:

A list of additional domains (host/authority header) that will be matched to this loadbalancer.

Domains are also used for SNI matching if the \`with\_sni\` is true Domains also indicate the list
of names for which DNS resolution will be done by VER.

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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3033120123312320-2021023120031103-3012223211200103-1223233001233102-0002130203232310-2031131132223201-3320200313313113-1201232232120212"></a>

<a id="canonical-2130210122011221-3200312212103113-1002023331333300-0330132233010011-2111001220320310-0001223010113011-0232302311303030-0221102020101333"></a>

## with_sni property — tcp_loadbalancer / 032010013222 / 5

Type: `"bool"`. Computed.

Set to true to enable TCP loadbalancer with SNI.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1213221212130210-3230200310213222-1220302223013213-0312200311003222-1012023322022121-2021033322221211-3022030201112112-2133303301210002"></a>

## Next pages — tcp_loadbalancer / 032010013222 / 6

- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2131020211211101-0213303111310032-0232312121101303-0303013213123313-0112311302213302-2012200033230321-1231231303121203-0121103301113123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123212303301223-0322131120120210-0223303032221200-0200211231123330-0333100311320103-3221233302302132-1201102230112302-2201030031011310"></a>

## service.advertise_options.advertise_in_cluster — advertise_in_cluster / 200131120032 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- service.advertise_options.advertise_in_cluster

<a id="canonical-1020200221333110-1133133033310022-3311231022333203-2300002231012130-2232113321103211-0221303322003212-3322211203120102-0210030220211110"></a>

Type: `"single"`. Computed.

Advertise the workload locally in-cluster.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"multi_ports\",\"port\"]"
}
```

<a id="canonical-3323001332031302-3022030133210000-3231022330311312-2322031323032321-1000303223232333-0311313330230221-1311311232021213-0033321322121010"></a>

## Direct properties — advertise_in_cluster / 200131120032 / 3

- [multi_ports](data-sources--workload--reference--group-008.md#canonical-0323130202203203-0212323222122012-0133122332003031-0311312333002330-0110212200031012-3030321303133231-2322211032002200-0003101011302311): complete subsection reference.

- [port](data-sources--workload--reference--group-008.md#canonical-2223131103313322-0013111330301232-2313200302013312-2300312202002333-1313130002133111-1321232321000201-1003101301002233-3322012010032123): complete subsection reference.

<a id="canonical-2320011120222200-0312132321223023-3302001120133222-3220301303030201-3302000010031322-0000220002123311-2121101020131122-1021011010101010"></a>

## Next pages — advertise_in_cluster / 200131120032 / 4

- [service.advertise_options.advertise_in_cluster.multi_ports](data-sources--workload--reference--group-008.md#canonical-0323130202203203-0212323222122012-0133122332003031-0311312333002330-0110212200031012-3030321303133231-2322211032002200-0003101011302311)
- [service.advertise_options.advertise_in_cluster.port](data-sources--workload--reference--group-008.md#canonical-2223131103313322-0013111330301232-2313200302013312-2300312202002333-1313130002133111-1321232321000201-1003101301002233-3322012010032123)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0323130202203203-0212323222122012-0133122332003031-0311312333002330-0110212200031012-3030321303133231-2322211032002200-0003101011302311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100102323320300-0310123212133030-2300003210303130-3112013233113000-1022030102213201-1202202120131032-0130332303013333-3331021211230220"></a>

## service.advertise_options.advertise_in_cluster.multi_ports — multi_ports / 123232303022 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-2131020211211101-0213303111310032-0232312121101303-0303013213123313-0112311302213302-2012200033230321-1231231303121203-0121103301113123)
- service.advertise_options.advertise_in_cluster.multi_ports

<a id="canonical-2121221332110112-0023100121302222-1133101033331001-2013030311103331-3120133231303310-1021322202132130-1203120023020132-1011100111331111"></a>

Type: `"single"`. Computed.

Multiple Ports. Multiple ports.

Upstream description:

Multiple ports.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1322331013002330-3113201133202001-0110233223031020-2031323233302332-2311023001123323-2001203123203113-0310331321222203-3001320031310300"></a>

## Direct properties — multi_ports / 123232303022 / 3

- [ports](data-sources--workload--reference--group-008.md#canonical-1200123030213102-3312221012212330-2122212021203232-3111213102133231-3300123331001200-1300033313011233-0113120231301331-3210113023313323): complete subsection reference.

<a id="canonical-0010322332333322-0133031231323332-1201000222211132-3303230330032101-3023312232223130-2012333221021303-1332111301331221-1320001002203100"></a>

## Next pages — multi_ports / 123232303022 / 4

- [service.advertise_options.advertise_in_cluster.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-1200123030213102-3312221012212330-2122212021203232-3111213102133231-3300123331001200-1300033313011233-0113120231301331-3210113023313323)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-2131020211211101-0213303111310032-0232312121101303-0303013213123313-0112311302213302-2012200033230321-1231231303121203-0121103301113123)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1200123030213102-3312221012212330-2122212021203232-3111213102133231-3300123331001200-1300033313011233-0113120231301331-3210113023313323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011021313032032-1132200231013222-0002321232313323-1102221031332321-0221012333111230-1323001001230033-2002133313022221-2300320030002310"></a>

## service.advertise_options.advertise_in_cluster.multi_ports.ports — ports / 022311201322 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-2131020211211101-0213303111310032-0232312121101303-0303013213123313-0112311302213302-2012200033230321-1231231303121203-0121103301113123)
- [service.advertise_options.advertise_in_cluster.multi_ports](data-sources--workload--reference--group-008.md#canonical-0323130202203203-0212323222122012-0133122332003031-0311312333002330-0110212200031012-3030321303133231-2322211032002200-0003101011302311)
- service.advertise_options.advertise_in_cluster.multi_ports.ports

<a id="canonical-3130302131113211-1233131322302313-3212131122223013-3301220312213012-2311101311012101-1330333013232310-1100210111220030-2013213121001323"></a>

Type: `"list"`. Computed.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2330022203330010-1120220031122010-0000013102231323-0221012330003123-3211222033212332-1232201312030301-1110020022010203-0332320102011121"></a>

## Direct properties — ports / 022311201322 / 3

- [info](data-sources--workload--reference--group-008.md#canonical-1311013020110322-3332312222230230-2111331120112121-1302023321211110-3121210102132002-3331300220020203-0032132323002103-2013103230232311): complete subsection reference.

<a id="canonical-2111222223331201-2120012010100122-3222203222123021-3333131113021033-3301301130201301-3200313132121311-1130002003221112-2131201333333221"></a>

<a id="canonical-2100000222301032-3122030211012201-0223213220321210-2320033011132201-0130001030231130-0112320102312322-3312313210112031-3330223310112002"></a>

## name property — ports / 022311201322 / 4

Type: `"string"`. Computed.

Name. Name of the Port.

Upstream description:

Name of the Port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-3223013211110133-3201032002101011-0212131203031331-0011023223232333-0011311130321020-3100211313032012-3232310132003322-3012011321301110"></a>

## Next pages — ports / 022311201322 / 5

- [service.advertise_options.advertise_in_cluster.multi_ports.ports.info](data-sources--workload--reference--group-008.md#canonical-1311013020110322-3332312222230230-2111331120112121-1302023321211110-3121210102132002-3331300220020203-0032132323002103-2013103230232311)
- [service.advertise_options.advertise_in_cluster.multi_ports](data-sources--workload--reference--group-008.md#canonical-0323130202203203-0212323222122012-0133122332003031-0311312333002330-0110212200031012-3030321303133231-2322211032002200-0003101011302311)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1311013020110322-3332312222230230-2111331120112121-1302023321211110-3121210102132002-3331300220020203-0032132323002103-2013103230232311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200213332230031-1120031213201103-3223031110023122-3012103121030212-0032332013312130-3033121131021130-1133113231020131-0001303112302323"></a>

## service.advertise_options.advertise_in_cluster.multi_ports.ports.info — info / 002012322300 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-2131020211211101-0213303111310032-0232312121101303-0303013213123313-0112311302213302-2012200033230321-1231231303121203-0121103301113123)
- [service.advertise_options.advertise_in_cluster.multi_ports](data-sources--workload--reference--group-008.md#canonical-0323130202203203-0212323222122012-0133122332003031-0311312333002330-0110212200031012-3030321303133231-2322211032002200-0003101011302311)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-1200123030213102-3312221012212330-2122212021203232-3111213102133231-3300123331001200-1300033313011233-0113120231301331-3210113023313323)
- service.advertise_options.advertise_in_cluster.multi_ports.ports.info

<a id="canonical-1123230230111032-0300201311021033-3121223200012230-3221032231113330-3001322102233113-3030032132022221-1112330300333201-3130013131010023"></a>

Type: `"single"`. Computed.

Port Information. Port information.

Upstream description:

Port information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

<a id="canonical-2031332222013012-0002002120022112-3021333111010021-3113111321221232-3222202122323331-1201010331231031-2313112211022123-0311322013020003"></a>

## Direct properties — info / 002012322300 / 3

<a id="canonical-3012332131231010-2310032113130330-2020030232311031-0013031010110333-3303013323202133-2322012013312102-2013001331021201-0122220020320113"></a>

<a id="canonical-2303312011123020-2013130123232231-2130023231020310-3331030120012233-0203000333332002-1303223111333133-1211112222003112-2300023003330313"></a>

## port property — info / 002012322300 / 4

Type: `"number"`. Computed.

Port. Port the workload can be reached on.

Upstream description:

Port the workload can be reached on.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1210130322123300-3033032321302210-0112010303001330-1332321303220322-0311111022231213-2330133002132000-2320330022013303-3221110113233311"></a>

<a id="canonical-2020010333112302-3023320310112310-1333221101113202-0023232310200021-2120123311232120-1011330230200100-0103112012302123-3303031021103013"></a>

## protocol property — info / 002012322300 / 5

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

TCP &#8203;- PROTOCOL\_HTTP: HTTP

HTTP &#8203;- PROTOCOL\_HTTP2: HTTP2

HTTP2 &#8203;- PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI

TLS with SNI &#8203;- PROTOCOL\_UDP: UDP

UDP.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [same_as_port](data-sources--workload--reference--group-008.md#canonical-3312230321301022-3201311033220031-0223210311200033-3110222212213013-3312031030333223-0311310302230130-1220121101111132-1003020123013313): complete subsection reference.

<a id="canonical-0321231311111130-3223320022223000-2333010133101221-3303323230220202-1233102100231030-0133201210123212-2312121222302003-2230021312120122"></a>

<a id="canonical-2133303010010233-3300323102331013-2311023230133333-1102103122300211-3000302201031332-0231113001322300-2201313121012232-3110230023120321"></a>

## target_port property — info / 002012322300 / 6

Type: `"number"`. Computed.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Upstream description:

Exclusive with \[same\_as\_port\] Port the workload is listening on.

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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3223020103303311-2223110011021131-2202133110210012-0030300212312122-3213003000102033-2012112233102212-2121002033311010-1110313332030101"></a>

## Next pages — info / 002012322300 / 7

- [service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port](data-sources--workload--reference--group-008.md#canonical-3312230321301022-3201311033220031-0223210311200033-3110222212213013-3312031030333223-0311310302230130-1220121101111132-1003020123013313)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-1200123030213102-3312221012212330-2122212021203232-3111213102133231-3300123331001200-1300033313011233-0113120231301331-3210113023313323)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3312230321301022-3201311033220031-0223210311200033-3110222212213013-3312031030333223-0311310302230130-1220121101111132-1003020123013313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012310303023030-0030111023103211-2330221233201021-0112232010202230-1101012333210333-0122121322130130-0221213231222121-2100331101123131"></a>

## service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port — same_as_port / 333130030132 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-2131020211211101-0213303111310032-0232312121101303-0303013213123313-0112311302213302-2012200033230321-1231231303121203-0121103301113123)
- [service.advertise_options.advertise_in_cluster.multi_ports](data-sources--workload--reference--group-008.md#canonical-0323130202203203-0212323222122012-0133122332003031-0311312333002330-0110212200031012-3030321303133231-2322211032002200-0003101011302311)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-1200123030213102-3312221012212330-2122212021203232-3111213102133231-3300123331001200-1300033313011233-0113120231301331-3210113023313323)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports.info](data-sources--workload--reference--group-008.md#canonical-1311013020110322-3332312222230230-2111331120112121-1302023321211110-3121210102132002-3331300220020203-0032132323002103-2013103230232311)
- service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port

<a id="canonical-1222301312020211-1302010021230103-0231103033233310-1002232210001132-3223001323133202-1131333101303121-3311213102330113-2321300121131203"></a>

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

<a id="canonical-1310220213330011-1123313310303301-1001122223222001-0220333332133013-3222332300132202-1312121120013222-0332331111330002-0312321003021102"></a>

## Direct properties — same_as_port / 333130030132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023103312123020-0202113223123231-2131122130012210-0031322100320312-0021223003021320-2301132333303221-3200122033201211-1213100130032330"></a>

## Next pages — same_as_port / 333130030132 / 4

- [service.advertise_options.advertise_in_cluster.multi_ports.ports.info](data-sources--workload--reference--group-008.md#canonical-1311013020110322-3332312222230230-2111331120112121-1302023321211110-3121210102132002-3331300220020203-0032132323002103-2013103230232311)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2223131103313322-0013111330301232-2313200302013312-2300312202002333-1313130002133111-1321232321000201-1003101301002233-3322012010032123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010033113232002-3033131033331003-2232320301202211-3131312103301332-1133013222312313-3331300003301310-1312123132110120-2131002032303210"></a>

## service.advertise_options.advertise_in_cluster.port — port / 210001230332 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-2131020211211101-0213303111310032-0232312121101303-0303013213123313-0112311302213302-2012200033230321-1231231303121203-0121103301113123)
- service.advertise_options.advertise_in_cluster.port

<a id="canonical-0322022311102321-3231000101121113-1200022100320232-1113222311021132-0122110023312332-3003203312221113-3112130302113310-0211201320320023"></a>

Type: `"single"`. Computed.

Port. Single port.

Upstream description:

Single port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2233022113300230-1330011020220112-0330010131233321-2313230331020312-1013331022230210-2101101033123221-3303132022131321-0133302301220322"></a>

## Direct properties — port / 210001230332 / 3

- [info](data-sources--workload--reference--group-008.md#canonical-1100203133233112-3312221110210102-1210332011321212-3231310230300330-0213111211021221-1023312011200221-3311112330000323-0230133210322200): complete subsection reference.

<a id="canonical-3100123301110302-1113321323220232-2212232201220002-1021131130223311-2000332223302202-0121313332013313-1020033333212220-0120001001101010"></a>

## Next pages — port / 210001230332 / 4

- [service.advertise_options.advertise_in_cluster.port.info](data-sources--workload--reference--group-008.md#canonical-1100203133233112-3312221110210102-1210332011321212-3231310230300330-0213111211021221-1023312011200221-3311112330000323-0230133210322200)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-2131020211211101-0213303111310032-0232312121101303-0303013213123313-0112311302213302-2012200033230321-1231231303121203-0121103301113123)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1100203133233112-3312221110210102-1210332011321212-3231310230300330-0213111211021221-1023312011200221-3311112330000323-0230133210322200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322210220213232-1310032113132020-3333032120111031-0331121033123123-0020210113122222-1002322202330103-3020332210322123-1313120031203312"></a>

## service.advertise_options.advertise_in_cluster.port.info — info / 132312230221 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-2131020211211101-0213303111310032-0232312121101303-0303013213123313-0112311302213302-2012200033230321-1231231303121203-0121103301113123)
- [service.advertise_options.advertise_in_cluster.port](data-sources--workload--reference--group-008.md#canonical-2223131103313322-0013111330301232-2313200302013312-2300312202002333-1313130002133111-1321232321000201-1003101301002233-3322012010032123)
- service.advertise_options.advertise_in_cluster.port.info

<a id="canonical-1213011013003202-2202122332032123-1302122123001122-2001213012321302-0111212131330220-0330213302113201-3011222022211321-2031110200201003"></a>

Type: `"single"`. Computed.

Port Information. Port information.

Upstream description:

Port information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

<a id="canonical-1232210001001103-1023211221313311-1333120101312122-1022200103203210-1011321331023213-3213211221321333-3231023103201013-1103122213002300"></a>

## Direct properties — info / 132312230221 / 3

<a id="canonical-0232011113133211-0212120302222113-0321331001120100-3112021320221001-0120221200232032-2211332200031232-1212313321230003-1032020323331033"></a>

<a id="canonical-1320103210330023-1102023323110320-0023011303202333-0100133220222303-3300123103301033-0002001003013011-2111331003123212-3121032333231002"></a>

## port property — info / 132312230221 / 4

Type: `"number"`. Computed.

Port. Port the workload can be reached on.

Upstream description:

Port the workload can be reached on.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3120010110202323-3230233003212030-1323312022002313-0202303222211001-3300033132001302-2030032311002332-2133313133302233-0312223012331022"></a>

<a id="canonical-1333132120130212-1100021023220213-1311013202112302-2323322211130201-1210302300312032-1133332110130331-1013221221131200-2302301033032031"></a>

## protocol property — info / 132312230221 / 5

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

TCP &#8203;- PROTOCOL\_HTTP: HTTP

HTTP &#8203;- PROTOCOL\_HTTP2: HTTP2

HTTP2 &#8203;- PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI

TLS with SNI &#8203;- PROTOCOL\_UDP: UDP

UDP.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [same_as_port](data-sources--workload--reference--group-008.md#canonical-1010313300133320-1112331112010212-0020130030331122-0213012013031110-3121131111013030-2110003013030323-2222333030231322-3220323222331000): complete subsection reference.

<a id="canonical-1210221302200113-0100012213010333-3122210331132330-1100020011030122-3220200233320231-0030230331030230-1120013022222002-3213103120302203"></a>

<a id="canonical-3300113330120232-2102211301100101-2033121311021110-0211202130210011-2322101203113202-0031101031013110-2030000112022231-1131101333311132"></a>

## target_port property — info / 132312230221 / 6

Type: `"number"`. Computed.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Upstream description:

Exclusive with \[same\_as\_port\] Port the workload is listening on.

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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2132203023223021-2200011020322031-2131102111121230-2020323013011212-3323231310310131-0330031032111010-0121121110013131-1221120130221332"></a>

## Next pages — info / 132312230221 / 7

- [service.advertise_options.advertise_in_cluster.port.info.same_as_port](data-sources--workload--reference--group-008.md#canonical-1010313300133320-1112331112010212-0020130030331122-0213012013031110-3121131111013030-2110003013030323-2222333030231322-3220323222331000)
- [service.advertise_options.advertise_in_cluster.port](data-sources--workload--reference--group-008.md#canonical-2223131103313322-0013111330301232-2313200302013312-2300312202002333-1313130002133111-1321232321000201-1003101301002233-3322012010032123)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1010313300133320-1112331112010212-0020130030331122-0213012013031110-3121131111013030-2110003013030323-2222333030231322-3220323222331000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210201131302233-3120133133313100-3311213010022200-3023113322221010-3001213302303100-2301333102122323-3120132012312122-1311103233322132"></a>

## service.advertise_options.advertise_in_cluster.port.info.same_as_port — same_as_port / 023133311333 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-2131020211211101-0213303111310032-0232312121101303-0303013213123313-0112311302213302-2012200033230321-1231231303121203-0121103301113123)
- [service.advertise_options.advertise_in_cluster.port](data-sources--workload--reference--group-008.md#canonical-2223131103313322-0013111330301232-2313200302013312-2300312202002333-1313130002133111-1321232321000201-1003101301002233-3322012010032123)
- [service.advertise_options.advertise_in_cluster.port.info](data-sources--workload--reference--group-008.md#canonical-1100203133233112-3312221110210102-1210332011321212-3231310230300330-0213111211021221-1023312011200221-3311112330000323-0230133210322200)
- service.advertise_options.advertise_in_cluster.port.info.same_as_port

<a id="canonical-2230333333331320-2010003031031312-1022321030120022-3222231012320321-2313331320210322-0202223331300013-3000000121200212-1021213220231031"></a>

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

<a id="canonical-0120220312302302-2321030211112200-2222010121210202-0020111010111021-3202103123122210-3002023231233100-3203020212220223-3221200030321223"></a>

## Direct properties — same_as_port / 023133311333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332113332303122-0200221022020131-0321013212322010-2302200330111121-2221320132030203-1010232301220213-2302221132123112-0012021311031110"></a>

## Next pages — same_as_port / 023133311333 / 4

- [service.advertise_options.advertise_in_cluster.port.info](data-sources--workload--reference--group-008.md#canonical-1100203133233112-3312221110210102-1210332011321212-3231310230300330-0213111211021221-1023312011200221-3311112330000323-0230133210322200)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311013231212230-3033311200000202-1123113003220210-0313110113303331-3001203001313130-1133023113232200-1202231200230100-2230332322220213"></a>

## service.advertise_options.advertise_on_public — advertise_on_public / 230120333113 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- service.advertise_options.advertise_on_public

<a id="canonical-0002310100113021-1120113232030032-1121211313221000-1232312123211121-3203212031312121-1213231033321122-2322100022223311-3122010310100130"></a>

Type: `"single"`. Computed.

Advertise this workload via loadbalancer on Internet with default VIP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"multi_ports\",\"port\"]"
}
```

<a id="canonical-3133301012323120-3300230323213332-2231312313220013-2130200031220321-2300331230123120-1331213120323220-3201120302120321-1033130310113210"></a>

## Direct properties — advertise_on_public / 230120333113 / 3

- [multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012): complete subsection reference.

- [port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332): complete subsection reference.

<a id="canonical-3212200111110200-1323130211200102-1231033303201301-2003132100121322-2012020123221322-3013221201102213-3020200133132213-3203202110113200"></a>

## Next pages — advertise_on_public / 230120333113 / 4

- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313323011121003-1232303312010301-1302021031030100-0122233322222201-0312131233132333-0330123333222332-1233220013331220-0002311311321031"></a>

## service.advertise_options.advertise_on_public.multi_ports — multi_ports / 222123200022 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- service.advertise_options.advertise_on_public.multi_ports

<a id="canonical-1132223020300320-0123012113003033-2003002310032331-3323123332203233-2203300101102310-3230200010130300-3030111230113103-0333300222121203"></a>

Type: `"single"`. Computed.

Advertise Multiple Ports. Advertise multiple ports.

Upstream description:

Advertise multiple ports.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3013113011220001-3030132010122122-3330031313030223-0133113211021123-1331110202200103-1301103001030320-0321112130031111-2001133231122310"></a>

## Direct properties — multi_ports / 222123200022 / 3

- [ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120): complete subsection reference.

<a id="canonical-2221032012000300-1302020201031230-2112233002120323-2022303322303033-0320213013011133-3111112132311012-2232333300333123-1011130230001232"></a>

## Next pages — multi_ports / 222123200022 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300223233002013-3321203031012123-0023120103303122-1112320301300213-1303301313220212-0123101032201023-3230032323231120-3001211220022212"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports — ports / 033123332120 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- service.advertise_options.advertise_on_public.multi_ports.ports

<a id="canonical-3312211221101100-2132031223232300-0333330213313200-1013321221003131-3200323100110020-0020313332211313-0212121312001210-1122220001230033"></a>

Type: `"list"`. Computed.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-0220010323032100-0030021303300203-2331013330012133-2113312132103223-0123332131103330-0021203011022231-1000331131223132-2100232013330301"></a>

## Direct properties — ports / 033123332120 / 3

- [http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301): complete subsection reference.

- [port](data-sources--workload--reference--group-012.md#canonical-0330230310331300-1331023123323012-0102032330100131-3131300133300013-3111110113101200-0300012331333331-3113303133230233-2331220200130303): complete subsection reference.

- [tcp_loadbalancer](data-sources--workload--reference--group-012.md#canonical-3003113231211313-2011322221013301-0221321233231300-0201132111020033-2320132012211021-0112000312001132-3200330301031110-1310012311133030): complete subsection reference.

<a id="canonical-3330310232323033-1322100003311002-3000123320310033-1033102011221131-3100030302211221-1210013231223332-2231100203012101-1223231012001121"></a>

## Next pages — ports / 033123332120 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port](data-sources--workload--reference--group-012.md#canonical-0330230310331300-1331023123323012-0102032330100131-3131300133300013-3111110113101200-0300012331333331-3113303133230233-2331220200130303)
- [service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer](data-sources--workload--reference--group-012.md#canonical-3003113231211313-2011322221013301-0221321233231300-0201132111020033-2320132012211021-0112000312001132-3200330301031110-1310012311133030)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020330301331112-2131120301331003-2322230321103011-0013220001022021-0020012203102202-1210312230332230-1132121213003212-3120310100003322"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer — http_loadbalancer / 213201031111 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer

<a id="canonical-3332021333333303-1012030210203000-0031233323121233-3122122311021212-2313232212222333-3310322233003302-2331023221010222-3110021002122222"></a>

Type: `"single"`. Computed.

Configuration parameter for http loadbalancer.

Upstream description:

HTTP/HTTPS Load balancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]",
  "x-ves-oneof-field-route_choice": "[\"default_route\",\"specific_routes\"]"
}
```

<a id="canonical-2333210202023331-0123123200013223-1322230322133011-2201233111003130-1022322132100002-1111203230033002-3132223310110120-1023101203110120"></a>

## Direct properties — http_loadbalancer / 213201031111 / 3

- [default_route](data-sources--workload--reference--group-008.md#canonical-2022000103331303-1031232003202312-0023231311103320-2310100000131002-3001332203021122-3121010312331311-1211030031023032-1312023030221100): complete subsection reference.

<a id="canonical-2032023210013012-2022300001020222-3232202003021011-2313011130110313-3221322322022333-1010321310101111-0113032311321003-0233022010013032"></a>

<a id="canonical-2222233111303111-1303001120233023-1200020022113013-3123210300000233-3031001212021003-3321302200011323-0301332133320210-2233221133013102"></a>

## domains property — http_loadbalancer / 213201031111 / 4

Type: `["list", "string"]`. Computed.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\` is invalid
Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the..

Upstream description:

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*.bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*.example.com\`\` will match \`\`bar.example.com\`\`
and \`\`baz-bar.example.com\`\` but not \`\`.example.com\`\`. The longest wildcards match first.
Wildcards must match a whole DNS label. E.g. \`\`\*.example.com\`\` and \*.bar.example.com are
valid, however \`\`\*bar.example.com\`\` or \`\`\*-bar.example.com\`\` is invalid

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [http](data-sources--workload--reference--group-008.md#canonical-1011120000322313-0203202303130222-0321020201312302-2103210110013011-2313120202130120-2023321230030113-0220312010313221-2101203130122322): complete subsection reference.

- [https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120): complete subsection reference.

- [https_auto_cert](data-sources--workload--reference--group-010.md#canonical-0330123332132230-2323033213311003-0100012323030300-0001333232231100-1311000102100031-3330120010112333-1002233203102323-3111300312022210): complete subsection reference.

- [specific_routes](data-sources--workload--reference--group-011.md#canonical-0300030213323103-0001012303030311-1001122302013331-2012030011300330-3111230203222132-2320002210311130-0230303000212113-2103000301201323): complete subsection reference.

<a id="canonical-1101313110103000-1122031231102231-0030133330101003-1113320313130002-0022313230310222-1330302130323233-0103011330000001-1220311221333333"></a>

## Next pages — http_loadbalancer / 213201031111 / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-008.md#canonical-2022000103331303-1031232003202312-0023231311103320-2310100000131002-3001332203021122-3121010312331311-1211030031023032-1312023030221100)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http](data-sources--workload--reference--group-008.md#canonical-1011120000322313-0203202303130222-0321020201312302-2103210110013011-2313120202130120-2023321230030113-0220312010313221-2101203130122322)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-008.md#canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-010.md#canonical-0330123332132230-2323033213311003-0100012323030300-0001333232231100-1311000102100031-3330120010112333-1002233203102323-3111300312022210)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-011.md#canonical-0300030213323103-0001012303030311-1001122302013331-2012030011300330-3111230203222132-2320002210311130-0230303000212113-2103000301201323)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2022000103331303-1031232003202312-0023231311103320-2310100000131002-3001332203021122-3121010312331311-1211030031023032-1312023030221100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313201201332300-1221313110302311-0301332002220023-0333023022333302-1123233113202310-0011320110130321-1103020102010012-0100001313300202"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route — default_route / 211300130100 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route

<a id="canonical-1020303222103210-3233332210121230-0111321121312333-1221011100320333-3312211303321202-0331232112111010-0232211202003233-2333112223111213"></a>

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

<a id="canonical-1223200311230133-2110320231220021-2121322113132322-3122032003332033-2310120030022330-2132101123300131-1313013031223203-3222303011203311"></a>

## Direct properties — default_route / 211300130100 / 3

- [auto_host_rewrite](data-sources--workload--reference--group-008.md#canonical-2110021010010110-2121003033102233-0033211230210001-0313011330302012-3201103033211133-3333133301000000-3031100031233302-0133303102030201): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-008.md#canonical-0111112200332001-3011102000302133-0130022000030203-3102321310101332-0210120202210311-3322332030012310-3013303101103203-3031120103213013): complete subsection reference.

<a id="canonical-0231222122303212-0322033211303233-1331313131122200-3011021321101112-2111323123223312-0012212120022110-1302110130123123-1222002003113113"></a>

<a id="canonical-3323120211103212-1003121330001210-3323313313212003-2130100100000313-0320212010301322-3013013102200022-3003330201212031-0203010332223203"></a>

## host_rewrite property — default_route / 211300130100 / 4

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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-0310001200001233-3002332033201122-2031203101301002-2211201233222230-3110023120133002-3020120020221113-0033201123032210-0223213200001103"></a>

## Next pages — default_route / 211300130100 / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite](data-sources--workload--reference--group-008.md#canonical-2110021010010110-2121003033102233-0033211230210001-0313011330302012-3201103033211133-3333133301000000-3031100031233302-0133303102030201)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite](data-sources--workload--reference--group-008.md#canonical-0111112200332001-3011102000302133-0130022000030203-3102321310101332-0210120202210311-3322332030012310-3013303101103203-3031120103213013)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2110021010010110-2121003033102233-0033211230210001-0313011330302012-3201103033211133-3333133301000000-3031100031233302-0133303102030201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202233021113200-2310122032221210-3101010213122311-2120330030131220-3331002001303030-0202103222021020-1121102313321001-0102022233321202"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite — auto_host_rewrite / 332310200323 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-008.md#canonical-2022000103331303-1031232003202312-0023231311103320-2310100000131002-3001332203021122-3121010312331311-1211030031023032-1312023030221100)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-0210012112013333-1110333332031012-0333120110101003-0001123301223211-3200221033333323-3212303003302031-2111100031002320-3111013130310331"></a>

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

<a id="canonical-3303102303212023-2330200321030101-3300233111322002-0021210303313033-2020112123321110-2030202113302202-0213203130312313-2103311132133103"></a>

## Direct properties — auto_host_rewrite / 332310200323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1001233333023301-3101210003111012-1203210032210213-3220223323331203-2130330201322320-1020111003303320-0212213100131101-1110233103111200"></a>

## Next pages — auto_host_rewrite / 332310200323 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-008.md#canonical-2022000103331303-1031232003202312-0023231311103320-2310100000131002-3001332203021122-3121010312331311-1211030031023032-1312023030221100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0111112200332001-3011102000302133-0130022000030203-3102321310101332-0210120202210311-3322332030012310-3013303101103203-3031120103213013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020322130032302-0122332030300322-0210300201210202-1010130113233311-3031023121310103-2222032213000101-0323301032230111-3012323033022230"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite — disable_host_rewrite / 030112031032 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-008.md#canonical-2022000103331303-1031232003202312-0023231311103320-2310100000131002-3001332203021122-3121010312331311-1211030031023032-1312023030221100)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-0202122312111132-1123313032132113-2120322120113031-3330033220312203-3113321121333012-2031331010120022-2322031033101011-3123203121303010"></a>

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

<a id="canonical-0200210331310113-0220331133032200-3103302013131100-0203030031322110-3110233323012211-0311132130101002-2223003222122221-0300032110100211"></a>

## Direct properties — disable_host_rewrite / 030112031032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310010321331031-0030000033301300-2330100323320230-1200100210132231-1000113330233222-0021231320131033-2112000020212302-1021111002323321"></a>

## Next pages — disable_host_rewrite / 030112031032 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-008.md#canonical-2022000103331303-1031232003202312-0023231311103320-2310100000131002-3001332203021122-3121010312331311-1211030031023032-1312023030221100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1011120000322313-0203202303130222-0321020201312302-2103210110013011-2313120202130120-2023321230030113-0220312010313221-2101203130122322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033103000103023-3011203120312103-3210321310311203-3330202001320300-3233011313312210-2323032122131211-1213331322101201-2323020213202022"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http — http / 332021000113 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http

<a id="canonical-3010100310331323-2011202303132332-0330311021303220-2222300331210232-2300203012300310-3010300111010111-2210022112032332-2220122320320013"></a>

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

<a id="canonical-0231133002000102-1223312211012000-3321130200100030-0001130222022030-3210110021010220-2311221101123210-0132311330103010-1013313311010331"></a>

## Direct properties — http / 332021000113 / 3

<a id="canonical-2302030233301032-0330133032110313-1232012310321120-1303201133330130-1320331001232000-1101211312320232-3130020322013230-2021301212330030"></a>

<a id="canonical-0313131033332133-2120112330113132-2332332311103322-1302200320233031-3310331312233023-1013021112201203-2102331230030113-0231333003233112"></a>

## dns_volterra_managed property — http / 332021000113 / 4

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

<a id="canonical-1012213022330331-2103312121023210-3131021032011001-0211221303120112-3133310320122312-2021300313011330-0321013212010112-1221312010303231"></a>

<a id="canonical-2122322132223331-1320100031013212-2210130233300213-0013201201003301-0101022232323312-2220333303231312-0300100102212332-0110001333012122"></a>

## port property — http / 332021000113 / 5

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

<a id="canonical-0031101203000320-3313010203010101-0322221203232020-1010111122020022-3220032001121002-1023103212310122-1321233203231221-0311200200131331"></a>

<a id="canonical-2202000331322001-3032221203202331-2030023302300302-1200232232211020-2202010232023321-2321331221021103-1103211220230023-1010032033200002"></a>

## port_ranges property — http / 332021000113 / 6

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

<a id="canonical-2320122101211220-1313223301203311-1330033010203033-3303000022022331-3003330323330200-1020010012130121-0230113323322100-1222332121021221"></a>

## Next pages — http / 332021000113 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233333332002003-0320022303213120-2100101220012001-0220132300323002-1323001302232330-2131001122221031-2010032021030112-1313021212013003"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https — https / 330302311213 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https

<a id="canonical-1222221113031211-2120110211330323-2030033210232211-2210123122130130-3111121210300013-0010110100122301-0101233133320132-2133211102100330"></a>

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

<a id="canonical-1100221220101231-2210303020301021-1301321103311002-1110110031102203-2111110012210033-0100000223113312-3020032020113102-0200210022000221"></a>

## Direct properties — https / 330302311213 / 3

<a id="canonical-0221012130000100-0032331222332332-2200122332102332-0010123120111132-0121230010313211-2002123122301212-1133201233222122-2130223113001202"></a>

<a id="canonical-1320103220023001-0031202200321210-3101121311013003-0303200330130210-3130301130312213-0010220000313112-2203022132032120-1122030133303301"></a>

## add_hsts property — https / 330302311213 / 4

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

<a id="canonical-2203020033001120-3101122002003031-2320103203120013-3011011010233213-2133130130122033-3121220220333201-0200132131201032-1222231132212211"></a>

<a id="canonical-0123231311010203-1333230112030000-0133323103302002-1102302120003230-1011212232032021-1211032310122133-1010323003121211-0000212322320331"></a>

## append_server_name property — https / 330302311213 / 5

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

- [coalescing_options](data-sources--workload--reference--group-009.md#canonical-1021033321200220-2122123222010113-3330010201103233-2103020220030330-2021200120310310-1302033032023022-1321103303100000-3212020230312013): complete subsection reference.

<a id="canonical-0023211102022100-2211010122300131-2021011312322221-1233230313103310-1132010212131020-2121130330223101-0232321302333132-1331310213333131"></a>
