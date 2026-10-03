---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1232020211032312-0303320033201230-0032121323122230-2233211133022332-3033110211102013-1113130222013012-2221133331023231-3301301020220031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121102132211101-2313303020302133-2230201311323133-2121222210132321-1010200022010021-0220100213323031-1231102013100032-2020230322213301"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response — route_direct_response / 230111122301 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-1030212223221121-2003313010132300-1020102130221221-1020323100331032-0232122121201212-1320322100030002-2100103213001013-3120331131112333)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-3111301021020303-0021100212023222-3230220100031310-3022311220211130-1233001333030033-3311110321112022-3232303003202302-1312332222222102)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-014.md#canonical-1023103121131031-3130012003123110-2300032111232121-2112312223010113-2322101332312303-1230133232332202-0010101333231222-2221211003012213)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

<a id="canonical-0200023210300012-1123220333300110-0123312023200221-0332230301221031-3113021211201322-1012312300323132-0023133332301101-2301203010002032"></a>

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

<a id="canonical-1220123232012311-2100323302200302-2131131201131232-3330210010122322-2333030123002023-1023000013131313-3003300232331320-1221123200011232"></a>

## Direct properties — route_direct_response / 230111122301 / 3

<a id="canonical-3301211222011111-2212220133023101-2212320313220332-3200112210230312-3130323020121310-3320303023233021-1300121223320100-1121112312013203"></a>

<a id="canonical-0210321313301100-2200030303310200-3001231320111222-0331112210213012-1023022323332110-2122023130220010-2331202131123212-0112003022213010"></a>

## response_body_encoded property — route_direct_response / 230111122301 / 4

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
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3200110201122023-2112330320130012-2333220313313032-2310211123013012-2012210113130221-1221332332312330-3302003101012210-1303110032111203"></a>

<a id="canonical-3001111311010003-1022202133323130-3123123223231103-1233203330132210-3111210033303313-0322332201032001-2030203002333002-0233013210210033"></a>

## response_code property — route_direct_response / 230111122301 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2013101210220022-0303012212301231-2123212321110010-3111332333103211-2112030133231122-2132120112302322-1111031322333220-0100122210133013"></a>

## Next pages — route_direct_response / 230111122301 / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-014.md#canonical-1023103121131031-3130012003123110-2300032111232121-2112312223010113-2322101332312303-1230133232332202-0010101333231222-2221211003012213)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1122120223023302-0122231111032311-2012130120211113-0220002121123103-1323223231322003-1313322120221001-1311133323202021-1333302222102121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313313112103121-1220012311112002-2212211231303312-3013111301312303-0101123111030101-3111312220001130-1011303130022102-2030312110222012"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route — redirect_route / 302103301320 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-1030212223221121-2003313010132300-1020102130221221-1020323100331032-0232122121201212-1320322100030002-2100103213001013-3120331131112333)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-3111301021020303-0021100212023222-3230220100031310-3022311220211130-1233001333030033-3311110321112022-3232303003202302-1312332222222102)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-1131101003100231-1101023200100121-3010030001120113-2310332112003331-0201303031121221-0113302131013300-0133222110223333-3023032101301311"></a>

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

<a id="canonical-2221312323213031-1222213130330302-0021202212032203-0003203113131323-2232223221201322-1320010003011021-3300202300322330-0120302201123213"></a>

## Direct properties — redirect_route / 302103301320 / 3

- [headers](data-sources--workload--reference--group-015.md#canonical-0003233122120000-0131230130100121-2201111101333330-2003001212220302-0333102221213012-2330122210312002-2200132020113001-3122033333100300): complete subsection reference.

<a id="canonical-2210332320213003-3120122220112003-2300122131000202-2113303330311101-3121331002030112-1231113000033332-3033030213200103-1122100323113132"></a>

<a id="canonical-2030023001102233-2132332123330203-2320320230100020-0000001302031213-0312212033113332-2211102103120003-0031032202203013-2023231313133333"></a>

## http_method property — redirect_route / 302103301320 / 4

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

- [incoming_port](data-sources--workload--reference--group-015.md#canonical-0220232021301133-1303133012210320-0323323222112210-0303323003313013-1031110231030322-3222121013232311-3312122300131312-3100132332211003): complete subsection reference.

- [path](data-sources--workload--reference--group-015.md#canonical-2003232300002311-0300100201302003-1120331231232232-1101302223212021-1033302230121303-3022002211332132-1310000133321220-0111100220312213): complete subsection reference.

- [route_redirect](data-sources--workload--reference--group-015.md#canonical-1112320121131220-0231312222012211-3302032310023212-3233030302323131-1101033331222010-0302223201200312-3030230100201022-1010313002202132): complete subsection reference.

<a id="canonical-0111022121301010-2032033203012213-2112202102101303-0011133111301112-0320332301002312-1112113212000121-1011321001312121-1232133003303212"></a>

## Next pages — redirect_route / 302103301320 / 5

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers](data-sources--workload--reference--group-015.md#canonical-0003233122120000-0131230130100121-2201111101333330-2003001212220302-0333102221213012-2330122210312002-2200132020113001-3122033333100300)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-015.md#canonical-0220232021301133-1303133012210320-0323323222112210-0303323003313013-1031110231030322-3222121013232311-3312122300131312-3100132332211003)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path](data-sources--workload--reference--group-015.md#canonical-2003232300002311-0300100201302003-1120331231232232-1101302223212021-1033302230121303-3022002211332132-1310000133321220-0111100220312213)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-015.md#canonical-1112320121131220-0231312222012211-3302032310023212-3233030302323131-1101033331222010-0302223201200312-3030230100201022-1010313002202132)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-3111301021020303-0021100212023222-3230220100031310-3022311220211130-1233001333030033-3311110321112022-3232303003202302-1312332222222102)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0003233122120000-0131230130100121-2201111101333330-2003001212220302-0333102221213012-2330122210312002-2200132020113001-3122033333100300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101123330003031-2113310303201210-3220330003023203-3210231010300322-1311230311323020-1011010223321010-0123322002333000-0220002132122232"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers — headers / 320120323113 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-1030212223221121-2003313010132300-1020102130221221-1020323100331032-0232122121201212-1320322100030002-2100103213001013-3120331131112333)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-3111301021020303-0021100212023222-3230220100031310-3022311220211130-1233001333030033-3311110321112022-3232303003202302-1312332222222102)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-015.md#canonical-1122120223023302-0122231111032311-2012130120211113-0220002121123103-1323223231322003-1313322120221001-1311133323202021-1333302222102121)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-2222220221123320-2013233133111032-2213102232021330-2131200213212232-0001331210302323-2100333111112303-3100012221213203-2302101031023132"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3211102330203130-3013110022012102-3123131220101233-3033301230302013-0122031011102300-2221102300232020-1112231013210130-2303302321113032"></a>

## Direct properties — headers / 320120323113 / 3

<a id="canonical-1010303120031301-3103333110202030-0312022122312301-3033212001000110-0031321002111333-1033011220330122-0323113001213112-0130231021222030"></a>

<a id="canonical-3320213331002333-0320000000213002-0233123012023203-1320223103221222-2320130333221112-1302321011012012-2030030201200322-3303010100131102"></a>

## exact property — headers / 320120323113 / 4

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-3203223110222212-3233332330312002-1032232211333112-3220232003133112-2312323323112020-2102100101100322-1203333030313132-2003322113133212"></a>

<a id="canonical-1233331031201322-1033012330210331-3113012302322022-2130103110100023-0320112232032120-3210333000101023-0230000200301231-0332023010300212"></a>

## invert_match property — headers / 320120323113 / 5

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

<a id="canonical-2331003320310011-0323011210111222-3121220112023201-2311130310030121-3201101021203220-3013301221312110-1122210122002302-2030023211031100"></a>

<a id="canonical-2003002001133131-2220230203332232-2003300333201231-0102222021111100-1201313030303123-1313333201133000-0332022011322210-1220001111332302"></a>

## name property — headers / 320120323113 / 6

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

<a id="canonical-3000311020320031-1332231232323111-2000113021031100-1102120032012202-1223001220033110-3330000112133223-3023202032232002-1121123312131101"></a>

<a id="canonical-1233012001203022-2131123230013323-0211202130021120-0022220201211220-1313333210130122-1100310102311122-0030320000300233-3320303001112332"></a>

## presence property — headers / 320120323113 / 7

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

<a id="canonical-0002313303300320-3203313033323121-2211113022110011-0211231030031020-1121020022312222-0300233002123022-0012220022112123-1120301033113101"></a>

<a id="canonical-0212303012300113-0022202022132333-3030312100223223-3132321310223333-0330310220310322-3130233200202302-1323101021013331-1202202212330320"></a>

## regular expression property — headers / 320120323113 / 8

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
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3121310210200132-1120321213330031-2100110013332110-2120330233010222-2201333100122211-3213320002232123-3310103130022100-2023202323013233"></a>

## Next pages — headers / 320120323113 / 9

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-015.md#canonical-1122120223023302-0122231111032311-2012130120211113-0220002121123103-1323223231322003-1313322120221001-1311133323202021-1333302222102121)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0220232021301133-1303133012210320-0323323222112210-0303323003313013-1031110231030322-3222121013232311-3312122300131312-3100132332211003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130320100232103-0023012213130003-3330111033020301-1200302113321201-2212033202021120-0123221020321112-3301132123120122-3022130210203203"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port — incoming_port / 123013010313 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-1030212223221121-2003313010132300-1020102130221221-1020323100331032-0232122121201212-1320322100030002-2100103213001013-3120331131112333)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-3111301021020303-0021100212023222-3230220100031310-3022311220211130-1233001333030033-3311110321112022-3232303003202302-1312332222222102)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-015.md#canonical-1122120223023302-0122231111032311-2012130120211113-0220002121123103-1323223231322003-1313322120221001-1311133323202021-1333302222102121)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port

<a id="canonical-2133101003231033-2311221210020223-2000033311121203-0310101122130231-0100131132301231-0331232010031220-1102033332103231-0032330231023220"></a>

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

<a id="canonical-1011133000022123-0321110323011001-1010022030132031-0220103302301311-2203331032322332-1312220331001102-0102213231011212-1100123302331300"></a>

## Direct properties — incoming_port / 123013010313 / 3

- [no_port_match](data-sources--workload--reference--group-015.md#canonical-2332301330310331-1012232133020010-1111330023023011-1003233300031022-3212022332010101-1020222232013203-2332300203333202-0000031333230302): complete subsection reference.

<a id="canonical-1012300210011021-0211203021002031-1112030011032303-3320333231233110-0123001332322113-2000211100222000-1213123131221021-0202120112022130"></a>

<a id="canonical-1123131211012020-0112023011031333-2213221200020122-1322233323021132-3213322230011021-1323330323111010-2330012131031132-1333320023231103"></a>

## port property — incoming_port / 123013010313 / 4

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

<a id="canonical-3122031012221221-1321021303332121-0000231132010032-0122211313033203-3201301032000021-0011101310313021-3130033132330100-0113002303231103"></a>

<a id="canonical-3000202330030123-3030121121101100-3223323100013310-3232300202232211-3202022220333331-3303222102030322-1123223020002232-1113333320111033"></a>

## port_ranges property — incoming_port / 123013010313 / 5

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

<a id="canonical-1201200333102230-0330122131111331-0103200330311233-0002013231303300-3013231201110021-1111202232223201-1023113322113010-0111311220001333"></a>

## Next pages — incoming_port / 123013010313 / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match](data-sources--workload--reference--group-015.md#canonical-2332301330310331-1012232133020010-1111330023023011-1003233300031022-3212022332010101-1020222232013203-2332300203333202-0000031333230302)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-015.md#canonical-1122120223023302-0122231111032311-2012130120211113-0220002121123103-1323223231322003-1313322120221001-1311133323202021-1333302222102121)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2332301330310331-1012232133020010-1111330023023011-1003233300031022-3212022332010101-1020222232013203-2332300203333202-0000031333230302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211011020032212-0312122211331010-0113320122310113-1110301221000132-3302013001201333-0021112231133310-0301101223330210-2332111030302101"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match — no_port_match / 110120130213 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-1030212223221121-2003313010132300-1020102130221221-1020323100331032-0232122121201212-1320322100030002-2100103213001013-3120331131112333)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-3111301021020303-0021100212023222-3230220100031310-3022311220211130-1233001333030033-3311110321112022-3232303003202302-1312332222222102)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-015.md#canonical-1122120223023302-0122231111032311-2012130120211113-0220002121123103-1323223231322003-1313322120221001-1311133323202021-1333302222102121)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-015.md#canonical-0220232021301133-1303133012210320-0323323222112210-0303323003313013-1031110231030322-3222121013232311-3312122300131312-3100132332211003)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-2020210003001110-3001333320213320-2233001132120133-1102120103110311-1311320003203123-3023301223101322-3122231223000022-3132203332002112"></a>

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

<a id="canonical-0112121202223102-1303122020211310-0133303120221013-0222113220101330-2312200032330220-2320032313121001-2231322323233200-2222222333213111"></a>

## Direct properties — no_port_match / 110120130213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231030002113123-0302300001231022-3131201331120113-2212120010132123-0232202111000023-2202233220003332-0133313010100002-3311030121202232"></a>

## Next pages — no_port_match / 110120130213 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-015.md#canonical-0220232021301133-1303133012210320-0323323222112210-0303323003313013-1031110231030322-3222121013232311-3312122300131312-3100132332211003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2003232300002311-0300100201302003-1120331231232232-1101302223212021-1033302230121303-3022002211332132-1310000133321220-0111100220312213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012210220312210-1021033133320321-0231010011112201-2113220113301310-3033303331313312-0001211233110313-3202233230312300-0310210020102113"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path — path / 101031103100 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-1030212223221121-2003313010132300-1020102130221221-1020323100331032-0232122121201212-1320322100030002-2100103213001013-3120331131112333)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-3111301021020303-0021100212023222-3230220100031310-3022311220211130-1233001333030033-3311110321112022-3232303003202302-1312332222222102)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-015.md#canonical-1122120223023302-0122231111032311-2012130120211113-0220002121123103-1323223231322003-1313322120221001-1311133323202021-1333302222102121)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path

<a id="canonical-1130100003301222-0101313111321112-0022131232003332-0132311101132321-2202233130233213-3113000103022201-0010321003002201-3311233022330112"></a>

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

<a id="canonical-2003121220303120-0231322122321220-1131310223123020-1023001120223330-3233332210110320-0300200113130212-1312133111221233-3201132330222122"></a>

## Direct properties — path / 101031103100 / 3

<a id="canonical-2233310210231132-0332133323102121-0110131021200113-2321212010310003-2213320200311320-2033211020030233-2102321233302230-0013230301132131"></a>

<a id="canonical-2102302103203233-2102330333021112-2221103032202232-2002332023031221-0303231331123323-3332033223101300-0302231231122300-0111333032001102"></a>

## path property — path / 101031103100 / 4

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

<a id="canonical-1303030022212300-3222300021021211-2332220302310133-1032011212011132-3103032220203321-1331121001010030-1032232232013121-1233123002201030"></a>

<a id="canonical-1320110311221021-2012003112310322-3333232031220023-1032030030330302-3030121233200032-3031323321003203-3330030330202132-3300220331122320"></a>

## prefix property — path / 101031103100 / 5

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

<a id="canonical-3031222021020210-0203022313121212-2313232131203230-2110000003103321-0012132232322021-3220211031000331-0310202000323112-0221112123202203"></a>

<a id="canonical-3332331112133032-1320021132010001-2012211003012321-2313100330212303-2220011112002112-1303003330213320-3130000320323303-1303212223022122"></a>

## regular expression property — path / 101031103100 / 6

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

<a id="canonical-2101320021101101-1112113012130311-3113102310303002-2000030313033032-2132113301020210-3111312123302200-1000320131001202-0311102030311131"></a>

## Next pages — path / 101031103100 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-015.md#canonical-1122120223023302-0122231111032311-2012130120211113-0220002121123103-1323223231322003-1313322120221001-1311133323202021-1333302222102121)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1112320121131220-0231312222012211-3302032310023212-3233030302323131-1101033331222010-0302223201200312-3030230100201022-1010313002202132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312110122101022-0023023012201303-1032313310200210-1211332302311010-0200112102303210-2331031013323301-3022300200202100-0023031132320311"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect — route_redirect / 100000333330 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-1030212223221121-2003313010132300-1020102130221221-1020323100331032-0232122121201212-1320322100030002-2100103213001013-3120331131112333)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-3111301021020303-0021100212023222-3230220100031310-3022311220211130-1233001333030033-3311110321112022-3232303003202302-1312332222222102)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-015.md#canonical-1122120223023302-0122231111032311-2012130120211113-0220002121123103-1323223231322003-1313322120221001-1311133323202021-1333302222102121)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect

<a id="canonical-3133023101313101-1000103021232121-0102130313031032-0221123200211002-0333220002232322-3223000323121232-2120011111223032-3011032033213210"></a>

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

<a id="canonical-3113322131002321-3213031003222020-1213321020200102-0103311001323131-3203013221130031-1312002021021330-1213213100031002-0201200120221123"></a>

## Direct properties — route_redirect / 100000333330 / 3

<a id="canonical-0031202310200232-1133101230130332-1230132212221200-2003000301323322-1000101002010301-0100112130003000-0333131000022210-1330132123311112"></a>

<a id="canonical-3222111100123323-3301100030321023-3023311323120321-1202010312033223-2133313113310323-1021133002131332-2132322020202233-1333201223331102"></a>

## host_redirect property — route_redirect / 100000333330 / 4

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

<a id="canonical-1123231233230222-0133223031012203-3200013121003212-2132312012010312-1001323113032010-0133210211221310-1110232101220211-0022032332131011"></a>

<a id="canonical-1010223302033311-1133011221000233-1212210003030011-2312200303032311-2023011110320202-1201231201311212-1323310113101132-0013002201212223"></a>

## path_redirect property — route_redirect / 100000333330 / 5

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

<a id="canonical-1233322122202210-3312113212331333-1111300220121311-1310022300202213-1020122231300000-0231031012103133-3013313002223333-2033132002213312"></a>

<a id="canonical-3232312233101123-2122132010021203-2123220002333021-3312020122222101-3301210103203120-2222120002200211-1030110302003130-2012123333023131"></a>

## prefix_rewrite property — route_redirect / 100000333330 / 6

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

<a id="canonical-0132313013311322-0120110220033103-1221122322310311-2132201232031020-0220101023030011-0313200230113222-2030001101103120-1010203203010110"></a>

<a id="canonical-3033222313313203-3123012032003003-0120033221303310-3110302330201101-3321133110002033-1302011301003302-0130223110302200-3013220332101332"></a>

## proto_redirect property — route_redirect / 100000333330 / 7

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
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  }
}
```

- [remove_all_params](data-sources--workload--reference--group-015.md#canonical-2202313330012201-2123031231231202-1103220220210130-2002013233003221-3011001330023113-2131322233033101-2011312113211101-2222101331220330): complete subsection reference.

<a id="canonical-2200012221100102-2331133210013023-1023331113103103-2323100000300222-3023333233102213-3100032013012211-1113301131021231-1210021222111032"></a>

<a id="canonical-3110121233003111-2001301232103103-1011202321231330-2122330030220103-0302223333003311-0201322120120003-2323223101310210-0113011333112022"></a>

## replace_params property — route_redirect / 100000333330 / 8

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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0130100333033100-2131222212021222-1321322310221202-2222200202332302-1023003333311312-0232000113331223-3310022303200022-3123032321301002"></a>

<a id="canonical-0001011302010332-1232021223020113-0311230230030312-1211123220111020-1031312103121032-2010103203033333-1120331322121131-3010111322312120"></a>

## response_code property — route_redirect / 100000333330 / 9

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
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

- [retain_all_params](data-sources--workload--reference--group-015.md#canonical-0102121313003231-3321303133202210-0333010112033023-2222333332001001-0033302110011323-2231021321301003-1202221323020133-3133133012100202): complete subsection reference.

<a id="canonical-2330010203210222-3110121101203233-0032112003021301-3301210033332011-1320100120333331-1300131110132322-1131022213033110-2211321131220122"></a>

## Next pages — route_redirect / 100000333330 / 10

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params](data-sources--workload--reference--group-015.md#canonical-2202313330012201-2123031231231202-1103220220210130-2002013233003221-3011001330023113-2131322233033101-2011312113211101-2222101331220330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params](data-sources--workload--reference--group-015.md#canonical-0102121313003231-3321303133202210-0333010112033023-2222333332001001-0033302110011323-2231021321301003-1202221323020133-3133133012100202)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-015.md#canonical-1122120223023302-0122231111032311-2012130120211113-0220002121123103-1323223231322003-1313322120221001-1311133323202021-1333302222102121)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2202313330012201-2123031231231202-1103220220210130-2002013233003221-3011001330023113-2131322233033101-2011312113211101-2222101331220330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123123012013333-2200133002002302-1301310020001103-3330111330300031-1031100112331113-2312020213033313-2301031030123321-1011032002013332"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params — remove_all_params / 130032210223 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-1030212223221121-2003313010132300-1020102130221221-1020323100331032-0232122121201212-1320322100030002-2100103213001013-3120331131112333)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-3111301021020303-0021100212023222-3230220100031310-3022311220211130-1233001333030033-3311110321112022-3232303003202302-1312332222222102)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-015.md#canonical-1122120223023302-0122231111032311-2012130120211113-0220002121123103-1323223231322003-1313322120221001-1311133323202021-1333302222102121)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-015.md#canonical-1112320121131220-0231312222012211-3302032310023212-3233030302323131-1101033331222010-0302223201200312-3030230100201022-1010313002202132)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-3222311303232132-2201123310321310-2223131130202101-0002301122000001-0000303310012322-2230122022111011-2311133302122332-2113010221030110"></a>

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

<a id="canonical-1023001023310321-3133231011333221-2130130021121112-3010201300013102-2322132003133031-2222331102111311-2213303123021202-2301123131031221"></a>

## Direct properties — remove_all_params / 130032210223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3200221222221032-0221223201333120-0313133312302223-1112021201100312-2003203322201100-0320333220123032-2023102003111133-2233220111022220"></a>

## Next pages — remove_all_params / 130032210223 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-015.md#canonical-1112320121131220-0231312222012211-3302032310023212-3233030302323131-1101033331222010-0302223201200312-3030230100201022-1010313002202132)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0102121313003231-3321303133202210-0333010112033023-2222333332001001-0033302110011323-2231021321301003-1202221323020133-3133133012100202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033323121302201-2112020121303002-0033232020021230-0220223201200310-3311202223230010-3122030000110103-2020202020212123-3003211122013002"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params — retain_all_params / 103331211312 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-1030212223221121-2003313010132300-1020102130221221-1020323100331032-0232122121201212-1320322100030002-2100103213001013-3120331131112333)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-3111301021020303-0021100212023222-3230220100031310-3022311220211130-1233001333030033-3311110321112022-3232303003202302-1312332222222102)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-015.md#canonical-1122120223023302-0122231111032311-2012130120211113-0220002121123103-1323223231322003-1313322120221001-1311133323202021-1333302222102121)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-015.md#canonical-1112320121131220-0231312222012211-3302032310023212-3233030302323131-1101033331222010-0302223201200312-3030230100201022-1010313002202132)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-2330201222023023-2012121203220300-0123333121103031-2210201302213032-3201320200000300-0132302230223121-0132312003121223-0132130211130203"></a>

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

<a id="canonical-2102102021032310-2012201310333010-1102310312112323-3332120220112303-0113123301121023-1102000303100202-3201330210103222-1202332202002001"></a>

## Direct properties — retain_all_params / 103331211312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113011030120123-0321201302103123-3020201123311023-0320123101130121-3123202313222222-2231023221030320-0122330013002013-2233300331002332"></a>

## Next pages — retain_all_params / 103331211312 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-015.md#canonical-1112320121131220-0231312222012211-3302032310023212-3233030302323131-1101033331222010-0302223201200312-3030230100201022-1010313002202132)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3233013323003331-2033101300323011-1300102212130223-0223123003201121-2100112211231333-2332133233020332-3023220212202022-2232220233313322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200011223122012-3201031121030323-1330032203330100-2201021100030203-2213213220223002-0101211113110033-2301331333323020-2202010011131132"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route — simple_route / 212030120001 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-1030212223221121-2003313010132300-1020102130221221-1020323100331032-0232122121201212-1320322100030002-2100103213001013-3120331131112333)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-3111301021020303-0021100212023222-3230220100031310-3022311220211130-1233001333030033-3311110321112022-3232303003202302-1312332222222102)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-0230220232221313-2211011333333023-3332031103311121-3300022023200222-0230003332300202-1221032311200300-2130033103320220-2023331233300001"></a>

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

<a id="canonical-1323212213130131-2122123033121121-3301220010311300-1210120011311031-1311023203300221-1112030313111201-2021321313021330-3121112010310203"></a>

## Direct properties — simple_route / 212030120001 / 3

- [auto_host_rewrite](data-sources--workload--reference--group-015.md#canonical-2023101232332030-3121001212202001-3103001313302031-1211023300000013-0320303213201002-1001012023121121-0020021100003301-3112121213103011): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-015.md#canonical-3332020010202213-2221113303131121-1221100112023200-3311332233201300-1311023110022123-0200020010313122-0013233102302200-2112331112032100): complete subsection reference.

<a id="canonical-0010020013110031-1002103000231221-1021233013220313-2023020330112302-0212012031023331-2112113101112203-3122120301203011-3213201023230232"></a>

<a id="canonical-0023320012331311-2123221120312311-2202201301213121-3213310312330001-3010221213121023-3032201222011313-0212121230132333-2013012230010101"></a>

## host_rewrite property — simple_route / 212030120001 / 4

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

<a id="canonical-3331122301232303-2203121010300210-1020032103120233-3111232130003133-1012113323211132-3311310202110313-0000302322221030-3010030211112333"></a>

<a id="canonical-0202022310321103-1103120333231010-2212110223302230-1300003330020302-2100322100230000-2200030333211122-2233300200130100-3011022332102001"></a>

## http_method property — simple_route / 212030120001 / 5

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

- [path](data-sources--workload--reference--group-015.md#canonical-2312121333133302-2133220023010323-0012113330323002-2012110133021023-3233210313332223-3211033010221333-3030010103333102-0031203220332020): complete subsection reference.

<a id="canonical-2112223002101203-1220231031113131-3022230323222222-3123313112120231-0011030013221232-1101002220111321-2032100222213020-1203313123021322"></a>

## Next pages — simple_route / 212030120001 / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite](data-sources--workload--reference--group-015.md#canonical-2023101232332030-3121001212202001-3103001313302031-1211023300000013-0320303213201002-1001012023121121-0020021100003301-3112121213103011)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite](data-sources--workload--reference--group-015.md#canonical-3332020010202213-2221113303131121-1221100112023200-3311332233201300-1311023110022123-0200020010313122-0013233102302200-2112331112032100)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path](data-sources--workload--reference--group-015.md#canonical-2312121333133302-2133220023010323-0012113330323002-2012110133021023-3233210313332223-3211033010221333-3030010103333102-0031203220332020)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-3111301021020303-0021100212023222-3230220100031310-3022311220211130-1233001333030033-3311110321112022-3232303003202302-1312332222222102)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2023101232332030-3121001212202001-3103001313302031-1211023300000013-0320303213201002-1001012023121121-0020021100003301-3112121213103011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130223303330313-1300130131110120-1131122310000303-2311112232133120-1220312323200112-2011212110312212-1023320302303331-2032313011210302"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite — auto_host_rewrite / 331133320323 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-1030212223221121-2003313010132300-1020102130221221-1020323100331032-0232122121201212-1320322100030002-2100103213001013-3120331131112333)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-3111301021020303-0021100212023222-3230220100031310-3022311220211130-1233001333030033-3311110321112022-3232303003202302-1312332222222102)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-015.md#canonical-3233013323003331-2033101300323011-1300102212130223-0223123003201121-2100112211231333-2332133233020332-3023220212202022-2232220233313322)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-1131310222021002-2223022100010112-2200331320122300-2020330110013030-1233210103312121-0020113212330232-2323002012011303-0101203101023102"></a>

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

<a id="canonical-1021230300012233-1103232321320112-2123112000023133-2102021001222113-1331110223113131-2112220211111322-1212000211221311-3323222323121132"></a>

## Direct properties — auto_host_rewrite / 331133320323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222121300330032-1313010202331012-1303312310002030-0200103302031011-0110212002011210-0110310031222023-2101020213010131-3223000222211133"></a>

## Next pages — auto_host_rewrite / 331133320323 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-015.md#canonical-3233013323003331-2033101300323011-1300102212130223-0223123003201121-2100112211231333-2332133233020332-3023220212202022-2232220233313322)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3332020010202213-2221113303131121-1221100112023200-3311332233201300-1311023110022123-0200020010313122-0013233102302200-2112331112032100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030111332230332-0231303030222310-0001232310013232-2322322223112232-3212323211233230-0012332022001021-2202013320223220-1311121213012220"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite — disable_host_rewrite / 120311121102 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-1030212223221121-2003313010132300-1020102130221221-1020323100331032-0232122121201212-1320322100030002-2100103213001013-3120331131112333)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-3111301021020303-0021100212023222-3230220100031310-3022311220211130-1233001333030033-3311110321112022-3232303003202302-1312332222222102)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-015.md#canonical-3233013323003331-2033101300323011-1300102212130223-0223123003201121-2100112211231333-2332133233020332-3023220212202022-2232220233313322)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-0200021201132330-3213101330132201-1332000020203231-2110222213201313-2022130322203020-0300031121330003-0200120330321122-1022201202112301"></a>

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

<a id="canonical-3331122220020003-1112302033013011-3213321220212133-1132030022232313-2220211233120011-2011131303213123-0022110303333320-3323003230213133"></a>

## Direct properties — disable_host_rewrite / 120311121102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012103003310133-0232123202322111-3303203222212132-2313201301310200-2122230212003002-2332203111322023-2133311201333111-1200311203121112"></a>

## Next pages — disable_host_rewrite / 120311121102 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-015.md#canonical-3233013323003331-2033101300323011-1300102212130223-0223123003201121-2100112211231333-2332133233020332-3023220212202022-2232220233313322)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2312121333133302-2133220023010323-0012113330323002-2012110133021023-3233210313332223-3211033010221333-3030010103333102-0031203220332020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313112322030000-3023312320323103-3311102111231121-2331020320111122-2330110231112323-3020111233221102-1020100030131031-1322001222012210"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path — path / 100101332233 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-012.md#canonical-1201010131131102-0122232302301201-1311103211301323-0002022102012002-1210302013111313-0223303123232013-2020000301211213-1320023310023200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-014.md#canonical-1030212223221121-2003313010132300-1020102130221221-1020323100331032-0232122121201212-1320322100030002-2100103213001013-3120331131112333)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-014.md#canonical-3111301021020303-0021100212023222-3230220100031310-3022311220211130-1233001333030033-3311110321112022-3232303003202302-1312332222222102)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-015.md#canonical-3233013323003331-2033101300323011-1300102212130223-0223123003201121-2100112211231333-2332133233020332-3023220212202022-2232220233313322)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-1200100123130030-2211220323033000-0013011201000210-2021302323020021-0301232231021013-3322210211213133-0321110032332133-1223022230131111"></a>

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

<a id="canonical-1023032330102302-3302202122303232-0130101102331222-1020003012032132-0303102012013133-0331302032211221-0222322013123230-2211311300230102"></a>

## Direct properties — path / 100101332233 / 3

<a id="canonical-0121222022332022-1023302023011011-0112213212311001-3101102101120023-2132203203102021-1031113021212100-2033000312322303-0100000210102311"></a>

<a id="canonical-3332020231331231-0231122310013002-2030323311130132-0210330311120011-3233112213231223-2332111300321000-2230201031021111-1013101330121023"></a>

## path property — path / 100101332233 / 4

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

<a id="canonical-1003121223110012-3301311322120320-0012313230031221-0030322010120000-2223121310221022-2230330131032002-2331031312110231-2103112330322112"></a>

<a id="canonical-1010301020122221-3300121310232133-0101300231300013-1100312111321123-2003011011202002-3113131310031132-0212001002213202-0320010321012101"></a>

## prefix property — path / 100101332233 / 5

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

<a id="canonical-0310112210120201-1102201202330023-1022023202230012-2031233311011100-3111303023000313-1002003203112331-1213022320030322-1000002023100023"></a>

<a id="canonical-0333212101311212-2120020232323003-1103212101220202-1333111131130011-3031221032233331-3000113222110002-3230121011231021-0101102330122102"></a>

## regular expression property — path / 100101332233 / 6

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

<a id="canonical-0000221103321130-1002323223010021-1210010302121201-0332223021103233-3011220230202332-1301312132210300-2302031210133232-1311011323123322"></a>

## Next pages — path / 100101332233 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-015.md#canonical-3233013323003331-2033101300323011-1300102212130223-0223123003201121-2100112211231333-2332133233020332-3023220212202022-2232220233313322)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0113001100231030-3133302301102130-2211201121032213-2012110310122033-1330211302020133-2012113020123300-0302331301232310-3311330330300223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122303102301120-2032033223302310-2120311103030312-3333311012102320-0110122312001223-3020020113322033-0200333031231330-0233033333111123"></a>

## service.advertise_options.advertise_on_public.port.port — port / 020210131213 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- service.advertise_options.advertise_on_public.port.port

<a id="canonical-3312132031100321-2312232301103211-2001232002213103-1323002213211221-1322210030100210-3002211330301101-0112103322223220-2001001302321232"></a>

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

<a id="canonical-1000013123021321-0033202022002012-0011121212033103-0111320020030320-1232323020030021-2122133331312122-1131213311130001-0302123033123122"></a>

## Direct properties — port / 020210131213 / 3

- [info](data-sources--workload--reference--group-015.md#canonical-3000322223020310-1220010213122311-2312132133323012-2201202021301210-2121301011230211-1303223020101202-1201111320201102-0003222003130010): complete subsection reference.

<a id="canonical-0223013120003232-2233133113111031-2222130021023233-3020032031110232-3300031330130123-1321121100110210-0203331010003212-3021111021233202"></a>

## Next pages — port / 020210131213 / 4

- [service.advertise_options.advertise_on_public.port.port.info](data-sources--workload--reference--group-015.md#canonical-3000322223020310-1220010213122311-2312132133323012-2201202021301210-2121301011230211-1303223020101202-1201111320201102-0003222003130010)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3000322223020310-1220010213122311-2312132133323012-2201202021301210-2121301011230211-1303223020101202-1201111320201102-0003222003130010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300313300211003-0301020323002212-2112110101130100-0221112020101020-0013313303320333-0111323212212322-0112230233032303-1032300002122303"></a>

## service.advertise_options.advertise_on_public.port.port.info — info / 300203000101 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.port](data-sources--workload--reference--group-015.md#canonical-0113001100231030-3133302301102130-2211201121032213-2012110310122033-1330211302020133-2012113020123300-0302331301232310-3311330330300223)
- service.advertise_options.advertise_on_public.port.port.info

<a id="canonical-2230303130332000-3203302003123321-3233211320222230-0310330212131102-1113300230030323-2311233122010321-0022301232021101-2301133321002232"></a>

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

<a id="canonical-0101110300112001-2210220011132331-0202330210201030-3031231030312020-0220020002320121-3103231202223300-3033033103323011-3122320222112230"></a>

## Direct properties — info / 300203000101 / 3

<a id="canonical-3212221031313322-2313023030111333-3000011211221033-2321333122033003-3212012023310221-3033323202232330-3000322021012100-2011133231031013"></a>

<a id="canonical-1011011001220103-1011100203331333-0100201101020132-3210322130001013-3003231002110010-0320311002202230-1321231011211022-3220303112023021"></a>

## port property — info / 300203000101 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1300131231323211-1202332203030032-1200001030130333-1232130200013321-2102223333030323-1220310023011031-3123203131230321-2313202101201232"></a>

<a id="canonical-0230132303010200-3233213302200100-1023212023023012-3120013333302021-0013322303100232-3033231013123013-3033213001101303-3012101322102121"></a>

## protocol property — info / 300203000101 / 5

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

- [same_as_port](data-sources--workload--reference--group-015.md#canonical-0111333033032000-0132013301102111-0323100301302113-3102032323002001-0212022300031331-1212113333032301-2003223000132001-3111300120213201): complete subsection reference.

<a id="canonical-0311221212123223-1003200110321211-3220020021301132-0010023211333322-3300220020032211-1122111231301002-0302110220323131-0322103221001223"></a>

<a id="canonical-1331013101331133-0330011220202111-2013330211001200-1300222103313010-0322331121021331-1323321132031111-1011202303030311-1220132111003300"></a>

## target_port property — info / 300203000101 / 6

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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1201322221230113-2112221033223002-3211120122103322-1331030030310311-3330032312322013-0013012200000311-1012011100003002-2131111110131321"></a>

## Next pages — info / 300203000101 / 7

- [service.advertise_options.advertise_on_public.port.port.info.same_as_port](data-sources--workload--reference--group-015.md#canonical-0111333033032000-0132013301102111-0323100301302113-3102032323002001-0212022300031331-1212113333032301-2003223000132001-3111300120213201)
- [service.advertise_options.advertise_on_public.port.port](data-sources--workload--reference--group-015.md#canonical-0113001100231030-3133302301102130-2211201121032213-2012110310122033-1330211302020133-2012113020123300-0302331301232310-3311330330300223)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0111333033032000-0132013301102111-0323100301302113-3102032323002001-0212022300031331-1212113333032301-2003223000132001-3111300120213201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113113131331110-1011230122221032-3021030001000232-1302020210031331-1321122110021301-1023312102110212-3303002311300031-0300331131010322"></a>

## service.advertise_options.advertise_on_public.port.port.info.same_as_port — same_as_port / 302112002020 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [service.advertise_options.advertise_on_public.port.port](data-sources--workload--reference--group-015.md#canonical-0113001100231030-3133302301102130-2211201121032213-2012110310122033-1330211302020133-2012113020123300-0302331301232310-3311330330300223)
- [service.advertise_options.advertise_on_public.port.port.info](data-sources--workload--reference--group-015.md#canonical-3000322223020310-1220010213122311-2312132133323012-2201202021301210-2121301011230211-1303223020101202-1201111320201102-0003222003130010)
- service.advertise_options.advertise_on_public.port.port.info.same_as_port

<a id="canonical-1232110030233130-0211033200131011-0103323120231203-3301232120121303-0233033131322301-3200103022300111-1013021012320320-1232320322103133"></a>

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

<a id="canonical-3002021200202001-0001022233223321-0323130301301111-3111130221232203-0332012213013103-0021201222113321-0312303001002232-2222201111120010"></a>

## Direct properties — same_as_port / 302112002020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021132120102111-1122010200320103-2302031302321203-2123133301010233-1110023213331323-3200231202232301-3023022001022131-3021330112321022"></a>

## Next pages — same_as_port / 302112002020 / 4

- [service.advertise_options.advertise_on_public.port.port.info](data-sources--workload--reference--group-015.md#canonical-3000322223020310-1220010213122311-2312132133323012-2201202021301210-2121301011230211-1303223020101202-1201111320201102-0003222003130010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0032120222121011-0103113003232030-3033201310032302-2121320210222122-2220130011002330-2032233123020302-1011100012101223-1313103122322332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121102223203013-1131221312130020-2100223113003001-1002031200300010-2023221221212301-2303222103130333-1302031212331210-0232223002012220"></a>

## service.advertise_options.advertise_on_public.port.tcp_loadbalancer — tcp_loadbalancer / 323222322320 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- service.advertise_options.advertise_on_public.port.tcp_loadbalancer

<a id="canonical-2132202023103102-0101213011202031-1233322313231311-0022123221100231-1212322121123011-3103132121100311-3310101221310220-1203110012013213"></a>

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

<a id="canonical-0312303332030131-3233323221331202-3113213313122123-0310012310222020-3310221332132030-3032130021330122-1010233112113220-0030210331122133"></a>

## Direct properties — tcp_loadbalancer / 323222322320 / 3

<a id="canonical-3101112103131010-1322031233121220-0201101000221110-3100112312110002-1022001113201221-0113201313113130-3102020012120312-0102213020332120"></a>

<a id="canonical-3203132030212132-1032001133011332-0031211130113332-1121131122333222-3130000312123020-3321310112322301-0213021000203300-2230100331332021"></a>

## domains property — tcp_loadbalancer / 323222322320 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2021222222033002-2232331302301121-2130031210033303-3323003211111030-3312032211311010-0001013222010200-3332231312320013-1220301303120221"></a>

<a id="canonical-0020220123000100-2211002132203231-2323233200033323-1102302012321031-1000001210213220-3032123300323230-0032023212031101-3331321303310210"></a>

## with_sni property — tcp_loadbalancer / 323222322320 / 5

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

<a id="canonical-0231323102002130-2132132003100132-2232101122312231-0103311320200111-0323030333102131-1003113213030220-3030033311321121-3133302012300201"></a>

## Next pages — tcp_loadbalancer / 323222322320 / 6

- [service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-012.md#canonical-2201022333000202-1302130200321220-3213032300002112-2121010232101103-0121112032111122-3003301203201212-0020213301103033-1313103313220332)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3303332030303022-2112310112031201-1232033023113213-2310022310310212-0130333030130000-2200322102123323-3201131322011030-3120330230033000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000311110313323-2030302300331200-3220021122103132-1023001111032202-2130222332201203-3300103111231221-1031022333303131-1031100223213103"></a>

## service.advertise_options.do_not_advertise — do_not_advertise / 222230200323 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- service.advertise_options.do_not_advertise

<a id="canonical-1101011012113320-2333201322230110-3021202022120211-1012200202033200-1211133131301102-2321030012032031-0033303033031300-0230311123002221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for do not advertise.

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

<a id="canonical-0202002232012132-3331232332131222-3223132001330111-1232211333300130-0313223300231113-2132021322332033-3200322011033112-2010012110110210"></a>

## Direct properties — do_not_advertise / 222230200323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302120332230313-3201201321330313-0322313011230120-2111001221211303-1011203212102302-3021303302321300-1023311001102321-1222121102011231"></a>

## Next pages — do_not_advertise / 222230200323 / 4

- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3123002002331230-0201233111020120-1333221333320003-3110223112030011-2003333010313113-2333211001212031-3333112223203131-0322131002310210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131132121320332-2313231202133210-1020101300300201-3021230221032010-3112113210221221-3231111300012212-0100111212310332-3110221133233021"></a>

## service.configuration — configuration / 133211030300 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- service.configuration

<a id="canonical-1333212211000030-2330301321131213-3223012322301032-2200003230222211-1110100300013203-3202230113111213-2003303201111322-1320111210231031"></a>

Type: `"single"`. Computed.

Configuration parameters of the workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2303203301103203-0301100132301201-2000130003231011-1232223121232122-2200302033301000-1031012033222313-3310213022031320-1310220013021033"></a>

## Direct properties — configuration / 133211030300 / 3

- [parameters](data-sources--workload--reference--group-015.md#canonical-0012011320103113-3033302112011313-2110032033231102-1120002333001312-0230133122323022-2322200011110000-1310321323010332-0020233203221233): complete subsection reference.

<a id="canonical-2101320333211000-2100121031331102-3202203302103032-2132203220221313-1213302033201031-3022022312331231-0210103000100312-1130001321320103"></a>

## Next pages — configuration / 133211030300 / 4

- [service.configuration.parameters](data-sources--workload--reference--group-015.md#canonical-0012011320103113-3033302112011313-2110032033231102-1120002333001312-0230133122323022-2322200011110000-1310321323010332-0020233203221233)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0012011320103113-3033302112011313-2110032033231102-1120002333001312-0230133122323022-2322200011110000-1310321323010332-0020233203221233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220121122210223-0112212303203010-0022331031030202-2122121102132212-2332312100232033-0010112121333100-1101322010202132-0103112132223230"></a>

## service.configuration.parameters — parameters / 112002301311 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.configuration](data-sources--workload--reference--group-015.md#canonical-3123002002331230-0201233111020120-1333221333320003-3110223112030011-2003333010313113-2333211001212031-3333112223203131-0322131002310210)
- service.configuration.parameters

<a id="canonical-1123303301120122-1223030313001030-3211023001012100-0202113332031113-1233011301213233-2202020113231112-3322303301022201-3200232031331231"></a>

Type: `"list"`. Computed.

Parameters. Parameters for the workload.

Upstream description:

Parameters for the workload.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2323101210022031-0231130133100031-0122021120122111-1131331230112101-2313313031033110-3013100330033002-1010020013013021-0322333020121023"></a>

## Direct properties — parameters / 112002301311 / 3

- [env_var](data-sources--workload--reference--group-015.md#canonical-0031002112012200-0030032032102202-3310321000323231-3122003333310023-1033031130310330-2332012312232220-1003130213031222-2112112201323313): complete subsection reference.

- [file](data-sources--workload--reference--group-015.md#canonical-3332330103203231-2232213211320013-3031002131033012-0231300111321033-0200101300032231-0102031133310322-1202321212212011-3303003200302311): complete subsection reference.

<a id="canonical-0200032111023313-0100033220032121-1010111122300302-2112121022021022-2010012212133112-0002110322102103-2231122013122130-3331122002003121"></a>

## Next pages — parameters / 112002301311 / 4

- [service.configuration.parameters.env_var](data-sources--workload--reference--group-015.md#canonical-0031002112012200-0030032032102202-3310321000323231-3122003333310023-1033031130310330-2332012312232220-1003130213031222-2112112201323313)
- [service.configuration.parameters.file](data-sources--workload--reference--group-015.md#canonical-3332330103203231-2232213211320013-3031002131033012-0231300111321033-0200101300032231-0102031133310322-1202321212212011-3303003200302311)
- [service.configuration](data-sources--workload--reference--group-015.md#canonical-3123002002331230-0201233111020120-1333221333320003-3110223112030011-2003333010313113-2333211001212031-3333112223203131-0322131002310210)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0031002112012200-0030032032102202-3310321000323231-3122003333310023-1033031130310330-2332012312232220-1003130213031222-2112112201323313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112000112330132-3020003303100132-0131211233200311-0012300302222013-0213132120120311-3131201310112301-2302111210310122-2322212011011211"></a>

## service.configuration.parameters.env_var — env_var / 130001320032 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.configuration](data-sources--workload--reference--group-015.md#canonical-3123002002331230-0201233111020120-1333221333320003-3110223112030011-2003333010313113-2333211001212031-3333112223203131-0322131002310210)
- [service.configuration.parameters](data-sources--workload--reference--group-015.md#canonical-0012011320103113-3033302112011313-2110032033231102-1120002333001312-0230133122323022-2322200011110000-1310321323010332-0020233203221233)
- service.configuration.parameters.env_var

<a id="canonical-0033320211222312-0313221133020122-1031303102003221-2332313220100100-1330233013110021-1233110121310021-0331012130221223-0100023003203130"></a>

Type: `"single"`. Computed.

Environment Variable. Environment Variable.

Upstream description:

Environment Variable.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0300011111322231-2323021133103000-1221223213210330-1033111302230323-1001320101211130-3002100013020302-3311122131220310-1322223012031112"></a>

## Direct properties — env_var / 130001320032 / 3

<a id="canonical-3303302023322203-1221221222221322-0322132333211333-2100213111101113-3020321302011221-2232302332330001-3001030132210023-0000301002300322"></a>

<a id="canonical-1031313330301321-0200211003332002-2323230130113331-2302222121102111-1231111300323220-0022333312201132-1221331001233101-0201210221123021"></a>

## name property — env_var / 130001320032 / 4

Type: `"string"`. Computed.

Name. Name of Environment Variable.

Upstream description:

Name of Environment Variable.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2131022102001222-0131023033102131-0113101020013133-0333220233112203-0002022122133233-3331211201310331-3220200002231031-3332032131011211"></a>

<a id="canonical-0213123221013033-2210110312130320-3121233110222312-0201223023032213-1031212233322233-3020020012001231-0003111223000021-3020223323302223"></a>

## value property — env_var / 130001320032 / 5

Type: `"string"`. Computed.

Value. Value of Environment Variable.

Upstream description:

Value of Environment Variable.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0020301003131123-2322000323203123-0010000213302303-2112011322321333-1023033231103333-1220301030033320-1303333103233332-2132200130332001"></a>

## Next pages — env_var / 130001320032 / 6

- [service.configuration.parameters](data-sources--workload--reference--group-015.md#canonical-0012011320103113-3033302112011313-2110032033231102-1120002333001312-0230133122323022-2322200011110000-1310321323010332-0020233203221233)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3332330103203231-2232213211320013-3031002131033012-0231300111321033-0200101300032231-0102031133310322-1202321212212011-3303003200302311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103110213001311-3013020122303213-3311320333202132-1100233113323000-1311203331211312-0331233120323021-3013031002302310-1333332310031012"></a>

## service.configuration.parameters.file — file / 233211032112 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.configuration](data-sources--workload--reference--group-015.md#canonical-3123002002331230-0201233111020120-1333221333320003-3110223112030011-2003333010313113-2333211001212031-3333112223203131-0322131002310210)
- [service.configuration.parameters](data-sources--workload--reference--group-015.md#canonical-0012011320103113-3033302112011313-2110032033231102-1120002333001312-0230133122323022-2322200011110000-1310321323010332-0020233203221233)
- service.configuration.parameters.file

<a id="canonical-3202112021320123-0023331322030323-3223011133301033-1102310210022001-0113030032033233-2103133210032203-0222021100323001-0112020302331021"></a>

Type: `"single"`. Computed.

Configuration File. Configuration File for the workload.

Upstream description:

Configuration File for the workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2123200312001321-2103330131013002-0003202000033333-1103020201101331-1321230330202123-3011121110332312-1220232130120012-1321231212220332"></a>

## Direct properties — file / 233211032112 / 3

<a id="canonical-1103022122020102-0310120331100330-3120202322110120-1320010303320010-3000121003300023-1120033211210312-1330330110223311-3323322133002110"></a>

<a id="canonical-0122202323313301-3332033112123002-2110303212211121-0320303013100101-3011013110012031-1323003203330321-3300320232221213-3331113330201211"></a>

## data property — file / 233211032112 / 4

Type: `"string"`. Computed.

Data. File data

Upstream description:

File data

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16384,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 16384,
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
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [mount](data-sources--workload--reference--group-015.md#canonical-1230212122322002-2121020302223322-1111021002133302-2021301201301010-1012022203330120-2211330111121130-3201330303120333-1123022211033230): complete subsection reference.

<a id="canonical-0201320131302022-1230031231030100-2001231213110100-2200121111223010-1230101103020030-2222222100111223-2110301131121011-0113210200010220"></a>

<a id="canonical-1231123102101331-1102011313303100-0003131231220333-0303002032000211-3123131020200032-1113301112102023-2133020012322211-3132131333312330"></a>

## name property — file / 233211032112 / 5

Type: `"string"`. Computed.

Name. Name of the file.

Upstream description:

Name of the file.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1010331001310233-2111002320212230-0122103000202111-1031033332013120-3122312202222002-2312211322230002-1103010002202200-1131102012132323"></a>

<a id="canonical-2222231230222302-2321330032130112-1013023122120110-3120023323111311-0232320322212212-1312221032210110-1331022101013123-2210211022300221"></a>

## volume_name property — file / 233211032112 / 6

Type: `"string"`. Computed.

Volume Name. Name of the Volume.

Upstream description:

Name of the Volume.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2010022212332103-1231301330223113-1003213033130121-0233201322230200-1232212031333330-3323331013301101-2112113323020023-3311100003303011"></a>

## Next pages — file / 233211032112 / 7

- [service.configuration.parameters.file.mount](data-sources--workload--reference--group-015.md#canonical-1230212122322002-2121020302223322-1111021002133302-2021301201301010-1012022203330120-2211330111121130-3201330303120333-1123022211033230)
- [service.configuration.parameters](data-sources--workload--reference--group-015.md#canonical-0012011320103113-3033302112011313-2110032033231102-1120002333001312-0230133122323022-2322200011110000-1310321323010332-0020233203221233)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1230212122322002-2121020302223322-1111021002133302-2021301201301010-1012022203330120-2211330111121130-3201330303120333-1123022211033230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022320123021312-3320202113020120-3101010330002031-1310311011113011-0311000312031303-2313212331001100-1010230123213302-1320112212200321"></a>

## service.configuration.parameters.file.mount — mount / 123321012233 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.configuration](data-sources--workload--reference--group-015.md#canonical-3123002002331230-0201233111020120-1333221333320003-3110223112030011-2003333010313113-2333211001212031-3333112223203131-0322131002310210)
- [service.configuration.parameters](data-sources--workload--reference--group-015.md#canonical-0012011320103113-3033302112011313-2110032033231102-1120002333001312-0230133122323022-2322200011110000-1310321323010332-0020233203221233)
- [service.configuration.parameters.file](data-sources--workload--reference--group-015.md#canonical-3332330103203231-2232213211320013-3031002131033012-0231300111321033-0200101300032231-0102031133310322-1202321212212011-3303003200302311)
- service.configuration.parameters.file.mount

<a id="canonical-3331122212302203-3100332222020132-1030023321013303-0010312122333302-0313212303310203-3023231312201112-0022213120300130-1003022300233202"></a>

Type: `"single"`. Computed.

Volume mount describes how volume is mounted inside a workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2022101033003102-2231222001231222-2221022022021300-0323113031113312-2131102311033030-2011201331033203-2100031122221121-1120003230320010"></a>

## Direct properties — mount / 123321012233 / 3

<a id="canonical-2231322312222202-0020031311322333-2222223201100200-0101321101122312-0103102222133232-0112010222120222-3132321321120122-1300203020031101"></a>

<a id="canonical-3201113133333010-2100101112033032-2302130301221000-3313213223312230-3330302330020233-3330233201112133-3230101123311200-2021333311030020"></a>

## mode property — mount / 123321012233 / 4

Type: `"string"`. Computed.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3011333330032312-1202321101333022-1330232133230100-1211230323002220-3122203123110122-3332302233212313-1011310121011300-3222223203203120"></a>

<a id="canonical-1220013313302031-2113000103031200-3011203033000212-2321200302301202-3031332233120131-0121130320321330-3020010022311210-2131221322013311"></a>

## mount_path property — mount / 123321012233 / 5

Type: `"string"`. Computed.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
    "pattern": "^[^:]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-3221312222213120-0110120122113320-2112113302111021-2323201202033102-1033232330211333-3012121302123312-0320121330031302-1123013311021223"></a>

<a id="canonical-0011033032103100-3213033201233133-1230331103202033-2102132123332102-3010113023002030-0133212032032320-1210332132010010-2120201102311211"></a>

## sub_path property — mount / 123321012233 / 6

Type: `"string"`. Computed.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1123223031301030-1323102010032232-0203331230230303-3100332033120330-3111002222330012-3010333123203122-3120312321233112-3212301311112200"></a>

## Next pages — mount / 123321012233 / 7

- [service.configuration.parameters.file](data-sources--workload--reference--group-015.md#canonical-3332330103203231-2232213211320013-3031002131033012-0231300111321033-0200101300032231-0102031133310322-1202321212212011-3303003200302311)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232333202302221-0233211020110100-1130133302101020-2312320223211203-0020220031021001-1313002033113332-2002100103023313-1123313222312230"></a>

## service.containers — containers / 303113011002 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- service.containers

<a id="canonical-1030103313112111-3203302313130121-1101132111201320-2302232210123111-1200221232312022-0330311210023211-3301330100012202-1223003001221031"></a>

Type: `"list"`. Computed.

Containers. Containers to use for service.

Upstream description:

Containers to use for service.

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-0120312101103001-3122102320123001-1021320301212322-2020331322230303-3022130002323300-3211230310022300-3133123220311200-2233332002222020"></a>

## Direct properties — containers / 303113011002 / 3

<a id="canonical-1221132310133302-2322320333211033-2210331111120003-2223301101331112-0012100010201223-0203100122130212-2102130210012012-3322013222213302"></a>

<a id="canonical-0130103001202100-0131010102023011-0223313113212320-2133012321201212-0310010002202122-0001302311210131-1230011320210101-2000221023300003"></a>

## args property — containers / 303113011002 / 4

Type: `["list", "string"]`. Computed.

Arguments to the entrypoint. Overrides the Docker image's CMD.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-2021131332311132-0230332233301300-2313302012322210-3201012111003131-0030031013322320-2101130021012002-1023212000300332-3232301320103130"></a>

<a id="canonical-2001021323023312-0211313210231103-2200111222021203-0301223030000311-3033111200201222-2212003311120310-3300002301302031-2322200012110122"></a>

## command property — containers / 303113011002 / 5

Type: `["list", "string"]`. Computed.

Command to execute. Overrides the Docker image's ENTRYPOINT.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

- [custom_flavor](data-sources--workload--reference--group-015.md#canonical-0220222331120113-2122123100010221-0310120331213320-0123133100313303-1201123030002230-0120232210311131-1210023200022301-1323320310110230): complete subsection reference.

- [default_flavor](data-sources--workload--reference--group-015.md#canonical-2233010122103202-0301112200120010-3023311330130020-0101312220031302-2233021333311333-1011212003200123-3331100212230113-1002313301313231): complete subsection reference.

<a id="canonical-2012003210032020-1023000122022221-3130212032320212-1033331223221312-1032310300313233-1101203302332121-2021033020300320-0002232302332133"></a>

<a id="canonical-1121133222101303-3202023322232113-0011212210123230-2110132210301102-2213113122121032-0332000313232312-2323033322130212-0212220100213103"></a>

## flavor property — containers / 303113011002 / 6

Type: `"string"`. Computed.

\[Enum:
CONTAINER\_FLAVOR\_TYPE\_TINY|CONTAINER\_FLAVOR\_TYPE\_MEDIUM|CONTAINER\_FLAVOR\_TYPE\_LARGE\]
Container Flavor type - CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny Tiny containers have limit of 0.1 vCPU
and 256 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium Medium containers have limit
of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_LARGE: Large Large containers
have.. Possible values are \`CONTAINER\_FLAVOR\_TYPE\_TINY\`, \`CONTAINER\_FLAVOR\_TYPE\_MEDIUM\`,
\`CONTAINER\_FLAVOR\_TYPE\_LARGE\`. Defaults to \`CONTAINER\_FLAVOR\_TYPE\_TINY\`.

Upstream description:

Container Flavor type

&#8203;- CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny

Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium

Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_LARGE: Large

Large containers have limit of 1 vCPU and 2048 MiB (mebibyte) memory.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTAINER_FLAVOR_TYPE_TINY",
  "enum": [
    "CONTAINER_FLAVOR_TYPE_TINY",
    "CONTAINER_FLAVOR_TYPE_MEDIUM",
    "CONTAINER_FLAVOR_TYPE_LARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [image](data-sources--workload--reference--group-015.md#canonical-2031322233230031-1232103120020002-3332210332031001-3331321000310203-0102013211002110-0032032121233323-3230101213231312-3030302330121333): complete subsection reference.

<a id="canonical-0320133032010211-1110202010013310-1023011210133203-2001020311112000-2123110000303332-3022120303211303-2321010323130311-3002031003311023"></a>

<a id="canonical-3330222112313032-3330102123220023-1323013322011031-3332231233130032-0112033300331311-1132112330211231-3000310321221231-0110311103312223"></a>

## init_container property — containers / 303113011002 / 7

Type: `"bool"`. Computed.

Specialized container that runs before application container and runs to completion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [liveness_check](data-sources--workload--reference--group-015.md#canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120): complete subsection reference.

<a id="canonical-3333133230213323-1012110220310030-1200323213133201-0213012113331021-1132201013130323-3330200023310031-3113321323323313-3031023332000301"></a>

<a id="canonical-1220131311001330-1323310000313110-1332111130231321-1310320323200203-0303001133211332-3023000233333100-0311010311123202-0321322323302210"></a>

## name property — containers / 303113011002 / 8

Type: `"string"`. Computed.

Name. Name of the container.

Upstream description:

Name of the container.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [readiness_check](data-sources--workload--reference--group-015.md#canonical-1121221330131033-2220320323103232-2110113031323312-1331130001103020-3223102313333101-1132213102032230-1020012000200321-2323233031030012): complete subsection reference.

<a id="canonical-1203011013102332-2201223221333012-3023022302210301-2213123330002202-1310302231310123-0320112121012210-0020103123223310-1133211321120100"></a>

## Next pages — containers / 303113011002 / 9

- [service.containers.custom_flavor](data-sources--workload--reference--group-015.md#canonical-0220222331120113-2122123100010221-0310120331213320-0123133100313303-1201123030002230-0120232210311131-1210023200022301-1323320310110230)
- [service.containers.default_flavor](data-sources--workload--reference--group-015.md#canonical-2233010122103202-0301112200120010-3023311330130020-0101312220031302-2233021333311333-1011212003200123-3331100212230113-1002313301313231)
- [service.containers.image](data-sources--workload--reference--group-015.md#canonical-2031322233230031-1232103120020002-3332210332031001-3331321000310203-0102013211002110-0032032121233323-3230101213231312-3030302330121333)
- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120)
- [service.containers.readiness_check](data-sources--workload--reference--group-015.md#canonical-1121221330131033-2220320323103232-2110113031323312-1331130001103020-3223102313333101-1132213102032230-1020012000200321-2323233031030012)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0220222331120113-2122123100010221-0310120331213320-0123133100313303-1201123030002230-0120232210311131-1210023200022301-1323320310110230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213003312031302-1302032232102113-0203023022130310-2021021021012310-3133031001332110-0122033330231021-1302302010101221-1230022312233332"></a>

## service.containers.custom_flavor — custom_flavor / 302310020131 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- service.containers.custom_flavor

<a id="canonical-1023312133103123-1132100121032322-3230011302030323-0220202100313011-2032030111112300-2300323003223212-1232013011330102-2133201133300221"></a>

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

<a id="canonical-0313031021120311-1122300200133200-2030312001012221-0100212202220231-1213003312201300-3022021021231201-2021312312313332-0303332323233013"></a>

## Direct properties — custom_flavor / 302310020131 / 3

<a id="canonical-2002333010332233-2320310121110012-0230012001231120-0013011222233123-1120013033133012-0332032213200222-2023302232033202-2122032330100032"></a>

<a id="canonical-2132201132132022-3132321122201120-3310221300033133-1200203103321002-0201211301002131-0122231030110020-1122101010131000-1030010032130030"></a>

## name property — custom_flavor / 302310020131 / 4

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

<a id="canonical-2032310230323001-1321301110020011-2311011200313123-3313021221121312-0120332133102313-3230223100220121-3323000232121200-0113301103122030"></a>

<a id="canonical-1312011120333322-1003133020313123-2203020010033011-2200201203202130-2320133200030020-2021110211323200-1013113332021312-0020322333002312"></a>

## namespace property — custom_flavor / 302310020131 / 5

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

<a id="canonical-2301330013021333-3333333110120022-2131331301022102-1322221120023320-1112012023310323-3231211330110013-0122312021320301-0212103323211213"></a>

<a id="canonical-0122123020302003-0301330123001201-1230232230100300-2133300323130213-0103201122220303-2313321221232121-0223000032101020-2000001101313123"></a>

## tenant property — custom_flavor / 302310020131 / 6

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

<a id="canonical-1332210020010323-1211211103021000-2312103030212101-3311021013033131-3013331131112232-3320311112021001-2023121101330221-3023022313222331"></a>

## Next pages — custom_flavor / 302310020131 / 7

- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2233010122103202-0301112200120010-3023311330130020-0101312220031302-2233021333311333-1011212003200123-3331100212230113-1002313301313231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110322001220302-0223202020331333-3120210223013320-1031102132212122-0120132330001110-3331203033030223-3321003332203201-3101313102001222"></a>

## service.containers.default_flavor — default_flavor / 130101220233 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- service.containers.default_flavor

<a id="canonical-1102131220000011-3030011122332331-3003211220100103-1103333000132313-2232101312012112-3311200312222100-0031033110203220-2133032020320300"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default flavor.

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

<a id="canonical-3103023333311302-3103202121320000-3030330112131101-1200232022122321-3331233203002031-3302333333311002-2310303123211101-0023110131201230"></a>

## Direct properties — default_flavor / 130101220233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100023123332211-0322231023010220-0313100010101202-0321100111212021-0110331332122130-3202002121302231-1030201321122022-0200020320012220"></a>

## Next pages — default_flavor / 130101220233 / 4

- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2031322233230031-1232103120020002-3332210332031001-3331321000310203-0102013211002110-0032032121233323-3230101213231312-3030302330121333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233223102202201-1232202133121211-2231011132310001-0100102203322311-2331310013020231-3312032201003102-1312221110222233-1202203310311010"></a>

## service.containers.image — image / 021133033130 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- service.containers.image

<a id="canonical-2031303113222323-1003133102310220-1123133330203012-3001301132110303-2000003031231002-0213022223213330-1000232323223322-0222321203101133"></a>

Type: `"single"`. Computed.

ImageType configures the image to use, how to pull the image, and the associated secrets to use if
any.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-registry_choice": "[\"container_registry\",\"public\"]"
}
```

<a id="canonical-1030030323200013-0211323212202010-1320322021030102-1221021201220000-0111110231010301-3133001223012003-2131213111312123-3211200123000303"></a>

## Direct properties — image / 021133033130 / 3

- [container_registry](data-sources--workload--reference--group-015.md#canonical-0200232232103113-0020033231030001-2112000331001000-1120312203111231-2131221223033001-0112022312302223-1232302032023123-3100021001322300): complete subsection reference.

<a id="canonical-2210332323133001-0213112231321210-0203110200331122-3003011031333220-2302203321231132-3230312013133123-1222223320321200-3023020123322003"></a>

<a id="canonical-2301013111333132-1210203103303032-2020202301102230-1011033121201202-2331033131223231-1120200023022203-0211222003112330-1303200321131003"></a>

## name property — image / 021133033130 / 4

Type: `"string"`. Computed.

Name is a container image which are usually given a name such as alpine, Ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed.

Upstream description:

Name is a container image which are usually given a name such as alpine, Ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [public](data-sources--workload--reference--group-015.md#canonical-1013002102221300-3120310221330313-2032123311323001-3102220200300221-3323322030321330-0010100030212200-2220330023130320-1320223112003200): complete subsection reference.

<a id="canonical-3220011220110231-0321232322030131-3010220332033203-1000323230230313-1032032130120001-1212132013131133-0110321221001013-0302321210303102"></a>

<a id="canonical-1123123200303030-0221010133020202-2022302232133201-1312223111020313-2312213302221003-0110200313031310-3032200322000123-0100233103300221"></a>

## pull_policy property — image / 021133033130 / 5

Type: `"string"`. Computed.

\[Enum:
IMAGE\_PULL\_POLICY\_DEFAULT|IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT|IMAGE\_PULL\_POLICY\_ALWAYS|IMAGE\_PULL\_POLICY\_NEVER\]
Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload - IMAGE\_PULL\_POLICY\_DEFAULT: Default Default will always pull image if :latest tag
is specified in image name. If :latest tag is not specified in image name, it will pull image only..
Possible values are \`IMAGE\_PULL\_POLICY\_DEFAULT\`, \`IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT\`,
\`IMAGE\_PULL\_POLICY\_ALWAYS\`, \`IMAGE\_PULL\_POLICY\_NEVER\`. Defaults to
\`IMAGE\_PULL\_POLICY\_DEFAULT\`.

Upstream description:

Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload

&#8203;- IMAGE\_PULL\_POLICY\_DEFAULT: Default

Default will always pull image if :latest tag is specified in image name. If :latest tag is not
specified in image name, it will pull image only if it does not already exist on the node &#8203;-
IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT: IfNotPresent

Only pull the image if it does not already exist on the node &#8203;- IMAGE\_PULL\_POLICY\_ALWAYS:
Always

Always pull the image &#8203;- IMAGE\_PULL\_POLICY\_NEVER: Never

Never pull the image.

Receipt-pinned upstream constraints:

```json
{
  "default": "IMAGE_PULL_POLICY_DEFAULT",
  "enum": [
    "IMAGE_PULL_POLICY_DEFAULT",
    "IMAGE_PULL_POLICY_IF_NOT_PRESENT",
    "IMAGE_PULL_POLICY_ALWAYS",
    "IMAGE_PULL_POLICY_NEVER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0120200020101202-1233111022202231-1023300123010031-2132231133200030-2303030000310000-1232113200023021-1213220230202201-2100000000121023"></a>

## Next pages — image / 021133033130 / 6

- [service.containers.image.container_registry](data-sources--workload--reference--group-015.md#canonical-0200232232103113-0020033231030001-2112000331001000-1120312203111231-2131221223033001-0112022312302223-1232302032023123-3100021001322300)
- [service.containers.image.public](data-sources--workload--reference--group-015.md#canonical-1013002102221300-3120310221330313-2032123311323001-3102220200300221-3323322030321330-0010100030212200-2220330023130320-1320223112003200)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0200232232103113-0020033231030001-2112000331001000-1120312203111231-2131221223033001-0112022312302223-1232302032023123-3100021001322300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321022301201233-3021223013211211-3221330113322203-0012122321000312-0332032232033000-1123030130212313-1330213211333131-1001121203303301"></a>

## service.containers.image.container_registry — container_registry / 111330301232 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.image](data-sources--workload--reference--group-015.md#canonical-2031322233230031-1232103120020002-3332210332031001-3331321000310203-0102013211002110-0032032121233323-3230101213231312-3030302330121333)
- service.containers.image.container_registry

<a id="canonical-2101203103200010-1100223010230030-0023320222102321-0001202130230111-0003223001330021-1030221300120201-1012312301110011-3110233012222303"></a>

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

<a id="canonical-3031203331133112-1111210033011102-0132021102303233-3313211220221011-0000302223100301-0212323123123003-0001200122131020-1121121220320311"></a>

## Direct properties — container_registry / 111330301232 / 3

<a id="canonical-1202131010300301-0212102223133130-3021020103000020-1202032332023330-2332021013020210-3112203302212100-0223003233200000-0013213003000310"></a>

<a id="canonical-2021030012122122-0301122332003101-2121101331330103-1203132320011310-3013011302313321-3122031001210302-0221001202101110-2131311310313330"></a>

## name property — container_registry / 111330301232 / 4

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

<a id="canonical-1032111120122122-0222002030220032-0123110130131203-0201133030310222-3302233210111231-3131232312302331-3211232302222322-1002010120332223"></a>

<a id="canonical-2123200230223312-2101311232020300-0223120301123220-3130002013211302-3331223120030302-0120131332002033-3000010223200222-2113322321231223"></a>

## namespace property — container_registry / 111330301232 / 5

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

<a id="canonical-3321223103031210-0011011110121332-3312000302213121-2221312310231110-2133032311233120-0322022020311331-1331021010001022-0322212312202023"></a>

<a id="canonical-1110020302020112-0221322233122022-3033203101031100-3120303203010201-2332010131100113-0311331313321333-0333011100220333-0032113320201122"></a>

## tenant property — container_registry / 111330301232 / 6

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

<a id="canonical-2233312303321023-0103010323033220-3012033203312121-0230211201312200-1030321321223302-0223301210000200-3001312110303311-2010213012233333"></a>

## Next pages — container_registry / 111330301232 / 7

- [service.containers.image](data-sources--workload--reference--group-015.md#canonical-2031322233230031-1232103120020002-3332210332031001-3331321000310203-0102013211002110-0032032121233323-3230101213231312-3030302330121333)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1013002102221300-3120310221330313-2032123311323001-3102220200300221-3323322030321330-0010100030212200-2220330023130320-1320223112003200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300203110313120-0020212013331323-3331300313220001-1122011301300331-2202223222201302-1201000233232320-1302022332311031-0012211220201312"></a>

## service.containers.image.public — public / 021333223003 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.image](data-sources--workload--reference--group-015.md#canonical-2031322233230031-1232103120020002-3332210332031001-3331321000310203-0102013211002110-0032032121233323-3230101213231312-3030302330121333)
- service.containers.image.public

<a id="canonical-3211210123200031-1122231210312111-3123123101313000-3220022113210112-1132210203301033-0010220133233303-2232311303103033-2103110311202021"></a>

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

<a id="canonical-2121021002233031-2332300221301332-1333330032000103-0132330303023031-3203323131013111-1032212202112113-1331313121200113-2022001220322020"></a>

## Direct properties — public / 021333223003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322323201233131-2222123022330302-1122203210213320-3221211121301201-3213312223233113-1320123313101111-2021200030312013-0333223320113333"></a>

## Next pages — public / 021333223003 / 4

- [service.containers.image](data-sources--workload--reference--group-015.md#canonical-2031322233230031-1232103120020002-3332210332031001-3331321000310203-0102013211002110-0032032121233323-3230101213231312-3030302330121333)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312113323331223-2221130130003232-0211221000122001-3332113332303201-2201012303330221-0111002122221121-2112021131123000-2300011102123210"></a>

## service.containers.liveness_check — liveness_check / 231133012120 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- service.containers.liveness_check

<a id="canonical-3002022213311203-1223201000312100-2013131001201332-1102111223132332-3013113313313132-1130313000010131-0202002101130312-2322002101010202"></a>

Type: `"single"`. Computed.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

<a id="canonical-3333323122021311-3223332202221322-3313020200300221-3323213010023111-3033103301301330-2133033132131222-0230132320322003-1223123031302311"></a>

## Direct properties — liveness_check / 231133012120 / 3

- [exec_health_check](data-sources--workload--reference--group-015.md#canonical-1332232120312312-0111300020001122-2231211200220311-2010220321113020-2313123020123031-0232031230123021-0311021231003023-1203331203000122): complete subsection reference.

<a id="canonical-0222230303310103-2102223013222131-0131031230120323-0010223231123100-2312110330222010-0122123113203210-1222022230102101-2321320232213223"></a>

<a id="canonical-1320312323003303-0022200323202110-1200203023021231-1010022320011033-1133001233120020-2112003210302321-0302011112031222-2322100223022002"></a>

## healthy_threshold property — liveness_check / 231133012120 / 4

Type: `"number"`. Computed.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](data-sources--workload--reference--group-015.md#canonical-1123231322133301-2311101311332033-2000320331330302-2030201003323030-0023112021030200-2033100032113223-3133033003133121-3120021323313200): complete subsection reference.

<a id="canonical-0121331330011121-0011011211130303-2002110300203210-0120333133002123-2033022323333121-3231112322211322-0332221122031033-3222321300322023"></a>

<a id="canonical-3131021233201011-1320333013002011-2013220003213301-1011230033332020-2100233321200001-2110130233201120-3203313213123033-1232300010200132"></a>

## initial_delay property — liveness_check / 231133012120 / 5

Type: `"number"`. Computed.

Number of seconds after the container has started before health checks are initiated.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-0331210002011332-0201102122331112-0112301032120001-3120131213112222-0231230001011122-2220230120321120-2103113213201013-3130033122213310"></a>

<a id="canonical-1232001002310300-2122110211313100-1001101113203122-0021300011210100-2130313012312112-2231023133131302-2331320211233210-3001020313210023"></a>

## interval property — liveness_check / 231133012120 / 6

Type: `"number"`. Computed.

Time interval in seconds between two health check requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [tcp_health_check](data-sources--workload--reference--group-015.md#canonical-2233021321103233-3232122010233002-1230103200310032-0022111223313331-0020133301023222-3022133030203123-2120022220102022-0001202030200101): complete subsection reference.

<a id="canonical-0210103221121330-3323002101313002-3330312312120012-1301022201300103-3010233200020013-0332022313310203-1303001000101031-2011312232033112"></a>

<a id="canonical-1120330200122001-1220112110012301-2223133033031200-1220330013002201-2031312321210012-2100122320230013-1230311131111010-3230320222301303"></a>

## timeout property — liveness_check / 231133012120 / 7

Type: `"number"`. Computed.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-0313210231121022-3200320222210222-1101223031113212-3031030010331012-2321232323332323-1112312220300013-2000312300023312-3313230033101121"></a>

<a id="canonical-2303302321010211-1312102223022133-3333100021130033-1320013001313323-2131011132220222-3320123030003001-0122201312030112-3202033223012000"></a>

## unhealthy_threshold property — liveness_check / 231133012120 / 8

Type: `"number"`. Computed.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-1300031010032233-3222133013133112-0311210032222213-0020213123233220-0201002322330102-2012131121200131-1320330213101311-0300212101101012"></a>

## Next pages — liveness_check / 231133012120 / 9

- [service.containers.liveness_check.exec_health_check](data-sources--workload--reference--group-015.md#canonical-1332232120312312-0111300020001122-2231211200220311-2010220321113020-2313123020123031-0232031230123021-0311021231003023-1203331203000122)
- [service.containers.liveness_check.http_health_check](data-sources--workload--reference--group-015.md#canonical-1123231322133301-2311101311332033-2000320331330302-2030201003323030-0023112021030200-2033100032113223-3133033003133121-3120021323313200)
- [service.containers.liveness_check.tcp_health_check](data-sources--workload--reference--group-015.md#canonical-2233021321103233-3232122010233002-1230103200310032-0022111223313331-0020133301023222-3022133030203123-2120022220102022-0001202030200101)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1332232120312312-0111300020001122-2231211200220311-2010220321113020-2313123020123031-0232031230123021-0311021231003023-1203331203000122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021131203313101-1103002322323310-0231133201200010-1120011310100102-0310010021321033-2200211202231310-2102321220231122-2311322013200023"></a>

## service.containers.liveness_check.exec_health_check — exec_health_check / 031132101311 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120)
- service.containers.liveness_check.exec_health_check

<a id="canonical-0221001223012333-3201303221203101-2010321102331101-1033323223123111-2121132010320011-2003020213220230-2111003103331203-3200121120310220"></a>

Type: `"single"`. Computed.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0300323003111302-2212310032301130-1102011132031021-0000121211310022-1331321320003223-1110202311110232-1110000133003000-0103112110132313"></a>

## Direct properties — exec_health_check / 031132101311 / 3

<a id="canonical-1302032013111211-1110313102032321-0132201103311032-1002132010223130-1110311302111102-0321013323202011-0112221213303011-1313221311300210"></a>

<a id="canonical-1300023012321131-1103121010302111-1111312020031022-1230200202320221-2323020131312223-3302133311113030-3131011013210323-2023322031023203"></a>

## command property — exec_health_check / 031132101311 / 4

Type: `["list", "string"]`. Computed.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0201013312023223-3121111010110303-2020212301231103-3303303100133223-2220131022101320-2031113022131310-2202030131131002-3230132103233302"></a>

## Next pages — exec_health_check / 031132101311 / 5

- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1123231322133301-2311101311332033-2000320331330302-2030201003323030-0023112021030200-2033100032113223-3133033003133121-3120021323313200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011312130033012-0232110011012220-1331332123221301-1233013300313133-3230133100023211-3101300000223231-3000232303112303-0012330231232111"></a>

## service.containers.liveness_check.http_health_check — http_health_check / 000201302323 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120)
- service.containers.liveness_check.http_health_check

<a id="canonical-0002032311221020-2333203000210031-1131223201102322-1033013000122110-2100311001200031-0303221133330222-2333100030111300-3110011232201232"></a>

Type: `"single"`. Computed.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1023201032212202-0130220102332001-2231210012310232-3333000332110231-1103020300130031-2003030022332002-2331303312123001-2013233031312302"></a>

## Direct properties — http_health_check / 000201302323 / 3

<a id="canonical-3211021011310231-0231333321232213-1223220112013213-1220103002132003-0233012003212233-0303113133210102-0033323000233030-0102330223233102"></a>

<a id="canonical-1322030130332010-1132232220010222-0101210321112133-0203222203300031-0312021122030220-2133000020121132-3003201320330300-1230200122130031"></a>

## headers property — http_health_check / 000201302323 / 4

Type: `["map", "string"]`. Computed.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 256,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "256",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "2048",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 2048,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-3221221130301320-3011001032303220-2101321310230122-2221021011112002-3312221032013300-3321133130103013-3222123111103112-3103222003202232"></a>

<a id="canonical-3321121023020333-3013200202320120-2223112001211001-0122333013202032-0123310130101023-0121322203320102-3102100230230220-0002323023300030"></a>

## host_header property — http_health_check / 000201302323 / 5

Type: `"string"`. Computed.

The value of the host header in the HTTP health check request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-0210210020320030-1112023323222313-2000213102102022-3213310031212030-0221013231133321-3003022322211301-1321011021200233-0012101101030133"></a>

<a id="canonical-0212233312011310-1220131232121021-0331302120102110-3312231331202022-3103202021221103-1011032022020200-0033233003133200-3031302132120102"></a>

## path property — http_health_check / 000201302323 / 6

Type: `"string"`. Computed.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

- [port](data-sources--workload--reference--group-015.md#canonical-2323332311033013-0102132333310313-1220303311232030-2323132012203211-0203132131332323-3022101301331332-1333333132132203-0022111011012021): complete subsection reference.

<a id="canonical-2001130113030111-2221232320323330-2132030233320103-2033322101212133-2333112302322203-3323010012201000-1123121100222233-3033121232212203"></a>

## Next pages — http_health_check / 000201302323 / 7

- [service.containers.liveness_check.http_health_check.port](data-sources--workload--reference--group-015.md#canonical-2323332311033013-0102132333310313-1220303311232030-2323132012203211-0203132131332323-3022101301331332-1333333132132203-0022111011012021)
- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2323332311033013-0102132333310313-1220303311232030-2323132012203211-0203132131332323-3022101301331332-1333333132132203-0022111011012021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202202222112333-2031112112200320-1012320300021203-0232130210300212-3303323201232021-1023103011112300-3013032321230312-2322321113033220"></a>

## service.containers.liveness_check.http_health_check.port — port / 330030321301 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120)
- [service.containers.liveness_check.http_health_check](data-sources--workload--reference--group-015.md#canonical-1123231322133301-2311101311332033-2000320331330302-2030201003323030-0023112021030200-2033100032113223-3133033003133121-3120021323313200)
- service.containers.liveness_check.http_health_check.port

<a id="canonical-1022310121330110-3101201030312222-2322302100021010-3003112021120110-3010202003102303-0213111120023200-0312010023310310-2002020122323210"></a>

Type: `"single"`. Computed.

Port. Port

Upstream description:

Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

<a id="canonical-1200131102113301-0110311003101022-2100030123000213-2120320011300312-3001100132312230-0211233030313311-1010301000202213-2321020131202332"></a>

## Direct properties — port / 330030321301 / 3

<a id="canonical-1331310130322021-0233130120330220-3322310123000010-3321122213212232-2102030230223101-2201101220233232-3230323131113100-2331302222132030"></a>

<a id="canonical-0213212130033223-0131121331203223-0330323222130133-3110302022230002-2201212211122321-0030321102201100-2300311330232003-0222203023233231"></a>

## name property — port / 330030321301 / 4

Type: `"string"`. Computed.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-2032210103013003-0301121102221330-0013022120033010-0321001220313013-0320222110221021-2221233233223300-0103030303033130-0331112223223133"></a>

<a id="canonical-1110020302220112-3321330300013330-3233110133320130-2220312131303031-1102311203323232-0020320310101211-1113212030331232-1120032213003111"></a>

## num property — port / 330030321301 / 5

Type: `"number"`. Computed.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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
    "minimum": 1
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

<a id="canonical-3222222003131310-3010212100021020-0331020232300231-0321211221321013-1301101030010103-0231302311303210-3330032211101100-1311123333201221"></a>

## Next pages — port / 330030321301 / 6

- [service.containers.liveness_check.http_health_check](data-sources--workload--reference--group-015.md#canonical-1123231322133301-2311101311332033-2000320331330302-2030201003323030-0023112021030200-2033100032113223-3133033003133121-3120021323313200)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2233021321103233-3232122010233002-1230103200310032-0022111223313331-0020133301023222-3022133030203123-2120022220102022-0001202030200101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111020322021332-1010221110003010-0213300101230131-1222200033030310-3020131032102131-0223223201221332-2232301220311123-3303131102310301"></a>

## service.containers.liveness_check.tcp_health_check — tcp_health_check / 333333133011 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120)
- service.containers.liveness_check.tcp_health_check

<a id="canonical-3130110123132001-3220123000232233-0021001322321110-1300220323223102-1111231230021320-1032011013030211-0203120002102103-3101132122233303"></a>

Type: `"single"`. Computed.

TCPHealthCheckType describes a health check based on opening a TCP connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0232332113002020-0120231233223010-0110131332330132-0221010011132200-2133320222110032-0302010213030302-1012332133301200-1133031030021023"></a>

## Direct properties — tcp_health_check / 333333133011 / 3

- [port](data-sources--workload--reference--group-015.md#canonical-2220312333010203-0013130000120301-3310011311323303-0313033321300332-1123031313331012-1221212300313320-3130211310110110-0000301031301023): complete subsection reference.

<a id="canonical-0221210213233013-3111121210123031-0203121333332302-1112301321002110-1123332320023103-2312100213102323-3323223203200213-2001233301003020"></a>

## Next pages — tcp_health_check / 333333133011 / 4

- [service.containers.liveness_check.tcp_health_check.port](data-sources--workload--reference--group-015.md#canonical-2220312333010203-0013130000120301-3310011311323303-0313033321300332-1123031313331012-1221212300313320-3130211310110110-0000301031301023)
- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2220312333010203-0013130000120301-3310011311323303-0313033321300332-1123031313331012-1221212300313320-3130211310110110-0000301031301023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303111022020101-2212113030121200-0311223131011023-1302132113222011-2301312103103330-0330211012112200-0032333222321121-2013130003132030"></a>

## service.containers.liveness_check.tcp_health_check.port — port / 032002212130 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.liveness_check](data-sources--workload--reference--group-015.md#canonical-2103112123301021-3000100001311232-2221202331012020-1010211113203123-1320111102121332-0300112332133120-1031003113232032-0101031323013120)
- [service.containers.liveness_check.tcp_health_check](data-sources--workload--reference--group-015.md#canonical-2233021321103233-3232122010233002-1230103200310032-0022111223313331-0020133301023222-3022133030203123-2120022220102022-0001202030200101)
- service.containers.liveness_check.tcp_health_check.port

<a id="canonical-1321301200221112-1010312321210231-0132113300122332-0220300331203203-3031312321023112-2103110012110313-2313201303200110-0021231011222000"></a>

Type: `"single"`. Computed.

Port. Port

Upstream description:

Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

<a id="canonical-0010311031002111-0113021312313220-2310210312210013-3200203031201301-2100232310212233-3310200031211312-3200320023113001-2220013100003221"></a>

## Direct properties — port / 032002212130 / 3

<a id="canonical-3220231210112013-2132101210030213-2012310333113103-3212321002232121-1212133003313311-0321200023112113-2112322133301310-0133320300033100"></a>

<a id="canonical-3313121332012131-3303233032333323-2312012111121202-2221320033313130-2330210122310110-2310113313022032-2201122032200232-0332100221320121"></a>

## name property — port / 032002212130 / 4

Type: `"string"`. Computed.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-2003230100201220-1120332200022202-3301003133102121-1133320003202232-1232001300100130-1211023000030132-0011031010011003-3022122333012330"></a>

<a id="canonical-0320221003222310-1213322110211033-0321001220132013-3332032213103000-0320201220133233-3111030201310230-3310013202321011-2022230310112112"></a>

## num property — port / 032002212130 / 5

Type: `"number"`. Computed.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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
    "minimum": 1
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

<a id="canonical-1322003112201312-2000302223200113-3301001020231312-0321030203210321-0202211300132003-1232111013322303-2032123120131030-2133211302003012"></a>

## Next pages — port / 032002212130 / 6

- [service.containers.liveness_check.tcp_health_check](data-sources--workload--reference--group-015.md#canonical-2233021321103233-3232122010233002-1230103200310032-0022111223313331-0020133301023222-3022133030203123-2120022220102022-0001202030200101)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1121221330131033-2220320323103232-2110113031323312-1331130001103020-3223102313333101-1132213102032230-1020012000200321-2323233031030012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000332221110231-1330120312120222-0200222202231102-2111133012020033-1021113221130333-0231001030232320-3330111312130303-2032223230201112"></a>

## service.containers.readiness_check — readiness_check / 103203312220 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- service.containers.readiness_check

<a id="canonical-2123302320010122-3032113300023011-2200330022002310-3113002320003321-3100203203323022-2311131112100322-0010301001131122-2300122111223223"></a>

Type: `"single"`. Computed.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

<a id="canonical-1120213122211330-1022221032323113-1021220311112031-2231132101031011-3331122201110303-3110111113222311-3233200330200211-2112132332322121"></a>

## Direct properties — readiness_check / 103203312220 / 3

- [exec_health_check](data-sources--workload--reference--group-015.md#canonical-3203123232102231-0203100231032212-1113300221012032-2002031012232310-3333313011131301-1220231231312223-3221201123120101-2103323030021200): complete subsection reference.

<a id="canonical-3220013222312332-0001311132121203-1233223211001311-1010023130231301-3220210022221203-2323120320233112-1000212030201223-3222300003203332"></a>

<a id="canonical-2133301312122023-2023312222223220-2213311313213312-0002213323121103-2023123002222131-1021212021032021-1303123023023201-2103133012330210"></a>

## healthy_threshold property — readiness_check / 103203312220 / 4

Type: `"number"`. Computed.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](data-sources--workload--reference--group-015.md#canonical-1223323322103201-2103302313231120-0103310320002203-2333222312011203-1020003312022100-0001111000211211-1111312111112030-3130210210222130): complete subsection reference.

<a id="canonical-2111222010330032-2121212123301333-2203123233120133-2010122233211031-3120222311000220-3010031310123031-0201310323333230-1120230033332122"></a>

<a id="canonical-3311302220001021-1220302111100233-1211023132210002-0032330030013011-2122221300122223-1110020330112321-3302211113212022-1213100212222111"></a>

## initial_delay property — readiness_check / 103203312220 / 5

Type: `"number"`. Computed.

Number of seconds after the container has started before health checks are initiated.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-0333003132311012-1023301110303202-2120001230213222-3331102320011212-2210003002021130-0202022000303332-2110302033130030-0112033031101300"></a>

<a id="canonical-1331210011332321-2012213330221213-1202301020032130-0322001310313221-0212312022130212-0011113322210211-0133002020333231-1002210112120132"></a>

## interval property — readiness_check / 103203312220 / 6

Type: `"number"`. Computed.

Time interval in seconds between two health check requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [tcp_health_check](data-sources--workload--reference--group-016.md#canonical-1212111010322100-1002213020102212-2130303313110301-0233301101332221-2321303033123301-3012032310002132-2003232123020003-0300221002312130): complete subsection reference.

<a id="canonical-0311121113221222-3211330013220332-1212022033030302-2113013100101023-3122333201133032-3031200031322123-0010113312312003-0201311312113031"></a>

<a id="canonical-3221321322100220-0202101121013212-1213121301231111-2323213200322213-3130200320323232-2230202121120101-3131000222002330-3203122000323111"></a>

## timeout property — readiness_check / 103203312220 / 7

Type: `"number"`. Computed.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-1302210132200212-1321133102201103-1232301131101202-3030201130212002-1303220223033223-3122111012111000-0220003221010200-3003130130231301"></a>

<a id="canonical-1200013133323110-0011030022222002-2123232203110020-3320311212132030-3002201132301122-1122120013102010-3011310322200233-0203132331203231"></a>

## unhealthy_threshold property — readiness_check / 103203312220 / 8

Type: `"number"`. Computed.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-1233313222320013-1102020110031021-0100203311303012-1303213012200023-0202002001311312-2200332210212130-2312202201011012-0133210112102011"></a>

## Next pages — readiness_check / 103203312220 / 9

- [service.containers.readiness_check.exec_health_check](data-sources--workload--reference--group-015.md#canonical-3203123232102231-0203100231032212-1113300221012032-2002031012232310-3333313011131301-1220231231312223-3221201123120101-2103323030021200)
- [service.containers.readiness_check.http_health_check](data-sources--workload--reference--group-015.md#canonical-1223323322103201-2103302313231120-0103310320002203-2333222312011203-1020003312022100-0001111000211211-1111312111112030-3130210210222130)
- [service.containers.readiness_check.tcp_health_check](data-sources--workload--reference--group-016.md#canonical-1212111010322100-1002213020102212-2130303313110301-0233301101332221-2321303033123301-3012032310002132-2003232123020003-0300221002312130)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3203123232102231-0203100231032212-1113300221012032-2002031012232310-3333313011131301-1220231231312223-3221201123120101-2103323030021200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120302201310110-1032312322003331-2201200110303301-3203321231111222-3211003013021102-3232112023020011-3212202233213130-2102101211113232"></a>

## service.containers.readiness_check.exec_health_check — exec_health_check / 221100102310 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.readiness_check](data-sources--workload--reference--group-015.md#canonical-1121221330131033-2220320323103232-2110113031323312-1331130001103020-3223102313333101-1132213102032230-1020012000200321-2323233031030012)
- service.containers.readiness_check.exec_health_check

<a id="canonical-3021311203312313-2300120122230002-3130012232023111-0013010010333313-1200103112310130-1332231122122012-1033200211310101-1322222300030213"></a>

Type: `"single"`. Computed.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3210310130332101-1331020011220130-2011233103111102-2322310222021201-2012223033230002-2030023003322130-1301122121013021-2200312013333133"></a>

## Direct properties — exec_health_check / 221100102310 / 3

<a id="canonical-3020211010213321-1310033030213022-3031121002132300-0000220032120221-3230100102221311-2212112210010232-2200321303021220-3102220300100000"></a>

<a id="canonical-2210222112000312-1022133221131111-3112312221022230-2103320231210021-2331100210331122-2101120333021333-0233200300201300-3122111222013222"></a>

## command property — exec_health_check / 221100102310 / 4

Type: `["list", "string"]`. Computed.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1101030300303130-0303011200113212-1122201331232031-1322322301210333-0113000121302100-1211100002120111-0032320131320121-0310333311330313"></a>

## Next pages — exec_health_check / 221100102310 / 5

- [service.containers.readiness_check](data-sources--workload--reference--group-015.md#canonical-1121221330131033-2220320323103232-2110113031323312-1331130001103020-3223102313333101-1132213102032230-1020012000200321-2323233031030012)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1223323322103201-2103302313231120-0103310320002203-2333222312011203-1020003312022100-0001111000211211-1111312111112030-3130210210222130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300020133302210-2201000300000333-1313131321002120-2223132000000321-0010222110031222-3200223112313103-1310321031130330-2133222201000233"></a>

## service.containers.readiness_check.http_health_check — http_health_check / 303132331312 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.readiness_check](data-sources--workload--reference--group-015.md#canonical-1121221330131033-2220320323103232-2110113031323312-1331130001103020-3223102313333101-1132213102032230-1020012000200321-2323233031030012)
- service.containers.readiness_check.http_health_check

<a id="canonical-2100303330220011-3321311300123220-2000020210010130-2020221221113332-0030332103303112-0311032112212203-2312212132112131-3012302102330333"></a>

Type: `"single"`. Computed.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2110301030231102-1223031233201203-2010223100011313-3311012110332013-2103300232233213-2202033332200133-3231011210230203-3101310003301023"></a>

## Direct properties — http_health_check / 303132331312 / 3

<a id="canonical-0223100133010033-0332103330232110-3212030131132133-2023113232311120-0200201030030231-0133132123133101-1010021313023020-1103032102331223"></a>

<a id="canonical-2033123102121311-3130012131131013-0101020201302312-0302333112103121-1233123110232210-3121222322100101-3230200112210012-2230331103321221"></a>

## headers property — http_health_check / 303132331312 / 4

Type: `["map", "string"]`. Computed.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 256,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "256",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "2048",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 2048,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-0023332301222323-0033332303111020-2212300331321221-0332310113201313-0030322230222313-2012332222132023-2103023113101111-0101233221120033"></a>

<a id="canonical-2120111133321001-1031201210020321-0211223002121122-2020000001211321-1131201232010003-3100331021031211-1302010323123033-1303011033131302"></a>

## host_header property — http_health_check / 303132331312 / 5

Type: `"string"`. Computed.

The value of the host header in the HTTP health check request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-0312210320012002-0310333110303232-1123330111203322-0332112321301211-1301230320120221-0333302220220201-1120331120221010-1121002011213222"></a>

<a id="canonical-0111233312310130-3332300231120310-1312331033110100-0311201113331331-0200231001332310-0312020023010000-1112010031330213-3320030123131110"></a>

## path property — http_health_check / 303132331312 / 6

Type: `"string"`. Computed.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

- [port](data-sources--workload--reference--group-015.md#canonical-2331301302232210-3223012223110031-1001322111123023-3230303200030321-2033220311223231-2321111202130233-2003303303100321-1230033120132123): complete subsection reference.

<a id="canonical-1320203013220110-2121113033011210-1300210101232203-2113223001301331-2223213033312111-1323112310103302-3331312122100203-2232131011022332"></a>

## Next pages — http_health_check / 303132331312 / 7

- [service.containers.readiness_check.http_health_check.port](data-sources--workload--reference--group-015.md#canonical-2331301302232210-3223012223110031-1001322111123023-3230303200030321-2033220311223231-2321111202130233-2003303303100321-1230033120132123)
- [service.containers.readiness_check](data-sources--workload--reference--group-015.md#canonical-1121221330131033-2220320323103232-2110113031323312-1331130001103020-3223102313333101-1132213102032230-1020012000200321-2323233031030012)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2331301302232210-3223012223110031-1001322111123023-3230303200030321-2033220311223231-2321111202130233-2003303303100321-1230033120132123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313012001110331-1130022311332113-3130021211130133-1020330011023013-3102020030133121-0322212302131031-3203103212312121-1331323233101133"></a>

## service.containers.readiness_check.http_health_check.port — port / 101320012132 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.containers](data-sources--workload--reference--group-015.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033)
- [service.containers.readiness_check](data-sources--workload--reference--group-015.md#canonical-1121221330131033-2220320323103232-2110113031323312-1331130001103020-3223102313333101-1132213102032230-1020012000200321-2323233031030012)
- [service.containers.readiness_check.http_health_check](data-sources--workload--reference--group-015.md#canonical-1223323322103201-2103302313231120-0103310320002203-2333222312011203-1020003312022100-0001111000211211-1111312111112030-3130210210222130)
- service.containers.readiness_check.http_health_check.port

<a id="canonical-1110322020311323-2102302110323013-1033221202230303-1221022101001233-2100021313023123-2330030333102103-2030103001320011-2132321202100230"></a>

Type: `"single"`. Computed.

Port. Port

Upstream description:

Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

<a id="canonical-2133220301103203-0021300030213232-2201000301200332-3111031023111302-0333220221122220-1231033212321111-2231033012130332-0233101301030110"></a>

## Direct properties — port / 101320012132 / 3

<a id="canonical-2213323113130122-3320301122330223-1303322311200122-2012322010000011-3213003133120131-1000323300310020-0102212202232231-2330230223013213"></a>

<a id="canonical-1002220232311102-0021013100203102-2313233211123131-3233211213013213-2322221001203333-0210212121012321-2300012121123033-2012221211130333"></a>

## name property — port / 101320012132 / 4

Type: `"string"`. Computed.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-3231012302200300-2212302213313010-2012003031302203-3232231323031221-1000100320300132-0023323111012300-3113120210311301-2131222012333031"></a>

<a id="canonical-2111002032200201-1223101332003210-3332213031302133-3321022031033311-0303230030133020-1030313230323111-1210133033233211-1003302303130310"></a>

## num property — port / 101320012132 / 5

Type: `"number"`. Computed.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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
    "minimum": 1
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
