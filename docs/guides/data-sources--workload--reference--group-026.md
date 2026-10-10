---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1200110130102320-0121231112030201-3130002232103131-0312211110213031-1221230122133313-3012231113202013-3232233232221320-2103301123023102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-0203020030131321-3103100120211300-3010122121322230-1011010220201113-0330123002023320-2021222213312112-1132031021311130-0131313231323203"></a>

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

<a id="canonical-3223212023330013-1002033001211322-2332320220320212-2112230220302133-1000120201331302-2022120301001222-3202233111113322-3033012203130133"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route`

- [headers](data-sources--workload--reference--group-026.md#canonical-0002223012100201-3003102011001020-2303313023222023-1110333221302102-3033322200312101-0211132203101131-2330331033223220-1310001331213131): complete subsection reference.

<a id="canonical-1211212211110331-1111112021002220-1311132322223120-3210013222021312-2023002031003001-2122212200231103-1210003232200300-2323033331331331"></a>

<a id="canonical-2300110303310201-0210023012012011-2011220310102012-1302010222032022-0320202113012221-3133311303202322-3131230312131200-2302130313131322"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.http_method` property

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

- [incoming_port](data-sources--workload--reference--group-026.md#canonical-2222313113132022-2121332331100221-2330313020303133-0101031100213001-2203320102032033-0310022313320301-0112311221001112-1121333320212312): complete subsection reference.

- [path](data-sources--workload--reference--group-026.md#canonical-0131331113031331-1200012222302323-2123220002113003-0313211302132300-2122100122002021-2321320023320202-0021013312121001-0123011120031323): complete subsection reference.

- [route_redirect](data-sources--workload--reference--group-026.md#canonical-1112113202100020-1220200322031332-3233033312002003-2010233020320223-0203132133300221-3300312131013013-1101213022023011-3120130300003232): complete subsection reference.

<a id="canonical-0002223012100201-3003102011001020-2303313023222023-1110333221302102-3033322200312101-0211132203101131-2330331033223220-1310001331213131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-026.md#canonical-1200110130102320-0121231112030201-3130002232103131-0312211110213031-1221230122133313-3012231113202013-3232233232221320-2103301123023102)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-1210111132333220-3122020300130002-1320331313230202-3333021033230212-1020210323312230-0011110223110011-1012322002221322-1020331203030013"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1120220320201122-2331033223302031-3110233211322003-2121301030002230-1201222120212111-0233111033221010-2021202132122032-2131023200003102"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers`

<a id="canonical-3311032311101122-0301302033321103-2231212202131122-1103033200002200-0103123301110200-3102213013030330-1320010310311310-0131130022211303"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers.exact` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2012012303330323-2110032333112031-0131303130123222-0130303200131101-2101012013121101-3023102012213330-2002003012110310-1020300133223003"></a>

<a id="canonical-3132033223111232-3011332300300013-0033131121213233-2012312310203203-3330202120303132-1111013323122320-2301310313013133-3010101210101220"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers.invert_match` property

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

<a id="canonical-2212331102313020-0000230323323112-1230130220021323-2032223032222002-0301133023210013-0333120031001000-0011032122030201-2000133220103310"></a>

<a id="canonical-3222120130033120-3320313102313002-2111013121100331-3132213301313322-0230103113100032-3220300132220231-0313220230201031-3112321320333131"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0003213123132232-0112331231202213-2023121103300303-0230302321012031-2210203121312132-0030122032002122-3321021013313022-3302000001230233"></a>

<a id="canonical-1103113321012002-1333130003232213-1230111012300331-2120303000100132-0010113333231301-2223022132033310-1223130311113221-3200323320310033"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers.presence` property

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

<a id="canonical-2003021001023031-2013330022312320-1230112123102111-0020202022010301-2133333232013110-0301320301101312-2012203111133232-3321313011232221"></a>

<a id="canonical-3302331321130333-0123300220132300-0131333303030333-2010310030302230-0233021002210322-2202330031310231-1121101322102000-2012020303033113"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers.regex` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2222313113132022-2121332331100221-2330313020303133-0101031100213001-2203320102032033-0310022313320301-0112311221001112-1121333320212312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-026.md#canonical-1200110130102320-0121231112030201-3130002232103131-0312211110213031-1221230122133313-3012231113202013-3232233232221320-2103301123023102)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port

<a id="canonical-1320202023132113-0220222100100130-2322331131323101-3321002320123031-0120113222201300-3313112101223120-0221030220031001-1032203020010000"></a>

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

<a id="canonical-0132101021000100-3112112321011003-0200023220120000-1330111120021001-1110100233022301-2133131302201221-2322310121021003-1322023130020333"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port`

- [no_port_match](data-sources--workload--reference--group-026.md#canonical-2003120010230221-0200210003031221-0313123221010133-0301201212331102-1000020213322132-2103032123103320-0030231211313331-1023322333003212): complete subsection reference.

<a id="canonical-1101233030230320-1230032332232131-1100111012032213-0121210333333012-0303101022220320-2202212223020231-2220120112031223-1033220232003230"></a>

<a id="canonical-0133133203313111-2311102300022111-1313113332213322-0231133110100311-2000221300201330-3123232320111133-1032013111100300-1310223113323312"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.port` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0111110320023133-1132123211010013-0003120032000023-2213220332130203-3010312121133223-1023102231231211-0100131022121033-2032021100122023"></a>

<a id="canonical-0113203310302003-2210022312311013-0300222031033103-0331302333110003-1202110121021132-3201302023210323-1130110130011103-2213022123311203"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.port_ranges` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2003120010230221-0200210003031221-0313123221010133-0301201212331102-1000020213322132-2103032123103320-0030231211313331-1023322333003212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-026.md#canonical-1200110130102320-0121231112030201-3130002232103131-0312211110213031-1221230122133313-3012231113202013-3232233232221320-2103301123023102)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-026.md#canonical-2222313113132022-2121332331100221-2330313020303133-0101031100213001-2203320102032033-0310022313320301-0112311221001112-1121333320212312)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-0003012303103312-1003121231010032-0100113233223101-3111002202230011-0233032320220120-0231032112031032-0223313021300031-1100330213101102"></a>

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

<a id="canonical-0131331113031331-1200012222302323-2123220002113003-0313211302132300-2122100122002021-2321320023320202-0021013312121001-0123011120031323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-026.md#canonical-1200110130102320-0121231112030201-3130002232103131-0312211110213031-1221230122133313-3012231113202013-3232233232221320-2103301123023102)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path

<a id="canonical-0211113020311330-1220130121301210-3122001023311330-3130233000332101-1003310002213101-2120222013020311-0330301302300001-0200100202333001"></a>

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

<a id="canonical-1221012003313021-0301323211000020-1322122013012302-3331111022133302-2213131002010031-0233121203310032-3333201220000022-2231100233320233"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path`

<a id="canonical-2220031010103300-2203320301000113-3203221020022302-0001110320221310-3310200210031320-2213031000303133-1111200331101310-2212001133002133"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path.path` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1021011303312030-1203102122233230-3300222133223311-2133221231313030-0313210302222133-0322223100102103-0120023020103321-2031101031301232"></a>

<a id="canonical-1320100010111110-0133032012301012-3233201221210230-2302201101030323-1303000301230233-0302213120320021-3231030212002131-2231022330302320"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path.prefix` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2333102131013310-3233131131031002-3212210002133320-3102130330322311-2311221301300211-3322031221212030-2333301122101202-2230103211222230"></a>

<a id="canonical-2010022210011332-3033000033122133-2232220033031313-0120030232112110-2113100030210001-2200030322233112-1131103130303233-0013303210303332"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path.regex` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1112113202100020-1220200322031332-3233033312002003-2010233020320223-0203132133300221-3300312131013013-1101213022023011-3120130300003232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-026.md#canonical-1200110130102320-0121231112030201-3130002232103131-0312211110213031-1221230122133313-3012231113202013-3232233232221320-2103301123023102)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect

<a id="canonical-3330333203120112-3023221131222202-0201222020132223-3023003130123302-2221300020023010-0333130322022110-2030133031023201-0001222110112112"></a>

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

<a id="canonical-1332022301320001-0331120022213023-3120213220201030-3110222333333100-3212121332131322-2201323323021333-1031311030300330-1213110130220031"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect`

<a id="canonical-1223012013323032-2030310020320311-2223313331100031-1111010001030122-2320022103113133-2220101110130201-2030032103010110-3332130102213030"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.host_redirect` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3300013102312213-3121021110331210-0202321222100332-1202132120202200-2222231003133203-1002300131103330-1121032010330000-1133313013221020"></a>

<a id="canonical-3012023121320032-0203321223113022-3313301313312033-3232112211312102-0220120213113323-1023030031230302-1110311022000130-3310021300221303"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.path_redirect` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2212010312311021-1312001200103002-1233322333033033-2213023330222101-0312110131232101-2022011123223101-0311122221111322-3321233311030213"></a>

<a id="canonical-0233313032103013-0133332233332203-3100230200332121-1232110132032033-0311133022323211-3120100203310212-0013033331220311-2223210220103213"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.prefix_rewrite` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3030001303122221-2010220021130120-0001221333201301-0220223031020030-2110223121033220-1310023113330032-0230222312111331-0201013020210223"></a>

<a id="canonical-0330032121203020-1201310303001323-2131033310112020-0232330022102122-3013220303300311-3210232013301301-3231010010212111-1332112022213313"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.proto_redirect` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [remove_all_params](data-sources--workload--reference--group-026.md#canonical-1202312233122001-0002211312301031-0032020230222233-0012030033212321-2322332011301031-0030211122333113-1330213021032321-2311033013122201): complete subsection reference.

<a id="canonical-0133020130010231-1013032213133310-1323121330301213-1032230110123130-2323221230313211-3131100031001113-1320110301303103-2321033323000131"></a>

<a id="canonical-2023330320301132-0100232202302020-1201010010312131-1222122011131131-0220022333013001-1321200113201122-0211300122311311-3011332200120021"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.replace_params` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1311010003012002-3211313321010102-3013320223213111-2112211320001130-1213312300320112-2332123100223132-0313300203030020-0012313010012001"></a>

<a id="canonical-0300030033222300-3333200100230132-3100312333213000-2330033123031201-2020331103130012-1000102123203310-2021101020201130-0212211313322012"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.response_code` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [retain_all_params](data-sources--workload--reference--group-026.md#canonical-1203203331321232-3322230313031311-3001013233101012-2212322131032101-3201120200232102-3302010011010010-3312103331122123-0133222020320130): complete subsection reference.

<a id="canonical-1202312233122001-0002211312301031-0032020230222233-0012030033212321-2322332011301031-0030211122333113-1330213021032321-2311033013122201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-026.md#canonical-1200110130102320-0121231112030201-3130002232103131-0312211110213031-1221230122133313-3012231113202013-3232233232221320-2103301123023102)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-026.md#canonical-1112113202100020-1220200322031332-3233033312002003-2010233020320223-0203132133300221-3300312131013013-1101213022023011-3120130300003232)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-1231311020012020-0033201030121000-1230312231003332-2211122203033321-3113003301022221-3222133202222102-1030012132222102-0223301021023010"></a>

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

<a id="canonical-1203203331321232-3322230313031311-3001013233101012-2212322131032101-3201120200232102-3302010011010010-3312103331122123-0133222020320130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-026.md#canonical-1200110130102320-0121231112030201-3130002232103131-0312211110213031-1221230122133313-3012231113202013-3232233232221320-2103301123023102)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-026.md#canonical-1112113202100020-1220200322031332-3233033312002003-2010233020320223-0203132133300221-3300312131013013-1101213022023011-3120130300003232)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-2321212123011121-3103212331201311-2022020133102031-0321113200331112-2230103123020320-3332123230012313-2031101132131222-3221332330200022"></a>

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

<a id="canonical-2023011001130113-1300333200123000-3030302321122132-3232120332212030-2310223313003023-3213131323211332-1213202213220213-1203103011331322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-2010213332021230-2010300100131113-1311020313120030-2310132312100203-1313130111221231-0331221123121210-0230302231032301-2023202311020120"></a>

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

<a id="canonical-1123333211012030-1133300130031033-2121111231312032-3102101133231010-1030201110232030-0123321113222211-1011201231123323-1103211210232212"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route`

- [auto_host_rewrite](data-sources--workload--reference--group-026.md#canonical-2010100110022233-2132121210133021-1302312032010111-3332031331212133-3233023321012131-2030300222320021-0122310302320303-0213222321130023): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-026.md#canonical-1031001333103223-3130232111130230-0132202102302022-0323001131223123-3013011332101010-1033212133221112-0132330230133013-3021322112331031): complete subsection reference.

<a id="canonical-3310002123021321-0230000101010320-0002010303030330-2223222301010002-2221121312313310-1311102132033032-3202123311122023-3312113302001032"></a>

<a id="canonical-3223122021032231-2102331311220031-0111311333023001-0230010200220011-1320013122111123-2003323000320102-1003221103102020-1301232123211201"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.host_rewrite` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2133031000001110-2031200001013220-0331003322222010-0321002311323330-3111213002302310-3131303200230313-2310013221121113-1211003023010130"></a>

<a id="canonical-2222311231233201-1213333311233321-2302002002322221-1321221322022133-2330021110100103-3332222231011133-0213213022233100-3003301132102130"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.http_method` property

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

- [path](data-sources--workload--reference--group-026.md#canonical-1000031102331323-3033331221100000-1222333311113312-1100121020313211-3322331120012133-2212322133310023-1300031011312113-1312302123333113): complete subsection reference.

<a id="canonical-2010100110022233-2132121210133021-1302312032010111-3332031331212133-3233023321012131-2030300222320021-0122310302320303-0213222321130023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-026.md#canonical-2023011001130113-1300333200123000-3030302321122132-3232120332212030-2310223313003023-3213131323211332-1213202213220213-1203103011331322)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-1032233233131201-1132122323302100-0022020222121310-0232230200300212-3310232032321121-0233300333312100-3112201012202320-1013120131302330"></a>

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

<a id="canonical-1031001333103223-3130232111130230-0132202102302022-0323001131223123-3013011332101010-1033212133221112-0132330230133013-3021322112331031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-026.md#canonical-2023011001130113-1300333200123000-3030302321122132-3232120332212030-2310223313003023-3213131323211332-1213202213220213-1203103011331322)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-1002210222310131-3232023113013130-3302121203132202-0020032032113310-0121333211021202-0301200032221132-3023003230022312-2211221330233201"></a>

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

<a id="canonical-1000031102331323-3033331221100000-1222333311113312-1100121020313211-3322331120012133-2212322133310023-1300031011312113-1312302123333113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-026.md#canonical-2023011001130113-1300333200123000-3030302321122132-3232120332212030-2310223313003023-3213131323211332-1213202213220213-1203103011331322)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-2022333200321012-3130230300213000-1213132331101321-3031213013330300-3301032020012231-3201133130031222-2032222102322220-1213311111113300"></a>

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

<a id="canonical-1200211112302201-3302000122213321-1030031210012323-1200020112000000-3202212020023000-2231133131010330-0200132002020323-3023313130220122"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path`

<a id="canonical-2031333230233130-2010012020033122-1000213031113132-1123221112112032-2311100230300020-1013101230100200-3132111311211112-3020113222111221"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path.path` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2302203201301101-0100233021221012-3000112320032232-2300331132123312-3112220213103320-2111123223320102-2133111101120331-2230121011213103"></a>

<a id="canonical-2320132111112302-1313233000003012-0310121320103203-1203002012211322-2100223311212323-1231033112203002-2321203201101122-1320020030212303"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path.prefix` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1310133221131223-1221311313132221-2000303200320112-3221002212002033-2213103111301211-2113300212302023-3202210101213002-2112331213022110"></a>

<a id="canonical-0332303230222323-2230023231120132-3101300033103311-2302020312321201-0010310001220301-0100203123333302-2131230110133111-3033101122220200"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path.regex` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0131322030120213-1100100231233311-1130301232021101-2213213330300032-0021030111212103-0111121122012212-3320321121032303-0300331333012133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- stateful_service.advertise_options.advertise_on_public.port.port

<a id="canonical-3131100120320210-2210031112200201-3233213113301013-0032010322122031-1133023100112330-1023331123132222-1101131002232121-1022310121132011"></a>

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

<a id="canonical-1223031220231203-2022201232001132-2010130131211300-0013322002030202-0120311103301313-3210001310013202-0211122132210031-3022113130110202"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.port`

- [info](data-sources--workload--reference--group-026.md#canonical-0211003120303131-3103210022333130-0123010000333331-0300332001313323-0223221100300100-3230230003001110-0232100122111312-0222330233121323): complete subsection reference.

<a id="canonical-0211003120303131-3103210022333130-0123010000333331-0300332001313323-0223221100300100-3230230003001110-0232100122111312-0222330233121323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.port.info` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.port](data-sources--workload--reference--group-026.md#canonical-0131322030120213-1100100231233311-1130301232021101-2213213330300032-0021030111212103-0111121122012212-3320321121032303-0300331333012133)
- stateful_service.advertise_options.advertise_on_public.port.port.info

<a id="canonical-1213101321231030-3311112301223212-3012111120002130-1221323033101013-0213010223030302-0310000132310132-1323303203100222-1013311303113130"></a>

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

<a id="canonical-3223232030010300-3300111212123032-3132122322212033-0323030032323330-2111033110022021-3011013313102002-0003001112202022-3133103111121112"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.port.info`

<a id="canonical-2310202233230000-2011000101300330-1311030323113113-1201201313302103-0110120110312202-2322120223310103-3212303233222131-1122300103200200"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.port.info.port` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1303012132021302-1202033011233202-0002023330301322-1230122001112130-2002211001222030-3112001110332233-2320132333212203-1222020130303232"></a>

<a id="canonical-2310030203331023-2303033020110200-0133110223000032-3122113011320222-2003030233012332-0300321020202021-3023222120232212-2223331030101221"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.port.info.protocol` property

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

- [same_as_port](data-sources--workload--reference--group-026.md#canonical-2013110230232312-1322113131120223-0122133103330033-1222313111101233-2102323012033312-0320303212010300-1223333010130122-3311213012012233): complete subsection reference.

<a id="canonical-1022102232112032-2322310112323233-1021111203211001-2320230122320210-1220210103231130-2201213201211233-2321030023300313-2211031130231100"></a>

<a id="canonical-1030321211323211-2233231200111120-1313011133332321-2100023033032231-2312021101012202-2202313121311102-3302233010100230-0103300331020330"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.port.info.target_port` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2013110230232312-1322113131120223-0122133103330033-1222313111101233-2102323012033312-0320303212010300-1223333010130122-3311213012012233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.port.info.same_as_port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.port](data-sources--workload--reference--group-026.md#canonical-0131322030120213-1100100231233311-1130301232021101-2213213330300032-0021030111212103-0111121122012212-3320321121032303-0300331333012133)
- [stateful_service.advertise_options.advertise_on_public.port.port.info](data-sources--workload--reference--group-026.md#canonical-0211003120303131-3103210022333130-0123010000333331-0300332001313323-0223221100300100-3230230003001110-0232100122111312-0222330233121323)
- stateful_service.advertise_options.advertise_on_public.port.port.info.same_as_port

<a id="canonical-1211303110110010-3331303001301201-0020331103120112-2310010111321131-3302232022033002-3122003002031122-0230122011013022-1021223313012202"></a>

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

<a id="canonical-1033132331311233-3020310332033313-0130113122310122-2013000233110030-2111310033030212-0200233001232302-0310022301231111-1011312300303021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.tcp_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- stateful_service.advertise_options.advertise_on_public.port.tcp_loadbalancer

<a id="canonical-2310112333221301-2303223331230203-3101013110003223-2312113032231330-2300200200012030-0003220101320320-1330303321220311-3023232122023232"></a>

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

<a id="canonical-1223321022032100-1200221221210322-1030310101113202-3022130123130221-2031331103332032-0201013233030311-3311203101002211-1213203301210020"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.tcp_loadbalancer`

<a id="canonical-0332112131300022-2123020223123122-2221201311223232-1302103323213031-1320000010020300-2332230223122223-0332111022002121-2012133202200111"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.tcp_loadbalancer.domains` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1220311210103300-2010003302203110-3232002122033232-1033132030112102-1313213213230120-2331300200233011-3211103222130033-3102120013212023"></a>

<a id="canonical-3232203323113131-0200233312033120-3203321013312213-2100312302122003-1023120302331232-1103320220302020-1233233300330112-2121212230303300"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.tcp_loadbalancer.with_sni` property

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

<a id="canonical-2130301222010233-1223132131303222-0322121021211321-2133011133112020-3312223303312212-3213023022123103-1032021220021202-2003132313121010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.do_not_advertise` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- stateful_service.advertise_options.do_not_advertise

<a id="canonical-2112010330313011-0122310020030213-0213000130033003-2232210001223321-3203310011221322-2003121023231131-0032310003111222-3132203221333102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for do not advertise.

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

<a id="canonical-2032211221320222-3233032312121000-0321111221221122-3331201113022312-0020000223303230-0020033200320022-2222023002102320-2200212222310100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.configuration` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- stateful_service.configuration

<a id="canonical-2032010130230211-2312203313331131-1002010211032113-0120313110021313-0132130310131331-0013312311320001-2301223310002020-2223230203133102"></a>

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

<a id="canonical-1130213320103330-3133300200221100-3013023120101231-2130222003003033-1110331320233200-2132123233203020-1021321322102112-3133212020313231"></a>

### Direct properties for `stateful_service.configuration`

- [parameters](data-sources--workload--reference--group-026.md#canonical-3122320321220010-2102312233010312-0232123013221122-2001223312112201-2022020012221002-0030032122333331-3230312102013102-2113113003030230): complete subsection reference.

<a id="canonical-3122320321220010-2102312233010312-0232123013221122-2001223312112201-2022020012221002-0030032122333331-3230312102013102-2113113003030230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.configuration.parameters` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.configuration](data-sources--workload--reference--group-026.md#canonical-2032211221320222-3233032312121000-0321111221221122-3331201113022312-0020000223303230-0020033200320022-2222023002102320-2200212222310100)
- stateful_service.configuration.parameters

<a id="canonical-2221223221023320-0121001022302203-0112103201111101-2111020231122122-2320112303112102-2331332031211220-2003303112002213-3002123111122132"></a>

Type: `"list"`. Computed.

Parameters. Parameters for the workload.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1302001231221322-2300112022102223-1102233131323301-3303210003213332-1102110200322132-0303113122321112-3131101121332031-3132123200021021"></a>

### Direct properties for `stateful_service.configuration.parameters`

- [env_var](data-sources--workload--reference--group-026.md#canonical-0231102013011210-1020102000313033-1100331131212012-2021110220223323-2013031320222203-3003031332222020-1233030121123123-1003220313112232): complete subsection reference.

- [file](data-sources--workload--reference--group-026.md#canonical-1100303330211012-2000203213110113-0211120230032130-3232032133233010-1303222303021102-3312020133132200-2030323321211133-2112032122000213): complete subsection reference.

<a id="canonical-0231102013011210-1020102000313033-1100331131212012-2021110220223323-2013031320222203-3003031332222020-1233030121123123-1003220313112232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.configuration.parameters.env_var` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.configuration](data-sources--workload--reference--group-026.md#canonical-2032211221320222-3233032312121000-0321111221221122-3331201113022312-0020000223303230-0020033200320022-2222023002102320-2200212222310100)
- [stateful_service.configuration.parameters](data-sources--workload--reference--group-026.md#canonical-3122320321220010-2102312233010312-0232123013221122-2001223312112201-2022020012221002-0030032122333331-3230312102013102-2113113003030230)
- stateful_service.configuration.parameters.env_var

<a id="canonical-0011002310111111-3030011210030130-0223002301103232-3100033012011123-3113222230321011-1131313321011332-2203122200210220-1021300230010120"></a>

Type: `"single"`. Computed.

Environment Variable. Environment Variable.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0011232010211010-3010312130233022-3213010130001113-1130323111333030-0011331323330011-0330001201320300-2002222102310221-3130022011020132"></a>

### Direct properties for `stateful_service.configuration.parameters.env_var`

<a id="canonical-0121033211321303-2103313001232113-0201011221033311-0221322002101033-3112021012132023-0313222022210003-3112210300012013-1110323002303233"></a>

#### `stateful_service.configuration.parameters.env_var.name` property

Type: `"string"`. Computed.

Name. Name of Environment Variable.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2312122102030121-2010101213121132-0032003220131131-3113200130222231-2122101033303021-2131103013021222-2333202001213303-3132300131012001"></a>

<a id="canonical-2100021000301001-1220001313110311-0331003212223331-3301300322000223-2202210130321102-0130333222021322-0113023121101103-0100000001302323"></a>

#### `stateful_service.configuration.parameters.env_var.value` property

Type: `"string"`. Computed.

Value. Value of Environment Variable.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1100303330211012-2000203213110113-0211120230032130-3232032133233010-1303222303021102-3312020133132200-2030323321211133-2112032122000213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.configuration.parameters.file` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.configuration](data-sources--workload--reference--group-026.md#canonical-2032211221320222-3233032312121000-0321111221221122-3331201113022312-0020000223303230-0020033200320022-2222023002102320-2200212222310100)
- [stateful_service.configuration.parameters](data-sources--workload--reference--group-026.md#canonical-3122320321220010-2102312233010312-0232123013221122-2001223312112201-2022020012221002-0030032122333331-3230312102013102-2113113003030230)
- stateful_service.configuration.parameters.file

<a id="canonical-0133221021302000-3101113220120231-2220223211303022-2333032033312303-3031121123303332-1322333002023312-2213132200221231-3211220013100013"></a>

Type: `"single"`. Computed.

Configuration File. Configuration File for the workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1230121331010203-1133301130212220-1000333320323122-1311101221220303-1220312202333301-3230303013110102-1032330033222230-1002022210003311"></a>

### Direct properties for `stateful_service.configuration.parameters.file`

<a id="canonical-2121122012212221-0200033103301323-2312102032122212-2321331333033303-2231123132122020-2210311233301313-1320101003003000-3110020301021232"></a>

#### `stateful_service.configuration.parameters.file.data` property

Type: `"string"`. Computed.

Data. File data

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [mount](data-sources--workload--reference--group-026.md#canonical-2202112023100030-3222012200111102-1223330020120000-0023022220301112-0123212122110013-0130120233332122-2000133223200130-2020203211020302): complete subsection reference.

<a id="canonical-2311312211003221-2320210210221213-1003233323031032-2011021111133233-3230232211301302-2300001322311223-3322011110223300-3311021111113011"></a>

<a id="canonical-0322123221302220-3332020210212210-3210301011222322-3102011202221121-2031212021302312-0200010310312002-2232220101213013-3223231012220011"></a>

#### `stateful_service.configuration.parameters.file.name` property

Type: `"string"`. Computed.

Name. Name of the file.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3112011311031222-2313113033113012-3322220303211232-1320311201123322-2122302112103033-3233133331020313-2221033032312030-3013131022003332"></a>

<a id="canonical-0333230322000001-0222020233233302-0322031003131022-3330320023310100-3133302000200020-1323301302012113-2323211113131000-1223232231022323"></a>

#### `stateful_service.configuration.parameters.file.volume_name` property

Type: `"string"`. Computed.

Volume Name. Name of the Volume.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2202112023100030-3222012200111102-1223330020120000-0023022220301112-0123212122110013-0130120233332122-2000133223200130-2020203211020302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.configuration.parameters.file.mount` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.configuration](data-sources--workload--reference--group-026.md#canonical-2032211221320222-3233032312121000-0321111221221122-3331201113022312-0020000223303230-0020033200320022-2222023002102320-2200212222310100)
- [stateful_service.configuration.parameters](data-sources--workload--reference--group-026.md#canonical-3122320321220010-2102312233010312-0232123013221122-2001223312112201-2022020012221002-0030032122333331-3230312102013102-2113113003030230)
- [stateful_service.configuration.parameters.file](data-sources--workload--reference--group-026.md#canonical-1100303330211012-2000203213110113-0211120230032130-3232032133233010-1303222303021102-3312020133132200-2030323321211133-2112032122000213)
- stateful_service.configuration.parameters.file.mount

<a id="canonical-0112223312213103-3320130013130031-3223123211013130-0130122023112031-0220011111021312-3221331331231020-0010202012011123-1022033201031123"></a>

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

<a id="canonical-1033323310100103-3213201003331033-3202313220311200-3021100001320013-0013011220321212-3230201210302231-3213001222333222-1311003221300103"></a>

### Direct properties for `stateful_service.configuration.parameters.file.mount`

<a id="canonical-3012232332003013-2333303131123001-2202123113320311-0022320120001313-2220220323301033-0200231131002201-3012330221031231-1223311010033123"></a>

#### `stateful_service.configuration.parameters.file.mount.mode` property

Type: `"string"`. Computed.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

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

<a id="canonical-0131122220210211-2032201201311032-2102312031230032-2210033200311101-2022010032300100-1132122311010113-3310111321023102-3103303312031320"></a>

<a id="canonical-2233303321323232-2101030213011333-3103310300002113-1102030011011211-3130001322110311-0301102012333331-2213001110233130-0110032321202220"></a>

#### `stateful_service.configuration.parameters.file.mount.mount_path` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2022320322100132-2022033011300313-3322233202103123-0011213102301301-0323213110032130-0022020302311333-2000213203002310-2020321120331302"></a>

<a id="canonical-0102213302120123-2022301020331101-2200310301001031-0130121030033333-2133013333001201-1031302213320000-2213312322202030-0313303300110213"></a>

#### `stateful_service.configuration.parameters.file.mount.sub_path` property

Type: `"string"`. Computed.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Additional upstream details:

Defaults to "" (volume's root).

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- stateful_service.containers

<a id="canonical-2231200033230033-1103320220110033-1200233000202013-1303312212212310-1112300321301320-0310220313000231-3230332102331312-1111323001002002"></a>

Type: `"list"`. Computed.

Containers. Containers to use for service.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1231233332000300-1233310110120330-1033201022133222-2321323030312333-0221103131232313-3223302011313232-1132230313200100-1330012332203222"></a>

### Direct properties for `stateful_service.containers`

<a id="canonical-1212332112232333-2213231212302013-2023131230302000-3112102013101311-3311121333300131-0120103011121312-1123131031302113-1232330021020203"></a>

#### `stateful_service.containers.args` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1223011320100121-1211012033023013-0233113301221012-3233133020030110-2202013003020311-0313213310223031-2220230200300211-2232210132202200"></a>

<a id="canonical-1230130113220012-1110133023123220-1320011201010103-3113232101302331-1320121000333231-2321230232131010-1102320103112232-1013300202200200"></a>

#### `stateful_service.containers.command` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [custom_flavor](data-sources--workload--reference--group-026.md#canonical-0111033232200132-3122032001133001-0120333133122101-3223322213333131-2201032200331211-2033202303323233-2330003230013031-3113232020100003): complete subsection reference.

- [default_flavor](data-sources--workload--reference--group-026.md#canonical-0131320010322031-0233112112321032-3231113121012203-3033232202221033-3201112132321233-1302020210022201-1000303301120220-1012231333220113): complete subsection reference.

<a id="canonical-2200120223013121-1220121212131033-1103112002330030-1111112200000010-0113030332002310-1203232103102100-2033120213020303-1201021310000313"></a>

<a id="canonical-0223310013203321-2002113321310120-3002212112310320-1130102210013322-3132011321221031-2331212311132332-2123231100311222-0302103223303100"></a>

#### `stateful_service.containers.flavor` property

Type: `"string"`. Computed.

\[Enum:
CONTAINER\_FLAVOR\_TYPE\_TINY|CONTAINER\_FLAVOR\_TYPE\_MEDIUM|CONTAINER\_FLAVOR\_TYPE\_LARGE\]
Container Flavor type - CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny Tiny containers have limit of 0.1 vCPU
and 256 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium Medium containers have limit
of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_LARGE: Large Large containers
have.. Possible values are \`CONTAINER\_FLAVOR\_TYPE\_TINY\`, \`CONTAINER\_FLAVOR\_TYPE\_MEDIUM\`,
\`CONTAINER\_FLAVOR\_TYPE\_LARGE\`. Defaults to \`CONTAINER\_FLAVOR\_TYPE\_TINY\`.

Additional upstream details:

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

- [image](data-sources--workload--reference--group-026.md#canonical-0021330333100321-2210032021313101-0103302100233222-1221130213013113-2331021231110011-0122333011022231-0333012303023310-3330131303211312): complete subsection reference.

<a id="canonical-3130203201033303-3212313110232121-0123233133030010-1303003332233300-0031122312300301-3231120202311020-1231103003101230-1032223113103123"></a>

<a id="canonical-0310001303330133-3203210312233320-0313111020001202-0310000032010002-2020311313301132-2213022010103330-3033201103102300-0012232100222110"></a>

#### `stateful_service.containers.init_container` property

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

- [liveness_check](data-sources--workload--reference--group-027.md#canonical-1000133100210330-0101332000310233-1022300301200130-1202221333012211-3002300330030301-2323010320303223-3322100021113321-2121321202120322): complete subsection reference.

<a id="canonical-3312103121313330-0010030033131020-0032001131313311-2221333121033013-0032111310032123-1030333212123011-3201210021003321-2320213232122232"></a>

<a id="canonical-1131212201322333-3001031302133220-2201120332111133-3123130312313330-1110111110321212-3133012121131000-0021033020031301-1210122102132103"></a>

#### `stateful_service.containers.name` property

Type: `"string"`. Computed.

Name. Name of the container.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [readiness_check](data-sources--workload--reference--group-027.md#canonical-1203230022130102-2123013130220100-0133020010102133-0210200010302332-2332232103111100-2033223132011311-0213121011301300-3330312202013233): complete subsection reference.

<a id="canonical-0111033232200132-3122032001133001-0120333133122101-3223322213333131-2201032200331211-2033202303323233-2330003230013031-3113232020100003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.custom_flavor` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- stateful_service.containers.custom_flavor

<a id="canonical-1212302001133323-2222112033123200-0102032020200201-3230313310110033-0032013303233333-1123201332323123-2121332103211013-2023321003323333"></a>

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

<a id="canonical-0110130322313113-3012310300233233-1110102313132221-0312123030022013-2303223020200202-3210110121222332-2133221122303310-1312001123133001"></a>

### Direct properties for `stateful_service.containers.custom_flavor`

<a id="canonical-1333302212022320-2133211113211203-0210112201202332-0030310103310322-1123301113010112-1210320122012212-1010210320311322-2220211300131012"></a>

#### `stateful_service.containers.custom_flavor.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3030113002112032-1232313120320003-1302011200310122-0313212222311120-2233120232010312-1233212131032013-0300301111132032-1112012310223011"></a>

<a id="canonical-0123322221321132-0002201212133330-1201230022000321-3230320313330331-2301020320232333-0032122211311312-1200013311311033-1013321011023012"></a>

#### `stateful_service.containers.custom_flavor.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1103313333331131-0331000310000333-2021100222233312-2221002231030021-3123213122102323-0022210311212323-0213032211101202-2100333112130130"></a>

<a id="canonical-3311030200000222-2133320323031113-1313023013331220-0210311011130102-1333023100033032-2030021130202012-1203033001302232-2201330300121112"></a>

#### `stateful_service.containers.custom_flavor.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0131320010322031-0233112112321032-3231113121012203-3033232202221033-3201112132321233-1302020210022201-1000303301120220-1012231333220113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.default_flavor` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- stateful_service.containers.default_flavor

<a id="canonical-2003021323113121-1101201023333010-1031202111200311-0201233202211133-3323312032220200-3303232321221022-0000122010130333-2133100030301321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default flavor.

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

<a id="canonical-0021330333100321-2210032021313101-0103302100233222-1221130213013113-2331021231110011-0122333011022231-0333012303023310-3330131303211312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.image` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- stateful_service.containers.image

<a id="canonical-2003021032010113-3302202113030113-0233331031232221-1030131132013102-3122110302213022-1002112201311301-0110030321001303-1030102211220313"></a>

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

<a id="canonical-1011031102202210-3330013132032011-2001300110312322-3233110123031220-3100213300220130-2230032102311230-2010321023012121-0212201331120303"></a>

### Direct properties for `stateful_service.containers.image`

- [container_registry](data-sources--workload--reference--group-026.md#canonical-3003323331101200-2333101030033122-1233302031331333-0212222132020201-0101002131213303-3022011030112201-0002302021102301-3012100310320110): complete subsection reference.

<a id="canonical-3310101233132332-0001313231321032-0202220310202113-3010332221110320-0323103331302333-2231202213001222-1131000032211302-3201212002303112"></a>

<a id="canonical-0003022322021112-3330321021300101-1131120023210031-3001300320233003-3200312113103112-0131123021010333-1033133003110303-0220302021231031"></a>

#### `stateful_service.containers.image.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [public](data-sources--workload--reference--group-027.md#canonical-3103202230232111-3001333120213010-1133130321013002-2223022030303023-1133301012110330-0133320232312002-0022120311113300-1323331231023030): complete subsection reference.

<a id="canonical-1012122102033323-1201112222332323-1032103112231213-2300110000322030-2230131313220201-1032033312332122-0031313003012102-1123133131122300"></a>

<a id="canonical-3000231010311101-0311010332212122-2133222032030103-1323103031303223-3022200120221102-2020233132100031-3122020303022101-2123010121032012"></a>

#### `stateful_service.containers.image.pull_policy` property

Type: `"string"`. Computed.

\[Enum:
IMAGE\_PULL\_POLICY\_DEFAULT|IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT|IMAGE\_PULL\_POLICY\_ALWAYS|IMAGE\_PULL\_POLICY\_NEVER\]
Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload - IMAGE\_PULL\_POLICY\_DEFAULT: Default Default will always pull image if :latest tag
is specified in image name. If :latest tag is not specified in image name, it will pull image only..
Possible values are \`IMAGE\_PULL\_POLICY\_DEFAULT\`, \`IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT\`,
\`IMAGE\_PULL\_POLICY\_ALWAYS\`, \`IMAGE\_PULL\_POLICY\_NEVER\`. Defaults to
\`IMAGE\_PULL\_POLICY\_DEFAULT\`.

Additional upstream details:

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

<a id="canonical-3003323331101200-2333101030033122-1233302031331333-0212222132020201-0101002131213303-3022011030112201-0002302021102301-3012100310320110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.image.container_registry` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- [stateful_service.containers.image](data-sources--workload--reference--group-026.md#canonical-0021330333100321-2210032021313101-0103302100233222-1221130213013113-2331021231110011-0122333011022231-0333012303023310-3330131303211312)
- stateful_service.containers.image.container_registry

<a id="canonical-3210122011321300-3132000322121023-2220013010012213-2123313021220220-3302103213002020-1313102200121221-1110310123011300-2123230100010220"></a>

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

<a id="canonical-1033130333112331-3302310022003232-3110321113200312-3033302002321112-3302131220101321-1231233103023332-3123101033222203-2010232022030301"></a>

### Direct properties for `stateful_service.containers.image.container_registry`

<a id="canonical-2301220030202021-3332221320032113-2211231133321033-3113213303023201-0313221130023313-2123002032330102-1001031233222202-2012231201302222"></a>

#### `stateful_service.containers.image.container_registry.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3212102201312333-0300001103131233-0313331013132010-1013002202201322-0133312201130202-1102101230211221-3333211303133323-3021322201323013"></a>
