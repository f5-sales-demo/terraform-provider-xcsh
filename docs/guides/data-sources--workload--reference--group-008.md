---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3113321211103300-3121123312121102-2103230121222320-0023030100232313-0030331001233112-3103020012203031-1211333322302002-3323010310210312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route

<a id="canonical-0311012010021310-1120121100032011-2332303021212100-1321310112321013-3231210213202033-2332232312030133-3033032033123231-1132111131333120"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3002331120302120-0112031231330002-0203030303322100-1201130100231022-1120210331300302-3002201230230113-0323000012110200-0013111311311123"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route`

- [headers](data-sources--workload--reference--group-008.md#canonical-3213031321232203-0200213113100310-1332113110320200-0322013221032123-1312302112121331-1122331203010123-2333231121233310-3100113321320200): complete subsection reference.

<a id="canonical-2202003133120203-2302231311312030-2322102101031022-0103301213110223-3310002330210211-1200213311300231-3030230123331002-1322020020110031"></a>

<a id="canonical-1221103231212003-3322120231120202-2302312321132111-2223310022320203-0023121211033133-2313023030303310-1330232313033031-2223121333213111"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.http_method` property

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

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

<a id="canonical-3213031321232203-0200213113100310-1332113110320200-0322013221032123-1312302112121331-1122331203010123-2333231121233310-3100113321320200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-008.md#canonical-3113321211103300-3121123312121102-2103230121222320-0023030100232313-0030331001233112-3103020012203031-1211333322302002-3323010310210312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers

<a id="canonical-2233312221220102-0212223222100300-3202032102211132-2331320222312022-2311332012133102-1323013233101122-1302230322213021-3223101322211130"></a>

Type: `"list"`. Computed.

Headers. List of (key, value) headers.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1101001322231013-2331032101130103-1031210012321320-1122202112322011-0031111212002113-1202220012002030-0233122301313011-0003012013301322"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers`

<a id="canonical-1121303313231201-3202210212211310-1001101200310232-2130001122002202-0100203300330033-1030133003230030-2322310312112203-1102300223230001"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers.exact` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1202203112121232-0310301301313100-0321220202123233-3200002202223000-1211200220303221-2133320202013133-0221101231021003-2123213332122331"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers.invert_match` property

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

<a id="canonical-3000330133221330-2202122312121031-3133210222312233-1010021302121101-0223310222332002-0133032233201331-0220130322021110-0321213113231131"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers.name` property

Type: `"string"`. Computed.

Name. Name of the header.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1301130111030110-2103201112020330-0023120311232321-3021120133220020-3311222121020110-3202313312212110-0321233020020011-1121100012222103"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers.presence` property

Type: `"bool"`. Computed.

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

<a id="canonical-2003201321213322-2320012223030011-0213010301123211-2201000302023123-0200000201210011-2100010023103333-1101132011202211-2002311230210310"></a>

<a id="canonical-2200202330032001-2201313330332121-0320032021113312-3123230123222320-2023012133220010-0312031201333211-0112200303120113-0033021210033122"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers.regex` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1220223331102112-3311030110330320-3312022010122030-0323002023012210-2211120031010111-3202221103032312-3113313102103213-0311102310032123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-008.md#canonical-3113321211103300-3121123312121102-2103230121222320-0023030100232313-0030331001233112-3103020012203031-1211333322302002-3323010310210312)
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

<a id="canonical-2100230113130233-1132300111201122-1000220312230213-0233032013222010-0000000231221130-3011130013133231-0232311122023031-3301231033100130"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port`

- [no_port_match](data-sources--workload--reference--group-008.md#canonical-2001211111311033-1222220210301102-0100200031231121-1222030230303223-1013230103331323-3303013212300223-3011212010131132-1220332003033010): complete subsection reference.

<a id="canonical-3203203021230303-2232112201331010-3202100102030130-3323131310032212-1030033022002000-0113223002303022-3133112132010013-3012202030122301"></a>

<a id="canonical-2313312120321032-2311323331123010-0001111232102203-3023321301303031-3301022011122012-1013332033130020-3003130203033031-1021112310300023"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.port` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0300213232031300-2201112312301101-1122220003100222-0101311201210213-2023232202030300-0210323231100111-1000031230112210-2000020030220031"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.port_ranges` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2001211111311033-1222220210301102-0100200031231121-1222030230303223-1013230103331323-3303013212300223-3011212010131132-1220332003033010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-008.md#canonical-3113321211103300-3121123312121102-2103230121222320-0023030100232313-0030331001233112-3103020012203031-1211333322302002-3323010310210312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](data-sources--workload--reference--group-008.md#canonical-1220223331102112-3311030110330320-3312022010122030-0323002023012210-2211120031010111-3202221103032312-3113313102103213-0311102310032123)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match

<a id="canonical-0321203022211013-3012130100111211-2323333322212111-2000231110023030-0201212033303222-0002322302013301-0032220022320033-0321101312231012"></a>

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

<a id="canonical-3112303221220112-3220200323222331-3031032321232031-3231213031300132-0222021122010321-3010200221120330-1303033103030000-0212303111301332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-008.md#canonical-3113321211103300-3121123312121102-2103230121222320-0023030100232313-0030331001233112-3103020012203031-1211333322302002-3323010310210312)
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

<a id="canonical-1223212303213202-3100223103230323-0010122313012200-2332230132032330-0321020203020303-0110211201210210-3212031221230200-3003230122231331"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path`

<a id="canonical-1313303201332102-0213230133121101-0132300023031223-3231120230013212-3230122131303203-3312320002103133-3111111211320312-0312223130000331"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path.path` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0231131202120002-0133130321231203-3312121020101012-0100121233000123-1102000132313121-2110231000110103-2002112330133000-3220223313212222"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path.prefix` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2102232121013100-2330211200103201-2021330013220322-1322330311220221-0303222102221203-2323332100122120-1201321100310333-0220200110203021"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1233012303212313-3101123331213120-2210130313232300-3010333233222310-0331231001203323-1000212221033112-1323002011222232-1030300300323101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-008.md#canonical-3113321211103300-3121123312121102-2103230121222320-0023030100232313-0030331001233112-3103020012203031-1211333322302002-3323010310210312)
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

<a id="canonical-3320102101222212-2302022102121133-2131023120230313-1132000131312230-1321332112120203-1330232302103033-3102031221301003-1003323120023113"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response`

<a id="canonical-3021012312102031-1132012203123121-1321221223101223-1111200022212220-2301233322031303-3331200131232331-2202322301321013-3003003212032110"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response.response_body_encoded` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2012332202300301-0130011120330302-1301101333332231-2010131111121113-2203113131112101-1202301111111330-0001133301011210-2232330110203120"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response.response_code` property

Type: `"number"`. Computed.

Response Code. Response code to send.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-3003323202321120-0230323212223020-0101102231001321-1032012232002121-3311233032110000-3200221013203101-0300010231121211-2023211302121113"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3112001111222131-0122020121312032-2212203120100010-1331010111101233-2301221331310233-0123301000232323-1333031010030202-3230123003310301"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route`

- [headers](data-sources--workload--reference--group-008.md#canonical-3303100213023330-1230312202133020-1001031231121212-1112011220021122-3010122020112100-3123312320220200-1302201130011331-3221230230323130): complete subsection reference.

<a id="canonical-0020121311120310-0130300113133231-3020301312202110-0101030231000310-1233011110300103-3302203201310233-2333023321311131-0133310302133023"></a>

<a id="canonical-1231321232233222-1221323323123010-3033200323330002-2213331301330132-3312323301120302-3213322033300023-2113011312021002-1200101323100231"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.http_method` property

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

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

<a id="canonical-3303100213023330-1230312202133020-1001031231121212-1112011220021122-3010122020112100-3123312320220200-1302201130011331-3221230230323130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-2300123323331132-0220210011123233-0231221302303011-0133202101310301-2323003130311031-0002233200212102-1220012122321030-0310000202201330"></a>

Type: `"list"`. Computed.

Headers. List of (key, value) headers.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2303120100302123-0302131012200120-0023000211332332-2121001002301101-0232001333113112-2331333133213012-0302111212030131-1302221233102130"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers`

<a id="canonical-1333203021032023-1221110132202132-0232032023332112-3133120003213212-1310203311313222-1203001230122133-3013211201200313-1311322332311310"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.exact` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0222310322103031-3131333312232023-3202210201231231-0120113122113031-2011013121330030-3213003301112223-1203022022013331-2200311312002003"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.invert_match` property

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

<a id="canonical-1133033000211100-2131031220102100-2222001222313021-1102233211313311-0101003110331313-0331021223030220-2021001010103101-0212301033110222"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.name` property

Type: `"string"`. Computed.

Name. Name of the header.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3303212210332320-3313021202102130-0100122302302210-0212220123021010-2201013013020321-3131212100032020-1331211200332120-0302030301000212"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.presence` property

Type: `"bool"`. Computed.

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

<a id="canonical-1012200301110001-2012310320233200-1021232013221230-3110320201033311-0120013231103311-0133030313203231-1000022012231111-0032221320212120"></a>

<a id="canonical-3322030221132231-0301111103212301-3201021010211233-3230100012310333-1303033312113131-1320232302322333-0323132211233130-3131133112020012"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.regex` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0203230113012003-3121132002233002-2001233201131311-3331123331010211-0210013332003122-2210120321310031-2301031011310030-3332203103302302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
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

<a id="canonical-2103113110201030-0300231131200003-2002021132123301-3120012233013103-1221003022233013-2330233313231301-1322302301101031-0120001123231321"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port`

- [no_port_match](data-sources--workload--reference--group-008.md#canonical-2210233331230201-3202222231103000-3111203200330120-1000010101133223-3001132311233120-3300330013100001-3011121302311210-2013010103233120): complete subsection reference.

<a id="canonical-3200221230112201-1011123030023213-1331220132030103-0321202103130012-2122223230123023-1233213302231222-3302122203313132-2230003231112220"></a>

<a id="canonical-1302102020223322-3013233232220130-1330123110120300-1011303303203201-1031103001231232-2022122121310313-1302032333212123-0323113321132232"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.port` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1123111121102310-1332220013231121-2221200200112032-0331012020323011-1101303103103020-0021031112113031-3222223321010230-0022233303203000"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.port_ranges` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2210233331230201-3202222231103000-3111203200330120-1000010101133223-3001132311233120-3300330013100001-3011121302311210-2013010103233120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-008.md#canonical-0203230113012003-3121132002233002-2001233201131311-3331123331010211-0210013332003122-2210120321310031-2301031011310030-3332203103302302)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-0231023231131103-2010112002030113-2121331003202330-1300332112232012-0213033103033132-1232012020001121-3022103230201302-3232130311312120"></a>

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

<a id="canonical-1310113111301010-2210003323310121-2301120032013120-0131032330221011-0133210022001300-1000331030120230-0001032300130033-3332232201201223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
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

<a id="canonical-2210313211221303-1123113231033102-3223001133121100-1233300022333011-3120010122221320-1333300122210100-0201332300101022-0131012023110231"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path`

<a id="canonical-3330122221232110-1230103100233031-0320120200013331-2130112203211131-1013210030323223-2233213213133222-3220013303130221-3221322211310331"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path.path` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3013100130011121-2223012003303230-0211123121203100-1030003120321213-1211302120233230-1002200320333333-1031223021102332-3213303213003000"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path.prefix` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0032032022123013-1030001220220012-1121220203120331-3110200312133222-3310032300131033-2100231033202332-1101222312132023-1101333123232010"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1313133332010100-0222322321110231-0031303330213222-1332232222223031-1101322302012120-3321000333301121-1321131021333010-2320221231230120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
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

<a id="canonical-2331330221110313-2311010300103220-2013031113212332-2223120133121111-1300302331312022-3200122031110130-0032220211003130-3100311132213111"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect`

<a id="canonical-2013312312232011-1200132203332001-1222210312002233-1300132221320313-0222200333222321-2333031312330213-0033202013101120-2101123312120110"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.host_redirect` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2102321210113031-2111001021203002-2120200113003120-0301231012232320-0331122003310310-3102112132032211-3321100200100211-1110331112033031"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.path_redirect` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3231132331101303-1322133203033120-3101120200032122-1011212110111333-2121201000102322-3330210323313003-1211000220123012-2210031101020300"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.prefix_rewrite` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1110212120311011-2212303213312112-3003131112313321-0310022311013032-2003212321210222-3232032100232233-2031103122331212-3321020331011031"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.proto_redirect` property

Type: `"string"`. Computed.

\[Enum: incoming-proto|http|https\] Swap protocol part of incoming URL in redirect URL The protocol
can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of
protocol is not done. Possible values are \`incoming-proto\`, \`http\`, \`https\`.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0111311122120012-1211033010222220-0023010123131123-1000300310122321-0323313200021012-2210203031012000-3202110120220120-2301013003333211"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.replace_params` property

Type: `"string"`. Computed.

Exclusive with \[remove\_all\_params retain\_all\_params\].

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3301033120133103-2222302012012120-1210103112230000-2003012213200222-3032202330022022-3130122021200033-1002111302301321-2011020112310001"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.response_code` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1231223201210110-2300312012213203-2111211220311001-3111231133301332-3123123302131303-2130311310101323-1222220330221211-0102332232032301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-008.md#canonical-1313133332010100-0222322321110231-0031303330213222-1332232222223031-1101322302012120-3321000333301121-1321131021333010-2320221231230120)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-1032100023030233-0013012003030303-3302101230011323-3031011330312313-3323002223232222-1331223003131111-1323113333013232-1320213032200310"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for remove all params.

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

<a id="canonical-2013033210103222-1320000312030300-1030201120230023-1201110333323030-3323310312300332-1211202023111330-1212320303322030-2132022101001030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-008.md#canonical-1201002300120012-1200100030133330-2133130220322113-0012100022212021-1102200022131300-2012230133300001-0010312003303310-3032220313100132)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-008.md#canonical-1313133332010100-0222322321110231-0031303330213222-1332232222223031-1101322302012120-3321000333301121-1321131021333010-2320221231230120)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-3332303011323101-2100302310301333-3132123110221332-2013101002301231-0311333001011203-2132111312100222-0111002010211302-1013031321110101"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for retain all params.

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

<a id="canonical-0310112122113210-2110103333320012-3100102321331121-2303102223302303-0033320103233333-1221123312232313-2132013023311033-1120032010201221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-0330201022113320-0323132220101330-2032312220323130-2301020233110231-2312102032133023-1320003233211201-1132303330101312-3233011221021322"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2001102012320333-3123030223123112-0113103232231211-1331331212220103-0021001333032212-3003330103022020-3203112211131203-3123022111230001"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route`

- [auto_host_rewrite](data-sources--workload--reference--group-008.md#canonical-3111112032111131-0201023202011221-3021010031211231-3031012110222322-2223032001201031-3331110213133032-1202332132112212-3313122313032122): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-008.md#canonical-0010313130021121-3113302100330202-3300213012332313-1133220113330323-1300322030033322-2203221132112213-2200302123222212-1020323223013031): complete subsection reference.

<a id="canonical-3003103113300323-2302303313222102-1210222101032213-0001033111212320-3200331101011223-3223010321002323-1132322022001020-1200331001130312"></a>

<a id="canonical-1132012131213003-0113001132313113-3213111022301222-2312112203302301-3111133113332021-0322031121310000-0233102320023303-1022330003000200"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.host_rewrite` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3001032322020310-2012211001000211-2331002323231030-0121020012010032-2131121010201203-2011211223231030-0322323212212223-2200212213323213"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.http_method` property

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

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

<a id="canonical-3111112032111131-0201023202011221-3021010031211231-3031012110222322-2223032001201031-3331110213133032-1202332132112212-3313122313032122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-008.md#canonical-0310112122113210-2110103333320012-3100102321331121-2303102223302303-0033320103233333-1221123312232313-2132013023311033-1120032010201221)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-2100120130130201-2023123221002110-3330001131021233-1230021232003203-3330330310312201-0103121010203202-0210023200110212-3320113301101223"></a>

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

<a id="canonical-0010313130021121-3113302100330202-3300213012332313-1133220113330323-1300322030033322-2203221132112213-2200302123222212-1020323223013031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-007.md#canonical-2003100000123021-0313322222311230-2033012233001130-0232233301233103-2122221223220113-0132133112200321-1132222233302121-3302230203200001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-007.md#canonical-3103111112211230-1120132001111211-0222302132231333-0310223330121330-1121130002100312-3002002002231100-0101322031122102-2232313033001212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-008.md#canonical-0310112122113210-2110103333320012-3100102321331121-2303102223302303-0033320103233333-1221123312232313-2132013023311033-1120032010201221)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-0003113333102010-3030031202112031-2323001011221223-2020213212110211-1012103023212102-1222211021000133-1020302130130120-0020133113133203"></a>

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

<a id="canonical-1103333221333300-3332320302022022-2130011200321223-2022010220200112-1313121333320301-3230012303211132-0021110133131330-2323223100203030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-006.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
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

<a id="canonical-2121313320033332-3221131101010300-2133030333320223-0123332231332022-3123313332212322-3222321120132231-2103102113302330-1012311021301233"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path`

<a id="canonical-3233300331021111-1112322113002212-1112113103033011-0203230032211012-1101021003010312-1100030130331013-3310221123132201-1331000232301310"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path.path` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1303223312130213-1022210202233230-2012122110312223-2023233210201333-2100103020012100-0220331201131233-1230110111231110-3230133310321323"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path.prefix` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3302222302133011-0303332013102112-1312210332301320-0030210122332201-1013301230103233-1010303310033331-0202230112031003-2202101100012202"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1330120103311032-3323002022332313-1230010032111002-0301100332330310-1013300322113000-3020110122120300-1230010001312321-1311000201122222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- service.advertise_options.advertise_custom.ports.port

<a id="canonical-0212102122101321-0101022310120302-3320221011210112-1212200300103123-2311003102102002-3200021323130221-0013003322311123-3200133130232201"></a>

Type: `"single"`. Computed.

Port. Port of the workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1202203313223131-3030000101123110-2331121211203333-1231023100132122-2121300333303330-2101232100200131-3120230310200233-2332031102021302"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.port`

- [info](data-sources--workload--reference--group-008.md#canonical-0110122133233020-0013313330132130-1321111101003321-0211310020013220-0110220200320011-1003211320330303-1033031311131333-2211112312023000): complete subsection reference.

<a id="canonical-1322201112212023-0332000202310201-0032232311303322-2010300122201011-1223120313323221-0010002031112123-1133230010022012-0331012101003220"></a>

<a id="canonical-1011303032331010-0023101010223101-1311211030310131-2311002020200313-2100212110131230-2031133233202300-2020221213032210-3102131201122321"></a>

#### `service.advertise_options.advertise_custom.ports.port.name` property

Type: `"string"`. Computed.

Name. Name of the Port.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0110122133233020-0013313330132130-1321111101003321-0211310020013220-0110220200320011-1003211320330303-1033031311131333-2211112312023000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.port.info` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.port](data-sources--workload--reference--group-008.md#canonical-1330120103311032-3323002022332313-1230010032111002-0301100332330310-1013300322113000-3020110122120300-1230010001312321-1311000201122222)
- service.advertise_options.advertise_custom.ports.port.info

<a id="canonical-2201301010121010-1011300020231210-2031003001213330-3000121031133210-0103230223011030-1010033320322210-0121122003203111-0331121200302330"></a>

Type: `"single"`. Computed.

Port Information. Port information.

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

<a id="canonical-3033201001301332-1111003111022323-1302022123003322-2311303020333213-0202321222230312-1000112130022313-1121001320331230-1320331113123203"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.port.info`

<a id="canonical-3302201320013322-2030331201100300-0211300301110321-1131312030100311-2121120320012223-2310221001213221-2012113223312112-3320331231012031"></a>

#### `service.advertise_options.advertise_custom.ports.port.info.port` property

Type: `"number"`. Computed.

Port. Port the workload can be reached on.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0030231132103322-1033211023032301-1232311300321123-1201110220012012-1131100203330323-3230303331333010-1313010133100102-1213300023330323"></a>

#### `service.advertise_options.advertise_custom.ports.port.info.protocol` property

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

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

<a id="canonical-1312132120220322-1221110331000010-3211223100322333-0201120210331000-1312120023302033-3231112113131003-1021130213222132-3013030122231111"></a>

#### `service.advertise_options.advertise_custom.ports.port.info.target_port` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1210011231223112-0112103212001003-2023321021321303-2232320101010331-0013332100311120-0332013310112322-1120132310101300-2120031112211310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.port.info.same_as_port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.port](data-sources--workload--reference--group-008.md#canonical-1330120103311032-3323002022332313-1230010032111002-0301100332330310-1013300322113000-3020110122120300-1230010001312321-1311000201122222)
- [service.advertise_options.advertise_custom.ports.port.info](data-sources--workload--reference--group-008.md#canonical-0110122133233020-0013313330132130-1321111101003321-0211310020013220-0110220200320011-1003211320330303-1033031311131333-2211112312023000)
- service.advertise_options.advertise_custom.ports.port.info.same_as_port

<a id="canonical-2210221320120312-3321022021333223-0323113332103322-1233000122302001-2121300220310021-0002203123333001-1323200333210210-1131201203023212"></a>

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

<a id="canonical-1213013231333121-3300100100220230-1131111132003220-3022122220020203-0200132131130103-3220333332311300-2220212322331102-1331120231121131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.tcp_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- service.advertise_options.advertise_custom.ports.tcp_loadbalancer

<a id="canonical-2211133012131001-2203212223030313-0131102131302203-3010322330230031-1102211232103122-1131110001212311-2123300311000321-1321223213100021"></a>

Type: `"single"`. Computed.

Configuration parameter for tcp loadbalancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0321011133200303-1123213210023002-1322302000212310-3200021001203102-2131200002010300-1032313322113133-0321003203301313-3132002211310213"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.tcp_loadbalancer`

<a id="canonical-0232330130313333-0302113210312332-0102003133012312-0102310200210112-2303212021323222-0232033020001130-2123132313210323-0033220000213201"></a>

#### `service.advertise_options.advertise_custom.ports.tcp_loadbalancer.domains` property

Type: `["list", "string"]`. Computed.

List of additional domains (host/authority header) that will be matched to this loadbalancer.
Domains are also used for SNI matching if the is true Domains also indicate the list of names for
which DNS resolution will be done by VER.

Additional upstream details:

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2001120033222022-0223103231210100-3103330022203220-3232333201223333-3131113113213210-0311001201320203-1330232322022300-2003313202002121"></a>

#### `service.advertise_options.advertise_custom.ports.tcp_loadbalancer.with_sni` property

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

<a id="canonical-2131020211211101-0213303111310032-0232312121101303-0303013213123313-0112311302213302-2012200033230321-1231231303121203-0121103301113123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster` properties

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

<a id="canonical-3123212303301223-0322131120120210-0223303032221200-0200211231123330-0333100311320103-3221233302302132-1201102230112302-2201030031011310"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster`

- [multi_ports](data-sources--workload--reference--group-008.md#canonical-0323130202203203-0212323222122012-0133122332003031-0311312333002330-0110212200031012-3030321303133231-2322211032002200-0003101011302311): complete subsection reference.

- [port](data-sources--workload--reference--group-008.md#canonical-2223131103313322-0013111330301232-2313200302013312-2300312202002333-1313130002133111-1321232321000201-1003101301002233-3322012010032123): complete subsection reference.

<a id="canonical-0323130202203203-0212323222122012-0133122332003031-0311312333002330-0110212200031012-3030321303133231-2322211032002200-0003101011302311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.multi_ports` properties

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

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1100102323320300-0310123212133030-2300003210303130-3112013233113000-1022030102213201-1202202120131032-0130332303013333-3331021211230220"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster.multi_ports`

- [ports](data-sources--workload--reference--group-008.md#canonical-1200123030213102-3312221012212330-2122212021203232-3111213102133231-3300123331001200-1300033313011233-0113120231301331-3210113023313323): complete subsection reference.

<a id="canonical-1200123030213102-3312221012212330-2122212021203232-3111213102133231-3300123331001200-1300033313011233-0113120231301331-3210113023313323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.multi_ports.ports` properties

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2011021313032032-1132200231013222-0002321232313323-1102221031332321-0221012333111230-1323001001230033-2002133313022221-2300320030002310"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster.multi_ports.ports`

- [info](data-sources--workload--reference--group-008.md#canonical-1311013020110322-3332312222230230-2111331120112121-1302023321211110-3121210102132002-3331300220020203-0032132323002103-2013103230232311): complete subsection reference.

<a id="canonical-2111222223331201-2120012010100122-3222203222123021-3333131113021033-3301301130201301-3200313132121311-1130002003221112-2131201333333221"></a>

<a id="canonical-2330022203330010-1120220031122010-0000013102231323-0221012330003123-3211222033212332-1232201312030301-1110020022010203-0332320102011121"></a>

#### `service.advertise_options.advertise_in_cluster.multi_ports.ports.name` property

Type: `"string"`. Computed.

Name. Name of the Port.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1311013020110322-3332312222230230-2111331120112121-1302023321211110-3121210102132002-3331300220020203-0032132323002103-2013103230232311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.multi_ports.ports.info` properties

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

<a id="canonical-0200213332230031-1120031213201103-3223031110023122-3012103121030212-0032332013312130-3033121131021130-1133113231020131-0001303112302323"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster.multi_ports.ports.info`

<a id="canonical-3012332131231010-2310032113130330-2020030232311031-0013031010110333-3303013323202133-2322012013312102-2013001331021201-0122220020320113"></a>

#### `service.advertise_options.advertise_in_cluster.multi_ports.ports.info.port` property

Type: `"number"`. Computed.

Port. Port the workload can be reached on.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2031332222013012-0002002120022112-3021333111010021-3113111321221232-3222202122323331-1201010331231031-2313112211022123-0311322013020003"></a>

#### `service.advertise_options.advertise_in_cluster.multi_ports.ports.info.protocol` property

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

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

<a id="canonical-2303312011123020-2013130123232231-2130023231020310-3331030120012233-0203000333332002-1303223111333133-1211112222003112-2300023003330313"></a>

#### `service.advertise_options.advertise_in_cluster.multi_ports.ports.info.target_port` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3312230321301022-3201311033220031-0223210311200033-3110222212213013-3312031030333223-0311310302230130-1220121101111132-1003020123013313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port` properties

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

<a id="canonical-2223131103313322-0013111330301232-2313200302013312-2300312202002333-1313130002133111-1321232321000201-1003101301002233-3322012010032123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.port` properties

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

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1010033113232002-3033131033331003-2232320301202211-3131312103301332-1133013222312313-3331300003301310-1312123132110120-2131002032303210"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster.port`

- [info](data-sources--workload--reference--group-008.md#canonical-1100203133233112-3312221110210102-1210332011321212-3231310230300330-0213111211021221-1023312011200221-3311112330000323-0230133210322200): complete subsection reference.

<a id="canonical-1100203133233112-3312221110210102-1210332011321212-3231310230300330-0213111211021221-1023312011200221-3311112330000323-0230133210322200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.port.info` properties

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

<a id="canonical-0322210220213232-1310032113132020-3333032120111031-0331121033123123-0020210113122222-1002322202330103-3020332210322123-1313120031203312"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster.port.info`

<a id="canonical-0232011113133211-0212120302222113-0321331001120100-3112021320221001-0120221200232032-2211332200031232-1212313321230003-1032020323331033"></a>

#### `service.advertise_options.advertise_in_cluster.port.info.port` property

Type: `"number"`. Computed.

Port. Port the workload can be reached on.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1232210001001103-1023211221313311-1333120101312122-1022200103203210-1011321331023213-3213211221321333-3231023103201013-1103122213002300"></a>

#### `service.advertise_options.advertise_in_cluster.port.info.protocol` property

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

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

<a id="canonical-1320103210330023-1102023323110320-0023011303202333-0100133220222303-3300123103301033-0002001003013011-2111331003123212-3121032333231002"></a>

#### `service.advertise_options.advertise_in_cluster.port.info.target_port` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1010313300133320-1112331112010212-0020130030331122-0213012013031110-3121131111013030-2110003013030323-2222333030231322-3220323222331000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.port.info.same_as_port` properties

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

<a id="canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- service.advertise_options.advertise_on_public

<a id="canonical-0002310100113021-1120113232030032-1121211313221000-1232312123211121-3203212031312121-1213231033321122-2322100022223311-3122010310100130"></a>

Type: `"single"`. Computed.

Advertise this workload via loadbalancer on internet with default VIP.

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

<a id="canonical-1311013231212230-3033311200000202-1123113003220210-0313110113303331-3001203001313130-1133023113232200-1202231200230100-2230332322220213"></a>

### Direct properties for `service.advertise_options.advertise_on_public`

- [multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012): complete subsection reference.

- [port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332): complete subsection reference.

<a id="canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports` properties

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

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2313323011121003-1232303312010301-1302021031030100-0122233322222201-0312131233132333-0330123333222332-1233220013331220-0002311311321031"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports`

- [ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120): complete subsection reference.

<a id="canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports` properties

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2300223233002013-3321203031012123-0023120103303122-1112320301300213-1303301313220212-0123101032201023-3230032323231120-3001211220022212"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports`

- [http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301): complete subsection reference.

- [port](data-sources--workload--reference--group-011.md#canonical-0330230310331300-1331023123323012-0102032330100131-3131300133300013-3111110113101200-0300012331333331-3113303133230233-2331220200130303): complete subsection reference.

- [tcp_loadbalancer](data-sources--workload--reference--group-011.md#canonical-3003113231211313-2011322221013301-0221321233231300-0201132111020033-2320132012211021-0112000312001132-3200330301031110-1310012311133030): complete subsection reference.

<a id="canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer` properties

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

Additional upstream details:

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

<a id="canonical-0020330301331112-2131120301331003-2322230321103011-0013220001022021-0020012203102202-1210312230332230-1132121213003212-3120310100003322"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer`

- [default_route](data-sources--workload--reference--group-008.md#canonical-2022000103331303-1031232003202312-0023231311103320-2310100000131002-3001332203021122-3121010312331311-1211030031023032-1312023030221100): complete subsection reference.

<a id="canonical-2032023210013012-2022300001020222-3232202003021011-2313011130110313-3221322322022333-1010321310101111-0113032311321003-0233022010013032"></a>

<a id="canonical-2333210202023331-0123123200013223-1322230322133011-2201233111003130-1022322132100002-1111203230033002-3132223310110120-1023101203110120"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.domains` property

Type: `["list", "string"]`. Computed.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\` is invalid
Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the..

Additional upstream details:

A list of domains (host/authority header) that will be matched to loadbalancer. Exact domain names:
\`\`www&#46;example.com\`\`. &#8203;2. Prefix domain wildcards: \`\`\*.example.com\`\` or
\`\`\*.bar.example.com\`\`. &#8203;3. Special wildcard \`\`\*\`\` matching any domain. Wildcard will
not match empty string. E.g. \`\`\*.example.com\`\` will match \`\`bar.example.com\`\` and
\`\`baz-bar.example.com\`\` but not \`\`.example.com\`\`. The longest wildcards match first.
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [specific_routes](data-sources--workload--reference--group-010.md#canonical-0300030213323103-0001012303030311-1001122302013331-2012030011300330-3111230203222132-2320002210311130-0230303000212113-2103000301201323): complete subsection reference.

<a id="canonical-2022000103331303-1031232003202312-0023231311103320-2310100000131002-3001332203021122-3121010312331311-1211030031023032-1312023030221100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route` properties

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

Additional upstream details:

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

<a id="canonical-0313201201332300-1221313110302311-0301332002220023-0333023022333302-1123233113202310-0011320110130321-1103020102010012-0100001313300202"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route`

- [auto_host_rewrite](data-sources--workload--reference--group-008.md#canonical-2110021010010110-2121003033102233-0033211230210001-0313011330302012-3201103033211133-3333133301000000-3031100031233302-0133303102030201): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-008.md#canonical-0111112200332001-3011102000302133-0130022000030203-3102321310101332-0210120202210311-3322332030012310-3013303101103203-3031120103213013): complete subsection reference.

<a id="canonical-0231222122303212-0322033211303233-1331313131122200-3011021321101112-2111323123223312-0012212120022110-1302110130123123-1222002003113113"></a>

<a id="canonical-1223200311230133-2110320231220021-2121322113132322-3122032003332033-2310120030022330-2132101123300131-1313013031223203-3222303011203311"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.host_rewrite` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2110021010010110-2121003033102233-0033211230210001-0313011330302012-3201103033211133-3333133301000000-3031100031233302-0133303102030201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite` properties

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

<a id="canonical-0111112200332001-3011102000302133-0130022000030203-3102321310101332-0210120202210311-3322332030012310-3013303101103203-3031120103213013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite` properties

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

<a id="canonical-1011120000322313-0203202303130222-0321020201312302-2103210110013011-2313120202130120-2023321230030113-0220312010313221-2101203130122322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http` properties

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

<a id="canonical-2033103000103023-3011203120312103-3210321310311203-3330202001320300-3233011313312210-2323032122131211-1213331322101201-2323020213202022"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http`

<a id="canonical-2302030233301032-0330133032110313-1232012310321120-1303201133330130-1320331001232000-1101211312320232-3130020322013230-2021301212330030"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http.dns_volterra_managed` property

Type: `"bool"`. Computed.

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

<a id="canonical-0231133002000102-1223312211012000-3321130200100030-0001130222022030-3210110021010220-2311221101123210-0132311330103010-1013313311010331"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http.port` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0313131033332133-2120112330113132-2332332311103322-1302200320233031-3310331312233023-1013021112201203-2102331230030113-0231333003233112"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1222321333022313-0131122211222120-3203023132311031-0223132013021320-3213031131013303-3013023321330231-0022112233333002-3110332101131120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https` properties

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

<a id="canonical-3233333332002003-0320022303213120-2100101220012001-0220132300323002-1323001302232330-2131001122221031-2010032021030112-1313021212013003"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https`

<a id="canonical-0221012130000100-0032331222332332-2200122332102332-0010123120111132-0121230010313211-2002123122301212-1133201233222122-2130223113001202"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.add_hsts` property

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

<a id="canonical-1100221220101231-2210303020301021-1301321103311002-1110110031102203-2111110012210033-0100000223113312-3020032020113102-0200210022000221"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.append_server_name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [coalescing_options](data-sources--workload--reference--group-008.md#canonical-1021033321200220-2122123222010113-3330010201103233-2103020220030330-2021200120310310-1302033032023022-1321103303100000-3212020230312013): complete subsection reference.

<a id="canonical-0023211102022100-2211010122300131-2021011312322221-1233230313103310-1132010212131020-2121130330223101-0232321302333132-1331310213333131"></a>

<a id="canonical-1320103220023001-0031202200321210-3101121311013003-0303200330130210-3130301130312213-0010220000313112-2203022132032120-1122030133303301"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.connection_idle_timeout` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [default_header](data-sources--workload--reference--group-008.md#canonical-1310111300201130-3300312200323211-1201232033332103-3222013101021223-0201122120303221-1310323210223030-1233300300012011-1003022011133120): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3001133332232322-1200321121321301-3200302113300201-3132003123223110-3211030323222301-0020231321013113-2123223023131023-3023220001323230): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-008.md#canonical-3202103311120012-2030332330132132-1320131200311223-1131012021212131-3031223333232020-0102201200031203-3031323220032133-2011203301101132): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-008.md#canonical-0302222110022333-1201233200130332-0000013212100201-1310010102120323-3312231232120302-3122110200301111-3212021331031120-0300033111013020): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-008.md#canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021): complete subsection reference.

<a id="canonical-3301230313221100-3100023123331103-1103023321220021-2112000223212102-0323003122012032-2232033000003020-2012001331010220-3011032111231030"></a>

<a id="canonical-0123231311010203-1333230112030000-0133323103302002-1102302120003230-1011212232032021-1211032310122133-1010323003121211-0000212322320331"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_redirect` property

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-1022320030333303-0030313102323203-2023213133322331-1001110102321313-1031220330111122-2020312112202202-0023230032031312-1300300133212303"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.port` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0331003313233211-2003023000302030-1012211211223333-2112301001320123-1002230200023223-0320220102131020-3121203333000333-0230021231310011"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0133210012103031-0121330302300010-3012201303121300-3202311212300112-2031131032013303-3321203102120133-2323130102232230-0030202102112011"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.server_name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1021033321200220-2122123222010113-3330010201103233-2103020220030330-2021200120310310-1302033032023022-1321103303100000-3212020230312013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options` properties

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

<a id="canonical-1000010001310223-1322322102203232-1031103232112201-2130210322003330-0100330331202131-3021213112303222-1211213300103123-3020233000102021"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options`

- [default_coalescing](data-sources--workload--reference--group-008.md#canonical-3200123202312121-2130211222213001-2030033011000112-3333201223200023-3010011103121031-0323231112232001-3020200220333112-3330331230311000): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-008.md#canonical-1122112303323203-3030011120030002-3312320202001030-3001223221203301-1222133322102300-2131130301122202-3102130201321233-1212112213300210): complete subsection reference.

<a id="canonical-3200123202312121-2130211222213001-2030033011000112-3333201223200023-3010011103121031-0323231112232001-3020200220333112-3330331230311000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-008.md#canonical-1021033321200220-2122123222010113-3330010201103233-2103020220030330-2021200120310310-1302033032023022-1321103303100000-3212020230312013)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-2202133203302123-2232332023103200-3201312013032213-2332313233002203-3003110213202201-2212033232011303-2011121221011222-1312112110120101"></a>

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

<a id="canonical-1122112303323203-3030011120030002-3312320202001030-3001223221203301-1222133322102300-2131130301122202-3102130201321233-1212112213300210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing` properties

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-008.md#canonical-1021033321200220-2122123222010113-3330010201103233-2103020220030330-2021200120310310-1302033032023022-1321103303100000-3212020230312013)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-2020113321312021-3221220311210321-0300123331302312-2230222021002011-2132132030102030-0200030220220032-1011223321332210-3110112300132233"></a>

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

<a id="canonical-1310111300201130-3300312200323211-1201232033332103-3222013101021223-0201122120303221-1310323210223030-1233300300012011-1003022011133120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header` properties

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

<a id="canonical-3001133332232322-1200321121321301-3200302113300201-3132003123223110-3211030323222301-0020231321013113-2123223023131023-3023220001323230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer` properties

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

<a id="canonical-3202103311120012-2030332330132132-1320131200311223-1131012021212131-3031223333232020-0102201200031203-3031323220032133-2011203301101132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize` properties

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

<a id="canonical-0302222110022333-1201233200130332-0000013212100201-1310010102120323-3312231232120302-3122110200301111-3212021331031120-0300033111013020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize` properties

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

<a id="canonical-1220000002121231-3310101111120203-0222002232113323-1133232022321203-0322003221301300-3320031030230332-2022301001030312-2132122232200021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options` properties

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

<a id="canonical-0012213301013320-2032202213010031-3000213221111101-3313200230030233-1322030300011011-0111312130310311-3310122210133102-2033123321321121"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options`

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-009.md#canonical-3210323323021023-2301233332012302-1132311331201003-0212201012311032-1233312303310322-1122313230211311-2013300200000020-0033323210303031): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-009.md#canonical-3300333110331030-2300220013021003-3211133121101120-3012231311030202-0312102232102223-1020202303022201-0030301332102213-0022223231322122): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-009.md#canonical-1303310302200120-1022220302203012-3030003321113020-2000202021010232-0130020133221331-3210002111033101-1312210103332222-1013230332320232): complete subsection reference.
