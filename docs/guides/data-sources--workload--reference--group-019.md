---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-0102110322203030-1313110001301021-1230020003210101-1311032001112021-1323322110321013-0202120001000002-3010211330200210-0321333121213230"></a>

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.http_method` property

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

- [incoming_port](data-sources--workload--reference--group-019.md#canonical-3301122102321222-0033021131303002-2222211330300331-1200201023320112-1330202321022320-3203120120231012-2030203012033023-3303010323023121): complete subsection reference.

- [path](data-sources--workload--reference--group-019.md#canonical-0000010220100122-1021121012222131-3131321121302120-2221202111321201-2022223321332233-3322231231332121-1222331003231112-1003220003011331): complete subsection reference.

- [route_redirect](data-sources--workload--reference--group-019.md#canonical-0220322303321013-3330333202112020-2313230033232012-2233101313031232-0332310331000000-1200021113023021-2021201122031020-1220131033033213): complete subsection reference.

<a id="canonical-1333302203002320-1331201010032332-2111120231210121-1321233130101102-0201303200002231-1131113212132310-2002330002132100-1020133012202321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-018.md#canonical-1321230220033300-2331122222200003-3002001213201022-1231111322210322-2122202133002220-2020331313130011-0023111323322033-3013101310010003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-1322300033113020-1223230320031202-2030313113012223-0113222000033323-1333111133113230-3312311212303100-1232003303220322-2303100121311001"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2113232022120300-1012101300103232-0300202013320213-0020322220221333-2132130132201303-0111033330113311-2100021131310300-3130311000000130"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers`

<a id="canonical-3231120121221100-0031031231320330-0200030331231203-1232013112013020-2322303130100100-3122020333123111-0112133211333333-3001113231113003"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.exact` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2222221203232112-2130333013313022-0300123031333022-0213012312002233-3301032100122231-3200320321103133-3311331320321121-2130133022113230"></a>

<a id="canonical-2223222323310230-0321013231013121-0023010033131032-3211301020122320-2110022110223012-1122320001012101-0221013032210301-0332223323321013"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.invert_match` property

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

<a id="canonical-3030001021110232-2011321020103000-0101113021013103-0001000300223232-2031030222201033-2212013002120101-1000032220310132-1123300201100331"></a>

<a id="canonical-0011000031030103-0220103323013333-1320131012303131-2003202300232112-1330321200210021-0233221212032212-3321131103321000-1001332132223312"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1310011333113222-2312322120320223-2131000311212212-1300000212112011-0022113122202220-1012130231330201-3120330310030112-1100213200102222"></a>

<a id="canonical-0202102100122300-0030313330331202-3020000133331112-2010210021120232-2232202001022212-2200300021303200-3220031023011000-0212132130023111"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.presence` property

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

<a id="canonical-3031022310231332-0311022300022331-3313233300303131-2222011323331110-1233333333330222-2133123113222310-0113001302122220-0020301001033200"></a>

<a id="canonical-2213220323133332-0201103013211021-0101013122200020-3203313013212003-3011111030233030-2131312031230312-0011032011020020-0213113212302010"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.regex` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3301122102321222-0033021131303002-2222211330300331-1200201023320112-1330202321022320-3203120120231012-2030203012033023-3303010323023121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-018.md#canonical-1321230220033300-2331122222200003-3002001213201022-1231111322210322-2122202133002220-2020331313130011-0023111323322033-3013101310010003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port

<a id="canonical-0030110330132230-1003130022232321-3103303120303003-2120320122303113-1222312110000220-0113010231133013-0121103103230133-0023001331323303"></a>

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

<a id="canonical-2321222033331220-2133333012302113-3201123030320001-1010033233200331-2112113231222220-3303310120333130-2112121013321110-0301132002132302"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port`

- [no_port_match](data-sources--workload--reference--group-019.md#canonical-2300211332011123-2310333111312201-3032221031032121-2320231000232123-3011212221221301-2220310322220133-1120100012022331-0200032211302311): complete subsection reference.

<a id="canonical-1102223233330130-1321233103023302-1033112011203011-1222221210133301-1012012232103122-0103301201122010-0132322311122122-2010023131233230"></a>

<a id="canonical-2323111011211230-1101131011320103-3301323013223001-2013330333103220-2103103131320213-3010030000311302-2301021303022220-3221030211332331"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.port` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3122123333100030-3223131323303301-3012011200131331-1320323102013020-2113200212012301-0201220313002111-0120031010033302-0313300010312211"></a>

<a id="canonical-2003202232313213-2331103321021233-1003330312213332-1333120230032330-3133113133101300-3202213321211032-3033301302223023-1301033001111020"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.port_ranges` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2300211332011123-2310333111312201-3032221031032121-2320231000232123-3011212221221301-2220310322220133-1120100012022331-0200032211302311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-018.md#canonical-1321230220033300-2331122222200003-3002001213201022-1231111322210322-2122202133002220-2020331313130011-0023111323322033-3013101310010003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-019.md#canonical-3301122102321222-0033021131303002-2222211330300331-1200201023320112-1330202321022320-3203120120231012-2030203012033023-3303010323023121)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-0032222021112012-3211202001101120-0002233113010103-2010131323001021-0131303223110100-3100303333303023-3010333232211101-0003300212222100"></a>

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

<a id="canonical-0000010220100122-1021121012222131-3131321121302120-2221202111321201-2022223321332233-3322231231332121-1222331003231112-1003220003011331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-018.md#canonical-1321230220033300-2331122222200003-3002001213201022-1231111322210322-2122202133002220-2020331313130011-0023111323322033-3013101310010003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path

<a id="canonical-3200122211021313-3310113122001001-3133012201211031-3232313211113303-1111303310131003-1222120101221112-1321020010113201-0311113131021000"></a>

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

<a id="canonical-0010002302322121-1001102222101033-2212031302000033-0022113122212222-2232303300313230-3301112203310211-1133122031121033-2303331220222133"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path`

<a id="canonical-0301130312113000-3310130302220033-1330232120210331-1113101022311102-3200221322023313-1132033333312210-3002102013322332-0123031133202011"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path.path` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3131002101013320-2313313320033311-3103102323222111-0302233311131021-3102210333130133-3300101020230200-0120133310022222-0132230032013020"></a>

<a id="canonical-2210211031131120-0222321112031313-1203101231301132-1221330311103310-0100033021222213-2013322221033010-3032001001111331-1002022002122013"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path.prefix` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1321321321332311-1130222220301012-1330110103021120-1312111012321103-3233032322103112-1322012221021230-3020000133203303-0112022121322210"></a>

<a id="canonical-1131031202210031-0102003013000202-1002200131232101-0112113031012100-3231032230133301-3130131302131212-0112121001131310-1230103022131321"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path.regex` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0220322303321013-3330333202112020-2313230033232012-2233101313031232-0332310331000000-1200021113023021-2021201122031020-1220131033033213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-018.md#canonical-1321230220033300-2331122222200003-3002001213201022-1231111322210322-2122202133002220-2020331313130011-0023111323322033-3013101310010003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect

<a id="canonical-1320113021023010-2312001313023222-2202003012100112-1022331031320210-2210003211100203-3000030310311323-2102231302310112-1230333011300113"></a>

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

<a id="canonical-3230300231313132-1213221100230332-1313123020011332-1201232201332332-0230132310202212-2112222113011330-1123022320003030-3233301000312331"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect`

<a id="canonical-3101020312310003-1312012203333220-0122020313230331-2030103302031112-3013011310322330-1331101313020121-3333203231332330-0212302133320302"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.host_redirect` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1012102133202031-0110113021203101-3313233231032211-2221232331220130-3221202003313001-3000122311300323-3223132121033221-2010111222030100"></a>

<a id="canonical-3200323233232223-2321210031331013-0221203012011003-2220030101033321-0003133232203321-2231022300202023-1222033130321111-2212231310211023"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.path_redirect` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0233101033123101-0310010230033111-0312212023301301-2311233232310320-3003323200333331-0023312023010112-2302020200333331-1101003001101231"></a>

<a id="canonical-2320202132100201-1221103001110233-3122132003303333-0221201121201113-3300331200121301-3220320213030331-3010210333132130-2202101201031121"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.prefix_rewrite` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2132331010213210-3200030231302013-3020120200332112-0332033231021131-2112313311112223-2332201230101010-3111230133232200-2232101213000103"></a>

<a id="canonical-0131122000331210-3203131120121221-3200203220332000-2001330332123102-0001230121122332-0312121022033031-0130000023013300-3023301332210300"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.proto_redirect` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [remove_all_params](data-sources--workload--reference--group-019.md#canonical-1221232231200232-3110122201031020-1021310013223020-0120302032231120-0301202130133103-0100132213020012-2122000330103111-0331333302211313): complete subsection reference.

<a id="canonical-3300012300120223-2021030103233011-3222200012301322-2231210013131012-1001302013300033-1200002221101233-2100232110230111-2201013232131202"></a>

<a id="canonical-1022100123120121-0132233011300231-0022310012211303-1030012010230120-1333003213212033-0231221000311103-0001110120303110-1202332012330312"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.replace_params` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0003313222131301-2220212323133320-3333220113131311-2001210330130112-0031011212131331-0030300120311301-0210121200311202-3010012020232202"></a>

<a id="canonical-3221300230032010-0131103222100232-1012230033230321-3020301310321303-3012222113123333-0200123310032322-1213332322233003-2110131211213231"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.response_code` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [retain_all_params](data-sources--workload--reference--group-019.md#canonical-3102201313311130-0122313310030000-3332302310120120-0302122131331003-1230320022100023-3203322120010232-3031332121210303-0203322321120011): complete subsection reference.

<a id="canonical-1221232231200232-3110122201031020-1021310013223020-0120302032231120-0301202130133103-0100132213020012-2122000330103111-0331333302211313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-018.md#canonical-1321230220033300-2331122222200003-3002001213201022-1231111322210322-2122202133002220-2020331313130011-0023111323322033-3013101310010003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-019.md#canonical-0220322303321013-3330333202112020-2313230033232012-2233101313031232-0332310331000000-1200021113023021-2021201122031020-1220131033033213)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-3031111233213333-2310323330002323-2011330030003003-3203301311121113-0112012322222002-2311313031012103-2232032133233113-0330211201100211"></a>

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

<a id="canonical-3102201313311130-0122313310030000-3332302310120120-0302122131331003-1230320022100023-3203322120010232-3031332121210303-0203322321120011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-018.md#canonical-1321230220033300-2331122222200003-3002001213201022-1231111322210322-2122202133002220-2020331313130011-0023111323322033-3013101310010003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-019.md#canonical-0220322303321013-3330333202112020-2313230033232012-2233101313031232-0332310331000000-1200021113023021-2021201122031020-1220131033033213)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-1033322130021331-3303323011010133-2033231100112010-1301013113210231-2203203211202301-2131030213233020-1221012201032222-1230230020210102"></a>

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

<a id="canonical-3212320223022021-3213110113203220-3102231111023223-3311013332312230-0223232223121323-2120320230310210-1222202102333233-0031000030332110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-0033012013311103-2032023021231111-0001231220120122-3113233231130000-0112132011210321-2031233023222032-0012112332011220-3301223230101201"></a>

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

<a id="canonical-2030011120301300-3312023032132213-3323230222000121-3112030000132110-1110220203021311-3303110302202102-2112200333213302-3320321033103210"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route`

- [auto_host_rewrite](data-sources--workload--reference--group-019.md#canonical-0310031312033132-0013132212232210-1311103011102011-1010303113033220-1130002211312302-2012102110002100-2102201112132103-2121310232300333): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-019.md#canonical-1031210023320202-2023102310230210-1021332100223120-1021032121200121-2001332021321003-0011300330200230-0011010001001110-0102103311100202): complete subsection reference.

<a id="canonical-2030311132010133-3303222211001200-2012033331122230-0132221203233013-3231122101032023-1120200321312302-1330031221123331-1103310333321031"></a>

<a id="canonical-1331101231222132-2132123020221202-3013332010332202-3312123130122120-1231122013322303-0203111200100302-0112101312221012-1202302033203213"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.host_rewrite` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3122203230211310-3101003011333210-0003031333112221-2213303313133010-3000123303012110-2001013122101130-2033103000310121-0123031131332000"></a>

<a id="canonical-3233332222021312-3331310322331030-0103200132013323-2132202231103101-1121302032110323-3311023302302333-1231010303200021-1130132322310222"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.http_method` property

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

- [path](data-sources--workload--reference--group-019.md#canonical-1022302023022221-3322313030113221-3223323320111011-0211332022102201-0300120200301201-2002203231131131-1232001223113300-3123330301132121): complete subsection reference.

<a id="canonical-0310031312033132-0013132212232210-1311103011102011-1010303113033220-1130002211312302-2012102110002100-2102201112132103-2121310232300333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-019.md#canonical-3212320223022021-3213110113203220-3102231111023223-3311013332312230-0223232223121323-2120320230310210-1222202102333233-0031000030332110)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-0213121102302113-3300031113323201-2132333000022010-0320031032132223-2312111111012202-2302112202302303-2203231203003300-1330302210221231"></a>

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

<a id="canonical-1031210023320202-2023102310230210-1021332100223120-1021032121200121-2001332021321003-0011300330200230-0011010001001110-0102103311100202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-019.md#canonical-3212320223022021-3213110113203220-3102231111023223-3311013332312230-0223232223121323-2120320230310210-1222202102333233-0031000030332110)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-1300223101031012-1222001033211102-3131132000110223-1310130233332201-1133100101023120-0221031101211132-1323003013013123-0303202313311132"></a>

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

<a id="canonical-1022302023022221-3322313030113221-3223323320111011-0211332022102201-0300120200301201-2002203231131131-1232001223113300-3123330301132121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-018.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-019.md#canonical-3212320223022021-3213110113203220-3102231111023223-3311013332312230-0223232223121323-2120320230310210-1222202102333233-0031000030332110)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-3001131313002111-3330010030220031-3100030123200311-3333200331331123-2110230301223110-1003100122201111-2203010332202023-1232033221030212"></a>

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

<a id="canonical-0213311310312230-0302330312222130-3211032220021230-3131333330211311-2000300331211130-2212201111233101-0032312020131333-2221333133221002"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path`

<a id="canonical-0220120332002300-1210210330201133-2121303331302110-2103230320213322-2331230130121222-3313211100121130-1312323301232023-0032231200001303"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path.path` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0223211112313100-1302000221233020-2203120220322020-0302102200311332-3002201211132132-3121000000322300-0333322111001201-0022032311322331"></a>

<a id="canonical-3101223213023232-2020102221032013-1232112302001201-3231030210221031-3300332232111220-0312012212233332-1010220000102313-1122202232313320"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path.prefix` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0211123032123321-3322113121331021-0023330223310300-2313132202332213-3331030031121033-3322003321323303-1033301031000301-2113313320110020"></a>

<a id="canonical-3302300213000233-0323123200211130-1212233033311002-1113003320313303-3130002223033200-2022322130021003-1322301030221301-0130002133010030"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path.regex` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3020203222030212-3033113130221110-2320320300332011-1001030302320102-2022321230310000-0112020113321222-3002102322302000-3221222122031101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- stateful_service.advertise_options.advertise_custom.ports.port

<a id="canonical-2210131221210323-2102332123301022-3030023210022010-2222100323132102-1201222121200112-2130003211311023-3023011112101200-3233011011211201"></a>

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

<a id="canonical-0110302003133110-2202311012201133-3030000133230111-1123212203201003-2222303321311010-0310300012321120-3313303103011001-3003111122112033"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.port`

- [info](data-sources--workload--reference--group-019.md#canonical-2110220130101010-3030301022012202-0021201311302013-2010021213133010-3003201320101102-3211032223013120-2011131020111320-2311023211121203): complete subsection reference.

<a id="canonical-2211132103001013-0321031203323201-1031211130213112-2002012010011211-0333120210232131-0023110031022003-3032322021202231-2000320221311220"></a>

<a id="canonical-1212233313031103-2031230320320130-1021221002012122-1033320121122232-2130130010002233-0102120110010220-0113310112000300-1332001233211333"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.port.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2110220130101010-3030301022012202-0021201311302013-2010021213133010-3003201320101102-3211032223013120-2011131020111320-2311023211121203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.port.info` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.port](data-sources--workload--reference--group-019.md#canonical-3020203222030212-3033113130221110-2320320300332011-1001030302320102-2022321230310000-0112020113321222-3002102322302000-3221222122031101)
- stateful_service.advertise_options.advertise_custom.ports.port.info

<a id="canonical-3033201230022300-2021012000313223-0220010313223313-1222002122013311-3120123203010213-3112301113203001-3102113020201330-2123320132132231"></a>

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

<a id="canonical-2330023100112122-3100223200031011-1333011013201311-1210323322233303-1303002203313032-1003021321031200-2131101233201223-1230301332103330"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.port.info`

<a id="canonical-3001323230200023-1112311232200123-2230021331021132-1213202003013300-2302133222301111-3112201321332000-0322101203311233-3223313111331202"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.port.info.port` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0122030321303313-3210220202021212-3123323110322001-1023033000223102-1312203312020012-0213202222021123-3033012101333000-1032320020033013"></a>

<a id="canonical-3103032120111231-0103122023131112-1001033010301223-2012200330131323-2101012222023233-0123122103112202-3313110013020010-1211233202332130"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.port.info.protocol` property

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

- [same_as_port](data-sources--workload--reference--group-019.md#canonical-0211113220011200-3210133211300211-1201013130300023-3122330312122331-0123110010131120-0010211132232013-0311310332011230-2031102332102132): complete subsection reference.

<a id="canonical-1232003030230030-3220011233330032-2121021222002333-1210002302132312-1300232223101103-1120103321320331-1230220132223003-1203102013322320"></a>

<a id="canonical-2013101130030322-2321010012033231-2301003030230132-0332313022112201-1210021330330233-2231110213213003-3100202113300103-0113333302002031"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.port.info.target_port` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0211113220011200-3210133211300211-1201013130300023-3122330312122331-0123110010131120-0010211132232013-0311310332011230-2031102332102132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.port.info.same_as_port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.port](data-sources--workload--reference--group-019.md#canonical-3020203222030212-3033113130221110-2320320300332011-1001030302320102-2022321230310000-0112020113321222-3002102322302000-3221222122031101)
- [stateful_service.advertise_options.advertise_custom.ports.port.info](data-sources--workload--reference--group-019.md#canonical-2110220130101010-3030301022012202-0021201311302013-2010021213133010-3003201320101102-3211032223013120-2011131020111320-2311023211121203)
- stateful_service.advertise_options.advertise_custom.ports.port.info.same_as_port

<a id="canonical-2130200213022223-1211022321031033-2330312301331203-0130321022232202-3321303011211332-2221333112032320-0011230002212111-0301313330320022"></a>

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

<a id="canonical-3021012100101331-3130231312320112-3123200110222222-2102132131213021-0312223301023023-0100202023203122-0211120233110023-0332201103002102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer

<a id="canonical-0131133031323001-2320102302000320-2121210222123002-2020121133301230-1201032302331211-0021232021012022-2231303230330330-3130010332322120"></a>

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

<a id="canonical-3002101230222310-2303311111123300-2022230021210232-2213020333232331-2203121211310231-3322111211111100-1330323331213132-2021303110321203"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer`

<a id="canonical-3131332321212221-1013321111231222-3002202023133221-3210312203001201-2312120012330033-3203332132311223-3121202003202002-3010030120220210"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer.domains` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1023213120333223-1113132031132312-3223303020232010-0303320011110122-1002102103312103-2300232002311331-1313022100202322-1311132102223211"></a>

<a id="canonical-1123003202123101-3210002033023130-1032100333111130-0320131002100320-0030223312103322-2220220202023312-1122202223133302-1031001313330231"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer.with_sni` property

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

<a id="canonical-2000012203330031-3013213320003321-1130203203021012-0301212320010201-2013312012132313-2201113011320310-2123010110120120-3030200031122122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_in_cluster` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- stateful_service.advertise_options.advertise_in_cluster

<a id="canonical-2021111220303001-1203012113223120-3013012212221100-1331001301133211-2331331130022012-0202231101102301-3320023232321110-1333003120103230"></a>

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

<a id="canonical-3233213033213032-0002230310002013-0202001311323313-3000221111320201-2301122021031121-2031322221110313-0332223212132011-3230233031030132"></a>

### Direct properties for `stateful_service.advertise_options.advertise_in_cluster`

- [multi_ports](data-sources--workload--reference--group-019.md#canonical-2000303120011023-0032222100323101-0130223122130220-0133032101133223-3220033300323002-1222031012303102-1233311123222313-3023312333113101): complete subsection reference.

- [port](data-sources--workload--reference--group-019.md#canonical-0132121001321323-0100130033213110-1210023033132202-1313133133131203-0323330100012311-1222313130120133-3020102231033032-1132100211233330): complete subsection reference.

<a id="canonical-2000303120011023-0032222100323101-0130223122130220-0133032101133223-3220033300323002-1222031012303102-1233311123222313-3023312333113101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_in_cluster.multi_ports` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-019.md#canonical-2000012203330031-3013213320003321-1130203203021012-0301212320010201-2013312012132313-2201113011320310-2123010110120120-3030200031122122)
- stateful_service.advertise_options.advertise_in_cluster.multi_ports

<a id="canonical-1032123113231112-1230221133330310-2033022202323203-3301320111113322-0213303211321333-0121123101312120-2002201131321000-2221313302113102"></a>

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

<a id="canonical-1332002210033230-2201000110223103-0201131003002021-0330220110003222-2201011231012131-0031130231100333-2213102322213122-1232210021300320"></a>

### Direct properties for `stateful_service.advertise_options.advertise_in_cluster.multi_ports`

- [ports](data-sources--workload--reference--group-019.md#canonical-1100003301200211-0221011010130010-3232032030331031-1000332023233021-0121021231320231-2121103331312012-1031200122010330-0231200011102302): complete subsection reference.

<a id="canonical-1100003301200211-0221011010130010-3232032030331031-1000332023233021-0121021231320231-2121103331312012-1031200122010330-0231200011102302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-019.md#canonical-2000012203330031-3013213320003321-1130203203021012-0301212320010201-2013312012132313-2201113011320310-2123010110120120-3030200031122122)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports](data-sources--workload--reference--group-019.md#canonical-2000303120011023-0032222100323101-0130223122130220-0133032101133223-3220033300323002-1222031012303102-1233311123222313-3023312333113101)
- stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports

<a id="canonical-2033011113321220-3112112333123002-3131220012033123-3232300323122211-0003321113220003-2113101232020210-2331202033231100-2232111132033222"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0211201200330230-3102013213320003-3322130232003021-0330303332033332-0321220121311000-1013113031211230-0202030323133303-1302112331313131"></a>

### Direct properties for `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports`

- [info](data-sources--workload--reference--group-019.md#canonical-2130000132031112-2232330332021231-3211010303112311-1310211220011201-2010100300010011-1301000202033103-0201220120333301-0121013130231200): complete subsection reference.

<a id="canonical-2012301111320131-2201132022110032-0120232112212313-3011020121031230-0230003120013103-0330112222210121-1033213123021020-0130020113321302"></a>

<a id="canonical-0010231113211110-0200032100321213-0232001333013222-3231333220201220-0232221003112021-2032322122301330-3023300331032313-2231201332011100"></a>

#### `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2130000132031112-2232330332021231-3211010303112311-1310211220011201-2010100300010011-1301000202033103-0201220120333301-0121013130231200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-019.md#canonical-2000012203330031-3013213320003321-1130203203021012-0301212320010201-2013312012132313-2201113011320310-2123010110120120-3030200031122122)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports](data-sources--workload--reference--group-019.md#canonical-2000303120011023-0032222100323101-0130223122130220-0133032101133223-3220033300323002-1222031012303102-1233311123222313-3023312333113101)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports](data-sources--workload--reference--group-019.md#canonical-1100003301200211-0221011010130010-3232032030331031-1000332023233021-0121021231320231-2121103331312012-1031200122010330-0231200011102302)
- stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info

<a id="canonical-3112133032101223-2230312212313202-1001322330121330-0200212001013102-3211210311132330-3310012013103200-2201013101030211-2102322223321202"></a>

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

<a id="canonical-2011032020230102-1032132023231133-1321203212220130-3301022232120213-3033120230030323-3233103331320031-1322021030103221-2021230210133002"></a>

### Direct properties for `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info`

<a id="canonical-1323210232213220-0330012021130303-1113223231110303-1011133010201112-3000333300222310-1331102132321212-2110332301121303-3210211213221332"></a>

#### `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info.port` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3321011132002210-0211310213221232-2032301010003031-3220011321301033-3300002303121203-1210230332220033-2133202303001122-0121200003321030"></a>

<a id="canonical-1310011011113312-2303210232222102-1322312103322312-2103311330231231-2210130331033233-0301010201302210-2202231323121300-1013230301102301"></a>

#### `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info.protocol` property

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

- [same_as_port](data-sources--workload--reference--group-019.md#canonical-2032310200320332-3020331003310203-2301020030101102-3332032123103032-2301333220000213-3121323223013123-3231111213210232-0303010123310220): complete subsection reference.

<a id="canonical-3312001103012221-1231322022310233-0102020122212220-0001302123032332-3032102002023202-0123233002123101-1303310023313131-0110301201120203"></a>

<a id="canonical-3323220212020112-2011301003200331-3121330322233000-1132132233330010-1001201031030221-0212320111321032-2310123021130021-3011321213200013"></a>

#### `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info.target_port` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2032310200320332-3020331003310203-2301020030101102-3332032123103032-2301333220000213-3121323223013123-3231111213210232-0303010123310220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-019.md#canonical-2000012203330031-3013213320003321-1130203203021012-0301212320010201-2013312012132313-2201113011320310-2123010110120120-3030200031122122)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports](data-sources--workload--reference--group-019.md#canonical-2000303120011023-0032222100323101-0130223122130220-0133032101133223-3220033300323002-1222031012303102-1233311123222313-3023312333113101)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports](data-sources--workload--reference--group-019.md#canonical-1100003301200211-0221011010130010-3232032030331031-1000332023233021-0121021231320231-2121103331312012-1031200122010330-0231200011102302)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info](data-sources--workload--reference--group-019.md#canonical-2130000132031112-2232330332021231-3211010303112311-1310211220011201-2010100300010011-1301000202033103-0201220120333301-0121013130231200)
- stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port

<a id="canonical-1300123302201202-2230210322000222-2123221331231330-1030301221132330-0301320133200030-3302121123132223-1232233130023323-2302031031012321"></a>

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

<a id="canonical-0132121001321323-0100130033213110-1210023033132202-1313133133131203-0323330100012311-1222313130120133-3020102231033032-1132100211233330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_in_cluster.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-019.md#canonical-2000012203330031-3013213320003321-1130203203021012-0301212320010201-2013312012132313-2201113011320310-2123010110120120-3030200031122122)
- stateful_service.advertise_options.advertise_in_cluster.port

<a id="canonical-1330023300302231-0133123120302303-1131010003233320-1230032213230113-1212120120220212-0221122220231300-1222303223131131-0311023222203320"></a>

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

<a id="canonical-2110213303322222-0013111101202021-1331303313123300-2021212323010003-1103311213011320-1103202102010320-0101121032121112-2030120022121322"></a>

### Direct properties for `stateful_service.advertise_options.advertise_in_cluster.port`

- [info](data-sources--workload--reference--group-020.md#canonical-0113223111021000-3232330101122312-2122223231322002-0111100001210313-3303112120020230-1122133031012002-2203013223322310-1011112213223001): complete subsection reference.
