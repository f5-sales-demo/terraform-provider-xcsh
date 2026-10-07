---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-2322222302103220-3130301133020213-3333302101310233-3203322110202330-3320300130212122-2122330133110212-1202130031103232-3030123312302111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-010.md#canonical-0300030213323103-0001012303030311-1001122302013331-2012030011300330-3111230203222132-2320002210311130-0230303000212113-2103000301201323)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-010.md#canonical-1301202201003331-1301211323121333-1102331202010031-3213321100131310-2302202331233111-3130322011130113-0132211212022230-2333033221200311)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-010.md#canonical-2312231312302003-3311222223331033-1223203202222102-0121132330001310-1321210111233132-3133101111112333-2032112010023113-1210121203303311)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

<a id="canonical-1211322133103210-2012101133213333-3330213112131230-1103312231312021-2301322313131110-3200013033302011-0331200020000110-3122200002210332"></a>

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

<a id="canonical-1123033121310010-3203220322102112-3033122313310211-1132120011002000-3103103203310130-2301023223221202-2103123032132032-3203301310231031"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response`

<a id="canonical-3311323201201332-3012112211321202-1312332323002112-3313033230112120-0023013101210312-3230123112032000-3132203302203031-0221202130000331"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response.response_body_encoded` property

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

<a id="canonical-1000112033323132-2321202121231113-0133303203230103-2322220122211100-1122303220011302-2020213131022100-3110112100031223-0022022033130320"></a>

<a id="canonical-0010202102020230-2332202203102331-2133111120223223-2231032311211121-0303221230223331-1003230122231310-1021232322301201-1012221103110331"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response.response_code` property

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

<a id="canonical-3003020211101001-3000023123101123-0021222231012330-1021131023101202-2203221032123211-2111031313310310-3033321301221212-1231101123221212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-010.md#canonical-0300030213323103-0001012303030311-1001122302013331-2012030011300330-3111230203222132-2320002210311130-0230303000212113-2103000301201323)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-010.md#canonical-1301202201003331-1301211323121333-1102331202010031-3213321100131310-2302202331233111-3130322011130113-0132211212022230-2333033221200311)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-3203031121323010-1322211303023211-0100032021020031-0103100023312211-1201220013310230-2023030302020331-0213033212303232-1011233300030333"></a>

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

<a id="canonical-0200133203223212-2220203111031201-2202332013102132-2122131221322030-0102301133220201-1303003310122020-3311010313033330-1301111120232120"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route`

- [headers](data-sources--workload--reference--group-011.md#canonical-1232030032133112-2031302313022121-0102121132113100-3102103213130202-2011222220200232-3023130013111232-3103023303202021-2132311010022031): complete subsection reference.

<a id="canonical-3320112021300332-1221211031200203-0100131011021230-0221231311321232-2012012223233001-1102101313120133-2213202012113011-0211200331210133"></a>

<a id="canonical-0003333003111321-1212200023311220-1302100202100121-3011002213331230-3210302120223320-2230221102323012-0211022200032332-2321323203033002"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.http_method` property

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

- [incoming_port](data-sources--workload--reference--group-011.md#canonical-3011333220023003-0200220301333130-2211303012032220-0132320201330133-1303002112321201-1101110221120321-1310313103031320-1230013121100012): complete subsection reference.

- [path](data-sources--workload--reference--group-011.md#canonical-0012113212132311-2110332223312330-2312232313023333-2213021322102020-1110303033222200-1103332300302130-3120302223310132-2311013202120233): complete subsection reference.

- [route_redirect](data-sources--workload--reference--group-011.md#canonical-2112310211001211-2233000313201203-2100010112230011-0201322213113301-1203221121302230-2221003323020321-0113123200320320-0020132133303313): complete subsection reference.

<a id="canonical-1232030032133112-2031302313022121-0102121132113100-3102103213130202-2011222220200232-3023130013111232-3103023303202021-2132311010022031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-010.md#canonical-0300030213323103-0001012303030311-1001122302013331-2012030011300330-3111230203222132-2320002210311130-0230303000212113-2103000301201323)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-010.md#canonical-1301202201003331-1301211323121333-1102331202010031-3213321100131310-2302202331233111-3130322011130113-0132211212022230-2333033221200311)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-011.md#canonical-3003020211101001-3000023123101123-0021222231012330-1021131023101202-2203221032123211-2111031313310310-3033321301221212-1231101123221212)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-3121222221123311-1301032323102130-1011012112322130-3120120333303130-3322003101311111-0110003122133111-1330213203233230-2220030120130132"></a>

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

<a id="canonical-0230331231002303-0311032003302210-1102200330330012-1233301213003331-1202333222001130-1220301011202113-2030321233120232-2323202320223013"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers`

<a id="canonical-2122312302211023-0113011230003000-3032322000333223-1330103323132110-2221130311010330-3230321310303203-1222121031331133-0233310213033212"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.exact` property

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

<a id="canonical-1003102120300230-0032212330313210-1333221230221300-3022333221122021-3321121313001223-1023230332321322-1023102300000102-2120113331313121"></a>

<a id="canonical-2321000031013133-3333302312201301-1021303022030133-2202132021223330-1301021320031133-1321013200011333-3201120102300033-2103110320212202"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.invert_match` property

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

<a id="canonical-2200031301312123-2332313330130233-0010012310320203-3102122322002313-3010300020130102-3230213230333022-0032210231323002-0311323132221120"></a>

<a id="canonical-1320331330121320-0012033213330132-1130002230230111-1131003031313311-1100330002031203-3122211023033113-2203001232231002-0330002322021222"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.name` property

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

<a id="canonical-0132301032101033-0212222301100012-2302131133030002-2100202012333211-1112010302233302-0100201332302311-3000131110210222-3012011030333232"></a>

<a id="canonical-0003223321233121-3323322103333002-0100310210121131-0011130030220200-2013231122213031-2022333211323123-1122031111022131-1110013031313120"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.presence` property

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

<a id="canonical-0132301030232320-1032302130331120-1220232031301333-2210013120203302-2312203103201033-1222202003200213-2201320102300130-1333000101000001"></a>

<a id="canonical-0203321021222010-2213313030220303-2121231310233233-1100130321203010-0222023003200211-0122231130133000-0233002202100223-3000311121011221"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.regex` property

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

<a id="canonical-3011333220023003-0200220301333130-2211303012032220-0132320201330133-1303002112321201-1101110221120321-1310313103031320-1230013121100012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-010.md#canonical-0300030213323103-0001012303030311-1001122302013331-2012030011300330-3111230203222132-2320002210311130-0230303000212113-2103000301201323)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-010.md#canonical-1301202201003331-1301211323121333-1102331202010031-3213321100131310-2302202331233111-3130322011130113-0132211212022230-2333033221200311)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-011.md#canonical-3003020211101001-3000023123101123-0021222231012330-1021131023101202-2203221032123211-2111031313310310-3033321301221212-1231101123221212)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port

<a id="canonical-0001313120132111-3032022303120032-0123011301020121-2310102300031300-2333322320122021-3012111230121220-0130313312130123-2320033000000112"></a>

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

<a id="canonical-0210311300113033-0132102233121113-3322221030310131-3132130022122303-3000221330221210-0131100012131200-2331220323030221-2200220200231131"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port`

- [no_port_match](data-sources--workload--reference--group-011.md#canonical-0131022300122223-2322321211221012-2003020233201323-3011000312201022-1203121203200211-1101230302022202-0321030011311213-2023032012221111): complete subsection reference.

<a id="canonical-0212322322232221-1302102023021121-3213213022313033-0012333211233320-3031101132131011-3100233020203012-0322300203202220-2310210312001132"></a>

<a id="canonical-3112031202023031-3031200022102111-1203100100211012-1331020232213203-1011101330000323-0022111223211332-1113011322203330-2133332023122030"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.port` property

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

<a id="canonical-3233211332013020-2200013222231001-1231011023022120-2122101200331112-2323330031322302-0002231102003012-1131000122132033-2120003321110300"></a>

<a id="canonical-1321232222230212-3322030200331310-1110101230331322-0111301130010300-1300020310313232-2001233021231023-2321111112323002-2332221303221221"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.port_ranges` property

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

<a id="canonical-0131022300122223-2322321211221012-2003020233201323-3011000312201022-1203121203200211-1101230302022202-0321030011311213-2023032012221111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-010.md#canonical-0300030213323103-0001012303030311-1001122302013331-2012030011300330-3111230203222132-2320002210311130-0230303000212113-2103000301201323)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-010.md#canonical-1301202201003331-1301211323121333-1102331202010031-3213321100131310-2302202331233111-3130322011130113-0132211212022230-2333033221200311)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-011.md#canonical-3003020211101001-3000023123101123-0021222231012330-1021131023101202-2203221032123211-2111031313310310-3033321301221212-1231101123221212)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-011.md#canonical-3011333220023003-0200220301333130-2211303012032220-0132320201330133-1303002112321201-1101110221120321-1310313103031320-1230013121100012)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-0023322010103002-3033000100201123-1301231201033200-1201102202201133-0201213131321020-1123022223221000-3323030312200130-1213020020322203"></a>

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

<a id="canonical-0012113212132311-2110332223312330-2312232313023333-2213021322102020-1110303033222200-1103332300302130-3120302223310132-2311013202120233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-010.md#canonical-0300030213323103-0001012303030311-1001122302013331-2012030011300330-3111230203222132-2320002210311130-0230303000212113-2103000301201323)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-010.md#canonical-1301202201003331-1301211323121333-1102331202010031-3213321100131310-2302202331233111-3130322011130113-0132211212022230-2333033221200311)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-011.md#canonical-3003020211101001-3000023123101123-0021222231012330-1021131023101202-2203221032123211-2111031313310310-3033321301221212-1231101123221212)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path

<a id="canonical-3311221133011022-2332330002130322-3103321032121301-1131100202333111-3020030031331023-0213130111130223-1320002012332301-1133301320000311"></a>

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

<a id="canonical-1232111203222110-2030123221021102-0322003333131032-1031220222333001-1231330012013123-2022212313231010-3231333311120301-1201211332121010"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path`

<a id="canonical-3130321232210122-1330113322012023-0322300200011222-1103301310131202-3310311321330210-1222320231111212-3232030321221232-2012200312031320"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path.path` property

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

<a id="canonical-1201002111331012-1013202203113012-3113301002333032-1033201320231123-0132133032303231-1300132302103230-2221020012002200-2101120222223010"></a>

<a id="canonical-2302013120003212-1033232001201013-1002020033000321-3313300210110323-2223111212120310-3202211332013303-3020320132222020-0203223201131320"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path.prefix` property

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

<a id="canonical-3211100302201103-1130313210300210-1103302021103201-1132202030130321-0213010220120321-0121021020010010-0333213213200100-3211203001100131"></a>

<a id="canonical-2200301333132002-3321331003223111-1133012003323110-1012330112200032-1102233020230023-1113132102021312-2121003223220210-3230323200002132"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path.regex` property

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

<a id="canonical-2112310211001211-2233000313201203-2100010112230011-0201322213113301-1203221121302230-2221003323020321-0113123200320320-0020132133303313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-010.md#canonical-0300030213323103-0001012303030311-1001122302013331-2012030011300330-3111230203222132-2320002210311130-0230303000212113-2103000301201323)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-010.md#canonical-1301202201003331-1301211323121333-1102331202010031-3213321100131310-2302202331233111-3130322011130113-0132211212022230-2333033221200311)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-011.md#canonical-3003020211101001-3000023123101123-0021222231012330-1021131023101202-2203221032123211-2111031313310310-3033321301221212-1231101123221212)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect

<a id="canonical-0022210003100321-0323000211310003-2103213121112032-3322332111133231-3232112120231023-0122310023212010-3101232002020231-3312331030013212"></a>

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

<a id="canonical-3121313333111110-3230100002230023-2030301320230120-0223030131131111-2021003320303123-2301111100323300-0122002102000022-2001300131302003"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect`

<a id="canonical-2110022123320122-0301313012111332-3200112130232103-2111130213323033-3102110212212020-0200202013113100-1200101322022220-2013322013213003"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.host_redirect` property

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

<a id="canonical-0023133121322231-0031120003033111-0213122010213202-3123021303233313-1032123112301132-0021011013311233-2200103120010031-0303013111132122"></a>

<a id="canonical-3323211213230003-2202233313333320-3102203301102231-0320113230132111-2013201111013233-3301000331101130-1203320113020120-3021310222221201"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.path_redirect` property

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

<a id="canonical-2011023232212033-1001310133033001-3321302221021022-3332303003312130-1211020332133313-2301010200313120-2011212032110213-3331022232002231"></a>

<a id="canonical-0110202010312313-2330123022303033-1200120332011102-3300212233203001-2210230201032320-2310233301112210-2022300232213010-1013121233330220"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.prefix_rewrite` property

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

<a id="canonical-2023002300100301-0011302011201202-3032130023022032-2122130113320231-2231330233203222-0102211311201121-2220333332000310-1123330021320332"></a>

<a id="canonical-2300120330210323-0123032033010332-3321333130032121-0202111122200301-2213132302131303-3212232130212200-0113121112122132-1031003200130123"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.proto_redirect` property

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

- [remove_all_params](data-sources--workload--reference--group-011.md#canonical-3113312320201301-2013103321230332-3032123133022302-1303321202033213-0222120222012232-0120320030130323-3031313123330223-2103001301220011): complete subsection reference.

<a id="canonical-0302100003223213-3003333032223023-1233112322100113-3112030300030321-0202112223032210-0212220231123013-3112022211112123-0320210321013332"></a>

<a id="canonical-2103132010233211-3300012210031000-2111301323000221-0001102222233002-0121330111310022-1210121331330312-0103311010100002-2100301121102302"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.replace_params` property

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

<a id="canonical-2022010112131120-2022001131113332-3222313102322211-0030301101310310-3312133032101012-3303131211211222-3213320230313300-3030330331203010"></a>

<a id="canonical-2231120333000333-1133311303122300-2032322223103302-2303122023120201-2201111031213100-1033113232023323-0102322010201230-0122130102223221"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.response_code` property

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

- [retain_all_params](data-sources--workload--reference--group-011.md#canonical-2320033003220310-3300301201121020-0000123101100132-3023001000100320-1113101133002122-2211133111032200-2220003232123232-2330213302311131): complete subsection reference.

<a id="canonical-3113312320201301-2013103321230332-3032123133022302-1303321202033213-0222120222012232-0120320030130323-3031313123330223-2103001301220011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-010.md#canonical-0300030213323103-0001012303030311-1001122302013331-2012030011300330-3111230203222132-2320002210311130-0230303000212113-2103000301201323)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-010.md#canonical-1301202201003331-1301211323121333-1102331202010031-3213321100131310-2302202331233111-3130322011130113-0132211212022230-2333033221200311)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-011.md#canonical-3003020211101001-3000023123101123-0021222231012330-1021131023101202-2203221032123211-2111031313310310-3033321301221212-1231101123221212)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-011.md#canonical-2112310211001211-2233000313201203-2100010112230011-0201322213113301-1203221121302230-2221003323020321-0113123200320320-0020132133303313)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-1132231020300201-2120201130132312-1233312231231111-1020201031212221-1232113331321221-1221132122011323-1003120210310213-0122022110032102"></a>

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

<a id="canonical-2320033003220310-3300301201121020-0000123101100132-3023001000100320-1113101133002122-2211133111032200-2220003232123232-2330213302311131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-010.md#canonical-0300030213323103-0001012303030311-1001122302013331-2012030011300330-3111230203222132-2320002210311130-0230303000212113-2103000301201323)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-010.md#canonical-1301202201003331-1301211323121333-1102331202010031-3213321100131310-2302202331233111-3130322011130113-0132211212022230-2333033221200311)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-011.md#canonical-3003020211101001-3000023123101123-0021222231012330-1021131023101202-2203221032123211-2111031313310310-3033321301221212-1231101123221212)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-011.md#canonical-2112310211001211-2233000313201203-2100010112230011-0201322213113301-1203221121302230-2221003323020321-0113123200320320-0020132133303313)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-3001122212223001-0111101120022231-0203200023112000-0130101112133311-1332020112313332-0223201311323012-2013211320203312-2201122200123221"></a>

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

<a id="canonical-2130221102331122-1012320122203102-3100201333233302-2132132123121323-3002222320133330-1313213230022210-3102113201233322-1302333032130331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-010.md#canonical-0300030213323103-0001012303030311-1001122302013331-2012030011300330-3111230203222132-2320002210311130-0230303000212113-2103000301201323)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-010.md#canonical-1301202201003331-1301211323121333-1102331202010031-3213321100131310-2302202331233111-3130322011130113-0132211212022230-2333033221200311)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-3031302110323132-2110023232130121-0130002211111310-2132303300333211-1331320230301332-1203330032030223-1232120111032033-1323303220120330"></a>

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

<a id="canonical-1323231131030102-0000121213333232-3301111330310122-3133230333201012-3221032332121010-2000202030310303-2002000302112110-0310012202201113"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route`

- [auto_host_rewrite](data-sources--workload--reference--group-011.md#canonical-1200302311013033-0020301230111312-1030312003322211-2332021023330133-3321110111323002-1332032010003013-0030101320103031-0103322330021320): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-011.md#canonical-3033213200223121-1113122232113123-2023130320130012-3211030131323323-1322011220231230-2032210213030233-2012021121201302-2131231030113130): complete subsection reference.

<a id="canonical-0331132203010021-2133031010120000-1100103123222313-2333220331212100-3101021030221330-0320311113032033-2313313122302103-2330011202233131"></a>

<a id="canonical-1111313021112230-1202100303321020-1113302010000230-0130100030010121-1312030303213110-1023220113103003-0303011312220331-2022321210320133"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.host_rewrite` property

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

<a id="canonical-3231203330113200-2030120303023233-1001020331010311-3112212322013200-0010201122131001-0313310121233011-1111111101322023-1200120323300331"></a>

<a id="canonical-0310033110022003-0102030303003101-3011332100202112-0031113131203121-1223012033322331-0031302022110221-1331002330213123-0013233203212322"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.http_method` property

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

- [path](data-sources--workload--reference--group-011.md#canonical-0022132312230200-3212130123203122-3203111112111330-0312012310120333-0013030101013122-0000102111221012-1220110232011210-3002103031131033): complete subsection reference.

<a id="canonical-1200302311013033-0020301230111312-1030312003322211-2332021023330133-3321110111323002-1332032010003013-0030101320103031-0103322330021320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-010.md#canonical-0300030213323103-0001012303030311-1001122302013331-2012030011300330-3111230203222132-2320002210311130-0230303000212113-2103000301201323)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-010.md#canonical-1301202201003331-1301211323121333-1102331202010031-3213321100131310-2302202331233111-3130322011130113-0132211212022230-2333033221200311)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-011.md#canonical-2130221102331122-1012320122203102-3100201333233302-2132132123121323-3002222320133330-1313213230022210-3102113201233322-1302333032130331)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-3122001311121203-3031331003231203-2030303133212322-2313223312202002-2032310320312033-2322111132120300-3023332223201201-3322100232012302"></a>

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

<a id="canonical-3033213200223121-1113122232113123-2023130320130012-3211030131323323-1322011220231230-2032210213030233-2012021121201302-2131231030113130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-010.md#canonical-0300030213323103-0001012303030311-1001122302013331-2012030011300330-3111230203222132-2320002210311130-0230303000212113-2103000301201323)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-010.md#canonical-1301202201003331-1301211323121333-1102331202010031-3213321100131310-2302202331233111-3130322011130113-0132211212022230-2333033221200311)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-011.md#canonical-2130221102331122-1012320122203102-3100201333233302-2132132123121323-3002222320133330-1313213230022210-3102113201233322-1302333032130331)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-3021323003123100-0331301103110110-3002310302111100-2011113100300323-0030031231303212-3023201003321020-3010132013023300-0022332220200120"></a>

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

<a id="canonical-0022132312230200-3212130123203122-3203111112111330-0312012310120333-0013030101013122-0000102111221012-1220110232011210-3002103031131033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-008.md#canonical-3023111120210202-0313023112032123-2033210311120213-3000313003000213-2203112012310120-3112303221031323-1313230112020333-0022201322102301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-010.md#canonical-0300030213323103-0001012303030311-1001122302013331-2012030011300330-3111230203222132-2320002210311130-0230303000212113-2103000301201323)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-010.md#canonical-1301202201003331-1301211323121333-1102331202010031-3213321100131310-2302202331233111-3130322011130113-0132211212022230-2333033221200311)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-011.md#canonical-2130221102331122-1012320122203102-3100201333233302-2132132123121323-3002222320133330-1313213230022210-3102113201233322-1302333032130331)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-3032330132123333-3020320312313102-0100332033210212-0010223013332312-0013101122100312-1010320123023211-2012330023012001-2113031202002323"></a>

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

<a id="canonical-2002011132110103-0333133111311102-0011132200301321-3212231111301212-1210132223111223-1211013022231022-3312103131233230-2220330320002333"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path`

<a id="canonical-2202222022122211-0021011233332321-0211220322231010-3001030013211013-3131110333302120-2303111000201330-1130123032120303-2030031133031020"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path.path` property

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

<a id="canonical-3211330110130300-0010322223221223-1120003222210003-3000122022200032-0202322203121230-3120121000330032-2122313101301232-2222311202311130"></a>

<a id="canonical-0010132220331022-2131212011103310-3323030332122310-0130323001323320-0322003032300321-0010301313002231-2330311110033001-3201101111322321"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path.prefix` property

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

<a id="canonical-1212211312203221-3322032231211312-1310122201330302-0321300122023003-0020212102111311-0032221111103311-3200020333232332-2103223030213233"></a>

<a id="canonical-3030303002231212-1203011310123002-2300220202322322-1322112101122100-1322310310232202-3032031203032133-1020111302322202-3002100332300201"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path.regex` property

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

<a id="canonical-0330230310331300-1331023123323012-0102032330100131-3131300133300013-3111110113101200-0300012331333331-3113303133230233-2331220200130303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- service.advertise_options.advertise_on_public.multi_ports.ports.port

<a id="canonical-0230311000110301-1232031230133031-1311333310101003-0022213023220212-1221222013122032-0230122313123121-0311031121020231-1301121023110332"></a>

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

<a id="canonical-1023022201212000-2100211212102020-0301110010121012-3321021130202212-0212312220010011-1003012131131213-1300211201303101-3113323101223002"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.port`

- [info](data-sources--workload--reference--group-011.md#canonical-2222012120323123-0001121202333302-3233130223021022-1300331303231101-1132120112122132-1103213132202300-2012013331310311-0302023121330301): complete subsection reference.

<a id="canonical-0120322032020110-2012220322032102-1102133322132011-1022132101221301-2331012213103322-3123220210032013-0030210003133122-2013113101333012"></a>

<a id="canonical-1031210311010233-3200203310003203-3300213310203321-1020101200100212-2203330110210233-1221123110003010-0211303332201133-0103012011320010"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.port.name` property

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

<a id="canonical-2222012120323123-0001121202333302-3233130223021022-1300331303231101-1132120112122132-1103213132202300-2012013331310311-0302023121330301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.port.info` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port](data-sources--workload--reference--group-011.md#canonical-0330230310331300-1331023123323012-0102032330100131-3131300133300013-3111110113101200-0300012331333331-3113303133230233-2331220200130303)
- service.advertise_options.advertise_on_public.multi_ports.ports.port.info

<a id="canonical-3322012032033203-0122103122301002-2210230101100022-3230302323302220-2202022311102333-2011210120230303-3131300210120331-0110011213130023"></a>

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

<a id="canonical-1032001201303130-0002211133302012-3321021023203031-1233021203103113-1012330121320110-0333000001020032-1301302311011031-3232213100120333"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.port.info`

<a id="canonical-1303220220021321-3022012122031111-1302012021121113-3103221332332111-0323211222010121-3322330003022102-2230021002320320-3120110301211332"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.port.info.port` property

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

<a id="canonical-1221321301121203-2102131232231101-3332210123332230-2021311100323332-0202200121303013-3000330322122302-0131220331111233-0301103110130001"></a>

<a id="canonical-3033030220302113-0001312111202132-1102021310233312-3011000113300030-0021222221333201-2123030203023120-3322101223322233-0223011310230130"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.port.info.protocol` property

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

- [same_as_port](data-sources--workload--reference--group-011.md#canonical-2211203133030111-3332202321102020-3323031200223121-3300233103103011-0132120033002121-0033123321310300-2113103021023130-2310010031330032): complete subsection reference.

<a id="canonical-0022120212321001-1030121301201233-0033030211003001-0202302312303102-1113302013112300-2120232103220201-1102303003332230-2310312100000330"></a>

<a id="canonical-0020313003222203-3230022123323202-0211332113010033-0003302113030103-3010130230311320-0332332202100120-3012303002132110-2110200002333030"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.port.info.target_port` property

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

<a id="canonical-2211203133030111-3332202321102020-3323031200223121-3300233103103011-0132120033002121-0033123321310300-2113103021023130-2310010031330032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port](data-sources--workload--reference--group-011.md#canonical-0330230310331300-1331023123323012-0102032330100131-3131300133300013-3111110113101200-0300012331333331-3113303133230233-2331220200130303)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port.info](data-sources--workload--reference--group-011.md#canonical-2222012120323123-0001121202333302-3233130223021022-1300331303231101-1132120112122132-1103213132202300-2012013331310311-0302023121330301)
- service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port

<a id="canonical-1330232103002033-0100303133020113-0213033002100103-1103122201023131-2101333131012021-3023201322113033-1011012122033310-2010220331320113"></a>

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

<a id="canonical-3003113231211313-2011322221013301-0221321233231300-0201132111020033-2320132012211021-0112000312001132-3200330301031110-1310012311133030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-008.md#canonical-3131133232121020-0203302310121231-1131113230212030-2111103313010330-2131230112223203-2311311213122322-2023221000002321-0122010021022012)
- [service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-008.md#canonical-0021133101121333-1210320223100220-3133030021012212-1301020332003023-2213120102203020-0023301100331122-2210121303201321-0211331333333120)
- service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer

<a id="canonical-1203233220212110-2311003230301133-0232213233130303-1101233203102220-3112113102000003-3233033221332120-0131020323122223-3303233322113231"></a>

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

<a id="canonical-3122302002110110-1223310222331230-3132130113312332-2133131110012332-3310230110022213-0101203232123110-0220122101233323-0220120011211002"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer`

<a id="canonical-2123300322012121-0023233231233201-3302200112032200-2003001031003320-2112230232203222-3212220032113010-3032002022331111-1030031011111322"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer.domains` property

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

<a id="canonical-0323220210122222-0333310000233222-0100310132020201-0310020222202330-2010202213222132-3222231022111301-1331321221310020-3100200112021100"></a>

<a id="canonical-3123311110220032-1320112330221313-3032003110113031-1320302323233201-0232333010331020-3333233020310232-2112220003123210-1221232332130113"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer.with_sni` property

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

<a id="canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- service.advertise_options.advertise_on_public.port

<a id="canonical-1103002223113210-1011033133113021-0010102010033111-3113131231213121-0131122003103223-2131300223013312-0311020201112312-0130002130123330"></a>

Type: `"single"`. Computed.

Advertise Port. Advertise single port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"http_loadbalancer\",\"tcp_loadbalancer\"]"
}
```

<a id="canonical-2120320202212213-2300202322333202-3003031010203113-2031023110302031-0232323101331103-2223333132303210-1013221323213310-2123120221312101"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port`

- [http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200): complete subsection reference.

- [port](data-sources--workload--reference--group-013.md#canonical-0113001100231030-3133302301102130-2211201121032213-2012110310122033-1330211302020133-2012113020123300-0302331301232310-3311330330300223): complete subsection reference.

- [tcp_loadbalancer](data-sources--workload--reference--group-013.md#canonical-0032120222121011-0103113003232030-3033201310032302-2121320210222122-2220130011002330-2032233123020302-1011100012101223-1313103122322332): complete subsection reference.

<a id="canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- service.advertise_options.advertise_on_public.port.http_loadbalancer

<a id="canonical-3100120112211102-1133310221030311-0033223102031120-3121203232330113-0222311103201003-2310101121323233-3113113310323331-1302200310112020"></a>

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

<a id="canonical-2213112100230320-2113012002333332-0323220231322100-0303313323103331-0100303330212320-3100011211333212-3103103202233002-0031103320213022"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer`

- [default_route](data-sources--workload--reference--group-011.md#canonical-2011112103113311-0021022003131201-1130232210322130-3010310110133133-1222133200311031-0203221120201320-0232101233111133-0130000201020203): complete subsection reference.

<a id="canonical-2100112110333313-1220220223310101-3031232031223203-3322201231321120-0112321330300123-0120312201311032-2120202303302211-2101132332202212"></a>

<a id="canonical-3012310331330012-1101033210220210-1010032322230310-3133231210123101-0300230133223233-3221201201201330-2303213100131323-1132122230201130"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.domains` property

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

- [http](data-sources--workload--reference--group-011.md#canonical-1212333320321331-0031212212133120-0111313221222323-0130130201212132-2022031123333023-1120133203320133-0113003012330130-3331313232312033): complete subsection reference.

- [https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122): complete subsection reference.

- [https_auto_cert](data-sources--workload--reference--group-012.md#canonical-3221332032012331-1012112033312231-1200002111331012-2001011223231300-0012132110021331-2132113113032321-3203030300133203-3302233232313100): complete subsection reference.

- [specific_routes](data-sources--workload--reference--group-013.md#canonical-1030212223221121-2003313010132300-1020102130221221-1020323100331032-0232122121201212-1320322100030002-2100103213001013-3120331131112333): complete subsection reference.

<a id="canonical-2011112103113311-0021022003131201-1130232210322130-3010310110133133-1222133200311031-0203221120201320-0232101233111133-0130000201020203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route

<a id="canonical-3210221010022001-3212331323331221-3231220020323323-1001223102310133-0012123233303230-1030031203102312-1221203132232221-1301102332101322"></a>

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

<a id="canonical-2122112101130131-3232121111031231-0031033113333010-0122030021333002-1123112021323322-2331210303303233-0131331000013301-1110313213011311"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route`

- [auto_host_rewrite](data-sources--workload--reference--group-011.md#canonical-3132130333122203-2020232211121312-1122200020211033-0131300113012232-1212022132013011-2303313313113321-1210001303022111-0023320103332203): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-011.md#canonical-1223202010102232-2003223132022123-2121112133220111-1133231003113333-0332310001210012-3102032130213110-0200332111203321-1210233000012320): complete subsection reference.

<a id="canonical-3232003210033010-2222120133302330-2330132123201121-3030311122122111-3022103333001312-3031022203230233-1031133201012301-2311130031201320"></a>

<a id="canonical-0300001030201130-0110223121031101-1313023231333012-2122213310102331-0030112121110322-0133211001310111-3210001212212012-1330103302300031"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.host_rewrite` property

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

<a id="canonical-3132130333122203-2020232211121312-1122200020211033-0131300113012232-1212022132013011-2303313313113321-1210001303022111-0023320103332203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](data-sources--workload--reference--group-011.md#canonical-2011112103113311-0021022003131201-1130232210322130-3010310110133133-1222133200311031-0203221120201320-0232101233111133-0130000201020203)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-2032121003321011-0222120031100033-1221130020122303-0011233021301212-3103121000301123-0011000323200233-1323322013130200-0031102331121301"></a>

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

<a id="canonical-1223202010102232-2003223132022123-2121112133220111-1133231003113333-0332310001210012-3102032130213110-0200332111203321-1210233000012320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](data-sources--workload--reference--group-011.md#canonical-2011112103113311-0021022003131201-1130232210322130-3010310110133133-1222133200311031-0203221120201320-0232101233111133-0130000201020203)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-1223323022010022-3311232303010121-1120122133300010-3131112132333201-0233010100231332-3213132033120013-3023320233320233-2121322120020331"></a>

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

<a id="canonical-1212333320321331-0031212212133120-0111313221222323-0130130201212132-2022031123333023-1120133203320133-0113003012330130-3331313232312033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.http` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.http

<a id="canonical-1210010032233000-0000011303000300-2102210122221113-1300011220322132-2002002032202131-3031320320111102-0310320110011322-1220030122212231"></a>

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

<a id="canonical-2010030110333000-2001110033033011-0212022303112220-1210321201121020-1120201012132103-2230102012212111-2012321201313233-2301320123011303"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.http`

<a id="canonical-2231230310101130-1103313111213100-1102022232132133-1301031300320113-1301330203000312-1120020323133232-2030102032001212-2222011203221332"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.http.dns_volterra_managed` property

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

<a id="canonical-1010021222102233-0222223123200232-3320201031320330-1212030222030111-1132100022212002-2320030113000323-2130220300110311-0030321233112223"></a>

<a id="canonical-1101231100023111-0113113333113233-3032113201032112-2311210311322120-1002100321313020-1202202031021123-1113102322202232-2231120331000111"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.http.port` property

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

<a id="canonical-2201100101313010-2221303322001111-0012213221232110-1111333332320320-0020333130233003-1132212203210001-0333103331301330-0101301230122022"></a>

<a id="canonical-3023313221013300-2232211003322333-3311333120300300-1211202001033303-1303312002312310-3220002113230201-1103003300303313-2323311200221213"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.http.port_ranges` property

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

<a id="canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https

<a id="canonical-0020123222023221-0003121310311300-1032131132010000-2321213113000101-0300122221333123-0212110101333331-1033022013002233-1313023332231300"></a>

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

<a id="canonical-1110203220330323-3320022233111031-1023220321100132-0330211112330332-2003222110103320-3331032210323203-2121023002103211-1330011121230031"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https`

<a id="canonical-1121022111003300-0122330102131310-3120230301231003-1133233020021223-3301312203122112-0203112230003321-3212031310220230-0010231233103000"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.add_hsts` property

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

<a id="canonical-0101331330200233-3131302222203212-1222021131001202-0313221132102133-2213120322311230-2122012330310110-2003122320022310-2202201223221110"></a>

<a id="canonical-0021212302330231-1222010132000121-2200133332232311-2300232223011310-0220212121132311-2020002230003011-0033113102132032-2230103111231020"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.append_server_name` property

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

- [coalescing_options](data-sources--workload--reference--group-011.md#canonical-0230013303021021-0331203200311333-1323230212321233-2120100120333023-0132133030032010-1331300220320331-3303022033022311-3112201130111321): complete subsection reference.

<a id="canonical-3220200203312310-3212313031302220-3333120331000010-3200220303332121-0121113011003200-0233031200020001-0102302213323323-1031223202123311"></a>

<a id="canonical-2032013301200102-0322200233003103-1200203330330000-3331312020111121-0131131231203311-1111133331303232-1210300113310110-2011200211111320"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.connection_idle_timeout` property

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

- [default_header](data-sources--workload--reference--group-011.md#canonical-3303013033220231-1233322321222132-3100133022301321-3302003333221123-1201021311111333-1101103210321202-1321233012322333-3133020122201033): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-011.md#canonical-2220012221003130-2131333230020230-3212320032110331-3120233110300233-2222232203200322-2313200220022230-3312321013003122-2232231010121231): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-011.md#canonical-3310212101313303-0211232303100020-0332302121110001-3030031300322203-2212313031221332-1021030231030313-0202002302020013-3211032313033000): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-011.md#canonical-2011112233003330-2030332211101120-0323212222213121-2113311001011232-2130020222033233-2102310102031303-1112223232323011-3131123303012201): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-011.md#canonical-0110313113231130-3320002222210031-0012202203321021-0020311322133013-2001022032311201-0123102123030220-0000333200021112-2320221302033100): complete subsection reference.

<a id="canonical-1121213111112100-1222113200321001-3021301011303102-3303111233201132-3110013333300312-0311301322302230-0331102011200121-0001203232033332"></a>

<a id="canonical-1010010120202302-3321121110313023-2332210200320012-0103003303201000-0121123312313011-3323110302213201-0320022012102303-0130021213020023"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_redirect` property

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

- [non_default_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1203232221031221-2133011011023111-1003123033102233-0321212010230202-1003301113220310-3112222211223301-2012211312132310-2202011212002320): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-011.md#canonical-0323132102131121-2302300122020021-3312200023130311-2221102131322111-0221213321301300-1203112322222300-1001331303303200-1320112210021020): complete subsection reference.

<a id="canonical-1323100000311102-3330233331030330-1031122031100323-3231223322213011-1010101211011010-1320111210303130-0123003112010101-3200013113310032"></a>

<a id="canonical-2033201321113333-0220130201101122-1132312120211000-1322322120032101-2301333231021020-0231211231220333-1210330023020201-0312310303222203"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.port` property

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

<a id="canonical-0320122212331302-2021200210230132-1100123320332312-2010123221112031-3112102300111202-3131333330302220-2120033000113310-1333100330220133"></a>

<a id="canonical-3102020112322011-2331212013130313-0133200112030311-3121303033011103-2332222221321223-3300330122200000-1212030011310012-2201023223221301"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.port_ranges` property

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

<a id="canonical-2313300131302200-0230123223230011-1323132210013011-3013100232202103-0130220223320122-3011110230000300-1212102012323322-1122323102201233"></a>

<a id="canonical-3331213312000233-3323013331332131-3313222111133303-2223220203123321-1013100332103001-2233211122003221-0232311012230001-2211110102320031"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.server_name` property

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

- [tls_cert_params](data-sources--workload--reference--group-011.md#canonical-0200231313301230-1020120000303023-2310213213031220-3211310113302002-2013232220213000-3320133031232013-1021231020010121-0000222220020220): complete subsection reference.

- [tls_parameters](data-sources--workload--reference--group-012.md#canonical-1103313212320303-2333031011021023-0021002220023123-0233132312230101-3120022121303232-0010310020022121-0031303022331012-0312231130212211): complete subsection reference.

<a id="canonical-0230013303021021-0331203200311333-1323230212321233-2120100120333023-0132133030032010-1331300220320331-3303022033022311-3112201130111321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options

<a id="canonical-0110133202212102-2000122213211103-2201302010001322-1121100200001221-1223232111011121-3111302331122200-0112323013220320-2310112223031031"></a>

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

<a id="canonical-3321320013001300-2011133221023212-3311000121112013-2030323220100102-3331101200020021-2022003220010122-1032231322213201-0102021010112020"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options`

- [default_coalescing](data-sources--workload--reference--group-011.md#canonical-1132212100101122-0002221220221332-1200312121320122-1012322211231002-1222111011313233-0013333313221222-2011111310311033-2221312301132112): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-011.md#canonical-3100301031321332-3022201211222023-1103200320100011-1011110120122311-1121300111000013-3023133300222203-3322311123132002-2220001212220311): complete subsection reference.

<a id="canonical-1132212100101122-0002221220221332-1200312121320122-1012322211231002-1222111011313233-0013333313221222-2011111310311033-2221312301132112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-011.md#canonical-0230013303021021-0331203200311333-1323230212321233-2120100120333023-0132133030032010-1331300220320331-3303022033022311-3112201130111321)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-0122210130002032-3032031131202111-0030101021012012-2001330220210333-0232102021202032-2023000200113302-3312120100330130-2312002012112331"></a>

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

<a id="canonical-3100301031321332-3022201211222023-1103200320100011-1011110120122311-1121300111000013-3023133300222203-3322311123132002-2220001212220311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-011.md#canonical-0230013303021021-0331203200311333-1323230212321233-2120100120333023-0132133030032010-1331300220320331-3303022033022311-3112201130111321)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-1022013101210303-0000220001301232-3021021111022312-1333032113110322-3031323010311313-1231103121000322-1333031022333000-3233331213312130"></a>

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

<a id="canonical-3303013033220231-1233322321222132-3100133022301321-3302003333221123-1201021311111333-1101103210321202-1321233012322333-3133020122201033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header

<a id="canonical-1222000012021220-3322000000110002-1302033323110002-0112222302333202-2211100223013101-0330103312230331-1000212331321120-1211130213211213"></a>

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

<a id="canonical-2220012221003130-2131333230020230-3212320032110331-3120233110300233-2222232203200322-2313200220022230-3312321013003122-2232231010121231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer

<a id="canonical-1003331233112232-1002133210032113-1210301120231113-3003112303333310-2110130020131011-0220302000231300-3023010011202301-0321000003322021"></a>

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

<a id="canonical-3310212101313303-0211232303100020-0332302121110001-3030031300322203-2212313031221332-1021030231030313-0202002302020013-3211032313033000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize

<a id="canonical-3213110212310133-1002102112003331-1100310321020221-2012011001020300-1210032103002131-2102212021012111-3203112112203033-1030230312313000"></a>

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

<a id="canonical-2011112233003330-2030332211101120-0323212222213121-2113311001011232-2130020222033233-2102310102031303-1112223232323011-3131123303012201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize

<a id="canonical-0331302033203220-0022323023102322-1213023013223031-3312331203021201-2222321112222022-1332103133232022-0222013000233223-2123103000313331"></a>

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

<a id="canonical-0110313113231130-3320002222210031-0012202203321021-0020311322133013-2001022032311201-0123102123030220-0000333200021112-2320221302033100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options

<a id="canonical-2112313023123321-1303310020003333-3230231010211230-2323230122300200-1121002213212023-1002203020131232-0022013310220111-2133211211231001"></a>

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

<a id="canonical-0311003021021322-3230310211232020-3201323321303212-3021312232010323-3303303133331132-0013212303223302-1020032031113020-3301021022320000"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options`

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-011.md#canonical-3211231110311030-2331001231120111-1231123022003220-0123210201002031-0330300120203103-1130101332331003-2220230122211000-3231303311233111): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-011.md#canonical-1312310123212123-0031211220330201-3012030211030300-3302203101230312-1221023131021102-3102131220011221-2300101121210333-3100013222300233): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-011.md#canonical-1301001112110020-3032312003103232-1233112133112121-2312232332132113-1102322220323331-2231103222132020-1333103113323010-3223310013102232): complete subsection reference.

<a id="canonical-3211231110311030-2331001231120111-1231123022003220-0123210201002031-0330300120203103-1130101332331003-2220230122211000-3231303311233111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-011.md#canonical-0110313113231130-3320002222210031-0012202203321021-0020311322133013-2001022032311201-0123102123030220-0000333200021112-2320221302033100)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-0001220120002321-3100133211303013-0131120310131202-3120221230031013-0102320122220111-0203010011330023-1112123120301323-1032330032023110"></a>

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

<a id="canonical-0013023003130111-2300020323312132-1102123112012223-0220002131222132-2210130233231203-1321200111002031-1311113000023201-2310212220120223"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](data-sources--workload--reference--group-011.md#canonical-0213232110302010-2201203320313001-1120231331133200-3002310001323321-3113232100322020-1221130003333323-3031230233220312-1031132032002003): complete subsection reference.

<a id="canonical-0213232110302010-2201203320313001-1120231331133200-3002310001323321-3113232100322020-1221130003333323-3031230233220312-1031132032002003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-011.md#canonical-0110313113231130-3320002222210031-0012202203321021-0020311322133013-2001022032311201-0123102123030220-0000333200021112-2320221302033100)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-011.md#canonical-3211231110311030-2331001231120111-1231123022003220-0123210201002031-0330300120203103-1130101332331003-2220230122211000-3231303311233111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-1231112303302133-2223100101112301-3103010123313310-1010221123311112-3223323113200102-3120111121023010-0202013233101030-1300133113020010"></a>

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

<a id="canonical-1122123203222332-3330100200101001-1212201312130122-1220132100320212-2233232231231103-2010300133231220-3120013100211120-3210331130300312"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](data-sources--workload--reference--group-011.md#canonical-1322123012301131-1320023130112123-3201120133313002-2223320302001023-1113232103220112-3300012323020012-2310121002210220-1221002113320011): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-011.md#canonical-2210120112001332-0002231321332313-3030200012213311-1013200223302321-0220331000032230-0322201210131323-1003121220332312-3330301320110030): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-011.md#canonical-0210211231213322-3220022103132333-2212210221202120-0031231200020010-0033330310233013-3312022012021232-1112122303233300-1010101333222200): complete subsection reference.

<a id="canonical-1322123012301131-1320023130112123-3201120133313002-2223320302001023-1113232103220112-3300012323020012-2310121002210220-1221002113320011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-011.md#canonical-0110313113231130-3320002222210031-0012202203321021-0020311322133013-2001022032311201-0123102123030220-0000333200021112-2320221302033100)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-011.md#canonical-3211231110311030-2331001231120111-1231123022003220-0123210201002031-0330300120203103-1130101332331003-2220230122211000-3231303311233111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-011.md#canonical-0213232110302010-2201203320313001-1120231331133200-3002310001323321-3113232100322020-1221130003333323-3031230233220312-1031132032002003)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-0320302323021022-2013121011302022-0130012331102232-0013012120130130-3002021121122232-0113303012003301-3321310032103013-1212312011332103"></a>

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

<a id="canonical-2210120112001332-0002231321332313-3030200012213311-1013200223302321-0220331000032230-0322201210131323-1003121220332312-3330301320110030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-011.md#canonical-0110313113231130-3320002222210031-0012202203321021-0020311322133013-2001022032311201-0123102123030220-0000333200021112-2320221302033100)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-011.md#canonical-3211231110311030-2331001231120111-1231123022003220-0123210201002031-0330300120203103-1130101332331003-2220230122211000-3231303311233111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-011.md#canonical-0213232110302010-2201203320313001-1120231331133200-3002310001323321-3113232100322020-1221130003333323-3031230233220312-1031132032002003)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-3221021320030103-1132300221031221-1322202312313320-0133101023331303-1203330131110110-2103200101310033-0212320002101212-2011033220003001"></a>

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

<a id="canonical-0210211231213322-3220022103132333-2212210221202120-0031231200020010-0033330310233013-3312022012021232-1112122303233300-1010101333222200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-011.md#canonical-0110313113231130-3320002222210031-0012202203321021-0020311322133013-2001022032311201-0123102123030220-0000333200021112-2320221302033100)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-011.md#canonical-3211231110311030-2331001231120111-1231123022003220-0123210201002031-0330300120203103-1130101332331003-2220230122211000-3231303311233111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-011.md#canonical-0213232110302010-2201203320313001-1120231331133200-3002310001323321-3113232100322020-1221130003333323-3031230233220312-1031132032002003)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-0331301332332312-3312213103130332-3322123300020201-1332022123323200-2202311202201031-3220300110031020-0321320130002203-3112101310321032"></a>

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

<a id="canonical-1312310123212123-0031211220330201-3012030211030300-3302203101230312-1221023131021102-3102131220011221-2300101121210333-3100013222300233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-011.md#canonical-0110313113231130-3320002222210031-0012202203321021-0020311322133013-2001022032311201-0123102123030220-0000333200021112-2320221302033100)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2113032201320311-0110022121010010-3131011110022133-2333233203121120-3212033130122102-1113010100113210-1120300021131120-0103333032222100"></a>

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

<a id="canonical-1301001112110020-3032312003103232-1233112133112121-2312232332132113-1102322220323331-2231103222132020-1333103113323010-3223310013102232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-011.md#canonical-0110313113231130-3320002222210031-0012202203321021-0020311322133013-2001022032311201-0123102123030220-0000333200021112-2320221302033100)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-1031032000132320-0102121020002301-3011133020103210-0330002103013230-2312123322031323-0133213131301103-0323202132002311-0222020310012022"></a>

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

<a id="canonical-1203232221031221-2133011011023111-1003123033102233-0321212010230202-1003301113220310-3112222211223301-2012211312132310-2202011212002320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-3200111023032221-3133332010313310-3111200321000332-0102101212312021-1333202331111030-0311210333003301-1003031203022033-3223032212023332"></a>

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

<a id="canonical-0323132102131121-2302300122020021-3312200023130311-2221102131322111-0221213321301300-1203112322222300-1001331303303200-1320112210021020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through

<a id="canonical-3132031230211132-3223202002113222-2210211020302021-3031331102001223-3213002323321310-0310330132320221-1103102232112131-0120202102031232"></a>

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

<a id="canonical-0200231313301230-1020120000303023-2310213213031220-3211310113302002-2013232220213000-3320133031232013-1021231020010121-0000222220020220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params

<a id="canonical-1200301020023112-1012333232112001-3023031221212213-2301303001032100-3301201321201113-2211220001032201-0200002212021130-3220231130103132"></a>

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

<a id="canonical-3130031320111222-2032212021200003-2010233202112113-3301113233222203-1310320230113212-2112213022133110-2300130222213131-1031332310133022"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params`

- [certificates](data-sources--workload--reference--group-011.md#canonical-3223220323301322-1123102312302301-1002101131100303-1231332130212133-2001223221333231-3212101320020333-3021232000111023-3102020110003123): complete subsection reference.

- [no_mtls](data-sources--workload--reference--group-011.md#canonical-3203112103001312-3200110030211300-1303213310310013-2020123211312003-0233330332121333-0203111231123131-3022200023221013-3212322210203210): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-011.md#canonical-0223321302121313-1211003020331101-0031211123212313-1020213230020013-3122020010011012-2030202203303333-2331003012303300-2322322110302313): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-012.md#canonical-1200133203102333-2120310032300312-2022020002031300-1003032200130313-1233303010202002-2210020001022213-2223231322331010-0101203013003021): complete subsection reference.

<a id="canonical-3223220323301322-1123102312302301-1002101131100303-1231332130212133-2001223221333231-3212101320020333-3021232000111023-3102020110003123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-011.md#canonical-0200231313301230-1020120000303023-2310213213031220-3211310113302002-2013232220213000-3320133031232013-1021231020010121-0000222220020220)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-0130003010200020-1230320222112022-3132302110220010-3210231223211312-1323023123200000-1312221011201021-1301130313232003-3032312032312102"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0213311313210113-3212021210202332-2310130330133111-0303321231102002-3221330333110313-1313323333330220-0223230121122321-3311011313200031"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates`

<a id="canonical-2331130200122332-0031212011230023-0110322211100213-2332020331132333-1020120103123120-0212323003230332-2133211020312110-0313032303110001"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0211210101123332-0112231120001301-0012121330101120-1200020032123112-2321232003300023-1312203122311033-1211232221333301-2121122302200203"></a>

<a id="canonical-0123221220222212-1133112230301320-0321002111200113-0201330111230010-2000030131311300-0300200222310302-1033111332123113-3202031330031310"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates.namespace` property

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

<a id="canonical-2033033223131101-2330131020022100-2200303211131102-0232102223000021-0232020303200210-3233023033031122-3112000022001110-2310310100330223"></a>

<a id="canonical-3113231203221300-0031330221200222-0023100330130203-2331330130313113-3010130222001303-1101301021200101-1202121202331302-3022200002113000"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3203112103001312-3200110030211300-1303213310310013-2020123211312003-0233330332121333-0203111231123131-3022200023221013-3212322210203210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-011.md#canonical-0200231313301230-1020120000303023-2310213213031220-3211310113302002-2013232220213000-3320133031232013-1021231020010121-0000222220020220)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-1001300201313201-2021123002213032-2321303021101322-1303121102300131-0322300002330211-2322331232020212-2020311303012220-1103121132013211"></a>

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

<a id="canonical-0223321302121313-1211003020331101-0031211123212313-1020213230020013-3122020010011012-2030202203303333-2331003012303300-2322322110302313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-011.md#canonical-0200231313301230-1020120000303023-2310213213031220-3211310113302002-2013232220213000-3320133031232013-1021231020010121-0000222220020220)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-2032230030132111-0012320002103222-2100202322331123-1202020123230332-2300300330233000-2222232132200100-3303220303120303-0213332023213311"></a>

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

<a id="canonical-0101230331313213-0212112010011130-3233033113313020-3232103313300130-3201033333223201-2233321231300222-1303303231032012-0023032321200002"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config`

- [custom_security](data-sources--workload--reference--group-011.md#canonical-1310302211113013-3200100223233311-0113313312221013-0102031301322303-2013222210012222-2122212331311230-0133232310310200-0020002101232133): complete subsection reference.

- [default_security](data-sources--workload--reference--group-012.md#canonical-0310003021212133-2232022023323231-0220300122102233-0220131213301323-1130220303332302-3331022122122203-1331113212203333-2303021231002312): complete subsection reference.

- [low_security](data-sources--workload--reference--group-012.md#canonical-2110311100331112-0321030323233302-0112133013123303-1131210031000131-0210223112121121-2202202312231230-2222320120303012-1033331300203323): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-012.md#canonical-3202201333202302-1021302212332222-3231031323110200-3212020311123023-3200033211033131-0010033322122113-1303211333311303-1320301012133332): complete subsection reference.

<a id="canonical-1310302211113013-3200100223233311-0113313312221013-0102031301322303-2013222210012222-2122212331311230-0133232310310200-0020002101232133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-011.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-011.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-011.md#canonical-1103203221301212-1112000122133022-0022122330012220-3331112112122231-2310223131213312-1303000213203301-0313313203133023-3311110220133122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-011.md#canonical-0200231313301230-1020120000303023-2310213213031220-3211310113302002-2013232220213000-3320133031232013-1021231020010121-0000222220020220)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-011.md#canonical-0223321302121313-1211003020331101-0031211123212313-1020213230020013-3122020010011012-2030202203303333-2331003012303300-2322322110302313)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-0300113103032001-2122012021322032-1113303020310020-3120312110202333-3022002320020031-1202131112121120-3133221200110112-0010122302232023"></a>

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
